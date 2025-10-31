# TlsCertificatesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTlsCert**](TlsCertificatesAPI.md#CreateTlsCert) | **POST** `/tls/certificates` | Create a TLS certificate
[**DeleteTlsCert**](TlsCertificatesAPI.md#DeleteTlsCert) | **DELETE** `/tls/certificates/{tls_certificate_id}` | Delete a TLS certificate
[**GetTlsCert**](TlsCertificatesAPI.md#GetTlsCert) | **GET** `/tls/certificates/{tls_certificate_id}` | Get a TLS certificate
[**GetTlsCertBlob**](TlsCertificatesAPI.md#GetTlsCertBlob) | **GET** `/tls/certificates/{tls_certificate_id}/blob` | Get a TLS certificate blob (Limited Availability)
[**ListTlsCerts**](TlsCertificatesAPI.md#ListTlsCerts) | **GET** `/tls/certificates` | List TLS certificates
[**UpdateTlsCert**](TlsCertificatesAPI.md#UpdateTlsCert) | **PATCH** `/tls/certificates/{tls_certificate_id}` | Update a TLS certificate



## CreateTlsCert

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
    tlsCertificate := *openapiclient.NewTlsCertificate() // TlsCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsCertificatesAPI.CreateTlsCert(ctx).TlsCertificate(tlsCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsCertificatesAPI.CreateTlsCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTlsCert`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TlsCertificatesAPI.CreateTlsCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCertificate** | [**TlsCertificate**](TlsCertificate.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTlsCert

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
    tlsCertificateId := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsCertificatesAPI.DeleteTlsCert(ctx, tlsCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsCertificatesAPI.DeleteTlsCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateId** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsCertRequest struct via the builder pattern


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


## GetTlsCert

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
    tlsCertificateId := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsCertificatesAPI.GetTlsCert(ctx, tlsCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsCertificatesAPI.GetTlsCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsCert`: TlsCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsCertificatesAPI.GetTlsCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateId** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TlsCertificateResponse**](TlsCertificateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetTlsCertBlob

Get a TLS certificate blob (Limited Availability)



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
    tlsCertificateId := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsCertificatesAPI.GetTlsCertBlob(ctx, tlsCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsCertificatesAPI.GetTlsCertBlob`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsCertBlob`: TlsCertificateBlobResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsCertificatesAPI.GetTlsCertBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateId** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsCertBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TlsCertificateBlobResponse**](TlsCertificateBlobResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTlsCerts

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
    filterInUse := "filterInUse_example" // string | Optional. Limit the returned certificates to those currently using Fastly to terminate TLS (that is, certificates associated with an activation). Permitted values: true, false. (optional)
    filterNotAfter := "filterNotAfter_example" // string | Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]=2020-05-05).  (optional)
    filterTlsDomainsId := "filterTlsDomainsId_example" // string | Limit the returned certificates to those that include the specific domain. (optional)
    include := "include_example" // string | Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`.  (optional)
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "-created_at")
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsCertificatesAPI.ListTlsCerts(ctx).FilterInUse(filterInUse).FilterNotAfter(filterNotAfter).FilterTlsDomainsId(filterTlsDomainsId).Include(include).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsCertificatesAPI.ListTlsCerts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsCerts`: TlsCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsCertificatesAPI.ListTlsCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInUse** | **string** | Optional. Limit the returned certificates to those currently using Fastly to terminate TLS (that is, certificates associated with an activation). Permitted values: true, false. |  **filterNotAfter** | **string** | Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]&#x3D;2020-05-05).  |  **filterTlsDomainsId** | **string** | Limit the returned certificates to those that include the specific domain. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`.  |  **sort** | **string** | The order in which to list the results. | [default to &quot;-created_at&quot;] **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TlsCertificatesResponse**](TlsCertificatesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateTlsCert

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
    tlsCertificateId := "tlsCertificateId_example" // string | Alphanumeric string identifying a TLS certificate.
    tlsCertificate := *openapiclient.NewTlsCertificate() // TlsCertificate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsCertificatesAPI.UpdateTlsCert(ctx, tlsCertificateId).TlsCertificate(tlsCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsCertificatesAPI.UpdateTlsCert`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTlsCert`: TlsCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsCertificatesAPI.UpdateTlsCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsCertificateId** | **string** | Alphanumeric string identifying a TLS certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTlsCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCertificate** | [**TlsCertificate**](TlsCertificate.md) |  | 

### Return type

[**TlsCertificateResponse**](TlsCertificateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

