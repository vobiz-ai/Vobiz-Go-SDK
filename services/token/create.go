package token

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// Create creates a JWT token.
func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	if params.IncomingAllow != nil || params.OutgoingAllow != nil {
		voicemap := make(map[string]interface{})
		if params.IncomingAllow != nil {
			voicemap["incoming_allow"] = params.IncomingAllow
		}
		if params.OutgoingAllow != nil {
			voicemap["outgoing_allow"] = params.OutgoingAllow
		}
		if len(voicemap) > 0 {
			params.Per = map[string]interface{}{"voice": voicemap}
		}
	}
	req, err := s.r.NewRequest("POST", params, "JWT/Token")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
