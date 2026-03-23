package accounts

func (s *Service) GetConcurrency() (*ConcurrencyResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernAccountPath(req, s.r.AuthID(), "concurrency")
	resp := &ConcurrencyResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
