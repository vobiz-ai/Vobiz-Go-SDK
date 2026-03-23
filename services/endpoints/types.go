package endpoints

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Endpoint represents a SIP endpoint.
type Endpoint struct {
	Alias       string `json:"alias,omitempty"`
	EndpointID  string `json:"endpoint_id,omitempty"`
	Password    string `json:"password,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	SIPURI      string `json:"sip_uri,omitempty"`
	Username    string `json:"username,omitempty"`
	AppID       string `json:"app_id,omitempty"`
}

// CreateParams holds parameters for creating an endpoint.
type CreateParams struct {
	Alias    string `json:"alias,omitempty"`
	Password string `json:"password,omitempty"`
	AppID    string `json:"app_id,omitempty"`
	Username string `json:"username,omitempty"`
}

// CreateResponse is the response from creating an endpoint.
type CreateResponse struct {
	models.BaseResponse
	Alias      string `json:"alias,omitempty"`
	EndpointID string `json:"endpoint_id,omitempty"`
	Username   string `json:"username,omitempty"`
}

// UpdateParams holds parameters for updating an endpoint.
type UpdateParams struct {
	Alias    string `json:"alias,omitempty"`
	Password string `json:"password,omitempty"`
	AppID    string `json:"app_id,omitempty"`
}

// UpdateResponse is the response from updating an endpoint.
type UpdateResponse = models.BaseResponse

// ListParams filters the endpoint list.
type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of endpoints.
type ListResponse struct {
	models.BaseListResponse
	Objects []*Endpoint `json:"objects"`
}
