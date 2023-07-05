# PurgeAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkPurgeTag**](PurgeAPI.md#BulkPurgeTag) | **POST** `/service/{service_id}/purge` | Purge multiple surrogate key tags
[**PurgeAll**](PurgeAPI.md#PurgeAll) | **POST** `/service/{service_id}/purge_all` | Purge everything from a service
[**PurgeSingleURL**](PurgeAPI.md#PurgeSingleURL) | **POST** `/purge/{cached_url}` | Purge a URL
[**PurgeTag**](PurgeAPI.md#PurgeTag) | **POST** `/service/{service_id}/purge/{surrogate_key}` | Purge by surrogate key tag



## BulkPurgeTag

Purge multiple surrogate key tags



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
    fastlySoftPurge := int32(1) // int32 | If present, this header triggers the purge to be 'soft', which marks the affected object as stale rather than making it inaccessible.  Typically set to \"1\" when used, but the value is not important. (optional)
    surrogateKey := "key_1 key_2 key_3" // string | Purge multiple surrogate key tags using a request header. Not required if a JSON POST body is specified. (optional)
    purgeResponse := *openapiclient.NewPurgeResponse() // PurgeResponse |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PurgeAPI.BulkPurgeTag(ctx, serviceID).FastlySoftPurge(fastlySoftPurge).SurrogateKey(surrogateKey).PurgeResponse(purgeResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurgeAPI.BulkPurgeTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkPurgeTag`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `PurgeAPI.BulkPurgeTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPurgeTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fastlySoftPurge** | **int32** | If present, this header triggers the purge to be &#39;soft&#39;, which marks the affected object as stale rather than making it inaccessible.  Typically set to \&quot;1\&quot; when used, but the value is not important. |  **surrogateKey** | **string** | Purge multiple surrogate key tags using a request header. Not required if a JSON POST body is specified. |  **purgeResponse** | [**PurgeResponse**](PurgeResponse.md) |  | 

### Return type

**map[string]string**

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## PurgeAll

Purge everything from a service



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PurgeAPI.PurgeAll(ctx, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurgeAPI.PurgeAll`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurgeAll`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `PurgeAPI.PurgeAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeAllRequest struct via the builder pattern


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


## PurgeSingleURL

Purge a URL



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
    cachedURL := "www.example.com/path/to/object-to-purge" // string | URL of object in cache to be purged.
    fastlySoftPurge := int32(1) // int32 | If present, this header triggers the purge to be 'soft', which marks the affected object as stale rather than making it inaccessible.  Typically set to \"1\" when used, but the value is not important. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PurgeAPI.PurgeSingleURL(ctx, cachedURL).FastlySoftPurge(fastlySoftPurge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurgeAPI.PurgeSingleURL`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurgeSingleURL`: PurgeResponse
    fmt.Fprintf(os.Stdout, "Response from `PurgeAPI.PurgeSingleURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cachedURL** | **string** | URL of object in cache to be purged. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeSingleURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fastlySoftPurge** | **int32** | If present, this header triggers the purge to be &#39;soft&#39;, which marks the affected object as stale rather than making it inaccessible.  Typically set to \&quot;1\&quot; when used, but the value is not important. | 

### Return type

[**PurgeResponse**](PurgeResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## PurgeTag

Purge by surrogate key tag



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
    surrogateKey := "key_1" // string | Surrogate keys are used to efficiently purge content from cache. Instead of purging your entire site or individual URLs, you can tag related assets (like all images and descriptions associated with a single product) with surrogate keys, and these grouped URLs can be purged in a single request.
    fastlySoftPurge := int32(1) // int32 | If present, this header triggers the purge to be 'soft', which marks the affected object as stale rather than making it inaccessible.  Typically set to \"1\" when used, but the value is not important. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PurgeAPI.PurgeTag(ctx, serviceID, surrogateKey).FastlySoftPurge(fastlySoftPurge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurgeAPI.PurgeTag`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurgeTag`: PurgeResponse
    fmt.Fprintf(os.Stdout, "Response from `PurgeAPI.PurgeTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**surrogateKey** | **string** | Surrogate keys are used to efficiently purge content from cache. Instead of purging your entire site or individual URLs, you can tag related assets (like all images and descriptions associated with a single product) with surrogate keys, and these grouped URLs can be purged in a single request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fastlySoftPurge** | **int32** | If present, this header triggers the purge to be &#39;soft&#39;, which marks the affected object as stale rather than making it inaccessible.  Typically set to \&quot;1\&quot; when used, but the value is not important. | 

### Return type

[**PurgeResponse**](PurgeResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
