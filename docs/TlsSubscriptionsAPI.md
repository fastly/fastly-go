# TlsSubscriptionsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalsignEmailChallenge**](TlsSubscriptionsAPI.md#CreateGlobalsignEmailChallenge) | **POST** `/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges` | Creates a GlobalSign email challenge.
[**CreateTlsSub**](TlsSubscriptionsAPI.md#CreateTlsSub) | **POST** `/tls/subscriptions` | Create a TLS subscription
[**DeleteGlobalsignEmailChallenge**](TlsSubscriptionsAPI.md#DeleteGlobalsignEmailChallenge) | **DELETE** `/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges/{globalsign_email_challenge_id}` | Delete a GlobalSign email challenge
[**DeleteTlsSub**](TlsSubscriptionsAPI.md#DeleteTlsSub) | **DELETE** `/tls/subscriptions/{tls_subscription_id}` | Delete a TLS subscription
[**GetTlsSub**](TlsSubscriptionsAPI.md#GetTlsSub) | **GET** `/tls/subscriptions/{tls_subscription_id}` | Get a TLS subscription
[**ListTlsSubs**](TlsSubscriptionsAPI.md#ListTlsSubs) | **GET** `/tls/subscriptions` | List TLS subscriptions
[**PatchTlsSub**](TlsSubscriptionsAPI.md#PatchTlsSub) | **PATCH** `/tls/subscriptions/{tls_subscription_id}` | Update a TLS subscription



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
    tlsSubscriptionId := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    tlsAuthorizationId := "tlsAuthorizationId_example" // string | Alphanumeric string identifying a TLS subscription.
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.CreateGlobalsignEmailChallenge(ctx, tlsSubscriptionId, tlsAuthorizationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.CreateGlobalsignEmailChallenge`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalsignEmailChallenge`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TlsSubscriptionsAPI.CreateGlobalsignEmailChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionId** | **string** | Alphanumeric string identifying a TLS subscription. | 
**tlsAuthorizationId** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalsignEmailChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateTlsSub

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
    tlsSubscription := *openapiclient.NewTlsSubscription() // TlsSubscription |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.CreateTlsSub(ctx).TlsSubscription(tlsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.CreateTlsSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTlsSub`: TlsSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsSubscriptionsAPI.CreateTlsSub`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsSubscription** | [**TlsSubscription**](TlsSubscription.md) |  | 

### Return type

[**TlsSubscriptionResponse**](TlsSubscriptionResponse.md)

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
    tlsSubscriptionId := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    tlsAuthorizationId := "tlsAuthorizationId_example" // string | Alphanumeric string identifying a TLS subscription.
    globalsignEmailChallengeId := "gU3guUGZzb2W9Euo4Mo0r" // string | Alphanumeric string identifying a GlobalSign email challenge.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.DeleteGlobalsignEmailChallenge(ctx, tlsSubscriptionId, tlsAuthorizationId, globalsignEmailChallengeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.DeleteGlobalsignEmailChallenge`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionId** | **string** | Alphanumeric string identifying a TLS subscription. | 
**tlsAuthorizationId** | **string** | Alphanumeric string identifying a TLS subscription. | 
**globalsignEmailChallengeId** | **string** | Alphanumeric string identifying a GlobalSign email challenge. | 

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


## DeleteTlsSub

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
    tlsSubscriptionId := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.DeleteTlsSub(ctx, tlsSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.DeleteTlsSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionId** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsSubRequest struct via the builder pattern


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


## GetTlsSub

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
    tlsSubscriptionId := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    include := "tls_authorizations" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.GetTlsSub(ctx, tlsSubscriptionId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.GetTlsSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsSub`: TlsSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsSubscriptionsAPI.GetTlsSub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionId** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  | 

### Return type

[**TlsSubscriptionResponse**](TlsSubscriptionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTlsSubs

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
    filterTlsDomainsId := "filterTlsDomainsId_example" // string | Limit the returned subscriptions to those that include the specific domain. (optional)
    filterHasActiveOrder := true // bool | Limit the returned subscriptions to those that have currently active orders. Permitted values: `true`.  (optional)
    filterCertificateAuthority := "filterCertificateAuthority_example" // string | Limit the returned subscriptions to a specific certification authority. Values may include `certainly`, `lets-encrypt`, or `globalsign`.  (optional)
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "-created_at")
    include := "tls_authorizations" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.ListTlsSubs(ctx).FilterState(filterState).FilterTlsDomainsId(filterTlsDomainsId).FilterHasActiveOrder(filterHasActiveOrder).FilterCertificateAuthority(filterCertificateAuthority).Sort(sort).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.ListTlsSubs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsSubs`: TlsSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsSubscriptionsAPI.ListTlsSubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsSubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterState** | **string** | Limit the returned subscriptions by state. Valid values are `pending`, `processing`, `issued`, `renewing`, and `failed`. Accepts parameters: `not` (e.g., `filter[state][not]&#x3D;renewing`).  |  **filterTlsDomainsId** | **string** | Limit the returned subscriptions to those that include the specific domain. |  **filterHasActiveOrder** | **bool** | Limit the returned subscriptions to those that have currently active orders. Permitted values: `true`.  |  **filterCertificateAuthority** | **string** | Limit the returned subscriptions to a specific certification authority. Values may include `certainly`, `lets-encrypt`, or `globalsign`.  |  **sort** | **string** | The order in which to list the results. | [default to &quot;-created_at&quot;] **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, `tls_authorizations.self_managed_http_challenge`, and `tls_certificates`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TlsSubscriptionsResponse**](TlsSubscriptionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## PatchTlsSub

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
    tlsSubscriptionId := "tlsSubscriptionId_example" // string | Alphanumeric string identifying a TLS subscription.
    force := true // bool | A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.  (optional)
    tlsSubscription := *openapiclient.NewTlsSubscription() // TlsSubscription |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsSubscriptionsAPI.PatchTlsSub(ctx, tlsSubscriptionId).Force(force).TlsSubscription(tlsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsSubscriptionsAPI.PatchTlsSub`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTlsSub`: TlsSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsSubscriptionsAPI.PatchTlsSub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsSubscriptionId** | **string** | Alphanumeric string identifying a TLS subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTlsSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain.  |  **tlsSubscription** | [**TlsSubscription**](TlsSubscription.md) |  | 

### Return type

[**TlsSubscriptionResponse**](TlsSubscriptionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

