# ServerAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePoolServer**](ServerAPI.md#CreatePoolServer) | **POST** `/service/{service_id}/pool/{pool_id}/server` | Add a server to a pool
[**DeletePoolServer**](ServerAPI.md#DeletePoolServer) | **DELETE** `/service/{service_id}/pool/{pool_id}/server/{server_id}` | Delete a server from a pool
[**GetPoolServer**](ServerAPI.md#GetPoolServer) | **GET** `/service/{service_id}/pool/{pool_id}/server/{server_id}` | Get a pool server
[**ListPoolServers**](ServerAPI.md#ListPoolServers) | **GET** `/service/{service_id}/pool/{pool_id}/servers` | List servers in a pool
[**UpdatePoolServer**](ServerAPI.md#UpdatePoolServer) | **PUT** `/service/{service_id}/pool/{pool_id}/server/{server_id}` | Update a server



## CreatePoolServer

Add a server to a pool



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    poolID := "poolId_example" // string | Alphanumeric string identifying a Pool.
    weight := int32(56) // int32 | Weight (`1-100`) used to load balance this server against others. (optional) (default to 100)
    maxConn := int32(56) // int32 | Maximum number of connections. If the value is `0`, it inherits the value from pool's `max_conn_default`. (optional) (default to 0)
    port := int32(56) // int32 | Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS. (optional) (default to 80)
    address := "address_example" // string | A hostname, IPv4, or IPv6 address for the server. Required. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    disabled := true // bool | Allows servers to be enabled and disabled in a pool. (optional) (default to false)
    overrideHost := "overrideHost_example" // string | The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting. (optional) (default to "null")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServerAPI.CreatePoolServer(ctx, serviceID, poolID).Weight(weight).MaxConn(maxConn).Port(port).Address(address).Comment(comment).Disabled(disabled).OverrideHost(overrideHost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.CreatePoolServer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePoolServer`: ServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerAPI.CreatePoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**poolID** | **string** | Alphanumeric string identifying a Pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **weight** | **int32** | Weight (`1-100`) used to load balance this server against others. | [default to 100] **maxConn** | **int32** | Maximum number of connections. If the value is `0`, it inherits the value from pool&#39;s `max_conn_default`. | [default to 0] **port** | **int32** | Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS. | [default to 80] **address** | **string** | A hostname, IPv4, or IPv6 address for the server. Required. |  **comment** | **string** | A freeform descriptive note. |  **disabled** | **bool** | Allows servers to be enabled and disabled in a pool. | [default to false] **overrideHost** | **string** | The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting. | [default to &quot;null&quot;]

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeletePoolServer

Delete a server from a pool



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    poolID := "poolId_example" // string | Alphanumeric string identifying a Pool.
    serverID := "serverId_example" // string | Alphanumeric string identifying a Server.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServerAPI.DeletePoolServer(ctx, serviceID, poolID, serverID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.DeletePoolServer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePoolServer`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ServerAPI.DeletePoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**poolID** | **string** | Alphanumeric string identifying a Pool. | 
**serverID** | **string** | Alphanumeric string identifying a Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetPoolServer

Get a pool server



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    poolID := "poolId_example" // string | Alphanumeric string identifying a Pool.
    serverID := "serverId_example" // string | Alphanumeric string identifying a Server.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServerAPI.GetPoolServer(ctx, serviceID, poolID, serverID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetPoolServer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoolServer`: ServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**poolID** | **string** | Alphanumeric string identifying a Pool. | 
**serverID** | **string** | Alphanumeric string identifying a Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListPoolServers

List servers in a pool



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    poolID := "poolId_example" // string | Alphanumeric string identifying a Pool.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServerAPI.ListPoolServers(ctx, serviceID, poolID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.ListPoolServers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPoolServers`: []ServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerAPI.ListPoolServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**poolID** | **string** | Alphanumeric string identifying a Pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPoolServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ServerResponse**](ServerResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdatePoolServer

Update a server



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    poolID := "poolId_example" // string | Alphanumeric string identifying a Pool.
    serverID := "serverId_example" // string | Alphanumeric string identifying a Server.
    weight := int32(56) // int32 | Weight (`1-100`) used to load balance this server against others. (optional) (default to 100)
    maxConn := int32(56) // int32 | Maximum number of connections. If the value is `0`, it inherits the value from pool's `max_conn_default`. (optional) (default to 0)
    port := int32(56) // int32 | Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS. (optional) (default to 80)
    address := "address_example" // string | A hostname, IPv4, or IPv6 address for the server. Required. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    disabled := true // bool | Allows servers to be enabled and disabled in a pool. (optional) (default to false)
    overrideHost := "overrideHost_example" // string | The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting. (optional) (default to "null")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServerAPI.UpdatePoolServer(ctx, serviceID, poolID, serverID).Weight(weight).MaxConn(maxConn).Port(port).Address(address).Comment(comment).Disabled(disabled).OverrideHost(overrideHost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.UpdatePoolServer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePoolServer`: ServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerAPI.UpdatePoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**poolID** | **string** | Alphanumeric string identifying a Pool. | 
**serverID** | **string** | Alphanumeric string identifying a Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **weight** | **int32** | Weight (`1-100`) used to load balance this server against others. | [default to 100] **maxConn** | **int32** | Maximum number of connections. If the value is `0`, it inherits the value from pool&#39;s `max_conn_default`. | [default to 0] **port** | **int32** | Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS. | [default to 80] **address** | **string** | A hostname, IPv4, or IPv6 address for the server. Required. |  **comment** | **string** | A freeform descriptive note. |  **disabled** | **bool** | Allows servers to be enabled and disabled in a pool. | [default to false] **overrideHost** | **string** | The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting. | [default to &quot;null&quot;]

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
