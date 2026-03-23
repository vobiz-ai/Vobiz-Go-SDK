package multipartycall

import (
	"net/http"

	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
)

// StartParticipantRecording starts recording a participant.
func (s *Service) StartParticipantRecording(pp ParticipantParams, params StartRecordingParams) (*StartRecordingResponse, error) {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("POST", params, "MultiPartyCall/%s/Participant/%s/Record", mpcID, pp.ParticipantId)
	if err != nil {
		return nil, err
	}
	resp := &StartRecordingResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopParticipantRecording stops recording a participant.
func (s *Service) StopParticipantRecording(pp ParticipantParams, params ...ParticipantRecordingParams) error {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	var req *http.Request
	var err error
	if len(params) == 0 {
		req, err = s.r.NewRequest("DELETE", nil, "MultiPartyCall/%s/Participant/%s/Record", mpcID, pp.ParticipantId)
	} else {
		req, err = s.r.NewRequest("DELETE", params[0], "MultiPartyCall/%s/Participant/%s/Record", mpcID, pp.ParticipantId)
	}
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// PauseParticipantRecording pauses recording a participant.
func (s *Service) PauseParticipantRecording(pp ParticipantParams, params ...ParticipantRecordingParams) error {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	var req *http.Request
	var err error
	if len(params) == 0 {
		req, err = s.r.NewRequest("POST", nil, "MultiPartyCall/%s/Participant/%s/Record/Pause", mpcID, pp.ParticipantId)
	} else {
		req, err = s.r.NewRequest("POST", params[0], "MultiPartyCall/%s/Participant/%s/Record/Pause", mpcID, pp.ParticipantId)
	}
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// ResumeParticipantRecording resumes recording a participant.
func (s *Service) ResumeParticipantRecording(pp ParticipantParams, params ...ParticipantRecordingParams) error {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	var req *http.Request
	var err error
	if len(params) == 0 {
		req, err = s.r.NewRequest("POST", nil, "MultiPartyCall/%s/Participant/%s/Record/Resume", mpcID, pp.ParticipantId)
	} else {
		req, err = s.r.NewRequest("POST", params[0], "MultiPartyCall/%s/Participant/%s/Record/Resume", mpcID, pp.ParticipantId)
	}
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
