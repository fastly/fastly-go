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

// SnippetAPI defines an interface for interacting with the resource.
type SnippetAPI interface {

	/*
		CreateSnippet Create a snippet

		Create a snippet for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateSnippetRequest
	*/
	CreateSnippet(ctx context.Context, serviceID string, versionID int32) APICreateSnippetRequest

	// CreateSnippetExecute executes the request
	//  @return SnippetResponse
	CreateSnippetExecute(r APICreateSnippetRequest) (*SnippetResponse, *http.Response, error)

	/*
		DeleteSnippet Delete a snippet

		Delete a specific snippet for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param name The name for the snippet.
		 @return APIDeleteSnippetRequest
	*/
	DeleteSnippet(ctx context.Context, serviceID string, versionID int32, name string) APIDeleteSnippetRequest

	// DeleteSnippetExecute executes the request
	//  @return InlineResponse200
	DeleteSnippetExecute(r APIDeleteSnippetRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetSnippet Get a versioned snippet

		Get a single snippet for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param name The name for the snippet.
		 @return APIGetSnippetRequest
	*/
	GetSnippet(ctx context.Context, serviceID string, versionID int32, name string) APIGetSnippetRequest

	// GetSnippetExecute executes the request
	//  @return SnippetResponse
	GetSnippetExecute(r APIGetSnippetRequest) (*SnippetResponse, *http.Response, error)

	/*
		GetSnippetDynamic Get a dynamic snippet

		Get a single dynamic snippet for a particular service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param id Alphanumeric string identifying a VCL Snippet.
		 @return APIGetSnippetDynamicRequest
	*/
	GetSnippetDynamic(ctx context.Context, serviceID string, id string) APIGetSnippetDynamicRequest

	// GetSnippetDynamicExecute executes the request
	//  @return SnippetResponse
	GetSnippetDynamicExecute(r APIGetSnippetDynamicRequest) (*SnippetResponse, *http.Response, error)

	/*
		ListSnippets List snippets

		List all snippets for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListSnippetsRequest
	*/
	ListSnippets(ctx context.Context, serviceID string, versionID int32) APIListSnippetsRequest

	// ListSnippetsExecute executes the request
	//  @return []SnippetResponse
	ListSnippetsExecute(r APIListSnippetsRequest) ([]SnippetResponse, *http.Response, error)

	/*
		UpdateSnippet Update a versioned snippet

		Update a specific snippet for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param name The name for the snippet.
		 @return APIUpdateSnippetRequest
	*/
	UpdateSnippet(ctx context.Context, serviceID string, versionID int32, name string) APIUpdateSnippetRequest

	// UpdateSnippetExecute executes the request
	//  @return SnippetResponse
	UpdateSnippetExecute(r APIUpdateSnippetRequest) (*SnippetResponse, *http.Response, error)

	/*
		UpdateSnippetDynamic Update a dynamic snippet

		Update a dynamic snippet for a particular service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param id Alphanumeric string identifying a VCL Snippet.
		 @return APIUpdateSnippetDynamicRequest
	*/
	UpdateSnippetDynamic(ctx context.Context, serviceID string, id string) APIUpdateSnippetDynamicRequest

	// UpdateSnippetDynamicExecute executes the request
	//  @return SnippetResponse
	UpdateSnippetDynamicExecute(r APIUpdateSnippetDynamicRequest) (*SnippetResponse, *http.Response, error)
}

// SnippetAPIService SnippetAPI service
type SnippetAPIService service

// APICreateSnippetRequest represents a request for the resource.
type APICreateSnippetRequest struct {
	ctx          context.Context
	APIService   SnippetAPI
	serviceID    string
	versionID    int32
	name         *string
	resourceType *string
	content      *string
	priority     *string
	dynamic      *string
}

// Name The name for the snippet.
func (r *APICreateSnippetRequest) Name(name string) *APICreateSnippetRequest {
	r.name = &name
	return r
}

// ResourceType The location in generated VCL where the snippet should be placed.
func (r *APICreateSnippetRequest) ResourceType(resourceType string) *APICreateSnippetRequest {
	r.resourceType = &resourceType
	return r
}

// Content The VCL code that specifies exactly what the snippet does.
func (r *APICreateSnippetRequest) Content(content string) *APICreateSnippetRequest {
	r.content = &content
	return r
}

// Priority Priority determines execution order. Lower numbers execute first.
func (r *APICreateSnippetRequest) Priority(priority string) *APICreateSnippetRequest {
	r.priority = &priority
	return r
}

// Dynamic Sets the snippet version.
func (r *APICreateSnippetRequest) Dynamic(dynamic string) *APICreateSnippetRequest {
	r.dynamic = &dynamic
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateSnippetRequest) Execute() (*SnippetResponse, *http.Response, error) {
	return r.APIService.CreateSnippetExecute(r)
}

/*
CreateSnippet Create a snippet

Create a snippet for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateSnippetRequest
*/
func (a *SnippetAPIService) CreateSnippet(ctx context.Context, serviceID string, versionID int32) APICreateSnippetRequest {
	return APICreateSnippetRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateSnippetExecute executes the request
//  @return SnippetResponse
func (a *SnippetAPIService) CreateSnippetExecute(r APICreateSnippetRequest) (*SnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SnippetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.CreateSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/snippet"
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
	if r.resourceType != nil {
		localVarFormParams.Add("type", parameterToString(*r.resourceType, ""))
	}
	if r.content != nil {
		localVarFormParams.Add("content", parameterToString(*r.content, ""))
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
	}
	if r.dynamic != nil {
		localVarFormParams.Add("dynamic", parameterToString(*r.dynamic, ""))
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

// APIDeleteSnippetRequest represents a request for the resource.
type APIDeleteSnippetRequest struct {
	ctx        context.Context
	APIService SnippetAPI
	serviceID  string
	versionID  int32
	name       string
}

// Execute calls the API using the request data configured.
func (r APIDeleteSnippetRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteSnippetExecute(r)
}

/*
DeleteSnippet Delete a snippet

Delete a specific snippet for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param name The name for the snippet.
 @return APIDeleteSnippetRequest
*/
func (a *SnippetAPIService) DeleteSnippet(ctx context.Context, serviceID string, versionID int32, name string) APIDeleteSnippetRequest {
	return APIDeleteSnippetRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		name:       name,
	}
}

// DeleteSnippetExecute executes the request
//  @return InlineResponse200
func (a *SnippetAPIService) DeleteSnippetExecute(r APIDeleteSnippetRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.DeleteSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/snippet/{name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"name"+"}", gourl.PathEscape(parameterToString(r.name, "")))

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

// APIGetSnippetRequest represents a request for the resource.
type APIGetSnippetRequest struct {
	ctx        context.Context
	APIService SnippetAPI
	serviceID  string
	versionID  int32
	name       string
}

// Execute calls the API using the request data configured.
func (r APIGetSnippetRequest) Execute() (*SnippetResponse, *http.Response, error) {
	return r.APIService.GetSnippetExecute(r)
}

/*
GetSnippet Get a versioned snippet

Get a single snippet for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param name The name for the snippet.
 @return APIGetSnippetRequest
*/
func (a *SnippetAPIService) GetSnippet(ctx context.Context, serviceID string, versionID int32, name string) APIGetSnippetRequest {
	return APIGetSnippetRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		name:       name,
	}
}

// GetSnippetExecute executes the request
//  @return SnippetResponse
func (a *SnippetAPIService) GetSnippetExecute(r APIGetSnippetRequest) (*SnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SnippetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.GetSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/snippet/{name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"name"+"}", gourl.PathEscape(parameterToString(r.name, "")))

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

// APIGetSnippetDynamicRequest represents a request for the resource.
type APIGetSnippetDynamicRequest struct {
	ctx        context.Context
	APIService SnippetAPI
	serviceID  string
	id         string
}

// Execute calls the API using the request data configured.
func (r APIGetSnippetDynamicRequest) Execute() (*SnippetResponse, *http.Response, error) {
	return r.APIService.GetSnippetDynamicExecute(r)
}

/*
GetSnippetDynamic Get a dynamic snippet

Get a single dynamic snippet for a particular service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param id Alphanumeric string identifying a VCL Snippet.
 @return APIGetSnippetDynamicRequest
*/
func (a *SnippetAPIService) GetSnippetDynamic(ctx context.Context, serviceID string, id string) APIGetSnippetDynamicRequest {
	return APIGetSnippetDynamicRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		id:         id,
	}
}

// GetSnippetDynamicExecute executes the request
//  @return SnippetResponse
func (a *SnippetAPIService) GetSnippetDynamicExecute(r APIGetSnippetDynamicRequest) (*SnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SnippetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.GetSnippetDynamic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/snippet/{id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"id"+"}", gourl.PathEscape(parameterToString(r.id, "")))

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

// APIListSnippetsRequest represents a request for the resource.
type APIListSnippetsRequest struct {
	ctx        context.Context
	APIService SnippetAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListSnippetsRequest) Execute() ([]SnippetResponse, *http.Response, error) {
	return r.APIService.ListSnippetsExecute(r)
}

/*
ListSnippets List snippets

List all snippets for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListSnippetsRequest
*/
func (a *SnippetAPIService) ListSnippets(ctx context.Context, serviceID string, versionID int32) APIListSnippetsRequest {
	return APIListSnippetsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListSnippetsExecute executes the request
//  @return []SnippetResponse
func (a *SnippetAPIService) ListSnippetsExecute(r APIListSnippetsRequest) ([]SnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []SnippetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.ListSnippets")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/snippet"
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

// APIUpdateSnippetRequest represents a request for the resource.
type APIUpdateSnippetRequest struct {
	ctx        context.Context
	APIService SnippetAPI
	serviceID  string
	versionID  int32
	name       string
}

// Execute calls the API using the request data configured.
func (r APIUpdateSnippetRequest) Execute() (*SnippetResponse, *http.Response, error) {
	return r.APIService.UpdateSnippetExecute(r)
}

/*
UpdateSnippet Update a versioned snippet

Update a specific snippet for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param name The name for the snippet.
 @return APIUpdateSnippetRequest
*/
func (a *SnippetAPIService) UpdateSnippet(ctx context.Context, serviceID string, versionID int32, name string) APIUpdateSnippetRequest {
	return APIUpdateSnippetRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		name:       name,
	}
}

// UpdateSnippetExecute executes the request
//  @return SnippetResponse
func (a *SnippetAPIService) UpdateSnippetExecute(r APIUpdateSnippetRequest) (*SnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SnippetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.UpdateSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/snippet/{name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"name"+"}", gourl.PathEscape(parameterToString(r.name, "")))

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

// APIUpdateSnippetDynamicRequest represents a request for the resource.
type APIUpdateSnippetDynamicRequest struct {
	ctx          context.Context
	APIService   SnippetAPI
	serviceID    string
	id           string
	name         *string
	resourceType *string
	content      *string
	priority     *string
	dynamic      *string
}

// Name The name for the snippet.
func (r *APIUpdateSnippetDynamicRequest) Name(name string) *APIUpdateSnippetDynamicRequest {
	r.name = &name
	return r
}

// ResourceType The location in generated VCL where the snippet should be placed.
func (r *APIUpdateSnippetDynamicRequest) ResourceType(resourceType string) *APIUpdateSnippetDynamicRequest {
	r.resourceType = &resourceType
	return r
}

// Content The VCL code that specifies exactly what the snippet does.
func (r *APIUpdateSnippetDynamicRequest) Content(content string) *APIUpdateSnippetDynamicRequest {
	r.content = &content
	return r
}

// Priority Priority determines execution order. Lower numbers execute first.
func (r *APIUpdateSnippetDynamicRequest) Priority(priority string) *APIUpdateSnippetDynamicRequest {
	r.priority = &priority
	return r
}

// Dynamic Sets the snippet version.
func (r *APIUpdateSnippetDynamicRequest) Dynamic(dynamic string) *APIUpdateSnippetDynamicRequest {
	r.dynamic = &dynamic
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateSnippetDynamicRequest) Execute() (*SnippetResponse, *http.Response, error) {
	return r.APIService.UpdateSnippetDynamicExecute(r)
}

/*
UpdateSnippetDynamic Update a dynamic snippet

Update a dynamic snippet for a particular service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param id Alphanumeric string identifying a VCL Snippet.
 @return APIUpdateSnippetDynamicRequest
*/
func (a *SnippetAPIService) UpdateSnippetDynamic(ctx context.Context, serviceID string, id string) APIUpdateSnippetDynamicRequest {
	return APIUpdateSnippetDynamicRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		id:         id,
	}
}

// UpdateSnippetDynamicExecute executes the request
//  @return SnippetResponse
func (a *SnippetAPIService) UpdateSnippetDynamicExecute(r APIUpdateSnippetDynamicRequest) (*SnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SnippetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnippetAPIService.UpdateSnippetDynamic")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/snippet/{id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"id"+"}", gourl.PathEscape(parameterToString(r.id, "")))

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
	if r.resourceType != nil {
		localVarFormParams.Add("type", parameterToString(*r.resourceType, ""))
	}
	if r.content != nil {
		localVarFormParams.Add("content", parameterToString(*r.content, ""))
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
	}
	if r.dynamic != nil {
		localVarFormParams.Add("dynamic", parameterToString(*r.dynamic, ""))
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
