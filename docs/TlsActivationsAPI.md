# TlsActivationsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTlsActivation**](TlsActivationsAPI.md#CreateTlsActivation) | **POST** `/tls/activations` | Enable TLS for a domain using a custom certificate
[**DeleteTlsActivation**](TlsActivationsAPI.md#DeleteTlsActivation) | **DELETE** `/tls/activations/{tls_activation_id}` | Disable TLS on a domain
[**GetTlsActivation**](TlsActivationsAPI.md#GetTlsActivation) | **GET** `/tls/activations/{tls_activation_id}` | Get a TLS activation
[**ListTlsActivations**](TlsActivationsAPI.md#ListTlsActivations) | **GET** `/tls/activations` | List TLS activations
[**UpdateTlsActivation**](TlsActivationsAPI.md#UpdateTlsActivation) | **PATCH** `/tls/activations/{tls_activation_id}` | Update a certificate



## CreateTlsActivation

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
    tlsActivation := *openapiclient.NewTlsActivation() // TlsActivation |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsActivationsAPI.CreateTlsActivation(ctx).TlsActivation(tlsActivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsActivationsAPI.CreateTlsActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTlsActivation`: TlsActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsActivationsAPI.CreateTlsActivation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsActivation** | [**TlsActivation**](TlsActivation.md) |  | 

### Return type

[**TlsActivationResponse**](TlsActivationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTlsActivation

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
    tlsActivationId := "tlsActivationId_example" // string | Alphanumeric string identifying a TLS activation.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsActivationsAPI.DeleteTlsActivation(ctx, tlsActivationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsActivationsAPI.DeleteTlsActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsActivationId** | **string** | Alphanumeric string identifying a TLS activation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsActivationRequest struct via the builder pattern


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


## GetTlsActivation

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
    tlsActivationId := "tlsActivationId_example" // string | Alphanumeric string identifying a TLS activation.
    include := "tls_certificate,tls_configuration,tls_domain" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsActivationsAPI.GetTlsActivation(ctx, tlsActivationId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsActivationsAPI.GetTlsActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsActivation`: TlsActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsActivationsAPI.GetTlsActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsActivationId** | **string** | Alphanumeric string identifying a TLS activation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  | 

### Return type

[**TlsActivationResponse**](TlsActivationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTlsActivations

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
    filterTlsCertificateId := "filterTlsCertificateId_example" // string | Limit the returned activations to a specific certificate. (optional)
    filterTlsConfigurationId := "filterTlsConfigurationId_example" // string | Limit the returned activations to a specific TLS configuration. (optional)
    filterTlsDomainId := "filterTlsDomainId_example" // string | Limit the returned rules to a specific domain name. (optional)
    include := "tls_certificate,tls_configuration,tls_domain" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsActivationsAPI.ListTlsActivations(ctx).FilterTlsCertificateId(filterTlsCertificateId).FilterTlsConfigurationId(filterTlsConfigurationId).FilterTlsDomainId(filterTlsDomainId).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsActivationsAPI.ListTlsActivations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsActivations`: TlsActivationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsActivationsAPI.ListTlsActivations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsActivationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTlsCertificateId** | **string** | Limit the returned activations to a specific certificate. |  **filterTlsConfigurationId** | **string** | Limit the returned activations to a specific TLS configuration. |  **filterTlsDomainId** | **string** | Limit the returned rules to a specific domain name. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_certificate`, `tls_configuration`, and `tls_domain`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TlsActivationsResponse**](TlsActivationsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateTlsActivation

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
    tlsActivationId := "tlsActivationId_example" // string | Alphanumeric string identifying a TLS activation.
    tlsActivation := *openapiclient.NewTlsActivation() // TlsActivation |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsActivationsAPI.UpdateTlsActivation(ctx, tlsActivationId).TlsActivation(tlsActivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsActivationsAPI.UpdateTlsActivation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTlsActivation`: TlsActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsActivationsAPI.UpdateTlsActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsActivationId** | **string** | Alphanumeric string identifying a TLS activation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTlsActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsActivation** | [**TlsActivation**](TlsActivation.md) |  | 

### Return type

[**TlsActivationResponse**](TlsActivationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

