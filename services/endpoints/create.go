package endpoints

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Create creates a new endpoint.
func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Endpoint")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
