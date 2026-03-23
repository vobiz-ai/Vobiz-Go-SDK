package multipartycall

import (
	"log"

	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/internal/utils"
)

// AddParticipant adds a participant to an MPC.
func (s *Service) AddParticipant(basic BasicParams, params AddParticipantParams) (*AddParticipantResponse, error) {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	if (params.From != "" || params.To != "") && params.CallUuid != "" {
		log.Fatal("cannot specify call_uuid when (from, to) is provided")
	}
	if params.From == "" && params.To == "" && params.CallUuid == "" {
		log.Fatal("specify either call_uuid or (from, to)")
	}
	if params.CallUuid == "" && (params.From == "" || params.To == "") {
		log.Fatal("specify (from, to) when not adding an existing call_uuid")
	}
	if params.CallerName == "" {
		params.CallerName = params.From
	}
	if len(params.CallerName) > 50 {
		log.Fatal("CallerName length must be in range [0,50]")
	}
	if params.RingTimeout == nil {
		params.RingTimeout = 45
	} else {
		utils.MultipleValidIntegers("RingTimeout", params.RingTimeout)
	}
	if params.DelayDial == nil {
		params.DelayDial = 0
	} else {
		utils.MultipleValidIntegers("DelayDial", params.DelayDial)
	}
	req, err := s.r.NewRequest("POST", params, "MultiPartyCall/%s/Participant", mpcID)
	if err != nil {
		return nil, err
	}
	resp := &AddParticipantResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// ListParticipants returns a paginated list of participants in an MPC.
func (s *Service) ListParticipants(basic BasicParams, params ListParticipantParams) (*ListParticipantsResponse, error) {
	mpcID := MakeMPCId(basic.MpcUuid, basic.FriendlyName)
	req, err := s.r.NewRequest("GET", params, "MultiPartyCall/%s/Participant", mpcID)
	if err != nil {
		return nil, err
	}
	resp := &ListParticipantsResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// GetParticipant retrieves a participant by ID.
func (s *Service) GetParticipant(pp ParticipantParams) (*Participant, error) {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("GET", nil, "MultiPartyCall/%s/Participant/%s", mpcID, pp.ParticipantId)
	if err != nil {
		return nil, err
	}
	resp := &Participant{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// UpdateParticipant updates a participant.
func (s *Service) UpdateParticipant(pp ParticipantParams, params UpdateParticipantParams) (*UpdateParticipantResponse, error) {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("POST", params, "MultiPartyCall/%s/Participant/%s", mpcID, pp.ParticipantId)
	if err != nil {
		return nil, err
	}
	resp := &UpdateParticipantResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// KickParticipant removes a participant from an MPC.
func (s *Service) KickParticipant(pp ParticipantParams) error {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("DELETE", nil, "MultiPartyCall/%s/Participant/%s", mpcID, pp.ParticipantId)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
