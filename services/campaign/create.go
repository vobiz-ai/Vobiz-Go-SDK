package campaign

// Create creates a new campaign.
func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "10dlc/Campaign")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Import imports an existing campaign.
func (s *Service) Import(params ImportParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "10dlc/Campaign/Import")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
