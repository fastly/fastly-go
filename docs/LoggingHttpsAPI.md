# LoggingHttpsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogHttps**](LoggingHttpsAPI.md#CreateLogHttps) | **POST** `/service/{service_id}/version/{version_id}/logging/https` | Create an HTTPS log endpoint
[**DeleteLogHttps**](LoggingHttpsAPI.md#DeleteLogHttps) | **DELETE** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Delete an HTTPS log endpoint
[**GetLogHttps**](LoggingHttpsAPI.md#GetLogHttps) | **GET** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Get an HTTPS log endpoint
[**ListLogHttps**](LoggingHttpsAPI.md#ListLogHttps) | **GET** `/service/{service_id}/version/{version_id}/logging/https` | List HTTPS log endpoints
[**UpdateLogHttps**](LoggingHttpsAPI.md#UpdateLogHttps) | **PUT** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Update an HTTPS log endpoint



## CreateLogHttps

Create an HTTPS log endpoint



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    logProcessingRegion := "logProcessingRegion_example" // string | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. (optional) (default to "none")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    requestMaxEntries := int32(56) // int32 | The maximum number of logs sent in one request. Defaults `0` (10k). (optional) (default to 0)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` (100MB). (optional) (default to 0)
    url := "url_example" // string | The URL to send logs to. Must use HTTPS. Required. (optional)
    contentType := "contentType_example" // string | Content type of the header sent with the request. (optional) (default to "null")
    headerName := "headerName_example" // string | Name of the custom header sent with the request. (optional) (default to "null")
    messageType := openapiclient.logging_message_type("classic") // LoggingMessageType |  (optional) (default to "classic")
    headerValue := "headerValue_example" // string | Value of the custom header sent with the request. (optional) (default to "null")
    method := "method_example" // string | HTTP method used for request. (optional) (default to "POST")
    jsonFormat := "jsonFormat_example" // string | Enforces valid JSON formatting for log entries. (optional)
    period := int32(56) // int32 | How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds. (optional) (default to 5)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHttpsAPI.CreateLogHttps(ctx, serviceId, versionId).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TlsCaCert(tlsCaCert).TlsClientCert(tlsClientCert).TlsClientKey(tlsClientKey).TlsHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).Url(url).ContentType(contentType).HeaderName(headerName).MessageType(messageType).HeaderValue(headerValue).Method(method).JsonFormat(jsonFormat).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHttpsAPI.CreateLogHttps`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogHttps`: LoggingHttpsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHttpsAPI.CreateLogHttps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogHttpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [default to 0] **url** | **string** | The URL to send logs to. Must use HTTPS. Required. |  **contentType** | **string** | Content type of the header sent with the request. | [default to &quot;null&quot;] **headerName** | **string** | Name of the custom header sent with the request. | [default to &quot;null&quot;] **messageType** | [**LoggingMessageType**](LoggingMessageType.md) |  | [default to &quot;classic&quot;] **headerValue** | **string** | Value of the custom header sent with the request. | [default to &quot;null&quot;] **method** | **string** | HTTP method used for request. | [default to &quot;POST&quot;] **jsonFormat** | **string** | Enforces valid JSON formatting for log entries. |  **period** | **int32** | How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds. | [default to 5]

### Return type

[**LoggingHttpsResponse**](LoggingHttpsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogHttps

Delete an HTTPS log endpoint



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    loggingHttpsName := "loggingHttpsName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHttpsAPI.DeleteLogHttps(ctx, serviceId, versionId, loggingHttpsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHttpsAPI.DeleteLogHttps`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogHttps`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingHttpsAPI.DeleteLogHttps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**loggingHttpsName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogHttpsRequest struct via the builder pattern


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


## GetLogHttps

Get an HTTPS log endpoint



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    loggingHttpsName := "loggingHttpsName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHttpsAPI.GetLogHttps(ctx, serviceId, versionId, loggingHttpsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHttpsAPI.GetLogHttps`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogHttps`: LoggingHttpsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHttpsAPI.GetLogHttps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**loggingHttpsName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogHttpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingHttpsResponse**](LoggingHttpsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogHttps

List HTTPS log endpoints



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHttpsAPI.ListLogHttps(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHttpsAPI.ListLogHttps`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogHttps`: []LoggingHttpsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHttpsAPI.ListLogHttps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogHttpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingHttpsResponse**](LoggingHttpsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogHttps

Update an HTTPS log endpoint



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    loggingHttpsName := "loggingHttpsName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    logProcessingRegion := "logProcessingRegion_example" // string | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. (optional) (default to "none")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    requestMaxEntries := int32(56) // int32 | The maximum number of logs sent in one request. Defaults `0` (10k). (optional) (default to 0)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` (100MB). (optional) (default to 0)
    url := "url_example" // string | The URL to send logs to. Must use HTTPS. Required. (optional)
    contentType := "contentType_example" // string | Content type of the header sent with the request. (optional) (default to "null")
    headerName := "headerName_example" // string | Name of the custom header sent with the request. (optional) (default to "null")
    messageType := openapiclient.logging_message_type("classic") // LoggingMessageType |  (optional) (default to LOGGINGMESSAGETYPE_CLASSIC)
    headerValue := "headerValue_example" // string | Value of the custom header sent with the request. (optional) (default to "null")
    method := "method_example" // string | HTTP method used for request. (optional) (default to "POST")
    jsonFormat := "jsonFormat_example" // string | Enforces valid JSON formatting for log entries. (optional)
    period := int32(56) // int32 | How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds. (optional) (default to 5)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHttpsAPI.UpdateLogHttps(ctx, serviceId, versionId, loggingHttpsName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TlsCaCert(tlsCaCert).TlsClientCert(tlsClientCert).TlsClientKey(tlsClientKey).TlsHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).Url(url).ContentType(contentType).HeaderName(headerName).MessageType(messageType).HeaderValue(headerValue).Method(method).JsonFormat(jsonFormat).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHttpsAPI.UpdateLogHttps`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogHttps`: LoggingHttpsResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHttpsAPI.UpdateLogHttps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**loggingHttpsName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogHttpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [default to 0] **url** | **string** | The URL to send logs to. Must use HTTPS. Required. |  **contentType** | **string** | Content type of the header sent with the request. | [default to &quot;null&quot;] **headerName** | **string** | Name of the custom header sent with the request. | [default to &quot;null&quot;] **messageType** | [**LoggingMessageType**](LoggingMessageType.md) |  | [default to LOGGINGMESSAGETYPE_CLASSIC] **headerValue** | **string** | Value of the custom header sent with the request. | [default to &quot;null&quot;] **method** | **string** | HTTP method used for request. | [default to &quot;POST&quot;] **jsonFormat** | **string** | Enforces valid JSON formatting for log entries. |  **period** | **int32** | How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds. | [default to 5]

### Return type

[**LoggingHttpsResponse**](LoggingHttpsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

