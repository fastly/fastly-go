# BotManagementAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotManagementGetBot**](BotManagementAPI.md#BotManagementGetBot) | **GET** `/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots/{bot_id}` | Get a bot
[**BotManagementGetCategory**](BotManagementAPI.md#BotManagementGetCategory) | **GET** `/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}` | Get a category
[**BotManagementGetWorkspace**](BotManagementAPI.md#BotManagementGetWorkspace) | **GET** `/bot-management/v1/workspaces/{workspace_id}` | Get a workspace
[**BotManagementListBots**](BotManagementAPI.md#BotManagementListBots) | **GET** `/bot-management/v1/workspaces/{workspace_id}/policy/bots` | List all bots in a workspace
[**BotManagementListCategories**](BotManagementAPI.md#BotManagementListCategories) | **GET** `/bot-management/v1/workspaces/{workspace_id}/policy/categories` | List categories
[**BotManagementListCategoryBots**](BotManagementAPI.md#BotManagementListCategoryBots) | **GET** `/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots` | List bots in a category
[**BotManagementListWorkspaces**](BotManagementAPI.md#BotManagementListWorkspaces) | **GET** `/bot-management/v1/workspaces` | List workspaces
[**BotManagementUpdateBot**](BotManagementAPI.md#BotManagementUpdateBot) | **PATCH** `/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots/{bot_id}` | Set a bot action
[**BotManagementUpdateCategory**](BotManagementAPI.md#BotManagementUpdateCategory) | **PATCH** `/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}` | Set a category action
[**BotManagementUpdateCategoryBots**](BotManagementAPI.md#BotManagementUpdateCategoryBots) | **PATCH** `/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots` | Set the action for all bots in a category
[**BotManagementUpdateWorkspace**](BotManagementAPI.md#BotManagementUpdateWorkspace) | **PATCH** `/bot-management/v1/workspaces/{workspace_id}` | Update a workspace



## BotManagementGetBot

Get a bot



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    categoryID := "ai_crawlers" // string | The ID of the bot category.
    botID := "gptbot" // string | The ID of the bot.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementGetBot(ctx, workspaceID, categoryID, botID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementGetBot`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementGetBot`: BotManagementBotResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementGetBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 
**categoryID** | **string** | The ID of the bot category. | 
**botID** | **string** | The ID of the bot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementGetBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotManagementBotResponse**](BotManagementBotResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementGetCategory

Get a category



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    categoryID := "ai_crawlers" // string | The ID of the bot category.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementGetCategory(ctx, workspaceID, categoryID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementGetCategory`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementGetCategory`: BotManagementCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementGetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 
**categoryID** | **string** | The ID of the bot category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotManagementCategoryResponse**](BotManagementCategoryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementGetWorkspace

Get a workspace



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementGetWorkspace(ctx, workspaceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementGetWorkspace`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementGetWorkspace`: BotManagementWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementGetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotManagementWorkspaceResponse**](BotManagementWorkspaceResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementListBots

List all bots in a workspace



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    cursor := "MTAwOjA" // string | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementListBots(ctx, workspaceID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementListBots`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementListBots`: BotManagementWorkspaceBotsResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementListBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementListBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **cursor** | **string** | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. | 

### Return type

[**BotManagementWorkspaceBotsResponse**](BotManagementWorkspaceBotsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementListCategories

List categories



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    cursor := "MTAwOjA" // string | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementListCategories(ctx, workspaceID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementListCategories`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementListCategories`: BotManagementCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementListCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementListCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **cursor** | **string** | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. | 

### Return type

[**BotManagementCategoriesResponse**](BotManagementCategoriesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementListCategoryBots

List bots in a category



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    categoryID := "ai_crawlers" // string | The ID of the bot category.
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    cursor := "MTAwOjA" // string | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementListCategoryBots(ctx, workspaceID, categoryID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementListCategoryBots`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementListCategoryBots`: BotManagementBotsResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementListCategoryBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 
**categoryID** | **string** | The ID of the bot category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementListCategoryBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **cursor** | **string** | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. | 

### Return type

[**BotManagementBotsResponse**](BotManagementBotsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementListWorkspaces

List workspaces



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
    serviceID := "6oMkSHsKCMBjQJPLHgMxu1" // string | Filter results to the workspace attached to this service. Returns an empty list if the service is not attached to a workspace. (optional)
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    cursor := "MTAwOjA" // string | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementListWorkspaces(ctx).ServiceID(serviceID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementListWorkspaces`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementListWorkspaces`: BotManagementWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceID** | **string** | Filter results to the workspace attached to this service. Returns an empty list if the service is not attached to a workspace. |  **limit** | **int32** | Limit how many results are returned. | [default to 100] **cursor** | **string** | Cursor for pagination. Use the `next_cursor` value from a previous response to request the next page of results. | 

### Return type

[**BotManagementWorkspacesResponse**](BotManagementWorkspacesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementUpdateBot

Set a bot action



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    categoryID := "ai_crawlers" // string | The ID of the bot category.
    botID := "gptbot" // string | The ID of the bot.
    botManagementBot := *openapiclient.NewBotManagementBot(*openapiclient.NewBotManagementBotAction("inherit")) // BotManagementBot | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementUpdateBot(ctx, workspaceID, categoryID, botID).BotManagementBot(botManagementBot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementUpdateBot`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementUpdateBot`: BotManagementBotResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementUpdateBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 
**categoryID** | **string** | The ID of the bot category. | 
**botID** | **string** | The ID of the bot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementUpdateBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botManagementBot** | [**BotManagementBot**](BotManagementBot.md) |  | 

### Return type

[**BotManagementBotResponse**](BotManagementBotResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementUpdateCategory

Set a category action



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    categoryID := "ai_crawlers" // string | The ID of the bot category.
    botManagementCategory := *openapiclient.NewBotManagementCategory(*openapiclient.NewBotManagementAction("block")) // BotManagementCategory | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementUpdateCategory(ctx, workspaceID, categoryID).BotManagementCategory(botManagementCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementUpdateCategory`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementUpdateCategory`: BotManagementCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementUpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 
**categoryID** | **string** | The ID of the bot category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botManagementCategory** | [**BotManagementCategory**](BotManagementCategory.md) |  | 

### Return type

[**BotManagementCategoryResponse**](BotManagementCategoryResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementUpdateCategoryBots

Set the action for all bots in a category



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    categoryID := "ai_crawlers" // string | The ID of the bot category.
    botManagementCategoryBots := *openapiclient.NewBotManagementCategoryBots(*openapiclient.NewBotManagementCategoryBotsAction("inherit")) // BotManagementCategoryBots | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementUpdateCategoryBots(ctx, workspaceID, categoryID).BotManagementCategoryBots(botManagementCategoryBots).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementUpdateCategoryBots`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementUpdateCategoryBots`: BotManagementBotsResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementUpdateCategoryBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 
**categoryID** | **string** | The ID of the bot category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementUpdateCategoryBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botManagementCategoryBots** | [**BotManagementCategoryBots**](BotManagementCategoryBots.md) |  | 

### Return type

[**BotManagementBotsResponse**](BotManagementBotsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## BotManagementUpdateWorkspace

Update a workspace



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
    workspaceID := "SU1Z0isxPaozGVKXdv0eY" // string | The ID of the bot management workspace.
    botManagementWorkspace := *openapiclient.NewBotManagementWorkspace() // BotManagementWorkspace | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BotManagementAPI.BotManagementUpdateWorkspace(ctx, workspaceID).BotManagementWorkspace(botManagementWorkspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotManagementAPI.BotManagementUpdateWorkspace`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotManagementUpdateWorkspace`: BotManagementWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `BotManagementAPI.BotManagementUpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceID** | **string** | The ID of the bot management workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotManagementUpdateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botManagementWorkspace** | [**BotManagementWorkspace**](BotManagementWorkspace.md) |  | 

### Return type

[**BotManagementWorkspaceResponse**](BotManagementWorkspaceResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
