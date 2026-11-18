# AiRuntimeControlProvidersAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArcGetBedrockSetup**](AiRuntimeControlProvidersAPI.md#ArcGetBedrockSetup) | **GET** `/ai-runtime-control/v1/providers/bedrock/setup` | Get Bedrock setup information
[**ArcGetModelsForProvider**](AiRuntimeControlProvidersAPI.md#ArcGetModelsForProvider) | **GET** `/ai-runtime-control/v1/providers/{provider_id}/models` | Get models for a provider
[**ArcGetSupportedProviders**](AiRuntimeControlProvidersAPI.md#ArcGetSupportedProviders) | **GET** `/ai-runtime-control/v1/providers` | Get supported providers



## ArcGetBedrockSetup

Get Bedrock setup information



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
    resp, r, err := apiClient.AiRuntimeControlProvidersAPI.ArcGetBedrockSetup(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProvidersAPI.ArcGetBedrockSetup`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetBedrockSetup`: ArcBedrockSetupResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProvidersAPI.ArcGetBedrockSetup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiArcGetBedrockSetupRequest struct via the builder pattern



### Return type

[**ArcBedrockSetupResponse**](ArcBedrockSetupResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcGetModelsForProvider

Get models for a provider



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
    providerID := "anthropic" // string | The ID of the provider.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.AiRuntimeControlProvidersAPI.ArcGetModelsForProvider(ctx, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProvidersAPI.ArcGetModelsForProvider`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetModelsForProvider`: ArcModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProvidersAPI.ArcGetModelsForProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerID** | **string** | The ID of the provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArcGetModelsForProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArcModelsResponse**](ArcModelsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ArcGetSupportedProviders

Get supported providers



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
    resp, r, err := apiClient.AiRuntimeControlProvidersAPI.ArcGetSupportedProviders(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiRuntimeControlProvidersAPI.ArcGetSupportedProviders`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArcGetSupportedProviders`: ArcProvidersResponse
    fmt.Fprintf(os.Stdout, "Response from `AiRuntimeControlProvidersAPI.ArcGetSupportedProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiArcGetSupportedProvidersRequest struct via the builder pattern



### Return type

[**ArcProvidersResponse**](ArcProvidersResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
