# InsightsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogInsights**](InsightsAPI.md#GetLogInsights) | **GET** `/observability/log-insights` | Retrieve log insights



## GetLogInsights

Retrieve log insights



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
    visualization := "top-url-by-requests" // string | 
    serviceId := "1jlmtMz1ncwA0KC3TBGD0X" // string | 
    start := "2024-01-01T14:30:23Z" // string | 
    end := "2024-01-05T14:30:23Z" // string | 
    pops := "pops_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)
    domainExactMatch := true // bool |  (optional)
    limit := float32(8.14) // float32 |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.InsightsAPI.GetLogInsights(ctx).Visualization(visualization).ServiceId(serviceId).Start(start).End(end).Pops(pops).Domain(domain).DomainExactMatch(domainExactMatch).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.GetLogInsights`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogInsights`: GetLogInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.GetLogInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visualization** | **string** |  |  **serviceId** | **string** |  |  **start** | **string** |  |  **end** | **string** |  |  **pops** | **string** |  |  **domain** | **string** |  |  **domainExactMatch** | **bool** |  |  **limit** | **float32** |  | 

### Return type

[**GetLogInsightsResponse**](GetLogInsightsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

