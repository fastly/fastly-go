# RealtimeAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatsLast120Seconds**](RealtimeAPI.md#GetStatsLast120Seconds) | **GET** `/v1/channel/{service_id}/ts/h` | Get real-time data for the last 120 seconds
[**GetStatsLast120SecondsLimitEntries**](RealtimeAPI.md#GetStatsLast120SecondsLimitEntries) | **GET** `/v1/channel/{service_id}/ts/h/limit/{max_entries}` | Get a limited number of real-time data entries
[**GetStatsLastSecond**](RealtimeAPI.md#GetStatsLastSecond) | **GET** `/v1/channel/{service_id}/ts/{timestamp_in_seconds}` | Get real-time data from specified time



## GetStatsLast120Seconds

Get real-time data for the last 120 seconds



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
    resp, r, err := apiClient.RealtimeAPI.GetStatsLast120Seconds(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealtimeAPI.GetStatsLast120Seconds`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsLast120Seconds`: Realtime
    fmt.Fprintf(os.Stdout, "Response from `RealtimeAPI.GetStatsLast120Seconds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsLast120SecondsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Realtime**](Realtime.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetStatsLast120SecondsLimitEntries

Get a limited number of real-time data entries



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
    maxEntries := int32(1) // int32 | Maximum number of results to show.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RealtimeAPI.GetStatsLast120SecondsLimitEntries(ctx, serviceID, maxEntries).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealtimeAPI.GetStatsLast120SecondsLimitEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsLast120SecondsLimitEntries`: Realtime
    fmt.Fprintf(os.Stdout, "Response from `RealtimeAPI.GetStatsLast120SecondsLimitEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**maxEntries** | **int32** | Maximum number of results to show. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsLast120SecondsLimitEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Realtime**](Realtime.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetStatsLastSecond

Get real-time data from specified time



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
    timestampInSeconds := int32(56) // int32 | Timestamp in seconds (Unix epoch time).

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RealtimeAPI.GetStatsLastSecond(ctx, serviceID, timestampInSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealtimeAPI.GetStatsLastSecond`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsLastSecond`: Realtime
    fmt.Fprintf(os.Stdout, "Response from `RealtimeAPI.GetStatsLastSecond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**timestampInSeconds** | **int32** | Timestamp in seconds (Unix epoch time). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsLastSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Realtime**](Realtime.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
