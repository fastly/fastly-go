# WafRulesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWafRule**](WafRulesAPI.md#GetWafRule) | **GET** `/waf/rules/{waf_rule_id}` | Get a rule
[**ListWafRules**](WafRulesAPI.md#ListWafRules) | **GET** `/waf/rules` | List available WAF rules



## GetWafRule

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
    include := "waf_tags,waf_rule_revisions" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_tags` and `waf_rule_revisions`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafRulesAPI.GetWafRule(ctx, wafRuleID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafRulesAPI.GetWafRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafRule`: WafRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `WafRulesAPI.GetWafRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_tags` and `waf_rule_revisions`.  | 

### Return type

[**WafRuleResponse**](WafRuleResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafRules

List available WAF rules



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
    filterModsecRuleID := "filterModsecRuleId_example" // string | Limit the returned rules to a specific ModSecurity rule ID. (optional)
    filterWafTagsName := "filterWafTagsName_example" // string | Limit the returned rules to a set linked to a tag by name. (optional)
    filterWafRuleRevisionsSource := "filterWafRuleRevisionsSource_example" // string | Limit the returned rules to a set linked to a source. (optional)
    filterWafFirewallIDNotMatch := "filterWafFirewallIDNotMatch_example" // string | Limit the returned rules to a set not included in the active firewall version for a firewall. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "waf_tags,waf_rule_revisions" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_tags` and `waf_rule_revisions`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafRulesAPI.ListWafRules(ctx).FilterModsecRuleID(filterModsecRuleID).FilterWafTagsName(filterWafTagsName).FilterWafRuleRevisionsSource(filterWafRuleRevisionsSource).FilterWafFirewallIDNotMatch(filterWafFirewallIDNotMatch).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafRulesAPI.ListWafRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafRules`: WafRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `WafRulesAPI.ListWafRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWafRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterModsecRuleID** | **string** | Limit the returned rules to a specific ModSecurity rule ID. |  **filterWafTagsName** | **string** | Limit the returned rules to a set linked to a tag by name. |  **filterWafRuleRevisionsSource** | **string** | Limit the returned rules to a set linked to a source. |  **filterWafFirewallIDNotMatch** | **string** | Limit the returned rules to a set not included in the active firewall version for a firewall. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_tags` and `waf_rule_revisions`.  | 

### Return type

[**WafRulesResponse**](WafRulesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
