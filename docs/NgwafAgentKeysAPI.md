# NgwafAgentKeysAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**NgwafListAgentKeys**](NgwafAgentKeysAPI.md#NgwafListAgentKeys) | **GET** `/ngwaf/v1/workspaces/{workspace_id}/agent-keys` | List agent keys for a workspace



## NgwafListAgentKeys

List agent keys for a workspace



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.NgwafAgentKeysAPI.NgwafListAgentKeys(ctx, workspaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgwafAgentKeysAPI.NgwafListAgentKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NgwafListAgentKeys`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `NgwafAgentKeysAPI.NgwafListAgentKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The ID of the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNgwafListAgentKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

