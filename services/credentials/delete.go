package credentials

func (s *Service) Delete(id string) error {
	req, err := s.r.NewRequest("DELETE", nil, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "credentials/"+id)
	return s.r.ExecuteRequest(req, nil)
}
