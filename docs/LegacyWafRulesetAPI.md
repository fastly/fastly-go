# LegacyWafRulesetAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWafRuleset**](LegacyWafRulesetAPI.md#GetWafRuleset) | **GET** `/service/{service_id}/wafs/{firewall_id}/ruleset` | Get a WAF ruleset
[**GetWafRulesetVcl**](LegacyWafRulesetAPI.md#GetWafRulesetVcl) | **GET** `/service/{service_id}/wafs/{firewall_id}/ruleset/preview` | Generate WAF ruleset VCL
[**UpdateWafRuleset**](LegacyWafRulesetAPI.md#UpdateWafRuleset) | **PATCH** `/service/{service_id}/wafs/{firewall_id}/ruleset` | Update a WAF ruleset



## GetWafRuleset

Get a WAF ruleset



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
    resp, r, err := apiClient.LegacyWafRulesetAPI.GetWafRuleset(ctx, serviceID, firewallID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRulesetAPI.GetWafRuleset`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafRuleset`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRulesetAPI.GetWafRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafRulesetRequest struct via the builder pattern


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


## GetWafRulesetVcl

Generate WAF ruleset VCL



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
    resp, r, err := apiClient.LegacyWafRulesetAPI.GetWafRulesetVcl(ctx, serviceID, firewallID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRulesetAPI.GetWafRulesetVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafRulesetVcl`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRulesetAPI.GetWafRulesetVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafRulesetVclRequest struct via the builder pattern


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


## UpdateWafRuleset

Update a WAF ruleset



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
    resp, r, err := apiClient.LegacyWafRulesetAPI.UpdateWafRuleset(ctx, serviceID, firewallID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRulesetAPI.UpdateWafRuleset`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafRuleset`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRulesetAPI.UpdateWafRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafRulesetRequest struct via the builder pattern


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
