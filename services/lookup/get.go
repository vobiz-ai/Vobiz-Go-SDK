package lookup

import (
	"encoding/json"

	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
)

// Get looks up a phone number.
func (s *Service) Get(number string, params Params) (*Response, error) {
	if params.Type == "" {
		params.Type = "carrier"
	}
	req, err := s.r.NewRequest("GET", params, "v1/Number/%s", number)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	if err := s.r.ExecuteRequest(req, resp, internalhttp.IsLookupRequest()); err != nil {
		return nil, newError(err.Error())
	}
	return resp, nil
}

func newError(body string) *Error {
	e := &Error{respBody: body}
	_ = json.Unmarshal([]byte(body), e)
	return e
}
