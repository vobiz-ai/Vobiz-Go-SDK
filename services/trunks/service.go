package trunks

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes trunk CRUD operations.
type Service struct{ r interfaces.Requester }

// New creates a new trunks service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }