package dao

import (
	"context"
	"time"

	"chick/api/oauth2"
	"chick/app/account/oauth2/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) GRPCAccessToken(ctx context.Context, key, secret, code string) (reply *oauth2.AccessTokenReply, err error) {
	req := &oauth2.AccessTokenReq{
		ClientId:     key,
		ClientSecret: secret,
		Code:         code,
	}
	token, err := d.oauthGrpc.AccessToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (d *Dao) TxInsertAccessToken(ctx context.Context, tx *gorm.DB, token *model.OauthAccessToken) error {
	return tx.Create(token).Error
}

func (d *Dao) TxInsertRefreshToken(ctx context.Context, tx *gorm.DB, token *model.OauthRefreshToken) error {
	return tx.Create(token).Error
}

func (d *Dao) TxDeleteRefreshToken(ctx context.Context, tx *gorm.DB, refreshTokenInfo *model.OauthRefreshToken) error {
	return tx.Table("oauth_refresh_token").Where("token = ?", refreshTokenInfo.Token).
		Updates(map[string]int{"is_deleted": 1}).Error
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	return "", nil
}

func (d *Dao) GetRefreshToken(ctx context.Context, req *model.OauthRefreshTokenReq) (*model.OauthRefreshToken, bool) {
	refreshTokenInfo := &model.OauthRefreshToken{}
	if err := d.DB.Where("token = ?", req.RefreshToken).
		Where("expires_at > ?", time.Now()).
		Where("is_deleted = ?", 0).
		First(&refreshTokenInfo).Error; err != nil {
		return nil, false
	}
	return refreshTokenInfo, true
}
