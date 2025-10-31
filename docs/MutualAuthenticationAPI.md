# MutualAuthenticationAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMutualTlsAuthentication**](MutualAuthenticationAPI.md#CreateMutualTlsAuthentication) | **POST** `/tls/mutual_authentications` | Create a Mutual Authentication
[**DeleteMutualTls**](MutualAuthenticationAPI.md#DeleteMutualTls) | **DELETE** `/tls/mutual_authentications/{mutual_authentication_id}` | Delete a Mutual TLS
[**GetMutualAuthentication**](MutualAuthenticationAPI.md#GetMutualAuthentication) | **GET** `/tls/mutual_authentications/{mutual_authentication_id}` | Get a Mutual Authentication
[**ListMutualAuthentications**](MutualAuthenticationAPI.md#ListMutualAuthentications) | **GET** `/tls/mutual_authentications` | List Mutual Authentications
[**PatchMutualAuthentication**](MutualAuthenticationAPI.md#PatchMutualAuthentication) | **PATCH** `/tls/mutual_authentications/{mutual_authentication_id}` | Update a Mutual Authentication



## CreateMutualTlsAuthentication

Create a Mutual Authentication



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
    mutualAuthentication := *openapiclient.NewMutualAuthentication() // MutualAuthentication |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.MutualAuthenticationAPI.CreateMutualTlsAuthentication(ctx).MutualAuthentication(mutualAuthentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutualAuthenticationAPI.CreateMutualTlsAuthentication`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMutualTlsAuthentication`: MutualAuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `MutualAuthenticationAPI.CreateMutualTlsAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMutualTlsAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mutualAuthentication** | [**MutualAuthentication**](MutualAuthentication.md) |  | 

### Return type

[**MutualAuthenticationResponse**](MutualAuthenticationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteMutualTls

Delete a Mutual TLS



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
    mutualAuthenticationId := "mutualAuthenticationId_example" // string | Alphanumeric string identifying a mutual authentication.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.MutualAuthenticationAPI.DeleteMutualTls(ctx, mutualAuthenticationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutualAuthenticationAPI.DeleteMutualTls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mutualAuthenticationId** | **string** | Alphanumeric string identifying a mutual authentication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMutualTlsRequest struct via the builder pattern


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


## GetMutualAuthentication

Get a Mutual Authentication



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
    mutualAuthenticationId := "mutualAuthenticationId_example" // string | Alphanumeric string identifying a mutual authentication.
    include := "include_example" // string | Comma-separated list of related objects to include (optional). Permitted values: `tls_activations`. Including TLS activations will provide you with the TLS domain names that are related to your Mutual TLS authentication.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.MutualAuthenticationAPI.GetMutualAuthentication(ctx, mutualAuthenticationId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutualAuthenticationAPI.GetMutualAuthentication`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMutualAuthentication`: MutualAuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `MutualAuthenticationAPI.GetMutualAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mutualAuthenticationId** | **string** | Alphanumeric string identifying a mutual authentication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Comma-separated list of related objects to include (optional). Permitted values: `tls_activations`. Including TLS activations will provide you with the TLS domain names that are related to your Mutual TLS authentication.  | 

### Return type

[**MutualAuthenticationResponse**](MutualAuthenticationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListMutualAuthentications

List Mutual Authentications



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
    include := "include_example" // string | Comma-separated list of related objects to include (optional). Permitted values: `tls_activations`. Including TLS activations will provide you with the TLS domain names that are related to your Mutual TLS authentication.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.MutualAuthenticationAPI.ListMutualAuthentications(ctx).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutualAuthenticationAPI.ListMutualAuthentications`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMutualAuthentications`: MutualAuthenticationsResponse
    fmt.Fprintf(os.Stdout, "Response from `MutualAuthenticationAPI.ListMutualAuthentications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMutualAuthenticationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Comma-separated list of related objects to include (optional). Permitted values: `tls_activations`. Including TLS activations will provide you with the TLS domain names that are related to your Mutual TLS authentication.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**MutualAuthenticationsResponse**](MutualAuthenticationsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## PatchMutualAuthentication

Update a Mutual Authentication



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
    mutualAuthenticationId := "mutualAuthenticationId_example" // string | Alphanumeric string identifying a mutual authentication.
    mutualAuthentication := *openapiclient.NewMutualAuthentication() // MutualAuthentication |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.MutualAuthenticationAPI.PatchMutualAuthentication(ctx, mutualAuthenticationId).MutualAuthentication(mutualAuthentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutualAuthenticationAPI.PatchMutualAuthentication`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMutualAuthentication`: MutualAuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `MutualAuthenticationAPI.PatchMutualAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mutualAuthenticationId** | **string** | Alphanumeric string identifying a mutual authentication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMutualAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mutualAuthentication** | [**MutualAuthentication**](MutualAuthentication.md) |  | 

### Return type

[**MutualAuthenticationResponse**](MutualAuthenticationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

