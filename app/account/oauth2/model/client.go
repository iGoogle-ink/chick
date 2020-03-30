package model

// db table
type (
	OauthClient struct {
		Id          int
		Key         string
		Secret      string
		RedirectUri string
	}
)

func (m *OauthClient) TableName() string {
	return "oauth_client"
}
