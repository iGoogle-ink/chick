package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"chick/app/account/oauth2/model"
	"chick/errno"

	"github.com/jinzhu/gorm"
	"gopkg.in/oauth2.v3/utils/uuid"
)

// GetAccessToken
func (s *Service) GetAccessToken(ctx context.Context, clientKey, clientSecret, code, grantType string) (*model.AccessTokenReply, error) {

	// 校验code
	dbCode, err := s.dao.GetCodeInfoByCode(ctx, code)
	if (err != nil && err != gorm.ErrRecordNotFound) || (dbCode == nil) {
		return nil, errno.InvalidCode
	}
	//校验code 过期
	if dbCode.ExpiresAt.Before(time.Now()) {
		return nil, errno.CodeExpired
	}
	// 校验 client_key, client_secret
	if ok, err := s.dao.CheckClientInfo(ctx, dbCode.ClientId, clientKey, clientSecret); err != nil {
		return nil, errno.InvalidCode
	} else {
		if !ok {
			return nil, errno.InvalidClient
		}
	}

	// 生成token
	access, refresh, _ := GenerateCode(clientKey, dbCode.UserId, true)
	expireAt := time.Now().Add(time.Hour * 24 * 7)
	// 插入token
	if err := s.dao.InsertAccessToken(ctx, access, refresh, dbCode.Scope, dbCode.ClientId, dbCode.UserId, expireAt); err != nil {
		return nil, errno.ServerErr
	}
	// 删除code
	if err := s.dao.DeleteCodeInfoByCode(ctx, code); err != nil {
		return nil, errno.ServerErr
	}
	reply := &model.AccessTokenReply{
		AccessToken:  access,
		ExpiresIn:    24 * 7 * 60,
		RefreshToken: refresh,
	}
	return reply, nil
}

func GenerateCode(clientKey string, userId int, isGenRefresh bool) (string, string, error) {

	buf := bytes.NewBufferString(clientKey)
	buf.WriteString(strconv.Itoa(userId))
	buf.WriteString(time.Now().String())
	buf.WriteString(uuid.Must(uuid.NewRandom()).String())

	hash := md5.New()
	hash.Write(buf.Bytes())
	accessToken := hex.EncodeToString(hash.Sum(nil))
	refreshToken := ""

	if isGenRefresh {
		buf.WriteString(uuid.Must(uuid.NewRandom()).String())
		refreshHash := md5.New()
		refreshHash.Write(buf.Bytes())
		refreshToken = hex.EncodeToString(refreshHash.Sum(nil))
	}

	return accessToken, refreshToken, nil
}
