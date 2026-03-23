// Package transcription provides the Vobiz Transcription API service.
package transcription

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes transcription API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new transcription Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
