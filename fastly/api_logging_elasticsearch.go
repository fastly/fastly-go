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

// LoggingElasticsearchAPI defines an interface for interacting with the resource.
type LoggingElasticsearchAPI interface {

	/*
	CreateLogElasticsearch Create an Elasticsearch log endpoint

	Create a Elasticsearch logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogElasticsearchRequest
	*/
	CreateLogElasticsearch(ctx context.Context, serviceID string, versionID int32) APICreateLogElasticsearchRequest

	// CreateLogElasticsearchExecute executes the request
	//  @return LoggingElasticsearchResponse
	CreateLogElasticsearchExecute(r APICreateLogElasticsearchRequest) (*LoggingElasticsearchResponse, *http.Response, error)

	/*
	DeleteLogElasticsearch Delete an Elasticsearch log endpoint

	Delete the Elasticsearch logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingElasticsearchName The name for the real-time logging configuration.
	 @return APIDeleteLogElasticsearchRequest
	*/
	DeleteLogElasticsearch(ctx context.Context, serviceID string, versionID int32, loggingElasticsearchName string) APIDeleteLogElasticsearchRequest

	// DeleteLogElasticsearchExecute executes the request
	//  @return InlineResponse200
	DeleteLogElasticsearchExecute(r APIDeleteLogElasticsearchRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogElasticsearch Get an Elasticsearch log endpoint

	Get the Elasticsearch logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingElasticsearchName The name for the real-time logging configuration.
	 @return APIGetLogElasticsearchRequest
	*/
	GetLogElasticsearch(ctx context.Context, serviceID string, versionID int32, loggingElasticsearchName string) APIGetLogElasticsearchRequest

	// GetLogElasticsearchExecute executes the request
	//  @return LoggingElasticsearchResponse
	GetLogElasticsearchExecute(r APIGetLogElasticsearchRequest) (*LoggingElasticsearchResponse, *http.Response, error)

	/*
	ListLogElasticsearch List Elasticsearch log endpoints

	List all of the Elasticsearch logging endpoints for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogElasticsearchRequest
	*/
	ListLogElasticsearch(ctx context.Context, serviceID string, versionID int32) APIListLogElasticsearchRequest

	// ListLogElasticsearchExecute executes the request
	//  @return []LoggingElasticsearchResponse
	ListLogElasticsearchExecute(r APIListLogElasticsearchRequest) ([]LoggingElasticsearchResponse, *http.Response, error)

	/*
	UpdateLogElasticsearch Update an Elasticsearch log endpoint

	Update the Elasticsearch logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingElasticsearchName The name for the real-time logging configuration.
	 @return APIUpdateLogElasticsearchRequest
	*/
	UpdateLogElasticsearch(ctx context.Context, serviceID string, versionID int32, loggingElasticsearchName string) APIUpdateLogElasticsearchRequest

	// UpdateLogElasticsearchExecute executes the request
	//  @return LoggingElasticsearchResponse
	UpdateLogElasticsearchExecute(r APIUpdateLogElasticsearchRequest) (*LoggingElasticsearchResponse, *http.Response, error)
}

// LoggingElasticsearchAPIService LoggingElasticsearchAPI service
type LoggingElasticsearchAPIService service

// APICreateLogElasticsearchRequest represents a request for the resource.
type APICreateLogElasticsearchRequest struct {
	ctx context.Context
	APIService LoggingElasticsearchAPI
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
	requestMaxEntries *int32
	requestMaxBytes *int32
	index *string
	url *string
	pipeline *string
	user *string
	password *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogElasticsearchRequest) Name(name string) *APICreateLogElasticsearchRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogElasticsearchRequest) Placement(placement string) *APICreateLogElasticsearchRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogElasticsearchRequest) ResponseCondition(responseCondition string) *APICreateLogElasticsearchRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Elasticsearch can ingest.
func (r *APICreateLogElasticsearchRequest) Format(format string) *APICreateLogElasticsearchRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogElasticsearchRequest) FormatVersion(formatVersion int32) *APICreateLogElasticsearchRequest {
	r.formatVersion = &formatVersion
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateLogElasticsearchRequest) TLSCaCert(tlsCaCert string) *APICreateLogElasticsearchRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogElasticsearchRequest) TLSClientCert(tlsClientCert string) *APICreateLogElasticsearchRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogElasticsearchRequest) TLSClientKey(tlsClientKey string) *APICreateLogElasticsearchRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APICreateLogElasticsearchRequest) TLSHostname(tlsHostname string) *APICreateLogElasticsearchRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APICreateLogElasticsearchRequest) RequestMaxEntries(requestMaxEntries int32) *APICreateLogElasticsearchRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}
// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APICreateLogElasticsearchRequest) RequestMaxBytes(requestMaxBytes int32) *APICreateLogElasticsearchRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}
// Index The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, &#x60;#{%F}&#x60; will interpolate as &#x60;YYYY-MM-DD&#x60; with today&#39;s date.
func (r *APICreateLogElasticsearchRequest) Index(index string) *APICreateLogElasticsearchRequest {
	r.index = &index
	return r
}
// URL The URL to stream logs to. Must use HTTPS.
func (r *APICreateLogElasticsearchRequest) URL(url string) *APICreateLogElasticsearchRequest {
	r.url = &url
	return r
}
// Pipeline The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html).
func (r *APICreateLogElasticsearchRequest) Pipeline(pipeline string) *APICreateLogElasticsearchRequest {
	r.pipeline = &pipeline
	return r
}
// User Basic Auth username.
func (r *APICreateLogElasticsearchRequest) User(user string) *APICreateLogElasticsearchRequest {
	r.user = &user
	return r
}
// Password Basic Auth password.
func (r *APICreateLogElasticsearchRequest) Password(password string) *APICreateLogElasticsearchRequest {
	r.password = &password
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogElasticsearchRequest) Execute() (*LoggingElasticsearchResponse, *http.Response, error) {
	return r.APIService.CreateLogElasticsearchExecute(r)
}

