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

// LoggingHoneycombAPI defines an interface for interacting with the resource.
type LoggingHoneycombAPI interface {

	/*
	CreateLogHoneycomb Create a Honeycomb log endpoint

	Create a Honeycomb logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogHoneycombRequest
	*/
	CreateLogHoneycomb(ctx context.Context, serviceID string, versionID int32) APICreateLogHoneycombRequest

	// CreateLogHoneycombExecute executes the request
	//  @return LoggingHoneycombResponse
	CreateLogHoneycombExecute(r APICreateLogHoneycombRequest) (*LoggingHoneycombResponse, *http.Response, error)

	/*
	DeleteLogHoneycomb Delete the Honeycomb log endpoint

	Delete the Honeycomb logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingHoneycombName The name for the real-time logging configuration.
	 @return APIDeleteLogHoneycombRequest
	*/
	DeleteLogHoneycomb(ctx context.Context, serviceID string, versionID int32, loggingHoneycombName string) APIDeleteLogHoneycombRequest

	// DeleteLogHoneycombExecute executes the request
	//  @return InlineResponse200
	DeleteLogHoneycombExecute(r APIDeleteLogHoneycombRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogHoneycomb Get a Honeycomb log endpoint

	Get the details of a Honeycomb logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingHoneycombName The name for the real-time logging configuration.
	 @return APIGetLogHoneycombRequest
	*/
	GetLogHoneycomb(ctx context.Context, serviceID string, versionID int32, loggingHoneycombName string) APIGetLogHoneycombRequest

	// GetLogHoneycombExecute executes the request
	//  @return LoggingHoneycombResponse
	GetLogHoneycombExecute(r APIGetLogHoneycombRequest) (*LoggingHoneycombResponse, *http.Response, error)

	/*
	ListLogHoneycomb List Honeycomb log endpoints

	List all of the Honeycomb logging objects for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogHoneycombRequest
	*/
	ListLogHoneycomb(ctx context.Context, serviceID string, versionID int32) APIListLogHoneycombRequest

	// ListLogHoneycombExecute executes the request
	//  @return []LoggingHoneycombResponse
	ListLogHoneycombExecute(r APIListLogHoneycombRequest) ([]LoggingHoneycombResponse, *http.Response, error)

	/*
	UpdateLogHoneycomb Update a Honeycomb log endpoint

	Update a Honeycomb logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingHoneycombName The name for the real-time logging configuration.
	 @return APIUpdateLogHoneycombRequest
	*/
	UpdateLogHoneycomb(ctx context.Context, serviceID string, versionID int32, loggingHoneycombName string) APIUpdateLogHoneycombRequest

	// UpdateLogHoneycombExecute executes the request
	//  @return LoggingHoneycombResponse
	UpdateLogHoneycombExecute(r APIUpdateLogHoneycombRequest) (*LoggingHoneycombResponse, *http.Response, error)
}

// LoggingHoneycombAPIService LoggingHoneycombAPI service
type LoggingHoneycombAPIService service

// APICreateLogHoneycombRequest represents a request for the resource.
type APICreateLogHoneycombRequest struct {
	ctx context.Context
	APIService LoggingHoneycombAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	dataset *string
	token *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogHoneycombRequest) Name(name string) *APICreateLogHoneycombRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogHoneycombRequest) Placement(placement string) *APICreateLogHoneycombRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogHoneycombRequest) ResponseCondition(responseCondition string) *APICreateLogHoneycombRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest.
func (r *APICreateLogHoneycombRequest) Format(format string) *APICreateLogHoneycombRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogHoneycombRequest) FormatVersion(formatVersion int32) *APICreateLogHoneycombRequest {
	r.formatVersion = &formatVersion
	return r
}
// Dataset The Honeycomb Dataset you want to log to.
func (r *APICreateLogHoneycombRequest) Dataset(dataset string) *APICreateLogHoneycombRequest {
	r.dataset = &dataset
	return r
}
// Token The Write Key from the Account page of your Honeycomb account.
func (r *APICreateLogHoneycombRequest) Token(token string) *APICreateLogHoneycombRequest {
	r.token = &token
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogHoneycombRequest) Execute() (*LoggingHoneycombResponse, *http.Response, error) {
	return r.APIService.CreateLogHoneycombExecute(r)
}

