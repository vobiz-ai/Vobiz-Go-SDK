package campaign

// Get retrieves a campaign by ID.
func (s *Service) Get(campaignID string) (*GetResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "10dlc/Campaign/%s", campaignID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
