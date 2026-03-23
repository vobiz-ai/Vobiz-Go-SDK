package brand

// Get retrieves a brand by ID.
func (s *Service) Get(brandID string) (*GetResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "10dlc/Brand/%s", brandID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
