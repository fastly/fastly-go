# LegacyWafTagAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLegacyWafTags**](LegacyWafTagAPI.md#ListLegacyWafTags) | **GET** `/wafs/tags` | List WAF tags



## ListLegacyWafTags

List WAF tags



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
    filterName := "filterName_example" // string | Limit the returned tags to a specific name. (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    include := "rules" // string | Include relationships. Optional, comma separated values. Permitted values: `rules`.  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafTagAPI.ListLegacyWafTags(ctx).FilterName(filterName).PageNumber(pageNumber).PageSize(pageSize).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafTagAPI.ListLegacyWafTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWafTags`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafTagAPI.ListLegacyWafTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWafTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string** | Limit the returned tags to a specific name. |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **include** | **string** | Include relationships. Optional, comma separated values. Permitted values: `rules`.  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
