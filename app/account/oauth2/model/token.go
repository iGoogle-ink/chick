package model

import (
	xtime "chick/pkg/time"
)

type (
	OauthAccessToken struct {
		Id        int `gorm:"primary_key"`
		ClientId  int
		UserId    int
		Access    string
		Refresh   string
		ExpiresAt xtime.Time
		Scope     string
		IsDeleted int
	}
)

func (m *OauthAccessToken) TableName() string {
	return "oauth_access_token"
}

type (
	AccessTokenReq struct {
		ClientId     string `form:"client_id" json:"client_id"`
		ClientSecret string `form:"client_secret" json:"client_secret"`
		Code         string `form:"code" json:"code"`
	}

	AccessTokenRsp struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenReq struct {
		ClientId     string `form:"client_id" json:"client_id"`
		ClientSecret string `form:"client_secret" json:"client_secret"`
		RefreshToken string `form:"refresh_token" json:"refresh_token"`
	}
)
