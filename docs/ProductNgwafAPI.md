# ProductNgwafAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProductNgwaf**](ProductNgwafAPI.md#DisableProductNgwaf) | **DELETE** `/enabled-products/v1/ngwaf/services/{service_id}` | Disable product
[**EnableProductNgwaf**](ProductNgwafAPI.md#EnableProductNgwaf) | **PUT** `/enabled-products/v1/ngwaf/services/{service_id}` | Enable product
[**GetProductNgwaf**](ProductNgwafAPI.md#GetProductNgwaf) | **GET** `/enabled-products/v1/ngwaf/services/{service_id}` | Get product enablement status
[**GetProductNgwafConfiguration**](ProductNgwafAPI.md#GetProductNgwafConfiguration) | **GET** `/enabled-products/v1/ngwaf/services/{service_id}/configuration` | Get configuration
[**GetServicesProductNgwaf**](ProductNgwafAPI.md#GetServicesProductNgwaf) | **GET** `/enabled-products/v1/ngwaf/services` | Get services with product enabled
[**SetProductNgwafConfiguration**](ProductNgwafAPI.md#SetProductNgwafConfiguration) | **PATCH** `/enabled-products/v1/ngwaf/services/{service_id}/configuration` | Update configuration



## DisableProductNgwaf

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
    resp, r, err := apiClient.ProductNgwafAPI.DisableProductNgwaf(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductNgwafAPI.DisableProductNgwaf`: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableProductNgwafRequest struct via the builder pattern


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


## EnableProductNgwaf

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
    ngwafRequestEnable := *openapiclient.NewNgwafRequestEnable("WorkspaceId_example") // NgwafRequestEnable |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductNgwafAPI.EnableProductNgwaf(ctx, serviceID).NgwafRequestEnable(ngwafRequestEnable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductNgwafAPI.EnableProductNgwaf`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableProductNgwaf`: NgwafResponseEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductNgwafAPI.EnableProductNgwaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProductNgwafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ngwafRequestEnable** | [**NgwafRequestEnable**](NgwafRequestEnable.md) |  | 

### Return type

[**NgwafResponseEnable**](NgwafResponseEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductNgwaf

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
    resp, r, err := apiClient.ProductNgwafAPI.GetProductNgwaf(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductNgwafAPI.GetProductNgwaf`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductNgwaf`: NgwafResponseEnable
    fmt.Fprintf(os.Stdout, "Response from `ProductNgwafAPI.GetProductNgwaf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductNgwafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NgwafResponseEnable**](NgwafResponseEnable.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetProductNgwafConfiguration

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
    resp, r, err := apiClient.ProductNgwafAPI.GetProductNgwafConfiguration(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductNgwafAPI.GetProductNgwafConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductNgwafConfiguration`: NgwafResponseConfigure
    fmt.Fprintf(os.Stdout, "Response from `ProductNgwafAPI.GetProductNgwafConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductNgwafConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NgwafResponseConfigure**](NgwafResponseConfigure.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetServicesProductNgwaf

Get services with product enabled



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
    resp, r, err := apiClient.ProductNgwafAPI.GetServicesProductNgwaf(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductNgwafAPI.GetServicesProductNgwaf`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesProductNgwaf`: NgwafResponseBodyGetAllServices
    fmt.Fprintf(os.Stdout, "Response from `ProductNgwafAPI.GetServicesProductNgwaf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesProductNgwafRequest struct via the builder pattern



### Return type

[**NgwafResponseBodyGetAllServices**](NgwafResponseBodyGetAllServices.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SetProductNgwafConfiguration

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
    ngwafRequestUpdateConfiguration := *openapiclient.NewNgwafRequestUpdateConfiguration() // NgwafRequestUpdateConfiguration |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ProductNgwafAPI.SetProductNgwafConfiguration(ctx, serviceID).NgwafRequestUpdateConfiguration(ngwafRequestUpdateConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductNgwafAPI.SetProductNgwafConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProductNgwafConfiguration`: NgwafResponseConfigure
    fmt.Fprintf(os.Stdout, "Response from `ProductNgwafAPI.SetProductNgwafConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProductNgwafConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ngwafRequestUpdateConfiguration** | [**NgwafRequestUpdateConfiguration**](NgwafRequestUpdateConfiguration.md) |  | 

### Return type

[**NgwafResponseConfigure**](NgwafResponseConfigure.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
