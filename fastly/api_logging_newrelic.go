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

// LoggingNewrelicAPI defines an interface for interacting with the resource.
type LoggingNewrelicAPI interface {

	/*
		CreateLogNewrelic Create a New Relic log endpoint

		Create a New Relic Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogNewrelicRequest
	*/
	CreateLogNewrelic(ctx context.Context, serviceID string, versionID int32) APICreateLogNewrelicRequest

	// CreateLogNewrelicExecute executes the request
	//  @return LoggingNewrelicResponse
	CreateLogNewrelicExecute(r APICreateLogNewrelicRequest) (*LoggingNewrelicResponse, *http.Response, error)

	/*
		DeleteLogNewrelic Delete a New Relic log endpoint

		Delete the New Relic Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingNewrelicName The name for the real-time logging configuration.
		 @return APIDeleteLogNewrelicRequest
	*/
	DeleteLogNewrelic(ctx context.Context, serviceID string, versionID int32, loggingNewrelicName string) APIDeleteLogNewrelicRequest

	// DeleteLogNewrelicExecute executes the request
	//  @return InlineResponse200
	DeleteLogNewrelicExecute(r APIDeleteLogNewrelicRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogNewrelic Get a New Relic log endpoint

		Get the details of a New Relic Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingNewrelicName The name for the real-time logging configuration.
		 @return APIGetLogNewrelicRequest
	*/
	GetLogNewrelic(ctx context.Context, serviceID string, versionID int32, loggingNewrelicName string) APIGetLogNewrelicRequest

	// GetLogNewrelicExecute executes the request
	//  @return LoggingNewrelicResponse
	GetLogNewrelicExecute(r APIGetLogNewrelicRequest) (*LoggingNewrelicResponse, *http.Response, error)

	/*
		ListLogNewrelic List New Relic log endpoints

		List all of the New Relic Logs logging objects for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogNewrelicRequest
	*/
	ListLogNewrelic(ctx context.Context, serviceID string, versionID int32) APIListLogNewrelicRequest

	// ListLogNewrelicExecute executes the request
	//  @return []LoggingNewrelicResponse
	ListLogNewrelicExecute(r APIListLogNewrelicRequest) ([]LoggingNewrelicResponse, *http.Response, error)

	/*
		UpdateLogNewrelic Update a New Relic log endpoint

		Update a New Relic Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingNewrelicName The name for the real-time logging configuration.
		 @return APIUpdateLogNewrelicRequest
	*/
	UpdateLogNewrelic(ctx context.Context, serviceID string, versionID int32, loggingNewrelicName string) APIUpdateLogNewrelicRequest

	// UpdateLogNewrelicExecute executes the request
	//  @return LoggingNewrelicResponse
	UpdateLogNewrelicExecute(r APIUpdateLogNewrelicRequest) (*LoggingNewrelicResponse, *http.Response, error)
}

// LoggingNewrelicAPIService LoggingNewrelicAPI service
type LoggingNewrelicAPIService service

// APICreateLogNewrelicRequest represents a request for the resource.
type APICreateLogNewrelicRequest struct {
	ctx                 context.Context
	APIService          LoggingNewrelicAPI
	serviceID           string
	versionID           int32
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	token               *string
	region              *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogNewrelicRequest) Name(name string) *APICreateLogNewrelicRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogNewrelicRequest) Placement(placement string) *APICreateLogNewrelicRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogNewrelicRequest) ResponseCondition(responseCondition string) *APICreateLogNewrelicRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that New Relic Logs can ingest.
func (r *APICreateLogNewrelicRequest) Format(format string) *APICreateLogNewrelicRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogNewrelicRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogNewrelicRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogNewrelicRequest) FormatVersion(formatVersion int32) *APICreateLogNewrelicRequest {
	r.formatVersion = &formatVersion
	return r
}

// Token The Insert API key from the Account page of your New Relic account. Required.
func (r *APICreateLogNewrelicRequest) Token(token string) *APICreateLogNewrelicRequest {
	r.token = &token
	return r
}

// Region The region to which to stream logs.
func (r *APICreateLogNewrelicRequest) Region(region string) *APICreateLogNewrelicRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogNewrelicRequest) Execute() (*LoggingNewrelicResponse, *http.Response, error) {
	return r.APIService.CreateLogNewrelicExecute(r)
}

