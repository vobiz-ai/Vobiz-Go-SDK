package endpoints

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves an endpoint by ID.
func (s *Service) Get(endpointID string) (*Endpoint, error) {
	req, err := s.r.NewRequest("GET", nil, "Endpoint/%s", endpointID)
	if err != nil {
		return nil, err
	}
	resp := &Endpoint{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
