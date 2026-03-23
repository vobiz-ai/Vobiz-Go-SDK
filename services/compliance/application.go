package compliance

// Create creates a new compliance application.
func (s *ApplicationService) Create(params CreateApplicationParams) (*ApplicationResponse, error) {
	req, err := s.r.NewRequest("POST", params, "ComplianceApplication")
	if err != nil {
		return nil, err
	}
	resp := &ApplicationResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Get retrieves a compliance application by ID.
func (s *ApplicationService) Get(applicationID string) (*ApplicationResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "ComplianceApplication/%s", applicationID)
	if err != nil {
		return nil, err
	}
	resp := &ApplicationResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List returns a paginated list of compliance applications.
func (s *ApplicationService) List(params ApplicationListParams) (*ListApplicationResponse, error) {
	req, err := s.r.NewRequest("GET", params, "ComplianceApplication")
	if err != nil {
		return nil, err
	}
	resp := &ListApplicationResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Update modifies a compliance application.
func (s *ApplicationService) Update(params UpdateApplicationParams) (*UpdateApplicationResponse, error) {
	req, err := s.r.NewRequest("POST", params, "ComplianceApplication/%s", params.ComplianceApplicationId)
	if err != nil {
		return nil, err
	}
	resp := &UpdateApplicationResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Delete removes a compliance application.
func (s *ApplicationService) Delete(applicationID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "ComplianceApplication/%s", applicationID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil)
}

// Submit submits a compliance application for review.
func (s *ApplicationService) Submit(applicationID string) (*ApplicationResponse, error) {
	req, err := s.r.NewRequest("POST", nil, "ComplianceApplication/%s/Submit", applicationID)
	if err != nil {
		return nil, err
	}
	resp := &ApplicationResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
