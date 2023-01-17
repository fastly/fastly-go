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

// LoggingSplunkAPI defines an interface for interacting with the resource.
type LoggingSplunkAPI interface {

	/*
	CreateLogSplunk Create a Splunk log endpoint

	Create a Splunk logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogSplunkRequest
	*/
	CreateLogSplunk(ctx context.Context, serviceID string, versionID int32) APICreateLogSplunkRequest

	// CreateLogSplunkExecute executes the request
	//  @return LoggingSplunkResponse
	CreateLogSplunkExecute(r APICreateLogSplunkRequest) (*LoggingSplunkResponse, *http.Response, error)

	/*
	DeleteLogSplunk Delete a Splunk log endpoint

	Delete the Splunk logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSplunkName The name for the real-time logging configuration.
	 @return APIDeleteLogSplunkRequest
	*/
	DeleteLogSplunk(ctx context.Context, serviceID string, versionID int32, loggingSplunkName string) APIDeleteLogSplunkRequest

	// DeleteLogSplunkExecute executes the request
	//  @return InlineResponse200
	DeleteLogSplunkExecute(r APIDeleteLogSplunkRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogSplunk Get a Splunk log endpoint

	Get the details for a Splunk logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSplunkName The name for the real-time logging configuration.
	 @return APIGetLogSplunkRequest
	*/
	GetLogSplunk(ctx context.Context, serviceID string, versionID int32, loggingSplunkName string) APIGetLogSplunkRequest

	// GetLogSplunkExecute executes the request
	//  @return LoggingSplunkResponse
	GetLogSplunkExecute(r APIGetLogSplunkRequest) (*LoggingSplunkResponse, *http.Response, error)

	/*
	ListLogSplunk List Splunk log endpoints

	List all of the Splunk logging objects for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogSplunkRequest
	*/
	ListLogSplunk(ctx context.Context, serviceID string, versionID int32) APIListLogSplunkRequest

	// ListLogSplunkExecute executes the request
	//  @return []LoggingSplunkResponse
	ListLogSplunkExecute(r APIListLogSplunkRequest) ([]LoggingSplunkResponse, *http.Response, error)

	/*
	UpdateLogSplunk Update a Splunk log endpoint

	Update the Splunk logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSplunkName The name for the real-time logging configuration.
	 @return APIUpdateLogSplunkRequest
	*/
	UpdateLogSplunk(ctx context.Context, serviceID string, versionID int32, loggingSplunkName string) APIUpdateLogSplunkRequest

	// UpdateLogSplunkExecute executes the request
	//  @return LoggingSplunkResponse
	UpdateLogSplunkExecute(r APIUpdateLogSplunkRequest) (*LoggingSplunkResponse, *http.Response, error)
}

// LoggingSplunkAPIService LoggingSplunkAPI service
type LoggingSplunkAPIService service

// APICreateLogSplunkRequest represents a request for the resource.
type APICreateLogSplunkRequest struct {
	ctx context.Context
	APIService LoggingSplunkAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	formatVersion *int32
	responseCondition *string
	format *string
	tlsCaCert *string
	tlsClientCert *string
	tlsClientKey *string
	tlsHostname *string
	requestMaxEntries *int32
	requestMaxBytes *int32
	url *string
	token *string
	useTLS *LoggingUseTLS
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogSplunkRequest) Name(name string) *APICreateLogSplunkRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogSplunkRequest) Placement(placement string) *APICreateLogSplunkRequest {
	r.placement = &placement
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogSplunkRequest) FormatVersion(formatVersion int32) *APICreateLogSplunkRequest {
	r.formatVersion = &formatVersion
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogSplunkRequest) ResponseCondition(responseCondition string) *APICreateLogSplunkRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogSplunkRequest) Format(format string) *APICreateLogSplunkRequest {
	r.format = &format
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateLogSplunkRequest) TLSCaCert(tlsCaCert string) *APICreateLogSplunkRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogSplunkRequest) TLSClientCert(tlsClientCert string) *APICreateLogSplunkRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogSplunkRequest) TLSClientKey(tlsClientKey string) *APICreateLogSplunkRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APICreateLogSplunkRequest) TLSHostname(tlsHostname string) *APICreateLogSplunkRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APICreateLogSplunkRequest) RequestMaxEntries(requestMaxEntries int32) *APICreateLogSplunkRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}
// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APICreateLogSplunkRequest) RequestMaxBytes(requestMaxBytes int32) *APICreateLogSplunkRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}
// URL The URL to post logs to.
func (r *APICreateLogSplunkRequest) URL(url string) *APICreateLogSplunkRequest {
	r.url = &url
	return r
}
// Token A Splunk token for use in posting logs over HTTP to your collector.
func (r *APICreateLogSplunkRequest) Token(token string) *APICreateLogSplunkRequest {
	r.token = &token
	return r
}
// UseTLS returns a pointer to a request.
func (r *APICreateLogSplunkRequest) UseTLS(useTLS LoggingUseTLS) *APICreateLogSplunkRequest {
	r.useTLS = &useTLS
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogSplunkRequest) Execute() (*LoggingSplunkResponse, *http.Response, error) {
	return r.APIService.CreateLogSplunkExecute(r)
}

