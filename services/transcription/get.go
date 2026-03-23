package transcription

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves a transcription by ID and type.
func (s *Service) Get(transcriptionID, transcriptionType string) (*GetResponse, error) {
	req, err := s.r.NewRequest("GET", GetParams{Type: transcriptionType}, "Transcription/%s", transcriptionID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
