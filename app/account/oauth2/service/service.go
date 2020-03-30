package service

import (
	"context"

	"chick/app/account/oauth2/conf"
	"chick/app/account/oauth2/dao"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

var ctx = context.Background()

type Service struct {
	dao *dao.Dao
	c   *conf.Config
}

func New(c *conf.Config, mysql *gorm.DB, redisCli *redis.ClusterClient) (srv *Service) {
	srv = &Service{
		dao: dao.New(mysql, redisCli),
		c:   c,
	}
	return srv
}

func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}
