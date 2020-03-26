package dao

import (
	"context"

	"chick/micro-svr/oauth2/model"

	oauth "gopkg.in/oauth2.v3"
)

func (d *Dao) GenerateAccessToken(ctx context.Context, id, secret, uri, scope string) (token *model.TokenInfo, err error) {
	tgr := &oauth.TokenGenerateRequest{
		ClientID:       id,
		ClientSecret:   secret,
		RedirectURI:    uri,
		Scope:          scope,
		AccessTokenExp: 7200,
	}
	tokenInfo, err := d.Oauth2Dao.Manager.GenerateAccessToken(oauth.AuthorizationCode, tgr)
	if err != nil {
		return nil, err

	}
	token = &model.TokenInfo{
		AccessToken: tokenInfo.GetAccess(),
		ExpiresIn:   int64(tokenInfo.GetAccessExpiresIn()),
	}
	return token, nil
}
