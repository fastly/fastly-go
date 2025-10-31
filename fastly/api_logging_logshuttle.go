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

// LoggingLogshuttleAPI defines an interface for interacting with the resource.
type LoggingLogshuttleAPI interface {

	/*
		CreateLogLogshuttle Create a Log Shuttle log endpoint

		Create a Log Shuttle logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APICreateLogLogshuttleRequest
	*/
	CreateLogLogshuttle(ctx context.Context, serviceId string, versionId int32) APICreateLogLogshuttleRequest

	// CreateLogLogshuttleExecute executes the request
	//  @return LoggingLogshuttleResponse
	CreateLogLogshuttleExecute(r APICreateLogLogshuttleRequest) (*LoggingLogshuttleResponse, *http.Response, error)

	/*
		DeleteLogLogshuttle Delete a Log Shuttle log endpoint

		Delete the Log Shuttle logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingLogshuttleName The name for the real-time logging configuration.
		 @return APIDeleteLogLogshuttleRequest
	*/
	DeleteLogLogshuttle(ctx context.Context, serviceId string, versionId int32, loggingLogshuttleName string) APIDeleteLogLogshuttleRequest

	// DeleteLogLogshuttleExecute executes the request
	//  @return InlineResponse200
	DeleteLogLogshuttleExecute(r APIDeleteLogLogshuttleRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogLogshuttle Get a Log Shuttle log endpoint

		Get the Log Shuttle logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingLogshuttleName The name for the real-time logging configuration.
		 @return APIGetLogLogshuttleRequest
	*/
	GetLogLogshuttle(ctx context.Context, serviceId string, versionId int32, loggingLogshuttleName string) APIGetLogLogshuttleRequest

	// GetLogLogshuttleExecute executes the request
	//  @return LoggingLogshuttleResponse
	GetLogLogshuttleExecute(r APIGetLogLogshuttleRequest) (*LoggingLogshuttleResponse, *http.Response, error)

	/*
		ListLogLogshuttle List Log Shuttle log endpoints

		List all of the Log Shuttle logging endpoints for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APIListLogLogshuttleRequest
	*/
	ListLogLogshuttle(ctx context.Context, serviceId string, versionId int32) APIListLogLogshuttleRequest

	// ListLogLogshuttleExecute executes the request
	//  @return []LoggingLogshuttleResponse
	ListLogLogshuttleExecute(r APIListLogLogshuttleRequest) ([]LoggingLogshuttleResponse, *http.Response, error)

	/*
		UpdateLogLogshuttle Update a Log Shuttle log endpoint

		Update the Log Shuttle logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingLogshuttleName The name for the real-time logging configuration.
		 @return APIUpdateLogLogshuttleRequest
	*/
	UpdateLogLogshuttle(ctx context.Context, serviceId string, versionId int32, loggingLogshuttleName string) APIUpdateLogLogshuttleRequest

	// UpdateLogLogshuttleExecute executes the request
	//  @return LoggingLogshuttleResponse
	UpdateLogLogshuttleExecute(r APIUpdateLogLogshuttleRequest) (*LoggingLogshuttleResponse, *http.Response, error)
}

// LoggingLogshuttleAPIService LoggingLogshuttleAPI service
type LoggingLogshuttleAPIService service

// APICreateLogLogshuttleRequest represents a request for the resource.
type APICreateLogLogshuttleRequest struct {
	ctx                 context.Context
	APIService          LoggingLogshuttleAPI
	serviceId           string
	versionId           int32
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	token               *string
	url                 *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogLogshuttleRequest) Name(name string) *APICreateLogLogshuttleRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogLogshuttleRequest) Placement(placement string) *APICreateLogLogshuttleRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogLogshuttleRequest) ResponseCondition(responseCondition string) *APICreateLogLogshuttleRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogLogshuttleRequest) Format(format string) *APICreateLogLogshuttleRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogLogshuttleRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogLogshuttleRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogLogshuttleRequest) FormatVersion(formatVersion int32) *APICreateLogLogshuttleRequest {
	r.formatVersion = &formatVersion
	return r
}

