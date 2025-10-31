# IamServiceGroupsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceGroupServices**](IamServiceGroupsAPI.md#AddServiceGroupServices) | **POST** `/service-groups/{service_group_id}/services` | Add services in a service group
[**CreateAServiceGroup**](IamServiceGroupsAPI.md#CreateAServiceGroup) | **POST** `/service-groups` | Create a service group
[**DeleteAServiceGroup**](IamServiceGroupsAPI.md#DeleteAServiceGroup) | **DELETE** `/service-groups/{service_group_id}` | Delete a service group
[**GetAServiceGroup**](IamServiceGroupsAPI.md#GetAServiceGroup) | **GET** `/service-groups/{service_group_id}` | Get a service group
[**ListServiceGroupServices**](IamServiceGroupsAPI.md#ListServiceGroupServices) | **GET** `/service-groups/{service_group_id}/services` | List services to a service group
[**ListServiceGroups**](IamServiceGroupsAPI.md#ListServiceGroups) | **GET** `/service-groups` | List service groups
[**RemoveServiceGroupServices**](IamServiceGroupsAPI.md#RemoveServiceGroupServices) | **DELETE** `/service-groups/{service_group_id}/services` | Remove services from a service group
[**UpdateAServiceGroup**](IamServiceGroupsAPI.md#UpdateAServiceGroup) | **PATCH** `/service-groups/{service_group_id}` | Update a service group



## AddServiceGroupServices

Add services in a service group



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
    serviceGroupId := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.AddServiceGroupServices(ctx, serviceGroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.AddServiceGroupServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServiceGroupServices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.AddServiceGroupServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupId** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServiceGroupServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateAServiceGroup

Create a service group



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.CreateAServiceGroup(ctx).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.CreateAServiceGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAServiceGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.CreateAServiceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAServiceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


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
    serviceGroupId := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.DeleteAServiceGroup(ctx, serviceGroupId).Execute()
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
**serviceGroupId** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAServiceGroupRequest struct via the builder pattern


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
    serviceGroupId := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.GetAServiceGroup(ctx, serviceGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.GetAServiceGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAServiceGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.GetAServiceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupId** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAServiceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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
    serviceGroupId := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.
    perPage := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    page := int32(1) // int32 | Current page. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.ListServiceGroupServices(ctx, serviceGroupId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.ListServiceGroupServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceGroupServices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.ListServiceGroupServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupId** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceGroupServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records per page. | [default to 20] **page** | **int32** | Current page. | 

### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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
    // response from `ListServiceGroups`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RemoveServiceGroupServices

Remove services from a service group



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
    serviceGroupId := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.RemoveServiceGroupServices(ctx, serviceGroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.RemoveServiceGroupServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupId** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServiceGroupServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateAServiceGroup

Update a service group



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
    serviceGroupId := "serviceGroupId_example" // string | Alphanumeric string identifying the service group.
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.IamServiceGroupsAPI.UpdateAServiceGroup(ctx, serviceGroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamServiceGroupsAPI.UpdateAServiceGroup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAServiceGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamServiceGroupsAPI.UpdateAServiceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceGroupId** | **string** | Alphanumeric string identifying the service group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAServiceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

