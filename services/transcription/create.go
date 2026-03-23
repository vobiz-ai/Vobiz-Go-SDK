package transcription

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Create creates a transcription for a recording.
func (s *Service) Create(params CreateParams) (map[string]interface{}, error) {
	body := struct {
		TranscriptionCallbackUrl string `json:"transcription_callback_url,omitempty"`
	}{params.TranscriptionCallbackUrl}
	req, err := s.r.NewRequest("POST", body, "Transcription/%s", params.RecordingID)
	if err != nil {
		return nil, err
	}
	resp := make(map[string]interface{})
	return resp, s.r.ExecuteRequest(req, &resp, internalhttp.IsVoiceRequest())
}
