# LoggingDigitaloceanAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogDigocean**](LoggingDigitaloceanAPI.md#CreateLogDigocean) | **POST** `/service/{service_id}/version/{version_id}/logging/digitalocean` | Create a DigitalOcean Spaces log endpoint
[**DeleteLogDigocean**](LoggingDigitaloceanAPI.md#DeleteLogDigocean) | **DELETE** `/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}` | Delete a DigitalOcean Spaces log endpoint
[**GetLogDigocean**](LoggingDigitaloceanAPI.md#GetLogDigocean) | **GET** `/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}` | Get a DigitalOcean Spaces log endpoint
[**ListLogDigocean**](LoggingDigitaloceanAPI.md#ListLogDigocean) | **GET** `/service/{service_id}/version/{version_id}/logging/digitalocean` | List DigitalOcean Spaces log endpoints
[**UpdateLogDigocean**](LoggingDigitaloceanAPI.md#UpdateLogDigocean) | **PUT** `/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}` | Update a DigitalOcean Spaces log endpoint



## CreateLogDigocean

Create a DigitalOcean Spaces log endpoint



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
    bucketName := "bucketName_example" // string | The name of the DigitalOcean Space. (optional)
    accessKey := "accessKey_example" // string | Your DigitalOcean Spaces account access key. (optional)
    secretKey := "secretKey_example" // string | Your DigitalOcean Spaces account secret key. (optional)
    domain := "domain_example" // string | The domain of the DigitalOcean Spaces endpoint. (optional) (default to "nyc3.digitaloceanspaces.com")
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingDigitaloceanAPI.CreateLogDigocean(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).BucketName(bucketName).AccessKey(accessKey).SecretKey(secretKey).Domain(domain).Path(path).PublicKey(publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingDigitaloceanAPI.CreateLogDigocean`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogDigocean`: LoggingDigitaloceanResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingDigitaloceanAPI.CreateLogDigocean`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogDigoceanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **bucketName** | **string** | The name of the DigitalOcean Space. |  **accessKey** | **string** | Your DigitalOcean Spaces account access key. |  **secretKey** | **string** | Your DigitalOcean Spaces account secret key. |  **domain** | **string** | The domain of the DigitalOcean Spaces endpoint. | [default to &quot;nyc3.digitaloceanspaces.com&quot;] **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;]

### Return type

[**LoggingDigitaloceanResponse**](LoggingDigitaloceanResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogDigocean

Delete a DigitalOcean Spaces log endpoint



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
    loggingDigitaloceanName := "loggingDigitaloceanName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingDigitaloceanAPI.DeleteLogDigocean(ctx, serviceID, versionID, loggingDigitaloceanName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingDigitaloceanAPI.DeleteLogDigocean`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogDigocean`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingDigitaloceanAPI.DeleteLogDigocean`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingDigitaloceanName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogDigoceanRequest struct via the builder pattern


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


## GetLogDigocean

Get a DigitalOcean Spaces log endpoint



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
    loggingDigitaloceanName := "loggingDigitaloceanName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingDigitaloceanAPI.GetLogDigocean(ctx, serviceID, versionID, loggingDigitaloceanName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingDigitaloceanAPI.GetLogDigocean`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogDigocean`: LoggingDigitaloceanResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingDigitaloceanAPI.GetLogDigocean`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingDigitaloceanName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogDigoceanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingDigitaloceanResponse**](LoggingDigitaloceanResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogDigocean

List DigitalOcean Spaces log endpoints



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
    resp, r, err := apiClient.LoggingDigitaloceanAPI.ListLogDigocean(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingDigitaloceanAPI.ListLogDigocean`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogDigocean`: []LoggingDigitaloceanResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingDigitaloceanAPI.ListLogDigocean`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogDigoceanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingDigitaloceanResponse**](LoggingDigitaloceanResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogDigocean

Update a DigitalOcean Spaces log endpoint



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
    loggingDigitaloceanName := "loggingDigitaloceanName_example" // string | The name for the real-time logging configuration.
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
    bucketName := "bucketName_example" // string | The name of the DigitalOcean Space. (optional)
    accessKey := "accessKey_example" // string | Your DigitalOcean Spaces account access key. (optional)
    secretKey := "secretKey_example" // string | Your DigitalOcean Spaces account secret key. (optional)
    domain := "domain_example" // string | The domain of the DigitalOcean Spaces endpoint. (optional) (default to "nyc3.digitaloceanspaces.com")
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingDigitaloceanAPI.UpdateLogDigocean(ctx, serviceID, versionID, loggingDigitaloceanName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).BucketName(bucketName).AccessKey(accessKey).SecretKey(secretKey).Domain(domain).Path(path).PublicKey(publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingDigitaloceanAPI.UpdateLogDigocean`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogDigocean`: LoggingDigitaloceanResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingDigitaloceanAPI.UpdateLogDigocean`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingDigitaloceanName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogDigoceanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **bucketName** | **string** | The name of the DigitalOcean Space. |  **accessKey** | **string** | Your DigitalOcean Spaces account access key. |  **secretKey** | **string** | Your DigitalOcean Spaces account secret key. |  **domain** | **string** | The domain of the DigitalOcean Spaces endpoint. | [default to &quot;nyc3.digitaloceanspaces.com&quot;] **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;]

### Return type

[**LoggingDigitaloceanResponse**](LoggingDigitaloceanResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
