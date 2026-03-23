// Package calls provides the Vobiz Calls API service.
package calls

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes all call-related API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new calls Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
