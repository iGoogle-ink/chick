package server

import (
	"context"
	"fmt"
	"time"

	"chick/api/oauth2"
	"chick/micro-svr/oauth2/conf"
	"chick/micro-svr/oauth2/service"
	"chick/pkg/util"
	"github.com/micro/go-micro/v2/server"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func Init(c *conf.Config, svr *service.Service) {
	newRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = c.Addrs
	})

	s := micro.NewService(
		micro.Name(c.Name),
		micro.Registry(newRegistry),
		micro.WrapHandler(logWrapper),
		micro.Version("latest"),
	)

	s.Init()

	_ = oauth2.RegisterOauth2Handler(s.Server(), svr)

	go func() {
		if err := s.Run(); err != nil {
			panic(err)
		}
	}()
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%s] server received request: %s", time.Now().Format(util.TimeLayout), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}
