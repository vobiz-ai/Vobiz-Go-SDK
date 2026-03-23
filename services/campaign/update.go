package campaign

// Update modifies a campaign.
func (s *Service) Update(campaignID string, params UpdateParams) (*GetResponse, error) {
	req, err := s.r.NewRequest("POST", params, "10dlc/Campaign/%s", campaignID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
