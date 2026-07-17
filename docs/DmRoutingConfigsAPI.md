# DmRoutingConfigsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateDmRoutingConfigDraft**](DmRoutingConfigsAPI.md#ActivateDmRoutingConfigDraft) | **POST** `/domain-management/v1/routing-configs/{config_id}/activate` | Activate the draft
[**CreateDmRoutingConfig**](DmRoutingConfigsAPI.md#CreateDmRoutingConfig) | **POST** `/domain-management/v1/routing-configs` | Create a routing config
[**CreateDmRoutingConfigPath**](DmRoutingConfigsAPI.md#CreateDmRoutingConfigPath) | **POST** `/domain-management/v1/routing-configs/{config_id}/paths` | Create a path
[**CreateDmRoutingConfigRule**](DmRoutingConfigsAPI.md#CreateDmRoutingConfigRule) | **POST** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules` | Create a rule
[**DeactivateDmRoutingConfig**](DmRoutingConfigsAPI.md#DeactivateDmRoutingConfig) | **POST** `/domain-management/v1/routing-configs/{config_id}/deactivate` | Deactivate a routing config
[**DeleteDmRoutingConfig**](DmRoutingConfigsAPI.md#DeleteDmRoutingConfig) | **DELETE** `/domain-management/v1/routing-configs/{config_id}` | Delete a routing config
[**DeleteDmRoutingConfigInactiveVersions**](DmRoutingConfigsAPI.md#DeleteDmRoutingConfigInactiveVersions) | **DELETE** `/domain-management/v1/routing-configs/{config_id}/versions/inactive` | Delete inactive versions
[**DeleteDmRoutingConfigPath**](DmRoutingConfigsAPI.md#DeleteDmRoutingConfigPath) | **DELETE** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}` | Delete a path
[**DeleteDmRoutingConfigRule**](DmRoutingConfigsAPI.md#DeleteDmRoutingConfigRule) | **DELETE** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules/{rule_id}` | Delete a rule
[**DiscardDmRoutingConfigDraft**](DmRoutingConfigsAPI.md#DiscardDmRoutingConfigDraft) | **DELETE** `/domain-management/v1/routing-configs/{config_id}/draft` | Discard the draft
[**GetDmRoutingConfig**](DmRoutingConfigsAPI.md#GetDmRoutingConfig) | **GET** `/domain-management/v1/routing-configs/{config_id}` | Get a routing config
[**GetDmRoutingConfigDraftDiff**](DmRoutingConfigsAPI.md#GetDmRoutingConfigDraftDiff) | **GET** `/domain-management/v1/routing-configs/{config_id}/draft/diff` | Get the draft diff
[**GetDmRoutingConfigPath**](DmRoutingConfigsAPI.md#GetDmRoutingConfigPath) | **GET** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}` | Get a path
[**GetDmRoutingConfigRule**](DmRoutingConfigsAPI.md#GetDmRoutingConfigRule) | **GET** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules/{rule_id}` | Get a rule
[**ListDmRoutingConfigPaths**](DmRoutingConfigsAPI.md#ListDmRoutingConfigPaths) | **GET** `/domain-management/v1/routing-configs/{config_id}/paths` | List paths
[**ListDmRoutingConfigRules**](DmRoutingConfigsAPI.md#ListDmRoutingConfigRules) | **GET** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules` | List rules
[**ListDmRoutingConfigVersions**](DmRoutingConfigsAPI.md#ListDmRoutingConfigVersions) | **GET** `/domain-management/v1/routing-configs/{config_id}/versions` | List versions
[**ListDmRoutingConfigs**](DmRoutingConfigsAPI.md#ListDmRoutingConfigs) | **GET** `/domain-management/v1/routing-configs` | List routing configs
[**ReactivateDmRoutingConfigVersion**](DmRoutingConfigsAPI.md#ReactivateDmRoutingConfigVersion) | **POST** `/domain-management/v1/routing-configs/{config_id}/versions/{version_id}/activate` | Reactivate a version
[**UpdateDmRoutingConfigDraft**](DmRoutingConfigsAPI.md#UpdateDmRoutingConfigDraft) | **PATCH** `/domain-management/v1/routing-configs/{config_id}/draft` | Update the draft
[**UpdateDmRoutingConfigPath**](DmRoutingConfigsAPI.md#UpdateDmRoutingConfigPath) | **PATCH** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}` | Update a path
[**UpdateDmRoutingConfigRule**](DmRoutingConfigsAPI.md#UpdateDmRoutingConfigRule) | **PATCH** `/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules/{rule_id}` | Update a rule



## ActivateDmRoutingConfigDraft

Activate the draft



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
    configId := "configId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.ActivateDmRoutingConfigDraft(ctx, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.ActivateDmRoutingConfigDraft`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateDmRoutingConfigDraft`: RoutingConfigVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.ActivateDmRoutingConfigDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateDmRoutingConfigDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoutingConfigVersionResponse**](RoutingConfigVersionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateDmRoutingConfig

Create a routing config



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
    routingConfig := *openapiclient.NewRoutingConfig("Name_example") // RoutingConfig |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.CreateDmRoutingConfig(ctx).RoutingConfig(routingConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.CreateDmRoutingConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmRoutingConfig`: RoutingConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.CreateDmRoutingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routingConfig** | [**RoutingConfig**](RoutingConfig.md) |  | 

### Return type

[**RoutingConfigResponse**](RoutingConfigResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateDmRoutingConfigPath

Create a path



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
    configId := "configId_example" // string | 
    pathCreate := *openapiclient.NewPathCreate("Path_example") // PathCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.CreateDmRoutingConfigPath(ctx, configId).PathCreate(pathCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.CreateDmRoutingConfigPath`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmRoutingConfigPath`: PathResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.CreateDmRoutingConfigPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmRoutingConfigPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pathCreate** | [**PathCreate**](PathCreate.md) |  | 

### Return type

[**PathResponse**](PathResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CreateDmRoutingConfigRule

Create a rule



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 
    ruleCreate := *openapiclient.NewRuleCreate(*openapiclient.NewAction(openapiclient.action_type("service"), "Value_example"), []openapiclient.RoutingConfigCondition{*openapiclient.NewRoutingConfigCondition(openapiclient.condition_type("header"), openapiclient.condition_operator("equals"), "Value_example")}) // RuleCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.CreateDmRoutingConfigRule(ctx, configId, pathId).RuleCreate(ruleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.CreateDmRoutingConfigRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmRoutingConfigRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.CreateDmRoutingConfigRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmRoutingConfigRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleCreate** | [**RuleCreate**](RuleCreate.md) |  | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeactivateDmRoutingConfig

Deactivate a routing config



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
    configId := "configId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.DeactivateDmRoutingConfig(ctx, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.DeactivateDmRoutingConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateDmRoutingConfig`: RoutingConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.DeactivateDmRoutingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateDmRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoutingConfigResponse**](RoutingConfigResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDmRoutingConfig

Delete a routing config



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
    configId := "configId_example" // string | 
    force := true // bool | When `true`, allows deleting a routing config that has an active version. This is destructive — traffic routing for any paths served by the config will stop immediately. (optional) (default to false)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.DeleteDmRoutingConfig(ctx, configId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.DeleteDmRoutingConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | When `true`, allows deleting a routing config that has an active version. This is destructive — traffic routing for any paths served by the config will stop immediately. | [default to false]

### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDmRoutingConfigInactiveVersions

Delete inactive versions



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
    configId := "configId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.DeleteDmRoutingConfigInactiveVersions(ctx, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.DeleteDmRoutingConfigInactiveVersions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmRoutingConfigInactiveVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDmRoutingConfigPath

Delete a path



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.DeleteDmRoutingConfigPath(ctx, configId, pathId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.DeleteDmRoutingConfigPath`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmRoutingConfigPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDmRoutingConfigRule

Delete a rule



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 
    ruleId := "ruleId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.DeleteDmRoutingConfigRule(ctx, configId, pathId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.DeleteDmRoutingConfigRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmRoutingConfigRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DiscardDmRoutingConfigDraft

Discard the draft



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
    configId := "configId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.DiscardDmRoutingConfigDraft(ctx, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.DiscardDmRoutingConfigDraft`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscardDmRoutingConfigDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDmRoutingConfig

Get a routing config



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
    configId := "configId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.GetDmRoutingConfig(ctx, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.GetDmRoutingConfig`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmRoutingConfig`: RoutingConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.GetDmRoutingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoutingConfigResponse**](RoutingConfigResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDmRoutingConfigDraftDiff

Get the draft diff



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
    configId := "configId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.GetDmRoutingConfigDraftDiff(ctx, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.GetDmRoutingConfigDraftDiff`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmRoutingConfigDraftDiff`: DraftDiff
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.GetDmRoutingConfigDraftDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmRoutingConfigDraftDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DraftDiff**](DraftDiff.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDmRoutingConfigPath

Get a path



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.GetDmRoutingConfigPath(ctx, configId, pathId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.GetDmRoutingConfigPath`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmRoutingConfigPath`: PathResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.GetDmRoutingConfigPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmRoutingConfigPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PathResponse**](PathResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetDmRoutingConfigRule

Get a rule



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 
    ruleId := "ruleId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.GetDmRoutingConfigRule(ctx, configId, pathId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.GetDmRoutingConfigRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmRoutingConfigRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.GetDmRoutingConfigRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmRoutingConfigRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDmRoutingConfigPaths

List paths



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
    configId := "configId_example" // string | 
    path := "path_example" // string | Filter results by path pattern. The match strategy is controlled by the `match` parameter. (optional)
    match := "match_example" // string | How to match the value of the `path` filter against existing path patterns. Has no effect unless `path` is also provided. (optional) (default to "exact")
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "-created_at")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.ListDmRoutingConfigPaths(ctx, configId).Path(path).Match(match).Sort(sort).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.ListDmRoutingConfigPaths`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDmRoutingConfigPaths`: PathsResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.ListDmRoutingConfigPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDmRoutingConfigPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Filter results by path pattern. The match strategy is controlled by the `match` parameter. |  **match** | **string** | How to match the value of the `path` filter against existing path patterns. Has no effect unless `path` is also provided. | [default to &quot;exact&quot;] **sort** | **string** | The order in which to list the results. | [default to &quot;-created_at&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20]

### Return type

[**PathsResponse**](PathsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDmRoutingConfigRules

List rules



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "position")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.ListDmRoutingConfigRules(ctx, configId, pathId).Sort(sort).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.ListDmRoutingConfigRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDmRoutingConfigRules`: RulesResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.ListDmRoutingConfigRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDmRoutingConfigRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | The order in which to list the results. | [default to &quot;position&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20]

### Return type

[**RulesResponse**](RulesResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDmRoutingConfigVersions

List versions



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
    configId := "configId_example" // string | 
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "-activated_at")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.ListDmRoutingConfigVersions(ctx, configId).Sort(sort).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.ListDmRoutingConfigVersions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDmRoutingConfigVersions`: VersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.ListDmRoutingConfigVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDmRoutingConfigVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | The order in which to list the results. | [default to &quot;-activated_at&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20]

### Return type

[**VersionsResponse**](VersionsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDmRoutingConfigs

List routing configs



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
    state := []string{"State_example"} // []string | Filter configs by lifecycle state. Accepts a comma-separated list of state values (e.g. `?state=active,active-with-draft`). Returns only configs whose current state matches one of the provided values. Returns 400 if any value is not a recognised state. (optional)
    sort := "sort_example" // string | The order in which to list the results. (optional) (default to "-created_at")
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.ListDmRoutingConfigs(ctx).State(state).Sort(sort).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.ListDmRoutingConfigs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDmRoutingConfigs`: RoutingConfigsResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.ListDmRoutingConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDmRoutingConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **[]string** | Filter configs by lifecycle state. Accepts a comma-separated list of state values (e.g. `?state&#x3D;active,active-with-draft`). Returns only configs whose current state matches one of the provided values. Returns 400 if any value is not a recognised state. |  **sort** | **string** | The order in which to list the results. | [default to &quot;-created_at&quot;] **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20]

### Return type

[**RoutingConfigsResponse**](RoutingConfigsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ReactivateDmRoutingConfigVersion

Reactivate a version



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
    configId := "configId_example" // string | 
    versionId := "versionId_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.ReactivateDmRoutingConfigVersion(ctx, configId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.ReactivateDmRoutingConfigVersion`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactivateDmRoutingConfigVersion`: RoutingConfigVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.ReactivateDmRoutingConfigVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateDmRoutingConfigVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoutingConfigVersionResponse**](RoutingConfigVersionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDmRoutingConfigDraft

Update the draft



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
    configId := "configId_example" // string | 
    draftUpdate := *openapiclient.NewDraftUpdate() // DraftUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.UpdateDmRoutingConfigDraft(ctx, configId).DraftUpdate(draftUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.UpdateDmRoutingConfigDraft`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmRoutingConfigDraft`: RoutingConfigVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.UpdateDmRoutingConfigDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmRoutingConfigDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **draftUpdate** | [**DraftUpdate**](DraftUpdate.md) |  | 

### Return type

[**RoutingConfigVersionResponse**](RoutingConfigVersionResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDmRoutingConfigPath

Update a path



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 
    pathUpdate := *openapiclient.NewPathUpdate() // PathUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.UpdateDmRoutingConfigPath(ctx, configId, pathId).PathUpdate(pathUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.UpdateDmRoutingConfigPath`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmRoutingConfigPath`: PathResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.UpdateDmRoutingConfigPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmRoutingConfigPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pathUpdate** | [**PathUpdate**](PathUpdate.md) |  | 

### Return type

[**PathResponse**](PathResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDmRoutingConfigRule

Update a rule



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
    configId := "configId_example" // string | 
    pathId := "pathId_example" // string | 
    ruleId := "ruleId_example" // string | 
    ruleUpdate := *openapiclient.NewRuleUpdate() // RuleUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DmRoutingConfigsAPI.UpdateDmRoutingConfigRule(ctx, configId, pathId, ruleId).RuleUpdate(ruleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmRoutingConfigsAPI.UpdateDmRoutingConfigRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmRoutingConfigRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DmRoutingConfigsAPI.UpdateDmRoutingConfigRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 
**pathId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmRoutingConfigRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleUpdate** | [**RuleUpdate**](RuleUpdate.md) |  | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

