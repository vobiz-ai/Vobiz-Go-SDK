package profile

// Get retrieves a profile by UUID.
func (s *Service) Get(profileUUID string) (*GetResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "Profile/%s", profileUUID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
