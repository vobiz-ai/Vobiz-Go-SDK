package originationuris

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes origination URI CRUD operations.
type Service struct{ r interfaces.Requester }

// New creates a new originationuris service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }