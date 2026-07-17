# ClientSideProtectionAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CspCreatePage**](ClientSideProtectionAPI.md#CspCreatePage) | **POST** `/client-side-protection/v1/pages` | Create page
[**CspCreatePolicy**](ClientSideProtectionAPI.md#CspCreatePolicy) | **POST** `/client-side-protection/v1/pages/{page_id}/policies` | Create policy
[**CspCreateWebsite**](ClientSideProtectionAPI.md#CspCreateWebsite) | **POST** `/client-side-protection/v1/websites` | Create website
[**CspDeletePage**](ClientSideProtectionAPI.md#CspDeletePage) | **DELETE** `/client-side-protection/v1/pages/{page_id}` | Delete page
[**CspDeleteWebsite**](ClientSideProtectionAPI.md#CspDeleteWebsite) | **DELETE** `/client-side-protection/v1/websites/{website_id}` | Delete website
[**CspGetPage**](ClientSideProtectionAPI.md#CspGetPage) | **GET** `/client-side-protection/v1/pages/{page_id}` | Get page
[**CspGetPolicy**](ClientSideProtectionAPI.md#CspGetPolicy) | **GET** `/client-side-protection/v1/pages/{page_id}/policies/{policy_id}` | Get policy
[**CspGetScript**](ClientSideProtectionAPI.md#CspGetScript) | **GET** `/client-side-protection/v1/pages/{page_id}/scripts/{script_id}` | Get script
[**CspGetWebsite**](ClientSideProtectionAPI.md#CspGetWebsite) | **GET** `/client-side-protection/v1/websites/{website_id}` | Get website
[**CspListHeaderEvents**](ClientSideProtectionAPI.md#CspListHeaderEvents) | **GET** `/client-side-protection/v1/pages/{page_id}/events` | List header events
[**CspListHeaders**](ClientSideProtectionAPI.md#CspListHeaders) | **GET** `/client-side-protection/v1/pages/{page_id}/headers` | List security headers
[**CspListPages**](ClientSideProtectionAPI.md#CspListPages) | **GET** `/client-side-protection/v1/pages` | List pages
[**CspListPolicies**](ClientSideProtectionAPI.md#CspListPolicies) | **GET** `/client-side-protection/v1/pages/{page_id}/policies` | List policies
[**CspListPolicyReports**](ClientSideProtectionAPI.md#CspListPolicyReports) | **GET** `/client-side-protection/v1/pages/{page_id}/policies/{policy_id}/reports` | List policy reports
[**CspListScripts**](ClientSideProtectionAPI.md#CspListScripts) | **GET** `/client-side-protection/v1/pages/{page_id}/scripts` | List scripts
[**CspListWebsites**](ClientSideProtectionAPI.md#CspListWebsites) | **GET** `/client-side-protection/v1/websites` | List websites
[**CspUpdatePage**](ClientSideProtectionAPI.md#CspUpdatePage) | **PATCH** `/client-side-protection/v1/pages/{page_id}` | Update page
[**CspUpdatePolicy**](ClientSideProtectionAPI.md#CspUpdatePolicy) | **PATCH** `/client-side-protection/v1/pages/{page_id}/policies/{policy_id}` | Update policy
[**CspUpdateScript**](ClientSideProtectionAPI.md#CspUpdateScript) | **PATCH** `/client-side-protection/v1/pages/{page_id}/scripts/{script_id}` | Update script
[**CspUpdateWebsite**](ClientSideProtectionAPI.md#CspUpdateWebsite) | **PATCH** `/client-side-protection/v1/websites/{website_id}` | Update website



## CspCreatePage

Create page



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
    pageCreate := *openapiclient.NewPageCreate("WebsiteId_example", "Name_example") // PageCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspCreatePage(ctx).PageCreate(pageCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspCreatePage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspCreatePage`: Page
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspCreatePage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCspCreatePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageCreate** | [**PageCreate**](PageCreate.md) |  | 

### Return type

[**Page**](Page.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspCreatePolicy

Create policy



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    policyCreate := *openapiclient.NewPolicyCreate("Name_example", "Mode_example") // PolicyCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspCreatePolicy(ctx, pageId).PolicyCreate(policyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspCreatePolicy`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspCreatePolicy`: Policy
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspCreatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyCreate** | [**PolicyCreate**](PolicyCreate.md) |  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspCreateWebsite

Create website



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
    websiteCreate := *openapiclient.NewWebsiteCreate("Domain_example") // WebsiteCreate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspCreateWebsite(ctx).WebsiteCreate(websiteCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspCreateWebsite`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspCreateWebsite`: Website
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspCreateWebsite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCspCreateWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteCreate** | [**WebsiteCreate**](WebsiteCreate.md) |  | 

### Return type

[**Website**](Website.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspDeletePage

Delete page



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspDeletePage(ctx, pageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspDeletePage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspDeletePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspDeleteWebsite

Delete website



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
    websiteId := "2Xk9JgPCkf1NzVsNmKrECp" // string | Website identifier

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspDeleteWebsite(ctx, websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspDeleteWebsite`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** | Website identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspDeleteWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspGetPage

Get page



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspGetPage(ctx, pageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspGetPage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspGetPage`: Page
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Page**](Page.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspGetPolicy

Get policy



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    policyId := "7Cp4OlUHqj6SfAwSrQwJHu" // string | Policy identifier

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspGetPolicy(ctx, pageId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspGetPolicy`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspGetPolicy`: Policy
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspGetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**policyId** | **string** | Policy identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Policy**](Policy.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspGetScript

Get script



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    scriptId := "5An2MjSFoh4QcYvQpNuHFs" // string | Script identifier

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspGetScript(ctx, pageId, scriptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspGetScript`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspGetScript`: Script
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspGetScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**scriptId** | **string** | Script identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspGetScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Script**](Script.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspGetWebsite

Get website



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
    websiteId := "2Xk9JgPCkf1NzVsNmKrECp" // string | Website identifier

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspGetWebsite(ctx, websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspGetWebsite`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspGetWebsite`: Website
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspGetWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** | Website identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspGetWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Website**](Website.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListHeaderEvents

List header events



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListHeaderEvents(ctx, pageId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListHeaderEvents`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListHeaderEvents`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListHeaderEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspListHeaderEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListHeaders

List security headers



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListHeaders(ctx, pageId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListHeaders`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListHeaders`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspListHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListPages

List pages



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
    websiteId := "websiteId_example" // string | Filter pages by website ID (optional)
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListPages(ctx).WebsiteId(websiteId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListPages`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListPages`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCspListPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteId** | **string** | Filter pages by website ID |  **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListPolicies

List policies



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListPolicies(ctx, pageId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListPolicies`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListPolicies`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListPolicyReports

List policy reports



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    policyId := "7Cp4OlUHqj6SfAwSrQwJHu" // string | Policy identifier
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListPolicyReports(ctx, pageId, policyId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListPolicyReports`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListPolicyReports`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListPolicyReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**policyId** | **string** | Policy identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspListPolicyReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListScripts

List scripts



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListScripts(ctx, pageId).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListScripts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListScripts`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspListScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspListWebsites

List websites



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
    limit := int32(100) // int32 | Limit how many results are returned. (optional) (default to 100)
    page := int32(1) // int32 | Page number of the collection to request. (optional) (default to 0)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspListWebsites(ctx).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspListWebsites`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspListWebsites`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspListWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCspListWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit how many results are returned. | [default to 100] **page** | **int32** | Page number of the collection to request. | [default to 0]

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspUpdatePage

Update page



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    pageUpdate := *openapiclient.NewPageUpdate() // PageUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspUpdatePage(ctx, pageId).PageUpdate(pageUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspUpdatePage`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspUpdatePage`: Page
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspUpdatePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspUpdatePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageUpdate** | [**PageUpdate**](PageUpdate.md) |  | 

### Return type

[**Page**](Page.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspUpdatePolicy

Update policy



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    policyId := "7Cp4OlUHqj6SfAwSrQwJHu" // string | Policy identifier
    policyUpdate := *openapiclient.NewPolicyUpdate() // PolicyUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspUpdatePolicy(ctx, pageId, policyId).PolicyUpdate(policyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspUpdatePolicy`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspUpdatePolicy`: Policy
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspUpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**policyId** | **string** | Policy identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyUpdate** | [**PolicyUpdate**](PolicyUpdate.md) |  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspUpdateScript

Update script



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
    pageId := "3Yl0KhQDlg2OaWtOnLsFDq" // string | Page identifier
    scriptId := "5An2MjSFoh4QcYvQpNuHFs" // string | Script identifier
    scriptUpdate := *openapiclient.NewScriptUpdate() // ScriptUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspUpdateScript(ctx, pageId, scriptId).ScriptUpdate(scriptUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspUpdateScript`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspUpdateScript`: Script
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspUpdateScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**scriptId** | **string** | Script identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspUpdateScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scriptUpdate** | [**ScriptUpdate**](ScriptUpdate.md) |  | 

### Return type

[**Script**](Script.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## CspUpdateWebsite

Update website



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
    websiteId := "2Xk9JgPCkf1NzVsNmKrECp" // string | Website identifier
    websiteUpdate := *openapiclient.NewWebsiteUpdate() // WebsiteUpdate |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.ClientSideProtectionAPI.CspUpdateWebsite(ctx, websiteId).WebsiteUpdate(websiteUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientSideProtectionAPI.CspUpdateWebsite`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CspUpdateWebsite`: Website
    fmt.Fprintf(os.Stdout, "Response from `ClientSideProtectionAPI.CspUpdateWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** | Website identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCspUpdateWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteUpdate** | [**WebsiteUpdate**](WebsiteUpdate.md) |  | 

### Return type

[**Website**](Website.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

