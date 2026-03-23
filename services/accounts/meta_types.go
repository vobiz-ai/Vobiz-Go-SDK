package accounts

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type TransactionListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type Transaction struct {
	ID          string `json:"id,omitempty"`
	Amount      string `json:"amount,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

type TransactionListResponse struct {
	ApiID   string         `json:"api_id"`
	Meta    *models.Meta   `json:"meta"`
	Objects []*Transaction `json:"objects"`
}

type BalanceResponse struct {
	ApiID    string `json:"api_id,omitempty"`
	Balance  string `json:"balance,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type ConcurrencyResponse struct {
	ApiID       string `json:"api_id,omitempty"`
	Concurrency int    `json:"concurrency,omitempty"`
}
