# DiffAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiffServiceVersions**](DiffAPI.md#DiffServiceVersions) | **GET** `/service/{service_id}/diff/from/{from_version_id}/to/{to_version_id}` | Diff two service versions



## DiffServiceVersions

Diff two service versions



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    fromVersionId := int32(1) // int32 | The version number of the service to which changes in the generated VCL are being compared. Can either be a positive number from 1 to your maximum version or a negative number from -1 down (-1 is latest version etc).
    toVersionId := int32(2) // int32 | The version number of the service from which changes in the generated VCL are being compared. Uses same numbering scheme as `from`.
    format := "format_example" // string | Optional method to format the diff field. (optional) (default to "text")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DiffAPI.DiffServiceVersions(ctx, serviceId, fromVersionId, toVersionId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiffAPI.DiffServiceVersions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiffServiceVersions`: DiffResponse
    fmt.Fprintf(os.Stdout, "Response from `DiffAPI.DiffServiceVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**fromVersionId** | **int32** | The version number of the service to which changes in the generated VCL are being compared. Can either be a positive number from 1 to your maximum version or a negative number from -1 down (-1 is latest version etc). | 
**toVersionId** | **int32** | The version number of the service from which changes in the generated VCL are being compared. Uses same numbering scheme as `from`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiffServiceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Optional method to format the diff field. | [default to &quot;text&quot;]

### Return type

[**DiffResponse**](DiffResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

