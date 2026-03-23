// Package tollfree provides the Vobiz Toll-Free Verification API service.
package tollfree

import (
	"time"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/models"
)

// Verification represents a toll-free verification request.
type Verification struct {
	UUID                  string    `json:"uuid"`
	ProfileUUID           string    `json:"profile_uuid"`
	Number                string    `json:"number"`
	Usecase               string    `json:"usecase"`
	UsecaseSummary        string    `json:"usecase_summary"`
	MessageSample         string    `json:"message_sample"`
	OptinImageURL         *string   `json:"optin_image_url"`
	OptinType             string    `json:"optin_type"`
	Volume                string    `json:"volume"`
	AdditionalInformation string    `json:"additional_information"`
	ExtraData             string    `json:"extra_data"`
	CallbackURL           string    `json:"callback_url"`
	CallbackMethod        string    `json:"callback_method"`
	Status                string    `json:"status"`
	ErrorMessage          string    `json:"error_message"`
	Created               time.Time `json:"created"`
	LastModified          time.Time `json:"last_modified"`
}

// CreateParams holds parameters for creating a toll-free verification.
type CreateParams struct {
	ProfileUUID           string `json:"profile_uuid,omitempty"`
	Usecase               string `json:"usecase,omitempty"`
	UsecaseSummary        string `json:"usecase_summary,omitempty"`
	MessageSample         string `json:"message_sample,omitempty"`
	OptInImageURL         string `json:"optin_image_url,omitempty"`
	OptInType             string `json:"optin_type,omitempty"`
	Volume                string `json:"volume,omitempty"`
	AdditionalInformation string `json:"additional_information,omitempty"`
	ExtraData             string `json:"extra_data,omitempty"`
	Number                string `json:"number,omitempty"`
	CallbackURL           string `json:"callback_url,omitempty"`
	CallbackMethod        string `json:"callback_method,omitempty"`
}

// UpdateParams holds parameters for updating a toll-free verification.
type UpdateParams struct {
	ProfileUUID           string `json:"profile_uuid,omitempty"`
	Usecase               string `json:"usecase,omitempty"`
	UsecaseSummary        string `json:"usecase_summary,omitempty"`
	MessageSample         string `json:"message_sample,omitempty"`
	OptInImageURL         string `json:"optin_image_url,omitempty"`
	OptInType             string `json:"optin_type,omitempty"`
	Volume                string `json:"volume,omitempty"`
	AdditionalInformation string `json:"additional_information,omitempty"`
	ExtraData             string `json:"extra_data,omitempty"`
	CallbackURL           string `json:"callback_url,omitempty"`
	CallbackMethod        string `json:"callback_method,omitempty"`
}

// ListParams filters toll-free verifications.
type ListParams struct {
	Number      string `url:"number,omitempty"`
	Status      string `url:"status,omitempty"`
	ProfileUUID string `url:"profile_uuid,omitempty"`
	CreatedGT   string `url:"created__gt,omitempty"`
	CreatedGTE  string `url:"created__gte,omitempty"`
	CreatedLT   string `url:"created__lt,omitempty"`
	CreatedLTE  string `url:"created__lte,omitempty"`
	Usecase     string `url:"usecase,omitempty"`
	Limit       int64  `url:"limit,omitempty"`
	Offset      int64  `url:"offset,omitempty"`
}

// CreateResponse is the response from creating a toll-free verification.
type CreateResponse struct {
	APIID   string `json:"api_id"`
	Message string `json:"message,omitempty"`
	UUID    string `json:"uuid"`
}

// Response is the generic response for update/delete operations.
type Response struct {
	APIID   string `json:"api_id"`
	Message string `json:"message,omitempty"`
}

// ListResponse wraps a paginated list of toll-free verifications.
type ListResponse struct {
	APIID   string         `json:"api_id"`
	Meta    *models.Meta   `json:"meta,omitempty"`
	Objects []Verification `json:"objects,omitempty"`
}

// Service exposes toll-free verification API operations.
type Service struct{ r interfaces.Requester }

// New creates a new tollfree Service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }

func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "TollfreeVerification")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Update(uuid string, params UpdateParams) (*Response, error) {
	req, err := s.r.NewRequest("POST", params, "TollfreeVerification/%s", uuid)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Get(uuid string) (*Verification, error) {
	req, err := s.r.NewRequest("GET", nil, "TollfreeVerification/%s", uuid)
	if err != nil {
		return nil, err
	}
	resp := &Verification{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "TollfreeVerification")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Delete(uuid string) error {
	req, err := s.r.NewRequest("DELETE", nil, "TollfreeVerification/%s", uuid)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil)
}
