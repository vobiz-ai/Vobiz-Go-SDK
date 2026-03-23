package originationuris

func (s *Service) Get(id string) (*URI, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "origination-uris/"+id)
	resp := &URI{}
	return resp, s.r.ExecuteRequest(req, resp)
}
