# LoggingBigqueryAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogBigquery**](LoggingBigqueryAPI.md#CreateLogBigquery) | **POST** `/service/{service_id}/version/{version_id}/logging/bigquery` | Create a BigQuery log endpoint
[**DeleteLogBigquery**](LoggingBigqueryAPI.md#DeleteLogBigquery) | **DELETE** `/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}` | Delete a BigQuery log endpoint
[**GetLogBigquery**](LoggingBigqueryAPI.md#GetLogBigquery) | **GET** `/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}` | Get a BigQuery log endpoint
[**ListLogBigquery**](LoggingBigqueryAPI.md#ListLogBigquery) | **GET** `/service/{service_id}/version/{version_id}/logging/bigquery` | List BigQuery log endpoints
[**UpdateLogBigquery**](LoggingBigqueryAPI.md#UpdateLogBigquery) | **PUT** `/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}` | Update a BigQuery log endpoint



## CreateLogBigquery

Create a BigQuery log endpoint



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
    name := "name_example" // string | The name of the BigQuery logging object. Used as a primary key for API access. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table. (optional)
    user := "user_example" // string | Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified. (optional)
    secretKey := "secretKey_example" // string | Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified. (optional)
    accountName := "accountName_example" // string | The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided. (optional)
    dataset := "dataset_example" // string | Your BigQuery dataset. (optional)
    table := "table_example" // string | Your BigQuery table. (optional)
    templateSuffix := "templateSuffix_example" // string | BigQuery table name suffix template. Optional. (optional)
    projectID := "projectId_example" // string | Your Google Cloud Platform project ID. Required (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingBigqueryAPI.CreateLogBigquery(ctx, serviceID, versionID).Name(name).Placement(placement).FormatVersion(formatVersion).ResponseCondition(responseCondition).Format(format).User(user).SecretKey(secretKey).AccountName(accountName).Dataset(dataset).Table(table).TemplateSuffix(templateSuffix).ProjectID(projectID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingBigqueryAPI.CreateLogBigquery`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogBigquery`: LoggingBigqueryResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingBigqueryAPI.CreateLogBigquery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogBigqueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the BigQuery logging object. Used as a primary key for API access. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table. |  **user** | **string** | Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified. |  **secretKey** | **string** | Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified. |  **accountName** | **string** | The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided. |  **dataset** | **string** | Your BigQuery dataset. |  **table** | **string** | Your BigQuery table. |  **templateSuffix** | **string** | BigQuery table name suffix template. Optional. |  **projectID** | **string** | Your Google Cloud Platform project ID. Required | 

### Return type

[**LoggingBigqueryResponse**](LoggingBigqueryResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteLogBigquery

Delete a BigQuery log endpoint



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
    loggingBigqueryName := "loggingBigqueryName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingBigqueryAPI.DeleteLogBigquery(ctx, serviceID, versionID, loggingBigqueryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingBigqueryAPI.DeleteLogBigquery`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogBigquery`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `LoggingBigqueryAPI.DeleteLogBigquery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingBigqueryName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogBigqueryRequest struct via the builder pattern


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


## GetLogBigquery

Get a BigQuery log endpoint



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
    loggingBigqueryName := "loggingBigqueryName_example" // string | The name for the real-time logging configuration.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingBigqueryAPI.GetLogBigquery(ctx, serviceID, versionID, loggingBigqueryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingBigqueryAPI.GetLogBigquery`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogBigquery`: LoggingBigqueryResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingBigqueryAPI.GetLogBigquery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingBigqueryName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogBigqueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggingBigqueryResponse**](LoggingBigqueryResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListLogBigquery

List BigQuery log endpoints



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
    resp, r, err := apiClient.LoggingBigqueryAPI.ListLogBigquery(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingBigqueryAPI.ListLogBigquery`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogBigquery`: []LoggingBigqueryResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingBigqueryAPI.ListLogBigquery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogBigqueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LoggingBigqueryResponse**](LoggingBigqueryResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateLogBigquery

Update a BigQuery log endpoint



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
    loggingBigqueryName := "loggingBigqueryName_example" // string | The name for the real-time logging configuration.
    name := "name_example" // string | The name of the BigQuery logging object. Used as a primary key for API access. (optional)
    placement := "placement_example" // string | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  (optional)
    formatVersion := int32(56) // int32 | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  (optional) (default to 2)
    responseCondition := "responseCondition_example" // string | The name of an existing condition in the configured endpoint, or leave blank to always execute. (optional)
    format := "format_example" // string | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table. (optional)
    user := "user_example" // string | Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified. (optional)
    secretKey := "secretKey_example" // string | Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified. (optional)
    accountName := "accountName_example" // string | The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided. (optional)
    dataset := "dataset_example" // string | Your BigQuery dataset. (optional)
    table := "table_example" // string | Your BigQuery table. (optional)
    templateSuffix := "templateSuffix_example" // string | BigQuery table name suffix template. Optional. (optional)
    projectID := "projectId_example" // string | Your Google Cloud Platform project ID. Required (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LoggingBigqueryAPI.UpdateLogBigquery(ctx, serviceID, versionID, loggingBigqueryName).Name(name).Placement(placement).FormatVersion(formatVersion).ResponseCondition(responseCondition).Format(format).User(user).SecretKey(secretKey).AccountName(accountName).Dataset(dataset).Table(table).TemplateSuffix(templateSuffix).ProjectID(projectID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingBigqueryAPI.UpdateLogBigquery`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogBigquery`: LoggingBigqueryResponse
    fmt.Fprintf(os.Stdout, "Response from `LoggingBigqueryAPI.UpdateLogBigquery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**loggingBigqueryName** | **string** | The name for the real-time logging configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogBigqueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the BigQuery logging object. Used as a primary key for API access. |  **placement** | **string** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  |  **formatVersion** | **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [default to 2] **responseCondition** | **string** | The name of an existing condition in the configured endpoint, or leave blank to always execute. |  **format** | **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table. |  **user** | **string** | Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified. |  **secretKey** | **string** | Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified. |  **accountName** | **string** | The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided. |  **dataset** | **string** | Your BigQuery dataset. |  **table** | **string** | Your BigQuery table. |  **templateSuffix** | **string** | BigQuery table name suffix template. Optional. |  **projectID** | **string** | Your Google Cloud Platform project ID. Required | 

### Return type

[**LoggingBigqueryResponse**](LoggingBigqueryResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
