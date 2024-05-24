# LegacyWafUpdateStatusAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWafUpdateStatus**](LegacyWafUpdateStatusAPI.md#GetWafUpdateStatus) | **GET** `/service/{service_id}/wafs/{firewall_id}/update_statuses/{update_status_id}` | Get the status of a WAF update
[**ListWafUpdateStatuses**](LegacyWafUpdateStatusAPI.md#ListWafUpdateStatuses) | **GET** `/service/{service_id}/wafs/{firewall_id}/update_statuses` | List update statuses



## GetWafUpdateStatus

Get the status of a WAF update



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
    updateStatusID := "updateStatusId_example" // string | Alphanumeric string identifying a WAF update status.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafUpdateStatusAPI.GetWafUpdateStatus(ctx, serviceID, firewallID, updateStatusID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafUpdateStatusAPI.GetWafUpdateStatus`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWafUpdateStatus`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafUpdateStatusAPI.GetWafUpdateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 
**updateStatusID** | **string** | Alphanumeric string identifying a WAF update status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWafUpdateStatusRequest struct via the builder pattern


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


## ListWafUpdateStatuses

List update statuses



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
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "waf" // string | Include relationships. Optional, comma separated values. Permitted values: `waf`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafUpdateStatusAPI.ListWafUpdateStatuses(ctx, serviceID, firewallID).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafUpdateStatusAPI.ListWafUpdateStatuses`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafUpdateStatuses`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafUpdateStatusAPI.ListWafUpdateStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**firewallID** | **string** | Alphanumeric string identifying a Firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafUpdateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional, comma separated values. Permitted values: `waf`.  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