/*
CreateLogElasticsearch Create an Elasticsearch log endpoint

Create a Elasticsearch logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogElasticsearchRequest
*/
func (a *LoggingElasticsearchAPIService) CreateLogElasticsearch(ctx context.Context, serviceID string, versionID int32) APICreateLogElasticsearchRequest {
	return APICreateLogElasticsearchRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogElasticsearchExecute executes the request
//  @return LoggingElasticsearchResponse
func (a *LoggingElasticsearchAPIService) CreateLogElasticsearchExecute(r APICreateLogElasticsearchRequest) (*LoggingElasticsearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingElasticsearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingElasticsearchAPIService.CreateLogElasticsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/elasticsearch"
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
	if r.index != nil {
		localVarFormParams.Add("index", parameterToString(*r.index, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
	}
	if r.pipeline != nil {
		localVarFormParams.Add("pipeline", parameterToString(*r.pipeline, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
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

// APIDeleteLogElasticsearchRequest represents a request for the resource.
type APIDeleteLogElasticsearchRequest struct {
	ctx context.Context
	APIService LoggingElasticsearchAPI
	serviceID string
	versionID int32
	loggingElasticsearchName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogElasticsearchRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogElasticsearchExecute(r)
}

/*
DeleteLogElasticsearch Delete an Elasticsearch log endpoint

Delete the Elasticsearch logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingElasticsearchName The name for the real-time logging configuration.
 @return APIDeleteLogElasticsearchRequest
*/
func (a *LoggingElasticsearchAPIService) DeleteLogElasticsearch(ctx context.Context, serviceID string, versionID int32, loggingElasticsearchName string) APIDeleteLogElasticsearchRequest {
	return APIDeleteLogElasticsearchRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingElasticsearchName: loggingElasticsearchName,
	}
}

// DeleteLogElasticsearchExecute executes the request
//  @return InlineResponse200
func (a *LoggingElasticsearchAPIService) DeleteLogElasticsearchExecute(r APIDeleteLogElasticsearchRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingElasticsearchAPIService.DeleteLogElasticsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_elasticsearch_name"+"}", gourl.PathEscape(parameterToString(r.loggingElasticsearchName, "")))

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

// APIGetLogElasticsearchRequest represents a request for the resource.
type APIGetLogElasticsearchRequest struct {
	ctx context.Context
	APIService LoggingElasticsearchAPI
	serviceID string
	versionID int32
	loggingElasticsearchName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogElasticsearchRequest) Execute() (*LoggingElasticsearchResponse, *http.Response, error) {
	return r.APIService.GetLogElasticsearchExecute(r)
}

/*
GetLogElasticsearch Get an Elasticsearch log endpoint

Get the Elasticsearch logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingElasticsearchName The name for the real-time logging configuration.
 @return APIGetLogElasticsearchRequest
*/
func (a *LoggingElasticsearchAPIService) GetLogElasticsearch(ctx context.Context, serviceID string, versionID int32, loggingElasticsearchName string) APIGetLogElasticsearchRequest {
	return APIGetLogElasticsearchRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingElasticsearchName: loggingElasticsearchName,
	}
}

// GetLogElasticsearchExecute executes the request
//  @return LoggingElasticsearchResponse
func (a *LoggingElasticsearchAPIService) GetLogElasticsearchExecute(r APIGetLogElasticsearchRequest) (*LoggingElasticsearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingElasticsearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingElasticsearchAPIService.GetLogElasticsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_elasticsearch_name"+"}", gourl.PathEscape(parameterToString(r.loggingElasticsearchName, "")))

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

// APIListLogElasticsearchRequest represents a request for the resource.
type APIListLogElasticsearchRequest struct {
	ctx context.Context
	APIService LoggingElasticsearchAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogElasticsearchRequest) Execute() ([]LoggingElasticsearchResponse, *http.Response, error) {
	return r.APIService.ListLogElasticsearchExecute(r)
}

/*
ListLogElasticsearch List Elasticsearch log endpoints

List all of the Elasticsearch logging endpoints for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogElasticsearchRequest
*/
func (a *LoggingElasticsearchAPIService) ListLogElasticsearch(ctx context.Context, serviceID string, versionID int32) APIListLogElasticsearchRequest {
	return APIListLogElasticsearchRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogElasticsearchExecute executes the request
//  @return []LoggingElasticsearchResponse
func (a *LoggingElasticsearchAPIService) ListLogElasticsearchExecute(r APIListLogElasticsearchRequest) ([]LoggingElasticsearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingElasticsearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingElasticsearchAPIService.ListLogElasticsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/elasticsearch"
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

// APIUpdateLogElasticsearchRequest represents a request for the resource.
type APIUpdateLogElasticsearchRequest struct {
	ctx context.Context
	APIService LoggingElasticsearchAPI
	serviceID string
	versionID int32
	loggingElasticsearchName string
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	tlsCaCert *string
	tlsClientCert *string
	tlsClientKey *string
	tlsHostname *string
	requestMaxEntries *int32
	requestMaxBytes *int32
	index *string
	url *string
	pipeline *string
	user *string
	password *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogElasticsearchRequest) Name(name string) *APIUpdateLogElasticsearchRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogElasticsearchRequest) Placement(placement string) *APIUpdateLogElasticsearchRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogElasticsearchRequest) ResponseCondition(responseCondition string) *APIUpdateLogElasticsearchRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Elasticsearch can ingest.
func (r *APIUpdateLogElasticsearchRequest) Format(format string) *APIUpdateLogElasticsearchRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogElasticsearchRequest) FormatVersion(formatVersion int32) *APIUpdateLogElasticsearchRequest {
	r.formatVersion = &formatVersion
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APIUpdateLogElasticsearchRequest) TLSCaCert(tlsCaCert string) *APIUpdateLogElasticsearchRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogElasticsearchRequest) TLSClientCert(tlsClientCert string) *APIUpdateLogElasticsearchRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APIUpdateLogElasticsearchRequest) TLSClientKey(tlsClientKey string) *APIUpdateLogElasticsearchRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APIUpdateLogElasticsearchRequest) TLSHostname(tlsHostname string) *APIUpdateLogElasticsearchRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// RequestMaxEntries The maximum number of logs sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APIUpdateLogElasticsearchRequest) RequestMaxEntries(requestMaxEntries int32) *APIUpdateLogElasticsearchRequest {
	r.requestMaxEntries = &requestMaxEntries
	return r
}
// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; for unbounded.
func (r *APIUpdateLogElasticsearchRequest) RequestMaxBytes(requestMaxBytes int32) *APIUpdateLogElasticsearchRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}
// Index The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, &#x60;#{%F}&#x60; will interpolate as &#x60;YYYY-MM-DD&#x60; with today&#39;s date.
func (r *APIUpdateLogElasticsearchRequest) Index(index string) *APIUpdateLogElasticsearchRequest {
	r.index = &index
	return r
}
// URL The URL to stream logs to. Must use HTTPS.
func (r *APIUpdateLogElasticsearchRequest) URL(url string) *APIUpdateLogElasticsearchRequest {
	r.url = &url
	return r
}
// Pipeline The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html).
func (r *APIUpdateLogElasticsearchRequest) Pipeline(pipeline string) *APIUpdateLogElasticsearchRequest {
	r.pipeline = &pipeline
	return r
}
// User Basic Auth username.
func (r *APIUpdateLogElasticsearchRequest) User(user string) *APIUpdateLogElasticsearchRequest {
	r.user = &user
	return r
}
// Password Basic Auth password.
func (r *APIUpdateLogElasticsearchRequest) Password(password string) *APIUpdateLogElasticsearchRequest {
	r.password = &password
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogElasticsearchRequest) Execute() (*LoggingElasticsearchResponse, *http.Response, error) {
	return r.APIService.UpdateLogElasticsearchExecute(r)
}

/*
UpdateLogElasticsearch Update an Elasticsearch log endpoint

Update the Elasticsearch logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingElasticsearchName The name for the real-time logging configuration.
 @return APIUpdateLogElasticsearchRequest
*/
func (a *LoggingElasticsearchAPIService) UpdateLogElasticsearch(ctx context.Context, serviceID string, versionID int32, loggingElasticsearchName string) APIUpdateLogElasticsearchRequest {
	return APIUpdateLogElasticsearchRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingElasticsearchName: loggingElasticsearchName,
	}
}

// UpdateLogElasticsearchExecute executes the request
//  @return LoggingElasticsearchResponse
func (a *LoggingElasticsearchAPIService) UpdateLogElasticsearchExecute(r APIUpdateLogElasticsearchRequest) (*LoggingElasticsearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingElasticsearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingElasticsearchAPIService.UpdateLogElasticsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_elasticsearch_name"+"}", gourl.PathEscape(parameterToString(r.loggingElasticsearchName, "")))

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
	if r.index != nil {
		localVarFormParams.Add("index", parameterToString(*r.index, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
	}
	if r.pipeline != nil {
		localVarFormParams.Add("pipeline", parameterToString(*r.pipeline, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
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
