package dao

import (
	"context"
	"strconv"

	"chick/micro-svr/oauth2/model"

	oauth "gopkg.in/oauth2.v3"
)

func (d *Dao) GenerateAccessToken(ctx context.Context, id, secret, code, userId string) (token *model.TokenInfo, err error) {
	tgr := &oauth.TokenGenerateRequest{
		ClientID:     id,
		ClientSecret: secret,
		Code:         code,
		UserID:       userId,
	}
	tokenInfo, err := d.Oauth2Dao.Manager.GenerateAccessToken(oauth.AuthorizationCode, tgr)
	if err != nil {
		return nil, err

	}
	token = &model.TokenInfo{
		AccessToken:  tokenInfo.GetAccess(),
		ExpiresIn:    strconv.FormatInt(int64(tokenInfo.GetAccessExpiresIn()), 10),
		RefreshToken: tokenInfo.GetRefresh(),
		Openid:       tokenInfo.GetUserID(),
	}
	return token, nil
}

func (d *Dao) VerifyAccessToken(ctx context.Context, access string) (openId string, err error) {
	tokenInfo, err := d.Oauth2Dao.Manager.LoadAccessToken(access)
	if err != nil {
		return "", err
	}
	return tokenInfo.GetUserID(), nil
}

func (d *Dao) RefreshAccessToken(ctx context.Context, id, secret, refresh string) (token *model.TokenInfo, err error) {
	tgr := &oauth.TokenGenerateRequest{
		ClientID:     id,
		ClientSecret: secret,
		Refresh:      refresh,
	}
	tokenInfo, err := d.Oauth2Dao.Manager.RefreshAccessToken(tgr)
	if err != nil {
		return nil, err
	}
	token = &model.TokenInfo{
		AccessToken:  tokenInfo.GetAccess(),
		ExpiresIn:    strconv.FormatInt(int64(tokenInfo.GetAccessExpiresIn()), 10),
		RefreshToken: tokenInfo.GetRefresh(),
		Openid:       tokenInfo.GetUserID(),
	}
	return token, nil
}

func (d *Dao) RemoveAccessToken(ctx context.Context, access string) (isOk bool, err error) {
	err = d.Oauth2Dao.Manager.RemoveAccessToken(access)
	if err != nil {
		return false, err
	}
	return true, nil
}
