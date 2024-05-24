# LegacyWafOwaspAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOwaspSettings**](LegacyWafOwaspAPI.md#CreateOwaspSettings) | **POST** `/service/{service_id}/wafs/{firewall_id}/owasp` | Create an OWASP settings object
[**GetOwaspSettings**](LegacyWafOwaspAPI.md#GetOwaspSettings) | **GET** `/service/{service_id}/wafs/{firewall_id}/owasp` | Get the OWASP settings object
[**UpdateOwaspSettings**](LegacyWafOwaspAPI.md#UpdateOwaspSettings) | **PATCH** `/service/{service_id}/wafs/{firewall_id}/owasp` | Update the OWASP settings object



## CreateOwaspSettings

Create an OWASP settings object



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafOwaspAPI.CreateOwaspSettings(ctx, serviceID, firewallID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafOwaspAPI.CreateOwaspSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOwaspSettings`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafOwaspAPI.CreateOwaspSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOwaspSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetOwaspSettings

Get the OWASP settings object



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafOwaspAPI.GetOwaspSettings(ctx, serviceID, firewallID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafOwaspAPI.GetOwaspSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOwaspSettings`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafOwaspAPI.GetOwaspSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwaspSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateOwaspSettings

Update the OWASP settings object



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafOwaspAPI.UpdateOwaspSettings(ctx, serviceID, firewallID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafOwaspAPI.UpdateOwaspSettings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOwaspSettings`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafOwaspAPI.UpdateOwaspSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwaspSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
