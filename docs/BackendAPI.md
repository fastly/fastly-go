# BackendAPI

All URIs are relative to *https://api.fastly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackend**](BackendAPI.md#CreateBackend) | **POST** `/service/{service_id}/version/{version_id}/backend` | Create a backend
[**DeleteBackend**](BackendAPI.md#DeleteBackend) | **DELETE** `/service/{service_id}/version/{version_id}/backend/{backend_name}` | Delete a backend
[**GetBackend**](BackendAPI.md#GetBackend) | **GET** `/service/{service_id}/version/{version_id}/backend/{backend_name}` | Describe a backend
[**ListBackends**](BackendAPI.md#ListBackends) | **GET** `/service/{service_id}/version/{version_id}/backend` | List backends
[**UpdateBackend**](BackendAPI.md#UpdateBackend) | **PUT** `/service/{service_id}/version/{version_id}/backend/{backend_name}` | Update a backend



## CreateBackend

Create a backend



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    address := "address_example" // string | A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend. (optional)
    autoLoadbalance := true // bool | Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don't have a `request_condition` will be selected based on their `weight`. (optional)
    betweenBytesTimeout := int32(56) // int32 | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`. (optional)
    clientCert := "clientCert_example" // string | Unused. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    connectTimeout := int32(56) // int32 | Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`. (optional)
    firstByteTimeout := int32(56) // int32 | Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`. (optional)
    healthcheck := "healthcheck_example" // string | The name of the healthcheck to use with this backend. (optional)
    hostname := "hostname_example" // string | The hostname of the backend. May be used as an alternative to `address` to set the backend location. (optional)
    ipv4 := "ipv4_example" // string | IPv4 address of the backend. May be used as an alternative to `address` to set the backend location. (optional)
    ipv6 := "ipv6_example" // string | IPv6 address of the backend. May be used as an alternative to `address` to set the backend location. (optional)
    keepaliveTime := int32(56) // int32 | How long in seconds to keep a persistent connection to the backend between requests. (optional)
    maxConn := int32(56) // int32 | Maximum number of concurrent connections this backend will accept. (optional)
    maxTLSVersion := "maxTLSVersion_example" // string | Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. (optional)
    minTLSVersion := "minTLSVersion_example" // string | Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. (optional)
    name := "name_example" // string | The name of the backend. (optional)
    overrideHost := "overrideHost_example" // string | If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL. (optional)
    port := int32(56) // int32 | Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request. (optional)
    requestCondition := "requestCondition_example" // string | Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests. (optional)
    shareKey := "shareKey_example" // string | Value that when shared across backends will enable those backends to share the same health check. (optional)
    shield := "shield_example" // string | Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding). (optional)
    sslCaCert := "sslCaCert_example" // string | CA certificate attached to origin. (optional)
    sslCertHostname := "sslCertHostname_example" // string | Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all. (optional)
    sslCheckCert := true // bool | Be strict on checking SSL certs. (optional) (default to true)
    sslCiphers := "sslCiphers_example" // string | List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. (optional)
    sslClientCert := "sslClientCert_example" // string | Client certificate attached to origin. (optional)
    sslClientKey := "sslClientKey_example" // string | Client key attached to origin. (optional)
    sslHostname := "sslHostname_example" // string | Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation. (optional)
    sslSniHostname := "sslSniHostname_example" // string | Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all. (optional)
    useSsl := true // bool | Whether or not to require TLS for connections to this backend. (optional)
    weight := int32(56) // int32 | Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BackendAPI.CreateBackend(ctx, serviceID, versionID).Address(address).AutoLoadbalance(autoLoadbalance).BetweenBytesTimeout(betweenBytesTimeout).ClientCert(clientCert).Comment(comment).ConnectTimeout(connectTimeout).FirstByteTimeout(firstByteTimeout).Healthcheck(healthcheck).Hostname(hostname).Ipv4(ipv4).Ipv6(ipv6).KeepaliveTime(keepaliveTime).MaxConn(maxConn).MaxTLSVersion(maxTLSVersion).MinTLSVersion(minTLSVersion).Name(name).OverrideHost(overrideHost).Port(port).RequestCondition(requestCondition).ShareKey(shareKey).Shield(shield).SslCaCert(sslCaCert).SslCertHostname(sslCertHostname).SslCheckCert(sslCheckCert).SslCiphers(sslCiphers).SslClientCert(sslClientCert).SslClientKey(sslClientKey).SslHostname(sslHostname).SslSniHostname(sslSniHostname).UseSsl(useSsl).Weight(weight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.CreateBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBackend`: BackendResponse
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.CreateBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend. |  **autoLoadbalance** | **bool** | Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don&#39;t have a `request_condition` will be selected based on their `weight`. |  **betweenBytesTimeout** | **int32** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`. |  **clientCert** | **string** | Unused. |  **comment** | **string** | A freeform descriptive note. |  **connectTimeout** | **int32** | Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`. |  **firstByteTimeout** | **int32** | Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`. |  **healthcheck** | **string** | The name of the healthcheck to use with this backend. |  **hostname** | **string** | The hostname of the backend. May be used as an alternative to `address` to set the backend location. |  **ipv4** | **string** | IPv4 address of the backend. May be used as an alternative to `address` to set the backend location. |  **ipv6** | **string** | IPv6 address of the backend. May be used as an alternative to `address` to set the backend location. |  **keepaliveTime** | **int32** | How long in seconds to keep a persistent connection to the backend between requests. |  **maxConn** | **int32** | Maximum number of concurrent connections this backend will accept. |  **maxTLSVersion** | **string** | Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. |  **minTLSVersion** | **string** | Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. |  **name** | **string** | The name of the backend. |  **overrideHost** | **string** | If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL. |  **port** | **int32** | Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request. |  **requestCondition** | **string** | Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests. |  **shareKey** | **string** | Value that when shared across backends will enable those backends to share the same health check. |  **shield** | **string** | Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding). |  **sslCaCert** | **string** | CA certificate attached to origin. |  **sslCertHostname** | **string** | Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all. |  **sslCheckCert** | **bool** | Be strict on checking SSL certs. | [default to true] **sslCiphers** | **string** | List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. |  **sslClientCert** | **string** | Client certificate attached to origin. |  **sslClientKey** | **string** | Client key attached to origin. |  **sslHostname** | **string** | Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation. |  **sslSniHostname** | **string** | Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all. |  **useSsl** | **bool** | Whether or not to require TLS for connections to this backend. |  **weight** | **int32** | Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true. | 

### Return type

[**BackendResponse**](BackendResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## DeleteBackend

Delete a backend



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    backendName := "backendName_example" // string | The name of the backend.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BackendAPI.DeleteBackend(ctx, serviceID, versionID, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.DeleteBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBackend`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.DeleteBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**backendName** | **string** | The name of the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## GetBackend

Describe a backend



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    backendName := "backendName_example" // string | The name of the backend.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BackendAPI.GetBackend(ctx, serviceID, versionID, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.GetBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackend`: BackendResponse
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.GetBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**backendName** | **string** | The name of the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackendResponse**](BackendResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## ListBackends

List backends



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
    versionID := int32(56) // int32 | Integer identifying a service version.

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BackendAPI.ListBackends(ctx, serviceID, versionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.ListBackends`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackends`: []BackendResponse
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.ListBackends`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BackendResponse**](BackendResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


## UpdateBackend

Update a backend



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
    versionID := int32(56) // int32 | Integer identifying a service version.
    backendName := "backendName_example" // string | The name of the backend.
    address := "address_example" // string | A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend. (optional)
    autoLoadbalance := true // bool | Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don't have a `request_condition` will be selected based on their `weight`. (optional)
    betweenBytesTimeout := int32(56) // int32 | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`. (optional)
    clientCert := "clientCert_example" // string | Unused. (optional)
    comment := "comment_example" // string | A freeform descriptive note. (optional)
    connectTimeout := int32(56) // int32 | Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`. (optional)
    firstByteTimeout := int32(56) // int32 | Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`. (optional)
    healthcheck := "healthcheck_example" // string | The name of the healthcheck to use with this backend. (optional)
    hostname := "hostname_example" // string | The hostname of the backend. May be used as an alternative to `address` to set the backend location. (optional)
    ipv4 := "ipv4_example" // string | IPv4 address of the backend. May be used as an alternative to `address` to set the backend location. (optional)
    ipv6 := "ipv6_example" // string | IPv6 address of the backend. May be used as an alternative to `address` to set the backend location. (optional)
    keepaliveTime := int32(56) // int32 | How long in seconds to keep a persistent connection to the backend between requests. (optional)
    maxConn := int32(56) // int32 | Maximum number of concurrent connections this backend will accept. (optional)
    maxTLSVersion := "maxTLSVersion_example" // string | Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. (optional)
    minTLSVersion := "minTLSVersion_example" // string | Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. (optional)
    name := "name_example" // string | The name of the backend. (optional)
    overrideHost := "overrideHost_example" // string | If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL. (optional)
    port := int32(56) // int32 | Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request. (optional)
    requestCondition := "requestCondition_example" // string | Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests. (optional)
    shareKey := "shareKey_example" // string | Value that when shared across backends will enable those backends to share the same health check. (optional)
    shield := "shield_example" // string | Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding). (optional)
    sslCaCert := "sslCaCert_example" // string | CA certificate attached to origin. (optional)
    sslCertHostname := "sslCertHostname_example" // string | Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all. (optional)
    sslCheckCert := true // bool | Be strict on checking SSL certs. (optional) (default to true)
    sslCiphers := "sslCiphers_example" // string | List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. (optional)
    sslClientCert := "sslClientCert_example" // string | Client certificate attached to origin. (optional)
    sslClientKey := "sslClientKey_example" // string | Client key attached to origin. (optional)
    sslHostname := "sslHostname_example" // string | Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation. (optional)
    sslSniHostname := "sslSniHostname_example" // string | Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all. (optional)
    useSsl := true // bool | Whether or not to require TLS for connections to this backend. (optional)
    weight := int32(56) // int32 | Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true. (optional)

    cfg := fastly.NewConfiguration()
    apiClient := fastly.NewAPIClient(cfg)
    ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")
    resp, r, err := apiClient.BackendAPI.UpdateBackend(ctx, serviceID, versionID, backendName).Address(address).AutoLoadbalance(autoLoadbalance).BetweenBytesTimeout(betweenBytesTimeout).ClientCert(clientCert).Comment(comment).ConnectTimeout(connectTimeout).FirstByteTimeout(firstByteTimeout).Healthcheck(healthcheck).Hostname(hostname).Ipv4(ipv4).Ipv6(ipv6).KeepaliveTime(keepaliveTime).MaxConn(maxConn).MaxTLSVersion(maxTLSVersion).MinTLSVersion(minTLSVersion).Name(name).OverrideHost(overrideHost).Port(port).RequestCondition(requestCondition).ShareKey(shareKey).Shield(shield).SslCaCert(sslCaCert).SslCertHostname(sslCertHostname).SslCheckCert(sslCheckCert).SslCiphers(sslCiphers).SslClientCert(sslClientCert).SslClientKey(sslClientKey).SslHostname(sslHostname).SslSniHostname(sslSniHostname).UseSsl(useSsl).Weight(weight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.UpdateBackend`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBackend`: BackendResponse
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.UpdateBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceID** | **string** | Alphanumeric string identifying the service. | 
**versionID** | **int32** | Integer identifying a service version. | 
**backendName** | **string** | The name of the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend. |  **autoLoadbalance** | **bool** | Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don&#39;t have a `request_condition` will be selected based on their `weight`. |  **betweenBytesTimeout** | **int32** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`. |  **clientCert** | **string** | Unused. |  **comment** | **string** | A freeform descriptive note. |  **connectTimeout** | **int32** | Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`. |  **firstByteTimeout** | **int32** | Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`. |  **healthcheck** | **string** | The name of the healthcheck to use with this backend. |  **hostname** | **string** | The hostname of the backend. May be used as an alternative to `address` to set the backend location. |  **ipv4** | **string** | IPv4 address of the backend. May be used as an alternative to `address` to set the backend location. |  **ipv6** | **string** | IPv6 address of the backend. May be used as an alternative to `address` to set the backend location. |  **keepaliveTime** | **int32** | How long in seconds to keep a persistent connection to the backend between requests. |  **maxConn** | **int32** | Maximum number of concurrent connections this backend will accept. |  **maxTLSVersion** | **string** | Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. |  **minTLSVersion** | **string** | Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. |  **name** | **string** | The name of the backend. |  **overrideHost** | **string** | If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL. |  **port** | **int32** | Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request. |  **requestCondition** | **string** | Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests. |  **shareKey** | **string** | Value that when shared across backends will enable those backends to share the same health check. |  **shield** | **string** | Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding). |  **sslCaCert** | **string** | CA certificate attached to origin. |  **sslCertHostname** | **string** | Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all. |  **sslCheckCert** | **bool** | Be strict on checking SSL certs. | [default to true] **sslCiphers** | **string** | List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. |  **sslClientCert** | **string** | Client certificate attached to origin. |  **sslClientKey** | **string** | Client key attached to origin. |  **sslHostname** | **string** | Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation. |  **sslSniHostname** | **string** | Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all. |  **useSsl** | **bool** | Whether or not to require TLS for connections to this backend. |  **weight** | **int32** | Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true. | 

### Return type

[**BackendResponse**](BackendResponse.md)

### Authorization

[API Token](https://developer.fastly.com/reference/api/#authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[Back to top](#) | [Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
