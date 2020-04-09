package server

import (
	"chick/api/user"
	"chick/micro-svr/user/conf"
	"chick/micro-svr/user/service"
	"chick/pkg/log"
	"chick/pkg/micro"

	"github.com/micro/go-micro/v2/server"
)

func Init(c *conf.Config, srv *service.Service) {
	micro.InitServer(c.Name, "latest", c.Registry, nil, func(s server.Server) {
		if err := user.RegisterUserHandler(s, srv); err != nil {
			log.Panic(err)
		}
	})
}
