package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"

	"chick/app/account/oauth2/model"
	"chick/errno"

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
		log.Printf("s.dao.GetClient(%s) err:%+v,\n", clientKey, err)
		return "", errno.New(-1, "查找Client失败")
	}
	// 生成code
	code := generateCode(userId, client.Key)

	codeCache := &model.CacheAuthCode{
		ClientKey:   client.Key,
		UserId:      userId,
		Code:        code,
		RedirectUri: redUri,
		Expires:     300, // 有效期 300s
		Scope:       scope,
	}
	// code 写入缓存
	if err = s.dao.AddCacheAuthCode(ctx, codeCache); err != nil {
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
