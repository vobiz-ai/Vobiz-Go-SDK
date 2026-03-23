// Package lookup provides the Vobiz Lookup API service.
package lookup

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes lookup API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new lookup Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
