package originationuris

func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "origination-uris")
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
