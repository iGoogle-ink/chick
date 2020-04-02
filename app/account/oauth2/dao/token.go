package dao

import (
	"context"

	"chick/api/oauth2"
	"chick/app/account/oauth2/model"
	xtime "chick/pkg/time"

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

func (d *Dao) AccessToken(ctx context.Context, refresh string) (token *model.OauthAccessToken, err error) {
	token = new(model.OauthAccessToken)
	err = d.DB.Select([]string{"id", "client_id", "user_id", "access", "refresh", "expires_at", "scope"}).
		Where("refresh = ?", refresh).
		Where("is_deleted = 0").
		First(token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (d *Dao) UpdateAccessToken(ctx context.Context, id int, access, refresh string, expireAt xtime.Time) error {
	ups := map[string]interface{}{
		"access":     access,
		"refresh":    refresh,
		"expires_at": expireAt,
	}
	return d.DB.Table("oauth_access_token").Select([]string{"access", "refresh", "expires_at"}).
		Where("id = ?", id).Updates(ups).Error
}

func (d *Dao) InsertAccessToken(ctx context.Context, token *model.OauthAccessToken) error {
	return d.DB.Create(token).Error
}

//func (d *Dao) TxInsertRefreshToken(ctx context.Context, tx *gorm.DB, token *model.OauthRefreshToken) error {
//	return tx.Create(token).Error
//}

//func (d *Dao) TxDeleteRefreshToken(ctx context.Context, tx *gorm.DB, refresh string) error {
//	return tx.Table("oauth_refresh_token").Where("token = ?", refresh).
//		Updates(map[string]int{"is_deleted": 1}).Error
//}

func (d *Dao) TxDeleteAccessToken(ctx context.Context, tx *gorm.DB, access string) error {
	return tx.Table("oauth_access_token").Select("is_deleted").Where("token = ?", access).
		Updates(map[string]int{"is_deleted": 1}).Error
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	return "", nil
}

//func (d *Dao) GetRefreshToken(ctx context.Context, refresh string) (*model.OauthRefreshToken, bool) {
//	refreshTokenInfo := &model.OauthRefreshToken{}
//	if err := d.DB.Where("token = ?", refresh).
//		Where("expires_at > ?", time.Now()).
//		Where("is_deleted = ?", 0).
//		First(refreshTokenInfo).Error; err != nil {
//		return nil, false
//	}
//	return refreshTokenInfo, true
//}
