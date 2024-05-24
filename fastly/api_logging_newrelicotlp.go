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

// LoggingNewrelicotlpAPI defines an interface for interacting with the resource.
type LoggingNewrelicotlpAPI interface {

	/*
	CreateLogNewrelicotlp Create a New Relic OTLP endpoint

	Create a New Relic OTLP logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogNewrelicotlpRequest
	*/
	CreateLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32) APICreateLogNewrelicotlpRequest

	// CreateLogNewrelicotlpExecute executes the request
	//  @return LoggingNewrelicotlpResponse
	CreateLogNewrelicotlpExecute(r APICreateLogNewrelicotlpRequest) (*LoggingNewrelicotlpResponse, *http.Response, error)

	/*
	DeleteLogNewrelicotlp Delete a New Relic OTLP endpoint

	Delete the New Relic OTLP logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingNewrelicotlpName The name for the real-time logging configuration.
	 @return APIDeleteLogNewrelicotlpRequest
	*/
	DeleteLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32, loggingNewrelicotlpName string) APIDeleteLogNewrelicotlpRequest

	// DeleteLogNewrelicotlpExecute executes the request
	//  @return InlineResponse200
	DeleteLogNewrelicotlpExecute(r APIDeleteLogNewrelicotlpRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogNewrelicotlp Get a New Relic OTLP endpoint

	Get the details of a New Relic OTLP logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingNewrelicotlpName The name for the real-time logging configuration.
	 @return APIGetLogNewrelicotlpRequest
	*/
	GetLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32, loggingNewrelicotlpName string) APIGetLogNewrelicotlpRequest

	// GetLogNewrelicotlpExecute executes the request
	//  @return LoggingNewrelicotlpResponse
	GetLogNewrelicotlpExecute(r APIGetLogNewrelicotlpRequest) (*LoggingNewrelicotlpResponse, *http.Response, error)

	/*
	ListLogNewrelicotlp List New Relic OTLP endpoints

	List all of the New Relic OTLP logging objects for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogNewrelicotlpRequest
	*/
	ListLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32) APIListLogNewrelicotlpRequest

	// ListLogNewrelicotlpExecute executes the request
	//  @return []LoggingNewrelicotlpResponse
	ListLogNewrelicotlpExecute(r APIListLogNewrelicotlpRequest) ([]LoggingNewrelicotlpResponse, *http.Response, error)

	/*
	UpdateLogNewrelicotlp Update a New Relic log endpoint

	Update a New Relic OTLP logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingNewrelicotlpName The name for the real-time logging configuration.
	 @return APIUpdateLogNewrelicotlpRequest
	*/
	UpdateLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32, loggingNewrelicotlpName string) APIUpdateLogNewrelicotlpRequest

	// UpdateLogNewrelicotlpExecute executes the request
	//  @return LoggingNewrelicotlpResponse
	UpdateLogNewrelicotlpExecute(r APIUpdateLogNewrelicotlpRequest) (*LoggingNewrelicotlpResponse, *http.Response, error)
}

// LoggingNewrelicotlpAPIService LoggingNewrelicotlpAPI service
type LoggingNewrelicotlpAPIService service

// APICreateLogNewrelicotlpRequest represents a request for the resource.
type APICreateLogNewrelicotlpRequest struct {
	ctx context.Context
	APIService LoggingNewrelicotlpAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	token *string
	region *string
	url *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogNewrelicotlpRequest) Name(name string) *APICreateLogNewrelicotlpRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogNewrelicotlpRequest) Placement(placement string) *APICreateLogNewrelicotlpRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogNewrelicotlpRequest) ResponseCondition(responseCondition string) *APICreateLogNewrelicotlpRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogNewrelicotlpRequest) Format(format string) *APICreateLogNewrelicotlpRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogNewrelicotlpRequest) FormatVersion(formatVersion int32) *APICreateLogNewrelicotlpRequest {
	r.formatVersion = &formatVersion
	return r
}
// Token The Insert API key from the Account page of your New Relic account. Required.
func (r *APICreateLogNewrelicotlpRequest) Token(token string) *APICreateLogNewrelicotlpRequest {
	r.token = &token
	return r
}
// Region The region to which to stream logs.
func (r *APICreateLogNewrelicotlpRequest) Region(region string) *APICreateLogNewrelicotlpRequest {
	r.region = &region
	return r
}
// URL (Optional) URL of the New Relic Trace Observer, if you are using New Relic Infinite Tracing.
func (r *APICreateLogNewrelicotlpRequest) URL(url string) *APICreateLogNewrelicotlpRequest {
	r.url = &url
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogNewrelicotlpRequest) Execute() (*LoggingNewrelicotlpResponse, *http.Response, error) {
	return r.APIService.CreateLogNewrelicotlpExecute(r)
}

