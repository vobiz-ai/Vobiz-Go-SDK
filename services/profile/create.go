package profile

// Create creates a new profile.
func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Profile")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
