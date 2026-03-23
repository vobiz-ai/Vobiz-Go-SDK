package verify

// Get retrieves a verify session by UUID.
func (s *Service) Get(sessionUUID string) (*Session, error) {
	req, err := s.r.NewRequest("GET", nil, "Verify/Session/%s", sessionUUID)
	if err != nil {
		return nil, err
	}
	resp := &Session{}
	return resp, s.r.ExecuteRequest(req, resp)
}
