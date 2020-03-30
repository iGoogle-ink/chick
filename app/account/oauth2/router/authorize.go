package router

import (
	"fmt"

	"chick/pkg/web"
	"github.com/gin-gonic/gin"
)

//func authorize(c *gin.Context) {
//	query := c.Request.URL.Query()
//	// 用户未登录，跳转到登录页面
//	session, _ := c.Cookie("session")
//	if session == "" {
//		//http://localhost:8082/account/oauth/authorize?client_id=i_am_client_id&response_type=code&redirect_uri=http://auth.mxchip.com/callback&state=9e76f36043cc927475d3de381019418e7471cbdb47c00266
//		params := query.Encode()
//		loginUrl, _ := url.Parse("http://localhost:8082/account/oauth/login")
//		loginUrl.Query().Set("client_id", query.Get("client_id"))
//		loginUrl.Query().Set("return_to", query.Get("redirect_uri"))
//		web.Redirect(c, loginUrl.String()+"?"+params)
//		return
//	}
//	// 已登录，获取用户id
//	userId := getUserId(session)
//
//	req := new(model.AuthorizeReq)
//	err := c.ShouldBind(req)
//	if err != nil {
//		web.JSON(c, nil, errno.RequestErr)
//		return
//	}
//	locationUrl, err := srv.Authorize(userId, req.ClientKey, req.ResponseType, req.RedirectUri, req.Scope, req.State)
//	if err != nil {
//		web.JSON(c, nil, err)
//		return
//	}
//	web.Redirect(c, locationUrl)
//}

func authorize(c *gin.Context) {
	err := oauthSrv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		web.JSON(c, nil, err)
	}
}

func callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println("code:", code)
	fmt.Println("state:", state)
}
