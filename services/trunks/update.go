package trunks

func (s *Service) Update(trunkID string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("PUT", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "trunks/"+trunkID)
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
