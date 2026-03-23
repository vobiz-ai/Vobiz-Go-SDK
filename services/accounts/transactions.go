package accounts

func (s *Service) GetTransactions(params TransactionListParams) (*TransactionListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "")
	if err != nil { return nil, err }
	setModernAccountPath(req, s.r.AuthID(), "transactions")
	resp := &TransactionListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
