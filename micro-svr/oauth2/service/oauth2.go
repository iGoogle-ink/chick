package service

import (
	"context"

	"chick/api/oauth2"
)

// AccessToken get access_token
func (s *Service) AccessToken(ctx context.Context, in *oauth2.AccessTokenReq, out *oauth2.AccessTokenReply) (err error) {
	// todo: get user openid

	token, err := s.dao.GenerateAccessToken(ctx, in.ClientId, in.ClientSecret, in.Code, "openid")
	if err != nil {
		return err
	}
	out.AccessToken = token.AccessToken
	out.ExpiresIn = token.ExpiresIn
	out.RefreshToken = token.RefreshToken
	out.Openid = token.Openid
	return nil
}

func (s *Service) VerifyToken(ctx context.Context, in *oauth2.VerifyTokenReq, out *oauth2.VerifyTokenReply) (err error) {
	isOk, _ := s.dao.VerifyAccessToken(ctx, in.AccessToken)
	out.IsOk = isOk
	return nil
}

func (s *Service) RefreshToken(ctx context.Context, in *oauth2.RefreshTokenReq, out *oauth2.RefreshTokenReply) (err error) {
	token, err := s.dao.RefreshAccessToken(ctx, in.ClientId, in.ClientSecret, in.RefreshToken)
	if err != nil {
		return err
	}
	out.AccessToken = token.AccessToken
	out.ExpiresIn = token.ExpiresIn
	out.RefreshToken = token.RefreshToken
	out.Openid = token.Openid
	return nil
}

func (s *Service) RemoveToken(ctx context.Context, in *oauth2.RemoveTokenReq, out *oauth2.RemoveTokenReply) (err error) {
	isOk, _ := s.dao.RemoveAccessToken(ctx, in.AccessToken)
	out.IsOk = isOk
	return nil
}
