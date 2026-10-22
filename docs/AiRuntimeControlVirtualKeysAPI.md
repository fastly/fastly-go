# AiRuntimeControlVirtualKeysAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArcCreateVirtualKey**](AiRuntimeControlVirtualKeysAPI.md#ArcCreateVirtualKey) | **POST** `/ai-runtime-control/v1/keys` | Create a virtual key
[**ArcDeleteVirtualKey**](AiRuntimeControlVirtualKeysAPI.md#ArcDeleteVirtualKey) | **DELETE** `/ai-runtime-control/v1/keys/{key_id}` | Delete a virtual key
[**ArcGetVirtualKey**](AiRuntimeControlVirtualKeysAPI.md#ArcGetVirtualKey) | **GET** `/ai-runtime-control/v1/keys/{key_id}` | Get a virtual key
[**ArcGetVirtualKeyFailover**](AiRuntimeControlVirtualKeysAPI.md#ArcGetVirtualKeyFailover) | **GET** `/ai-runtime-control/v1/keys/{key_id}/failover` | Get a virtual key&#39;s failover policy
[**ArcListVirtualKeys**](AiRuntimeControlVirtualKeysAPI.md#ArcListVirtualKeys) | **GET** `/ai-runtime-control/v1/keys` | List virtual keys
[**ArcRotateVirtualKey**](AiRuntimeControlVirtualKeysAPI.md#ArcRotateVirtualKey) | **POST** `/ai-runtime-control/v1/keys/{key_id}/rotate` | Rotate a virtual key
[**ArcUpdateVirtualKey**](AiRuntimeControlVirtualKeysAPI.md#ArcUpdateVirtualKey) | **PATCH** `/ai-runtime-control/v1/keys/{key_id}` | Update a virtual key
[**ArcUpdateVirtualKeyFailover**](AiRuntimeControlVirtualKeysAPI.md#ArcUpdateVirtualKeyFailover) | **PUT** `/ai-runtime-control/v1/keys/{key_id}/failover` | Update a virtual key&#39;s failover policy



## ArcCreateVirtualKey

Create a virtual key



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
    arcVirtualKey := *openapiclient.NewArcVirtualKey("production-anthropic-key") // ArcVirtualKey | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcCreateVirtualKey(ctx).ArcVirtualKey(arcVirtualKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcCreateVirtualKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcCreateVirtualKey`: ArcCreateVirtualKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcCreateVirtualKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcCreateVirtualKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcVirtualKey** | [**ArcVirtualKey**](ArcVirtualKey.md) |  | 

### Return type

[**ArcCreateVirtualKeyResponse**](ArcCreateVirtualKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcDeleteVirtualKey

Delete a virtual key



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
    keyId := "7fVVQb1nZaJgnzCn4B4v3B" // string | The ID of the virtual key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcDeleteVirtualKey(ctx, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcDeleteVirtualKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the virtual key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcDeleteVirtualKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcGetVirtualKey

Get a virtual key



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
    keyId := "7fVVQb1nZaJgnzCn4B4v3B" // string | The ID of the virtual key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcGetVirtualKey(ctx, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcGetVirtualKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetVirtualKey`: ArcVirtualKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcGetVirtualKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the virtual key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcGetVirtualKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArcVirtualKeyResponse**](ArcVirtualKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcGetVirtualKeyFailover

Get a virtual key's failover policy



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
    keyId := "7fVVQb1nZaJgnzCn4B4v3B" // string | The ID of the virtual key.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcGetVirtualKeyFailover(ctx, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcGetVirtualKeyFailover`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetVirtualKeyFailover`: ArcVirtualKeyFailoverResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcGetVirtualKeyFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the virtual key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcGetVirtualKeyFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArcVirtualKeyFailoverResponse**](ArcVirtualKeyFailoverResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcListVirtualKeys

List virtual keys



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
    includeDeleted := true // bool | Include deleted virtual keys in the results. (optional) (default to false)
    search := "production" // string | Filter results to virtual keys whose name contains this substring. (optional)
    cursor := "MTA6MA" // string | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results, or the `prev_cursor` value to request the previous page.  (optional)
    limit := int32(10) // int32 | Limit how many results are returned. (optional) (default to 10)
    sort := "created_at" // string | Sort field. Prefix with `-` for descending order (e.g., `sort=-created_at`).  (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcListVirtualKeys(ctx).IncludeDeleted(includeDeleted).Search(search).Cursor(cursor).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcListVirtualKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcListVirtualKeys`: ArcVirtualKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcListVirtualKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcListVirtualKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeleted** | **bool** | Include deleted virtual keys in the results. | [default to false] **search** | **string** | Filter results to virtual keys whose name contains this substring. |  **cursor** | **string** | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results, or the `prev_cursor` value to request the previous page.  |  **limit** | **int32** | Limit how many results are returned. | [default to 10] **sort** | **string** | Sort field. Prefix with `-` for descending order (e.g., `sort&#x3D;-created_at`).  | [default to &quot;created_at&quot;]

### Return type

[**ArcVirtualKeysResponse**](ArcVirtualKeysResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcRotateVirtualKey

Rotate a virtual key



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
    keyId := "7fVVQb1nZaJgnzCn4B4v3B" // string | The ID of the virtual key.
    arcRotateVirtualKeyRequest := *openapiclient.NewArcRotateVirtualKeyRequest() // ArcRotateVirtualKeyRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcRotateVirtualKey(ctx, keyId).ArcRotateVirtualKeyRequest(arcRotateVirtualKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcRotateVirtualKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcRotateVirtualKey`: ArcRotateVirtualKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcRotateVirtualKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the virtual key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcRotateVirtualKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcRotateVirtualKeyRequest** | [**ArcRotateVirtualKeyRequest**](ArcRotateVirtualKeyRequest.md) |  | 

### Return type

[**ArcRotateVirtualKeyResponse**](ArcRotateVirtualKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcUpdateVirtualKey

Update a virtual key



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
    keyId := "7fVVQb1nZaJgnzCn4B4v3B" // string | The ID of the virtual key.
    arcUpdateVirtualKeyRequest := *openapiclient.NewArcUpdateVirtualKeyRequest() // ArcUpdateVirtualKeyRequest | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcUpdateVirtualKey(ctx, keyId).ArcUpdateVirtualKeyRequest(arcUpdateVirtualKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcUpdateVirtualKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcUpdateVirtualKey`: ArcUpdateVirtualKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcUpdateVirtualKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the virtual key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcUpdateVirtualKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcUpdateVirtualKeyRequest** | [**ArcUpdateVirtualKeyRequest**](ArcUpdateVirtualKeyRequest.md) |  | 

### Return type

[**ArcUpdateVirtualKeyResponse**](ArcUpdateVirtualKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcUpdateVirtualKeyFailover

Update a virtual key's failover policy



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
    keyId := "7fVVQb1nZaJgnzCn4B4v3B" // string | The ID of the virtual key.
    arcVirtualKeyFailoverRequest := *openapiclient.NewArcVirtualKeyFailoverRequest([]openapiclient.ArcFailoverTarget{*openapiclient.NewArcFailoverTarget("pTfnUUBcU1yv8SoRJn5PZ2")}) // ArcVirtualKeyFailoverRequest | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlVirtualKeysAPI.ArcUpdateVirtualKeyFailover(ctx, keyId).ArcVirtualKeyFailoverRequest(arcVirtualKeyFailoverRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlVirtualKeysAPI.ArcUpdateVirtualKeyFailover`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcUpdateVirtualKeyFailover`: ArcVirtualKeyFailoverResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlVirtualKeysAPI.ArcUpdateVirtualKeyFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the virtual key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcUpdateVirtualKeyFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcVirtualKeyFailoverRequest** | [**ArcVirtualKeyFailoverRequest**](ArcVirtualKeyFailoverRequest.md) |  | 

### Return type

[**ArcVirtualKeyFailoverResponse**](ArcVirtualKeyFailoverResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

