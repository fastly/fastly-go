# TlsPrivateKeysAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTlsKey**](TlsPrivateKeysAPI.md#CreateTlsKey) | **POST** `/tls/private_keys` | Create a TLS private key
[**DeleteTlsKey**](TlsPrivateKeysAPI.md#DeleteTlsKey) | **DELETE** `/tls/private_keys/{tls_private_key_id}` | Delete a TLS private key
[**GetTlsKey**](TlsPrivateKeysAPI.md#GetTlsKey) | **GET** `/tls/private_keys/{tls_private_key_id}` | Get a TLS private key
[**ListTlsKeys**](TlsPrivateKeysAPI.md#ListTlsKeys) | **GET** `/tls/private_keys` | List TLS private keys



## CreateTlsKey

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
    tlsPrivateKey := *openapiclient.NewTlsPrivateKey() // TlsPrivateKey |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsPrivateKeysAPI.CreateTlsKey(ctx).TlsPrivateKey(tlsPrivateKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsPrivateKeysAPI.CreateTlsKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTlsKey`: TlsPrivateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsPrivateKeysAPI.CreateTlsKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsPrivateKey** | [**TlsPrivateKey**](TlsPrivateKey.md) |  | 

### Return type

[**TlsPrivateKeyResponse**](TlsPrivateKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteTlsKey

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
    tlsPrivateKeyId := "tlsPrivateKeyId_example" // string | Alphanumeric string identifying a private Key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsPrivateKeysAPI.DeleteTlsKey(ctx, tlsPrivateKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsPrivateKeysAPI.DeleteTlsKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsPrivateKeyId** | **string** | Alphanumeric string identifying a private Key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsKeyRequest struct via the builder pattern


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


## GetTlsKey

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
    tlsPrivateKeyId := "tlsPrivateKeyId_example" // string | Alphanumeric string identifying a private Key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsPrivateKeysAPI.GetTlsKey(ctx, tlsPrivateKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsPrivateKeysAPI.GetTlsKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsKey`: TlsPrivateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsPrivateKeysAPI.GetTlsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsPrivateKeyId** | **string** | Alphanumeric string identifying a private Key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TlsPrivateKeyResponse**](TlsPrivateKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTlsKeys

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
    resp, r, err := apiClient.TlsPrivateKeysAPI.ListTlsKeys(ctx).FilterInUse(filterInUse).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsPrivateKeysAPI.ListTlsKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsKeys`: TlsPrivateKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsPrivateKeysAPI.ListTlsKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInUse** | **string** | Limit the returned keys to those without any matching TLS certificates. The only valid value is false. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TlsPrivateKeysResponse**](TlsPrivateKeysResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

