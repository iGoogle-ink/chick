package service

import (
	"chick/app-svr/app-admin/model"
)

// Login login
func (s *Service) Login() (rsp *model.LoginRsp, err error) {

	user, err := s.dao.UserInfo(ctx, 1)
	if err != nil {
		return nil, err
	}
	rsp = &model.LoginRsp{
		Token: user.Name,
	}
	return rsp, nil
}
