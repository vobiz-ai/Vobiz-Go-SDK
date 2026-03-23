package applications

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Update modifies an application.
func (s *Service) Update(appID string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Application/%s", appID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
