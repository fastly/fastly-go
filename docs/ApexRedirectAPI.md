# ApexRedirectAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApexRedirect**](ApexRedirectAPI.md#CreateApexRedirect) | **POST** `/service/{service_id}/version/{version_id}/apex-redirects` | Create an apex redirect
[**DeleteApexRedirect**](ApexRedirectAPI.md#DeleteApexRedirect) | **DELETE** `/apex-redirects/{apex_redirect_id}` | Delete an apex redirect
[**GetApexRedirect**](ApexRedirectAPI.md#GetApexRedirect) | **GET** `/apex-redirects/{apex_redirect_id}` | Get an apex redirect
[**ListApexRedirects**](ApexRedirectAPI.md#ListApexRedirects) | **GET** `/service/{service_id}/version/{version_id}/apex-redirects` | List apex redirects
[**UpdateApexRedirect**](ApexRedirectAPI.md#UpdateApexRedirect) | **PUT** `/apex-redirects/{apex_redirect_id}` | Update an apex redirect



## CreateApexRedirect

Create an apex redirect



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionID := int32(56) // int32 | Integer identifying a service version.
    serviceID2 := "serviceId_example" // string |  (optional)
    version := int32(56) // int32 |  (optional)
    createdAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    deletedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    updatedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    statusCode := int32(56) // int32 | HTTP status code used to redirect the client. (optional)
    domains := []string{"Inner_example"} // []string | Array of apex domains that should redirect to their WWW subdomain. (optional)
    featureRevision := int32(56) // int32 | Revision number of the apex redirect feature implementation. Defaults to the most recent revision. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApexRedirectAPI.CreateApexRedirect(ctx, serviceID, versionID).ServiceID2(serviceID2).Version(version).CreatedAt(createdAt).DeletedAt(deletedAt).UpdatedAt(updatedAt).StatusCode(statusCode).Domains(domains).FeatureRevision(featureRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApexRedirectAPI.CreateApexRedirect`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApexRedirect`: ApexRedirect
    fmt.Fprintf(os.Stdout, "Response from `ApexRedirectAPI.CreateApexRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApexRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceID2** | **string** |  |  **version** | **int32** |  |  **createdAt** | **time.Time** | Date and time in ISO 8601 format. |  **deletedAt** | **time.Time** | Date and time in ISO 8601 format. |  **updatedAt** | **time.Time** | Date and time in ISO 8601 format. |  **statusCode** | **int32** | HTTP status code used to redirect the client. |  **domains** | **[]string** | Array of apex domains that should redirect to their WWW subdomain. |  **featureRevision** | **int32** | Revision number of the apex redirect feature implementation. Defaults to the most recent revision. | 

### Return type

[**ApexRedirect**](ApexRedirect.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteApexRedirect

Delete an apex redirect



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
    apexRedirectID := "6QI9o6ZZrSP3y9gI0OhMwZ" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApexRedirectAPI.DeleteApexRedirect(ctx, apexRedirectID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApexRedirectAPI.DeleteApexRedirect`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApexRedirect`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ApexRedirectAPI.DeleteApexRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apexRedirectID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApexRedirectRequest struct via the builder pattern


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


## GetApexRedirect

Get an apex redirect



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
    apexRedirectID := "6QI9o6ZZrSP3y9gI0OhMwZ" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApexRedirectAPI.GetApexRedirect(ctx, apexRedirectID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApexRedirectAPI.GetApexRedirect`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApexRedirect`: ApexRedirect
    fmt.Fprintf(os.Stdout, "Response from `ApexRedirectAPI.GetApexRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apexRedirectID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApexRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApexRedirect**](ApexRedirect.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListApexRedirects

List apex redirects



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
    resp, r, err := apiClient.ApexRedirectAPI.ListApexRedirects(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApexRedirectAPI.ListApexRedirects`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApexRedirects`: []ApexRedirect
    fmt.Fprintf(os.Stdout, "Response from `ApexRedirectAPI.ListApexRedirects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApexRedirectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApexRedirect**](ApexRedirect.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateApexRedirect

Update an apex redirect



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    apexRedirectID := "6QI9o6ZZrSP3y9gI0OhMwZ" // string | 
    serviceID := "serviceId_example" // string |  (optional)
    version := int32(56) // int32 |  (optional)
    createdAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    deletedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    updatedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    statusCode := int32(56) // int32 | HTTP status code used to redirect the client. (optional)
    domains := []string{"Inner_example"} // []string | Array of apex domains that should redirect to their WWW subdomain. (optional)
    featureRevision := int32(56) // int32 | Revision number of the apex redirect feature implementation. Defaults to the most recent revision. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApexRedirectAPI.UpdateApexRedirect(ctx, apexRedirectID).ServiceID(serviceID).Version(version).CreatedAt(createdAt).DeletedAt(deletedAt).UpdatedAt(updatedAt).StatusCode(statusCode).Domains(domains).FeatureRevision(featureRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApexRedirectAPI.UpdateApexRedirect`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApexRedirect`: ApexRedirect
    fmt.Fprintf(os.Stdout, "Response from `ApexRedirectAPI.UpdateApexRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apexRedirectID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApexRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceID** | **string** |  |  **version** | **int32** |  |  **createdAt** | **time.Time** | Date and time in ISO 8601 format. |  **deletedAt** | **time.Time** | Date and time in ISO 8601 format. |  **updatedAt** | **time.Time** | Date and time in ISO 8601 format. |  **statusCode** | **int32** | HTTP status code used to redirect the client. |  **domains** | **[]string** | Array of apex domains that should redirect to their WWW subdomain. |  **featureRevision** | **int32** | Revision number of the apex redirect feature implementation. Defaults to the most recent revision. | 

### Return type

[**ApexRedirect**](ApexRedirect.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
