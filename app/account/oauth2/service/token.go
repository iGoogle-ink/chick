package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"chick/app/account/oauth2/conf"
	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/log"
	xtime "chick/pkg/time"

	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
)

func (s *Service) GRPCAccessToken(ctx context.Context) {
	var (
		key    = "la4rDuxPhBoIDyAq"
		secret = "GJQishYkUAQhKUOvYIfpnytbagATZI4L"
		code   = "a92a1382899bef688549b3a8244a2b0d"
	)
	token, err := s.dao.GRPCAccessToken(ctx, key, secret, code)
	if err != nil {
		log.Error("s.dao.GRPCAccessToken:", err)
		return
	}
	log.Debug("token.AccessToken:", token.AccessToken)
	log.Debug("token.RefreshToken:", token.RefreshToken)
	log.Debug("token.ExpiresIn:", token.ExpiresIn)
}

// AccessToken
func (s *Service) AccessToken(ctx context.Context, req *model.AccessTokenReq) (*model.AccessTokenRsp, error) {
	// 获取code
	code, err := s.dao.CacheAuthCode(ctx, req.Code)
	if err != nil {
		if err == redis.Nil {
			return nil, errno.CodeExpired
		}
		return nil, errno.InvalidCode
	}
	// 获取Client信息
	client, err := s.dao.GetClient(ctx, req.ClientId)
	if err != nil {
		return nil, errno.InvalidClient
	}
	// 验证 key 和 secret
	if !(req.ClientId == client.Key && req.ClientSecret == client.Secret) {
		return nil, errno.InvalidClient
	}
	// 生成token
	access := generateToken(req.ClientId, code.UserId)
	refresh := generateToken(req.ClientId, code.UserId)
	// 有效期
	expireAt := xtime.Time(time.Now().Add(time.Second * time.Duration(conf.Conf.TokenExpiresIn)).Unix())
	// 插入access_token
	accessToken := &model.OauthAccessToken{
		ClientId:  client.Id,
		UserId:    code.UserId,
		Access:    access,
		Refresh:   refresh,
		ExpiresAt: expireAt,
		Scope:     code.Scope,
	}
	if err := s.dao.InsertAccessToken(ctx, accessToken); err != nil {
		return nil, err
	}
	// 返回 Rsp
	rsp := &model.AccessTokenRsp{
		AccessToken:  access,
		ExpiresIn:    conf.Conf.TokenExpiresIn,
		RefreshToken: refresh,
	}
	return rsp, nil
}

func (s *Service) RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (*model.AccessTokenRsp, error) {
	// 获取Client信息
	client, err := s.dao.GetClient(ctx, req.ClientId)
	if err != nil {
		return nil, errno.InvalidClient
	}
	// 验证 key 和 secret
	if !(req.ClientId == client.Key && req.ClientSecret == client.Secret) {
		return nil, errno.InvalidClient
	}
	// 获取原有的 AccessToken 信息
	at, err := s.dao.AccessToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, errno.InvalidRefreshToken
	}
	// 验证 RefreshToken 有效性
	if !at.CheckRefreshToken(req.RefreshToken) {
		return nil, errno.InvalidRefreshToken
	}
	// 生成token
	access := generateToken(req.ClientId, at.UserId)
	refresh := generateToken(req.ClientId, at.UserId)
	// 有效期
	expireAt := xtime.Time(time.Now().Add(time.Second * time.Duration(conf.Conf.TokenExpiresIn)).Unix())
	// 更新 access_token
	if err = s.dao.UpdateAccessToken(ctx, at.Id, access, refresh, expireAt); err != nil {
		return nil, errno.New(423, "刷新AccessToken失败")
	}
	// 返回Rsp
	rsp := &model.AccessTokenRsp{
		AccessToken:  access,
		ExpiresIn:    conf.Conf.TokenExpiresIn,
		RefreshToken: refresh,
	}
	return rsp, nil
}

func generateToken(clientKey string, userId int) (token string) {
	buf := bytes.NewBufferString(clientKey)
	buf.WriteString(strconv.Itoa(userId))
	buf.WriteString(time.Now().String())
	buf.WriteString(uuid.New().String())
	hash := md5.New()
	hash.Write(buf.Bytes())
	return hex.EncodeToString(hash.Sum(nil))
}
