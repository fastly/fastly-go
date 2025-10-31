# EventsAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

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
    eventId := "eventId_example" // string | Alphanumeric string identifying an event.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EventsAPI.GetEvent(ctx, eventId).Execute()
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
**eventId** | **string** | Alphanumeric string identifying an event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

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
    filterCustomerId := "x4xCwxxJxGCx123Rx5xTx" // string | Limit the results returned to a specific customer. (optional)
    filterEventType := "filterEventType_example" // string | Limit the returned events to a specific `event_type`. (optional)
    filterServiceId := "filterServiceId_example" // string | Limit the results returned to a specific service. (optional)
    filterUserId := "filterUserId_example" // string | Limit the results returned to a specific user. (optional)
    filterTokenId := "filterTokenId_example" // string | Limit the returned events to a specific token. (optional)
    filterCreatedAt := "filterCreatedAt_example" // string | Limit the returned events to a specific time frame. Accepts sub-parameters: lt, lte, gt, gte (e.g., filter[created_at][gt]=2022-01-12).  (optional)
    filterCreatedAtLte := "filterCreatedAtLte_example" // string | Return events on and before a date and time in ISO 8601 format.  (optional)
    filterCreatedAtLt := "filterCreatedAtLt_example" // string | Return events before a date and time in ISO 8601 format.  (optional)
    filterCreatedAtGte := "filterCreatedAtGte_example" // string | Return events on and after a date and time in ISO 8601 format.  (optional)
    filterCreatedAtGt := "filterCreatedAtGt_example" // string | Return events after a date and time in ISO 8601 format.  (optional)
    pageNumber := int32(1) // int32 | Current page. (optional)
    pageSize := int32(20) // int32 | Number of records per page. (optional) (default to 20)
    sort := "created_at" // string | The order in which to list the results by creation date. (optional) (default to "created_at")

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.EventsAPI.ListEvents(ctx).FilterCustomerId(filterCustomerId).FilterEventType(filterEventType).FilterServiceId(filterServiceId).FilterUserId(filterUserId).FilterTokenId(filterTokenId).FilterCreatedAt(filterCreatedAt).FilterCreatedAtLte(filterCreatedAtLte).FilterCreatedAtLt(filterCreatedAtLt).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtGt(filterCreatedAtGt).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
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
 **filterCustomerId** | **string** | Limit the results returned to a specific customer. |  **filterEventType** | **string** | Limit the returned events to a specific `event_type`. |  **filterServiceId** | **string** | Limit the results returned to a specific service. |  **filterUserId** | **string** | Limit the results returned to a specific user. |  **filterTokenId** | **string** | Limit the returned events to a specific token. |  **filterCreatedAt** | **string** | Limit the returned events to a specific time frame. Accepts sub-parameters: lt, lte, gt, gte (e.g., filter[created_at][gt]&#x3D;2022-01-12).  |  **filterCreatedAtLte** | **string** | Return events on and before a date and time in ISO 8601 format.  |  **filterCreatedAtLt** | **string** | Return events before a date and time in ISO 8601 format.  |  **filterCreatedAtGte** | **string** | Return events on and after a date and time in ISO 8601 format.  |  **filterCreatedAtGt** | **string** | Return events after a date and time in ISO 8601 format.  |  **pageNumber** | **int32** | Current page. |  **pageSize** | **int32** | Number of records per page. | [default to 20] **sort** | **string** | The order in which to list the results by creation date. | [default to &quot;created_at&quot;]

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

