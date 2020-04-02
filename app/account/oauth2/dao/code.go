package dao

//func (d *Dao) InsertCode(ctx context.Context, oauthCode *model.OauthAuthCode) error {
//	return d.DB.Select([]string{"client_id", "user_id", "code", "redirect_uri", "scope", "expires_at"}).Create(oauthCode).Error
//}

//func (d *Dao) GetCodeInfoByCode(ctx context.Context, code string) (*model.OauthAuthCode, error) {
//	dbCode := &model.OauthAuthCode{}
//	if err := d.DB.Select([]string{"id", "client_id", "user_id", "code", "expires_at", "scope"}).
//		Where("code = ? and is_deleted = ?", code, 0).First(dbCode).Error; err != nil {
//		return nil, err
//	}
//	return dbCode, nil
//}

//func (d *Dao) ValidateClient(ctx context.Context, clientKey, clientSecret string) (bool, error) {
//	clientInfo := &model.OauthClient{}
//	if err := d.DB.Select([]string{"key", "secret"}).Where("key = ?", clientKey).First(clientInfo).Error; err != nil {
//		return false, err
//	}
//	return clientKey == clientInfo.Key && clientSecret == clientInfo.Secret, nil
//}

//func (d *Dao) TxDeleteCodeInfoByCode(ctx context.Context, tx *gorm.DB, code string) error {
//	return tx.Table("oauth_auth_code").Where("code = ?", code).
//		Updates(map[string]int{"is_deleted": 1}).Error
//}
