package protoweb

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

type ServiceDesc struct {
	ServiceName string
	HandlerType interface{}
	Methods     []MethodDesc
	Streams     []StreamDesc
	Metadata    string
}

type methodHandler func(srv interface{}, w http.ResponseWriter, r *http.Request, params httprouter.Params, interceptor grpc.UnaryServerInterceptor) (interface{}, error)

type MethodDesc struct {
	MethodName string
	Path       string
	HttpMethod string
	Handler    methodHandler
}

type StreamDesc struct {
	StreamName    string
	Path          string
	Handler       grpc.StreamHandler
	ServerStreams bool
	ClientStreams bool
}

type serviceInfo struct {
	serviceImpl interface{}
	methods     map[string]*MethodDesc
	streams     map[string]*StreamDesc
	mdata       interface{}
}
