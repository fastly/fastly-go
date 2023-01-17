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

// IamUserGroupsAPI defines an interface for interacting with the resource.
type IamUserGroupsAPI interface {

	/*
	DeleteAUserGroup Delete a user group

	Delete a user group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param userGroupID Alphanumeric string identifying the user group.
	 @return APIDeleteAUserGroupRequest
	*/
	DeleteAUserGroup(ctx context.Context, userGroupID string) APIDeleteAUserGroupRequest

	// DeleteAUserGroupExecute executes the request
	DeleteAUserGroupExecute(r APIDeleteAUserGroupRequest) (*http.Response, error)

	/*
	GetAUserGroup Get a user group

	Get a user group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param userGroupID Alphanumeric string identifying the user group.
	 @return APIGetAUserGroupRequest
	*/
	GetAUserGroup(ctx context.Context, userGroupID string) APIGetAUserGroupRequest

	// GetAUserGroupExecute executes the request
	//  @return map[string]any
	GetAUserGroupExecute(r APIGetAUserGroupRequest) (map[string]any, *http.Response, error)

	/*
	ListUserGroupMembers List members of a user group

	List members of a user group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param userGroupID Alphanumeric string identifying the user group.
	 @return APIListUserGroupMembersRequest
	*/
	ListUserGroupMembers(ctx context.Context, userGroupID string) APIListUserGroupMembersRequest

	// ListUserGroupMembersExecute executes the request
	//  @return map[string]any
	ListUserGroupMembersExecute(r APIListUserGroupMembersRequest) (map[string]any, *http.Response, error)

	/*
	ListUserGroupRoles List roles in a user group

	List roles in a user group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param userGroupID Alphanumeric string identifying the user group.
	 @return APIListUserGroupRolesRequest
	*/
	ListUserGroupRoles(ctx context.Context, userGroupID string) APIListUserGroupRolesRequest

	// ListUserGroupRolesExecute executes the request
	//  @return map[string]any
	ListUserGroupRolesExecute(r APIListUserGroupRolesRequest) (map[string]any, *http.Response, error)

	/*
	ListUserGroupServiceGroups List service groups in a user group

	List service groups in a user group.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param userGroupID Alphanumeric string identifying the user group.
	 @return APIListUserGroupServiceGroupsRequest
	*/
	ListUserGroupServiceGroups(ctx context.Context, userGroupID string) APIListUserGroupServiceGroupsRequest

	// ListUserGroupServiceGroupsExecute executes the request
	//  @return map[string]any
	ListUserGroupServiceGroupsExecute(r APIListUserGroupServiceGroupsRequest) (map[string]any, *http.Response, error)

	/*
	ListUserGroups List user groups

	List all user groups.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListUserGroupsRequest
	*/
	ListUserGroups(ctx context.Context) APIListUserGroupsRequest

	// ListUserGroupsExecute executes the request
	//  @return map[string]any
	ListUserGroupsExecute(r APIListUserGroupsRequest) (map[string]any, *http.Response, error)
}

// IamUserGroupsAPIService IamUserGroupsAPI service
type IamUserGroupsAPIService service

// APIDeleteAUserGroupRequest represents a request for the resource.
type APIDeleteAUserGroupRequest struct {
	ctx context.Context
	APIService IamUserGroupsAPI
	userGroupID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteAUserGroupRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteAUserGroupExecute(r)
}

/*
DeleteAUserGroup Delete a user group

Delete a user group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGroupID Alphanumeric string identifying the user group.
 @return APIDeleteAUserGroupRequest
*/
func (a *IamUserGroupsAPIService) DeleteAUserGroup(ctx context.Context, userGroupID string) APIDeleteAUserGroupRequest {
	return APIDeleteAUserGroupRequest{
		APIService: a,
		ctx: ctx,
		userGroupID: userGroupID,
	}
}

// DeleteAUserGroupExecute executes the request
func (a *IamUserGroupsAPIService) DeleteAUserGroupExecute(r APIDeleteAUserGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamUserGroupsAPIService.DeleteAUserGroup")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user-groups/{user_group_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_group_id"+"}", gourl.PathEscape(parameterToString(r.userGroupID, "")))

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

// APIGetAUserGroupRequest represents a request for the resource.
type APIGetAUserGroupRequest struct {
	ctx context.Context
	APIService IamUserGroupsAPI
	userGroupID string
}


// Execute calls the API using the request data configured.
func (r APIGetAUserGroupRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetAUserGroupExecute(r)
}

/*
GetAUserGroup Get a user group

Get a user group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGroupID Alphanumeric string identifying the user group.
 @return APIGetAUserGroupRequest
*/
func (a *IamUserGroupsAPIService) GetAUserGroup(ctx context.Context, userGroupID string) APIGetAUserGroupRequest {
	return APIGetAUserGroupRequest{
		APIService: a,
		ctx: ctx,
		userGroupID: userGroupID,
	}
}

