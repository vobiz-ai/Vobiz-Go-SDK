// Package multipartycall provides the Vobiz Multi-Party Call API service.
package multipartycall

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes multi-party call API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new multipartycall Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
