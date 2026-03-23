package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Hangup terminates a live call.
func (s *Service) Hangup(callID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Call/%s", callID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// CancelRequest cancels a pending outbound call request.
func (s *Service) CancelRequest(requestID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Request/%s", requestID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
