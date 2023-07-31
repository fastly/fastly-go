# DomainAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDomain**](DomainAPI.md#CheckDomain) | **GET** `/service/{service_id}/version/{version_id}/domain/{domain_name}/check` | Validate DNS configuration for a single domain on a service
[**CheckDomains**](DomainAPI.md#CheckDomains) | **GET** `/service/{service_id}/version/{version_id}/domain/check_all` | Validate DNS configuration for all domains on a service
[**CreateDomain**](DomainAPI.md#CreateDomain) | **POST** `/service/{service_id}/version/{version_id}/domain` | Add a domain name to a service
[**DeleteDomain**](DomainAPI.md#DeleteDomain) | **DELETE** `/service/{service_id}/version/{version_id}/domain/{domain_name}` | Remove a domain from a service
[**GetDomain**](DomainAPI.md#GetDomain) | **GET** `/service/{service_id}/version/{version_id}/domain/{domain_name}` | Describe a domain
[**ListDomains**](DomainAPI.md#ListDomains) | **GET** `/service/{service_id}/version/{version_id}/domain` | List domains
[**UpdateDomain**](DomainAPI.md#UpdateDomain) | **PUT** `/service/{service_id}/version/{version_id}/domain/{domain_name}` | Update a domain



## CheckDomain

Validate DNS configuration for a single domain on a service



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
    domainName := "domainName_example" // string | The name of the domain or domains associated with this service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainAPI.CheckDomain(ctx, serviceID, versionID, domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CheckDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDomain`: []any
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CheckDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**domainName** | **string** | The name of the domain or domains associated with this service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CheckDomains

Validate DNS configuration for all domains on a service



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
    resp, r, err := apiClient.DomainAPI.CheckDomains(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CheckDomains`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDomains`: [][]any
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CheckDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[][]any**](array.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateDomain

Add a domain name to a service



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
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    name := "name_example" // string | The name of the domain or domains associated with this service. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainAPI.CreateDomain(ctx, serviceID, versionID).Comment(comment).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CreateDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomain`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CreateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **string** | A freeform descriptive note. |  **name** | **string** | The name of the domain or domains associated with this service. | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDomain

Remove a domain from a service



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
    domainName := "domainName_example" // string | The name of the domain or domains associated with this service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainAPI.DeleteDomain(ctx, serviceID, versionID, domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DeleteDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomain`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**domainName** | **string** | The name of the domain or domains associated with this service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


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


## GetDomain

Describe a domain



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
    domainName := "domainName_example" // string | The name of the domain or domains associated with this service.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainAPI.GetDomain(ctx, serviceID, versionID, domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**domainName** | **string** | The name of the domain or domains associated with this service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDomains

List domains



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
    resp, r, err := apiClient.DomainAPI.ListDomains(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.ListDomains`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomains`: []DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.ListDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainResponse**](DomainResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDomain

Update a domain



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
    domainName := "domainName_example" // string | The name of the domain or domains associated with this service.
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    name := "name_example" // string | The name of the domain or domains associated with this service. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainAPI.UpdateDomain(ctx, serviceID, versionID, domainName).Comment(comment).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.UpdateDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomain`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**domainName** | **string** | The name of the domain or domains associated with this service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **string** | A freeform descriptive note. |  **name** | **string** | The name of the domain or domains associated with this service. | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
