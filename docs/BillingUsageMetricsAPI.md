# BillingUsageMetricsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceLevelUsage**](BillingUsageMetricsAPI.md#GetServiceLevelUsage) | **GET** `/billing/v2/account_customers/{customer_id}/service-usage-metrics` | Retrieve service-level usage metrics for a product.
[**GetServiceLevelUsageTypes**](BillingUsageMetricsAPI.md#GetServiceLevelUsageTypes) | **GET** `/billing/v2/account_customers/{customer_id}/service-usage-types` | Retrieve product usage types for a customer.



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
    customerID := "customerId_example" // string | Alphanumeric string identifying the customer.
    productID := "productId_example" // string | The product identifier for the metrics returned (e.g., `cdn_usage`). This field is not required for CSV requests.
    usageTypeName := "usageTypeName_example" // string | The usage type name for the metrics returned (e.g., `North America Requests`). This field is not required for CSV requests.
    timeGranularity := "timeGranularity_example" // string | 
    startDate := "2023-01-01" // string |  (optional)
    endDate := "2023-01-31" // string |  (optional)
    startMonth := "2023-01" // string |  (optional)
    endMonth := "2023-03" // string |  (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 100. (optional) (default to "5")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingUsageMetricsAPI.GetServiceLevelUsage(ctx, customerID).ProductID(productID).UsageTypeName(usageTypeName).TimeGranularity(timeGranularity).StartDate(startDate).EndDate(endDate).StartMonth(startMonth).EndMonth(endMonth).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingUsageMetricsAPI.GetServiceLevelUsage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceLevelUsage`: Serviceusagemetrics
    fmt.Fprintf(os.Stdout, "Response from `BillingUsageMetricsAPI.GetServiceLevelUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceLevelUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productID** | **string** | The product identifier for the metrics returned (e.g., `cdn_usage`). This field is not required for CSV requests. |  **usageTypeName** | **string** | The usage type name for the metrics returned (e.g., `North America Requests`). This field is not required for CSV requests. |  **timeGranularity** | **string** |  |  **startDate** | **string** |  |  **endDate** | **string** |  |  **startMonth** | **string** |  |  **endMonth** | **string** |  |  **limit** | **string** | Number of results per page. The maximum is 100. | [default to &quot;5&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. | 

### Return type

[**Serviceusagemetrics**](Serviceusagemetrics.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetServiceLevelUsageTypes

Retrieve product usage types for a customer.



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
    customerID := "customerId_example" // string | Alphanumeric string identifying the customer.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingUsageMetricsAPI.GetServiceLevelUsageTypes(ctx, customerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingUsageMetricsAPI.GetServiceLevelUsageTypes`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceLevelUsageTypes`: Serviceusagetypes
    fmt.Fprintf(os.Stdout, "Response from `BillingUsageMetricsAPI.GetServiceLevelUsageTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceLevelUsageTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Serviceusagetypes**](Serviceusagetypes.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
