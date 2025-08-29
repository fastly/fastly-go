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
	"time"
)

// Linger please
var (
	_ context.Context
)

// PoolAPI defines an interface for interacting with the resource.
type PoolAPI interface {

	/*
		CreateServerPool Create a server pool

		Creates a pool for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateServerPoolRequest
	*/
	CreateServerPool(ctx context.Context, serviceID string, versionID int32) APICreateServerPoolRequest

	// CreateServerPoolExecute executes the request
	//  @return PoolResponsePost
	CreateServerPoolExecute(r APICreateServerPoolRequest) (*PoolResponsePost, *http.Response, error)

	/*
		DeleteServerPool Delete a server pool

		Deletes a specific pool for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param poolName Name for the Pool.
		 @return APIDeleteServerPoolRequest
	*/
	DeleteServerPool(ctx context.Context, serviceID string, versionID int32, poolName string) APIDeleteServerPoolRequest

	// DeleteServerPoolExecute executes the request
	//  @return InlineResponse200
	DeleteServerPoolExecute(r APIDeleteServerPoolRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetServerPool Get a server pool

		Gets a single pool for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param poolName Name for the Pool.
		 @return APIGetServerPoolRequest
	*/
	GetServerPool(ctx context.Context, serviceID string, versionID int32, poolName string) APIGetServerPoolRequest

	// GetServerPoolExecute executes the request
	//  @return PoolResponse
	GetServerPoolExecute(r APIGetServerPoolRequest) (*PoolResponse, *http.Response, error)

	/*
		ListServerPools List server pools

		Lists all pools for a particular service and pool.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListServerPoolsRequest
	*/
	ListServerPools(ctx context.Context, serviceID string, versionID int32) APIListServerPoolsRequest

	// ListServerPoolsExecute executes the request
	//  @return []PoolResponse
	ListServerPoolsExecute(r APIListServerPoolsRequest) ([]PoolResponse, *http.Response, error)

	/*
		UpdateServerPool Update a server pool

		Updates a specific pool for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param poolName Name for the Pool.
		 @return APIUpdateServerPoolRequest
	*/
	UpdateServerPool(ctx context.Context, serviceID string, versionID int32, poolName string) APIUpdateServerPoolRequest

	// UpdateServerPoolExecute executes the request
	//  @return PoolResponse
	UpdateServerPoolExecute(r APIUpdateServerPoolRequest) (*PoolResponse, *http.Response, error)
}

// PoolAPIService PoolAPI service
type PoolAPIService service

// APICreateServerPoolRequest represents a request for the resource.
type APICreateServerPoolRequest struct {
	ctx                 context.Context
	APIService          PoolAPI
	serviceID           string
	versionID           int32
	tlsCaCert           *string
	tlsClientCert       *string
	tlsClientKey        *string
	tlsCertHostname     *string
	useTLS              *int32
	createdAt           *time.Time
	deletedAt           *time.Time
	updatedAt           *time.Time
	serviceID2          *string
	version             *string
	name                *string
	shield              *string
	requestCondition    *string
	tlsCiphers          *string
	tlsSniHostname      *string
	minTLSVersion       *int32
	maxTLSVersion       *int32
	healthcheck         *string
	comment             *string
	resourceType        *string
	overrideHost        *string
	betweenBytesTimeout *int32
	connectTimeout      *int32
	firstByteTimeout    *int32
	maxConnDefault      *int32
	quorum              *int32
	tlsCheckCert        *int32
}

// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateServerPoolRequest) TLSCaCert(tlsCaCert string) *APICreateServerPoolRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}

// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateServerPoolRequest) TLSClientCert(tlsClientCert string) *APICreateServerPoolRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}

// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateServerPoolRequest) TLSClientKey(tlsClientKey string) *APICreateServerPoolRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}

// TLSCertHostname The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
func (r *APICreateServerPoolRequest) TLSCertHostname(tlsCertHostname string) *APICreateServerPoolRequest {
	r.tlsCertHostname = &tlsCertHostname
	return r
}

// UseTLS Whether to use TLS.
func (r *APICreateServerPoolRequest) UseTLS(useTLS int32) *APICreateServerPoolRequest {
	r.useTLS = &useTLS
	return r
}

// CreatedAt Date and time in ISO 8601 format.
func (r *APICreateServerPoolRequest) CreatedAt(createdAt time.Time) *APICreateServerPoolRequest {
	r.createdAt = &createdAt
	return r
}

// DeletedAt Date and time in ISO 8601 format.
func (r *APICreateServerPoolRequest) DeletedAt(deletedAt time.Time) *APICreateServerPoolRequest {
	r.deletedAt = &deletedAt
	return r
}

// UpdatedAt Date and time in ISO 8601 format.
func (r *APICreateServerPoolRequest) UpdatedAt(updatedAt time.Time) *APICreateServerPoolRequest {
	r.updatedAt = &updatedAt
	return r
}

// ServiceID2 returns a pointer to a request.
func (r *APICreateServerPoolRequest) ServiceID2(serviceID2 string) *APICreateServerPoolRequest {
	r.serviceID2 = &serviceID2
	return r
}

// Version returns a pointer to a request.
func (r *APICreateServerPoolRequest) Version(version string) *APICreateServerPoolRequest {
	r.version = &version
	return r
}

// Name Name for the Pool.
func (r *APICreateServerPoolRequest) Name(name string) *APICreateServerPoolRequest {
	r.name = &name
	return r
}

// Shield Selected POP to serve as a shield for the servers. Defaults to &#x60;null&#x60; meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding.
func (r *APICreateServerPoolRequest) Shield(shield string) *APICreateServerPoolRequest {
	r.shield = &shield
	return r
}

// RequestCondition Condition which, if met, will select this configuration during a request. Optional.
func (r *APICreateServerPoolRequest) RequestCondition(requestCondition string) *APICreateServerPoolRequest {
	r.requestCondition = &requestCondition
	return r
}

// TLSCiphers List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional.
func (r *APICreateServerPoolRequest) TLSCiphers(tlsCiphers string) *APICreateServerPoolRequest {
	r.tlsCiphers = &tlsCiphers
	return r
}

// TLSSniHostname SNI hostname. Optional.
func (r *APICreateServerPoolRequest) TLSSniHostname(tlsSniHostname string) *APICreateServerPoolRequest {
	r.tlsSniHostname = &tlsSniHostname
	return r
}

// MinTLSVersion Minimum allowed TLS version on connections to this server. Optional.
func (r *APICreateServerPoolRequest) MinTLSVersion(minTLSVersion int32) *APICreateServerPoolRequest {
	r.minTLSVersion = &minTLSVersion
	return r
}

// MaxTLSVersion Maximum allowed TLS version on connections to this server. Optional.
func (r *APICreateServerPoolRequest) MaxTLSVersion(maxTLSVersion int32) *APICreateServerPoolRequest {
	r.maxTLSVersion = &maxTLSVersion
	return r
}

// Healthcheck Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools.
func (r *APICreateServerPoolRequest) Healthcheck(healthcheck string) *APICreateServerPoolRequest {
	r.healthcheck = &healthcheck
	return r
}

// Comment A freeform descriptive note.
func (r *APICreateServerPoolRequest) Comment(comment string) *APICreateServerPoolRequest {
	r.comment = &comment
	return r
}

// ResourceType What type of load balance group to use.
func (r *APICreateServerPoolRequest) ResourceType(resourceType string) *APICreateServerPoolRequest {
	r.resourceType = &resourceType
	return r
}

// OverrideHost The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to &#x60;null&#x60; meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting.
func (r *APICreateServerPoolRequest) OverrideHost(overrideHost string) *APICreateServerPoolRequest {
	r.overrideHost = &overrideHost
	return r
}

// BetweenBytesTimeout Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using &#x60;bereq.between_bytes_timeout&#x60;.
func (r *APICreateServerPoolRequest) BetweenBytesTimeout(betweenBytesTimeout int32) *APICreateServerPoolRequest {
	r.betweenBytesTimeout = &betweenBytesTimeout
	return r
}

// ConnectTimeout How long to wait for a timeout in milliseconds. Optional.
func (r *APICreateServerPoolRequest) ConnectTimeout(connectTimeout int32) *APICreateServerPoolRequest {
	r.connectTimeout = &connectTimeout
	return r
}

// FirstByteTimeout How long to wait for the first byte in milliseconds. Optional.
func (r *APICreateServerPoolRequest) FirstByteTimeout(firstByteTimeout int32) *APICreateServerPoolRequest {
	r.firstByteTimeout = &firstByteTimeout
	return r
}

// MaxConnDefault Maximum number of connections. Optional.
func (r *APICreateServerPoolRequest) MaxConnDefault(maxConnDefault int32) *APICreateServerPoolRequest {
	r.maxConnDefault = &maxConnDefault
	return r
}

// Quorum Percentage of capacity (&#x60;0-100&#x60;) that needs to be operationally available for a pool to be considered up.
func (r *APICreateServerPoolRequest) Quorum(quorum int32) *APICreateServerPoolRequest {
	r.quorum = &quorum
	return r
}

// TLSCheckCert Be strict on checking TLS certs. Optional.
func (r *APICreateServerPoolRequest) TLSCheckCert(tlsCheckCert int32) *APICreateServerPoolRequest {
	r.tlsCheckCert = &tlsCheckCert
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateServerPoolRequest) Execute() (*PoolResponsePost, *http.Response, error) {
	return r.APIService.CreateServerPoolExecute(r)
}

/*
CreateServerPool Create a server pool

Creates a pool for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateServerPoolRequest
*/
func (a *PoolAPIService) CreateServerPool(ctx context.Context, serviceID string, versionID int32) APICreateServerPoolRequest {
	return APICreateServerPoolRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateServerPoolExecute executes the request
//  @return PoolResponsePost
func (a *PoolAPIService) CreateServerPoolExecute(r APICreateServerPoolRequest) (*PoolResponsePost, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PoolResponsePost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolAPIService.CreateServerPool")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/pool"
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
	if r.tlsCaCert != nil {
		localVarFormParams.Add("tls_ca_cert", parameterToString(*r.tlsCaCert, ""))
	}
	if r.tlsClientCert != nil {
		localVarFormParams.Add("tls_client_cert", parameterToString(*r.tlsClientCert, ""))
	}
	if r.tlsClientKey != nil {
		localVarFormParams.Add("tls_client_key", parameterToString(*r.tlsClientKey, ""))
	}
	if r.tlsCertHostname != nil {
		localVarFormParams.Add("tls_cert_hostname", parameterToString(*r.tlsCertHostname, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
	}
	if r.createdAt != nil {
		localVarFormParams.Add("created_at", parameterToString(*r.createdAt, ""))
	}
	if r.deletedAt != nil {
		localVarFormParams.Add("deleted_at", parameterToString(*r.deletedAt, ""))
	}
	if r.updatedAt != nil {
		localVarFormParams.Add("updated_at", parameterToString(*r.updatedAt, ""))
	}
	if r.serviceID2 != nil {
		paramJSON, err := parameterToJSON(*r.serviceID2)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("service_id", paramJSON)
	}
	if r.version != nil {
		paramJSON, err := parameterToJSON(*r.version)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("version", paramJSON)
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.shield != nil {
		localVarFormParams.Add("shield", parameterToString(*r.shield, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.tlsCiphers != nil {
		localVarFormParams.Add("tls_ciphers", parameterToString(*r.tlsCiphers, ""))
	}
	if r.tlsSniHostname != nil {
		localVarFormParams.Add("tls_sni_hostname", parameterToString(*r.tlsSniHostname, ""))
	}
	if r.minTLSVersion != nil {
		localVarFormParams.Add("min_tls_version", parameterToString(*r.minTLSVersion, ""))
	}
	if r.maxTLSVersion != nil {
		localVarFormParams.Add("max_tls_version", parameterToString(*r.maxTLSVersion, ""))
	}
	if r.healthcheck != nil {
		localVarFormParams.Add("healthcheck", parameterToString(*r.healthcheck, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.resourceType != nil {
		localVarFormParams.Add("type", parameterToString(*r.resourceType, ""))
	}
	if r.overrideHost != nil {
		localVarFormParams.Add("override_host", parameterToString(*r.overrideHost, ""))
	}
	if r.betweenBytesTimeout != nil {
		localVarFormParams.Add("between_bytes_timeout", parameterToString(*r.betweenBytesTimeout, ""))
	}
	if r.connectTimeout != nil {
		localVarFormParams.Add("connect_timeout", parameterToString(*r.connectTimeout, ""))
	}
	if r.firstByteTimeout != nil {
		localVarFormParams.Add("first_byte_timeout", parameterToString(*r.firstByteTimeout, ""))
	}
	if r.maxConnDefault != nil {
		localVarFormParams.Add("max_conn_default", parameterToString(*r.maxConnDefault, ""))
	}
	if r.quorum != nil {
		localVarFormParams.Add("quorum", parameterToString(*r.quorum, ""))
	}
	if r.tlsCheckCert != nil {
		localVarFormParams.Add("tls_check_cert", parameterToString(*r.tlsCheckCert, ""))
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

// APIDeleteServerPoolRequest represents a request for the resource.
type APIDeleteServerPoolRequest struct {
	ctx        context.Context
	APIService PoolAPI
	serviceID  string
	versionID  int32
	poolName   string
}

// Execute calls the API using the request data configured.
func (r APIDeleteServerPoolRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteServerPoolExecute(r)
}

/*
DeleteServerPool Delete a server pool

Deletes a specific pool for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param poolName Name for the Pool.
 @return APIDeleteServerPoolRequest
*/
func (a *PoolAPIService) DeleteServerPool(ctx context.Context, serviceID string, versionID int32, poolName string) APIDeleteServerPoolRequest {
	return APIDeleteServerPoolRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		poolName:   poolName,
	}
}

// DeleteServerPoolExecute executes the request
//  @return InlineResponse200
func (a *PoolAPIService) DeleteServerPoolExecute(r APIDeleteServerPoolRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolAPIService.DeleteServerPool")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/pool/{pool_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_name"+"}", gourl.PathEscape(parameterToString(r.poolName, "")))

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

// APIGetServerPoolRequest represents a request for the resource.
type APIGetServerPoolRequest struct {
	ctx        context.Context
	APIService PoolAPI
	serviceID  string
	versionID  int32
	poolName   string
}

// Execute calls the API using the request data configured.
func (r APIGetServerPoolRequest) Execute() (*PoolResponse, *http.Response, error) {
	return r.APIService.GetServerPoolExecute(r)
}

/*
GetServerPool Get a server pool

Gets a single pool for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param poolName Name for the Pool.
 @return APIGetServerPoolRequest
*/
func (a *PoolAPIService) GetServerPool(ctx context.Context, serviceID string, versionID int32, poolName string) APIGetServerPoolRequest {
	return APIGetServerPoolRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		poolName:   poolName,
	}
}

// GetServerPoolExecute executes the request
//  @return PoolResponse
func (a *PoolAPIService) GetServerPoolExecute(r APIGetServerPoolRequest) (*PoolResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PoolResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolAPIService.GetServerPool")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/pool/{pool_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_name"+"}", gourl.PathEscape(parameterToString(r.poolName, "")))

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

// APIListServerPoolsRequest represents a request for the resource.
type APIListServerPoolsRequest struct {
	ctx        context.Context
	APIService PoolAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListServerPoolsRequest) Execute() ([]PoolResponse, *http.Response, error) {
	return r.APIService.ListServerPoolsExecute(r)
}

/*
ListServerPools List server pools

Lists all pools for a particular service and pool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListServerPoolsRequest
*/
func (a *PoolAPIService) ListServerPools(ctx context.Context, serviceID string, versionID int32) APIListServerPoolsRequest {
	return APIListServerPoolsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListServerPoolsExecute executes the request
//  @return []PoolResponse
func (a *PoolAPIService) ListServerPoolsExecute(r APIListServerPoolsRequest) ([]PoolResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []PoolResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolAPIService.ListServerPools")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/pool"
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

// APIUpdateServerPoolRequest represents a request for the resource.
type APIUpdateServerPoolRequest struct {
	ctx                 context.Context
	APIService          PoolAPI
	serviceID           string
	versionID           int32
	poolName            string
	tlsCaCert           *string
	tlsClientCert       *string
	tlsClientKey        *string
	tlsCertHostname     *string
	useTLS              *int32
	createdAt           *time.Time
	deletedAt           *time.Time
	updatedAt           *time.Time
	serviceID2          *string
	version             *string
	name                *string
	shield              *string
	requestCondition    *string
	tlsCiphers          *string
	tlsSniHostname      *string
	minTLSVersion       *int32
	maxTLSVersion       *int32
	healthcheck         *string
	comment             *string
	resourceType        *string
	overrideHost        *string
	betweenBytesTimeout *int32
	connectTimeout      *int32
	firstByteTimeout    *int32
	maxConnDefault      *int32
	quorum              *int32
	tlsCheckCert        *int32
}

// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APIUpdateServerPoolRequest) TLSCaCert(tlsCaCert string) *APIUpdateServerPoolRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}

// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateServerPoolRequest) TLSClientCert(tlsClientCert string) *APIUpdateServerPoolRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}

// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateServerPoolRequest) TLSClientKey(tlsClientKey string) *APIUpdateServerPoolRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}

// TLSCertHostname The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
func (r *APIUpdateServerPoolRequest) TLSCertHostname(tlsCertHostname string) *APIUpdateServerPoolRequest {
	r.tlsCertHostname = &tlsCertHostname
	return r
}

// UseTLS Whether to use TLS.
func (r *APIUpdateServerPoolRequest) UseTLS(useTLS int32) *APIUpdateServerPoolRequest {
	r.useTLS = &useTLS
	return r
}

// CreatedAt Date and time in ISO 8601 format.
func (r *APIUpdateServerPoolRequest) CreatedAt(createdAt time.Time) *APIUpdateServerPoolRequest {
	r.createdAt = &createdAt
	return r
}

// DeletedAt Date and time in ISO 8601 format.
func (r *APIUpdateServerPoolRequest) DeletedAt(deletedAt time.Time) *APIUpdateServerPoolRequest {
	r.deletedAt = &deletedAt
	return r
}

// UpdatedAt Date and time in ISO 8601 format.
func (r *APIUpdateServerPoolRequest) UpdatedAt(updatedAt time.Time) *APIUpdateServerPoolRequest {
	r.updatedAt = &updatedAt
	return r
}

// ServiceID2 returns a pointer to a request.
func (r *APIUpdateServerPoolRequest) ServiceID2(serviceID2 string) *APIUpdateServerPoolRequest {
	r.serviceID2 = &serviceID2
	return r
}

// Version returns a pointer to a request.
func (r *APIUpdateServerPoolRequest) Version(version string) *APIUpdateServerPoolRequest {
	r.version = &version
	return r
}

// Name Name for the Pool.
func (r *APIUpdateServerPoolRequest) Name(name string) *APIUpdateServerPoolRequest {
	r.name = &name
	return r
}

// Shield Selected POP to serve as a shield for the servers. Defaults to &#x60;null&#x60; meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding.
func (r *APIUpdateServerPoolRequest) Shield(shield string) *APIUpdateServerPoolRequest {
	r.shield = &shield
	return r
}

// RequestCondition Condition which, if met, will select this configuration during a request. Optional.
func (r *APIUpdateServerPoolRequest) RequestCondition(requestCondition string) *APIUpdateServerPoolRequest {
	r.requestCondition = &requestCondition
	return r
}

// TLSCiphers List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional.
func (r *APIUpdateServerPoolRequest) TLSCiphers(tlsCiphers string) *APIUpdateServerPoolRequest {
	r.tlsCiphers = &tlsCiphers
	return r
}

// TLSSniHostname SNI hostname. Optional.
func (r *APIUpdateServerPoolRequest) TLSSniHostname(tlsSniHostname string) *APIUpdateServerPoolRequest {
	r.tlsSniHostname = &tlsSniHostname
	return r
}

// MinTLSVersion Minimum allowed TLS version on connections to this server. Optional.
func (r *APIUpdateServerPoolRequest) MinTLSVersion(minTLSVersion int32) *APIUpdateServerPoolRequest {
	r.minTLSVersion = &minTLSVersion
	return r
}

// MaxTLSVersion Maximum allowed TLS version on connections to this server. Optional.
func (r *APIUpdateServerPoolRequest) MaxTLSVersion(maxTLSVersion int32) *APIUpdateServerPoolRequest {
	r.maxTLSVersion = &maxTLSVersion
	return r
}

// Healthcheck Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools.
func (r *APIUpdateServerPoolRequest) Healthcheck(healthcheck string) *APIUpdateServerPoolRequest {
	r.healthcheck = &healthcheck
	return r
}

// Comment A freeform descriptive note.
func (r *APIUpdateServerPoolRequest) Comment(comment string) *APIUpdateServerPoolRequest {
	r.comment = &comment
	return r
}

// ResourceType What type of load balance group to use.
func (r *APIUpdateServerPoolRequest) ResourceType(resourceType string) *APIUpdateServerPoolRequest {
	r.resourceType = &resourceType
	return r
}

// OverrideHost The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to &#x60;null&#x60; meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting.
func (r *APIUpdateServerPoolRequest) OverrideHost(overrideHost string) *APIUpdateServerPoolRequest {
	r.overrideHost = &overrideHost
	return r
}

// BetweenBytesTimeout Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using &#x60;bereq.between_bytes_timeout&#x60;.
func (r *APIUpdateServerPoolRequest) BetweenBytesTimeout(betweenBytesTimeout int32) *APIUpdateServerPoolRequest {
	r.betweenBytesTimeout = &betweenBytesTimeout
	return r
}

// ConnectTimeout How long to wait for a timeout in milliseconds. Optional.
func (r *APIUpdateServerPoolRequest) ConnectTimeout(connectTimeout int32) *APIUpdateServerPoolRequest {
	r.connectTimeout = &connectTimeout
	return r
}

// FirstByteTimeout How long to wait for the first byte in milliseconds. Optional.
func (r *APIUpdateServerPoolRequest) FirstByteTimeout(firstByteTimeout int32) *APIUpdateServerPoolRequest {
	r.firstByteTimeout = &firstByteTimeout
	return r
}

// MaxConnDefault Maximum number of connections. Optional.
func (r *APIUpdateServerPoolRequest) MaxConnDefault(maxConnDefault int32) *APIUpdateServerPoolRequest {
	r.maxConnDefault = &maxConnDefault
	return r
}

// Quorum Percentage of capacity (&#x60;0-100&#x60;) that needs to be operationally available for a pool to be considered up.
func (r *APIUpdateServerPoolRequest) Quorum(quorum int32) *APIUpdateServerPoolRequest {
	r.quorum = &quorum
	return r
}

// TLSCheckCert Be strict on checking TLS certs. Optional.
func (r *APIUpdateServerPoolRequest) TLSCheckCert(tlsCheckCert int32) *APIUpdateServerPoolRequest {
	r.tlsCheckCert = &tlsCheckCert
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateServerPoolRequest) Execute() (*PoolResponse, *http.Response, error) {
	return r.APIService.UpdateServerPoolExecute(r)
}

/*
UpdateServerPool Update a server pool

Updates a specific pool for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param poolName Name for the Pool.
 @return APIUpdateServerPoolRequest
*/
func (a *PoolAPIService) UpdateServerPool(ctx context.Context, serviceID string, versionID int32, poolName string) APIUpdateServerPoolRequest {
	return APIUpdateServerPoolRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		poolName:   poolName,
	}
}

// UpdateServerPoolExecute executes the request
//  @return PoolResponse
func (a *PoolAPIService) UpdateServerPoolExecute(r APIUpdateServerPoolRequest) (*PoolResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PoolResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolAPIService.UpdateServerPool")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/pool/{pool_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_name"+"}", gourl.PathEscape(parameterToString(r.poolName, "")))

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
	if r.tlsCaCert != nil {
		localVarFormParams.Add("tls_ca_cert", parameterToString(*r.tlsCaCert, ""))
	}
	if r.tlsClientCert != nil {
		localVarFormParams.Add("tls_client_cert", parameterToString(*r.tlsClientCert, ""))
	}
	if r.tlsClientKey != nil {
		localVarFormParams.Add("tls_client_key", parameterToString(*r.tlsClientKey, ""))
	}
	if r.tlsCertHostname != nil {
		localVarFormParams.Add("tls_cert_hostname", parameterToString(*r.tlsCertHostname, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
	}
	if r.createdAt != nil {
		localVarFormParams.Add("created_at", parameterToString(*r.createdAt, ""))
	}
	if r.deletedAt != nil {
		localVarFormParams.Add("deleted_at", parameterToString(*r.deletedAt, ""))
	}
	if r.updatedAt != nil {
		localVarFormParams.Add("updated_at", parameterToString(*r.updatedAt, ""))
	}
	if r.serviceID2 != nil {
		paramJSON, err := parameterToJSON(*r.serviceID2)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("service_id", paramJSON)
	}
	if r.version != nil {
		paramJSON, err := parameterToJSON(*r.version)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("version", paramJSON)
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.shield != nil {
		localVarFormParams.Add("shield", parameterToString(*r.shield, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.tlsCiphers != nil {
		localVarFormParams.Add("tls_ciphers", parameterToString(*r.tlsCiphers, ""))
	}
	if r.tlsSniHostname != nil {
		localVarFormParams.Add("tls_sni_hostname", parameterToString(*r.tlsSniHostname, ""))
	}
	if r.minTLSVersion != nil {
		localVarFormParams.Add("min_tls_version", parameterToString(*r.minTLSVersion, ""))
	}
	if r.maxTLSVersion != nil {
		localVarFormParams.Add("max_tls_version", parameterToString(*r.maxTLSVersion, ""))
	}
	if r.healthcheck != nil {
		localVarFormParams.Add("healthcheck", parameterToString(*r.healthcheck, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.resourceType != nil {
		localVarFormParams.Add("type", parameterToString(*r.resourceType, ""))
	}
	if r.overrideHost != nil {
		localVarFormParams.Add("override_host", parameterToString(*r.overrideHost, ""))
	}
	if r.betweenBytesTimeout != nil {
		localVarFormParams.Add("between_bytes_timeout", parameterToString(*r.betweenBytesTimeout, ""))
	}
	if r.connectTimeout != nil {
		localVarFormParams.Add("connect_timeout", parameterToString(*r.connectTimeout, ""))
	}
	if r.firstByteTimeout != nil {
		localVarFormParams.Add("first_byte_timeout", parameterToString(*r.firstByteTimeout, ""))
	}
	if r.maxConnDefault != nil {
		localVarFormParams.Add("max_conn_default", parameterToString(*r.maxConnDefault, ""))
	}
	if r.quorum != nil {
		localVarFormParams.Add("quorum", parameterToString(*r.quorum, ""))
	}
	if r.tlsCheckCert != nil {
		localVarFormParams.Add("tls_check_cert", parameterToString(*r.tlsCheckCert, ""))
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
