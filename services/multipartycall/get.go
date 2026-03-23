package multipartycall

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Get retrieves an MPC by UUID or friendly name.
func (s *Service) Get(basic BasicParams) (*GetResponse, error) {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("GET", nil, "MultiPartyCall/%s", mpcID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
