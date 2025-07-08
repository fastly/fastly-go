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

// LoggingPapertrailAPI defines an interface for interacting with the resource.
type LoggingPapertrailAPI interface {

	/*
		CreateLogPapertrail Create a Papertrail log endpoint

		Create a Papertrail for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogPapertrailRequest
	*/
	CreateLogPapertrail(ctx context.Context, serviceID string, versionID int32) APICreateLogPapertrailRequest

	// CreateLogPapertrailExecute executes the request
	//  @return LoggingPapertrailResponse
	CreateLogPapertrailExecute(r APICreateLogPapertrailRequest) (*LoggingPapertrailResponse, *http.Response, error)

	/*
		DeleteLogPapertrail Delete a Papertrail log endpoint

		Delete the Papertrail for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingPapertrailName The name for the real-time logging configuration.
		 @return APIDeleteLogPapertrailRequest
	*/
	DeleteLogPapertrail(ctx context.Context, serviceID string, versionID int32, loggingPapertrailName string) APIDeleteLogPapertrailRequest

	// DeleteLogPapertrailExecute executes the request
	//  @return InlineResponse200
	DeleteLogPapertrailExecute(r APIDeleteLogPapertrailRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogPapertrail Get a Papertrail log endpoint

		Get the Papertrail for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingPapertrailName The name for the real-time logging configuration.
		 @return APIGetLogPapertrailRequest
	*/
	GetLogPapertrail(ctx context.Context, serviceID string, versionID int32, loggingPapertrailName string) APIGetLogPapertrailRequest

	// GetLogPapertrailExecute executes the request
	//  @return LoggingPapertrailResponse
	GetLogPapertrailExecute(r APIGetLogPapertrailRequest) (*LoggingPapertrailResponse, *http.Response, error)

	/*
		ListLogPapertrail List Papertrail log endpoints

		List all of the Papertrails for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogPapertrailRequest
	*/
	ListLogPapertrail(ctx context.Context, serviceID string, versionID int32) APIListLogPapertrailRequest

	// ListLogPapertrailExecute executes the request
	//  @return []LoggingPapertrailResponse
	ListLogPapertrailExecute(r APIListLogPapertrailRequest) ([]LoggingPapertrailResponse, *http.Response, error)

	/*
		UpdateLogPapertrail Update a Papertrail log endpoint

		Update the Papertrail for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingPapertrailName The name for the real-time logging configuration.
		 @return APIUpdateLogPapertrailRequest
	*/
	UpdateLogPapertrail(ctx context.Context, serviceID string, versionID int32, loggingPapertrailName string) APIUpdateLogPapertrailRequest

	// UpdateLogPapertrailExecute executes the request
	//  @return LoggingPapertrailResponse
	UpdateLogPapertrailExecute(r APIUpdateLogPapertrailRequest) (*LoggingPapertrailResponse, *http.Response, error)
}

// LoggingPapertrailAPIService LoggingPapertrailAPI service
type LoggingPapertrailAPIService service

// APICreateLogPapertrailRequest represents a request for the resource.
type APICreateLogPapertrailRequest struct {
	ctx                 context.Context
	APIService          LoggingPapertrailAPI
	serviceID           string
	versionID           int32
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	address             *string
	port                *int32
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogPapertrailRequest) Name(name string) *APICreateLogPapertrailRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogPapertrailRequest) Placement(placement string) *APICreateLogPapertrailRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogPapertrailRequest) ResponseCondition(responseCondition string) *APICreateLogPapertrailRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogPapertrailRequest) Format(format string) *APICreateLogPapertrailRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogPapertrailRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogPapertrailRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogPapertrailRequest) FormatVersion(formatVersion int32) *APICreateLogPapertrailRequest {
	r.formatVersion = &formatVersion
	return r
}

// Address A hostname or IPv4 address.
func (r *APICreateLogPapertrailRequest) Address(address string) *APICreateLogPapertrailRequest {
	r.address = &address
	return r
}

// Port The port number.
func (r *APICreateLogPapertrailRequest) Port(port int32) *APICreateLogPapertrailRequest {
	r.port = &port
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogPapertrailRequest) Execute() (*LoggingPapertrailResponse, *http.Response, error) {
	return r.APIService.CreateLogPapertrailExecute(r)
}

