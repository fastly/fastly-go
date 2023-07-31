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

// RequestSettingsAPI defines an interface for interacting with the resource.
type RequestSettingsAPI interface {

	/*
	CreateRequestSettings Create a Request Settings object

	Creates a new Request Settings object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateRequestSettingsRequest
	*/
	CreateRequestSettings(ctx context.Context, serviceID string, versionID int32) APICreateRequestSettingsRequest

	// CreateRequestSettingsExecute executes the request
	//  @return RequestSettingsResponse
	CreateRequestSettingsExecute(r APICreateRequestSettingsRequest) (*RequestSettingsResponse, *http.Response, error)

	/*
	DeleteRequestSettings Delete a Request Settings object

	Removes the specified Request Settings object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param requestSettingsName Name for the request settings.
	 @return APIDeleteRequestSettingsRequest
	*/
	DeleteRequestSettings(ctx context.Context, serviceID string, versionID int32, requestSettingsName string) APIDeleteRequestSettingsRequest

	// DeleteRequestSettingsExecute executes the request
	//  @return InlineResponse200
	DeleteRequestSettingsExecute(r APIDeleteRequestSettingsRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetRequestSettings Get a Request Settings object

	Gets the specified Request Settings object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param requestSettingsName Name for the request settings.
	 @return APIGetRequestSettingsRequest
	*/
	GetRequestSettings(ctx context.Context, serviceID string, versionID int32, requestSettingsName string) APIGetRequestSettingsRequest

	// GetRequestSettingsExecute executes the request
	//  @return RequestSettingsResponse
	GetRequestSettingsExecute(r APIGetRequestSettingsRequest) (*RequestSettingsResponse, *http.Response, error)

	/*
	ListRequestSettings List Request Settings objects

	Returns a list of all Request Settings objects for the given service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListRequestSettingsRequest
	*/
	ListRequestSettings(ctx context.Context, serviceID string, versionID int32) APIListRequestSettingsRequest

	// ListRequestSettingsExecute executes the request
	//  @return []RequestSettingsResponse
	ListRequestSettingsExecute(r APIListRequestSettingsRequest) ([]RequestSettingsResponse, *http.Response, error)

	/*
	UpdateRequestSettings Update a Request Settings object

	Updates the specified Request Settings object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param requestSettingsName Name for the request settings.
	 @return APIUpdateRequestSettingsRequest
	*/
	UpdateRequestSettings(ctx context.Context, serviceID string, versionID int32, requestSettingsName string) APIUpdateRequestSettingsRequest

	// UpdateRequestSettingsExecute executes the request
	//  @return RequestSettingsResponse
	UpdateRequestSettingsExecute(r APIUpdateRequestSettingsRequest) (*RequestSettingsResponse, *http.Response, error)
}

// RequestSettingsAPIService RequestSettingsAPI service
type RequestSettingsAPIService service

// APICreateRequestSettingsRequest represents a request for the resource.
type APICreateRequestSettingsRequest struct {
	ctx context.Context
	APIService RequestSettingsAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APICreateRequestSettingsRequest) Execute() (*RequestSettingsResponse, *http.Response, error) {
	return r.APIService.CreateRequestSettingsExecute(r)
}

/*
CreateRequestSettings Create a Request Settings object

Creates a new Request Settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateRequestSettingsRequest
*/
func (a *RequestSettingsAPIService) CreateRequestSettings(ctx context.Context, serviceID string, versionID int32) APICreateRequestSettingsRequest {
	return APICreateRequestSettingsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateRequestSettingsExecute executes the request
//  @return RequestSettingsResponse
func (a *RequestSettingsAPIService) CreateRequestSettingsExecute(r APICreateRequestSettingsRequest) (*RequestSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *RequestSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestSettingsAPIService.CreateRequestSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/request_settings"
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

// APIDeleteRequestSettingsRequest represents a request for the resource.
type APIDeleteRequestSettingsRequest struct {
	ctx context.Context
	APIService RequestSettingsAPI
	serviceID string
	versionID int32
	requestSettingsName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteRequestSettingsRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteRequestSettingsExecute(r)
}

/*
DeleteRequestSettings Delete a Request Settings object

Removes the specified Request Settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param requestSettingsName Name for the request settings.
 @return APIDeleteRequestSettingsRequest
*/
func (a *RequestSettingsAPIService) DeleteRequestSettings(ctx context.Context, serviceID string, versionID int32, requestSettingsName string) APIDeleteRequestSettingsRequest {
	return APIDeleteRequestSettingsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		requestSettingsName: requestSettingsName,
	}
}

// DeleteRequestSettingsExecute executes the request
//  @return InlineResponse200
func (a *RequestSettingsAPIService) DeleteRequestSettingsExecute(r APIDeleteRequestSettingsRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestSettingsAPIService.DeleteRequestSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"request_settings_name"+"}", gourl.PathEscape(parameterToString(r.requestSettingsName, "")))

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

// APIGetRequestSettingsRequest represents a request for the resource.
type APIGetRequestSettingsRequest struct {
	ctx context.Context
	APIService RequestSettingsAPI
	serviceID string
	versionID int32
	requestSettingsName string
}


// Execute calls the API using the request data configured.
func (r APIGetRequestSettingsRequest) Execute() (*RequestSettingsResponse, *http.Response, error) {
	return r.APIService.GetRequestSettingsExecute(r)
}

/*
GetRequestSettings Get a Request Settings object

Gets the specified Request Settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param requestSettingsName Name for the request settings.
 @return APIGetRequestSettingsRequest
*/
func (a *RequestSettingsAPIService) GetRequestSettings(ctx context.Context, serviceID string, versionID int32, requestSettingsName string) APIGetRequestSettingsRequest {
	return APIGetRequestSettingsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		requestSettingsName: requestSettingsName,
	}
}

// GetRequestSettingsExecute executes the request
//  @return RequestSettingsResponse
func (a *RequestSettingsAPIService) GetRequestSettingsExecute(r APIGetRequestSettingsRequest) (*RequestSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *RequestSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestSettingsAPIService.GetRequestSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"request_settings_name"+"}", gourl.PathEscape(parameterToString(r.requestSettingsName, "")))

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

// APIListRequestSettingsRequest represents a request for the resource.
type APIListRequestSettingsRequest struct {
	ctx context.Context
	APIService RequestSettingsAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListRequestSettingsRequest) Execute() ([]RequestSettingsResponse, *http.Response, error) {
	return r.APIService.ListRequestSettingsExecute(r)
}

/*
ListRequestSettings List Request Settings objects

Returns a list of all Request Settings objects for the given service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListRequestSettingsRequest
*/
func (a *RequestSettingsAPIService) ListRequestSettings(ctx context.Context, serviceID string, versionID int32) APIListRequestSettingsRequest {
	return APIListRequestSettingsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListRequestSettingsExecute executes the request
//  @return []RequestSettingsResponse
func (a *RequestSettingsAPIService) ListRequestSettingsExecute(r APIListRequestSettingsRequest) ([]RequestSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []RequestSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestSettingsAPIService.ListRequestSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/request_settings"
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

// APIUpdateRequestSettingsRequest represents a request for the resource.
type APIUpdateRequestSettingsRequest struct {
	ctx context.Context
	APIService RequestSettingsAPI
	serviceID string
	versionID int32
	requestSettingsName string
	action *string
	defaultHost *string
	hashKeys *string
	name *string
	requestCondition *string
	xff *string
	bypassBusyWait *int32
	forceMiss *int32
	forceSsl *int32
	geoHeaders *int32
	maxStaleAge *int32
	timerSupport *int32
}

// Action Allows you to terminate request handling and immediately perform an action.
func (r *APIUpdateRequestSettingsRequest) Action(action string) *APIUpdateRequestSettingsRequest {
	r.action = &action
	return r
}
// DefaultHost Sets the host header.
func (r *APIUpdateRequestSettingsRequest) DefaultHost(defaultHost string) *APIUpdateRequestSettingsRequest {
	r.defaultHost = &defaultHost
	return r
}
// HashKeys Comma separated list of varnish request object fields that should be in the hash key.
func (r *APIUpdateRequestSettingsRequest) HashKeys(hashKeys string) *APIUpdateRequestSettingsRequest {
	r.hashKeys = &hashKeys
	return r
}
// Name Name for the request settings.
func (r *APIUpdateRequestSettingsRequest) Name(name string) *APIUpdateRequestSettingsRequest {
	r.name = &name
	return r
}
// RequestCondition Condition which, if met, will select this configuration during a request. Optional.
func (r *APIUpdateRequestSettingsRequest) RequestCondition(requestCondition string) *APIUpdateRequestSettingsRequest {
	r.requestCondition = &requestCondition
	return r
}
// Xff Short for X-Forwarded-For.
func (r *APIUpdateRequestSettingsRequest) Xff(xff string) *APIUpdateRequestSettingsRequest {
	r.xff = &xff
	return r
}
// BypassBusyWait Disable collapsed forwarding, so you don&#39;t wait for other objects to origin.
func (r *APIUpdateRequestSettingsRequest) BypassBusyWait(bypassBusyWait int32) *APIUpdateRequestSettingsRequest {
	r.bypassBusyWait = &bypassBusyWait
	return r
}
// ForceMiss Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable.
func (r *APIUpdateRequestSettingsRequest) ForceMiss(forceMiss int32) *APIUpdateRequestSettingsRequest {
	r.forceMiss = &forceMiss
	return r
}
// ForceSsl Forces the request use SSL (redirects a non-SSL to SSL).
func (r *APIUpdateRequestSettingsRequest) ForceSsl(forceSsl int32) *APIUpdateRequestSettingsRequest {
	r.forceSsl = &forceSsl
	return r
}
// GeoHeaders Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.
func (r *APIUpdateRequestSettingsRequest) GeoHeaders(geoHeaders int32) *APIUpdateRequestSettingsRequest {
	r.geoHeaders = &geoHeaders
	return r
}
// MaxStaleAge How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.
func (r *APIUpdateRequestSettingsRequest) MaxStaleAge(maxStaleAge int32) *APIUpdateRequestSettingsRequest {
	r.maxStaleAge = &maxStaleAge
	return r
}
// TimerSupport Injects the X-Timer info into the request for viewing origin fetch durations.
func (r *APIUpdateRequestSettingsRequest) TimerSupport(timerSupport int32) *APIUpdateRequestSettingsRequest {
	r.timerSupport = &timerSupport
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateRequestSettingsRequest) Execute() (*RequestSettingsResponse, *http.Response, error) {
	return r.APIService.UpdateRequestSettingsExecute(r)
}

/*
UpdateRequestSettings Update a Request Settings object

Updates the specified Request Settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param requestSettingsName Name for the request settings.
 @return APIUpdateRequestSettingsRequest
*/
func (a *RequestSettingsAPIService) UpdateRequestSettings(ctx context.Context, serviceID string, versionID int32, requestSettingsName string) APIUpdateRequestSettingsRequest {
	return APIUpdateRequestSettingsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		requestSettingsName: requestSettingsName,
	}
}

// UpdateRequestSettingsExecute executes the request
//  @return RequestSettingsResponse
func (a *RequestSettingsAPIService) UpdateRequestSettingsExecute(r APIUpdateRequestSettingsRequest) (*RequestSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *RequestSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestSettingsAPIService.UpdateRequestSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"request_settings_name"+"}", gourl.PathEscape(parameterToString(r.requestSettingsName, "")))

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
	if r.action != nil {
		localVarFormParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.defaultHost != nil {
		localVarFormParams.Add("default_host", parameterToString(*r.defaultHost, ""))
	}
	if r.hashKeys != nil {
		localVarFormParams.Add("hash_keys", parameterToString(*r.hashKeys, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.xff != nil {
		localVarFormParams.Add("xff", parameterToString(*r.xff, ""))
	}
	if r.bypassBusyWait != nil {
		localVarFormParams.Add("bypass_busy_wait", parameterToString(*r.bypassBusyWait, ""))
	}
	if r.forceMiss != nil {
		localVarFormParams.Add("force_miss", parameterToString(*r.forceMiss, ""))
	}
	if r.forceSsl != nil {
		localVarFormParams.Add("force_ssl", parameterToString(*r.forceSsl, ""))
	}
	if r.geoHeaders != nil {
		localVarFormParams.Add("geo_headers", parameterToString(*r.geoHeaders, ""))
	}
	if r.maxStaleAge != nil {
		localVarFormParams.Add("max_stale_age", parameterToString(*r.maxStaleAge, ""))
	}
	if r.timerSupport != nil {
		localVarFormParams.Add("timer_support", parameterToString(*r.timerSupport, ""))
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
