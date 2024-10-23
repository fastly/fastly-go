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

// LoggingHTTPSAPI defines an interface for interacting with the resource.
type LoggingHTTPSAPI interface {

	/*
		CreateLogHTTPS Create an HTTPS log endpoint

		Create an HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogHTTPSRequest
	*/
	CreateLogHTTPS(ctx context.Context, serviceID string, versionID int32) APICreateLogHTTPSRequest

	// CreateLogHTTPSExecute executes the request
	//  @return LoggingHTTPSResponse
	CreateLogHTTPSExecute(r APICreateLogHTTPSRequest) (*LoggingHTTPSResponse, *http.Response, error)

	/*
		DeleteLogHTTPS Delete an HTTPS log endpoint

		Delete the HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingHTTPSName The name for the real-time logging configuration.
		 @return APIDeleteLogHTTPSRequest
	*/
	DeleteLogHTTPS(ctx context.Context, serviceID string, versionID int32, loggingHTTPSName string) APIDeleteLogHTTPSRequest

	// DeleteLogHTTPSExecute executes the request
	//  @return InlineResponse200
	DeleteLogHTTPSExecute(r APIDeleteLogHTTPSRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogHTTPS Get an HTTPS log endpoint

		Get the HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingHTTPSName The name for the real-time logging configuration.
		 @return APIGetLogHTTPSRequest
	*/
	GetLogHTTPS(ctx context.Context, serviceID string, versionID int32, loggingHTTPSName string) APIGetLogHTTPSRequest

	// GetLogHTTPSExecute executes the request
	//  @return LoggingHTTPSResponse
	GetLogHTTPSExecute(r APIGetLogHTTPSRequest) (*LoggingHTTPSResponse, *http.Response, error)

	/*
		ListLogHTTPS List HTTPS log endpoints

		List all of the HTTPS objects for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogHTTPSRequest
	*/
	ListLogHTTPS(ctx context.Context, serviceID string, versionID int32) APIListLogHTTPSRequest

	// ListLogHTTPSExecute executes the request
	//  @return []LoggingHTTPSResponse
	ListLogHTTPSExecute(r APIListLogHTTPSRequest) ([]LoggingHTTPSResponse, *http.Response, error)

	/*
		UpdateLogHTTPS Update an HTTPS log endpoint

		Update the HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingHTTPSName The name for the real-time logging configuration.
		 @return APIUpdateLogHTTPSRequest
	*/
	UpdateLogHTTPS(ctx context.Context, serviceID string, versionID int32, loggingHTTPSName string) APIUpdateLogHTTPSRequest

	// UpdateLogHTTPSExecute executes the request
	//  @return LoggingHTTPSResponse
	UpdateLogHTTPSExecute(r APIUpdateLogHTTPSRequest) (*LoggingHTTPSResponse, *http.Response, error)
}

// LoggingHTTPSAPIService LoggingHTTPSAPI service
type LoggingHTTPSAPIService service

// APICreateLogHTTPSRequest represents a request for the resource.
type APICreateLogHTTPSRequest struct {
	ctx               context.Context
	APIService        LoggingHTTPSAPI
	serviceID         string
	versionID         int32
	name              *string
	placement         *string
	responseCondition *string
	format            *string
	formatVersion     *int32
	tlsCaCert         *string
	tlsClientCert     *string
	tlsClientKey      *string
	tlsHostname       *string
	requestMaxEntries *int32
	requestMaxBytes   *int32
	url               *string
	contentType       *string
	headerName        *string
	messageType       *LoggingMessageType
	headerValue       *string
	method            *string
	jsonFormat        *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogHTTPSRequest) Name(name string) *APICreateLogHTTPSRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogHTTPSRequest) Placement(placement string) *APICreateLogHTTPSRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogHTTPSRequest) ResponseCondition(responseCondition string) *APICreateLogHTTPSRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogHTTPSRequest) Format(format string) *APICreateLogHTTPSRequest {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogHTTPSRequest) FormatVersion(formatVersion int32) *APICreateLogHTTPSRequest {
	r.formatVersion = &formatVersion
	return r
}

// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateLogHTTPSRequest) TLSCaCert(tlsCaCert string) *APICreateLogHTTPSRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}

// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogHTTPSRequest) TLSClientCert(tlsClientCert string) *APICreateLogHTTPSRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}

// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogHTTPSRequest) TLSClientKey(tlsClientKey string) *APICreateLogHTTPSRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}

// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APICreateLogHTTPSRequest) TLSHostname(tlsHostname string) *APICreateLogHTTPSRequest {
	r.tlsHostname = &tlsHostname
	return r
}

// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; (10k).
func (r *APICreateLogHTTPSRequest) RequestMaxEntries(requestMaxEntries int32) *APICreateLogHTTPSRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}

// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; (100MB).
func (r *APICreateLogHTTPSRequest) RequestMaxBytes(requestMaxBytes int32) *APICreateLogHTTPSRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}

// URL The URL to send logs to. Must use HTTPS. Required.
func (r *APICreateLogHTTPSRequest) URL(url string) *APICreateLogHTTPSRequest {
	r.url = &url
	return r
}

// ContentType Content type of the header sent with the request.
func (r *APICreateLogHTTPSRequest) ContentType(contentType string) *APICreateLogHTTPSRequest {
	r.contentType = &contentType
	return r
}

// HeaderName Name of the custom header sent with the request.
func (r *APICreateLogHTTPSRequest) HeaderName(headerName string) *APICreateLogHTTPSRequest {
	r.headerName = &headerName
	return r
}

// MessageType returns a pointer to a request.
func (r *APICreateLogHTTPSRequest) MessageType(messageType LoggingMessageType) *APICreateLogHTTPSRequest {
	r.messageType = &messageType
	return r
}

// HeaderValue Value of the custom header sent with the request.
func (r *APICreateLogHTTPSRequest) HeaderValue(headerValue string) *APICreateLogHTTPSRequest {
	r.headerValue = &headerValue
	return r
}

// Method HTTP method used for request.
func (r *APICreateLogHTTPSRequest) Method(method string) *APICreateLogHTTPSRequest {
	r.method = &method
	return r
}

// JSONFormat Enforces valid JSON formatting for log entries.
func (r *APICreateLogHTTPSRequest) JSONFormat(jsonFormat string) *APICreateLogHTTPSRequest {
	r.jsonFormat = &jsonFormat
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogHTTPSRequest) Execute() (*LoggingHTTPSResponse, *http.Response, error) {
	return r.APIService.CreateLogHTTPSExecute(r)
}

