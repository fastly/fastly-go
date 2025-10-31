# IamRolesAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamV1RolesGet**](IamRolesAPI.md#IamV1RolesGet) | **GET** `/iam/v1/roles/{role_id}` | Get IAM role by ID
[**IamV1RolesList**](IamRolesAPI.md#IamV1RolesList) | **GET** `/iam/v1/roles` | List IAM roles



## IamV1RolesGet

Get IAM role by ID



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
    roleId := "Q4rXe9vN1szK8a2fUjYtWp" // string | Alphanumeric string identifying the role.
    include := "include_example" // string | Include related data (i.e., permissions). (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamRolesAPI.IamV1RolesGet(ctx, roleId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamRolesAPI.IamV1RolesGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamV1RolesGet`: IamV1RoleResponse
    fmt.Fprintf(os.Stdout, "Response from `IamRolesAPI.IamV1RolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Alphanumeric string identifying the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamV1RolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Include related data (i.e., permissions). | 

### Return type

[**IamV1RoleResponse**](IamV1RoleResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## IamV1RolesList

List IAM roles



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
    limit := int32(56) // int32 | Number of results per page. The maximum is 1000. (optional) (default to 100)
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamRolesAPI.IamV1RolesList(ctx).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamRolesAPI.IamV1RolesList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamV1RolesList`: IamV1RoleListResponse
    fmt.Fprintf(os.Stdout, "Response from `IamRolesAPI.IamV1RolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamV1RolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results per page. The maximum is 1000. | [default to 100] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. | 

### Return type

[**IamV1RoleListResponse**](IamV1RoleListResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

