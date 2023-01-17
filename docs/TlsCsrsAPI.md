# TLSCsrsAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCsr**](TlsCsrsAPI.md#CreateCsr) | **POST** `/tls/certificate_signing_requests` | Create CSR



## CreateCsr

Create CSR



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
    tlsCsr := *openapiclient.NewTLSCsr() // TLSCsr |  (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.TLSCsrsAPI.CreateCsr(ctx).TLSCsr(tlsCsr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TLSCsrsAPI.CreateCsr`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCsr`: TLSCsrResponse
    fmt.Fprintf(os.Stdout, "Response from `TLSCsrsAPI.CreateCsr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCsr** | [**TLSCsr**](TlsCsr.md) |  | 

### Return type

[**TLSCsrResponse**](TlsCsrResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
