# KvStoreItemAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvStoreDeleteItem**](KvStoreItemAPI.md#KvStoreDeleteItem) | **DELETE** `/resources/stores/kv/{store_id}/keys/{key}` | Delete an item.
[**KvStoreGetItem**](KvStoreItemAPI.md#KvStoreGetItem) | **GET** `/resources/stores/kv/{store_id}/keys/{key}` | Get an item.
[**KvStoreListItemKeys**](KvStoreItemAPI.md#KvStoreListItemKeys) | **GET** `/resources/stores/kv/{store_id}/keys` | List item keys.
[**KvStoreUpsertItem**](KvStoreItemAPI.md#KvStoreUpsertItem) | **PUT** `/resources/stores/kv/{store_id}/keys/{key}` | Insert or update an item.



## KvStoreDeleteItem

Delete an item.



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
    key := "key_example" // string | 
    ifGenerationMatch := int32(56) // int32 |  (optional)
    force := true // bool |  (optional) (default to false)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreItemAPI.KvStoreDeleteItem(ctx, storeId, key).IfGenerationMatch(ifGenerationMatch).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.KvStoreDeleteItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreDeleteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifGenerationMatch** | **int32** |  |  **force** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## KvStoreGetItem

Get an item.



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
    key := "key_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreItemAPI.KvStoreGetItem(ctx, storeId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.KvStoreGetItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KvStoreGetItem`: string
    fmt.Fprintf(os.Stdout, "Response from `KvStoreItemAPI.KvStoreGetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## KvStoreListItemKeys

List item keys.



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
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 100)
    prefix := "prefix_example" // string |  (optional)
    consistency := "consistency_example" // string |  (optional) (default to "strong")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreItemAPI.KvStoreListItemKeys(ctx, storeId).Cursor(cursor).Limit(limit).Prefix(prefix).Consistency(consistency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.KvStoreListItemKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KvStoreListItemKeys`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `KvStoreItemAPI.KvStoreListItemKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreListItemKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 100] **prefix** | **string** |  |  **consistency** | **string** |  | [default to &quot;strong&quot;]

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## KvStoreUpsertItem

Insert or update an item.



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
    key := "key_example" // string | 
    ifGenerationMatch := int32(56) // int32 |  (optional)
    timeToLiveSec := int32(56) // int32 |  (optional)
    metadata := "metadata_example" // string |  (optional)
    add := true // bool |  (optional) (default to false)
    append := true // bool |  (optional) (default to false)
    prepend := true // bool |  (optional) (default to false)
    backgroundFetch := true // bool |  (optional) (default to false)
    body := string(BYTE_ARRAY_DATA_HERE) // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreItemAPI.KvStoreUpsertItem(ctx, storeId, key).IfGenerationMatch(ifGenerationMatch).TimeToLiveSec(timeToLiveSec).Metadata(metadata).Add(add).Append(append).Prepend(prepend).BackgroundFetch(backgroundFetch).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.KvStoreUpsertItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStoreUpsertItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifGenerationMatch** | **int32** |  |  **timeToLiveSec** | **int32** |  |  **metadata** | **string** |  |  **add** | **bool** |  | [default to false] **append** | **bool** |  | [default to false] **prepend** | **bool** |  | [default to false] **backgroundFetch** | **bool** |  | [default to false] **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

