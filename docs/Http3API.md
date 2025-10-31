# Http3API

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHttp3**](Http3API.md#CreateHttp3) | **POST** `/service/{service_id}/version/{version_id}/http3` | Enable support for HTTP/3
[**DeleteHttp3**](Http3API.md#DeleteHttp3) | **DELETE** `/service/{service_id}/version/{version_id}/http3` | Disable support for HTTP/3
[**GetHttp3**](Http3API.md#GetHttp3) | **GET** `/service/{service_id}/version/{version_id}/http3` | Get HTTP/3 status



## CreateHttp3

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    serviceId2 := "serviceId_example" // string |  (optional)
    version := int32(56) // int32 |  (optional)
    createdAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    deletedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    updatedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    featureRevision := int32(56) // int32 | Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.Http3API.CreateHttp3(ctx, serviceId, versionId).ServiceId2(serviceId2).Version(version).CreatedAt(createdAt).DeletedAt(deletedAt).UpdatedAt(updatedAt).FeatureRevision(featureRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Http3API.CreateHttp3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHttp3`: Http3
    fmt.Fprintf(os.Stdout, "Response from `Http3API.CreateHttp3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHttp3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId2** | **string** |  |  **version** | **int32** |  |  **createdAt** | **time.Time** | Date and time in ISO 8601 format. |  **deletedAt** | **time.Time** | Date and time in ISO 8601 format. |  **updatedAt** | **time.Time** | Date and time in ISO 8601 format. |  **featureRevision** | **int32** | Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision. | 

### Return type

[**Http3**](Http3.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteHttp3

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.Http3API.DeleteHttp3(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Http3API.DeleteHttp3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHttp3`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `Http3API.DeleteHttp3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHttp3Request struct via the builder pattern


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


## GetHttp3

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.Http3API.GetHttp3(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Http3API.GetHttp3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttp3`: Http3
    fmt.Fprintf(os.Stdout, "Response from `Http3API.GetHttp3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttp3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Http3**](Http3.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

