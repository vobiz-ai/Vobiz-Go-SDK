// Package endpoints provides the Vobiz Endpoints API service.
package endpoints

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes endpoint API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new endpoints Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
