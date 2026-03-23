package subaccounts

func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "sub-accounts")
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
