# ImageOptimizerDefaultSettingsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultSettings**](ImageOptimizerDefaultSettingsAPI.md#GetDefaultSettings) | **GET** `/service/{service_id}/version/{version_id}/image_optimizer_default_settings` | Get current Image Optimizer Default Settings
[**UpdateDefaultSettings**](ImageOptimizerDefaultSettingsAPI.md#UpdateDefaultSettings) | **PATCH** `/service/{service_id}/version/{version_id}/image_optimizer_default_settings` | Update Image Optimizer Default Settings



## GetDefaultSettings

Get current Image Optimizer Default Settings



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
    resp, r, err := apiClient.ImageOptimizerDefaultSettingsAPI.GetDefaultSettings(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageOptimizerDefaultSettingsAPI.GetDefaultSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultSettings`: DefaultSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageOptimizerDefaultSettingsAPI.GetDefaultSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefaultSettingsResponse**](DefaultSettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDefaultSettings

Update Image Optimizer Default Settings



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
    defaultSettings := *openapiclient.NewDefaultSettings() // DefaultSettings |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ImageOptimizerDefaultSettingsAPI.UpdateDefaultSettings(ctx, serviceId, versionId).DefaultSettings(defaultSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageOptimizerDefaultSettingsAPI.UpdateDefaultSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefaultSettings`: DefaultSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageOptimizerDefaultSettingsAPI.UpdateDefaultSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **defaultSettings** | [**DefaultSettings**](DefaultSettings.md) |  | 

### Return type

[**DefaultSettingsResponse**](DefaultSettingsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

