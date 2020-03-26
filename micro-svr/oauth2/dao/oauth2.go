package dao

import (
	"context"
	"strconv"

	"chick/micro-svr/oauth2/model"

	oauth "gopkg.in/oauth2.v3"
)

func (d *Dao) GenerateAccessToken(ctx context.Context, id, secret, code string) (token *model.TokenInfo, err error) {
	tgr := &oauth.TokenGenerateRequest{
		ClientID:     id,
		ClientSecret: secret,
		Code:         code,
	}
	tokenInfo, err := d.Oauth2Dao.Manager.GenerateAccessToken(oauth.AuthorizationCode, tgr)
	if err != nil {
		return nil, err

	}
	token = &model.TokenInfo{
		AccessToken: tokenInfo.GetAccess(),
		ExpiresIn:   strconv.FormatInt(int64(tokenInfo.GetAccessExpiresIn()), 10),
	}
	return token, nil
}

func (d *Dao) VerifyAccessToken(ctx context.Context) {
	d.Oauth2Dao.Manager.LoadAccessToken()

}
