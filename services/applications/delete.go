package applications

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Delete removes an application.
func (s *Service) Delete(appID string, data ...DeleteParams) error {
	var optionalParams interface{}
	if len(data) > 0 {
		optionalParams = data[0]
	}
	req, err := s.r.NewRequest("DELETE", optionalParams, "Application/%s", appID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
