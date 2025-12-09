# MetricsPlatformAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlatformMetricsServiceHistorical**](MetricsPlatformAPI.md#GetPlatformMetricsServiceHistorical) | **GET** `/metrics/platform/services/{service_id}/{granularity}` | Get historical time series metrics for a single service



## GetPlatformMetricsServiceHistorical

Get historical time series metrics for a single service



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
    granularity := "daily" // string | Duration of sample windows.
    from := "2025-06-08T00:00:00.000Z" // string | A valid RFC-8339-formatted date and time indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `granularity` value. (optional)
    to := "2025-08-02T00:00:00.000Z" // string | A valid RFC-8339-formatted date and time indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `granularity` value. (optional)
    metric := "ttfb_edge_p95_us,ttfb_edge_p99_us" // string | The metric(s) to retrieve. Multiple values should be comma-separated. (optional)
    metricSet := "ttfb" // string | The metric set(s) to retrieve. Multiple values should be comma-separated. (optional)
    groupBy := "region" // string | Field to group_by in the query. For example, `group_by=region` will return entries for grouped by timestamp and region.  (optional)
    region := "usa" // string | Limit query to one or more specific geographic regions. Values should be comma-separated.  (optional)
    datacenter := "SJC,STP" // string | Limit query to one or more specific POPs. Values should be comma-separated. (optional)
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 10000. (optional) (default to "1000")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.MetricsPlatformAPI.GetPlatformMetricsServiceHistorical(ctx, serviceId, granularity).From(from).To(to).Metric(metric).MetricSet(metricSet).GroupBy(groupBy).Region(region).Datacenter(datacenter).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsPlatformAPI.GetPlatformMetricsServiceHistorical`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlatformMetricsServiceHistorical`: PlatformMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsPlatformAPI.GetPlatformMetricsServiceHistorical`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**granularity** | **string** | Duration of sample windows. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformMetricsServiceHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A valid RFC-8339-formatted date and time indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `granularity` value. |  **to** | **string** | A valid RFC-8339-formatted date and time indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `granularity` value. |  **metric** | **string** | The metric(s) to retrieve. Multiple values should be comma-separated. |  **metricSet** | **string** | The metric set(s) to retrieve. Multiple values should be comma-separated. |  **groupBy** | **string** | Field to group_by in the query. For example, `group_by&#x3D;region` will return entries for grouped by timestamp and region.  |  **region** | **string** | Limit query to one or more specific geographic regions. Values should be comma-separated.  |  **datacenter** | **string** | Limit query to one or more specific POPs. Values should be comma-separated. |  **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **string** | Number of results per page. The maximum is 10000. | [default to &quot;1000&quot;]

### Return type

[**PlatformMetricsResponse**](PlatformMetricsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

