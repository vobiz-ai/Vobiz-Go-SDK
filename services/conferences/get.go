package conferences

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves a conference by ID.
func (s *Service) Get(conferenceID string) (*Conference, error) {
	req, err := s.r.NewRequest("GET", nil, "Conference/%s", conferenceID)
	if err != nil {
		return nil, err
	}
	resp := &Conference{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// IDList returns a list of all conference IDs.
func (s *Service) IDList() (*IDListResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "Conference")
	if err != nil {
		return nil, err
	}
	resp := &IDListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
