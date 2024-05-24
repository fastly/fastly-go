# LoggingS3API

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogAwsS3**](LoggingS3API.md#CreateLogAwsS3) | **POST** `/service/{service_id}/version/{version_id}/logging/s3` | Create an AWS S3 log endpoint
[**DeleteLogAwsS3**](LoggingS3API.md#DeleteLogAwsS3) | **DELETE** `/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}` | Delete an AWS S3 log endpoint
[**GetLogAwsS3**](LoggingS3API.md#GetLogAwsS3) | **GET** `/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}` | Get an AWS S3 log endpoint
[**ListLogAwsS3**](LoggingS3API.md#ListLogAwsS3) | **GET** `/service/{service_id}/version/{version_id}/logging/s3` | List AWS S3 log endpoints
[**UpdateLogAwsS3**](LoggingS3API.md#UpdateLogAwsS3) | **PUT** `/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}` | Update an AWS S3 log endpoint



## CreateLogAwsS3

Create an AWS S3 log endpoint



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
    accessKey := "accessKey_example" // string | The access key for your S3 account. Not required if `iam_role` is provided. (optional)
    acl := "acl_example" // string | The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information. (optional)
    bucketName := "bucketName_example" // string | The bucket name for S3 account. (optional)
    domain := "domain_example" // string | The domain of the Amazon S3 endpoint. (optional)
    iamRole := "iamRole_example" // string | The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided. (optional)
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    redundancy := "redundancy_example" // string | The S3 redundancy level. (optional) (default to "null")
    secretKey := "secretKey_example" // string | The secret key for your S3 account. Not required if `iam_role` is provided. (optional)
    serverSideEncryptionKmsKeyID := "serverSideEncryptionKmsKeyId_example" // string | Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`. (optional) (default to "null")
    serverSideEncryption := "serverSideEncryption_example" // string | Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption. (optional) (default to "null")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingS3API.CreateLogAwsS3(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).AccessKey(accessKey).ACL(acl).BucketName(bucketName).Domain(domain).IamRole(iamRole).Path(path).PublicKey(publicKey).Redundancy(redundancy).SecretKey(secretKey).ServerSideEncryptionKmsKeyID(serverSideEncryptionKmsKeyID).ServerSideEncryption(serverSideEncryption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingS3API.CreateLogAwsS3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogAwsS3`: LoggingS3Response
    fmt.Fprintf(os.Stdout, "Response from `LoggingS3API.CreateLogAwsS3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogAwsS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **accessKey** | **string** | The access key for your S3 account. Not required if `iam_role` is provided. |  **acl** | **string** | The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information. |  **bucketName** | **string** | The bucket name for S3 account. |  **domain** | **string** | The domain of the Amazon S3 endpoint. |  **iamRole** | **string** | The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided. |  **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **redundancy** | **string** | The S3 redundancy level. | [default to &quot;null&quot;] **secretKey** | **string** | The secret key for your S3 account. Not required if `iam_role` is provided. |  **serverSideEncryptionKmsKeyID** | **string** | Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`. | [default to &quot;null&quot;] **serverSideEncryption** | **string** | Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption. | [default to &quot;null&quot;]

### Return type

[**LoggingS3Response**](LoggingS3Response.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogAwsS3

Delete an AWS S3 log endpoint



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
    loggingS3Name := "loggingS3Name_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingS3API.DeleteLogAwsS3(ctx, serviceID, versionID, loggingS3Name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingS3API.DeleteLogAwsS3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogAwsS3`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingS3API.DeleteLogAwsS3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingS3Name** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogAwsS3Request struct via the builder pattern


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


## GetLogAwsS3

Get an AWS S3 log endpoint



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
    loggingS3Name := "loggingS3Name_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingS3API.GetLogAwsS3(ctx, serviceID, versionID, loggingS3Name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingS3API.GetLogAwsS3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogAwsS3`: LoggingS3Response
    fmt.Fprintf(os.Stdout, "Response from `LoggingS3API.GetLogAwsS3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingS3Name** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogAwsS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingS3Response**](LoggingS3Response.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogAwsS3

List AWS S3 log endpoints



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
    resp, r, err := apiClient.LoggingS3API.ListLogAwsS3(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingS3API.ListLogAwsS3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogAwsS3`: []LoggingS3Response
    fmt.Fprintf(os.Stdout, "Response from `LoggingS3API.ListLogAwsS3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogAwsS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingS3Response**](LoggingS3Response.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogAwsS3

Update an AWS S3 log endpoint



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
    loggingS3Name := "loggingS3Name_example" // string | The name for the real-time logging configuration.
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
    accessKey := "accessKey_example" // string | The access key for your S3 account. Not required if `iam_role` is provided. (optional)
    acl := "acl_example" // string | The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information. (optional)
    bucketName := "bucketName_example" // string | The bucket name for S3 account. (optional)
    domain := "domain_example" // string | The domain of the Amazon S3 endpoint. (optional)
    iamRole := "iamRole_example" // string | The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided. (optional)
    path := "path_example" // string | The path to upload logs to. (optional) (default to "null")
    publicKey := "publicKey_example" // string | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. (optional) (default to "null")
    redundancy := "redundancy_example" // string | The S3 redundancy level. (optional) (default to "null")
    secretKey := "secretKey_example" // string | The secret key for your S3 account. Not required if `iam_role` is provided. (optional)
    serverSideEncryptionKmsKeyID := "serverSideEncryptionKmsKeyId_example" // string | Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`. (optional) (default to "null")
    serverSideEncryption := "serverSideEncryption_example" // string | Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption. (optional) (default to "null")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingS3API.UpdateLogAwsS3(ctx, serviceID, versionID, loggingS3Name).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).FormatVersion(formatVersion).MessageType(messageType).TimestampFormat(timestampFormat).CompressionCodec(compressionCodec).Period(period).GzipLevel(gzipLevel).AccessKey(accessKey).ACL(acl).BucketName(bucketName).Domain(domain).IamRole(iamRole).Path(path).PublicKey(publicKey).Redundancy(redundancy).SecretKey(secretKey).ServerSideEncryptionKmsKeyID(serverSideEncryptionKmsKeyID).ServerSideEncryption(serverSideEncryption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingS3API.UpdateLogAwsS3`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogAwsS3`: LoggingS3Response
    fmt.Fprintf(os.Stdout, "Response from `LoggingS3API.UpdateLogAwsS3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingS3Name** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogAwsS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **messageType** | **string** | How the message should be formatted. | [default to &quot;classic&quot;] **timestampFormat** | **string** | A timestamp format |  **compressionCodec** | **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. |  **period** | **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [default to 3600] **gzipLevel** | **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [default to 0] **accessKey** | **string** | The access key for your S3 account. Not required if `iam_role` is provided. |  **acl** | **string** | The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information. |  **bucketName** | **string** | The bucket name for S3 account. |  **domain** | **string** | The domain of the Amazon S3 endpoint. |  **iamRole** | **string** | The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided. |  **path** | **string** | The path to upload logs to. | [default to &quot;null&quot;] **publicKey** | **string** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [default to &quot;null&quot;] **redundancy** | **string** | The S3 redundancy level. | [default to &quot;null&quot;] **secretKey** | **string** | The secret key for your S3 account. Not required if `iam_role` is provided. |  **serverSideEncryptionKmsKeyID** | **string** | Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`. | [default to &quot;null&quot;] **serverSideEncryption** | **string** | Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption. | [default to &quot;null&quot;]

### Return type

[**LoggingS3Response**](LoggingS3Response.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
