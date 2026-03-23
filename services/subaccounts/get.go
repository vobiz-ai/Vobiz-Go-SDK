package subaccounts

func (s *Service) Get(authID string) (*Subaccount, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "sub-accounts/"+authID)
	resp := &Subaccount{}
	return resp, s.r.ExecuteRequest(req, resp)
}
