# IamUserGroupsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserGroupMembers**](IamUserGroupsAPI.md#AddUserGroupMembers) | **POST** `/user-groups/{user_group_id}/members` | Add members to a user group
[**AddUserGroupRoles**](IamUserGroupsAPI.md#AddUserGroupRoles) | **POST** `/user-groups/{user_group_id}/roles` | Add roles to a user group
[**AddUserGroupServiceGroups**](IamUserGroupsAPI.md#AddUserGroupServiceGroups) | **POST** `/user-groups/{user_group_id}/service-groups` | Add service groups to a user group
[**CreateAUserGroup**](IamUserGroupsAPI.md#CreateAUserGroup) | **POST** `/user-groups` | Create a user group
[**DeleteAUserGroup**](IamUserGroupsAPI.md#DeleteAUserGroup) | **DELETE** `/user-groups/{user_group_id}` | Delete a user group
[**GetAUserGroup**](IamUserGroupsAPI.md#GetAUserGroup) | **GET** `/user-groups/{user_group_id}` | Get a user group
[**ListUserGroupMembers**](IamUserGroupsAPI.md#ListUserGroupMembers) | **GET** `/user-groups/{user_group_id}/members` | List members of a user group
[**ListUserGroupRoles**](IamUserGroupsAPI.md#ListUserGroupRoles) | **GET** `/user-groups/{user_group_id}/roles` | List roles in a user group
[**ListUserGroupServiceGroups**](IamUserGroupsAPI.md#ListUserGroupServiceGroups) | **GET** `/user-groups/{user_group_id}/service-groups` | List service groups in a user group
[**ListUserGroups**](IamUserGroupsAPI.md#ListUserGroups) | **GET** `/user-groups` | List user groups
[**RemoveUserGroupMembers**](IamUserGroupsAPI.md#RemoveUserGroupMembers) | **DELETE** `/user-groups/{user_group_id}/members` | Remove members of a user group
[**RemoveUserGroupRoles**](IamUserGroupsAPI.md#RemoveUserGroupRoles) | **DELETE** `/user-groups/{user_group_id}/roles` | Remove roles from a user group
[**RemoveUserGroupServiceGroups**](IamUserGroupsAPI.md#RemoveUserGroupServiceGroups) | **DELETE** `/user-groups/{user_group_id}/service-groups` | Remove service groups from a user group
[**UpdateAUserGroup**](IamUserGroupsAPI.md#UpdateAUserGroup) | **PATCH** `/user-groups/{user_group_id}` | Update a user group



## AddUserGroupMembers

Add members to a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.AddUserGroupMembers(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.AddUserGroupMembers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserGroupMembers`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.AddUserGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## AddUserGroupRoles

Add roles to a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.AddUserGroupRoles(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.AddUserGroupRoles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserGroupRoles`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.AddUserGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## AddUserGroupServiceGroups

Add service groups to a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.AddUserGroupServiceGroups(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.AddUserGroupServiceGroups`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserGroupServiceGroups`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.AddUserGroupServiceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupServiceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateAUserGroup

Create a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.CreateAUserGroup(ctx).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.CreateAUserGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAUserGroup`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.CreateAUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RemoveUserGroupMembers

Remove members of a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.RemoveUserGroupMembers(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.RemoveUserGroupMembers`: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RemoveUserGroupRoles

Remove roles from a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.RemoveUserGroupRoles(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.RemoveUserGroupRoles`: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveUserGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RemoveUserGroupServiceGroups

Remove service groups from a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.RemoveUserGroupServiceGroups(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.RemoveUserGroupServiceGroups`: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveUserGroupServiceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateAUserGroup

Update a user group



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
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamUserGroupsAPI.UpdateAUserGroup(ctx, userGroupID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserGroupsAPI.UpdateAUserGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAUserGroup`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamUserGroupsAPI.UpdateAUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupID** | **string** | Alphanumeric string identifying the user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
