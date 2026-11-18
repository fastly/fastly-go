# CacheSettingsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCacheSettings**](CacheSettingsAPI.md#CreateCacheSettings) | **POST** `/service/{service_id}/version/{version_id}/cache_settings` | Create a cache settings object
[**DeleteCacheSettings**](CacheSettingsAPI.md#DeleteCacheSettings) | **DELETE** `/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}` | Delete a cache settings object
[**GetCacheSettings**](CacheSettingsAPI.md#GetCacheSettings) | **GET** `/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}` | Get a cache settings object
[**ListCacheSettings**](CacheSettingsAPI.md#ListCacheSettings) | **GET** `/service/{service_id}/version/{version_id}/cache_settings` | List cache settings objects
[**UpdateCacheSettings**](CacheSettingsAPI.md#UpdateCacheSettings) | **PUT** `/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}` | Update a cache settings object



## CreateCacheSettings

Create a cache settings object



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
    action := "action_example" // string | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  (optional)
    cacheCondition := "cacheCondition_example" // string | Name of the cache condition controlling when this configuration applies. (optional)
    name := "name_example" // string | Name for the cache settings object. (optional)
    staleTTL := "staleTtl_example" // string | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as 'stale if error'). (optional)
    ttl := "ttl_example" // string | Maximum time to consider the object fresh in the cache (the cache 'time to live'). (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CacheSettingsAPI.CreateCacheSettings(ctx, serviceID, versionID).Action(action).CacheCondition(cacheCondition).Name(name).StaleTTL(staleTTL).TTL(ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.CreateCacheSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCacheSettings`: CacheSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.CreateCacheSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCacheSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  |  **cacheCondition** | **string** | Name of the cache condition controlling when this configuration applies. |  **name** | **string** | Name for the cache settings object. |  **staleTTL** | **string** | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;). |  **ttl** | **string** | Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;). | 

### Return type

[**CacheSettingResponse**](CacheSettingResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteCacheSettings

Delete a cache settings object



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
    cacheSettingsName := "cacheSettingsName_example" // string | Name for the cache settings object.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CacheSettingsAPI.DeleteCacheSettings(ctx, serviceID, versionID, cacheSettingsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.DeleteCacheSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCacheSettings`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.DeleteCacheSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**cacheSettingsName** | **string** | Name for the cache settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCacheSettingsRequest struct via the builder pattern


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


## GetCacheSettings

Get a cache settings object



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
    cacheSettingsName := "cacheSettingsName_example" // string | Name for the cache settings object.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CacheSettingsAPI.GetCacheSettings(ctx, serviceID, versionID, cacheSettingsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.GetCacheSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCacheSettings`: CacheSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.GetCacheSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**cacheSettingsName** | **string** | Name for the cache settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCacheSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CacheSettingResponse**](CacheSettingResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListCacheSettings

List cache settings objects



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
    resp, r, err := apiClient.CacheSettingsAPI.ListCacheSettings(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.ListCacheSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCacheSettings`: []CacheSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.ListCacheSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCacheSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CacheSettingResponse**](CacheSettingResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateCacheSettings

Update a cache settings object



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
    cacheSettingsName := "cacheSettingsName_example" // string | Name for the cache settings object.
    action := "action_example" // string | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  (optional)
    cacheCondition := "cacheCondition_example" // string | Name of the cache condition controlling when this configuration applies. (optional)
    name := "name_example" // string | Name for the cache settings object. (optional)
    staleTTL := "staleTtl_example" // string | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as 'stale if error'). (optional)
    ttl := "ttl_example" // string | Maximum time to consider the object fresh in the cache (the cache 'time to live'). (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CacheSettingsAPI.UpdateCacheSettings(ctx, serviceID, versionID, cacheSettingsName).Action(action).CacheCondition(cacheCondition).Name(name).StaleTTL(staleTTL).TTL(ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.UpdateCacheSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCacheSettings`: CacheSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.UpdateCacheSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**cacheSettingsName** | **string** | Name for the cache settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCacheSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  |  **cacheCondition** | **string** | Name of the cache condition controlling when this configuration applies. |  **name** | **string** | Name for the cache settings object. |  **staleTTL** | **string** | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;). |  **ttl** | **string** | Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;). | 

### Return type

[**CacheSettingResponse**](CacheSettingResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
