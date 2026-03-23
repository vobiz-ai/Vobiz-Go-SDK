package profile

// Delete removes a profile.
func (s *Service) Delete(profileUUID string) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Profile/%s", profileUUID)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
