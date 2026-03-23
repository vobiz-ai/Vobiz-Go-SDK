package trunks

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type Trunk struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	CPSLimit    int    `json:"cps_limit,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type CreateParams struct {
	Name     string `json:"name,omitempty"`
	CPSLimit int    `json:"cps_limit,omitempty"`
}

type UpdateParams struct {
	Name     string `json:"name,omitempty"`
	CPSLimit int    `json:"cps_limit,omitempty"`
}

type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type ListResponse struct {
	ApiID   string         `json:"api_id"`
	Meta    *models.Meta   `json:"meta"`
	Objects []*Trunk       `json:"objects"`
}

type CreateResponse = models.BaseResponse

type UpdateResponse = models.BaseResponse
