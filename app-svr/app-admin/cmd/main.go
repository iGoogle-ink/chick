package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chick/app-svr/app-admin/conf"
	"chick/app-svr/app-admin/router"
	"chick/app-svr/app-admin/service"
)

func main() {
	err := conf.Parse()
	if err != nil {
		panic(err)
	}

	srv := service.New(conf.Conf)

	router.Init(conf.Conf, srv)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			log.Printf("get a signal %s, stop the oauth2 process\n", si.String())
			srv.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
