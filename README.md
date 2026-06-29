# Vobiz Go SDK

The official Go SDK for [Vobiz](https://vobiz.ai) — an AI-first voice and telephony API platform for builders. Use it to make and control calls, manage SIP trunks, provision phone numbers, run conferences, and capture recordings, all from idiomatic Go.

## Quick links

- **Documentation:** https://docs.vobiz.ai
- **Dashboard:** https://console.vobiz.ai
- **Full API reference:** [`./reference.md`](./reference.md)

## Installation

```sh
go get github.com/vobiz-ai/Vobiz-Go-SDK
```

Requires Go 1.21 or later.

## Authentication

Vobiz authenticates requests with an **Auth ID** and an **Auth Token** (sent as the `X-Auth-ID` and `X-Auth-Token` headers). You can find both in the [Vobiz Console](https://console.vobiz.ai).

Configure the client with `option.WithAPIKey` (your Auth ID) and `option.WithAuthToken` (your Auth Token):

```go
client := client.NewClient(
    option.WithAPIKey("<AUTH_ID>"),
    option.WithAuthToken("<AUTH_TOKEN>"),
)
```

We recommend loading credentials from environment variables rather than hardcoding them:

```go
client := client.NewClient(
    option.WithAPIKey(os.Getenv("VOBIZ_AUTH_ID")),
    option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
)
```

## Quickstart

Place an outbound call. When the call connects, Vobiz fetches your XML instructions from `AnswerURL`:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    vobiz "github.com/vobiz-ai/Vobiz-Go-SDK"
    "github.com/vobiz-ai/Vobiz-Go-SDK/client"
    "github.com/vobiz-ai/Vobiz-Go-SDK/option"
)

func main() {
    c := client.NewClient(
        option.WithAPIKey(os.Getenv("VOBIZ_AUTH_ID")),
        option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
    )

    response, err := c.Calls.MakeCall(context.TODO(), &vobiz.MakeCallRequest{
        AuthID:       "MA_XXXXXX",
        From:         "14155551234",
        To:           "+919876543210",
        AnswerURL:    "https://example.com/answer",
        AnswerMethod: "POST",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Call queued: %+v\n", response)
}
```

See [`./reference.md`](./reference.md) for the complete list of resources and methods (Account, Balance, Calls, and more).

## Error handling

API calls that return a non-success status code produce structured errors compatible with `errors.Is` and `errors.As`:

```go
response, err := client.Calls.MakeCall(ctx, request)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, &apiError) {
        // Inspect apiError ...
    }
    return err
}
```

## Request options

Options adapt the client's behavior. Pass them to `NewClient` to apply on every request, or to an individual call to override per request.

```go
// Applied to every request.
client := client.NewClient(
    option.WithAPIKey(os.Getenv("VOBIZ_AUTH_ID")),
    option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
    option.WithHTTPClient(&http.Client{Timeout: 5 * time.Second}),
)

// Applied to a single request.
response, err := client.Calls.MakeCall(ctx, request, option.WithMaxAttempts(1))
```

The SDK retries failed requests automatically with exponential backoff (default: 2 attempts) and supports per-request timeouts via the standard `context` library. Point the client at a different environment with `option.WithBaseURL`.

## Other SDKs

Vobiz publishes idiomatic SDKs for every major language:

| Language   | Repository |
|------------|------------|
| TypeScript | https://github.com/vobiz-ai/Vobiz-Node-SDK |
| Python     | https://github.com/vobiz-ai/Vobiz-Python-SDK |
| Java       | https://github.com/vobiz-ai/Vobiz-Java-SDK |
| Ruby       | https://github.com/vobiz-ai/Vobiz-Ruby-SDK |
| C#         | https://github.com/vobiz-ai/Vobiz-Csharp-sdk |
| PHP        | https://github.com/vobiz-ai/Vobiz-PHP-SDK |

## License

Released under the [MIT License](./LICENSE).
