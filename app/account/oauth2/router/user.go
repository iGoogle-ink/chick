package router

import (
	"net/http"

	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/web"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	req := new(model.LoginReq)
	err := c.ShouldBind(req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	params := c.Request.URL.Query().Encode()

	session, err := srv.Login(c, req)
	if err != nil {
		web.JSON(c, nil, err)
		return
	}
	c.SetCookie("session", session, 30*24*60*60, "/", "localhost:8082", http.SameSiteDefaultMode, false, true)
	web.Redirect(c, "http://localhost:8082/account/static/auth?"+params)

}

func register(c *gin.Context) {
	req := new(model.RegisterReq)
	err := c.ShouldBind(req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	params := c.Request.URL.Query().Encode()

	srv.Register(c, req)
}
