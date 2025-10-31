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

// LoggingHttpsAPI defines an interface for interacting with the resource.
type LoggingHttpsAPI interface {

	/*
		CreateLogHttps Create an HTTPS log endpoint

		Create an HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APICreateLogHttpsRequest
	*/
	CreateLogHttps(ctx context.Context, serviceId string, versionId int32) APICreateLogHttpsRequest

	// CreateLogHttpsExecute executes the request
	//  @return LoggingHttpsResponse
	CreateLogHttpsExecute(r APICreateLogHttpsRequest) (*LoggingHttpsResponse, *http.Response, error)

	/*
		DeleteLogHttps Delete an HTTPS log endpoint

		Delete the HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingHttpsName The name for the real-time logging configuration.
		 @return APIDeleteLogHttpsRequest
	*/
	DeleteLogHttps(ctx context.Context, serviceId string, versionId int32, loggingHttpsName string) APIDeleteLogHttpsRequest

	// DeleteLogHttpsExecute executes the request
	//  @return InlineResponse200
	DeleteLogHttpsExecute(r APIDeleteLogHttpsRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogHttps Get an HTTPS log endpoint

		Get the HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingHttpsName The name for the real-time logging configuration.
		 @return APIGetLogHttpsRequest
	*/
	GetLogHttps(ctx context.Context, serviceId string, versionId int32, loggingHttpsName string) APIGetLogHttpsRequest

	// GetLogHttpsExecute executes the request
	//  @return LoggingHttpsResponse
	GetLogHttpsExecute(r APIGetLogHttpsRequest) (*LoggingHttpsResponse, *http.Response, error)

	/*
		ListLogHttps List HTTPS log endpoints

		List all of the HTTPS objects for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APIListLogHttpsRequest
	*/
	ListLogHttps(ctx context.Context, serviceId string, versionId int32) APIListLogHttpsRequest

	// ListLogHttpsExecute executes the request
	//  @return []LoggingHttpsResponse
	ListLogHttpsExecute(r APIListLogHttpsRequest) ([]LoggingHttpsResponse, *http.Response, error)

	/*
		UpdateLogHttps Update an HTTPS log endpoint

		Update the HTTPS object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingHttpsName The name for the real-time logging configuration.
		 @return APIUpdateLogHttpsRequest
	*/
	UpdateLogHttps(ctx context.Context, serviceId string, versionId int32, loggingHttpsName string) APIUpdateLogHttpsRequest

	// UpdateLogHttpsExecute executes the request
	//  @return LoggingHttpsResponse
	UpdateLogHttpsExecute(r APIUpdateLogHttpsRequest) (*LoggingHttpsResponse, *http.Response, error)
}

// LoggingHttpsAPIService LoggingHttpsAPI service
type LoggingHttpsAPIService service

// APICreateLogHttpsRequest represents a request for the resource.
type APICreateLogHttpsRequest struct {
	ctx                 context.Context
	APIService          LoggingHttpsAPI
	serviceId           string
	versionId           int32
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	tlsCaCert           *string
	tlsClientCert       *string
	tlsClientKey        *string
	tlsHostname         *string
	requestMaxEntries   *int32
	requestMaxBytes     *int32
	url                 *string
	contentType         *string
	headerName          *string
	messageType         *LoggingMessageType
	headerValue         *string
	method              *string
	jsonFormat          *string
	period              *int32
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogHttpsRequest) Name(name string) *APICreateLogHttpsRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogHttpsRequest) Placement(placement string) *APICreateLogHttpsRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogHttpsRequest) ResponseCondition(responseCondition string) *APICreateLogHttpsRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogHttpsRequest) Format(format string) *APICreateLogHttpsRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogHttpsRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogHttpsRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogHttpsRequest) FormatVersion(formatVersion int32) *APICreateLogHttpsRequest {
	r.formatVersion = &formatVersion
	return r
}

// TlsCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateLogHttpsRequest) TlsCaCert(tlsCaCert string) *APICreateLogHttpsRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}

// TlsClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogHttpsRequest) TlsClientCert(tlsClientCert string) *APICreateLogHttpsRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}

// TlsClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogHttpsRequest) TlsClientKey(tlsClientKey string) *APICreateLogHttpsRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}

// TlsHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APICreateLogHttpsRequest) TlsHostname(tlsHostname string) *APICreateLogHttpsRequest {
	r.tlsHostname = &tlsHostname
	return r
}

// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; (10k).
func (r *APICreateLogHttpsRequest) RequestMaxEntries(requestMaxEntries int32) *APICreateLogHttpsRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}

// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; (100MB).
func (r *APICreateLogHttpsRequest) RequestMaxBytes(requestMaxBytes int32) *APICreateLogHttpsRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}

// Url The URL to send logs to. Must use HTTPS. Required.
func (r *APICreateLogHttpsRequest) Url(url string) *APICreateLogHttpsRequest {
	r.url = &url
	return r
}

// ContentType Content type of the header sent with the request.
func (r *APICreateLogHttpsRequest) ContentType(contentType string) *APICreateLogHttpsRequest {
	r.contentType = &contentType
	return r
}

// HeaderName Name of the custom header sent with the request.
func (r *APICreateLogHttpsRequest) HeaderName(headerName string) *APICreateLogHttpsRequest {
	r.headerName = &headerName
	return r
}

// MessageType returns a pointer to a request.
func (r *APICreateLogHttpsRequest) MessageType(messageType LoggingMessageType) *APICreateLogHttpsRequest {
	r.messageType = &messageType
	return r
}

// HeaderValue Value of the custom header sent with the request.
func (r *APICreateLogHttpsRequest) HeaderValue(headerValue string) *APICreateLogHttpsRequest {
	r.headerValue = &headerValue
	return r
}

// Method HTTP method used for request.
func (r *APICreateLogHttpsRequest) Method(method string) *APICreateLogHttpsRequest {
	r.method = &method
	return r
}

// JsonFormat Enforces valid JSON formatting for log entries.
func (r *APICreateLogHttpsRequest) JsonFormat(jsonFormat string) *APICreateLogHttpsRequest {
	r.jsonFormat = &jsonFormat
	return r
}

// Period How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of &#x60;0&#x60; sends logs at the same interval as the default, which is &#x60;5&#x60; seconds.
func (r *APICreateLogHttpsRequest) Period(period int32) *APICreateLogHttpsRequest {
	r.period = &period
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogHttpsRequest) Execute() (*LoggingHttpsResponse, *http.Response, error) {
	return r.APIService.CreateLogHttpsExecute(r)
}

