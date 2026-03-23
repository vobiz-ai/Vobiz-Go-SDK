// Package brand provides the Vobiz 10DLC Brand API service.
package brand

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes brand API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new brand Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
