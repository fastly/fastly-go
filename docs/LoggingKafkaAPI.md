# LoggingKafkaAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogKafka**](LoggingKafkaAPI.md#CreateLogKafka) | **POST** `/service/{service_id}/version/{version_id}/logging/kafka` | Create a Kafka log endpoint
[**DeleteLogKafka**](LoggingKafkaAPI.md#DeleteLogKafka) | **DELETE** `/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}` | Delete the Kafka log endpoint
[**GetLogKafka**](LoggingKafkaAPI.md#GetLogKafka) | **GET** `/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}` | Get a Kafka log endpoint
[**ListLogKafka**](LoggingKafkaAPI.md#ListLogKafka) | **GET** `/service/{service_id}/version/{version_id}/logging/kafka` | List Kafka log endpoints
[**UpdateLogKafka**](LoggingKafkaAPI.md#UpdateLogKafka) | **PUT** `/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}` | Update the Kafka log endpoint



## CreateLogKafka

Create a Kafka log endpoint



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
    topic := "topic_example" // string | The Kafka topic to send logs to. Required. (optional)
    brokers := "brokers_example" // string | A comma-separated list of IP addresses or hostnames of Kafka brokers. Required. (optional)
    compressionCodec := "compressionCodec_example" // string | The codec used for compression of your logs. (optional)
    requiredAcks := int32(56) // int32 | The number of acknowledgements a leader must receive before a write is considered successful. (optional) (default to 1)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` (no limit). (optional) (default to 0)
    parseLogKeyvals := true // bool | Enables parsing of key=value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers). (optional)
    authMethod := "authMethod_example" // string | SASL authentication method. (optional)
    user := "user_example" // string | SASL user. (optional)
    password := "password_example" // string | SASL password. (optional)
    useTLS := openapiclient.logging_use_tls_string("0") // LoggingUseTLSString |  (optional) (default to "0")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingKafkaAPI.CreateLogKafka(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).Topic(topic).Brokers(brokers).CompressionCodec(compressionCodec).RequiredAcks(requiredAcks).RequestMaxBytes(requestMaxBytes).ParseLogKeyvals(parseLogKeyvals).AuthMethod(authMethod).User(user).Password(password).UseTLS(useTLS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingKafkaAPI.CreateLogKafka`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogKafka`: LoggingKafkaResponsePost
    fmt.Fprintf(os.Stdout, "Response from `LoggingKafkaAPI.CreateLogKafka`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [default to &quot;%h %l %u %t \&quot;%r\&quot; %&amp;gt;s %b&quot;] **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **topic** | **string** | The Kafka topic to send logs to. Required. |  **brokers** | **string** | A comma-separated list of IP addresses or hostnames of Kafka brokers. Required. |  **compressionCodec** | **string** | The codec used for compression of your logs. |  **requiredAcks** | **int32** | The number of acknowledgements a leader must receive before a write is considered successful. | [default to 1] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` (no limit). | [default to 0] **parseLogKeyvals** | **bool** | Enables parsing of key&#x3D;value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers). |  **authMethod** | **string** | SASL authentication method. |  **user** | **string** | SASL user. |  **password** | **string** | SASL password. |  **useTLS** | [**LoggingUseTLSString**](LoggingUseTLSString.md) |  | [default to &quot;0&quot;]

### Return type

[**LoggingKafkaResponsePost**](LoggingKafkaResponsePost.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogKafka

Delete the Kafka log endpoint



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
    loggingKafkaName := "loggingKafkaName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingKafkaAPI.DeleteLogKafka(ctx, serviceID, versionID, loggingKafkaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingKafkaAPI.DeleteLogKafka`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogKafka`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingKafkaAPI.DeleteLogKafka`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingKafkaName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogKafkaRequest struct via the builder pattern


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


## GetLogKafka

Get a Kafka log endpoint



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
    loggingKafkaName := "loggingKafkaName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingKafkaAPI.GetLogKafka(ctx, serviceID, versionID, loggingKafkaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingKafkaAPI.GetLogKafka`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogKafka`: LoggingKafkaResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingKafkaAPI.GetLogKafka`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingKafkaName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingKafkaResponse**](LoggingKafkaResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogKafka

List Kafka log endpoints



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
    resp, r, err := apiClient.LoggingKafkaAPI.ListLogKafka(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingKafkaAPI.ListLogKafka`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogKafka`: []LoggingKafkaResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingKafkaAPI.ListLogKafka`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingKafkaResponse**](LoggingKafkaResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogKafka

Update the Kafka log endpoint



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
    loggingKafkaName := "loggingKafkaName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingKafkaAPI.UpdateLogKafka(ctx, serviceID, versionID, loggingKafkaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingKafkaAPI.UpdateLogKafka`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogKafka`: LoggingKafkaResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingKafkaAPI.UpdateLogKafka`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingKafkaName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingKafkaResponse**](LoggingKafkaResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
