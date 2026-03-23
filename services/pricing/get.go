package pricing

// Get retrieves pricing for a country by ISO code.
func (s *Service) Get(countryISO string) (*Pricing, error) {
	req, err := s.r.NewRequest("GET", struct {
		CountryISO string `url:"country_iso"`
	}{countryISO}, "Pricing")
	if err != nil {
		return nil, err
	}
	resp := &Pricing{}
	return resp, s.r.ExecuteRequest(req, resp)
}
