# DomainOwnershipsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDomainOwnerships**](DomainOwnershipsAPI.md#ListDomainOwnerships) | **GET** `/domain-ownerships` | List domain-ownerships



## ListDomainOwnerships

List domain-ownerships



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
    resp, r, err := apiClient.DomainOwnershipsAPI.ListDomainOwnerships(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainOwnershipsAPI.ListDomainOwnerships`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomainOwnerships`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DomainOwnershipsAPI.ListDomainOwnerships`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainOwnershipsRequest struct via the builder pattern



### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
