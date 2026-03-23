# Vobiz Go SDK

Official Go SDK for the [Vobiz](https://vobiz.ai) Voice & Telephony API.

Make outbound calls, manage SIP trunks, phone numbers, conferences, recordings, and build dynamic call flows with VobizXML — all from Go.

---

## Installation

```bash
go get github.com/vobiz-ai/vobiz-go-sdk
```

---

## Authentication

All requests require your **Auth ID** and **Auth Token** from the [Vobiz Console](https://console.vobiz.ai).

```go
import "github.com/vobiz-ai/vobiz-go-sdk/client"

c := client.New("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN")
```

Or set environment variables:

```bash
export VOBIZ_AUTH_ID=your_auth_id
export VOBIZ_AUTH_TOKEN=your_auth_token
```

---

## Quick Start

### Make an Outbound Call

```go
package main

import (
    "fmt"
    "github.com/vobiz-ai/vobiz-go-sdk/client"
    "github.com/vobiz-ai/vobiz-go-sdk/services/calls"
)

func main() {
    c := client.New("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN")

    call, err := calls.New(c).Create(&calls.CreateParams{
        From:         "+1xxxxxxxxxx",
        To:           "+1xxxxxxxxxx",
        AnswerURL:    "https://yourserver.com/answer",
        AnswerMethod: "POST",
        HangupURL:    "https://yourserver.com/hangup",
        HangupMethod: "POST",
    })
    if err != nil {
        panic(err)
    }
    fmt.Println("Call UUID:", call.RequestUUID)
}
```

---

## Features

### Calls

```go
svc := calls.New(c)

// List live calls
liveCalls, _ := svc.List(&calls.ListParams{Status: "live"})

// Get a specific call
call, _ := svc.Get("call-uuid")

// Speak text on a live call
svc.Speak("call-uuid", &calls.SpeakParams{
    Text:     "Hello, this is a test.",
    Voice:    "WOMAN",
    Language: "en-US",
    Legs:     "aleg",
})

// Play audio on a live call
svc.Play("call-uuid", &calls.PlayParams{
    URLs: []string{"https://example.com/audio.mp3"},
    Legs: "aleg",
})

// Start recording
svc.Record("call-uuid", &calls.RecordParams{
    TimeLimit:  60,
    FileFormat: "mp3",
})

// Send DTMF digits
svc.DTMF("call-uuid", &calls.DTMFParams{
    Digits: "1234",
    Leg:    "aleg",
})

// Transfer call
svc.Transfer("call-uuid", &calls.TransferParams{
    Legs:        "aleg",
    ALegURL:     "https://yourserver.com/transfer",
    ALegMethod:  "POST",
})

// Hang up
svc.Delete("call-uuid")
```

### Conferences

```go
import "github.com/vobiz-ai/vobiz-go-sdk/services/conferences"

svc := conferences.New(c)

// List conferences
list, _ := svc.List()

// Kick a member
svc.Member.Delete("conference-name", "member-id")

// Mute a member
svc.Member.Mute("conference-name", "member-id")

// Record a conference
svc.Record("conference-name", &conferences.RecordParams{
    FileFormat: "mp3",
})
```

### Recordings

```go
import "github.com/vobiz-ai/vobiz-go-sdk/services/recordings"

svc := recordings.New(c)

// List recordings
list, _ := svc.List(nil)

// Get a recording
rec, _ := svc.Get("recording-uuid")

// Delete a recording
svc.Delete("recording-uuid")
```

### Account

```go
import "github.com/vobiz-ai/vobiz-go-sdk/services/accounts"

svc := accounts.New(c)

// Get account details
profile, _ := svc.Get()

// Get balance
balance, _ := svc.Balance()
```

### Phone Numbers

```go
import "github.com/vobiz-ai/vobiz-go-sdk/services/numbers"

svc := numbers.New(c)

// List available numbers
available, _ := svc.Search(&numbers.SearchParams{
    CountryISO: "US",
    Type:       "local",
})

// Purchase a number
purchased, _ := svc.Purchase(&numbers.PurchaseParams{
    Number: "+1xxxxxxxxxx",
})
```

### Applications

```go
import "github.com/vobiz-ai/vobiz-go-sdk/services/applications"

svc := applications.New(c)

// Create an application
app, _ := svc.Create(&applications.CreateParams{
    AppName:      "My App",
    AnswerURL:    "https://yourserver.com/answer",
    HangupURL:    "https://yourserver.com/hangup",
})

// List applications
list, _ := svc.List(nil)
```

---

## VobizXML

Build dynamic call flow responses using the XML builder:

```go
import "github.com/vobiz-ai/vobiz-go-sdk/xml"

response := xml.NewResponse()
response.AddSpeak("Welcome to Vobiz!", "WOMAN", "en-US")
response.AddRecord(xml.RecordParams{
    Action:     "https://yourserver.com/recording",
    MaxLength:  30,
    PlayBeep:   true,
})

fmt.Println(response.ToXML())
```

Output:
```xml
<Response>
  <Speak voice="WOMAN" language="en-US">Welcome to Vobiz!</Speak>
  <Record action="https://yourserver.com/recording" maxLength="30" playBeep="true"/>
</Response>
```

---

## Sub-Accounts

```go
import "github.com/vobiz-ai/vobiz-go-sdk/services/accounts"

svc := accounts.New(c)

// Create a sub-account
sub, _ := svc.Subaccounts.Create(&accounts.SubaccountParams{
    Name: "Sub Account 1",
})

// List sub-accounts
list, _ := svc.Subaccounts.List()
```

---

## Error Handling

```go
call, err := calls.New(c).Create(&calls.CreateParams{...})
if err != nil {
    // err contains HTTP status and API error message
    fmt.Println("Error:", err)
}
```

---

## Requirements

- Go 1.21+

---

## Support

- Docs: [https://docs.vobiz.ai](https://docs.vobiz.ai)
- Console: [https://console.vobiz.ai](https://console.vobiz.ai)
- Email: support@vobiz.ai
