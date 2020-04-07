package dao

import (
	"chick/micro-svr/user/conf"
	"chick/pkg/orm"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB       *gorm.DB
	redisCli *redis.ClusterClient
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		DB:       orm.InitMySQL(c.MySQL),
		redisCli: orm.InitRedisCluster(c.Redis),
	}
	return
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
