package recordings

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Delete removes a recording.
func (s *Service) Delete(recordingID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Recording/%s", recordingID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
