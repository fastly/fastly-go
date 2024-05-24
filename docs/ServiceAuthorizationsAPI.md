# ServiceAuthorizationsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAuthorization**](ServiceAuthorizationsAPI.md#CreateServiceAuthorization) | **POST** `/service-authorizations` | Create service authorization
[**DeleteServiceAuthorization**](ServiceAuthorizationsAPI.md#DeleteServiceAuthorization) | **DELETE** `/service-authorizations/{service_authorization_id}` | Delete service authorization
[**DeleteServiceAuthorization2**](ServiceAuthorizationsAPI.md#DeleteServiceAuthorization2) | **DELETE** `/service-authorizations` | Delete service authorizations
[**ListServiceAuthorization**](ServiceAuthorizationsAPI.md#ListServiceAuthorization) | **GET** `/service-authorizations` | List service authorizations
[**ShowServiceAuthorization**](ServiceAuthorizationsAPI.md#ShowServiceAuthorization) | **GET** `/service-authorizations/{service_authorization_id}` | Show service authorization
[**UpdateServiceAuthorization**](ServiceAuthorizationsAPI.md#UpdateServiceAuthorization) | **PATCH** `/service-authorizations/{service_authorization_id}` | Update service authorization
[**UpdateServiceAuthorization2**](ServiceAuthorizationsAPI.md#UpdateServiceAuthorization2) | **PATCH** `/service-authorizations` | Update service authorizations



## CreateServiceAuthorization

Create service authorization



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
    serviceAuthorization := *openapiclient.NewServiceAuthorization() // ServiceAuthorization |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServiceAuthorizationsAPI.CreateServiceAuthorization(ctx).ServiceAuthorization(serviceAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.CreateServiceAuthorization`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAuthorization`: ServiceAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthorizationsAPI.CreateServiceAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAuthorization** | [**ServiceAuthorization**](ServiceAuthorization.md) |  | 

### Return type

[**ServiceAuthorizationResponse**](ServiceAuthorizationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteServiceAuthorization

Delete service authorization



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
    serviceAuthorizationID := "serviceAuthorizationId_example" // string | Alphanumeric string identifying a service authorization.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServiceAuthorizationsAPI.DeleteServiceAuthorization(ctx, serviceAuthorizationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.DeleteServiceAuthorization`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAuthorizationID** | **string** | Alphanumeric string identifying a service authorization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAuthorizationRequest struct via the builder pattern


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


## DeleteServiceAuthorization2

Delete service authorizations



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
    resp, r, err := apiClient.ServiceAuthorizationsAPI.DeleteServiceAuthorization2(ctx).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.DeleteServiceAuthorization2`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceAuthorization2`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthorizationsAPI.DeleteServiceAuthorization2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAuthorization2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json; ext=bulk
- **Accept**: application/vnd.api+json; ext=bulk

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListServiceAuthorization

List service authorizations



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
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServiceAuthorizationsAPI.ListServiceAuthorization(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.ListServiceAuthorization`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceAuthorization`: ServiceAuthorizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthorizationsAPI.ListServiceAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20]

### Return type

[**ServiceAuthorizationsResponse**](ServiceAuthorizationsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ShowServiceAuthorization

Show service authorization



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
    serviceAuthorizationID := "serviceAuthorizationId_example" // string | Alphanumeric string identifying a service authorization.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServiceAuthorizationsAPI.ShowServiceAuthorization(ctx, serviceAuthorizationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.ShowServiceAuthorization`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowServiceAuthorization`: ServiceAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthorizationsAPI.ShowServiceAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAuthorizationID** | **string** | Alphanumeric string identifying a service authorization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServiceAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAuthorizationResponse**](ServiceAuthorizationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateServiceAuthorization

Update service authorization



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
    serviceAuthorizationID := "serviceAuthorizationId_example" // string | Alphanumeric string identifying a service authorization.
    serviceAuthorization := *openapiclient.NewServiceAuthorization() // ServiceAuthorization |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ServiceAuthorizationsAPI.UpdateServiceAuthorization(ctx, serviceAuthorizationID).ServiceAuthorization(serviceAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.UpdateServiceAuthorization`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAuthorization`: ServiceAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthorizationsAPI.UpdateServiceAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAuthorizationID** | **string** | Alphanumeric string identifying a service authorization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAuthorization** | [**ServiceAuthorization**](ServiceAuthorization.md) |  | 

### Return type

[**ServiceAuthorizationResponse**](ServiceAuthorizationResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateServiceAuthorization2

Update service authorizations



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
    resp, r, err := apiClient.ServiceAuthorizationsAPI.UpdateServiceAuthorization2(ctx).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthorizationsAPI.UpdateServiceAuthorization2`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAuthorization2`: ServiceAuthorizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthorizationsAPI.UpdateServiceAuthorization2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAuthorization2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

[**ServiceAuthorizationsResponse**](ServiceAuthorizationsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json; ext=bulk
- **Accept**: application/vnd.api+json; ext=bulk

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
