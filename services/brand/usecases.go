package brand

// Usecases retrieves use cases for a brand.
func (s *Service) Usecases(brandID string) (*UsecaseResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "10dlc/Brand/%s/usecases", brandID)
	if err != nil {
		return nil, err
	}
	resp := &UsecaseResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
