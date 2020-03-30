package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			log.Printf("get a signal %s, stop the admin process\n", si.String())

			srv.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
