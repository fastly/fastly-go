# ACLEntryAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateACLEntries**](AclEntryAPI.md#BulkUpdateACLEntries) | **PATCH** `/service/{service_id}/acl/{acl_id}/entries` | Update multiple ACL entries
[**CreateACLEntry**](AclEntryAPI.md#CreateACLEntry) | **POST** `/service/{service_id}/acl/{acl_id}/entry` | Create an ACL entry
[**DeleteACLEntry**](AclEntryAPI.md#DeleteACLEntry) | **DELETE** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Delete an ACL entry
[**GetACLEntry**](AclEntryAPI.md#GetACLEntry) | **GET** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Describe an ACL entry
[**ListACLEntries**](AclEntryAPI.md#ListACLEntries) | **GET** `/service/{service_id}/acl/{acl_id}/entries` | List ACL entries
[**UpdateACLEntry**](AclEntryAPI.md#UpdateACLEntry) | **PATCH** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Update an ACL entry



## BulkUpdateACLEntries

Update multiple ACL entries



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
    aclID := "aclId_example" // string | Alphanumeric string identifying a ACL.
    bulkUpdateACLEntriesRequest := *openapiclient.NewBulkUpdateACLEntriesRequest() // BulkUpdateACLEntriesRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLEntryAPI.BulkUpdateACLEntries(ctx, serviceID, aclID).BulkUpdateACLEntriesRequest(bulkUpdateACLEntriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLEntryAPI.BulkUpdateACLEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateACLEntries`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ACLEntryAPI.BulkUpdateACLEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**aclID** | **string** | Alphanumeric string identifying a ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateACLEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateACLEntriesRequest** | [**BulkUpdateACLEntriesRequest**](BulkUpdateACLEntriesRequest.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateACLEntry

Create an ACL entry



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
    aclID := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntry := *openapiclient.NewACLEntry() // ACLEntry |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLEntryAPI.CreateACLEntry(ctx, serviceID, aclID).ACLEntry(aclEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLEntryAPI.CreateACLEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateACLEntry`: ACLEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLEntryAPI.CreateACLEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**aclID** | **string** | Alphanumeric string identifying a ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateACLEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclEntry** | [**ACLEntry**](AclEntry.md) |  | 

### Return type

[**ACLEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteACLEntry

Delete an ACL entry



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
    aclID := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntryID := "aclEntryId_example" // string | Alphanumeric string identifying an ACL Entry.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLEntryAPI.DeleteACLEntry(ctx, serviceID, aclID, aclEntryID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLEntryAPI.DeleteACLEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteACLEntry`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ACLEntryAPI.DeleteACLEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**aclID** | **string** | Alphanumeric string identifying a ACL. | 
**aclEntryID** | **string** | Alphanumeric string identifying an ACL Entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACLEntryRequest struct via the builder pattern


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


## GetACLEntry

Describe an ACL entry



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
    aclID := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntryID := "aclEntryId_example" // string | Alphanumeric string identifying an ACL Entry.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLEntryAPI.GetACLEntry(ctx, serviceID, aclID, aclEntryID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLEntryAPI.GetACLEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACLEntry`: ACLEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLEntryAPI.GetACLEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**aclID** | **string** | Alphanumeric string identifying a ACL. | 
**aclEntryID** | **string** | Alphanumeric string identifying an ACL Entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACLEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ACLEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListACLEntries

List ACL entries



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
    aclID := "aclId_example" // string | Alphanumeric string identifying a ACL.
    page := int32(1) // int32 | Current page. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created" // string | Field on which to sort. (optional) (default to "created")
    direction := "ascend" // string | Direction in which to sort results. (optional) (default to "ascend")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLEntryAPI.ListACLEntries(ctx, serviceID, aclID).Page(page).PerPage(perPage).Sort(sort).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLEntryAPI.ListACLEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListACLEntries`: []ACLEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLEntryAPI.ListACLEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**aclID** | **string** | Alphanumeric string identifying a ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListACLEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page. |  **perPage** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | Field on which to sort. | [default to &quot;created&quot;] **direction** | **string** | Direction in which to sort results. | [default to &quot;ascend&quot;]

### Return type

[**[]ACLEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateACLEntry

Update an ACL entry



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
    aclID := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntryID := "aclEntryId_example" // string | Alphanumeric string identifying an ACL Entry.
    aclEntry := *openapiclient.NewACLEntry() // ACLEntry |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ACLEntryAPI.UpdateACLEntry(ctx, serviceID, aclID, aclEntryID).ACLEntry(aclEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLEntryAPI.UpdateACLEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateACLEntry`: ACLEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `ACLEntryAPI.UpdateACLEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**aclID** | **string** | Alphanumeric string identifying a ACL. | 
**aclEntryID** | **string** | Alphanumeric string identifying an ACL Entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateACLEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclEntry** | [**ACLEntry**](AclEntry.md) |  | 

### Return type

[**ACLEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
