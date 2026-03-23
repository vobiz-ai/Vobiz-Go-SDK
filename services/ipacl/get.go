package ipacl

func (s *Service) Get(id string) (*Entry, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "ip-acl/"+id)
	resp := &Entry{}
	return resp, s.r.ExecuteRequest(req, resp)
}
