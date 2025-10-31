# TlsBulkCertificatesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBulkTlsCert**](TlsBulkCertificatesAPI.md#DeleteBulkTlsCert) | **DELETE** `/tls/bulk/certificates/{certificate_id}` | Delete a certificate
[**GetTlsBulkCert**](TlsBulkCertificatesAPI.md#GetTlsBulkCert) | **GET** `/tls/bulk/certificates/{certificate_id}` | Get a certificate
[**ListTlsBulkCerts**](TlsBulkCertificatesAPI.md#ListTlsBulkCerts) | **GET** `/tls/bulk/certificates` | List certificates
[**UpdateBulkTlsCert**](TlsBulkCertificatesAPI.md#UpdateBulkTlsCert) | **PATCH** `/tls/bulk/certificates/{certificate_id}` | Update a certificate
[**UploadTlsBulkCert**](TlsBulkCertificatesAPI.md#UploadTlsBulkCert) | **POST** `/tls/bulk/certificates` | Upload a certificate



## DeleteBulkTlsCert

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
    certificateId := "certificateId_example" // string | Alphanumeric string identifying a TLS bulk certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsBulkCertificatesAPI.DeleteBulkTlsCert(ctx, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsBulkCertificatesAPI.DeleteBulkTlsCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** | Alphanumeric string identifying a TLS bulk certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkTlsCertRequest struct via the builder pattern


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


## GetTlsBulkCert

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
    certificateId := "certificateId_example" // string | Alphanumeric string identifying a TLS bulk certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsBulkCertificatesAPI.GetTlsBulkCert(ctx, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsBulkCertificatesAPI.GetTlsBulkCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsBulkCert`: TlsBulkCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsBulkCertificatesAPI.GetTlsBulkCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** | Alphanumeric string identifying a TLS bulk certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsBulkCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TlsBulkCertificateResponse**](TlsBulkCertificateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTlsBulkCerts

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
    filterTlsDomainId := "filterTlsDomainId_example" // string | Filter certificates by their matching, fully-qualified domain name. (optional)
    filterNotBefore := "filterNotBefore_example" // string | Filter the returned certificates by not_before date in UTC.  Accepts parameters: lt, lte, gt, gte (e.g., filter[not_before][gte]=2020-05-05).  (optional)
    filterNotAfter := "filterNotAfter_example" // string | Filter the returned certificates by expiry date in UTC.  Accepts parameters: lt, lte, gt, gte (e.g., filter[not_after][lte]=2020-05-05).  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsBulkCertificatesAPI.ListTlsBulkCerts(ctx).FilterTlsDomainId(filterTlsDomainId).FilterNotBefore(filterNotBefore).FilterNotAfter(filterNotAfter).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsBulkCertificatesAPI.ListTlsBulkCerts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsBulkCerts`: TlsBulkCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsBulkCertificatesAPI.ListTlsBulkCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsBulkCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTlsDomainId** | **string** | Filter certificates by their matching, fully-qualified domain name. |  **filterNotBefore** | **string** | Filter the returned certificates by not_before date in UTC.  Accepts parameters: lt, lte, gt, gte (e.g., filter[not_before][gte]&#x3D;2020-05-05).  |  **filterNotAfter** | **string** | Filter the returned certificates by expiry date in UTC.  Accepts parameters: lt, lte, gt, gte (e.g., filter[not_after][lte]&#x3D;2020-05-05).  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**TlsBulkCertificatesResponse**](TlsBulkCertificatesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateBulkTlsCert

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
    certificateId := "certificateId_example" // string | Alphanumeric string identifying a TLS bulk certificate.
    tlsBulkCertificate := *openapiclient.NewTlsBulkCertificate() // TlsBulkCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsBulkCertificatesAPI.UpdateBulkTlsCert(ctx, certificateId).TlsBulkCertificate(tlsBulkCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsBulkCertificatesAPI.UpdateBulkTlsCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBulkTlsCert`: TlsBulkCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsBulkCertificatesAPI.UpdateBulkTlsCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** | Alphanumeric string identifying a TLS bulk certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkTlsCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsBulkCertificate** | [**TlsBulkCertificate**](TlsBulkCertificate.md) |  | 

### Return type

[**TlsBulkCertificateResponse**](TlsBulkCertificateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UploadTlsBulkCert

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
    tlsBulkCertificate := *openapiclient.NewTlsBulkCertificate() // TlsBulkCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsBulkCertificatesAPI.UploadTlsBulkCert(ctx).TlsBulkCertificate(tlsBulkCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsBulkCertificatesAPI.UploadTlsBulkCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTlsBulkCert`: TlsBulkCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsBulkCertificatesAPI.UploadTlsBulkCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTlsBulkCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsBulkCertificate** | [**TlsBulkCertificate**](TlsBulkCertificate.md) |  | 

### Return type

[**TlsBulkCertificateResponse**](TlsBulkCertificateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

