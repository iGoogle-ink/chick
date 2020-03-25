package dao

import "chick/app-svr/app-admin/model"

func (d *Dao) UserInfo(id int) (user *model.MxUser, err error) {
	user = new(model.MxUser)
	if err = d.DB.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
