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

// IamRolesAPI defines an interface for interacting with the resource.
type IamRolesAPI interface {

	/*
	AddRolePermissions Add permissions to a role

	Add permissions to a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param roleID Alphanumeric string identifying the role.
	 @return APIAddRolePermissionsRequest
	*/
	AddRolePermissions(ctx context.Context, roleID string) APIAddRolePermissionsRequest

	// AddRolePermissionsExecute executes the request
	//  @return map[string]any
	AddRolePermissionsExecute(r APIAddRolePermissionsRequest) (map[string]any, *http.Response, error)

	/*
	CreateARole Create a role

	Create a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateARoleRequest
	*/
	CreateARole(ctx context.Context) APICreateARoleRequest

	// CreateARoleExecute executes the request
	//  @return map[string]any
	CreateARoleExecute(r APICreateARoleRequest) (map[string]any, *http.Response, error)

	/*
	DeleteARole Delete a role

	Delete a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param roleID Alphanumeric string identifying the role.
	 @return APIDeleteARoleRequest
	*/
	DeleteARole(ctx context.Context, roleID string) APIDeleteARoleRequest

	// DeleteARoleExecute executes the request
	DeleteARoleExecute(r APIDeleteARoleRequest) (*http.Response, error)

	/*
	GetARole Get a role

	Get a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param roleID Alphanumeric string identifying the role.
	 @return APIGetARoleRequest
	*/
	GetARole(ctx context.Context, roleID string) APIGetARoleRequest

	// GetARoleExecute executes the request
	//  @return map[string]any
	GetARoleExecute(r APIGetARoleRequest) (map[string]any, *http.Response, error)

	/*
	ListRolePermissions List permissions in a role

	List all permissions in a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param roleID Alphanumeric string identifying the role.
	 @return APIListRolePermissionsRequest
	*/
	ListRolePermissions(ctx context.Context, roleID string) APIListRolePermissionsRequest

	// ListRolePermissionsExecute executes the request
	//  @return map[string]any
	ListRolePermissionsExecute(r APIListRolePermissionsRequest) (map[string]any, *http.Response, error)

	/*
	ListRoles List roles

	List all roles.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListRolesRequest
	*/
	ListRoles(ctx context.Context) APIListRolesRequest

	// ListRolesExecute executes the request
	//  @return map[string]any
	ListRolesExecute(r APIListRolesRequest) (map[string]any, *http.Response, error)

	/*
	RemoveRolePermissions Remove permissions from a role

	Remove permissions from a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param roleID Alphanumeric string identifying the role.
	 @return APIRemoveRolePermissionsRequest
	*/
	RemoveRolePermissions(ctx context.Context, roleID string) APIRemoveRolePermissionsRequest

	// RemoveRolePermissionsExecute executes the request
	RemoveRolePermissionsExecute(r APIRemoveRolePermissionsRequest) (*http.Response, error)

	/*
	UpdateARole Update a role

	Update a role.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param roleID Alphanumeric string identifying the role.
	 @return APIUpdateARoleRequest
	*/
	UpdateARole(ctx context.Context, roleID string) APIUpdateARoleRequest

	// UpdateARoleExecute executes the request
	//  @return map[string]any
	UpdateARoleExecute(r APIUpdateARoleRequest) (map[string]any, *http.Response, error)
}

// IamRolesAPIService IamRolesAPI service
type IamRolesAPIService service

// APIAddRolePermissionsRequest represents a request for the resource.
type APIAddRolePermissionsRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	roleID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIAddRolePermissionsRequest) RequestBody(requestBody map[string]map[string]any) *APIAddRolePermissionsRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIAddRolePermissionsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.AddRolePermissionsExecute(r)
}

