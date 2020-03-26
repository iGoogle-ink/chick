package service

import (
	"context"

	"chick/app/web-svr/admin/conf"
	"chick/app/web-svr/admin/dao"
	"chick/pkg/http"

	"github.com/gin-gonic/gin"
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

func (s *Service) Verify(c *gin.Context) {
	req := c.Request
	token := req.Header.Get("Authorization")

	_, err := s.dao.VerifyToken(ctx, token)
	if err != nil {
		http.JSON(c, nil, err)
		return
	}

}
