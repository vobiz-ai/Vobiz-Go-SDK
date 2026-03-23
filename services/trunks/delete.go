package trunks

func (s *Service) Delete(trunkID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "trunks/"+trunkID)
	return s.r.ExecuteRequest(req, nil)
}
