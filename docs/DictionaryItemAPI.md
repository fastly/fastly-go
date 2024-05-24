# DictionaryItemAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateDictionaryItem**](DictionaryItemAPI.md#BulkUpdateDictionaryItem) | **PATCH** `/service/{service_id}/dictionary/{dictionary_id}/items` | Update multiple entries in an edge dictionary
[**CreateDictionaryItem**](DictionaryItemAPI.md#CreateDictionaryItem) | **POST** `/service/{service_id}/dictionary/{dictionary_id}/item` | Create an entry in an edge dictionary
[**DeleteDictionaryItem**](DictionaryItemAPI.md#DeleteDictionaryItem) | **DELETE** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Delete an item from an edge dictionary
[**GetDictionaryItem**](DictionaryItemAPI.md#GetDictionaryItem) | **GET** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Get an item from an edge dictionary
[**ListDictionaryItems**](DictionaryItemAPI.md#ListDictionaryItems) | **GET** `/service/{service_id}/dictionary/{dictionary_id}/items` | List items in an edge dictionary
[**UpdateDictionaryItem**](DictionaryItemAPI.md#UpdateDictionaryItem) | **PATCH** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Update an entry in an edge dictionary
[**UpsertDictionaryItem**](DictionaryItemAPI.md#UpsertDictionaryItem) | **PUT** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Insert or update an entry in an edge dictionary



## BulkUpdateDictionaryItem

Update multiple entries in an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    bulkUpdateDictionaryListRequest := *openapiclient.NewBulkUpdateDictionaryListRequest() // BulkUpdateDictionaryListRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.BulkUpdateDictionaryItem(ctx, serviceID, dictionaryID).BulkUpdateDictionaryListRequest(bulkUpdateDictionaryListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.BulkUpdateDictionaryItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateDictionaryItem`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.BulkUpdateDictionaryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateDictionaryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateDictionaryListRequest** | [**BulkUpdateDictionaryListRequest**](BulkUpdateDictionaryListRequest.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateDictionaryItem

Create an entry in an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    itemKey := "itemKey_example" // string | Item key, maximum 256 characters. (optional)
    itemValue := "itemValue_example" // string | Item value, maximum 8000 characters. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.CreateDictionaryItem(ctx, serviceID, dictionaryID).ItemKey(itemKey).ItemValue(itemValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.CreateDictionaryItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDictionaryItem`: DictionaryItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.CreateDictionaryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDictionaryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemKey** | **string** | Item key, maximum 256 characters. |  **itemValue** | **string** | Item value, maximum 8000 characters. | 

### Return type

[**DictionaryItemResponse**](DictionaryItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDictionaryItem

Delete an item from an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    dictionaryItemKey := "dictionaryItemKey_example" // string | Item key, maximum 256 characters.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.DeleteDictionaryItem(ctx, serviceID, dictionaryID, dictionaryItemKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.DeleteDictionaryItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDictionaryItem`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.DeleteDictionaryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 
**dictionaryItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDictionaryItemRequest struct via the builder pattern


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


## GetDictionaryItem

Get an item from an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    dictionaryItemKey := "dictionaryItemKey_example" // string | Item key, maximum 256 characters.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.GetDictionaryItem(ctx, serviceID, dictionaryID, dictionaryItemKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.GetDictionaryItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDictionaryItem`: DictionaryItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.GetDictionaryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 
**dictionaryItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDictionaryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DictionaryItemResponse**](DictionaryItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDictionaryItems

List items in an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    page := int32(1) // int32 | Current page. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created" // string | Field on which to sort. (optional) (default to "created")
    direction := "ascend" // string | Direction in which to sort results. (optional) (default to "ascend")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.ListDictionaryItems(ctx, serviceID, dictionaryID).Page(page).PerPage(perPage).Sort(sort).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.ListDictionaryItems`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDictionaryItems`: []DictionaryItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.ListDictionaryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDictionaryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page. |  **perPage** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | Field on which to sort. | [default to &quot;created&quot;] **direction** | **string** | Direction in which to sort results. | [default to &quot;ascend&quot;]

### Return type

[**[]DictionaryItemResponse**](DictionaryItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDictionaryItem

Update an entry in an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    dictionaryItemKey := "dictionaryItemKey_example" // string | Item key, maximum 256 characters.
    itemKey := "itemKey_example" // string | Item key, maximum 256 characters. (optional)
    itemValue := "itemValue_example" // string | Item value, maximum 8000 characters. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.UpdateDictionaryItem(ctx, serviceID, dictionaryID, dictionaryItemKey).ItemKey(itemKey).ItemValue(itemValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.UpdateDictionaryItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDictionaryItem`: DictionaryItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.UpdateDictionaryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 
**dictionaryItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDictionaryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemKey** | **string** | Item key, maximum 256 characters. |  **itemValue** | **string** | Item value, maximum 8000 characters. | 

### Return type

[**DictionaryItemResponse**](DictionaryItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpsertDictionaryItem

Insert or update an entry in an edge dictionary



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.
    dictionaryItemKey := "dictionaryItemKey_example" // string | Item key, maximum 256 characters.
    itemKey := "itemKey_example" // string | Item key, maximum 256 characters. (optional)
    itemValue := "itemValue_example" // string | Item value, maximum 8000 characters. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryItemAPI.UpsertDictionaryItem(ctx, serviceID, dictionaryID, dictionaryItemKey).ItemKey(itemKey).ItemValue(itemValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryItemAPI.UpsertDictionaryItem`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertDictionaryItem`: DictionaryItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryItemAPI.UpsertDictionaryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 
**dictionaryItemKey** | **string** | Item key, maximum 256 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertDictionaryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemKey** | **string** | Item key, maximum 256 characters. |  **itemValue** | **string** | Item value, maximum 8000 characters. | 

### Return type

[**DictionaryItemResponse**](DictionaryItemResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
