package config

type Registry struct {
	Etcd *Etcd
}

type Etcd struct {
	Addrs []string
}

type Broker struct {
	Nsq   *Nsq
	Kafka *Kafka
}

type Nsq struct {
	Addrs   []string
	Topic   string
	Channel string
}

type Kafka struct {
	Addrs []string
	Topic string
}

type OauthClient struct {
	Id     string
	Secret string
	Domain string
	UserId string
}
