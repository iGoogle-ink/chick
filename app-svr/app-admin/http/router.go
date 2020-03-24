package http

import "github.com/gin-gonic/gin"

func router(g *gin.Engine) {
	g.Group("api")
}
