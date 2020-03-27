package main

import (
	"chick/app/account/oauth2/conf"
	"chick/app/account/oauth2/router"
	"chick/app/account/oauth2/service"
)

func main() {
	err := conf.Parse()
	if err != nil {
		panic(err)
	}

	srv := service.New(conf.Conf)

	router.Init(conf.Conf, srv)

}
