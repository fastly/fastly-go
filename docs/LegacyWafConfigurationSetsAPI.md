# LegacyWafConfigurationSetsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListWafConfigSets**](LegacyWafConfigurationSetsAPI.md#ListWafConfigSets) | **GET** `/wafs/configuration_sets` | List configuration sets
[**ListWafsConfigSet**](LegacyWafConfigurationSetsAPI.md#ListWafsConfigSet) | **GET** `/wafs/configuration_sets/{configuration_set_id}/relationships/wafs` | List WAFs currently using a configuration set
[**UseWafConfigSet**](LegacyWafConfigurationSetsAPI.md#UseWafConfigSet) | **PATCH** `/wafs/configuration_sets/{configuration_set_id}/relationships/wafs` | Apply a configuration set to a WAF



## ListWafConfigSets

List configuration sets



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

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafConfigurationSetsAPI.ListWafConfigSets(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafConfigurationSetsAPI.ListWafConfigSets`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafConfigSets`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafConfigurationSetsAPI.ListWafConfigSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWafConfigSetsRequest struct via the builder pattern



### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListWafsConfigSet

List WAFs currently using a configuration set



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
    configurationSetID := "configurationSetId_example" // string | Alphanumeric string identifying a WAF configuration set.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafConfigurationSetsAPI.ListWafsConfigSet(ctx, configurationSetID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafConfigurationSetsAPI.ListWafsConfigSet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWafsConfigSet`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafConfigurationSetsAPI.ListWafsConfigSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationSetID** | **string** | Alphanumeric string identifying a WAF configuration set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafsConfigSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UseWafConfigSet

Apply a configuration set to a WAF



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
    configurationSetID := "configurationSetId_example" // string | Alphanumeric string identifying a WAF configuration set.
    requestBody := map[string]map[string]any{"key": map[string]any(123)} // map[string]map[string]any |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.LegacyWafConfigurationSetsAPI.UseWafConfigSet(ctx, configurationSetID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWafConfigurationSetsAPI.UseWafConfigSet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UseWafConfigSet`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `LegacyWafConfigurationSetsAPI.UseWafConfigSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationSetID** | **string** | Alphanumeric string identifying a WAF configuration set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUseWafConfigSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]any** |  | 

### Return type

**map[string]any**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
