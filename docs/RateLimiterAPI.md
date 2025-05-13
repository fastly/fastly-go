# RateLimiterAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRateLimiter**](RateLimiterAPI.md#CreateRateLimiter) | **POST** `/service/{service_id}/version/{version_id}/rate-limiters` | Create a rate limiter
[**DeleteRateLimiter**](RateLimiterAPI.md#DeleteRateLimiter) | **DELETE** `/rate-limiters/{rate_limiter_id}` | Delete a rate limiter
[**GetRateLimiter**](RateLimiterAPI.md#GetRateLimiter) | **GET** `/rate-limiters/{rate_limiter_id}` | Get a rate limiter
[**ListRateLimiters**](RateLimiterAPI.md#ListRateLimiters) | **GET** `/service/{service_id}/version/{version_id}/rate-limiters` | List rate limiters
[**UpdateRateLimiter**](RateLimiterAPI.md#UpdateRateLimiter) | **PUT** `/rate-limiters/{rate_limiter_id}` | Update a rate limiter



## CreateRateLimiter

Create a rate limiter



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
    name := "name_example" // string | A human readable name for the rate limiting rule. (optional)
    uriDictionaryName := "uriDictionaryName_example" // string | The name of a Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited. (optional)
    httpMethods := []string{"Inner_example"} // []string | Array of HTTP methods to apply rate limiting to. (optional)
    rpsLimit := int32(56) // int32 | Upper limit of requests per second allowed by the rate limiter. (optional)
    windowSize := int32(56) // int32 | Number of seconds during which the RPS limit must be exceeded in order to trigger a violation. (optional)
    clientKey := []string{"Inner_example"} // []string | Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`. (optional)
    penaltyBoxDuration := int32(56) // int32 | Length of time in minutes that the rate limiter is in effect after the initial violation is detected. (optional)
    action := "action_example" // string | The action to take when a rate limiter violation is detected. (optional)
    responseObjectName := "responseObjectName_example" // string | Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration. (optional)
    loggerType := "loggerType_example" // string | Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries. (optional)
    featureRevision := int32(56) // int32 | Revision number of the rate limiting feature implementation. Defaults to the most recent revision. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RateLimiterAPI.CreateRateLimiter(ctx, serviceID, versionID).Name(name).URIDictionaryName(uriDictionaryName).HTTPMethods(httpMethods).RpsLimit(rpsLimit).WindowSize(windowSize).ClientKey(clientKey).PenaltyBoxDuration(penaltyBoxDuration).Action(action).ResponseObjectName(responseObjectName).LoggerType(loggerType).FeatureRevision(featureRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimiterAPI.CreateRateLimiter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRateLimiter`: RateLimiterResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimiterAPI.CreateRateLimiter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateLimiterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | A human readable name for the rate limiting rule. |  **uriDictionaryName** | **string** | The name of a Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited. |  **httpMethods** | **[]string** | Array of HTTP methods to apply rate limiting to. |  **rpsLimit** | **int32** | Upper limit of requests per second allowed by the rate limiter. |  **windowSize** | **int32** | Number of seconds during which the RPS limit must be exceeded in order to trigger a violation. |  **clientKey** | **[]string** | Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`. |  **penaltyBoxDuration** | **int32** | Length of time in minutes that the rate limiter is in effect after the initial violation is detected. |  **action** | **string** | The action to take when a rate limiter violation is detected. |  **responseObjectName** | **string** | Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration. |  **loggerType** | **string** | Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries. |  **featureRevision** | **int32** | Revision number of the rate limiting feature implementation. Defaults to the most recent revision. | 

### Return type

[**RateLimiterResponse**](RateLimiterResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateRateLimiter

Update a rate limiter



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
    name := "name_example" // string | A human readable name for the rate limiting rule. (optional)
    uriDictionaryName := "uriDictionaryName_example" // string | The name of a Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited. (optional)
    httpMethods := []string{"Inner_example"} // []string | Array of HTTP methods to apply rate limiting to. (optional)
    rpsLimit := int32(56) // int32 | Upper limit of requests per second allowed by the rate limiter. (optional)
    windowSize := int32(56) // int32 | Number of seconds during which the RPS limit must be exceeded in order to trigger a violation. (optional)
    clientKey := []string{"Inner_example"} // []string | Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`. (optional)
    penaltyBoxDuration := int32(56) // int32 | Length of time in minutes that the rate limiter is in effect after the initial violation is detected. (optional)
    action := "action_example" // string | The action to take when a rate limiter violation is detected. (optional)
    responseObjectName := "responseObjectName_example" // string | Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration. (optional)
    loggerType := "loggerType_example" // string | Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries. (optional)
    featureRevision := int32(56) // int32 | Revision number of the rate limiting feature implementation. Defaults to the most recent revision. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RateLimiterAPI.UpdateRateLimiter(ctx, rateLimiterID).Name(name).URIDictionaryName(uriDictionaryName).HTTPMethods(httpMethods).RpsLimit(rpsLimit).WindowSize(windowSize).ClientKey(clientKey).PenaltyBoxDuration(penaltyBoxDuration).Action(action).ResponseObjectName(responseObjectName).LoggerType(loggerType).FeatureRevision(featureRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimiterAPI.UpdateRateLimiter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRateLimiter`: RateLimiterResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimiterAPI.UpdateRateLimiter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateLimiterID** | **string** | Alphanumeric string identifying the rate limiter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRateLimiterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | A human readable name for the rate limiting rule. |  **uriDictionaryName** | **string** | The name of a Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited. |  **httpMethods** | **[]string** | Array of HTTP methods to apply rate limiting to. |  **rpsLimit** | **int32** | Upper limit of requests per second allowed by the rate limiter. |  **windowSize** | **int32** | Number of seconds during which the RPS limit must be exceeded in order to trigger a violation. |  **clientKey** | **[]string** | Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`. |  **penaltyBoxDuration** | **int32** | Length of time in minutes that the rate limiter is in effect after the initial violation is detected. |  **action** | **string** | The action to take when a rate limiter violation is detected. |  **responseObjectName** | **string** | Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration. |  **loggerType** | **string** | Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries. |  **featureRevision** | **int32** | Revision number of the rate limiting feature implementation. Defaults to the most recent revision. | 

### Return type

[**RateLimiterResponse**](RateLimiterResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
