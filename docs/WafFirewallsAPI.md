# WafFirewallsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWafFirewall**](WafFirewallsAPI.md#CreateWafFirewall) | **POST** `/waf/firewalls` | Create a firewall
[**DeleteWafFirewall**](WafFirewallsAPI.md#DeleteWafFirewall) | **DELETE** `/waf/firewalls/{firewall_id}` | Delete a firewall
[**GetWafFirewall**](WafFirewallsAPI.md#GetWafFirewall) | **GET** `/waf/firewalls/{firewall_id}` | Get a firewall
[**ListWafFirewalls**](WafFirewallsAPI.md#ListWafFirewalls) | **GET** `/waf/firewalls` | List firewalls
[**UpdateWafFirewall**](WafFirewallsAPI.md#UpdateWafFirewall) | **PATCH** `/waf/firewalls/{firewall_id}` | Update a firewall



## CreateWafFirewall

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
    wafFirewall := *openapiclient.NewWafFirewall() // WafFirewall |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallsAPI.CreateWafFirewall(ctx).WafFirewall(wafFirewall).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallsAPI.CreateWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWafFirewall`: WafFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallsAPI.CreateWafFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafFirewall** | [**WafFirewall**](WafFirewall.md) |  | 

### Return type

[**WafFirewallResponse**](WafFirewallResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteWafFirewall

Delete a firewall



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
    wafFirewall := *openapiclient.NewWafFirewall() // WafFirewall |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallsAPI.DeleteWafFirewall(ctx, firewallID).WafFirewall(wafFirewall).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallsAPI.DeleteWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafFirewall** | [**WafFirewall**](WafFirewall.md) |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetWafFirewall

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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    filterServiceVersionNumber := "filterServiceVersionNumber_example" // string | Limit the results returned to a specific service version. (optional)
    include := "include_example" // string | Include related objects. Optional. (optional) (default to "waf_firewall_versions")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallsAPI.GetWafFirewall(ctx, firewallID).FilterServiceVersionNumber(filterServiceVersionNumber).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallsAPI.GetWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafFirewall`: WafFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallsAPI.GetWafFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterServiceVersionNumber** | **string** | Limit the results returned to a specific service version. |  **include** | **string** | Include related objects. Optional. | [default to &quot;waf_firewall_versions&quot;]

### Return type

[**WafFirewallResponse**](WafFirewallResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafFirewalls

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
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    filterServiceID := "filterServiceId_example" // string | Limit the results returned to a specific service. (optional)
    filterServiceVersionNumber := "filterServiceVersionNumber_example" // string | Limit the results returned to a specific service version. (optional)
    include := "include_example" // string | Include related objects. Optional. (optional) (default to "waf_firewall_versions")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallsAPI.ListWafFirewalls(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterServiceID(filterServiceID).FilterServiceVersionNumber(filterServiceVersionNumber).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallsAPI.ListWafFirewalls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafFirewalls`: WafFirewallsResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallsAPI.ListWafFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWafFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **filterServiceID** | **string** | Limit the results returned to a specific service. |  **filterServiceVersionNumber** | **string** | Limit the results returned to a specific service version. |  **include** | **string** | Include related objects. Optional. | [default to &quot;waf_firewall_versions&quot;]

### Return type

[**WafFirewallsResponse**](WafFirewallsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateWafFirewall

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
    firewallID := "firewallId_example" // string | Alphanumeric string identifying a WAF Firewall.
    wafFirewall := *openapiclient.NewWafFirewall() // WafFirewall |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.WafFirewallsAPI.UpdateWafFirewall(ctx, firewallID).WafFirewall(wafFirewall).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WafFirewallsAPI.UpdateWafFirewall`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWafFirewall`: WafFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `WafFirewallsAPI.UpdateWafFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallID** | **string** | Alphanumeric string identifying a WAF Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafFirewall** | [**WafFirewall**](WafFirewall.md) |  | 

### Return type

[**WafFirewallResponse**](WafFirewallResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
