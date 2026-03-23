package numbers

// Create buys a phone number from the catalog.
func (s *PhoneNumberService) Create(number string, params PhoneNumberCreateParams) (*PhoneNumberCreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "PhoneNumber/%s", number)
	if err != nil {
		return nil, err
	}
	resp := &PhoneNumberCreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List returns a paginated list of available phone numbers from the catalog.
func (s *PhoneNumberService) List(params PhoneNumberListParams) (*PhoneNumberListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "PhoneNumber")
	if err != nil {
		return nil, err
	}
	resp := &PhoneNumberListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