/*
CreateLogNewrelicotlp Create a New Relic OTLP endpoint

Create a New Relic OTLP logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogNewrelicotlpRequest
*/
func (a *LoggingNewrelicotlpAPIService) CreateLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32) APICreateLogNewrelicotlpRequest {
	return APICreateLogNewrelicotlpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogNewrelicotlpExecute executes the request
//  @return LoggingNewrelicotlpResponse
func (a *LoggingNewrelicotlpAPIService) CreateLogNewrelicotlpExecute(r APICreateLogNewrelicotlpRequest) (*LoggingNewrelicotlpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingNewrelicotlpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicotlpAPIService.CreateLogNewrelicotlp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelicotlp"
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
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
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

// APIDeleteLogNewrelicotlpRequest represents a request for the resource.
type APIDeleteLogNewrelicotlpRequest struct {
	ctx context.Context
	APIService LoggingNewrelicotlpAPI
	serviceID string
	versionID int32
	loggingNewrelicotlpName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogNewrelicotlpRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogNewrelicotlpExecute(r)
}

/*
DeleteLogNewrelicotlp Delete a New Relic OTLP endpoint

Delete the New Relic OTLP logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingNewrelicotlpName The name for the real-time logging configuration.
 @return APIDeleteLogNewrelicotlpRequest
*/
func (a *LoggingNewrelicotlpAPIService) DeleteLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32, loggingNewrelicotlpName string) APIDeleteLogNewrelicotlpRequest {
	return APIDeleteLogNewrelicotlpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingNewrelicotlpName: loggingNewrelicotlpName,
	}
}

// DeleteLogNewrelicotlpExecute executes the request
//  @return InlineResponse200
func (a *LoggingNewrelicotlpAPIService) DeleteLogNewrelicotlpExecute(r APIDeleteLogNewrelicotlpRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicotlpAPIService.DeleteLogNewrelicotlp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelicotlp/{logging_newrelicotlp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_newrelicotlp_name"+"}", gourl.PathEscape(parameterToString(r.loggingNewrelicotlpName, "")))

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

// APIGetLogNewrelicotlpRequest represents a request for the resource.
type APIGetLogNewrelicotlpRequest struct {
	ctx context.Context
	APIService LoggingNewrelicotlpAPI
	serviceID string
	versionID int32
	loggingNewrelicotlpName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogNewrelicotlpRequest) Execute() (*LoggingNewrelicotlpResponse, *http.Response, error) {
	return r.APIService.GetLogNewrelicotlpExecute(r)
}

/*
GetLogNewrelicotlp Get a New Relic OTLP endpoint

Get the details of a New Relic OTLP logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingNewrelicotlpName The name for the real-time logging configuration.
 @return APIGetLogNewrelicotlpRequest
*/
func (a *LoggingNewrelicotlpAPIService) GetLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32, loggingNewrelicotlpName string) APIGetLogNewrelicotlpRequest {
	return APIGetLogNewrelicotlpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingNewrelicotlpName: loggingNewrelicotlpName,
	}
}

