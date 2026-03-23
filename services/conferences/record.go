package conferences

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Record starts recording a conference.
func (s *Service) Record(conferenceID string, params RecordParams) (*RecordResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Conference/%s/Record", conferenceID)
	if err != nil {
		return nil, err
	}
	resp := &RecordResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// RecordStop stops recording a conference.
func (s *Service) RecordStop(conferenceID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s/Record", conferenceID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
