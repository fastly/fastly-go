# VclAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomVcl**](VclAPI.md#CreateCustomVcl) | **POST** `/service/{service_id}/version/{version_id}/vcl` | Create a custom VCL file
[**DeleteCustomVcl**](VclAPI.md#DeleteCustomVcl) | **DELETE** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}` | Delete a custom VCL file
[**GetCustomVcl**](VclAPI.md#GetCustomVcl) | **GET** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}` | Get a custom VCL file
[**GetCustomVclBoilerplate**](VclAPI.md#GetCustomVclBoilerplate) | **GET** `/service/{service_id}/version/{version_id}/boilerplate` | Get boilerplate VCL
[**GetCustomVclGenerated**](VclAPI.md#GetCustomVclGenerated) | **GET** `/service/{service_id}/version/{version_id}/generated_vcl` | Get the generated VCL for a service
[**GetCustomVclRaw**](VclAPI.md#GetCustomVclRaw) | **GET** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}/download` | Download a custom VCL file
[**ListCustomVcl**](VclAPI.md#ListCustomVcl) | **GET** `/service/{service_id}/version/{version_id}/vcl` | List custom VCL files
[**SetCustomVclMain**](VclAPI.md#SetCustomVclMain) | **PUT** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}/main` | Set a custom VCL file as main
[**UpdateCustomVcl**](VclAPI.md#UpdateCustomVcl) | **PUT** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}` | Update a custom VCL file



## CreateCustomVcl

Create a custom VCL file



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
    content := "content_example" // string | The VCL code to be included. (optional)
    main := true // bool | Set to `true` when this is the main VCL, otherwise `false`. (optional)
    name := "name_example" // string | The name of this VCL. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclAPI.CreateCustomVcl(ctx, serviceID, versionID).Content(content).Main(main).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.CreateCustomVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomVcl`: VclResponse
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.CreateCustomVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomVclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | The VCL code to be included. |  **main** | **bool** | Set to `true` when this is the main VCL, otherwise `false`. |  **name** | **string** | The name of this VCL. | 

### Return type

[**VclResponse**](VclResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteCustomVcl

Delete a custom VCL file



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
    vclName := "vclName_example" // string | The name of this VCL.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclAPI.DeleteCustomVcl(ctx, serviceID, versionID, vclName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.DeleteCustomVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomVcl`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.DeleteCustomVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**vclName** | **string** | The name of this VCL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomVclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetCustomVcl

Get a custom VCL file



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
    vclName := "vclName_example" // string | The name of this VCL.
    noContent := "noContent_example" // string | Omit VCL content. (optional) (default to "0")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclAPI.GetCustomVcl(ctx, serviceID, versionID, vclName).NoContent(noContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.GetCustomVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomVcl`: VclResponse
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.GetCustomVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**vclName** | **string** | The name of this VCL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomVclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noContent** | **string** | Omit VCL content. | [default to &quot;0&quot;]

### Return type

[**VclResponse**](VclResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetCustomVclBoilerplate

Get boilerplate VCL



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
    resp, r, err := apiClient.VclAPI.GetCustomVclBoilerplate(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.GetCustomVclBoilerplate`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomVclBoilerplate`: string
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.GetCustomVclBoilerplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomVclBoilerplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetCustomVclGenerated

Get the generated VCL for a service



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
    resp, r, err := apiClient.VclAPI.GetCustomVclGenerated(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.GetCustomVclGenerated`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomVclGenerated`: VclResponse
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.GetCustomVclGenerated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomVclGeneratedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VclResponse**](VclResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetCustomVclRaw

Download a custom VCL file



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
    vclName := "vclName_example" // string | The name of this VCL.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclAPI.GetCustomVclRaw(ctx, serviceID, versionID, vclName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.GetCustomVclRaw`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomVclRaw`: string
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.GetCustomVclRaw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**vclName** | **string** | The name of this VCL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomVclRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListCustomVcl

List custom VCL files



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
    resp, r, err := apiClient.VclAPI.ListCustomVcl(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.ListCustomVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomVcl`: []VclResponse
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.ListCustomVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomVclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]VclResponse**](VclResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SetCustomVclMain

Set a custom VCL file as main



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
    vclName := "vclName_example" // string | The name of this VCL.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclAPI.SetCustomVclMain(ctx, serviceID, versionID, vclName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.SetCustomVclMain`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCustomVclMain`: VclResponse
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.SetCustomVclMain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**vclName** | **string** | The name of this VCL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCustomVclMainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VclResponse**](VclResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateCustomVcl

Update a custom VCL file



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
    vclName := "vclName_example" // string | The name of this VCL.
    content := "content_example" // string | The VCL code to be included. (optional)
    main := true // bool | Set to `true` when this is the main VCL, otherwise `false`. (optional)
    name := "name_example" // string | The name of this VCL. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclAPI.UpdateCustomVcl(ctx, serviceID, versionID, vclName).Content(content).Main(main).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclAPI.UpdateCustomVcl`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomVcl`: VclResponse
    fmt.Fprintf(os.Stdout, "Response from `VclAPI.UpdateCustomVcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**vclName** | **string** | The name of this VCL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomVclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | The VCL code to be included. |  **main** | **bool** | Set to `true` when this is the main VCL, otherwise `false`. |  **name** | **string** | The name of this VCL. | 

### Return type

[**VclResponse**](VclResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