/*
CreateLogHoneycomb Create a Honeycomb log endpoint

Create a Honeycomb logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogHoneycombRequest
*/
func (a *LoggingHoneycombAPIService) CreateLogHoneycomb(ctx context.Context, serviceID string, versionID int32) APICreateLogHoneycombRequest {
	return APICreateLogHoneycombRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogHoneycombExecute executes the request
//  @return LoggingHoneycombResponse
func (a *LoggingHoneycombAPIService) CreateLogHoneycombExecute(r APICreateLogHoneycombRequest) (*LoggingHoneycombResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingHoneycombResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHoneycombAPIService.CreateLogHoneycomb")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/honeycomb"
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
	if r.dataset != nil {
		localVarFormParams.Add("dataset", parameterToString(*r.dataset, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
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

// APIDeleteLogHoneycombRequest represents a request for the resource.
type APIDeleteLogHoneycombRequest struct {
	ctx context.Context
	APIService LoggingHoneycombAPI
	serviceID string
	versionID int32
	loggingHoneycombName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogHoneycombRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogHoneycombExecute(r)
}

/*
DeleteLogHoneycomb Delete the Honeycomb log endpoint

Delete the Honeycomb logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingHoneycombName The name for the real-time logging configuration.
 @return APIDeleteLogHoneycombRequest
*/
func (a *LoggingHoneycombAPIService) DeleteLogHoneycomb(ctx context.Context, serviceID string, versionID int32, loggingHoneycombName string) APIDeleteLogHoneycombRequest {
	return APIDeleteLogHoneycombRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingHoneycombName: loggingHoneycombName,
	}
}

// DeleteLogHoneycombExecute executes the request
//  @return InlineResponse200
func (a *LoggingHoneycombAPIService) DeleteLogHoneycombExecute(r APIDeleteLogHoneycombRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHoneycombAPIService.DeleteLogHoneycomb")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_honeycomb_name"+"}", gourl.PathEscape(parameterToString(r.loggingHoneycombName, "")))

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

// APIGetLogHoneycombRequest represents a request for the resource.
type APIGetLogHoneycombRequest struct {
	ctx context.Context
	APIService LoggingHoneycombAPI
	serviceID string
	versionID int32
	loggingHoneycombName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogHoneycombRequest) Execute() (*LoggingHoneycombResponse, *http.Response, error) {
	return r.APIService.GetLogHoneycombExecute(r)
}

/*
GetLogHoneycomb Get a Honeycomb log endpoint

Get the details of a Honeycomb logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingHoneycombName The name for the real-time logging configuration.
 @return APIGetLogHoneycombRequest
*/
func (a *LoggingHoneycombAPIService) GetLogHoneycomb(ctx context.Context, serviceID string, versionID int32, loggingHoneycombName string) APIGetLogHoneycombRequest {
	return APIGetLogHoneycombRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingHoneycombName: loggingHoneycombName,
	}
}

// GetLogHoneycombExecute executes the request
//  @return LoggingHoneycombResponse
func (a *LoggingHoneycombAPIService) GetLogHoneycombExecute(r APIGetLogHoneycombRequest) (*LoggingHoneycombResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingHoneycombResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHoneycombAPIService.GetLogHoneycomb")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_honeycomb_name"+"}", gourl.PathEscape(parameterToString(r.loggingHoneycombName, "")))

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

// APIListLogHoneycombRequest represents a request for the resource.
type APIListLogHoneycombRequest struct {
	ctx context.Context
	APIService LoggingHoneycombAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogHoneycombRequest) Execute() ([]LoggingHoneycombResponse, *http.Response, error) {
	return r.APIService.ListLogHoneycombExecute(r)
}

/*
ListLogHoneycomb List Honeycomb log endpoints

List all of the Honeycomb logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogHoneycombRequest
*/
func (a *LoggingHoneycombAPIService) ListLogHoneycomb(ctx context.Context, serviceID string, versionID int32) APIListLogHoneycombRequest {
	return APIListLogHoneycombRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogHoneycombExecute executes the request
//  @return []LoggingHoneycombResponse
func (a *LoggingHoneycombAPIService) ListLogHoneycombExecute(r APIListLogHoneycombRequest) ([]LoggingHoneycombResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingHoneycombResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHoneycombAPIService.ListLogHoneycomb")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/honeycomb"
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

// APIUpdateLogHoneycombRequest represents a request for the resource.
type APIUpdateLogHoneycombRequest struct {
	ctx context.Context
	APIService LoggingHoneycombAPI
	serviceID string
	versionID int32
	loggingHoneycombName string
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	dataset *string
	token *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogHoneycombRequest) Name(name string) *APIUpdateLogHoneycombRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogHoneycombRequest) Placement(placement string) *APIUpdateLogHoneycombRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogHoneycombRequest) ResponseCondition(responseCondition string) *APIUpdateLogHoneycombRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest.
func (r *APIUpdateLogHoneycombRequest) Format(format string) *APIUpdateLogHoneycombRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogHoneycombRequest) FormatVersion(formatVersion int32) *APIUpdateLogHoneycombRequest {
	r.formatVersion = &formatVersion
	return r
}
// Dataset The Honeycomb Dataset you want to log to.
func (r *APIUpdateLogHoneycombRequest) Dataset(dataset string) *APIUpdateLogHoneycombRequest {
	r.dataset = &dataset
	return r
}
// Token The Write Key from the Account page of your Honeycomb account.
func (r *APIUpdateLogHoneycombRequest) Token(token string) *APIUpdateLogHoneycombRequest {
	r.token = &token
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogHoneycombRequest) Execute() (*LoggingHoneycombResponse, *http.Response, error) {
	return r.APIService.UpdateLogHoneycombExecute(r)
}

/*
UpdateLogHoneycomb Update a Honeycomb log endpoint

Update a Honeycomb logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingHoneycombName The name for the real-time logging configuration.
 @return APIUpdateLogHoneycombRequest
*/
func (a *LoggingHoneycombAPIService) UpdateLogHoneycomb(ctx context.Context, serviceID string, versionID int32, loggingHoneycombName string) APIUpdateLogHoneycombRequest {
	return APIUpdateLogHoneycombRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingHoneycombName: loggingHoneycombName,
	}
}

// UpdateLogHoneycombExecute executes the request
//  @return LoggingHoneycombResponse
func (a *LoggingHoneycombAPIService) UpdateLogHoneycombExecute(r APIUpdateLogHoneycombRequest) (*LoggingHoneycombResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingHoneycombResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingHoneycombAPIService.UpdateLogHoneycomb")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_honeycomb_name"+"}", gourl.PathEscape(parameterToString(r.loggingHoneycombName, "")))

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
	if r.dataset != nil {
		localVarFormParams.Add("dataset", parameterToString(*r.dataset, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
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
