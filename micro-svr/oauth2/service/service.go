package service

import (
	"chick/micro-svr/oauth2/conf"
	"chick/micro-svr/oauth2/dao"
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
