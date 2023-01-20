# ObjectStoreAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](ObjectStoreAPI.md#CreateStore) | **POST** `/resources/stores/object` | Create an object store.
[**DeleteKeyFromStore**](ObjectStoreAPI.md#DeleteKeyFromStore) | **DELETE** `/resources/stores/object/{store_id}/keys/{key_name}` | Delete object store key.
[**DeleteStore**](ObjectStoreAPI.md#DeleteStore) | **DELETE** `/resources/stores/object/{store_id}` | Delete an object store.
[**GetKeys**](ObjectStoreAPI.md#GetKeys) | **GET** `/resources/stores/object/{store_id}/keys` | List object store keys.
[**GetStore**](ObjectStoreAPI.md#GetStore) | **GET** `/resources/stores/object/{store_id}` | Describe an object store.
[**GetStores**](ObjectStoreAPI.md#GetStores) | **GET** `/resources/stores/object` | List object stores.
[**GetValueForKey**](ObjectStoreAPI.md#GetValueForKey) | **GET** `/resources/stores/object/{store_id}/keys/{key_name}` | Get object store key value.
[**SetValueForKey**](ObjectStoreAPI.md#SetValueForKey) | **PUT** `/resources/stores/object/{store_id}/keys/{key_name}` | Insert object store key-value.



## CreateStore

Create an object store.



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
    store := *openapiclient.NewStore() // Store |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStoreAPI.CreateStore(ctx).Store(store).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.CreateStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStore`: StoreResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreAPI.CreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **store** | [**Store**](Store.md) |  | 

### Return type

[**StoreResponse**](StoreResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteKeyFromStore

Delete object store key.



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
    keyName := "keyName_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStoreAPI.DeleteKeyFromStore(ctx, storeID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.DeleteKeyFromStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyFromStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteStore

Delete an object store.



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
    resp, r, err := apiClient.ObjectStoreAPI.DeleteStore(ctx, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.DeleteStore`: %v\n", err)
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

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetKeys

List object store keys.



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
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 100)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStoreAPI.GetKeys(ctx, storeID).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.GetKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeys`: KeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreAPI.GetKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 100]

### Return type

[**KeyResponse**](KeyResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetStore

Describe an object store.



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
    resp, r, err := apiClient.ObjectStoreAPI.GetStore(ctx, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.GetStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStore`: StoreResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreAPI.GetStore`: %v\n", resp)
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

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetStores

List object stores.



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
    limit := int32(56) // int32 |  (optional) (default to 100)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStoreAPI.GetStores(ctx).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.GetStores`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStores`: GetStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreAPI.GetStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 100]

### Return type

[**GetStoresResponse**](GetStoresResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetValueForKey

Get object store key value.



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
    keyName := "keyName_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStoreAPI.GetValueForKey(ctx, storeID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.GetValueForKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValueForKey`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreAPI.GetValueForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValueForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SetValueForKey

Insert object store key-value.



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
    keyName := "keyName_example" // string | 
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStoreAPI.SetValueForKey(ctx, storeID, keyName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreAPI.SetValueForKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetValueForKey`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreAPI.SetValueForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetValueForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/octet-stream

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
