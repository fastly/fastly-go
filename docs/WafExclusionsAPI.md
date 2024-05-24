# WafExclusionsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWafRuleExclusion**](WafExclusionsAPI.md#CreateWafRuleExclusion) | **POST** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions` | Create a WAF rule exclusion
[**DeleteWafRuleExclusion**](WafExclusionsAPI.md#DeleteWafRuleExclusion) | **DELETE** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}` | Delete a WAF rule exclusion
[**GetWafRuleExclusion**](WafExclusionsAPI.md#GetWafRuleExclusion) | **GET** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}` | Get a WAF rule exclusion
[**ListWafRuleExclusions**](WafExclusionsAPI.md#ListWafRuleExclusions) | **GET** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions` | List WAF rule exclusions
[**UpdateWafRuleExclusion**](WafExclusionsAPI.md#UpdateWafRuleExclusion) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}` | Update a WAF rule exclusion



## CreateWafRuleExclusion

Create a WAF rule exclusion



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
    wafExclusion := *openapiclient.NewWafExclusion() // WafExclusion |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafExclusionsAPI.CreateWafRuleExclusion(ctx, firewallID, firewallVersionNumber).WafExclusion(wafExclusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafExclusionsAPI.CreateWafRuleExclusion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWafRuleExclusion`: WafExclusionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafExclusionsAPI.CreateWafRuleExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafRuleExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafExclusion** | [**WafExclusion**](WafExclusion.md) |  | 

### Return type

[**WafExclusionResponse**](WafExclusionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteWafRuleExclusion

Delete a WAF rule exclusion



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
    exclusionNumber := int32(56) // int32 | A numeric ID identifying a WAF exclusion.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafExclusionsAPI.DeleteWafRuleExclusion(ctx, firewallID, firewallVersionNumber, exclusionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafExclusionsAPI.DeleteWafRuleExclusion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 
**exclusionNumber** | **int32** | A numeric ID identifying a WAF exclusion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWafRuleExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetWafRuleExclusion

Get a WAF rule exclusion



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
    exclusionNumber := int32(56) // int32 | A numeric ID identifying a WAF exclusion.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafExclusionsAPI.GetWafRuleExclusion(ctx, firewallID, firewallVersionNumber, exclusionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafExclusionsAPI.GetWafRuleExclusion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafRuleExclusion`: WafExclusionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafExclusionsAPI.GetWafRuleExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 
**exclusionNumber** | **int32** | A numeric ID identifying a WAF exclusion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafRuleExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WafExclusionResponse**](WafExclusionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafRuleExclusions

List WAF rule exclusions



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
    filterExclusionType := "filterExclusionType_example" // string | Filters the results based on this exclusion type. (optional)
    filterName := "filterName_example" // string | Filters the results based on name. (optional)
    filterWafRulesModsecRuleID := int32(56) // int32 | Filters the results based on this ModSecurity rule ID. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "waf_rules" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_rules` and `waf_rule_revisions`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafExclusionsAPI.ListWafRuleExclusions(ctx, firewallID, firewallVersionNumber).FilterExclusionType(filterExclusionType).FilterName(filterName).FilterWafRulesModsecRuleID(filterWafRulesModsecRuleID).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafExclusionsAPI.ListWafRuleExclusions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafRuleExclusions`: WafExclusionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WafExclusionsAPI.ListWafRuleExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafRuleExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterExclusionType** | **string** | Filters the results based on this exclusion type. |  **filterName** | **string** | Filters the results based on name. |  **filterWafRulesModsecRuleID** | **int32** | Filters the results based on this ModSecurity rule ID. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_rules` and `waf_rule_revisions`.  | 

### Return type

[**WafExclusionsResponse**](WafExclusionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateWafRuleExclusion

Update a WAF rule exclusion



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
    exclusionNumber := int32(56) // int32 | A numeric ID identifying a WAF exclusion.
    wafExclusion := *openapiclient.NewWafExclusion() // WafExclusion |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafExclusionsAPI.UpdateWafRuleExclusion(ctx, firewallID, firewallVersionNumber, exclusionNumber).WafExclusion(wafExclusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafExclusionsAPI.UpdateWafRuleExclusion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafRuleExclusion`: WafExclusionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafExclusionsAPI.UpdateWafRuleExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 
**firewallVersionNumber** | **int32** | Integer identifying a WAF firewall version. | 
**exclusionNumber** | **int32** | A numeric ID identifying a WAF exclusion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafRuleExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafExclusion** | [**WafExclusion**](WafExclusion.md) |  | 

### Return type

[**WafExclusionResponse**](WafExclusionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
