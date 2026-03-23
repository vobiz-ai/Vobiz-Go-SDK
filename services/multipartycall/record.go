package multipartycall

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// StartRecording starts recording an MPC.
func (s *Service) StartRecording(basic BasicParams, params StartRecordingParams) (*StartRecordingResponse, error) {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("POST", params, "MultiPartyCall/%s/Record", mpcID)
	if err != nil {
		return nil, err
	}
	resp := &StartRecordingResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopRecording stops recording an MPC.
func (s *Service) StopRecording(basic BasicParams) error {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("DELETE", nil, "MultiPartyCall/%s/Record", mpcID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// PauseRecording pauses recording an MPC.
func (s *Service) PauseRecording(basic BasicParams) error {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("POST", nil, "MultiPartyCall/%s/Record/Pause", mpcID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// ResumeRecording resumes recording an MPC.
func (s *Service) ResumeRecording(basic BasicParams) error {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("POST", nil, "MultiPartyCall/%s/Record/Resume", mpcID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
