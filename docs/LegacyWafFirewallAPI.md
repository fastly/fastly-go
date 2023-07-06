# LegacyWafFirewallAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLegacyWafFirewallService**](LegacyWafFirewallAPI.md#CreateLegacyWafFirewallService) | **POST** `/service/{service_id}/version/{version_id}/wafs` | Create a firewall
[**DisableLegacyWafFirewall**](LegacyWafFirewallAPI.md#DisableLegacyWafFirewall) | **PATCH** `/wafs/{firewall_id}/disable` | Disable a firewall
[**EnableLegacyWafFirewall**](LegacyWafFirewallAPI.md#EnableLegacyWafFirewall) | **PATCH** `/wafs/{firewall_id}/enable` | Enable a firewall
[**GetLegacyWafFirewall**](LegacyWafFirewallAPI.md#GetLegacyWafFirewall) | **GET** `/wafs/{firewall_id}` | Get a firewall object
[**GetLegacyWafFirewallService**](LegacyWafFirewallAPI.md#GetLegacyWafFirewallService) | **GET** `/service/{service_id}/version/{version_id}/wafs/{firewall_id}` | Get a firewall
[**ListLegacyWafFirewalls**](LegacyWafFirewallAPI.md#ListLegacyWafFirewalls) | **GET** `/wafs` | List active firewalls
[**ListLegacyWafFirewallsService**](LegacyWafFirewallAPI.md#ListLegacyWafFirewallsService) | **GET** `/service/{service_id}/version/{version_id}/wafs` | List firewalls
[**UpdateLegacyWafFirewallService**](LegacyWafFirewallAPI.md#UpdateLegacyWafFirewallService) | **PATCH** `/service/{service_id}/version/{version_id}/wafs/{firewall_id}` | Update a firewall



## CreateLegacyWafFirewallService

Create a firewall



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.CreateLegacyWafFirewallService(ctx, serviceID, versionID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.CreateLegacyWafFirewallService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLegacyWafFirewallService`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.CreateLegacyWafFirewallService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLegacyWafFirewallServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DisableLegacyWafFirewall

Disable a firewall



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.DisableLegacyWafFirewall(ctx, firewallID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.DisableLegacyWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableLegacyWafFirewall`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.DisableLegacyWafFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableLegacyWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## EnableLegacyWafFirewall

Enable a firewall



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.EnableLegacyWafFirewall(ctx, firewallID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.EnableLegacyWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableLegacyWafFirewall`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.EnableLegacyWafFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableLegacyWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetLegacyWafFirewall

Get a firewall object



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.
    include := "configuration_set" // string | Include relationships. Optional, comma separated values. Permitted values: `configuration_set`, `owasp`, `rules`, `rule_statuses`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.GetLegacyWafFirewall(ctx, firewallID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.GetLegacyWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegacyWafFirewall`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.GetLegacyWafFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include relationships. Optional, comma separated values. Permitted values: `configuration_set`, `owasp`, `rules`, `rule_statuses`.  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetLegacyWafFirewallService

Get a firewall



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.GetLegacyWafFirewallService(ctx, serviceID, versionID, firewallID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.GetLegacyWafFirewallService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegacyWafFirewallService`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.GetLegacyWafFirewallService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyWafFirewallServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLegacyWafFirewalls

List active firewalls



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
    filterRulesRuleID := "filterRulesRuleId_example" // string | Limit the returned firewalls to a specific rule ID. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "configuration_set" // string | Include relationships. Optional, comma separated values. Permitted values: `configuration_set`, `owasp`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.ListLegacyWafFirewalls(ctx).FilterRulesRuleID(filterRulesRuleID).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.ListLegacyWafFirewalls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWafFirewalls`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.ListLegacyWafFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWafFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRulesRuleID** | **string** | Limit the returned firewalls to a specific rule ID. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional, comma separated values. Permitted values: `configuration_set`, `owasp`.  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLegacyWafFirewallsService

List firewalls



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
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "configuration_set" // string | Include relationships. Optional, comma separated values. Permitted values: `configuration_set`, `owasp`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.ListLegacyWafFirewallsService(ctx, serviceID, versionID).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.ListLegacyWafFirewallsService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWafFirewallsService`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.ListLegacyWafFirewallsService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWafFirewallsServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional, comma separated values. Permitted values: `configuration_set`, `owasp`.  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLegacyWafFirewallService

Update a firewall



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a Firewall.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafFirewallAPI.UpdateLegacyWafFirewallService(ctx, serviceID, versionID, firewallID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafFirewallAPI.UpdateLegacyWafFirewallService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLegacyWafFirewallService`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafFirewallAPI.UpdateLegacyWafFirewallService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLegacyWafFirewallServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
