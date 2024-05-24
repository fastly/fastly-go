# DictionaryAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDictionary**](DictionaryAPI.md#CreateDictionary) | **POST** `/service/{service_id}/version/{version_id}/dictionary` | Create an edge dictionary
[**DeleteDictionary**](DictionaryAPI.md#DeleteDictionary) | **DELETE** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_name}` | Delete an edge dictionary
[**GetDictionary**](DictionaryAPI.md#GetDictionary) | **GET** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_name}` | Get an edge dictionary
[**ListDictionaries**](DictionaryAPI.md#ListDictionaries) | **GET** `/service/{service_id}/version/{version_id}/dictionary` | List edge dictionaries
[**UpdateDictionary**](DictionaryAPI.md#UpdateDictionary) | **PUT** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_name}` | Update an edge dictionary



## CreateDictionary

Create an edge dictionary



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    name := "name_example" // string | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). (optional)
    writeOnly := true // bool | Determines if items in the dictionary are readable or not. (optional) (default to false)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryAPI.CreateDictionary(ctx, serviceID, versionID).Name(name).WriteOnly(writeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.CreateDictionary`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDictionary`: DictionaryResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.CreateDictionary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). |  **writeOnly** | **bool** | Determines if items in the dictionary are readable or not. | [default to false]

### Return type

[**DictionaryResponse**](DictionaryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDictionary

Delete an edge dictionary



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    dictionaryName := "dictionaryName_example" // string | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryAPI.DeleteDictionary(ctx, serviceID, versionID, dictionaryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.DeleteDictionary`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDictionary`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.DeleteDictionary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**dictionaryName** | **string** | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDictionaryRequest struct via the builder pattern


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


## GetDictionary

Get an edge dictionary



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    dictionaryName := "dictionaryName_example" // string | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryAPI.GetDictionary(ctx, serviceID, versionID, dictionaryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.GetDictionary`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDictionary`: DictionaryResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.GetDictionary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**dictionaryName** | **string** | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DictionaryResponse**](DictionaryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDictionaries

List edge dictionaries



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
    versionID := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryAPI.ListDictionaries(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.ListDictionaries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDictionaries`: []DictionaryResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.ListDictionaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDictionariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DictionaryResponse**](DictionaryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDictionary

Update an edge dictionary



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    dictionaryName := "dictionaryName_example" // string | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).
    name := "name_example" // string | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). (optional)
    writeOnly := true // bool | Determines if items in the dictionary are readable or not. (optional) (default to false)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryAPI.UpdateDictionary(ctx, serviceID, versionID, dictionaryName).Name(name).WriteOnly(writeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.UpdateDictionary`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDictionary`: DictionaryResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.UpdateDictionary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**dictionaryName** | **string** | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). |  **writeOnly** | **bool** | Determines if items in the dictionary are readable or not. | [default to false]

### Return type

[**DictionaryResponse**](DictionaryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
