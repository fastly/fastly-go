# PublicIpListAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFastlyIps**](PublicIpListAPI.md#ListFastlyIps) | **GET** `/public-ip-list` | List Fastly&#39;s public IPs



## ListFastlyIps

List Fastly's public IPs



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
    resp, r, err := apiClient.PublicIpListAPI.ListFastlyIps(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIpListAPI.ListFastlyIps`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFastlyIps`: PublicIpList
    fmt.Fprintf(os.Stdout, "Response from `PublicIpListAPI.ListFastlyIps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFastlyIpsRequest struct via the builder pattern



### Return type

[**PublicIpList**](PublicIpList.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