/*
CreateLogHTTPS Create an HTTPS log endpoint

Create an HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogHTTPSRequest
*/
func (a *LoggingHTTPSAPIService) CreateLogHTTPS(ctx context.Context, serviceID string, versionID int32) APICreateLogHTTPSRequest {
	return APICreateLogHTTPSRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogHTTPSExecute executes the request
//  @return LoggingHTTPSResponse
func (a *LoggingHTTPSAPIService) CreateLogHTTPSExecute(r APICreateLogHTTPSRequest) (*LoggingHTTPSResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingHTTPSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHTTPSAPIService.CreateLogHTTPS")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https"
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
	if r.requestMaxEntries != nil {
		localVarFormParams.Add("request_max_entries", parameterToString(*r.requestMaxEntries, ""))
	}
	if r.requestMaxBytes != nil {
		localVarFormParams.Add("request_max_bytes", parameterToString(*r.requestMaxBytes, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
	}
	if r.contentType != nil {
		localVarFormParams.Add("content_type", parameterToString(*r.contentType, ""))
	}
	if r.headerName != nil {
		localVarFormParams.Add("header_name", parameterToString(*r.headerName, ""))
	}
	if r.messageType != nil {
		localVarFormParams.Add("message_type", parameterToString(*r.messageType, ""))
	}
	if r.headerValue != nil {
		localVarFormParams.Add("header_value", parameterToString(*r.headerValue, ""))
	}
	if r.method != nil {
		localVarFormParams.Add("method", parameterToString(*r.method, ""))
	}
	if r.jsonFormat != nil {
		localVarFormParams.Add("json_format", parameterToString(*r.jsonFormat, ""))
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

// APIDeleteLogHTTPSRequest represents a request for the resource.
type APIDeleteLogHTTPSRequest struct {
	ctx              context.Context
	APIService       LoggingHTTPSAPI
	serviceID        string
	versionID        int32
	loggingHTTPSName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogHTTPSRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogHTTPSExecute(r)
}

/*
DeleteLogHTTPS Delete an HTTPS log endpoint

Delete the HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingHTTPSName The name for the real-time logging configuration.
 @return APIDeleteLogHTTPSRequest
*/
func (a *LoggingHTTPSAPIService) DeleteLogHTTPS(ctx context.Context, serviceID string, versionID int32, loggingHTTPSName string) APIDeleteLogHTTPSRequest {
	return APIDeleteLogHTTPSRequest{
		APIService:       a,
		ctx:              ctx,
		serviceID:        serviceID,
		versionID:        versionID,
		loggingHTTPSName: loggingHTTPSName,
	}
}

// DeleteLogHTTPSExecute executes the request
//  @return InlineResponse200
func (a *LoggingHTTPSAPIService) DeleteLogHTTPSExecute(r APIDeleteLogHTTPSRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHTTPSAPIService.DeleteLogHTTPS")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_https_name"+"}", gourl.PathEscape(parameterToString(r.loggingHTTPSName, "")))

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

// APIGetLogHTTPSRequest represents a request for the resource.
type APIGetLogHTTPSRequest struct {
	ctx              context.Context
	APIService       LoggingHTTPSAPI
	serviceID        string
	versionID        int32
	loggingHTTPSName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogHTTPSRequest) Execute() (*LoggingHTTPSResponse, *http.Response, error) {
	return r.APIService.GetLogHTTPSExecute(r)
}

/*
GetLogHTTPS Get an HTTPS log endpoint

Get the HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingHTTPSName The name for the real-time logging configuration.
 @return APIGetLogHTTPSRequest
*/
func (a *LoggingHTTPSAPIService) GetLogHTTPS(ctx context.Context, serviceID string, versionID int32, loggingHTTPSName string) APIGetLogHTTPSRequest {
	return APIGetLogHTTPSRequest{
		APIService:       a,
		ctx:              ctx,
		serviceID:        serviceID,
		versionID:        versionID,
		loggingHTTPSName: loggingHTTPSName,
	}
}

// GetLogHTTPSExecute executes the request
//  @return LoggingHTTPSResponse
func (a *LoggingHTTPSAPIService) GetLogHTTPSExecute(r APIGetLogHTTPSRequest) (*LoggingHTTPSResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingHTTPSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHTTPSAPIService.GetLogHTTPS")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_https_name"+"}", gourl.PathEscape(parameterToString(r.loggingHTTPSName, "")))

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

// APIListLogHTTPSRequest represents a request for the resource.
type APIListLogHTTPSRequest struct {
	ctx        context.Context
	APIService LoggingHTTPSAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogHTTPSRequest) Execute() ([]LoggingHTTPSResponse, *http.Response, error) {
	return r.APIService.ListLogHTTPSExecute(r)
}

/*
ListLogHTTPS List HTTPS log endpoints

List all of the HTTPS objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogHTTPSRequest
*/
func (a *LoggingHTTPSAPIService) ListLogHTTPS(ctx context.Context, serviceID string, versionID int32) APIListLogHTTPSRequest {
	return APIListLogHTTPSRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogHTTPSExecute executes the request
//  @return []LoggingHTTPSResponse
func (a *LoggingHTTPSAPIService) ListLogHTTPSExecute(r APIListLogHTTPSRequest) ([]LoggingHTTPSResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingHTTPSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHTTPSAPIService.ListLogHTTPS")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https"
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

// APIUpdateLogHTTPSRequest represents a request for the resource.
type APIUpdateLogHTTPSRequest struct {
	ctx               context.Context
	APIService        LoggingHTTPSAPI
	serviceID         string
	versionID         int32
	loggingHTTPSName  string
	name              *string
	placement         *string
	responseCondition *string
	format            *string
	formatVersion     *int32
	tlsCaCert         *string
	tlsClientCert     *string
	tlsClientKey      *string
	tlsHostname       *string
	requestMaxEntries *int32
	requestMaxBytes   *int32
	url               *string
	contentType       *string
	headerName        *string
	messageType       *LoggingMessageType
	headerValue       *string
	method            *string
	jsonFormat        *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogHTTPSRequest) Name(name string) *APIUpdateLogHTTPSRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogHTTPSRequest) Placement(placement string) *APIUpdateLogHTTPSRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogHTTPSRequest) ResponseCondition(responseCondition string) *APIUpdateLogHTTPSRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogHTTPSRequest) Format(format string) *APIUpdateLogHTTPSRequest {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogHTTPSRequest) FormatVersion(formatVersion int32) *APIUpdateLogHTTPSRequest {
	r.formatVersion = &formatVersion
	return r
}

// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APIUpdateLogHTTPSRequest) TLSCaCert(tlsCaCert string) *APIUpdateLogHTTPSRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}

// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogHTTPSRequest) TLSClientCert(tlsClientCert string) *APIUpdateLogHTTPSRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}

// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogHTTPSRequest) TLSClientKey(tlsClientKey string) *APIUpdateLogHTTPSRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}

// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APIUpdateLogHTTPSRequest) TLSHostname(tlsHostname string) *APIUpdateLogHTTPSRequest {
	r.tlsHostname = &tlsHostname
	return r
}

// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; (10k).
func (r *APIUpdateLogHTTPSRequest) RequestMaxEntries(requestMaxEntries int32) *APIUpdateLogHTTPSRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}

// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; (100MB).
func (r *APIUpdateLogHTTPSRequest) RequestMaxBytes(requestMaxBytes int32) *APIUpdateLogHTTPSRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}

// URL The URL to send logs to. Must use HTTPS. Required.
func (r *APIUpdateLogHTTPSRequest) URL(url string) *APIUpdateLogHTTPSRequest {
	r.url = &url
	return r
}

// ContentType Content type of the header sent with the request.
func (r *APIUpdateLogHTTPSRequest) ContentType(contentType string) *APIUpdateLogHTTPSRequest {
	r.contentType = &contentType
	return r
}

// HeaderName Name of the custom header sent with the request.
func (r *APIUpdateLogHTTPSRequest) HeaderName(headerName string) *APIUpdateLogHTTPSRequest {
	r.headerName = &headerName
	return r
}

// MessageType returns a pointer to a request.
func (r *APIUpdateLogHTTPSRequest) MessageType(messageType LoggingMessageType) *APIUpdateLogHTTPSRequest {
	r.messageType = &messageType
	return r
}

// HeaderValue Value of the custom header sent with the request.
func (r *APIUpdateLogHTTPSRequest) HeaderValue(headerValue string) *APIUpdateLogHTTPSRequest {
	r.headerValue = &headerValue
	return r
}

// Method HTTP method used for request.
func (r *APIUpdateLogHTTPSRequest) Method(method string) *APIUpdateLogHTTPSRequest {
	r.method = &method
	return r
}

// JSONFormat Enforces valid JSON formatting for log entries.
func (r *APIUpdateLogHTTPSRequest) JSONFormat(jsonFormat string) *APIUpdateLogHTTPSRequest {
	r.jsonFormat = &jsonFormat
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogHTTPSRequest) Execute() (*LoggingHTTPSResponse, *http.Response, error) {
	return r.APIService.UpdateLogHTTPSExecute(r)
}

/*
UpdateLogHTTPS Update an HTTPS log endpoint

Update the HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingHTTPSName The name for the real-time logging configuration.
 @return APIUpdateLogHTTPSRequest
*/
func (a *LoggingHTTPSAPIService) UpdateLogHTTPS(ctx context.Context, serviceID string, versionID int32, loggingHTTPSName string) APIUpdateLogHTTPSRequest {
	return APIUpdateLogHTTPSRequest{
		APIService:       a,
		ctx:              ctx,
		serviceID:        serviceID,
		versionID:        versionID,
		loggingHTTPSName: loggingHTTPSName,
	}
}

// UpdateLogHTTPSExecute executes the request
//  @return LoggingHTTPSResponse
func (a *LoggingHTTPSAPIService) UpdateLogHTTPSExecute(r APIUpdateLogHTTPSRequest) (*LoggingHTTPSResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingHTTPSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHTTPSAPIService.UpdateLogHTTPS")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_https_name"+"}", gourl.PathEscape(parameterToString(r.loggingHTTPSName, "")))

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
	if r.requestMaxEntries != nil {
		localVarFormParams.Add("request_max_entries", parameterToString(*r.requestMaxEntries, ""))
	}
	if r.requestMaxBytes != nil {
		localVarFormParams.Add("request_max_bytes", parameterToString(*r.requestMaxBytes, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
	}
	if r.contentType != nil {
		localVarFormParams.Add("content_type", parameterToString(*r.contentType, ""))
	}
	if r.headerName != nil {
		localVarFormParams.Add("header_name", parameterToString(*r.headerName, ""))
	}
	if r.messageType != nil {
		localVarFormParams.Add("message_type", parameterToString(*r.messageType, ""))
	}
	if r.headerValue != nil {
		localVarFormParams.Add("header_value", parameterToString(*r.headerValue, ""))
	}
	if r.method != nil {
		localVarFormParams.Add("method", parameterToString(*r.method, ""))
	}
	if r.jsonFormat != nil {
		localVarFormParams.Add("json_format", parameterToString(*r.jsonFormat, ""))
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
