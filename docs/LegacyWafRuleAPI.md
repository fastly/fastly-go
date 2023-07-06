# LegacyWafRuleAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLegacyWafFirewallRuleVcl**](LegacyWafRuleAPI.md#GetLegacyWafFirewallRuleVcl) | **GET** `/wafs/{firewall_id}/rules/{waf_rule_id}/vcl` | Get VCL for a rule associated with a firewall
[**GetLegacyWafRule**](LegacyWafRuleAPI.md#GetLegacyWafRule) | **GET** `/wafs/rules/{waf_rule_id}` | Get a rule
[**GetLegacyWafRuleVcl**](LegacyWafRuleAPI.md#GetLegacyWafRuleVcl) | **GET** `/wafs/rules/{waf_rule_id}/vcl` | Get VCL for a rule
[**ListLegacyWafRules**](LegacyWafRuleAPI.md#ListLegacyWafRules) | **GET** `/wafs/rules` | List rules in the latest configuration set



## GetLegacyWafFirewallRuleVcl

Get VCL for a rule associated with a firewall



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
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleAPI.GetLegacyWafFirewallRuleVcl(ctx, firewallID, wafRuleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleAPI.GetLegacyWafFirewallRuleVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegacyWafFirewallRuleVcl`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleAPI.GetLegacyWafFirewallRuleVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyWafFirewallRuleVclRequest struct via the builder pattern


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


## GetLegacyWafRule

Get a rule



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
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.
    filterConfigurationSetID := "filterConfigurationSetId_example" // string | Optional. Limit rule to a specific configuration set or pass \"all\" to search all configuration sets, including stale ones. (optional)
    include := "tags" // string | Include relationships. Optional. Comma separated values. Permitted values: `tags`, `rule_statuses`, `source`, and `vcl`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleAPI.GetLegacyWafRule(ctx, wafRuleID).FilterConfigurationSetID(filterConfigurationSetID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleAPI.GetLegacyWafRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegacyWafRule`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleAPI.GetLegacyWafRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyWafRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterConfigurationSetID** | **string** | Optional. Limit rule to a specific configuration set or pass \&quot;all\&quot; to search all configuration sets, including stale ones. |  **include** | **string** | Include relationships. Optional. Comma separated values. Permitted values: `tags`, `rule_statuses`, `source`, and `vcl`.  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetLegacyWafRuleVcl

Get VCL for a rule



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
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleAPI.GetLegacyWafRuleVcl(ctx, wafRuleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleAPI.GetLegacyWafRuleVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegacyWafRuleVcl`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleAPI.GetLegacyWafRuleVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyWafRuleVclRequest struct via the builder pattern


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


## ListLegacyWafRules

List rules in the latest configuration set



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
    filterRuleID := "filterRuleId_example" // string | Limit the returned rules to a specific rule ID. (optional)
    filterSeverity := "filterSeverity_example" // string | Limit the returned rules to a specific severity. (optional)
    filterTagsName := "filterTagsName_example" // string | Limit the returned rules to a set linked to a tag by name. (optional)
    filterConfigurationSetID := "filterConfigurationSetId_example" // string | Optional. Limit rules to specific configuration set or pass \"all\" to search all configuration sets, including stale ones. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "include_example" // string | Include relationships. Optional. Comma separated values. Permitted values: `tags`, `rule_statuses`, and `source`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleAPI.ListLegacyWafRules(ctx).FilterRuleID(filterRuleID).FilterSeverity(filterSeverity).FilterTagsName(filterTagsName).FilterConfigurationSetID(filterConfigurationSetID).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleAPI.ListLegacyWafRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWafRules`: []map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleAPI.ListLegacyWafRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWafRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRuleID** | **string** | Limit the returned rules to a specific rule ID. |  **filterSeverity** | **string** | Limit the returned rules to a specific severity. |  **filterTagsName** | **string** | Limit the returned rules to a set linked to a tag by name. |  **filterConfigurationSetID** | **string** | Optional. Limit rules to specific configuration set or pass \&quot;all\&quot; to search all configuration sets, including stale ones. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional. Comma separated values. Permitted values: `tags`, `rule_statuses`, and `source`.  | 

### Return type

**[]map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
