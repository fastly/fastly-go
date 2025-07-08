# LoggingElasticsearchAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogElasticsearch**](LoggingElasticsearchAPI.md#CreateLogElasticsearch) | **POST** `/service/{service_id}/version/{version_id}/logging/elasticsearch` | Create an Elasticsearch log endpoint
[**DeleteLogElasticsearch**](LoggingElasticsearchAPI.md#DeleteLogElasticsearch) | **DELETE** `/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}` | Delete an Elasticsearch log endpoint
[**GetLogElasticsearch**](LoggingElasticsearchAPI.md#GetLogElasticsearch) | **GET** `/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}` | Get an Elasticsearch log endpoint
[**ListLogElasticsearch**](LoggingElasticsearchAPI.md#ListLogElasticsearch) | **GET** `/service/{service_id}/version/{version_id}/logging/elasticsearch` | List Elasticsearch log endpoints
[**UpdateLogElasticsearch**](LoggingElasticsearchAPI.md#UpdateLogElasticsearch) | **PUT** `/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}` | Update an Elasticsearch log endpoint



## CreateLogElasticsearch

Create an Elasticsearch log endpoint



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
    format := "format_example" // string | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest. (optional)
    logProcessingRegion := "logProcessingRegion_example" // string | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. (optional) (default to "none")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    requestMaxEntries := int32(56) // int32 | The maximum number of logs sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    index := "index_example" // string | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today's date. (optional)
    url := "url_example" // string | The URL to stream logs to. Must use HTTPS. (optional)
    pipeline := "pipeline_example" // string | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). (optional)
    user := "user_example" // string | Basic Auth username. (optional)
    password := "password_example" // string | Basic Auth password. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingElasticsearchAPI.CreateLogElasticsearch(ctx, serviceID, versionID).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).Index(index).URL(url).Pipeline(pipeline).User(user).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingElasticsearchAPI.CreateLogElasticsearch`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogElasticsearch`: LoggingElasticsearchResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingElasticsearchAPI.CreateLogElasticsearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogElasticsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest. |  **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [default to 0] **index** | **string** | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today&#39;s date. |  **url** | **string** | The URL to stream logs to. Must use HTTPS. |  **pipeline** | **string** | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). |  **user** | **string** | Basic Auth username. |  **password** | **string** | Basic Auth password. | 

### Return type

[**LoggingElasticsearchResponse**](LoggingElasticsearchResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogElasticsearch

Delete an Elasticsearch log endpoint



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
    loggingElasticsearchName := "loggingElasticsearchName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingElasticsearchAPI.DeleteLogElasticsearch(ctx, serviceID, versionID, loggingElasticsearchName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingElasticsearchAPI.DeleteLogElasticsearch`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogElasticsearch`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingElasticsearchAPI.DeleteLogElasticsearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingElasticsearchName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogElasticsearchRequest struct via the builder pattern


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


## GetLogElasticsearch

Get an Elasticsearch log endpoint



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
    loggingElasticsearchName := "loggingElasticsearchName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingElasticsearchAPI.GetLogElasticsearch(ctx, serviceID, versionID, loggingElasticsearchName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingElasticsearchAPI.GetLogElasticsearch`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogElasticsearch`: LoggingElasticsearchResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingElasticsearchAPI.GetLogElasticsearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingElasticsearchName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogElasticsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingElasticsearchResponse**](LoggingElasticsearchResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogElasticsearch

List Elasticsearch log endpoints



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
    resp, r, err := apiClient.LoggingElasticsearchAPI.ListLogElasticsearch(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingElasticsearchAPI.ListLogElasticsearch`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogElasticsearch`: []LoggingElasticsearchResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingElasticsearchAPI.ListLogElasticsearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogElasticsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingElasticsearchResponse**](LoggingElasticsearchResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogElasticsearch

Update an Elasticsearch log endpoint



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
    loggingElasticsearchName := "loggingElasticsearchName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name for the real-time logging configuration. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest. (optional)
    logProcessingRegion := "logProcessingRegion_example" // string | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. (optional) (default to "none")
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsHostname := "tlsHostname_example" // string | The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. (optional) (default to "null")
    requestMaxEntries := int32(56) // int32 | The maximum number of logs sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    requestMaxBytes := int32(56) // int32 | The maximum number of bytes sent in one request. Defaults `0` for unbounded. (optional) (default to 0)
    index := "index_example" // string | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today's date. (optional)
    url := "url_example" // string | The URL to stream logs to. Must use HTTPS. (optional)
    pipeline := "pipeline_example" // string | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). (optional)
    user := "user_example" // string | Basic Auth username. (optional)
    password := "password_example" // string | Basic Auth password. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingElasticsearchAPI.UpdateLogElasticsearch(ctx, serviceID, versionID, loggingElasticsearchName).Name(name).Placement(placement).ResponseCondition(responseCondition).Format(format).LogProcessingRegion(logProcessingRegion).FormatVersion(formatVersion).TLSCaCert(tlsCaCert).TLSClientCert(tlsClientCert).TLSClientKey(tlsClientKey).TLSHostname(tlsHostname).RequestMaxEntries(requestMaxEntries).RequestMaxBytes(requestMaxBytes).Index(index).URL(url).Pipeline(pipeline).User(user).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingElasticsearchAPI.UpdateLogElasticsearch`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogElasticsearch`: LoggingElasticsearchResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingElasticsearchAPI.UpdateLogElasticsearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingElasticsearchName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogElasticsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the real-time logging configuration. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest. |  **logProcessingRegion** | **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [default to &quot;none&quot;] **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsHostname** | **string** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [default to &quot;null&quot;] **requestMaxEntries** | **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [default to 0] **requestMaxBytes** | **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [default to 0] **index** | **string** | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today&#39;s date. |  **url** | **string** | The URL to stream logs to. Must use HTTPS. |  **pipeline** | **string** | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). |  **user** | **string** | Basic Auth username. |  **password** | **string** | Basic Auth password. | 

### Return type

[**LoggingElasticsearchResponse**](LoggingElasticsearchResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
