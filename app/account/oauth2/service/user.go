package service

import (
	"context"
	"log"

	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/util"
	"github.com/jinzhu/gorm"
)

func (s *Service) Login(ctx context.Context, req *model.LoginReq) (session string, err error) {

	user, err := s.dao.CloudUserInfo(ctx, req.Uname)
	if err != nil {
		return "", errno.New(-1, "账号不存在")
	}
	if req.Uname != user.Uname {
		return "", errno.New(-2, "密码错误")
	}
	return util.GetRandomString(32), nil
}

func (s *Service) Register(ctx context.Context, req *model.RegisterReq) (err error) {
	// 从数据库中查找，看是否已经存在此用户名
	user, err := s.dao.CloudUserInfo(ctx, req.Uname)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("s.dao.CloudUserInfo(%s),error:%+v.\n", req.Uname, err)
		return errno.New(-1, "创建失败")
	}
	if user != nil {
		return errno.New(-2, "用户已存在")
	}
	// 创建新用户
	newUser := new(model.CloudUser)
	newUser.CopyFrom(req)
	_, err = s.dao.InsertCloudUser(ctx, newUser)
	if err != nil {
		log.Printf("s.dao.InsertCloudUser(%v),error:%+v.\n", newUser, err)
		return errno.New(-3, "创建失败")
	}
	return nil
}
