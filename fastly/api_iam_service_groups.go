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

// IamServiceGroupsAPI defines an interface for interacting with the resource.
type IamServiceGroupsAPI interface {

	/*
	DeleteAServiceGroup Delete a service group

	Delete a service group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceGroupID Alphanumeric string identifying the service group.
	 @return APIDeleteAServiceGroupRequest
	*/
	DeleteAServiceGroup(ctx context.Context, serviceGroupID string) APIDeleteAServiceGroupRequest

	// DeleteAServiceGroupExecute executes the request
	DeleteAServiceGroupExecute(r APIDeleteAServiceGroupRequest) (*http.Response, error)

	/*
	GetAServiceGroup Get a service group

	Get a service group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceGroupID Alphanumeric string identifying the service group.
	 @return APIGetAServiceGroupRequest
	*/
	GetAServiceGroup(ctx context.Context, serviceGroupID string) APIGetAServiceGroupRequest

	// GetAServiceGroupExecute executes the request
	//  @return map[string]any
	GetAServiceGroupExecute(r APIGetAServiceGroupRequest) (map[string]any, *http.Response, error)

	/*
	ListServiceGroupServices List services to a service group

	List services to a service group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceGroupID Alphanumeric string identifying the service group.
	 @return APIListServiceGroupServicesRequest
	*/
	ListServiceGroupServices(ctx context.Context, serviceGroupID string) APIListServiceGroupServicesRequest

	// ListServiceGroupServicesExecute executes the request
	//  @return map[string]any
	ListServiceGroupServicesExecute(r APIListServiceGroupServicesRequest) (map[string]any, *http.Response, error)

	/*
	ListServiceGroups List service groups

	List all service groups.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListServiceGroupsRequest
	*/
	ListServiceGroups(ctx context.Context) APIListServiceGroupsRequest

	// ListServiceGroupsExecute executes the request
	//  @return map[string]any
	ListServiceGroupsExecute(r APIListServiceGroupsRequest) (map[string]any, *http.Response, error)
}

// IamServiceGroupsAPIService IamServiceGroupsAPI service
type IamServiceGroupsAPIService service

// APIDeleteAServiceGroupRequest represents a request for the resource.
type APIDeleteAServiceGroupRequest struct {
	ctx context.Context
	APIService IamServiceGroupsAPI
	serviceGroupID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteAServiceGroupRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteAServiceGroupExecute(r)
}

/*
DeleteAServiceGroup Delete a service group

Delete a service group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceGroupID Alphanumeric string identifying the service group.
 @return APIDeleteAServiceGroupRequest
*/
func (a *IamServiceGroupsAPIService) DeleteAServiceGroup(ctx context.Context, serviceGroupID string) APIDeleteAServiceGroupRequest {
	return APIDeleteAServiceGroupRequest{
		APIService: a,
		ctx: ctx,
		serviceGroupID: serviceGroupID,
	}
}

// DeleteAServiceGroupExecute executes the request
func (a *IamServiceGroupsAPIService) DeleteAServiceGroupExecute(r APIDeleteAServiceGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamServiceGroupsAPIService.DeleteAServiceGroup")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-groups/{service_group_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_group_id"+"}", gourl.PathEscape(parameterToString(r.serviceGroupID, "")))

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
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

	return localVarHTTPResponse, nil
}

// APIGetAServiceGroupRequest represents a request for the resource.
type APIGetAServiceGroupRequest struct {
	ctx context.Context
	APIService IamServiceGroupsAPI
	serviceGroupID string
}


// Execute calls the API using the request data configured.
func (r APIGetAServiceGroupRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetAServiceGroupExecute(r)
}

/*
GetAServiceGroup Get a service group

Get a service group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceGroupID Alphanumeric string identifying the service group.
 @return APIGetAServiceGroupRequest
*/
func (a *IamServiceGroupsAPIService) GetAServiceGroup(ctx context.Context, serviceGroupID string) APIGetAServiceGroupRequest {
	return APIGetAServiceGroupRequest{
		APIService: a,
		ctx: ctx,
		serviceGroupID: serviceGroupID,
	}
}

// GetAServiceGroupExecute executes the request
//  @return map[string]any
func (a *IamServiceGroupsAPIService) GetAServiceGroupExecute(r APIGetAServiceGroupRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamServiceGroupsAPIService.GetAServiceGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-groups/{service_group_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_group_id"+"}", gourl.PathEscape(parameterToString(r.serviceGroupID, "")))

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

// APIListServiceGroupServicesRequest represents a request for the resource.
type APIListServiceGroupServicesRequest struct {
	ctx context.Context
	APIService IamServiceGroupsAPI
	serviceGroupID string
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListServiceGroupServicesRequest) PerPage(perPage int32) *APIListServiceGroupServicesRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListServiceGroupServicesRequest) Page(page int32) *APIListServiceGroupServicesRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListServiceGroupServicesRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListServiceGroupServicesExecute(r)
}

/*
ListServiceGroupServices List services to a service group

List services to a service group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceGroupID Alphanumeric string identifying the service group.
 @return APIListServiceGroupServicesRequest
*/
func (a *IamServiceGroupsAPIService) ListServiceGroupServices(ctx context.Context, serviceGroupID string) APIListServiceGroupServicesRequest {
	return APIListServiceGroupServicesRequest{
		APIService: a,
		ctx: ctx,
		serviceGroupID: serviceGroupID,
	}
}

// ListServiceGroupServicesExecute executes the request
//  @return map[string]any
func (a *IamServiceGroupsAPIService) ListServiceGroupServicesExecute(r APIListServiceGroupServicesRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamServiceGroupsAPIService.ListServiceGroupServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-groups/{service_group_id}/services"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_group_id"+"}", gourl.PathEscape(parameterToString(r.serviceGroupID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
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

// APIListServiceGroupsRequest represents a request for the resource.
type APIListServiceGroupsRequest struct {
	ctx context.Context
	APIService IamServiceGroupsAPI
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListServiceGroupsRequest) PerPage(perPage int32) *APIListServiceGroupsRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListServiceGroupsRequest) Page(page int32) *APIListServiceGroupsRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListServiceGroupsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListServiceGroupsExecute(r)
}

/*
ListServiceGroups List service groups

List all service groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListServiceGroupsRequest
*/
func (a *IamServiceGroupsAPIService) ListServiceGroups(ctx context.Context) APIListServiceGroupsRequest {
	return APIListServiceGroupsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListServiceGroupsExecute executes the request
//  @return map[string]any
func (a *IamServiceGroupsAPIService) ListServiceGroupsExecute(r APIListServiceGroupsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamServiceGroupsAPIService.ListServiceGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
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
