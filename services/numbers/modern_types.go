package numbers

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type ModernNumber struct {
	Number       string `json:"number,omitempty"`
	E164         string `json:"e164,omitempty"`
	Country      string `json:"country,omitempty"`
	Region       string `json:"region,omitempty"`
	Application  string `json:"application,omitempty"`
	ResourceURI  string `json:"resource_uri,omitempty"`
}

type ModernListParams struct {
	Page               int  `url:"page,omitempty"`
	PerPage            int  `url:"per_page,omitempty"`
	IncludeSubaccounts bool `url:"include_subaccounts,omitempty"`
}

type ModernListResponse struct {
	ApiID   string         `json:"api_id"`
	Meta    *models.Meta   `json:"meta"`
	Objects []*ModernNumber `json:"objects"`
}

type InventoryListParams struct {
	Country string `url:"country,omitempty"`
	Region  string `url:"region,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
}

type PurchaseFromInventoryParams struct {
	E164     string `json:"e164"`
	Currency string `json:"currency,omitempty"`
}

type LinkApplicationParams struct {
	ApplicationID string `json:"application_id"`
}
