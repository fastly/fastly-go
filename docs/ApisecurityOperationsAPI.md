# ApisecurityOperationsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSecurityBulkAddTagsToOperations**](ApisecurityOperationsAPI.md#ApiSecurityBulkAddTagsToOperations) | **POST** `/api-security/v1/services/{service_id}/operations-bulk-tags` | Bulk add tags to operations
[**ApiSecurityBulkCreateOperations**](ApisecurityOperationsAPI.md#ApiSecurityBulkCreateOperations) | **POST** `/api-security/v1/services/{service_id}/operations-bulk` | Bulk create operations
[**ApiSecurityBulkDeleteOperations**](ApisecurityOperationsAPI.md#ApiSecurityBulkDeleteOperations) | **DELETE** `/api-security/v1/services/{service_id}/operations-bulk` | Bulk delete operations
[**ApiSecurityCreateOperation**](ApisecurityOperationsAPI.md#ApiSecurityCreateOperation) | **POST** `/api-security/v1/services/{service_id}/operations` | Create operation
[**ApiSecurityCreateOperationTag**](ApisecurityOperationsAPI.md#ApiSecurityCreateOperationTag) | **POST** `/api-security/v1/services/{service_id}/tags` | Create operation tag
[**ApiSecurityDeleteOperation**](ApisecurityOperationsAPI.md#ApiSecurityDeleteOperation) | **DELETE** `/api-security/v1/services/{service_id}/operations/{operation_id}` | Delete operation
[**ApiSecurityDeleteOperationTag**](ApisecurityOperationsAPI.md#ApiSecurityDeleteOperationTag) | **DELETE** `/api-security/v1/services/{service_id}/tags/{tag_id}` | Delete operation tag
[**ApiSecurityGetOperation**](ApisecurityOperationsAPI.md#ApiSecurityGetOperation) | **GET** `/api-security/v1/services/{service_id}/operations/{operation_id}` | Retrieve operation
[**ApiSecurityGetOperationTag**](ApisecurityOperationsAPI.md#ApiSecurityGetOperationTag) | **GET** `/api-security/v1/services/{service_id}/tags/{tag_id}` | Retrieve operation tag
[**ApiSecurityListDiscoveredOperations**](ApisecurityOperationsAPI.md#ApiSecurityListDiscoveredOperations) | **GET** `/api-security/v1/services/{service_id}/discovered-operations` | List discovered operations
[**ApiSecurityListOperationTags**](ApisecurityOperationsAPI.md#ApiSecurityListOperationTags) | **GET** `/api-security/v1/services/{service_id}/tags` | List operation tags
[**ApiSecurityListOperations**](ApisecurityOperationsAPI.md#ApiSecurityListOperations) | **GET** `/api-security/v1/services/{service_id}/operations` | List operations
[**ApiSecurityUpdateOperation**](ApisecurityOperationsAPI.md#ApiSecurityUpdateOperation) | **PATCH** `/api-security/v1/services/{service_id}/operations/{operation_id}` | Update operation
[**ApiSecurityUpdateOperationTag**](ApisecurityOperationsAPI.md#ApiSecurityUpdateOperationTag) | **PATCH** `/api-security/v1/services/{service_id}/tags/{tag_id}` | Update operation tag



## ApiSecurityBulkAddTagsToOperations

Bulk add tags to operations



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationBulkAddTags := *openapiclient.NewOperationBulkAddTags([]string{"op_abc123def456"}, []string{"tag_abc123def456"}) // OperationBulkAddTags |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityBulkAddTagsToOperations(ctx, serviceId).OperationBulkAddTags(operationBulkAddTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityBulkAddTagsToOperations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityBulkAddTagsToOperations`: InlineResponse2071
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityBulkAddTagsToOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityBulkAddTagsToOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationBulkAddTags** | [**OperationBulkAddTags**](OperationBulkAddTags.md) |  | 

### Return type

[**InlineResponse2071**](InlineResponse2071.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityBulkCreateOperations

Bulk create operations



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationBulkCreate := *openapiclient.NewOperationBulkCreate([]openapiclient.OperationBulkCreateOperations{*openapiclient.NewOperationBulkCreateOperations("GET", "www.example.com", "/api/v1/users/{var1}")}) // OperationBulkCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityBulkCreateOperations(ctx, serviceId).OperationBulkCreate(operationBulkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityBulkCreateOperations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityBulkCreateOperations`: InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityBulkCreateOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityBulkCreateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationBulkCreate** | [**OperationBulkCreate**](OperationBulkCreate.md) |  | 

### Return type

[**InlineResponse207**](InlineResponse207.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityBulkDeleteOperations

Bulk delete operations



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationBulkDelete := *openapiclient.NewOperationBulkDelete([]string{"op_abc123def456"}) // OperationBulkDelete |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityBulkDeleteOperations(ctx, serviceId).OperationBulkDelete(operationBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityBulkDeleteOperations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityBulkDeleteOperations`: InlineResponse2071
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityBulkDeleteOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityBulkDeleteOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationBulkDelete** | [**OperationBulkDelete**](OperationBulkDelete.md) |  | 

### Return type

[**InlineResponse2071**](InlineResponse2071.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityCreateOperation

Create operation



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationCreate := *openapiclient.NewOperationCreate("GET", "www.example.com", "/api/v1/users/{var1}") // OperationCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityCreateOperation(ctx, serviceId).OperationCreate(operationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityCreateOperation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityCreateOperation`: OperationGet
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityCreateOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityCreateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCreate** | [**OperationCreate**](OperationCreate.md) |  | 

### Return type

[**OperationGet**](OperationGet.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityCreateOperationTag

Create operation tag



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    tagCreate := *openapiclient.NewTagCreate() // TagCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityCreateOperationTag(ctx, serviceId).TagCreate(tagCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityCreateOperationTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityCreateOperationTag`: TagGet
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityCreateOperationTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityCreateOperationTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagCreate** | [**TagCreate**](TagCreate.md) |  | 

### Return type

[**TagGet**](TagGet.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityDeleteOperation

Delete operation



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationId := "op_abc123def456" // string | The unique identifier of the operation.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityDeleteOperation(ctx, serviceId, operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityDeleteOperation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**operationId** | **string** | The unique identifier of the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityDeleteOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityDeleteOperationTag

Delete operation tag



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    tagId := "tag_abc123def456" // string | The unique identifier of the operation tag.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityDeleteOperationTag(ctx, serviceId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityDeleteOperationTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**tagId** | **string** | The unique identifier of the operation tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityDeleteOperationTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityGetOperation

Retrieve operation



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationId := "op_abc123def456" // string | The unique identifier of the operation.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityGetOperation(ctx, serviceId, operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityGetOperation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityGetOperation`: OperationGet
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityGetOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**operationId** | **string** | The unique identifier of the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityGetOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationGet**](OperationGet.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityGetOperationTag

Retrieve operation tag



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    tagId := "tag_abc123def456" // string | The unique identifier of the operation tag.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityGetOperationTag(ctx, serviceId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityGetOperationTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityGetOperationTag`: TagGet
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityGetOperationTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**tagId** | **string** | The unique identifier of the operation tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityGetOperationTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagGet**](TagGet.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityListDiscoveredOperations

List discovered operations



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    method := []string{"Method_example"} // []string | Filter operations by HTTP method. (optional)
    domain := []string{"Inner_example"} // []string | Filter operations by fully-qualified domain name (exact match). (optional)
    path := "/api/v1/users" // string | Filter operations by path (exact match). (optional)
    limit := int32(100) // int32 | The maximum number of operations to return per page. (optional) (default to 100)
    page := int32(1) // int32 | The page number to return. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityListDiscoveredOperations(ctx, serviceId).Method(method).Domain(domain).Path(path).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityListDiscoveredOperations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityListDiscoveredOperations`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityListDiscoveredOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityListDiscoveredOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **method** | **[]string** | Filter operations by HTTP method. |  **domain** | **[]string** | Filter operations by fully-qualified domain name (exact match). |  **path** | **string** | Filter operations by path (exact match). |  **limit** | **int32** | The maximum number of operations to return per page. | [default to 100] **page** | **int32** | The page number to return. | [default to 0]

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityListOperationTags

List operation tags



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    limit := int32(100) // int32 | The maximum number of operations to return per page. (optional) (default to 100)
    page := int32(1) // int32 | The page number to return. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityListOperationTags(ctx, serviceId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityListOperationTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityListOperationTags`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityListOperationTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityListOperationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of operations to return per page. | [default to 100] **page** | **int32** | The page number to return. | [default to 0]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityListOperations

List operations



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    tagId := "tag_abc123def456" // string | Filter operations by operation tag ID. Only operations associated with this operation tag will be returned. (optional)
    status := "SAVED" // string | Filter operations by status. Defaults to SAVED if omitted. (optional) (default to "SAVED")
    method := []string{"Method_example"} // []string | Filter operations by HTTP method. (optional)
    domain := []string{"Inner_example"} // []string | Filter operations by fully-qualified domain name (exact match). (optional)
    path := "/api/v1/users" // string | Filter operations by path (exact match). (optional)
    limit := int32(100) // int32 | The maximum number of operations to return per page. (optional) (default to 100)
    page := int32(1) // int32 | The page number to return. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityListOperations(ctx, serviceId).TagId(tagId).Status(status).Method(method).Domain(domain).Path(path).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityListOperations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityListOperations`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityListOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagId** | **string** | Filter operations by operation tag ID. Only operations associated with this operation tag will be returned. |  **status** | **string** | Filter operations by status. Defaults to SAVED if omitted. | [default to &quot;SAVED&quot;] **method** | **[]string** | Filter operations by HTTP method. |  **domain** | **[]string** | Filter operations by fully-qualified domain name (exact match). |  **path** | **string** | Filter operations by path (exact match). |  **limit** | **int32** | The maximum number of operations to return per page. | [default to 100] **page** | **int32** | The page number to return. | [default to 0]

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityUpdateOperation

Update operation



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    operationId := "op_abc123def456" // string | The unique identifier of the operation.
    operationUpdate := *openapiclient.NewOperationUpdate() // OperationUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityUpdateOperation(ctx, serviceId, operationId).OperationUpdate(operationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityUpdateOperation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityUpdateOperation`: OperationGet
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityUpdateOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**operationId** | **string** | The unique identifier of the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityUpdateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationUpdate** | [**OperationUpdate**](OperationUpdate.md) |  | 

### Return type

[**OperationGet**](OperationGet.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ApiSecurityUpdateOperationTag

Update operation tag



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
    serviceId := "3NeCFuZNP1v0iyJ2vmYQI6" // string | The unique identifier of the service.
    tagId := "tag_abc123def456" // string | The unique identifier of the operation tag.
    body := TagBase(987) // TagBase |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ApisecurityOperationsAPI.ApiSecurityUpdateOperationTag(ctx, serviceId, tagId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisecurityOperationsAPI.ApiSecurityUpdateOperationTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSecurityUpdateOperationTag`: TagGet
    fmt.Fprintf(os.Stdout, "Response from `ApisecurityOperationsAPI.ApiSecurityUpdateOperationTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**tagId** | **string** | The unique identifier of the operation tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSecurityUpdateOperationTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **TagBase** |  | 

### Return type

[**TagGet**](TagGet.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

