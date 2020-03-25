package router

import (
	"chick/pkg/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// todo: bind some parameter

	rsp, err := srv.Login()

	http.JSON(c, rsp, err)
}
