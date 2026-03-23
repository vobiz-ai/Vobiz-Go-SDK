// Package applications provides the Vobiz Applications API service.
package applications

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes application API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new applications Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
