# KvStoreItemAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeyFromStore**](KvStoreItemAPI.md#DeleteKeyFromStore) | **DELETE** `/resources/stores/kv/{store_id}/keys/{key_name}` | Delete kv store item.
[**GetKeys**](KvStoreItemAPI.md#GetKeys) | **GET** `/resources/stores/kv/{store_id}/keys` | List kv store keys.
[**GetValueForKey**](KvStoreItemAPI.md#GetValueForKey) | **GET** `/resources/stores/kv/{store_id}/keys/{key_name}` | Get the value of an kv store item
[**SetValueForKey**](KvStoreItemAPI.md#SetValueForKey) | **PUT** `/resources/stores/kv/{store_id}/keys/{key_name}` | Insert an item into an kv store



## DeleteKeyFromStore

Delete kv store item.



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
    resp, r, err := apiClient.KvStoreItemAPI.DeleteKeyFromStore(ctx, storeID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.DeleteKeyFromStore`: %v\n", err)
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


## GetKeys

List kv store keys.



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
    prefix := "prefix_example" // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreItemAPI.GetKeys(ctx, storeID).Cursor(cursor).Limit(limit).Prefix(prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.GetKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeys`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `KvStoreItemAPI.GetKeys`: %v\n", resp)
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
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 100] **prefix** | **string** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetValueForKey

Get the value of an kv store item



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
    resp, r, err := apiClient.KvStoreItemAPI.GetValueForKey(ctx, storeID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.GetValueForKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValueForKey`: string
    fmt.Fprintf(os.Stdout, "Response from `KvStoreItemAPI.GetValueForKey`: %v\n", resp)
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

**string**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SetValueForKey

Insert an item into an kv store



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
    ifGenerationMatch := int32(56) // int32 |  (optional)
    timeToLiveSec := int32(56) // int32 |  (optional)
    metadata := "metadata_example" // string |  (optional)
    add := true // bool |  (optional)
    append := true // bool |  (optional)
    prepend := true // bool |  (optional)
    backgroundFetch := true // bool |  (optional)
    body := string(BYTE_ARRAY_DATA_HERE) // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.KvStoreItemAPI.SetValueForKey(ctx, storeID, keyName).IfGenerationMatch(ifGenerationMatch).TimeToLiveSec(timeToLiveSec).Metadata(metadata).Add(add).Append(append).Prepend(prepend).BackgroundFetch(backgroundFetch).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvStoreItemAPI.SetValueForKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetValueForKey`: string
    fmt.Fprintf(os.Stdout, "Response from `KvStoreItemAPI.SetValueForKey`: %v\n", resp)
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
 **ifGenerationMatch** | **int32** |  |  **timeToLiveSec** | **int32** |  |  **metadata** | **string** |  |  **add** | **bool** |  |  **append** | **bool** |  |  **prepend** | **bool** |  |  **backgroundFetch** | **bool** |  |  **body** | **string** |  | 

### Return type

**string**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/octet-stream

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
