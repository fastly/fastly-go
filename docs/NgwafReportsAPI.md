# NgwafReportsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttacksReport**](NgwafReportsAPI.md#GetAttacksReport) | **GET** `/ngwaf/v1/reports/attacks` | Get attacks report
[**GetSignalsReport**](NgwafReportsAPI.md#GetSignalsReport) | **GET** `/ngwaf/v1/reports/signals` | Get signals report



## GetAttacksReport

Get attacks report



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
    from := time.Now() // time.Time | The start of a time range in RFC 3339 format.
    to := time.Now() // time.Time | The end of a time range in RFC 3339 format. Defaults to the current time. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NgwafReportsAPI.GetAttacksReport(ctx).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgwafReportsAPI.GetAttacksReport`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttacksReport`: ListAttackReport
    fmt.Fprintf(os.Stdout, "Response from `NgwafReportsAPI.GetAttacksReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttacksReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | The start of a time range in RFC 3339 format. |  **to** | **time.Time** | The end of a time range in RFC 3339 format. Defaults to the current time. | 

### Return type

[**ListAttackReport**](ListAttackReport.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetSignalsReport

Get signals report



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
    from := time.Now() // time.Time | The start of a time range in RFC 3339 format.
    to := time.Now() // time.Time | The end of a time range in RFC 3339 format. Defaults to the current time. (optional)
    signalType := "all" // string | The type of signal (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NgwafReportsAPI.GetSignalsReport(ctx).From(from).To(to).SignalType(signalType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgwafReportsAPI.GetSignalsReport`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignalsReport`: ListSignalReport
    fmt.Fprintf(os.Stdout, "Response from `NgwafReportsAPI.GetSignalsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | The start of a time range in RFC 3339 format. |  **to** | **time.Time** | The end of a time range in RFC 3339 format. Defaults to the current time. |  **signalType** | **string** | The type of signal | 

### Return type

[**ListSignalReport**](ListSignalReport.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
