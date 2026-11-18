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

// HeaderAPI defines an interface for interacting with the resource.
type HeaderAPI interface {

	/*
		CreateHeaderObject Create a Header object

		Creates a new Header object.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateHeaderObjectRequest
	*/
	CreateHeaderObject(ctx context.Context, serviceID string, versionID int32) APICreateHeaderObjectRequest

	// CreateHeaderObjectExecute executes the request
	//  @return HeaderResponse
	CreateHeaderObjectExecute(r APICreateHeaderObjectRequest) (*HeaderResponse, *http.Response, error)

	/*
		DeleteHeaderObject Delete a Header object

		Deletes a Header object by name.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param headerName A handle to refer to this Header object.
		 @return APIDeleteHeaderObjectRequest
	*/
	DeleteHeaderObject(ctx context.Context, serviceID string, versionID int32, headerName string) APIDeleteHeaderObjectRequest

	// DeleteHeaderObjectExecute executes the request
	//  @return InlineResponse200
	DeleteHeaderObjectExecute(r APIDeleteHeaderObjectRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetHeaderObject Get a Header object

		Retrieves a Header object by name.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param headerName A handle to refer to this Header object.
		 @return APIGetHeaderObjectRequest
	*/
	GetHeaderObject(ctx context.Context, serviceID string, versionID int32, headerName string) APIGetHeaderObjectRequest

	// GetHeaderObjectExecute executes the request
	//  @return HeaderResponse
	GetHeaderObjectExecute(r APIGetHeaderObjectRequest) (*HeaderResponse, *http.Response, error)

	/*
		ListHeaderObjects List Header objects

		Retrieves all Header objects for a particular Version of a Service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListHeaderObjectsRequest
	*/
	ListHeaderObjects(ctx context.Context, serviceID string, versionID int32) APIListHeaderObjectsRequest

	// ListHeaderObjectsExecute executes the request
	//  @return []HeaderResponse
	ListHeaderObjectsExecute(r APIListHeaderObjectsRequest) ([]HeaderResponse, *http.Response, error)

	/*
		UpdateHeaderObject Update a Header object

		Modifies an existing Header object by name.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param headerName A handle to refer to this Header object.
		 @return APIUpdateHeaderObjectRequest
	*/
	UpdateHeaderObject(ctx context.Context, serviceID string, versionID int32, headerName string) APIUpdateHeaderObjectRequest

	// UpdateHeaderObjectExecute executes the request
	//  @return HeaderResponse
	UpdateHeaderObjectExecute(r APIUpdateHeaderObjectRequest) (*HeaderResponse, *http.Response, error)
}

// HeaderAPIService HeaderAPI service
type HeaderAPIService service

// APICreateHeaderObjectRequest represents a request for the resource.
type APICreateHeaderObjectRequest struct {
	ctx               context.Context
	APIService        HeaderAPI
	serviceID         string
	versionID         int32
	action            *string
	cacheCondition    *string
	dst               *string
	name              *string
	regex             *string
	requestCondition  *string
	responseCondition *string
	src               *string
	substitution      *string
	resourceType      *string
	ignoreIfSet       *string
	priority          *string
}

// Action Accepts a string value.
func (r *APICreateHeaderObjectRequest) Action(action string) *APICreateHeaderObjectRequest {
	r.action = &action
	return r
}

// CacheCondition Name of the cache condition controlling when this configuration applies.
func (r *APICreateHeaderObjectRequest) CacheCondition(cacheCondition string) *APICreateHeaderObjectRequest {
	r.cacheCondition = &cacheCondition
	return r
}

// Dst Header to set.
func (r *APICreateHeaderObjectRequest) Dst(dst string) *APICreateHeaderObjectRequest {
	r.dst = &dst
	return r
}

// Name A handle to refer to this Header object.
func (r *APICreateHeaderObjectRequest) Name(name string) *APICreateHeaderObjectRequest {
	r.name = &name
	return r
}

// Regex Regular expression to use. Only applies to &#x60;regex&#x60; and &#x60;regex_repeat&#x60; actions.
func (r *APICreateHeaderObjectRequest) Regex(regex string) *APICreateHeaderObjectRequest {
	r.regex = &regex
	return r
}

// RequestCondition Condition which, if met, will select this configuration during a request. Optional.
func (r *APICreateHeaderObjectRequest) RequestCondition(requestCondition string) *APICreateHeaderObjectRequest {
	r.requestCondition = &requestCondition
	return r
}

// ResponseCondition Optional name of a response condition to apply.
func (r *APICreateHeaderObjectRequest) ResponseCondition(responseCondition string) *APICreateHeaderObjectRequest {
	r.responseCondition = &responseCondition
	return r
}

// Src Variable to be used as a source for the header content. Does not apply to &#x60;delete&#x60; action.
func (r *APICreateHeaderObjectRequest) Src(src string) *APICreateHeaderObjectRequest {
	r.src = &src
	return r
}

// Substitution Value to substitute in place of regular expression. Only applies to &#x60;regex&#x60; and &#x60;regex_repeat&#x60; actions.
func (r *APICreateHeaderObjectRequest) Substitution(substitution string) *APICreateHeaderObjectRequest {
	r.substitution = &substitution
	return r
}

// ResourceType Accepts a string value.
func (r *APICreateHeaderObjectRequest) ResourceType(resourceType string) *APICreateHeaderObjectRequest {
	r.resourceType = &resourceType
	return r
}

// IgnoreIfSet Don&#39;t add the header if it is added already. Only applies to &#39;set&#39; action. Numerical value (\\\&quot;0\\\&quot; &#x3D; false, \\\&quot;1\\\&quot; &#x3D; true)
func (r *APICreateHeaderObjectRequest) IgnoreIfSet(ignoreIfSet string) *APICreateHeaderObjectRequest {
	r.ignoreIfSet = &ignoreIfSet
	return r
}

// Priority Priority determines execution order. Lower numbers execute first.
func (r *APICreateHeaderObjectRequest) Priority(priority string) *APICreateHeaderObjectRequest {
	r.priority = &priority
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateHeaderObjectRequest) Execute() (*HeaderResponse, *http.Response, error) {
	return r.APIService.CreateHeaderObjectExecute(r)
}

/*
CreateHeaderObject Create a Header object

Creates a new Header object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateHeaderObjectRequest
*/
func (a *HeaderAPIService) CreateHeaderObject(ctx context.Context, serviceID string, versionID int32) APICreateHeaderObjectRequest {
	return APICreateHeaderObjectRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateHeaderObjectExecute executes the request
//  @return HeaderResponse
func (a *HeaderAPIService) CreateHeaderObjectExecute(r APICreateHeaderObjectRequest) (*HeaderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HeaderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HeaderAPIService.CreateHeaderObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/header"
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
	if r.action != nil {
		localVarFormParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.cacheCondition != nil {
		localVarFormParams.Add("cache_condition", parameterToString(*r.cacheCondition, ""))
	}
	if r.dst != nil {
		localVarFormParams.Add("dst", parameterToString(*r.dst, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.regex != nil {
		localVarFormParams.Add("regex", parameterToString(*r.regex, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.responseCondition != nil {
		paramJSON, err := parameterToJSON(*r.responseCondition)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("response_condition", paramJSON)
	}
	if r.src != nil {
		localVarFormParams.Add("src", parameterToString(*r.src, ""))
	}
	if r.substitution != nil {
		localVarFormParams.Add("substitution", parameterToString(*r.substitution, ""))
	}
	if r.resourceType != nil {
		localVarFormParams.Add("type", parameterToString(*r.resourceType, ""))
	}
	if r.ignoreIfSet != nil {
		localVarFormParams.Add("ignore_if_set", parameterToString(*r.ignoreIfSet, ""))
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
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

// APIDeleteHeaderObjectRequest represents a request for the resource.
type APIDeleteHeaderObjectRequest struct {
	ctx        context.Context
	APIService HeaderAPI
	serviceID  string
	versionID  int32
	headerName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteHeaderObjectRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteHeaderObjectExecute(r)
}

/*
DeleteHeaderObject Delete a Header object

Deletes a Header object by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param headerName A handle to refer to this Header object.
 @return APIDeleteHeaderObjectRequest
*/
func (a *HeaderAPIService) DeleteHeaderObject(ctx context.Context, serviceID string, versionID int32, headerName string) APIDeleteHeaderObjectRequest {
	return APIDeleteHeaderObjectRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		headerName: headerName,
	}
}

// DeleteHeaderObjectExecute executes the request
//  @return InlineResponse200
func (a *HeaderAPIService) DeleteHeaderObjectExecute(r APIDeleteHeaderObjectRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HeaderAPIService.DeleteHeaderObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/header/{header_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"header_name"+"}", gourl.PathEscape(parameterToString(r.headerName, "")))

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

// APIGetHeaderObjectRequest represents a request for the resource.
type APIGetHeaderObjectRequest struct {
	ctx        context.Context
	APIService HeaderAPI
	serviceID  string
	versionID  int32
	headerName string
}

// Execute calls the API using the request data configured.
func (r APIGetHeaderObjectRequest) Execute() (*HeaderResponse, *http.Response, error) {
	return r.APIService.GetHeaderObjectExecute(r)
}

/*
GetHeaderObject Get a Header object

Retrieves a Header object by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param headerName A handle to refer to this Header object.
 @return APIGetHeaderObjectRequest
*/
func (a *HeaderAPIService) GetHeaderObject(ctx context.Context, serviceID string, versionID int32, headerName string) APIGetHeaderObjectRequest {
	return APIGetHeaderObjectRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		headerName: headerName,
	}
}

// GetHeaderObjectExecute executes the request
//  @return HeaderResponse
func (a *HeaderAPIService) GetHeaderObjectExecute(r APIGetHeaderObjectRequest) (*HeaderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HeaderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HeaderAPIService.GetHeaderObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/header/{header_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"header_name"+"}", gourl.PathEscape(parameterToString(r.headerName, "")))

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

// APIListHeaderObjectsRequest represents a request for the resource.
type APIListHeaderObjectsRequest struct {
	ctx        context.Context
	APIService HeaderAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListHeaderObjectsRequest) Execute() ([]HeaderResponse, *http.Response, error) {
	return r.APIService.ListHeaderObjectsExecute(r)
}

/*
ListHeaderObjects List Header objects

Retrieves all Header objects for a particular Version of a Service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListHeaderObjectsRequest
*/
func (a *HeaderAPIService) ListHeaderObjects(ctx context.Context, serviceID string, versionID int32) APIListHeaderObjectsRequest {
	return APIListHeaderObjectsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListHeaderObjectsExecute executes the request
//  @return []HeaderResponse
func (a *HeaderAPIService) ListHeaderObjectsExecute(r APIListHeaderObjectsRequest) ([]HeaderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []HeaderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HeaderAPIService.ListHeaderObjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/header"
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

// APIUpdateHeaderObjectRequest represents a request for the resource.
type APIUpdateHeaderObjectRequest struct {
	ctx               context.Context
	APIService        HeaderAPI
	serviceID         string
	versionID         int32
	headerName        string
	action            *string
	cacheCondition    *string
	dst               *string
	name              *string
	regex             *string
	requestCondition  *string
	responseCondition *string
	src               *string
	substitution      *string
	resourceType      *string
	ignoreIfSet       *string
	priority          *string
}

// Action Accepts a string value.
func (r *APIUpdateHeaderObjectRequest) Action(action string) *APIUpdateHeaderObjectRequest {
	r.action = &action
	return r
}

// CacheCondition Name of the cache condition controlling when this configuration applies.
func (r *APIUpdateHeaderObjectRequest) CacheCondition(cacheCondition string) *APIUpdateHeaderObjectRequest {
	r.cacheCondition = &cacheCondition
	return r
}

// Dst Header to set.
func (r *APIUpdateHeaderObjectRequest) Dst(dst string) *APIUpdateHeaderObjectRequest {
	r.dst = &dst
	return r
}

// Name A handle to refer to this Header object.
func (r *APIUpdateHeaderObjectRequest) Name(name string) *APIUpdateHeaderObjectRequest {
	r.name = &name
	return r
}

// Regex Regular expression to use. Only applies to &#x60;regex&#x60; and &#x60;regex_repeat&#x60; actions.
func (r *APIUpdateHeaderObjectRequest) Regex(regex string) *APIUpdateHeaderObjectRequest {
	r.regex = &regex
	return r
}

// RequestCondition Condition which, if met, will select this configuration during a request. Optional.
func (r *APIUpdateHeaderObjectRequest) RequestCondition(requestCondition string) *APIUpdateHeaderObjectRequest {
	r.requestCondition = &requestCondition
	return r
}

// ResponseCondition Optional name of a response condition to apply.
func (r *APIUpdateHeaderObjectRequest) ResponseCondition(responseCondition string) *APIUpdateHeaderObjectRequest {
	r.responseCondition = &responseCondition
	return r
}

// Src Variable to be used as a source for the header content. Does not apply to &#x60;delete&#x60; action.
func (r *APIUpdateHeaderObjectRequest) Src(src string) *APIUpdateHeaderObjectRequest {
	r.src = &src
	return r
}

// Substitution Value to substitute in place of regular expression. Only applies to &#x60;regex&#x60; and &#x60;regex_repeat&#x60; actions.
func (r *APIUpdateHeaderObjectRequest) Substitution(substitution string) *APIUpdateHeaderObjectRequest {
	r.substitution = &substitution
	return r
}

// ResourceType Accepts a string value.
func (r *APIUpdateHeaderObjectRequest) ResourceType(resourceType string) *APIUpdateHeaderObjectRequest {
	r.resourceType = &resourceType
	return r
}

// IgnoreIfSet Don&#39;t add the header if it is added already. Only applies to &#39;set&#39; action. Numerical value (\\\&quot;0\\\&quot; &#x3D; false, \\\&quot;1\\\&quot; &#x3D; true)
func (r *APIUpdateHeaderObjectRequest) IgnoreIfSet(ignoreIfSet string) *APIUpdateHeaderObjectRequest {
	r.ignoreIfSet = &ignoreIfSet
	return r
}

// Priority Priority determines execution order. Lower numbers execute first.
func (r *APIUpdateHeaderObjectRequest) Priority(priority string) *APIUpdateHeaderObjectRequest {
	r.priority = &priority
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateHeaderObjectRequest) Execute() (*HeaderResponse, *http.Response, error) {
	return r.APIService.UpdateHeaderObjectExecute(r)
}

/*
UpdateHeaderObject Update a Header object

Modifies an existing Header object by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param headerName A handle to refer to this Header object.
 @return APIUpdateHeaderObjectRequest
*/
func (a *HeaderAPIService) UpdateHeaderObject(ctx context.Context, serviceID string, versionID int32, headerName string) APIUpdateHeaderObjectRequest {
	return APIUpdateHeaderObjectRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
		headerName: headerName,
	}
}

// UpdateHeaderObjectExecute executes the request
//  @return HeaderResponse
func (a *HeaderAPIService) UpdateHeaderObjectExecute(r APIUpdateHeaderObjectRequest) (*HeaderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HeaderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HeaderAPIService.UpdateHeaderObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/header/{header_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"header_name"+"}", gourl.PathEscape(parameterToString(r.headerName, "")))

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
	if r.cacheCondition != nil {
		localVarFormParams.Add("cache_condition", parameterToString(*r.cacheCondition, ""))
	}
	if r.dst != nil {
		localVarFormParams.Add("dst", parameterToString(*r.dst, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.regex != nil {
		localVarFormParams.Add("regex", parameterToString(*r.regex, ""))
	}
	if r.requestCondition != nil {
		localVarFormParams.Add("request_condition", parameterToString(*r.requestCondition, ""))
	}
	if r.responseCondition != nil {
		paramJSON, err := parameterToJSON(*r.responseCondition)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("response_condition", paramJSON)
	}
	if r.src != nil {
		localVarFormParams.Add("src", parameterToString(*r.src, ""))
	}
	if r.substitution != nil {
		localVarFormParams.Add("substitution", parameterToString(*r.substitution, ""))
	}
	if r.resourceType != nil {
		localVarFormParams.Add("type", parameterToString(*r.resourceType, ""))
	}
	if r.ignoreIfSet != nil {
		localVarFormParams.Add("ignore_if_set", parameterToString(*r.ignoreIfSet, ""))
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
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
