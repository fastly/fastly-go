# BillingAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoice**](BillingAPI.md#GetInvoice) | **GET** `/billing/v2/year/{year}/month/{month}` | Get an invoice
[**GetInvoiceByID**](BillingAPI.md#GetInvoiceByID) | **GET** `/billing/v2/account_customers/{customer_id}/invoices/{invoice_id}` | Get an invoice
[**GetInvoiceMtd**](BillingAPI.md#GetInvoiceMtd) | **GET** `/billing/v2/account_customers/{customer_id}/mtd_invoice` | Get month-to-date billing estimate



## GetInvoice

Get an invoice



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
    month := "05" // string | 2-digit month.
    year := "2020" // string | 4-digit year.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAPI.GetInvoice(ctx, month, year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetInvoice`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: BillingResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** | 2-digit month. | 
**year** | **string** | 4-digit year. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingResponse**](BillingResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/pdf

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetInvoiceByID

Get an invoice



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
    customerID := "customerId_example" // string | Alphanumeric string identifying the customer.
    invoiceID := "invoiceId_example" // string | Alphanumeric string identifying the invoice.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAPI.GetInvoiceByID(ctx, customerID, invoiceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetInvoiceByID`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceByID`: BillingResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetInvoiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 
**invoiceID** | **string** | Alphanumeric string identifying the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingResponse**](BillingResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/pdf

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetInvoiceMtd

Get month-to-date billing estimate



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
    customerID := "customerId_example" // string | Alphanumeric string identifying the customer.
    month := "05" // string | 2-digit month. (optional)
    year := "2020" // string | 4-digit year. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAPI.GetInvoiceMtd(ctx, customerID).Month(month).Year(year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetInvoiceMtd`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceMtd`: BillingEstimateResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetInvoiceMtd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceMtdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | 2-digit month. |  **year** | **string** | 4-digit year. | 

### Return type

[**BillingEstimateResponse**](BillingEstimateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
