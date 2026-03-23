// Package recordings provides the Vobiz Recordings API service.
package recordings

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes recording API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new recordings Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
