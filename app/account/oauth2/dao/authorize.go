package dao

import (
	"context"

	"chick/app/account/oauth2/model"
)

func (d *Dao) InsertAuthorizationCode(ctx context.Context, oauthCode *model.OauthAuthCode) error {
	return d.DB.Select([]string{"client_id", "user_id", "code", "redirect_uri", "scope", "expires_at"}).Create(oauthCode).Error
}

func (d *Dao) AuthorizeCode() {
}
