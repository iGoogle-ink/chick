package router

import (
	"net/http"

	"chick/app/web-svr/admin/conf"
	"chick/app/web-svr/admin/service"
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var (
	srv *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	srv = s

	g := web.InitServer(c.HttpPort)

	initRoute(g.Gin)

	g.Start()
}

func initRoute(g *gin.Engine) {
	g.GET("/admin/ping", ping)

	admin := g.Group("/admin")
	{
		admin.POST("/login", login)
		mx := admin.Group("/xxx")
		{
			mx.GET("/info")
		}
	}
}

func ping(c *gin.Context) {
	c.Render(http.StatusOK, render.JSON{})
}
