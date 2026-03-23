// Package pricing provides the Vobiz Pricing API service.
package pricing

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes pricing API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new pricing Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
