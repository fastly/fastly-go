# BillingUsageMetricsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceLevelUsage**](BillingUsageMetricsAPI.md#GetServiceLevelUsage) | **GET** `/billing/v3/service-usage-metrics` | Retrieve service-level usage metrics for services with non-zero usage units.
[**GetUsageMetrics**](BillingUsageMetricsAPI.md#GetUsageMetrics) | **GET** `/billing/v3/usage-metrics` | Get monthly usage metrics



## GetServiceLevelUsage

Retrieve service-level usage metrics for services with non-zero usage units.



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
    productID := "productId_example" // string | The product identifier for the metrics returned (e.g., `cdn_usage`). This should be used along with `usage_type_name`. (optional)
    service := "service_example" // string | The service identifier for the metrics being requested. (optional)
    usageTypeName := "usageTypeName_example" // string | The usage type name for the metrics returned (e.g., `North America Requests`). This should be used along with `product_id`. (optional)
    startMonth := "2023-01" // string |  (optional)
    endMonth := "2023-03" // string |  (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 10000. (optional) (default to "1000")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingUsageMetricsAPI.GetServiceLevelUsage(ctx).ProductID(productID).Service(service).UsageTypeName(usageTypeName).StartMonth(startMonth).EndMonth(endMonth).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingUsageMetricsAPI.GetServiceLevelUsage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceLevelUsage`: Serviceusagemetrics
    fmt.Fprintf(os.Stdout, "Response from `BillingUsageMetricsAPI.GetServiceLevelUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceLevelUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productID** | **string** | The product identifier for the metrics returned (e.g., `cdn_usage`). This should be used along with `usage_type_name`. |  **service** | **string** | The service identifier for the metrics being requested. |  **usageTypeName** | **string** | The usage type name for the metrics returned (e.g., `North America Requests`). This should be used along with `product_id`. |  **startMonth** | **string** |  |  **endMonth** | **string** |  |  **limit** | **string** | Number of results per page. The maximum is 10000. | [default to &quot;1000&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. | 

### Return type

[**Serviceusagemetrics**](Serviceusagemetrics.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetUsageMetrics

Get monthly usage metrics



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
    startMonth := "2024-05" // string | 
    endMonth := "2024-06" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingUsageMetricsAPI.GetUsageMetrics(ctx).StartMonth(startMonth).EndMonth(endMonth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingUsageMetricsAPI.GetUsageMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageMetrics`: Usagemetric
    fmt.Fprintf(os.Stdout, "Response from `BillingUsageMetricsAPI.GetUsageMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startMonth** | **string** |  |  **endMonth** | **string** |  | 

### Return type

[**Usagemetric**](Usagemetric.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
