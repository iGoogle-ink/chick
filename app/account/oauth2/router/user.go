package router

import (
	"net/http"
	"net/url"

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
	c.SetCookie("session", session, 30*24*60*60, "/", "localhost:2233", http.SameSiteDefaultMode, false, true)
	web.Redirect(c, "http://localhost:2233/account/static/auth?"+params)
}

func register(c *gin.Context) {
	req := new(model.RegisterReq)
	err := c.ShouldBind(req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	params := c.Request.URL.Query().Encode()
	if err = srv.Register(c, req); err != nil {
		web.JSON(c, nil, err)
		return
	}
	loginUrl, _ := url.Parse("http://localhost:2233/account/static/login")
	web.Redirect(c, loginUrl.String()+"?"+params)
}
