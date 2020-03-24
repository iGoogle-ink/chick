package server

import (
	"chick/micro-svr/oauth2/conf"
	"chick/micro-svr/oauth2/service"
	"chick/proto/oauth2"

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
		micro.Version("latest"),
	)

	oauth2.RegisterOauth2Handler(s.Server(), svr)

	go func() {
		if err := s.Run(); err != nil {
			panic(err)
		}
	}()
}
