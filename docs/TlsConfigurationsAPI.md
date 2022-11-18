# TLSConfigurationsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTLSConfig**](TlsConfigurationsAPI.md#GetTLSConfig) | **GET** `/tls/configurations/{tls_configuration_id}` | Get a TLS configuration
[**ListTLSConfigs**](TlsConfigurationsAPI.md#ListTLSConfigs) | **GET** `/tls/configurations` | List TLS configurations
[**UpdateTLSConfig**](TlsConfigurationsAPI.md#UpdateTLSConfig) | **PATCH** `/tls/configurations/{tls_configuration_id}` | Update a TLS configuration



## GetTLSConfig

Get a TLS configuration



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
    tlsConfigurationID := "tlsConfigurationId_example" // string | Alphanumeric string identifying a TLS configuration.
    include := "dns_records" // string | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSConfigurationsAPI.GetTLSConfig(ctx, tlsConfigurationID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSConfigurationsAPI.GetTLSConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLSConfig`: TLSConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSConfigurationsAPI.GetTLSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsConfigurationID** | **string** | Alphanumeric string identifying a TLS configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  | 

### Return type

[**TLSConfigurationResponse**](TlsConfigurationResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTLSConfigs

List TLS configurations



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
    filterBulk := "filterBulk_example" // string | Optionally filters by the bulk attribute. (optional)
    include := "dns_records" // string | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSConfigurationsAPI.ListTLSConfigs(ctx).FilterBulk(filterBulk).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSConfigurationsAPI.ListTLSConfigs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTLSConfigs`: TLSConfigurationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSConfigurationsAPI.ListTLSConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTLSConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBulk** | **string** | Optionally filters by the bulk attribute. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TLSConfigurationsResponse**](TlsConfigurationsResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateTLSConfig

Update a TLS configuration



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
    tlsConfigurationID := "tlsConfigurationId_example" // string | Alphanumeric string identifying a TLS configuration.
    tlsConfiguration := *openapiclient.NewTLSConfiguration() // TLSConfiguration |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSConfigurationsAPI.UpdateTLSConfig(ctx, tlsConfigurationID).TLSConfiguration(tlsConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSConfigurationsAPI.UpdateTLSConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTLSConfig`: TLSConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSConfigurationsAPI.UpdateTLSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsConfigurationID** | **string** | Alphanumeric string identifying a TLS configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTLSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsConfiguration** | [**TLSConfiguration**](TlsConfiguration.md) |  | 

### Return type

[**TLSConfigurationResponse**](TlsConfigurationResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
