# BillingAddressAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBillingAddr**](BillingAddressAPI.md#AddBillingAddr) | **POST** `/customer/{customer_id}/billing_address` | Add a billing address to a customer
[**DeleteBillingAddr**](BillingAddressAPI.md#DeleteBillingAddr) | **DELETE** `/customer/{customer_id}/billing_address` | Delete a billing address
[**GetBillingAddr**](BillingAddressAPI.md#GetBillingAddr) | **GET** `/customer/{customer_id}/billing_address` | Get a billing address
[**UpdateBillingAddr**](BillingAddressAPI.md#UpdateBillingAddr) | **PATCH** `/customer/{customer_id}/billing_address` | Update a billing address



## AddBillingAddr

Add a billing address to a customer



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
    billingAddressRequest := *openapiclient.NewBillingAddressRequest() // BillingAddressRequest | Billing address (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAddressAPI.AddBillingAddr(ctx, customerID).BillingAddressRequest(billingAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAddressAPI.AddBillingAddr`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBillingAddr`: BillingAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingAddressAPI.AddBillingAddr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBillingAddrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingAddressRequest** | [**BillingAddressRequest**](BillingAddressRequest.md) | Billing address | 

### Return type

[**BillingAddressResponse**](BillingAddressResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteBillingAddr

Delete a billing address



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAddressAPI.DeleteBillingAddr(ctx, customerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAddressAPI.DeleteBillingAddr`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingAddrRequest struct via the builder pattern


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


## GetBillingAddr

Get a billing address



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAddressAPI.GetBillingAddr(ctx, customerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAddressAPI.GetBillingAddr`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingAddr`: BillingAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingAddressAPI.GetBillingAddr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAddrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingAddressResponse**](BillingAddressResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateBillingAddr

Update a billing address



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
    updateBillingAddressRequest := *openapiclient.NewUpdateBillingAddressRequest() // UpdateBillingAddressRequest | One or more billing address attributes (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BillingAddressAPI.UpdateBillingAddr(ctx, customerID).UpdateBillingAddressRequest(updateBillingAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingAddressAPI.UpdateBillingAddr`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBillingAddr`: BillingAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingAddressAPI.UpdateBillingAddr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillingAddrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateBillingAddressRequest** | [**UpdateBillingAddressRequest**](UpdateBillingAddressRequest.md) | One or more billing address attributes | 

### Return type

[**BillingAddressResponse**](BillingAddressResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
