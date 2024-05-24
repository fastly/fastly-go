# SecretStoreItemAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](SecretStoreItemAPI.md#CreateSecret) | **POST** `/resources/stores/secret/{store_id}/secrets` | Create a new secret in a store.
[**DeleteSecret**](SecretStoreItemAPI.md#DeleteSecret) | **DELETE** `/resources/stores/secret/{store_id}/secrets/{secret_name}` | Delete a secret from a store.
[**GetSecret**](SecretStoreItemAPI.md#GetSecret) | **GET** `/resources/stores/secret/{store_id}/secrets/{secret_name}` | Get secret metadata.
[**GetSecrets**](SecretStoreItemAPI.md#GetSecrets) | **GET** `/resources/stores/secret/{store_id}/secrets` | List secrets within a store.
[**MustRecreateSecret**](SecretStoreItemAPI.md#MustRecreateSecret) | **PATCH** `/resources/stores/secret/{store_id}/secrets` | Recreate a secret in a store.
[**RecreateSecret**](SecretStoreItemAPI.md#RecreateSecret) | **PUT** `/resources/stores/secret/{store_id}/secrets` | Create or recreate a secret in a store.



## CreateSecret

Create a new secret in a store.



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
    secret := *openapiclient.NewSecret() // Secret |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreItemAPI.CreateSecret(ctx, storeID).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreItemAPI.CreateSecret`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreItemAPI.CreateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | [**Secret**](Secret.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteSecret

Delete a secret from a store.



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
    secretName := "secretName_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreItemAPI.DeleteSecret(ctx, storeID, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreItemAPI.DeleteSecret`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 
**secretName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


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


## GetSecret

Get secret metadata.



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
    secretName := "secretName_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreItemAPI.GetSecret(ctx, storeID, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreItemAPI.GetSecret`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreItemAPI.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 
**secretName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetSecrets

List secrets within a store.



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
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := "limit_example" // string | Number of results per page. The maximum is 200. (optional) (default to "100")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreItemAPI.GetSecrets(ctx, storeID).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreItemAPI.GetSecrets`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecrets`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreItemAPI.GetSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **string** | Number of results per page. The maximum is 200. | [default to &quot;100&quot;]

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## MustRecreateSecret

Recreate a secret in a store.



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
    secret := *openapiclient.NewSecret() // Secret |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreItemAPI.MustRecreateSecret(ctx, storeID).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreItemAPI.MustRecreateSecret`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MustRecreateSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreItemAPI.MustRecreateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMustRecreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | [**Secret**](Secret.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RecreateSecret

Create or recreate a secret in a store.



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
    secret := *openapiclient.NewSecret() // Secret |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SecretStoreItemAPI.RecreateSecret(ctx, storeID).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreItemAPI.RecreateSecret`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecreateSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreItemAPI.RecreateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | [**Secret**](Secret.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
