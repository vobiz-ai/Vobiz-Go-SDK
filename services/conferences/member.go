package conferences

import internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"

// MemberHangup hangs up a conference member.
func (s *Service) MemberHangup(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s/Member/%s", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberKick kicks a member from the conference.
func (s *Service) MemberKick(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("POST", nil, "Conference/%s/Member/%s/Kick", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberMute mutes a conference member.
func (s *Service) MemberMute(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("POST", nil, "Conference/%s/Member/%s/Mute", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberUnmute unmutes a conference member.
func (s *Service) MemberUnmute(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s/Member/%s/Mute", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberDeaf deafens a conference member.
func (s *Service) MemberDeaf(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("POST", nil, "Conference/%s/Member/%s/Deaf", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberUndeaf undeafens a conference member.
func (s *Service) MemberUndeaf(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s/Member/%s/Deaf", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberPlay plays audio to a conference member.
func (s *Service) MemberPlay(conferenceID, memberID, url string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("POST", struct {
		Url string `json:"url"`
	}{url}, "Conference/%s/Member/%s/Play", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberPlayStop stops playing audio to a conference member.
func (s *Service) MemberPlayStop(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s/Member/%s/Play", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberSpeak speaks text to a conference member using TTS.
func (s *Service) MemberSpeak(conferenceID, memberID string, params MemberSpeakParams) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Conference/%s/Member/%s/Speak", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

// MemberSpeakStop stops speaking to a conference member.
func (s *Service) MemberSpeakStop(conferenceID, memberID string) (*MemberActionResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Conference/%s/Member/%s/Speak", conferenceID, memberID)
	if err != nil {
		return nil, err
	}
	resp := &MemberActionResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
