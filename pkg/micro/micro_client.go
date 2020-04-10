package micro

import (
	"chick/pkg/config"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func InitClient(cliName, version string, registry *config.Registry, broker *config.Broker, fn func(client.Client)) {
	etcdRegistry := etcdRegistry(registry)
	var s micro.Service
	if broker != nil {
		nsqBroker := nsqBroker(broker)
		s = micro.NewService(
			micro.Name(cliName),
			micro.Registry(etcdRegistry),
			micro.Broker(nsqBroker),
			micro.WrapClient(logClientWrap),
			micro.Version(version),
		)
	} else {
		s = micro.NewService(
			micro.Name(cliName),
			micro.Registry(etcdRegistry),
			micro.WrapClient(logClientWrap),
			micro.Version(version),
		)
	}
	//s.Init()
	fn(s.Client())
}