// GetAUserGroupExecute executes the request
//  @return map[string]any
func (a *IamUserGroupsAPIService) GetAUserGroupExecute(r APIGetAUserGroupRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamUserGroupsAPIService.GetAUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user-groups/{user_group_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_group_id"+"}", gourl.PathEscape(parameterToString(r.userGroupID, "")))

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

// APIListUserGroupMembersRequest represents a request for the resource.
type APIListUserGroupMembersRequest struct {
	ctx context.Context
	APIService IamUserGroupsAPI
	userGroupID string
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListUserGroupMembersRequest) PerPage(perPage int32) *APIListUserGroupMembersRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListUserGroupMembersRequest) Page(page int32) *APIListUserGroupMembersRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListUserGroupMembersRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListUserGroupMembersExecute(r)
}

/*
ListUserGroupMembers List members of a user group

List members of a user group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGroupID Alphanumeric string identifying the user group.
 @return APIListUserGroupMembersRequest
*/
func (a *IamUserGroupsAPIService) ListUserGroupMembers(ctx context.Context, userGroupID string) APIListUserGroupMembersRequest {
	return APIListUserGroupMembersRequest{
		APIService: a,
		ctx: ctx,
		userGroupID: userGroupID,
	}
}

// ListUserGroupMembersExecute executes the request
//  @return map[string]any
func (a *IamUserGroupsAPIService) ListUserGroupMembersExecute(r APIListUserGroupMembersRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamUserGroupsAPIService.ListUserGroupMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user-groups/{user_group_id}/members"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_group_id"+"}", gourl.PathEscape(parameterToString(r.userGroupID, "")))

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

// APIListUserGroupRolesRequest represents a request for the resource.
type APIListUserGroupRolesRequest struct {
	ctx context.Context
	APIService IamUserGroupsAPI
	userGroupID string
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListUserGroupRolesRequest) PerPage(perPage int32) *APIListUserGroupRolesRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListUserGroupRolesRequest) Page(page int32) *APIListUserGroupRolesRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListUserGroupRolesRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListUserGroupRolesExecute(r)
}

/*
ListUserGroupRoles List roles in a user group

List roles in a user group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGroupID Alphanumeric string identifying the user group.
 @return APIListUserGroupRolesRequest
*/
func (a *IamUserGroupsAPIService) ListUserGroupRoles(ctx context.Context, userGroupID string) APIListUserGroupRolesRequest {
	return APIListUserGroupRolesRequest{
		APIService: a,
		ctx: ctx,
		userGroupID: userGroupID,
	}
}

// ListUserGroupRolesExecute executes the request
//  @return map[string]any
func (a *IamUserGroupsAPIService) ListUserGroupRolesExecute(r APIListUserGroupRolesRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamUserGroupsAPIService.ListUserGroupRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user-groups/{user_group_id}/roles"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_group_id"+"}", gourl.PathEscape(parameterToString(r.userGroupID, "")))

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

// APIListUserGroupServiceGroupsRequest represents a request for the resource.
type APIListUserGroupServiceGroupsRequest struct {
	ctx context.Context
	APIService IamUserGroupsAPI
	userGroupID string
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListUserGroupServiceGroupsRequest) PerPage(perPage int32) *APIListUserGroupServiceGroupsRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListUserGroupServiceGroupsRequest) Page(page int32) *APIListUserGroupServiceGroupsRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListUserGroupServiceGroupsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListUserGroupServiceGroupsExecute(r)
}

/*
ListUserGroupServiceGroups List service groups in a user group

List service groups in a user group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGroupID Alphanumeric string identifying the user group.
 @return APIListUserGroupServiceGroupsRequest
*/
func (a *IamUserGroupsAPIService) ListUserGroupServiceGroups(ctx context.Context, userGroupID string) APIListUserGroupServiceGroupsRequest {
	return APIListUserGroupServiceGroupsRequest{
		APIService: a,
		ctx: ctx,
		userGroupID: userGroupID,
	}
}

// ListUserGroupServiceGroupsExecute executes the request
//  @return map[string]any
func (a *IamUserGroupsAPIService) ListUserGroupServiceGroupsExecute(r APIListUserGroupServiceGroupsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamUserGroupsAPIService.ListUserGroupServiceGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user-groups/{user_group_id}/service-groups"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_group_id"+"}", gourl.PathEscape(parameterToString(r.userGroupID, "")))

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

// APIListUserGroupsRequest represents a request for the resource.
type APIListUserGroupsRequest struct {
	ctx context.Context
	APIService IamUserGroupsAPI
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListUserGroupsRequest) PerPage(perPage int32) *APIListUserGroupsRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListUserGroupsRequest) Page(page int32) *APIListUserGroupsRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListUserGroupsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListUserGroupsExecute(r)
}

/*
ListUserGroups List user groups

List all user groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListUserGroupsRequest
*/
func (a *IamUserGroupsAPIService) ListUserGroups(ctx context.Context) APIListUserGroupsRequest {
	return APIListUserGroupsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListUserGroupsExecute executes the request
//  @return map[string]any
func (a *IamUserGroupsAPIService) ListUserGroupsExecute(r APIListUserGroupsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamUserGroupsAPIService.ListUserGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user-groups"

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
