package model

type (
	AuthorizeReq struct {
		ClientKey    string `form:"client_id"`
		ResponseType string `form:"response_type"`
		RedirectUri  string `form:"redirect_uri"`
		Scope        string `form:"scope"`
		State        string `form:"state"`
	}
)
