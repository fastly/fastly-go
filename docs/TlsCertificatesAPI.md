# TLSCertificatesAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTLSCert**](TlsCertificatesAPI.md#CreateTLSCert) | **POST** `/tls/certificates` | Create a TLS certificate
[**DeleteTLSCert**](TlsCertificatesAPI.md#DeleteTLSCert) | **DELETE** `/tls/certificates/{tls_certificate_id}` | Delete a TLS certificate
[**GetTLSCert**](TlsCertificatesAPI.md#GetTLSCert) | **GET** `/tls/certificates/{tls_certificate_id}` | Get a TLS certificate
[**ListTLSCerts**](TlsCertificatesAPI.md#ListTLSCerts) | **GET** `/tls/certificates` | List TLS certificates
[**UpdateTLSCert**](TlsCertificatesAPI.md#UpdateTLSCert) | **PATCH** `/tls/certificates/{tls_certificate_id}` | Update a TLS certificate



## CreateTLSCert

Create a TLS certificate



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
    tlsCertificate := *openapiclient.NewTLSCertificate() // TLSCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSCertificatesAPI.CreateTLSCert(ctx).TLSCertificate(tlsCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSCertificatesAPI.CreateTLSCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTLSCert`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `TLSCertificatesAPI.CreateTLSCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTLSCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCertificate** | [**TLSCertificate**](TlsCertificate.md) |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTLSCert

Delete a TLS certificate



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
    tlsCertificateID := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSCertificatesAPI.DeleteTLSCert(ctx, tlsCertificateID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSCertificatesAPI.DeleteTLSCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateID** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTLSCertRequest struct via the builder pattern


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


## GetTLSCert

Get a TLS certificate



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
    tlsCertificateID := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSCertificatesAPI.GetTLSCert(ctx, tlsCertificateID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSCertificatesAPI.GetTLSCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLSCert`: TLSCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSCertificatesAPI.GetTLSCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateID** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TLSCertificateResponse**](TlsCertificateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTLSCerts

List TLS certificates



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
    filterNotAfter := "filterNotAfter_example" // string | Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]=2020-05-05).  (optional)
    filterTLSDomainsID := "filterTLSDomainsId_example" // string | Limit the returned certificates to those that include the specific domain. (optional)
    include := "include_example" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSCertificatesAPI.ListTLSCerts(ctx).FilterNotAfter(filterNotAfter).FilterTLSDomainsID(filterTLSDomainsID).Include(include).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSCertificatesAPI.ListTLSCerts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSCerts`: TLSCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSCertificatesAPI.ListTLSCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterNotAfter** | **string** | Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]&#x3D;2020-05-05).  |  **filterTLSDomainsID** | **string** | Limit the returned certificates to those that include the specific domain. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**TLSCertificatesResponse**](TlsCertificatesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateTLSCert

Update a TLS certificate



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
    tlsCertificateID := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.
    tlsCertificate := *openapiclient.NewTLSCertificate() // TLSCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSCertificatesAPI.UpdateTLSCert(ctx, tlsCertificateID).TLSCertificate(tlsCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSCertificatesAPI.UpdateTLSCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTLSCert`: TLSCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSCertificatesAPI.UpdateTLSCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateID** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTLSCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCertificate** | [**TLSCertificate**](TlsCertificate.md) |  | 

### Return type

[**TLSCertificateResponse**](TlsCertificateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
