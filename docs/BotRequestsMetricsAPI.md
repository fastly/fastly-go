# BotRequestsMetricsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBotRequestMetricsServiceHistorical**](BotRequestsMetricsAPI.md#GetBotRequestMetricsServiceHistorical) | **GET** `/metrics/bots/entities/{service_id}/{granularity}` | Get per-service historical time series metrics for bot requests broken down by bot-type.



## GetBotRequestMetricsServiceHistorical

Get per-service historical time series metrics for bot requests broken down by bot-type.



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
    granularity := "minutely" // string | Duration of sample windows.
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    from := "2025-06-08T00:00:00.000Z" // string | A valid RFC-8339-formatted date and time indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `granularity` value. (optional)
    to := "2025-08-02T00:00:00.000Z" // string | A valid RFC-8339-formatted date and time indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `granularity` value. (optional)
    groupBy := "bot_id" // string | Field to group by in the query. For example, `group_by=bot_id` will return entries grouped by `timestamp` and `bot_id`.  (optional)
    datacenter := "SJC,STP" // string | Limit query to one or more specific POPs. Values should be comma-separated. (optional)
    bot := "bot_example" // string | Filter results based by Bot ID. (optional)
    region := "usa" // string | Limit query to one or more specific geographic regions. Values should be comma-separated.  (optional)
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    metric := "status_4xx,status_5xx,ipv6,http3" // string | One or more metrics to retrieve. Multiple values should be comma-separated. (optional)
    metricSet := "content_guard" // string | The metric set to retrieve. (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 10000. (optional) (default to "1000")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotRequestsMetricsAPI.GetBotRequestMetricsServiceHistorical(ctx, granularity, serviceID).From(from).To(to).GroupBy(groupBy).Datacenter(datacenter).Bot(bot).Region(region).Cursor(cursor).Metric(metric).MetricSet(metricSet).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotRequestsMetricsAPI.GetBotRequestMetricsServiceHistorical`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBotRequestMetricsServiceHistorical`: BotRequestMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `BotRequestsMetricsAPI.GetBotRequestMetricsServiceHistorical`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**granularity** | **string** | Duration of sample windows. | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotRequestMetricsServiceHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A valid RFC-8339-formatted date and time indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided `granularity` value. |  **to** | **string** | A valid RFC-8339-formatted date and time indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided `granularity` value. |  **groupBy** | **string** | Field to group by in the query. For example, `group_by&#x3D;bot_id` will return entries grouped by `timestamp` and `bot_id`.  |  **datacenter** | **string** | Limit query to one or more specific POPs. Values should be comma-separated. |  **bot** | **string** | Filter results based by Bot ID. |  **region** | **string** | Limit query to one or more specific geographic regions. Values should be comma-separated.  |  **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **metric** | **string** | One or more metrics to retrieve. Multiple values should be comma-separated. |  **metricSet** | **string** | The metric set to retrieve. |  **limit** | **string** | Number of results per page. The maximum is 10000. | [default to &quot;1000&quot;]

### Return type

[**BotRequestMetricsResponse**](BotRequestMetricsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
