# LoggingEndpointErrorsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogEndpointErrors**](LoggingEndpointErrorsAPI.md#GetLogEndpointErrors) | **GET** `/observability/service/{service_id}/logging/errors` | Stream Log Endpoint Errors



## GetLogEndpointErrors

Stream Log Endpoint Errors



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
    serviceId := "SU1Z0isxPaozGVKXdv0eY" // string | 
    from := int64(1756123200) // int64 |  (optional)
    to := int64(1756209600) // int64 |  (optional)
    filterEndpoint := "MyS3,BigQuery" // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingEndpointErrorsAPI.GetLogEndpointErrors(ctx, serviceId).From(from).To(to).FilterEndpoint(filterEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingEndpointErrorsAPI.GetLogEndpointErrors`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogEndpointErrors`: string
    fmt.Fprintf(os.Stdout, "Response from `LoggingEndpointErrorsAPI.GetLogEndpointErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogEndpointErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  |  **to** | **int64** |  |  **filterEndpoint** | **string** |  | 

### Return type

**string**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

