// Package media provides the Vobiz Media API service.
package media

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes media API operations.
type Service struct {
	r         interfaces.Requester
	authID    string
	authToken string
	baseURL   string
}

// New creates a new media Service.
func New(r interfaces.Requester) *Service {
	return &Service{
		r:       r,
		authID:  r.AuthID(),
		baseURL: "https://api.vobiz.ai",
	}
}
