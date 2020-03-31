package model

import "time"

// db table
type (
	OauthClient struct {
		Id          int
		Key         string
		Secret      string
		RedirectUri string
		IsDeleted   int
		Mtime       time.Time
	}
)

func (m *OauthClient) TableName() string {
	return "oauth_client"
}
