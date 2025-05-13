# SecretStoreAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientKey**](SecretStoreAPI.md#ClientKey) | **POST** `/resources/stores/secret/client-key` | Create new client key
[**CreateSecretStore**](SecretStoreAPI.md#CreateSecretStore) | **POST** `/resources/stores/secret` | Create new secret store
[**DeleteSecretStore**](SecretStoreAPI.md#DeleteSecretStore) | **DELETE** `/resources/stores/secret/{store_id}` | Delete secret store
[**GetSecretStore**](SecretStoreAPI.md#GetSecretStore) | **GET** `/resources/stores/secret/{store_id}` | Get secret store by ID
[**GetSecretStores**](SecretStoreAPI.md#GetSecretStores) | **GET** `/resources/stores/secret` | Get all secret stores
[**SigningKey**](SecretStoreAPI.md#SigningKey) | **GET** `/resources/stores/secret/signing-key` | Get public key



## ClientKey

Create new client key



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreAPI.ClientKey(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreAPI.ClientKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientKey`: ClientKey
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreAPI.ClientKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClientKeyRequest struct via the builder pattern



### Return type

[**ClientKey**](ClientKey.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateSecretStore

Create new secret store



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
    secretStore := *openapiclient.NewSecretStore() // SecretStore |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreAPI.CreateSecretStore(ctx).SecretStore(secretStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreAPI.CreateSecretStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecretStore`: SecretStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreAPI.CreateSecretStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secretStore** | [**SecretStore**](SecretStore.md) |  | 

### Return type

[**SecretStoreResponse**](SecretStoreResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteSecretStore

Delete secret store



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
    storeID := "storeId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreAPI.DeleteSecretStore(ctx, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreAPI.DeleteSecretStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretStoreRequest struct via the builder pattern


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


## GetSecretStore

Get secret store by ID



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
    storeID := "storeId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreAPI.GetSecretStore(ctx, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreAPI.GetSecretStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretStore`: SecretStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreAPI.GetSecretStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretStoreResponse**](SecretStoreResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetSecretStores

Get all secret stores



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
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 200. (optional) (default to "100")
    name := "name_example" // string | Returns a one-element array containing the details for the named secret store. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreAPI.GetSecretStores(ctx).Cursor(cursor).Limit(limit).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreAPI.GetSecretStores`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretStores`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreAPI.GetSecretStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **string** | Number of results per page. The maximum is 200. | [default to &quot;100&quot;] **name** | **string** | Returns a one-element array containing the details for the named secret store. | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SigningKey

Get public key



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreAPI.SigningKey(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreAPI.SigningKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SigningKey`: SigningKey
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreAPI.SigningKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSigningKeyRequest struct via the builder pattern



### Return type

[**SigningKey**](SigningKey.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
