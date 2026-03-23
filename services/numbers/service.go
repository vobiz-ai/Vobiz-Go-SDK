// Package numbers provides the Vobiz Numbers and PhoneNumbers API services.
package numbers

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes number management operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new numbers Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}

// PhoneNumberService exposes phone number catalog operations.
type PhoneNumberService struct {
	r interfaces.Requester
}

// NewPhoneNumber creates a new PhoneNumberService.
func NewPhoneNumber(r interfaces.Requester) *PhoneNumberService {
	return &PhoneNumberService{r: r}
}