// Token The data authentication token associated with this endpoint.
func (r *APICreateLogLogshuttleRequest) Token(token string) *APICreateLogLogshuttleRequest {
	r.token = &token
	return r
}

// Url The URL to stream logs to.
func (r *APICreateLogLogshuttleRequest) Url(url string) *APICreateLogLogshuttleRequest {
	r.url = &url
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogLogshuttleRequest) Execute() (*LoggingLogshuttleResponse, *http.Response, error) {
	return r.APIService.CreateLogLogshuttleExecute(r)
}

/*
CreateLogLogshuttle Create a Log Shuttle log endpoint

Create a Log Shuttle logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APICreateLogLogshuttleRequest
*/
func (a *LoggingLogshuttleAPIService) CreateLogLogshuttle(ctx context.Context, serviceId string, versionId int32) APICreateLogLogshuttleRequest {
	return APICreateLogLogshuttleRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// CreateLogLogshuttleExecute executes the request
//  @return LoggingLogshuttleResponse
func (a *LoggingLogshuttleAPIService) CreateLogLogshuttleExecute(r APICreateLogLogshuttleRequest) (*LoggingLogshuttleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingLogshuttleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogshuttleAPIService.CreateLogLogshuttle")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logshuttle"
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
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
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

// APIDeleteLogLogshuttleRequest represents a request for the resource.
type APIDeleteLogLogshuttleRequest struct {
	ctx                   context.Context
	APIService            LoggingLogshuttleAPI
	serviceId             string
	versionId             int32
	loggingLogshuttleName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogLogshuttleRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogLogshuttleExecute(r)
}

/*
DeleteLogLogshuttle Delete a Log Shuttle log endpoint

Delete the Log Shuttle logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingLogshuttleName The name for the real-time logging configuration.
 @return APIDeleteLogLogshuttleRequest
*/
func (a *LoggingLogshuttleAPIService) DeleteLogLogshuttle(ctx context.Context, serviceId string, versionId int32, loggingLogshuttleName string) APIDeleteLogLogshuttleRequest {
	return APIDeleteLogLogshuttleRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceId:             serviceId,
		versionId:             versionId,
		loggingLogshuttleName: loggingLogshuttleName,
	}
}

// DeleteLogLogshuttleExecute executes the request
//  @return InlineResponse200
func (a *LoggingLogshuttleAPIService) DeleteLogLogshuttleExecute(r APIDeleteLogLogshuttleRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogshuttleAPIService.DeleteLogLogshuttle")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logshuttle/{logging_logshuttle_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_logshuttle_name"+"}", gourl.PathEscape(parameterToString(r.loggingLogshuttleName, "")))

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

// APIGetLogLogshuttleRequest represents a request for the resource.
type APIGetLogLogshuttleRequest struct {
	ctx                   context.Context
	APIService            LoggingLogshuttleAPI
	serviceId             string
	versionId             int32
	loggingLogshuttleName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogLogshuttleRequest) Execute() (*LoggingLogshuttleResponse, *http.Response, error) {
	return r.APIService.GetLogLogshuttleExecute(r)
}

/*
GetLogLogshuttle Get a Log Shuttle log endpoint

Get the Log Shuttle logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingLogshuttleName The name for the real-time logging configuration.
 @return APIGetLogLogshuttleRequest
*/
func (a *LoggingLogshuttleAPIService) GetLogLogshuttle(ctx context.Context, serviceId string, versionId int32, loggingLogshuttleName string) APIGetLogLogshuttleRequest {
	return APIGetLogLogshuttleRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceId:             serviceId,
		versionId:             versionId,
		loggingLogshuttleName: loggingLogshuttleName,
	}
}

// GetLogLogshuttleExecute executes the request
//  @return LoggingLogshuttleResponse
func (a *LoggingLogshuttleAPIService) GetLogLogshuttleExecute(r APIGetLogLogshuttleRequest) (*LoggingLogshuttleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingLogshuttleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogshuttleAPIService.GetLogLogshuttle")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logshuttle/{logging_logshuttle_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_logshuttle_name"+"}", gourl.PathEscape(parameterToString(r.loggingLogshuttleName, "")))

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

