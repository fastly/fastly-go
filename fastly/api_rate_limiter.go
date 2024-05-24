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

// RateLimiterAPI defines an interface for interacting with the resource.
type RateLimiterAPI interface {

	/*
	CreateRateLimiter Create a rate limiter

	Create a rate limiter for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateRateLimiterRequest
	*/
	CreateRateLimiter(ctx context.Context, serviceID string, versionID int32) APICreateRateLimiterRequest

	// CreateRateLimiterExecute executes the request
	//  @return RateLimiterResponse
	CreateRateLimiterExecute(r APICreateRateLimiterRequest) (*RateLimiterResponse, *http.Response, error)

	/*
	DeleteRateLimiter Delete a rate limiter

	Delete a rate limiter by its ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param rateLimiterID Alphanumeric string identifying the rate limiter.
	 @return APIDeleteRateLimiterRequest
	*/
	DeleteRateLimiter(ctx context.Context, rateLimiterID string) APIDeleteRateLimiterRequest

	// DeleteRateLimiterExecute executes the request
	//  @return InlineResponse200
	DeleteRateLimiterExecute(r APIDeleteRateLimiterRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetRateLimiter Get a rate limiter

	Get a rate limiter by its ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param rateLimiterID Alphanumeric string identifying the rate limiter.
	 @return APIGetRateLimiterRequest
	*/
	GetRateLimiter(ctx context.Context, rateLimiterID string) APIGetRateLimiterRequest

	// GetRateLimiterExecute executes the request
	//  @return RateLimiterResponse
	GetRateLimiterExecute(r APIGetRateLimiterRequest) (*RateLimiterResponse, *http.Response, error)

	/*
	ListRateLimiters List rate limiters

	List all rate limiters for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListRateLimitersRequest
	*/
	ListRateLimiters(ctx context.Context, serviceID string, versionID int32) APIListRateLimitersRequest

	// ListRateLimitersExecute executes the request
	//  @return []RateLimiterResponse
	ListRateLimitersExecute(r APIListRateLimitersRequest) ([]RateLimiterResponse, *http.Response, error)

	/*
	UpdateRateLimiter Update a rate limiter

	Update a rate limiter by its ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param rateLimiterID Alphanumeric string identifying the rate limiter.
	 @return APIUpdateRateLimiterRequest
	*/
	UpdateRateLimiter(ctx context.Context, rateLimiterID string) APIUpdateRateLimiterRequest

	// UpdateRateLimiterExecute executes the request
	//  @return RateLimiterResponse
	UpdateRateLimiterExecute(r APIUpdateRateLimiterRequest) (*RateLimiterResponse, *http.Response, error)
}

// RateLimiterAPIService RateLimiterAPI service
type RateLimiterAPIService service

// APICreateRateLimiterRequest represents a request for the resource.
type APICreateRateLimiterRequest struct {
	ctx context.Context
	APIService RateLimiterAPI
	serviceID string
	versionID int32
	name *string
	uriDictionaryName *string
	httpMethods *[]string
	rpsLimit *int32
	windowSize *int32
	clientKey *[]string
	penaltyBoxDuration *int32
	action *string
	responseObjectName *string
	loggerType *string
	featureRevision *int32
}

// Name A human readable name for the rate limiting rule.
func (r *APICreateRateLimiterRequest) Name(name string) *APICreateRateLimiterRequest {
	r.name = &name
	return r
}
// URIDictionaryName The name of an Edge Dictionary containing URIs as keys. If not defined or &#x60;null&#x60;, all origin URIs will be rate limited.
func (r *APICreateRateLimiterRequest) URIDictionaryName(uriDictionaryName string) *APICreateRateLimiterRequest {
	r.uriDictionaryName = &uriDictionaryName
	return r
}
// HTTPMethods Array of HTTP methods to apply rate limiting to.
func (r *APICreateRateLimiterRequest) HTTPMethods(httpMethods []string) *APICreateRateLimiterRequest {
	r.httpMethods = &httpMethods
	return r
}
// RpsLimit Upper limit of requests per second allowed by the rate limiter.
func (r *APICreateRateLimiterRequest) RpsLimit(rpsLimit int32) *APICreateRateLimiterRequest {
	r.rpsLimit = &rpsLimit
	return r
}
// WindowSize Number of seconds during which the RPS limit must be exceeded in order to trigger a violation.
func (r *APICreateRateLimiterRequest) WindowSize(windowSize int32) *APICreateRateLimiterRequest {
	r.windowSize = &windowSize
	return r
}
// ClientKey Array of VCL variables used to generate a counter key to identify a client. Example variables include &#x60;req.http.Fastly-Client-IP&#x60;, &#x60;req.http.User-Agent&#x60;, or a custom header like &#x60;req.http.API-Key&#x60;.
func (r *APICreateRateLimiterRequest) ClientKey(clientKey []string) *APICreateRateLimiterRequest {
	r.clientKey = &clientKey
	return r
}
// PenaltyBoxDuration Length of time in minutes that the rate limiter is in effect after the initial violation is detected.
func (r *APICreateRateLimiterRequest) PenaltyBoxDuration(penaltyBoxDuration int32) *APICreateRateLimiterRequest {
	r.penaltyBoxDuration = &penaltyBoxDuration
	return r
}
// Action The action to take when a rate limiter violation is detected.
func (r *APICreateRateLimiterRequest) Action(action string) *APICreateRateLimiterRequest {
	r.action = &action
	return r
}
// ResponseObjectName Name of existing response object. Required if &#x60;action&#x60; is &#x60;response_object&#x60;. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration.
func (r *APICreateRateLimiterRequest) ResponseObjectName(responseObjectName string) *APICreateRateLimiterRequest {
	r.responseObjectName = &responseObjectName
	return r
}
// LoggerType Name of the type of logging endpoint to be used when action is &#x60;log_only&#x60;. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries.
func (r *APICreateRateLimiterRequest) LoggerType(loggerType string) *APICreateRateLimiterRequest {
	r.loggerType = &loggerType
	return r
}
// FeatureRevision Revision number of the rate limiting feature implementation. Defaults to the most recent revision.
func (r *APICreateRateLimiterRequest) FeatureRevision(featureRevision int32) *APICreateRateLimiterRequest {
	r.featureRevision = &featureRevision
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateRateLimiterRequest) Execute() (*RateLimiterResponse, *http.Response, error) {
	return r.APIService.CreateRateLimiterExecute(r)
}

/*
CreateRateLimiter Create a rate limiter

Create a rate limiter for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateRateLimiterRequest
*/
func (a *RateLimiterAPIService) CreateRateLimiter(ctx context.Context, serviceID string, versionID int32) APICreateRateLimiterRequest {
	return APICreateRateLimiterRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateRateLimiterExecute executes the request
//  @return RateLimiterResponse
func (a *RateLimiterAPIService) CreateRateLimiterExecute(r APICreateRateLimiterRequest) (*RateLimiterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *RateLimiterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimiterAPIService.CreateRateLimiter")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/rate-limiters"
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
	if r.uriDictionaryName != nil {
		localVarFormParams.Add("uri_dictionary_name", parameterToString(*r.uriDictionaryName, ""))
	}
	if r.httpMethods != nil {
		localVarFormParams.Add("http_methods", parameterToString(*r.httpMethods, "csv"))
	}
	if r.rpsLimit != nil {
		localVarFormParams.Add("rps_limit", parameterToString(*r.rpsLimit, ""))
	}
	if r.windowSize != nil {
		localVarFormParams.Add("window_size", parameterToString(*r.windowSize, ""))
	}
	if r.clientKey != nil {
		localVarFormParams.Add("client_key", parameterToString(*r.clientKey, "csv"))
	}
	if r.penaltyBoxDuration != nil {
		localVarFormParams.Add("penalty_box_duration", parameterToString(*r.penaltyBoxDuration, ""))
	}
	if r.action != nil {
		localVarFormParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.responseObjectName != nil {
		localVarFormParams.Add("response_object_name", parameterToString(*r.responseObjectName, ""))
	}
	if r.loggerType != nil {
		localVarFormParams.Add("logger_type", parameterToString(*r.loggerType, ""))
	}
	if r.featureRevision != nil {
		localVarFormParams.Add("feature_revision", parameterToString(*r.featureRevision, ""))
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

// APIDeleteRateLimiterRequest represents a request for the resource.
type APIDeleteRateLimiterRequest struct {
	ctx context.Context
	APIService RateLimiterAPI
	rateLimiterID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteRateLimiterRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteRateLimiterExecute(r)
}

/*
DeleteRateLimiter Delete a rate limiter

Delete a rate limiter by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rateLimiterID Alphanumeric string identifying the rate limiter.
 @return APIDeleteRateLimiterRequest
*/
func (a *RateLimiterAPIService) DeleteRateLimiter(ctx context.Context, rateLimiterID string) APIDeleteRateLimiterRequest {
	return APIDeleteRateLimiterRequest{
		APIService: a,
		ctx: ctx,
		rateLimiterID: rateLimiterID,
	}
}

// DeleteRateLimiterExecute executes the request
//  @return InlineResponse200
func (a *RateLimiterAPIService) DeleteRateLimiterExecute(r APIDeleteRateLimiterRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimiterAPIService.DeleteRateLimiter")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rate-limiters/{rate_limiter_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rate_limiter_id"+"}", gourl.PathEscape(parameterToString(r.rateLimiterID, "")))

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

// APIGetRateLimiterRequest represents a request for the resource.
type APIGetRateLimiterRequest struct {
	ctx context.Context
	APIService RateLimiterAPI
	rateLimiterID string
}


// Execute calls the API using the request data configured.
func (r APIGetRateLimiterRequest) Execute() (*RateLimiterResponse, *http.Response, error) {
	return r.APIService.GetRateLimiterExecute(r)
}

/*
GetRateLimiter Get a rate limiter

Get a rate limiter by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rateLimiterID Alphanumeric string identifying the rate limiter.
 @return APIGetRateLimiterRequest
*/
func (a *RateLimiterAPIService) GetRateLimiter(ctx context.Context, rateLimiterID string) APIGetRateLimiterRequest {
	return APIGetRateLimiterRequest{
		APIService: a,
		ctx: ctx,
		rateLimiterID: rateLimiterID,
	}
}

// GetRateLimiterExecute executes the request
//  @return RateLimiterResponse
func (a *RateLimiterAPIService) GetRateLimiterExecute(r APIGetRateLimiterRequest) (*RateLimiterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *RateLimiterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimiterAPIService.GetRateLimiter")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rate-limiters/{rate_limiter_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rate_limiter_id"+"}", gourl.PathEscape(parameterToString(r.rateLimiterID, "")))

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

// APIListRateLimitersRequest represents a request for the resource.
type APIListRateLimitersRequest struct {
	ctx context.Context
	APIService RateLimiterAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListRateLimitersRequest) Execute() ([]RateLimiterResponse, *http.Response, error) {
	return r.APIService.ListRateLimitersExecute(r)
}

/*
ListRateLimiters List rate limiters

List all rate limiters for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListRateLimitersRequest
*/
func (a *RateLimiterAPIService) ListRateLimiters(ctx context.Context, serviceID string, versionID int32) APIListRateLimitersRequest {
	return APIListRateLimitersRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListRateLimitersExecute executes the request
//  @return []RateLimiterResponse
func (a *RateLimiterAPIService) ListRateLimitersExecute(r APIListRateLimitersRequest) ([]RateLimiterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []RateLimiterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimiterAPIService.ListRateLimiters")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/rate-limiters"
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

// APIUpdateRateLimiterRequest represents a request for the resource.
type APIUpdateRateLimiterRequest struct {
	ctx context.Context
	APIService RateLimiterAPI
	rateLimiterID string
	name *string
	uriDictionaryName *string
	httpMethods *[]string
	rpsLimit *int32
	windowSize *int32
	clientKey *[]string
	penaltyBoxDuration *int32
	action *string
	responseObjectName *string
	loggerType *string
	featureRevision *int32
}

// Name A human readable name for the rate limiting rule.
func (r *APIUpdateRateLimiterRequest) Name(name string) *APIUpdateRateLimiterRequest {
	r.name = &name
	return r
}
// URIDictionaryName The name of an Edge Dictionary containing URIs as keys. If not defined or &#x60;null&#x60;, all origin URIs will be rate limited.
func (r *APIUpdateRateLimiterRequest) URIDictionaryName(uriDictionaryName string) *APIUpdateRateLimiterRequest {
	r.uriDictionaryName = &uriDictionaryName
	return r
}
// HTTPMethods Array of HTTP methods to apply rate limiting to.
func (r *APIUpdateRateLimiterRequest) HTTPMethods(httpMethods []string) *APIUpdateRateLimiterRequest {
	r.httpMethods = &httpMethods
	return r
}
// RpsLimit Upper limit of requests per second allowed by the rate limiter.
func (r *APIUpdateRateLimiterRequest) RpsLimit(rpsLimit int32) *APIUpdateRateLimiterRequest {
	r.rpsLimit = &rpsLimit
	return r
}
// WindowSize Number of seconds during which the RPS limit must be exceeded in order to trigger a violation.
func (r *APIUpdateRateLimiterRequest) WindowSize(windowSize int32) *APIUpdateRateLimiterRequest {
	r.windowSize = &windowSize
	return r
}
// ClientKey Array of VCL variables used to generate a counter key to identify a client. Example variables include &#x60;req.http.Fastly-Client-IP&#x60;, &#x60;req.http.User-Agent&#x60;, or a custom header like &#x60;req.http.API-Key&#x60;.
func (r *APIUpdateRateLimiterRequest) ClientKey(clientKey []string) *APIUpdateRateLimiterRequest {
	r.clientKey = &clientKey
	return r
}
// PenaltyBoxDuration Length of time in minutes that the rate limiter is in effect after the initial violation is detected.
func (r *APIUpdateRateLimiterRequest) PenaltyBoxDuration(penaltyBoxDuration int32) *APIUpdateRateLimiterRequest {
	r.penaltyBoxDuration = &penaltyBoxDuration
	return r
}
// Action The action to take when a rate limiter violation is detected.
func (r *APIUpdateRateLimiterRequest) Action(action string) *APIUpdateRateLimiterRequest {
	r.action = &action
	return r
}
// ResponseObjectName Name of existing response object. Required if &#x60;action&#x60; is &#x60;response_object&#x60;. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration.
func (r *APIUpdateRateLimiterRequest) ResponseObjectName(responseObjectName string) *APIUpdateRateLimiterRequest {
	r.responseObjectName = &responseObjectName
	return r
}
// LoggerType Name of the type of logging endpoint to be used when action is &#x60;log_only&#x60;. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries.
func (r *APIUpdateRateLimiterRequest) LoggerType(loggerType string) *APIUpdateRateLimiterRequest {
	r.loggerType = &loggerType
	return r
}
// FeatureRevision Revision number of the rate limiting feature implementation. Defaults to the most recent revision.
func (r *APIUpdateRateLimiterRequest) FeatureRevision(featureRevision int32) *APIUpdateRateLimiterRequest {
	r.featureRevision = &featureRevision
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateRateLimiterRequest) Execute() (*RateLimiterResponse, *http.Response, error) {
	return r.APIService.UpdateRateLimiterExecute(r)
}

/*
UpdateRateLimiter Update a rate limiter

Update a rate limiter by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rateLimiterID Alphanumeric string identifying the rate limiter.
 @return APIUpdateRateLimiterRequest
*/
func (a *RateLimiterAPIService) UpdateRateLimiter(ctx context.Context, rateLimiterID string) APIUpdateRateLimiterRequest {
	return APIUpdateRateLimiterRequest{
		APIService: a,
		ctx: ctx,
		rateLimiterID: rateLimiterID,
	}
}

// UpdateRateLimiterExecute executes the request
//  @return RateLimiterResponse
func (a *RateLimiterAPIService) UpdateRateLimiterExecute(r APIUpdateRateLimiterRequest) (*RateLimiterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *RateLimiterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimiterAPIService.UpdateRateLimiter")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rate-limiters/{rate_limiter_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rate_limiter_id"+"}", gourl.PathEscape(parameterToString(r.rateLimiterID, "")))

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
	if r.uriDictionaryName != nil {
		localVarFormParams.Add("uri_dictionary_name", parameterToString(*r.uriDictionaryName, ""))
	}
	if r.httpMethods != nil {
		localVarFormParams.Add("http_methods", parameterToString(*r.httpMethods, "csv"))
	}
	if r.rpsLimit != nil {
		localVarFormParams.Add("rps_limit", parameterToString(*r.rpsLimit, ""))
	}
	if r.windowSize != nil {
		localVarFormParams.Add("window_size", parameterToString(*r.windowSize, ""))
	}
	if r.clientKey != nil {
		localVarFormParams.Add("client_key", parameterToString(*r.clientKey, "csv"))
	}
	if r.penaltyBoxDuration != nil {
		localVarFormParams.Add("penalty_box_duration", parameterToString(*r.penaltyBoxDuration, ""))
	}
	if r.action != nil {
		localVarFormParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.responseObjectName != nil {
		localVarFormParams.Add("response_object_name", parameterToString(*r.responseObjectName, ""))
	}
	if r.loggerType != nil {
		localVarFormParams.Add("logger_type", parameterToString(*r.loggerType, ""))
	}
	if r.featureRevision != nil {
		localVarFormParams.Add("feature_revision", parameterToString(*r.featureRevision, ""))
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
