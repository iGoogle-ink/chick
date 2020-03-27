package model

type AuthorizeReq struct {
	ClientId     string `form:"client_id"`
	ResponseType string `form:"response_type"`
	RedirectUri  string `form:"redirect_uri"`
	State        string `form:"state"`
}

type AuthorizeRsp struct {
}
