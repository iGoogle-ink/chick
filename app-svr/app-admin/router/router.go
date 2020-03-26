package router

import (
	"net/http"

	"chick/app-svr/app-admin/conf"
	"chick/app-svr/app-admin/service"
	xhttp "chick/pkg/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var (
	srv *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	srv = s

	g := xhttp.InitServer(c.Port)

	initRoute(g.Gin)

	g.Start()
}

func initRoute(g *gin.Engine) {
	g.GET("/admin/ping", ping)

	admin := g.Group("/admin")
	{
		admin.POST("/login", login)
		mx := admin.Group("/xxx", srv.Verify)
		{
			mx.GET("/info")
		}
	}
}

func ping(c *gin.Context) {
	c.Render(http.StatusOK, render.JSON{})
}
