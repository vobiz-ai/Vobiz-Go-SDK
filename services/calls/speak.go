package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Speak plays TTS audio on a live call.
func (s *Service) Speak(callID string, params SpeakParams) (*SpeakResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Call/%s/Speak", callID)
	if err != nil {
		return nil, err
	}
	resp := &SpeakResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopSpeaking stops active TTS on a call.
func (s *Service) StopSpeaking(callID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Call/%s/Speak", callID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
