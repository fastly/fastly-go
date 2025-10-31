# DmDomainsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDmDomain**](DmDomainsAPI.md#CreateDmDomain) | **POST** `/domain-management/v1/domains` | Create a domain
[**DeleteDmDomain**](DmDomainsAPI.md#DeleteDmDomain) | **DELETE** `/domain-management/v1/domains/{domain_id}` | Delete a domain
[**GetDmDomain**](DmDomainsAPI.md#GetDmDomain) | **GET** `/domain-management/v1/domains/{domain_id}` | Get a domain
[**ListDmDomains**](DmDomainsAPI.md#ListDmDomains) | **GET** `/domain-management/v1/domains` | List domains
[**UpdateDmDomain**](DmDomainsAPI.md#UpdateDmDomain) | **PATCH** `/domain-management/v1/domains/{domain_id}` | Update a domain



## CreateDmDomain

Create a domain



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
    requestBodyForCreate := *openapiclient.NewRequestBodyForCreate("Fqdn_example") // RequestBodyForCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmDomainsAPI.CreateDmDomain(ctx).RequestBodyForCreate(requestBodyForCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmDomainsAPI.CreateDmDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmDomain`: SuccessfulResponseAsObject
    fmt.Fprintf(os.Stdout, "Response from `DmDomainsAPI.CreateDmDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBodyForCreate** | [**RequestBodyForCreate**](RequestBodyForCreate.md) |  | 

### Return type

[**SuccessfulResponseAsObject**](SuccessfulResponseAsObject.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDmDomain

Delete a domain



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
    domainId := "domainId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmDomainsAPI.DeleteDmDomain(ctx, domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmDomainsAPI.DeleteDmDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDmDomain

Get a domain



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
    domainId := "domainId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmDomainsAPI.GetDmDomain(ctx, domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmDomainsAPI.GetDmDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmDomain`: SuccessfulResponseAsObject
    fmt.Fprintf(os.Stdout, "Response from `DmDomainsAPI.GetDmDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessfulResponseAsObject**](SuccessfulResponseAsObject.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDmDomains

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
    fqdn := "fqdn_example" // string |  (optional)
    serviceId := "serviceId_example" // string | Filter results based on a service_id. (optional)
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "fqdn")
    activated := true // bool |  (optional)
    verified := true // bool |  (optional)
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmDomainsAPI.ListDmDomains(ctx).Fqdn(fqdn).ServiceId(serviceId).Sort(sort).Activated(activated).Verified(verified).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmDomainsAPI.ListDmDomains`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDmDomains`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `DmDomainsAPI.ListDmDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDmDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fqdn** | **string** |  |  **serviceId** | **string** | Filter results based on a service_id. |  **sort** | **string** | The order in which to list the results. | [default to &quot;fqdn&quot;] **activated** | **bool** |  |  **verified** | **bool** |  |  **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20]

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDmDomain

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
    domainId := "domainId_example" // string | 
    requestBodyForUpdate := *openapiclient.NewRequestBodyForUpdate() // RequestBodyForUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmDomainsAPI.UpdateDmDomain(ctx, domainId).RequestBodyForUpdate(requestBodyForUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmDomainsAPI.UpdateDmDomain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmDomain`: SuccessfulResponseAsObject
    fmt.Fprintf(os.Stdout, "Response from `DmDomainsAPI.UpdateDmDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBodyForUpdate** | [**RequestBodyForUpdate**](RequestBodyForUpdate.md) |  | 

### Return type

[**SuccessfulResponseAsObject**](SuccessfulResponseAsObject.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

