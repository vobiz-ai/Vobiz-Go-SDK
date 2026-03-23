package brand

// Delete removes a brand.
func (s *Service) Delete(brandID string) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "10dlc/Brand/%s", brandID)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
