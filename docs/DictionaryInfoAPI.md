# DictionaryInfoAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDictionaryInfo**](DictionaryInfoAPI.md#GetDictionaryInfo) | **GET** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_id}/info` | Get edge dictionary metadata



## GetDictionaryInfo

Get edge dictionary metadata



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    dictionaryID := "dictionaryId_example" // string | Alphanumeric string identifying a Dictionary.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DictionaryInfoAPI.GetDictionaryInfo(ctx, serviceID, versionID, dictionaryID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryInfoAPI.GetDictionaryInfo`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDictionaryInfo`: DictionaryInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `DictionaryInfoAPI.GetDictionaryInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**dictionaryID** | **string** | Alphanumeric string identifying a Dictionary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDictionaryInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DictionaryInfoResponse**](DictionaryInfoResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
