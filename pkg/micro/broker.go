package micro

import (
	"chick/pkg/config"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/nsq/v2"
)

func nsqBroker(c *config.Broker) broker.Broker {
	return nsq.NewBroker(func(options *broker.Options) {
		options.Addrs = c.Nsq.Addrs
	})
}
