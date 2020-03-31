package dao

import (
	"chick/app/account/oauth2/model"
	"context"
)

func (d *Dao) GetCodeInfoByCode(ctx context.Context, code string) (*model.OauthAuthCode, error) {
	dbCode := &model.OauthAuthCode{}
	if err := d.DB.Select([]string{"id", "client_id", "user_id", "code", "expires_at", "scope"}).
		Where("code = ? and is_deleted = ?", code, 0).Order("ctime desc").First(dbCode).Error; err != nil {
		return nil, err
	}

	return dbCode, nil
}

func (d *Dao) DeleteCodeInfoByCode(ctx context.Context, code string) error {

	return d.DB.Table("oauth_auth_code").Where("code = ?", code).
		Updates(map[string]int{"is_deleted": 1}).Error
}

func (d *Dao) CheckClientInfo(ctx context.Context, clientId int, clientKey string, clientSecret string) (bool, error) {

	clientInfo := &model.OauthClient{}
	if err := d.DB.Where("id = ?", clientId).First(clientInfo).Error; err != nil {
		return false, err
	}
	return clientKey == clientInfo.Key && clientSecret == clientInfo.Secret, nil
}
