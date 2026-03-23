// Package token provides the Vobiz JWT Token API service.
package token

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes token API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new token Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
