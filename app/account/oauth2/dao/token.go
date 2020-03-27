package dao

import (
	"context"
	"log"

	"chick/api/oauth2"
)

func (d *Dao) AccessToken(ctx context.Context, clientId, clientSecret, code string) (err error) {
	req := &oauth2.AccessTokenReq{ClientId: clientId, ClientSecret: clientSecret, Code: code}
	token, err := d.oauthGrpc.AccessToken(ctx, req)
	if err != nil {
		return err
	}
	log.Println("d.oauthGrpc.AccessToken：", token.AccessToken)
	log.Println("d.oauthGrpc.AccessToken：", token.ExpiresIn)
	log.Println("d.oauthGrpc.AccessToken：", token.RefreshToken)
	log.Println("d.oauthGrpc.AccessToken：", token.Openid)
	return nil
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	req := &oauth2.VerifyTokenReq{AccessToken: access}
	token, err := d.oauthGrpc.VerifyToken(ctx, req)
	if err != nil {
		return "", err
	}
	log.Println("d.oauthGrpc.VerifyToken:", token.Openid)
	return token.Openid, nil
}
