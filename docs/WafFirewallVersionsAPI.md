# WafFirewallVersionsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneWafFirewallVersion**](WafFirewallVersionsAPI.md#CloneWafFirewallVersion) | **PUT** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/clone` | Clone a firewall version
[**CreateWafFirewallVersion**](WafFirewallVersionsAPI.md#CreateWafFirewallVersion) | **POST** `/waf/firewalls/{firewall_id}/versions` | Create a firewall version
[**DeployActivateWafFirewallVersion**](WafFirewallVersionsAPI.md#DeployActivateWafFirewallVersion) | **PUT** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/activate` | Deploy or activate a firewall version
[**GetWafFirewallVersion**](WafFirewallVersionsAPI.md#GetWafFirewallVersion) | **GET** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}` | Get a firewall version
[**ListWafFirewallVersions**](WafFirewallVersionsAPI.md#ListWafFirewallVersions) | **GET** `/waf/firewalls/{firewall_id}/versions` | List firewall versions
[**UpdateWafFirewallVersion**](WafFirewallVersionsAPI.md#UpdateWafFirewallVersion) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}` | Update a firewall version



## CloneWafFirewallVersion

Clone a firewall version



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    firewallVersionNumber := int32(56) // int32 | Integer identifying a WAF firewall version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallVersionsAPI.CloneWafFirewallVersion(ctx, firewallID, firewallVersionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallVersionsAPI.CloneWafFirewallVersion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneWafFirewallVersion`: WafFirewallVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallVersionsAPI.CloneWafFirewallVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneWafFirewallVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WafFirewallVersionResponse**](WafFirewallVersionResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateWafFirewallVersion

Create a firewall version



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    wafFirewallVersion := *openapiclient.NewWafFirewallVersion() // WafFirewallVersion |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallVersionsAPI.CreateWafFirewallVersion(ctx, firewallID).WafFirewallVersion(wafFirewallVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallVersionsAPI.CreateWafFirewallVersion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWafFirewallVersion`: WafFirewallVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallVersionsAPI.CreateWafFirewallVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafFirewallVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafFirewallVersion** | [**WafFirewallVersion**](WafFirewallVersion.md) |  | 

### Return type

[**WafFirewallVersionResponse**](WafFirewallVersionResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeployActivateWafFirewallVersion

Deploy or activate a firewall version



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    firewallVersionNumber := int32(56) // int32 | Integer identifying a WAF firewall version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallVersionsAPI.DeployActivateWafFirewallVersion(ctx, firewallID, firewallVersionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallVersionsAPI.DeployActivateWafFirewallVersion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployActivateWafFirewallVersion`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallVersionsAPI.DeployActivateWafFirewallVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployActivateWafFirewallVersionRequest struct via the builder pattern


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


## GetWafFirewallVersion

Get a firewall version



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    firewallVersionNumber := int32(56) // int32 | Integer identifying a WAF firewall version.
    include := "waf_firewall,waf_active_rules" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_firewall` and `waf_active_rules`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallVersionsAPI.GetWafFirewallVersion(ctx, firewallID, firewallVersionNumber).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallVersionsAPI.GetWafFirewallVersion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafFirewallVersion`: WafFirewallVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallVersionsAPI.GetWafFirewallVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafFirewallVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_firewall` and `waf_active_rules`.  | 

### Return type

[**WafFirewallVersionResponse**](WafFirewallVersionResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafFirewallVersions

List firewall versions



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    include := "waf_firewall" // string | Include relationships. Optional. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallVersionsAPI.ListWafFirewallVersions(ctx, firewallID).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallVersionsAPI.ListWafFirewallVersions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafFirewallVersions`: WafFirewallVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallVersionsAPI.ListWafFirewallVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafFirewallVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include relationships. Optional. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**WafFirewallVersionsResponse**](WafFirewallVersionsResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateWafFirewallVersion

Update a firewall version



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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    firewallVersionNumber := int32(56) // int32 | Integer identifying a WAF firewall version.
    wafFirewallVersion := *openapiclient.NewWafFirewallVersion() // WafFirewallVersion |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallVersionsAPI.UpdateWafFirewallVersion(ctx, firewallID, firewallVersionNumber).WafFirewallVersion(wafFirewallVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallVersionsAPI.UpdateWafFirewallVersion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafFirewallVersion`: WafFirewallVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallVersionsAPI.UpdateWafFirewallVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafFirewallVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafFirewallVersion** | [**WafFirewallVersion**](WafFirewallVersion.md) |  | 

### Return type

[**WafFirewallVersionResponse**](WafFirewallVersionResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
