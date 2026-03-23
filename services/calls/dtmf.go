package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// SendDigits sends DTMF digits on a live call.
func (s *Service) SendDigits(callID string, params DTMFParams) (*DTMFResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Call/%s/DTMF", callID)
	if err != nil {
		return nil, err
	}
	resp := &DTMFResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
