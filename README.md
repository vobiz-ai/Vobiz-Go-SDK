# Vobiz Go Library

Typed Go client for the Vobiz programmable voice and SIP-trunking API, with a
`vobizxml` subpackage for building call-control documents.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/vobiz-ai/Vobiz-Go-SDK.svg)](https://pkg.go.dev/github.com/vobiz-ai/Vobiz-Go-SDK)
[![Go](https://img.shields.io/badge/go-1.21%2B-00ADD8.svg)](https://go.dev/)
[![Docs](https://img.shields.io/badge/docs-docs.vobiz.ai-3b82f6.svg)](https://docs.vobiz.ai)
[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=Vobiz%2FGo)

## Overview

The Vobiz Go library is the official client for the Vobiz REST API. It covers the
whole platform surface: placing and controlling calls, live-call inspection, in-call
actions such as text-to-speech, audio playback and DTMF, recordings, call detail
records, phone-number inventory, SIP trunks and endpoints, conferences,
applications, sub-accounts and KYC, IP access control lists, balance and the partner
API.

The client is generated from the Vobiz OpenAPI specification with
[Fern](https://buildwithfern.com), so every sub-client, method, request struct and
response struct tracks the published API. Configuration uses the functional-options
pattern — `option.WithToken`, `option.WithBaseURL`, `option.WithHTTPClient` and
friends — applied either once on the client or per call.

The module has a deliberately small dependency footprint: `github.com/google/uuid` at
runtime, plus `stretchr/testify` for the generated tests. Every method takes a
`context.Context` as its first argument, so cancellation and deadlines work exactly
the way you expect from the standard library.

Alongside the API client the module ships
[`vobizxml`](https://pkg.go.dev/github.com/vobiz-ai/Vobiz-Go-SDK/vobizxml), a
self-contained builder for VobizXML — the XML call-control documents Vobiz fetches
from your `answer_url` when a call connects. It mirrors the `plivoxml` builder shape
and emits XML byte-identical to the Python and Node builders.

At the end of a first integration you should be able to place an outbound call from
Go, serve a VobizXML document that speaks a prompt and collects a DTMF digit, watch
the call in the live-calls list, and read the resulting CDR.

## Installation

```sh
go get github.com/vobiz-ai/Vobiz-Go-SDK
```

Requires Go 1.21 or newer. The module has no tagged releases yet, so `go get`
resolves a pseudo-version pinned to the latest commit on the default branch. To pin
one explicitly:

```sh
go get github.com/vobiz-ai/Vobiz-Go-SDK@<commit-sha>
```

Three import paths are in play, and it is worth naming them all at once:

```go
import (
    vobiz "github.com/vobiz-ai/Vobiz-Go-SDK"                 // request/response types, errors
    client "github.com/vobiz-ai/Vobiz-Go-SDK/client"         // the root client
    option "github.com/vobiz-ai/Vobiz-Go-SDK/option"         // functional options
    vobizxml "github.com/vobiz-ai/Vobiz-Go-SDK/vobizxml"     // XML builder
)
```

## Authentication

Vobiz identifies your account with an **Auth ID** and an **Auth Token**, and
authorises the request with a **bearer token**. All three are supplied as options:

| Header | Option | Purpose |
| --- | --- | --- |
| `X-Auth-ID` | `option.WithAuthID` | Identifies the account or sub-account |
| `X-Auth-Token` | `option.WithAuthToken` | Account secret paired with the Auth ID |
| `Authorization: Bearer <token>` | `option.WithToken` | Bearer credential for the request |

```go
package main

import (
    "os"

    client "github.com/vobiz-ai/Vobiz-Go-SDK/client"
    option "github.com/vobiz-ai/Vobiz-Go-SDK/option"
)

func newVobizClient() *client.Client {
    return client.NewClient(
        option.WithToken(os.Getenv("VOBIZ_TOKEN")),
        option.WithAuthID(os.Getenv("VOBIZ_AUTH_ID")),
        option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
    )
}
```

For credentials that rotate, `option.WithTokenFunc` takes a
`func() (string, error)` that is resolved per request, so you can refresh a
short-lived token without rebuilding the client:

```go
c := client.NewClient(
    option.WithTokenFunc(func() (string, error) { return tokenStore.Current() }),
    option.WithAuthID(os.Getenv("VOBIZ_AUTH_ID")),
    option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
)
```

Any option can also be passed to an individual call, which overrides the client-level
value for that request only — that is how a parent account acts on behalf of a
sub-account:

```go
resp, err := c.Cdr.ListCdrs(ctx, req,
    option.WithAuthID(subAuthID),
    option.WithAuthToken(subAuthToken),
)
```

Note the separate `AuthID` field on request structs. That is the account the
operation acts on and it goes into the URL path, whereas `option.WithAuthID` sets the
`X-Auth-ID` header. They are usually the same value.

Keep credentials in environment variables or a secrets manager, never in source. Sign
up and find your credentials at [vobiz.ai](https://vobiz.ai); the credential model is
documented at [docs.vobiz.ai/api-reference](https://docs.vobiz.ai/api-reference).

## Quickstart

Place an outbound call. Vobiz dials `To`, and when the callee answers it fetches your
`AnswerURL` for a VobizXML document describing what should happen next.

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    vobiz "github.com/vobiz-ai/Vobiz-Go-SDK"
    client "github.com/vobiz-ai/Vobiz-Go-SDK/client"
    option "github.com/vobiz-ai/Vobiz-Go-SDK/option"
)

func main() {
    authID := os.Getenv("VOBIZ_AUTH_ID")

    c := client.NewClient(
        option.WithToken(os.Getenv("VOBIZ_TOKEN")),
        option.WithAuthID(authID),
        option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
        option.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
    )

    ctx := context.Background()

    resp, err := c.Calls.MakeCall(ctx, &vobiz.MakeCallRequest{
        AuthID:       authID,
        From:         "14155551234",
        To:           "+15550003333",
        AnswerURL:    "https://example.com/answer",
        AnswerMethod: "POST",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", resp)
}
```

Naming conventions worth internalising before you write the second call:

- **Struct fields are Go-idiomatic PascalCase with initialisms capitalised** —
  `AuthID`, `CallUUID`, `AnswerURL`, `RecordingID`, `SipCallID`, `PerPage`. The JSON
  and query-string names underneath stay snake_case, which the struct tags handle.
- **Client construction uses option functions**, not a config struct.
- **Optional fields are pointers.** Use the helpers in the root package —
  `vobiz.Int(30)`, `vobiz.String("auto")`, `vobiz.Bool(true)` — or an enum's `.Ptr()`
  method.
- `To` accepts multiple destinations separated by `<`, fanning a single request out
  to up to 1000 destinations, for example `"+15550003333<+15550004444"`.

## Common operations

Every snippet below reuses `c`, `ctx` and `authID` from the quickstart. Signatures
come from the generated sub-clients in this module; the exhaustive list is in
[`reference.md`](./reference.md).

### List live calls

`Status` is required. Use the generated constants rather than raw strings.

```go
live, err := c.LiveCalls.ListLiveCalls(ctx, &vobiz.ListLiveCallsRequest{
    AuthID: authID,
    Status: vobiz.ListLiveCallsRequestStatusLive,   // or ...StatusQueued
})

detail, err := c.LiveCalls.GetLiveCall(ctx, &vobiz.GetLiveCallRequest{
    AuthID:   authID,
    CallUUID: "cdr_XXXXXXXXXX",
    Status:   vobiz.GetLiveCallRequestStatusLive,
})
```

`ListQueuedCalls` and `GetQueuedCall` mirror these for the queued set.

### Hang up a call

```go
err := c.LiveCalls.HangupCall(ctx, &vobiz.HangupCallRequest{
    AuthID:   authID,
    CallUUID: "cdr_XXXXXXXXXX",
})
```

### Speak text and play audio into a live call

```go
err := c.SpeakText.Call(ctx, &vobiz.SpeakTextCallRequest{
    AuthID:   authID,
    CallUUID: "cdr_XXXXXXXXXX",
    Text:     "Your driver is two minutes away.",
    Legs:     vobiz.SpeakTextCallRequestLegsAleg.Ptr(),
})

err = c.PlayAudio.Call(ctx, &vobiz.PlayAudioCallRequest{
    AuthID:   authID,
    CallUUID: "cdr_XXXXXXXXXX",
    URLs:     "https://cdn.example.com/hold-music.mp3",
    Loop:     vobiz.Bool(true),
})
```

`SpeakText.StopSpeakCall` and `PlayAudio.StopAudioCall` take just `AuthID` and
`CallUUID`. In-call action methods return only an `error` — there is no response
body to unwrap.

### Send DTMF

```go
err := c.Dtmf.SendDtmf(ctx, &vobiz.SendDtmfRequest{
    AuthID:   authID,
    CallUUID: "cdr_XXXXXXXXXX",
    Digits:   "1234#",
    Leg:      vobiz.SendDtmfRequestLegAleg.Ptr(),   // Aleg | Bleg | Both
})
```

### Record a call and fetch the recording

```go
resp, err := c.RecordCalls.StartRecording(ctx, &vobiz.StartRecordingRequest{
    AuthID:            authID,
    CallUUID:          "cdr_XXXXXXXXXX",
    FileFormat:        vobiz.StartRecordingRequestFileFormatMp3.Ptr(),   // Mp3 | Wav
    TimeLimit:         vobiz.Int(600),
    TranscriptionType: vobiz.String("auto"),
    CallbackURL:       vobiz.String("https://example.com/recording-ready"),
})

err = c.RecordCalls.StopRecording(ctx, &vobiz.StopRecordingRequest{
    AuthID:   authID,
    CallUUID: "cdr_XXXXXXXXXX",
})

recs, err := c.Recordings.ListRecordings(ctx, &vobiz.ListRecordingsRequest{
    AuthID: authID,
    Limit:  vobiz.Int(20),
    Offset: vobiz.Int(0),
})
```

`StartRecording` returns `(any, error)` because the response shape is not modelled;
`StopRecording` returns only an `error`. `Recordings.GetRecording` and
`Recordings.DeleteRecording` take `AuthID` and `RecordingID`.

### Query call detail records

`StartDate` and `EndDate` are `*time.Time` values serialised as `YYYY-MM-DD`, and
each is required when the other is set.

```go
start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC)

page, err := c.Cdr.ListCdrs(ctx, &vobiz.ListCdrsRequest{
    AuthID:        authID,
    StartDate:     &start,
    EndDate:       &end,
    CallDirection: vobiz.ListCdrsRequestCallDirectionOutbound.Ptr(),
    MinDuration:   vobiz.Int(30),
    Page:          vobiz.Int(1),
    PerPage:       vobiz.Int(100),   // max 100
})

recent, err := c.Cdr.ListRecentCdrs(ctx, &vobiz.ListRecentCdrsRequest{
    AuthID: authID,
    Limit:  vobiz.Int(25),
})

one, err := c.Cdr.GetCdr(ctx, &vobiz.GetCdrRequest{
    AuthID: authID,
    CallID: "CALL_ID",
})
```

`SearchCdrs` takes the same filters as `ListCdrs`, and `ExportCdrs` takes the same
filters without the paging fields. Other filters on all three: `FromNumber`,
`ToNumber`, `SipCallID`, `BridgeUUID`, `HangupCause`, `HangupDisposition`, `Context`,
`CampaignID` and free-text `Search`.

### Other resource groups

The same pattern — `c.<Group>.<Method>(ctx, &vobiz.<Method>Request{...})` — covers
`Account`, `Balance`, `PhoneNumbers`, `Applications`, `Trunks`, `Endpoints`,
`Credentials`, `Conference`, `Conferences`, `ConferenceMembers`,
`ConferenceRecording`, `AudioStreams`, `SubAccounts`, `SubAccountKyc`,
`SubAccountKycTestMode`, `BulkOperations`, `IPAccessControlList`, `OriginationURI`
and `PartnerAPI`. See [`reference.md`](./reference.md) for every signature.

## VobizXML

The `vobizxml` subpackage builds the call-control documents Vobiz fetches from your
`answer_url`. It imports nothing beyond `fmt` and `strings`, so you can use it in an
HTTP handler without constructing an API client.

```go
package main

import (
    "fmt"

    "github.com/vobiz-ai/Vobiz-Go-SDK/vobizxml"
)

func main() {
    r := vobizxml.NewResponse()

    g := r.AddGather(
        vobizxml.Attr("action", "https://example.com/menu"),
        vobizxml.Attr("method", "POST"),
        vobizxml.Attr("inputType", "dtmf"),
        vobizxml.Attr("numDigits", 1),
        vobizxml.Attr("executionTimeout", 10),
    )
    g.AddSpeak("Press 1 for sales, or 2 for support.")

    r.AddSpeak("We did not receive any input. Goodbye.")
    r.AddHangup()

    fmt.Println(r.String())
}
```

That prints:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Gather action="https://example.com/menu" method="POST" inputType="dtmf" numDigits="1" executionTimeout="10">
        <Speak>Press 1 for sales, or 2 for support.</Speak>
    </Gather>
    <Speak>We did not receive any input. Goodbye.</Speak>
    <Hangup/>
</Response>
```

Points worth knowing:

- **Attributes are set with `vobizxml.Attr(name, value)`**, using the camelCase
  VobizXML name verbatim — `inputType`, `executionTimeout`, `numDigits`, `callerId`,
  `startConferenceOnEnter`, `sendDigits`, `audioTrack`. `<Gather>` uses
  `executionTimeout`, never `timeout`.
- **Attribute order is preserved.** Attributes are held in an ordered slice rather
  than a map, so the rendered order matches the order you set them.
- **Values are `any`.** Booleans render as `true`/`false`, everything else goes
  through `fmt.Sprint`, and text plus attribute values are XML escaped for you.
- **`Add*` helpers return the child element**, so you can keep nesting:
  `r.AddDial().AddNumber("+15550003333")`.
- **`AddSpeakSSML` injects raw, unescaped content** into `<Speak>` when you need SSML
  markup.
- **`String()` renders pretty-printed with the XML declaration; `StringCompact()`
  renders it on a single line.**

Builder types: `Response`, `Speak`, `Play`, `Wait`, `Gather`, `Dial`, `Number`,
`User`, `Record`, `Conference`, `DTMF`, `Redirect`, `Hangup`, `PreAnswer` and
`Stream`.

Migrating from Plivo? `AddGetDigits()` and `AddGetInput()` are kept as aliases for
`AddGather()`.

Serving it from `net/http`:

```go
http.HandleFunc("/answer", func(w http.ResponseWriter, req *http.Request) {
    r := vobizxml.NewResponse()
    r.AddSpeak("Thanks for calling. Connecting you now.")
    r.AddDial().AddNumber("+15550003333")

    w.Header().Set("Content-Type", "application/xml")
    fmt.Fprint(w, r.String())
})
```

## Configuration

### Environments and base URL

The client targets production by default. Use `option.WithBaseURL` to override it —
useful for a proxy, a gateway, or a local mock:

```go
c := client.NewClient(
    option.WithBaseURL(vobiz.Environments.Production),   // https://api.vobiz.ai
)

test := client.NewClient(
    option.WithBaseURL("http://localhost:8080"),
)
```

### Request options

Every option can be applied to the whole client or to a single call. Per-call options
win.

| Option | Purpose |
| --- | --- |
| `option.WithToken(string)` | Bearer credential |
| `option.WithTokenFunc(func() (string, error))` | Bearer credential resolved per request |
| `option.WithAuthID(string)` | `X-Auth-ID` header |
| `option.WithAuthToken(string)` | `X-Auth-Token` header |
| `option.WithBaseURL(string)` | Override the API base URL |
| `option.WithHTTPClient(core.HTTPClient)` | Supply your own `*http.Client` |
| `option.WithHTTPHeader(http.Header)` | Extra headers on the request |
| `option.WithQueryParameters(url.Values)` | Extra query string parameters |
| `option.WithBodyProperties(map[string]interface{})` | Extra JSON body properties |
| `option.WithMaxAttempts(uint)` | Retry limit (default 2) |
| `option.WithoutRetries()` | Disable retries entirely |
| `option.WithMaxStreamBufSize(int)` | Buffer size for streaming responses |
| `option.WithMaxStreamReconnectAttempts(uint)` | Reconnect limit for streams |
| `option.WithoutStreamReconnection()` | Disable stream reconnection |

### HTTP client and timeouts

> Providing your own `*http.Client` is recommended. Otherwise `http.DefaultClient` is
> used, and your client will wait indefinitely for a response unless the
> per-request, context-based timeout is used.

```go
c := client.NewClient(
    option.WithToken(os.Getenv("VOBIZ_TOKEN")),
    option.WithAuthID(os.Getenv("VOBIZ_AUTH_ID")),
    option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
    option.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
)
```

Per-call deadlines use the standard context library:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resp, err := c.Cdr.ListCdrs(ctx, req)
```

### Retries

Requests are retried with exponential backoff on **408**, **429** and all **5xx**
responses, twice by default. If a `Retry-After` header is present the SDK honours its
value exactly instead of the backoff schedule.

```go
c := client.NewClient(option.WithMaxAttempts(1))

resp, err := c.Calls.MakeCall(ctx, req, option.WithMaxAttempts(1))
```

`option.WithoutRetries()` disables retries entirely — worth using for non-idempotent
operations where a duplicate would be worse than a failure.

### Explicit null

Optional fields are pointers and are omitted when `nil`. To send an explicit JSON
`null`, use the generated setter, which flips a bit in the request's private
`explicitFields` bitmask so the field is serialised rather than omitted:

```go
req := &vobiz.StartRecordingRequest{AuthID: authID, CallUUID: callUUID}
req.SetCallbackURL(nil)   // serialises "callback_url": null
```

## Error handling

Non-2xx responses return structured error types that embed `*core.APIError`, so they
work with `errors.Is` and `errors.As`:

```go
import (
    "errors"

    vobiz "github.com/vobiz-ai/Vobiz-Go-SDK"
    core "github.com/vobiz-ai/Vobiz-Go-SDK/core"
)

resp, err := c.Calls.MakeCall(ctx, req)
if err != nil {
    var apiErr *core.APIError
    if errors.As(err, &apiErr) {
        log.Printf("status=%d headers=%v", apiErr.StatusCode, apiErr.Header)
    }
    return err
}
```

Status-specific types let you branch without inspecting the status code:

| Type | Status | Typical cause |
| --- | --- | --- |
| `vobiz.BadRequestError` | 400 | Malformed or missing parameters |
| `vobiz.UnauthorizedError` | 401 | Wrong Auth ID, Auth Token or bearer token |
| `vobiz.ForbiddenError` | 403 | Credentials valid, operation not permitted |
| `vobiz.NotFoundError` | 404 | Unknown call UUID, recording ID or account |
| `vobiz.ConflictError` | 409 | Resource already exists or is in use |
| `vobiz.UnprocessableEntityError` | 422 | Understood but semantically invalid |
| `vobiz.TooManyRequestsError` | 429 | Rate limited |
| `vobiz.InternalServerError` | 500 | Server-side failure |

```go
var notFound *vobiz.NotFoundError
var rateLimited *vobiz.TooManyRequestsError

switch {
case errors.As(err, &notFound):
    // the call has already ended
case errors.As(err, &rateLimited):
    // back off; inspect rateLimited.Header for retry hints
}
```

Each type carries a `Body any` field holding the decoded response payload, and
`Unwrap()` returns the embedded `*core.APIError`.

### Response headers

Access the raw HTTP response through the `WithRawResponse` field on any sub-client —
useful for reading rate-limit headers:

```go
raw, err := c.Cdr.WithRawResponse.ListCdrs(ctx, req)
if err != nil {
    return err
}
fmt.Printf("headers: %v status: %d\n", raw.Header, raw.StatusCode)
fmt.Printf("body: %+v\n", raw.Body)
```

## Pagination and concurrency

Listing methods paginate explicitly; there is no auto-paging iterator, so you drive
the loop yourself. Two conventions are in use:

- **`Page` / `PerPage`** — `Cdr.ListCdrs`, `Cdr.SearchCdrs` (`PerPage` max 100)
- **`Limit` / `Offset`** — `Recordings.ListRecordings`; `Cdr.ListRecentCdrs` takes
  `Limit` only

```go
for page := 1; ; page++ {
    result, err := c.Cdr.ListCdrs(ctx, &vobiz.ListCdrsRequest{
        AuthID:  authID,
        Page:    vobiz.Int(page),
        PerPage: vobiz.Int(100),
    })
    if err != nil {
        return err
    }
    if len(result.Data) == 0 {
        break
    }
    // ... handle result.Data
}
```

The client is safe to share across goroutines — build one and reuse it, so the
underlying `*http.Client` can pool connections. Bound your fan-out with a semaphore
or an `errgroup` limit so you do not trip the rate limit:

```go
g, gctx := errgroup.WithContext(ctx)
g.SetLimit(8)
for _, uuid := range callUUIDs {
    uuid := uuid
    g.Go(func() error {
        return c.LiveCalls.HangupCall(gctx, &vobiz.HangupCallRequest{
            AuthID: authID, CallUUID: uuid,
        })
    })
}
err := g.Wait()
```

## Troubleshooting

| Symptom | Likely cause | Fix |
| --- | --- | --- |
| `cannot use "aleg" (untyped string constant) as *vobiz.SendDtmfRequestLeg value` | Optional enum fields are pointers to a named type, not strings | Use `vobiz.SendDtmfRequestLegAleg.Ptr()` |
| `cannot use 30 (untyped int constant) as *int value` | Optional scalar fields are pointers | Wrap with the root-package helpers: `vobiz.Int(30)`, `vobiz.String("auto")`, `vobiz.Bool(true)` |
| `undefined: vobiz.Client` | The root package holds types; the client lives in the `client` subpackage | Import `github.com/vobiz-ai/Vobiz-Go-SDK/client` and call `client.NewClient(...)` |
| A request hangs forever with no error | `http.DefaultClient` has no timeout | Pass `option.WithHTTPClient(&http.Client{Timeout: 30 * time.Second})`, or use a `context.WithTimeout` |
| `401 Unauthorized` on every call | Auth ID, Auth Token or bearer token is wrong, or points at a different environment | Re-check all three; confirm `option.WithBaseURL` targets the intended host |
| `404 Not Found` from `GetLiveCall` or `HangupCall` | The call has already ended, so it is no longer in the live-call set | Treat 404 as "already finished"; look it up with `c.Cdr.GetCdr(...)` instead |
| `429 Too Many Requests` during a bulk loop | Requests are issued faster than the account's rate allowance | Bound goroutine fan-out; the SDK already retries twice and honours `Retry-After` |
| An optional field you set to the zero value never reaches the API | Zero and `nil` values are omitted during serialisation | Use the generated setter, e.g. `req.SetTimeLimit(0)`, to force the field to be sent |
| `go get` picks an unexpected pseudo-version | The module has no semantic-version tags yet | Pin a commit: `go get github.com/vobiz-ai/Vobiz-Go-SDK@<commit-sha>` |
| `<Gather>` never fires the `action` callback | `timeout` was used instead of `executionTimeout` | Pass `vobizxml.Attr("executionTimeout", 10)` |
| VobizXML renders as escaped text in the browser | The response was served as `text/html` | Set `Content-Type: application/xml` before writing `r.String()` |

## Other Vobiz SDKs

| Language | Repository | Package name |
| --- | --- | --- |
| Python | [Vobiz-Python-SDK](https://github.com/vobiz-ai/Vobiz-Python-SDK) | `vobiz` |
| Node.js / TypeScript | [Vobiz-Node-SDK](https://github.com/vobiz-ai/Vobiz-Node-SDK) | `@vobiz/sdk` |
| Ruby | [Vobiz-Ruby-SDK](https://github.com/vobiz-ai/Vobiz-Ruby-SDK) | `vobiz` |
| C# / .NET | [Vobiz-Csharp-sdk](https://github.com/vobiz-ai/Vobiz-Csharp-sdk) | `Vobiz` |

All of them are generated from the same OpenAPI specification, so resource groups and
method names line up across languages once you allow for naming conventions.

## Versioning and stability

The module has no semantic-version tags yet, so `go get` resolves a pseudo-version
from the default branch. Pin an exact commit in production and review the diff before
upgrading.

The API surface is regenerated from the Vobiz OpenAPI specification, so sub-client and
method names can change when the specification changes. `vobizxml` is hand-written and
follows the `plivoxml` shape; it is the more stable half of the module.

## Roadmap

> Planned improvements to this library. Ideas and pull requests are welcome —
> open an issue to discuss anything here.

- [ ] Publish semantic-version tags (`v0.1.0` onward) so `go get` resolves a real
      release rather than a pseudo-version, and Go module proxies can cache it.
- [ ] Adopt semantic versioning guarantees from `v1.0.0`, with a documented
      deprecation window for generated method renames.
- [ ] Auto-paging iterators for `Cdr.ListCdrs` and `Recordings.ListRecordings`, so
      callers stop hand-rolling `for` loops.
- [ ] Surface the Vobiz error code and message as typed fields on the error structs
      rather than an untyped `Body any`.
- [ ] Webhook signature verification helpers, so `answer_url` and callback handlers
      can validate that a request genuinely came from Vobiz.
- [ ] Runnable examples under `examples/` covering outbound calling, an IVR answer
      handler and CDR export, so `go run` gets you a working call.
- [ ] Extend test coverage to the `vobizxml` builder alongside the generated
      request/response tests.

## Contributing

While we value open-source contributions to this SDK, this library is generated
programmatically. Additions made directly to this library would have to be moved over
to our generation code, otherwise they would be overwritten upon the next generated
release. Feel free to open a PR as a proof of concept, but know that we will not be
able to merge it as-is. We suggest opening an issue first to discuss with us!

On the other hand, contributions to the README and to the hand-written
[`vobizxml`](./vobizxml) package are always very welcome. See
[CONTRIBUTING.md](./CONTRIBUTING.md) for details.

To check your changes locally:

```sh
go build ./...
go test ./...
go vet ./...
```

## License

Released under the [MIT License](./LICENSE) © Vobiz.

MIT is permissive: you may use, modify, and redistribute this code, including in
closed-source commercial products, provided the copyright notice and licence text
are retained. There is no warranty. If your organisation needs a different
licensing arrangement, contact [piyush@vobiz.ai](mailto:piyush@vobiz.ai).

## Built by Team Vobiz

[Vobiz](https://vobiz.ai) is a programmable voice and SIP-trunking platform for
voice APIs, SIP trunking, and AI voice agents. This repository is built and
maintained by the Vobiz team.

**Maintainer:** Piyush Sahoo — [piyush@vobiz.ai](mailto:piyush@vobiz.ai) · [LinkedIn](https://www.linkedin.com/in/piyush-s713/)

Questions, or want to talk through an integration? Open an issue on this repo,
or reach out directly at [piyush@vobiz.ai](mailto:piyush@vobiz.ai).

**Useful links:** [Docs](https://docs.vobiz.ai) · [API reference](https://docs.vobiz.ai/api-reference) · [Sign up](https://vobiz.ai)
