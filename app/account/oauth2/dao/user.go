package dao

import (
	"context"

	"chick/app/account/oauth2/model"
)

func (d *Dao) CloudUserInfo(ctx context.Context, uname string) (user *model.CloudUser, err error) {
	user = new(model.CloudUser)
	err = d.DB.Select([]string{"uname", "passwd", "phone", "mtime"}).
		Where("is_deleted = 0").
		Where("uname = ?", uname).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *Dao) InsertCloudUser(ctx context.Context, user *model.CloudUser) (id int, err error) {
	err = d.DB.Create(user).Error
	if err != nil {
		return -1, nil
	}
	return user.Id, nil
}
