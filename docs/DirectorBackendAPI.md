# DirectorBackendAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectorBackend**](DirectorBackendAPI.md#CreateDirectorBackend) | **POST** `/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}` | Create a director-backend relationship
[**DeleteDirectorBackend**](DirectorBackendAPI.md#DeleteDirectorBackend) | **DELETE** `/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}` | Delete a director-backend relationship
[**GetDirectorBackend**](DirectorBackendAPI.md#GetDirectorBackend) | **GET** `/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}` | Get a director-backend relationship



## CreateDirectorBackend

Create a director-backend relationship



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
    directorName := "directorName_example" // string | Name for the Director.
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    backendName := "backendName_example" // string | The name of the backend.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorBackendAPI.CreateDirectorBackend(ctx, directorName, serviceID, versionID, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorBackendAPI.CreateDirectorBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectorBackend`: DirectorBackend
    fmt.Fprintf(os.Stdout, "Response from `DirectorBackendAPI.CreateDirectorBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directorName** | **string** | Name for the Director. | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**backendName** | **string** | The name of the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectorBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectorBackend**](DirectorBackend.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDirectorBackend

Delete a director-backend relationship



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
    directorName := "directorName_example" // string | Name for the Director.
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    backendName := "backendName_example" // string | The name of the backend.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorBackendAPI.DeleteDirectorBackend(ctx, directorName, serviceID, versionID, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorBackendAPI.DeleteDirectorBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDirectorBackend`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DirectorBackendAPI.DeleteDirectorBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directorName** | **string** | Name for the Director. | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**backendName** | **string** | The name of the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectorBackendRequest struct via the builder pattern


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


## GetDirectorBackend

Get a director-backend relationship



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
    directorName := "directorName_example" // string | Name for the Director.
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    backendName := "backendName_example" // string | The name of the backend.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorBackendAPI.GetDirectorBackend(ctx, directorName, serviceID, versionID, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorBackendAPI.GetDirectorBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectorBackend`: DirectorBackend
    fmt.Fprintf(os.Stdout, "Response from `DirectorBackendAPI.GetDirectorBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directorName** | **string** | Name for the Director. | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**backendName** | **string** | The name of the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectorBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectorBackend**](DirectorBackend.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
