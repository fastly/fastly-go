# LoggingSyslogAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogSyslog**](LoggingSyslogAPI.md#CreateLogSyslog) | **POST** `/service/{service_id}/version/{version_id}/logging/syslog` | Create a syslog log endpoint
[**DeleteLogSyslog**](LoggingSyslogAPI.md#DeleteLogSyslog) | **DELETE** `/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}` | Delete a syslog log endpoint
[**GetLogSyslog**](LoggingSyslogAPI.md#GetLogSyslog) | **GET** `/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}` | Get a syslog log endpoint
[**ListLogSyslog**](LoggingSyslogAPI.md#ListLogSyslog) | **GET** `/service/{service_id}/version/{version_id}/logging/syslog` | List Syslog log endpoints
[**UpdateLogSyslog**](LoggingSyslogAPI.md#UpdateLogSyslog) | **PUT** `/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}` | Update a syslog log endpoint



## CreateLogSyslog

Create a syslog log endpoint



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
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    address := "address_example" // string | A hostname or IPv4 address. (optional)
    port := int32(56) // int32 | The port number. (optional) (default to 514)
    messageType := openapiclient.logging_message_type("classic") // LoggingMessageType |  (optional) (default to LOGGINGMESSAGETYPE_CLASSIC)
    hostname := "hostname_example" // string | The hostname used for the syslog endpoint. (optional)
    ipv4 := "ipv4_example" // string | The IPv4 address used for the syslog endpoint. (optional)
    token := "token_example" // string | Whether to prepend each message with a specific token. (optional) (default to "null")
    useTLS := openapiclient.logging_use_tls(0) // LoggingUseTLS |  (optional) (default to LOGGINGUSETLS_no_tls)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSyslogAPI.CreateLogSyslog(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).Address(address).Port(port).MessageType(messageType).Hostname(hostname).Ipv4(ipv4).Token(token).UseTLS(useTLS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSyslogAPI.CreateLogSyslog`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogSyslog`: LoggingSyslogResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSyslogAPI.CreateLogSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogSyslogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **address** | **string** | A hostname or IPv4 address. |  **port** | **int32** | The port number. | [default to 514] **messageType** | [**LoggingMessageType**](LoggingMessageType.md) |  | [default to LOGGINGMESSAGETYPE_CLASSIC] **hostname** | **string** | The hostname used for the syslog endpoint. |  **ipv4** | **string** | The IPv4 address used for the syslog endpoint. |  **token** | **string** | Whether to prepend each message with a specific token. | [default to &quot;null&quot;] **useTLS** | [**LoggingUseTLS**](LoggingUseTLS.md) |  | [default to LOGGINGUSETLS_no_tls]

### Return type

[**LoggingSyslogResponse**](LoggingSyslogResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogSyslog

Delete a syslog log endpoint



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
    loggingSyslogName := "loggingSyslogName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSyslogAPI.DeleteLogSyslog(ctx, serviceID, versionID, loggingSyslogName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSyslogAPI.DeleteLogSyslog`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogSyslog`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingSyslogAPI.DeleteLogSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSyslogName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogSyslogRequest struct via the builder pattern


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


## GetLogSyslog

Get a syslog log endpoint



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
    loggingSyslogName := "loggingSyslogName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSyslogAPI.GetLogSyslog(ctx, serviceID, versionID, loggingSyslogName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSyslogAPI.GetLogSyslog`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSyslog`: LoggingSyslogResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSyslogAPI.GetLogSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSyslogName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSyslogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingSyslogResponse**](LoggingSyslogResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogSyslog

List Syslog log endpoints



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
    resp, r, err := apiClient.LoggingSyslogAPI.ListLogSyslog(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSyslogAPI.ListLogSyslog`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogSyslog`: []LoggingSyslogResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSyslogAPI.ListLogSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogSyslogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingSyslogResponse**](LoggingSyslogResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogSyslog

Update a syslog log endpoint



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
    loggingSyslogName := "loggingSyslogName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    address := "address_example" // string | A hostname or IPv4 address. (optional)
    port := int32(56) // int32 | The port number. (optional) (default to 514)
    messageType := openapiclient.logging_message_type("classic") // LoggingMessageType |  (optional) (default to LOGGINGMESSAGETYPE_CLASSIC)
    hostname := "hostname_example" // string | The hostname used for the syslog endpoint. (optional)
    ipv4 := "ipv4_example" // string | The IPv4 address used for the syslog endpoint. (optional)
    token := "token_example" // string | Whether to prepend each message with a specific token. (optional) (default to "null")
    useTLS := openapiclient.logging_use_tls(0) // LoggingUseTLS |  (optional) (default to LOGGINGUSETLS_no_tls)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSyslogAPI.UpdateLogSyslog(ctx, serviceID, versionID, loggingSyslogName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).Address(address).Port(port).MessageType(messageType).Hostname(hostname).Ipv4(ipv4).Token(token).UseTLS(useTLS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSyslogAPI.UpdateLogSyslog`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogSyslog`: LoggingSyslogResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSyslogAPI.UpdateLogSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSyslogName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSyslogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **address** | **string** | A hostname or IPv4 address. |  **port** | **int32** | The port number. | [default to 514] **messageType** | [**LoggingMessageType**](LoggingMessageType.md) |  | [default to LOGGINGMESSAGETYPE_CLASSIC] **hostname** | **string** | The hostname used for the syslog endpoint. |  **ipv4** | **string** | The IPv4 address used for the syslog endpoint. |  **token** | **string** | Whether to prepend each message with a specific token. | [default to &quot;null&quot;] **useTLS** | [**LoggingUseTLS**](LoggingUseTLS.md) |  | [default to LOGGINGUSETLS_no_tls]

### Return type

[**LoggingSyslogResponse**](LoggingSyslogResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
