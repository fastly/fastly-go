# LoggingGrafanacloudlogsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogGrafanacloudlogs**](LoggingGrafanacloudlogsAPI.md#CreateLogGrafanacloudlogs) | **POST** `/service/{service_id}/version/{version_id}/logging/grafanacloudlogs` | Create a Grafana Cloud Logs log endpoint
[**DeleteLogGrafanacloudlogs**](LoggingGrafanacloudlogsAPI.md#DeleteLogGrafanacloudlogs) | **DELETE** `/service/{service_id}/version/{version_id}/logging/grafanacloudlogs/{logging_grafanacloudlogs_name}` | Delete the Grafana Cloud Logs log endpoint
[**GetLogGrafanacloudlogs**](LoggingGrafanacloudlogsAPI.md#GetLogGrafanacloudlogs) | **GET** `/service/{service_id}/version/{version_id}/logging/grafanacloudlogs/{logging_grafanacloudlogs_name}` | Get a Grafana Cloud Logs log endpoint
[**ListLogGrafanacloudlogs**](LoggingGrafanacloudlogsAPI.md#ListLogGrafanacloudlogs) | **GET** `/service/{service_id}/version/{version_id}/logging/grafanacloudlogs` | List Grafana Cloud Logs log endpoints
[**UpdateLogGrafanacloudlogs**](LoggingGrafanacloudlogsAPI.md#UpdateLogGrafanacloudlogs) | **PUT** `/service/{service_id}/version/{version_id}/logging/grafanacloudlogs/{logging_grafanacloudlogs_name}` | Update a Grafana Cloud Logs log endpoint



## CreateLogGrafanacloudlogs

Create a Grafana Cloud Logs log endpoint



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
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    user := "user_example" // string | The Grafana Cloud Logs Dataset you want to log to. (optional)
    url := "url_example" // string | The URL of the Loki instance in your Grafana stack. (optional)
    token := "token_example" // string | The Grafana Access Policy token with `logs:write` access scoped to your Loki instance. (optional)
    index := "index_example" // string | The Stream Labels, a JSON string used to identify the stream. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingGrafanacloudlogsAPI.CreateLogGrafanacloudlogs(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).User(user).URL(url).Token(token).Index(index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingGrafanacloudlogsAPI.CreateLogGrafanacloudlogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogGrafanacloudlogs`: LoggingGrafanacloudlogsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingGrafanacloudlogsAPI.CreateLogGrafanacloudlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogGrafanacloudlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **user** | **string** | The Grafana Cloud Logs Dataset you want to log to. |  **url** | **string** | The URL of the Loki instance in your Grafana stack. |  **token** | **string** | The Grafana Access Policy token with `logs:write` access scoped to your Loki instance. |  **index** | **string** | The Stream Labels, a JSON string used to identify the stream. | 

### Return type

[**LoggingGrafanacloudlogsResponse**](LoggingGrafanacloudlogsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogGrafanacloudlogs

Delete the Grafana Cloud Logs log endpoint



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
    loggingGrafanacloudlogsName := "loggingGrafanacloudlogsName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingGrafanacloudlogsAPI.DeleteLogGrafanacloudlogs(ctx, serviceID, versionID, loggingGrafanacloudlogsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingGrafanacloudlogsAPI.DeleteLogGrafanacloudlogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogGrafanacloudlogs`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingGrafanacloudlogsAPI.DeleteLogGrafanacloudlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingGrafanacloudlogsName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogGrafanacloudlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetLogGrafanacloudlogs

Get a Grafana Cloud Logs log endpoint



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
    loggingGrafanacloudlogsName := "loggingGrafanacloudlogsName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingGrafanacloudlogsAPI.GetLogGrafanacloudlogs(ctx, serviceID, versionID, loggingGrafanacloudlogsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingGrafanacloudlogsAPI.GetLogGrafanacloudlogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogGrafanacloudlogs`: LoggingGrafanacloudlogsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingGrafanacloudlogsAPI.GetLogGrafanacloudlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingGrafanacloudlogsName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogGrafanacloudlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingGrafanacloudlogsResponse**](LoggingGrafanacloudlogsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogGrafanacloudlogs

List Grafana Cloud Logs log endpoints



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
    resp, r, err := apiClient.LoggingGrafanacloudlogsAPI.ListLogGrafanacloudlogs(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingGrafanacloudlogsAPI.ListLogGrafanacloudlogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogGrafanacloudlogs`: []LoggingGrafanacloudlogsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingGrafanacloudlogsAPI.ListLogGrafanacloudlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogGrafanacloudlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingGrafanacloudlogsResponse**](LoggingGrafanacloudlogsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogGrafanacloudlogs

Update a Grafana Cloud Logs log endpoint



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
    loggingGrafanacloudlogsName := "loggingGrafanacloudlogsName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    user := "user_example" // string | The Grafana Cloud Logs Dataset you want to log to. (optional)
    url := "url_example" // string | The URL of the Loki instance in your Grafana stack. (optional)
    token := "token_example" // string | The Grafana Access Policy token with `logs:write` access scoped to your Loki instance. (optional)
    index := "index_example" // string | The Stream Labels, a JSON string used to identify the stream. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingGrafanacloudlogsAPI.UpdateLogGrafanacloudlogs(ctx, serviceID, versionID, loggingGrafanacloudlogsName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).User(user).URL(url).Token(token).Index(index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingGrafanacloudlogsAPI.UpdateLogGrafanacloudlogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogGrafanacloudlogs`: LoggingGrafanacloudlogsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingGrafanacloudlogsAPI.UpdateLogGrafanacloudlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingGrafanacloudlogsName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogGrafanacloudlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **user** | **string** | The Grafana Cloud Logs Dataset you want to log to. |  **url** | **string** | The URL of the Loki instance in your Grafana stack. |  **token** | **string** | The Grafana Access Policy token with `logs:write` access scoped to your Loki instance. |  **index** | **string** | The Stream Labels, a JSON string used to identify the stream. | 

### Return type

[**LoggingGrafanacloudlogsResponse**](LoggingGrafanacloudlogsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
