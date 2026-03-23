package recordings

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves a recording by ID.
func (s *Service) Get(recordingID string) (*Recording, error) {
	req, err := s.r.NewRequest("GET", nil, "Recording/%s", recordingID)
	if err != nil {
		return nil, err
	}
	resp := &Recording{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
