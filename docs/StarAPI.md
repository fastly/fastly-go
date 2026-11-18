# StarAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceStar**](StarAPI.md#CreateServiceStar) | **POST** `/stars` | Create a star
[**DeleteServiceStar**](StarAPI.md#DeleteServiceStar) | **DELETE** `/stars/{star_id}` | Delete a star
[**GetServiceStar**](StarAPI.md#GetServiceStar) | **GET** `/stars/{star_id}` | Get a star
[**ListServiceStars**](StarAPI.md#ListServiceStars) | **GET** `/stars` | List stars



## CreateServiceStar

Create a star



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
    star := *openapiclient.NewStar() // Star |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.StarAPI.CreateServiceStar(ctx).Star(star).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StarAPI.CreateServiceStar`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceStar`: StarResponse
    fmt.Fprintf(os.Stdout, "Response from `StarAPI.CreateServiceStar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **star** | [**Star**](Star.md) |  | 

### Return type

[**StarResponse**](StarResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteServiceStar

Delete a star



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
    starID := "starId_example" // string | Alphanumeric string identifying a star.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.StarAPI.DeleteServiceStar(ctx, starID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StarAPI.DeleteServiceStar`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**starID** | **string** | Alphanumeric string identifying a star. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceStarRequest struct via the builder pattern


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


## GetServiceStar

Get a star



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
    starID := "starId_example" // string | Alphanumeric string identifying a star.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.StarAPI.GetServiceStar(ctx, starID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StarAPI.GetServiceStar`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceStar`: StarResponse
    fmt.Fprintf(os.Stdout, "Response from `StarAPI.GetServiceStar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**starID** | **string** | Alphanumeric string identifying a star. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StarResponse**](StarResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListServiceStars

List stars



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.StarAPI.ListServiceStars(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StarAPI.ListServiceStars`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceStars`: Pagination
    fmt.Fprintf(os.Stdout, "Response from `StarAPI.ListServiceStars`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceStarsRequest struct via the builder pattern



### Return type

[**Pagination**](pagination.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
