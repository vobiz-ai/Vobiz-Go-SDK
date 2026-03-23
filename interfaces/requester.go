// Package interfaces defines core SDK interfaces to avoid circular dependencies.
package interfaces

import "net/http"

// Requester is the interface every service depends on.
// It is satisfied by *client.Client, which makes services easy to test with a mock.
type Requester interface {
	NewRequest(method string, params interface{}, pathFmt string, pathArgs ...interface{}) (*http.Request, error)
	ExecuteRequest(req *http.Request, dest interface{}, extra ...map[string]interface{}) error
	AuthID() string
}
