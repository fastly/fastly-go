# WafActiveRulesAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateWafActiveRules**](WafActiveRulesAPI.md#BulkUpdateWafActiveRules) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/bulk` | Update multiple active rules
[**CreateWafActiveRule**](WafActiveRulesAPI.md#CreateWafActiveRule) | **POST** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules` | Add a rule to a WAF as an active rule
[**CreateWafActiveRulesTag**](WafActiveRulesAPI.md#CreateWafActiveRulesTag) | **POST** `/waf/firewalls/{firewall_id}/versions/{version_id}/tags/{waf_tag_name}/active-rules` | Create active rules by tag
[**DeleteWafActiveRule**](WafActiveRulesAPI.md#DeleteWafActiveRule) | **DELETE** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}` | Delete an active rule
[**GetWafActiveRule**](WafActiveRulesAPI.md#GetWafActiveRule) | **GET** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}` | Get an active WAF rule object
[**ListWafActiveRules**](WafActiveRulesAPI.md#ListWafActiveRules) | **GET** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules` | List active rules on a WAF
[**UpdateWafActiveRule**](WafActiveRulesAPI.md#UpdateWafActiveRule) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}` | Update an active rule



## BulkUpdateWafActiveRules

Update multiple active rules



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    body := WafActiveRuleData({"type":"waf_active_rule","attributes":{"revision":"latest"}}) // WafActiveRuleData |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.BulkUpdateWafActiveRules(ctx, firewallID, versionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.BulkUpdateWafActiveRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateWafActiveRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **WafActiveRuleData** |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateWafActiveRule

Add a rule to a WAF as an active rule



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    wafActiveRule := *openapiclient.NewWafActiveRule() // WafActiveRule |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.CreateWafActiveRule(ctx, firewallID, versionID).WafActiveRule(wafActiveRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.CreateWafActiveRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWafActiveRule`: WafActiveRuleCreationResponse
    fmt.Fprintf(os.Stdout, "Response from `WafActiveRulesAPI.CreateWafActiveRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafActiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafActiveRule** | [**WafActiveRule**](WafActiveRule.md) |  | 

### Return type

[**WafActiveRuleCreationResponse**](WafActiveRuleCreationResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/vnd.api+json; ext=bulk
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateWafActiveRulesTag

Create active rules by tag



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    wafTagName := "wafTagName_example" // string | Name of the tag.
    wafActiveRule := *openapiclient.NewWafActiveRule() // WafActiveRule |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.CreateWafActiveRulesTag(ctx, firewallID, versionID, wafTagName).WafActiveRule(wafActiveRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.CreateWafActiveRulesTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 
**wafTagName** | **string** | Name of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafActiveRulesTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafActiveRule** | [**WafActiveRule**](WafActiveRule.md) |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteWafActiveRule

Delete an active rule



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.DeleteWafActiveRule(ctx, firewallID, versionID, wafRuleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.DeleteWafActiveRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWafActiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetWafActiveRule

Get an active WAF rule object



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.
    include := "waf_rule_revision,waf_firewall_version" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule_revision` and `waf_firewall_version`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.GetWafActiveRule(ctx, firewallID, versionID, wafRuleID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.GetWafActiveRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafActiveRule`: WafActiveRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `WafActiveRulesAPI.GetWafActiveRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafActiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule_revision` and `waf_firewall_version`.  | 

### Return type

[**WafActiveRuleResponse**](WafActiveRuleResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafActiveRules

List active rules on a WAF



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    filterStatus := "filterStatus_example" // string | Limit results to active rules with the specified status. (optional)
    filterWafRuleRevisionMessage := "filterWafRuleRevisionMessage_example" // string | Limit results to active rules with the specified message. (optional)
    filterWafRuleRevisionModsecRuleID := "filterWafRuleRevisionModsecRuleId_example" // string | Limit results to active rules that represent the specified ModSecurity modsec_rule_id. (optional)
    filterOutdated := "filterOutdated_example" // string | Limit results to active rules referencing an outdated rule revision. (optional)
    include := "waf_rule_revision,waf_firewall_version" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule_revision` and `waf_firewall_version`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.ListWafActiveRules(ctx, firewallID, versionID).FilterStatus(filterStatus).FilterWafRuleRevisionMessage(filterWafRuleRevisionMessage).FilterWafRuleRevisionModsecRuleID(filterWafRuleRevisionModsecRuleID).FilterOutdated(filterOutdated).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.ListWafActiveRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafActiveRules`: WafActiveRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `WafActiveRulesAPI.ListWafActiveRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafActiveRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Limit results to active rules with the specified status. |  **filterWafRuleRevisionMessage** | **string** | Limit results to active rules with the specified message. |  **filterWafRuleRevisionModsecRuleID** | **string** | Limit results to active rules that represent the specified ModSecurity modsec_rule_id. |  **filterOutdated** | **string** | Limit results to active rules referencing an outdated rule revision. |  **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule_revision` and `waf_firewall_version`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**WafActiveRulesResponse**](WafActiveRulesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateWafActiveRule

Update an active rule



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.
    wafActiveRule := *openapiclient.NewWafActiveRule() // WafActiveRule |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafActiveRulesAPI.UpdateWafActiveRule(ctx, firewallID, versionID, wafRuleID).WafActiveRule(wafActiveRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafActiveRulesAPI.UpdateWafActiveRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafActiveRule`: WafActiveRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `WafActiveRulesAPI.UpdateWafActiveRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**versionID** | **int32** | Integer identifying a service version. | 
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafActiveRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafActiveRule** | [**WafActiveRule**](WafActiveRule.md) |  | 

### Return type

[**WafActiveRuleResponse**](WafActiveRuleResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
