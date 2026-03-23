package http

import (
	"io"
	"net/http"
	"strings"
)

// MockHTTPClient captures outgoing requests and returns preconfigured responses.
type MockHTTPClient struct {
	LastRequest *http.Request
	Response    *http.Response
	Err         error
}

// Do implements the request executor behavior.
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.LastRequest = req
	if m.Err != nil {
		return nil, m.Err
	}
	if m.Response != nil {
		if m.Response.Body == nil {
			m.Response.Body = io.NopCloser(strings.NewReader(""))
		}
		return m.Response, nil
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("{}")),
		Header:     make(http.Header),
	}, nil
}

type mockRoundTripper struct {
	mock *MockHTTPClient
}

func (mrt mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return mrt.mock.Do(req)
}

// NewMockedHTTPClient adapts MockHTTPClient into a standard *http.Client.
func NewMockedHTTPClient(mock *MockHTTPClient) *http.Client {
	return &http.Client{Transport: mockRoundTripper{mock: mock}}
}
