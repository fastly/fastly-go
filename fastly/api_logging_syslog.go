// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

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

// LoggingSyslogAPI defines an interface for interacting with the resource.
type LoggingSyslogAPI interface {

	/*
	CreateLogSyslog Create a syslog log endpoint

	Create a Syslog for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogSyslogRequest
	*/
	CreateLogSyslog(ctx context.Context, serviceID string, versionID int32) APICreateLogSyslogRequest

	// CreateLogSyslogExecute executes the request
	//  @return LoggingSyslogResponse
	CreateLogSyslogExecute(r APICreateLogSyslogRequest) (*LoggingSyslogResponse, *http.Response, error)

	/*
	DeleteLogSyslog Delete a syslog log endpoint

	Delete the Syslog for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSyslogName The name for the real-time logging configuration.
	 @return APIDeleteLogSyslogRequest
	*/
	DeleteLogSyslog(ctx context.Context, serviceID string, versionID int32, loggingSyslogName string) APIDeleteLogSyslogRequest

	// DeleteLogSyslogExecute executes the request
	//  @return InlineResponse200
	DeleteLogSyslogExecute(r APIDeleteLogSyslogRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogSyslog Get a syslog log endpoint

	Get the Syslog for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSyslogName The name for the real-time logging configuration.
	 @return APIGetLogSyslogRequest
	*/
	GetLogSyslog(ctx context.Context, serviceID string, versionID int32, loggingSyslogName string) APIGetLogSyslogRequest

	// GetLogSyslogExecute executes the request
	//  @return LoggingSyslogResponse
	GetLogSyslogExecute(r APIGetLogSyslogRequest) (*LoggingSyslogResponse, *http.Response, error)

	/*
	ListLogSyslog List Syslog log endpoints

	List all of the Syslogs for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogSyslogRequest
	*/
	ListLogSyslog(ctx context.Context, serviceID string, versionID int32) APIListLogSyslogRequest

	// ListLogSyslogExecute executes the request
	//  @return []LoggingSyslogResponse
	ListLogSyslogExecute(r APIListLogSyslogRequest) ([]LoggingSyslogResponse, *http.Response, error)

	/*
	UpdateLogSyslog Update a syslog log endpoint

	Update the Syslog for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSyslogName The name for the real-time logging configuration.
	 @return APIUpdateLogSyslogRequest
	*/
	UpdateLogSyslog(ctx context.Context, serviceID string, versionID int32, loggingSyslogName string) APIUpdateLogSyslogRequest

	// UpdateLogSyslogExecute executes the request
	//  @return LoggingSyslogResponse
	UpdateLogSyslogExecute(r APIUpdateLogSyslogRequest) (*LoggingSyslogResponse, *http.Response, error)
}

// LoggingSyslogAPIService LoggingSyslogAPI service
type LoggingSyslogAPIService service

// APICreateLogSyslogRequest represents a request for the resource.
type APICreateLogSyslogRequest struct {
	ctx context.Context
	APIService LoggingSyslogAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	tlsCaCert *string
	tlsClientCert *string
	tlsClientKey *string
	tlsHostname *string
	address *string
	port *int32
	messageType *LoggingMessageType
	hostname *string
	ipv4 *string
	token *string
	useTLS *LoggingUseTLS
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogSyslogRequest) Name(name string) *APICreateLogSyslogRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogSyslogRequest) Placement(placement string) *APICreateLogSyslogRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogSyslogRequest) ResponseCondition(responseCondition string) *APICreateLogSyslogRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogSyslogRequest) Format(format string) *APICreateLogSyslogRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogSyslogRequest) FormatVersion(formatVersion int32) *APICreateLogSyslogRequest {
	r.formatVersion = &formatVersion
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateLogSyslogRequest) TLSCaCert(tlsCaCert string) *APICreateLogSyslogRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogSyslogRequest) TLSClientCert(tlsClientCert string) *APICreateLogSyslogRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogSyslogRequest) TLSClientKey(tlsClientKey string) *APICreateLogSyslogRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APICreateLogSyslogRequest) TLSHostname(tlsHostname string) *APICreateLogSyslogRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// Address A hostname or IPv4 address.
func (r *APICreateLogSyslogRequest) Address(address string) *APICreateLogSyslogRequest {
	r.address = &address
	return r
}
// Port The port number.
func (r *APICreateLogSyslogRequest) Port(port int32) *APICreateLogSyslogRequest {
	r.port = &port
	return r
}
// MessageType returns a pointer to a request.
func (r *APICreateLogSyslogRequest) MessageType(messageType LoggingMessageType) *APICreateLogSyslogRequest {
	r.messageType = &messageType
	return r
}
// Hostname The hostname used for the syslog endpoint.
func (r *APICreateLogSyslogRequest) Hostname(hostname string) *APICreateLogSyslogRequest {
	r.hostname = &hostname
	return r
}
// Ipv4 The IPv4 address used for the syslog endpoint.
func (r *APICreateLogSyslogRequest) Ipv4(ipv4 string) *APICreateLogSyslogRequest {
	r.ipv4 = &ipv4
	return r
}
// Token Whether to prepend each message with a specific token.
func (r *APICreateLogSyslogRequest) Token(token string) *APICreateLogSyslogRequest {
	r.token = &token
	return r
}
// UseTLS returns a pointer to a request.
func (r *APICreateLogSyslogRequest) UseTLS(useTLS LoggingUseTLS) *APICreateLogSyslogRequest {
	r.useTLS = &useTLS
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogSyslogRequest) Execute() (*LoggingSyslogResponse, *http.Response, error) {
	return r.APIService.CreateLogSyslogExecute(r)
}

