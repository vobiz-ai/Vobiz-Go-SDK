// Package powerpack provides the Vobiz Powerpack API service.
package powerpack

import (
	"fmt"
	"strings"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/numbers"
)

// Priority holds number pool priorities.
type Priority struct {
	Priority1 *string `json:"priority1"`
	Priority2 *string `json:"priority2"`
	Priority3 *string `json:"priority3"`
}

// NumberPriority holds priority configuration per service type and country.
type NumberPriority struct {
	ServiceType string   `json:"service_type"`
	CountryISO  string   `json:"country_iso"`
	Priority    Priority `json:"priority"`
}

// Powerpack represents a Vobiz Powerpack.
type Powerpack struct {
	UUID             string           `json:"uuid,omitempty"`
	Name             string           `json:"name,omitempty"`
	StickySender     bool             `json:"sticky_sender,omitempty"`
	LocalConnect     bool             `json:"local_connect,omitempty"`
	ApplicationType  string           `json:"application_type,omitempty"`
	ApplicationID    string           `json:"application_id,omitempty"`
	NumberPoolUUID   string           `json:"number_pool,omitempty"`
	CreatedOn        string           `json:"created_on,omitempty"`
	NumberPriorities []NumberPriority `json:"number_priority,omitempty"`
}

// CreateParams holds parameters for creating a powerpack.
type CreateParams struct {
	Name             string           `json:"name,omitempty"`
	StickySender     bool             `json:"sticky_sender,omitempty"`
	LocalConnect     bool             `json:"local_connect,omitempty"`
	ApplicationType  string           `json:"application_type,omitempty"`
	ApplicationID    string           `json:"application_id,omitempty"`
	NumberPriorities []NumberPriority `json:"number_priority,omitempty"`
}

// UpdateParams holds parameters for updating a powerpack.
type UpdateParams struct {
	Name             string           `json:"name,omitempty"`
	StickySender     bool             `json:"sticky_sender,omitempty"`
	LocalConnect     bool             `json:"local_connect,omitempty"`
	ApplicationType  string           `json:"application_type,omitempty"`
	ApplicationID    string           `json:"application_id,omitempty"`
	NumberPriorities []NumberPriority `json:"number_priority,omitempty"`
}

