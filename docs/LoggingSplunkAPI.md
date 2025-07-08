# LoggingSplunkAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogSplunk**](LoggingSplunkAPI.md#CreateLogSplunk) | **POST** `/service/{service_id}/version/{version_id}/logging/splunk` | Create a Splunk log endpoint
[**DeleteLogSplunk**](LoggingSplunkAPI.md#DeleteLogSplunk) | **DELETE** `/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}` | Delete a Splunk log endpoint
[**GetLogSplunk**](LoggingSplunkAPI.md#GetLogSplunk) | **GET** `/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}` | Get a Splunk log endpoint
[**ListLogSplunk**](LoggingSplunkAPI.md#ListLogSplunk) | **GET** `/service/{service_id}/version/{version_id}/logging/splunk` | List Splunk log endpoints
[**UpdateLogSplunk**](LoggingSplunkAPI.md#UpdateLogSplunk) | **PUT** `/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}` | Update a Splunk log endpoint



## CreateLogSplunk

Create a Splunk log endpoint



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
    format := "format_example" // string | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    logProcessingRegion := "logProcessingRegion_example" // string | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. (optional) (default to "none")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    requestMaxEntries := int32(56) // int32 | The maximum number of logs sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    url := "url_example" // string | The URL to post logs to. (optional)
    token := "token_example" // string | A Splunk token for use in posting logs over HTTP to your collector. (optional)
    useTLS := openapiclient.logging_use_tls_string("0") // LoggingUseTLSString |  (optional) (default to LOGGINGUSETLSSTRING_no_tls)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSplunkAPI.CreateLogSplunk(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).URL(url).Token(token).UseTLS(useTLS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSplunkAPI.CreateLogSplunk`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogSplunk`: LoggingSplunkResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSplunkAPI.CreateLogSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogSplunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [default to 0] **url** | **string** | The URL to post logs to. |  **token** | **string** | A Splunk token for use in posting logs over HTTP to your collector. |  **useTLS** | [**LoggingUseTLSString**](LoggingUseTLSString.md) |  | [default to LOGGINGUSETLSSTRING_no_tls]

### Return type

[**LoggingSplunkResponse**](LoggingSplunkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogSplunk

Delete a Splunk log endpoint



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
    loggingSplunkName := "loggingSplunkName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSplunkAPI.DeleteLogSplunk(ctx, serviceID, versionID, loggingSplunkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSplunkAPI.DeleteLogSplunk`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogSplunk`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingSplunkAPI.DeleteLogSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSplunkName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogSplunkRequest struct via the builder pattern


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


## GetLogSplunk

Get a Splunk log endpoint



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
    loggingSplunkName := "loggingSplunkName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSplunkAPI.GetLogSplunk(ctx, serviceID, versionID, loggingSplunkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSplunkAPI.GetLogSplunk`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSplunk`: LoggingSplunkResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSplunkAPI.GetLogSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSplunkName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSplunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingSplunkResponse**](LoggingSplunkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogSplunk

List Splunk log endpoints



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
    resp, r, err := apiClient.LoggingSplunkAPI.ListLogSplunk(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSplunkAPI.ListLogSplunk`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogSplunk`: []LoggingSplunkResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSplunkAPI.ListLogSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogSplunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingSplunkResponse**](LoggingSplunkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogSplunk

Update a Splunk log endpoint



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
    loggingSplunkName := "loggingSplunkName_example" // string | The name for the real-time logging configuration.
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
    requestMaxEntries := int32(56) // int32 | The maximum number of logs sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    url := "url_example" // string | The URL to post logs to. (optional)
    token := "token_example" // string | A Splunk token for use in posting logs over HTTP to your collector. (optional)
    useTLS := openapiclient.logging_use_tls_string("0") // LoggingUseTLSString |  (optional) (default to LOGGINGUSETLSSTRING_no_tls)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSplunkAPI.UpdateLogSplunk(ctx, serviceID, versionID, loggingSplunkName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).URL(url).Token(token).UseTLS(useTLS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSplunkAPI.UpdateLogSplunk`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogSplunk`: LoggingSplunkResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSplunkAPI.UpdateLogSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSplunkName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSplunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [default to 0] **url** | **string** | The URL to post logs to. |  **token** | **string** | A Splunk token for use in posting logs over HTTP to your collector. |  **useTLS** | [**LoggingUseTLSString**](LoggingUseTLSString.md) |  | [default to LOGGINGUSETLSSTRING_no_tls]

### Return type

[**LoggingSplunkResponse**](LoggingSplunkResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
