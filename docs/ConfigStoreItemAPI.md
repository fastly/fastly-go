# ConfigStoreItemAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateConfigStoreItem**](ConfigStoreItemAPI.md#BulkUpdateConfigStoreItem) | **PATCH** `/resources/stores/config/{config_store_id}/items` | Update multiple entries in a config store
[**CreateConfigStoreItem**](ConfigStoreItemAPI.md#CreateConfigStoreItem) | **POST** `/resources/stores/config/{config_store_id}/item` | Create an entry in a config store
[**DeleteConfigStoreItem**](ConfigStoreItemAPI.md#DeleteConfigStoreItem) | **DELETE** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Delete an item from a config store
[**GetConfigStoreItem**](ConfigStoreItemAPI.md#GetConfigStoreItem) | **GET** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Get an item from a config store
[**ListConfigStoreItems**](ConfigStoreItemAPI.md#ListConfigStoreItems) | **GET** `/resources/stores/config/{config_store_id}/items` | List items in a config store
[**UpdateConfigStoreItem**](ConfigStoreItemAPI.md#UpdateConfigStoreItem) | **PATCH** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Update an entry in a config store
[**UpsertConfigStoreItem**](ConfigStoreItemAPI.md#UpsertConfigStoreItem) | **PUT** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Insert or update an entry in a config store



## BulkUpdateConfigStoreItem

Update multiple entries in a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    bulkUpdateConfigStoreListRequest := *openapiclient.NewBulkUpdateConfigStoreListRequest() // BulkUpdateConfigStoreListRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.BulkUpdateConfigStoreItem(ctx, configStoreId).BulkUpdateConfigStoreListRequest(bulkUpdateConfigStoreListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.BulkUpdateConfigStoreItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateConfigStoreItem`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.BulkUpdateConfigStoreItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateConfigStoreItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateConfigStoreListRequest** | [**BulkUpdateConfigStoreListRequest**](BulkUpdateConfigStoreListRequest.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateConfigStoreItem

Create an entry in a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    itemKey := "itemKey_example" // string | Item key, maximum 256 characters. (optional)
    itemValue := "itemValue_example" // string | Item value, maximum 8000 characters. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.CreateConfigStoreItem(ctx, configStoreId).ItemKey(itemKey).ItemValue(itemValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.CreateConfigStoreItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigStoreItem`: ConfigStoreItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.CreateConfigStoreItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigStoreItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemKey** | **string** | Item key, maximum 256 characters. |  **itemValue** | **string** | Item value, maximum 8000 characters. | 

### Return type

[**ConfigStoreItemResponse**](ConfigStoreItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteConfigStoreItem

Delete an item from a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    configStoreItemKey := "configStoreItemKey_example" // string | Item key, maximum 256 characters.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.DeleteConfigStoreItem(ctx, configStoreId, configStoreItemKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.DeleteConfigStoreItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfigStoreItem`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.DeleteConfigStoreItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 
**configStoreItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigStoreItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetConfigStoreItem

Get an item from a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    configStoreItemKey := "configStoreItemKey_example" // string | Item key, maximum 256 characters.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.GetConfigStoreItem(ctx, configStoreId, configStoreItemKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.GetConfigStoreItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigStoreItem`: ConfigStoreItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.GetConfigStoreItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 
**configStoreItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigStoreItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigStoreItemResponse**](ConfigStoreItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListConfigStoreItems

List items in a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.ListConfigStoreItems(ctx, configStoreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.ListConfigStoreItems`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigStoreItems`: []ConfigStoreItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.ListConfigStoreItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigStoreItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigStoreItemResponse**](ConfigStoreItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateConfigStoreItem

Update an entry in a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    configStoreItemKey := "configStoreItemKey_example" // string | Item key, maximum 256 characters.
    itemKey := "itemKey_example" // string | Item key, maximum 256 characters. (optional)
    itemValue := "itemValue_example" // string | Item value, maximum 8000 characters. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.UpdateConfigStoreItem(ctx, configStoreId, configStoreItemKey).ItemKey(itemKey).ItemValue(itemValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.UpdateConfigStoreItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigStoreItem`: ConfigStoreItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.UpdateConfigStoreItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 
**configStoreItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigStoreItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemKey** | **string** | Item key, maximum 256 characters. |  **itemValue** | **string** | Item value, maximum 8000 characters. | 

### Return type

[**ConfigStoreItemResponse**](ConfigStoreItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpsertConfigStoreItem

Insert or update an entry in a config store



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
    configStoreId := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    configStoreItemKey := "configStoreItemKey_example" // string | Item key, maximum 256 characters.
    itemKey := "itemKey_example" // string | Item key, maximum 256 characters. (optional)
    itemValue := "itemValue_example" // string | Item value, maximum 8000 characters. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreItemAPI.UpsertConfigStoreItem(ctx, configStoreId, configStoreItemKey).ItemKey(itemKey).ItemValue(itemValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreItemAPI.UpsertConfigStoreItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertConfigStoreItem`: ConfigStoreItemResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreItemAPI.UpsertConfigStoreItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreId** | **string** | An alphanumeric string identifying the config store. | 
**configStoreItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertConfigStoreItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemKey** | **string** | Item key, maximum 256 characters. |  **itemValue** | **string** | Item value, maximum 8000 characters. | 

### Return type

[**ConfigStoreItemResponse**](ConfigStoreItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

