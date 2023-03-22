# ObjectStoreItemAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeyFromStore**](ObjectStoreItemAPI.md#DeleteKeyFromStore) | **DELETE** `/resources/stores/object/{store_id}/keys/{key_name}` | Delete object store item.
[**GetKeys**](ObjectStoreItemAPI.md#GetKeys) | **GET** `/resources/stores/object/{store_id}/keys` | List object store keys.
[**GetValueForKey**](ObjectStoreItemAPI.md#GetValueForKey) | **GET** `/resources/stores/object/{store_id}/keys/{key_name}` | Get the value of an object store item
[**SetValueForKey**](ObjectStoreItemAPI.md#SetValueForKey) | **PUT** `/resources/stores/object/{store_id}/keys/{key_name}` | Insert an item into an object store



## DeleteKeyFromStore

Delete object store item.



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
    resp, r, err := apiClient.ObjectStoreItemAPI.DeleteKeyFromStore(ctx, storeID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreItemAPI.DeleteKeyFromStore`: %v\n", err)
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
    resp, r, err := apiClient.ObjectStoreItemAPI.GetKeys(ctx, storeID).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreItemAPI.GetKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeys`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreItemAPI.GetKeys`: %v\n", resp)
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

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetValueForKey

Get the value of an object store item



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
    resp, r, err := apiClient.ObjectStoreItemAPI.GetValueForKey(ctx, storeID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreItemAPI.GetValueForKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValueForKey`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreItemAPI.GetValueForKey`: %v\n", resp)
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

Insert an item into an object store



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
    resp, r, err := apiClient.ObjectStoreItemAPI.SetValueForKey(ctx, storeID, keyName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoreItemAPI.SetValueForKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetValueForKey`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoreItemAPI.SetValueForKey`: %v\n", resp)
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
