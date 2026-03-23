package numbers

// Delete removes a number from the account.
func (s *Service) Delete(numberID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Number/%s", numberID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil)
}
