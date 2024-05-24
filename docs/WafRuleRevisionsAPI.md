# WafRuleRevisionsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWafRuleRevision**](WafRuleRevisionsAPI.md#GetWafRuleRevision) | **GET** `/waf/rules/{waf_rule_id}/revisions/{waf_rule_revision_number}` | Get a revision of a rule
[**ListWafRuleRevisions**](WafRuleRevisionsAPI.md#ListWafRuleRevisions) | **GET** `/waf/rules/{waf_rule_id}/revisions` | List revisions for a rule



## GetWafRuleRevision

Get a revision of a rule



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
    wafRuleRevisionNumber := int32(56) // int32 | Revision number.
    include := "source,vcl,waf_rule" // string | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule`, `vcl`, and `source`. The `vcl` and `source` relationships show the WAF VCL and corresponding ModSecurity source. These fields are blank unless the relationship is included.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafRuleRevisionsAPI.GetWafRuleRevision(ctx, wafRuleID, wafRuleRevisionNumber).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafRuleRevisionsAPI.GetWafRuleRevision`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafRuleRevision`: WafRuleRevisionResponse
    fmt.Fprintf(os.Stdout, "Response from `WafRuleRevisionsAPI.GetWafRuleRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 
**wafRuleRevisionNumber** | **int32** | Revision number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafRuleRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include relationships. Optional, comma-separated values. Permitted values: `waf_rule`, `vcl`, and `source`. The `vcl` and `source` relationships show the WAF VCL and corresponding ModSecurity source. These fields are blank unless the relationship is included.  | 

### Return type

[**WafRuleRevisionResponse**](WafRuleRevisionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafRuleRevisions

List revisions for a rule



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
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "waf_rule" // string | Include relationships. Optional. (optional) (default to "waf_rule")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafRuleRevisionsAPI.ListWafRuleRevisions(ctx, wafRuleID).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafRuleRevisionsAPI.ListWafRuleRevisions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafRuleRevisions`: WafRuleRevisionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WafRuleRevisionsAPI.ListWafRuleRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleID** | **string** | Alphanumeric string identifying a WAF rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafRuleRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional. | [default to &quot;waf_rule&quot;]

### Return type

[**WafRuleRevisionsResponse**](WafRuleRevisionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
