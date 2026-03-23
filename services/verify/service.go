// Package verify provides the Vobiz Verify Session API service.
package verify

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes verify session API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new verify Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
