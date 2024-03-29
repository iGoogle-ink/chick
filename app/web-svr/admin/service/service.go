package service

import (
	"context"

	"chick/app/web-svr/admin/conf"
	"chick/app/web-svr/admin/dao"
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
