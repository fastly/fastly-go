# OriginInspectorHistoricalAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOriginInspectorHistorical**](OriginInspectorHistoricalAPI.md#GetOriginInspectorHistorical) | **GET** `/metrics/origins/services/{service_id}` | Get historical origin data for a service



## GetOriginInspectorHistorical

Get historical origin data for a service



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
    start := "2021-08-01T00:00:00.000Z" // string | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `downsample` value. (optional)
    end := "2020-08-02T00:00:00.000Z" // string | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `downsample` value. (optional)
    downsample := "hour" // string | Duration of sample windows. (optional) (default to "hour")
    metric := "resp_body_bytes,status_2xx" // string | The metrics to retrieve. Multiple values should be comma-separated. (optional) (default to "responses")
    groupBy := "host" // string | Dimensions to return in the query. Multiple dimensions may be separated by commas. For example, `group_by=host` will return one timeseries for every origin host, as a total across all POPs.  (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 200. (optional) (default to "100")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    region := "usa" // string | Limit query to one or more specific geographic regions. Values should be comma-separated.  (optional)
    datacenter := "SJC,STP" // string | Limit query to one or more specific POPs. Values should be comma-separated. (optional)
    host := "origin_1,origin_2" // string | Limit query to one or more specific origin hosts. Values should be comma-separated. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.OriginInspectorHistoricalAPI.GetOriginInspectorHistorical(ctx, serviceID).Start(start).End(end).Downsample(downsample).Metric(metric).GroupBy(groupBy).Limit(limit).Cursor(cursor).Region(region).Datacenter(datacenter).Host(host).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginInspectorHistoricalAPI.GetOriginInspectorHistorical`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOriginInspectorHistorical`: HistoricalOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `OriginInspectorHistoricalAPI.GetOriginInspectorHistorical`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginInspectorHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `downsample` value. |  **end** | **string** | A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `downsample` value. |  **downsample** | **string** | Duration of sample windows. | [default to &quot;hour&quot;] **metric** | **string** | The metrics to retrieve. Multiple values should be comma-separated. | [default to &quot;responses&quot;] **groupBy** | **string** | Dimensions to return in the query. Multiple dimensions may be separated by commas. For example, `group_by&#x3D;host` will return one timeseries for every origin host, as a total across all POPs.  |  **limit** | **string** | Number of results per page. The maximum is 200. | [default to &quot;100&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **region** | **string** | Limit query to one or more specific geographic regions. Values should be comma-separated.  |  **datacenter** | **string** | Limit query to one or more specific POPs. Values should be comma-separated. |  **host** | **string** | Limit query to one or more specific origin hosts. Values should be comma-separated. | 

### Return type

[**HistoricalOriginsResponse**](HistoricalOriginsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
