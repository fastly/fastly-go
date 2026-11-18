# AutomationTokensAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutomationToken**](AutomationTokensAPI.md#CreateAutomationToken) | **POST** `/automation-tokens` | Create Automation Token
[**GetAutomationTokenID**](AutomationTokensAPI.md#GetAutomationTokenID) | **GET** `/automation-tokens/{id}` | Retrieve an Automation Token by ID
[**GetAutomationTokensIDServices**](AutomationTokensAPI.md#GetAutomationTokensIDServices) | **GET** `/automation-tokens/{id}/services` | List Automation Token Services
[**GetAutomationTokensIDWorkspaces**](AutomationTokensAPI.md#GetAutomationTokensIDWorkspaces) | **GET** `/automation-tokens/{id}/workspaces` | List Automation Token Workspaces
[**ListAutomationTokens**](AutomationTokensAPI.md#ListAutomationTokens) | **GET** `/automation-tokens` | List Customer Automation Tokens
[**RevokeAutomationTokenID**](AutomationTokensAPI.md#RevokeAutomationTokenID) | **DELETE** `/automation-tokens/{id}` | Revoke an Automation Token by ID



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


## GetAutomationTokenID

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
    resp, r, err := apiClient.AutomationTokensAPI.GetAutomationTokenID(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.GetAutomationTokenID`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationTokenID`: AutomationTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.GetAutomationTokenID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationTokenIDRequest struct via the builder pattern


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


## GetAutomationTokensIDServices

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
    resp, r, err := apiClient.AutomationTokensAPI.GetAutomationTokensIDServices(ctx, id).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.GetAutomationTokensIDServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationTokensIDServices`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.GetAutomationTokensIDServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationTokensIDServicesRequest struct via the builder pattern


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


## GetAutomationTokensIDWorkspaces

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
    resp, r, err := apiClient.AutomationTokensAPI.GetAutomationTokensIDWorkspaces(ctx, id).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.GetAutomationTokensIDWorkspaces`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationTokensIDWorkspaces`: AutomationTokenWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.GetAutomationTokensIDWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The alphanumeric string identifying the automation token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationTokensIDWorkspacesRequest struct via the builder pattern


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


## RevokeAutomationTokenID

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
    resp, r, err := apiClient.AutomationTokensAPI.RevokeAutomationTokenID(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationTokensAPI.RevokeAutomationTokenID`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeAutomationTokenID`: AutomationTokenErrorResponse
    fmt.Fprintf(os.Stdout, "Response from `AutomationTokensAPI.RevokeAutomationTokenID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAutomationTokenIDRequest struct via the builder pattern


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
