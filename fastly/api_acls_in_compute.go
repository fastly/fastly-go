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

// ACLsInComputeAPI defines an interface for interacting with the resource.
type ACLsInComputeAPI interface {

	/*
		ComputeACLCreateAcls Create a new ACL

		Create a new ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIComputeACLCreateAclsRequest
	*/
	ComputeACLCreateAcls(ctx context.Context) APIComputeACLCreateAclsRequest

	// ComputeACLCreateAclsExecute executes the request
	//  @return ComputeACLCreateAclsResponse
	ComputeACLCreateAclsExecute(r APIComputeACLCreateAclsRequest) (*ComputeACLCreateAclsResponse, *http.Response, error)

	/*
		ComputeACLDeleteSAclID Delete an ACL

		Delete an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclID
		 @return APIComputeACLDeleteSAclIDRequest
	*/
	ComputeACLDeleteSAclID(ctx context.Context, aclID string) APIComputeACLDeleteSAclIDRequest

	// ComputeACLDeleteSAclIDExecute executes the request
	ComputeACLDeleteSAclIDExecute(r APIComputeACLDeleteSAclIDRequest) (*http.Response, error)

	/*
		ComputeACLListAclEntries List an ACL

		List an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclID
		 @return APIComputeACLListAclEntriesRequest
	*/
	ComputeACLListAclEntries(ctx context.Context, aclID string) APIComputeACLListAclEntriesRequest

	// ComputeACLListAclEntriesExecute executes the request
	//  @return ComputeACLListEntries
	ComputeACLListAclEntriesExecute(r APIComputeACLListAclEntriesRequest) (*ComputeACLListEntries, *http.Response, error)

	/*
		ComputeACLListAcls List ACLs

		List all ACLs.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIComputeACLListAclsRequest
	*/
	ComputeACLListAcls(ctx context.Context) APIComputeACLListAclsRequest

	// ComputeACLListAclsExecute executes the request
	//  @return ComputeACLList
	ComputeACLListAclsExecute(r APIComputeACLListAclsRequest) (*ComputeACLList, *http.Response, error)

	/*
		ComputeACLListAclsSAclID Describe an ACL

		Describe an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclID
		 @return APIComputeACLListAclsSAclIDRequest
	*/
	ComputeACLListAclsSAclID(ctx context.Context, aclID string) APIComputeACLListAclsSAclIDRequest

	// ComputeACLListAclsSAclIDExecute executes the request
	//  @return ComputeACLCreateAclsResponse
	ComputeACLListAclsSAclIDExecute(r APIComputeACLListAclsSAclIDRequest) (*ComputeACLCreateAclsResponse, *http.Response, error)

	/*
		ComputeACLLookupAcls Lookup an ACL

		Find a matching ACL entry for an IP address.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclID
		 @param aclIP
		 @return APIComputeACLLookupAclsRequest
	*/
	ComputeACLLookupAcls(ctx context.Context, aclID string, aclIP string) APIComputeACLLookupAclsRequest

	// ComputeACLLookupAclsExecute executes the request
	//  @return ComputeACLLookup
	ComputeACLLookupAclsExecute(r APIComputeACLLookupAclsRequest) (*ComputeACLLookup, *http.Response, error)

	/*
		ComputeACLUpdateAcls Update an ACL

		Update an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclID
		 @return APIComputeACLUpdateAclsRequest
	*/
	ComputeACLUpdateAcls(ctx context.Context, aclID string) APIComputeACLUpdateAclsRequest

	// ComputeACLUpdateAclsExecute executes the request
	ComputeACLUpdateAclsExecute(r APIComputeACLUpdateAclsRequest) (*http.Response, error)
}

// ACLsInComputeAPIService ACLsInComputeAPI service
type ACLsInComputeAPIService service

// APIComputeACLCreateAclsRequest represents a request for the resource.
type APIComputeACLCreateAclsRequest struct {
	ctx                         context.Context
	APIService                  ACLsInComputeAPI
	computeACLCreateAclsRequest *ComputeACLCreateAclsRequest
}