/*
CreateLogHttps Create an HTTPS log endpoint

Create an HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APICreateLogHttpsRequest
*/
func (a *LoggingHttpsAPIService) CreateLogHttps(ctx context.Context, serviceId string, versionId int32) APICreateLogHttpsRequest {
	return APICreateLogHttpsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// CreateLogHttpsExecute executes the request
//  @return LoggingHttpsResponse
func (a *LoggingHttpsAPIService) CreateLogHttpsExecute(r APICreateLogHttpsRequest) (*LoggingHttpsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingHttpsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHttpsAPIService.CreateLogHttps")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

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
	if r.logProcessingRegion != nil {
		localVarFormParams.Add("log_processing_region", parameterToString(*r.logProcessingRegion, ""))
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
	if r.period != nil {
		localVarFormParams.Add("period", parameterToString(*r.period, ""))
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

// APIDeleteLogHttpsRequest represents a request for the resource.
type APIDeleteLogHttpsRequest struct {
	ctx              context.Context
	APIService       LoggingHttpsAPI
	serviceId        string
	versionId        int32
	loggingHttpsName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogHttpsRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogHttpsExecute(r)
}

/*
DeleteLogHttps Delete an HTTPS log endpoint

Delete the HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingHttpsName The name for the real-time logging configuration.
 @return APIDeleteLogHttpsRequest
*/
func (a *LoggingHttpsAPIService) DeleteLogHttps(ctx context.Context, serviceId string, versionId int32, loggingHttpsName string) APIDeleteLogHttpsRequest {
	return APIDeleteLogHttpsRequest{
		APIService:       a,
		ctx:              ctx,
		serviceId:        serviceId,
		versionId:        versionId,
		loggingHttpsName: loggingHttpsName,
	}
}

// DeleteLogHttpsExecute executes the request
//  @return InlineResponse200
func (a *LoggingHttpsAPIService) DeleteLogHttpsExecute(r APIDeleteLogHttpsRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHttpsAPIService.DeleteLogHttps")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_https_name"+"}", gourl.PathEscape(parameterToString(r.loggingHttpsName, "")))

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

// APIGetLogHttpsRequest represents a request for the resource.
type APIGetLogHttpsRequest struct {
	ctx              context.Context
	APIService       LoggingHttpsAPI
	serviceId        string
	versionId        int32
	loggingHttpsName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogHttpsRequest) Execute() (*LoggingHttpsResponse, *http.Response, error) {
	return r.APIService.GetLogHttpsExecute(r)
}

/*
GetLogHttps Get an HTTPS log endpoint

Get the HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingHttpsName The name for the real-time logging configuration.
 @return APIGetLogHttpsRequest
*/
func (a *LoggingHttpsAPIService) GetLogHttps(ctx context.Context, serviceId string, versionId int32, loggingHttpsName string) APIGetLogHttpsRequest {
	return APIGetLogHttpsRequest{
		APIService:       a,
		ctx:              ctx,
		serviceId:        serviceId,
		versionId:        versionId,
		loggingHttpsName: loggingHttpsName,
	}
}

// GetLogHttpsExecute executes the request
//  @return LoggingHttpsResponse
func (a *LoggingHttpsAPIService) GetLogHttpsExecute(r APIGetLogHttpsRequest) (*LoggingHttpsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingHttpsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHttpsAPIService.GetLogHttps")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_https_name"+"}", gourl.PathEscape(parameterToString(r.loggingHttpsName, "")))

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

// APIListLogHttpsRequest represents a request for the resource.
type APIListLogHttpsRequest struct {
	ctx        context.Context
	APIService LoggingHttpsAPI
	serviceId  string
	versionId  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogHttpsRequest) Execute() ([]LoggingHttpsResponse, *http.Response, error) {
	return r.APIService.ListLogHttpsExecute(r)
}

/*
ListLogHttps List HTTPS log endpoints

List all of the HTTPS objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APIListLogHttpsRequest
*/
func (a *LoggingHttpsAPIService) ListLogHttps(ctx context.Context, serviceId string, versionId int32) APIListLogHttpsRequest {
	return APIListLogHttpsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// ListLogHttpsExecute executes the request
//  @return []LoggingHttpsResponse
func (a *LoggingHttpsAPIService) ListLogHttpsExecute(r APIListLogHttpsRequest) ([]LoggingHttpsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingHttpsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHttpsAPIService.ListLogHttps")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

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

// APIUpdateLogHttpsRequest represents a request for the resource.
type APIUpdateLogHttpsRequest struct {
	ctx                 context.Context
	APIService          LoggingHttpsAPI
	serviceId           string
	versionId           int32
	loggingHttpsName    string
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	tlsCaCert           *string
	tlsClientCert       *string
	tlsClientKey        *string
	tlsHostname         *string
	requestMaxEntries   *int32
	requestMaxBytes     *int32
	url                 *string
	contentType         *string
	headerName          *string
	messageType         *LoggingMessageType
	headerValue         *string
	method              *string
	jsonFormat          *string
	period              *int32
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogHttpsRequest) Name(name string) *APIUpdateLogHttpsRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogHttpsRequest) Placement(placement string) *APIUpdateLogHttpsRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogHttpsRequest) ResponseCondition(responseCondition string) *APIUpdateLogHttpsRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APIUpdateLogHttpsRequest) Format(format string) *APIUpdateLogHttpsRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogHttpsRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogHttpsRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogHttpsRequest) FormatVersion(formatVersion int32) *APIUpdateLogHttpsRequest {
	r.formatVersion = &formatVersion
	return r
}

// TlsCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APIUpdateLogHttpsRequest) TlsCaCert(tlsCaCert string) *APIUpdateLogHttpsRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}

// TlsClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogHttpsRequest) TlsClientCert(tlsClientCert string) *APIUpdateLogHttpsRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}

// TlsClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogHttpsRequest) TlsClientKey(tlsClientKey string) *APIUpdateLogHttpsRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}

// TlsHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APIUpdateLogHttpsRequest) TlsHostname(tlsHostname string) *APIUpdateLogHttpsRequest {
	r.tlsHostname = &tlsHostname
	return r
}

// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; (10k).
func (r *APIUpdateLogHttpsRequest) RequestMaxEntries(requestMaxEntries int32) *APIUpdateLogHttpsRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}

// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; (100MB).
func (r *APIUpdateLogHttpsRequest) RequestMaxBytes(requestMaxBytes int32) *APIUpdateLogHttpsRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}

// Url The URL to send logs to. Must use HTTPS. Required.
func (r *APIUpdateLogHttpsRequest) Url(url string) *APIUpdateLogHttpsRequest {
	r.url = &url
	return r
}

// ContentType Content type of the header sent with the request.
func (r *APIUpdateLogHttpsRequest) ContentType(contentType string) *APIUpdateLogHttpsRequest {
	r.contentType = &contentType
	return r
}

// HeaderName Name of the custom header sent with the request.
func (r *APIUpdateLogHttpsRequest) HeaderName(headerName string) *APIUpdateLogHttpsRequest {
	r.headerName = &headerName
	return r
}

// MessageType returns a pointer to a request.
func (r *APIUpdateLogHttpsRequest) MessageType(messageType LoggingMessageType) *APIUpdateLogHttpsRequest {
	r.messageType = &messageType
	return r
}

// HeaderValue Value of the custom header sent with the request.
func (r *APIUpdateLogHttpsRequest) HeaderValue(headerValue string) *APIUpdateLogHttpsRequest {
	r.headerValue = &headerValue
	return r
}

// Method HTTP method used for request.
func (r *APIUpdateLogHttpsRequest) Method(method string) *APIUpdateLogHttpsRequest {
	r.method = &method
	return r
}

// JsonFormat Enforces valid JSON formatting for log entries.
func (r *APIUpdateLogHttpsRequest) JsonFormat(jsonFormat string) *APIUpdateLogHttpsRequest {
	r.jsonFormat = &jsonFormat
	return r
}

// Period How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of &#x60;0&#x60; sends logs at the same interval as the default, which is &#x60;5&#x60; seconds.
func (r *APIUpdateLogHttpsRequest) Period(period int32) *APIUpdateLogHttpsRequest {
	r.period = &period
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogHttpsRequest) Execute() (*LoggingHttpsResponse, *http.Response, error) {
	return r.APIService.UpdateLogHttpsExecute(r)
}

/*
UpdateLogHttps Update an HTTPS log endpoint

Update the HTTPS object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingHttpsName The name for the real-time logging configuration.
 @return APIUpdateLogHttpsRequest
*/
func (a *LoggingHttpsAPIService) UpdateLogHttps(ctx context.Context, serviceId string, versionId int32, loggingHttpsName string) APIUpdateLogHttpsRequest {
	return APIUpdateLogHttpsRequest{
		APIService:       a,
		ctx:              ctx,
		serviceId:        serviceId,
		versionId:        versionId,
		loggingHttpsName: loggingHttpsName,
	}
}

// UpdateLogHttpsExecute executes the request
//  @return LoggingHttpsResponse
func (a *LoggingHttpsAPIService) UpdateLogHttpsExecute(r APIUpdateLogHttpsRequest) (*LoggingHttpsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingHttpsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHttpsAPIService.UpdateLogHttps")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_https_name"+"}", gourl.PathEscape(parameterToString(r.loggingHttpsName, "")))

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
	if r.logProcessingRegion != nil {
		localVarFormParams.Add("log_processing_region", parameterToString(*r.logProcessingRegion, ""))
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
	if r.period != nil {
		localVarFormParams.Add("period", parameterToString(*r.period, ""))
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
