package dao

import (
	"chick/api/oauth2"
	"chick/app/account/oauth2/conf"
	"chick/pkg/micro"
	"chick/pkg/orm"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/client"
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

func initOauthGrpc(c *conf.Config) (cli oauth2.Oauth2Service) {
	micro.InitClient(c.Name, "latest", c.Registry, nil, func(c client.Client) {
		cli = oauth2.NewOauth2Service(ServiceOauth2, c)
	})
	return cli
}