/*
CreateLogSplunk Create a Splunk log endpoint

Create a Splunk logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogSplunkRequest
*/
func (a *LoggingSplunkAPIService) CreateLogSplunk(ctx context.Context, serviceID string, versionID int32) APICreateLogSplunkRequest {
	return APICreateLogSplunkRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogSplunkExecute executes the request
//  @return LoggingSplunkResponse
func (a *LoggingSplunkAPIService) CreateLogSplunkExecute(r APICreateLogSplunkRequest) (*LoggingSplunkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSplunkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSplunkAPIService.CreateLogSplunk")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/splunk"
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
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
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
	if r.requestMaxEntries != nil {
		localVarFormParams.Add("request_max_entries", parameterToString(*r.requestMaxEntries, ""))
	}
	if r.requestMaxBytes != nil {
		localVarFormParams.Add("request_max_bytes", parameterToString(*r.requestMaxBytes, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
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

// APIDeleteLogSplunkRequest represents a request for the resource.
type APIDeleteLogSplunkRequest struct {
	ctx context.Context
	APIService LoggingSplunkAPI
	serviceID string
	versionID int32
	loggingSplunkName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogSplunkRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogSplunkExecute(r)
}

/*
DeleteLogSplunk Delete a Splunk log endpoint

Delete the Splunk logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSplunkName The name for the real-time logging configuration.
 @return APIDeleteLogSplunkRequest
*/
func (a *LoggingSplunkAPIService) DeleteLogSplunk(ctx context.Context, serviceID string, versionID int32, loggingSplunkName string) APIDeleteLogSplunkRequest {
	return APIDeleteLogSplunkRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSplunkName: loggingSplunkName,
	}
}

// DeleteLogSplunkExecute executes the request
//  @return InlineResponse200
func (a *LoggingSplunkAPIService) DeleteLogSplunkExecute(r APIDeleteLogSplunkRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSplunkAPIService.DeleteLogSplunk")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_splunk_name"+"}", gourl.PathEscape(parameterToString(r.loggingSplunkName, "")))

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

// APIGetLogSplunkRequest represents a request for the resource.
type APIGetLogSplunkRequest struct {
	ctx context.Context
	APIService LoggingSplunkAPI
	serviceID string
	versionID int32
	loggingSplunkName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogSplunkRequest) Execute() (*LoggingSplunkResponse, *http.Response, error) {
	return r.APIService.GetLogSplunkExecute(r)
}

/*
GetLogSplunk Get a Splunk log endpoint

Get the details for a Splunk logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSplunkName The name for the real-time logging configuration.
 @return APIGetLogSplunkRequest
*/
func (a *LoggingSplunkAPIService) GetLogSplunk(ctx context.Context, serviceID string, versionID int32, loggingSplunkName string) APIGetLogSplunkRequest {
	return APIGetLogSplunkRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSplunkName: loggingSplunkName,
	}
}

// GetLogSplunkExecute executes the request
//  @return LoggingSplunkResponse
func (a *LoggingSplunkAPIService) GetLogSplunkExecute(r APIGetLogSplunkRequest) (*LoggingSplunkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSplunkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSplunkAPIService.GetLogSplunk")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_splunk_name"+"}", gourl.PathEscape(parameterToString(r.loggingSplunkName, "")))

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

// APIListLogSplunkRequest represents a request for the resource.
type APIListLogSplunkRequest struct {
	ctx context.Context
	APIService LoggingSplunkAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogSplunkRequest) Execute() ([]LoggingSplunkResponse, *http.Response, error) {
	return r.APIService.ListLogSplunkExecute(r)
}

/*
ListLogSplunk List Splunk log endpoints

List all of the Splunk logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogSplunkRequest
*/
func (a *LoggingSplunkAPIService) ListLogSplunk(ctx context.Context, serviceID string, versionID int32) APIListLogSplunkRequest {
	return APIListLogSplunkRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogSplunkExecute executes the request
//  @return []LoggingSplunkResponse
func (a *LoggingSplunkAPIService) ListLogSplunkExecute(r APIListLogSplunkRequest) ([]LoggingSplunkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingSplunkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSplunkAPIService.ListLogSplunk")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/splunk"
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

// APIUpdateLogSplunkRequest represents a request for the resource.
type APIUpdateLogSplunkRequest struct {
	ctx context.Context
	APIService LoggingSplunkAPI
	serviceID string
	versionID int32
	loggingSplunkName string
	name *string
	placement *string
	formatVersion *int32
	responseCondition *string
	format *string
	tlsCaCert *string
	tlsClientCert *string
	tlsClientKey *string
	tlsHostname *string
	requestMaxEntries *int32
	requestMaxBytes *int32
	url *string
	token *string
	useTLS *LoggingUseTLS
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogSplunkRequest) Name(name string) *APIUpdateLogSplunkRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogSplunkRequest) Placement(placement string) *APIUpdateLogSplunkRequest {
	r.placement = &placement
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogSplunkRequest) FormatVersion(formatVersion int32) *APIUpdateLogSplunkRequest {
	r.formatVersion = &formatVersion
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogSplunkRequest) ResponseCondition(responseCondition string) *APIUpdateLogSplunkRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogSplunkRequest) Format(format string) *APIUpdateLogSplunkRequest {
	r.format = &format
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APIUpdateLogSplunkRequest) TLSCaCert(tlsCaCert string) *APIUpdateLogSplunkRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogSplunkRequest) TLSClientCert(tlsClientCert string) *APIUpdateLogSplunkRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogSplunkRequest) TLSClientKey(tlsClientKey string) *APIUpdateLogSplunkRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APIUpdateLogSplunkRequest) TLSHostname(tlsHostname string) *APIUpdateLogSplunkRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APIUpdateLogSplunkRequest) RequestMaxEntries(requestMaxEntries int32) *APIUpdateLogSplunkRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}
// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APIUpdateLogSplunkRequest) RequestMaxBytes(requestMaxBytes int32) *APIUpdateLogSplunkRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}
// URL The URL to post logs to.
func (r *APIUpdateLogSplunkRequest) URL(url string) *APIUpdateLogSplunkRequest {
	r.url = &url
	return r
}
// Token A Splunk token for use in posting logs over HTTP to your collector.
func (r *APIUpdateLogSplunkRequest) Token(token string) *APIUpdateLogSplunkRequest {
	r.token = &token
	return r
}
// UseTLS returns a pointer to a request.
func (r *APIUpdateLogSplunkRequest) UseTLS(useTLS LoggingUseTLS) *APIUpdateLogSplunkRequest {
	r.useTLS = &useTLS
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogSplunkRequest) Execute() (*LoggingSplunkResponse, *http.Response, error) {
	return r.APIService.UpdateLogSplunkExecute(r)
}

/*
UpdateLogSplunk Update a Splunk log endpoint

Update the Splunk logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSplunkName The name for the real-time logging configuration.
 @return APIUpdateLogSplunkRequest
*/
func (a *LoggingSplunkAPIService) UpdateLogSplunk(ctx context.Context, serviceID string, versionID int32, loggingSplunkName string) APIUpdateLogSplunkRequest {
	return APIUpdateLogSplunkRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSplunkName: loggingSplunkName,
	}
}

// UpdateLogSplunkExecute executes the request
//  @return LoggingSplunkResponse
func (a *LoggingSplunkAPIService) UpdateLogSplunkExecute(r APIUpdateLogSplunkRequest) (*LoggingSplunkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSplunkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSplunkAPIService.UpdateLogSplunk")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_splunk_name"+"}", gourl.PathEscape(parameterToString(r.loggingSplunkName, "")))

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
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
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
	if r.requestMaxEntries != nil {
		localVarFormParams.Add("request_max_entries", parameterToString(*r.requestMaxEntries, ""))
	}
	if r.requestMaxBytes != nil {
		localVarFormParams.Add("request_max_bytes", parameterToString(*r.requestMaxBytes, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
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