// APIListLogLogshuttleRequest represents a request for the resource.
type APIListLogLogshuttleRequest struct {
	ctx        context.Context
	APIService LoggingLogshuttleAPI
	serviceId  string
	versionId  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogLogshuttleRequest) Execute() ([]LoggingLogshuttleResponse, *http.Response, error) {
	return r.APIService.ListLogLogshuttleExecute(r)
}

/*
ListLogLogshuttle List Log Shuttle log endpoints

List all of the Log Shuttle logging endpoints for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APIListLogLogshuttleRequest
*/
func (a *LoggingLogshuttleAPIService) ListLogLogshuttle(ctx context.Context, serviceId string, versionId int32) APIListLogLogshuttleRequest {
	return APIListLogLogshuttleRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// ListLogLogshuttleExecute executes the request
//  @return []LoggingLogshuttleResponse
func (a *LoggingLogshuttleAPIService) ListLogLogshuttleExecute(r APIListLogLogshuttleRequest) ([]LoggingLogshuttleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingLogshuttleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogshuttleAPIService.ListLogLogshuttle")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logshuttle"
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

// APIUpdateLogLogshuttleRequest represents a request for the resource.
type APIUpdateLogLogshuttleRequest struct {
	ctx                   context.Context
	APIService            LoggingLogshuttleAPI
	serviceId             string
	versionId             int32
	loggingLogshuttleName string
	name                  *string
	placement             *string
	responseCondition     *string
	format                *string
	logProcessingRegion   *string
	formatVersion         *int32
	token                 *string
	url                   *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogLogshuttleRequest) Name(name string) *APIUpdateLogLogshuttleRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogLogshuttleRequest) Placement(placement string) *APIUpdateLogLogshuttleRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogLogshuttleRequest) ResponseCondition(responseCondition string) *APIUpdateLogLogshuttleRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APIUpdateLogLogshuttleRequest) Format(format string) *APIUpdateLogLogshuttleRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogLogshuttleRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogLogshuttleRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogLogshuttleRequest) FormatVersion(formatVersion int32) *APIUpdateLogLogshuttleRequest {
	r.formatVersion = &formatVersion
	return r
}

// Token The data authentication token associated with this endpoint.
func (r *APIUpdateLogLogshuttleRequest) Token(token string) *APIUpdateLogLogshuttleRequest {
	r.token = &token
	return r
}

// Url The URL to stream logs to.
func (r *APIUpdateLogLogshuttleRequest) Url(url string) *APIUpdateLogLogshuttleRequest {
	r.url = &url
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogLogshuttleRequest) Execute() (*LoggingLogshuttleResponse, *http.Response, error) {
	return r.APIService.UpdateLogLogshuttleExecute(r)
}

/*
UpdateLogLogshuttle Update a Log Shuttle log endpoint

Update the Log Shuttle logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingLogshuttleName The name for the real-time logging configuration.
 @return APIUpdateLogLogshuttleRequest
*/
func (a *LoggingLogshuttleAPIService) UpdateLogLogshuttle(ctx context.Context, serviceId string, versionId int32, loggingLogshuttleName string) APIUpdateLogLogshuttleRequest {
	return APIUpdateLogLogshuttleRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceId:             serviceId,
		versionId:             versionId,
		loggingLogshuttleName: loggingLogshuttleName,
	}
}

// UpdateLogLogshuttleExecute executes the request
//  @return LoggingLogshuttleResponse
func (a *LoggingLogshuttleAPIService) UpdateLogLogshuttleExecute(r APIUpdateLogLogshuttleRequest) (*LoggingLogshuttleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingLogshuttleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingLogshuttleAPIService.UpdateLogLogshuttle")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/logshuttle/{logging_logshuttle_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_logshuttle_name"+"}", gourl.PathEscape(parameterToString(r.loggingLogshuttleName, "")))

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
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
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