// GetLogNewrelicotlpExecute executes the request
//  @return LoggingNewrelicotlpResponse
func (a *LoggingNewrelicotlpAPIService) GetLogNewrelicotlpExecute(r APIGetLogNewrelicotlpRequest) (*LoggingNewrelicotlpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingNewrelicotlpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicotlpAPIService.GetLogNewrelicotlp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelicotlp/{logging_newrelicotlp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_newrelicotlp_name"+"}", gourl.PathEscape(parameterToString(r.loggingNewrelicotlpName, "")))

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

// APIListLogNewrelicotlpRequest represents a request for the resource.
type APIListLogNewrelicotlpRequest struct {
	ctx context.Context
	APIService LoggingNewrelicotlpAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogNewrelicotlpRequest) Execute() ([]LoggingNewrelicotlpResponse, *http.Response, error) {
	return r.APIService.ListLogNewrelicotlpExecute(r)
}

/*
ListLogNewrelicotlp List New Relic OTLP endpoints

List all of the New Relic OTLP logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogNewrelicotlpRequest
*/
func (a *LoggingNewrelicotlpAPIService) ListLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32) APIListLogNewrelicotlpRequest {
	return APIListLogNewrelicotlpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogNewrelicotlpExecute executes the request
//  @return []LoggingNewrelicotlpResponse
func (a *LoggingNewrelicotlpAPIService) ListLogNewrelicotlpExecute(r APIListLogNewrelicotlpRequest) ([]LoggingNewrelicotlpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingNewrelicotlpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicotlpAPIService.ListLogNewrelicotlp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelicotlp"
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

// APIUpdateLogNewrelicotlpRequest represents a request for the resource.
type APIUpdateLogNewrelicotlpRequest struct {
	ctx context.Context
	APIService LoggingNewrelicotlpAPI
	serviceID string
	versionID int32
	loggingNewrelicotlpName string
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	token *string
	region *string
	url *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogNewrelicotlpRequest) Name(name string) *APIUpdateLogNewrelicotlpRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogNewrelicotlpRequest) Placement(placement string) *APIUpdateLogNewrelicotlpRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogNewrelicotlpRequest) ResponseCondition(responseCondition string) *APIUpdateLogNewrelicotlpRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogNewrelicotlpRequest) Format(format string) *APIUpdateLogNewrelicotlpRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogNewrelicotlpRequest) FormatVersion(formatVersion int32) *APIUpdateLogNewrelicotlpRequest {
	r.formatVersion = &formatVersion
	return r
}
// Token The Insert API key from the Account page of your New Relic account. Required.
func (r *APIUpdateLogNewrelicotlpRequest) Token(token string) *APIUpdateLogNewrelicotlpRequest {
	r.token = &token
	return r
}
// Region The region to which to stream logs.
func (r *APIUpdateLogNewrelicotlpRequest) Region(region string) *APIUpdateLogNewrelicotlpRequest {
	r.region = &region
	return r
}
// URL (Optional) URL of the New Relic Trace Observer, if you are using New Relic Infinite Tracing.
func (r *APIUpdateLogNewrelicotlpRequest) URL(url string) *APIUpdateLogNewrelicotlpRequest {
	r.url = &url
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogNewrelicotlpRequest) Execute() (*LoggingNewrelicotlpResponse, *http.Response, error) {
	return r.APIService.UpdateLogNewrelicotlpExecute(r)
}

/*
UpdateLogNewrelicotlp Update a New Relic log endpoint

Update a New Relic OTLP logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingNewrelicotlpName The name for the real-time logging configuration.
 @return APIUpdateLogNewrelicotlpRequest
*/
func (a *LoggingNewrelicotlpAPIService) UpdateLogNewrelicotlp(ctx context.Context, serviceID string, versionID int32, loggingNewrelicotlpName string) APIUpdateLogNewrelicotlpRequest {
	return APIUpdateLogNewrelicotlpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingNewrelicotlpName: loggingNewrelicotlpName,
	}
}

// UpdateLogNewrelicotlpExecute executes the request
//  @return LoggingNewrelicotlpResponse
func (a *LoggingNewrelicotlpAPIService) UpdateLogNewrelicotlpExecute(r APIUpdateLogNewrelicotlpRequest) (*LoggingNewrelicotlpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingNewrelicotlpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingNewrelicotlpAPIService.UpdateLogNewrelicotlp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/newrelicotlp/{logging_newrelicotlp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_newrelicotlp_name"+"}", gourl.PathEscape(parameterToString(r.loggingNewrelicotlpName, "")))

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
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
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
