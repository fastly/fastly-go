# ACLAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateACL**](AclAPI.md#CreateACL) | **POST** `/service/{service_id}/version/{version_id}/acl` | Create a new ACL
[**DeleteACL**](AclAPI.md#DeleteACL) | **DELETE** `/service/{service_id}/version/{version_id}/acl/{acl_name}` | Delete an ACL
[**GetACL**](AclAPI.md#GetACL) | **GET** `/service/{service_id}/version/{version_id}/acl/{acl_name}` | Describe an ACL
[**ListACLs**](AclAPI.md#ListACLs) | **GET** `/service/{service_id}/version/{version_id}/acl` | List ACLs
[**UpdateACL**](AclAPI.md#UpdateACL) | **PUT** `/service/{service_id}/version/{version_id}/acl/{acl_name}` | Update an ACL



## CreateACL

Create a new ACL



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
    name := "name_example" // string | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLAPI.CreateACL(ctx, serviceID, versionID).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.CreateACL`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateACL`: ACLResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLAPI.CreateACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. | 

### Return type

[**ACLResponse**](AclResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteACL

Delete an ACL



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
    aclName := "aclName_example" // string | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLAPI.DeleteACL(ctx, serviceID, versionID, aclName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.DeleteACL`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteACL`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ACLAPI.DeleteACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**aclName** | **string** | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACLRequest struct via the builder pattern


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


## GetACL

Describe an ACL



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
    aclName := "aclName_example" // string | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLAPI.GetACL(ctx, serviceID, versionID, aclName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.GetACL`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACL`: ACLResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLAPI.GetACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**aclName** | **string** | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ACLResponse**](AclResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListACLs

List ACLs



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
    resp, r, err := apiClient.ACLAPI.ListACLs(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ListACLs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListACLs`: []ACLResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ListACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ACLResponse**](AclResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateACL

Update an ACL



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
    aclName := "aclName_example" // string | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace.
    name := "name_example" // string | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLAPI.UpdateACL(ctx, serviceID, versionID, aclName).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.UpdateACL`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateACL`: ACLResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLAPI.UpdateACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**aclName** | **string** | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. | 

### Return type

[**ACLResponse**](AclResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
