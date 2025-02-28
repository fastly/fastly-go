# ProductOriginInspectorAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductOriginInspector**](ProductOriginInspectorAPI.md#DisableProductOriginInspector) | **DELETE** `/enabled-products/v1/origin_inspector/services/{service_id}` | Disable product
[**EnableProductOriginInspector**](ProductOriginInspectorAPI.md#EnableProductOriginInspector) | **PUT** `/enabled-products/v1/origin_inspector/services/{service_id}` | Enable product
[**GetProductOriginInspector**](ProductOriginInspectorAPI.md#GetProductOriginInspector) | **GET** `/enabled-products/v1/origin_inspector/services/{service_id}` | Get product enablement status



## DisableProductOriginInspector

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
    resp, r, err := apiClient.ProductOriginInspectorAPI.DisableProductOriginInspector(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductOriginInspectorAPI.DisableProductOriginInspector`: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableProductOriginInspectorRequest struct via the builder pattern


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


## EnableProductOriginInspector

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
    resp, r, err := apiClient.ProductOriginInspectorAPI.EnableProductOriginInspector(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductOriginInspectorAPI.EnableProductOriginInspector`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableProductOriginInspector`: OriginInspectorResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductOriginInspectorAPI.EnableProductOriginInspector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProductOriginInspectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OriginInspectorResponseBodyEnable**](OriginInspectorResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductOriginInspector

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
    resp, r, err := apiClient.ProductOriginInspectorAPI.GetProductOriginInspector(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductOriginInspectorAPI.GetProductOriginInspector`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductOriginInspector`: OriginInspectorResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductOriginInspectorAPI.GetProductOriginInspector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductOriginInspectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OriginInspectorResponseBodyEnable**](OriginInspectorResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
