package credentials

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type Credential struct {
	ID          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type CreateParams struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UpdateParams struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Enabled  *bool  `json:"enabled,omitempty"`
}

type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type ListResponse struct {
	ApiID   string       `json:"api_id"`
	Meta    *models.Meta `json:"meta"`
	Objects []*Credential `json:"objects"`
}

type CreateResponse = models.BaseResponse

type UpdateResponse = models.BaseResponse
