# ObservabilityCustomDashboardsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](ObservabilityCustomDashboardsAPI.md#CreateDashboard) | **POST** `/observability/dashboards` | Create a new dashboard
[**DeleteDashboard**](ObservabilityCustomDashboardsAPI.md#DeleteDashboard) | **DELETE** `/observability/dashboards/{dashboard_id}` | Delete an existing dashboard
[**GetDashboard**](ObservabilityCustomDashboardsAPI.md#GetDashboard) | **GET** `/observability/dashboards/{dashboard_id}` | Retrieve a dashboard by ID
[**ListDashboards**](ObservabilityCustomDashboardsAPI.md#ListDashboards) | **GET** `/observability/dashboards` | List all custom dashboards
[**UpdateDashboard**](ObservabilityCustomDashboardsAPI.md#UpdateDashboard) | **PATCH** `/observability/dashboards/{dashboard_id}` | Update an existing dashboard



## CreateDashboard

Create a new dashboard



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
    createDashboardRequest := *openapiclient.NewCreateDashboardRequest("Name_example") // CreateDashboardRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObservabilityCustomDashboardsAPI.CreateDashboard(ctx).CreateDashboardRequest(createDashboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityCustomDashboardsAPI.CreateDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `ObservabilityCustomDashboardsAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDashboardRequest** | [**CreateDashboardRequest**](CreateDashboardRequest.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteDashboard

Delete an existing dashboard



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
    dashboardID := "2eGFXF4F4kTxd4gU39Bg3e" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObservabilityCustomDashboardsAPI.DeleteDashboard(ctx, dashboardID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityCustomDashboardsAPI.DeleteDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## GetDashboard

Retrieve a dashboard by ID



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
    dashboardID := "2eGFXF4F4kTxd4gU39Bg3e" // string | 

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObservabilityCustomDashboardsAPI.GetDashboard(ctx, dashboardID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityCustomDashboardsAPI.GetDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `ObservabilityCustomDashboardsAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListDashboards

List all custom dashboards



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
    resp, r, err := apiClient.ObservabilityCustomDashboardsAPI.ListDashboards(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityCustomDashboardsAPI.ListDashboards`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboards`: ListDashboardsResponse
    fmt.Fprintf(os.Stdout, "Response from `ObservabilityCustomDashboardsAPI.ListDashboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern



### Return type

[**ListDashboardsResponse**](ListDashboardsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateDashboard

Update an existing dashboard



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
    dashboardID := "2eGFXF4F4kTxd4gU39Bg3e" // string | 
    updateDashboardRequest := *openapiclient.NewUpdateDashboardRequest() // UpdateDashboardRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ObservabilityCustomDashboardsAPI.UpdateDashboard(ctx, dashboardID).UpdateDashboardRequest(updateDashboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityCustomDashboardsAPI.UpdateDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `ObservabilityCustomDashboardsAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDashboardRequest** | [**UpdateDashboardRequest**](UpdateDashboardRequest.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
