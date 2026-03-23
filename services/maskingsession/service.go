// Package maskingsession provides the Vobiz Masking Session API service.
package maskingsession

import (
	"github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"
	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/models"
)

// VoiceInteraction holds details of a voice interaction in a masking session.
type VoiceInteraction struct {
	StartTime              string  `json:"start_time,omitempty"`
	EndTime                string  `json:"end_time,omitempty"`
	FirstPartyResourceUrl  string  `json:"first_party_resource_url,omitempty"`
	SecondPartyResourceUrl string  `json:"second_party_resource_url,omitempty"`
	FirstPartyStatus       string  `json:"first_party_status,omitempty"`
	SecondPartyStatus      string  `json:"second_party_status,omitempty"`
	Type                   string  `json:"type,omitempty"`
	TotalCallAmount        float64 `json:"total_call_amount,omitempty"`
	CallBilledDuration     int     `json:"call_billed_duration,omitempty"`
	RecordingResourceUrl   string  `json:"recording_resource_url,omitempty"`
	AuthID                 string  `json:"auth_id,omitempty"`
	TotalCallCount         int     `json:"total_call_count"`
	Duration               float64 `json:"duration"`
}

// Session represents a masking session.
type Session struct {
	FirstParty                   string             `json:"first_party,omitempty"`
	SecondParty                  string             `json:"second_party,omitempty"`
	VirtualNumber                string             `json:"virtual_number,omitempty"`
	Status                       string             `json:"status,omitempty"`
	InitiateCallToFirstParty     bool               `json:"initiate_call_to_first_party,omitempty"`
	SessionUUID                  string             `json:"session_uuid,omitempty"`
	CallbackUrl                  string             `json:"callback_url,omitempty"`
	CallbackMethod               string             `json:"callback_method,omitempty"`
	CreatedAt                    string             `json:"created_time,omitempty"`
	UpdatedAt                    string             `json:"modified_time,omitempty"`
	ExpiryAt                     string             `json:"expiry_time,omitempty"`
	Duration                     int64              `json:"duration,omitempty"`
	SessionCreationAmount        int64              `json:"amount"`
	CallTimeLimit                int64              `json:"call_time_limit,omitempty"`
	RingTimeout                  int64              `json:"ring_timeout,omitempty"`
	FirstPartyPlayUrl            string             `json:"first_party_play_url,omitempty"`
	SecondPartyPlayUrl           string             `json:"second_party_play_url,omitempty"`
	Record                       bool               `json:"record,omitempty"`
	RecordFileFormat             string             `json:"record_file_format,omitempty"`
	RecordingCallbackUrl         string             `json:"recording_callback_url,omitempty"`
	RecordingCallbackMethod      string             `json:"recording_callback_method,omitempty"`
	Interaction                  []VoiceInteraction `json:"interaction"`
	TotalCallAmount              float64            `json:"total_call_amount"`
	TotalCallCount               int                `json:"total_call_count"`
	TotalCallBilledDuration      int                `json:"total_call_billed_duration"`
	TotalSessionAmount           float64            `json:"total_session_amount"`
	LastInteractionTime          string             `json:"last_interaction_time"`
	IsPinAuthenticationRequired  bool               `json:"is_pin_authentication_required"`
	GeneratePin                  bool               `json:"generate_pin"`
	GeneratePinLength            int64              `json:"generate_pin_length"`
	FirstPartyPin                string             `json:"first_party_pin"`
	SecondPartyPin               string             `json:"second_party_pin"`
	PinPromptPlay                string             `json:"pin_prompt_play"`
	PinRetry                     int64              `json:"pin_retry"`
	PinRetryWait                 int64              `json:"pin_retry_wait"`
	IncorrectPinPlay             string             `json:"incorrect_pin_play"`
	UnknownCallerPlay            string             `json:"unknown_caller_play"`
	VirtualNumberCooloffPeriod   int                `json:"virtual_number_cooloff_period,omitempty"`
	ForcePinAuthentication       bool               `json:"force_pin_authentication"`
	CreateSessionWithSingleParty bool               `json:"create_session_with_single_Party"`
}

// CreateParams holds parameters for creating a masking session.
type CreateParams struct {
	FirstParty                   string `json:"first_party,omitempty"`
	SecondParty                  string `json:"second_party,omitempty"`
	SessionExpiry                int    `json:"session_expiry"`
	CallTimeLimit                int    `json:"call_time_limit,omitempty"`
	Record                       bool   `json:"record,omitempty"`
	RecordFileFormat             string `json:"record_file_format,omitempty"`
	RecordingCallbackUrl         string `json:"recording_callback_url,omitempty"`
	InitiateCallToFirstParty     bool   `json:"initiate_call_to_first_party,omitempty"`
	CallbackUrl                  string `json:"callback_url,omitempty"`
	CallbackMethod               string `json:"callback_method,omitempty"`
	RingTimeout                  int64  `json:"ring_timeout,omitempty"`
	FirstPartyPlayUrl            string `json:"first_party_play_url,omitempty"`
	SecondPartyPlayUrl           string `json:"second_party_play_url,omitempty"`
	RecordingCallbackMethod      string `json:"recording_callback_method,omitempty"`
	IsPinAuthenticationRequired  bool   `json:"is_pin_authentication_required,omitempty"`
	GeneratePin                  bool   `json:"generate_pin,omitempty"`
	GeneratePinLength            int64  `json:"generate_pin_length,omitempty"`
	FirstPartyPin                string `json:"first_party_pin,omitempty"`
	SecondPartyPin               string `json:"second_party_pin,omitempty"`
	PinPromptPlay                string `json:"pin_prompt_play,omitempty"`
	PinRetry                     int64  `json:"pin_retry,omitempty"`
	PinRetryWait                 int64  `json:"pin_retry_wait,omitempty"`
	IncorrectPinPlay             string `json:"incorrect_pin_play,omitempty"`
	UnknownCallerPlay            string `json:"unknown_caller_play,omitempty"`
	SubAccount                   string `json:"subaccount,omitempty"`
	GeoMatch                     *bool  `json:"geomatch,omitempty"`
	VirtualNumberCooloffPeriod   int    `json:"virtual_number_cooloff_period,omitempty"`
	ForcePinAuthentication       bool   `json:"force_pin_authentication,omitempty"`
	CreateSessionWithSingleParty bool   `json:"create_session_with_single_Party,omitempty"`
}