/*
CreateLogPapertrail Create a Papertrail log endpoint

Create a Papertrail for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogPapertrailRequest
*/
func (a *LoggingPapertrailAPIService) CreateLogPapertrail(ctx context.Context, serviceID string, versionID int32) APICreateLogPapertrailRequest {
	return APICreateLogPapertrailRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogPapertrailExecute executes the request
//  @return LoggingPapertrailResponse
func (a *LoggingPapertrailAPIService) CreateLogPapertrailExecute(r APICreateLogPapertrailRequest) (*LoggingPapertrailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingPapertrailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPapertrailAPIService.CreateLogPapertrail")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/papertrail"
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
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
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

// APIDeleteLogPapertrailRequest represents a request for the resource.
type APIDeleteLogPapertrailRequest struct {
	ctx                   context.Context
	APIService            LoggingPapertrailAPI
	serviceID             string
	versionID             int32
	loggingPapertrailName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogPapertrailRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogPapertrailExecute(r)
}

/*
DeleteLogPapertrail Delete a Papertrail log endpoint

Delete the Papertrail for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingPapertrailName The name for the real-time logging configuration.
 @return APIDeleteLogPapertrailRequest
*/
func (a *LoggingPapertrailAPIService) DeleteLogPapertrail(ctx context.Context, serviceID string, versionID int32, loggingPapertrailName string) APIDeleteLogPapertrailRequest {
	return APIDeleteLogPapertrailRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceID:             serviceID,
		versionID:             versionID,
		loggingPapertrailName: loggingPapertrailName,
	}
}

// DeleteLogPapertrailExecute executes the request
//  @return InlineResponse200
func (a *LoggingPapertrailAPIService) DeleteLogPapertrailExecute(r APIDeleteLogPapertrailRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPapertrailAPIService.DeleteLogPapertrail")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/papertrail/{logging_papertrail_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_papertrail_name"+"}", gourl.PathEscape(parameterToString(r.loggingPapertrailName, "")))

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

// APIGetLogPapertrailRequest represents a request for the resource.
type APIGetLogPapertrailRequest struct {
	ctx                   context.Context
	APIService            LoggingPapertrailAPI
	serviceID             string
	versionID             int32
	loggingPapertrailName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogPapertrailRequest) Execute() (*LoggingPapertrailResponse, *http.Response, error) {
	return r.APIService.GetLogPapertrailExecute(r)
}

/*
GetLogPapertrail Get a Papertrail log endpoint

Get the Papertrail for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingPapertrailName The name for the real-time logging configuration.
 @return APIGetLogPapertrailRequest
*/
func (a *LoggingPapertrailAPIService) GetLogPapertrail(ctx context.Context, serviceID string, versionID int32, loggingPapertrailName string) APIGetLogPapertrailRequest {
	return APIGetLogPapertrailRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceID:             serviceID,
		versionID:             versionID,
		loggingPapertrailName: loggingPapertrailName,
	}
}

// GetLogPapertrailExecute executes the request
//  @return LoggingPapertrailResponse
func (a *LoggingPapertrailAPIService) GetLogPapertrailExecute(r APIGetLogPapertrailRequest) (*LoggingPapertrailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingPapertrailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPapertrailAPIService.GetLogPapertrail")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/papertrail/{logging_papertrail_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_papertrail_name"+"}", gourl.PathEscape(parameterToString(r.loggingPapertrailName, "")))

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

// APIListLogPapertrailRequest represents a request for the resource.
type APIListLogPapertrailRequest struct {
	ctx        context.Context
	APIService LoggingPapertrailAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogPapertrailRequest) Execute() ([]LoggingPapertrailResponse, *http.Response, error) {
	return r.APIService.ListLogPapertrailExecute(r)
}

/*
ListLogPapertrail List Papertrail log endpoints

List all of the Papertrails for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogPapertrailRequest
*/
func (a *LoggingPapertrailAPIService) ListLogPapertrail(ctx context.Context, serviceID string, versionID int32) APIListLogPapertrailRequest {
	return APIListLogPapertrailRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogPapertrailExecute executes the request
//  @return []LoggingPapertrailResponse
func (a *LoggingPapertrailAPIService) ListLogPapertrailExecute(r APIListLogPapertrailRequest) ([]LoggingPapertrailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingPapertrailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPapertrailAPIService.ListLogPapertrail")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/papertrail"
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

// APIUpdateLogPapertrailRequest represents a request for the resource.
type APIUpdateLogPapertrailRequest struct {
	ctx                   context.Context
	APIService            LoggingPapertrailAPI
	serviceID             string
	versionID             int32
	loggingPapertrailName string
	name                  *string
	placement             *string
	responseCondition     *string
	format                *string
	logProcessingRegion   *string
	formatVersion         *int32
	address               *string
	port                  *int32
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogPapertrailRequest) Name(name string) *APIUpdateLogPapertrailRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogPapertrailRequest) Placement(placement string) *APIUpdateLogPapertrailRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogPapertrailRequest) ResponseCondition(responseCondition string) *APIUpdateLogPapertrailRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APIUpdateLogPapertrailRequest) Format(format string) *APIUpdateLogPapertrailRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogPapertrailRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogPapertrailRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogPapertrailRequest) FormatVersion(formatVersion int32) *APIUpdateLogPapertrailRequest {
	r.formatVersion = &formatVersion
	return r
}

// Address A hostname or IPv4 address.
func (r *APIUpdateLogPapertrailRequest) Address(address string) *APIUpdateLogPapertrailRequest {
	r.address = &address
	return r
}

// Port The port number.
func (r *APIUpdateLogPapertrailRequest) Port(port int32) *APIUpdateLogPapertrailRequest {
	r.port = &port
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogPapertrailRequest) Execute() (*LoggingPapertrailResponse, *http.Response, error) {
	return r.APIService.UpdateLogPapertrailExecute(r)
}

/*
UpdateLogPapertrail Update a Papertrail log endpoint

Update the Papertrail for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingPapertrailName The name for the real-time logging configuration.
 @return APIUpdateLogPapertrailRequest
*/
func (a *LoggingPapertrailAPIService) UpdateLogPapertrail(ctx context.Context, serviceID string, versionID int32, loggingPapertrailName string) APIUpdateLogPapertrailRequest {
	return APIUpdateLogPapertrailRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceID:             serviceID,
		versionID:             versionID,
		loggingPapertrailName: loggingPapertrailName,
	}
}

// UpdateLogPapertrailExecute executes the request
//  @return LoggingPapertrailResponse
func (a *LoggingPapertrailAPIService) UpdateLogPapertrailExecute(r APIUpdateLogPapertrailRequest) (*LoggingPapertrailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingPapertrailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPapertrailAPIService.UpdateLogPapertrail")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/papertrail/{logging_papertrail_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_papertrail_name"+"}", gourl.PathEscape(parameterToString(r.loggingPapertrailName, "")))

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
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
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
