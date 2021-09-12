package protoweb

import (
	"context"
	"net"

	"github.com/gobwas/ws/wsutil"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type serverStream struct {
	ctx  context.Context
	conn net.Conn
}

func newServerStream(ctx context.Context, conn net.Conn) *serverStream {
	return &serverStream{
		ctx:  ctx,
		conn: conn,
	}
}

func (ss *serverStream) SetHeader(metadata.MD) error {
	return nil
}

func (ss *serverStream) SendHeader(metadata.MD) error {
	return nil
}

func (ss *serverStream) SetTrailer(metadata.MD) {
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func (ss *serverStream) SendMsg(m interface{}) error {
	b, err := protojsonMarshalOptions.Marshal(m.(proto.Message))
	if err != nil {
		return err
	}
	return wsutil.WriteServerText(ss.conn, b)
}

func (ss *serverStream) RecvMsg(m interface{}) error {
	b, err := wsutil.ReadClientText(ss.conn)
	if err != nil {
		return err
	}
	return protojsonUnmarshalOptions.Unmarshal(b, m.(proto.Message))
}
