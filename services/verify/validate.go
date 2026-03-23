package verify

// Validate validates an OTP for a verify session.
func (s *Service) Validate(sessionUUID string, params ValidateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Verify/Session/%s", sessionUUID)
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
