package dao

import (
	"context"
	"fmt"

	"chick/api/oauth2"
	"chick/app/account/oauth2/conf"
	"chick/pkg/orm"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceOauth2 = "micro.service.oauth2"
)

type Dao struct {
	DB        *gorm.DB
	redis     *redis.ClusterClient
	oauthGrpc oauth2.Oauth2Service
}

func New(c *conf.Config) (d *Dao) {
	fmt.Println("redisConfig:", c.Redis.Addrs)
	d = &Dao{
		DB:        orm.InitMySQL(c.MySQL),
		redis:     orm.InitRedisCluster(c.Redis),
		oauthGrpc: initOauthGrpc(c),
	}
	return d
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

func initOauthGrpc(c *conf.Config) oauth2.Oauth2Service {
	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = c.Registry.Etcd.Addrs
	})
	service := micro.NewService(
		micro.Name(c.Name),
		micro.Registry(etcdRegistry),
		micro.WrapClient(logWrap),
	)
	service.Init()
	return oauth2.NewOauth2Service(ServiceOauth2, service.Client())
}
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request to service: %s endpoint: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}
