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
func (s *Service) GetAccessToken(ctx context.Context, req *model.AccessTokenReq) (*model.AccessTokenRsp, error) {

	// 校验code
	dbCode, err := s.dao.GetCodeInfoByCode(ctx, req.Code)
	if err != nil {
		return nil, errno.InvalidCode
	}
	//校验code 过期
	if dbCode.ExpiresAt.Before(time.Now()) {
		return nil, errno.CodeExpired
	}
	// 校验 client_key, client_secret
	ok, err := s.dao.CheckClientInfo(ctx, dbCode.ClientId, req.ClientId, req.ClientSecret)
	if err != nil {
		return nil, errno.InvalidCode
	}
	if !ok {
		return nil, errno.InvalidClient
	}

	// 生成token
	access, refresh, _ := generateToken(req.ClientId, dbCode.UserId, true)
	expireAt := time.Now().Add(time.Hour * 24 * 7)

	err = s.dao.DB.Transaction(func(tx *gorm.DB) error {
		// 插入access_token
		accessToken := &model.OauthAccessToken{
			ClientId:  dbCode.Id,
			UserId:    dbCode.UserId,
			Token:     access,
			ExpiresAt: expireAt,
			Scope:     dbCode.Scope,
		}
		err := s.dao.TxInsertAccessToken(tx, accessToken)
		if err != nil {
			return err
		}
		// 插入refresh_token
		refreshToken := &model.OauthRefreshToken{
			ClientId:  dbCode.Id,
			UserId:    dbCode.UserId,
			Token:     refresh,
			ExpiresAt: expireAt,
			Scope:     dbCode.Scope,
		}
		err = s.dao.TxInsertRefreshToken(tx, refreshToken)
		if err != nil {
			return err
		}
		// 删除code
		if err = s.dao.DeleteCodeInfoByCode(ctx, req.Code); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	rsp := &model.AccessTokenRsp{
		AccessToken:  access,
		ExpiresIn:    24 * 7 * 60,
		RefreshToken: refresh,
	}
	return rsp, nil
}

func generateToken(clientKey string, userId int, isGenRefresh bool) (string, string, error) {
	var refreshToken string

	buf := bytes.NewBufferString(clientKey)
	buf.WriteString(strconv.Itoa(userId))
	buf.WriteString(time.Now().String())
	buf.WriteString(uuid.Must(uuid.NewRandom()).String())

	hash := md5.New()
	hash.Write(buf.Bytes())
	accessToken := hex.EncodeToString(hash.Sum(nil))

	if isGenRefresh {
		buf.WriteString(uuid.Must(uuid.NewRandom()).String())
		refreshHash := md5.New()
		refreshHash.Write(buf.Bytes())
		refreshToken = hex.EncodeToString(refreshHash.Sum(nil))
	}
	return accessToken, refreshToken, nil
}
