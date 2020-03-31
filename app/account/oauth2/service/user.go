package service

import (
	"context"

	"chick/app/account/oauth2/model"
	"chick/errno"
	"chick/pkg/util"
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

	newUser := new(model.CloudUser)
	newUser.CopyFrom(req)
	s.dao.InsertCloudUser(ctx, newUser)
}
