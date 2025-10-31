# SettingsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceSettings**](SettingsAPI.md#GetServiceSettings) | **GET** `/service/{service_id}/version/{version_id}/settings` | Get service settings
[**UpdateServiceSettings**](SettingsAPI.md#UpdateServiceSettings) | **PUT** `/service/{service_id}/version/{version_id}/settings` | Update service settings



## GetServiceSettings

Get service settings



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
    resp, r, err := apiClient.SettingsAPI.GetServiceSettings(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetServiceSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceSettings`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetServiceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateServiceSettings

Update service settings



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
    generalDefaultHost := "generalDefaultHost_example" // string | The default host name for the version. (optional)
    generalDefaultTtl := int32(56) // int32 | The default time-to-live (TTL) for the version. (optional)
    generalStaleIfError := true // bool | Enables serving a stale object if there is an error. (optional) (default to false)
    generalStaleIfErrorTtl := int32(56) // int32 | The default time-to-live (TTL) for serving the stale object for the version. (optional) (default to 43200)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SettingsAPI.UpdateServiceSettings(ctx, serviceId, versionId).GeneralDefaultHost(generalDefaultHost).GeneralDefaultTtl(generalDefaultTtl).GeneralStaleIfError(generalStaleIfError).GeneralStaleIfErrorTtl(generalStaleIfErrorTtl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateServiceSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceSettings`: SettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UpdateServiceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generalDefaultHost** | **string** | The default host name for the version. |  **generalDefaultTtl** | **int32** | The default time-to-live (TTL) for the version. |  **generalStaleIfError** | **bool** | Enables serving a stale object if there is an error. | [default to false] **generalStaleIfErrorTtl** | **int32** | The default time-to-live (TTL) for serving the stale object for the version. | [default to 43200]

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

