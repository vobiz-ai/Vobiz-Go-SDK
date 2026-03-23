package applications

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves an application by ID.
func (s *Service) Get(appID string) (*Application, error) {
	req, err := s.r.NewRequest("GET", nil, "Application/%s", appID)
	if err != nil {
		return nil, err
	}
	resp := &Application{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
