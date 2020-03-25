package service

import (
	"context"

	"chick/api/oauth2"
)

func (s *Service) AccessToken(ctx context.Context, in *oauth2.AccessTokenReq, reply *oauth2.AccessTokenReply) (err error) {

	return nil
}

func (s *Service) User(ctx context.Context, in *oauth2.UserReq, reply *oauth2.UserReply) (err error) {

	return nil
}

func (s *Service) UserOpenID(ctx context.Context, in *oauth2.UserOpenIDReq, reply *oauth2.UserOpenIDReply) (err error) {

	return nil
}
