package subaccounts

func (s *Service) Delete(authID string, params ...DeleteParams) error {
	var p interface{}
	if len(params) > 0 { p = params[0] }
	req, err := s.r.NewRequest("DELETE", p, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "sub-accounts/"+authID)
	return s.r.ExecuteRequest(req, nil)
}
