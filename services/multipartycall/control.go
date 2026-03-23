package multipartycall

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Start starts an MPC.
func (s *Service) Start(basic BasicParams) error {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("POST", map[string]string{"status": "active"}, "MultiPartyCall/%s", mpcID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// Stop stops an MPC.
func (s *Service) Stop(basic BasicParams) error {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("DELETE", nil, "MultiPartyCall/%s", mpcID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
