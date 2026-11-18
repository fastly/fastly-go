# ProductAiAcceleratorAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductAiAccelerator**](ProductAiAcceleratorAPI.md#DisableProductAiAccelerator) | **DELETE** `/enabled-products/v1/ai_accelerator` | Disable product
[**EnableAiAccelerator**](ProductAiAcceleratorAPI.md#EnableAiAccelerator) | **PUT** `/enabled-products/v1/ai_accelerator` | Enable product
[**GetAiAccelerator**](ProductAiAcceleratorAPI.md#GetAiAccelerator) | **GET** `/enabled-products/v1/ai_accelerator` | Get product enablement status



## DisableProductAiAccelerator

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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductAiAcceleratorAPI.DisableProductAiAccelerator(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAiAcceleratorAPI.DisableProductAiAccelerator`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductAiAcceleratorRequest struct via the builder pattern



### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## EnableAiAccelerator

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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductAiAcceleratorAPI.EnableAiAccelerator(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAiAcceleratorAPI.EnableAiAccelerator`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAiAccelerator`: AiAcceleratorResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductAiAcceleratorAPI.EnableAiAccelerator`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAiAcceleratorRequest struct via the builder pattern



### Return type

[**AiAcceleratorResponseBodyEnable**](AiAcceleratorResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetAiAccelerator

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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductAiAcceleratorAPI.GetAiAccelerator(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAiAcceleratorAPI.GetAiAccelerator`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiAccelerator`: AiAcceleratorResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductAiAcceleratorAPI.GetAiAccelerator`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAcceleratorRequest struct via the builder pattern



### Return type

[**AiAcceleratorResponseBodyEnable**](AiAcceleratorResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
