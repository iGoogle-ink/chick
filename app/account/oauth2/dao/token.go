package dao

import (
	"chick/app/account/oauth2/model"
	"context"

	"github.com/jinzhu/gorm"
)

func (d *Dao) TxInsertAccessToken(ctx context.Context, tx *gorm.DB, token *model.OauthAccessToken) error {
	return tx.Create(token).Error
}

func (d *Dao) TxInsertRefreshToken(ctx context.Context, tx *gorm.DB, token *model.OauthRefreshToken) error {
	return tx.Create(token).Error
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	return "", nil
}
