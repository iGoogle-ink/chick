package server

import (
	"chick/api/oauth2"
	"chick/micro-svr/oauth2/conf"
	"chick/micro-svr/oauth2/service"
	"chick/pkg/server"
)

func Init(c *conf.Config, svr *service.Service) {

	s, err := server.InitServer(c.Name, "latest", c.Registry)
	if err != nil {
		panic(err)
	}

	_ = oauth2.RegisterOauth2Handler(s.Server.Server(), svr)

	s.Start()
}
