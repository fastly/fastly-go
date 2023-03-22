# IamServiceGroupsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAServiceGroup**](IamServiceGroupsAPI.md#DeleteAServiceGroup) | **DELETE** `/service-groups/{service_group_id}` | Delete a service group
[**GetAServiceGroup**](IamServiceGroupsAPI.md#GetAServiceGroup) | **GET** `/service-groups/{service_group_id}` | Get a service group
[**ListServiceGroupServices**](IamServiceGroupsAPI.md#ListServiceGroupServices) | **GET** `/service-groups/{service_group_id}/services` | List services to a service group
[**ListServiceGroups**](IamServiceGroupsAPI.md#ListServiceGroups) | **GET** `/service-groups` | List service groups



## DeleteAServiceGroup

Delete a service group



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
    serviceGroupID := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.DeleteAServiceGroup(ctx, serviceGroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.DeleteAServiceGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupID** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAServiceGroupRequest struct via the builder pattern


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


## GetAServiceGroup

Get a service group



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
    serviceGroupID := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.GetAServiceGroup(ctx, serviceGroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.GetAServiceGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAServiceGroup`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.GetAServiceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupID** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAServiceGroupRequest struct via the builder pattern


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


## ListServiceGroupServices

List services to a service group



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
    serviceGroupID := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    page := int32(1) // int32 | Current page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.ListServiceGroupServices(ctx, serviceGroupID).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.ListServiceGroupServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceGroupServices`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.ListServiceGroupServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupID** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceGroupServicesRequest struct via the builder pattern


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


## ListServiceGroups

List service groups



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
    resp, r, err := apiClient.IamServiceGroupsAPI.ListServiceGroups(ctx).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.ListServiceGroups`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceGroups`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.ListServiceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceGroupsRequest struct via the builder pattern


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
