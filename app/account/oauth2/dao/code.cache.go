package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"chick/app/account/oauth2/model"
)

func codeCacheKey(code string) string {
	return fmt.Sprintf("AUTH_CODE_%s", code)
}

func (d *Dao) CacheAuthCode(ctx context.Context, code string) (codeInfo *model.CacheAuthCode, err error) {
	key := codeCacheKey(code)
	bs, err := d.redis.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	codeInfo = new(model.CacheAuthCode)
	if err = json.Unmarshal(bs, codeInfo); err != nil {
		return codeInfo, nil
	}
	return nil, err
}

func (d *Dao) AddCacheAuthCode(ctx context.Context, code *model.CacheAuthCode) error {
	key := codeCacheKey(code.Code)
	codeBs, _ := json.Marshal(code)
	return d.redis.Set(key, codeBs, time.Duration(code.Expires)*time.Second).Err()
}
