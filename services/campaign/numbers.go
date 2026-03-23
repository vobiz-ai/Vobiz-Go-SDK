package campaign

// NumberLink links numbers to a campaign.
func (s *Service) NumberLink(campaignID string, params NumberLinkParams) (*NumberLinkUnlinkResponse, error) {
	req, err := s.r.NewRequest("POST", params, "10dlc/Campaign/%s/Number", campaignID)
	if err != nil {
		return nil, err
	}
	resp := &NumberLinkUnlinkResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// NumberGet retrieves a specific number from a campaign.
func (s *Service) NumberGet(campaignID, number string) (*NumberGetResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "10dlc/Campaign/%s/Number/%s", campaignID, number)
	if err != nil {
		return nil, err
	}
	resp := &NumberGetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// NumbersGet retrieves all numbers for a campaign.
func (s *Service) NumbersGet(campaignID string, params NumbersGetParams) (*NumberGetResponse, error) {
	req, err := s.r.NewRequest("GET", params, "10dlc/Campaign/%s/Number", campaignID)
	if err != nil {
		return nil, err
	}
	resp := &NumberGetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// NumberUnlink unlinks a number from a campaign.
func (s *Service) NumberUnlink(campaignID, number string, params NumberUnlinkParams) (*NumberLinkUnlinkResponse, error) {
	req, err := s.r.NewRequest("DELETE", params, "10dlc/Campaign/%s/Number/%s", campaignID, number)
	if err != nil {
		return nil, err
	}
	resp := &NumberLinkUnlinkResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
