# LoggingAzureblobAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogAzure**](LoggingAzureblobAPI.md#CreateLogAzure) | **POST** `/service/{service_id}/version/{version_id}/logging/azureblob` | Create an Azure Blob Storage log endpoint
[**DeleteLogAzure**](LoggingAzureblobAPI.md#DeleteLogAzure) | **DELETE** `/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}` | Delete the Azure Blob Storage log endpoint
[**GetLogAzure**](LoggingAzureblobAPI.md#GetLogAzure) | **GET** `/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}` | Get an Azure Blob Storage log endpoint
[**ListLogAzure**](LoggingAzureblobAPI.md#ListLogAzure) | **GET** `/service/{service_id}/version/{version_id}/logging/azureblob` | List Azure Blob Storage log endpoints
[**UpdateLogAzure**](LoggingAzureblobAPI.md#UpdateLogAzure) | **PUT** `/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}` | Update an Azure Blob Storage log endpoint



## CreateLogAzure

Create an Azure Blob Storage log endpoint



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
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    accountName := "accountName_example" // string | The unique Azure Blob Storage namespace in which your data objects are stored. Required. (optional)
    container := "container_example" // string | The name of the Azure Blob Storage container in which to store logs. Required. (optional)
    sasToken := "sasToken_example" // string | The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required. (optional)
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    fileMaxBytes := int32(56) // int32 | The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.) (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingAzureblobAPI.CreateLogAzure(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).Path(path).AccountName(accountName).Container(container).SasToken(sasToken).PublicKey(publicKey).FileMaxBytes(fileMaxBytes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingAzureblobAPI.CreateLogAzure`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogAzure`: LoggingAzureblobResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingAzureblobAPI.CreateLogAzure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **accountName** | **string** | The unique Azure Blob Storage namespace in which your data objects are stored. Required. |  **container** | **string** | The name of the Azure Blob Storage container in which to store logs. Required. |  **sasToken** | **string** | The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required. |  **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **fileMaxBytes** | **int32** | The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.) | 

### Return type

[**LoggingAzureblobResponse**](LoggingAzureblobResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogAzure

Delete the Azure Blob Storage log endpoint



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
    loggingAzureblobName := "loggingAzureblobName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingAzureblobAPI.DeleteLogAzure(ctx, serviceID, versionID, loggingAzureblobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingAzureblobAPI.DeleteLogAzure`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogAzure`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingAzureblobAPI.DeleteLogAzure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingAzureblobName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogAzureRequest struct via the builder pattern


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


## GetLogAzure

Get an Azure Blob Storage log endpoint



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
    loggingAzureblobName := "loggingAzureblobName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingAzureblobAPI.GetLogAzure(ctx, serviceID, versionID, loggingAzureblobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingAzureblobAPI.GetLogAzure`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogAzure`: LoggingAzureblobResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingAzureblobAPI.GetLogAzure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingAzureblobName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingAzureblobResponse**](LoggingAzureblobResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogAzure

List Azure Blob Storage log endpoints



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
    resp, r, err := apiClient.LoggingAzureblobAPI.ListLogAzure(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingAzureblobAPI.ListLogAzure`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogAzure`: []LoggingAzureblobResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingAzureblobAPI.ListLogAzure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingAzureblobResponse**](LoggingAzureblobResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogAzure

Update an Azure Blob Storage log endpoint



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
    loggingAzureblobName := "loggingAzureblobName_example" // string | The name for the real-time logging configuration.
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
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    accountName := "accountName_example" // string | The unique Azure Blob Storage namespace in which your data objects are stored. Required. (optional)
    container := "container_example" // string | The name of the Azure Blob Storage container in which to store logs. Required. (optional)
    sasToken := "sasToken_example" // string | The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required. (optional)
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    fileMaxBytes := int32(56) // int32 | The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.) (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingAzureblobAPI.UpdateLogAzure(ctx, serviceID, versionID, loggingAzureblobName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).Path(path).AccountName(accountName).Container(container).SasToken(sasToken).PublicKey(publicKey).FileMaxBytes(fileMaxBytes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingAzureblobAPI.UpdateLogAzure`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogAzure`: LoggingAzureblobResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingAzureblobAPI.UpdateLogAzure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingAzureblobName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **accountName** | **string** | The unique Azure Blob Storage namespace in which your data objects are stored. Required. |  **container** | **string** | The name of the Azure Blob Storage container in which to store logs. Required. |  **sasToken** | **string** | The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required. |  **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **fileMaxBytes** | **int32** | The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.) | 

### Return type

[**LoggingAzureblobResponse**](LoggingAzureblobResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
