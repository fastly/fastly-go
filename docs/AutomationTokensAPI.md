# AutomationTokensAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutomationToken**](AutomationTokensAPI.md#CreateAutomationToken) | **POST** `/automation-tokens` | Create Automation Token
[**GetAutomationTokenId**](AutomationTokensAPI.md#GetAutomationTokenId) | **GET** `/automation-tokens/{id}` | Retrieve an Automation Token by ID
[**GetAutomationTokensIdServices**](AutomationTokensAPI.md#GetAutomationTokensIdServices) | **GET** `/automation-tokens/{id}/services` | List Automation Token Services
[**GetAutomationTokensIdWorkspaces**](AutomationTokensAPI.md#GetAutomationTokensIdWorkspaces) | **GET** `/automation-tokens/{id}/workspaces` | List Automation Token Workspaces
[**ListAutomationTokens**](AutomationTokensAPI.md#ListAutomationTokens) | **GET** `/automation-tokens` | List Customer Automation Tokens
[**RevokeAutomationTokenId**](AutomationTokensAPI.md#RevokeAutomationTokenId) | **DELETE** `/automation-tokens/{id}` | Revoke an Automation Token by ID



## CreateAutomationToken

Create Automation Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/fastly/fastly-go/fastly"
)

func main() {
    automationTokenCreateRequest := *openapiclient.NewAutomationTokenCreateRequest("Name_example", time.Now()) // AutomationTokenCreateRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AutomationTokensAPI.CreateAutomationToken(ctx).AutomationTokenCreateRequest(automationTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.CreateAutomationToken`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomationToken`: AutomationTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.CreateAutomationToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationTokenCreateRequest** | [**AutomationTokenCreateRequest**](AutomationTokenCreateRequest.md) |  | 

### Return type

[**AutomationTokenCreateResponse**](AutomationTokenCreateResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetAutomationTokenId

Retrieve an Automation Token by ID



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
    id := "id_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AutomationTokensAPI.GetAutomationTokenId(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.GetAutomationTokenId`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationTokenId`: AutomationTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.GetAutomationTokenId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationTokenIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationTokenResponse**](AutomationTokenResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetAutomationTokensIdServices

List Automation Token Services



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
    id := "id_example" // string | 
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AutomationTokensAPI.GetAutomationTokensIdServices(ctx, id).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.GetAutomationTokensIdServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationTokensIdServices`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.GetAutomationTokensIdServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationTokensIdServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  |  **page** | **int32** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetAutomationTokensIdWorkspaces

List Automation Token Workspaces



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
    id := "id_example" // string | The alphanumeric string identifying the automation token.
    limit := int32(56) // int32 | Maximum number of workspaces to return per page. (optional) (default to 20)
    cursor := "cursor_example" // string | Pagination cursor from a previous response. Pass the `next_cursor` value to retrieve the next page, or `previous_cursor` to go back. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AutomationTokensAPI.GetAutomationTokensIdWorkspaces(ctx, id).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.GetAutomationTokensIdWorkspaces`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationTokensIdWorkspaces`: AutomationTokenWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.GetAutomationTokensIdWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The alphanumeric string identifying the automation token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationTokensIdWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of workspaces to return per page. | [default to 20] **cursor** | **string** | Pagination cursor from a previous response. Pass the `next_cursor` value to retrieve the next page, or `previous_cursor` to go back. | 

### Return type

[**AutomationTokenWorkspacesResponse**](AutomationTokenWorkspacesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListAutomationTokens

List Customer Automation Tokens



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
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AutomationTokensAPI.ListAutomationTokens(ctx).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.ListAutomationTokens`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutomationTokens`: []AutomationTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.ListAutomationTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAutomationTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  |  **page** | **int32** |  | 

### Return type

[**[]AutomationTokenResponse**](AutomationTokenResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RevokeAutomationTokenId

Revoke an Automation Token by ID



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
    id := "id_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AutomationTokensAPI.RevokeAutomationTokenId(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.RevokeAutomationTokenId`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeAutomationTokenId`: AutomationTokenErrorResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.RevokeAutomationTokenId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAutomationTokenIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationTokenErrorResponse**](AutomationTokenErrorResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

