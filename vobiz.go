// Package vobiz is the public entry point for the Vobiz Go SDK.
//
// Quick start:
//
//	import "github.com/vobiz/all-vobiz-sdk/go-sdk"
//
//	c, err := vobiz.NewClient("AUTH_ID", "AUTH_TOKEN", nil)
//	call, err := c.Calls.Create(calls.CreateParams{
//	    From:      "+14155551234",
//	    To:        "+14155555678",
//	    AnswerURL: "https://example.com/answer",
//	})
package vobiz

import (
	"github.com/vobiz/all-vobiz-sdk/go-sdk/client"
	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
)

// Client is the top-level SDK client. Re-exported from the client package for
// convenience so callers only need to import this package.
type Client = client.Client

// Options configures the SDK client. Re-exported for convenience.
type Options = client.Options

// NewClient creates a fully initialised Vobiz SDK client.
//
// authID and authToken may be empty; the constructor will fall back to the
// VOBIZ_AUTH_ID and VOBIZ_AUTH_TOKEN environment variables.
func NewClient(authID, authToken string, opts *Options) (*Client, error) {
	return client.NewClient(authID, authToken, opts)
}

// SDKVersion is the current SDK version string.
const SDKVersion = internalhttp.SDKVersion
