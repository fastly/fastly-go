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

// ACLEntryAPI defines an interface for interacting with the resource.
type ACLEntryAPI interface {

	/*
	BulkUpdateACLEntries Update multiple ACL entries

	Update multiple ACL entries on the same ACL. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 entries. [Contact support](https://support.fastly.com/) to discuss raising this limit.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param aclID Alphanumeric string identifying a ACL.
	 @return APIBulkUpdateACLEntriesRequest
	*/
	BulkUpdateACLEntries(ctx context.Context, serviceID string, aclID string) APIBulkUpdateACLEntriesRequest

	// BulkUpdateACLEntriesExecute executes the request
	//  @return InlineResponse200
	BulkUpdateACLEntriesExecute(r APIBulkUpdateACLEntriesRequest) (*InlineResponse200, *http.Response, error)

	/*
	CreateACLEntry Create an ACL entry

	Add an ACL entry to an ACL.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param aclID Alphanumeric string identifying a ACL.
	 @return APICreateACLEntryRequest
	*/
	CreateACLEntry(ctx context.Context, serviceID string, aclID string) APICreateACLEntryRequest

	// CreateACLEntryExecute executes the request
	//  @return ACLEntryResponse
	CreateACLEntryExecute(r APICreateACLEntryRequest) (*ACLEntryResponse, *http.Response, error)

	/*
	DeleteACLEntry Delete an ACL entry

	Delete an ACL entry from a specified ACL.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param aclID Alphanumeric string identifying a ACL.
	 @param aclEntryID Alphanumeric string identifying an ACL Entry.
	 @return APIDeleteACLEntryRequest
	*/
	DeleteACLEntry(ctx context.Context, serviceID string, aclID string, aclEntryID string) APIDeleteACLEntryRequest

	// DeleteACLEntryExecute executes the request
	//  @return InlineResponse200
	DeleteACLEntryExecute(r APIDeleteACLEntryRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetACLEntry Describe an ACL entry

	Retrieve a single ACL entry.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param aclID Alphanumeric string identifying a ACL.
	 @param aclEntryID Alphanumeric string identifying an ACL Entry.
	 @return APIGetACLEntryRequest
	*/
	GetACLEntry(ctx context.Context, serviceID string, aclID string, aclEntryID string) APIGetACLEntryRequest

	// GetACLEntryExecute executes the request
	//  @return ACLEntryResponse
	GetACLEntryExecute(r APIGetACLEntryRequest) (*ACLEntryResponse, *http.Response, error)

	/*
	ListACLEntries List ACL entries

	List ACL entries for a specified ACL.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param aclID Alphanumeric string identifying a ACL.
	 @return APIListACLEntriesRequest
	*/
	ListACLEntries(ctx context.Context, serviceID string, aclID string) APIListACLEntriesRequest

	// ListACLEntriesExecute executes the request
	//  @return []ACLEntryResponse
	ListACLEntriesExecute(r APIListACLEntriesRequest) ([]ACLEntryResponse, *http.Response, error)

	/*
	UpdateACLEntry Update an ACL entry

	Update an ACL entry for a specified ACL.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param aclID Alphanumeric string identifying a ACL.
	 @param aclEntryID Alphanumeric string identifying an ACL Entry.
	 @return APIUpdateACLEntryRequest
	*/
	UpdateACLEntry(ctx context.Context, serviceID string, aclID string, aclEntryID string) APIUpdateACLEntryRequest

	// UpdateACLEntryExecute executes the request
	//  @return ACLEntryResponse
	UpdateACLEntryExecute(r APIUpdateACLEntryRequest) (*ACLEntryResponse, *http.Response, error)
}

// ACLEntryAPIService ACLEntryAPI service
type ACLEntryAPIService service

// APIBulkUpdateACLEntriesRequest represents a request for the resource.
type APIBulkUpdateACLEntriesRequest struct {
	ctx context.Context
	APIService ACLEntryAPI
	serviceID string
	aclID string
	bulkUpdateACLEntriesRequest *BulkUpdateACLEntriesRequest
}

// BulkUpdateACLEntriesRequest returns a pointer to a request.
func (r *APIBulkUpdateACLEntriesRequest) BulkUpdateACLEntriesRequest(bulkUpdateACLEntriesRequest BulkUpdateACLEntriesRequest) *APIBulkUpdateACLEntriesRequest {
	r.bulkUpdateACLEntriesRequest = &bulkUpdateACLEntriesRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkUpdateACLEntriesRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.BulkUpdateACLEntriesExecute(r)
}

/*
BulkUpdateACLEntries Update multiple ACL entries

Update multiple ACL entries on the same ACL. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 entries. [Contact support](https://support.fastly.com/) to discuss raising this limit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param aclID Alphanumeric string identifying a ACL.
 @return APIBulkUpdateACLEntriesRequest
*/
func (a *ACLEntryAPIService) BulkUpdateACLEntries(ctx context.Context, serviceID string, aclID string) APIBulkUpdateACLEntriesRequest {
	return APIBulkUpdateACLEntriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		aclID: aclID,
	}
}

// BulkUpdateACLEntriesExecute executes the request
//  @return InlineResponse200
func (a *ACLEntryAPIService) BulkUpdateACLEntriesExecute(r APIBulkUpdateACLEntriesRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLEntryAPIService.BulkUpdateACLEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkUpdateACLEntriesRequest
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

// APICreateACLEntryRequest represents a request for the resource.
type APICreateACLEntryRequest struct {
	ctx context.Context
	APIService ACLEntryAPI
	serviceID string
	aclID string
	aclEntry *ACLEntry
}

// ACLEntry returns a pointer to a request.
func (r *APICreateACLEntryRequest) ACLEntry(aclEntry ACLEntry) *APICreateACLEntryRequest {
	r.aclEntry = &aclEntry
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateACLEntryRequest) Execute() (*ACLEntryResponse, *http.Response, error) {
	return r.APIService.CreateACLEntryExecute(r)
}

/*
CreateACLEntry Create an ACL entry

Add an ACL entry to an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param aclID Alphanumeric string identifying a ACL.
 @return APICreateACLEntryRequest
*/
func (a *ACLEntryAPIService) CreateACLEntry(ctx context.Context, serviceID string, aclID string) APICreateACLEntryRequest {
	return APICreateACLEntryRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		aclID: aclID,
	}
}

// CreateACLEntryExecute executes the request
//  @return ACLEntryResponse
func (a *ACLEntryAPIService) CreateACLEntryExecute(r APICreateACLEntryRequest) (*ACLEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ACLEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLEntryAPIService.CreateACLEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.aclEntry
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

// APIDeleteACLEntryRequest represents a request for the resource.
type APIDeleteACLEntryRequest struct {
	ctx context.Context
	APIService ACLEntryAPI
	serviceID string
	aclID string
	aclEntryID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteACLEntryRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteACLEntryExecute(r)
}

/*
DeleteACLEntry Delete an ACL entry

Delete an ACL entry from a specified ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param aclID Alphanumeric string identifying a ACL.
 @param aclEntryID Alphanumeric string identifying an ACL Entry.
 @return APIDeleteACLEntryRequest
*/
func (a *ACLEntryAPIService) DeleteACLEntry(ctx context.Context, serviceID string, aclID string, aclEntryID string) APIDeleteACLEntryRequest {
	return APIDeleteACLEntryRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		aclID: aclID,
		aclEntryID: aclEntryID,
	}
}

// DeleteACLEntryExecute executes the request
//  @return InlineResponse200
func (a *ACLEntryAPIService) DeleteACLEntryExecute(r APIDeleteACLEntryRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLEntryAPIService.DeleteACLEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_entry_id"+"}", gourl.PathEscape(parameterToString(r.aclEntryID, "")))

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

// APIGetACLEntryRequest represents a request for the resource.
type APIGetACLEntryRequest struct {
	ctx context.Context
	APIService ACLEntryAPI
	serviceID string
	aclID string
	aclEntryID string
}


// Execute calls the API using the request data configured.
func (r APIGetACLEntryRequest) Execute() (*ACLEntryResponse, *http.Response, error) {
	return r.APIService.GetACLEntryExecute(r)
}

/*
GetACLEntry Describe an ACL entry

Retrieve a single ACL entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param aclID Alphanumeric string identifying a ACL.
 @param aclEntryID Alphanumeric string identifying an ACL Entry.
 @return APIGetACLEntryRequest
*/
func (a *ACLEntryAPIService) GetACLEntry(ctx context.Context, serviceID string, aclID string, aclEntryID string) APIGetACLEntryRequest {
	return APIGetACLEntryRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		aclID: aclID,
		aclEntryID: aclEntryID,
	}
}

// GetACLEntryExecute executes the request
//  @return ACLEntryResponse
func (a *ACLEntryAPIService) GetACLEntryExecute(r APIGetACLEntryRequest) (*ACLEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ACLEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLEntryAPIService.GetACLEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_entry_id"+"}", gourl.PathEscape(parameterToString(r.aclEntryID, "")))

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

// APIListACLEntriesRequest represents a request for the resource.
type APIListACLEntriesRequest struct {
	ctx context.Context
	APIService ACLEntryAPI
	serviceID string
	aclID string
	page *int32
	perPage *int32
	sort *string
	direction *string
}

// Page Current page.
func (r *APIListACLEntriesRequest) Page(page int32) *APIListACLEntriesRequest {
	r.page = &page
	return r
}
// PerPage Number of records per page.
func (r *APIListACLEntriesRequest) PerPage(perPage int32) *APIListACLEntriesRequest {
	r.perPage = &perPage
	return r
}
// Sort Field on which to sort.
func (r *APIListACLEntriesRequest) Sort(sort string) *APIListACLEntriesRequest {
	r.sort = &sort
	return r
}
// Direction Direction in which to sort results.
func (r *APIListACLEntriesRequest) Direction(direction string) *APIListACLEntriesRequest {
	r.direction = &direction
	return r
}

// Execute calls the API using the request data configured.
func (r APIListACLEntriesRequest) Execute() ([]ACLEntryResponse, *http.Response, error) {
	return r.APIService.ListACLEntriesExecute(r)
}

/*
ListACLEntries List ACL entries

List ACL entries for a specified ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param aclID Alphanumeric string identifying a ACL.
 @return APIListACLEntriesRequest
*/
func (a *ACLEntryAPIService) ListACLEntries(ctx context.Context, serviceID string, aclID string) APIListACLEntriesRequest {
	return APIListACLEntriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		aclID: aclID,
	}
}

// ListACLEntriesExecute executes the request
//  @return []ACLEntryResponse
func (a *ACLEntryAPIService) ListACLEntriesExecute(r APIListACLEntriesRequest) ([]ACLEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []ACLEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLEntryAPIService.ListACLEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
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

// APIUpdateACLEntryRequest represents a request for the resource.
type APIUpdateACLEntryRequest struct {
	ctx context.Context
	APIService ACLEntryAPI
	serviceID string
	aclID string
	aclEntryID string
	aclEntry *ACLEntry
}

// ACLEntry returns a pointer to a request.
func (r *APIUpdateACLEntryRequest) ACLEntry(aclEntry ACLEntry) *APIUpdateACLEntryRequest {
	r.aclEntry = &aclEntry
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateACLEntryRequest) Execute() (*ACLEntryResponse, *http.Response, error) {
	return r.APIService.UpdateACLEntryExecute(r)
}

/*
UpdateACLEntry Update an ACL entry

Update an ACL entry for a specified ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param aclID Alphanumeric string identifying a ACL.
 @param aclEntryID Alphanumeric string identifying an ACL Entry.
 @return APIUpdateACLEntryRequest
*/
func (a *ACLEntryAPIService) UpdateACLEntry(ctx context.Context, serviceID string, aclID string, aclEntryID string) APIUpdateACLEntryRequest {
	return APIUpdateACLEntryRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		aclID: aclID,
		aclEntryID: aclEntryID,
	}
}

// UpdateACLEntryExecute executes the request
//  @return ACLEntryResponse
func (a *ACLEntryAPIService) UpdateACLEntryExecute(r APIUpdateACLEntryRequest) (*ACLEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ACLEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ACLEntryAPIService.UpdateACLEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_entry_id"+"}", gourl.PathEscape(parameterToString(r.aclEntryID, "")))

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
	localVarPostBody = r.aclEntry
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
