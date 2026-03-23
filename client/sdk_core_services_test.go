package client

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/calls"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/messages"
)

type mockTransport struct {
	lastRequest *http.Request
	response    *http.Response
	err         error
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.lastRequest = req
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}

func TestCoreServicesRequestConstructionAndParsing(t *testing.T) {
	t.Run("calls create", func(t *testing.T) {
		mock := &mockTransport{response: &http.Response{StatusCode: http.StatusCreated, Body: io.NopCloser(strings.NewReader(`{"message":"queued","api_id":"api-1","request_uuid":"req-1"}`)), Header: make(http.Header)}}
		c, err := NewClient("AUTH_ID", "AUTH_TOKEN", &Options{HTTPClient: &http.Client{Transport: mock}})
		if err != nil {
			t.Fatalf("new client: %v", err)
		}

		resp, err := c.Calls.Create(calls.CreateParams{From: "+14155550000", To: "+14155550001", AnswerURL: "https://example.com/answer"})
		if err != nil {
			t.Fatalf("create call: %v", err)
		}
		if mock.lastRequest.Method != http.MethodPost {
			t.Fatalf("method = %s", mock.lastRequest.Method)
		}
		if mock.lastRequest.URL.Path != "/api/v1/Account/AUTH_ID/Call/" {
			t.Fatalf("path = %s", mock.lastRequest.URL.Path)
		}
		body, _ := io.ReadAll(mock.lastRequest.Body)
		if !strings.Contains(string(body), `"from":"+14155550000"`) {
			t.Fatalf("unexpected body: %s", string(body))
		}
		if resp.ApiID != "api-1" {
			t.Fatalf("api_id = %s", resp.ApiID)
		}
	})

	t.Run("messages send", func(t *testing.T) {
		mock := &mockTransport{response: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"message":"sent","api_id":"api-2","message_uuid":["msg-1"],"error":""}`)), Header: make(http.Header)}}
		c, err := NewClient("AUTH_ID", "AUTH_TOKEN", &Options{HTTPClient: &http.Client{Transport: mock}})
		if err != nil {
			t.Fatalf("new client: %v", err)
		}
		resp, err := c.Messages.Send(messages.SendParams{Src: "+14155550000", Dst: "+14155550001", Text: "hello"})
		if err != nil {
			t.Fatalf("send message: %v", err)
		}
		if mock.lastRequest.Method != http.MethodPost || mock.lastRequest.URL.Path != "/api/v1/Account/AUTH_ID/Message/" {
			t.Fatalf("unexpected request: %s %s", mock.lastRequest.Method, mock.lastRequest.URL.Path)
		}
		body, _ := io.ReadAll(mock.lastRequest.Body)
		if !strings.Contains(string(body), `"text":"hello"`) {
			t.Fatalf("unexpected body: %s", string(body))
		}
		if resp.ApiID != "api-2" {
			t.Fatalf("api_id = %s", resp.ApiID)
		}
	})

	t.Run("numbers get", func(t *testing.T) {
		mock := &mockTransport{response: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"number":"+14155550000","api_id":"api-3"}`)), Header: make(http.Header)}}
		c, err := NewClient("AUTH_ID", "AUTH_TOKEN", &Options{HTTPClient: &http.Client{Transport: mock}})
		if err != nil {
			t.Fatalf("new client: %v", err)
		}
		resp, err := c.Numbers.Get("+14155550000")
		if err != nil {
			t.Fatalf("get number: %v", err)
		}
		if mock.lastRequest.Method != http.MethodGet || mock.lastRequest.URL.Path != "/api/v1/Account/AUTH_ID/Number/+14155550000/" {
			t.Fatalf("unexpected request: %s %s", mock.lastRequest.Method, mock.lastRequest.URL.Path)
		}
		if resp.Number != "+14155550000" {
			t.Fatalf("number = %s", resp.Number)
		}
	})
}
