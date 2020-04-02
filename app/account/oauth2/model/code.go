package model

type (
	CacheAuthCode struct {
		ClientKey   string `json:"client_key"`
		UserId      int    `json:"user_id"`
		Code        string `json:"code"`
		RedirectUri string `json:"redirect_uri"`
		Expires     int64  `json:"expires"`
		Scope       string `json:"scope"`
	}
)

//func (m *OauthAuthCode) TableName() string {
//	return "oauth_auth_code"
//}
