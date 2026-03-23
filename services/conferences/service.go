// Package conferences provides the Vobiz Conferences API service.
package conferences

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes conference API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new conferences Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
