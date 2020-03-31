package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"chick/app/account/oauth2/model"
	"chick/errno"
	xtime "chick/pkg/time"
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
	code := generateCode(userId, client)

	codeInsert := &model.OauthAuthCode{
		ClientId:    client.Id,
		UserId:      userId,
		Code:        code,
		RedirectUri: redUri,
		ExpiresAt:   xtime.Time(time.Now().Add(time.Duration(5) * time.Minute).Unix()),
		Scope:       scope,
	}
	// code 写入数据库
	err = s.dao.InsertCode(ctx, codeInsert)
	if err != nil {
		return "", errno.New(-2, "生成Code失败")
	}
	return redUri + "?code=" + code + "&state=" + state, nil
}

func generateCode(userId int, client *model.OauthClient) (code string) {
	hash := md5.New()
	hash.Write([]byte(client.Key))
	hash.Write([]byte(strconv.Itoa(userId)))
	hash.Write([]byte(time.Now().String()))
	return hex.EncodeToString(hash.Sum(nil))
}
