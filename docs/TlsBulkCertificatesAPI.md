# TLSBulkCertificatesAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBulkTLSCert**](TlsBulkCertificatesAPI.md#DeleteBulkTLSCert) | **DELETE** `/tls/bulk/certificates/{certificate_id}` | Delete a certificate
[**GetTLSBulkCert**](TlsBulkCertificatesAPI.md#GetTLSBulkCert) | **GET** `/tls/bulk/certificates/{certificate_id}` | Get a certificate
[**ListTLSBulkCerts**](TlsBulkCertificatesAPI.md#ListTLSBulkCerts) | **GET** `/tls/bulk/certificates` | List certificates
[**UpdateBulkTLSCert**](TlsBulkCertificatesAPI.md#UpdateBulkTLSCert) | **PATCH** `/tls/bulk/certificates/{certificate_id}` | Update a certificate
[**UploadTLSBulkCert**](TlsBulkCertificatesAPI.md#UploadTLSBulkCert) | **POST** `/tls/bulk/certificates` | Upload a certificate



## DeleteBulkTLSCert

Delete a certificate



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
    certificateID := "certificateId_example" // string | Alphanumeric string identifying a TLS bulk certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSBulkCertificatesAPI.DeleteBulkTLSCert(ctx, certificateID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSBulkCertificatesAPI.DeleteBulkTLSCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateID** | **string** | Alphanumeric string identifying a TLS bulk certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkTLSCertRequest struct via the builder pattern


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


## GetTLSBulkCert

Get a certificate



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
    certificateID := "certificateId_example" // string | Alphanumeric string identifying a TLS bulk certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSBulkCertificatesAPI.GetTLSBulkCert(ctx, certificateID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSBulkCertificatesAPI.GetTLSBulkCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLSBulkCert`: TLSBulkCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSBulkCertificatesAPI.GetTLSBulkCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateID** | **string** | Alphanumeric string identifying a TLS bulk certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSBulkCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TLSBulkCertificateResponse**](TlsBulkCertificateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTLSBulkCerts

List certificates



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
    filterTLSDomainID := "filterTLSDomainId_example" // string | Filter certificates by their matching, fully-qualified domain name. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSBulkCertificatesAPI.ListTLSBulkCerts(ctx).FilterTLSDomainID(filterTLSDomainID).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSBulkCertificatesAPI.ListTLSBulkCerts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSBulkCerts`: TLSBulkCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSBulkCertificatesAPI.ListTLSBulkCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSBulkCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTLSDomainID** | **string** | Filter certificates by their matching, fully-qualified domain name. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**TLSBulkCertificatesResponse**](TlsBulkCertificatesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateBulkTLSCert

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
    certificateID := "certificateId_example" // string | Alphanumeric string identifying a TLS bulk certificate.
    tlsBulkCertificate := *openapiclient.NewTLSBulkCertificate() // TLSBulkCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSBulkCertificatesAPI.UpdateBulkTLSCert(ctx, certificateID).TLSBulkCertificate(tlsBulkCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSBulkCertificatesAPI.UpdateBulkTLSCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBulkTLSCert`: TLSBulkCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSBulkCertificatesAPI.UpdateBulkTLSCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateID** | **string** | Alphanumeric string identifying a TLS bulk certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkTLSCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsBulkCertificate** | [**TLSBulkCertificate**](TlsBulkCertificate.md) |  | 

### Return type

[**TLSBulkCertificateResponse**](TlsBulkCertificateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UploadTLSBulkCert

Upload a certificate



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
    tlsBulkCertificate := *openapiclient.NewTLSBulkCertificate() // TLSBulkCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSBulkCertificatesAPI.UploadTLSBulkCert(ctx).TLSBulkCertificate(tlsBulkCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSBulkCertificatesAPI.UploadTLSBulkCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTLSBulkCert`: TLSBulkCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSBulkCertificatesAPI.UploadTLSBulkCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTLSBulkCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsBulkCertificate** | [**TLSBulkCertificate**](TlsBulkCertificate.md) |  | 

### Return type

[**TLSBulkCertificateResponse**](TlsBulkCertificateResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
