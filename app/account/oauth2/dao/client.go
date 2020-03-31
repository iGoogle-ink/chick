package dao

import (
	"context"

	"chick/app/account/oauth2/model"
)

func (d *Dao) GetClient(ctx context.Context, clientKey string) (clientInfo *model.OauthClient, err error) {
	clientInfo = new(model.OauthClient)
	if err = d.DB.Select([]string{"id", "`key`", "secret", "redirect_uri"}).Where("`key` = ?", clientKey).First(clientInfo).Error; err != nil {
		return nil, err
	}
	return clientInfo, nil
}
