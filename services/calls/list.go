package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// List returns a paginated list of calls.
func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "Call")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// ListLive returns UUIDs of all live calls.
func (s *Service) ListLive(filters ...LiveCallFilters) (*LiveCallIDListResponse, error) {
	params := LiveCallFilters{Status: "live"}
	if len(filters) > 0 {
		params = filters[0]
		params.Status = "live"
	}
	req, err := s.r.NewRequest("GET", params, "Call")
	if err != nil {
		return nil, err
	}
	resp := &LiveCallIDListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// ListQueued returns UUIDs of all queued calls.
func (s *Service) ListQueued() (*QueuedCallIDListResponse, error) {
	req, err := s.r.NewRequest("GET", struct {
		Status string `url:"status"`
	}{"queued"}, "Call")
	if err != nil {
		return nil, err
	}
	resp := &QueuedCallIDListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