// ComputeACLCreateAclsRequest returns a pointer to a request.
func (r *APIComputeACLCreateAclsRequest) ComputeACLCreateAclsRequest(computeACLCreateAclsRequest ComputeACLCreateAclsRequest) *APIComputeACLCreateAclsRequest {
	r.computeACLCreateAclsRequest = &computeACLCreateAclsRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIComputeACLCreateAclsRequest) Execute() (*ComputeACLCreateAclsResponse, *http.Response, error) {
	return r.APIService.ComputeACLCreateAclsExecute(r)
}

/*
ComputeACLCreateAcls Create a new ACL

Create a new ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIComputeACLCreateAclsRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLCreateAcls(ctx context.Context) APIComputeACLCreateAclsRequest {
	return APIComputeACLCreateAclsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ComputeACLCreateAclsExecute executes the request
//  @return ComputeACLCreateAclsResponse
func (a *ACLsInComputeAPIService) ComputeACLCreateAclsExecute(r APIComputeACLCreateAclsRequest) (*ComputeACLCreateAclsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeACLCreateAclsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLCreateAcls")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls"

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
	localVarPostBody = r.computeACLCreateAclsRequest
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

// APIComputeACLDeleteSAclIDRequest represents a request for the resource.
type APIComputeACLDeleteSAclIDRequest struct {
	ctx        context.Context
	APIService ACLsInComputeAPI
	aclID      string
}

// Execute calls the API using the request data configured.
func (r APIComputeACLDeleteSAclIDRequest) Execute() (*http.Response, error) {
	return r.APIService.ComputeACLDeleteSAclIDExecute(r)
}

/*
ComputeACLDeleteSAclID Delete an ACL

Delete an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclID
 @return APIComputeACLDeleteSAclIDRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLDeleteSAclID(ctx context.Context, aclID string) APIComputeACLDeleteSAclIDRequest {
	return APIComputeACLDeleteSAclIDRequest{
		APIService: a,
		ctx:        ctx,
		aclID:      aclID,
	}
}

// ComputeACLDeleteSAclIDExecute executes the request
func (a *ACLsInComputeAPIService) ComputeACLDeleteSAclIDExecute(r APIComputeACLDeleteSAclIDRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLDeleteSAclID")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))

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

// APIComputeACLListAclEntriesRequest represents a request for the resource.
type APIComputeACLListAclEntriesRequest struct {
	ctx        context.Context
	APIService ACLsInComputeAPI
	aclID      string
	cursor     *string
	limit      *int32
}

// Cursor returns a pointer to a request.
func (r *APIComputeACLListAclEntriesRequest) Cursor(cursor string) *APIComputeACLListAclEntriesRequest {
	r.cursor = &cursor
	return r
}

// Limit returns a pointer to a request.
func (r *APIComputeACLListAclEntriesRequest) Limit(limit int32) *APIComputeACLListAclEntriesRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIComputeACLListAclEntriesRequest) Execute() (*ComputeACLListEntries, *http.Response, error) {
	return r.APIService.ComputeACLListAclEntriesExecute(r)
}

/*
ComputeACLListAclEntries List an ACL

List an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclID
 @return APIComputeACLListAclEntriesRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLListAclEntries(ctx context.Context, aclID string) APIComputeACLListAclEntriesRequest {
	return APIComputeACLListAclEntriesRequest{
		APIService: a,
		ctx:        ctx,
		aclID:      aclID,
	}
}

// ComputeACLListAclEntriesExecute executes the request
//  @return ComputeACLListEntries
func (a *ACLsInComputeAPIService) ComputeACLListAclEntriesExecute(r APIComputeACLListAclEntriesRequest) (*ComputeACLListEntries, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeACLListEntries
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLListAclEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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

// APIComputeACLListAclsRequest represents a request for the resource.
type APIComputeACLListAclsRequest struct {
	ctx        context.Context
	APIService ACLsInComputeAPI
}

// Execute calls the API using the request data configured.
func (r APIComputeACLListAclsRequest) Execute() (*ComputeACLList, *http.Response, error) {
	return r.APIService.ComputeACLListAclsExecute(r)
}

/*
ComputeACLListAcls List ACLs

List all ACLs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIComputeACLListAclsRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLListAcls(ctx context.Context) APIComputeACLListAclsRequest {
	return APIComputeACLListAclsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ComputeACLListAclsExecute executes the request
//  @return ComputeACLList
func (a *ACLsInComputeAPIService) ComputeACLListAclsExecute(r APIComputeACLListAclsRequest) (*ComputeACLList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeACLList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLListAcls")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls"

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

// APIComputeACLListAclsSAclIDRequest represents a request for the resource.
type APIComputeACLListAclsSAclIDRequest struct {
	ctx        context.Context
	APIService ACLsInComputeAPI
	aclID      string
}

// Execute calls the API using the request data configured.
func (r APIComputeACLListAclsSAclIDRequest) Execute() (*ComputeACLCreateAclsResponse, *http.Response, error) {
	return r.APIService.ComputeACLListAclsSAclIDExecute(r)
}

/*
ComputeACLListAclsSAclID Describe an ACL

Describe an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclID
 @return APIComputeACLListAclsSAclIDRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLListAclsSAclID(ctx context.Context, aclID string) APIComputeACLListAclsSAclIDRequest {
	return APIComputeACLListAclsSAclIDRequest{
		APIService: a,
		ctx:        ctx,
		aclID:      aclID,
	}
}

// ComputeACLListAclsSAclIDExecute executes the request
//  @return ComputeACLCreateAclsResponse
func (a *ACLsInComputeAPIService) ComputeACLListAclsSAclIDExecute(r APIComputeACLListAclsSAclIDRequest) (*ComputeACLCreateAclsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeACLCreateAclsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLListAclsSAclID")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))

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

// APIComputeACLLookupAclsRequest represents a request for the resource.
type APIComputeACLLookupAclsRequest struct {
	ctx        context.Context
	APIService ACLsInComputeAPI
	aclID      string
	aclIP      string
}

// Execute calls the API using the request data configured.
func (r APIComputeACLLookupAclsRequest) Execute() (*ComputeACLLookup, *http.Response, error) {
	return r.APIService.ComputeACLLookupAclsExecute(r)
}

/*
ComputeACLLookupAcls Lookup an ACL

Find a matching ACL entry for an IP address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclID
 @param aclIP
 @return APIComputeACLLookupAclsRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLLookupAcls(ctx context.Context, aclID string, aclIP string) APIComputeACLLookupAclsRequest {
	return APIComputeACLLookupAclsRequest{
		APIService: a,
		ctx:        ctx,
		aclID:      aclID,
		aclIP:      aclIP,
	}
}

// ComputeACLLookupAclsExecute executes the request
//  @return ComputeACLLookup
func (a *ACLsInComputeAPIService) ComputeACLLookupAclsExecute(r APIComputeACLLookupAclsRequest) (*ComputeACLLookup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeACLLookup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLLookupAcls")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}/entry/{acl_ip}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_ip"+"}", gourl.PathEscape(parameterToString(r.aclIP, "")))

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

// APIComputeACLUpdateAclsRequest represents a request for the resource.
type APIComputeACLUpdateAclsRequest struct {
	ctx              context.Context
	APIService       ACLsInComputeAPI
	aclID            string
	computeACLUpdate *ComputeACLUpdate
}

// ComputeACLUpdate returns a pointer to a request.
func (r *APIComputeACLUpdateAclsRequest) ComputeACLUpdate(computeACLUpdate ComputeACLUpdate) *APIComputeACLUpdateAclsRequest {
	r.computeACLUpdate = &computeACLUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIComputeACLUpdateAclsRequest) Execute() (*http.Response, error) {
	return r.APIService.ComputeACLUpdateAclsExecute(r)
}

/*
ComputeACLUpdateAcls Update an ACL

Update an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclID
 @return APIComputeACLUpdateAclsRequest
*/
func (a *ACLsInComputeAPIService) ComputeACLUpdateAcls(ctx context.Context, aclID string) APIComputeACLUpdateAclsRequest {
	return APIComputeACLUpdateAclsRequest{
		APIService: a,
		ctx:        ctx,
		aclID:      aclID,
	}
}

// ComputeACLUpdateAclsExecute executes the request
func (a *ACLsInComputeAPIService) ComputeACLUpdateAclsExecute(r APIComputeACLUpdateAclsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLsInComputeAPIService.ComputeACLUpdateAcls")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))

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
	localVarPostBody = r.computeACLUpdate
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
