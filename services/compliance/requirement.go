package compliance

// Get retrieves a compliance requirement by ID.
func (s *RequirementService) Get(requirementID string) (*GetRequirementResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "ComplianceRequirement/%s", requirementID)
	if err != nil {
		return nil, err
	}
	resp := &GetRequirementResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List returns compliance requirements based on filters.
func (s *RequirementService) List(params ListRequirementParams) (*GetRequirementResponse, error) {
	req, err := s.r.NewRequest("GET", params, "ComplianceRequirement")
	if err != nil {
		return nil, err
	}
	resp := &GetRequirementResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
