# Vobiz Go SDK

The official Go SDK for [Vobiz](https://vobiz.ai) — an AI-first voice and telephony API platform for builders. Use it to make and control calls, manage SIP trunks, provision phone numbers, run conferences, and capture recordings, all from idiomatic, type-safe Go. This SDK provides a comprehensive, production-ready interface to the Vobiz API, enabling you to integrate powerful voice capabilities into your Go applications with ease.

## Quick links

- **Documentation:** [docs.vobiz.ai](https://docs.vobiz.ai)
- **Dashboard / Console:** [console.vobiz.ai](https://console.vobiz.ai)
- **Full API reference:** [`./reference.md`](./reference.md)

## Features

The Vobiz Go SDK provides access to the full suite of Vobiz API capabilities, including:

- **Programmatic Call Control:** Initiate outbound calls, manage live calls (hang up, transfer), and respond to inbound call events.
- **SIP Trunk Management:** Create, retrieve, update, and delete SIP trunks for flexible voice routing.
- **Phone Number Provisioning:** List available numbers, purchase new numbers, and manage number assignments to trunks or sub-accounts.
- **Conference Management:** Create and control multi-party conference calls, including muting, deafening, and kicking participants.
- **Call Recording:** Start and stop call recordings, retrieve recording details, and manage stored recordings.
- **Audio Manipulation & TTS:** Play audio files or convert text-to-speech (TTS) into live calls.
- **DTMF Tones:** Send Dual-Tone Multi-Frequency (DTMF) tones during active calls.
- **Sub-Account & KYC Management:** Programmatically manage sub-accounts, run PAN/GSTIN/DigiLocker KYC verifications, or test with mock KYC sessions.

## Requirements

The Vobiz Go SDK requires Go version **`1.21`** or later.

## Installation

To add the Vobiz Go SDK to your project, use `go get`:

```sh
go get github.com/vobiz-ai/Vobiz-Go-SDK
```

## Authentication

Vobiz authenticates requests using an **Auth ID** and an **Auth Token**. You can find your unique credentials in the [Vobiz Console](https://console.vobiz.ai).

To configure the client, use `option.WithAPIKey` for your Auth ID and `option.WithAuthToken` for your Auth Token. These options map internally to the `X-Auth-ID` and `X-Auth-Token` HTTP headers. For production environments, it is highly recommended to load your credentials from environment variables:

```go
package main

import (
    "os"

    "github.com/vobiz-ai/Vobiz-Go-SDK/client"
    "github.com/vobiz-ai/Vobiz-Go-SDK/option"
)

func main() {
    c := client.NewClient(
        option.WithAPIKey(os.Getenv("VOBIZ_AUTH_ID")),
        option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
    )
    
    // Your API calls here
}
```

## Quickstart

This quickstart demonstrates how to make an outbound call using the Vobiz Go SDK. When the call connects, Vobiz will fetch XML instructions from your specified `AnswerURL`.

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
        AuthID:       os.Getenv("VOBIZ_AUTH_ID"),
        From:         "14155551234",
        To:           "+919876543210",
        AnswerURL:    "https://example.com/answer",
        AnswerMethod: "POST",
    })
    if err != nil {
        log.Fatalf("Error making call: %v", err)
    }

    fmt.Printf("Call queued successfully: %+v\n", response)
}
```

## Common operations

### Retrieve account balance

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

    balanceResponse, err := c.Balance.GetBalance(context.TODO(), &vobiz.GetBalanceRequest{
        AuthID:   os.Getenv("VOBIZ_AUTH_ID"),
        Currency: "USD",
    })
    if err != nil {
        log.Fatalf("Error retrieving balance: %v", err)
    }

    fmt.Printf("Account Balance: %+v\n", balanceResponse)
}
```

### List Call Detail Records (CDRs) with filters

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

    request := &vobiz.ListCdrsRequest{
        AuthID:     os.Getenv("VOBIZ_AUTH_ID"),
        FromNumber: vobiz.String("9876543210"),
        ToNumber:   vobiz.String("1234567890"),
        StartDate: vobiz.Time(
            vobiz.MustParseDate("2026-03-01"),
        ),
        EndDate: vobiz.Time(
            vobiz.MustParseDate("2026-03-17"),
        ),
        MinDuration: vobiz.Int(10),
    }

    response, err := c.Cdr.ListCdrs(context.TODO(), request)
    if err != nil {
        log.Fatalf("Error listing CDRs: %v", err)
    }

    fmt.Printf("CDRs: %+v\n", response)
}
```

### Create a SIP Trunk

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

    request := &vobiz.CreateTrunkRequest{
        AuthID:             os.Getenv("VOBIZ_AUTH_ID"),
        Name:               "My Outbound Trunk",
        TrunkType:          "OUTBOUND",
        MaxConcurrentCalls: 10,
    }

    response, err := c.Trunks.CreateTrunk(context.TODO(), request)
    if err != nil {
        log.Fatalf("Error creating SIP trunk: %v", err)
    }

    fmt.Printf("Trunk created: %+v\n", response)
}
```

