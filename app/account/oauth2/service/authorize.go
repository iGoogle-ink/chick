package service

// Authorize
func (s *Service) Authorize(clientId, rspType, reUri, state string) {
	// step1: find client info by client id

	s.dao.GrantAuthorizationCode(ctx)
}
