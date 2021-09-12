package protoweb

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/metadata"
)

var (
	ErrIllegalHeaderWrite = errors.New("transport: the stream is done or WriteHeader was already called")
)

type transportStream struct {
	method  string
	w       http.ResponseWriter
	r       *http.Request
	isSent  bool
	trailer metadata.MD
}

func newTransportStream(name string, w http.ResponseWriter, r *http.Request) *transportStream {
	return &transportStream{
		method:  name,
		w:       w,
		r:       r,
		trailer: metadata.MD{},
	}
}

func (s *transportStream) Method() string {
	return s.method
}

func (s *transportStream) SetHeader(md metadata.MD) error {
	if s.isSent {
		return ErrIllegalHeaderWrite
	}
	writeMetadataToHeader(md, s.w.Header())
	return nil
}

func (s *transportStream) SendHeader(md metadata.MD) error {
	if s.isSent {
		return ErrIllegalHeaderWrite
	}
	writeMetadataToHeader(md, s.w.Header())
	return nil
}

func (s *transportStream) SetTrailer(md metadata.MD) error {
	s.trailer = metadata.Join(s.trailer, md)
	return nil
}

func (s *transportStream) writeTrailer() {
	writeMetadataToHeader(s.trailer, s.w.Header())
}

func writeMetadataToHeader(md metadata.MD, h http.Header) {
	for k, vv := range md {
		for i := range vv {
			h.Add(k, vv[i])
		}
	}
}
