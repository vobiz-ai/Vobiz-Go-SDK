package numbers

// Update modifies a number.
func (s *Service) Update(numberID string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Number/%s", numberID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
