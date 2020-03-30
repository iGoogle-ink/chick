package service

import "chick/errno"

// Authorize
func (s *Service) Authorize(userId int, clientKey, rspType, reUri, state string) (code string, err error) {
	// step1: find client info by client id
	if rspType != "code" {
		return "", errno.RequestErr
	}

	s.dao.AuthorizationCode(ctx, userId, clientKey, reUri)
}
