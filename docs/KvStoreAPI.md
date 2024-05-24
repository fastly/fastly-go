# KvStoreAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](KvStoreAPI.md#CreateStore) | **POST** `/resources/stores/kv` | Create a KV store.
[**DeleteStore**](KvStoreAPI.md#DeleteStore) | **DELETE** `/resources/stores/kv/{store_id}` | Delete a KV store.
[**GetStore**](KvStoreAPI.md#GetStore) | **GET** `/resources/stores/kv/{store_id}` | Describe a KV store.
[**GetStores**](KvStoreAPI.md#GetStores) | **GET** `/resources/stores/kv` | List KV stores.



## CreateStore

Create a KV store.



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
    location := "location_example" // string |  (optional)
    store := *openapiclient.NewStore() // Store |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.CreateStore(ctx).Location(location).Store(store).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.CreateStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStore`: StoreResponse
    fmt.Fprintf(os.Stdout, "Response from `KvStoreAPI.CreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** |  |  **store** | [**Store**](Store.md) |  | 

### Return type

[**StoreResponse**](StoreResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteStore

Delete a KV store.



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
    storeID := "storeId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.DeleteStore(ctx, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.DeleteStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreRequest struct via the builder pattern


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


## GetStore

Describe a KV store.



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
    storeID := "storeId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.GetStore(ctx, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.GetStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStore`: StoreResponse
    fmt.Fprintf(os.Stdout, "Response from `KvStoreAPI.GetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StoreResponse**](StoreResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetStores

List KV stores.



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
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 1000)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.GetStores(ctx).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.GetStores`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStores`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `KvStoreAPI.GetStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 1000]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
