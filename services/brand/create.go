package brand

// Create creates a new brand.
func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "10dlc/Brand")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