// UpdateParams holds parameters for updating a masking session.
type UpdateParams struct {
	FirstParty                   string `json:"first_party,omitempty"`
	SecondParty                  string `json:"second_party,omitempty"`
	SessionExpiry                int64  `json:"session_expiry,omitempty"`
	CallTimeLimit                int64  `json:"call_time_limit,omitempty"`
	Record                       bool   `json:"record,omitempty"`
	RecordFileFormat             string `json:"record_file_format,omitempty"`
	RecordingCallbackUrl         string `json:"recording_callback_url,omitempty"`
	CallbackUrl                  string `json:"callback_url,omitempty"`
	CallbackMethod               string `json:"callback_method,omitempty"`
	RingTimeout                  int64  `json:"ring_timeout,omitempty"`
	FirstPartyPlayUrl            string `json:"first_party_play_url,omitempty"`
	SecondPartyPlayUrl           string `json:"second_party_play_url,omitempty"`
	RecordingCallbackMethod      string `json:"recording_callback_method,omitempty"`
	SubAccount                   string `json:"subaccount,omitempty"`
	GeoMatch                     *bool  `json:"geomatch,omitempty"`
	CreateSessionWithSingleParty bool   `json:"create_session_with_single_Party,omitempty"`
}

// ListParams filters masking sessions.
type ListParams struct {
	FirstParty                string `url:"first_party,omitempty"`
	SecondParty               string `url:"second_party,omitempty"`
	VirtualNumber             string `url:"virtual_number,omitempty"`
	Status                    string `url:"status,omitempty"`
	CreatedTimeEquals         string `url:"created_time,omitempty"`
	CreatedTimeLessThan       string `url:"created_time__lt,omitempty"`
	CreatedTimeGreaterThan    string `url:"created_time__gt,omitempty"`
	CreatedTimeLessOrEqual    string `url:"created_time__lte,omitempty"`
	CreatedTimeGreaterOrEqual string `url:"created_time__gte,omitempty"`
	ExpiryTimeEquals          string `url:"expiry_time,omitempty"`
	ExpiryTimeLessThan        string `url:"expiry_time__lt,omitempty"`
	ExpiryTimeGreaterThan     string `url:"expiry_time__gt,omitempty"`
	ExpiryTimeLessOrEqual     string `url:"expiry_time__lte,omitempty"`
	ExpiryTimeGreaterOrEqual  string `url:"expiry_time__gte,omitempty"`
	DurationEquals            int64  `url:"duration,omitempty"`
	DurationLessThan          int64  `url:"duration__lt,omitempty"`
	DurationGreaterThan       int64  `url:"duration__gt,omitempty"`
	DurationLessOrEqual       int64  `url:"duration__lte,omitempty"`
	DurationGreaterOrEqual    int64  `url:"duration__gte,omitempty"`
	Limit                     int64  `url:"limit,omitempty"`
	Offset                    int64  `url:"offset,omitempty"`
	SubAccount                string `url:"subaccount,omitempty"`
}

// CreateResponse is the response from creating a masking session.
type CreateResponse struct {
	ApiID         string  `json:"api_id,omitempty"`
	SessionUUID   string  `json:"session_uuid,omitempty"`
	VirtualNumber string  `json:"virtual_number,omitempty"`
	Message       string  `json:"message,omitempty"`
	Session       Session `json:"session,omitempty"`
}

// DeleteResponse is the response from deleting a masking session.
type DeleteResponse struct {
	ApiID   string `json:"api_id,omitempty"`
	Message string `json:"message,omitempty"`
}

// UpdateResponse is the response from updating a masking session.
type UpdateResponse struct {
	ApiID   string  `json:"api_id,omitempty"`
	Message string  `json:"message,omitempty"`
	Session Session `json:"session,omitempty"`
}

// GetResponse is the response from fetching a masking session.
type GetResponse struct {
	ApiID    string  `json:"api_id,omitempty"`
	Response Session `json:"response,omitempty"`
}

// ListSessionResponse wraps a paginated list of sessions.
type ListSessionResponse struct {
	Meta    models.Meta `json:"meta"`
	Objects []Session   `json:"objects"`
}

// ListResponse wraps the list API response.
type ListResponse struct {
	ApiID    string              `json:"api_id"`
	Response ListSessionResponse `json:"response"`
}

// Service exposes masking session API operations.
type Service struct{ r interfaces.Requester }

// New creates a new masking session Service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }

func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Masking/Session")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) Delete(sessionID string) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Masking/Session/%s", sessionID)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) Update(sessionID string, params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Masking/Session/%s", sessionID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) Get(sessionID string) (*GetResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "Masking/Session/%s", sessionID)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "Masking/Session")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
