# DomainResearchAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainStatus**](DomainResearchAPI.md#DomainStatus) | **GET** `/domain-management/v1/tools/status` | Domain status
[**SuggestDomains**](DomainResearchAPI.md#SuggestDomains) | **GET** `/domain-management/v1/tools/suggest` | Suggest domains



## DomainStatus

Domain status



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
    domain := "acmecoffee.shop" // string | 
    scope := "estimate" // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainResearchAPI.DomainStatus(ctx).Domain(domain).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainResearchAPI.DomainStatus`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainStatus`: Status
    fmt.Fprintf(os.Stdout, "Response from `DomainResearchAPI.DomainStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** |  |  **scope** | **string** |  | 

### Return type

[**Status**](Status.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## SuggestDomains

Suggest domains



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
    query := "foo%20bar" // string | 
    defaults := "club" // string |  (optional)
    keywords := "food,kitchen" // string |  (optional)
    location := "de" // string |  (optional)
    vendor := "dnsimple.com" // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DomainResearchAPI.SuggestDomains(ctx).Query(query).Defaults(defaults).Keywords(keywords).Location(location).Vendor(vendor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainResearchAPI.SuggestDomains`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestDomains`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `DomainResearchAPI.SuggestDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  |  **defaults** | **string** |  |  **keywords** | **string** |  |  **location** | **string** |  |  **vendor** | **string** |  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

