package subaccounts

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type Subaccount struct {
	Account     string `json:"account,omitempty"`
	ApiID       string `json:"api_id,omitempty"`
	AuthID      string `json:"auth_id,omitempty"`
	AuthToken   string `json:"auth_token,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Name        string `json:"name,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type CreateParams struct {
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type UpdateParams = CreateParams

type DeleteParams struct {
	Cascade bool `json:"cascade,omitempty"`
}

type CreateResponse struct {
	models.BaseResponse
	AuthId    string `json:"auth_id"`
	AuthToken string `json:"auth_token"`
}

type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type ListResponse struct {
	models.BaseListResponse
	Objects []Subaccount `json:"objects"`
}

type UpdateResponse = models.BaseResponse
