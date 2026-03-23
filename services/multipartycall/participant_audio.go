package multipartycall

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// StartPlayAudio plays audio to a participant.
func (s *Service) StartPlayAudio(pp ParticipantParams, params AudioParams) (*AudioResponse, error) {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("POST", params, "MultiPartyCall/%s/Member/%s/Play", mpcID, pp.ParticipantId)
	if err != nil {
		return nil, err
	}
	resp := &AudioResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopPlayAudio stops playing audio to a participant.
func (s *Service) StopPlayAudio(pp ParticipantParams) error {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("DELETE", nil, "MultiPartyCall/%s/Member/%s/Play", mpcID, pp.ParticipantId)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

// StartSpeak speaks text to a participant using TTS.
func (s *Service) StartSpeak(pp ParticipantParams, params SpeakParams) (*AudioResponse, error) {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("POST", params, "MultiPartyCall/%s/Member/%s/Speak", mpcID, pp.ParticipantId)
	if err != nil {
		return nil, err
	}
	resp := &AudioResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// StopSpeak stops speaking to a participant.
func (s *Service) StopSpeak(pp ParticipantParams) error {
	mpcID := MakeMPCId(pp.MpcUuid, pp.FriendlyName)
	req, err := s.r.NewRequest("DELETE", nil, "MultiPartyCall/%s/Member/%s/Speak", mpcID, pp.ParticipantId)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}
