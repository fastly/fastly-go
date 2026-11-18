# ObjectStorageAccessKeysAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](ObjectStorageAccessKeysAPI.md#CreateAccessKey) | **POST** `/resources/object-storage/access-keys` | Create an access key
[**DeleteAccessKey**](ObjectStorageAccessKeysAPI.md#DeleteAccessKey) | **DELETE** `/resources/object-storage/access-keys/{access_key}` | Delete an access key
[**GetAccessKey**](ObjectStorageAccessKeysAPI.md#GetAccessKey) | **GET** `/resources/object-storage/access-keys/{access_key}` | Get an access key
[**ListAccessKeys**](ObjectStorageAccessKeysAPI.md#ListAccessKeys) | **GET** `/resources/object-storage/access-keys` | List access keys



## CreateAccessKey

Create an access key



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
    accessKey := *openapiclient.NewAccessKey("Description_example", "Permission_example") // AccessKey |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStorageAccessKeysAPI.CreateAccessKey(ctx).AccessKey(accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAccessKeysAPI.CreateAccessKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessKey`: AccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAccessKeysAPI.CreateAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessKey** | [**AccessKey**](AccessKey.md) |  | 

### Return type

[**AccessKeyResponse**](AccessKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteAccessKey

Delete an access key



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
    accessKey := "accessKey_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStorageAccessKeysAPI.DeleteAccessKey(ctx, accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAccessKeysAPI.DeleteAccessKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessKeyRequest struct via the builder pattern


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


## GetAccessKey

Get an access key



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
    accessKey := "accessKey_example" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObjectStorageAccessKeysAPI.GetAccessKey(ctx, accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAccessKeysAPI.GetAccessKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessKey`: AccessKey
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAccessKeysAPI.GetAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListAccessKeys

List access keys



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
    resp, r, err := apiClient.ObjectStorageAccessKeysAPI.ListAccessKeys(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAccessKeysAPI.ListAccessKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessKeys`: AccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAccessKeysAPI.ListAccessKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessKeysRequest struct via the builder pattern



### Return type

[**AccessKeyResponse**](AccessKeyResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
