package originationuris

func (s *Service) Update(id string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("PUT", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "origination-uris/"+id)
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
