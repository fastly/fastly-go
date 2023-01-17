# VclDiffAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VclDiffServiceVersions**](VclDiffAPI.md#VclDiffServiceVersions) | **GET** `/service/{service_id}/vcl/diff/from/{from_version_id}/to/{to_version_id}` | Get a comparison of the VCL changes between two service versions



## VclDiffServiceVersions

Get a comparison of the VCL changes between two service versions



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
    fromVersionID := int32(1) // int32 | The version number of the service to which changes in the generated VCL are being compared. Can either be a positive number from 1 to your maximum version or a negative number from -1 down (-1 is latest version etc).
    toVersionID := int32(2) // int32 | The version number of the service from which changes in the generated VCL are being compared. Uses same numbering scheme as `from`.
    format := "format_example" // string | Optional method to format the diff field. (optional) (default to "text")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.VclDiffAPI.VclDiffServiceVersions(ctx, serviceID, fromVersionID, toVersionID).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VclDiffAPI.VclDiffServiceVersions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VclDiffServiceVersions`: VclDiff
    fmt.Fprintf(os.Stdout, "Response from `VclDiffAPI.VclDiffServiceVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**fromVersionID** | **int32** | The version number of the service to which changes in the generated VCL are being compared. Can either be a positive number from 1 to your maximum version or a negative number from -1 down (-1 is latest version etc). | 
**toVersionID** | **int32** | The version number of the service from which changes in the generated VCL are being compared. Uses same numbering scheme as `from`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVclDiffServiceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Optional method to format the diff field. | [default to &quot;text&quot;]

### Return type

[**VclDiff**](VclDiff.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
