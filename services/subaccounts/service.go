package subaccounts

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes modern sub-accounts API operations.
type Service struct{ r interfaces.Requester }

// New creates a new subaccounts service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }