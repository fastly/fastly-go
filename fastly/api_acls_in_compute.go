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

// AclsInComputeAPI defines an interface for interacting with the resource.
type AclsInComputeAPI interface {

	/*
		ComputeAclCreateAcls Create a new ACL

		Create a new ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIComputeAclCreateAclsRequest
	*/
	ComputeAclCreateAcls(ctx context.Context) APIComputeAclCreateAclsRequest

	// ComputeAclCreateAclsExecute executes the request
	//  @return ComputeAclCreateAclsResponse
	ComputeAclCreateAclsExecute(r APIComputeAclCreateAclsRequest) (*ComputeAclCreateAclsResponse, *http.Response, error)

	/*
		ComputeAclDeleteSAclId Delete an ACL

		Delete an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclId
		 @return APIComputeAclDeleteSAclIdRequest
	*/
	ComputeAclDeleteSAclId(ctx context.Context, aclId string) APIComputeAclDeleteSAclIdRequest

	// ComputeAclDeleteSAclIdExecute executes the request
	ComputeAclDeleteSAclIdExecute(r APIComputeAclDeleteSAclIdRequest) (*http.Response, error)

	/*
		ComputeAclListAclEntries List an ACL

		List an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclId
		 @return APIComputeAclListAclEntriesRequest
	*/
	ComputeAclListAclEntries(ctx context.Context, aclId string) APIComputeAclListAclEntriesRequest

	// ComputeAclListAclEntriesExecute executes the request
	//  @return ComputeAclListEntries
	ComputeAclListAclEntriesExecute(r APIComputeAclListAclEntriesRequest) (*ComputeAclListEntries, *http.Response, error)

	/*
		ComputeAclListAcls List ACLs

		List all ACLs.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIComputeAclListAclsRequest
	*/
	ComputeAclListAcls(ctx context.Context) APIComputeAclListAclsRequest

	// ComputeAclListAclsExecute executes the request
	//  @return ComputeAclList
	ComputeAclListAclsExecute(r APIComputeAclListAclsRequest) (*ComputeAclList, *http.Response, error)

	/*
		ComputeAclListAclsSAclId Describe an ACL

		Describe an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclId
		 @return APIComputeAclListAclsSAclIdRequest
	*/
	ComputeAclListAclsSAclId(ctx context.Context, aclId string) APIComputeAclListAclsSAclIdRequest

	// ComputeAclListAclsSAclIdExecute executes the request
	//  @return ComputeAclCreateAclsResponse
	ComputeAclListAclsSAclIdExecute(r APIComputeAclListAclsSAclIdRequest) (*ComputeAclCreateAclsResponse, *http.Response, error)

	/*
		ComputeAclLookupAcls Lookup an ACL

		Find a matching ACL entry for an IP address.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclId
		 @param aclIp
		 @return APIComputeAclLookupAclsRequest
	*/
	ComputeAclLookupAcls(ctx context.Context, aclId string, aclIp string) APIComputeAclLookupAclsRequest

	// ComputeAclLookupAclsExecute executes the request
	//  @return ComputeAclLookup
	ComputeAclLookupAclsExecute(r APIComputeAclLookupAclsRequest) (*ComputeAclLookup, *http.Response, error)

	/*
		ComputeAclUpdateAcls Update an ACL

		Update an ACL entry with a new operation or action, this allows you to modify an existing entry or delete it.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param aclId
		 @return APIComputeAclUpdateAclsRequest
	*/
	ComputeAclUpdateAcls(ctx context.Context, aclId string) APIComputeAclUpdateAclsRequest

	// ComputeAclUpdateAclsExecute executes the request
	ComputeAclUpdateAclsExecute(r APIComputeAclUpdateAclsRequest) (*http.Response, error)
}

// AclsInComputeAPIService AclsInComputeAPI service
type AclsInComputeAPIService service

// APIComputeAclCreateAclsRequest represents a request for the resource.
type APIComputeAclCreateAclsRequest struct {
	ctx                         context.Context
	APIService                  AclsInComputeAPI
	computeAclCreateAclsRequest *ComputeAclCreateAclsRequest
}

// ComputeAclCreateAclsRequest returns a pointer to a request.
func (r *APIComputeAclCreateAclsRequest) ComputeAclCreateAclsRequest(computeAclCreateAclsRequest ComputeAclCreateAclsRequest) *APIComputeAclCreateAclsRequest {
	r.computeAclCreateAclsRequest = &computeAclCreateAclsRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIComputeAclCreateAclsRequest) Execute() (*ComputeAclCreateAclsResponse, *http.Response, error) {
	return r.APIService.ComputeAclCreateAclsExecute(r)
}

/*
ComputeAclCreateAcls Create a new ACL

Create a new ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIComputeAclCreateAclsRequest
*/
func (a *AclsInComputeAPIService) ComputeAclCreateAcls(ctx context.Context) APIComputeAclCreateAclsRequest {
	return APIComputeAclCreateAclsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ComputeAclCreateAclsExecute executes the request
//  @return ComputeAclCreateAclsResponse
func (a *AclsInComputeAPIService) ComputeAclCreateAclsExecute(r APIComputeAclCreateAclsRequest) (*ComputeAclCreateAclsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeAclCreateAclsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclCreateAcls")
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
	localVarPostBody = r.computeAclCreateAclsRequest
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

// APIComputeAclDeleteSAclIdRequest represents a request for the resource.
type APIComputeAclDeleteSAclIdRequest struct {
	ctx        context.Context
	APIService AclsInComputeAPI
	aclId      string
}

// Execute calls the API using the request data configured.
func (r APIComputeAclDeleteSAclIdRequest) Execute() (*http.Response, error) {
	return r.APIService.ComputeAclDeleteSAclIdExecute(r)
}

/*
ComputeAclDeleteSAclId Delete an ACL

Delete an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclId
 @return APIComputeAclDeleteSAclIdRequest
*/
func (a *AclsInComputeAPIService) ComputeAclDeleteSAclId(ctx context.Context, aclId string) APIComputeAclDeleteSAclIdRequest {
	return APIComputeAclDeleteSAclIdRequest{
		APIService: a,
		ctx:        ctx,
		aclId:      aclId,
	}
}

// ComputeAclDeleteSAclIdExecute executes the request
func (a *AclsInComputeAPIService) ComputeAclDeleteSAclIdExecute(r APIComputeAclDeleteSAclIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclDeleteSAclId")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))

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

// APIComputeAclListAclEntriesRequest represents a request for the resource.
type APIComputeAclListAclEntriesRequest struct {
	ctx        context.Context
	APIService AclsInComputeAPI
	aclId      string
	cursor     *string
	limit      *int32
}

// Cursor returns a pointer to a request.
func (r *APIComputeAclListAclEntriesRequest) Cursor(cursor string) *APIComputeAclListAclEntriesRequest {
	r.cursor = &cursor
	return r
}

// Limit returns a pointer to a request.
func (r *APIComputeAclListAclEntriesRequest) Limit(limit int32) *APIComputeAclListAclEntriesRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIComputeAclListAclEntriesRequest) Execute() (*ComputeAclListEntries, *http.Response, error) {
	return r.APIService.ComputeAclListAclEntriesExecute(r)
}

/*
ComputeAclListAclEntries List an ACL

List an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclId
 @return APIComputeAclListAclEntriesRequest
*/
func (a *AclsInComputeAPIService) ComputeAclListAclEntries(ctx context.Context, aclId string) APIComputeAclListAclEntriesRequest {
	return APIComputeAclListAclEntriesRequest{
		APIService: a,
		ctx:        ctx,
		aclId:      aclId,
	}
}

// ComputeAclListAclEntriesExecute executes the request
//  @return ComputeAclListEntries
func (a *AclsInComputeAPIService) ComputeAclListAclEntriesExecute(r APIComputeAclListAclEntriesRequest) (*ComputeAclListEntries, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeAclListEntries
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclListAclEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))

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

// APIComputeAclListAclsRequest represents a request for the resource.
type APIComputeAclListAclsRequest struct {
	ctx        context.Context
	APIService AclsInComputeAPI
}

// Execute calls the API using the request data configured.
func (r APIComputeAclListAclsRequest) Execute() (*ComputeAclList, *http.Response, error) {
	return r.APIService.ComputeAclListAclsExecute(r)
}

/*
ComputeAclListAcls List ACLs

List all ACLs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIComputeAclListAclsRequest
*/
func (a *AclsInComputeAPIService) ComputeAclListAcls(ctx context.Context) APIComputeAclListAclsRequest {
	return APIComputeAclListAclsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ComputeAclListAclsExecute executes the request
//  @return ComputeAclList
func (a *AclsInComputeAPIService) ComputeAclListAclsExecute(r APIComputeAclListAclsRequest) (*ComputeAclList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeAclList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclListAcls")
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

// APIComputeAclListAclsSAclIdRequest represents a request for the resource.
type APIComputeAclListAclsSAclIdRequest struct {
	ctx        context.Context
	APIService AclsInComputeAPI
	aclId      string
}

// Execute calls the API using the request data configured.
func (r APIComputeAclListAclsSAclIdRequest) Execute() (*ComputeAclCreateAclsResponse, *http.Response, error) {
	return r.APIService.ComputeAclListAclsSAclIdExecute(r)
}

/*
ComputeAclListAclsSAclId Describe an ACL

Describe an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclId
 @return APIComputeAclListAclsSAclIdRequest
*/
func (a *AclsInComputeAPIService) ComputeAclListAclsSAclId(ctx context.Context, aclId string) APIComputeAclListAclsSAclIdRequest {
	return APIComputeAclListAclsSAclIdRequest{
		APIService: a,
		ctx:        ctx,
		aclId:      aclId,
	}
}

// ComputeAclListAclsSAclIdExecute executes the request
//  @return ComputeAclCreateAclsResponse
func (a *AclsInComputeAPIService) ComputeAclListAclsSAclIdExecute(r APIComputeAclListAclsSAclIdRequest) (*ComputeAclCreateAclsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeAclCreateAclsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclListAclsSAclId")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))

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

// APIComputeAclLookupAclsRequest represents a request for the resource.
type APIComputeAclLookupAclsRequest struct {
	ctx        context.Context
	APIService AclsInComputeAPI
	aclId      string
	aclIp      string
}

// Execute calls the API using the request data configured.
func (r APIComputeAclLookupAclsRequest) Execute() (*ComputeAclLookup, *http.Response, error) {
	return r.APIService.ComputeAclLookupAclsExecute(r)
}

/*
ComputeAclLookupAcls Lookup an ACL

Find a matching ACL entry for an IP address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclId
 @param aclIp
 @return APIComputeAclLookupAclsRequest
*/
func (a *AclsInComputeAPIService) ComputeAclLookupAcls(ctx context.Context, aclId string, aclIp string) APIComputeAclLookupAclsRequest {
	return APIComputeAclLookupAclsRequest{
		APIService: a,
		ctx:        ctx,
		aclId:      aclId,
		aclIp:      aclIp,
	}
}

// ComputeAclLookupAclsExecute executes the request
//  @return ComputeAclLookup
func (a *AclsInComputeAPIService) ComputeAclLookupAclsExecute(r APIComputeAclLookupAclsRequest) (*ComputeAclLookup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ComputeAclLookup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclLookupAcls")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}/entry/{acl_ip}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_ip"+"}", gourl.PathEscape(parameterToString(r.aclIp, "")))

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

// APIComputeAclUpdateAclsRequest represents a request for the resource.
type APIComputeAclUpdateAclsRequest struct {
	ctx              context.Context
	APIService       AclsInComputeAPI
	aclId            string
	computeAclUpdate *ComputeAclUpdate
}

// ComputeAclUpdate returns a pointer to a request.
func (r *APIComputeAclUpdateAclsRequest) ComputeAclUpdate(computeAclUpdate ComputeAclUpdate) *APIComputeAclUpdateAclsRequest {
	r.computeAclUpdate = &computeAclUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIComputeAclUpdateAclsRequest) Execute() (*http.Response, error) {
	return r.APIService.ComputeAclUpdateAclsExecute(r)
}

/*
ComputeAclUpdateAcls Update an ACL

Update an ACL entry with a new operation or action, this allows you to modify an existing entry or delete it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param aclId
 @return APIComputeAclUpdateAclsRequest
*/
func (a *AclsInComputeAPIService) ComputeAclUpdateAcls(ctx context.Context, aclId string) APIComputeAclUpdateAclsRequest {
	return APIComputeAclUpdateAclsRequest{
		APIService: a,
		ctx:        ctx,
		aclId:      aclId,
	}
}

// ComputeAclUpdateAclsExecute executes the request
func (a *AclsInComputeAPIService) ComputeAclUpdateAclsExecute(r APIComputeAclUpdateAclsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsInComputeAPIService.ComputeAclUpdateAcls")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/acls/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))

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
	localVarPostBody = r.computeAclUpdate
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
