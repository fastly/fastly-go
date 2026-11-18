# TLSPrivateKeysAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTLSKey**](TlsPrivateKeysAPI.md#CreateTLSKey) | **POST** `/tls/private_keys` | Create a TLS private key
[**DeleteTLSKey**](TlsPrivateKeysAPI.md#DeleteTLSKey) | **DELETE** `/tls/private_keys/{tls_private_key_id}` | Delete a TLS private key
[**GetTLSKey**](TlsPrivateKeysAPI.md#GetTLSKey) | **GET** `/tls/private_keys/{tls_private_key_id}` | Get a TLS private key
[**ListTLSKeys**](TlsPrivateKeysAPI.md#ListTLSKeys) | **GET** `/tls/private_keys` | List TLS private keys



## CreateTLSKey

Create a TLS private key



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
    tlsPrivateKey := *openapiclient.NewTLSPrivateKey() // TLSPrivateKey |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSPrivateKeysAPI.CreateTLSKey(ctx).TLSPrivateKey(tlsPrivateKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSPrivateKeysAPI.CreateTLSKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTLSKey`: TLSPrivateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSPrivateKeysAPI.CreateTLSKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTLSKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsPrivateKey** | [**TLSPrivateKey**](TlsPrivateKey.md) |  | 

### Return type

[**TLSPrivateKeyResponse**](TlsPrivateKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTLSKey

Delete a TLS private key



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
    tlsPrivateKeyID := "tlsPrivateKeyId_example" // string | Alphanumeric string identifying a private Key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSPrivateKeysAPI.DeleteTLSKey(ctx, tlsPrivateKeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSPrivateKeysAPI.DeleteTLSKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsPrivateKeyID** | **string** | Alphanumeric string identifying a private Key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTLSKeyRequest struct via the builder pattern


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


## GetTLSKey

Get a TLS private key



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
    tlsPrivateKeyID := "tlsPrivateKeyId_example" // string | Alphanumeric string identifying a private Key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSPrivateKeysAPI.GetTLSKey(ctx, tlsPrivateKeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSPrivateKeysAPI.GetTLSKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLSKey`: TLSPrivateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSPrivateKeysAPI.GetTLSKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsPrivateKeyID** | **string** | Alphanumeric string identifying a private Key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TLSPrivateKeyResponse**](TlsPrivateKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTLSKeys

List TLS private keys



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
    filterInUse := "filterInUse_example" // string | Limit the returned keys to those without any matching TLS certificates. The only valid value is false. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSPrivateKeysAPI.ListTLSKeys(ctx).FilterInUse(filterInUse).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSPrivateKeysAPI.ListTLSKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSKeys`: TLSPrivateKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSPrivateKeysAPI.ListTLSKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInUse** | **string** | Limit the returned keys to those without any matching TLS certificates. The only valid value is false. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TLSPrivateKeysResponse**](TlsPrivateKeysResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
