# AclEntryAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateAclEntries**](AclEntryAPI.md#BulkUpdateAclEntries) | **PATCH** `/service/{service_id}/acl/{acl_id}/entries` | Update multiple ACL entries
[**CreateAclEntry**](AclEntryAPI.md#CreateAclEntry) | **POST** `/service/{service_id}/acl/{acl_id}/entry` | Create an ACL entry
[**DeleteAclEntry**](AclEntryAPI.md#DeleteAclEntry) | **DELETE** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Delete an ACL entry
[**GetAclEntry**](AclEntryAPI.md#GetAclEntry) | **GET** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Describe an ACL entry
[**ListAclEntries**](AclEntryAPI.md#ListAclEntries) | **GET** `/service/{service_id}/acl/{acl_id}/entries` | List ACL entries
[**UpdateAclEntry**](AclEntryAPI.md#UpdateAclEntry) | **PATCH** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Update an ACL entry



## BulkUpdateAclEntries

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    aclId := "aclId_example" // string | Alphanumeric string identifying a ACL.
    bulkUpdateAclEntriesRequest := *openapiclient.NewBulkUpdateAclEntriesRequest() // BulkUpdateAclEntriesRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclEntryAPI.BulkUpdateAclEntries(ctx, serviceId, aclId).BulkUpdateAclEntriesRequest(bulkUpdateAclEntriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclEntryAPI.BulkUpdateAclEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateAclEntries`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AclEntryAPI.BulkUpdateAclEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**aclId** | **string** | Alphanumeric string identifying a ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateAclEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateAclEntriesRequest** | [**BulkUpdateAclEntriesRequest**](BulkUpdateAclEntriesRequest.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateAclEntry

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    aclId := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntry := *openapiclient.NewAclEntry() // AclEntry |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclEntryAPI.CreateAclEntry(ctx, serviceId, aclId).AclEntry(aclEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclEntryAPI.CreateAclEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAclEntry`: AclEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `AclEntryAPI.CreateAclEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**aclId** | **string** | Alphanumeric string identifying a ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAclEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclEntry** | [**AclEntry**](AclEntry.md) |  | 

### Return type

[**AclEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteAclEntry

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    aclId := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntryId := "aclEntryId_example" // string | Alphanumeric string identifying an ACL Entry.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclEntryAPI.DeleteAclEntry(ctx, serviceId, aclId, aclEntryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclEntryAPI.DeleteAclEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAclEntry`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AclEntryAPI.DeleteAclEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**aclId** | **string** | Alphanumeric string identifying a ACL. | 
**aclEntryId** | **string** | Alphanumeric string identifying an ACL Entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAclEntryRequest struct via the builder pattern


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


## GetAclEntry

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    aclId := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntryId := "aclEntryId_example" // string | Alphanumeric string identifying an ACL Entry.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclEntryAPI.GetAclEntry(ctx, serviceId, aclId, aclEntryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclEntryAPI.GetAclEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAclEntry`: AclEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `AclEntryAPI.GetAclEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**aclId** | **string** | Alphanumeric string identifying a ACL. | 
**aclEntryId** | **string** | Alphanumeric string identifying an ACL Entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAclEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AclEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListAclEntries

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    aclId := "aclId_example" // string | Alphanumeric string identifying a ACL.
    page := int32(1) // int32 | Current page. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created" // string | Field on which to sort. (optional) (default to "created")
    direction := "ascend" // string | Direction in which to sort results. (optional) (default to "ascend")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclEntryAPI.ListAclEntries(ctx, serviceId, aclId).Page(page).PerPage(perPage).Sort(sort).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclEntryAPI.ListAclEntries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAclEntries`: []AclEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `AclEntryAPI.ListAclEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**aclId** | **string** | Alphanumeric string identifying a ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAclEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page. |  **perPage** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | Field on which to sort. | [default to &quot;created&quot;] **direction** | **string** | Direction in which to sort results. | [default to &quot;ascend&quot;]

### Return type

[**[]AclEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateAclEntry

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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    aclId := "aclId_example" // string | Alphanumeric string identifying a ACL.
    aclEntryId := "aclEntryId_example" // string | Alphanumeric string identifying an ACL Entry.
    aclEntry := *openapiclient.NewAclEntry() // AclEntry |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AclEntryAPI.UpdateAclEntry(ctx, serviceId, aclId, aclEntryId).AclEntry(aclEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclEntryAPI.UpdateAclEntry`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAclEntry`: AclEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `AclEntryAPI.UpdateAclEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**aclId** | **string** | Alphanumeric string identifying a ACL. | 
**aclEntryId** | **string** | Alphanumeric string identifying an ACL Entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAclEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclEntry** | [**AclEntry**](AclEntry.md) |  | 

### Return type

[**AclEntryResponse**](AclEntryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

