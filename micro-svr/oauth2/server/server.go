package server

import (
	"context"
	"fmt"
	"time"

	"chick/api/oauth2"
	"chick/micro-svr/oauth2/conf"
	"chick/micro-svr/oauth2/service"
	"chick/pkg/util"

	"github.com/micro/go-micro/v2"
	//"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"
	//"github.com/micro/go-plugins/broker/nsq/v2"
)

func Init(c *conf.Config, svr *service.Service) {
	newRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = c.Registry.Etcd.Addrs
	})
	//nsqBroker := nsq.NewBroker(func(options *broker.Options) {
	//	options.Addrs = c.Broker.Nsq.Addr
	//})

	s := micro.NewService(
		micro.Name(c.Name),
		micro.Registry(newRegistry),
		//micro.Broker(nsqBroker),
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
