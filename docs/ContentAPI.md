# ContentAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentCheck**](ContentAPI.md#ContentCheck) | **GET** `/content/edge_check` | Check status of content in each POP&#39;s cache



## ContentCheck

Check status of content in each POP's cache



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
    url := "https://www.example.com/foo/bar" // string | Full URL (host and path) to check on all nodes. if protocol is omitted, http will be assumed. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ContentAPI.ContentCheck(ctx).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.ContentCheck`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentCheck`: []Content
    fmt.Fprintf(os.Stdout, "Response from `ContentAPI.ContentCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | Full URL (host and path) to check on all nodes. if protocol is omitted, http will be assumed. | 

### Return type

[**[]Content**](Content.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

