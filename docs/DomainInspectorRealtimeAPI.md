# DomainInspectorRealtimeAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDomainInspectorLast120Seconds**](DomainInspectorRealtimeAPI.md#GetDomainInspectorLast120Seconds) | **GET** `/v1/domains/{service_id}/ts/h` | Get real-time domain data for the last 120 seconds
[**GetDomainInspectorLastMaxEntries**](DomainInspectorRealtimeAPI.md#GetDomainInspectorLastMaxEntries) | **GET** `/v1/domains/{service_id}/ts/h/limit/{max_entries}` | Get a limited number of real-time domain data entries
[**GetDomainInspectorLastSecond**](DomainInspectorRealtimeAPI.md#GetDomainInspectorLastSecond) | **GET** `/v1/domains/{service_id}/ts/{start_timestamp}` | Get real-time domain data from a specified time



## GetDomainInspectorLast120Seconds

Get real-time domain data for the last 120 seconds



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainInspectorRealtimeAPI.GetDomainInspectorLast120Seconds(ctx, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainInspectorRealtimeAPI.GetDomainInspectorLast120Seconds`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainInspectorLast120Seconds`: DomainInspector
    fmt.Fprintf(os.Stdout, "Response from `DomainInspectorRealtimeAPI.GetDomainInspectorLast120Seconds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainInspectorLast120SecondsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainInspector**](DomainInspector.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDomainInspectorLastMaxEntries

Get a limited number of real-time domain data entries



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    maxEntries := int32(1) // int32 | Maximum number of results to show.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainInspectorRealtimeAPI.GetDomainInspectorLastMaxEntries(ctx, serviceId, maxEntries).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainInspectorRealtimeAPI.GetDomainInspectorLastMaxEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainInspectorLastMaxEntries`: DomainInspector
    fmt.Fprintf(os.Stdout, "Response from `DomainInspectorRealtimeAPI.GetDomainInspectorLastMaxEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**maxEntries** | **int32** | Maximum number of results to show. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainInspectorLastMaxEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainInspector**](DomainInspector.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDomainInspectorLastSecond

Get real-time domain data from a specified time



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    startTimestamp := int32(56) // int32 | Timestamp in seconds (Unix epoch time).

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainInspectorRealtimeAPI.GetDomainInspectorLastSecond(ctx, serviceId, startTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainInspectorRealtimeAPI.GetDomainInspectorLastSecond`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainInspectorLastSecond`: DomainInspector
    fmt.Fprintf(os.Stdout, "Response from `DomainInspectorRealtimeAPI.GetDomainInspectorLastSecond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**startTimestamp** | **int32** | Timestamp in seconds (Unix epoch time). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainInspectorLastSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainInspector**](DomainInspector.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

