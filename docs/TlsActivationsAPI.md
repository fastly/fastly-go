# TLSActivationsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTLSActivation**](TlsActivationsAPI.md#CreateTLSActivation) | **POST** `/tls/activations` | Enable TLS for a domain using a custom certificate
[**DeleteTLSActivation**](TlsActivationsAPI.md#DeleteTLSActivation) | **DELETE** `/tls/activations/{tls_activation_id}` | Disable TLS on a domain
[**GetTLSActivation**](TlsActivationsAPI.md#GetTLSActivation) | **GET** `/tls/activations/{tls_activation_id}` | Get a TLS activation
[**ListTLSActivations**](TlsActivationsAPI.md#ListTLSActivations) | **GET** `/tls/activations` | List TLS activations
[**UpdateTLSActivation**](TlsActivationsAPI.md#UpdateTLSActivation) | **PATCH** `/tls/activations/{tls_activation_id}` | Update a certificate



## CreateTLSActivation

Enable TLS for a domain using a custom certificate



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
    tlsActivation := *openapiclient.NewTLSActivation() // TLSActivation |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSActivationsAPI.CreateTLSActivation(ctx).TLSActivation(tlsActivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSActivationsAPI.CreateTLSActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTLSActivation`: TLSActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSActivationsAPI.CreateTLSActivation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTLSActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsActivation** | [**TLSActivation**](TlsActivation.md) |  | 

### Return type

[**TLSActivationResponse**](TlsActivationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTLSActivation

Disable TLS on a domain



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
    tlsActivationID := "tlsActivationId_example" // string | Alphanumeric string identifying a TLS activation.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSActivationsAPI.DeleteTLSActivation(ctx, tlsActivationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSActivationsAPI.DeleteTLSActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsActivationID** | **string** | Alphanumeric string identifying a TLS activation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTLSActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetTLSActivation

Get a TLS activation



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
    tlsActivationID := "tlsActivationId_example" // string | Alphanumeric string identifying a TLS activation.
    include := "tls_certificate,tls_configuration,tls_domain" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSActivationsAPI.GetTLSActivation(ctx, tlsActivationID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSActivationsAPI.GetTLSActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLSActivation`: TLSActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSActivationsAPI.GetTLSActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsActivationID** | **string** | Alphanumeric string identifying a TLS activation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  | 

### Return type

[**TLSActivationResponse**](TlsActivationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTLSActivations

List TLS activations



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
    filterTLSCertificateID := "filterTLSCertificateId_example" // string | Limit the returned activations to a specific certificate. (optional)
    filterTLSConfigurationID := "filterTLSConfigurationId_example" // string | Limit the returned activations to a specific TLS configuration. (optional)
    filterTLSDomainID := "filterTLSDomainId_example" // string | Limit the returned rules to a specific domain name. (optional)
    include := "tls_certificate,tls_configuration,tls_domain" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSActivationsAPI.ListTLSActivations(ctx).FilterTLSCertificateID(filterTLSCertificateID).FilterTLSConfigurationID(filterTLSConfigurationID).FilterTLSDomainID(filterTLSDomainID).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSActivationsAPI.ListTLSActivations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSActivations`: TLSActivationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSActivationsAPI.ListTLSActivations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSActivationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTLSCertificateID** | **string** | Limit the returned activations to a specific certificate. |  **filterTLSConfigurationID** | **string** | Limit the returned activations to a specific TLS configuration. |  **filterTLSDomainID** | **string** | Limit the returned rules to a specific domain name. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TLSActivationsResponse**](TlsActivationsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateTLSActivation

Update a certificate



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
    tlsActivationID := "tlsActivationId_example" // string | Alphanumeric string identifying a TLS activation.
    tlsActivation := *openapiclient.NewTLSActivation() // TLSActivation |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSActivationsAPI.UpdateTLSActivation(ctx, tlsActivationID).TLSActivation(tlsActivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSActivationsAPI.UpdateTLSActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTLSActivation`: TLSActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSActivationsAPI.UpdateTLSActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsActivationID** | **string** | Alphanumeric string identifying a TLS activation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTLSActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsActivation** | [**TLSActivation**](TlsActivation.md) |  | 

### Return type

[**TLSActivationResponse**](TlsActivationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
