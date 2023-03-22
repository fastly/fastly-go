# HistoricalAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistStats**](HistoricalAPI.md#GetHistStats) | **GET** `/stats` | Get historical stats
[**GetHistStatsAggregated**](HistoricalAPI.md#GetHistStatsAggregated) | **GET** `/stats/aggregate` | Get aggregated historical stats
[**GetHistStatsField**](HistoricalAPI.md#GetHistStatsField) | **GET** `/stats/field/{field}` | Get historical stats for a single field
[**GetHistStatsService**](HistoricalAPI.md#GetHistStatsService) | **GET** `/stats/service/{service_id}` | Get historical stats for a single service
[**GetHistStatsServiceField**](HistoricalAPI.md#GetHistStatsServiceField) | **GET** `/stats/service/{service_id}/field/{field}` | Get historical stats for a single service/field combination
[**GetRegions**](HistoricalAPI.md#GetRegions) | **GET** `/stats/regions` | Get region codes
[**GetUsage**](HistoricalAPI.md#GetUsage) | **GET** `/stats/usage` | Get usage statistics
[**GetUsageMonth**](HistoricalAPI.md#GetUsageMonth) | **GET** `/stats/usage_by_month` | Get month-to-date usage statistics
[**GetUsageService**](HistoricalAPI.md#GetUsageService) | **GET** `/stats/usage_by_service` | Get usage statistics per service



## GetHistStats

Get historical stats



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
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")
    by := "day" // string | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  (optional) (default to "day")
    region := "usa" // string | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetHistStats(ctx).From(from).To(to).By(by).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetHistStats`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistStats`: HistoricalResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetHistStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;] **by** | **string** | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  | [default to &quot;day&quot;] **region** | **string** | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  | 

### Return type

[**HistoricalResponse**](HistoricalResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetHistStatsAggregated

Get aggregated historical stats



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
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")
    by := "day" // string | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  (optional) (default to "day")
    region := "usa" // string | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetHistStatsAggregated(ctx).From(from).To(to).By(by).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetHistStatsAggregated`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistStatsAggregated`: HistoricalAggregateResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetHistStatsAggregated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistStatsAggregatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;] **by** | **string** | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  | [default to &quot;day&quot;] **region** | **string** | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  | 

### Return type

[**HistoricalAggregateResponse**](HistoricalAggregateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetHistStatsField

Get historical stats for a single field



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
    field := "hit_ratio" // string | Name of the stats field.
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")
    by := "day" // string | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  (optional) (default to "day")
    region := "usa" // string | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetHistStatsField(ctx, field).From(from).To(to).By(by).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetHistStatsField`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistStatsField`: HistoricalFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetHistStatsField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**field** | **string** | Name of the stats field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistStatsFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;] **by** | **string** | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  | [default to &quot;day&quot;] **region** | **string** | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  | 

### Return type

[**HistoricalFieldResponse**](HistoricalFieldResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetHistStatsService

Get historical stats for a single service



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
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")
    by := "day" // string | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  (optional) (default to "day")
    region := "usa" // string | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetHistStatsService(ctx, serviceID).From(from).To(to).By(by).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetHistStatsService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistStatsService`: HistoricalAggregateResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetHistStatsService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistStatsServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;] **by** | **string** | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  | [default to &quot;day&quot;] **region** | **string** | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  | 

### Return type

[**HistoricalAggregateResponse**](HistoricalAggregateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetHistStatsServiceField

Get historical stats for a single service/field combination



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
    field := "hit_ratio" // string | Name of the stats field.
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")
    by := "day" // string | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  (optional) (default to "day")
    region := "usa" // string | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetHistStatsServiceField(ctx, serviceID, field).From(from).To(to).By(by).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetHistStatsServiceField`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistStatsServiceField`: HistoricalFieldAggregateResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetHistStatsServiceField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**field** | **string** | Name of the stats field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistStatsServiceFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;] **by** | **string** | Duration of sample windows. One of:   * `hour` - Group data by hour.   * `minute` - Group data by minute.   * `day` - Group data by day.  | [default to &quot;day&quot;] **region** | **string** | Limit query to a specific geographic region. One of:   * `usa` - North America.   * `europe` - Europe.   * `anzac` - Australia and New Zealand.   * `asia` - Asia.   * `asia_india` - India.   * `asia_southkorea` - South Korea.   * `africa_std` - Africa.   * `southamerica_std` - South America.  | 

### Return type

[**HistoricalFieldAggregateResponse**](HistoricalFieldAggregateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetRegions

Get region codes



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetRegions(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetRegions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegions`: HistoricalRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern



### Return type

[**HistoricalRegionsResponse**](HistoricalRegionsResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetUsage

Get usage statistics



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
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetUsage(ctx).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetUsage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsage`: HistoricalUsageAggregateResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;]

### Return type

[**HistoricalUsageAggregateResponse**](HistoricalUsageAggregateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetUsageMonth

Get month-to-date usage statistics



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
    year := "2020" // string | 4-digit year. (optional)
    month := "05" // string | 2-digit month. (optional)
    billableUnits := true // bool | If `true`, return results as billable units. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetUsageMonth(ctx).Year(year).Month(month).BillableUnits(billableUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetUsageMonth`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageMonth`: HistoricalUsageMonthResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetUsageMonth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **string** | 4-digit year. |  **month** | **string** | 2-digit month. |  **billableUnits** | **bool** | If `true`, return results as billable units. | 

### Return type

[**HistoricalUsageMonthResponse**](HistoricalUsageMonthResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetUsageService

Get usage statistics per service



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
    from := "2020-04-09T18:14:30Z" // string | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.  (optional)
    to := "2020-04-09T18:14:30Z" // string | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  (optional) (default to "now")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HistoricalAPI.GetUsageService(ctx).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetUsageService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageService`: HistoricalUsageServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetUsageService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of `by`.  |  **to** | **string** | Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.  | [default to &quot;now&quot;]

### Return type

[**HistoricalUsageServiceResponse**](HistoricalUsageServiceResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
