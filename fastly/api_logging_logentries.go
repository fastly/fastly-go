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

// LoggingLogentriesAPI defines an interface for interacting with the resource.
type LoggingLogentriesAPI interface {

	/*
	CreateLogLogentries Create a Logentries log endpoint

	Create a Logentry for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogLogentriesRequest
	*/
	CreateLogLogentries(ctx context.Context, serviceID string, versionID int32) APICreateLogLogentriesRequest

	// CreateLogLogentriesExecute executes the request
	//  @return LoggingLogentriesResponse
	CreateLogLogentriesExecute(r APICreateLogLogentriesRequest) (*LoggingLogentriesResponse, *http.Response, error)

	/*
	DeleteLogLogentries Delete a Logentries log endpoint

	Delete the Logentry for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingLogentriesName The name for the real-time logging configuration.
	 @return APIDeleteLogLogentriesRequest
	*/
	DeleteLogLogentries(ctx context.Context, serviceID string, versionID int32, loggingLogentriesName string) APIDeleteLogLogentriesRequest

	// DeleteLogLogentriesExecute executes the request
	//  @return InlineResponse200
	DeleteLogLogentriesExecute(r APIDeleteLogLogentriesRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogLogentries Get a Logentries log endpoint

	Get the Logentry for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingLogentriesName The name for the real-time logging configuration.
	 @return APIGetLogLogentriesRequest
	*/
	GetLogLogentries(ctx context.Context, serviceID string, versionID int32, loggingLogentriesName string) APIGetLogLogentriesRequest

	// GetLogLogentriesExecute executes the request
	//  @return LoggingLogentriesResponse
	GetLogLogentriesExecute(r APIGetLogLogentriesRequest) (*LoggingLogentriesResponse, *http.Response, error)

	/*
	ListLogLogentries List Logentries log endpoints

	List all of the Logentries for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogLogentriesRequest
	*/
	ListLogLogentries(ctx context.Context, serviceID string, versionID int32) APIListLogLogentriesRequest

	// ListLogLogentriesExecute executes the request
	//  @return []LoggingLogentriesResponse
	ListLogLogentriesExecute(r APIListLogLogentriesRequest) ([]LoggingLogentriesResponse, *http.Response, error)

	/*
	UpdateLogLogentries Update a Logentries log endpoint

	Update the Logentry for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingLogentriesName The name for the real-time logging configuration.
	 @return APIUpdateLogLogentriesRequest
	*/
	UpdateLogLogentries(ctx context.Context, serviceID string, versionID int32, loggingLogentriesName string) APIUpdateLogLogentriesRequest

	// UpdateLogLogentriesExecute executes the request
	//  @return LoggingLogentriesResponse
	UpdateLogLogentriesExecute(r APIUpdateLogLogentriesRequest) (*LoggingLogentriesResponse, *http.Response, error)
}

// LoggingLogentriesAPIService LoggingLogentriesAPI service
type LoggingLogentriesAPIService service

// APICreateLogLogentriesRequest represents a request for the resource.
type APICreateLogLogentriesRequest struct {
	ctx context.Context
	APIService LoggingLogentriesAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	formatVersion *int32
	responseCondition *string
	format *string
	port *int32
	token *string
	useTLS *LoggingUseTLS
	region *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogLogentriesRequest) Name(name string) *APICreateLogLogentriesRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogLogentriesRequest) Placement(placement string) *APICreateLogLogentriesRequest {
	r.placement = &placement
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogLogentriesRequest) FormatVersion(formatVersion int32) *APICreateLogLogentriesRequest {
	r.formatVersion = &formatVersion
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogLogentriesRequest) ResponseCondition(responseCondition string) *APICreateLogLogentriesRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogLogentriesRequest) Format(format string) *APICreateLogLogentriesRequest {
	r.format = &format
	return r
}
// Port The port number.
func (r *APICreateLogLogentriesRequest) Port(port int32) *APICreateLogLogentriesRequest {
	r.port = &port
	return r
}
// Token Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)).
func (r *APICreateLogLogentriesRequest) Token(token string) *APICreateLogLogentriesRequest {
	r.token = &token
	return r
}
// UseTLS returns a pointer to a request.
func (r *APICreateLogLogentriesRequest) UseTLS(useTLS LoggingUseTLS) *APICreateLogLogentriesRequest {
	r.useTLS = &useTLS
	return r
}
// Region The region to which to stream logs.
func (r *APICreateLogLogentriesRequest) Region(region string) *APICreateLogLogentriesRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogLogentriesRequest) Execute() (*LoggingLogentriesResponse, *http.Response, error) {
	return r.APIService.CreateLogLogentriesExecute(r)
}

