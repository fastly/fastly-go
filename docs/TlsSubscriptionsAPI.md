# TLSSubscriptionsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalsignEmailChallenge**](TlsSubscriptionsAPI.md#CreateGlobalsignEmailChallenge) | **POST** `/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges` | Creates a GlobalSign email challenge.
[**CreateTLSSub**](TlsSubscriptionsAPI.md#CreateTLSSub) | **POST** `/tls/subscriptions` | Create a TLS subscription
[**DeleteGlobalsignEmailChallenge**](TlsSubscriptionsAPI.md#DeleteGlobalsignEmailChallenge) | **DELETE** `/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges/{globalsign_email_challenge_id}` | Delete a GlobalSign email challenge
[**DeleteTLSSub**](TlsSubscriptionsAPI.md#DeleteTLSSub) | **DELETE** `/tls/subscriptions/{tls_subscription_id}` | Delete a TLS subscription
[**GetTLSSub**](TlsSubscriptionsAPI.md#GetTLSSub) | **GET** `/tls/subscriptions/{tls_subscription_id}` | Get a TLS subscription
[**ListTLSSubs**](TlsSubscriptionsAPI.md#ListTLSSubs) | **GET** `/tls/subscriptions` | List TLS subscriptions
[**PatchTLSSub**](TlsSubscriptionsAPI.md#PatchTLSSub) | **PATCH** `/tls/subscriptions/{tls_subscription_id}` | Update a TLS subscription



## CreateGlobalsignEmailChallenge

Creates a GlobalSign email challenge.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    tlsSubscriptionID := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    tlsAuthorizationID := "tlsAuthorizationId_example" // string | Alphanumeric string identifying a TLS subscription.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.CreateGlobalsignEmailChallenge(ctx, tlsSubscriptionID, tlsAuthorizationID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.CreateGlobalsignEmailChallenge`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalsignEmailChallenge`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `TLSSubscriptionsAPI.CreateGlobalsignEmailChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionID** | **string** | Alphanumeric string identifying a TLS subscription. | 
**tlsAuthorizationID** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalsignEmailChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateTLSSub

Create a TLS subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    force := true // bool | A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.  (optional)
    tlsSubscription := *openapiclient.NewTLSSubscription() // TLSSubscription |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.CreateTLSSub(ctx).Force(force).TLSSubscription(tlsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.CreateTLSSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTLSSub`: TLSSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSSubscriptionsAPI.CreateTLSSub`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTLSSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.  |  **tlsSubscription** | [**TLSSubscription**](TlsSubscription.md) |  | 

### Return type

[**TLSSubscriptionResponse**](TlsSubscriptionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteGlobalsignEmailChallenge

Delete a GlobalSign email challenge



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    tlsSubscriptionID := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    tlsAuthorizationID := "tlsAuthorizationId_example" // string | Alphanumeric string identifying a TLS subscription.
    globalsignEmailChallengeID := "gU3guUGZzb2W9Euo4Mo0r" // string | Alphanumeric string identifying a GlobalSign email challenge.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.DeleteGlobalsignEmailChallenge(ctx, tlsSubscriptionID, tlsAuthorizationID, globalsignEmailChallengeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.DeleteGlobalsignEmailChallenge`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionID** | **string** | Alphanumeric string identifying a TLS subscription. | 
**tlsAuthorizationID** | **string** | Alphanumeric string identifying a TLS subscription. | 
**globalsignEmailChallengeID** | **string** | Alphanumeric string identifying a GlobalSign email challenge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalsignEmailChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTLSSub

Delete a TLS subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    tlsSubscriptionID := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.DeleteTLSSub(ctx, tlsSubscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.DeleteTLSSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionID** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTLSSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetTLSSub

Get a TLS subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    tlsSubscriptionID := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    include := "tls_authorizations" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.GetTLSSub(ctx, tlsSubscriptionID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.GetTLSSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLSSub`: TLSSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSSubscriptionsAPI.GetTLSSub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionID** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  | 

### Return type

[**TLSSubscriptionResponse**](TlsSubscriptionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTLSSubs

List TLS subscriptions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    filterState := "filterState_example" // string | Limit the returned subscriptions by state. Valid values are `pending`, `processing`, `issued`, `renewing`, and `failed`. Accepts parameters: `not` (e.g., `filter[state][not]=renewing`).  (optional)
    filterTLSDomainsID := "filterTLSDomainsId_example" // string | Limit the returned subscriptions to those that include the specific domain. (optional)
    filterHasActiveOrder := true // bool | Limit the returned subscriptions to those that have currently active orders. Permitted values: `true`.  (optional)
    filterCertificateAuthority := "filterCertificateAuthority_example" // string | Limit the returned subscriptions to a specific certification authority. Values may include `certainly`, `lets-encrypt`, or `globalsign`.  (optional)
    include := "tls_authorizations" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.ListTLSSubs(ctx).FilterState(filterState).FilterTLSDomainsID(filterTLSDomainsID).FilterHasActiveOrder(filterHasActiveOrder).FilterCertificateAuthority(filterCertificateAuthority).Include(include).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.ListTLSSubs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSSubs`: TLSSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSSubscriptionsAPI.ListTLSSubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSSubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterState** | **string** | Limit the returned subscriptions by state. Valid values are `pending`, `processing`, `issued`, `renewing`, and `failed`. Accepts parameters: `not` (e.g., `filter[state][not]&#x3D;renewing`).  |  **filterTLSDomainsID** | **string** | Limit the returned subscriptions to those that include the specific domain. |  **filterHasActiveOrder** | **bool** | Limit the returned subscriptions to those that have currently active orders. Permitted values: `true`.  |  **filterCertificateAuthority** | **string** | Limit the returned subscriptions to a specific certification authority. Values may include `certainly`, `lets-encrypt`, or `globalsign`.  |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**TLSSubscriptionsResponse**](TlsSubscriptionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## PatchTLSSub

Update a TLS subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    tlsSubscriptionID := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    force := true // bool | A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.  (optional)
    tlsSubscription := *openapiclient.NewTLSSubscription() // TLSSubscription |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSSubscriptionsAPI.PatchTLSSub(ctx, tlsSubscriptionID).Force(force).TLSSubscription(tlsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSSubscriptionsAPI.PatchTLSSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTLSSub`: TLSSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSSubscriptionsAPI.PatchTLSSub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionID** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTLSSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.  |  **tlsSubscription** | [**TLSSubscription**](TlsSubscription.md) |  | 

### Return type

[**TLSSubscriptionResponse**](TlsSubscriptionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
