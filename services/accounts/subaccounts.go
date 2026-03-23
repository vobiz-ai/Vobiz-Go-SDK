package accounts

// Create creates a new subaccount.
func (s *SubaccountService) Create(params SubaccountCreateParams) (*SubaccountCreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Subaccount")
	if err != nil {
		return nil, err
	}
	resp := &SubaccountCreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Get retrieves a subaccount by auth ID.
func (s *SubaccountService) Get(subAuthID string) (*Subaccount, error) {
	req, err := s.r.NewRequest("GET", nil, "Subaccount/%s", subAuthID)
	if err != nil {
		return nil, err
	}
	resp := &Subaccount{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List returns a paginated list of subaccounts.
func (s *SubaccountService) List(params SubaccountListParams) (*SubaccountList, error) {
	req, err := s.r.NewRequest("GET", params, "Subaccount")
	if err != nil {
		return nil, err
	}
	resp := &SubaccountList{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Update modifies a subaccount.
func (s *SubaccountService) Update(subAuthID string, params SubaccountUpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Subaccount/%s", subAuthID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Delete removes a subaccount.
func (s *SubaccountService) Delete(subAuthID string, params ...SubaccountDeleteParams) error {
	var p interface{}
	if len(params) > 0 {
		p = params[0]
	}
	req, err := s.r.NewRequest("DELETE", p, "Subaccount/%s", subAuthID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil)
}
