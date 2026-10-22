# AiRuntimeControlSessionsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArcGetSessionLogs**](AiRuntimeControlSessionsAPI.md#ArcGetSessionLogs) | **GET** `/ai-runtime-control/v1/sessions` | Get session logs



## ArcGetSessionLogs

Get session logs



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
    to := time.Now() // time.Time | The end of the time range, in RFC 3339 format. Defaults to `now`. (optional)
    cursor := "MTAwOjA" // string | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results, or the `prev_cursor` value to request the previous page.  (optional)
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    sort := "sort_example" // string | Sort field. Prefix with `-` for descending order (e.g., `sort=-created_at`).  (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlSessionsAPI.ArcGetSessionLogs(ctx).Key(key).Provider(provider).Model(model).From(from).To(to).Cursor(cursor).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlSessionsAPI.ArcGetSessionLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetSessionLogs`: ArcSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlSessionsAPI.ArcGetSessionLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcGetSessionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Filter results to a specific virtual key ID. |  **provider** | **string** | Filter results by AI provider connection ID. |  **model** | **string** | Filter results by AI model identifier. |  **from** | **time.Time** | The start of the time range, in RFC 3339 format. |  **to** | **time.Time** | The end of the time range, in RFC 3339 format. Defaults to `now`. |  **cursor** | **string** | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results, or the `prev_cursor` value to request the previous page.  |  **limit** | **int32** | Limit how many results are returned. | [default to 100] **sort** | **string** | Sort field. Prefix with `-` for descending order (e.g., `sort&#x3D;-created_at`).  | [default to &quot;created_at&quot;]

### Return type

[**ArcSessionsResponse**](ArcSessionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

