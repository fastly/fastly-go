# ResponseObjectAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteResponseObject**](ResponseObjectAPI.md#DeleteResponseObject) | **DELETE** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Delete a Response Object
[**GetResponseObject**](ResponseObjectAPI.md#GetResponseObject) | **GET** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Get a Response object
[**ListResponseObjects**](ResponseObjectAPI.md#ListResponseObjects) | **GET** `/service/{service_id}/version/{version_id}/response_object` | List Response objects



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    responseObjectName := "responseObjectName_example" // string | Name for the request settings.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.DeleteResponseObject(ctx, serviceID, versionID, responseObjectName).Execute()
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
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**responseObjectName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResponseObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    responseObjectName := "responseObjectName_example" // string | Name for the request settings.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.GetResponseObject(ctx, serviceID, versionID, responseObjectName).Execute()
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
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**responseObjectName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponseObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseObjectResponse**](ResponseObjectResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ResponseObjectAPI.ListResponseObjects(ctx, serviceID, versionID).Execute()
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
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResponseObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ResponseObjectResponse**](ResponseObjectResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
