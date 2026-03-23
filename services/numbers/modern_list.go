package numbers

// ListModern lists numbers from modern /numbers endpoint.
func (s *Service) ListModern(params ModernListParams) (*ModernListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "numbers")
	resp := &ModernListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// ListInventory lists modern inventory numbers.
func (s *Service) ListInventory(params InventoryListParams) (*ModernListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "inventory/numbers")
	resp := &ModernListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
