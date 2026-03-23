package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Play plays audio on a live call.
func (s *Service) Play(callID string, params PlayParams) (*PlayResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Call/%s/Play", callID)
	if err != nil {
		return nil, err
	}
	resp := &PlayResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopPlaying stops active audio playback.
func (s *Service) StopPlaying(callID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Call/%s/Play", callID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
