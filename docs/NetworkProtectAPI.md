# NetworkProtectAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkProtectEndpoint**](NetworkProtectAPI.md#CreateNetworkProtectEndpoint) | **POST** `/network-protect/v1/endpoints` | Create a Network Protect endpoint
[**CreateNetworkProtectNetwork**](NetworkProtectAPI.md#CreateNetworkProtectNetwork) | **POST** `/network-protect/v1/networks` | Create a Network Protect network
[**CreateNetworkProtectPrefixList**](NetworkProtectAPI.md#CreateNetworkProtectPrefixList) | **POST** `/network-protect/v1/prefix-lists` | Create a Network Protect prefix list
[**CreateNetworkProtectRuleSet**](NetworkProtectAPI.md#CreateNetworkProtectRuleSet) | **POST** `/network-protect/v1/rule-sets` | Create a Network Protect rule set
[**DeleteNetworkProtectEndpoint**](NetworkProtectAPI.md#DeleteNetworkProtectEndpoint) | **DELETE** `/network-protect/v1/endpoints/{endpoint_id}` | Delete a Network Protect endpoint
[**DeleteNetworkProtectNetwork**](NetworkProtectAPI.md#DeleteNetworkProtectNetwork) | **DELETE** `/network-protect/v1/networks/{network_id}` | Delete a Network Protect network
[**DeleteNetworkProtectPrefixList**](NetworkProtectAPI.md#DeleteNetworkProtectPrefixList) | **DELETE** `/network-protect/v1/prefix-lists/{prefix_list_id}` | Delete a Network Protect prefix list
[**DeleteNetworkProtectRuleSet**](NetworkProtectAPI.md#DeleteNetworkProtectRuleSet) | **DELETE** `/network-protect/v1/rule-sets/{rule_set_id}` | Delete a Network Protect rule set
[**GetNetworkProtectEndpoint**](NetworkProtectAPI.md#GetNetworkProtectEndpoint) | **GET** `/network-protect/v1/endpoints/{endpoint_id}` | Get a Network Protect endpoint
[**GetNetworkProtectNetwork**](NetworkProtectAPI.md#GetNetworkProtectNetwork) | **GET** `/network-protect/v1/networks/{network_id}` | Get a Network Protect network
[**GetNetworkProtectPrefixList**](NetworkProtectAPI.md#GetNetworkProtectPrefixList) | **GET** `/network-protect/v1/prefix-lists/{prefix_list_id}` | Get a Network Protect prefix list
[**GetNetworkProtectRuleSet**](NetworkProtectAPI.md#GetNetworkProtectRuleSet) | **GET** `/network-protect/v1/rule-sets/{rule_set_id}` | Get a Network Protect rule set
[**ListNetworkProtectEndpoints**](NetworkProtectAPI.md#ListNetworkProtectEndpoints) | **GET** `/network-protect/v1/endpoints` | List Network Protect endpoints
[**ListNetworkProtectNetworks**](NetworkProtectAPI.md#ListNetworkProtectNetworks) | **GET** `/network-protect/v1/networks` | List Network Protect networks
[**ListNetworkProtectPrefixLists**](NetworkProtectAPI.md#ListNetworkProtectPrefixLists) | **GET** `/network-protect/v1/prefix-lists` | List Network Protect prefix lists
[**ListNetworkProtectRuleSets**](NetworkProtectAPI.md#ListNetworkProtectRuleSets) | **GET** `/network-protect/v1/rule-sets` | List Network Protect rule sets
[**UpdateNetworkProtectEndpoint**](NetworkProtectAPI.md#UpdateNetworkProtectEndpoint) | **PATCH** `/network-protect/v1/endpoints/{endpoint_id}` | Update a Network Protect endpoint
[**UpdateNetworkProtectNetwork**](NetworkProtectAPI.md#UpdateNetworkProtectNetwork) | **PATCH** `/network-protect/v1/networks/{network_id}` | Update a Network Protect network
[**UpdateNetworkProtectPrefixList**](NetworkProtectAPI.md#UpdateNetworkProtectPrefixList) | **PATCH** `/network-protect/v1/prefix-lists/{prefix_list_id}` | Update a Network Protect prefix list
[**UpdateNetworkProtectRuleSet**](NetworkProtectAPI.md#UpdateNetworkProtectRuleSet) | **PATCH** `/network-protect/v1/rule-sets/{rule_set_id}` | Update a Network Protect rule set



## CreateNetworkProtectEndpoint

Create a Network Protect endpoint



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
    networkProtectEndpointCreate := *openapiclient.NewNetworkProtectEndpointCreate("Transit network", "192.0.2.10") // NetworkProtectEndpointCreate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.CreateNetworkProtectEndpoint(ctx).NetworkProtectEndpointCreate(networkProtectEndpointCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.CreateNetworkProtectEndpoint`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkProtectEndpoint`: NetworkProtectEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.CreateNetworkProtectEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProtectEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectEndpointCreate** | [**NetworkProtectEndpointCreate**](NetworkProtectEndpointCreate.md) |  | 

### Return type

[**NetworkProtectEndpointResponse**](NetworkProtectEndpointResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateNetworkProtectNetwork

Create a Network Protect network



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
    networkProtectNetworkCreate := *openapiclient.NewNetworkProtectNetworkCreate("Transit network", "198.51.100.0/24") // NetworkProtectNetworkCreate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.CreateNetworkProtectNetwork(ctx).NetworkProtectNetworkCreate(networkProtectNetworkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.CreateNetworkProtectNetwork`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkProtectNetwork`: NetworkProtectNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.CreateNetworkProtectNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProtectNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectNetworkCreate** | [**NetworkProtectNetworkCreate**](NetworkProtectNetworkCreate.md) |  | 

### Return type

[**NetworkProtectNetworkResponse**](NetworkProtectNetworkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateNetworkProtectPrefixList

Create a Network Protect prefix list



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
    networkProtectPrefixListCreate := *openapiclient.NewNetworkProtectPrefixListCreate("Transit network", openapiclient.network_protect_prefix_list_family("v4"), []string{"198.51.100.0/24"}) // NetworkProtectPrefixListCreate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.CreateNetworkProtectPrefixList(ctx).NetworkProtectPrefixListCreate(networkProtectPrefixListCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.CreateNetworkProtectPrefixList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkProtectPrefixList`: NetworkProtectPrefixListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.CreateNetworkProtectPrefixList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProtectPrefixListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectPrefixListCreate** | [**NetworkProtectPrefixListCreate**](NetworkProtectPrefixListCreate.md) |  | 

### Return type

[**NetworkProtectPrefixListResponse**](NetworkProtectPrefixListResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateNetworkProtectRuleSet

Create a Network Protect rule set



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
    networkProtectRuleSetCreate := *openapiclient.NewNetworkProtectRuleSetCreate(openapiclient.network_protect_rule_set_default_action("allow")) // NetworkProtectRuleSetCreate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.CreateNetworkProtectRuleSet(ctx).NetworkProtectRuleSetCreate(networkProtectRuleSetCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.CreateNetworkProtectRuleSet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkProtectRuleSet`: NetworkProtectRuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.CreateNetworkProtectRuleSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProtectRuleSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectRuleSetCreate** | [**NetworkProtectRuleSetCreate**](NetworkProtectRuleSetCreate.md) |  | 

### Return type

[**NetworkProtectRuleSetResponse**](NetworkProtectRuleSetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteNetworkProtectEndpoint

Delete a Network Protect endpoint



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
    endpointId := "DBOEWyMXebAir5qwjWEwK1" // string | Unique identifier of the Network Protect endpoint.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.DeleteNetworkProtectEndpoint(ctx, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.DeleteNetworkProtectEndpoint`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** | Unique identifier of the Network Protect endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProtectEndpointRequest struct via the builder pattern


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


## DeleteNetworkProtectNetwork

Delete a Network Protect network



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
    networkId := "an7NWnhOjjOkwSNvrwphTd" // string | Unique identifier of the Network Protect network.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.DeleteNetworkProtectNetwork(ctx, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.DeleteNetworkProtectNetwork`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Unique identifier of the Network Protect network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProtectNetworkRequest struct via the builder pattern


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


## DeleteNetworkProtectPrefixList

Delete a Network Protect prefix list



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
    prefixListId := "QoavyX1BHOMGQm2zmVyEbS" // string | Unique identifier of the Network Protect prefix list.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.DeleteNetworkProtectPrefixList(ctx, prefixListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.DeleteNetworkProtectPrefixList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefixListId** | **string** | Unique identifier of the Network Protect prefix list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProtectPrefixListRequest struct via the builder pattern


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


## DeleteNetworkProtectRuleSet

Delete a Network Protect rule set



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
    ruleSetId := "Z4xQoBznwNb5GJfdJbp58Y" // string | Unique identifier of the Network Protect rule set.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.DeleteNetworkProtectRuleSet(ctx, ruleSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.DeleteNetworkProtectRuleSet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleSetId** | **string** | Unique identifier of the Network Protect rule set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProtectRuleSetRequest struct via the builder pattern


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


## GetNetworkProtectEndpoint

Get a Network Protect endpoint



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
    endpointId := "DBOEWyMXebAir5qwjWEwK1" // string | Unique identifier of the Network Protect endpoint.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.GetNetworkProtectEndpoint(ctx, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.GetNetworkProtectEndpoint`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProtectEndpoint`: NetworkProtectEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.GetNetworkProtectEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** | Unique identifier of the Network Protect endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProtectEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkProtectEndpointResponse**](NetworkProtectEndpointResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetNetworkProtectNetwork

Get a Network Protect network



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
    networkId := "an7NWnhOjjOkwSNvrwphTd" // string | Unique identifier of the Network Protect network.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.GetNetworkProtectNetwork(ctx, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.GetNetworkProtectNetwork`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProtectNetwork`: NetworkProtectNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.GetNetworkProtectNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Unique identifier of the Network Protect network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProtectNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkProtectNetworkResponse**](NetworkProtectNetworkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetNetworkProtectPrefixList

Get a Network Protect prefix list



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
    prefixListId := "QoavyX1BHOMGQm2zmVyEbS" // string | Unique identifier of the Network Protect prefix list.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.GetNetworkProtectPrefixList(ctx, prefixListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.GetNetworkProtectPrefixList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProtectPrefixList`: NetworkProtectPrefixListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.GetNetworkProtectPrefixList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefixListId** | **string** | Unique identifier of the Network Protect prefix list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProtectPrefixListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkProtectPrefixListResponse**](NetworkProtectPrefixListResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetNetworkProtectRuleSet

Get a Network Protect rule set



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
    ruleSetId := "Z4xQoBznwNb5GJfdJbp58Y" // string | Unique identifier of the Network Protect rule set.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.GetNetworkProtectRuleSet(ctx, ruleSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.GetNetworkProtectRuleSet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProtectRuleSet`: NetworkProtectRuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.GetNetworkProtectRuleSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleSetId** | **string** | Unique identifier of the Network Protect rule set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProtectRuleSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkProtectRuleSetResponse**](NetworkProtectRuleSetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListNetworkProtectEndpoints

List Network Protect endpoints



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
    limit := int64(100) // int64 | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. (optional) (default to 100)
    cursor := "network-cursor-next" // string | Continuation value from a previous response's meta.next_cursor. Omit this parameter to request the first page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.ListNetworkProtectEndpoints(ctx).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.ListNetworkProtectEndpoints`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkProtectEndpoints`: NetworkProtectEndpointsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.ListNetworkProtectEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkProtectEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. | [default to 100] **cursor** | **string** | Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page. | 

### Return type

[**NetworkProtectEndpointsResponse**](NetworkProtectEndpointsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListNetworkProtectNetworks

List Network Protect networks



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
    limit := int64(100) // int64 | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. (optional) (default to 100)
    cursor := "network-cursor-next" // string | Continuation value from a previous response's meta.next_cursor. Omit this parameter to request the first page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.ListNetworkProtectNetworks(ctx).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.ListNetworkProtectNetworks`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkProtectNetworks`: NetworkProtectNetworksResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.ListNetworkProtectNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkProtectNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. | [default to 100] **cursor** | **string** | Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page. | 

### Return type

[**NetworkProtectNetworksResponse**](NetworkProtectNetworksResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListNetworkProtectPrefixLists

List Network Protect prefix lists



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
    limit := int64(100) // int64 | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. (optional) (default to 100)
    cursor := "network-cursor-next" // string | Continuation value from a previous response's meta.next_cursor. Omit this parameter to request the first page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.ListNetworkProtectPrefixLists(ctx).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.ListNetworkProtectPrefixLists`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkProtectPrefixLists`: NetworkProtectPrefixListsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.ListNetworkProtectPrefixLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkProtectPrefixListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. | [default to 100] **cursor** | **string** | Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page. | 

### Return type

[**NetworkProtectPrefixListsResponse**](NetworkProtectPrefixListsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListNetworkProtectRuleSets

List Network Protect rule sets



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
    limit := int64(100) // int64 | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. (optional) (default to 100)
    cursor := "network-cursor-next" // string | Continuation value from a previous response's meta.next_cursor. Omit this parameter to request the first page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.ListNetworkProtectRuleSets(ctx).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.ListNetworkProtectRuleSets`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkProtectRuleSets`: NetworkProtectRuleSetsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.ListNetworkProtectRuleSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkProtectRuleSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a `400 Bad Request` error. | [default to 100] **cursor** | **string** | Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page. | 

### Return type

[**NetworkProtectRuleSetsResponse**](NetworkProtectRuleSetsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateNetworkProtectEndpoint

Update a Network Protect endpoint



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
    endpointId := "DBOEWyMXebAir5qwjWEwK1" // string | Unique identifier of the Network Protect endpoint.
    networkProtectEndpointUpdate := *openapiclient.NewNetworkProtectEndpointUpdate("Transit network") // NetworkProtectEndpointUpdate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.UpdateNetworkProtectEndpoint(ctx, endpointId).NetworkProtectEndpointUpdate(networkProtectEndpointUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.UpdateNetworkProtectEndpoint`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkProtectEndpoint`: NetworkProtectEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.UpdateNetworkProtectEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** | Unique identifier of the Network Protect endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProtectEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectEndpointUpdate** | [**NetworkProtectEndpointUpdate**](NetworkProtectEndpointUpdate.md) |  | 

### Return type

[**NetworkProtectEndpointResponse**](NetworkProtectEndpointResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateNetworkProtectNetwork

Update a Network Protect network



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
    networkId := "an7NWnhOjjOkwSNvrwphTd" // string | Unique identifier of the Network Protect network.
    networkProtectNetworkUpdate := *openapiclient.NewNetworkProtectNetworkUpdate() // NetworkProtectNetworkUpdate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.UpdateNetworkProtectNetwork(ctx, networkId).NetworkProtectNetworkUpdate(networkProtectNetworkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.UpdateNetworkProtectNetwork`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkProtectNetwork`: NetworkProtectNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.UpdateNetworkProtectNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Unique identifier of the Network Protect network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProtectNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectNetworkUpdate** | [**NetworkProtectNetworkUpdate**](NetworkProtectNetworkUpdate.md) |  | 

### Return type

[**NetworkProtectNetworkResponse**](NetworkProtectNetworkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateNetworkProtectPrefixList

Update a Network Protect prefix list



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
    prefixListId := "QoavyX1BHOMGQm2zmVyEbS" // string | Unique identifier of the Network Protect prefix list.
    networkProtectPrefixListUpdate := *openapiclient.NewNetworkProtectPrefixListUpdate() // NetworkProtectPrefixListUpdate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.UpdateNetworkProtectPrefixList(ctx, prefixListId).NetworkProtectPrefixListUpdate(networkProtectPrefixListUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.UpdateNetworkProtectPrefixList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkProtectPrefixList`: NetworkProtectPrefixListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.UpdateNetworkProtectPrefixList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefixListId** | **string** | Unique identifier of the Network Protect prefix list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProtectPrefixListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectPrefixListUpdate** | [**NetworkProtectPrefixListUpdate**](NetworkProtectPrefixListUpdate.md) |  | 

### Return type

[**NetworkProtectPrefixListResponse**](NetworkProtectPrefixListResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateNetworkProtectRuleSet

Update a Network Protect rule set



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
    ruleSetId := "Z4xQoBznwNb5GJfdJbp58Y" // string | Unique identifier of the Network Protect rule set.
    networkProtectRuleSetUpdate := *openapiclient.NewNetworkProtectRuleSetUpdate() // NetworkProtectRuleSetUpdate | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NetworkProtectAPI.UpdateNetworkProtectRuleSet(ctx, ruleSetId).NetworkProtectRuleSetUpdate(networkProtectRuleSetUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkProtectAPI.UpdateNetworkProtectRuleSet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkProtectRuleSet`: NetworkProtectRuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkProtectAPI.UpdateNetworkProtectRuleSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleSetId** | **string** | Unique identifier of the Network Protect rule set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProtectRuleSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkProtectRuleSetUpdate** | [**NetworkProtectRuleSetUpdate**](NetworkProtectRuleSetUpdate.md) |  | 

### Return type

[**NetworkProtectRuleSetResponse**](NetworkProtectRuleSetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