/*
CreateLogLogentries Create a Logentries log endpoint

Create a Logentry for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogLogentriesRequest
*/
func (a *LoggingLogentriesAPIService) CreateLogLogentries(ctx context.Context, serviceID string, versionID int32) APICreateLogLogentriesRequest {
	return APICreateLogLogentriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogLogentriesExecute executes the request
//  @return LoggingLogentriesResponse
func (a *LoggingLogentriesAPIService) CreateLogLogentriesExecute(r APICreateLogLogentriesRequest) (*LoggingLogentriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingLogentriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogentriesAPIService.CreateLogLogentries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logentries"
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
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
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

// APIDeleteLogLogentriesRequest represents a request for the resource.
type APIDeleteLogLogentriesRequest struct {
	ctx context.Context
	APIService LoggingLogentriesAPI
	serviceID string
	versionID int32
	loggingLogentriesName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogLogentriesRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogLogentriesExecute(r)
}

/*
DeleteLogLogentries Delete a Logentries log endpoint

Delete the Logentry for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingLogentriesName The name for the real-time logging configuration.
 @return APIDeleteLogLogentriesRequest
*/
func (a *LoggingLogentriesAPIService) DeleteLogLogentries(ctx context.Context, serviceID string, versionID int32, loggingLogentriesName string) APIDeleteLogLogentriesRequest {
	return APIDeleteLogLogentriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingLogentriesName: loggingLogentriesName,
	}
}

// DeleteLogLogentriesExecute executes the request
//  @return InlineResponse200
func (a *LoggingLogentriesAPIService) DeleteLogLogentriesExecute(r APIDeleteLogLogentriesRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogentriesAPIService.DeleteLogLogentries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_logentries_name"+"}", gourl.PathEscape(parameterToString(r.loggingLogentriesName, "")))

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

// APIGetLogLogentriesRequest represents a request for the resource.
type APIGetLogLogentriesRequest struct {
	ctx context.Context
	APIService LoggingLogentriesAPI
	serviceID string
	versionID int32
	loggingLogentriesName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogLogentriesRequest) Execute() (*LoggingLogentriesResponse, *http.Response, error) {
	return r.APIService.GetLogLogentriesExecute(r)
}

/*
GetLogLogentries Get a Logentries log endpoint

Get the Logentry for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingLogentriesName The name for the real-time logging configuration.
 @return APIGetLogLogentriesRequest
*/
func (a *LoggingLogentriesAPIService) GetLogLogentries(ctx context.Context, serviceID string, versionID int32, loggingLogentriesName string) APIGetLogLogentriesRequest {
	return APIGetLogLogentriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingLogentriesName: loggingLogentriesName,
	}
}

// GetLogLogentriesExecute executes the request
//  @return LoggingLogentriesResponse
func (a *LoggingLogentriesAPIService) GetLogLogentriesExecute(r APIGetLogLogentriesRequest) (*LoggingLogentriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingLogentriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogentriesAPIService.GetLogLogentries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_logentries_name"+"}", gourl.PathEscape(parameterToString(r.loggingLogentriesName, "")))

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

// APIListLogLogentriesRequest represents a request for the resource.
type APIListLogLogentriesRequest struct {
	ctx context.Context
	APIService LoggingLogentriesAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogLogentriesRequest) Execute() ([]LoggingLogentriesResponse, *http.Response, error) {
	return r.APIService.ListLogLogentriesExecute(r)
}

/*
ListLogLogentries List Logentries log endpoints

List all of the Logentries for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogLogentriesRequest
*/
func (a *LoggingLogentriesAPIService) ListLogLogentries(ctx context.Context, serviceID string, versionID int32) APIListLogLogentriesRequest {
	return APIListLogLogentriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogLogentriesExecute executes the request
//  @return []LoggingLogentriesResponse
func (a *LoggingLogentriesAPIService) ListLogLogentriesExecute(r APIListLogLogentriesRequest) ([]LoggingLogentriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingLogentriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogentriesAPIService.ListLogLogentries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logentries"
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

// APIUpdateLogLogentriesRequest represents a request for the resource.
type APIUpdateLogLogentriesRequest struct {
	ctx context.Context
	APIService LoggingLogentriesAPI
	serviceID string
	versionID int32
	loggingLogentriesName string
	name *string
	placement *string
	formatVersion *int32
	responseCondition *string
	format *string
	port *int32
	token *string
	useTLS *LoggingUseTLS
	region *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogLogentriesRequest) Name(name string) *APIUpdateLogLogentriesRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogLogentriesRequest) Placement(placement string) *APIUpdateLogLogentriesRequest {
	r.placement = &placement
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogLogentriesRequest) FormatVersion(formatVersion int32) *APIUpdateLogLogentriesRequest {
	r.formatVersion = &formatVersion
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogLogentriesRequest) ResponseCondition(responseCondition string) *APIUpdateLogLogentriesRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogLogentriesRequest) Format(format string) *APIUpdateLogLogentriesRequest {
	r.format = &format
	return r
}
// Port The port number.
func (r *APIUpdateLogLogentriesRequest) Port(port int32) *APIUpdateLogLogentriesRequest {
	r.port = &port
	return r
}
// Token Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)).
func (r *APIUpdateLogLogentriesRequest) Token(token string) *APIUpdateLogLogentriesRequest {
	r.token = &token
	return r
}
// UseTLS returns a pointer to a request.
func (r *APIUpdateLogLogentriesRequest) UseTLS(useTLS LoggingUseTLS) *APIUpdateLogLogentriesRequest {
	r.useTLS = &useTLS
	return r
}
// Region The region to which to stream logs.
func (r *APIUpdateLogLogentriesRequest) Region(region string) *APIUpdateLogLogentriesRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogLogentriesRequest) Execute() (*LoggingLogentriesResponse, *http.Response, error) {
	return r.APIService.UpdateLogLogentriesExecute(r)
}

/*
UpdateLogLogentries Update a Logentries log endpoint

Update the Logentry for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingLogentriesName The name for the real-time logging configuration.
 @return APIUpdateLogLogentriesRequest
*/
func (a *LoggingLogentriesAPIService) UpdateLogLogentries(ctx context.Context, serviceID string, versionID int32, loggingLogentriesName string) APIUpdateLogLogentriesRequest {
	return APIUpdateLogLogentriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingLogentriesName: loggingLogentriesName,
	}
}

// UpdateLogLogentriesExecute executes the request
//  @return LoggingLogentriesResponse
func (a *LoggingLogentriesAPIService) UpdateLogLogentriesExecute(r APIUpdateLogLogentriesRequest) (*LoggingLogentriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingLogentriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogentriesAPIService.UpdateLogLogentries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_logentries_name"+"}", gourl.PathEscape(parameterToString(r.loggingLogentriesName, "")))

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
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
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
