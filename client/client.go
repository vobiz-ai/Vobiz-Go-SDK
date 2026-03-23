// Package client wires together all service instances and exposes the public
// entry point for the Vobiz Go SDK.
package client

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"

	internalhttp "github.com/vobiz/all-vobiz-sdk/go-sdk/internal/http"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/accounts"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/applications"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/brand"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/callfeedback"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/calls"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/campaign"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/cdr"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/compliance"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/conferences"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/credentials"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/endpoints"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/ipacl"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/lookup"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/maskingsession"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/media"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/messages"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/multipartycall"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/numbers"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/originationuris"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/powerpack"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/pricing"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/profile"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/recordings"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/subaccounts"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/token"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/tollfree"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/trunks"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/enduserscompliance"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/transcription"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/verify"
	"github.com/vobiz/all-vobiz-sdk/go-sdk/services/verifycallerid"
)

const (
	baseURLString    = "https://api.vobiz.ai/"
	baseRequestFmt   = "/api/v1/Account/%s/"
)

// Options allows callers to customise the underlying HTTP client (proxy, timeout, etc.).
type Options struct {
	HTTPClient *http.Client
}

// Client is the top-level SDK client. Instantiate it with NewClient.
//
// Usage:
//
//	c, err := client.NewClient("AUTH_ID", "AUTH_TOKEN", nil)
//	call, err := c.Calls.Create(calls.CreateParams{...})
type Client struct {
	base *internalhttp.BaseClient

	// Core voice/messaging services
	Calls          *calls.Service
	Messages       *messages.Service
	Recordings     *recordings.Service
	Transcription  *transcription.Service
	Conferences    *conferences.Service
	Endpoints      *endpoints.Service
	Numbers        *numbers.Service
	PhoneNumbers   *numbers.PhoneNumberService
	Trunks         *trunks.Service
	Credentials    *credentials.Service
	IpACL          *ipacl.Service
	OriginationURIs *originationuris.Service
	Applications   *applications.Service
	Accounts       *accounts.Service
	Subaccounts    *accounts.SubaccountService
	SubAccounts    *subaccounts.Service
	Cdr            *cdr.Service
	Pricing        *pricing.Service
	Token          *token.Service
	Lookup         *lookup.Service
	Media          *media.Service
	Powerpack      *powerpack.Service
	MaskingSession *maskingsession.Service
	VerifySession  *verify.Service
	VerifyCallerID *verifycallerid.Service
	TollFreeVerification *tollfree.Service
	CallFeedback   *callfeedback.Service
	MultiPartyCall *multipartycall.Service

	// 10DLC / compliance
	Brand                   *brand.Service
	Campaign                *campaign.Service
	Profile                 *profile.Service
	ComplianceDocuments     *compliance.DocumentService
	ComplianceDocumentTypes *compliance.DocumentTypeService
	ComplianceRequirements  *compliance.RequirementService
	ComplianceApplications  *compliance.ApplicationService
	EndUsers                *enduserscompliance.Service
}

// NewClient creates and returns a fully initialised *Client.
// authID and authToken may be empty strings; in that case the constructor
// falls back to the VOBIZ_AUTH_ID / VOBIZ_AUTH_TOKEN environment variables.
func NewClient(authID, authToken string, opts *Options) (*Client, error) {
	if authID == "" {
		authID = os.Getenv("VOBIZ_AUTH_ID")
	}
	if authToken == "" {
		authToken = os.Getenv("VOBIZ_AUTH_TOKEN")
	}

	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return nil, fmt.Errorf("vobiz: invalid base URL: %w", err)
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	if opts != nil && opts.HTTPClient != nil {
		httpClient = opts.HTTPClient
	}

	base := &internalhttp.BaseClient{
		HTTPClient: httpClient,
		AuthID:     authID,
		AuthToken:  authToken,
		BaseURL:    baseURL,
		UserAgent:  fmt.Sprintf("vobiz-go/%s (Go: %s)", internalhttp.SDKVersion, runtime.Version()),
	}

	c := &Client{base: base}
	c.wire()
	return c, nil
}

// NewRequest is a convenience wrapper that prepends the account auth-ID to the
// path and delegates to BaseClient.NewRequest.
func (c *Client) NewRequest(method string, params interface{}, pathFmt string, pathArgs ...interface{}) (*http.Request, error) {
	args := append([]interface{}{c.base.AuthID}, pathArgs...)
	return c.base.NewRequest(method, params, baseRequestFmt, "%s/"+pathFmt, args...)
}

// ExecuteRequest delegates to BaseClient.ExecuteRequest.
func (c *Client) ExecuteRequest(req *http.Request, dest interface{}, extra ...map[string]interface{}) error {
	return c.base.ExecuteRequest(req, dest, extra...)
}

// AuthID returns the configured auth ID (useful for services that need it directly).
func (c *Client) AuthID() string { return c.base.AuthID }

// BaseURL returns the configured base URL.
func (c *Client) BaseURL() *url.URL { return c.base.BaseURL }

// wire initialises every service with a reference to this client.
func (c *Client) wire() {
	c.Calls          = calls.New(c)
	c.Messages       = messages.New(c)
	c.Recordings     = recordings.New(c)
	c.Transcription  = transcription.New(c)
	c.Conferences    = conferences.New(c)
	c.Endpoints      = endpoints.New(c)
	c.Numbers        = numbers.New(c)
	c.PhoneNumbers   = numbers.NewPhoneNumber(c)
	c.Trunks         = trunks.New(c)
	c.Credentials    = credentials.New(c)
	c.IpACL          = ipacl.New(c)
	c.OriginationURIs = originationuris.New(c)
	c.Applications   = applications.New(c)
	c.Accounts       = accounts.New(c)
	c.Subaccounts    = accounts.NewSubaccount(c)
	c.SubAccounts    = subaccounts.New(c)
	c.Cdr            = cdr.New(c)
	c.Pricing        = pricing.New(c)
	c.Token          = token.New(c)
	c.Lookup         = lookup.New(c)
	c.Media          = media.New(c)
	c.Powerpack      = powerpack.New(c)
	c.MaskingSession = maskingsession.New(c)
	c.VerifySession  = verify.New(c)
	c.VerifyCallerID = verifycallerid.New(c)
	c.TollFreeVerification = tollfree.New(c)
	c.CallFeedback   = callfeedback.New(c)
	c.MultiPartyCall = multipartycall.New(c)
	c.Brand          = brand.New(c)
	c.Campaign       = campaign.New(c)
	c.Profile        = profile.New(c)
	c.ComplianceDocuments     = compliance.NewDocument(c)
	c.ComplianceDocumentTypes = compliance.NewDocumentType(c)
	c.ComplianceRequirements  = compliance.NewRequirement(c)
	c.ComplianceApplications  = compliance.NewApplication(c)
	c.EndUsers                = enduserscompliance.New(c)
}
