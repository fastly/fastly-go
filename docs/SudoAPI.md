# SudoAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestSudoAccess**](SudoAPI.md#RequestSudoAccess) | **POST** `/sudo` | Request Sudo access



## RequestSudoAccess

Request Sudo access



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
    sudoRequest := *openapiclient.NewSudoRequest("Username_example", "Password_example") // SudoRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.SudoAPI.RequestSudoAccess(ctx).SudoRequest(sudoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SudoAPI.RequestSudoAccess`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestSudoAccess`: SudoResponse
    fmt.Fprintf(os.Stdout, "Response from `SudoAPI.RequestSudoAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSudoAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sudoRequest** | [**SudoRequest**](SudoRequest.md) |  | 

### Return type

[**SudoResponse**](SudoResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.api+json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

