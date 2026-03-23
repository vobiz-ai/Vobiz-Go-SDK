package conferences

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Delete removes a conference.
func (s *Service) Delete(conferenceID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s", conferenceID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// DeleteAll removes all conferences.
func (s *Service) DeleteAll() error {
	req, err := s.r.NewRequest("DELETE", nil, "Conference")
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
