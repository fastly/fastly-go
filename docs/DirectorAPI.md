# DirectorAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirector**](DirectorAPI.md#CreateDirector) | **POST** `/service/{service_id}/version/{version_id}/director` | Create a director
[**DeleteDirector**](DirectorAPI.md#DeleteDirector) | **DELETE** `/service/{service_id}/version/{version_id}/director/{director_name}` | Delete a director
[**GetDirector**](DirectorAPI.md#GetDirector) | **GET** `/service/{service_id}/version/{version_id}/director/{director_name}` | Get a director
[**ListDirectors**](DirectorAPI.md#ListDirectors) | **GET** `/service/{service_id}/version/{version_id}/director` | List directors
[**UpdateDirector**](DirectorAPI.md#UpdateDirector) | **PUT** `/service/{service_id}/version/{version_id}/director/{director_name}` | Update a director



## CreateDirector

Create a director



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
    backends := []openapiclient.Backend{*openapiclient.NewBackend()} // []Backend | List of backends associated to a director. (optional)
    capacity := int32(56) // int32 | Unused. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    name := "name_example" // string | Name for the Director. (optional)
    quorum := int32(56) // int32 | The percentage of capacity that needs to be up for a director to be considered up. `0` to `100`. (optional) (default to 75)
    shield := "shield_example" // string | Selected POP to serve as a shield for the backends. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding. (optional) (default to "null")
    resourceType := int32(56) // int32 | What type of load balance group to use. (optional) (default to 1)
    retries := int32(56) // int32 | How many backends to search if it fails. (optional) (default to 5)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorAPI.CreateDirector(ctx, serviceID, versionID).Backends(backends).Capacity(capacity).Comment(comment).Name(name).Quorum(quorum).Shield(shield).ResourceType(resourceType).Retries(retries).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorAPI.CreateDirector`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirector`: DirectorResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectorAPI.CreateDirector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backends** | [**[]Backend**](Backend.md) | List of backends associated to a director. |  **capacity** | **int32** | Unused. |  **comment** | **string** | A freeform descriptive note. |  **name** | **string** | Name for the Director. |  **quorum** | **int32** | The percentage of capacity that needs to be up for a director to be considered up. `0` to `100`. | [default to 75] **shield** | **string** | Selected POP to serve as a shield for the backends. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [default to &quot;null&quot;] **resourceType** | **int32** | What type of load balance group to use. | [default to 1] **retries** | **int32** | How many backends to search if it fails. | [default to 5]

### Return type

[**DirectorResponse**](DirectorResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDirector

Delete a director



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
    directorName := "directorName_example" // string | Name for the Director.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorAPI.DeleteDirector(ctx, serviceID, versionID, directorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorAPI.DeleteDirector`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDirector`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DirectorAPI.DeleteDirector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**directorName** | **string** | Name for the Director. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectorRequest struct via the builder pattern


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


## GetDirector

Get a director



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
    directorName := "directorName_example" // string | Name for the Director.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorAPI.GetDirector(ctx, serviceID, versionID, directorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorAPI.GetDirector`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirector`: DirectorResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectorAPI.GetDirector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**directorName** | **string** | Name for the Director. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectorResponse**](DirectorResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDirectors

List directors



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
    resp, r, err := apiClient.DirectorAPI.ListDirectors(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorAPI.ListDirectors`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDirectors`: []DirectorResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectorAPI.ListDirectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDirectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DirectorResponse**](DirectorResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDirector

Update a director



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
    directorName := "directorName_example" // string | Name for the Director.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DirectorAPI.UpdateDirector(ctx, serviceID, versionID, directorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectorAPI.UpdateDirector`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDirector`: DirectorResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectorAPI.UpdateDirector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**directorName** | **string** | Name for the Director. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDirectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectorResponse**](DirectorResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
