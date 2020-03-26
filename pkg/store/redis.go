package store

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/utils/uuid"
)

//var (
//	_ oauth2.TokenStore = &TokenStore{}
//)

func NewRedisClusterStoreWithCli(cli *redis.ClusterClient, prefix ...string) *TokenStore {
	store := &TokenStore{
		cli: cli,
	}
	if len(prefix) > 0 {
		store.prefix = prefix[0]
	}
	return store
}

type TokenStore struct {
	cli    *redis.ClusterClient
	prefix string
}

func (s *TokenStore) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", s.prefix, key)
}

func (s *TokenStore) checkError(result redis.Cmder) (bool, error) {
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

// remove
func (s *TokenStore) remove(key string) error {
	result := s.cli.Del(s.wrapperKey(key))
	_, err := s.checkError(result)
	return err
}

func (s *TokenStore) removeToken(tokenString string, isRefresh bool) error {
	basicID, err := s.getBasicID(tokenString)
	if err != nil {
		return err
	} else if basicID == "" {
		return nil
	}

	err = s.remove(tokenString)
	if err != nil {
		return err
	}

	token, err := s.getToken(basicID)
	if err != nil {
		return err
	} else if token == nil {
		return nil
	}

	checkToken := token.GetRefresh()
	if isRefresh {
		checkToken = token.GetAccess()
	}
	iresult := s.cli.Exists(s.wrapperKey(checkToken))
	if err := iresult.Err(); err != nil && err != redis.Nil {
		return err
	} else if iresult.Val() == 0 {
		return s.remove(basicID)
	}

	return nil
}

func (s *TokenStore) parseToken(result *redis.StringCmd) (oauth2.TokenInfo, error) {
	if ok, err := s.checkError(result); err != nil {
		return nil, err
	} else if ok {
		return nil, nil
	}

	buf, err := result.Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var token models.Token
	if err := json.Unmarshal(buf, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *TokenStore) getToken(key string) (oauth2.TokenInfo, error) {
	result := s.cli.Get(s.wrapperKey(key))
	return s.parseToken(result)
}

func (s *TokenStore) parseBasicID(result *redis.StringCmd) (string, error) {
	if ok, err := s.checkError(result); err != nil {
		return "", err
	} else if ok {
		return "", nil
	}
	return result.Val(), nil
}

func (s *TokenStore) getBasicID(token string) (string, error) {
	result := s.cli.Get(s.wrapperKey(token))
	return s.parseBasicID(result)
}

// Create Create and store the new token information
func (s *TokenStore) Create(info oauth2.TokenInfo) error {
	ct := time.Now()
	jv, err := json.Marshal(info)
	if err != nil {
		return err
	}

	pipe := s.cli.TxPipeline()
	if code := info.GetCode(); code != "" {
		pipe.Set(s.wrapperKey(code), jv, info.GetCodeExpiresIn())
	} else {
		basicID := uuid.Must(uuid.NewRandom()).String()
		aexp := info.GetAccessExpiresIn()
		rexp := aexp

		if refresh := info.GetRefresh(); refresh != "" {
			rexp = info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()).Sub(ct)
			if aexp.Seconds() > rexp.Seconds() {
				aexp = rexp
			}
			pipe.Set(s.wrapperKey(refresh), basicID, rexp)
		}

		pipe.Set(s.wrapperKey(info.GetAccess()), basicID, aexp)
		pipe.Set(s.wrapperKey(basicID), jv, rexp)
	}

	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

// RemoveByCode Use the authorization code to delete the token information
func (s *TokenStore) RemoveByCode(code string) error {
	return s.remove(code)
}

// RemoveByAccess Use the access token to delete the token information
func (s *TokenStore) RemoveByAccess(access string) error {
	return s.removeToken(access, false)
}

// RemoveByRefresh Use the refresh token to delete the token information
func (s *TokenStore) RemoveByRefresh(refresh string) error {
	return s.removeToken(refresh, false)
}

// GetByCode Use the authorization code for token information data
func (s *TokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	return s.getToken(code)
}

// GetByAccess Use the access token for token information data
func (s *TokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	basicID, err := s.getBasicID(access)
	if err != nil || basicID == "" {
		return nil, err
	}
	return s.getToken(basicID)
}

// GetByRefresh Use the refresh token for token information data
func (s *TokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	basicID, err := s.getBasicID(refresh)
	if err != nil || basicID == "" {
		return nil, err
	}
	return s.getToken(basicID)
}
