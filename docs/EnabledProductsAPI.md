# EnabledProductsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProduct**](EnabledProductsAPI.md#DisableProduct) | **DELETE** `/enabled-products/v1/{product_id}/services/{service_id}` | Disable a product
[**EnableProduct**](EnabledProductsAPI.md#EnableProduct) | **PUT** `/enabled-products/v1/{product_id}/services/{service_id}` | Enable a product
[**GetEnabledProduct**](EnabledProductsAPI.md#GetEnabledProduct) | **GET** `/enabled-products/v1/{product_id}/services/{service_id}` | Get enabled product
[**GetProductConfiguration**](EnabledProductsAPI.md#GetProductConfiguration) | **GET** `/enabled-products/v1/{product_id}/services/{service_id}/configuration` | Get configuration for a product
[**SetProductConfiguration**](EnabledProductsAPI.md#SetProductConfiguration) | **PATCH** `/enabled-products/v1/{product_id}/services/{service_id}/configuration` | Update configuration for a product



## DisableProduct

Disable a product



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
    productID := "ngwaf" // string | 
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EnabledProductsAPI.DisableProduct(ctx, productID, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnabledProductsAPI.DisableProduct`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productID** | **string** |  | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductRequest struct via the builder pattern


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


## EnableProduct

Enable a product



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
    productID := "ngwaf" // string | 
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    setWorkspaceID := *openapiclient.NewSetWorkspaceID() // SetWorkspaceID |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EnabledProductsAPI.EnableProduct(ctx, productID, serviceID).SetWorkspaceID(setWorkspaceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnabledProductsAPI.EnableProduct`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableProduct`: EnabledProductResponse
    fmt.Fprintf(os.Stdout, "Response from `EnabledProductsAPI.EnableProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productID** | **string** |  | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setWorkspaceID** | [**SetWorkspaceID**](SetWorkspaceID.md) |  | 

### Return type

[**EnabledProductResponse**](EnabledProductResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetEnabledProduct

Get enabled product



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
    productID := "ngwaf" // string | 
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EnabledProductsAPI.GetEnabledProduct(ctx, productID, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnabledProductsAPI.GetEnabledProduct`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnabledProduct`: EnabledProductResponse
    fmt.Fprintf(os.Stdout, "Response from `EnabledProductsAPI.GetEnabledProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productID** | **string** |  | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnabledProductResponse**](EnabledProductResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductConfiguration

Get configuration for a product



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
    productID := "ngwaf" // string | 
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EnabledProductsAPI.GetProductConfiguration(ctx, productID, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnabledProductsAPI.GetProductConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductConfiguration`: ConfiguredProductResponse
    fmt.Fprintf(os.Stdout, "Response from `EnabledProductsAPI.GetProductConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productID** | **string** |  | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfiguredProductResponse**](ConfiguredProductResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SetProductConfiguration

Update configuration for a product



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
    productID := "ngwaf" // string | 
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    setConfiguration := *openapiclient.NewSetConfiguration() // SetConfiguration |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EnabledProductsAPI.SetProductConfiguration(ctx, productID, serviceID).SetConfiguration(setConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnabledProductsAPI.SetProductConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProductConfiguration`: ConfiguredProductResponse
    fmt.Fprintf(os.Stdout, "Response from `EnabledProductsAPI.SetProductConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productID** | **string** |  | 
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProductConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setConfiguration** | [**SetConfiguration**](SetConfiguration.md) |  | 

### Return type

[**ConfiguredProductResponse**](ConfiguredProductResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
