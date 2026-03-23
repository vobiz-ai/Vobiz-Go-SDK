package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Update modifies routing for a live call.
func (s *Service) Update(callID string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Call/%s", callID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
