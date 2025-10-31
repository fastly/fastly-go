# BillingInvoicesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceByInvoiceId**](BillingInvoicesAPI.md#GetInvoiceByInvoiceId) | **GET** `/billing/v3/invoices/{invoice_id}` | Get invoice by ID.
[**GetMonthToDateInvoice**](BillingInvoicesAPI.md#GetMonthToDateInvoice) | **GET** `/billing/v3/invoices/month-to-date` | Get month-to-date invoice.
[**ListInvoices**](BillingInvoicesAPI.md#ListInvoices) | **GET** `/billing/v3/invoices` | List of invoices.



## GetInvoiceByInvoiceId

Get invoice by ID.



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
    invoiceId := int32(4183280) // int32 | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingInvoicesAPI.GetInvoiceByInvoiceId(ctx, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInvoicesAPI.GetInvoiceByInvoiceId`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceByInvoiceId`: EomInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingInvoicesAPI.GetInvoiceByInvoiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceByInvoiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EomInvoiceResponse**](EomInvoiceResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetMonthToDateInvoice

Get month-to-date invoice.



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingInvoicesAPI.GetMonthToDateInvoice(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInvoicesAPI.GetMonthToDateInvoice`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonthToDateInvoice`: MtdInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingInvoicesAPI.GetMonthToDateInvoice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonthToDateInvoiceRequest struct via the builder pattern



### Return type

[**MtdInvoiceResponse**](MtdInvoiceResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListInvoices

List of invoices.



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
    billingStartDate := "2023-01-01" // string |  (optional)
    billingEndDate := "2023-01-31" // string |  (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 200. (optional) (default to "100")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingInvoicesAPI.ListInvoices(ctx).BillingStartDate(billingStartDate).BillingEndDate(billingEndDate).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingInvoicesAPI.ListInvoices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: ListEomInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingInvoicesAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingStartDate** | **string** |  |  **billingEndDate** | **string** |  |  **limit** | **string** | Number of results per page. The maximum is 200. | [default to &quot;100&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. | 

### Return type

[**ListEomInvoicesResponse**](ListEomInvoicesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

