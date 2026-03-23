package messages_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/client"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/messages"
)

type messagesMockTransport struct {
	lastRequest *http.Request
	response    *http.Response
	err         error
}

func (m *messagesMockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.lastRequest = req
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}

func TestMessagesSend_BuildsRequestAndParsesResponse_Clean(t *testing.T) {
	mock := &messagesMockTransport{response: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"message":"sent","api_id":"api-2","message_uuid":["msg-1"],"error":""}`)), Header: make(http.Header)}}
	c, err := client.NewClient("AUTH_ID", "AUTH_TOKEN", &client.Options{HTTPClient: &http.Client{Transport: mock}})
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
	if !strings.Contains(string(body), `"src":"+14155550000"`) || !strings.Contains(string(body), `"dst":"+14155550001"`) || !strings.Contains(string(body), `"text":"hello"`) {
		t.Fatalf("unexpected body: %s", string(body))
	}
	if resp.ApiID != "api-2" || len(resp.MessageUUID) != 1 || resp.MessageUUID[0] != "msg-1" {
		t.Fatalf("unexpected parsed response: %+v", resp)
	}
}
