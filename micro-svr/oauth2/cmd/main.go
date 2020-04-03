package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chick/micro-svr/oauth2/conf"
	"chick/micro-svr/oauth2/server"
	"chick/micro-svr/oauth2/service"
)

func main() {

	err := conf.Parse()
	if err != nil {
		panic(err)
	}

	srv := service.New(conf.Conf)

	server.Init(conf.Conf, srv)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			log.Printf("%s: get a signal %s, stop the oauth2 process\n", conf.Conf.Name, si.String())
			srv.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
