// Package profile provides the Vobiz Profile API service.
package profile

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes profile API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new profile Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
