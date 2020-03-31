package router

import (
	"net/http"

	"chick/errno"
	"chick/pkg/web"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	uname := c.Query("username")
	pwd := c.Query("password")

	if uname == "mxchip" && pwd == "mxchip" {
		c.SetCookie("session", "session-fumingming", 30*24*60*60, "/", "localhost:8082", http.SameSiteDefaultMode, false, true)
		web.Redirect(c, "http://localhost:8082/account/static/auth")
		return
	}
	web.JSON(c, nil, errno.Unauthorized)
}
