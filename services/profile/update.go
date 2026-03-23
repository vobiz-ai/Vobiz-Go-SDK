package profile

// Update modifies a profile.
func (s *Service) Update(profileUUID string, params UpdateParams) (*GetResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Profile/%s", profileUUID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
