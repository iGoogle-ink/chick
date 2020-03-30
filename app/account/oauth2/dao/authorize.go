package dao

import (
	"context"

	"chick/app/account/oauth2/model"
)

func (d *Dao) InsertAuthorizationCode(ctx context.Context, oauthCode *model.OauthAuthCode) error {
	return d.DB.Create(oauthCode).Error
}

func (d *Dao) AuthorizeCode() {
	d.oauthDao.HandleAuthorizeRequest()
}
