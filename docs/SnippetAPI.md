# SnippetAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnippet**](SnippetAPI.md#CreateSnippet) | **POST** `/service/{service_id}/version/{version_id}/snippet` | Create a snippet
[**DeleteSnippet**](SnippetAPI.md#DeleteSnippet) | **DELETE** `/service/{service_id}/version/{version_id}/snippet/{snippet_name}` | Delete a snippet
[**GetSnippet**](SnippetAPI.md#GetSnippet) | **GET** `/service/{service_id}/version/{version_id}/snippet/{snippet_name}` | Get a versioned snippet
[**GetSnippetDynamic**](SnippetAPI.md#GetSnippetDynamic) | **GET** `/service/{service_id}/snippet/{snippet_id}` | Get a dynamic snippet
[**ListSnippets**](SnippetAPI.md#ListSnippets) | **GET** `/service/{service_id}/version/{version_id}/snippet` | List snippets
[**UpdateSnippet**](SnippetAPI.md#UpdateSnippet) | **PUT** `/service/{service_id}/version/{version_id}/snippet/{snippet_name}` | Update a versioned snippet
[**UpdateSnippetDynamic**](SnippetAPI.md#UpdateSnippetDynamic) | **PUT** `/service/{service_id}/snippet/{snippet_id}` | Update a dynamic snippet



## CreateSnippet

Create a snippet



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
    name := "name_example" // string | The name for the snippet. (optional)
    resourceType := "resourceType_example" // string | The location in generated VCL where the snippet should be placed. (optional)
    content := "content_example" // string | The VCL code that specifies exactly what the snippet does. (optional)
    priority := "priority_example" // string | Priority determines execution order. Lower numbers execute first. (optional) (default to "100")
    dynamic := "dynamic_example" // string | Sets the snippet version. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SnippetAPI.CreateSnippet(ctx, serviceID, versionID).Name(name).ResourceType(resourceType).Content(content).Priority(priority).Dynamic(dynamic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.CreateSnippet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnippet`: SnippetResponsePost
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.CreateSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the snippet. |  **resourceType** | **string** | The location in generated VCL where the snippet should be placed. |  **content** | **string** | The VCL code that specifies exactly what the snippet does. |  **priority** | **string** | Priority determines execution order. Lower numbers execute first. | [default to &quot;100&quot;] **dynamic** | **string** | Sets the snippet version. | 

### Return type

[**SnippetResponsePost**](SnippetResponsePost.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteSnippet

Delete a snippet



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
    snippetName := "snippetName_example" // string | The name for the snippet.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SnippetAPI.DeleteSnippet(ctx, serviceID, versionID, snippetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.DeleteSnippet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnippet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.DeleteSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**snippetName** | **string** | The name for the snippet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnippetRequest struct via the builder pattern


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


## GetSnippet

Get a versioned snippet



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
    snippetName := "snippetName_example" // string | The name for the snippet.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SnippetAPI.GetSnippet(ctx, serviceID, versionID, snippetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.GetSnippet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnippet`: SnippetResponse
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.GetSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**snippetName** | **string** | The name for the snippet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnippetResponse**](SnippetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetSnippetDynamic

Get a dynamic snippet



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
    snippetID := "snippetId_example" // string | Alphanumeric string identifying a VCL Snippet.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SnippetAPI.GetSnippetDynamic(ctx, serviceID, snippetID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.GetSnippetDynamic`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnippetDynamic`: SnippetResponse
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.GetSnippetDynamic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**snippetID** | **string** | Alphanumeric string identifying a VCL Snippet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetDynamicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnippetResponse**](SnippetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListSnippets

List snippets



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
    resp, r, err := apiClient.SnippetAPI.ListSnippets(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.ListSnippets`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSnippets`: []SnippetResponse
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.ListSnippets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SnippetResponse**](SnippetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateSnippet

Update a versioned snippet



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
    snippetName := "snippetName_example" // string | The name for the snippet.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SnippetAPI.UpdateSnippet(ctx, serviceID, versionID, snippetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.UpdateSnippet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSnippet`: SnippetResponse
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.UpdateSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**snippetName** | **string** | The name for the snippet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnippetResponse**](SnippetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateSnippetDynamic

Update a dynamic snippet



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
    snippetID := "snippetId_example" // string | Alphanumeric string identifying a VCL Snippet.
    name := "name_example" // string | The name for the snippet. (optional)
    resourceType := "resourceType_example" // string | The location in generated VCL where the snippet should be placed. (optional)
    content := "content_example" // string | The VCL code that specifies exactly what the snippet does. (optional)
    priority := "priority_example" // string | Priority determines execution order. Lower numbers execute first. (optional) (default to "100")
    dynamic := "dynamic_example" // string | Sets the snippet version. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SnippetAPI.UpdateSnippetDynamic(ctx, serviceID, snippetID).Name(name).ResourceType(resourceType).Content(content).Priority(priority).Dynamic(dynamic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnippetAPI.UpdateSnippetDynamic`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSnippetDynamic`: SnippetResponse
    fmt.Fprintf(os.Stdout, "Response from `SnippetAPI.UpdateSnippetDynamic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**snippetID** | **string** | Alphanumeric string identifying a VCL Snippet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnippetDynamicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name for the snippet. |  **resourceType** | **string** | The location in generated VCL where the snippet should be placed. |  **content** | **string** | The VCL code that specifies exactly what the snippet does. |  **priority** | **string** | Priority determines execution order. Lower numbers execute first. | [default to &quot;100&quot;] **dynamic** | **string** | Sets the snippet version. | 

### Return type

[**SnippetResponse**](SnippetResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
