# RequestSettingsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestSettings**](RequestSettingsAPI.md#CreateRequestSettings) | **POST** `/service/{service_id}/version/{version_id}/request_settings` | Create a Request Settings object
[**DeleteRequestSettings**](RequestSettingsAPI.md#DeleteRequestSettings) | **DELETE** `/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}` | Delete a Request Settings object
[**GetRequestSettings**](RequestSettingsAPI.md#GetRequestSettings) | **GET** `/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}` | Get a Request Settings object
[**ListRequestSettings**](RequestSettingsAPI.md#ListRequestSettings) | **GET** `/service/{service_id}/version/{version_id}/request_settings` | List Request Settings objects
[**UpdateRequestSettings**](RequestSettingsAPI.md#UpdateRequestSettings) | **PUT** `/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}` | Update a Request Settings object



## CreateRequestSettings

Create a Request Settings object



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
    resp, r, err := apiClient.RequestSettingsAPI.CreateRequestSettings(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.CreateRequestSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestSettings`: RequestSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.CreateRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestSettingsResponse**](RequestSettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteRequestSettings

Delete a Request Settings object



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
    requestSettingsName := "requestSettingsName_example" // string | Name for the request settings.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RequestSettingsAPI.DeleteRequestSettings(ctx, serviceId, versionId, requestSettingsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.DeleteRequestSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRequestSettings`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.DeleteRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**requestSettingsName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestSettingsRequest struct via the builder pattern


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


## GetRequestSettings

Get a Request Settings object



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
    requestSettingsName := "requestSettingsName_example" // string | Name for the request settings.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RequestSettingsAPI.GetRequestSettings(ctx, serviceId, versionId, requestSettingsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.GetRequestSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestSettings`: RequestSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.GetRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**requestSettingsName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestSettingsResponse**](RequestSettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListRequestSettings

List Request Settings objects



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
    resp, r, err := apiClient.RequestSettingsAPI.ListRequestSettings(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.ListRequestSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestSettings`: []RequestSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.ListRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RequestSettingsResponse**](RequestSettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateRequestSettings

Update a Request Settings object



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
    requestSettingsName := "requestSettingsName_example" // string | Name for the request settings.
    action := "action_example" // string | Allows you to terminate request handling and immediately perform an action. (optional)
    defaultHost := "defaultHost_example" // string | Sets the host header. (optional)
    hashKeys := "hashKeys_example" // string | Comma separated list of varnish request object fields that should be in the hash key. (optional)
    name := "name_example" // string | Name for the request settings. (optional)
    requestCondition := "requestCondition_example" // string | Condition which, if met, will select this configuration during a request. Optional. (optional)
    xff := "xff_example" // string | Short for X-Forwarded-For. (optional)
    bypassBusyWait := int32(56) // int32 | Disable collapsed forwarding, so you don't wait for other objects to origin. (optional)
    forceMiss := int32(56) // int32 | Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable. (optional)
    forceSsl := int32(56) // int32 | Forces the request use SSL (redirects a non-SSL to SSL). (optional)
    geoHeaders := int32(56) // int32 | Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers. (optional)
    maxStaleAge := int32(56) // int32 | How old an object is allowed to be to serve stale-if-error or stale-while-revalidate. (optional)
    timerSupport := int32(56) // int32 | Injects the X-Timer info into the request for viewing origin fetch durations. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.RequestSettingsAPI.UpdateRequestSettings(ctx, serviceId, versionId, requestSettingsName).Action(action).DefaultHost(defaultHost).HashKeys(hashKeys).Name(name).RequestCondition(requestCondition).Xff(xff).BypassBusyWait(bypassBusyWait).ForceMiss(forceMiss).ForceSsl(forceSsl).GeoHeaders(geoHeaders).MaxStaleAge(maxStaleAge).TimerSupport(timerSupport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.UpdateRequestSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestSettings`: RequestSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.UpdateRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**requestSettingsName** | **string** | Name for the request settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Allows you to terminate request handling and immediately perform an action. |  **defaultHost** | **string** | Sets the host header. |  **hashKeys** | **string** | Comma separated list of varnish request object fields that should be in the hash key. |  **name** | **string** | Name for the request settings. |  **requestCondition** | **string** | Condition which, if met, will select this configuration during a request. Optional. |  **xff** | **string** | Short for X-Forwarded-For. |  **bypassBusyWait** | **int32** | Disable collapsed forwarding, so you don&#39;t wait for other objects to origin. |  **forceMiss** | **int32** | Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable. |  **forceSsl** | **int32** | Forces the request use SSL (redirects a non-SSL to SSL). |  **geoHeaders** | **int32** | Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers. |  **maxStaleAge** | **int32** | How old an object is allowed to be to serve stale-if-error or stale-while-revalidate. |  **timerSupport** | **int32** | Injects the X-Timer info into the request for viewing origin fetch durations. | 

### Return type

[**RequestSettingsResponse**](RequestSettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

