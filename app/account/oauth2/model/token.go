package model

import (
	"time"
)

type (
	AccessTokenReq struct {
		ClientId     string `json:"client_id" binding:"required"`
		ClientSecret string `json:"client_secret" binding:"required"`
		Code         string `json:"code" binding:"required"`
		GrantType    string `json:"grant_type" binding:"required"`
	}

	AccessTokenReply struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	OauthAccessToken struct {
		Id        int `gorm:"primary_key"`
		ClientId  int
		UserId    int
		Token     string
		ExpiresAt time.Time
		Scope     string
		IsDeleted bool
	}
	OauthRefreshToken struct {
		Id        int `gorm:"primary_key"`
		ClientId  int
		UserId    int
		Token     string
		ExpiresAt time.Time
		Scope     string
		IsDeleted bool
	}
)

func (m *OauthAccessToken) TableName() string {
	return "oauth_access_token"
}
func (m *OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}
