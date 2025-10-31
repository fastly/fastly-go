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

// UserAPI defines an interface for interacting with the resource.
type UserAPI interface {

	/*
		CreateUser Create a user

		Create a user.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateUserRequest

		Deprecated
	*/
	CreateUser(ctx context.Context) APICreateUserRequest

	// CreateUserExecute executes the request
	//  @return UserResponse
	// Deprecated
	CreateUserExecute(r APICreateUserRequest) (*UserResponse, *http.Response, error)

	/*
		DeleteUser Delete a user

		Delete a user.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param userId Alphanumeric string identifying the user.
		 @return APIDeleteUserRequest
	*/
	DeleteUser(ctx context.Context, userId string) APIDeleteUserRequest

	// DeleteUserExecute executes the request
	//  @return InlineResponse200
	DeleteUserExecute(r APIDeleteUserRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetCurrentUser Get the current user

		Get the logged in user.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetCurrentUserRequest
	*/
	GetCurrentUser(ctx context.Context) APIGetCurrentUserRequest

	// GetCurrentUserExecute executes the request
	//  @return UserResponse
	GetCurrentUserExecute(r APIGetCurrentUserRequest) (*UserResponse, *http.Response, error)

	/*
		GetUser Get a user

		Get a specific user.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param userId Alphanumeric string identifying the user.
		 @return APIGetUserRequest
	*/
	GetUser(ctx context.Context, userId string) APIGetUserRequest

	// GetUserExecute executes the request
	//  @return UserResponse
	GetUserExecute(r APIGetUserRequest) (*UserResponse, *http.Response, error)

	/*
		RequestPasswordReset Request a password reset

		Requests a password reset for the specified user.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param userLogin The login associated with the user (typically, an email address).
		 @return APIRequestPasswordResetRequest
	*/
	RequestPasswordReset(ctx context.Context, userLogin string) APIRequestPasswordResetRequest

	// RequestPasswordResetExecute executes the request
	//  @return InlineResponse200
	RequestPasswordResetExecute(r APIRequestPasswordResetRequest) (*InlineResponse200, *http.Response, error)

	/*
		UpdateUser Update a user

		Update a user. Only users with the role of `superuser` can make changes to other users on the account. Non-superusers may use this endpoint to make changes to their own account. Two-factor attributes are not editable via this endpoint.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param userId Alphanumeric string identifying the user.
		 @return APIUpdateUserRequest
	*/
	UpdateUser(ctx context.Context, userId string) APIUpdateUserRequest

	// UpdateUserExecute executes the request
	//  @return UserResponse
	UpdateUserExecute(r APIUpdateUserRequest) (*UserResponse, *http.Response, error)

	/*
		UpdateUserPassword Update the user's password

		Update the user's password to a new one.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIUpdateUserPasswordRequest
	*/
	UpdateUserPassword(ctx context.Context) APIUpdateUserPasswordRequest

	// UpdateUserPasswordExecute executes the request
	//  @return UserResponse
	UpdateUserPasswordExecute(r APIUpdateUserPasswordRequest) (*UserResponse, *http.Response, error)
}

// UserAPIService UserAPI service
type UserAPIService service

// APICreateUserRequest represents a request for the resource.
type APICreateUserRequest struct {
	ctx                    context.Context
	APIService             UserAPI
	login                  *string
	name                   *string
	limitServices          *bool
	locked                 *bool
	requireNewPassword     *bool
	role                   *RoleUser
	roles                  *[]string
	twoFactorAuthEnabled   *bool
	twoFactorSetupRequired *bool
}

// Login returns a pointer to a request.
func (r *APICreateUserRequest) Login(login string) *APICreateUserRequest {
	r.login = &login
	return r
}

// Name The real life name of the user.
func (r *APICreateUserRequest) Name(name string) *APICreateUserRequest {
	r.name = &name
	return r
}

// LimitServices Indicates that the user has limited access to the customer&#39;s services.
func (r *APICreateUserRequest) LimitServices(limitServices bool) *APICreateUserRequest {
	r.limitServices = &limitServices
	return r
}

// Locked Indicates whether the is account is locked for editing or not.
func (r *APICreateUserRequest) Locked(locked bool) *APICreateUserRequest {
	r.locked = &locked
	return r
}

// RequireNewPassword Indicates if a new password is required at next login.
func (r *APICreateUserRequest) RequireNewPassword(requireNewPassword bool) *APICreateUserRequest {
	r.requireNewPassword = &requireNewPassword
	return r
}

// Role returns a pointer to a request.
func (r *APICreateUserRequest) Role(role RoleUser) *APICreateUserRequest {
	r.role = &role
	return r
}

// Roles A list of role IDs assigned to the user.
func (r *APICreateUserRequest) Roles(roles []string) *APICreateUserRequest {
	r.roles = &roles
	return r
}

// TwoFactorAuthEnabled Indicates if 2FA is enabled on the user.
func (r *APICreateUserRequest) TwoFactorAuthEnabled(twoFactorAuthEnabled bool) *APICreateUserRequest {
	r.twoFactorAuthEnabled = &twoFactorAuthEnabled
	return r
}

// TwoFactorSetupRequired Indicates if 2FA is required by the user&#39;s customer account.
func (r *APICreateUserRequest) TwoFactorSetupRequired(twoFactorSetupRequired bool) *APICreateUserRequest {
	r.twoFactorSetupRequired = &twoFactorSetupRequired
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateUserRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.APIService.CreateUserExecute(r)
}

/*
CreateUser Create a user

Create a user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateUserRequest

Deprecated
*/
func (a *UserAPIService) CreateUser(ctx context.Context) APICreateUserRequest {
	return APICreateUserRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateUserExecute executes the request
//  @return UserResponse
// Deprecated
func (a *UserAPIService) CreateUserExecute(r APICreateUserRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.CreateUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user"

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
	if r.login != nil {
		paramJson, err := parameterToJSON(*r.login)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("login", paramJson)
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.limitServices != nil {
		localVarFormParams.Add("limit_services", parameterToString(*r.limitServices, ""))
	}
	if r.locked != nil {
		localVarFormParams.Add("locked", parameterToString(*r.locked, ""))
	}
	if r.requireNewPassword != nil {
		localVarFormParams.Add("require_new_password", parameterToString(*r.requireNewPassword, ""))
	}
	if r.role != nil {
		localVarFormParams.Add("role", parameterToString(*r.role, ""))
	}
	if r.roles != nil {
		localVarFormParams.Add("roles", parameterToString(*r.roles, "csv"))
	}
	if r.twoFactorAuthEnabled != nil {
		localVarFormParams.Add("two_factor_auth_enabled", parameterToString(*r.twoFactorAuthEnabled, ""))
	}
	if r.twoFactorSetupRequired != nil {
		localVarFormParams.Add("two_factor_setup_required", parameterToString(*r.twoFactorSetupRequired, ""))
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

// APIDeleteUserRequest represents a request for the resource.
type APIDeleteUserRequest struct {
	ctx        context.Context
	APIService UserAPI
	userId     string
}

// Execute calls the API using the request data configured.
func (r APIDeleteUserRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteUserExecute(r)
}

/*
DeleteUser Delete a user

Delete a user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId Alphanumeric string identifying the user.
 @return APIDeleteUserRequest
*/
func (a *UserAPIService) DeleteUser(ctx context.Context, userId string) APIDeleteUserRequest {
	return APIDeleteUserRequest{
		APIService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// DeleteUserExecute executes the request
//  @return InlineResponse200
func (a *UserAPIService) DeleteUserExecute(r APIDeleteUserRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.DeleteUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_id"+"}", gourl.PathEscape(parameterToString(r.userId, "")))

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

// APIGetCurrentUserRequest represents a request for the resource.
type APIGetCurrentUserRequest struct {
	ctx        context.Context
	APIService UserAPI
}

// Execute calls the API using the request data configured.
func (r APIGetCurrentUserRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.APIService.GetCurrentUserExecute(r)
}

/*
GetCurrentUser Get the current user

Get the logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetCurrentUserRequest
*/
func (a *UserAPIService) GetCurrentUser(ctx context.Context) APIGetCurrentUserRequest {
	return APIGetCurrentUserRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetCurrentUserExecute executes the request
//  @return UserResponse
func (a *UserAPIService) GetCurrentUserExecute(r APIGetCurrentUserRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.GetCurrentUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/current_user"

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

// APIGetUserRequest represents a request for the resource.
type APIGetUserRequest struct {
	ctx        context.Context
	APIService UserAPI
	userId     string
}

// Execute calls the API using the request data configured.
func (r APIGetUserRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.APIService.GetUserExecute(r)
}

/*
GetUser Get a user

Get a specific user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId Alphanumeric string identifying the user.
 @return APIGetUserRequest
*/
func (a *UserAPIService) GetUser(ctx context.Context, userId string) APIGetUserRequest {
	return APIGetUserRequest{
		APIService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// GetUserExecute executes the request
//  @return UserResponse
func (a *UserAPIService) GetUserExecute(r APIGetUserRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_id"+"}", gourl.PathEscape(parameterToString(r.userId, "")))

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

// APIRequestPasswordResetRequest represents a request for the resource.
type APIRequestPasswordResetRequest struct {
	ctx        context.Context
	APIService UserAPI
	userLogin  string
}

// Execute calls the API using the request data configured.
func (r APIRequestPasswordResetRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.RequestPasswordResetExecute(r)
}

/*
RequestPasswordReset Request a password reset

Requests a password reset for the specified user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userLogin The login associated with the user (typically, an email address).
 @return APIRequestPasswordResetRequest
*/
func (a *UserAPIService) RequestPasswordReset(ctx context.Context, userLogin string) APIRequestPasswordResetRequest {
	return APIRequestPasswordResetRequest{
		APIService: a,
		ctx:        ctx,
		userLogin:  userLogin,
	}
}

// RequestPasswordResetExecute executes the request
//  @return InlineResponse200
func (a *UserAPIService) RequestPasswordResetExecute(r APIRequestPasswordResetRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.RequestPasswordReset")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_login}/password/request_reset"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_login"+"}", gourl.PathEscape(parameterToString(r.userLogin, "")))

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

// APIUpdateUserRequest represents a request for the resource.
type APIUpdateUserRequest struct {
	ctx                    context.Context
	APIService             UserAPI
	userId                 string
	login                  *string
	name                   *string
	limitServices          *bool
	locked                 *bool
	requireNewPassword     *bool
	role                   *RoleUser
	roles                  *[]string
	twoFactorAuthEnabled   *bool
	twoFactorSetupRequired *bool
}

// Login returns a pointer to a request.
func (r *APIUpdateUserRequest) Login(login string) *APIUpdateUserRequest {
	r.login = &login
	return r
}

// Name The real life name of the user.
func (r *APIUpdateUserRequest) Name(name string) *APIUpdateUserRequest {
	r.name = &name
	return r
}

// LimitServices Indicates that the user has limited access to the customer&#39;s services.
func (r *APIUpdateUserRequest) LimitServices(limitServices bool) *APIUpdateUserRequest {
	r.limitServices = &limitServices
	return r
}

// Locked Indicates whether the is account is locked for editing or not.
func (r *APIUpdateUserRequest) Locked(locked bool) *APIUpdateUserRequest {
	r.locked = &locked
	return r
}

// RequireNewPassword Indicates if a new password is required at next login.
func (r *APIUpdateUserRequest) RequireNewPassword(requireNewPassword bool) *APIUpdateUserRequest {
	r.requireNewPassword = &requireNewPassword
	return r
}

// Role returns a pointer to a request.
func (r *APIUpdateUserRequest) Role(role RoleUser) *APIUpdateUserRequest {
	r.role = &role
	return r
}

// Roles A list of role IDs assigned to the user.
func (r *APIUpdateUserRequest) Roles(roles []string) *APIUpdateUserRequest {
	r.roles = &roles
	return r
}

// TwoFactorAuthEnabled Indicates if 2FA is enabled on the user.
func (r *APIUpdateUserRequest) TwoFactorAuthEnabled(twoFactorAuthEnabled bool) *APIUpdateUserRequest {
	r.twoFactorAuthEnabled = &twoFactorAuthEnabled
	return r
}

// TwoFactorSetupRequired Indicates if 2FA is required by the user&#39;s customer account.
func (r *APIUpdateUserRequest) TwoFactorSetupRequired(twoFactorSetupRequired bool) *APIUpdateUserRequest {
	r.twoFactorSetupRequired = &twoFactorSetupRequired
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateUserRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.APIService.UpdateUserExecute(r)
}

/*
UpdateUser Update a user

Update a user. Only users with the role of `superuser` can make changes to other users on the account. Non-superusers may use this endpoint to make changes to their own account. Two-factor attributes are not editable via this endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId Alphanumeric string identifying the user.
 @return APIUpdateUserRequest
*/
func (a *UserAPIService) UpdateUser(ctx context.Context, userId string) APIUpdateUserRequest {
	return APIUpdateUserRequest{
		APIService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// UpdateUserExecute executes the request
//  @return UserResponse
func (a *UserAPIService) UpdateUserExecute(r APIUpdateUserRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.UpdateUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"user_id"+"}", gourl.PathEscape(parameterToString(r.userId, "")))

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
	if r.login != nil {
		paramJson, err := parameterToJSON(*r.login)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("login", paramJson)
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.limitServices != nil {
		localVarFormParams.Add("limit_services", parameterToString(*r.limitServices, ""))
	}
	if r.locked != nil {
		localVarFormParams.Add("locked", parameterToString(*r.locked, ""))
	}
	if r.requireNewPassword != nil {
		localVarFormParams.Add("require_new_password", parameterToString(*r.requireNewPassword, ""))
	}
	if r.role != nil {
		localVarFormParams.Add("role", parameterToString(*r.role, ""))
	}
	if r.roles != nil {
		localVarFormParams.Add("roles", parameterToString(*r.roles, "csv"))
	}
	if r.twoFactorAuthEnabled != nil {
		localVarFormParams.Add("two_factor_auth_enabled", parameterToString(*r.twoFactorAuthEnabled, ""))
	}
	if r.twoFactorSetupRequired != nil {
		localVarFormParams.Add("two_factor_setup_required", parameterToString(*r.twoFactorSetupRequired, ""))
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

// APIUpdateUserPasswordRequest represents a request for the resource.
type APIUpdateUserPasswordRequest struct {
	ctx         context.Context
	APIService  UserAPI
	oldPassword *string
	newPassword *string
}

// OldPassword The user&#39;s current password.
func (r *APIUpdateUserPasswordRequest) OldPassword(oldPassword string) *APIUpdateUserPasswordRequest {
	r.oldPassword = &oldPassword
	return r
}

// NewPassword The user&#39;s new password.
func (r *APIUpdateUserPasswordRequest) NewPassword(newPassword string) *APIUpdateUserPasswordRequest {
	r.newPassword = &newPassword
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateUserPasswordRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.APIService.UpdateUserPasswordExecute(r)
}

/*
UpdateUserPassword Update the user's password

Update the user's password to a new one.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIUpdateUserPasswordRequest
*/
func (a *UserAPIService) UpdateUserPassword(ctx context.Context) APIUpdateUserPasswordRequest {
	return APIUpdateUserPasswordRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// UpdateUserPasswordExecute executes the request
//  @return UserResponse
func (a *UserAPIService) UpdateUserPasswordExecute(r APIUpdateUserPasswordRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.UpdateUserPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/current_user/password"

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
	if r.oldPassword != nil {
		localVarFormParams.Add("old_password", parameterToString(*r.oldPassword, ""))
	}
	if r.newPassword != nil {
		localVarFormParams.Add("new_password", parameterToString(*r.newPassword, ""))
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
