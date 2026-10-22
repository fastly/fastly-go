# AiRuntimeControlUsageMetricsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArcExportUsageMetrics**](AiRuntimeControlUsageMetricsAPI.md#ArcExportUsageMetrics) | **GET** `/ai-runtime-control/v1/usage-metrics/export` | Export usage metrics
[**ArcGetUsageMetrics**](AiRuntimeControlUsageMetricsAPI.md#ArcGetUsageMetrics) | **GET** `/ai-runtime-control/v1/usage-metrics` | Get usage metrics



## ArcExportUsageMetrics

Export usage metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    key := "7fVVQb1nZaJgnzCn4B4v3B" // string | Filter results to a specific virtual key ID. (optional)
    provider := "pTfnUUBcU1yv8SoRJn5PZ2" // string | Filter results by AI provider connection ID. (optional)
    model := "claude-sonnet-4-20250514" // string | Filter results by AI model identifier. (optional)
    from := time.Now() // time.Time | The start of the time range, in RFC 3339 format. (optional)
    to := time.Now() // time.Time | The end of the time range, in RFC 3339 format. Defaults to now. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlUsageMetricsAPI.ArcExportUsageMetrics(ctx).Key(key).Provider(provider).Model(model).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlUsageMetricsAPI.ArcExportUsageMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcExportUsageMetrics`: ArcUsageMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlUsageMetricsAPI.ArcExportUsageMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcExportUsageMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Filter results to a specific virtual key ID. |  **provider** | **string** | Filter results by AI provider connection ID. |  **model** | **string** | Filter results by AI model identifier. |  **from** | **time.Time** | The start of the time range, in RFC 3339 format. |  **to** | **time.Time** | The end of the time range, in RFC 3339 format. Defaults to now. | 

### Return type

[**ArcUsageMetricsResponse**](ArcUsageMetricsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcGetUsageMetrics

Get usage metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    key := "7fVVQb1nZaJgnzCn4B4v3B" // string | Filter results to a specific virtual key ID. (optional)
    provider := "pTfnUUBcU1yv8SoRJn5PZ2" // string | Filter results by AI provider connection ID. (optional)
    model := "claude-sonnet-4-20250514" // string | Filter results by AI model identifier. (optional)
    from := time.Now() // time.Time | The start of the time range, in RFC 3339 format. (optional)
    to := time.Now() // time.Time | The end of the time range, in RFC 3339 format. Defaults to now. (optional)
    granularity := "daily" // string | The time granularity to aggregate usage by. (optional) (default to "daily")
    sort := "date" // string | Sort field. Prefix with `-` for descending order (e.g., `sort=-date`).  (optional) (default to "date")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlUsageMetricsAPI.ArcGetUsageMetrics(ctx).Key(key).Provider(provider).Model(model).From(from).To(to).Granularity(granularity).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlUsageMetricsAPI.ArcGetUsageMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetUsageMetrics`: ArcUsageMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlUsageMetricsAPI.ArcGetUsageMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcGetUsageMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Filter results to a specific virtual key ID. |  **provider** | **string** | Filter results by AI provider connection ID. |  **model** | **string** | Filter results by AI model identifier. |  **from** | **time.Time** | The start of the time range, in RFC 3339 format. |  **to** | **time.Time** | The end of the time range, in RFC 3339 format. Defaults to now. |  **granularity** | **string** | The time granularity to aggregate usage by. | [default to &quot;daily&quot;] **sort** | **string** | Sort field. Prefix with `-` for descending order (e.g., `sort&#x3D;-date`).  | [default to &quot;date&quot;]

### Return type

[**ArcUsageMetricsResponse**](ArcUsageMetricsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

