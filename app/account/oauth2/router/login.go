package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {

	c.SetCookie("user_session", "session-fumingming", 30*24*60*60, "/", "localhost:8082", http.SameSiteDefaultMode, false, true)
	c.Redirect(302, "http://localhost:8082/account/oauth/callback?code=nihaowoshicode&state=qwer1234asdf")
}