### Verify Sub-Account PAN (KYC)

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

    request := &vobiz.VerifySubaccountPanRequest{
        SubAuthID: "SA_XXXXXX",
        Pan:       "ABCDE1234F",
    }

    result, err := c.SubAccountKyc.VerifySubaccountPan(context.TODO(), request)
    if err != nil {
        log.Fatalf("Error verifying PAN: %v", err)
    }

    fmt.Printf("Verification Result: %+v\n", result)
}
```

## Configuration

The Vobiz Go SDK client is designed to be highly configurable. You can pass functional options to `client.NewClient()` to customize the client's behavior.

### Timeouts and Contexts

All API methods accept a standard Go `context.Context` as their first parameter. This allows you to easily manage request timeouts, cancellations, and deadlines within your application flow:

```go
package main

import (
    "context"
    "log"
    "os"
    "time"

    vobiz "github.com/vobiz-ai/Vobiz-Go-SDK"
    "github.com/vobiz-ai/Vobiz-Go-SDK/client"
    "github.com/vobiz-ai/Vobiz-Go-SDK/option"
)

func main() {
    c := client.NewClient(
        option.WithAPIKey(os.Getenv("VOBIZ_AUTH_ID")),
        option.WithAuthToken(os.Getenv("VOBIZ_AUTH_TOKEN")),
    )

    // Set a strict 5-second timeout for the API call
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := c.Balance.GetBalance(ctx, &vobiz.GetBalanceRequest{
        AuthID:   os.Getenv("VOBIZ_AUTH_ID"),
        Currency: "USD",
    })
    if err != nil {
        log.Fatalf("Request failed or timed out: %v", err)
    }
}
```

## Error handling

The Vobiz Go SDK returns standard Go `error` values. If the API returns a non-2xx HTTP status code, the SDK returns an error detailing the failure. Always check the returned `err` value before attempting to access the response object.

```go
response, err := c.Balance.GetBalance(context.TODO(), &vobiz.GetBalanceRequest{
    AuthID:   os.Getenv("VOBIZ_AUTH_ID"),
    Currency: "USD",
})
if err != nil {
    // Handle error (e.g., network timeout, invalid credentials, or API-side validation error)
    log.Printf("API call failed: %v", err)
    return
}
```

## Pagination

List endpoints (such as `Cdr.ListCdrs`, `Balance.ListTransactions`, `PhoneNumbers.ListNumbers`, etc.) support pagination via `limit`, `offset`, `page`, and `perPage` parameters.

```go
request := &vobiz.ListCdrsRequest{
    AuthID:  os.Getenv("VOBIZ_AUTH_ID"),
    Page:    vobiz.Int(2),
    PerPage: vobiz.Int(50),
}
```

## Other Vobiz SDKs

Vobiz provides official SDKs for several popular programming languages:

| Language | Repository |
| :--- | :--- |
| **TypeScript / Node.js** | [vobiz-ai/Vobiz-Node-SDK](https://github.com/vobiz-ai/Vobiz-Node-SDK) |
| **Python** | [vobiz-ai/Vobiz-Python-SDK](https://github.com/vobiz-ai/Vobiz-Python-SDK) |
| **Ruby** | [vobiz-ai/Vobiz-Ruby-SDK](https://github.com/vobiz-ai/Vobiz-Ruby-SDK) |
| **C# (.NET)** | [vobiz-ai/Vobiz-Csharp-sdk](https://github.com/vobiz-ai/Vobiz-Csharp-sdk) |
| **Java** | [vobiz-ai/Vobiz-Java-SDK](https://github.com/vobiz-ai/Vobiz-Java-SDK) |
| **PHP** | [vobiz-ai/Vobiz-PHP-SDK](https://github.com/vobiz-ai/Vobiz-PHP-SDK) |

## Support

- **Detailed Documentation:** [docs.vobiz.ai](https://docs.vobiz.ai)
- **Vobiz Console:** [console.vobiz.ai](https://console.vobiz.ai)

## License

This SDK is distributed under the [MIT License](LICENSE).
