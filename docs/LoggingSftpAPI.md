# LoggingSftpAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogSftp**](LoggingSftpAPI.md#CreateLogSftp) | **POST** `/service/{service_id}/version/{version_id}/logging/sftp` | Create an SFTP log endpoint
[**DeleteLogSftp**](LoggingSftpAPI.md#DeleteLogSftp) | **DELETE** `/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}` | Delete an SFTP log endpoint
[**GetLogSftp**](LoggingSftpAPI.md#GetLogSftp) | **GET** `/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}` | Get an SFTP log endpoint
[**ListLogSftp**](LoggingSftpAPI.md#ListLogSftp) | **GET** `/service/{service_id}/version/{version_id}/logging/sftp` | List SFTP log endpoints
[**UpdateLogSftp**](LoggingSftpAPI.md#UpdateLogSftp) | **PUT** `/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}` | Update an SFTP log endpoint



## CreateLogSftp

Create an SFTP log endpoint



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
    messageType := "messageType_example" // string | How the message should be formatted. (optional) (default to "classic")
    timestampFormat := "timestampFormat_example" // string | A timestamp format (optional)
    compressionCodec := "compressionCodec_example" // string | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. (optional)
    period := int32(56) // int32 | How frequently log files are finalized so they can be available for reading (in seconds). (optional) (default to 3600)
    gzipLevel := int32(56) // int32 | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. (optional) (default to 0)
    address := "address_example" // string | A hostname or IPv4 address. (optional)
    port := int32(56) // int32 | The port number. (optional) (default to 22)
    password := "password_example" // string | The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. (optional)
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    secretKey := "secretKey_example" // string | The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. (optional) (default to "null")
    sshKnownHosts := "sshKnownHosts_example" // string | A list of host keys for all hosts we can connect to over SFTP. (optional)
    user := "user_example" // string | The username for the server. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSftpAPI.CreateLogSftp(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).Address(address).Port(port).Password(password).Path(path).PublicKey(publicKey).SecretKey(secretKey).SSHKnownHosts(sshKnownHosts).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSftpAPI.CreateLogSftp`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogSftp`: LoggingSftpResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSftpAPI.CreateLogSftp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogSftpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **address** | **string** | A hostname or IPv4 address. |  **port** | **int32** | The port number. | [default to 22] **password** | **string** | The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. |  **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **secretKey** | **string** | The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. | [default to &quot;null&quot;] **sshKnownHosts** | **string** | A list of host keys for all hosts we can connect to over SFTP. |  **user** | **string** | The username for the server. | 

### Return type

[**LoggingSftpResponse**](LoggingSftpResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogSftp

Delete an SFTP log endpoint



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
    loggingSftpName := "loggingSftpName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSftpAPI.DeleteLogSftp(ctx, serviceID, versionID, loggingSftpName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSftpAPI.DeleteLogSftp`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogSftp`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingSftpAPI.DeleteLogSftp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSftpName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogSftpRequest struct via the builder pattern


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


## GetLogSftp

Get an SFTP log endpoint



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
    loggingSftpName := "loggingSftpName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSftpAPI.GetLogSftp(ctx, serviceID, versionID, loggingSftpName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSftpAPI.GetLogSftp`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSftp`: LoggingSftpResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSftpAPI.GetLogSftp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSftpName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSftpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingSftpResponse**](LoggingSftpResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogSftp

List SFTP log endpoints



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
    resp, r, err := apiClient.LoggingSftpAPI.ListLogSftp(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSftpAPI.ListLogSftp`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogSftp`: []LoggingSftpResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSftpAPI.ListLogSftp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogSftpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingSftpResponse**](LoggingSftpResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogSftp

Update an SFTP log endpoint



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
    loggingSftpName := "loggingSftpName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). (optional) (default to "%h %l %u %t \"%r\" %&gt;s %b")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    messageType := "messageType_example" // string | How the message should be formatted. (optional) (default to "classic")
    timestampFormat := "timestampFormat_example" // string | A timestamp format (optional)
    compressionCodec := "compressionCodec_example" // string | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. (optional)
    period := int32(56) // int32 | How frequently log files are finalized so they can be available for reading (in seconds). (optional) (default to 3600)
    gzipLevel := int32(56) // int32 | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. (optional) (default to 0)
    address := "address_example" // string | A hostname or IPv4 address. (optional)
    port := int32(56) // int32 | The port number. (optional) (default to 22)
    password := "password_example" // string | The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. (optional)
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    secretKey := "secretKey_example" // string | The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. (optional) (default to "null")
    sshKnownHosts := "sshKnownHosts_example" // string | A list of host keys for all hosts we can connect to over SFTP. (optional)
    user := "user_example" // string | The username for the server. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingSftpAPI.UpdateLogSftp(ctx, serviceID, versionID, loggingSftpName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).Address(address).Port(port).Password(password).Path(path).PublicKey(publicKey).SecretKey(secretKey).SSHKnownHosts(sshKnownHosts).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingSftpAPI.UpdateLogSftp`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogSftp`: LoggingSftpResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingSftpAPI.UpdateLogSftp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingSftpName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSftpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **address** | **string** | A hostname or IPv4 address. |  **port** | **int32** | The port number. | [default to 22] **password** | **string** | The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. |  **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **secretKey** | **string** | The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. | [default to &quot;null&quot;] **sshKnownHosts** | **string** | A list of host keys for all hosts we can connect to over SFTP. |  **user** | **string** | The username for the server. | 

### Return type

[**LoggingSftpResponse**](LoggingSftpResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
