package protoweb

import (
	"context"
	"io"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"sync"

	"github.com/gobwas/ws"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	protojsonMarshalOptions = protojson.MarshalOptions{
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	protojsonUnmarshalOptions = protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
)

type ServiceRegistrar interface {
	RegisterService(desc *ServiceDesc, impl interface{})
}

type Server struct {
	mu       sync.Mutex
	router   *httprouter.Router
	upgrader *ws.HTTPUpgrader
	services map[string]*serviceInfo

	logger *zap.SugaredLogger

	unaryInterceptor  grpc.UnaryServerInterceptor
	streamInterceptor grpc.StreamServerInterceptor
}

func NewServer() *Server {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &Server{
		router:   httprouter.New(),
		upgrader: &ws.HTTPUpgrader{},
		services: map[string]*serviceInfo{},

		logger: logger.Sugar(),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) RegisterService(sd *ServiceDesc, ss interface{}) {
	if ss != nil {
		ht := reflect.TypeOf(sd.HandlerType).Elem()
		st := reflect.TypeOf(ss)
		if !st.Implements(ht) {
			s.logger.Fatalf("proto-web: Server.RegisterService found the handler of type %v that does not satisfy %v", st, ht)
		}
	}
	s.register(sd, ss)
}

func (s *Server) register(sd *ServiceDesc, ss interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.logger.Debugf("RegisterService(%q)", sd.ServiceName)

	if _, ok := s.services[sd.ServiceName]; ok {
		s.logger.Fatalf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName)
	}

	info := &serviceInfo{
		serviceImpl: ss,
		methods:     make(map[string]*MethodDesc),
		streams:     make(map[string]*StreamDesc),
		mdata:       sd.Metadata,
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		info.methods[d.MethodName] = d
		s.router.Handle(d.HttpMethod, d.Path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			s.processUnaryRequest(w, r, params, info, d)
		})
	}
	for i := range sd.Streams {
		d := &sd.Streams[i]
		info.streams[d.StreamName] = d
		s.router.GET(d.Path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			s.processStreamRequest(w, r, params, info, d)
		})
	}
	s.services[sd.ServiceName] = info
}

func (s *Server) processUnaryRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params, si *serviceInfo, md *MethodDesc) (err error) {
	ctx := r.Context()
	pr := &peer.Peer{}
	{
		host, port, _ := net.SplitHostPort(r.RemoteAddr)
		ip := net.ParseIP(host)
		p, _ := strconv.ParseInt(port, 10, 32)
		pr.Addr = &net.TCPAddr{
			IP:   ip,
			Port: int(p),
		}
	}
	ctx = peer.NewContext(ctx, pr)

	transport := newTransportStream(md.MethodName, w, r)
	r = r.WithContext(grpc.NewContextWithServerTransportStream(ctx, transport))

	var resp interface{}
	resp, err = md.Handler(si.serviceImpl, w, r, params, s.unaryInterceptor)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			// Convert appErr if it is not a grpc status error.
			err = status.Error(codes.Unknown, err.Error())
			st, _ = status.FromError(err)
		}
		httpStatus := http.StatusInternalServerError
		if he, ok := err.(interface {
			HTTPStatus() int
		}); ok {
			httpStatus = he.HTTPStatus()
		}
		w.WriteHeader(httpStatus)
		b, _ := protojsonMarshalOptions.Marshal(st.Proto())
		_, _ = w.Write(b)
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := protojsonMarshalOptions.Marshal(resp.(proto.Message))
		_, _ = w.Write(b)
	}

	transport.isSent = true
	transport.writeTrailer()
	return nil
}

func (s *Server) processStreamRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params, si *serviceInfo, sd *StreamDesc) (err error) {
	ctx := r.Context()
	pr := &peer.Peer{}
	{
		host, port, _ := net.SplitHostPort(r.RemoteAddr)
		ip := net.ParseIP(host)
		p, _ := strconv.ParseInt(port, 10, 32)
		pr.Addr = &net.TCPAddr{
			IP:   ip,
			Port: int(p),
		}
	}
	ctx = peer.NewContext(ctx, pr)

	transport := newTransportStream(sd.StreamName, w, r)
	ctx = grpc.NewContextWithServerTransportStream(ctx, transport)

	conn, _, _, err := s.upgrader.Upgrade(r, w)
	if err != nil {
		st := status.New(codes.Unknown, err.Error())
		err = st.Err()
		b, _ := protojsonMarshalOptions.Marshal(st.Proto())
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(b)
		return
	}

	ss := newServerStream(ctx, conn)

	if s.streamInterceptor == nil {
		err = sd.Handler(si.serviceImpl, ss)
	} else {
		info := &grpc.StreamServerInfo{
			FullMethod:     sd.StreamName,
			IsClientStream: sd.ClientStreams,
			IsServerStream: sd.ServerStreams,
		}
		err = s.streamInterceptor(si.serviceImpl, ss, info, sd.Handler)
	}
	return nil
}

func toRPCErr(err error) error {
	if err == nil || err == io.EOF {
		return err
	}
	if err == io.ErrUnexpectedEOF {
		return status.Error(codes.Internal, err.Error())
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	switch err.(type) {
	default:
		switch err {
		case context.DeadlineExceeded:
			return status.Error(codes.DeadlineExceeded, err.Error())
		case context.Canceled:
			return status.Error(codes.Canceled, err.Error())
		}
	}
	return status.Error(codes.Unknown, err.Error())
}
