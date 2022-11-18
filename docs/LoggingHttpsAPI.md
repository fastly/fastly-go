# LoggingHTTPSAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogHTTPS**](LoggingHTTPSAPI.md#CreateLogHTTPS) | **POST** `/service/{service_id}/version/{version_id}/logging/https` | Create an HTTPS log endpoint
[**DeleteLogHTTPS**](LoggingHTTPSAPI.md#DeleteLogHTTPS) | **DELETE** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Delete an HTTPS log endpoint
[**GetLogHTTPS**](LoggingHTTPSAPI.md#GetLogHTTPS) | **GET** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Get an HTTPS log endpoint
[**ListLogHTTPS**](LoggingHTTPSAPI.md#ListLogHTTPS) | **GET** `/service/{service_id}/version/{version_id}/logging/https` | List HTTPS log endpoints
[**UpdateLogHTTPS**](LoggingHTTPSAPI.md#UpdateLogHTTPS) | **PUT** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Update an HTTPS log endpoint



## CreateLogHTTPS

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHTTPSAPI.CreateLogHTTPS(ctx, serviceID, versionID).Name(name).Placement(placement).FormatVersion(formatVersion).ResponseCondition(responseCondition).Format(format).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).URL(url).ContentType(contentType).HeaderName(headerName).MessageType(messageType).HeaderValue(headerValue).Method(method).JSONFormat(jsonFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHTTPSAPI.CreateLogHTTPS`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogHTTPS`: LoggingHTTPSResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHTTPSAPI.CreateLogHTTPS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogHTTPSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [default to 0] **url** | **string** | The URL to send logs to. Must use HTTPS. Required. |  **contentType** | **string** | Content type of the header sent with the request. | [default to &quot;null&quot;] **headerName** | **string** | Name of the custom header sent with the request. | [default to &quot;null&quot;] **messageType** | [**LoggingMessageType**](LoggingMessageType.md) |  | [default to &quot;classic&quot;] **headerValue** | **string** | Value of the custom header sent with the request. | [default to &quot;null&quot;] **method** | **string** | HTTP method used for request. | [default to &quot;POST&quot;] **jsonFormat** | **string** | Enforces valid JSON formatting for log entries. | 

### Return type

[**LoggingHTTPSResponse**](LoggingHTTPSResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogHTTPS

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    loggingHTTPSName := "loggingHTTPSName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHTTPSAPI.DeleteLogHTTPS(ctx, serviceID, versionID, loggingHTTPSName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHTTPSAPI.DeleteLogHTTPS`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogHTTPS`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingHTTPSAPI.DeleteLogHTTPS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingHTTPSName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogHTTPSRequest struct via the builder pattern


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


## GetLogHTTPS

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    loggingHTTPSName := "loggingHTTPSName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHTTPSAPI.GetLogHTTPS(ctx, serviceID, versionID, loggingHTTPSName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHTTPSAPI.GetLogHTTPS`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogHTTPS`: LoggingHTTPSResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHTTPSAPI.GetLogHTTPS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingHTTPSName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogHTTPSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingHTTPSResponse**](LoggingHTTPSResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogHTTPS

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHTTPSAPI.ListLogHTTPS(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHTTPSAPI.ListLogHTTPS`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogHTTPS`: []LoggingHTTPSResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHTTPSAPI.ListLogHTTPS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogHTTPSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingHTTPSResponse**](LoggingHTTPSResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogHTTPS

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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    loggingHTTPSName := "loggingHTTPSName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingHTTPSAPI.UpdateLogHTTPS(ctx, serviceID, versionID, loggingHTTPSName).Name(name).Placement(placement).FormatVersion(formatVersion).ResponseCondition(responseCondition).Format(format).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).URL(url).ContentType(contentType).HeaderName(headerName).MessageType(messageType).HeaderValue(headerValue).Method(method).JSONFormat(jsonFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingHTTPSAPI.UpdateLogHTTPS`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogHTTPS`: LoggingHTTPSResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingHTTPSAPI.UpdateLogHTTPS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingHTTPSName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogHTTPSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [default to 0] **url** | **string** | The URL to send logs to. Must use HTTPS. Required. |  **contentType** | **string** | Content type of the header sent with the request. | [default to &quot;null&quot;] **headerName** | **string** | Name of the custom header sent with the request. | [default to &quot;null&quot;] **messageType** | [**LoggingMessageType**](LoggingMessageType.md) |  | [default to LOGGINGMESSAGETYPE_CLASSIC] **headerValue** | **string** | Value of the custom header sent with the request. | [default to &quot;null&quot;] **method** | **string** | HTTP method used for request. | [default to &quot;POST&quot;] **jsonFormat** | **string** | Enforces valid JSON formatting for log entries. | 

### Return type

[**LoggingHTTPSResponse**](LoggingHTTPSResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
