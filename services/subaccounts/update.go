package subaccounts

func (s *Service) Update(authID string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("PUT", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "sub-accounts/"+authID)
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
