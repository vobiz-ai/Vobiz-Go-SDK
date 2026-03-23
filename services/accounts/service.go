// Package accounts provides the Vobiz Account and Subaccount API services.
package accounts

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes account-level operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new accounts Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}

// SubaccountService exposes subaccount CRUD operations.
type SubaccountService struct {
	r interfaces.Requester
}

// NewSubaccount creates a new SubaccountService.
func NewSubaccount(r interfaces.Requester) *SubaccountService {
	return &SubaccountService{r: r}
}
