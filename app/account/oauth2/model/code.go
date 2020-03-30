package model

import "time"

type (
	OauthAuthCode struct {
		Id          int `gorm:"primary_key"`
		ClientId    int
		UserId      int
		Code        string
		RedirectUri string
		ExpiresAt   time.Time
		Scope       string
	}
)

func (m *OauthAuthCode) TableName() string {
	return "oauth_auth_code"
}
