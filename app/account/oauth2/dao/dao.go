package dao

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB    *gorm.DB
	redis *redis.ClusterClient
}

func New(mysql *gorm.DB, redisCli *redis.ClusterClient) (d *Dao) {
	d = &Dao{
		DB:    mysql,
		redis: redisCli,
	}
	return d
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
