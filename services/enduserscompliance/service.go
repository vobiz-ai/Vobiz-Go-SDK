// Package enduserscompliance provides the Vobiz End User compliance API service.
package enduserscompliance

import (
	"time"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/models"
)

// EndUser represents a compliance end user.
type EndUser struct {
	CreatedAt   time.Time `json:"created_at"`
	EndUserID   string    `json:"end_user_id"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	EndUserType string    `json:"end_user_type"`
}

// CreateEndUserResponse is the response from creating an end user.
type CreateEndUserResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	EndUserID   string    `json:"end_user_id"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	EndUserType string    `json:"end_user_type"`
	APIID       string    `json:"api_id"`
	Message     string    `json:"message"`
}

// ListParams filters end users.
type ListParams struct {
	Limit       int    `url:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty"`
	Name        string `url:"name,omitempty"`
	LastName    string `url:"last_name,omitempty"`
	EndUserType string `url:"end_user_type,omitempty"`
}

// ListResponse wraps a paginated list of end users.
type ListResponse struct {
	models.BaseListResponse
	Objects []EndUser `json:"objects"`
}

// CreateParams holds parameters for creating an end user.
type CreateParams struct {
	Name        string `json:"name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	EndUserType string `json:"end_user_type,omitempty"`
}

// UpdateParams holds parameters for updating an end user.
type UpdateParams struct {
	CreateParams
	EndUserID string `json:"end_user_id"`
}

// UpdateResponse is the response from updating an end user.
type UpdateResponse = models.BaseResponse

// Service exposes end user compliance API operations.
type Service struct{ r interfaces.Requester }

// New creates a new end users compliance Service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }

func (s *Service) Get(endUserID string) (*EndUser, error) {
	req, err := s.r.NewRequest("GET", nil, "EndUser/%s", endUserID)
	if err != nil {
		return nil, err
	}
	resp := &EndUser{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) List(params ListParams) (*ListResponse, error) {
	req, err := s.r.NewRequest("GET", params, "EndUser")
	if err != nil {
		return nil, err
	}
	resp := &ListResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Create(params CreateParams) (*CreateEndUserResponse, error) {
	req, err := s.r.NewRequest("POST", params, "EndUser")
	if err != nil {
		return nil, err
	}
	resp := &CreateEndUserResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Update(params UpdateParams) (*UpdateResponse, error) {
	req, err := s.r.NewRequest("POST", params, "EndUser/%s", params.EndUserID)
	if err != nil {
		return nil, err
	}
	resp := &UpdateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

func (s *Service) Delete(endUserID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "EndUser/%s", endUserID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil)
}
