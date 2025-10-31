# ResponseObjectAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResponseObject**](ResponseObjectAPI.md#CreateResponseObject) | **POST** `/service/{service_id}/version/{version_id}/response_object` | Create a Response object
[**DeleteResponseObject**](ResponseObjectAPI.md#DeleteResponseObject) | **DELETE** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Delete a Response Object
[**GetResponseObject**](ResponseObjectAPI.md#GetResponseObject) | **GET** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Get a Response object
[**ListResponseObjects**](ResponseObjectAPI.md#ListResponseObjects) | **GET** `/service/{service_id}/version/{version_id}/response_object` | List Response objects
[**UpdateResponseObject**](ResponseObjectAPI.md#UpdateResponseObject) | **PUT** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Update a Response object



## CreateResponseObject

Create a Response object



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
    createResponseObjectRequest := *openapiclient.NewCreateResponseObjectRequest() // CreateResponseObjectRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.CreateResponseObject(ctx, serviceId, versionId).CreateResponseObjectRequest(createResponseObjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResponseObjectAPI.CreateResponseObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResponseObject`: ResponseObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ResponseObjectAPI.CreateResponseObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResponseObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResponseObjectRequest** | [**CreateResponseObjectRequest**](CreateResponseObjectRequest.md) |  | 

### Return type

[**ResponseObjectResponse**](ResponseObjectResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteResponseObject

Delete a Response Object



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
    responseObjectName := "responseObjectName_example" // string | Name for the request settings.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.DeleteResponseObject(ctx, serviceId, versionId, responseObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResponseObjectAPI.DeleteResponseObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteResponseObject`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ResponseObjectAPI.DeleteResponseObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**responseObjectName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResponseObjectRequest struct via the builder pattern


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


## GetResponseObject

Get a Response object



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
    responseObjectName := "responseObjectName_example" // string | Name for the request settings.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.GetResponseObject(ctx, serviceId, versionId, responseObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResponseObjectAPI.GetResponseObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResponseObject`: ResponseObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ResponseObjectAPI.GetResponseObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**responseObjectName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponseObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseObjectResponse**](ResponseObjectResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListResponseObjects

List Response objects



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
    resp, r, err := apiClient.ResponseObjectAPI.ListResponseObjects(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResponseObjectAPI.ListResponseObjects`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResponseObjects`: []ResponseObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ResponseObjectAPI.ListResponseObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResponseObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ResponseObjectResponse**](ResponseObjectResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateResponseObject

Update a Response object



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
    responseObjectName := "responseObjectName_example" // string | Name for the request settings.
    createResponseObjectRequest := *openapiclient.NewCreateResponseObjectRequest() // CreateResponseObjectRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.UpdateResponseObject(ctx, serviceId, versionId, responseObjectName).CreateResponseObjectRequest(createResponseObjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResponseObjectAPI.UpdateResponseObject`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResponseObject`: ResponseObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ResponseObjectAPI.UpdateResponseObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**responseObjectName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResponseObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResponseObjectRequest** | [**CreateResponseObjectRequest**](CreateResponseObjectRequest.md) |  | 

### Return type

[**ResponseObjectResponse**](ResponseObjectResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

