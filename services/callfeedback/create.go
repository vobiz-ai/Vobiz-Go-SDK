package callfeedback

import (
	"errors"
	"fmt"

	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
)

// Create submits feedback for a call.
func (s *Service) Create(params CreateParams) (*CreateResponse, error) {
	if params.CallUUID == "" {
		return nil, errors.New("CallUUID cannot be empty")
	}
	if params.Rating == nil {
		return nil, errors.New("Rating cannot be nil")
	}

	feedbackPath := fmt.Sprintf(internalhttp.CallInsightsFeedbackPath, params.CallUUID)
	extra := map[string]interface{}{
		internalhttp.CallInsightsParams: map[string]interface{}{
			internalhttp.CallInsightsRequestPath: feedbackPath,
		},
	}

	req, err := s.r.NewRequest("POST", params, "Call", extra)
	if err != nil {
		return nil, err
	}
	resp := &CreateResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
