package micro

import (
	"context"

	"chick/pkg/log"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Infof("server [%s] received request: %s", req.Service(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}

func logClientWrap(c client.Client) client.Client {
	return &logClientWrapper{c}
}

type logClientWrapper struct {
	client.Client
}

func (l *logClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Infof("client request to service: [%s] endpoint: [%s].", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}