/*
CreateLogNewrelic Create a New Relic log endpoint

Create a New Relic Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogNewrelicRequest
*/
func (a *LoggingNewrelicAPIService) CreateLogNewrelic(ctx context.Context, serviceID string, versionID int32) APICreateLogNewrelicRequest {
	return APICreateLogNewrelicRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogNewrelicExecute executes the request
//  @return LoggingNewrelicResponse
func (a *LoggingNewrelicAPIService) CreateLogNewrelicExecute(r APICreateLogNewrelicRequest) (*LoggingNewrelicResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingNewrelicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicAPIService.CreateLogNewrelic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelic"
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
	if r.logProcessingRegion != nil {
		localVarFormParams.Add("log_processing_region", parameterToString(*r.logProcessingRegion, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
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

// APIDeleteLogNewrelicRequest represents a request for the resource.
type APIDeleteLogNewrelicRequest struct {
	ctx                 context.Context
	APIService          LoggingNewrelicAPI
	serviceID           string
	versionID           int32
	loggingNewrelicName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogNewrelicRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogNewrelicExecute(r)
}

/*
DeleteLogNewrelic Delete a New Relic log endpoint

Delete the New Relic Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingNewrelicName The name for the real-time logging configuration.
 @return APIDeleteLogNewrelicRequest
*/
func (a *LoggingNewrelicAPIService) DeleteLogNewrelic(ctx context.Context, serviceID string, versionID int32, loggingNewrelicName string) APIDeleteLogNewrelicRequest {
	return APIDeleteLogNewrelicRequest{
		APIService:          a,
		ctx:                 ctx,
		serviceID:           serviceID,
		versionID:           versionID,
		loggingNewrelicName: loggingNewrelicName,
	}
}

// DeleteLogNewrelicExecute executes the request
//  @return InlineResponse200
func (a *LoggingNewrelicAPIService) DeleteLogNewrelicExecute(r APIDeleteLogNewrelicRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicAPIService.DeleteLogNewrelic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelic/{logging_newrelic_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_newrelic_name"+"}", gourl.PathEscape(parameterToString(r.loggingNewrelicName, "")))

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

// APIGetLogNewrelicRequest represents a request for the resource.
type APIGetLogNewrelicRequest struct {
	ctx                 context.Context
	APIService          LoggingNewrelicAPI
	serviceID           string
	versionID           int32
	loggingNewrelicName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogNewrelicRequest) Execute() (*LoggingNewrelicResponse, *http.Response, error) {
	return r.APIService.GetLogNewrelicExecute(r)
}

/*
GetLogNewrelic Get a New Relic log endpoint

Get the details of a New Relic Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingNewrelicName The name for the real-time logging configuration.
 @return APIGetLogNewrelicRequest
*/
func (a *LoggingNewrelicAPIService) GetLogNewrelic(ctx context.Context, serviceID string, versionID int32, loggingNewrelicName string) APIGetLogNewrelicRequest {
	return APIGetLogNewrelicRequest{
		APIService:          a,
		ctx:                 ctx,
		serviceID:           serviceID,
		versionID:           versionID,
		loggingNewrelicName: loggingNewrelicName,
	}
}

// GetLogNewrelicExecute executes the request
//  @return LoggingNewrelicResponse
func (a *LoggingNewrelicAPIService) GetLogNewrelicExecute(r APIGetLogNewrelicRequest) (*LoggingNewrelicResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingNewrelicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicAPIService.GetLogNewrelic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelic/{logging_newrelic_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_newrelic_name"+"}", gourl.PathEscape(parameterToString(r.loggingNewrelicName, "")))

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

// APIListLogNewrelicRequest represents a request for the resource.
type APIListLogNewrelicRequest struct {
	ctx        context.Context
	APIService LoggingNewrelicAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogNewrelicRequest) Execute() ([]LoggingNewrelicResponse, *http.Response, error) {
	return r.APIService.ListLogNewrelicExecute(r)
}

/*
ListLogNewrelic List New Relic log endpoints

List all of the New Relic Logs logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogNewrelicRequest
*/
func (a *LoggingNewrelicAPIService) ListLogNewrelic(ctx context.Context, serviceID string, versionID int32) APIListLogNewrelicRequest {
	return APIListLogNewrelicRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogNewrelicExecute executes the request
//  @return []LoggingNewrelicResponse
func (a *LoggingNewrelicAPIService) ListLogNewrelicExecute(r APIListLogNewrelicRequest) ([]LoggingNewrelicResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingNewrelicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicAPIService.ListLogNewrelic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelic"
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

// APIUpdateLogNewrelicRequest represents a request for the resource.
type APIUpdateLogNewrelicRequest struct {
	ctx                 context.Context
	APIService          LoggingNewrelicAPI
	serviceID           string
	versionID           int32
	loggingNewrelicName string
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	token               *string
	region              *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogNewrelicRequest) Name(name string) *APIUpdateLogNewrelicRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogNewrelicRequest) Placement(placement string) *APIUpdateLogNewrelicRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogNewrelicRequest) ResponseCondition(responseCondition string) *APIUpdateLogNewrelicRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that New Relic Logs can ingest.
func (r *APIUpdateLogNewrelicRequest) Format(format string) *APIUpdateLogNewrelicRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogNewrelicRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogNewrelicRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogNewrelicRequest) FormatVersion(formatVersion int32) *APIUpdateLogNewrelicRequest {
	r.formatVersion = &formatVersion
	return r
}

// Token The Insert API key from the Account page of your New Relic account. Required.
func (r *APIUpdateLogNewrelicRequest) Token(token string) *APIUpdateLogNewrelicRequest {
	r.token = &token
	return r
}

// Region The region to which to stream logs.
func (r *APIUpdateLogNewrelicRequest) Region(region string) *APIUpdateLogNewrelicRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogNewrelicRequest) Execute() (*LoggingNewrelicResponse, *http.Response, error) {
	return r.APIService.UpdateLogNewrelicExecute(r)
}

/*
UpdateLogNewrelic Update a New Relic log endpoint

Update a New Relic Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingNewrelicName The name for the real-time logging configuration.
 @return APIUpdateLogNewrelicRequest
*/
func (a *LoggingNewrelicAPIService) UpdateLogNewrelic(ctx context.Context, serviceID string, versionID int32, loggingNewrelicName string) APIUpdateLogNewrelicRequest {
	return APIUpdateLogNewrelicRequest{
		APIService:          a,
		ctx:                 ctx,
		serviceID:           serviceID,
		versionID:           versionID,
		loggingNewrelicName: loggingNewrelicName,
	}
}

// UpdateLogNewrelicExecute executes the request
//  @return LoggingNewrelicResponse
func (a *LoggingNewrelicAPIService) UpdateLogNewrelicExecute(r APIUpdateLogNewrelicRequest) (*LoggingNewrelicResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingNewrelicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicAPIService.UpdateLogNewrelic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelic/{logging_newrelic_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_newrelic_name"+"}", gourl.PathEscape(parameterToString(r.loggingNewrelicName, "")))

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
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
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
