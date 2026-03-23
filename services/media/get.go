package media

// Get retrieves a single media item by ID.
func (s *Service) Get(mediaID string) (*Media, error) {
	req, err := s.r.NewRequest("GET", nil, "Media/%s", mediaID)
	if err != nil {
		return nil, err
	}
	resp := &Media{}
	return resp, s.r.ExecuteRequest(req, resp)
}
