# GzipAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGzipConfig**](GzipAPI.md#CreateGzipConfig) | **POST** `/service/{service_id}/version/{version_id}/gzip` | Create a gzip configuration
[**DeleteGzipConfig**](GzipAPI.md#DeleteGzipConfig) | **DELETE** `/service/{service_id}/version/{version_id}/gzip/{gzip_name}` | Delete a gzip configuration
[**GetGzipConfigs**](GzipAPI.md#GetGzipConfigs) | **GET** `/service/{service_id}/version/{version_id}/gzip/{gzip_name}` | Get a gzip configuration
[**ListGzipConfigs**](GzipAPI.md#ListGzipConfigs) | **GET** `/service/{service_id}/version/{version_id}/gzip` | List gzip configurations
[**UpdateGzipConfig**](GzipAPI.md#UpdateGzipConfig) | **PUT** `/service/{service_id}/version/{version_id}/gzip/{gzip_name}` | Update a gzip configuration



## CreateGzipConfig

Create a gzip configuration



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
    cacheCondition := "cacheCondition_example" // string | Name of the cache condition controlling when this configuration applies. (optional)
    contentTypes := "contentTypes_example" // string | Space-separated list of content types to compress. If you omit this field a default list will be used. (optional)
    extensions := "extensions_example" // string | Space-separated list of file extensions to compress. If you omit this field a default list will be used. (optional)
    name := "name_example" // string | Name of the gzip configuration. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.GzipAPI.CreateGzipConfig(ctx, serviceID, versionID).CacheCondition(cacheCondition).ContentTypes(contentTypes).Extensions(extensions).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GzipAPI.CreateGzipConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGzipConfig`: GzipResponse
    fmt.Fprintf(os.Stdout, "Response from `GzipAPI.CreateGzipConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGzipConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cacheCondition** | **string** | Name of the cache condition controlling when this configuration applies. |  **contentTypes** | **string** | Space-separated list of content types to compress. If you omit this field a default list will be used. |  **extensions** | **string** | Space-separated list of file extensions to compress. If you omit this field a default list will be used. |  **name** | **string** | Name of the gzip configuration. | 

### Return type

[**GzipResponse**](GzipResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteGzipConfig

Delete a gzip configuration



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
    gzipName := "gzipName_example" // string | Name of the gzip configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.GzipAPI.DeleteGzipConfig(ctx, serviceID, versionID, gzipName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GzipAPI.DeleteGzipConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGzipConfig`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `GzipAPI.DeleteGzipConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**gzipName** | **string** | Name of the gzip configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGzipConfigRequest struct via the builder pattern


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


## GetGzipConfigs

Get a gzip configuration



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
    gzipName := "gzipName_example" // string | Name of the gzip configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.GzipAPI.GetGzipConfigs(ctx, serviceID, versionID, gzipName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GzipAPI.GetGzipConfigs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGzipConfigs`: GzipResponse
    fmt.Fprintf(os.Stdout, "Response from `GzipAPI.GetGzipConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**gzipName** | **string** | Name of the gzip configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGzipConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GzipResponse**](GzipResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListGzipConfigs

List gzip configurations



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
    resp, r, err := apiClient.GzipAPI.ListGzipConfigs(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GzipAPI.ListGzipConfigs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGzipConfigs`: []GzipResponse
    fmt.Fprintf(os.Stdout, "Response from `GzipAPI.ListGzipConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGzipConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GzipResponse**](GzipResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateGzipConfig

Update a gzip configuration



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
    gzipName := "gzipName_example" // string | Name of the gzip configuration.
    cacheCondition := "cacheCondition_example" // string | Name of the cache condition controlling when this configuration applies. (optional)
    contentTypes := "contentTypes_example" // string | Space-separated list of content types to compress. If you omit this field a default list will be used. (optional)
    extensions := "extensions_example" // string | Space-separated list of file extensions to compress. If you omit this field a default list will be used. (optional)
    name := "name_example" // string | Name of the gzip configuration. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.GzipAPI.UpdateGzipConfig(ctx, serviceID, versionID, gzipName).CacheCondition(cacheCondition).ContentTypes(contentTypes).Extensions(extensions).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GzipAPI.UpdateGzipConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGzipConfig`: GzipResponse
    fmt.Fprintf(os.Stdout, "Response from `GzipAPI.UpdateGzipConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**gzipName** | **string** | Name of the gzip configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGzipConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cacheCondition** | **string** | Name of the cache condition controlling when this configuration applies. |  **contentTypes** | **string** | Space-separated list of content types to compress. If you omit this field a default list will be used. |  **extensions** | **string** | Space-separated list of file extensions to compress. If you omit this field a default list will be used. |  **name** | **string** | Name of the gzip configuration. | 

### Return type

[**GzipResponse**](GzipResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
