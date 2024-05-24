# HeaderAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHeaderObject**](HeaderAPI.md#CreateHeaderObject) | **POST** `/service/{service_id}/version/{version_id}/header` | Create a Header object
[**DeleteHeaderObject**](HeaderAPI.md#DeleteHeaderObject) | **DELETE** `/service/{service_id}/version/{version_id}/header/{header_name}` | Delete a Header object
[**GetHeaderObject**](HeaderAPI.md#GetHeaderObject) | **GET** `/service/{service_id}/version/{version_id}/header/{header_name}` | Get a Header object
[**ListHeaderObjects**](HeaderAPI.md#ListHeaderObjects) | **GET** `/service/{service_id}/version/{version_id}/header` | List Header objects
[**UpdateHeaderObject**](HeaderAPI.md#UpdateHeaderObject) | **PUT** `/service/{service_id}/version/{version_id}/header/{header_name}` | Update a Header object



## CreateHeaderObject

Create a Header object



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
    action := "action_example" // string | Accepts a string value. (optional)
    cacheCondition := "cacheCondition_example" // string | Name of the cache condition controlling when this configuration applies. (optional)
    dst := "dst_example" // string | Header to set. (optional)
    name := "name_example" // string | A handle to refer to this Header object. (optional)
    regex := "regex_example" // string | Regular expression to use. Only applies to `regex` and `regex_repeat` actions. (optional)
    requestCondition := "requestCondition_example" // string | Condition which, if met, will select this configuration during a request. Optional. (optional)
    responseCondition := "responseCondition_example" // string | Optional name of a response condition to apply. (optional)
    src := "src_example" // string | Variable to be used as a source for the header content. Does not apply to `delete` action. (optional)
    substitution := "substitution_example" // string | Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions. (optional)
    resourceType := "resourceType_example" // string | Accepts a string value. (optional)
    ignoreIfSet := int32(56) // int32 | Don't add the header if it is added already. Only applies to 'set' action. (optional)
    priority := int32(56) // int32 | Priority determines execution order. Lower numbers execute first. (optional) (default to 100)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HeaderAPI.CreateHeaderObject(ctx, serviceID, versionID).Action(action).CacheCondition(cacheCondition).Dst(dst).Name(name).Regex(regex).RequestCondition(requestCondition).ResponseCondition(responseCondition).Src(src).Substitution(substitution).ResourceType(resourceType).IgnoreIfSet(ignoreIfSet).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeaderAPI.CreateHeaderObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHeaderObject`: HeaderResponse
    fmt.Fprintf(os.Stdout, "Response from `HeaderAPI.CreateHeaderObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHeaderObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Accepts a string value. |  **cacheCondition** | **string** | Name of the cache condition controlling when this configuration applies. |  **dst** | **string** | Header to set. |  **name** | **string** | A handle to refer to this Header object. |  **regex** | **string** | Regular expression to use. Only applies to `regex` and `regex_repeat` actions. |  **requestCondition** | **string** | Condition which, if met, will select this configuration during a request. Optional. |  **responseCondition** | **string** | Optional name of a response condition to apply. |  **src** | **string** | Variable to be used as a source for the header content. Does not apply to `delete` action. |  **substitution** | **string** | Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions. |  **resourceType** | **string** | Accepts a string value. |  **ignoreIfSet** | **int32** | Don&#39;t add the header if it is added already. Only applies to &#39;set&#39; action. |  **priority** | **int32** | Priority determines execution order. Lower numbers execute first. | [default to 100]

### Return type

[**HeaderResponse**](HeaderResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteHeaderObject

Delete a Header object



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
    headerName := "headerName_example" // string | A handle to refer to this Header object.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HeaderAPI.DeleteHeaderObject(ctx, serviceID, versionID, headerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeaderAPI.DeleteHeaderObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHeaderObject`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `HeaderAPI.DeleteHeaderObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**headerName** | **string** | A handle to refer to this Header object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHeaderObjectRequest struct via the builder pattern


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


## GetHeaderObject

Get a Header object



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
    headerName := "headerName_example" // string | A handle to refer to this Header object.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HeaderAPI.GetHeaderObject(ctx, serviceID, versionID, headerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeaderAPI.GetHeaderObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeaderObject`: HeaderResponse
    fmt.Fprintf(os.Stdout, "Response from `HeaderAPI.GetHeaderObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**headerName** | **string** | A handle to refer to this Header object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeaderObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HeaderResponse**](HeaderResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListHeaderObjects

List Header objects



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
    resp, r, err := apiClient.HeaderAPI.ListHeaderObjects(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeaderAPI.ListHeaderObjects`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHeaderObjects`: []HeaderResponse
    fmt.Fprintf(os.Stdout, "Response from `HeaderAPI.ListHeaderObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHeaderObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HeaderResponse**](HeaderResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateHeaderObject

Update a Header object



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
    headerName := "headerName_example" // string | A handle to refer to this Header object.
    action := "action_example" // string | Accepts a string value. (optional)
    cacheCondition := "cacheCondition_example" // string | Name of the cache condition controlling when this configuration applies. (optional)
    dst := "dst_example" // string | Header to set. (optional)
    name := "name_example" // string | A handle to refer to this Header object. (optional)
    regex := "regex_example" // string | Regular expression to use. Only applies to `regex` and `regex_repeat` actions. (optional)
    requestCondition := "requestCondition_example" // string | Condition which, if met, will select this configuration during a request. Optional. (optional)
    responseCondition := "responseCondition_example" // string | Optional name of a response condition to apply. (optional)
    src := "src_example" // string | Variable to be used as a source for the header content. Does not apply to `delete` action. (optional)
    substitution := "substitution_example" // string | Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions. (optional)
    resourceType := "resourceType_example" // string | Accepts a string value. (optional)
    ignoreIfSet := int32(56) // int32 | Don't add the header if it is added already. Only applies to 'set' action. (optional)
    priority := int32(56) // int32 | Priority determines execution order. Lower numbers execute first. (optional) (default to 100)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HeaderAPI.UpdateHeaderObject(ctx, serviceID, versionID, headerName).Action(action).CacheCondition(cacheCondition).Dst(dst).Name(name).Regex(regex).RequestCondition(requestCondition).ResponseCondition(responseCondition).Src(src).Substitution(substitution).ResourceType(resourceType).IgnoreIfSet(ignoreIfSet).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeaderAPI.UpdateHeaderObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHeaderObject`: HeaderResponse
    fmt.Fprintf(os.Stdout, "Response from `HeaderAPI.UpdateHeaderObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**headerName** | **string** | A handle to refer to this Header object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHeaderObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Accepts a string value. |  **cacheCondition** | **string** | Name of the cache condition controlling when this configuration applies. |  **dst** | **string** | Header to set. |  **name** | **string** | A handle to refer to this Header object. |  **regex** | **string** | Regular expression to use. Only applies to `regex` and `regex_repeat` actions. |  **requestCondition** | **string** | Condition which, if met, will select this configuration during a request. Optional. |  **responseCondition** | **string** | Optional name of a response condition to apply. |  **src** | **string** | Variable to be used as a source for the header content. Does not apply to `delete` action. |  **substitution** | **string** | Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions. |  **resourceType** | **string** | Accepts a string value. |  **ignoreIfSet** | **int32** | Don&#39;t add the header if it is added already. Only applies to &#39;set&#39; action. |  **priority** | **int32** | Priority determines execution order. Lower numbers execute first. | [default to 100]

### Return type

[**HeaderResponse**](HeaderResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
