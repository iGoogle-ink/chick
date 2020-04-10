package micro

import (
	"chick/pkg/config"
	"chick/pkg/log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

func InitServer(name, version string, registry *config.Registry, broker *config.Broker, fn func(s server.Server)) {
	if registry == nil {
		log.Panicf("InitServer(%s) register is nil", name)
	}
	etcdRegistry := etcdRegistry(registry)
	var s micro.Service
	if broker != nil {
		nsqBroker := nsqBroker(broker)
		s = micro.NewService(
			micro.Name(name),
			micro.Registry(etcdRegistry),
			micro.Broker(nsqBroker),
			micro.WrapHandler(logWrapper),
			micro.Version(version),
		)
	} else {
		s = micro.NewService(
			micro.Name(name),
			micro.Registry(etcdRegistry),
			micro.WrapHandler(logWrapper),
			micro.Version(version),
		)
	}
	//s.Init()
	fn(s.Server())

	go func() {
		if err := s.Run(); err != nil {
			log.Panicf("[%s] micro server run error(%+v).", name, err)
		}
	}()
}
