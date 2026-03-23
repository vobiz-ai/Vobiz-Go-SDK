package numbers

// Get retrieves a number by ID.
func (s *Service) Get(numberID string) (*Number, error) {
	req, err := s.r.NewRequest("GET", nil, "Number/%s", numberID)
	if err != nil {
		return nil, err
	}
	resp := &Number{}
	return resp, s.r.ExecuteRequest(req, resp)
}
