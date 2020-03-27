package router

import (
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// todo: bind some parameter

	rsp, err := srv.Login()

	web.JSON(c, rsp, err)
}
