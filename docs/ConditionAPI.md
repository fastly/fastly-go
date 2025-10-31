# ConditionAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCondition**](ConditionAPI.md#CreateCondition) | **POST** `/service/{service_id}/version/{version_id}/condition` | Create a condition
[**DeleteCondition**](ConditionAPI.md#DeleteCondition) | **DELETE** `/service/{service_id}/version/{version_id}/condition/{condition_name}` | Delete a condition
[**GetCondition**](ConditionAPI.md#GetCondition) | **GET** `/service/{service_id}/version/{version_id}/condition/{condition_name}` | Describe a condition
[**ListConditions**](ConditionAPI.md#ListConditions) | **GET** `/service/{service_id}/version/{version_id}/condition` | List conditions
[**UpdateCondition**](ConditionAPI.md#UpdateCondition) | **PUT** `/service/{service_id}/version/{version_id}/condition/{condition_name}` | Update a condition



## CreateCondition

Create a condition



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    name := "name_example" // string | Name of the condition. Required. (optional)
    priority := "priority_example" // string | A numeric string. Priority determines execution order. Lower numbers execute first. (optional) (default to "100")
    statement := "statement_example" // string | A conditional expression in VCL used to determine if the condition is met. (optional)
    serviceId2 := "serviceId_example" // string |  (optional)
    version := "version_example" // string | A numeric string that represents the service version. (optional)
    type_ := "type__example" // string | Type of the condition. Required. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConditionAPI.CreateCondition(ctx, serviceId, versionId).Comment(comment).Name(name).Priority(priority).Statement(statement).ServiceId2(serviceId2).Version(version).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionAPI.CreateCondition`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCondition`: ConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConditionAPI.CreateCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **string** | A freeform descriptive note. |  **name** | **string** | Name of the condition. Required. |  **priority** | **string** | A numeric string. Priority determines execution order. Lower numbers execute first. | [default to &quot;100&quot;] **statement** | **string** | A conditional expression in VCL used to determine if the condition is met. |  **serviceId2** | **string** |  |  **version** | **string** | A numeric string that represents the service version. |  **type_** | **string** | Type of the condition. Required. | 

### Return type

[**ConditionResponse**](ConditionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteCondition

Delete a condition



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    conditionName := "conditionName_example" // string | Name of the condition. Required.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConditionAPI.DeleteCondition(ctx, serviceId, versionId, conditionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionAPI.DeleteCondition`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCondition`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ConditionAPI.DeleteCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**conditionName** | **string** | Name of the condition. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConditionRequest struct via the builder pattern


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


## GetCondition

Describe a condition



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    conditionName := "conditionName_example" // string | Name of the condition. Required.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConditionAPI.GetCondition(ctx, serviceId, versionId, conditionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionAPI.GetCondition`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCondition`: ConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConditionAPI.GetCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**conditionName** | **string** | Name of the condition. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConditionResponse**](ConditionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListConditions

List conditions



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConditionAPI.ListConditions(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionAPI.ListConditions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditions`: []ConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConditionAPI.ListConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConditionResponse**](ConditionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateCondition

Update a condition



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    conditionName := "conditionName_example" // string | Name of the condition. Required.
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    name := "name_example" // string | Name of the condition. Required. (optional)
    priority := "priority_example" // string | A numeric string. Priority determines execution order. Lower numbers execute first. (optional) (default to "100")
    statement := "statement_example" // string | A conditional expression in VCL used to determine if the condition is met. (optional)
    serviceId2 := "serviceId_example" // string |  (optional)
    version := "version_example" // string | A numeric string that represents the service version. (optional)
    type_ := "type__example" // string | Type of the condition. Required. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConditionAPI.UpdateCondition(ctx, serviceId, versionId, conditionName).Comment(comment).Name(name).Priority(priority).Statement(statement).ServiceId2(serviceId2).Version(version).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionAPI.UpdateCondition`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCondition`: ConditionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConditionAPI.UpdateCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**conditionName** | **string** | Name of the condition. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **string** | A freeform descriptive note. |  **name** | **string** | Name of the condition. Required. |  **priority** | **string** | A numeric string. Priority determines execution order. Lower numbers execute first. | [default to &quot;100&quot;] **statement** | **string** | A conditional expression in VCL used to determine if the condition is met. |  **serviceId2** | **string** |  |  **version** | **string** | A numeric string that represents the service version. |  **type_** | **string** | Type of the condition. Required. | 

### Return type

[**ConditionResponse**](ConditionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

