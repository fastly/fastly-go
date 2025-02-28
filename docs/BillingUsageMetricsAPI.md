# BillingUsageMetricsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceLevelUsage**](BillingUsageMetricsAPI.md#GetServiceLevelUsage) | **GET** `/billing/v3/service-usage-metrics` | Retrieve service-level usage metrics for a product.
[**GetUsageMetrics**](BillingUsageMetricsAPI.md#GetUsageMetrics) | **GET** `/billing/v3/usage-metrics` | Get monthly usage metrics



## GetServiceLevelUsage

Retrieve service-level usage metrics for a product.



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
    productID := "productId_example" // string | The product identifier for the metrics returned (e.g., `cdn_usage`). (optional)
    usageTypeName := "usageTypeName_example" // string | The usage type name for the metrics returned (e.g., `North America Requests`). (optional)
    startMonth := "2023-01" // string |  (optional)
    endMonth := "2023-03" // string |  (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 100. (optional) (default to "5")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingUsageMetricsAPI.GetServiceLevelUsage(ctx).ProductID(productID).UsageTypeName(usageTypeName).StartMonth(startMonth).EndMonth(endMonth).Limit(limit).Cursor(cursor).Execute()
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
 **productID** | **string** | The product identifier for the metrics returned (e.g., `cdn_usage`). |  **usageTypeName** | **string** | The usage type name for the metrics returned (e.g., `North America Requests`). |  **startMonth** | **string** |  |  **endMonth** | **string** |  |  **limit** | **string** | Number of results per page. The maximum is 100. | [default to &quot;5&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. | 

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
