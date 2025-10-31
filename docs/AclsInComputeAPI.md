# AclsInComputeAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeAclCreateAcls**](AclsInComputeAPI.md#ComputeAclCreateAcls) | **POST** `/resources/acls` | Create a new ACL
[**ComputeAclDeleteSAclId**](AclsInComputeAPI.md#ComputeAclDeleteSAclId) | **DELETE** `/resources/acls/{acl_id}` | Delete an ACL
[**ComputeAclListAclEntries**](AclsInComputeAPI.md#ComputeAclListAclEntries) | **GET** `/resources/acls/{acl_id}/entries` | List an ACL
[**ComputeAclListAcls**](AclsInComputeAPI.md#ComputeAclListAcls) | **GET** `/resources/acls` | List ACLs
[**ComputeAclListAclsSAclId**](AclsInComputeAPI.md#ComputeAclListAclsSAclId) | **GET** `/resources/acls/{acl_id}` | Describe an ACL
[**ComputeAclLookupAcls**](AclsInComputeAPI.md#ComputeAclLookupAcls) | **GET** `/resources/acls/{acl_id}/entry/{acl_ip}` | Lookup an ACL
[**ComputeAclUpdateAcls**](AclsInComputeAPI.md#ComputeAclUpdateAcls) | **PATCH** `/resources/acls/{acl_id}/entries` | Update an ACL



## ComputeAclCreateAcls

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
    computeAclCreateAclsRequest := *openapiclient.NewComputeAclCreateAclsRequest() // ComputeAclCreateAclsRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclCreateAcls(ctx).ComputeAclCreateAclsRequest(computeAclCreateAclsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclCreateAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeAclCreateAcls`: ComputeAclCreateAclsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclsInComputeAPI.ComputeAclCreateAcls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclCreateAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computeAclCreateAclsRequest** | [**ComputeAclCreateAclsRequest**](ComputeAclCreateAclsRequest.md) |  | 

### Return type

[**ComputeAclCreateAclsResponse**](ComputeAclCreateAclsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeAclDeleteSAclId

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
    aclId := "aclId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclDeleteSAclId(ctx, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclDeleteSAclId`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclDeleteSAclIdRequest struct via the builder pattern


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


## ComputeAclListAclEntries

List an ACL



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
    aclId := "aclId_example" // string | 
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 100)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclListAclEntries(ctx, aclId).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclListAclEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeAclListAclEntries`: ComputeAclListEntries
    fmt.Fprintf(os.Stdout, "Response from `AclsInComputeAPI.ComputeAclListAclEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclListAclEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 100]

### Return type

[**ComputeAclListEntries**](ComputeAclListEntries.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeAclListAcls

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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclListAcls(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclListAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeAclListAcls`: ComputeAclList
    fmt.Fprintf(os.Stdout, "Response from `AclsInComputeAPI.ComputeAclListAcls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclListAclsRequest struct via the builder pattern



### Return type

[**ComputeAclList**](ComputeAclList.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeAclListAclsSAclId

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
    aclId := "aclId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclListAclsSAclId(ctx, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclListAclsSAclId`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeAclListAclsSAclId`: ComputeAclCreateAclsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclsInComputeAPI.ComputeAclListAclsSAclId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclListAclsSAclIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputeAclCreateAclsResponse**](ComputeAclCreateAclsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeAclLookupAcls

Lookup an ACL



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
    aclId := "aclId_example" // string | 
    aclIp := "aclIp_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclLookupAcls(ctx, aclId, aclIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclLookupAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeAclLookupAcls`: ComputeAclLookup
    fmt.Fprintf(os.Stdout, "Response from `AclsInComputeAPI.ComputeAclLookupAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclId** | **string** |  | 
**aclIp** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclLookupAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputeAclLookup**](ComputeAclLookup.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeAclUpdateAcls

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
    aclId := "aclId_example" // string | 
    computeAclUpdate := *openapiclient.NewComputeAclUpdate() // ComputeAclUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclsInComputeAPI.ComputeAclUpdateAcls(ctx, aclId).ComputeAclUpdate(computeAclUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsInComputeAPI.ComputeAclUpdateAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeAclUpdateAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computeAclUpdate** | [**ComputeAclUpdate**](ComputeAclUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

