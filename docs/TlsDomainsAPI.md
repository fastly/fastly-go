# TLSDomainsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTLSDomains**](TlsDomainsAPI.md#ListTLSDomains) | **GET** `/tls/domains` | List TLS domains



## ListTLSDomains

List TLS domains



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
    filterInUse := "filterInUse_example" // string | Optional. Limit the returned domains to those currently using Fastly to terminate TLS with SNI (that is, domains considered \"in use\") Permitted values: true, false. (optional)
    filterTLSCertificatesID := "filterTLSCertificatesId_example" // string | Optional. Limit the returned domains to those listed in the given TLS certificate's SAN list. (optional)
    filterTLSSubscriptionsID := "filterTLSSubscriptionsId_example" // string | Optional. Limit the returned domains to those for a given TLS subscription. (optional)
    include := "include_example" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`, `tls_certificates`, `tls_subscriptions`, `tls_subscriptions.tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, and `tls_authorizations.self_managed_http_challenge`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSDomainsAPI.ListTLSDomains(ctx).FilterInUse(filterInUse).FilterTLSCertificatesID(filterTLSCertificatesID).FilterTLSSubscriptionsID(filterTLSSubscriptionsID).Include(include).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSDomainsAPI.ListTLSDomains`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSDomains`: TLSDomainsResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSDomainsAPI.ListTLSDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInUse** | **string** | Optional. Limit the returned domains to those currently using Fastly to terminate TLS with SNI (that is, domains considered \&quot;in use\&quot;) Permitted values: true, false. |  **filterTLSCertificatesID** | **string** | Optional. Limit the returned domains to those listed in the given TLS certificate&#39;s SAN list. |  **filterTLSSubscriptionsID** | **string** | Optional. Limit the returned domains to those for a given TLS subscription. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`, `tls_certificates`, `tls_subscriptions`, `tls_subscriptions.tls_authorizations`, `tls_authorizations.globalsign_email_challenge`, and `tls_authorizations.self_managed_http_challenge`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**TLSDomainsResponse**](TlsDomainsResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
