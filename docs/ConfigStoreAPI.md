# ConfigStoreAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigStore**](ConfigStoreAPI.md#CreateConfigStore) | **POST** `/resources/stores/config` | Create a config store
[**DeleteConfigStore**](ConfigStoreAPI.md#DeleteConfigStore) | **DELETE** `/resources/stores/config/{config_store_id}` | Delete a config store
[**GetConfigStore**](ConfigStoreAPI.md#GetConfigStore) | **GET** `/resources/stores/config/{config_store_id}` | Describe a config store
[**GetConfigStoreInfo**](ConfigStoreAPI.md#GetConfigStoreInfo) | **GET** `/resources/stores/config/{config_store_id}/info` | Get config store metadata
[**ListConfigStoreServices**](ConfigStoreAPI.md#ListConfigStoreServices) | **GET** `/resources/stores/config/{config_store_id}/services` | List linked services
[**ListConfigStores**](ConfigStoreAPI.md#ListConfigStores) | **GET** `/resources/stores/config` | List config stores
[**UpdateConfigStore**](ConfigStoreAPI.md#UpdateConfigStore) | **PUT** `/resources/stores/config/{config_store_id}` | Update a config store



## CreateConfigStore

Create a config store



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
    name := "name_example" // string | The name of the config store. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreAPI.CreateConfigStore(ctx).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.CreateConfigStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigStore`: ConfigStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.CreateConfigStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the config store. | 

### Return type

[**ConfigStoreResponse**](ConfigStoreResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteConfigStore

Delete a config store



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
    configStoreID := "configStoreId_example" // string | An alphanumeric string identifying the config store.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreAPI.DeleteConfigStore(ctx, configStoreID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.DeleteConfigStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfigStore`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.DeleteConfigStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreID** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigStoreRequest struct via the builder pattern


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


## GetConfigStore

Describe a config store



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
    configStoreID := "configStoreId_example" // string | An alphanumeric string identifying the config store.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreAPI.GetConfigStore(ctx, configStoreID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.GetConfigStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigStore`: ConfigStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.GetConfigStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreID** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigStoreResponse**](ConfigStoreResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetConfigStoreInfo

Get config store metadata



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
    configStoreID := "configStoreId_example" // string | An alphanumeric string identifying the config store.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreAPI.GetConfigStoreInfo(ctx, configStoreID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.GetConfigStoreInfo`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigStoreInfo`: ConfigStoreInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.GetConfigStoreInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreID** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigStoreInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigStoreInfoResponse**](ConfigStoreInfoResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListConfigStoreServices

List linked services



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
    configStoreID := "configStoreId_example" // string | An alphanumeric string identifying the config store.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreAPI.ListConfigStoreServices(ctx, configStoreID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.ListConfigStoreServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigStoreServices`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.ListConfigStoreServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreID** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigStoreServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListConfigStores

List config stores



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
    resp, r, err := apiClient.ConfigStoreAPI.ListConfigStores(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.ListConfigStores`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigStores`: []ConfigStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.ListConfigStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigStoresRequest struct via the builder pattern



### Return type

[**[]ConfigStoreResponse**](ConfigStoreResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateConfigStore

Update a config store



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
    configStoreID := "configStoreId_example" // string | An alphanumeric string identifying the config store.
    name := "name_example" // string | The name of the config store. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ConfigStoreAPI.UpdateConfigStore(ctx, configStoreID).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.UpdateConfigStore`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigStore`: ConfigStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.UpdateConfigStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configStoreID** | **string** | An alphanumeric string identifying the config store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the config store. | 

### Return type

[**ConfigStoreResponse**](ConfigStoreResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
