# NgwafSimulateAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**NgwafSimulateWafRequest**](NgwafSimulateAPI.md#NgwafSimulateWafRequest) | **POST** `/ngwaf/v1/workspaces/{workspace_id}/simulate` | Simulate a WAF request



## NgwafSimulateWafRequest

Simulate a WAF request



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
    workspaceId := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the workspace.
    wafSimulateRequest := *openapiclient.NewWafSimulateRequest("POST /login HTTP/1.1
Host: example.com
Content-Type: application/x-www-form-urlencoded

username=admin&password=1' OR '1'='1") // WafSimulateRequest | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NgwafSimulateAPI.NgwafSimulateWafRequest(ctx, workspaceId).WafSimulateRequest(wafSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgwafSimulateAPI.NgwafSimulateWafRequest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NgwafSimulateWafRequest`: WafSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `NgwafSimulateAPI.NgwafSimulateWafRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The ID of the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNgwafSimulateWafRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wafSimulateRequest** | [**WafSimulateRequest**](WafSimulateRequest.md) |  | 

### Return type

[**WafSimulateResponse**](WafSimulateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

