package calls

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Call represents a completed call record.
type Call struct {
	AnswerTime        string `json:"answer_time,omitempty"`
	BillDuration      int64  `json:"bill_duration,omitempty"`
	BilledDuration    int64  `json:"billed_duration,omitempty"`
	CallDirection     string `json:"call_direction,omitempty"`
	CallDuration      int64  `json:"call_duration,omitempty"`
	CallState         string `json:"call_state,omitempty"`
	CallUUID          string `json:"call_uuid,omitempty"`
	ConferenceUUID    string `json:"conference_uuid,omitempty"`
	EndTime           string `json:"end_time,omitempty"`
	FromNumber        string `json:"from_number,omitempty"`
	HangupCauseCode   int64  `json:"hangup_cause_code,omitempty"`
	HangupCauseName   string `json:"hangup_cause_name,omitempty"`
	HangupSource      string `json:"hangup_source,omitempty"`
	InitiationTime    string `json:"initiation_time,omitempty"`
	ParentCallUUID    string `json:"parent_call_uuid,omitempty"`
	ResourceURI       string `json:"resource_uri,omitempty"`
	ToNumber          string `json:"to_number,omitempty"`
	TotalAmount       string `json:"total_amount,omitempty"`
	TotalRate         string `json:"total_rate,omitempty"`
	StirVerification  string `json:"stir_verification,omitempty"`
	VoiceNetworkGroup string `json:"voice_network_group,omitempty"`
	StirAttestation   string `json:"stir_attestation,omitempty"`
	SourceIp          string `json:"source_ip,omitempty"`
	CnamLookup        string `json:"cnam_lookup,omitempty"`
}

// LiveCall represents an in-progress call.
type LiveCall struct {
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	AnswerURL        string `json:"answer_url,omitempty"`
	CallUUID         string `json:"call_uuid,omitempty"`
	CallerName       string `json:"caller_name,omitempty"`
	ParentCallUUID   string `json:"parent_call_uuid,omitempty"`
	SessionStart     string `json:"session_start,omitempty"`
	CallStatus       string `json:"call_status,omitempty"`
	StirVerification string `json:"stir_verification,omitempty"`
	StirAttestation  string `json:"stir_attestation,omitempty"`
}

// QueuedCall represents a queued (not yet connected) call.
type QueuedCall struct {
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	Status      string `json:"call_status,omitempty"`
	CallUUID    string `json:"call_uuid,omitempty"`
	CallerName  string `json:"caller_name,omitempty"`
	APIID       string `json:"api_id,omitempty"`
	Direction   string `json:"direction,omitempty"`
	RequestUUID string `json:"request_uuid,omitempty"`
}

// CreateParams are the parameters for creating an outbound call.
type CreateParams struct {
	From                   string `json:"from,omitempty"`
	To                     string `json:"to,omitempty"`
	AnswerURL              string `json:"answer_url,omitempty"`
	AnswerMethod           string `json:"answer_method,omitempty"`
	RingURL                string `json:"ring_url,omitempty"`
	RingMethod             string `json:"ring_method,omitempty"`
	HangupURL              string `json:"hangup_url,omitempty"`
	HangupMethod           string `json:"hangup_method,omitempty"`
	FallbackURL            string `json:"fallback_url,omitempty"`
	FallbackMethod         string `json:"fallback_method,omitempty"`
	CallerName             string `json:"caller_name,omitempty"`
	SendDigits             string `json:"send_digits,omitempty"`
	SendOnPreanswer        bool   `json:"send_on_preanswer,omitempty"`
	TimeLimit              int64  `json:"time_limit,omitempty"`
	HangupOnRing           int64  `json:"hangup_on_ring,omitempty"`
	MachineDetection       string `json:"machine_detection,omitempty"`
	MachineDetectionTime   int64  `json:"machine_detection_time,omitempty"`
	MachineDetectionUrl    string `json:"machine_detection_url,omitempty"`
	MachineDetectionMethod string `json:"machine_detection_method,omitempty"`
	SipHeaders             string `json:"sip_headers,omitempty"`
	RingTimeout            int64  `json:"ring_timeout,omitempty"`
}

