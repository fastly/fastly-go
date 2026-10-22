# ProductAiRuntimeControlAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductAiRuntimeControl**](ProductAiRuntimeControlAPI.md#DisableProductAiRuntimeControl) | **DELETE** `/enabled-products/v1/ai_runtime_control` | Disable product
[**EnableProductAiRuntimeControl**](ProductAiRuntimeControlAPI.md#EnableProductAiRuntimeControl) | **PUT** `/enabled-products/v1/ai_runtime_control` | Enable product
[**GetProductAiRuntimeControl**](ProductAiRuntimeControlAPI.md#GetProductAiRuntimeControl) | **GET** `/enabled-products/v1/ai_runtime_control` | Get product enablement status



## DisableProductAiRuntimeControl

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
    resp, r, err := apiClient.ProductAiRuntimeControlAPI.DisableProductAiRuntimeControl(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAiRuntimeControlAPI.DisableProductAiRuntimeControl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductAiRuntimeControlRequest struct via the builder pattern



### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## EnableProductAiRuntimeControl

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
    resp, r, err := apiClient.ProductAiRuntimeControlAPI.EnableProductAiRuntimeControl(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAiRuntimeControlAPI.EnableProductAiRuntimeControl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableProductAiRuntimeControl`: AiRuntimeControlResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductAiRuntimeControlAPI.EnableProductAiRuntimeControl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProductAiRuntimeControlRequest struct via the builder pattern



### Return type

[**AiRuntimeControlResponseBodyEnable**](AiRuntimeControlResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductAiRuntimeControl

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
    resp, r, err := apiClient.ProductAiRuntimeControlAPI.GetProductAiRuntimeControl(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAiRuntimeControlAPI.GetProductAiRuntimeControl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductAiRuntimeControl`: AiRuntimeControlResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductAiRuntimeControlAPI.GetProductAiRuntimeControl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAiRuntimeControlRequest struct via the builder pattern



### Return type

[**AiRuntimeControlResponseBodyEnable**](AiRuntimeControlResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

