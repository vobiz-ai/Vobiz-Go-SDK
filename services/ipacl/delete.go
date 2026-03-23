package ipacl

func (s *Service) Delete(id string) error {
	req, err := s.r.NewRequest("DELETE", nil, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "ip-acl/"+id)
	return s.r.ExecuteRequest(req, nil)
}
