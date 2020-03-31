package model

import (
	"time"
)

type (
	AccessTokenReq struct {
		ClientId     string `form:"client_id" json:"client_id" binding:"required"`
		ClientSecret string `form:"client_secret" json:"client_secret" binding:"required"`
		Code         string `form:"code" json:"code" binding:"required"`
		GrantType    string `form:"grant_type" json:"grant_type" binding:"required"`
	}

	AccessTokenRsp struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenId       string `json:"open_id"`
	}

	OauthAccessToken struct {
		Id        int `gorm:"primary_key"`
		ClientId  int
		UserId    int
		Token     string
		ExpiresAt time.Time
		Scope     string
		IsDeleted int
	}
	OauthRefreshToken struct {
		Id        int `gorm:"primary_key"`
		ClientId  int
		UserId    int
		Token     string
		ExpiresAt time.Time
		Scope     string
		IsDeleted int
	}

	OauthRefreshTokenReq struct {
		ClientId     string `form:"client_id" json:"client_id"`
		GrantType    string `form:"grant_type" json:"grant_type"`
		RefreshToken string `form:"refresh_token" json:"refresh_token"`
	}
)

func (m *OauthAccessToken) TableName() string {
	return "oauth_access_token"
}
func (m *OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}
