package dao

import (
	"chick/micro-svr/oauth2/conf"
	"chick/pkg/orm"
	xStore "chick/pkg/store"

	"gopkg.in/oauth2.v3/models"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type Dao struct {
	DB        *gorm.DB
	Redis     *redis.ClusterClient
	Oauth2Dao *server.Server
}

func New(c *conf.Config) (d *Dao) {
	redisCluster := orm.InitRedisCluster(c.Redis)

	d = &Dao{
		DB:        orm.InitMySQL(c.MySQL),
		Redis:     redisCluster,
		Oauth2Dao: newOauth2Dao(c.Clients, redisCluster),
	}
	return
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

func newOauth2Dao(clients []*xStore.ClientInfo, redisCli *redis.ClusterClient) *server.Server {
	mgr := manage.NewDefaultManager()

	mgr.MustTokenStorage(xStore.NewRedisClusterStoreWithCli(redisCli), nil)

	clientStore := store.NewClientStore()

	for _, v := range clients {
		clientStore.Set(v.Id, &models.Client{
			ID:     v.Id,
			Secret: v.Secret,
			Domain: v.Domain,
			UserID: v.UserId,
		})
	}

	mgr.MapClientStorage(clientStore)

	mgr.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	return server.NewDefaultServer(mgr)
}
