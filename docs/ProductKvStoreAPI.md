# ProductKvStoreAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductKvStore**](ProductKvStoreAPI.md#DisableProductKvStore) | **DELETE** `/enabled-products/v1/kv_store` | Disable product
[**EnableKvStore**](ProductKvStoreAPI.md#EnableKvStore) | **PUT** `/enabled-products/v1/kv_store` | Enable product
[**GetKvStore**](ProductKvStoreAPI.md#GetKvStore) | **GET** `/enabled-products/v1/kv_store` | Get product enablement status



## DisableProductKvStore

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
    resp, r, err := apiClient.ProductKvStoreAPI.DisableProductKvStore(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductKvStoreAPI.DisableProductKvStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductKvStoreRequest struct via the builder pattern



### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## EnableKvStore

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
    resp, r, err := apiClient.ProductKvStoreAPI.EnableKvStore(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductKvStoreAPI.EnableKvStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableKvStore`: KvStoreResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductKvStoreAPI.EnableKvStore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableKvStoreRequest struct via the builder pattern



### Return type

[**KvStoreResponseBodyEnable**](KvStoreResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetKvStore

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
    resp, r, err := apiClient.ProductKvStoreAPI.GetKvStore(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductKvStoreAPI.GetKvStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKvStore`: KvStoreResponseBodyEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductKvStoreAPI.GetKvStore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKvStoreRequest struct via the builder pattern



### Return type

[**KvStoreResponseBodyEnable**](KvStoreResponseBodyEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

