package campaign

// Delete removes a campaign.
func (s *Service) Delete(campaignID string) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "10dlc/Campaign/%s", campaignID)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
