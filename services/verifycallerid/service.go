// Package verifycallerid provides the Vobiz Verified Caller ID API service.
package verifycallerid

import (
	"time"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"
	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/models"
)

// InitiateParams holds parameters for initiating a caller ID verification.
type InitiateParams struct {
	PhoneNumber string `json:"phone_number"`
	Alias       string `json:"alias"`
	Channel     string `json:"channel"`
	Country     string `json:"country"`
	SubAccount  string `json:"subaccount"`
}

// InitiateResponse is the response from initiating a verification.
type InitiateResponse struct {
	ApiID            string `json:"api_id,omitempty"`
	Message          string `json:"message,omitempty"`
	VerificationUUID string `json:"verification_uuid,omitempty"`
}

// VerifyResponse is the response from verifying a caller ID.
type VerifyResponse struct {
	Alias            string    `json:"alias,omitempty"`
	ApiID            string    `json:"api_id,omitempty"`
	Channel          string    `json:"channel"`
	Country          string    `json:"country"`
	CreatedAt        time.Time `json:"created_at"`
	PhoneNumber      string    `json:"phone_number"`
	VerificationUUID string    `json:"verification_uuid"`
	SubAccount       string    `json:"subaccount,omitempty"`
}

// GetResponse is the response from fetching a verified caller ID.
type GetResponse struct {
	Alias            string    `json:"alias,omitempty"`
	ApiID            string    `json:"api_id,omitempty"`
	Country          string    `json:"country"`
	CreatedAt        time.Time `json:"created_at"`
	ModifiedAt       time.Time `json:"modified_at"`
	PhoneNumber      string    `json:"phone_number"`
	SubAccount       string    `json:"subaccount,omitempty"`
	VerificationUUID string    `json:"verification_uuid"`
}

// ListItem represents a single verified caller ID in a list.
type ListItem struct {
	Alias            string    `json:"alias,omitempty"`
	Country          string    `json:"country"`
	CreatedAt        time.Time `json:"created_at"`
	ModifiedAt       time.Time `json:"modified_at"`
	PhoneNumber      string    `json:"phone_number"`
	ResourceUri      string    `json:"resource_uri,omitempty"`
	SubAccount       string    `json:"subaccount,omitempty"`
	VerificationUUID string    `json:"verification_uuid"`
}

// ListParams filters verified caller IDs.
type ListParams struct {
	Country    string `url:"country,omitempty"`
	SubAccount string `url:"subaccount,omitempty"`
	Alias      string `url:"alias,omitempty"`
	Limit      int64  `url:"limit,omitempty"`
	Offset     int64  `url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of verified caller IDs.
type ListResponse struct {
	ApiID   string      `json:"api_id,omitempty"`
	Meta    models.Meta `json:"meta"`
	Objects []ListItem  `json:"objects"`
}

// UpdateParams holds parameters for updating a verified caller ID.
type UpdateParams struct {
	Alias      string `json:"alias,omitempty"`
	SubAccount string `json:"subaccount,omitempty"`
}

// Service exposes verified caller ID API operations.
type Service struct{ r interfaces.Requester }

// New creates a new verifycallerid Service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }

func (s *Service) Initiate(params InitiateParams) (*InitiateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "VerifiedCallerId")
	if err != nil {
		return nil, err
	}
	resp := &InitiateResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) Verify(verificationUUID, otp string) (*VerifyResponse, error) {
	req, err := s.r.NewRequest("POST", map[string]string{"otp": otp}, "VerifiedCallerId/Verification/%s", verificationUUID)
	if err != nil {
		return nil, err
	}
	resp := &VerifyResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) Delete(phoneNumber string) error {
	req, err := s.r.NewRequest("DELETE", nil, "VerifiedCallerId/%s", phoneNumber)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil, internalhttp.IsVoiceRequest())
}

func (s *Service) Update(phoneNumber string, params UpdateParams) (*GetResponse, error) {
	req, err := s.r.NewRequest("POST", params, "VerifiedCallerId/%s", phoneNumber)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) Get(phoneNumber string) (*GetResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "VerifiedCallerId/%s", phoneNumber)
	if err != nil {
		return nil, err
	}
	resp := &GetResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}

func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "VerifiedCallerId")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp, internalhttp.IsVoiceRequest())
}
