package cdr

// Get fetches CDR records with optional filters.
func (s *Service) Get(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "cdr")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List is an alias for Get.
func (s *Service) List(params ListParams) (*ListResponse, error) {
	return s.Get(params)
}
