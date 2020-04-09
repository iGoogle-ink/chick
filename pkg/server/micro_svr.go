package server

import (
	"chick/pkg/config"
	"chick/pkg/log"

	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
)

type MicroServer struct {
	Server    micro.Service
	MicroName string
	Version   string
}

func InitServer(name, version string, registry *config.Registry, broker ...*config.Broker) (svr *MicroServer, err error) {
	if registry == nil {
		return nil, errors.New("register is nil")
	}
	etcdRegistry := etcdRegistry(registry)
	var s micro.Service
	if len(broker) > 0 {
		nsqBroker := nsqBroker(broker[0])
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

	s.Init()

	svr = &MicroServer{
		Server:    s,
		MicroName: name,
		Version:   version,
	}
	return svr, nil
}

func (m *MicroServer) Start() {
	go func() {
		if err := m.Server.Run(); err != nil {
			log.Panicf("micro server run error(%+v).", err)
		}
	}()
}
