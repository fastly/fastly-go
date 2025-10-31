# TlsConfigurationsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTlsConfig**](TlsConfigurationsAPI.md#GetTlsConfig) | **GET** `/tls/configurations/{tls_configuration_id}` | Get a TLS configuration
[**ListTlsConfigs**](TlsConfigurationsAPI.md#ListTlsConfigs) | **GET** `/tls/configurations` | List TLS configurations
[**UpdateTlsConfig**](TlsConfigurationsAPI.md#UpdateTlsConfig) | **PATCH** `/tls/configurations/{tls_configuration_id}` | Update a TLS configuration



## GetTlsConfig

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
    tlsConfigurationId := "tlsConfigurationId_example" // string | Alphanumeric string identifying a TLS configuration.
    include := "dns_records" // string | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsConfigurationsAPI.GetTlsConfig(ctx, tlsConfigurationId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsConfigurationsAPI.GetTlsConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTlsConfig`: TlsConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsConfigurationsAPI.GetTlsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsConfigurationId** | **string** | Alphanumeric string identifying a TLS configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  | 

### Return type

[**TlsConfigurationResponse**](TlsConfigurationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTlsConfigs

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
    resp, r, err := apiClient.TlsConfigurationsAPI.ListTlsConfigs(ctx).FilterBulk(filterBulk).Include(include).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsConfigurationsAPI.ListTlsConfigs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsConfigs`: TlsConfigurationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsConfigurationsAPI.ListTlsConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBulk** | **string** | Optionally filters by the bulk attribute. |  **include** | **string** | Include related objects. Optional, comma-separated values. Permitted values: `dns_records`.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**TlsConfigurationsResponse**](TlsConfigurationsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateTlsConfig

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
    tlsConfigurationId := "tlsConfigurationId_example" // string | Alphanumeric string identifying a TLS configuration.
    tlsConfiguration := *openapiclient.NewTlsConfiguration() // TlsConfiguration |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TlsConfigurationsAPI.UpdateTlsConfig(ctx, tlsConfigurationId).TlsConfiguration(tlsConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TlsConfigurationsAPI.UpdateTlsConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTlsConfig`: TlsConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `TlsConfigurationsAPI.UpdateTlsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tlsConfigurationId** | **string** | Alphanumeric string identifying a TLS configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTlsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsConfiguration** | [**TlsConfiguration**](TlsConfiguration.md) |  | 

### Return type

[**TlsConfigurationResponse**](TlsConfigurationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

