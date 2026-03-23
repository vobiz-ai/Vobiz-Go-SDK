// Package http provides the core HTTP transport layer for the Vobiz SDK.
// It handles request construction, auth headers, JSON encoding/decoding,
// and voice-request retry logic.
package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	SDKVersion   = "7.59.7"
	LookupBaseURL = "lookup.vobiz.com"

	// CallInsights constants
	CallInsightsBaseURL      = "stats.vobiz.com"
	CallInsightsFeedbackPath = "v1/Call/%s/Feedback/"
	CallInsightsRequestPath  = "call_insights_feedback_path"
	CallInsightsParams       = "call_insights_params"
)

// BaseClient holds HTTP transport state and auth credentials.
// All service types embed or reference a *BaseClient.
type BaseClient struct {
	HTTPClient *http.Client

	AuthID    string
	AuthToken string

	BaseURL   *url.URL
	UserAgent string

	RequestInterceptor  func(r *http.Request)
	ResponseInterceptor func(r *http.Response)
}

// NewRequest builds an authenticated *http.Request.
//
//   - method          – HTTP verb
//   - params          – struct/map encoded as query string (GET) or JSON body (POST/DELETE)
//   - basePathFmt     – printf template for the full path, e.g. "/api/v1/Account/%s/"
//   - pathFmt         – printf template appended after the base, e.g. "Call/%s"
//   - pathArgs        – values substituted into pathFmt
func (c *BaseClient) NewRequest(
	method string,
	params interface{},
	basePathFmt string,
	pathFmt string,
	pathArgs ...interface{},
) (*http.Request, error) {
	if c == nil || c.HTTPClient == nil {
		return nil, errors.New("BaseClient and HTTPClient must not be nil")
	}

	// Detect Call Insights requests (special host override).
	isCallInsights := false
	var callInsightsPath string
	for i, arg := range pathArgs {
		if !isCallInsights {
			isCallInsights, callInsightsPath = detectCallInsights(arg)
		}
		if arg == nil || arg == "" {
			return nil, fmt.Errorf("path parameter #%d is nil/empty", i)
		}
	}

	reqURL := *c.BaseURL
	reqURL.Path = fmt.Sprintf(basePathFmt, fmt.Sprintf(pathFmt, pathArgs...))

	if isCallInsights {
		reqURL.Host = CallInsightsBaseURL
		reqURL.Path = callInsightsPath
	}

	var buf bytes.Buffer
	if method == "GET" {
		vals, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		reqURL.RawQuery = vals.Encode()
	} else {
		if params != nil {
			rv := reflect.ValueOf(params)
			if rv.Kind() != reflect.Map || !rv.IsNil() {
				if err := json.NewEncoder(&buf).Encode(params); err != nil {
					return nil, err
				}
			}
		}
	}

	req, err := http.NewRequest(method, reqURL.String(), &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-ID", c.AuthID)
	req.Header.Set("X-Auth-Token", c.AuthToken)

	if c.RequestInterceptor != nil {
		c.RequestInterceptor(req)
	}

	return req, nil
}

// ExecuteRequest sends req, decodes the response body into dest (if non-nil),
// and handles voice-request retry logic via the extra map.
//
// Recognised extra keys:
//   - "is_voice_request" – enables host-failover retry
//   - "is_lookup_request" – overrides host to LookupBaseURL
//   - "retry"            – current retry count (int)
func (c *BaseClient) ExecuteRequest(req *http.Request, dest interface{}, extra ...map[string]interface{}) error {
	isVoice := false
	if len(extra) > 0 {
		if _, ok := extra[0]["is_voice_request"]; ok {
			isVoice = true
			req.URL.Host = voiceHost(extra[0])
			req.Host = req.URL.Host
			req.URL.Scheme = "https"
		}
		if _, ok := extra[0]["is_lookup_request"]; ok {
			if req.URL.Host == "api.vobiz.com" { // unit-test sentinel
				req.URL.Host = LookupBaseURL
				req.Host = LookupBaseURL
				req.URL.Scheme = "https"
			}
		}
	}

	// Preserve body for retries.
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if c.ResponseInterceptor != nil {
		c.ResponseInterceptor(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}

	// Voice retry on 5xx.
	if isVoice && resp.StatusCode >= 500 {
		retryCount, _ := extra[0]["retry"].(int)
		if retryCount >= 2 {
			return errors.New(string(data))
		}
		extra[0]["retry"] = retryCount + 1
		retryReq, _ := http.NewRequest(req.Method, req.URL.String(), bytes.NewReader(bodyBytes))
		retryReq.Header.Set("User-Agent", c.UserAgent)
		retryReq.Header.Set("Content-Type", "application/json")
		retryReq.Header.Set("X-Auth-ID", c.AuthID)
		retryReq.Header.Set("X-Auth-Token", c.AuthToken)
		return c.ExecuteRequest(retryReq, dest, extra...)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if dest != nil {
			return json.Unmarshal(data, dest)
		}
		return nil
	}
	if string(data) == "{}" && resp.StatusCode == 404 {
		return errors.New("resource not found: " + resp.Status)
	}
	return errors.New(string(data))
}

// IsVoiceRequest returns the extra map needed to mark a request as a voice request.
func IsVoiceRequest() map[string]interface{} {
	return map[string]interface{}{
		"is_voice_request": true,
		"retry":            0,
	}
}

// IsLookupRequest returns the extra map needed to mark a request as a lookup request.
func IsLookupRequest() map[string]interface{} {
	return map[string]interface{}{
		"is_lookup_request": true,
	}
}

// ---- private helpers -------------------------------------------------------

var (
	voiceBaseURL          = "api.vobiz.ai"
	voiceBaseURLFallback1 = "api.vobiz.ai"
	voiceBaseURLFallback2 = "api.vobiz.ai"
)

func voiceHost(extra map[string]interface{}) string {
	retry, _ := extra["retry"].(int)
	switch retry {
	case 1:
		return voiceBaseURLFallback1
	case 2:
		return voiceBaseURLFallback2
	default:
		return voiceBaseURL
	}
}

func detectCallInsights(param interface{}) (bool, string) {
	if param == nil {
		return false, ""
	}
	rv := reflect.ValueOf(param)
	if rv.Kind() != reflect.Map || rv.Type().Key().Kind() != reflect.String {
		return false, ""
	}
	m, ok := param.(map[string]interface{})
	if !ok {
		return false, ""
	}
	inner, ok := m[CallInsightsParams].(map[string]interface{})
	if !ok {
		return false, ""
	}
	path, ok := inner[CallInsightsRequestPath].(string)
	return ok, path
}
