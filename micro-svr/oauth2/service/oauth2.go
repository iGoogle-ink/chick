package service

import (
	"context"

	"chick/api/oauth2"
)

// InsertAccessToken 获取 access token
func (s *Service) AccessToken(ctx context.Context, in *oauth2.AccessTokenReq, out *oauth2.AccessTokenReply) (err error) {
	// todo: get user openid
	out.AccessToken = "e110e93a549281a731ae694714789cf6"
	out.ExpiresIn = "1521312134"
	out.RefreshToken = "0acb42bcc1b6070e5be24628f6fac748"
	out.Openid = "111222"
	return nil
}

// VerifyToken 验证 access token 并返回 openid
func (s *Service) VerifyToken(ctx context.Context, in *oauth2.VerifyTokenReq, out *oauth2.VerifyTokenReply) (err error) {
	openId, err := s.dao.VerifyAccessToken(ctx, in.AccessToken)
	if err != nil {
		return err
	}
	out.Openid = openId
	return nil
}

// RefreshToken 刷新 access token
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

// RemoveToken 删除 access token
func (s *Service) RemoveToken(ctx context.Context, in *oauth2.RemoveTokenReq, out *oauth2.RemoveTokenReply) (err error) {
	isOk, _ := s.dao.RemoveAccessToken(ctx, in.AccessToken)
	out.IsOk = isOk
	return nil
}

// RemoveToken 删除 access token
func (s *Service) AuthorizationCode(ctx context.Context, in *oauth2.AuthorizationCodeReq, out *oauth2.AuthorizationCodeReply) (err error) {
	return nil
}
