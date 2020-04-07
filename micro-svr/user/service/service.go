package service

import (
	"context"

	"chick/micro-svr/user/conf"
	"chick/micro-svr/user/dao"
)

var ctx = context.Background()

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Config) (srv *Service) {
	srv = &Service{
		dao: dao.New(c),
	}
	return srv
}

func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}
