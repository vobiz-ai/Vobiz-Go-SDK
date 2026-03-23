package accounts

// GetProfile fetches account profile from /api/v1/auth/me.
func (s *Service) GetProfile() (map[string]interface{}, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	req.URL.Path = "/api/v1/auth/me"
	resp := map[string]interface{}{}
	return resp, s.r.ExecuteRequest(req, &resp)
}
