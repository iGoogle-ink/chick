package server

import (
	"chick/pkg/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func etcdRegistry(c *config.Registry) (reg registry.Registry) {
	return etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = c.Etcd.Addrs
	})
}
