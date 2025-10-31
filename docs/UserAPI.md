# UserAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserAPI.md#CreateUser) | **POST** `/user` | Create a user
[**DeleteUser**](UserAPI.md#DeleteUser) | **DELETE** `/user/{user_id}` | Delete a user
[**GetCurrentUser**](UserAPI.md#GetCurrentUser) | **GET** `/current_user` | Get the current user
[**GetUser**](UserAPI.md#GetUser) | **GET** `/user/{user_id}` | Get a user
[**RequestPasswordReset**](UserAPI.md#RequestPasswordReset) | **POST** `/user/{user_login}/password/request_reset` | Request a password reset
[**UpdateUser**](UserAPI.md#UpdateUser) | **PUT** `/user/{user_id}` | Update a user
[**UpdateUserPassword**](UserAPI.md#UpdateUserPassword) | **POST** `/current_user/password` | Update the user&#39;s password



## CreateUser

Create a user



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
    login := "login_example" // string |  (optional)
    name := "name_example" // string | The real life name of the user. (optional)
    limitServices := true // bool | Indicates that the user has limited access to the customer's services. (optional)
    locked := true // bool | Indicates whether the is account is locked for editing or not. (optional)
    requireNewPassword := true // bool | Indicates if a new password is required at next login. (optional)
    role := openapiclient.role_user("user") // RoleUser |  (optional)
    roles := []string{"6bKsDElwPt8vZXCArszK9x"} // []string | A list of role IDs assigned to the user. (optional)
    twoFactorAuthEnabled := true // bool | Indicates if 2FA is enabled on the user. (optional)
    twoFactorSetupRequired := true // bool | Indicates if 2FA is required by the user's customer account. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.UserAPI.CreateUser(ctx).Login(login).Name(name).LimitServices(limitServices).Locked(locked).RequireNewPassword(requireNewPassword).Role(role).Roles(roles).TwoFactorAuthEnabled(twoFactorAuthEnabled).TwoFactorSetupRequired(twoFactorSetupRequired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | **string** |  |  **name** | **string** | The real life name of the user. |  **limitServices** | **bool** | Indicates that the user has limited access to the customer&#39;s services. |  **locked** | **bool** | Indicates whether the is account is locked for editing or not. |  **requireNewPassword** | **bool** | Indicates if a new password is required at next login. |  **role** | [**RoleUser**](RoleUser.md) |  |  **roles** | **[]string** | A list of role IDs assigned to the user. |  **twoFactorAuthEnabled** | **bool** | Indicates if 2FA is enabled on the user. |  **twoFactorSetupRequired** | **bool** | Indicates if 2FA is required by the user&#39;s customer account. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteUser

Delete a user



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
    userId := "userId_example" // string | Alphanumeric string identifying the user.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.UserAPI.DeleteUser(ctx, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Alphanumeric string identifying the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetCurrentUser

Get the current user



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
    resp, r, err := apiClient.UserAPI.GetCurrentUser(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetCurrentUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern



### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetUser

Get a user



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
    userId := "userId_example" // string | Alphanumeric string identifying the user.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.UserAPI.GetUser(ctx, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Alphanumeric string identifying the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## RequestPasswordReset

Request a password reset



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
    userLogin := "userLogin_example" // string | The login associated with the user (typically, an email address).

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.UserAPI.RequestPasswordReset(ctx, userLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RequestPasswordReset`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestPasswordReset`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.RequestPasswordReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userLogin** | **string** | The login associated with the user (typically, an email address). | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateUser

Update a user



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
    userId := "userId_example" // string | Alphanumeric string identifying the user.
    login := "login_example" // string |  (optional)
    name := "name_example" // string | The real life name of the user. (optional)
    limitServices := true // bool | Indicates that the user has limited access to the customer's services. (optional)
    locked := true // bool | Indicates whether the is account is locked for editing or not. (optional)
    requireNewPassword := true // bool | Indicates if a new password is required at next login. (optional)
    role := openapiclient.role_user("user") // RoleUser |  (optional)
    roles := []string{"6bKsDElwPt8vZXCArszK9x"} // []string | A list of role IDs assigned to the user. (optional)
    twoFactorAuthEnabled := true // bool | Indicates if 2FA is enabled on the user. (optional)
    twoFactorSetupRequired := true // bool | Indicates if 2FA is required by the user's customer account. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.UserAPI.UpdateUser(ctx, userId).Login(login).Name(name).LimitServices(limitServices).Locked(locked).RequireNewPassword(requireNewPassword).Role(role).Roles(roles).TwoFactorAuthEnabled(twoFactorAuthEnabled).TwoFactorSetupRequired(twoFactorSetupRequired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Alphanumeric string identifying the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | **string** |  |  **name** | **string** | The real life name of the user. |  **limitServices** | **bool** | Indicates that the user has limited access to the customer&#39;s services. |  **locked** | **bool** | Indicates whether the is account is locked for editing or not. |  **requireNewPassword** | **bool** | Indicates if a new password is required at next login. |  **role** | [**RoleUser**](RoleUser.md) |  |  **roles** | **[]string** | A list of role IDs assigned to the user. |  **twoFactorAuthEnabled** | **bool** | Indicates if 2FA is enabled on the user. |  **twoFactorSetupRequired** | **bool** | Indicates if 2FA is required by the user&#39;s customer account. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateUserPassword

Update the user's password



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
    oldPassword := "oldPassword_example" // string | The user's current password. (optional)
    newPassword := "newPassword_example" // string | The user's new password. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.UserAPI.UpdateUserPassword(ctx).OldPassword(oldPassword).NewPassword(newPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserPassword`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPassword`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldPassword** | **string** | The user&#39;s current password. |  **newPassword** | **string** | The user&#39;s new password. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

