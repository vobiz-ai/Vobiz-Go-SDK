// Package callfeedback provides the Vobiz Call Feedback (Call Insights) API service.
package callfeedback

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// Service exposes call feedback API operations.
type Service struct {
	r interfaces.Requester
}

// New creates a new callfeedback Service.
func New(r interfaces.Requester) *Service {
	return &Service{r: r}
}
