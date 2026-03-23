package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Record starts recording a live call.
func (s *Service) Record(callID string, params RecordParams) (*RecordResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Call/%s/Record", callID)
	if err != nil {
		return nil, err
	}
	resp := &RecordResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopRecording stops an active recording.
func (s *Service) StopRecording(callID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Call/%s/Record", callID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
