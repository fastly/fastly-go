# LoggingHoneycombAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogHoneycomb**](LoggingHoneycombAPI.md#CreateLogHoneycomb) | **POST** `/service/{service_id}/version/{version_id}/logging/honeycomb` | Create a Honeycomb log endpoint
[**DeleteLogHoneycomb**](LoggingHoneycombAPI.md#DeleteLogHoneycomb) | **DELETE** `/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}` | Delete the Honeycomb log endpoint
[**GetLogHoneycomb**](LoggingHoneycombAPI.md#GetLogHoneycomb) | **GET** `/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}` | Get a Honeycomb log endpoint
[**ListLogHoneycomb**](LoggingHoneycombAPI.md#ListLogHoneycomb) | **GET** `/service/{service_id}/version/{version_id}/logging/honeycomb` | List Honeycomb log endpoints
[**UpdateLogHoneycomb**](LoggingHoneycombAPI.md#UpdateLogHoneycomb) | **PUT** `/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}` | Update a Honeycomb log endpoint



## CreateLogHoneycomb

Create a Honeycomb log endpoint



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest. (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    dataset := "dataset_example" // string | The Honeycomb Dataset you want to log to. (optional)
    token := "token_example" // string | The Write Key from the Account page of your Honeycomb account. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHoneycombAPI.CreateLogHoneycomb(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).Dataset(dataset).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHoneycombAPI.CreateLogHoneycomb`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogHoneycomb`: LoggingHoneycombResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHoneycombAPI.CreateLogHoneycomb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogHoneycombRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest. |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **dataset** | **string** | The Honeycomb Dataset you want to log to. |  **token** | **string** | The Write Key from the Account page of your Honeycomb account. | 

### Return type

[**LoggingHoneycombResponse**](LoggingHoneycombResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogHoneycomb

Delete the Honeycomb log endpoint



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    loggingHoneycombName := "loggingHoneycombName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHoneycombAPI.DeleteLogHoneycomb(ctx, serviceID, versionID, loggingHoneycombName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHoneycombAPI.DeleteLogHoneycomb`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogHoneycomb`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingHoneycombAPI.DeleteLogHoneycomb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingHoneycombName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogHoneycombRequest struct via the builder pattern


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


## GetLogHoneycomb

Get a Honeycomb log endpoint



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    loggingHoneycombName := "loggingHoneycombName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHoneycombAPI.GetLogHoneycomb(ctx, serviceID, versionID, loggingHoneycombName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHoneycombAPI.GetLogHoneycomb`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogHoneycomb`: LoggingHoneycombResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHoneycombAPI.GetLogHoneycomb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingHoneycombName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogHoneycombRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingHoneycombResponse**](LoggingHoneycombResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogHoneycomb

List Honeycomb log endpoints



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
    versionID := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHoneycombAPI.ListLogHoneycomb(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHoneycombAPI.ListLogHoneycomb`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogHoneycomb`: []LoggingHoneycombResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHoneycombAPI.ListLogHoneycomb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogHoneycombRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingHoneycombResponse**](LoggingHoneycombResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogHoneycomb

Update a Honeycomb log endpoint



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    loggingHoneycombName := "loggingHoneycombName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest. (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    dataset := "dataset_example" // string | The Honeycomb Dataset you want to log to. (optional)
    token := "token_example" // string | The Write Key from the Account page of your Honeycomb account. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHoneycombAPI.UpdateLogHoneycomb(ctx, serviceID, versionID, loggingHoneycombName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).Dataset(dataset).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHoneycombAPI.UpdateLogHoneycomb`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogHoneycomb`: LoggingHoneycombResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHoneycombAPI.UpdateLogHoneycomb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingHoneycombName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogHoneycombRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest. |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **dataset** | **string** | The Honeycomb Dataset you want to log to. |  **token** | **string** | The Write Key from the Account page of your Honeycomb account. | 

### Return type

[**LoggingHoneycombResponse**](LoggingHoneycombResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
