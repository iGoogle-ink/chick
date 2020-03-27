package service

import (
	"chick/app/web-svr/admin/model"
)

// Login login
func (s *Service) Login() (rsp *model.LoginRsp, err error) {

	user, err := s.dao.UserInfo(ctx, 1)
	if err != nil {
		return nil, err
	}
	rsp = new(model.LoginRsp)

	rsp.CopyFromUser(user)

	return rsp, nil
}
