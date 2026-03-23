package accounts

// Get retrieves the current account details.
func (s *Service) Get() (*Account, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil {
		return nil, err
	}
	resp := &Account{}
	return resp, s.r.ExecuteRequest(req, resp)
}
