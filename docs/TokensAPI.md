# TokensAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenCurrent**](TokensAPI.md#GetTokenCurrent) | **GET** `/tokens/self` | Get the current token
[**ListTokensCustomer**](TokensAPI.md#ListTokensCustomer) | **GET** `/customer/{customer_id}/tokens` | List tokens for a customer
[**ListTokensUser**](TokensAPI.md#ListTokensUser) | **GET** `/tokens` | List tokens for the authenticated user
[**RevokeToken**](TokensAPI.md#RevokeToken) | **DELETE** `/tokens/{token_id}` | Revoke a token
[**RevokeTokenCurrent**](TokensAPI.md#RevokeTokenCurrent) | **DELETE** `/tokens/self` | Revoke the current token



## GetTokenCurrent

Get the current token



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
    resp, r, err := apiClient.TokensAPI.GetTokenCurrent(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokenCurrent`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenCurrent`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokenCurrent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenCurrentRequest struct via the builder pattern



### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTokensCustomer

List tokens for a customer



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
    customerID := "customerId_example" // string | Alphanumeric string identifying the customer.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TokensAPI.ListTokensCustomer(ctx, customerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListTokensCustomer`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokensCustomer`: []TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListTokensCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string** | Alphanumeric string identifying the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokensCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListTokensUser

List tokens for the authenticated user



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
    resp, r, err := apiClient.TokensAPI.ListTokensUser(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListTokensUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokensUser`: []TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListTokensUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTokensUserRequest struct via the builder pattern



### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RevokeToken

Revoke a token



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
    tokenID := "tokenId_example" // string | Alphanumeric string identifying a token.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TokensAPI.RevokeToken(ctx, tokenID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.RevokeToken`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenID** | **string** | Alphanumeric string identifying a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RevokeTokenCurrent

Revoke the current token



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
    resp, r, err := apiClient.TokensAPI.RevokeTokenCurrent(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.RevokeTokenCurrent`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenCurrentRequest struct via the builder pattern



### Return type

 (empty response body)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
