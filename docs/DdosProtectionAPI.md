# DdosProtectionAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**DdosProtectionEventGet**](DdosProtectionAPI.md#DdosProtectionEventGet) | **GET** `/ddos-protection/v1/events/{event_id}` | Get event by ID
[**DdosProtectionEventList**](DdosProtectionAPI.md#DdosProtectionEventList) | **GET** `/ddos-protection/v1/events` | Get events
[**DdosProtectionEventRuleList**](DdosProtectionAPI.md#DdosProtectionEventRuleList) | **GET** `/ddos-protection/v1/events/{event_id}/rules` | Get all rules for an event
[**DdosProtectionRuleGet**](DdosProtectionAPI.md#DdosProtectionRuleGet) | **GET** `/ddos-protection/v1/rules/{rule_id}` | Get a rule by ID
[**DdosProtectionTrafficStatsRuleGet**](DdosProtectionAPI.md#DdosProtectionTrafficStatsRuleGet) | **GET** `/ddos-protection/v1/events/{event_id}/rules/{rule_id}/traffic-stats` | Get traffic stats for a rule



## DdosProtectionEventGet

Get event by ID



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
    eventID := "eventId_example" // string | Unique ID of the event.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DdosProtectionAPI.DdosProtectionEventGet(ctx, eventID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DdosProtectionAPI.DdosProtectionEventGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosProtectionEventGet`: DdosProtectionEvent
    fmt.Fprintf(os.Stdout, "Response from `DdosProtectionAPI.DdosProtectionEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventID** | **string** | Unique ID of the event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosProtectionEventGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosProtectionEvent**](DdosProtectionEvent.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DdosProtectionEventList

Get events



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
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)
    serviceID := "serviceId_example" // string | Filter results based on a service_id. (optional)
    from := time.Now() // time.Time | Represents the start of a date-time range expressed in RFC 3339 format. (optional)
    to := time.Now() // time.Time | Represents the end of a date-time range expressed in RFC 3339 format. (optional)
    name := "name_example" // string |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DdosProtectionAPI.DdosProtectionEventList(ctx).Cursor(cursor).Limit(limit).ServiceID(serviceID).From(from).To(to).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DdosProtectionAPI.DdosProtectionEventList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosProtectionEventList`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `DdosProtectionAPI.DdosProtectionEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDdosProtectionEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20] **serviceID** | **string** | Filter results based on a service_id. |  **from** | **time.Time** | Represents the start of a date-time range expressed in RFC 3339 format. |  **to** | **time.Time** | Represents the end of a date-time range expressed in RFC 3339 format. |  **name** | **string** |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DdosProtectionEventRuleList

Get all rules for an event



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
    eventID := "eventId_example" // string | Unique ID of the event.
    cursor := "cursor_example" // string | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. (optional)
    limit := int32(56) // int32 | Limit how many results are returned. (optional) (default to 20)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DdosProtectionAPI.DdosProtectionEventRuleList(ctx, eventID).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DdosProtectionAPI.DdosProtectionEventRuleList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosProtectionEventRuleList`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `DdosProtectionAPI.DdosProtectionEventRuleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventID** | **string** | Unique ID of the event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosProtectionEventRuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor value from the `next_cursor` field of a previous response, used to retrieve the next page. To request the first page, this should be empty. |  **limit** | **int32** | Limit how many results are returned. | [default to 20]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DdosProtectionRuleGet

Get a rule by ID



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
    ruleID := "ruleId_example" // string | Unique ID of the rule.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DdosProtectionAPI.DdosProtectionRuleGet(ctx, ruleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DdosProtectionAPI.DdosProtectionRuleGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosProtectionRuleGet`: DdosProtectionRule
    fmt.Fprintf(os.Stdout, "Response from `DdosProtectionAPI.DdosProtectionRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleID** | **string** | Unique ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosProtectionRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosProtectionRule**](DdosProtectionRule.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DdosProtectionTrafficStatsRuleGet

Get traffic stats for a rule



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
    eventID := "eventId_example" // string | Unique ID of the event.
    ruleID := "ruleId_example" // string | Unique ID of the rule.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.DdosProtectionAPI.DdosProtectionTrafficStatsRuleGet(ctx, eventID, ruleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DdosProtectionAPI.DdosProtectionTrafficStatsRuleGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosProtectionTrafficStatsRuleGet`: DdosProtectionTrafficStats
    fmt.Fprintf(os.Stdout, "Response from `DdosProtectionAPI.DdosProtectionTrafficStatsRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventID** | **string** | Unique ID of the event. | 
**ruleID** | **string** | Unique ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosProtectionTrafficStatsRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosProtectionTrafficStats**](DdosProtectionTrafficStats.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
