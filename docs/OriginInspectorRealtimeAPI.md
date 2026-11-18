# OriginInspectorRealtimeAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOriginInspectorLast120Seconds**](OriginInspectorRealtimeAPI.md#GetOriginInspectorLast120Seconds) | **GET** `/v1/origins/{service_id}/ts/h` | Get real-time origin data for the last 120 seconds
[**GetOriginInspectorLastMaxEntries**](OriginInspectorRealtimeAPI.md#GetOriginInspectorLastMaxEntries) | **GET** `/v1/origins/{service_id}/ts/h/limit/{max_entries}` | Get a limited number of real-time origin data entries
[**GetOriginInspectorLastSecond**](OriginInspectorRealtimeAPI.md#GetOriginInspectorLastSecond) | **GET** `/v1/origins/{service_id}/ts/{start_timestamp}` | Get real-time origin data from specific time.



## GetOriginInspectorLast120Seconds

Get real-time origin data for the last 120 seconds



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.OriginInspectorRealtimeAPI.GetOriginInspectorLast120Seconds(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginInspectorRealtimeAPI.GetOriginInspectorLast120Seconds`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOriginInspectorLast120Seconds`: OriginInspector
    fmt.Fprintf(os.Stdout, "Response from `OriginInspectorRealtimeAPI.GetOriginInspectorLast120Seconds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginInspectorLast120SecondsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OriginInspector**](OriginInspector.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetOriginInspectorLastMaxEntries

Get a limited number of real-time origin data entries



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
    maxEntries := int32(1) // int32 | Maximum number of results to display.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.OriginInspectorRealtimeAPI.GetOriginInspectorLastMaxEntries(ctx, serviceID, maxEntries).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginInspectorRealtimeAPI.GetOriginInspectorLastMaxEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOriginInspectorLastMaxEntries`: OriginInspector
    fmt.Fprintf(os.Stdout, "Response from `OriginInspectorRealtimeAPI.GetOriginInspectorLastMaxEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**maxEntries** | **int32** | Maximum number of results to display. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginInspectorLastMaxEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OriginInspector**](OriginInspector.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetOriginInspectorLastSecond

Get real-time origin data from specific time.



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
    startTimestamp := int32(56) // int32 | Timestamp in seconds (Unix epoch time).

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.OriginInspectorRealtimeAPI.GetOriginInspectorLastSecond(ctx, serviceID, startTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginInspectorRealtimeAPI.GetOriginInspectorLastSecond`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOriginInspectorLastSecond`: OriginInspector
    fmt.Fprintf(os.Stdout, "Response from `OriginInspectorRealtimeAPI.GetOriginInspectorLastSecond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**startTimestamp** | **int32** | Timestamp in seconds (Unix epoch time). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginInspectorLastSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OriginInspector**](OriginInspector.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
