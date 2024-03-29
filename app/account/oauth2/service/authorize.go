package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/log"

	"github.com/google/uuid"
)

// Authorize
func (s *Service) Authorize(ctx context.Context, userId int, clientKey, rspType, redUri, scope, state string) (locationUrl string, err error) {
	if rspType != "code" {
		return "", errno.RequestErr
	}
	// 查找Client信息
	client, err := s.dao.GetClient(ctx, clientKey)
	if err != nil {
		log.Errorf("s.dao.GetClient(%s) err:%+v.", clientKey, err)
		return "", errno.New(-1, "查找Client失败")
	}
	// 生成code
	code := generateCode(userId, client.Key)

	codeCache := &model.CacheAuthCode{
		ClientKey:   client.Key,
		UserId:      userId,
		Code:        code,
		RedirectUri: redUri,
		Expires:     600, // 有效期 600s
		Scope:       scope,
	}
	// code 写入缓存
	if err = s.dao.AddCacheAuthCode(ctx, codeCache); err != nil {
		log.Errorf("s.dao.AddCacheAuthCode(%v) error:%+v.", codeCache, err)
		return "", errno.New(-2, "生成Code失败")
	}
	return redUri + "?code=" + code + "&state=" + state, nil
}

func generateCode(userId int, clientKey string) (code string) {
	hash := md5.New()
	hash.Write([]byte(clientKey))
	hash.Write([]byte(strconv.Itoa(userId)))
	hash.Write([]byte(uuid.New().String()))
	return hex.EncodeToString(hash.Sum(nil))
}
