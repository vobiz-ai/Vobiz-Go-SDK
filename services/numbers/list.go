package numbers

// List returns a paginated list of numbers.
func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "Number")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
