# RateLimiterAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRateLimiter**](RateLimiterAPI.md#DeleteRateLimiter) | **DELETE** `/rate-limiters/{rate_limiter_id}` | Delete a rate limiter
[**GetRateLimiter**](RateLimiterAPI.md#GetRateLimiter) | **GET** `/rate-limiters/{rate_limiter_id}` | Get a rate limiter
[**ListRateLimiters**](RateLimiterAPI.md#ListRateLimiters) | **GET** `/service/{service_id}/version/{version_id}/rate-limiters` | List rate limiters



## DeleteRateLimiter

Delete a rate limiter



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
    rateLimiterID := "rateLimiterId_example" // string | Alphanumeric string identifying the rate limiter.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RateLimiterAPI.DeleteRateLimiter(ctx, rateLimiterID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimiterAPI.DeleteRateLimiter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRateLimiter`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `RateLimiterAPI.DeleteRateLimiter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateLimiterID** | **string** | Alphanumeric string identifying the rate limiter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRateLimiterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetRateLimiter

Get a rate limiter



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
    rateLimiterID := "rateLimiterId_example" // string | Alphanumeric string identifying the rate limiter.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RateLimiterAPI.GetRateLimiter(ctx, rateLimiterID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimiterAPI.GetRateLimiter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimiter`: RateLimiterResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimiterAPI.GetRateLimiter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateLimiterID** | **string** | Alphanumeric string identifying the rate limiter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimiterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RateLimiterResponse**](RateLimiterResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListRateLimiters

List rate limiters



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RateLimiterAPI.ListRateLimiters(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimiterAPI.ListRateLimiters`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRateLimiters`: []RateLimiterResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimiterAPI.ListRateLimiters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRateLimitersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RateLimiterResponse**](RateLimiterResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
