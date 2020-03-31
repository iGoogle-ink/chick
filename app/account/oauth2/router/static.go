package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func authHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "auth.html", nil)
}

func registerHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
