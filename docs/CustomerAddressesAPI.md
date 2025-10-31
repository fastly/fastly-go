# CustomerAddressesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerAddress**](CustomerAddressesAPI.md#CreateCustomerAddress) | **POST** `/billing/v3/customer-addresses` | Creates an address associated with a customer account.
[**ListCustomerAddresses**](CustomerAddressesAPI.md#ListCustomerAddresses) | **GET** `/billing/v3/customer-addresses` | Return the list of addresses associated with a customer account.
[**UpdateCustomerAddress**](CustomerAddressesAPI.md#UpdateCustomerAddress) | **PUT** `/billing/v3/customer-addresses/{type}` | Updates an address associated with a customer account.



## CreateCustomerAddress

Creates an address associated with a customer account.



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
    customerAddress := *openapiclient.NewCustomerAddress() // CustomerAddress | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CustomerAddressesAPI.CreateCustomerAddress(ctx).CustomerAddress(customerAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesAPI.CreateCustomerAddress`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerAddress`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `CustomerAddressesAPI.CreateCustomerAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerAddress** | [**CustomerAddress**](CustomerAddress.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListCustomerAddresses

Return the list of addresses associated with a customer account.



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
    resp, r, err := apiClient.CustomerAddressesAPI.ListCustomerAddresses(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesAPI.ListCustomerAddresses`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomerAddresses`: ListCustomerAddressesResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerAddressesAPI.ListCustomerAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerAddressesRequest struct via the builder pattern



### Return type

[**ListCustomerAddressesResponse**](ListCustomerAddressesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateCustomerAddress

Updates an address associated with a customer account.



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
    type_ := "type__example" // string | Alphanumeric type of the address being modified.
    customerAddress := *openapiclient.NewCustomerAddress() // CustomerAddress | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CustomerAddressesAPI.UpdateCustomerAddress(ctx, type_).CustomerAddress(customerAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAddressesAPI.UpdateCustomerAddress`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Alphanumeric type of the address being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerAddress** | [**CustomerAddress**](CustomerAddress.md) |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