// CreateResponse is returned when a call is successfully initiated.
type CreateResponse struct {
	Message     string      `json:"message"`
	ApiID       string      `json:"api_id"`
	RequestUUID interface{} `json:"request_uuid"`
}

// ListParams filters the call list endpoint.
type ListParams struct {
	Subaccount                 string `json:"subaccount,omitempty"          url:"subaccount,omitempty"`
	CallDirection              string `json:"call_direction,omitempty"      url:"call_direction,omitempty"`
	FromNumber                 string `json:"from_number,omitempty"         url:"from_number,omitempty"`
	ToNumber                   string `json:"to_number,omitempty"           url:"to_number,omitempty"`
	ParentCallUUID             string `json:"parent_call_uuid,omitempty"    url:"parent_call_uuid,omitempty"`
	EndTimeEquals              string `json:"end_time,omitempty"            url:"end_time,omitempty"`
	HangupCauseCode            int64  `json:"hangup_cause_code,omitempty"   url:"hangup_cause_code,omitempty"`
	HangupSource               string `json:"hangup_source,omitempty"       url:"hangup_source,omitempty"`
	EndTimeLessThan            string `json:"end_time__lt,omitempty"        url:"end_time__lt,omitempty"`
	EndTimeGreaterThan         string `json:"end_time__gt,omitempty"        url:"end_time__gt,omitempty"`
	EndTimeLessOrEqual         string `json:"end_time__lte,omitempty"       url:"end_time__lte,omitempty"`
	EndTimeGreaterOrEqual      string `json:"end_time__gte,omitempty"       url:"end_time__gte,omitempty"`
	BillDurationEquals         string `json:"bill_duration,omitempty"       url:"bill_duration,omitempty"`
	BillDurationLessThan       string `json:"bill_duration__lt,omitempty"   url:"bill_duration__lt,omitempty"`
	BillDurationGreaterThan    string `json:"bill_duration__gt,omitempty"   url:"bill_duration__gt,omitempty"`
	BillDurationLessOrEqual    string `json:"bill_duration__lte,omitempty"  url:"bill_duration__lte,omitempty"`
	BillDurationGreaterOrEqual string `json:"bill_duration__gte,omitempty"  url:"bill_duration__gte,omitempty"`
	Limit                      int64  `json:"limit,omitempty"               url:"limit,omitempty"`
	Offset                     int64  `json:"offset,omitempty"              url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of calls.
type ListResponse struct {
	ApiID   string        `json:"api_id"`
	Meta    *models.Meta  `json:"meta"`
	Objects []*Call       `json:"objects"`
}

// UpdateParams modifies a live call's routing.
type UpdateParams struct {
	Legs       string `json:"legs,omitempty"`
	AlegURL    string `json:"aleg_url,omitempty"`
	AlegMethod string `json:"aleg_method,omitempty"`
	BlegURL    string `json:"bleg_url,omitempty"`
	BlegMethod string `json:"bleg_method,omitempty"`
}

// UpdateResponse is returned after updating a call.
type UpdateResponse struct {
	ApiID   string `json:"api_id"`
	Message string `json:"message"`
}

// RecordParams configures call recording.
type RecordParams struct {
	TimeLimit           int64  `json:"time_limit,omitempty"`
	FileFormat          string `json:"file_format,omitempty"`
	TranscriptionType   string `json:"transcription_type,omitempty"`
	TranscriptionURL    string `json:"transcription_url,omitempty"`
	TranscriptionMethod string `json:"transcription_method,omitempty"`
	CallbackURL         string `json:"callback_url,omitempty"`
	CallbackMethod      string `json:"callback_method,omitempty"`
	RecordChannelType   string `json:"record_channel_type,omitempty"`
}

// RecordResponse is returned after starting a recording.
type RecordResponse struct {
	Message     string `json:"message,omitempty"`
	URL         string `json:"url,omitempty"`
	APIID       string `json:"api_id,omitempty"`
	RecordingID string `json:"recording_id,omitempty"`
}

// StreamParams configures a media stream.
type StreamParams struct {
	ServiceUrl           string `json:"service_url,omitempty"`
	Bidirectional        bool   `json:"bidirectional,omitempty"`
	AudioTrack           string `json:"audio_track,omitempty"`
	StreamTimeout        int64  `json:"stream_timeout,omitempty"`
	StatusCallbackUrl    string `json:"status_callback_url,omitempty"`
	StatusCallbackMethod string `json:"status_callback_method,omitempty"`
	ContentType          string `json:"content_type,omitempty"`
	ExtraHeaders         string `json:"extra_headers,omitempty"`
	CxBot                bool   `json:"cx_bot,omitempty"`
}

// StreamResponse is returned after starting a stream.
type StreamResponse struct {
	Message  string `json:"message,omitempty"`
	APIID    string `json:"api_id,omitempty"`
	StreamID string `json:"stream_id,omitempty"`
}

// StreamGetAllObject is a single stream record in a list response.
type StreamGetAllObject struct {
	AudioTrack          string `json:"audio_track"`
	Bidirectional       bool   `json:"bidirectional"`
	BilledAmount        string `json:"billed_amount"`
	BillDuration        int64  `json:"bill_duration"`
	CallUUID            string `json:"call_uuid"`
	CreatedAt           string `json:"created_at"`
	EndTime             string `json:"end_time"`
	VobizAuthId         string `json:"vobiz_auth_id"`
	ResourceURI         string `json:"resource_uri"`
	RoundedBillDuration string `json:"rounded_bill_duration"`
	ServiceURL          string `json:"service_url"`
	StartTime           string `json:"start_time"`
	Status              string `json:"status"`
	StatusCallbackURL   string `json:"status_callback_url"`
	StreamID            string `json:"stream_id"`
}

// StreamGetAll is the list response for streams on a call.
type StreamGetAll struct {
	ApiID   string               `json:"api_id,omitempty"`
	Meta    models.Meta          `json:"meta,omitempty"`
	Objects []StreamGetAllObject `json:"objects"`
}

// PlayParams configures audio playback on a live call.
type PlayParams struct {
	URLs   string `json:"urls"`
	Length string `json:"length,omitempty"`
	Legs   string `json:"legs,omitempty"`
	Loop   bool   `json:"loop,omitempty"`
	Mix    bool   `json:"mix,omitempty"`
}

// PlayResponse is returned after starting playback.
type PlayResponse struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// SpeakParams configures TTS on a live call.
type SpeakParams struct {
	Text           string `json:"text"`
	Voice          string `json:"length,omitempty"`
	Language       string `json:"language,omitempty"`
	Legs           string `json:"legs,omitempty"`
	Loop           bool   `json:"loop,omitempty"`
	Mix            bool   `json:"mix,omitempty"`
	Type           string `json:"type,omitempty"`
	CallbackURL    string `json:"callback_url,omitempty"`
	CallbackMethod string `json:"callback_method,omitempty"`
}

// SpeakResponse is returned after starting TTS.
type SpeakResponse struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// DTMFParams sends DTMF digits on a live call.
type DTMFParams struct {
	Digits string `json:"digits"`
	Legs   string `json:"legs,omitempty"`
}

// DTMFResponse is returned after sending DTMF.
type DTMFResponse struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// LiveCallFilters filters the live-call list.
type LiveCallFilters struct {
	CallDirection string `json:"call_direction,omitempty" url:"call_direction,omitempty"`
	FromNumber    string `json:"from_number,omitempty"    url:"from_number,omitempty"`
	ToNumber      string `json:"to_number,omitempty"      url:"to_number,omitempty"`
	Status        string `json:"status,omitempty"         url:"status,omitempty"`
}

// LiveCallIDListResponse holds a list of live call UUIDs.
type LiveCallIDListResponse struct {
	APIID string   `json:"api_id"`
	Calls []string `json:"calls"`
}

// QueuedCallIDListResponse holds a list of queued call UUIDs.
type QueuedCallIDListResponse struct {
	APIID string   `json:"api_id"`
	Calls []string `json:"calls"`
}
