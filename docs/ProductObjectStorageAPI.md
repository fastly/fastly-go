# ProductObjectStorageAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductObjectStorage**](ProductObjectStorageAPI.md#DisableProductObjectStorage) | **DELETE** `/enabled-products/v1/object_storage` | Disable product
[**EnableObjectStorage**](ProductObjectStorageAPI.md#EnableObjectStorage) | **PUT** `/enabled-products/v1/object_storage` | Enable product
[**GetObjectStorage**](ProductObjectStorageAPI.md#GetObjectStorage) | **GET** `/enabled-products/v1/object_storage` | Get product enablement status



## DisableProductObjectStorage

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
    resp, r, err := apiClient.ProductObjectStorageAPI.DisableProductObjectStorage(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductObjectStorageAPI.DisableProductObjectStorage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductObjectStorageRequest struct via the builder pattern



### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## EnableObjectStorage

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
    resp, r, err := apiClient.ProductObjectStorageAPI.EnableObjectStorage(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductObjectStorageAPI.EnableObjectStorage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableObjectStorage`: ObjectStorageResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductObjectStorageAPI.EnableObjectStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableObjectStorageRequest struct via the builder pattern



### Return type

[**ObjectStorageResponseBodyEnable**](ObjectStorageResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetObjectStorage

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
    resp, r, err := apiClient.ProductObjectStorageAPI.GetObjectStorage(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductObjectStorageAPI.GetObjectStorage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorage`: ObjectStorageResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductObjectStorageAPI.GetObjectStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageRequest struct via the builder pattern



### Return type

[**ObjectStorageResponseBodyEnable**](ObjectStorageResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
