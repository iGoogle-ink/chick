package service

import (
	"context"

	"chick/micro-svr/oauth2/proto"
)

func (s *Service) AccessToken(ctx context.Context, in *proto.AccessTokenReq, reply *proto.AccessTokenReply) (err error) {

	return nil
}

func (s *Service) User(ctx context.Context, in *proto.UserReq, reply *proto.UserReply) (err error) {

	return nil
}

func (s *Service) UserOpenID(ctx context.Context, in *proto.UserOpenIDReq, reply *proto.UserOpenIDReply) (err error) {

	return nil
}
