# Reference
## Account
<details><summary><code>client.Account.RetrieveAccount() -> *vobiz.RetrieveAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve complete account details including pricing tier and credentials.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Account.RetrieveAccount(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Account.GetConcurrency(AuthID) -> *vobiz.GetConcurrencyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the current concurrent call usage and configured limits.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetConcurrencyRequest{
        AuthID: "MA_XXXXXX",
    }
client.Account.GetConcurrency(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Account.PreviewChannelPricing(AuthID) -> *vobiz.ChannelPricingPreview</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Calculate the monthly price for CPS or concurrent-call capacity without purchasing capacity or debiting the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.PreviewChannelPricingRequest{
        AuthID: "MA_XXXX",
        ResourceType: vobiz.CapacityResourceTypeConcurrentCalls,
        Quantity: 30,
    }
client.Account.PreviewChannelPricing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Target account Auth ID. An account can preview only its own pricing; administrators may act for another account.
    
</dd>
</dl>

<dl>
<dd>

**resourceType:** `*vobiz.CapacityResourceType` — Capacity type to price.
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `int` — Capacity quantity to price. Pricing-tier block and quantity rules also apply.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Account.CreateChannelSubscription(AuthID, request) -> *vobiz.ChannelSubscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Purchase recurring CPS or concurrent-call capacity. A successful request immediately debits the first monthly charge and activates a subscription that renews every 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ChannelSubscriptionRequest{
        AuthID: "MA_XXXX",
        ResourceType: vobiz.CapacityResourceTypeConcurrentCalls,
        Quantity: 30,
    }
client.Account.CreateChannelSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Target account Auth ID. An account can purchase only for itself; administrators may act for another account.
    
</dd>
</dl>

<dl>
<dd>

**resourceType:** `*vobiz.CapacityResourceType` 
    
</dd>
</dl>

<dl>
<dd>

**quantity:** `int` — Capacity quantity to purchase. Pricing-tier block and quantity rules also apply.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Balance
<details><summary><code>client.Balance.GetBalance(AuthID, Currency) -> *vobiz.GetBalanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the current account balance for a specific currency.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetBalanceRequest{
        AuthID: "MA_XXXXXX",
        Currency: "INR",
    }
client.Balance.GetBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` — Currency code (e.g. INR, USD)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Balance.ListTransactions(AuthID) -> *vobiz.ListTransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve paginated transaction history for the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListTransactionsRequest{
        AuthID: "MA_XXXXXX",
    }
client.Balance.ListTransactions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Live Calls
<details><summary><code>client.LiveCalls.ListQueuedCalls(AuthID) -> *vobiz.ListQueuedCallsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all queued (pending, not yet connected) calls on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListQueuedCallsRequest{
        AuthID: "MA_XXXXXX",
        Status: vobiz.ListQueuedCallsRequestStatusLive,
    }
client.LiveCalls.ListQueuedCalls(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vobiz.ListQueuedCallsRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LiveCalls.ListLiveCalls(AuthID) -> *vobiz.ListLiveCallsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all currently active (live) calls on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListLiveCallsRequest{
        AuthID: "MA_XXXXXX",
        Status: vobiz.ListLiveCallsRequestStatusLive,
    }
client.LiveCalls.ListLiveCalls(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vobiz.ListLiveCallsRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LiveCalls.GetLiveCall(AuthID, CallUUID) -> *vobiz.GetLiveCallResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of a specific live or queued call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetLiveCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "cdr_XXXXXXXXXX",
        Status: vobiz.GetLiveCallRequestStatusLive,
    }
client.LiveCalls.GetLiveCall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vobiz.GetLiveCallRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LiveCalls.HangupCall(AuthID, CallUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminate an active call by its UUID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.HangupCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
    }
client.LiveCalls.HangupCall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LiveCalls.GetQueuedCall(AuthID, CallUUID) -> *vobiz.GetQueuedCallResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of a specific queued (pending) call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetQueuedCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "cdr_XXXXXXXXXX",
        Status: vobiz.GetQueuedCallRequestStatusLive,
    }
client.LiveCalls.GetQueuedCall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vobiz.GetQueuedCallRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Calls
<details><summary><code>client.Calls.MakeCall(AuthID, request) -> *vobiz.MakeCallResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Initiate an outbound call to a PSTN number or SIP endpoint.
Use `<` to separate multiple destinations (max 1000).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MakeCallRequest{
        AuthID: "MA_XXXXXX",
        From: "14155551234",
        To: "+919876543210",
        AnswerURL: "https://example.com/answer",
        AnswerMethod: "POST",
    }
client.Calls.MakeCall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**from:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**answerURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**answerMethod:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CDR
<details><summary><code>client.Cdr.ListCdrs(AuthID) -> *vobiz.ListCdrsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all CDRs for your account. Supports filtering by phone numbers,
date range, call direction, duration, and pagination.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListCdrsRequest{
        AuthID: "MA_XXXXXX",
        FromNumber: vobiz.String(
            "9876543210",
        ),
        ToNumber: vobiz.String(
            "1234567890",
        ),
        StartDate: vobiz.Time(
            vobiz.MustParseDate(
                "2026-03-01",
            ),
        ),
        EndDate: vobiz.Time(
            vobiz.MustParseDate(
                "2026-03-17",
            ),
        ),
        MinDuration: vobiz.Int(
            10,
        ),
        SipCallID: vobiz.String(
            "dD1qwu5VZ5iK3ed5u3uspjY5RKL",
        ),
        BridgeUUID: vobiz.String(
            "4b7ae653-f40d-42f1-b582-6b05dfcd0c0a",
        ),
        HangupCause: vobiz.String(
            "NORMAL_CLEARING",
        ),
        HangupDisposition: vobiz.String(
            "send_refuse",
        ),
        Context: vobiz.String(
            "sip-trunking",
        ),
    }
client.Cdr.ListCdrs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**fromNumber:** `*string` — Filter by the originating phone number (caller).
    
</dd>
</dl>

<dl>
<dd>

**toNumber:** `*string` — Filter by the destination phone number (callee).
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Beginning of the search period (YYYY-MM-DD). Required when using `end_date`.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End of the search period (YYYY-MM-DD). Required when using `start_date`.
    
</dd>
</dl>

<dl>
<dd>

**callDirection:** `*vobiz.ListCdrsRequestCallDirection` — Filter by direction.
    
</dd>
</dl>

<dl>
<dd>

**minDuration:** `*int` — Minimum call duration in seconds. Excludes calls shorter than this value.
    
</dd>
</dl>

<dl>
<dd>

**sipCallID:** `*string` — Filter by the SIP Call-ID of the call (matches the cdr's sip_call_id field).
    
</dd>
</dl>

<dl>
<dd>

**bridgeUUID:** `*string` — Filter by the UUID of the bridged leg (matches the cdr's bridge_uuid field).
    
</dd>
</dl>

<dl>
<dd>

**hangupCause:** `*string` — Filter by telephony hangup cause, e.g. NORMAL_CLEARING.
    
</dd>
</dl>

<dl>
<dd>

**hangupDisposition:** `*string` — Filter by how the leg was released, e.g. send_refuse.
    
</dd>
</dl>

<dl>
<dd>

**context:** `*string` — Filter by the call context, e.g. sip-trunking.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — Filter by the campaign identifier associated with the call.
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Free-text search across CDR fields (numbers, IDs, etc.).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number for paginated results.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of records per page. Max: 100.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cdr.SearchCdrs(AuthID) -> *vobiz.SearchCdrsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Identical filters to the list endpoint, but the response also includes a
`filter_summary` object describing the active filters applied.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.SearchCdrsRequest{
        AuthID: "MA_XXXXXX",
        FromNumber: vobiz.String(
            "9876543210",
        ),
        ToNumber: vobiz.String(
            "1234567890",
        ),
        StartDate: vobiz.Time(
            vobiz.MustParseDate(
                "2026-03-01",
            ),
        ),
        EndDate: vobiz.Time(
            vobiz.MustParseDate(
                "2026-03-17",
            ),
        ),
        MinDuration: vobiz.Int(
            10,
        ),
        SipCallID: vobiz.String(
            "dD1qwu5VZ5iK3ed5u3uspjY5RKL",
        ),
        BridgeUUID: vobiz.String(
            "4b7ae653-f40d-42f1-b582-6b05dfcd0c0a",
        ),
        HangupCause: vobiz.String(
            "NORMAL_CLEARING",
        ),
        HangupDisposition: vobiz.String(
            "send_refuse",
        ),
        Context: vobiz.String(
            "sip-trunking",
        ),
    }
client.Cdr.SearchCdrs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**fromNumber:** `*string` — Filter by the originating phone number (caller).
    
</dd>
</dl>

<dl>
<dd>

**toNumber:** `*string` — Filter by the destination phone number (callee).
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Beginning of the search period (YYYY-MM-DD). Required when using `end_date`.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End of the search period (YYYY-MM-DD). Required when using `start_date`.
    
</dd>
</dl>

<dl>
<dd>

**callDirection:** `*vobiz.SearchCdrsRequestCallDirection` — Filter by direction.
    
</dd>
</dl>

<dl>
<dd>

**minDuration:** `*int` — Minimum call duration in seconds. Excludes calls shorter than this value.
    
</dd>
</dl>

<dl>
<dd>

**sipCallID:** `*string` — Filter by the SIP Call-ID of the call (matches the cdr's sip_call_id field).
    
</dd>
</dl>

<dl>
<dd>

**bridgeUUID:** `*string` — Filter by the UUID of the bridged leg (matches the cdr's bridge_uuid field).
    
</dd>
</dl>

<dl>
<dd>

**hangupCause:** `*string` — Filter by telephony hangup cause, e.g. NORMAL_CLEARING.
    
</dd>
</dl>

<dl>
<dd>

**hangupDisposition:** `*string` — Filter by how the leg was released, e.g. send_refuse.
    
</dd>
</dl>

<dl>
<dd>

**context:** `*string` — Filter by the call context, e.g. sip-trunking.
    
</dd>
</dl>

<dl>
<dd>

**campaignID:** `*string` — Filter by the campaign identifier associated with the call.
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Free-text search across CDR fields (numbers, IDs, etc.).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number for paginated results.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of records per page. Max: 100.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cdr.ListRecentCdrs(AuthID) -> *vobiz.ListRecentCdrsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the most recent CDRs for your account without requiring a date range.
Default 20 records; use `limit` to retrieve more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListRecentCdrsRequest{
        AuthID: "MA_XXXXXX",
        Limit: vobiz.Int(
            50,
        ),
    }
client.Cdr.ListRecentCdrs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of recent CDRs to return.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cdr.GetCdr(AuthID, CallID) -> *vobiz.GetCdrResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the CDR for a specific completed call using its `call_id`.
Useful when you have a `call_id` from a callback or previous API response.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetCdrRequest{
        AuthID: "MA_XXXXXX",
        CallID: "abc123-def456-ghi789",
    }
client.Cdr.GetCdr(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callID:** `string` — The unique call ID of the completed call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Accounts
<details><summary><code>client.SubAccounts.ListSubaccounts(AuthID) -> *vobiz.ListSubaccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all sub-accounts under the master account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListSubaccountsRequest{
        AuthID: "MA_XXXXXX",
    }
client.SubAccounts.ListSubaccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccounts.CreateSubaccount(AuthID, request) -> *vobiz.CreateSubaccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new sub-account under the master account.

Set `kyc_mode` to control how the sub-account is verified:

- `personal_use` *(default)* - the sub-account inherits the parent's
  KYC; no separate verification is required.
- `customer_use` - the sub-account must complete its own KYC before it
  can place calls. A fresh `customer_use` sub-account is returned with
  `kyc_calls_blocked: true`. `customer_use` **requires** `email`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateSubaccountRequest{
        AuthID: "MA_XXXXXX",
        Name: "Customer Co",
        Email: vobiz.String(
            "customer@example.com",
        ),
        Password: vobiz.String(
            "Customer@12345",
        ),
        KycMode: vobiz.CreateSubaccountRequestKycModeCustomerUse.Ptr(),
        BusinessType: vobiz.CreateSubaccountRequestBusinessTypePrivateLimited.Ptr(),
    }
client.SubAccounts.CreateSubaccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Human-readable name for the sub-account.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Required when `kyc_mode` is `customer_use`.
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` — Login password for the sub-account.
    
</dd>
</dl>

<dl>
<dd>

**kycMode:** `*vobiz.CreateSubaccountRequestKycMode` 

`personal_use` inherits parent KYC. `customer_use` requires
the sub-account to complete its own KYC and requires `email`.
    
</dd>
</dl>

<dl>
<dd>

**businessType:** `*vobiz.CreateSubaccountRequestBusinessType` — Legal constitution of the customer. Drives which KYC documents are required.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccounts.RetrieveSubaccount(AuthID, SubAuthID) -> *vobiz.RetrieveSubaccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of a specific sub-account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.RetrieveSubaccountRequest{
        AuthID: "MA_XXXXXX",
        SubAuthID: "SA_XXXXXX",
    }
client.SubAccounts.RetrieveSubaccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**subAuthID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccounts.UpdateSubaccount(AuthID, SubAuthID, request) -> *vobiz.UpdateSubaccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the name or status of a sub-account, or change its `kyc_mode`.

Promoting an existing sub-account to `customer_use` requires the
sub-account to already have an `email` (otherwise `400`). On any
`kyc_mode` change, `kyc_calls_blocked` is re-derived from the
sub-account's current KYC state.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateSubaccountRequest{
        AuthID: "MA_XXXXXX",
        SubAuthID: "sub_auth_id",
        KycMode: vobiz.UpdateSubaccountRequestKycModeCustomerUse.Ptr(),
    }
client.SubAccounts.UpdateSubaccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**subAuthID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**kycMode:** `*vobiz.UpdateSubaccountRequestKycMode` — Change the verification mode. Promoting to `customer_use` requires the sub-account to have an `email`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccounts.DeleteSubaccount(AuthID, SubAuthID) -> *vobiz.DeleteSubaccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a sub-account and revoke its credentials.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteSubaccountRequest{
        AuthID: "MA_XXXXXX",
        SubAuthID: "sub_auth_id",
    }
client.SubAccounts.DeleteSubaccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**subAuthID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Account KYC
<details><summary><code>client.SubAccountKyc.GetSubaccountKycStatus(SubAuthID) -> *vobiz.SubAccountKycStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the aggregated KYC state for a `customer_use` sub-account —
which verifications have passed, whether calls are still blocked, and
the business type. The caller must be the parent main account that owns
the sub-account (or an admin).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetSubaccountKycStatusRequest{
        SubAuthID: "SA_XXXXXX",
    }
client.SubAccountKyc.GetSubaccountKycStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.VerifySubaccountPan(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Runs a real PAN verification (Perfios) for the sub-account. `pan` must
be exactly 10 characters. Persists a `kyc_verifications` row and
recomputes the sub-account's aggregated `kyc_status`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.VerifySubaccountPanRequest{
        SubAuthID: "SA_XXXXXX",
        Pan: "ABCDE1234F",
    }
client.SubAccountKyc.VerifySubaccountPan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**pan:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.VerifySubaccountGst(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Runs a real GSTIN verification. `gstin` must be a 15-character GSTIN.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.VerifySubaccountGstRequest{
        SubAuthID: "SA_XXXXXX",
        Gstin: "29AAJCN5983D1Z0",
    }
client.SubAccountKyc.VerifySubaccountGst(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**gstin:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.SearchSubaccountCin(SubAuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Name-based CIN lookup. Returns candidate company matches; pick one and
pass it to [CIN confirm](#operation/confirm-subaccount-cin).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.SearchSubaccountCinRequest{
        SubAuthID: "SA_XXXXXX",
        CompanyName: "ACME PRIVATE LIMITED",
    }
client.SubAccountKyc.SearchSubaccountCin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**companyName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.ConfirmSubaccountCin(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Confirm the CIN selected from the search results.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ConfirmSubaccountCinRequest{
        SubAuthID: "SA_XXXXXX",
        CompanyName: "ACME PRIVATE LIMITED",
        SelectedCin: "U72900KA2024PTC123456",
    }
client.SubAccountKyc.ConfirmSubaccountCin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**companyName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**selectedCin:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.SubaccountDigilockerInitiate(SubAuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the DigiLocker authorization link and an `access_request_id`.
The customer completes the OAuth flow on the DigiLocker portal, after
which you finalize with
[DigiLocker verify](#operation/subaccount-digilocker-verify).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.SubaccountDigilockerInitiateRequest{
        SubAuthID: "SA_XXXXXX",
        RedirectURL: "https://partner.example.com/kyc/callback",
    }
client.SubAccountKyc.SubaccountDigilockerInitiate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**oauthState:** `*string` — Opaque value echoed back on the redirect for CSRF protection.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.SubaccountDigilockerVerify(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Finalize Aadhaar via DigiLocker after the customer completes OAuth.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.SubaccountDigilockerVerifyRequest{
        SubAuthID: "SA_XXXXXX",
        AccessRequestID: "AR_xxxxxxxx",
    }
client.SubAccountKyc.SubaccountDigilockerVerify(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**accessRequestID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**linkedNumber:** `*string` — Optional. Binds the Aadhaar to a specific number (92-series).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKyc.CreateSubaccountKycSession(SubAuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a Vobiz-hosted KYC session for the sub-account. With
`flow_type=email` (default) Vobiz emails the customer a signed link
(from `kyc@vobiz.ai`, hosted at `kyc.vobiz.ai`) and `customer_email` is
required. With `flow_type=redirect`, omit `customer_email`, pass a
`redirect_url`, and the `widget_url` is returned directly for an inline
redirect.

This is the sub-account–scoped equivalent of the partner-level
[KYC Sessions](/partner/api/kyc-sessions) endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateSubaccountKycSessionRequest{
        SubAuthID: "SA_XXXXXX",
        AccountAuthID: "SA_XXXXXX",
        FlowType: vobiz.CreateSubaccountKycSessionRequestFlowTypeEmail,
        CustomerEmail: vobiz.String(
            "customer@example.com",
        ),
        WebhookURL: vobiz.String(
            "https://your-app.example.com/kyc/webhook",
        ),
        ExpiresInDays: vobiz.Int(
            30,
        ),
    }
client.SubAccountKyc.CreateSubaccountKycSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**accountAuthID:** `string` — The sub-account's auth_id (typically equal to the path `sub_auth_id`).
    
</dd>
</dl>

<dl>
<dd>

**flowType:** `*vobiz.CreateSubaccountKycSessionRequestFlowType` 
    
</dd>
</dl>

<dl>
<dd>

**customerEmail:** `*string` — Required when `flow_type` is `email`.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `*string` 

Required when `flow_type` is `redirect`. After verification the customer's
browser is sent to this URL.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — HTTPS endpoint VoBiz POSTs the KYC result to. Omit it and no callbacks are sent.
    
</dd>
</dl>

<dl>
<dd>

**expiresInDays:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Account KYC (Test Mode)
<details><summary><code>client.SubAccountKycTestMode.MockVerifySubaccountPan(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mock PAN verification - never hits the provider. Magic `pan` inputs:

| Input | Outcome |
|---|---|
| `TESTSUCCESS0001` | verified |
| `TESTFAIL0001` | failed |
| `TESTERROR0001` | HTTP 500 |
| `TESTPENDING001` | pending (finalize as verified) |
| `TESTPENDING_FAIL` | pending (finalize as failed) |

Persists a real `kyc_verifications` row and recomputes `kyc_status`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockVerifySubaccountPanRequest{
        SubAuthID: "SA_XXXXXX",
        Pan: "TESTSUCCESS0001",
    }
client.SubAccountKycTestMode.MockVerifySubaccountPan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**pan:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKycTestMode.MockVerifySubaccountGst(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mock GST verification. Same magic-input matrix as [Mock verify PAN](#operation/mock-verify-subaccount-pan).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockVerifySubaccountGstRequest{
        SubAuthID: "SA_XXXXXX",
        Gstin: "TESTSUCCESS0001GST",
    }
client.SubAccountKycTestMode.MockVerifySubaccountGst(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**gstin:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKycTestMode.MockSearchSubaccountCin(SubAuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns deterministic fake company matches.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockSearchSubaccountCinRequest{
        SubAuthID: "SA_XXXXXX",
        CompanyName: "ACME",
    }
client.SubAccountKycTestMode.MockSearchSubaccountCin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**companyName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKycTestMode.MockConfirmSubaccountCin(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Succeeds when `selected_cin` starts with `U72900KA2024PTC123456`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockConfirmSubaccountCinRequest{
        SubAuthID: "SA_XXXXXX",
        CompanyName: "ACME",
        SelectedCin: "U72900KA2024PTC123456",
    }
client.SubAccountKycTestMode.MockConfirmSubaccountCin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**companyName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**selectedCin:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKycTestMode.MockSubaccountDigilockerInitiate(SubAuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a deterministic `access_request_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockSubaccountDigilockerInitiateRequest{
        SubAuthID: "SA_XXXXXX",
        RedirectURL: "https://partner.example.com/kyc/callback",
    }
client.SubAccountKycTestMode.MockSubaccountDigilockerInitiate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKycTestMode.MockSubaccountDigilockerVerify(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

`access_request_id` `MOCK_AR_SUCCESS` → verified; `MOCK_AR_FAIL` → failed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockSubaccountDigilockerVerifyRequest{
        SubAuthID: "SA_XXXXXX",
        AccessRequestID: vobiz.MockSubaccountDigilockerVerifyRequestAccessRequestIDMockArSuccess,
    }
client.SubAccountKycTestMode.MockSubaccountDigilockerVerify(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**accessRequestID:** `*vobiz.MockSubaccountDigilockerVerifyRequestAccessRequestID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountKycTestMode.MockFinalizePendingKyc(SubAuthID, request) -> *vobiz.KycVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Promotes the most recent **pending** mock verification of the given
type to a terminal outcome - this drives the async (`TESTPENDING…`)
path without webhooks. `verification_type` ∈ `pan | aadhaar | gst | cin`;
`outcome` ∈ `verified | failed`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MockFinalizePendingKycRequest{
        SubAuthID: "SA_XXXXXX",
        VerificationType: vobiz.MockFinalizePendingKycRequestVerificationTypePan,
        Outcome: vobiz.MockFinalizePendingKycRequestOutcomeVerified,
    }
client.SubAccountKycTestMode.MockFinalizePendingKyc(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAuthID:** `string` — The sub-account's Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**verificationType:** `*vobiz.MockFinalizePendingKycRequestVerificationType` 
    
</dd>
</dl>

<dl>
<dd>

**outcome:** `*vobiz.MockFinalizePendingKycRequestOutcome` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Phone Numbers
<details><summary><code>client.PhoneNumbers.ListNumbers(AuthID) -> *vobiz.ListNumbersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all phone numbers on your account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListNumbersRequest{
        AuthID: "MA_XXXXXX",
        Search: vobiz.String(
            "+919876543210",
        ),
    }
client.PhoneNumbers.ListNumbers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number, starting at 1
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of phone numbers to return per page
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Filter by phone number. Include the country code and URL-encode a leading plus sign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.UnrentNumber(AuthID, E164) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Release a phone number from your account. By default, the number enters
`pending_release` for a 24-hour cooldown. You can cancel the release during
that window. Set `immediate=true` to skip the cooldown; an immediate release
cannot be cancelled.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UnrentNumberRequest{
        AuthID: "MA_XXXXXX",
        E164: "919876543210",
    }
client.PhoneNumbers.UnrentNumber(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**e164:** `string` — Phone number in E.164 format (without the +)
    
</dd>
</dl>

<dl>
<dd>

**immediate:** `*bool` — Skip the 24-hour cooldown and release the number immediately.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.CancelNumberRelease(AccountID, E164) -> *vobiz.CancelNumberReleaseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel a pending number release during the 24-hour cooldown. The number is
restored to `active`, the cooldown timer is cleared, and the release fee is
refunded. Any trunk or voice application detached by the release is not
re-attached automatically.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CancelNumberReleaseRequest{
        AccountID: "MA_XXXXXX",
        E164: "%2B919876543210",
    }
client.PhoneNumbers.CancelNumberRelease(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Your account Auth ID.
    
</dd>
</dl>

<dl>
<dd>

**e164:** `string` — The URL-encoded phone number in E.164 format. Encode `+` as `%2B`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.ListInventoryNumbers(AuthID) -> *vobiz.ListInventoryNumbersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Browse available phone numbers in inventory that are not assigned to
any account. Only numbers with `status='active'` and `auth_id=NULL`
are returned. These numbers are ready to be purchased.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListInventoryNumbersRequest{
        AuthID: "MA_XXXXXX",
        Country: vobiz.String(
            "IN",
        ),
        Exclude: vobiz.String(
            "9180,9192",
        ),
    }
client.PhoneNumbers.ListInventoryNumbers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Filter by country code (e.g., "US", "IN").
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Substring match against the E.164 number (e.g., "80" matches "+918065...").
    
</dd>
</dl>

<dl>
<dd>

**exclude:** `*string` — One or more E.164 prefixes to remove from results. Include the country code (e.g. "9180" for India +91 80-series, "1415" for US +1 415); a leading "+" is optional. Matched against the full E.164 form, so it works for any country. Accepts a comma-separated list ("9180,9192") or repeated params ("exclude=9180&exclude=9192"), and the two forms can be combined. It is ANDed with all other filters, so it takes priority over `search`; duplicates are de-duplicated silently and `total` reflects the filtered result set.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.PurchaseFromInventory(AuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Purchase a phone number from inventory and assign it to your account.
Debits your account balance for the setup fee and monthly fee. For
sub-accounts (SA_), the parent master account (MA_) is charged.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.PurchaseFromInventoryRequest{
        AuthID: "MA_XXXXXX",
        E164: "+919876543210",
        Currency: vobiz.String(
            "USD",
        ),
    }
client.PhoneNumbers.PurchaseFromInventory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**e164:** `string` — Phone number to purchase in E.164 format.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Currency for transaction. Defaults to the number's currency or "USD".
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.AssignNumberToTrunk(AuthID, PhoneNumber, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign a phone number to a specific SIP trunk. Once assigned, all
inbound calls to that phone number will be routed through the
designated trunk. The phone number must be URL-encoded; use `%2B`
instead of `+` (e.g., `%2B912271264217`).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.AssignNumberToTrunkRequest{
        AuthID: "MA_XXXXXX",
        PhoneNumber: "%2B912271264217",
        TrunkGroupID: "e3e55a78-1234-5678-90ab-cdef12345678",
    }
client.PhoneNumbers.AssignNumberToTrunk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `string` — The phone number to assign, URL-encoded (use %2B instead of +).
    
</dd>
</dl>

<dl>
<dd>

**trunkGroupID:** `string` — The UUID of the trunk to assign this number to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.UnassignNumberFromTrunk(AuthID, PhoneNumber) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove the assignment between a phone number and a SIP trunk. After
unassignment, the number remains in your account inventory but will
no longer route inbound calls through the previously assigned trunk.
URL-encode the phone number (use `%2B` instead of `+`).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UnassignNumberFromTrunkRequest{
        AuthID: "MA_XXXXXX",
        PhoneNumber: "%2B912271264217",
    }
client.PhoneNumbers.UnassignNumberFromTrunk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `string` — The phone number to unassign, URL-encoded (use %2B instead of +).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.GetNumberHealth(AuthID, E164) -> *vobiz.GetNumberHealthResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the health & analytics dashboard for one of your numbers: current
status, spam flag, and call metrics over the selected window (total and
answered calls, answer rate, minutes, average duration) plus a per-period
time series of snapshots.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetNumberHealthRequest{
        AuthID: "MA_XXXXXX",
        E164: "%2B919876543210",
        Days: vobiz.Int(
            30,
        ),
    }
client.PhoneNumbers.GetNumberHealth(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**e164:** `string` — The number in E.164, URL-encoded (use %2B instead of +).
    
</dd>
</dl>

<dl>
<dd>

**granularity:** `*vobiz.GetNumberHealthRequestGranularity` — Snapshot granularity.
    
</dd>
</dl>

<dl>
<dd>

**days:** `*int` — Size of the window (in days) for the summary and snapshots.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.AssignDidToSubaccount(AuthID, E164, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign a parent-pool DID to a sub-account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.AssignDidToSubaccountRequest{
        AuthID: "MA_XXXXXX",
        E164: "%2B919876543210",
        SubAccountID: "SA_XXXXXX",
    }
client.PhoneNumbers.AssignDidToSubaccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**e164:** `string` — The number in E.164, URL-encoded (use %2B instead of +).
    
</dd>
</dl>

<dl>
<dd>

**subAccountID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneNumbers.UnassignDidFromSubaccount(AuthID, E164) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Move the DID back to the parent pool.

A **15-day cool-off** is enforced: if the DID had a call within the last
15 days, the request is rejected with `409` and a
`did_cool_off_in_effect` error that includes `cool_off_until` and
`cool_off_remaining_seconds`. Never-used DIDs (`last_call_at` is `NULL`)
move back immediately.

Admins can bypass the cool-off with `?force=true` (see below); the
bypass writes a `did_assignment_audit` row and requires an
admin-role account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UnassignDidFromSubaccountRequest{
        AuthID: "MA_XXXXXX",
        E164: "%2B919876543210",
        Force: vobiz.Bool(
            true,
        ),
    }
client.PhoneNumbers.UnassignDidFromSubaccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**e164:** `string` — The number in E.164, URL-encoded (use %2B instead of +).
    
</dd>
</dl>

<dl>
<dd>

**force:** `*bool` 

Admin-only cool-off bypass. Requires an admin-role account
(enforced at the gateway) and writes a `did_assignment_audit` row.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Trunks
<details><summary><code>client.Trunks.ListTrunks(AuthID) -> *vobiz.ListTrunksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all SIP trunks configured on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListTrunksRequest{
        AuthID: "MA_XXXXXX",
    }
client.Trunks.ListTrunks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Trunks.CreateTrunk(AuthID, request) -> *vobiz.CreateTrunkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new SIP trunk for inbound or outbound calling.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateTrunkRequest{
        AuthID: "MA_XXXXXX",
        Name: "Retell AI SIP",
        TrunkDirection: vobiz.CreateTrunkRequestTrunkDirectionOutbound.Ptr(),
        Transport: vobiz.CreateTrunkRequestTransportUDP.Ptr(),
        ConcurrentCallsLimit: vobiz.Int(
            50,
        ),
        CpsLimit: vobiz.Int(
            15,
        ),
        CredentialUUID: vobiz.String(
            "b1e2...",
        ),
        IpaclUUID: vobiz.String(
            "c3d4...",
        ),
        Recording: vobiz.Bool(
            true,
        ),
        EnableTranscription: vobiz.Bool(
            true,
        ),
        WebhookURL: vobiz.String(
            "https://example.com/vobiz/webhook",
        ),
        WebhookMethod: vobiz.CreateTrunkRequestWebhookMethodPost.Ptr(),
    }
client.Trunks.CreateTrunk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Trunk name.
    
</dd>
</dl>

<dl>
<dd>

**trunkDirection:** `*vobiz.CreateTrunkRequestTrunkDirection` — Direction of the trunk - **`inbound` or `outbound` only** (a trunk is one direction, not both).
    
</dd>
</dl>

<dl>
<dd>

**trunkStatus:** `*vobiz.CreateTrunkRequestTrunkStatus` — Trunk status - `enabled` or `disabled` (note: not `active`).
    
</dd>
</dl>

<dl>
<dd>

**secure:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**trunkDomain:** `*string` — SIP domain. Auto-generated as `{first8ofUUID}.sip.vobiz.ai` if omitted.
    
</dd>
</dl>

<dl>
<dd>

**transport:** `*vobiz.CreateTrunkRequestTransport` 
    
</dd>
</dl>

<dl>
<dd>

**inboundDestination:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**concurrentCallsLimit:** `*int` — Stored on the trunk. The **enforced** concurrency limit is account-level (account base + channel subscriptions), not this field.
    
</dd>
</dl>

<dl>
<dd>

**cpsLimit:** `*int` — Stored on the trunk. The **enforced** CPS is account-level, not this field.
    
</dd>
</dl>

<dl>
<dd>

**credentialUUID:** `*string` — Attach an existing SIP credential (username / password / realm) by UUID.
    
</dd>
</dl>

<dl>
<dd>

**ipaclUUID:** `*string` — Attach an existing IP access-control list (IP-based auth) by UUID.
    
</dd>
</dl>

<dl>
<dd>

**primaryUriUuid:** `*string` — Primary origination URI UUID.
    
</dd>
</dl>

<dl>
<dd>

**fallbackUriUuid:** `*string` — Fallback origination URI UUID.
    
</dd>
</dl>

<dl>
<dd>

**recording:** `*bool` — Enable call recording.
    
</dd>
</dl>

<dl>
<dd>

**enableTranscription:** `*bool` — Auto-transcribe recordings when `recording=true`.
    
</dd>
</dl>

<dl>
<dd>

**piiRedaction:** `*bool` — Redact PII from transcriptions.
    
</dd>
</dl>

<dl>
<dd>

**piiEntityTypes:** `*string` — Comma-separated list of entity types to redact.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` 

Customer webhook for call-admission events (`CallInitiated` / `Hangup`).
Must be a valid **public** http/https URL. SSRF-validated - localhost,
private (RFC1918), and cloud-metadata (`169.254.169.254`) URLs are
rejected with `invalid webhook_url`. See [Trunk Webhooks](/trunks/webhook).
    
</dd>
</dl>

<dl>
<dd>

**webhookMethod:** `*vobiz.CreateTrunkRequestWebhookMethod` — HTTP method for the webhook callback.
    
</dd>
</dl>

<dl>
<dd>

**recordingWebhookEnabled:** `*bool` — Fire a `recording.completed` webhook to `webhook_url` after a recording is saved.
    
</dd>
</dl>

<dl>
<dd>

**username:** `*string` — Deprecated - use `credential_uuid`.
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` — Deprecated - use `credential_uuid`.
    
</dd>
</dl>

<dl>
<dd>

**ipWhitelist:** `[]string` — Deprecated - use `ipacl_uuid`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Trunks.RetrieveTrunk(AuthID, TrunkID) -> *vobiz.RetrieveTrunkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details of a specific SIP trunk.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.RetrieveTrunkRequest{
        AuthID: "MA_XXXXXX",
        TrunkID: "trunk_XXXXXX",
    }
client.Trunks.RetrieveTrunk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**trunkID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Trunks.UpdateTrunk(AuthID, TrunkID, request) -> *vobiz.UpdateTrunkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a SIP trunk's name, configuration, or status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateTrunkRequest{
        AuthID: "MA_XXXXXX",
        TrunkID: "trunk_id",
    }
client.Trunks.UpdateTrunk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**trunkID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**trunkDirection:** `*vobiz.UpdateTrunkRequestTrunkDirection` — Direction of the trunk - `inbound` or `outbound` only.
    
</dd>
</dl>

<dl>
<dd>

**trunkStatus:** `*vobiz.UpdateTrunkRequestTrunkStatus` 
    
</dd>
</dl>

<dl>
<dd>

**secure:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**trunkDomain:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**transport:** `*vobiz.UpdateTrunkRequestTransport` 
    
</dd>
</dl>

<dl>
<dd>

**inboundDestination:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**concurrentCallsLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cpsLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**credentialUUID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ipaclUUID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**primaryUriUuid:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**fallbackUriUuid:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**recording:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**enableTranscription:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**piiRedaction:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**piiEntityTypes:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — Customer webhook for call-admission events (`CallInitiated` / `Hangup`). Public http/https URL; SSRF-validated. See [Trunk Webhooks](/trunks/webhook).
    
</dd>
</dl>

<dl>
<dd>

**webhookMethod:** `*vobiz.UpdateTrunkRequestWebhookMethod` 
    
</dd>
</dl>

<dl>
<dd>

**recordingWebhookEnabled:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Trunks.DeleteTrunk(AuthID, TrunkID) -> *string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a SIP trunk.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteTrunkRequest{
        AuthID: "MA_XXXXXX",
        TrunkID: "trunk_id",
    }
client.Trunks.DeleteTrunk(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**trunkID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Conference
<details><summary><code>client.Conference.KickMember(AuthID, ConferenceName, MemberID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove one or more participants from a conference while allowing their XML flow to continue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.KickMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.Conference.KickMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conference.HangupMember(AuthID, ConferenceName, MemberID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminate one or more active conference member calls. A normal active-member request disconnects the member. If a member was kicked, continued its XML flow, and rejoined with the same numeric member ID, confirm removal through conference exit or call hangup callbacks.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.HangupMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.Conference.HangupMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conference.PlayAudioMember(AuthID, ConferenceName, MemberID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Play an audio file to a specific conference member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.PlayAudioMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
        URL: "https://example.com/audio.mp3",
    }
client.Conference.PlayAudioMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — URL of the audio file to play
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conference.StopAudioMember(AuthID, ConferenceName, MemberID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop audio playback for a specific conference member.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StopAudioMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.Conference.StopAudioMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conference.DeafMember(AuthID, ConferenceName, MemberID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Prevent a conference member from hearing other participants.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeafMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.Conference.DeafMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conference.UndeafMember(AuthID, ConferenceName, MemberID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Restore a conference member's ability to hear other participants.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UndeafMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.Conference.UndeafMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RecordCalls
<details><summary><code>client.RecordCalls.StartRecording(AuthID, CallUUID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Begin recording an active call. Set format, enable transcription, and configure a callback URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StartRecordingRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "cdr_XXXXXXXXXX",
        TimeLimit: vobiz.Int(
            120,
        ),
        FileFormat: vobiz.StartRecordingRequestFileFormatMp3.Ptr(),
    }
client.RecordCalls.StartRecording(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**timeLimit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**fileFormat:** `*vobiz.StartRecordingRequestFileFormat` 
    
</dd>
</dl>

<dl>
<dd>

**transcriptionType:** `*string` — Set to `auto` to enable transcription
    
</dd>
</dl>

<dl>
<dd>

**callbackURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**recordChannelType:** `*vobiz.StartRecordingRequestRecordChannelType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RecordCalls.StopRecording(AuthID, CallUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop an active recording on an in-progress call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StopRecordingRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
    }
client.RecordCalls.StopRecording(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PlayAudio
<details><summary><code>client.PlayAudio.Call(AuthID, CallUUID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Play an audio file to a live call leg.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.PlayAudioCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
        URLs: "https://example.com/audio.mp3",
        Legs: vobiz.PlayAudioCallRequestLegsAleg.Ptr(),
    }
client.PlayAudio.Call(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**urls:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**legs:** `*vobiz.PlayAudioCallRequestLegs` 
    
</dd>
</dl>

<dl>
<dd>

**loop:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PlayAudio.StopAudioCall(AuthID, CallUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop audio playing on a live call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StopAudioCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
    }
client.PlayAudio.StopAudioCall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SpeakText
<details><summary><code>client.SpeakText.Call(AuthID, CallUUID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert text to speech and play it on a live call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.SpeakTextCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
        Text: "Hello, your appointment is confirmed for tomorrow at 3 PM.",
        Voice: vobiz.String(
            "WOMAN",
        ),
        Language: vobiz.String(
            "en-US",
        ),
    }
client.SpeakText.Call(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**text:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**voice:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**language:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**legs:** `*vobiz.SpeakTextCallRequestLegs` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpeakText.StopSpeakCall(AuthID, CallUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop ongoing TTS playback on a live call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StopSpeakCallRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
    }
client.SpeakText.StopSpeakCall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Dtmf
<details><summary><code>client.Dtmf.SendDtmf(AuthID, CallUUID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send DTMF (keypad) tones on an active call. Use `w` for 0.5s pause, `W` for 1s pause.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.SendDtmfRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
        Digits: "1234",
        Leg: vobiz.SendDtmfRequestLegAleg.Ptr(),
    }
client.Dtmf.SendDtmf(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**digits:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**leg:** `*vobiz.SendDtmfRequestLeg` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AudioStreams
<details><summary><code>client.AudioStreams.ListStreams(AuthID, CallUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all audio streams on a live call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListStreamsRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
    }
client.AudioStreams.ListStreams(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AudioStreams.StartStream(AuthID, CallUUID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Start streaming raw audio from a live call to a WebSocket URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StartStreamRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
        ServiceURL: "wss://your-server.com/ws",
        Bidirectional: vobiz.Bool(
            true,
        ),
        AudioTrack: vobiz.StartStreamRequestAudioTrackBoth.Ptr(),
    }
client.AudioStreams.StartStream(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**serviceURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**bidirectional:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**audioTrack:** `*vobiz.StartStreamRequestAudioTrack` 
    
</dd>
</dl>

<dl>
<dd>

**audioFormat:** `*vobiz.StartStreamRequestAudioFormat` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AudioStreams.GetStream(AuthID, CallUUID, StreamID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details of a specific audio stream.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetStreamRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
        StreamID: "stream_id",
    }
client.AudioStreams.GetStream(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**streamID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AudioStreams.StopStream(AuthID, CallUUID, StreamID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop a specific audio stream on a live call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StopStreamRequest{
        AuthID: "MA_XXXXXX",
        CallUUID: "call_uuid",
        StreamID: "stream_id",
    }
client.AudioStreams.StopStream(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**callUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**streamID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Conferences
<details><summary><code>client.Conferences.ListConferences(AuthID) -> *vobiz.ListConferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve conference room names reported by the API. An empty array is inconclusive and can occur while conferences are active. Maintain your own room registry for authoritative discovery, billing, cleanup, and destructive workflows.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListConferencesRequest{
        AuthID: "MA_XXXXXX",
    }
client.Conferences.ListConferences(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conferences.DeleteAllConferences(AuthID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminate all active conference rooms.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteAllConferencesRequest{
        AuthID: "MA_XXXXXX",
    }
client.Conferences.DeleteAllConferences(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conferences.GetConference(AuthID, ConferenceName) -> *vobiz.GetConferenceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific conference room. A live conference can currently return a 200 response with an error payload instead of conference details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetConferenceRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "My Conf Room",
    }
client.Conferences.GetConference(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conferences.DeleteConference(AuthID, ConferenceName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminate a specific conference room and disconnect all members.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteConferenceRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
    }
client.Conferences.DeleteConference(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ConferenceMembers
<details><summary><code>client.ConferenceMembers.MuteMember(AuthID, ConferenceName, MemberID) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Prevent a member from speaking. Use `all` as member_id to mute everyone.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.MuteMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.ConferenceMembers.MuteMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` — Member ID, comma-separated IDs, or `all`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConferenceMembers.UnmuteMember(AuthID, ConferenceName, MemberID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Allow a muted member to speak again.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UnmuteMemberRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
        MemberID: "member_id",
    }
client.ConferenceMembers.UnmuteMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ConferenceRecording
<details><summary><code>client.ConferenceRecording.StartConferenceRecording(AuthID, ConferenceName, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queue recording for all audio in a conference room. The response does not include a recording ID or download URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StartConferenceRecordingRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
    }
client.ConferenceRecording.StartConferenceRecording(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fileFormat:** `*vobiz.StartConferenceRecordingRequestFileFormat` 
    
</dd>
</dl>

<dl>
<dd>

**callbackURL:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConferenceRecording.StopConferenceRecording(AuthID, ConferenceName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop recording a conference room.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.StopConferenceRecordingRequest{
        AuthID: "MA_XXXXXX",
        ConferenceName: "conference_name",
    }
client.ConferenceRecording.StopConferenceRecording(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**conferenceName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Recordings
<details><summary><code>client.Recordings.ListRecordings(AuthID) -> *vobiz.ListRecordingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all call recordings on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListRecordingsRequest{
        AuthID: "MA_XXXXXX",
    }
client.Recordings.ListRecordings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Recordings.GetRecording(AuthID, RecordingID) -> *vobiz.GetRecordingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details and download URL for a specific recording.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetRecordingRequest{
        AuthID: "MA_XXXXXX",
        RecordingID: "rec_XXXXXXXXXX",
    }
client.Recordings.GetRecording(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**recordingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Recordings.DeleteRecording(AuthID, RecordingID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a recording from the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteRecordingRequest{
        AuthID: "MA_XXXXXX",
        RecordingID: "recording_id",
    }
client.Recordings.DeleteRecording(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**recordingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Credentials
<details><summary><code>client.Credentials.CreateCredential(AuthID, request) -> *vobiz.CreateCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create SIP credentials for trunk authentication.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateCredentialRequest{
        AuthID: "MA_XXXXXX",
        Username: "myuser",
        Password: "securepassword123",
    }
client.Credentials.CreateCredential(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**username:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.ListCredentials(AuthID) -> *vobiz.ListCredentialsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all SIP credentials on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListCredentialsRequest{
        AuthID: "MA_XXXXXX",
    }
client.Credentials.ListCredentials(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.UpdateCredential(AuthID, CredentialID, request) -> *vobiz.UpdateCredentialResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the password for an existing SIP credential.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateCredentialRequest{
        AuthID: "MA_XXXXXX",
        CredentialID: "credential_id",
        Password: "password",
    }
client.Credentials.UpdateCredential(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**credentialID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Credentials.DeleteCredential(AuthID, CredentialID) -> *string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an existing SIP credential.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteCredentialRequest{
        AuthID: "MA_XXXXXX",
        CredentialID: "credential_id",
    }
client.Credentials.DeleteCredential(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**credentialID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## IPAccessControlList
<details><summary><code>client.IPAccessControlList.CreateIpAcl(AuthID, request) -> *vobiz.CreateIpAclResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add an IP access control rule to restrict SIP trunk access.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateIpAclRequest{
        AuthID: "MA_XXXXXX",
        Name: "Office IP",
        IPAddress: "ip_address",
    }
client.IPAccessControlList.CreateIpAcl(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IPAccessControlList.ListIpAcls(AuthID) -> *vobiz.ListIpAclsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all IP access control rules on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListIpAclsRequest{
        AuthID: "MA_XXXXXX",
    }
client.IPAccessControlList.ListIpAcls(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IPAccessControlList.UpdateIpAcl(AuthID, IpAclId, request) -> *vobiz.UpdateIpAclResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing IP access control rule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateIpAclRequest{
        AuthID: "MA_XXXXXX",
        IpAclId: "ip_acl_id",
        Name: "name",
        IPAddress: "ip_address",
    }
client.IPAccessControlList.UpdateIpAcl(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**ipAclId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IPAccessControlList.DeleteIpAcl(AuthID, IpAclId) -> *string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove an IP access control rule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteIpAclRequest{
        AuthID: "MA_XXXXXX",
        IpAclId: "ip_acl_id",
    }
client.IPAccessControlList.DeleteIpAcl(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**ipAclId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## OriginationURI
<details><summary><code>client.OriginationURI.CreateOriginationURI(AuthID, request) -> *vobiz.CreateOriginationURIResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add an inbound SIP endpoint (origination URI) to a trunk.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateOriginationURIRequest{
        AuthID: "MA_XXXXXX",
        Name: "Primary SBC",
        SipURI: "sip:sbc.example.com",
        Priority: 1,
    }
client.OriginationURI.CreateOriginationURI(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**sipURI:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**priority:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OriginationURI.ListOriginationURIs(AuthID) -> *vobiz.ListOriginationURIsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all origination URIs on the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListOriginationURIsRequest{
        AuthID: "MA_XXXXXX",
    }
client.OriginationURI.ListOriginationURIs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OriginationURI.UpdateOriginationURI(AuthID, UriId, request) -> *vobiz.UpdateOriginationURIResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing origination URI.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateOriginationURIRequest{
        AuthID: "MA_XXXXXX",
        UriId: "uri_id",
        Name: "name",
        Priority: 1,
    }
client.OriginationURI.UpdateOriginationURI(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**uriId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**priority:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OriginationURI.DeleteOriginationURI(AuthID, UriId) -> *string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an origination URI from a trunk.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteOriginationURIRequest{
        AuthID: "MA_XXXXXX",
        UriId: "uri_id",
    }
client.OriginationURI.DeleteOriginationURI(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**uriId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Applications
<details><summary><code>client.Applications.ListApplications(AuthID) -> *vobiz.ListApplicationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details of all applications created under your Vobiz account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListApplicationsRequest{
        AuthID: "MA_XXXXXX",
    }
client.Applications.ListApplications(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.CreateApplication(AuthID, request) -> *vobiz.CreateApplicationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an Application with webhook URLs for call handling.
Creating an application is usually a first step, after which you
attach the application to either a number or an endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateApplicationRequest{
        AuthID: "MA_XXXXXX",
        AppName: "My Voice Application",
        AnswerURL: "https://example.com/answer",
        AnswerMethod: "POST",
    }
client.Applications.CreateApplication(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**appName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**answerURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**answerMethod:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.RetrieveApplication(AuthID, AppID) -> *vobiz.RetrieveApplicationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details of a particular application by passing the app_id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.RetrieveApplicationRequest{
        AuthID: "MA_XXXXXX",
        AppID: "12345678",
    }
client.Applications.RetrieveApplication(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**appID:** `string` — Unique identifier for the application
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.UpdateApplication(AuthID, AppID, request) -> *vobiz.UpdateApplicationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify an application using this API. You can update any subset of
fields (partial update). Fields not provided will remain unchanged.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateApplicationRequest{
        AuthID: "MA_XXXXXX",
        AppID: "12345678",
        AppName: "Updated Application Name",
        DefaultNumberApp: true,
    }
client.Applications.UpdateApplication(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**appID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**appName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultNumberApp:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.DeleteApplication(AuthID, AppID) -> *string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete an Application. If the application is associated
with phone numbers, the deletion may be blocked unless those
associations are removed first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteApplicationRequest{
        AuthID: "MA_XXXXXX",
        AppID: "12345678",
    }
client.Applications.DeleteApplication(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**appID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Endpoints
<details><summary><code>client.Endpoints.ListEndpoints(AuthID) -> *vobiz.ListEndpointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a paginated list of all SIP endpoints in your account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListEndpointsRequest{
        AuthID: "MA_XXXXXX",
    }
client.Endpoints.ListEndpoints(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**usernameContains:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**usernameExact:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**usernameStartswith:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**aliasContains:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**aliasExact:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**applicationIDExact:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**applicationIDIsnull:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**subAccount:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpoints.CreateEndpoint(AuthID, request) -> *vobiz.CreateEndpointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new SIP endpoint that can be used to make and receive calls
through IP phones, softphones, or SIP clients. Each endpoint is
assigned a unique SIP URI and endpoint ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateEndpointRequest{
        AuthID: "MA_XXXXXX",
        Username: "john_doe",
        Password: "SecurePassword123!",
        Alias: "John's Desktop Phone",
        Application: 12345678,
    }
client.Endpoints.CreateEndpoint(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**username:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**alias:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpoints.RetrieveEndpoint(AuthID, EndpointID) -> *vobiz.RetrieveEndpointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of an existing endpoint. The response includes
all endpoint attributes and, if the endpoint is currently registered
on a SIP client, additional registration details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.RetrieveEndpointRequest{
        AuthID: "MA_XXXXXX",
        EndpointID: "87654321",
    }
client.Endpoints.RetrieveEndpoint(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**endpointID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpoints.UpdateEndpoint(AuthID, EndpointID, request) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing endpoint's configuration. You can change the
password, alias, or attached application. The fields `username`,
`endpoint_id`, `domain`, `allow_same_domain`, `allow_other_domains`,
`allow_phones`, and `allow_apps` are locked after creation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.UpdateEndpointRequest{
        AuthID: "MA_XXXXXX",
        EndpointID: "87654321",
        Alias: "John's Updated Desktop Phone",
        Password: "NewSecurePassword456!",
    }
client.Endpoints.UpdateEndpoint(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**endpointID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**alias:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Endpoints.DeleteEndpoint(AuthID, EndpointID) -> *string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete an endpoint from your Vobiz account. Once deleted,
the SIP URI will no longer be accessible, and any devices registered
with this endpoint will be disconnected.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.DeleteEndpointRequest{
        AuthID: "MA_XXXXXX",
        EndpointID: "87654321",
    }
client.Endpoints.DeleteEndpoint(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authID:** `string` — Your account Auth ID
    
</dd>
</dl>

<dl>
<dd>

**endpointID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Partner API
<details><summary><code>client.PartnerAPI.GetPartnerProfile() -> *vobiz.GetPartnerProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the authenticated partner's profile and balance.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PartnerAPI.GetPartnerProfile(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.GetPartnerDashboard() -> *vobiz.GetPartnerDashboardResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Aggregated partner metrics - total customers, active accounts, balance
held across the partner ecosystem, MTD revenue, etc.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PartnerAPI.GetPartnerDashboard(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.ListCustomerAccounts() -> *vobiz.ListCustomerAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all customer sub-accounts under your partner account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListCustomerAccountsRequest{}
client.PartnerAPI.ListCustomerAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Substring match on name or email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.CreateCustomerAccount(request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new customer sub-account under your partner account. VoBiz
emails the customer their login credentials and (separately) a KYC link
via the kyc-sessions endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateCustomerAccountRequest{
        Name: "John Doe",
        Email: "john@example.com",
        Phone: "+919876543210",
        Password: "SecurePass123!",
        Country: "IN",
    }
client.PartnerAPI.CreateCustomerAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Customer's full name.
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` — Customer's email. KYC link is sent here.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `string` — E.164 format.
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` — Min 8 chars, must include a number and a special character.
    
</dd>
</dl>

<dl>
<dd>

**company:** `*string` — Legal company or business name.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — ISO 2-letter country code.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.PartnerTransferBalance(CustomerAuthID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Atomically debits your partner master balance and credits the customer's
wallet. Both legs are recorded in each account's ledger. Transfers are
**permanent and cannot be reversed.**
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.PartnerTransferBalanceRequest{
        CustomerAuthID: "MA_ZKITB8Z2",
        Amount: 500,
        Currency: "INR",
    }
client.PartnerAPI.PartnerTransferBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerAuthID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `float64` — Positive decimal. Your master balance must be ≥ this amount.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `string` — Must match your partner account currency.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Note for your records. Appears in both ledgers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.ListCustomerTransactions(CustomerAuthID) -> *vobiz.ListCustomerTransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the customer's transaction ledger. Filter by date range or
transaction type. Useful for billing reconciliation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListCustomerTransactionsRequest{
        CustomerAuthID: "customer_auth_id",
        FromDate: vobiz.Time(
            vobiz.MustParseDate(
                "2026-03-01",
            ),
        ),
        ToDate: vobiz.Time(
            vobiz.MustParseDate(
                "2026-03-31",
            ),
        ),
    }
client.PartnerAPI.ListCustomerTransactions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerAuthID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fromDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**toDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**transactionType:** `*vobiz.ListCustomerTransactionsRequestTransactionType` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.ListCustomerCdrs(CustomerAuthID) -> *vobiz.ListCustomerCdrsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Look up any customer's call history. Same filter set as the
customer-side CDR endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListCustomerCdrsRequest{
        CustomerAuthID: "customer_auth_id",
        HangupCause: vobiz.String(
            "NO_ANSWER",
        ),
    }
client.PartnerAPI.ListCustomerCdrs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerAuthID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**callDirection:** `*vobiz.ListCustomerCdrsRequestCallDirection` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*vobiz.ListCustomerCdrsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**minDuration:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**hangupCause:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.ListCustomerNumbers(CustomerAuthID) -> *vobiz.ListCustomerNumbersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Phone numbers currently assigned to a customer account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListCustomerNumbersRequest{
        CustomerAuthID: "customer_auth_id",
    }
client.PartnerAPI.ListCustomerNumbers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customerAuthID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Substring match against the E.164 number.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.ListKycSessions() -> *vobiz.ListKycSessionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the authenticated partner's KYC sessions. Filter the list by
session status or customer account, and use `page` and `size` to
paginate the results.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ListKycSessionsRequest{}
client.PartnerAPI.ListKycSessions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**status:** `*vobiz.ListKycSessionsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**accountAuthID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.CreateKycSession(request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Triggers VoBiz to email a KYC link to the customer. KYC is OTP-based
(PAN + Aadhaar via DigiLocker for individuals, PAN + GSTIN for
companies). No document uploads required.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.CreateKycSessionRequest{
        AccountAuthID: "MA_ZKITB8Z2",
    }
client.PartnerAPI.CreateKycSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountAuthID:** `string` — Customer's auth_id (from create-customer-account).
    
</dd>
</dl>

<dl>
<dd>

**flowType:** `*vobiz.CreateKycSessionRequestFlowType` 

Delivery mode. `email` (default) emails the customer the KYC link.
`redirect` returns a `widget_url` in the response for immediate redirect.
    
</dd>
</dl>

<dl>
<dd>

**customerEmail:** `*string` — Required when `flow_type` is `email`. Ignored otherwise.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `*string` 

Required when `flow_type` is `redirect`. After verification the customer's
browser is sent to this URL with query params `session_id`, `status`, `auth_id`.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — VoBiz POSTs the KYC result here.
    
</dd>
</dl>

<dl>
<dd>

**expiresInDays:** `*int` — Days before the KYC link expires.
    
</dd>
</dl>

<dl>
<dd>

**reminderSchedule:** `[]*vobiz.CreateKycSessionRequestReminderScheduleItem` — Auto reminder emails before expiry. Email flow only.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Free-form key/value object echoed back on GET and webhooks.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.GetKycSession(SessionID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the current status and available details for one KYC session
owned by the authenticated partner.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.GetKycSessionRequest{
        SessionID: "session_id",
    }
client.PartnerAPI.GetKycSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.RevokeKycSession(SessionID, request) -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels an outstanding KYC session. Customer can no longer use the link.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.RevokeKycSessionRequest{
        SessionID: "session_id",
    }
client.PartnerAPI.RevokeKycSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PartnerAPI.ResendKycSession(SessionID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Re-dispatches the KYC link to the customer. Rate-limited to once per 30 minutes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &vobiz.ResendKycSessionRequest{
        SessionID: "session_id",
    }
client.PartnerAPI.ResendKycSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

