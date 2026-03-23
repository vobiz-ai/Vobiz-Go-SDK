package accounts

func (s *Service) GetBalance(currency string) (*BalanceResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernAccountPath(req, s.r.AuthID(), "balance/"+currency)
	resp := &BalanceResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
