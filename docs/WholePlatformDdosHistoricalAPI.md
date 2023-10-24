# WholePlatformDdosHistoricalAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlatformDdosHistorical**](WholePlatformDdosHistoricalAPI.md#GetPlatformDdosHistorical) | **GET** `/metrics/platform/ddos` | Get historical DDoS metrics for the entire Fastly platform



## GetPlatformDdosHistorical

Get historical DDoS metrics for the entire Fastly platform



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
    start := "2021-08-01T00:00:00.000Z" // string | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `downsample` value. (optional)
    end := "2020-08-02T00:00:00.000Z" // string | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `downsample` value. (optional)
    downsample := "hour" // string | Duration of sample windows. (optional) (default to "hour")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WholePlatformDdosHistoricalAPI.GetPlatformDdosHistorical(ctx).Start(start).End(end).Downsample(downsample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholePlatformDdosHistoricalAPI.GetPlatformDdosHistorical`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlatformDdosHistorical`: PlatformDdosResponse
    fmt.Fprintf(os.Stdout, "Response from `WholePlatformDdosHistoricalAPI.GetPlatformDdosHistorical`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformDdosHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `downsample` value. |  **end** | **string** | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `downsample` value. |  **downsample** | **string** | Duration of sample windows. | [default to &quot;hour&quot;]

### Return type

[**PlatformDdosResponse**](PlatformDdosResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
