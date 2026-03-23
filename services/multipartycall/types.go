package multipartycall

// Meta holds pagination metadata for MPC lists.
type Meta struct {
	Previous   *string
	Next       *string
	TotalCount int64 `json:"count"`
	Offset     int64
	Limit      int64
}

// BasicParams identifies an MPC by UUID or friendly name.
type BasicParams struct {
	MpcUuid      string
	FriendlyName string
}

// ParticipantParams identifies an MPC participant.
type ParticipantParams struct {
	MpcUuid       string
	FriendlyName  string
	ParticipantId string
}

// AddParticipantParams holds parameters for adding a participant.
type AddParticipantParams struct {
	Role                           string      `json:"role,omitempty"`
	From                           string      `json:"from,omitempty"`
	To                             string      `json:"to,omitempty"`
	CallUuid                       string      `json:"call_uuid,omitempty"`
	CallerName                     string      `json:"caller_name,omitempty"`
	CallStatusCallbackUrl          string      `json:"call_status_callback_url,omitempty"`
	CallStatusCallbackMethod       string      `json:"call_status_callback_method,omitempty"`
	SipHeaders                     string      `json:"sip_headers,omitempty"`
	ConfirmKey                     string      `json:"confirm_key,omitempty"`
	ConfirmKeySoundUrl             string      `json:"confirm_key_sound_url,omitempty"`
	ConfirmKeySoundMethod          string      `json:"confirm_key_sound_method,omitempty"`
	DialMusic                      string      `json:"dial_music,omitempty"`
	RingTimeout                    interface{} `json:"ring_timeout,omitempty"`
	DelayDial                      interface{} `json:"delay_dial,omitempty"`
	MaxDuration                    int64       `json:"max_duration,omitempty"`
	MaxParticipants                int64       `json:"max_participants,omitempty"`
	WaitMusicUrl                   string      `json:"wait_music_url,omitempty"`
	WaitMusicMethod                string      `json:"wait_music_method,omitempty"`
	AgentHoldMusicUrl              string      `json:"agent_hold_music_url,omitempty"`
	AgentHoldMusicMethod           string      `json:"agent_hold_music_method,omitempty"`
	CustomerHoldMusicUrl           string      `json:"customer_hold_music_url,omitempty"`
	CustomerHoldMusicMethod        string      `json:"customer_hold_music_method,omitempty"`
	RecordingCallbackUrl           string      `json:"recording_callback_url,omitempty"`
	RecordingCallbackMethod        string      `json:"recording_callback_method,omitempty"`
	StatusCallbackUrl              string      `json:"status_callback_url,omitempty"`
	StatusCallbackMethod           string      `json:"status_callback_method,omitempty"`
	OnExitActionUrl                string      `json:"on_exit_action_url,omitempty"`
	OnExitActionMethod             string      `json:"on_exit_action_method,omitempty"`
	Record                         bool        `json:"record,omitempty"`
	RecordFileFormat               string      `json:"record_file_format,omitempty"`
	StatusCallbackEvents           string      `json:"status_callback_events,omitempty"`
	StayAlone                      bool        `json:"stay_alone,omitempty"`
	CoachMode                      bool        `json:"coach_mode,omitempty"`
	Mute                           bool        `json:"mute,omitempty"`
	Hold                           bool        `json:"hold,omitempty"`
	StartMpcOnEnter                *bool       `json:"start_mpc_on_enter,omitempty"`
	EndMpcOnExit                   bool        `json:"end_mpc_on_exit,omitempty"`
	RelayDtmfInputs                bool        `json:"relay_dtmf_inputs,omitempty"`
	EnterSound                     string      `json:"enter_sound,omitempty"`
	EnterSoundMethod               string      `json:"enter_sound_method,omitempty"`
	ExitSound                      string      `json:"exit_sound,omitempty"`
	ExitSoundMethod                string      `json:"exit_sound_method,omitempty"`
	StartRecordingAudio            string      `json:"start_recording_audio,omitempty"`
	StartRecordingAudioMethod      string      `json:"start_recording_audio_method,omitempty"`
	StopRecordingAudio             string      `json:"stop_recording_audio,omitempty"`
	StopRecordingAudioMethod       string      `json:"stop_recording_audio_method,omitempty"`
	RecordMinMemberCount           int64       `json:"record_min_member_count,omitempty"`
	AgentHoldMusic                 string      `json:"agent_hold_music,omitempty"`
	CustomerHoldMusic              string      `json:"customer_hold_music,omitempty"`
	CreateMpcWithSingleParticipant *bool       `json:"create_mpc_with_single_participant,omitempty"`
	SendDigits                     string      `json:"send_digits,omitempty"`
	SendOnPreanswer                bool        `json:"send_on_preanswer,omitempty"`
	TranscriptionUrl               string      `json:"transcription_url,omitempty"`
	Transcript                     bool        `json:"transcript,omitempty"`
	RecordParticipantTrack         bool        `json:"record_participant_track"`
}

// ListParams filters MPC list results.
type ListParams struct {
	SubAccount           string `url:"sub_account,omitempty"`
	FriendlyName         string `url:"friendly_name,omitempty"`
	Status               string `url:"status,omitempty"`
	TerminationCauseCode int64  `url:"termination_cause_code,omitempty"`
	EndTimeGt            string `url:"end_time__gt,omitempty"`
	EndTimeGte           string `url:"end_time__gte,omitempty"`
	EndTimeLt            string `url:"end_time__lt,omitempty"`
	EndTimeLte           string `url:"end_time__lte,omitempty"`
	CreationTimeGt       string `url:"creation_time__gt,omitempty"`
	CreationTimeGte      string `url:"creation_time__gte,omitempty"`
	CreationTimeLt       string `url:"creation_time__lt,omitempty"`
	CreationTimeLte      string `url:"creation_time__lte,omitempty"`
	Limit                int64  `url:"limit,omitempty"`
	Offset               int64  `url:"offset,omitempty"`
}

