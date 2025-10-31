# ContactAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContacts**](ContactAPI.md#CreateContacts) | **POST** `/customer/{customer_id}/contacts` | Add a new customer contact
[**DeleteContact**](ContactAPI.md#DeleteContact) | **DELETE** `/customer/{customer_id}/contacts/{contact_id}` | Delete a contact
[**ListContacts**](ContactAPI.md#ListContacts) | **GET** `/customer/{customer_id}/contacts` | List contacts



## CreateContacts

Add a new customer contact



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
    userId := "userId_example" // string | The alphanumeric string representing the user for this customer contact. (optional)
    contactType := "contactType_example" // string | The type of contact. (optional)
    name := "name_example" // string | The name of this contact, when user_id is not provided. (optional)
    email := "email_example" // string | The email of this contact, when a user_id is not provided. (optional)
    phone := "phone_example" // string | The phone number for this contact. Required for primary, technical, and security contact types. (optional)
    customerId2 := "customerId_example" // string | The alphanumeric string representing the customer for this customer contact. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ContactAPI.CreateContacts(ctx, customerId).UserId(userId).ContactType(contactType).Name(name).Email(email).Phone(phone).CustomerId2(customerId2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.CreateContacts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContacts`: ContactResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactAPI.CreateContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The alphanumeric string representing the user for this customer contact. |  **contactType** | **string** | The type of contact. |  **name** | **string** | The name of this contact, when user_id is not provided. |  **email** | **string** | The email of this contact, when a user_id is not provided. |  **phone** | **string** | The phone number for this contact. Required for primary, technical, and security contact types. |  **customerId2** | **string** | The alphanumeric string representing the customer for this customer contact. | 

### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteContact

Delete a contact



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
    contactId := "contactId_example" // string | An alphanumeric string identifying the customer contact.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ContactAPI.DeleteContact(ctx, customerId, contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.DeleteContact`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContact`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ContactAPI.DeleteContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 
**contactId** | **string** | An alphanumeric string identifying the customer contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


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


## ListContacts

List contacts



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
    resp, r, err := apiClient.ContactAPI.ListContacts(ctx, customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.ListContacts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContacts`: []SchemasContactResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactAPI.ListContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SchemasContactResponse**](SchemasContactResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

