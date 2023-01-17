# IamUserGroupsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAUserGroup**](IamUserGroupsAPI.md#DeleteAUserGroup) | **DELETE** `/user-groups/{user_group_id}` | Delete a user group
[**GetAUserGroup**](IamUserGroupsAPI.md#GetAUserGroup) | **GET** `/user-groups/{user_group_id}` | Get a user group
[**ListUserGroupMembers**](IamUserGroupsAPI.md#ListUserGroupMembers) | **GET** `/user-groups/{user_group_id}/members` | List members of a user group
[**ListUserGroupRoles**](IamUserGroupsAPI.md#ListUserGroupRoles) | **GET** `/user-groups/{user_group_id}/roles` | List roles in a user group
[**ListUserGroupServiceGroups**](IamUserGroupsAPI.md#ListUserGroupServiceGroups) | **GET** `/user-groups/{user_group_id}/service-groups` | List service groups in a user group
[**ListUserGroups**](IamUserGroupsAPI.md#ListUserGroups) | **GET** `/user-groups` | List user groups



## DeleteAUserGroup

Delete a user group



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
    userGroupID := "userGroupId_example" // string | Alphanumeric string identifying the user group.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.DeleteAUserGroup(ctx, userGroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.DeleteAUserGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetAUserGroup

Get a user group



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
    userGroupID := "userGroupId_example" // string | Alphanumeric string identifying the user group.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.GetAUserGroup(ctx, userGroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.GetAUserGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAUserGroup`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.GetAUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListUserGroupMembers

List members of a user group



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
    userGroupID := "userGroupId_example" // string | Alphanumeric string identifying the user group.
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    page := int32(1) // int32 | Current page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.ListUserGroupMembers(ctx, userGroupID).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.ListUserGroupMembers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGroupMembers`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.ListUserGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records per page. | [default to 20] **page** | **int32** | Current page. | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListUserGroupRoles

List roles in a user group



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
    userGroupID := "userGroupId_example" // string | Alphanumeric string identifying the user group.
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    page := int32(1) // int32 | Current page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.ListUserGroupRoles(ctx, userGroupID).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.ListUserGroupRoles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGroupRoles`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.ListUserGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records per page. | [default to 20] **page** | **int32** | Current page. | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListUserGroupServiceGroups

List service groups in a user group



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
    userGroupID := "userGroupId_example" // string | Alphanumeric string identifying the user group.
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    page := int32(1) // int32 | Current page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.ListUserGroupServiceGroups(ctx, userGroupID).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.ListUserGroupServiceGroups`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGroupServiceGroups`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.ListUserGroupServiceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupServiceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records per page. | [default to 20] **page** | **int32** | Current page. | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListUserGroups

List user groups



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
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    page := int32(1) // int32 | Current page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.ListUserGroups(ctx).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.ListUserGroups`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGroups`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.ListUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records per page. | [default to 20] **page** | **int32** | Current page. | 

### Return type

**map[string]any**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
