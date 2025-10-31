# PoolAPI

> [!NOTE]
> All URIs are relative to `https://api.fastly.com`

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerPool**](PoolAPI.md#CreateServerPool) | **POST** `/service/{service_id}/version/{version_id}/pool` | Create a server pool
[**DeleteServerPool**](PoolAPI.md#DeleteServerPool) | **DELETE** `/service/{service_id}/version/{version_id}/pool/{pool_name}` | Delete a server pool
[**GetServerPool**](PoolAPI.md#GetServerPool) | **GET** `/service/{service_id}/version/{version_id}/pool/{pool_name}` | Get a server pool
[**ListServerPools**](PoolAPI.md#ListServerPools) | **GET** `/service/{service_id}/version/{version_id}/pool` | List server pools
[**UpdateServerPool**](PoolAPI.md#UpdateServerPool) | **PUT** `/service/{service_id}/version/{version_id}/pool/{pool_name}` | Update a server pool



## CreateServerPool

Create a server pool



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsCertHostname := "tlsCertHostname_example" // string | The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). (optional) (default to "null")
    useTls := int32(56) // int32 | Whether to use TLS. (optional) (default to 0)
    createdAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    deletedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    updatedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    serviceId2 := "serviceId_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    name := "name_example" // string | Name for the Pool. (optional)
    shield := "shield_example" // string | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. (optional) (default to "null")
    requestCondition := "requestCondition_example" // string | Condition which, if met, will select this configuration during a request. Optional. (optional)
    tlsCiphers := "tlsCiphers_example" // string | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. (optional)
    tlsSniHostname := "tlsSniHostname_example" // string | SNI hostname. Optional. (optional)
    minTlsVersion := int32(56) // int32 | Minimum allowed TLS version on connections to this server. Optional. (optional)
    maxTlsVersion := int32(56) // int32 | Maximum allowed TLS version on connections to this server. Optional. (optional)
    healthcheck := "healthcheck_example" // string | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    type_ := "type__example" // string | What type of load balance group to use. (optional)
    overrideHost := "overrideHost_example" // string | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. (optional) (default to "null")
    betweenBytesTimeout := int32(56) // int32 | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`. (optional) (default to 10000)
    connectTimeout := int32(56) // int32 | How long to wait for a timeout in milliseconds. Optional. (optional)
    firstByteTimeout := int32(56) // int32 | How long to wait for the first byte in milliseconds. Optional. (optional)
    maxConnDefault := int32(56) // int32 | Maximum number of connections. Optional. (optional) (default to 200)
    quorum := int32(56) // int32 | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. (optional) (default to 75)
    tlsCheckCert := int32(56) // int32 | Be strict on checking TLS certs. Optional. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PoolAPI.CreateServerPool(ctx, serviceId, versionId).TlsCaCert(tlsCaCert).TlsClientCert(tlsClientCert).TlsClientKey(tlsClientKey).TlsCertHostname(tlsCertHostname).UseTls(useTls).CreatedAt(createdAt).DeletedAt(deletedAt).UpdatedAt(updatedAt).ServiceId2(serviceId2).Version(version).Name(name).Shield(shield).RequestCondition(requestCondition).TlsCiphers(tlsCiphers).TlsSniHostname(tlsSniHostname).MinTlsVersion(minTlsVersion).MaxTlsVersion(maxTlsVersion).Healthcheck(healthcheck).Comment(comment).Type_(type_).OverrideHost(overrideHost).BetweenBytesTimeout(betweenBytesTimeout).ConnectTimeout(connectTimeout).FirstByteTimeout(firstByteTimeout).MaxConnDefault(maxConnDefault).Quorum(quorum).TlsCheckCert(tlsCheckCert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.CreateServerPool`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServerPool`: PoolResponsePost
    fmt.Fprintf(os.Stdout, "Response from `PoolAPI.CreateServerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsCertHostname** | **string** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [default to &quot;null&quot;] **useTls** | **int32** | Whether to use TLS. | [default to 0] **createdAt** | **time.Time** | Date and time in ISO 8601 format. |  **deletedAt** | **time.Time** | Date and time in ISO 8601 format. |  **updatedAt** | **time.Time** | Date and time in ISO 8601 format. |  **serviceId2** | **string** |  |  **version** | **string** |  |  **name** | **string** | Name for the Pool. |  **shield** | **string** | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [default to &quot;null&quot;] **requestCondition** | **string** | Condition which, if met, will select this configuration during a request. Optional. |  **tlsCiphers** | **string** | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. |  **tlsSniHostname** | **string** | SNI hostname. Optional. |  **minTlsVersion** | **int32** | Minimum allowed TLS version on connections to this server. Optional. |  **maxTlsVersion** | **int32** | Maximum allowed TLS version on connections to this server. Optional. |  **healthcheck** | **string** | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. |  **comment** | **string** | A freeform descriptive note. |  **type_** | **string** | What type of load balance group to use. |  **overrideHost** | **string** | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [default to &quot;null&quot;] **betweenBytesTimeout** | **int32** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`. | [default to 10000] **connectTimeout** | **int32** | How long to wait for a timeout in milliseconds. Optional. |  **firstByteTimeout** | **int32** | How long to wait for the first byte in milliseconds. Optional. |  **maxConnDefault** | **int32** | Maximum number of connections. Optional. | [default to 200] **quorum** | **int32** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [default to 75] **tlsCheckCert** | **int32** | Be strict on checking TLS certs. Optional. | 

### Return type

[**PoolResponsePost**](PoolResponsePost.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteServerPool

Delete a server pool



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    poolName := "poolName_example" // string | Name for the Pool.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PoolAPI.DeleteServerPool(ctx, serviceId, versionId, poolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.DeleteServerPool`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServerPool`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `PoolAPI.DeleteServerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**poolName** | **string** | Name for the Pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerPoolRequest struct via the builder pattern


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


## GetServerPool

Get a server pool



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    poolName := "poolName_example" // string | Name for the Pool.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PoolAPI.GetServerPool(ctx, serviceId, versionId, poolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.GetServerPool`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerPool`: PoolResponse
    fmt.Fprintf(os.Stdout, "Response from `PoolAPI.GetServerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**poolName** | **string** | Name for the Pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResponse**](PoolResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListServerPools

List server pools



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PoolAPI.ListServerPools(ctx, serviceId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ListServerPools`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServerPools`: []PoolResponse
    fmt.Fprintf(os.Stdout, "Response from `PoolAPI.ListServerPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PoolResponse**](PoolResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateServerPool

Update a server pool



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
    serviceId := "serviceId_example" // string | Alphanumeric string identifying the service.
    versionId := int32(56) // int32 | Integer identifying a service version.
    poolName := "poolName_example" // string | Name for the Pool.
    tlsCaCert := "tlsCaCert_example" // string | A secure certificate to authenticate a server with. Must be in PEM format. (optional) (default to "null")
    tlsClientCert := "tlsClientCert_example" // string | The client certificate used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsClientKey := "tlsClientKey_example" // string | The client private key used to make authenticated requests. Must be in PEM format. (optional) (default to "null")
    tlsCertHostname := "tlsCertHostname_example" // string | The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). (optional) (default to "null")
    useTls := int32(56) // int32 | Whether to use TLS. (optional) (default to 0)
    createdAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    deletedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    updatedAt := time.Now() // time.Time | Date and time in ISO 8601 format. (optional)
    serviceId2 := "serviceId_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    name := "name_example" // string | Name for the Pool. (optional)
    shield := "shield_example" // string | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. (optional) (default to "null")
    requestCondition := "requestCondition_example" // string | Condition which, if met, will select this configuration during a request. Optional. (optional)
    tlsCiphers := "tlsCiphers_example" // string | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. (optional)
    tlsSniHostname := "tlsSniHostname_example" // string | SNI hostname. Optional. (optional)
    minTlsVersion := int32(56) // int32 | Minimum allowed TLS version on connections to this server. Optional. (optional)
    maxTlsVersion := int32(56) // int32 | Maximum allowed TLS version on connections to this server. Optional. (optional)
    healthcheck := "healthcheck_example" // string | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    type_ := "type__example" // string | What type of load balance group to use. (optional)
    overrideHost := "overrideHost_example" // string | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. (optional) (default to "null")
    betweenBytesTimeout := int32(56) // int32 | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`. (optional) (default to 10000)
    connectTimeout := int32(56) // int32 | How long to wait for a timeout in milliseconds. Optional. (optional)
    firstByteTimeout := int32(56) // int32 | How long to wait for the first byte in milliseconds. Optional. (optional)
    maxConnDefault := int32(56) // int32 | Maximum number of connections. Optional. (optional) (default to 200)
    quorum := int32(56) // int32 | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. (optional) (default to 75)
    tlsCheckCert := int32(56) // int32 | Be strict on checking TLS certs. Optional. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.PoolAPI.UpdateServerPool(ctx, serviceId, versionId, poolName).TlsCaCert(tlsCaCert).TlsClientCert(tlsClientCert).TlsClientKey(tlsClientKey).TlsCertHostname(tlsCertHostname).UseTls(useTls).CreatedAt(createdAt).DeletedAt(deletedAt).UpdatedAt(updatedAt).ServiceId2(serviceId2).Version(version).Name(name).Shield(shield).RequestCondition(requestCondition).TlsCiphers(tlsCiphers).TlsSniHostname(tlsSniHostname).MinTlsVersion(minTlsVersion).MaxTlsVersion(maxTlsVersion).Healthcheck(healthcheck).Comment(comment).Type_(type_).OverrideHost(overrideHost).BetweenBytesTimeout(betweenBytesTimeout).ConnectTimeout(connectTimeout).FirstByteTimeout(firstByteTimeout).MaxConnDefault(maxConnDefault).Quorum(quorum).TlsCheckCert(tlsCheckCert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.UpdateServerPool`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerPool`: PoolResponse
    fmt.Fprintf(os.Stdout, "Response from `PoolAPI.UpdateServerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Alphanumeric string identifying the service. | 
**versionId** | **int32** | Integer identifying a service version. | 
**poolName** | **string** | Name for the Pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlsCaCert** | **string** | A secure certificate to authenticate a server with. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientCert** | **string** | The client certificate used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsClientKey** | **string** | The client private key used to make authenticated requests. Must be in PEM format. | [default to &quot;null&quot;] **tlsCertHostname** | **string** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [default to &quot;null&quot;] **useTls** | **int32** | Whether to use TLS. | [default to 0] **createdAt** | **time.Time** | Date and time in ISO 8601 format. |  **deletedAt** | **time.Time** | Date and time in ISO 8601 format. |  **updatedAt** | **time.Time** | Date and time in ISO 8601 format. |  **serviceId2** | **string** |  |  **version** | **string** |  |  **name** | **string** | Name for the Pool. |  **shield** | **string** | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [default to &quot;null&quot;] **requestCondition** | **string** | Condition which, if met, will select this configuration during a request. Optional. |  **tlsCiphers** | **string** | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. |  **tlsSniHostname** | **string** | SNI hostname. Optional. |  **minTlsVersion** | **int32** | Minimum allowed TLS version on connections to this server. Optional. |  **maxTlsVersion** | **int32** | Maximum allowed TLS version on connections to this server. Optional. |  **healthcheck** | **string** | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. |  **comment** | **string** | A freeform descriptive note. |  **type_** | **string** | What type of load balance group to use. |  **overrideHost** | **string** | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [default to &quot;null&quot;] **betweenBytesTimeout** | **int32** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`. | [default to 10000] **connectTimeout** | **int32** | How long to wait for a timeout in milliseconds. Optional. |  **firstByteTimeout** | **int32** | How long to wait for the first byte in milliseconds. Optional. |  **maxConnDefault** | **int32** | Maximum number of connections. Optional. | [default to 200] **quorum** | **int32** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [default to 75] **tlsCheckCert** | **int32** | Be strict on checking TLS certs. Optional. | 

### Return type

[**PoolResponse**](PoolResponse.md)

### Authorization

[API Token](https://www.fastly.com/documentation/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

