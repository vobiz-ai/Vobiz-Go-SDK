// Package messages provides the Vobiz Messages API service.
package messages

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes all messaging API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new messages Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