/*
AddRolePermissions Add permissions to a role

Add permissions to a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Alphanumeric string identifying the role.
 @return APIAddRolePermissionsRequest
*/
func (a *IamRolesAPIService) AddRolePermissions(ctx context.Context, roleID string) APIAddRolePermissionsRequest {
	return APIAddRolePermissionsRequest{
		APIService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// AddRolePermissionsExecute executes the request
//  @return map[string]any
func (a *IamRolesAPIService) AddRolePermissionsExecute(r APIAddRolePermissionsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.AddRolePermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles/{role_id}/permissions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
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

// APICreateARoleRequest represents a request for the resource.
type APICreateARoleRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APICreateARoleRequest) RequestBody(requestBody map[string]map[string]any) *APICreateARoleRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateARoleRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.CreateARoleExecute(r)
}

/*
CreateARole Create a role

Create a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateARoleRequest
*/
func (a *IamRolesAPIService) CreateARole(ctx context.Context) APICreateARoleRequest {
	return APICreateARoleRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateARoleExecute executes the request
//  @return map[string]any
func (a *IamRolesAPIService) CreateARoleExecute(r APICreateARoleRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.CreateARole")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
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

// APIDeleteARoleRequest represents a request for the resource.
type APIDeleteARoleRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	roleID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteARoleRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteARoleExecute(r)
}

/*
DeleteARole Delete a role

Delete a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Alphanumeric string identifying the role.
 @return APIDeleteARoleRequest
*/
func (a *IamRolesAPIService) DeleteARole(ctx context.Context, roleID string) APIDeleteARoleRequest {
	return APIDeleteARoleRequest{
		APIService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// DeleteARoleExecute executes the request
func (a *IamRolesAPIService) DeleteARoleExecute(r APIDeleteARoleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.DeleteARole")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleID, "")))

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

// APIGetARoleRequest represents a request for the resource.
type APIGetARoleRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	roleID string
}


// Execute calls the API using the request data configured.
func (r APIGetARoleRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetARoleExecute(r)
}

/*
GetARole Get a role

Get a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Alphanumeric string identifying the role.
 @return APIGetARoleRequest
*/
func (a *IamRolesAPIService) GetARole(ctx context.Context, roleID string) APIGetARoleRequest {
	return APIGetARoleRequest{
		APIService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// GetARoleExecute executes the request
//  @return map[string]any
func (a *IamRolesAPIService) GetARoleExecute(r APIGetARoleRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.GetARole")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleID, "")))

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

// APIListRolePermissionsRequest represents a request for the resource.
type APIListRolePermissionsRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	roleID string
}


// Execute calls the API using the request data configured.
func (r APIListRolePermissionsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListRolePermissionsExecute(r)
}

/*
ListRolePermissions List permissions in a role

List all permissions in a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Alphanumeric string identifying the role.
 @return APIListRolePermissionsRequest
*/
func (a *IamRolesAPIService) ListRolePermissions(ctx context.Context, roleID string) APIListRolePermissionsRequest {
	return APIListRolePermissionsRequest{
		APIService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// ListRolePermissionsExecute executes the request
//  @return map[string]any
func (a *IamRolesAPIService) ListRolePermissionsExecute(r APIListRolePermissionsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.ListRolePermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles/{role_id}/permissions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleID, "")))

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

// APIListRolesRequest represents a request for the resource.
type APIListRolesRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	perPage *int32
	page *int32
}

// PerPage Number of records per page.
func (r *APIListRolesRequest) PerPage(perPage int32) *APIListRolesRequest {
	r.perPage = &perPage
	return r
}
// Page Current page.
func (r *APIListRolesRequest) Page(page int32) *APIListRolesRequest {
	r.page = &page
	return r
}

// Execute calls the API using the request data configured.
func (r APIListRolesRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListRolesExecute(r)
}

/*
ListRoles List roles

List all roles.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListRolesRequest
*/
func (a *IamRolesAPIService) ListRoles(ctx context.Context) APIListRolesRequest {
	return APIListRolesRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListRolesExecute executes the request
//  @return map[string]any
func (a *IamRolesAPIService) ListRolesExecute(r APIListRolesRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.ListRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles"

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

// APIRemoveRolePermissionsRequest represents a request for the resource.
type APIRemoveRolePermissionsRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	roleID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIRemoveRolePermissionsRequest) RequestBody(requestBody map[string]map[string]any) *APIRemoveRolePermissionsRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIRemoveRolePermissionsRequest) Execute() (*http.Response, error) {
	return r.APIService.RemoveRolePermissionsExecute(r)
}

/*
RemoveRolePermissions Remove permissions from a role

Remove permissions from a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Alphanumeric string identifying the role.
 @return APIRemoveRolePermissionsRequest
*/
func (a *IamRolesAPIService) RemoveRolePermissions(ctx context.Context, roleID string) APIRemoveRolePermissionsRequest {
	return APIRemoveRolePermissionsRequest{
		APIService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// RemoveRolePermissionsExecute executes the request
func (a *IamRolesAPIService) RemoveRolePermissionsExecute(r APIRemoveRolePermissionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.RemoveRolePermissions")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles/{role_id}/permissions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
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

// APIUpdateARoleRequest represents a request for the resource.
type APIUpdateARoleRequest struct {
	ctx context.Context
	APIService IamRolesAPI
	roleID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIUpdateARoleRequest) RequestBody(requestBody map[string]map[string]any) *APIUpdateARoleRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateARoleRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.UpdateARoleExecute(r)
}

/*
UpdateARole Update a role

Update a role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Alphanumeric string identifying the role.
 @return APIUpdateARoleRequest
*/
func (a *IamRolesAPIService) UpdateARole(ctx context.Context, roleID string) APIUpdateARoleRequest {
	return APIUpdateARoleRequest{
		APIService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// UpdateARoleExecute executes the request
//  @return map[string]any
func (a *IamRolesAPIService) UpdateARoleExecute(r APIUpdateARoleRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.UpdateARole")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
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
