package service

import (
	"chick/app-svr/app-admin/conf"
	"chick/app-svr/app-admin/dao"
)

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
