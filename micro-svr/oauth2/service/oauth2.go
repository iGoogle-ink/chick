package service

import (
	"context"

	"chick/api/oauth2"
)

// AccessToken get access_token
func (s *Service) AccessToken(ctx context.Context, in *oauth2.AccessTokenReq, reply *oauth2.AccessTokenReply) (err error) {
	token, err := s.dao.GenerateAccessToken(ctx, in.ClientId, in.ClientSecret, in.Code)
	if err != nil {
		return err
	}
	reply.AccessToken = token.AccessToken
	reply.ExpiresIn = token.ExpiresIn
	return nil
}

func (s *Service) VerifyToken(ctx context.Context, in *oauth2.VerifyTokenReq, reply *oauth2.VerifyTokenReply) (err error) {

	return nil
}

func (s *Service) RefreshToken(ctx context.Context, in *oauth2.RefreshTokenReq, reply *oauth2.RefreshTokenReply) (err error) {

	return nil
}

func (s *Service) RemoveToken(ctx context.Context, in *oauth2.RemoveTokenReq, reply *oauth2.RemoveTokenReply) (err error) {

	return nil
}
