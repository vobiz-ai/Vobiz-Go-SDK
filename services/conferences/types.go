package conferences

// Conference represents a conference object.
type Conference struct {
	ConferenceName        string   `json:"conference_name,omitempty"`
	ConferenceRunTime     string   `json:"conference_run_time,omitempty"`
	ConferenceMemberCount string   `json:"conference_member_count,omitempty"`
	Members               []Member `json:"members,omitempty"`
}

// Member represents a conference participant.
type Member struct {
	Muted      bool   `json:"muted,omitempty"`
	MemberID   string `json:"member_id,omitempty"`
	Deaf       bool   `json:"deaf,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	CallerName string `json:"caller_name,omitempty"`
	Direction  string `json:"direction,omitempty"`
	CallUUID   string `json:"call_uuid,omitempty"`
	JoinTime   string `json:"join_time,omitempty"`
}

// IDListResponse holds a list of conference IDs.
type IDListResponse struct {
	ApiID       string   `json:"api_id"`
	Conferences []string `json:"conferences"`
}

// RecordParams configures conference recording.
type RecordParams struct {
	TimeLimit           int64  `json:"time_limit,omitempty"`
	FileFormat          string `json:"file_format,omitempty"`
	TranscriptionType   string `json:"transcription_type,omitempty"`
	TranscriptionUrl    string `json:"transcription_url,omitempty"`
	TranscriptionMethod string `json:"transcription_method,omitempty"`
	CallbackUrl         string `json:"callback_url,omitempty"`
	CallbackMethod      string `json:"callback_method,omitempty"`
}

// RecordResponse is the response from starting a recording.
type RecordResponse struct {
	Message string `json:"message,omitempty"`
	Url     string `json:"url,omitempty"`
}

// MemberActionResponse is the generic member action response.
type MemberActionResponse struct {
	Message  string   `json:"message"`
	APIID    string   `json:"api_id"`
	MemberID []string `json:"member_id"`
}

// MemberSpeakParams configures TTS for a member.
type MemberSpeakParams struct {
	Text     string `json:"text"`
	Voice    string `json:"voice,omitempty"`
	Language string `json:"language,omitempty"`
}
