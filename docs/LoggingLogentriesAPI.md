# LoggingLogentriesAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogLogentries**](LoggingLogentriesAPI.md#CreateLogLogentries) | **POST** `/service/{service_id}/version/{version_id}/logging/logentries` | Create a Logentries log endpoint
[**DeleteLogLogentries**](LoggingLogentriesAPI.md#DeleteLogLogentries) | **DELETE** `/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}` | Delete a Logentries log endpoint
[**GetLogLogentries**](LoggingLogentriesAPI.md#GetLogLogentries) | **GET** `/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}` | Get a Logentries log endpoint
[**ListLogLogentries**](LoggingLogentriesAPI.md#ListLogLogentries) | **GET** `/service/{service_id}/version/{version_id}/logging/logentries` | List Logentries log endpoints
[**UpdateLogLogentries**](LoggingLogentriesAPI.md#UpdateLogLogentries) | **PUT** `/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}` | Update a Logentries log endpoint



## CreateLogLogentries

Create a Logentries log endpoint



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
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    port := int32(56) // int32 | The port number. (optional) (default to 20000)
    token := "token_example" // string | Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)). (optional)
    useTLS := openapiclient.logging_use_tls(0) // LoggingUseTLS |  (optional) (default to LOGGINGUSETLS_no_tls)
    region := "region_example" // string | The region to which to stream logs. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingLogentriesAPI.CreateLogLogentries(ctx, serviceID, versionID).Name(name).Placement(placement).FormatVersion(formatVersion).ResponseCondition(responseCondition).Format(format).Port(port).Token(token).UseTLS(useTLS).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingLogentriesAPI.CreateLogLogentries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogLogentries`: LoggingLogentriesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingLogentriesAPI.CreateLogLogentries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogLogentriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **port** | **int32** | The port number. | [default to 20000] **token** | **string** | Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)). |  **useTLS** | [**LoggingUseTLS**](LoggingUseTLS.md) |  | [default to LOGGINGUSETLS_no_tls] **region** | **string** | The region to which to stream logs. | 

### Return type

[**LoggingLogentriesResponse**](LoggingLogentriesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogLogentries

Delete a Logentries log endpoint



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
    loggingLogentriesName := "loggingLogentriesName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingLogentriesAPI.DeleteLogLogentries(ctx, serviceID, versionID, loggingLogentriesName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingLogentriesAPI.DeleteLogLogentries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogLogentries`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingLogentriesAPI.DeleteLogLogentries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingLogentriesName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogLogentriesRequest struct via the builder pattern


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


## GetLogLogentries

Get a Logentries log endpoint



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
    loggingLogentriesName := "loggingLogentriesName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingLogentriesAPI.GetLogLogentries(ctx, serviceID, versionID, loggingLogentriesName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingLogentriesAPI.GetLogLogentries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogLogentries`: LoggingLogentriesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingLogentriesAPI.GetLogLogentries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingLogentriesName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogLogentriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingLogentriesResponse**](LoggingLogentriesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogLogentries

List Logentries log endpoints



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
    resp, r, err := apiClient.LoggingLogentriesAPI.ListLogLogentries(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingLogentriesAPI.ListLogLogentries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogLogentries`: []LoggingLogentriesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingLogentriesAPI.ListLogLogentries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogLogentriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingLogentriesResponse**](LoggingLogentriesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogLogentries

Update a Logentries log endpoint



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
    loggingLogentriesName := "loggingLogentriesName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    port := int32(56) // int32 | The port number. (optional) (default to 20000)
    token := "token_example" // string | Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)). (optional)
    useTLS := openapiclient.logging_use_tls(0) // LoggingUseTLS |  (optional) (default to LOGGINGUSETLS_no_tls)
    region := "region_example" // string | The region to which to stream logs. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingLogentriesAPI.UpdateLogLogentries(ctx, serviceID, versionID, loggingLogentriesName).Name(name).Placement(placement).FormatVersion(formatVersion).ResponseCondition(responseCondition).Format(format).Port(port).Token(token).UseTLS(useTLS).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingLogentriesAPI.UpdateLogLogentries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogLogentries`: LoggingLogentriesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingLogentriesAPI.UpdateLogLogentries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingLogentriesName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogLogentriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **port** | **int32** | The port number. | [default to 20000] **token** | **string** | Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)). |  **useTLS** | [**LoggingUseTLS**](LoggingUseTLS.md) |  | [default to LOGGINGUSETLS_no_tls] **region** | **string** | The region to which to stream logs. | 

### Return type

[**LoggingLogentriesResponse**](LoggingLogentriesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
