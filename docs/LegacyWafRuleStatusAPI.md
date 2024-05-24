# LegacyWafRuleStatusAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWafFirewallRuleStatus**](LegacyWafRuleStatusAPI.md#GetWafFirewallRuleStatus) | **GET** `/service/{service_id}/wafs/{firewall_id}/rules/{waf_rule_id}/rule_status` | Get the status of a rule on a firewall
[**ListWafFirewallRuleStatuses**](LegacyWafRuleStatusAPI.md#ListWafFirewallRuleStatuses) | **GET** `/service/{service_id}/wafs/{firewall_id}/rule_statuses` | List rule statuses
[**UpdateWafFirewallRuleStatus**](LegacyWafRuleStatusAPI.md#UpdateWafFirewallRuleStatus) | **PATCH** `/service/{service_id}/wafs/{firewall_id}/rules/{waf_rule_id}/rule_status` | Update the status of a rule
[**UpdateWafFirewallRuleStatusesTag**](LegacyWafRuleStatusAPI.md#UpdateWafFirewallRuleStatusesTag) | **POST** `/service/{service_id}/wafs/{firewall_id}/rule_statuses` | Create or update status of a tagged group of rules



## GetWafFirewallRuleStatus

Get the status of a rule on a firewall



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
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleStatusAPI.GetWafFirewallRuleStatus(ctx, serviceID, firewallID, wafRuleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleStatusAPI.GetWafFirewallRuleStatus`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafFirewallRuleStatus`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleStatusAPI.GetWafFirewallRuleStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafFirewallRuleStatusRequest struct via the builder pattern


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


## ListWafFirewallRuleStatuses

List rule statuses



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
    filterStatus := "filterStatus_example" // string | Limit results to rule statuses with the specified status. (optional)
    filterRuleMessage := "filterRuleMessage_example" // string | Limit results to rule statuses whose rules have the specified message. (optional)
    filterRuleRuleID := "filterRuleRuleId_example" // string | Limit results to rule statuses whose rules represent the specified ModSecurity rule_id. (optional)
    filterRuleTags := "filterRuleTags_example" // string | Limit results to rule statuses whose rules relate to the specified tag IDs. (optional)
    filterRuleTagsName := "application-FBC Market" // string | Limit results to rule statuses whose rules related to the named tags. (optional)
    include := "include_example" // string | Include relationships. Optional, comma separated values. Permitted values: `tags`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleStatusAPI.ListWafFirewallRuleStatuses(ctx, serviceID, firewallID).FilterStatus(filterStatus).FilterRuleMessage(filterRuleMessage).FilterRuleRuleID(filterRuleRuleID).FilterRuleTags(filterRuleTags).FilterRuleTagsName(filterRuleTagsName).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleStatusAPI.ListWafFirewallRuleStatuses`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafFirewallRuleStatuses`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleStatusAPI.ListWafFirewallRuleStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafFirewallRuleStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Limit results to rule statuses with the specified status. |  **filterRuleMessage** | **string** | Limit results to rule statuses whose rules have the specified message. |  **filterRuleRuleID** | **string** | Limit results to rule statuses whose rules represent the specified ModSecurity rule_id. |  **filterRuleTags** | **string** | Limit results to rule statuses whose rules relate to the specified tag IDs. |  **filterRuleTagsName** | **string** | Limit results to rule statuses whose rules related to the named tags. |  **include** | **string** | Include relationships. Optional, comma separated values. Permitted values: `tags`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateWafFirewallRuleStatus

Update the status of a rule



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
    wafRuleID := "wafRuleId_example" // string | Alphanumeric string identifying a WAF rule.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleStatusAPI.UpdateWafFirewallRuleStatus(ctx, serviceID, firewallID, wafRuleID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleStatusAPI.UpdateWafFirewallRuleStatus`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafFirewallRuleStatus`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleStatusAPI.UpdateWafFirewallRuleStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafFirewallRuleStatusRequest struct via the builder pattern


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


## UpdateWafFirewallRuleStatusesTag

Create or update status of a tagged group of rules



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
    name := "name_example" // string | The tag name to use to determine the set of rules to update. For example, OWASP or language-php. (optional)
    force := "force_example" // string | Whether or not to update rule statuses for disabled rules. Optional. (optional)
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafRuleStatusAPI.UpdateWafFirewallRuleStatusesTag(ctx, serviceID, firewallID).Name(name).Force(force).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafRuleStatusAPI.UpdateWafFirewallRuleStatusesTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafFirewallRuleStatusesTag`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafRuleStatusAPI.UpdateWafFirewallRuleStatusesTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafFirewallRuleStatusesTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The tag name to use to determine the set of rules to update. For example, OWASP or language-php. |  **force** | **string** | Whether or not to update rule statuses for disabled rules. Optional. |  **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
