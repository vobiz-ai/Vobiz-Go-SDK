package verify

// List returns a paginated list of verify sessions.
func (s *Service) List(params ListParams) (*SessionList, error) {
	req, err := s.r.NewRequest("GET", params, "Verify/Session")
	if err != nil {
		return nil, err
	}
	resp := &SessionList{}
	return resp, s.r.ExecuteRequest(req, resp)
}
