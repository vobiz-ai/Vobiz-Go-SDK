package media

// List returns a paginated list of media.
func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "Media")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
