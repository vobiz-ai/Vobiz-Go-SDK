package credentials

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes credential CRUD operations.
type Service struct{ r interfaces.Requester }

// New creates a new credentials service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }