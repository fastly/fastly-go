# PublishAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**Publish**](PublishAPI.md#Publish) | **POST** `/service/{service_id}/publish/` | Send messages to Fanout subscribers



## Publish

Send messages to Fanout subscribers



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
    serviceID := "serviceId_example" // string | Alphanumeric string identifying the service.
    publishRequest := *openapiclient.NewPublishRequest([]openapiclient.PublishItem{*openapiclient.NewPublishItem("Channel_example", *openapiclient.NewPublishItemFormats())}) // PublishRequest |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PublishAPI.Publish(ctx, serviceID).PublishRequest(publishRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishAPI.Publish`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Publish`: string
    fmt.Fprintf(os.Stdout, "Response from `PublishAPI.Publish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishRequest** | [**PublishRequest**](PublishRequest.md) |  | 

### Return type

**string**

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