/*
CreateLogSyslog Create a syslog log endpoint

Create a Syslog for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogSyslogRequest
*/
func (a *LoggingSyslogAPIService) CreateLogSyslog(ctx context.Context, serviceID string, versionID int32) APICreateLogSyslogRequest {
	return APICreateLogSyslogRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogSyslogExecute executes the request
//  @return LoggingSyslogResponse
func (a *LoggingSyslogAPIService) CreateLogSyslogExecute(r APICreateLogSyslogRequest) (*LoggingSyslogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSyslogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSyslogAPIService.CreateLogSyslog")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/syslog"
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
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.placement != nil {
		localVarFormParams.Add("placement", parameterToString(*r.placement, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
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
	if r.tlsHostname != nil {
		localVarFormParams.Add("tls_hostname", parameterToString(*r.tlsHostname, ""))
	}
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.messageType != nil {
		localVarFormParams.Add("message_type", parameterToString(*r.messageType, ""))
	}
	if r.hostname != nil {
		localVarFormParams.Add("hostname", parameterToString(*r.hostname, ""))
	}
	if r.ipv4 != nil {
		localVarFormParams.Add("ipv4", parameterToString(*r.ipv4, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
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

// APIDeleteLogSyslogRequest represents a request for the resource.
type APIDeleteLogSyslogRequest struct {
	ctx context.Context
	APIService LoggingSyslogAPI
	serviceID string
	versionID int32
	loggingSyslogName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogSyslogRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogSyslogExecute(r)
}

/*
DeleteLogSyslog Delete a syslog log endpoint

Delete the Syslog for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSyslogName The name for the real-time logging configuration.
 @return APIDeleteLogSyslogRequest
*/
func (a *LoggingSyslogAPIService) DeleteLogSyslog(ctx context.Context, serviceID string, versionID int32, loggingSyslogName string) APIDeleteLogSyslogRequest {
	return APIDeleteLogSyslogRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSyslogName: loggingSyslogName,
	}
}

// DeleteLogSyslogExecute executes the request
//  @return InlineResponse200
func (a *LoggingSyslogAPIService) DeleteLogSyslogExecute(r APIDeleteLogSyslogRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSyslogAPIService.DeleteLogSyslog")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_syslog_name"+"}", gourl.PathEscape(parameterToString(r.loggingSyslogName, "")))

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

// APIGetLogSyslogRequest represents a request for the resource.
type APIGetLogSyslogRequest struct {
	ctx context.Context
	APIService LoggingSyslogAPI
	serviceID string
	versionID int32
	loggingSyslogName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogSyslogRequest) Execute() (*LoggingSyslogResponse, *http.Response, error) {
	return r.APIService.GetLogSyslogExecute(r)
}

/*
GetLogSyslog Get a syslog log endpoint

Get the Syslog for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSyslogName The name for the real-time logging configuration.
 @return APIGetLogSyslogRequest
*/
func (a *LoggingSyslogAPIService) GetLogSyslog(ctx context.Context, serviceID string, versionID int32, loggingSyslogName string) APIGetLogSyslogRequest {
	return APIGetLogSyslogRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSyslogName: loggingSyslogName,
	}
}

// GetLogSyslogExecute executes the request
//  @return LoggingSyslogResponse
func (a *LoggingSyslogAPIService) GetLogSyslogExecute(r APIGetLogSyslogRequest) (*LoggingSyslogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSyslogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSyslogAPIService.GetLogSyslog")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_syslog_name"+"}", gourl.PathEscape(parameterToString(r.loggingSyslogName, "")))

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

// APIListLogSyslogRequest represents a request for the resource.
type APIListLogSyslogRequest struct {
	ctx context.Context
	APIService LoggingSyslogAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogSyslogRequest) Execute() ([]LoggingSyslogResponse, *http.Response, error) {
	return r.APIService.ListLogSyslogExecute(r)
}

/*
ListLogSyslog List Syslog log endpoints

List all of the Syslogs for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogSyslogRequest
*/
func (a *LoggingSyslogAPIService) ListLogSyslog(ctx context.Context, serviceID string, versionID int32) APIListLogSyslogRequest {
	return APIListLogSyslogRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogSyslogExecute executes the request
//  @return []LoggingSyslogResponse
func (a *LoggingSyslogAPIService) ListLogSyslogExecute(r APIListLogSyslogRequest) ([]LoggingSyslogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingSyslogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSyslogAPIService.ListLogSyslog")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/syslog"
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

// APIUpdateLogSyslogRequest represents a request for the resource.
type APIUpdateLogSyslogRequest struct {
	ctx context.Context
	APIService LoggingSyslogAPI
	serviceID string
	versionID int32
	loggingSyslogName string
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	tlsCaCert *string
	tlsClientCert *string
	tlsClientKey *string
	tlsHostname *string
	address *string
	port *int32
	messageType *LoggingMessageType
	hostname *string
	ipv4 *string
	token *string
	useTLS *LoggingUseTLS
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogSyslogRequest) Name(name string) *APIUpdateLogSyslogRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogSyslogRequest) Placement(placement string) *APIUpdateLogSyslogRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogSyslogRequest) ResponseCondition(responseCondition string) *APIUpdateLogSyslogRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogSyslogRequest) Format(format string) *APIUpdateLogSyslogRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogSyslogRequest) FormatVersion(formatVersion int32) *APIUpdateLogSyslogRequest {
	r.formatVersion = &formatVersion
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APIUpdateLogSyslogRequest) TLSCaCert(tlsCaCert string) *APIUpdateLogSyslogRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogSyslogRequest) TLSClientCert(tlsClientCert string) *APIUpdateLogSyslogRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogSyslogRequest) TLSClientKey(tlsClientKey string) *APIUpdateLogSyslogRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APIUpdateLogSyslogRequest) TLSHostname(tlsHostname string) *APIUpdateLogSyslogRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// Address A hostname or IPv4 address.
func (r *APIUpdateLogSyslogRequest) Address(address string) *APIUpdateLogSyslogRequest {
	r.address = &address
	return r
}
// Port The port number.
func (r *APIUpdateLogSyslogRequest) Port(port int32) *APIUpdateLogSyslogRequest {
	r.port = &port
	return r
}
// MessageType returns a pointer to a request.
func (r *APIUpdateLogSyslogRequest) MessageType(messageType LoggingMessageType) *APIUpdateLogSyslogRequest {
	r.messageType = &messageType
	return r
}
// Hostname The hostname used for the syslog endpoint.
func (r *APIUpdateLogSyslogRequest) Hostname(hostname string) *APIUpdateLogSyslogRequest {
	r.hostname = &hostname
	return r
}
// Ipv4 The IPv4 address used for the syslog endpoint.
func (r *APIUpdateLogSyslogRequest) Ipv4(ipv4 string) *APIUpdateLogSyslogRequest {
	r.ipv4 = &ipv4
	return r
}
// Token Whether to prepend each message with a specific token.
func (r *APIUpdateLogSyslogRequest) Token(token string) *APIUpdateLogSyslogRequest {
	r.token = &token
	return r
}
// UseTLS returns a pointer to a request.
func (r *APIUpdateLogSyslogRequest) UseTLS(useTLS LoggingUseTLS) *APIUpdateLogSyslogRequest {
	r.useTLS = &useTLS
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogSyslogRequest) Execute() (*LoggingSyslogResponse, *http.Response, error) {
	return r.APIService.UpdateLogSyslogExecute(r)
}

/*
UpdateLogSyslog Update a syslog log endpoint

Update the Syslog for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSyslogName The name for the real-time logging configuration.
 @return APIUpdateLogSyslogRequest
*/
func (a *LoggingSyslogAPIService) UpdateLogSyslog(ctx context.Context, serviceID string, versionID int32, loggingSyslogName string) APIUpdateLogSyslogRequest {
	return APIUpdateLogSyslogRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSyslogName: loggingSyslogName,
	}
}

// UpdateLogSyslogExecute executes the request
//  @return LoggingSyslogResponse
func (a *LoggingSyslogAPIService) UpdateLogSyslogExecute(r APIUpdateLogSyslogRequest) (*LoggingSyslogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSyslogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSyslogAPIService.UpdateLogSyslog")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_syslog_name"+"}", gourl.PathEscape(parameterToString(r.loggingSyslogName, "")))

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
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.placement != nil {
		localVarFormParams.Add("placement", parameterToString(*r.placement, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
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
	if r.tlsHostname != nil {
		localVarFormParams.Add("tls_hostname", parameterToString(*r.tlsHostname, ""))
	}
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.messageType != nil {
		localVarFormParams.Add("message_type", parameterToString(*r.messageType, ""))
	}
	if r.hostname != nil {
		localVarFormParams.Add("hostname", parameterToString(*r.hostname, ""))
	}
	if r.ipv4 != nil {
		localVarFormParams.Add("ipv4", parameterToString(*r.ipv4, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
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
