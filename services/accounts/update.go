package accounts

// Update modifies account details.
func (s *Service) Update(params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "")
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
