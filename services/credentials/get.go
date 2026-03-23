package credentials

func (s *Service) Get(id string) (*Credential, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "credentials/"+id)
	resp := &Credential{}
	return resp, s.r.ExecuteRequest(req, resp)
}
