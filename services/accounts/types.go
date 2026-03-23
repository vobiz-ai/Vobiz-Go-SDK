package accounts

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Account represents a Vobiz account.
type Account struct {
	AccountType  string `json:"account_type,omitempty"`
	Address      string `json:"address,omitempty"`
	ApiID        string `json:"api_id,omitempty"`
	AuthID       string `json:"auth_id,omitempty"`
	AutoRecharge bool   `json:"auto_recharge,omitempty"`
	BillingMode  string `json:"billing_mode,omitempty"`
	CashCredits  string `json:"cash_credits,omitempty"`
	City         string `json:"city,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceURI  string `json:"resource_uri,omitempty"`
	State        string `json:"state,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
}

// UpdateParams are the fields that can be modified on an account.
type UpdateParams struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
}

// UpdateResponse is returned after updating an account or subaccount.
type UpdateResponse = models.BaseResponse

// Subaccount represents a Vobiz subaccount.
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

// SubaccountCreateParams are the parameters for creating a subaccount.
type SubaccountCreateParams struct {
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

// SubaccountUpdateParams are the parameters for updating a subaccount.
type SubaccountUpdateParams = SubaccountCreateParams

// SubaccountDeleteParams controls cascade deletion.
type SubaccountDeleteParams struct {
	Cascade bool `json:"cascade,omitempty"`
}

// SubaccountCreateResponse is returned after creating a subaccount.
type SubaccountCreateResponse struct {
	models.BaseResponse
	AuthId    string `json:"auth_id"`
	AuthToken string `json:"auth_token"`
}

// SubaccountListParams filters the subaccount list.
type SubaccountListParams struct {
	Limit  int `json:"limit,omitempty"  url:"limit,omitempty"`
	Offset int `json:"offset,omitempty" url:"offset,omitempty"`
}

// SubaccountList wraps a paginated list of subaccounts.
type SubaccountList struct {
	models.BaseListResponse
	Objects []Subaccount `json:"objects"`
}
