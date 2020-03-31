package model

import (
	xtime "chick/pkg/time"
)

// db table
type (
	OauthClient struct {
		Id          int
		Key         string
		Secret      string
		RedirectUri string
		IsDeleted   int
		Mtime       xtime.Time
	}
)

func (m *OauthClient) TableName() string {
	return "oauth_client"
}
