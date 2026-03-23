package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves a completed call by UUID.
func (s *Service) Get(callID string) (*Call, error) {
	req, err := s.r.NewRequest("GET", nil, "Call/%s", callID)
	if err != nil {
		return nil, err
	}
	resp := &Call{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// GetLive retrieves a specific live call.
func (s *Service) GetLive(callID string) (*LiveCall, error) {
	req, err := s.r.NewRequest("GET", struct {
		Status string `url:"status"`
	}{"live"}, "Call/%s", callID)
	if err != nil {
		return nil, err
	}
	resp := &LiveCall{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// GetQueued retrieves a specific queued call.
func (s *Service) GetQueued(callID string) (*QueuedCall, error) {
	req, err := s.r.NewRequest("GET", struct {
		Status string `url:"status"`
	}{"queued"}, "Call/%s", callID)
	if err != nil {
		return nil, err
	}
	resp := &QueuedCall{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
