# HealthcheckAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHealthcheck**](HealthcheckAPI.md#CreateHealthcheck) | **POST** `/service/{service_id}/version/{version_id}/healthcheck` | Create a health check
[**DeleteHealthcheck**](HealthcheckAPI.md#DeleteHealthcheck) | **DELETE** `/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}` | Delete a health check
[**GetHealthcheck**](HealthcheckAPI.md#GetHealthcheck) | **GET** `/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}` | Get a health check
[**ListHealthchecks**](HealthcheckAPI.md#ListHealthchecks) | **GET** `/service/{service_id}/version/{version_id}/healthcheck` | List health checks
[**UpdateHealthcheck**](HealthcheckAPI.md#UpdateHealthcheck) | **PUT** `/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}` | Update a health check



## CreateHealthcheck

Create a health check



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
    checkInterval := int32(56) // int32 | How often to run the health check in milliseconds. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    expectedResponse := int32(56) // int32 | The status code expected from the host. (optional)
    headers := []string{"Inner_example"} // []string | Array of custom headers that will be added to the health check probes. (optional)
    host := "host_example" // string | Which host to check. (optional)
    httpVersion := "httpVersion_example" // string | Whether to use version 1.0 or 1.1 HTTP. (optional)
    initial := int32(56) // int32 | When loading a config, the initial number of probes to be seen as OK. (optional)
    method := "method_example" // string | Which HTTP method to use. (optional)
    name := "name_example" // string | The name of the health check. (optional)
    path := "path_example" // string | The path to check. (optional)
    threshold := int32(56) // int32 | How many health checks must succeed to be considered healthy. (optional)
    timeout := int32(56) // int32 | Timeout in milliseconds. (optional)
    window := int32(56) // int32 | The number of most recent health check queries to keep for this health check. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HealthcheckAPI.CreateHealthcheck(ctx, serviceID, versionID).CheckInterval(checkInterval).Comment(comment).ExpectedResponse(expectedResponse).Headers(headers).Host(host).HTTPVersion(httpVersion).Initial(initial).Method(method).Name(name).Path(path).Threshold(threshold).Timeout(timeout).Window(window).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.CreateHealthcheck`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHealthcheck`: HealthcheckResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.CreateHealthcheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHealthcheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkInterval** | **int32** | How often to run the health check in milliseconds. |  **comment** | **string** | A freeform descriptive note. |  **expectedResponse** | **int32** | The status code expected from the host. |  **headers** | **[]string** | Array of custom headers that will be added to the health check probes. |  **host** | **string** | Which host to check. |  **httpVersion** | **string** | Whether to use version 1.0 or 1.1 HTTP. |  **initial** | **int32** | When loading a config, the initial number of probes to be seen as OK. |  **method** | **string** | Which HTTP method to use. |  **name** | **string** | The name of the health check. |  **path** | **string** | The path to check. |  **threshold** | **int32** | How many health checks must succeed to be considered healthy. |  **timeout** | **int32** | Timeout in milliseconds. |  **window** | **int32** | The number of most recent health check queries to keep for this health check. | 

### Return type

[**HealthcheckResponse**](HealthcheckResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteHealthcheck

Delete a health check



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
    healthcheckName := "healthcheckName_example" // string | The name of the health check.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HealthcheckAPI.DeleteHealthcheck(ctx, serviceID, versionID, healthcheckName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.DeleteHealthcheck`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHealthcheck`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.DeleteHealthcheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**healthcheckName** | **string** | The name of the health check. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHealthcheckRequest struct via the builder pattern


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


## GetHealthcheck

Get a health check



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
    healthcheckName := "healthcheckName_example" // string | The name of the health check.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HealthcheckAPI.GetHealthcheck(ctx, serviceID, versionID, healthcheckName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.GetHealthcheck`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthcheck`: HealthcheckResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.GetHealthcheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**healthcheckName** | **string** | The name of the health check. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthcheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HealthcheckResponse**](HealthcheckResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListHealthchecks

List health checks



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
    resp, r, err := apiClient.HealthcheckAPI.ListHealthchecks(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.ListHealthchecks`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHealthchecks`: []HealthcheckResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.ListHealthchecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHealthchecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HealthcheckResponse**](HealthcheckResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateHealthcheck

Update a health check



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
    healthcheckName := "healthcheckName_example" // string | The name of the health check.
    checkInterval := int32(56) // int32 | How often to run the health check in milliseconds. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    expectedResponse := int32(56) // int32 | The status code expected from the host. (optional)
    headers := []string{"Inner_example"} // []string | Array of custom headers that will be added to the health check probes. (optional)
    host := "host_example" // string | Which host to check. (optional)
    httpVersion := "httpVersion_example" // string | Whether to use version 1.0 or 1.1 HTTP. (optional)
    initial := int32(56) // int32 | When loading a config, the initial number of probes to be seen as OK. (optional)
    method := "method_example" // string | Which HTTP method to use. (optional)
    name := "name_example" // string | The name of the health check. (optional)
    path := "path_example" // string | The path to check. (optional)
    threshold := int32(56) // int32 | How many health checks must succeed to be considered healthy. (optional)
    timeout := int32(56) // int32 | Timeout in milliseconds. (optional)
    window := int32(56) // int32 | The number of most recent health check queries to keep for this health check. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.HealthcheckAPI.UpdateHealthcheck(ctx, serviceID, versionID, healthcheckName).CheckInterval(checkInterval).Comment(comment).ExpectedResponse(expectedResponse).Headers(headers).Host(host).HTTPVersion(httpVersion).Initial(initial).Method(method).Name(name).Path(path).Threshold(threshold).Timeout(timeout).Window(window).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.UpdateHealthcheck`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHealthcheck`: HealthcheckResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.UpdateHealthcheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**healthcheckName** | **string** | The name of the health check. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHealthcheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkInterval** | **int32** | How often to run the health check in milliseconds. |  **comment** | **string** | A freeform descriptive note. |  **expectedResponse** | **int32** | The status code expected from the host. |  **headers** | **[]string** | Array of custom headers that will be added to the health check probes. |  **host** | **string** | Which host to check. |  **httpVersion** | **string** | Whether to use version 1.0 or 1.1 HTTP. |  **initial** | **int32** | When loading a config, the initial number of probes to be seen as OK. |  **method** | **string** | Which HTTP method to use. |  **name** | **string** | The name of the health check. |  **path** | **string** | The path to check. |  **threshold** | **int32** | How many health checks must succeed to be considered healthy. |  **timeout** | **int32** | Timeout in milliseconds. |  **window** | **int32** | The number of most recent health check queries to keep for this health check. | 

### Return type

[**HealthcheckResponse**](HealthcheckResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
