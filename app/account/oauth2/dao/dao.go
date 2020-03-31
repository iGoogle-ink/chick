package dao

import (
	"chick/app/account/oauth2/conf"
	"chick/pkg/orm"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB    *gorm.DB
	redis *redis.ClusterClient
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		DB:    orm.InitMySQL(c.MySQL),
		redis: orm.InitRedisCluster(c.Redis),
	}
	return d
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
