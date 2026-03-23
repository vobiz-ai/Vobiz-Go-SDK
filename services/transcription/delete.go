package transcription

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Delete removes a transcription.
func (s *Service) Delete(transcriptionID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Transcription/%s", transcriptionID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
