// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/)

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	gourl "net/url"
	"strconv"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// BackendAPI defines an interface for interacting with the resource.
type BackendAPI interface {

	/*
		CreateBackend Create a backend

		Create a backend for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateBackendRequest
	*/
	CreateBackend(ctx context.Context, serviceID string, versionID int32) APICreateBackendRequest

	// CreateBackendExecute executes the request
	//  @return BackendResponse
	CreateBackendExecute(r APICreateBackendRequest) (*BackendResponse, *http.Response, error)

	/*
		DeleteBackend Delete a backend

		Delete the backend for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param backendName The name of the backend.
		 @return APIDeleteBackendRequest
	*/
	DeleteBackend(ctx context.Context, serviceID string, versionID int32, backendName string) APIDeleteBackendRequest

	// DeleteBackendExecute executes the request
	//  @return InlineResponse200
	DeleteBackendExecute(r APIDeleteBackendRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetBackend Describe a backend

		Get the backend for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param backendName The name of the backend.
		 @return APIGetBackendRequest
	*/
	GetBackend(ctx context.Context, serviceID string, versionID int32, backendName string) APIGetBackendRequest

	// GetBackendExecute executes the request
	//  @return BackendResponse
	GetBackendExecute(r APIGetBackendRequest) (*BackendResponse, *http.Response, error)

	/*
		ListBackends List backends

		List all backends for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListBackendsRequest
	*/
	ListBackends(ctx context.Context, serviceID string, versionID int32) APIListBackendsRequest

	// ListBackendsExecute executes the request
	//  @return []BackendResponse
	ListBackendsExecute(r APIListBackendsRequest) ([]BackendResponse, *http.Response, error)

	/*
		UpdateBackend Update a backend

		Update the backend for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param backendName The name of the backend.
		 @return APIUpdateBackendRequest
	*/
	UpdateBackend(ctx context.Context, serviceID string, versionID int32, backendName string) APIUpdateBackendRequest

	// UpdateBackendExecute executes the request
	//  @return BackendResponse
	UpdateBackendExecute(r APIUpdateBackendRequest) (*BackendResponse, *http.Response, error)
}

// BackendAPIService BackendAPI service
type BackendAPIService service

// APICreateBackendRequest represents a request for the resource.
type APICreateBackendRequest struct {
	ctx                  context.Context
	APIService           BackendAPI
	serviceID            string
	versionID            int32
	address              *string
	autoLoadbalance      *bool
	betweenBytesTimeout  *int32
	clientCert           *string
	comment              *string
	connectTimeout       *int32
	firstByteTimeout     *int32
	healthcheck          *string
	hostname             *string
	ipv4                 *string
	ipv6                 *string
	keepaliveTime        *int32
	maxConn              *int32
	maxTLSVersion        *string
	minTLSVersion        *string
	name                 *string
	overrideHost         *string
	port                 *int32
	preferIpv6           *bool
	requestCondition     *string
	shareKey             *string
	shield               *string
	sslCaCert            *string
	sslCertHostname      *string
	sslCheckCert         *bool
	sslCiphers           *string
	sslClientCert        *string
	sslClientKey         *string
	sslHostname          *string
	sslSniHostname       *string
	tcpKeepaliveEnable   *bool
	tcpKeepaliveInterval *int32
	tcpKeepaliveProbes   *int32
	tcpKeepaliveTime     *int32
	useSsl               *bool
	weight               *int32
}

// Address A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend.
func (r *APICreateBackendRequest) Address(address string) *APICreateBackendRequest {
	r.address = &address
	return r
}

// AutoLoadbalance Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don&#39;t have a &#x60;request_condition&#x60; will be selected based on their &#x60;weight&#x60;.
func (r *APICreateBackendRequest) AutoLoadbalance(autoLoadbalance bool) *APICreateBackendRequest {
	r.autoLoadbalance = &autoLoadbalance
	return r
}

// BetweenBytesTimeout Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using &#x60;bereq.between_bytes_timeout&#x60;.
func (r *APICreateBackendRequest) BetweenBytesTimeout(betweenBytesTimeout int32) *APICreateBackendRequest {
	r.betweenBytesTimeout = &betweenBytesTimeout
	return r
}

// ClientCert Unused.
func (r *APICreateBackendRequest) ClientCert(clientCert string) *APICreateBackendRequest {
	r.clientCert = &clientCert
	return r
}

// Comment A freeform descriptive note.
func (r *APICreateBackendRequest) Comment(comment string) *APICreateBackendRequest {
	r.comment = &comment
	return r
}

// ConnectTimeout Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthetic &#x60;503&#x60; response will be presented instead. May be set at runtime using &#x60;bereq.connect_timeout&#x60;.
func (r *APICreateBackendRequest) ConnectTimeout(connectTimeout int32) *APICreateBackendRequest {
	r.connectTimeout = &connectTimeout
	return r
}

// FirstByteTimeout Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthetic &#x60;503&#x60; response will be presented instead. May be set at runtime using &#x60;bereq.first_byte_timeout&#x60;.
func (r *APICreateBackendRequest) FirstByteTimeout(firstByteTimeout int32) *APICreateBackendRequest {
	r.firstByteTimeout = &firstByteTimeout
	return r
}

// Healthcheck The name of the healthcheck to use with this backend.
func (r *APICreateBackendRequest) Healthcheck(healthcheck string) *APICreateBackendRequest {
	r.healthcheck = &healthcheck
	return r
}

// Hostname The hostname of the backend. May be used as an alternative to &#x60;address&#x60; to set the backend location.
func (r *APICreateBackendRequest) Hostname(hostname string) *APICreateBackendRequest {
	r.hostname = &hostname
	return r
}

// Ipv4 IPv4 address of the backend. May be used as an alternative to &#x60;address&#x60; to set the backend location.
func (r *APICreateBackendRequest) Ipv4(ipv4 string) *APICreateBackendRequest {
	r.ipv4 = &ipv4
	return r
}

// Ipv6 IPv6 address of the backend. May be used as an alternative to &#x60;address&#x60; to set the backend location.
func (r *APICreateBackendRequest) Ipv6(ipv6 string) *APICreateBackendRequest {
	r.ipv6 = &ipv6
	return r
}

// KeepaliveTime How long in seconds to keep a persistent connection to the backend between requests. By default, Varnish keeps connections open as long as it can.
func (r *APICreateBackendRequest) KeepaliveTime(keepaliveTime int32) *APICreateBackendRequest {
	r.keepaliveTime = &keepaliveTime
	return r
}

// MaxConn Maximum number of concurrent connections this backend will accept.
func (r *APICreateBackendRequest) MaxConn(maxConn int32) *APICreateBackendRequest {
	r.maxConn = &maxConn
	return r
}

// MaxTLSVersion Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic &#x60;503&#x60; error response will be generated.
func (r *APICreateBackendRequest) MaxTLSVersion(maxTLSVersion string) *APICreateBackendRequest {
	r.maxTLSVersion = &maxTLSVersion
	return r
}

// MinTLSVersion Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic &#x60;503&#x60; error response will be generated.
func (r *APICreateBackendRequest) MinTLSVersion(minTLSVersion string) *APICreateBackendRequest {
	r.minTLSVersion = &minTLSVersion
	return r
}

// Name The name of the backend.
func (r *APICreateBackendRequest) Name(name string) *APICreateBackendRequest {
	r.name = &name
	return r
}

// OverrideHost If set, will replace the client-supplied HTTP &#x60;Host&#x60; header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing &#x60;bereq.http.Host&#x60; in VCL.
func (r *APICreateBackendRequest) OverrideHost(overrideHost string) *APICreateBackendRequest {
	r.overrideHost = &overrideHost
	return r
}

// Port Port on which the backend server is listening for connections from Fastly. Setting &#x60;port&#x60; to 80 or 443 will also set &#x60;use_ssl&#x60; automatically (to false and true respectively), unless explicitly overridden by setting &#x60;use_ssl&#x60; in the same request.
func (r *APICreateBackendRequest) Port(port int32) *APICreateBackendRequest {
	r.port = &port
	return r
}

// PreferIpv6 Prefer IPv6 connections for DNS hostname lookups.
func (r *APICreateBackendRequest) PreferIpv6(preferIpv6 bool) *APICreateBackendRequest {
	r.preferIpv6 = &preferIpv6
	return r
}

// RequestCondition Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any &#x60;auto_loadbalance&#x60; setting. By default, the first backend added to a service is selected for all requests.
func (r *APICreateBackendRequest) RequestCondition(requestCondition string) *APICreateBackendRequest {
	r.requestCondition = &requestCondition
	return r
}

// ShareKey Value that when shared across backends will enable those backends to share the same health check.
func (r *APICreateBackendRequest) ShareKey(shareKey string) *APICreateBackendRequest {
	r.shareKey = &shareKey
	return r
}

// Shield Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding).
func (r *APICreateBackendRequest) Shield(shield string) *APICreateBackendRequest {
	r.shield = &shield
	return r
}

// SslCaCert CA certificate attached to origin.
func (r *APICreateBackendRequest) SslCaCert(sslCaCert string) *APICreateBackendRequest {
	r.sslCaCert = &sslCaCert
	return r
}

// SslCertHostname Overrides &#x60;ssl_hostname&#x60;, but only for cert verification. Does not affect SNI at all.
func (r *APICreateBackendRequest) SslCertHostname(sslCertHostname string) *APICreateBackendRequest {
	r.sslCertHostname = &sslCertHostname
	return r
}

// SslCheckCert Be strict on checking SSL certs.
func (r *APICreateBackendRequest) SslCheckCert(sslCheckCert bool) *APICreateBackendRequest {
	r.sslCheckCert = &sslCheckCert
	return r
}

// SslCiphers List of [OpenSSL ciphers](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic &#x60;503&#x60; error response will be generated.
func (r *APICreateBackendRequest) SslCiphers(sslCiphers string) *APICreateBackendRequest {
	r.sslCiphers = &sslCiphers
	return r
}

// SslClientCert Client certificate attached to origin.
func (r *APICreateBackendRequest) SslClientCert(sslClientCert string) *APICreateBackendRequest {
	r.sslClientCert = &sslClientCert
	return r
}

// SslClientKey Client key attached to origin.
func (r *APICreateBackendRequest) SslClientKey(sslClientKey string) *APICreateBackendRequest {
	r.sslClientKey = &sslClientKey
	return r
}

// SslHostname Use &#x60;ssl_cert_hostname&#x60; and &#x60;ssl_sni_hostname&#x60; to configure certificate validation.
func (r *APICreateBackendRequest) SslHostname(sslHostname string) *APICreateBackendRequest {
	r.sslHostname = &sslHostname
	return r
}

// SslSniHostname Overrides &#x60;ssl_hostname&#x60;, but only for SNI in the handshake. Does not affect cert validation at all.
func (r *APICreateBackendRequest) SslSniHostname(sslSniHostname string) *APICreateBackendRequest {
	r.sslSniHostname = &sslSniHostname
	return r
}

// TcpKeepaliveEnable Whether to enable TCP keepalives for backend connections. Varnish defaults to using keepalives if this is unspecified.
func (r *APICreateBackendRequest) TcpKeepaliveEnable(tcpKeepaliveEnable bool) *APICreateBackendRequest {
	r.tcpKeepaliveEnable = &tcpKeepaliveEnable
	return r
}

// TcpKeepaliveInterval Interval in seconds between subsequent keepalive probes.
func (r *APICreateBackendRequest) TcpKeepaliveInterval(tcpKeepaliveInterval int32) *APICreateBackendRequest {
	r.tcpKeepaliveInterval = &tcpKeepaliveInterval
	return r
}

// TcpKeepaliveProbes Number of unacknowledged probes to send before considering the connection dead.
func (r *APICreateBackendRequest) TcpKeepaliveProbes(tcpKeepaliveProbes int32) *APICreateBackendRequest {
	r.tcpKeepaliveProbes = &tcpKeepaliveProbes
	return r
}

// TcpKeepaliveTime Interval in seconds between the last data packet sent and the first keepalive probe.
func (r *APICreateBackendRequest) TcpKeepaliveTime(tcpKeepaliveTime int32) *APICreateBackendRequest {
	r.tcpKeepaliveTime = &tcpKeepaliveTime
	return r
}

// UseSsl Whether or not to require TLS for connections to this backend.
func (r *APICreateBackendRequest) UseSsl(useSsl bool) *APICreateBackendRequest {
	r.useSsl = &useSsl
	return r
}

// Weight Weight used to load balance this backend against others. May be any positive integer. If &#x60;auto_loadbalance&#x60; is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have &#x60;auto_loadbalance&#x60; set to true.
func (r *APICreateBackendRequest) Weight(weight int32) *APICreateBackendRequest {
	r.weight = &weight
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateBackendRequest) Execute() (*BackendResponse, *http.Response, error) {
	return r.APIService.CreateBackendExecute(r)
}

/*
CreateBackend Create a backend

Create a backend for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateBackendRequest
*/
func (a *BackendAPIService) CreateBackend(ctx context.Context, serviceID string, versionID int32) APICreateBackendRequest {
	return APICreateBackendRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateBackendExecute executes the request
//  @return BackendResponse
func (a *BackendAPIService) CreateBackendExecute(r APICreateBackendRequest) (*BackendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BackendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackendAPIService.CreateBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/backend"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.autoLoadbalance != nil {
		localVarFormParams.Add("auto_loadbalance", parameterToString(*r.autoLoadbalance, ""))
	}
	if r.betweenBytesTimeout != nil {
		localVarFormParams.Add("between_bytes_timeout", parameterToString(*r.betweenBytesTimeout, ""))
	}
	if r.clientCert != nil {
		localVarFormParams.Add("client_cert", parameterToString(*r.clientCert, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.connectTimeout != nil {
		localVarFormParams.Add("connect_timeout", parameterToString(*r.connectTimeout, ""))
	}
	if r.firstByteTimeout != nil {
		localVarFormParams.Add("first_byte_timeout", parameterToString(*r.firstByteTimeout, ""))
	}
	if r.healthcheck != nil {
		localVarFormParams.Add("healthcheck", parameterToString(*r.healthcheck, ""))
	}
	if r.hostname != nil {
		localVarFormParams.Add("hostname", parameterToString(*r.hostname, ""))
	}
	if r.ipv4 != nil {
		localVarFormParams.Add("ipv4", parameterToString(*r.ipv4, ""))
	}
	if r.ipv6 != nil {
		localVarFormParams.Add("ipv6", parameterToString(*r.ipv6, ""))
	}
	if r.keepaliveTime != nil {
		localVarFormParams.Add("keepalive_time", parameterToString(*r.keepaliveTime, ""))
	}
	if r.maxConn != nil {
		localVarFormParams.Add("max_conn", parameterToString(*r.maxConn, ""))
	}
	if r.maxTLSVersion != nil {
		localVarFormParams.Add("max_tls_version", parameterToString(*r.maxTLSVersion, ""))
	}
	if r.minTLSVersion != nil {
		localVarFormParams.Add("min_tls_version", parameterToString(*r.minTLSVersion, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.overrideHost != nil {
		localVarFormParams.Add("override_host", parameterToString(*r.overrideHost, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.preferIpv6 != nil {
		localVarFormParams.Add("prefer_ipv6", parameterToString(*r.preferIpv6, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.shareKey != nil {
		localVarFormParams.Add("share_key", parameterToString(*r.shareKey, ""))
	}
	if r.shield != nil {
		localVarFormParams.Add("shield", parameterToString(*r.shield, ""))
	}
	if r.sslCaCert != nil {
		localVarFormParams.Add("ssl_ca_cert", parameterToString(*r.sslCaCert, ""))
	}
	if r.sslCertHostname != nil {
		localVarFormParams.Add("ssl_cert_hostname", parameterToString(*r.sslCertHostname, ""))
	}
	if r.sslCheckCert != nil {
		localVarFormParams.Add("ssl_check_cert", parameterToString(*r.sslCheckCert, ""))
	}
	if r.sslCiphers != nil {
		localVarFormParams.Add("ssl_ciphers", parameterToString(*r.sslCiphers, ""))
	}
	if r.sslClientCert != nil {
		localVarFormParams.Add("ssl_client_cert", parameterToString(*r.sslClientCert, ""))
	}
	if r.sslClientKey != nil {
		localVarFormParams.Add("ssl_client_key", parameterToString(*r.sslClientKey, ""))
	}
	if r.sslHostname != nil {
		localVarFormParams.Add("ssl_hostname", parameterToString(*r.sslHostname, ""))
	}
	if r.sslSniHostname != nil {
		localVarFormParams.Add("ssl_sni_hostname", parameterToString(*r.sslSniHostname, ""))
	}
	if r.tcpKeepaliveEnable != nil {
		localVarFormParams.Add("tcp_keepalive_enable", parameterToString(*r.tcpKeepaliveEnable, ""))
	}
	if r.tcpKeepaliveInterval != nil {
		localVarFormParams.Add("tcp_keepalive_interval", parameterToString(*r.tcpKeepaliveInterval, ""))
	}
	if r.tcpKeepaliveProbes != nil {
		localVarFormParams.Add("tcp_keepalive_probes", parameterToString(*r.tcpKeepaliveProbes, ""))
	}
	if r.tcpKeepaliveTime != nil {
		localVarFormParams.Add("tcp_keepalive_time", parameterToString(*r.tcpKeepaliveTime, ""))
	}
	if r.useSsl != nil {
		localVarFormParams.Add("use_ssl", parameterToString(*r.useSsl, ""))
	}
	if r.weight != nil {
		localVarFormParams.Add("weight", parameterToString(*r.weight, ""))
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Fastly-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarHTTPResponse.Request.Method != http.MethodGet && localVarHTTPResponse.Request.Method != http.MethodHead {
		if remaining := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				a.client.RateLimitRemaining = i
			}
		}
		if reset := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Reset"); reset != "" {
			if i, err := strconv.Atoi(reset); err == nil {
				a.client.RateLimitReset = i
			}
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// APIDeleteBackendRequest represents a request for the resource.
type APIDeleteBackendRequest struct {
	ctx         context.Context
	APIService  BackendAPI
	serviceID   string
	versionID   int32
	backendName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteBackendRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteBackendExecute(r)
}

/*
DeleteBackend Delete a backend

Delete the backend for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param backendName The name of the backend.
 @return APIDeleteBackendRequest
*/
func (a *BackendAPIService) DeleteBackend(ctx context.Context, serviceID string, versionID int32, backendName string) APIDeleteBackendRequest {
	return APIDeleteBackendRequest{
		APIService:  a,
		ctx:         ctx,
		serviceID:   serviceID,
		versionID:   versionID,
		backendName: backendName,
	}
}

// DeleteBackendExecute executes the request
//  @return InlineResponse200
func (a *BackendAPIService) DeleteBackendExecute(r APIDeleteBackendRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackendAPIService.DeleteBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/backend/{backend_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"backend_name"+"}", gourl.PathEscape(parameterToString(r.backendName, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Fastly-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarHTTPResponse.Request.Method != http.MethodGet && localVarHTTPResponse.Request.Method != http.MethodHead {
		if remaining := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				a.client.RateLimitRemaining = i
			}
		}
		if reset := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Reset"); reset != "" {
			if i, err := strconv.Atoi(reset); err == nil {
				a.client.RateLimitReset = i
			}
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// APIGetBackendRequest represents a request for the resource.
type APIGetBackendRequest struct {
	ctx         context.Context
	APIService  BackendAPI
	serviceID   string
	versionID   int32
	backendName string
}

// Execute calls the API using the request data configured.
func (r APIGetBackendRequest) Execute() (*BackendResponse, *http.Response, error) {
	return r.APIService.GetBackendExecute(r)
}

/*
GetBackend Describe a backend

Get the backend for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param backendName The name of the backend.
 @return APIGetBackendRequest
*/
func (a *BackendAPIService) GetBackend(ctx context.Context, serviceID string, versionID int32, backendName string) APIGetBackendRequest {
	return APIGetBackendRequest{
		APIService:  a,
		ctx:         ctx,
		serviceID:   serviceID,
		versionID:   versionID,
		backendName: backendName,
	}
}

// GetBackendExecute executes the request
//  @return BackendResponse
func (a *BackendAPIService) GetBackendExecute(r APIGetBackendRequest) (*BackendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BackendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackendAPIService.GetBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/backend/{backend_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"backend_name"+"}", gourl.PathEscape(parameterToString(r.backendName, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Fastly-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarHTTPResponse.Request.Method != http.MethodGet && localVarHTTPResponse.Request.Method != http.MethodHead {
		if remaining := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				a.client.RateLimitRemaining = i
			}
		}
		if reset := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Reset"); reset != "" {
			if i, err := strconv.Atoi(reset); err == nil {
				a.client.RateLimitReset = i
			}
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// APIListBackendsRequest represents a request for the resource.
type APIListBackendsRequest struct {
	ctx        context.Context
	APIService BackendAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListBackendsRequest) Execute() ([]BackendResponse, *http.Response, error) {
	return r.APIService.ListBackendsExecute(r)
}

/*
ListBackends List backends

List all backends for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListBackendsRequest
*/
func (a *BackendAPIService) ListBackends(ctx context.Context, serviceID string, versionID int32) APIListBackendsRequest {
	return APIListBackendsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListBackendsExecute executes the request
//  @return []BackendResponse
func (a *BackendAPIService) ListBackendsExecute(r APIListBackendsRequest) ([]BackendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []BackendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackendAPIService.ListBackends")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/backend"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Fastly-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarHTTPResponse.Request.Method != http.MethodGet && localVarHTTPResponse.Request.Method != http.MethodHead {
		if remaining := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				a.client.RateLimitRemaining = i
			}
		}
		if reset := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Reset"); reset != "" {
			if i, err := strconv.Atoi(reset); err == nil {
				a.client.RateLimitReset = i
			}
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// APIUpdateBackendRequest represents a request for the resource.
type APIUpdateBackendRequest struct {
	ctx                  context.Context
	APIService           BackendAPI
	serviceID            string
	versionID            int32
	backendName          string
	address              *string
	autoLoadbalance      *bool
	betweenBytesTimeout  *int32
	clientCert           *string
	comment              *string
	connectTimeout       *int32
	firstByteTimeout     *int32
	healthcheck          *string
	hostname             *string
	ipv4                 *string
	ipv6                 *string
	keepaliveTime        *int32
	maxConn              *int32
	maxTLSVersion        *string
	minTLSVersion        *string
	name                 *string
	overrideHost         *string
	port                 *int32
	preferIpv6           *bool
	requestCondition     *string
	shareKey             *string
	shield               *string
	sslCaCert            *string
	sslCertHostname      *string
	sslCheckCert         *bool
	sslCiphers           *string
	sslClientCert        *string
	sslClientKey         *string
	sslHostname          *string
	sslSniHostname       *string
	tcpKeepaliveEnable   *bool
	tcpKeepaliveInterval *int32
	tcpKeepaliveProbes   *int32
	tcpKeepaliveTime     *int32
	useSsl               *bool
	weight               *int32
}

// Address A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend.
func (r *APIUpdateBackendRequest) Address(address string) *APIUpdateBackendRequest {
	r.address = &address
	return r
}

// AutoLoadbalance Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don&#39;t have a &#x60;request_condition&#x60; will be selected based on their &#x60;weight&#x60;.
func (r *APIUpdateBackendRequest) AutoLoadbalance(autoLoadbalance bool) *APIUpdateBackendRequest {
	r.autoLoadbalance = &autoLoadbalance
	return r
}

// BetweenBytesTimeout Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using &#x60;bereq.between_bytes_timeout&#x60;.
func (r *APIUpdateBackendRequest) BetweenBytesTimeout(betweenBytesTimeout int32) *APIUpdateBackendRequest {
	r.betweenBytesTimeout = &betweenBytesTimeout
	return r
}

// ClientCert Unused.
func (r *APIUpdateBackendRequest) ClientCert(clientCert string) *APIUpdateBackendRequest {
	r.clientCert = &clientCert
	return r
}

// Comment A freeform descriptive note.
func (r *APIUpdateBackendRequest) Comment(comment string) *APIUpdateBackendRequest {
	r.comment = &comment
	return r
}

// ConnectTimeout Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthetic &#x60;503&#x60; response will be presented instead. May be set at runtime using &#x60;bereq.connect_timeout&#x60;.
func (r *APIUpdateBackendRequest) ConnectTimeout(connectTimeout int32) *APIUpdateBackendRequest {
	r.connectTimeout = &connectTimeout
	return r
}

// FirstByteTimeout Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthetic &#x60;503&#x60; response will be presented instead. May be set at runtime using &#x60;bereq.first_byte_timeout&#x60;.
func (r *APIUpdateBackendRequest) FirstByteTimeout(firstByteTimeout int32) *APIUpdateBackendRequest {
	r.firstByteTimeout = &firstByteTimeout
	return r
}

// Healthcheck The name of the healthcheck to use with this backend.
func (r *APIUpdateBackendRequest) Healthcheck(healthcheck string) *APIUpdateBackendRequest {
	r.healthcheck = &healthcheck
	return r
}

// Hostname The hostname of the backend. May be used as an alternative to &#x60;address&#x60; to set the backend location.
func (r *APIUpdateBackendRequest) Hostname(hostname string) *APIUpdateBackendRequest {
	r.hostname = &hostname
	return r
}

// Ipv4 IPv4 address of the backend. May be used as an alternative to &#x60;address&#x60; to set the backend location.
func (r *APIUpdateBackendRequest) Ipv4(ipv4 string) *APIUpdateBackendRequest {
	r.ipv4 = &ipv4
	return r
}

// Ipv6 IPv6 address of the backend. May be used as an alternative to &#x60;address&#x60; to set the backend location.
func (r *APIUpdateBackendRequest) Ipv6(ipv6 string) *APIUpdateBackendRequest {
	r.ipv6 = &ipv6
	return r
}

// KeepaliveTime How long in seconds to keep a persistent connection to the backend between requests. By default, Varnish keeps connections open as long as it can.
func (r *APIUpdateBackendRequest) KeepaliveTime(keepaliveTime int32) *APIUpdateBackendRequest {
	r.keepaliveTime = &keepaliveTime
	return r
}

// MaxConn Maximum number of concurrent connections this backend will accept.
func (r *APIUpdateBackendRequest) MaxConn(maxConn int32) *APIUpdateBackendRequest {
	r.maxConn = &maxConn
	return r
}

// MaxTLSVersion Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic &#x60;503&#x60; error response will be generated.
func (r *APIUpdateBackendRequest) MaxTLSVersion(maxTLSVersion string) *APIUpdateBackendRequest {
	r.maxTLSVersion = &maxTLSVersion
	return r
}

// MinTLSVersion Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic &#x60;503&#x60; error response will be generated.
func (r *APIUpdateBackendRequest) MinTLSVersion(minTLSVersion string) *APIUpdateBackendRequest {
	r.minTLSVersion = &minTLSVersion
	return r
}

// Name The name of the backend.
func (r *APIUpdateBackendRequest) Name(name string) *APIUpdateBackendRequest {
	r.name = &name
	return r
}

// OverrideHost If set, will replace the client-supplied HTTP &#x60;Host&#x60; header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing &#x60;bereq.http.Host&#x60; in VCL.
func (r *APIUpdateBackendRequest) OverrideHost(overrideHost string) *APIUpdateBackendRequest {
	r.overrideHost = &overrideHost
	return r
}

// Port Port on which the backend server is listening for connections from Fastly. Setting &#x60;port&#x60; to 80 or 443 will also set &#x60;use_ssl&#x60; automatically (to false and true respectively), unless explicitly overridden by setting &#x60;use_ssl&#x60; in the same request.
func (r *APIUpdateBackendRequest) Port(port int32) *APIUpdateBackendRequest {
	r.port = &port
	return r
}

// PreferIpv6 Prefer IPv6 connections for DNS hostname lookups.
func (r *APIUpdateBackendRequest) PreferIpv6(preferIpv6 bool) *APIUpdateBackendRequest {
	r.preferIpv6 = &preferIpv6
	return r
}

// RequestCondition Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any &#x60;auto_loadbalance&#x60; setting. By default, the first backend added to a service is selected for all requests.
func (r *APIUpdateBackendRequest) RequestCondition(requestCondition string) *APIUpdateBackendRequest {
	r.requestCondition = &requestCondition
	return r
}

// ShareKey Value that when shared across backends will enable those backends to share the same health check.
func (r *APIUpdateBackendRequest) ShareKey(shareKey string) *APIUpdateBackendRequest {
	r.shareKey = &shareKey
	return r
}

// Shield Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding).
func (r *APIUpdateBackendRequest) Shield(shield string) *APIUpdateBackendRequest {
	r.shield = &shield
	return r
}

// SslCaCert CA certificate attached to origin.
func (r *APIUpdateBackendRequest) SslCaCert(sslCaCert string) *APIUpdateBackendRequest {
	r.sslCaCert = &sslCaCert
	return r
}

// SslCertHostname Overrides &#x60;ssl_hostname&#x60;, but only for cert verification. Does not affect SNI at all.
func (r *APIUpdateBackendRequest) SslCertHostname(sslCertHostname string) *APIUpdateBackendRequest {
	r.sslCertHostname = &sslCertHostname
	return r
}

// SslCheckCert Be strict on checking SSL certs.
func (r *APIUpdateBackendRequest) SslCheckCert(sslCheckCert bool) *APIUpdateBackendRequest {
	r.sslCheckCert = &sslCheckCert
	return r
}

// SslCiphers List of [OpenSSL ciphers](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic &#x60;503&#x60; error response will be generated.
func (r *APIUpdateBackendRequest) SslCiphers(sslCiphers string) *APIUpdateBackendRequest {
	r.sslCiphers = &sslCiphers
	return r
}

// SslClientCert Client certificate attached to origin.
func (r *APIUpdateBackendRequest) SslClientCert(sslClientCert string) *APIUpdateBackendRequest {
	r.sslClientCert = &sslClientCert
	return r
}

// SslClientKey Client key attached to origin.
func (r *APIUpdateBackendRequest) SslClientKey(sslClientKey string) *APIUpdateBackendRequest {
	r.sslClientKey = &sslClientKey
	return r
}

// SslHostname Use &#x60;ssl_cert_hostname&#x60; and &#x60;ssl_sni_hostname&#x60; to configure certificate validation.
func (r *APIUpdateBackendRequest) SslHostname(sslHostname string) *APIUpdateBackendRequest {
	r.sslHostname = &sslHostname
	return r
}

// SslSniHostname Overrides &#x60;ssl_hostname&#x60;, but only for SNI in the handshake. Does not affect cert validation at all.
func (r *APIUpdateBackendRequest) SslSniHostname(sslSniHostname string) *APIUpdateBackendRequest {
	r.sslSniHostname = &sslSniHostname
	return r
}

// TcpKeepaliveEnable Whether to enable TCP keepalives for backend connections. Varnish defaults to using keepalives if this is unspecified.
func (r *APIUpdateBackendRequest) TcpKeepaliveEnable(tcpKeepaliveEnable bool) *APIUpdateBackendRequest {
	r.tcpKeepaliveEnable = &tcpKeepaliveEnable
	return r
}

// TcpKeepaliveInterval Interval in seconds between subsequent keepalive probes.
func (r *APIUpdateBackendRequest) TcpKeepaliveInterval(tcpKeepaliveInterval int32) *APIUpdateBackendRequest {
	r.tcpKeepaliveInterval = &tcpKeepaliveInterval
	return r
}

// TcpKeepaliveProbes Number of unacknowledged probes to send before considering the connection dead.
func (r *APIUpdateBackendRequest) TcpKeepaliveProbes(tcpKeepaliveProbes int32) *APIUpdateBackendRequest {
	r.tcpKeepaliveProbes = &tcpKeepaliveProbes
	return r
}

// TcpKeepaliveTime Interval in seconds between the last data packet sent and the first keepalive probe.
func (r *APIUpdateBackendRequest) TcpKeepaliveTime(tcpKeepaliveTime int32) *APIUpdateBackendRequest {
	r.tcpKeepaliveTime = &tcpKeepaliveTime
	return r
}

// UseSsl Whether or not to require TLS for connections to this backend.
func (r *APIUpdateBackendRequest) UseSsl(useSsl bool) *APIUpdateBackendRequest {
	r.useSsl = &useSsl
	return r
}

// Weight Weight used to load balance this backend against others. May be any positive integer. If &#x60;auto_loadbalance&#x60; is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have &#x60;auto_loadbalance&#x60; set to true.
func (r *APIUpdateBackendRequest) Weight(weight int32) *APIUpdateBackendRequest {
	r.weight = &weight
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateBackendRequest) Execute() (*BackendResponse, *http.Response, error) {
	return r.APIService.UpdateBackendExecute(r)
}

/*
UpdateBackend Update a backend

Update the backend for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param backendName The name of the backend.
 @return APIUpdateBackendRequest
*/
func (a *BackendAPIService) UpdateBackend(ctx context.Context, serviceID string, versionID int32, backendName string) APIUpdateBackendRequest {
	return APIUpdateBackendRequest{
		APIService:  a,
		ctx:         ctx,
		serviceID:   serviceID,
		versionID:   versionID,
		backendName: backendName,
	}
}

// UpdateBackendExecute executes the request
//  @return BackendResponse
func (a *BackendAPIService) UpdateBackendExecute(r APIUpdateBackendRequest) (*BackendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BackendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackendAPIService.UpdateBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/backend/{backend_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"backend_name"+"}", gourl.PathEscape(parameterToString(r.backendName, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.autoLoadbalance != nil {
		localVarFormParams.Add("auto_loadbalance", parameterToString(*r.autoLoadbalance, ""))
	}
	if r.betweenBytesTimeout != nil {
		localVarFormParams.Add("between_bytes_timeout", parameterToString(*r.betweenBytesTimeout, ""))
	}
	if r.clientCert != nil {
		localVarFormParams.Add("client_cert", parameterToString(*r.clientCert, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.connectTimeout != nil {
		localVarFormParams.Add("connect_timeout", parameterToString(*r.connectTimeout, ""))
	}
	if r.firstByteTimeout != nil {
		localVarFormParams.Add("first_byte_timeout", parameterToString(*r.firstByteTimeout, ""))
	}
	if r.healthcheck != nil {
		localVarFormParams.Add("healthcheck", parameterToString(*r.healthcheck, ""))
	}
	if r.hostname != nil {
		localVarFormParams.Add("hostname", parameterToString(*r.hostname, ""))
	}
	if r.ipv4 != nil {
		localVarFormParams.Add("ipv4", parameterToString(*r.ipv4, ""))
	}
	if r.ipv6 != nil {
		localVarFormParams.Add("ipv6", parameterToString(*r.ipv6, ""))
	}
	if r.keepaliveTime != nil {
		localVarFormParams.Add("keepalive_time", parameterToString(*r.keepaliveTime, ""))
	}
	if r.maxConn != nil {
		localVarFormParams.Add("max_conn", parameterToString(*r.maxConn, ""))
	}
	if r.maxTLSVersion != nil {
		localVarFormParams.Add("max_tls_version", parameterToString(*r.maxTLSVersion, ""))
	}
	if r.minTLSVersion != nil {
		localVarFormParams.Add("min_tls_version", parameterToString(*r.minTLSVersion, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.overrideHost != nil {
		localVarFormParams.Add("override_host", parameterToString(*r.overrideHost, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.preferIpv6 != nil {
		localVarFormParams.Add("prefer_ipv6", parameterToString(*r.preferIpv6, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.shareKey != nil {
		localVarFormParams.Add("share_key", parameterToString(*r.shareKey, ""))
	}
	if r.shield != nil {
		localVarFormParams.Add("shield", parameterToString(*r.shield, ""))
	}
	if r.sslCaCert != nil {
		localVarFormParams.Add("ssl_ca_cert", parameterToString(*r.sslCaCert, ""))
	}
	if r.sslCertHostname != nil {
		localVarFormParams.Add("ssl_cert_hostname", parameterToString(*r.sslCertHostname, ""))
	}
	if r.sslCheckCert != nil {
		localVarFormParams.Add("ssl_check_cert", parameterToString(*r.sslCheckCert, ""))
	}
	if r.sslCiphers != nil {
		localVarFormParams.Add("ssl_ciphers", parameterToString(*r.sslCiphers, ""))
	}
	if r.sslClientCert != nil {
		localVarFormParams.Add("ssl_client_cert", parameterToString(*r.sslClientCert, ""))
	}
	if r.sslClientKey != nil {
		localVarFormParams.Add("ssl_client_key", parameterToString(*r.sslClientKey, ""))
	}
	if r.sslHostname != nil {
		localVarFormParams.Add("ssl_hostname", parameterToString(*r.sslHostname, ""))
	}
	if r.sslSniHostname != nil {
		localVarFormParams.Add("ssl_sni_hostname", parameterToString(*r.sslSniHostname, ""))
	}
	if r.tcpKeepaliveEnable != nil {
		localVarFormParams.Add("tcp_keepalive_enable", parameterToString(*r.tcpKeepaliveEnable, ""))
	}
	if r.tcpKeepaliveInterval != nil {
		localVarFormParams.Add("tcp_keepalive_interval", parameterToString(*r.tcpKeepaliveInterval, ""))
	}
	if r.tcpKeepaliveProbes != nil {
		localVarFormParams.Add("tcp_keepalive_probes", parameterToString(*r.tcpKeepaliveProbes, ""))
	}
	if r.tcpKeepaliveTime != nil {
		localVarFormParams.Add("tcp_keepalive_time", parameterToString(*r.tcpKeepaliveTime, ""))
	}
	if r.useSsl != nil {
		localVarFormParams.Add("use_ssl", parameterToString(*r.useSsl, ""))
	}
	if r.weight != nil {
		localVarFormParams.Add("weight", parameterToString(*r.weight, ""))
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Fastly-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarHTTPResponse.Request.Method != http.MethodGet && localVarHTTPResponse.Request.Method != http.MethodHead {
		if remaining := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				a.client.RateLimitRemaining = i
			}
		}
		if reset := localVarHTTPResponse.Header.Get("Fastly-RateLimit-Reset"); reset != "" {
			if i, err := strconv.Atoi(reset); err == nil {
				a.client.RateLimitReset = i
			}
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
