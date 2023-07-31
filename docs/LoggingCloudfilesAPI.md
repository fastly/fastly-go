# LoggingCloudfilesAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogCloudfiles**](LoggingCloudfilesAPI.md#CreateLogCloudfiles) | **POST** `/service/{service_id}/version/{version_id}/logging/cloudfiles` | Create a Cloud Files log endpoint
[**DeleteLogCloudfiles**](LoggingCloudfilesAPI.md#DeleteLogCloudfiles) | **DELETE** `/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}` | Delete the Cloud Files log endpoint
[**GetLogCloudfiles**](LoggingCloudfilesAPI.md#GetLogCloudfiles) | **GET** `/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}` | Get a Cloud Files log endpoint
[**ListLogCloudfiles**](LoggingCloudfilesAPI.md#ListLogCloudfiles) | **GET** `/service/{service_id}/version/{version_id}/logging/cloudfiles` | List Cloud Files log endpoints
[**UpdateLogCloudfiles**](LoggingCloudfilesAPI.md#UpdateLogCloudfiles) | **PUT** `/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}` | Update the Cloud Files log endpoint



## CreateLogCloudfiles

Create a Cloud Files log endpoint



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
    accessKey := "accessKey_example" // string | Your Cloud Files account access key. (optional)
    bucketName := "bucketName_example" // string | The name of your Cloud Files container. (optional)
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    region := "region_example" // string | The region to stream logs to. (optional)
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    user := "user_example" // string | The username for your Cloud Files account. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingCloudfilesAPI.CreateLogCloudfiles(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).AccessKey(accessKey).BucketName(bucketName).Path(path).Region(region).PublicKey(publicKey).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingCloudfilesAPI.CreateLogCloudfiles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogCloudfiles`: LoggingCloudfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingCloudfilesAPI.CreateLogCloudfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogCloudfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **accessKey** | **string** | Your Cloud Files account access key. |  **bucketName** | **string** | The name of your Cloud Files container. |  **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **region** | **string** | The region to stream logs to. |  **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **user** | **string** | The username for your Cloud Files account. | 

### Return type

[**LoggingCloudfilesResponse**](LoggingCloudfilesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogCloudfiles

Delete the Cloud Files log endpoint



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
    loggingCloudfilesName := "loggingCloudfilesName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingCloudfilesAPI.DeleteLogCloudfiles(ctx, serviceID, versionID, loggingCloudfilesName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingCloudfilesAPI.DeleteLogCloudfiles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogCloudfiles`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingCloudfilesAPI.DeleteLogCloudfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingCloudfilesName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogCloudfilesRequest struct via the builder pattern


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


## GetLogCloudfiles

Get a Cloud Files log endpoint



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
    loggingCloudfilesName := "loggingCloudfilesName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingCloudfilesAPI.GetLogCloudfiles(ctx, serviceID, versionID, loggingCloudfilesName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingCloudfilesAPI.GetLogCloudfiles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogCloudfiles`: LoggingCloudfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingCloudfilesAPI.GetLogCloudfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingCloudfilesName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogCloudfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingCloudfilesResponse**](LoggingCloudfilesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogCloudfiles

List Cloud Files log endpoints



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
    resp, r, err := apiClient.LoggingCloudfilesAPI.ListLogCloudfiles(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingCloudfilesAPI.ListLogCloudfiles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogCloudfiles`: []LoggingCloudfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingCloudfilesAPI.ListLogCloudfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogCloudfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingCloudfilesResponse**](LoggingCloudfilesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogCloudfiles

Update the Cloud Files log endpoint



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
    loggingCloudfilesName := "loggingCloudfilesName_example" // string | The name for the real-time logging configuration.
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
    accessKey := "accessKey_example" // string | Your Cloud Files account access key. (optional)
    bucketName := "bucketName_example" // string | The name of your Cloud Files container. (optional)
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    region := "region_example" // string | The region to stream logs to. (optional)
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    user := "user_example" // string | The username for your Cloud Files account. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingCloudfilesAPI.UpdateLogCloudfiles(ctx, serviceID, versionID, loggingCloudfilesName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).AccessKey(accessKey).BucketName(bucketName).Path(path).Region(region).PublicKey(publicKey).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingCloudfilesAPI.UpdateLogCloudfiles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogCloudfiles`: LoggingCloudfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingCloudfilesAPI.UpdateLogCloudfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingCloudfilesName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogCloudfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **accessKey** | **string** | Your Cloud Files account access key. |  **bucketName** | **string** | The name of your Cloud Files container. |  **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **region** | **string** | The region to stream logs to. |  **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **user** | **string** | The username for your Cloud Files account. | 

### Return type

[**LoggingCloudfilesResponse**](LoggingCloudfilesResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
