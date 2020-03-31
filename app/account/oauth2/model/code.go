package model

import (
	xtime "chick/pkg/time"
)

type (
	OauthAuthCode struct {
		Id          int `gorm:"primary_key"`
		ClientId    int
		UserId      int
		Code        string
		RedirectUri string
		ExpiresAt   xtime.Time
		Scope       string
		IsDeleted   int
		Mtime       xtime.Time
	}
)

func (m *OauthAuthCode) TableName() string {
	return "oauth_auth_code"
}
