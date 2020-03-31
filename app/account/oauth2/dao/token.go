package dao

import (
	"chick/app/account/oauth2/model"

	"github.com/jinzhu/gorm"

	"context"
)

func (d *Dao) TxInsertAccessToken(tx *gorm.DB, token *model.OauthAccessToken) error {
	if err := d.DB.Create(token).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) TxInsertRefreshToken(tx *gorm.DB, token *model.OauthRefreshToken) error {
	if err := d.DB.Create(token).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	return "", nil
}
