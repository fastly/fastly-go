# CustomerAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomer**](CustomerAPI.md#DeleteCustomer) | **DELETE** `/customer/{customer_id}` | Delete a customer
[**GetCustomer**](CustomerAPI.md#GetCustomer) | **GET** `/customer/{customer_id}` | Get a customer
[**GetLoggedInCustomer**](CustomerAPI.md#GetLoggedInCustomer) | **GET** `/current_customer` | Get the logged in customer
[**ListUsers**](CustomerAPI.md#ListUsers) | **GET** `/customer/{customer_id}/users` | List users
[**UpdateCustomer**](CustomerAPI.md#UpdateCustomer) | **PUT** `/customer/{customer_id}` | Update a customer



## DeleteCustomer

Delete a customer



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
    customerId := "customerId_example" // string | Alphanumeric string identifying the customer.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CustomerAPI.DeleteCustomer(ctx, customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.DeleteCustomer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomer`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.DeleteCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetCustomer

Get a customer



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
    customerId := "customerId_example" // string | Alphanumeric string identifying the customer.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CustomerAPI.GetCustomer(ctx, customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetCustomer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: CustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetLoggedInCustomer

Get the logged in customer



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
    resp, r, err := apiClient.CustomerAPI.GetLoggedInCustomer(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetLoggedInCustomer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoggedInCustomer`: CustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetLoggedInCustomer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInCustomerRequest struct via the builder pattern



### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListUsers

List users



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
    customerId := "customerId_example" // string | Alphanumeric string identifying the customer.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CustomerAPI.ListUsers(ctx, customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.ListUsers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []SchemasUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SchemasUserResponse**](SchemasUserResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateCustomer

Update a customer



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
    customerId := "customerId_example" // string | Alphanumeric string identifying the customer.
    billingContactId := "billingContactId_example" // string | The alphanumeric string representing the primary billing contact. (optional)
    billingNetworkType := "billingNetworkType_example" // string | Customer's current network revenue type. (optional)
    billingRef := "billingRef_example" // string | Used for adding purchased orders to customer's account. (optional)
    canConfigureWordpress := true // bool | Whether this customer can view or edit wordpress. (optional)
    canResetPasswords := true // bool | Whether this customer can reset passwords. (optional)
    canUploadVcl := true // bool | Whether this customer can upload VCL. (optional)
    force2fa := true // bool | Specifies whether 2FA is forced or not forced on the customer account. Logs out non-2FA users once 2FA is force enabled. (optional)
    forceSso := true // bool | Specifies whether SSO is forced or not forced on the customer account. (optional)
    hasAccountPanel := true // bool | Specifies whether the account has access or does not have access to the account panel. (optional)
    hasImprovedEvents := true // bool | Specifies whether the account has access or does not have access to the improved events. (optional)
    hasImprovedSslConfig := true // bool | Whether this customer can view or edit the SSL config. (optional)
    hasOpenstackLogging := true // bool | Specifies whether the account has enabled or not enabled openstack logging. (optional)
    hasPci := true // bool | Specifies whether the account can edit PCI for a service. (optional)
    hasPciPasswords := true // bool | Specifies whether PCI passwords are required for the account. (optional)
    ipWhitelist := "ipWhitelist_example" // string | The range of IP addresses authorized to access the customer account. (optional)
    legalContactId := "legalContactId_example" // string | The alphanumeric string identifying the account's legal contact. (optional)
    name := "name_example" // string | The name of the customer, generally the company name. (optional)
    ownerId := "ownerId_example" // string | The alphanumeric string identifying the account owner. (optional)
    phoneNumber := "phoneNumber_example" // string | The phone number associated with the account. (optional)
    postalAddress := "postalAddress_example" // string | The postal address associated with the account. (optional)
    pricingPlan := "pricingPlan_example" // string | The pricing plan this customer is under. (optional)
    pricingPlanId := "pricingPlanId_example" // string | The alphanumeric string identifying the pricing plan. (optional)
    securityContactId := "securityContactId_example" // string | The alphanumeric string identifying the account's security contact. (optional)
    technicalContactId := "technicalContactId_example" // string | The alphanumeric string identifying the account's technical contact. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.CustomerAPI.UpdateCustomer(ctx, customerId).BillingContactId(billingContactId).BillingNetworkType(billingNetworkType).BillingRef(billingRef).CanConfigureWordpress(canConfigureWordpress).CanResetPasswords(canResetPasswords).CanUploadVcl(canUploadVcl).Force2fa(force2fa).ForceSso(forceSso).HasAccountPanel(hasAccountPanel).HasImprovedEvents(hasImprovedEvents).HasImprovedSslConfig(hasImprovedSslConfig).HasOpenstackLogging(hasOpenstackLogging).HasPci(hasPci).HasPciPasswords(hasPciPasswords).IpWhitelist(ipWhitelist).LegalContactId(legalContactId).Name(name).OwnerId(ownerId).PhoneNumber(phoneNumber).PostalAddress(postalAddress).PricingPlan(pricingPlan).PricingPlanId(pricingPlanId).SecurityContactId(securityContactId).TechnicalContactId(technicalContactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.UpdateCustomer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: CustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingContactId** | **string** | The alphanumeric string representing the primary billing contact. |  **billingNetworkType** | **string** | Customer&#39;s current network revenue type. |  **billingRef** | **string** | Used for adding purchased orders to customer&#39;s account. |  **canConfigureWordpress** | **bool** | Whether this customer can view or edit wordpress. |  **canResetPasswords** | **bool** | Whether this customer can reset passwords. |  **canUploadVcl** | **bool** | Whether this customer can upload VCL. |  **force2fa** | **bool** | Specifies whether 2FA is forced or not forced on the customer account. Logs out non-2FA users once 2FA is force enabled. |  **forceSso** | **bool** | Specifies whether SSO is forced or not forced on the customer account. |  **hasAccountPanel** | **bool** | Specifies whether the account has access or does not have access to the account panel. |  **hasImprovedEvents** | **bool** | Specifies whether the account has access or does not have access to the improved events. |  **hasImprovedSslConfig** | **bool** | Whether this customer can view or edit the SSL config. |  **hasOpenstackLogging** | **bool** | Specifies whether the account has enabled or not enabled openstack logging. |  **hasPci** | **bool** | Specifies whether the account can edit PCI for a service. |  **hasPciPasswords** | **bool** | Specifies whether PCI passwords are required for the account. |  **ipWhitelist** | **string** | The range of IP addresses authorized to access the customer account. |  **legalContactId** | **string** | The alphanumeric string identifying the account&#39;s legal contact. |  **name** | **string** | The name of the customer, generally the company name. |  **ownerId** | **string** | The alphanumeric string identifying the account owner. |  **phoneNumber** | **string** | The phone number associated with the account. |  **postalAddress** | **string** | The postal address associated with the account. |  **pricingPlan** | **string** | The pricing plan this customer is under. |  **pricingPlanId** | **string** | The alphanumeric string identifying the pricing plan. |  **securityContactId** | **string** | The alphanumeric string identifying the account&#39;s security contact. |  **technicalContactId** | **string** | The alphanumeric string identifying the account&#39;s technical contact. | 

### Return type

[**CustomerResponse**](CustomerResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

