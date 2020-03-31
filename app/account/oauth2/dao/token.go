package dao

import (
	"context"
)

func (d *Dao) AccessToken(ctx context.Context, clientId, clientSecret, code string) (err error) {
	//req := &oauth2.AccessTokenReq{ClientId: clientId, ClientSecret: clientSecret, Code: code}
	//if err != nil {
	//	return err
	//}
	//log.Println("d.oauthGrpc.AccessToken：", token.AccessToken)
	//log.Println("d.oauthGrpc.AccessToken：", token.ExpiresIn)
	//log.Println("d.oauthGrpc.AccessToken：", token.RefreshToken)
	//log.Println("d.oauthGrpc.AccessToken：", token.Openid)
	return nil
}

func (d *Dao) VerifyToken(ctx context.Context, access string) (openid string, err error) {
	//req := &oauth2.VerifyTokenReq{AccessToken: access}
	//if err != nil {
	//	return "", err
	//}
	return "token.Openid", nil
}
