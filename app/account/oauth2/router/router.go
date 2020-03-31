package router

import (
	"net/http"

	"chick/app/account/oauth2/conf"
	"chick/app/account/oauth2/service"
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
)

var (
	srv *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	srv = s

	g := web.InitServer(c.Port)

	g.Gin.LoadHTMLGlob("static/*")

	initRoute(g.Gin)

	g.Start()
}

func initRoute(g *gin.Engine) {
	g.GET("/account/ping", ping)

	acc := g.Group("/account")
	{
		oa := acc.Group("/oauth")
		{
			oa.GET("/authorize", authorize)
			oa.GET("/login", login)
			oa.POST("/token", token)
			oa.POST("/refresh", refreshToken)
		}
		html := acc.Group("/static")
		{
			html.GET("/login", loginHtml)
			html.GET("/auth", authHtml)

		}
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong"})
}

func getUserId(session string) (id int) {

	return 1
}
