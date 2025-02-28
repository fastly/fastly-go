# ProductDdosProtectionAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductDdosProtection**](ProductDdosProtectionAPI.md#DisableProductDdosProtection) | **DELETE** `/enabled-products/v1/ddos_protection/services/{service_id}` | Disable product
[**EnableProductDdosProtection**](ProductDdosProtectionAPI.md#EnableProductDdosProtection) | **PUT** `/enabled-products/v1/ddos_protection/services/{service_id}` | Enable product
[**GetProductDdosProtection**](ProductDdosProtectionAPI.md#GetProductDdosProtection) | **GET** `/enabled-products/v1/ddos_protection/services/{service_id}` | Get product enablement status
[**GetProductDdosProtectionConfiguration**](ProductDdosProtectionAPI.md#GetProductDdosProtectionConfiguration) | **GET** `/enabled-products/v1/ddos_protection/services/{service_id}/configuration` | Get configuration
[**SetProductDdosProtectionConfiguration**](ProductDdosProtectionAPI.md#SetProductDdosProtectionConfiguration) | **PATCH** `/enabled-products/v1/ddos_protection/services/{service_id}/configuration` | Update configuration



## DisableProductDdosProtection

Disable product



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductDdosProtectionAPI.DisableProductDdosProtection(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductDdosProtectionAPI.DisableProductDdosProtection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProductDdosProtectionRequest struct via the builder pattern


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


## EnableProductDdosProtection

Enable product



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductDdosProtectionAPI.EnableProductDdosProtection(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductDdosProtectionAPI.EnableProductDdosProtection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableProductDdosProtection`: DdosProtectionResponseEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductDdosProtectionAPI.EnableProductDdosProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProductDdosProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosProtectionResponseEnable**](DdosProtectionResponseEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductDdosProtection

Get product enablement status



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductDdosProtectionAPI.GetProductDdosProtection(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductDdosProtectionAPI.GetProductDdosProtection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductDdosProtection`: DdosProtectionResponseEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductDdosProtectionAPI.GetProductDdosProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductDdosProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosProtectionResponseEnable**](DdosProtectionResponseEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductDdosProtectionConfiguration

Get configuration



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductDdosProtectionAPI.GetProductDdosProtectionConfiguration(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductDdosProtectionAPI.GetProductDdosProtectionConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductDdosProtectionConfiguration`: DdosProtectionResponseConfigure
    fmt.Fprintf(os.Stdout, "Response from `ProductDdosProtectionAPI.GetProductDdosProtectionConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductDdosProtectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosProtectionResponseConfigure**](DdosProtectionResponseConfigure.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SetProductDdosProtectionConfiguration

Update configuration



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
    ddosProtectionRequestUpdateConfiguration := *openapiclient.NewDdosProtectionRequestUpdateConfiguration("Mode_example") // DdosProtectionRequestUpdateConfiguration |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductDdosProtectionAPI.SetProductDdosProtectionConfiguration(ctx, serviceID).DdosProtectionRequestUpdateConfiguration(ddosProtectionRequestUpdateConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductDdosProtectionAPI.SetProductDdosProtectionConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProductDdosProtectionConfiguration`: DdosProtectionResponseConfigure
    fmt.Fprintf(os.Stdout, "Response from `ProductDdosProtectionAPI.SetProductDdosProtectionConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProductDdosProtectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ddosProtectionRequestUpdateConfiguration** | [**DdosProtectionRequestUpdateConfiguration**](DdosProtectionRequestUpdateConfiguration.md) |  | 

### Return type

[**DdosProtectionResponseConfigure**](DdosProtectionResponseConfigure.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
