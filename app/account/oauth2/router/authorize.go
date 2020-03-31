package router

import (
	"fmt"
	"net/url"

	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/web"

	"github.com/gin-gonic/gin"
)

func authorize(c *gin.Context) {
	query := c.Request.URL.Query()
	// 用户未登录，跳转到登录页面
	session, _ := c.Cookie("session")
	fmt.Println("session:", session)

	if session == "" {
		params := query.Encode()
		loginUrl, _ := url.Parse("http://localhost:8082/account/static/login")
		loginUrl.Query().Set("client_id", query.Get("client_id"))
		loginUrl.Query().Set("return_to", query.Get("redirect_uri"))
		web.Redirect(c, loginUrl.String()+"?"+params)
		return
	}
	// 已登录，获取用户id
	userId := checkAndGetUserId(session)

	req := new(model.AuthorizeReq)
	err := c.ShouldBind(req)
	if err != nil {
		web.JSON(c, nil, errno.RequestErr)
		return
	}
	fmt.Println("req:", req)
	locationUrl, err := srv.Authorize(c, userId, req.ClientKey, req.ResponseType, req.RedirectUri, req.Scope, req.State)
	if err != nil {
		web.JSON(c, nil, err)
		return
	}
	fmt.Println("locationUrl:", locationUrl)
	web.Redirect(c, locationUrl)
}

//func authorize(c *gin.Context) {
//	oauthSrv.HandleTokenRequest(c.Writer, c.Request)
//	err := oauthSrv.HandleAuthorizeRequest(c.Writer, c.Request)
//	if err != nil {
//		web.JSON(c, nil, err)
//	}
//}

func callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println("code:", code)
	fmt.Println("state:", state)
}
