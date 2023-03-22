# EventsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsAPI.md#GetEvent) | **GET** `/events/{event_id}` | Get an event
[**ListEvents**](EventsAPI.md#ListEvents) | **GET** `/events` | List events



## GetEvent

Get an event



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
    eventID := "eventId_example" // string | Alphanumeric string identifying an event.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EventsAPI.GetEvent(ctx, eventID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: EventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventID** | **string** | Alphanumeric string identifying an event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListEvents

List events



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
    filterCustomerID := "x4xCwxxJxGCx123Rx5xTx" // string | Limit the results returned to a specific customer. (optional)
    filterEventType := "filterEventType_example" // string | Limit the returned events to a specific `event_type`. (optional)
    filterServiceID := "filterServiceId_example" // string | Limit the results returned to a specific service. (optional)
    filterUserID := "filterUserId_example" // string | Limit the results returned to a specific user. (optional)
    filterTokenID := "filterTokenId_example" // string | Limit the returned events to a specific token. (optional)
    filterCreatedAt := "filterCreatedAt_example" // string | Limit the returned events to a specific time frame. Accepts sub-parameters: lt, lte, gt, gte (e.g., filter[created_at][gt]=2022-01-12).  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EventsAPI.ListEvents(ctx).FilterCustomerID(filterCustomerID).FilterEventType(filterEventType).FilterServiceID(filterServiceID).FilterUserID(filterUserID).FilterTokenID(filterTokenID).FilterCreatedAt(filterCreatedAt).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEvents`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: EventsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCustomerID** | **string** | Limit the results returned to a specific customer. |  **filterEventType** | **string** | Limit the returned events to a specific `event_type`. |  **filterServiceID** | **string** | Limit the results returned to a specific service. |  **filterUserID** | **string** | Limit the results returned to a specific user. |  **filterTokenID** | **string** | Limit the returned events to a specific token. |  **filterCreatedAt** | **string** | Limit the returned events to a specific time frame. Accepts sub-parameters: lt, lte, gt, gte (e.g., filter[created_at][gt]&#x3D;2022-01-12).  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
