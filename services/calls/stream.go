package calls

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Stream starts a media stream on a live call.
func (s *Service) Stream(callID string, params StreamParams) (*StreamResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Call/%s/Stream", callID)
	if err != nil {
		return nil, err
	}
	resp := &StreamResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopAllStreams stops all active streams on a call.
func (s *Service) StopAllStreams(callID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Call/%s/Stream", callID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// StopStream stops a specific stream.
func (s *Service) StopStream(callID, streamID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "Call/%s/Stream/%s", callID, streamID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// GetAllStreams lists all streams on a call.
func (s *Service) GetAllStreams(callID string) (*StreamGetAll, error) {
	req, err := s.r.NewRequest("GET", nil, "Call/%s/Stream", callID)
	if err != nil {
		return nil, err
	}
	resp := &StreamGetAll{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// GetStream retrieves a specific stream.
func (s *Service) GetStream(callID, streamID string) (*StreamGetAllObject, error) {
	req, err := s.r.NewRequest("GET", nil, "Call/%s/Stream/%s", callID, streamID)
	if err != nil {
		return nil, err
	}
	resp := &StreamGetAllObject{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
