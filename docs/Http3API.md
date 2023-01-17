# HTTP3API

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTP3**](Http3API.md#CreateHTTP3) | **POST** `/service/{service_id}/version/{version_id}/http3` | Enable support for HTTP/3
[**DeleteHTTP3**](Http3API.md#DeleteHTTP3) | **DELETE** `/service/{service_id}/version/{version_id}/http3` | Disable support for HTTP/3
[**GetHTTP3**](Http3API.md#GetHTTP3) | **GET** `/service/{service_id}/version/{version_id}/http3` | Get HTTP/3 status



## CreateHTTP3

Enable support for HTTP/3



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    serviceID2 := "serviceId_example" // string |  (optional)
    version := int32(56) // int32 |  (optional)
    createdAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    deletedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    updatedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    featureRevision := int32(56) // int32 | Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HTTP3API.CreateHTTP3(ctx, serviceID, versionID).ServiceID2(serviceID2).Version(version).CreatedAt(createdAt).DeletedAt(deletedAt).UpdatedAt(updatedAt).FeatureRevision(featureRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HTTP3API.CreateHTTP3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHTTP3`: HTTP3
    fmt.Fprintf(os.Stdout, "Response from `HTTP3API.CreateHTTP3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHTTP3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceID2** | **string** |  |  **version** | **int32** |  |  **createdAt** | **time.Time** | Date and time in ISO 8601 format. |  **deletedAt** | **time.Time** | Date and time in ISO 8601 format. |  **updatedAt** | **time.Time** | Date and time in ISO 8601 format. |  **featureRevision** | **int32** | Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision. | 

### Return type

[**HTTP3**](Http3.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteHTTP3

Disable support for HTTP/3



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
    resp, r, err := apiClient.HTTP3API.DeleteHTTP3(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HTTP3API.DeleteHTTP3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHTTP3`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `HTTP3API.DeleteHTTP3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHTTP3Request struct via the builder pattern


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


## GetHTTP3

Get HTTP/3 status



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
    resp, r, err := apiClient.HTTP3API.GetHTTP3(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HTTP3API.GetHTTP3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHTTP3`: HTTP3
    fmt.Fprintf(os.Stdout, "Response from `HTTP3API.GetHTTP3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHTTP3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HTTP3**](Http3.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
