package endpoints

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// List returns a paginated list of endpoints.
func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "Endpoint")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