// ListParticipantParams filters participant list results.
type ListParticipantParams struct {
	CallUuid string `url:"call_uuid,omitempty"`
}

// UpdateParticipantParams holds parameters for updating a participant.
type UpdateParticipantParams struct {
	CoachMode *bool `json:"coach_mode,omitempty"`
	Mute      *bool `json:"mute,omitempty"`
	Hold      *bool `json:"hold,omitempty"`
}

// StartRecordingParams holds parameters for starting a recording.
type StartRecordingParams struct {
	FileFormat              string `json:"file_format,omitempty"`
	RecordingCallbackUrl    string `json:"recording_callback_url,omitempty"`
	RecordingCallbackMethod string `json:"recording_callback_method,omitempty"`
	TranscriptionUrl        string `json:"transcription_url,omitempty"`
	Transcript              bool   `json:"transcript,omitempty"`
	RecordTrackType         string `json:"record_track_type"`
}

// ParticipantRecordingParams holds optional recording params for participant recording.
type ParticipantRecordingParams struct {
	RecordTrackType string `json:"record_track_type"`
}

// GetResponse represents a single MPC.
type GetResponse struct {
	BilledAmount         string `json:"billed_amount,omitempty"`
	BilledDuration       int64  `json:"billed_duration,omitempty"`
	CreationTime         string `json:"creation_time,omitempty"`
	Duration             int64  `json:"duration,omitempty"`
	EndTime              string `json:"end_time,omitempty"`
	FriendlyName         string `json:"friendly_name,omitempty"`
	MpcUuid              string `json:"mpc_uuid,omitempty"`
	Participants         string `json:"participants,omitempty"`
	Recording            string `json:"recording,omitempty"`
	ResourceUri          string `json:"resource_uri,omitempty"`
	StartTime            string `json:"start_time,omitempty"`
	Status               string `json:"status,omitempty"`
	StayAlone            bool   `json:"stay_alone,omitempty"`
	SubAccount           string `json:"sub_account,omitempty"`
	TerminationCause     string `json:"termination_cause,omitempty"`
	TerminationCauseCode int64  `json:"termination_cause_code,omitempty"`
}

// ListResponse wraps a paginated list of MPCs.
type ListResponse struct {
	ApiID   string         `json:"api_id"`
	Meta    *Meta          `json:"meta"`
	Objects []*GetResponse `json:"objects"`
}

// AddParticipantResponse is the response from adding a participant.
type AddParticipantResponse struct {
	ApiID       string `json:"api_id"`
	Calls       []struct {
		To       string `json:"to,omitempty"`
		From     string `json:"from,omitempty"`
		CallUuid string `json:"call_uuid,omitempty"`
	} `json:"calls"`
	Message     string `json:"message,omitempty"`
	RequestUuid string `json:"request_uuid,omitempty"`
}

// StartRecordingResponse is the response from starting a recording.
type StartRecordingResponse struct {
	ApiID        string `json:"api_id"`
	Message      string `json:"message,omitempty"`
	RecordingId  string `json:"recording_id,omitempty"`
	RecordingUrl string `json:"recording_url,omitempty"`
}

// Participant represents an MPC participant.
type Participant struct {
	BilledAmount    string `json:"billed_amount,omitempty"`
	BilledDuration  int64  `json:"billed_duration,omitempty"`
	CallUuid        string `json:"call_uuid,omitempty"`
	CoachMode       bool   `json:"coach_mode,omitempty"`
	Duration        int64  `json:"duration,omitempty"`
	EndMpcOnExit    bool   `json:"end_mpc_on_exit,omitempty"`
	ExitCause       string `json:"exit_cause,omitempty"`
	ExitTime        string `json:"exit_time,omitempty"`
	Hold            bool   `json:"hold,omitempty"`
	JoinTime        string `json:"join_time,omitempty"`
	MemberAddress   string `json:"member_address,omitempty"`
	MemberId        string `json:"member_id,omitempty"`
	MpcUuid         string `json:"mpc_uuid,omitempty"`
	Mute            bool   `json:"mute,omitempty"`
	ResourceUri     string `json:"resource_uri,omitempty"`
	Role            string `json:"role,omitempty"`
	StartMpcOnEnter bool   `json:"start_mpc_on_enter,omitempty"`
}

// ListParticipantsResponse wraps a paginated list of participants.
type ListParticipantsResponse struct {
	ApiID   string         `json:"api_id"`
	Meta    *Meta          `json:"meta"`
	Objects []*Participant `json:"objects"`
}

// UpdateParticipantResponse is the response from updating a participant.
type UpdateParticipantResponse struct {
	ApiID     string `json:"api_id"`
	CoachMode string `json:"coach_mode,omitempty"`
	Mute      string `json:"mute,omitempty"`
	Hold      string `json:"hold,omitempty"`
}

// AudioParams holds a URL for audio playback.
type AudioParams struct {
	Url string `json:"url"`
}

// AudioResponse is the response from audio play/speak operations.
type AudioResponse struct {
	APIID        string   `json:"api_id"`
	Message      string   `json:"message"`
	MemberId     []string `json:"mpcMemberId,omitempty"`
	FriendlyName string   `json:"mpcName,omitempty"`
}

// SpeakParams holds TTS parameters.
type SpeakParams struct {
	Text           string `json:"text"`
	Voice          string `json:"voice"`
	Language       string `json:"language"`
	Mix            bool   `json:"mix"`
	Type           string `json:"type,omitempty"`
	CallbackURL    string `json:"callback_url,omitempty"`
	CallbackMethod string `json:"callback_method,omitempty"`
}
