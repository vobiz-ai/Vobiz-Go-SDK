package originationuris

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type URI struct {
	ID          string `json:"id,omitempty"`
	URI         string `json:"uri,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type CreateParams struct {
	URI      string `json:"uri,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

type UpdateParams struct {
	URI      string `json:"uri,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type ListResponse struct {
	ApiID   string      `json:"api_id"`
	Meta    *models.Meta `json:"meta"`
	Objects []*URI      `json:"objects"`
}

type CreateResponse = models.BaseResponse

type UpdateResponse = models.BaseResponse
