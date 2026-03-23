package trunks

func (s *Service) Get(trunkID string) (*Trunk, error) {
	req, err := s.r.NewRequest("GET", nil, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "trunks/"+trunkID)
	resp := &Trunk{}
	return resp, s.r.ExecuteRequest(req, resp)
}
