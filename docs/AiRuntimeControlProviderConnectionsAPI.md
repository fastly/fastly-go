# AiRuntimeControlProviderConnectionsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArcCreateProviderConnection**](AiRuntimeControlProviderConnectionsAPI.md#ArcCreateProviderConnection) | **POST** `/ai-runtime-control/v1/provider-connections` | Create a provider connection
[**ArcDeleteProviderConnection**](AiRuntimeControlProviderConnectionsAPI.md#ArcDeleteProviderConnection) | **DELETE** `/ai-runtime-control/v1/provider-connections/{id}` | Delete a provider connection
[**ArcGetProviderConnection**](AiRuntimeControlProviderConnectionsAPI.md#ArcGetProviderConnection) | **GET** `/ai-runtime-control/v1/provider-connections/{id}` | Get a provider connection
[**ArcListProviderConnections**](AiRuntimeControlProviderConnectionsAPI.md#ArcListProviderConnections) | **GET** `/ai-runtime-control/v1/provider-connections` | List provider connections
[**ArcUpdateProviderConnection**](AiRuntimeControlProviderConnectionsAPI.md#ArcUpdateProviderConnection) | **PATCH** `/ai-runtime-control/v1/provider-connections/{id}` | Update a provider connection



## ArcCreateProviderConnection

Create a provider connection



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
    arcProviderConnectionRequest := openapiclient.arc_provider_connection_request{ArcProviderConnectionApiKey: openapiclient.NewArcProviderConnectionApiKey("my-openai-connection", []string{"Models_example"}, "https://api.openai.com", "sk-...")} // ArcProviderConnectionRequest | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlProviderConnectionsAPI.ArcCreateProviderConnection(ctx).ArcProviderConnectionRequest(arcProviderConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProviderConnectionsAPI.ArcCreateProviderConnection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcCreateProviderConnection`: ArcProviderConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProviderConnectionsAPI.ArcCreateProviderConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcCreateProviderConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcProviderConnectionRequest** | [**ArcProviderConnectionRequest**](ArcProviderConnectionRequest.md) |  | 

### Return type

[**ArcProviderConnectionResponse**](ArcProviderConnectionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcDeleteProviderConnection

Delete a provider connection



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
    id := "pTfnUUBcU1yv8SoRJn5PZ2" // string | The ID of the provider connection.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlProviderConnectionsAPI.ArcDeleteProviderConnection(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProviderConnectionsAPI.ArcDeleteProviderConnection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the provider connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcDeleteProviderConnectionRequest struct via the builder pattern


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


## ArcGetProviderConnection

Get a provider connection



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
    id := "pTfnUUBcU1yv8SoRJn5PZ2" // string | The ID of the provider connection.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlProviderConnectionsAPI.ArcGetProviderConnection(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProviderConnectionsAPI.ArcGetProviderConnection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetProviderConnection`: ArcGetProviderConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProviderConnectionsAPI.ArcGetProviderConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the provider connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcGetProviderConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArcGetProviderConnectionResponse**](ArcGetProviderConnectionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcListProviderConnections

List provider connections



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
    includeDeleted := true // bool | Include deleted provider connections in the results. Deleted connections have a `deleted_at` value. (optional) (default to false)
    sort := "created_at" // string | Sort field. Prefix with `-` for descending order (e.g., `sort=-name`).  (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlProviderConnectionsAPI.ArcListProviderConnections(ctx).IncludeDeleted(includeDeleted).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProviderConnectionsAPI.ArcListProviderConnections`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcListProviderConnections`: ArcProviderConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProviderConnectionsAPI.ArcListProviderConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArcListProviderConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeleted** | **bool** | Include deleted provider connections in the results. Deleted connections have a `deleted_at` value. | [default to false] **sort** | **string** | Sort field. Prefix with `-` for descending order (e.g., `sort&#x3D;-name`).  | [default to &quot;created_at&quot;]

### Return type

[**ArcProviderConnectionsResponse**](ArcProviderConnectionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcUpdateProviderConnection

Update a provider connection



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
    id := "pTfnUUBcU1yv8SoRJn5PZ2" // string | The ID of the provider connection.
    arcUpdateProviderConnectionRequest := *openapiclient.NewArcUpdateProviderConnectionRequest() // ArcUpdateProviderConnectionRequest | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlProviderConnectionsAPI.ArcUpdateProviderConnection(ctx, id).ArcUpdateProviderConnectionRequest(arcUpdateProviderConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProviderConnectionsAPI.ArcUpdateProviderConnection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcUpdateProviderConnection`: ArcProviderConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProviderConnectionsAPI.ArcUpdateProviderConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the provider connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcUpdateProviderConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcUpdateProviderConnectionRequest** | [**ArcUpdateProviderConnectionRequest**](ArcUpdateProviderConnectionRequest.md) |  | 

### Return type

[**ArcProviderConnectionResponse**](ArcProviderConnectionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

