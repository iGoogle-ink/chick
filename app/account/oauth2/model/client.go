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
