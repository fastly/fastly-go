# KvStoreAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvStoreCreate**](KvStoreAPI.md#KvStoreCreate) | **POST** `/resources/stores/kv` | Create a KV store.
[**KvStoreDelete**](KvStoreAPI.md#KvStoreDelete) | **DELETE** `/resources/stores/kv/{store_id}` | Delete a KV store.
[**KvStoreGet**](KvStoreAPI.md#KvStoreGet) | **GET** `/resources/stores/kv/{store_id}` | Describe a KV store.
[**KvStoreList**](KvStoreAPI.md#KvStoreList) | **GET** `/resources/stores/kv` | List all KV stores.
[**KvStorePut**](KvStoreAPI.md#KvStorePut) | **PUT** `/resources/stores/kv/{store_id}` | Update a KV store.



## KvStoreCreate

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
    kvStoreRequestCreateOrUpdate := *openapiclient.NewKvStoreRequestCreateOrUpdate("Name_example") // KvStoreRequestCreateOrUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.KvStoreCreate(ctx).Location(location).KvStoreRequestCreateOrUpdate(kvStoreRequestCreateOrUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.KvStoreCreate`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KvStoreCreate`: KvStoreDetails
    fmt.Fprintf(os.Stdout, "Response from `KvStoreAPI.KvStoreCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** |  |  **kvStoreRequestCreateOrUpdate** | [**KvStoreRequestCreateOrUpdate**](KvStoreRequestCreateOrUpdate.md) |  | 

### Return type

[**KvStoreDetails**](KvStoreDetails.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## KvStoreDelete

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
    storeId := "storeId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.KvStoreDelete(ctx, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.KvStoreDelete`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreDeleteRequest struct via the builder pattern


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


## KvStoreGet

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
    storeId := "storeId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.KvStoreGet(ctx, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.KvStoreGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KvStoreGet`: KvStoreDetails
    fmt.Fprintf(os.Stdout, "Response from `KvStoreAPI.KvStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KvStoreDetails**](KvStoreDetails.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## KvStoreList

List all KV stores.



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
    name := "name_example" // string | Returns a one-element array containing the details for the named KV store. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.KvStoreList(ctx).Cursor(cursor).Limit(limit).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.KvStoreList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KvStoreList`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `KvStoreAPI.KvStoreList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 1000] **name** | **string** | Returns a one-element array containing the details for the named KV store. | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## KvStorePut

Update a KV store.



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
    storeId := "storeId_example" // string | 
    kvStoreRequestCreateOrUpdate := *openapiclient.NewKvStoreRequestCreateOrUpdate("Name_example") // KvStoreRequestCreateOrUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreAPI.KvStorePut(ctx, storeId).KvStoreRequestCreateOrUpdate(kvStoreRequestCreateOrUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreAPI.KvStorePut`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStorePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvStoreRequestCreateOrUpdate** | [**KvStoreRequestCreateOrUpdate**](KvStoreRequestCreateOrUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

