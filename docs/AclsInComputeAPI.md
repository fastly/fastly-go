# ACLsInComputeAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeACLCreateAcls**](AclsInComputeAPI.md#ComputeACLCreateAcls) | **POST** `/resources/acls` | Create a new ACL
[**ComputeACLDeleteSAclID**](AclsInComputeAPI.md#ComputeACLDeleteSAclID) | **DELETE** `/resources/acls/{acl_id}` | Delete an ACL
[**ComputeACLListAclEntries**](AclsInComputeAPI.md#ComputeACLListAclEntries) | **GET** `/resources/acls/{acl_id}/entries` | List an ACL
[**ComputeACLListAcls**](AclsInComputeAPI.md#ComputeACLListAcls) | **GET** `/resources/acls` | List ACLs
[**ComputeACLListAclsSAclID**](AclsInComputeAPI.md#ComputeACLListAclsSAclID) | **GET** `/resources/acls/{acl_id}` | Describe an ACL
[**ComputeACLLookupAcls**](AclsInComputeAPI.md#ComputeACLLookupAcls) | **GET** `/resources/acls/{acl_id}/entry/{acl_ip}` | Lookup an ACL
[**ComputeACLUpdateAcls**](AclsInComputeAPI.md#ComputeACLUpdateAcls) | **PATCH** `/resources/acls/{acl_id}/entries` | Update an ACL



## ComputeACLCreateAcls

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
    computeACLCreateAclsRequest := *openapiclient.NewComputeACLCreateAclsRequest() // ComputeACLCreateAclsRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLCreateAcls(ctx).ComputeACLCreateAclsRequest(computeACLCreateAclsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLCreateAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeACLCreateAcls`: ComputeACLCreateAclsResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLsInComputeAPI.ComputeACLCreateAcls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLCreateAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computeACLCreateAclsRequest** | [**ComputeACLCreateAclsRequest**](ComputeACLCreateAclsRequest.md) |  | 

### Return type

[**ComputeACLCreateAclsResponse**](ComputeACLCreateAclsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeACLDeleteSAclID

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
    aclID := "aclId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLDeleteSAclID(ctx, aclID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLDeleteSAclID`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLDeleteSAclIDRequest struct via the builder pattern


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


## ComputeACLListAclEntries

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
    aclID := "aclId_example" // string | 
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 100)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLListAclEntries(ctx, aclID).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLListAclEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeACLListAclEntries`: ComputeACLListEntries
    fmt.Fprintf(os.Stdout, "Response from `ACLsInComputeAPI.ComputeACLListAclEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLListAclEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  |  **limit** | **int32** |  | [default to 100]

### Return type

[**ComputeACLListEntries**](ComputeACLListEntries.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeACLListAcls

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
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLListAcls(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLListAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeACLListAcls`: []ComputeACLCreateAclsResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLsInComputeAPI.ComputeACLListAcls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLListAclsRequest struct via the builder pattern



### Return type

[**[]ComputeACLCreateAclsResponse**](ComputeACLCreateAclsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeACLListAclsSAclID

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
    aclID := "aclId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLListAclsSAclID(ctx, aclID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLListAclsSAclID`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeACLListAclsSAclID`: ComputeACLCreateAclsResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLsInComputeAPI.ComputeACLListAclsSAclID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLListAclsSAclIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputeACLCreateAclsResponse**](ComputeACLCreateAclsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeACLLookupAcls

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
    aclID := "aclId_example" // string | 
    aclIP := "aclIp_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLLookupAcls(ctx, aclID, aclIP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLLookupAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeACLLookupAcls`: ComputeACLLookup
    fmt.Fprintf(os.Stdout, "Response from `ACLsInComputeAPI.ComputeACLLookupAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclID** | **string** |  | 
**aclIP** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLLookupAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputeACLLookup**](ComputeACLLookup.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ComputeACLUpdateAcls

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
    aclID := "aclId_example" // string | 
    computeACLUpdateEntry := []openapiclient.ComputeACLUpdateEntry{*openapiclient.NewComputeACLUpdateEntry()} // []ComputeACLUpdateEntry |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLsInComputeAPI.ComputeACLUpdateAcls(ctx, aclID).ComputeACLUpdateEntry(computeACLUpdateEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLsInComputeAPI.ComputeACLUpdateAcls`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeACLUpdateAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computeACLUpdateEntry** | [**[]ComputeACLUpdateEntry**](ComputeACLUpdateEntry.md) |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
