package dao

import (
	"chick/app/account/oauth2/model"
	"context"
	"time"
)

func (d *Dao) InsertAccessToken(ctx context.Context, access, refersh, scope string, clientId, userId int, expiresAt time.Time) (err error) {
	accessToken := &model.OauthAccessToken{
		ClientId:  clientId,
		UserId:    userId,
		Token:     access,
		ExpiresAt: expiresAt,
		Scope:     scope,
		IsDeleted: false,
	}

	if err := d.DB.Create(accessToken).Error; err != nil {
		return err
	}

	refershToken := &model.OauthRefreshToken{

		ClientId:  clientId,
		UserId:    userId,
		Token:     refersh,
		ExpiresAt: expiresAt,
		Scope:     scope,
		IsDeleted: false,
	}
	if err := d.DB.Create(refershToken).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	return "", nil
}
