# ProductBrotliCompressionAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductBrotliCompression**](ProductBrotliCompressionAPI.md#DisableProductBrotliCompression) | **DELETE** `/enabled-products/v1/brotli_compression/services/{service_id}` | Disable product
[**EnableProductBrotliCompression**](ProductBrotliCompressionAPI.md#EnableProductBrotliCompression) | **PUT** `/enabled-products/v1/brotli_compression/services/{service_id}` | Enable product
[**GetProductBrotliCompression**](ProductBrotliCompressionAPI.md#GetProductBrotliCompression) | **GET** `/enabled-products/v1/brotli_compression/services/{service_id}` | Get product enablement status



## DisableProductBrotliCompression

Disable product



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductBrotliCompressionAPI.DisableProductBrotliCompression(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductBrotliCompressionAPI.DisableProductBrotliCompression`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductBrotliCompressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## EnableProductBrotliCompression

Enable product



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductBrotliCompressionAPI.EnableProductBrotliCompression(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductBrotliCompressionAPI.EnableProductBrotliCompression`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableProductBrotliCompression`: BrotliCompressionResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductBrotliCompressionAPI.EnableProductBrotliCompression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProductBrotliCompressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrotliCompressionResponseBodyEnable**](BrotliCompressionResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductBrotliCompression

Get product enablement status



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductBrotliCompressionAPI.GetProductBrotliCompression(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductBrotliCompressionAPI.GetProductBrotliCompression`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductBrotliCompression`: BrotliCompressionResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductBrotliCompressionAPI.GetProductBrotliCompression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductBrotliCompressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrotliCompressionResponseBodyEnable**](BrotliCompressionResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
