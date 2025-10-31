# StatsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceStats**](StatsAPI.md#GetServiceStats) | **GET** `/service/{service_id}/stats/summary` | Get stats for a service



## GetServiceStats

Get stats for a service



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
    month := "05" // string | 2-digit month. (optional)
    year := "2020" // string | 4-digit year. (optional)
    startTime := int32(1608560817) // int32 | Epoch timestamp. Limits the results returned. (optional)
    endTime := int32(1608560817) // int32 | Epoch timestamp. Limits the results returned. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.StatsAPI.GetServiceStats(ctx, serviceId).Month(month).Year(year).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetServiceStats`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceStats`: Stats
    fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetServiceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | 2-digit month. |  **year** | **string** | 4-digit year. |  **startTime** | **int32** | Epoch timestamp. Limits the results returned. |  **endTime** | **int32** | Epoch timestamp. Limits the results returned. | 

### Return type

[**Stats**](Stats.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

