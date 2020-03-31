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
			oa.Any("/authorize", authorize)   // 请求授权
			oa.POST("/login", login)          // 授权登录
			oa.POST("/register", register)    // 注册
			oa.POST("/token", token)          // 换取AccessToken
			oa.POST("/refresh", refreshToken) // 刷新AccessToken
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

func checkAndGetUserId(session string) (id int) {

	return 1
}
