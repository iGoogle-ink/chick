package model

type AuthorizeReq struct {
	ClientKey    string `form:"client_key"`
	ResponseType string `form:"response_type"`
	RedirectUri  string `form:"redirect_uri"`
	State        string `form:"state"`
}

type AuthorizeRsp struct {
}
