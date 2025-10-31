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
		IamV1RolesGet Get IAM role by ID

		Retrieve a single IAM role by its unique identifier.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param roleId Alphanumeric string identifying the role.
		 @return APIIamV1RolesGetRequest
	*/
	IamV1RolesGet(ctx context.Context, roleId string) APIIamV1RolesGetRequest

	// IamV1RolesGetExecute executes the request
	//  @return IamV1RoleResponse
	IamV1RolesGetExecute(r APIIamV1RolesGetRequest) (*IamV1RoleResponse, *http.Response, error)

	/*
		IamV1RolesList List IAM roles

		Retrieve a paginated list of IAM roles available in the account.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIIamV1RolesListRequest
	*/
	IamV1RolesList(ctx context.Context) APIIamV1RolesListRequest

	// IamV1RolesListExecute executes the request
	//  @return IamV1RoleListResponse
	IamV1RolesListExecute(r APIIamV1RolesListRequest) (*IamV1RoleListResponse, *http.Response, error)
}

// IamRolesAPIService IamRolesAPI service
type IamRolesAPIService service

// APIIamV1RolesGetRequest represents a request for the resource.
type APIIamV1RolesGetRequest struct {
	ctx        context.Context
	APIService IamRolesAPI
	roleId     string
	include    *string
}

// Include Include related data (i.e., permissions).
func (r *APIIamV1RolesGetRequest) Include(include string) *APIIamV1RolesGetRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIIamV1RolesGetRequest) Execute() (*IamV1RoleResponse, *http.Response, error) {
	return r.APIService.IamV1RolesGetExecute(r)
}

/*
IamV1RolesGet Get IAM role by ID

Retrieve a single IAM role by its unique identifier.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleId Alphanumeric string identifying the role.
 @return APIIamV1RolesGetRequest
*/
func (a *IamRolesAPIService) IamV1RolesGet(ctx context.Context, roleId string) APIIamV1RolesGetRequest {
	return APIIamV1RolesGetRequest{
		APIService: a,
		ctx:        ctx,
		roleId:     roleId,
	}
}

// IamV1RolesGetExecute executes the request
//  @return IamV1RoleResponse
func (a *IamRolesAPIService) IamV1RolesGetExecute(r APIIamV1RolesGetRequest) (*IamV1RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *IamV1RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.IamV1RolesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", gourl.PathEscape(parameterToString(r.roleId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
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

// APIIamV1RolesListRequest represents a request for the resource.
type APIIamV1RolesListRequest struct {
	ctx        context.Context
	APIService IamRolesAPI
	limit      *int32
	cursor     *string
}

// Limit Number of results per page. The maximum is 1000.
func (r *APIIamV1RolesListRequest) Limit(limit int32) *APIIamV1RolesListRequest {
	r.limit = &limit
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIIamV1RolesListRequest) Cursor(cursor string) *APIIamV1RolesListRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIIamV1RolesListRequest) Execute() (*IamV1RoleListResponse, *http.Response, error) {
	return r.APIService.IamV1RolesListExecute(r)
}

/*
IamV1RolesList List IAM roles

Retrieve a paginated list of IAM roles available in the account.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIIamV1RolesListRequest
*/
func (a *IamRolesAPIService) IamV1RolesList(ctx context.Context) APIIamV1RolesListRequest {
	return APIIamV1RolesListRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// IamV1RolesListExecute executes the request
//  @return IamV1RoleListResponse
func (a *IamRolesAPIService) IamV1RolesListExecute(r APIIamV1RolesListRequest) (*IamV1RoleListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *IamV1RoleListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamRolesAPIService.IamV1RolesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