// ListParams filters powerpack list results.
type ListParams struct {
	Limit   int    `url:"limit,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Service string `url:"service,omitempty"`
}

// SearchParams filters number pool searches.
type SearchParams struct {
	StartsWith  string `url:"starts_with,omitempty"`
	CountryISO2 string `url:"country_iso2,omitempty"`
	Type        string `url:"type,omitempty"`
	Limit       string `url:"limit,omitempty"`
	Offset      string `url:"offset,omitempty"`
	Service     string `url:"service,omitempty"`
}

// Meta holds pagination metadata.
type Meta struct {
	Previous   *string
	Next       *string
	TotalCount int `json:"total_count"`
	Offset     int
	Limit      int
}

// BaseListResponse is the base for list responses.
type BaseListResponse struct {
	ApiID string `json:"api_id"`
	Meta  Meta   `json:"meta"`
}

// NumberPoolEntry represents a number in the pool.
type NumberPoolEntry struct {
	NumberPoolUUID              string `json:"number_pool_uuid,omitempty"`
	Number                      string `json:"number,omitempty"`
	Type                        string `json:"Type,omitempty"`
	CountryISO2                 string `json:"country_iso2,omitempty"`
	Service                     string `json:"service,omitempty"`
	AddedOn                     string `json:"added_on,omitempty"`
	AccountPhoneNumberResource  string `json:"account_phone_number_resource,omitempty"`
}

// NumberResponse is the response from number pool operations.
type NumberResponse struct {
	ApiID string `json:"api_id,omitempty"`
	NumberPoolEntry
	Error string `json:"error,omitempty"`
}

// ShortCode represents a short code in the pool.
type ShortCode struct {
	NumberPoolUUID string `json:"number_pool_uuid,omitempty"`
	Shortcode      string `json:"shortcode,omitempty"`
	CountryISO2    string `json:"country_iso2,omitempty"`
	AddedOn        string `json:"added_on,omitempty"`
	Service        string `json:"service,omitempty"`
}

// Tollfree represents a toll-free number in the pool.
type Tollfree struct {
	NumberPoolUUID string `json:"number_pool_uuid,omitempty"`
	Tollfree       string `json:"number,omitempty"`
	CountryISO2    string `json:"country_iso2,omitempty"`
	AddedOn        string `json:"added_on,omitempty"`
	Service        string `json:"service,omitempty"`
}

// ShortCodeResponse wraps a list of short codes.
type ShortCodeResponse struct {
	BaseListResponse
	Objects []ShortCode `json:"objects"`
}

// TollfreeResponse wraps a list of toll-free numbers.
type TollfreeResponse struct {
	BaseListResponse
	Objects []Tollfree `json:"objects"`
}

// FindShortCodeResponse is the response from finding a short code.
type FindShortCodeResponse struct {
	ApiID string `json:"api_id,omitempty"`
	ShortCode
	Error string `json:"error,omitempty"`
}

// FindTollfreeResponse is the response from finding a toll-free number.
type FindTollfreeResponse struct {
	ApiID string `json:"api_id,omitempty"`
	Tollfree
	Error string `json:"error,omitempty"`
}

// CreateResponse is the response from creating a powerpack.
type CreateResponse struct {
	ApiID string  `json:"api_id,omitempty"`
	Error *string `json:"error,omitempty"`
	Powerpack
}

// UpdateResponse is the response from updating a powerpack.
type UpdateResponse = CreateResponse

// ListResponse wraps a paginated list of powerpacks.
type ListResponse struct {
	BaseListResponse
	Objects []Powerpack `json:"objects"`
}

// NumbersListResponse wraps a paginated list of pool numbers.
type NumbersListResponse struct {
	BaseListResponse
	Objects []NumberPoolEntry `json:"objects"`
}

// DeleteResponse is the response from deleting a powerpack.
type DeleteResponse struct {
	ApiID    string `json:"api_id,omitempty"`
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}

// DeleteParams holds optional parameters for deleting a powerpack.
type DeleteParams struct {
	UnrentNumbers bool `json:"unrent_numbers,omitempty"`
}

// NumberRemoveParams holds optional parameters for removing a number.
type NumberRemoveParams struct {
	Unrent bool `json:"unrent,omitempty"`
}

// AddNumberOptions holds optional parameters for adding a number.
type AddNumberOptions struct {
	Service string `json:"service,omitempty"`
}

// BuyNumberParams holds parameters for buying and adding a number.
type BuyNumberParams struct {
	Number      string
	CountryISO2 string
	Type        string
	Region      string
	Pattern     string
	Service     string
}

// Service exposes powerpack API operations.
// After Get(), the returned *Service has the Powerpack embedded so number pool
// operations can be called on it directly.
type Service struct {
	r            interfaces.Requester
	phoneNumbers *numbers.PhoneNumberService
	Powerpack
}

// New creates a new powerpack Service.
func New(r interfaces.Requester) *Service {
	return &Service{
		r:            r,
		phoneNumbers: numbers.NewPhoneNumber(r),
	}
}

func (s *Service) poolUUID() string {
	parts := strings.Split(s.Powerpack.NumberPoolUUID, "/")
	if len(parts) >= 7 {
		return parts[6]
	}
	return s.Powerpack.NumberPoolUUID
}

func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "Powerpack/")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Get(powerpackUUID string) (*Service, error) {
	req, err := s.r.NewRequest("GET", nil, "Powerpack/%s/", powerpackUUID)
	if err != nil {
		return nil, err
	}
	body := &CreateResponse{}
	if err := s.r.ExecuteRequest(req, body); err != nil {
		return nil, err
	}
	result := &Service{r: s.r, phoneNumbers: s.phoneNumbers, Powerpack: body.Powerpack}
	return result, nil
}

func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Powerpack")
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Update(params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Powerpack/%s", s.Powerpack.UUID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Delete(params DeleteParams) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", params, "Powerpack/%s", s.Powerpack.UUID)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) ListNumbers(params SearchParams) (*NumbersListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "NumberPool/%s/Number/", s.poolUUID())
	if err != nil {
		return nil, err
	}
	resp := &NumbersListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) CountNumbers(params SearchParams) (int, error) {
	resp, err := s.ListNumbers(params)
	if err != nil {
		return 0, err
	}
	return resp.Meta.TotalCount, nil
}

func (s *Service) FindNumber(number string) (*NumberResponse, error) {
	return s.FindNumberWithOptions(number, AddNumberOptions{})
}

func (s *Service) FindNumberWithOptions(number string, opts AddNumberOptions) (*NumberResponse, error) {
	req, err := s.r.NewRequest("GET", opts, "NumberPool/%s/Number/%s/", s.poolUUID(), number)
	if err != nil {
		return nil, err
	}
	resp := &NumberResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) AddNumber(number string) (*NumberResponse, error) {
	return s.AddNumberWithOptions(number, AddNumberOptions{})
}

func (s *Service) AddNumberWithOptions(number string, opts AddNumberOptions) (*NumberResponse, error) {
	req, err := s.r.NewRequest("POST", opts, "NumberPool/%s/Number/%s", s.poolUUID(), number)
	if err != nil {
		return nil, err
	}
	resp := &NumberResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) AddTollfree(tollfree string) (*NumberResponse, error) {
	req, err := s.r.NewRequest("POST", nil, "NumberPool/%s/Tollfree/%s", s.poolUUID(), tollfree)
	if err != nil {
		return nil, err
	}
	resp := &NumberResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) RemoveNumber(number string, params NumberRemoveParams) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", params, "NumberPool/%s/Number/%s", s.poolUUID(), number)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) RemoveTollfree(tollfree string, params NumberRemoveParams) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", params, "NumberPool/%s/Tollfree/%s", s.poolUUID(), tollfree)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) RemoveShortcode(shortcode string) (*DeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "NumberPool/%s/Shortcode/%s", s.poolUUID(), shortcode)
	if err != nil {
		return nil, err
	}
	resp := &DeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) ListShortcodes() (*ShortCodeResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "NumberPool/%s/Shortcode", s.poolUUID())
	if err != nil {
		return nil, err
	}
	resp := &ShortCodeResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) ListTollfree() (*TollfreeResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "NumberPool/%s/Tollfree", s.poolUUID())
	if err != nil {
		return nil, err
	}
	resp := &TollfreeResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) FindShortcode(shortcode string) (*FindShortCodeResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "NumberPool/%s/Shortcode/%s/", s.poolUUID(), shortcode)
	if err != nil {
		return nil, err
	}
	resp := &FindShortCodeResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) FindTollfree(tollfree string) (*FindTollfreeResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "NumberPool/%s/Tollfree/%s/", s.poolUUID(), tollfree)
	if err != nil {
		return nil, err
	}
	resp := &FindTollfreeResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) BuyAddNumber(params BuyNumberParams) (*NumberResponse, error) {
	payload := struct {
		Rent    string `json:"rent"`
		Service string `json:"service,omitempty"`
	}{"true", params.Service}

	if params.Number != "" {
		req, err := s.r.NewRequest("POST", payload, "NumberPool/%s/Number/%s", s.poolUUID(), params.Number)
		if err != nil {
			return nil, err
		}
		resp := &NumberResponse{}
		return resp, s.r.ExecuteRequest(req, resp)
	}

	phoneResp, err := s.phoneNumbers.List(numbers.PhoneNumberListParams{
		CountryISO: params.CountryISO2,
		Type:       params.Type,
		Region:     params.Region,
		Pattern:    params.Pattern,
		Services:   params.Service,
	})
	if err != nil {
		return nil, err
	}
	if len(phoneResp.Objects) == 0 {
		return &NumberResponse{}, nil
	}
	number := phoneResp.Objects[0].Number
	req, err := s.r.NewRequest("POST", payload, "NumberPool/%s/Number/%s", s.poolUUID(), number)
	if err != nil {
		return nil, err
	}
	resp := &NumberResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Ensure fmt is used (for poolUUID fallback path).
var _ = fmt.Sprintf
