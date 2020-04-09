package server

import (
	"chick/api/user"
	"chick/micro-svr/user/conf"
	"chick/micro-svr/user/service"
	"chick/pkg/log"
	"chick/pkg/server"
)

func Init(c *conf.Config, srv *service.Service) {
	s, err := server.InitServer(c.Name, "latest", c.Registry)
	if err != nil {
		log.Panic(err)
	}

	_ = user.RegisterUserHandler(s.Server.Server(), srv)

	s.Start()
}
