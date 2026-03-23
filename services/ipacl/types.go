package ipacl

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type Entry struct {
	ID          string `json:"id,omitempty"`
	IPAddress   string `json:"ip_address,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type CreateParams struct {
	IPAddress   string `json:"ip_address,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateParams struct {
	IPAddress   string `json:"ip_address,omitempty"`
	Description string `json:"description,omitempty"`
}

type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type ListResponse struct {
	ApiID   string      `json:"api_id"`
	Meta    *models.Meta `json:"meta"`
	Objects []*Entry    `json:"objects"`
}

type CreateResponse = models.BaseResponse

type UpdateResponse = models.BaseResponse
