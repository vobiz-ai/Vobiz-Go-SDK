// Package campaign provides the Vobiz 10DLC Campaign API service.
package campaign

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes campaign API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new campaign Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
