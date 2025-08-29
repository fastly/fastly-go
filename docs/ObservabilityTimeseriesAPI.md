# ObservabilityTimeseriesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimeseriesGet**](ObservabilityTimeseriesAPI.md#TimeseriesGet) | **GET** `/observability/timeseries` | Retrieve observability data as a time series



## TimeseriesGet

Retrieve observability data as a time series



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
    source := "logs" // string | 
    from := "2024-01-03T16:00:00Z" // string | 
    to := "2024-01-03T18:00:00Z" // string | 
    granularity := "hour" // string | 
    series := "avg[response_time],p99[response_time]" // string | 
    dimensions := "dimensions_example" // string |  (optional)
    filter := "filter[response_status]=200" // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObservabilityTimeseriesAPI.TimeseriesGet(ctx).Source(source).From(from).To(to).Granularity(granularity).Series(series).Dimensions(dimensions).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityTimeseriesAPI.TimeseriesGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeseriesGet`: TimeseriesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ObservabilityTimeseriesAPI.TimeseriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimeseriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** |  |  **from** | **string** |  |  **to** | **string** |  |  **granularity** | **string** |  |  **series** | **string** |  |  **dimensions** | **string** |  |  **filter** | **string** |  | 

### Return type

[**TimeseriesGetResponse**](TimeseriesGetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
