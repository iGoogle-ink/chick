package service

import (
	"chick/app/account/oauth2/conf"
	"chick/app/account/oauth2/dao"
)

type Service struct {
	dao *dao.Dao
	c   *conf.Config
}

func New(c *conf.Config) (srv *Service) {
	srv = &Service{
		dao: dao.New(c),
		c:   c,
	}
	return srv
}

func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}
