package numbers

// UnrentModern releases number via modern /numbers/{number} route.
func (s *Service) UnrentModern(number string) error {
	req, err := s.r.NewRequest("DELETE", nil, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "numbers/"+escNumber(number))
	return s.r.ExecuteRequest(req, nil)
}
