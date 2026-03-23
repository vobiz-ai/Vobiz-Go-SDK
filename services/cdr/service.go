// Package cdr provides the Vobiz CDR (Call Detail Records) API service.
package cdr

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes CDR API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new CDR Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
