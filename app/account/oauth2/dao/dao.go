package dao

import (
	"context"
	"fmt"

	"chick/api/oauth2"
	"chick/app/account/oauth2/conf"
	"chick/pkg/orm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB        *gorm.DB
	redis     *redis.ClusterClient
	oauthGrpc oauth2.Oauth2Service
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		DB:        orm.InitMySQL(c.MySQL),
		redis:     orm.InitRedisCluster(c.Redis),
		oauthGrpc: initOauthGrpc(),
	}
	return d
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

func initOauthGrpc() oauth2.Oauth2Service {
	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Name("app.account.oauth2"),
		micro.Registry(etcdRegistry),
		micro.WrapClient(logWrap),
	)
	service.Init()
	return oauth2.NewOauth2Service("micro.service.oauth2", service.Client())
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
