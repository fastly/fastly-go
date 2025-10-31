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

// AclEntryAPI defines an interface for interacting with the resource.
type AclEntryAPI interface {

	/*
		BulkUpdateAclEntries Update multiple ACL entries

		Update multiple ACL entries on the same ACL. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 entries. [Contact support](https://support.fastly.com/) to discuss raising this limit.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param aclId Alphanumeric string identifying a ACL.
		 @return APIBulkUpdateAclEntriesRequest
	*/
	BulkUpdateAclEntries(ctx context.Context, serviceId string, aclId string) APIBulkUpdateAclEntriesRequest

	// BulkUpdateAclEntriesExecute executes the request
	//  @return InlineResponse200
	BulkUpdateAclEntriesExecute(r APIBulkUpdateAclEntriesRequest) (*InlineResponse200, *http.Response, error)

	/*
		CreateAclEntry Create an ACL entry

		Add an ACL entry to an ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param aclId Alphanumeric string identifying a ACL.
		 @return APICreateAclEntryRequest
	*/
	CreateAclEntry(ctx context.Context, serviceId string, aclId string) APICreateAclEntryRequest

	// CreateAclEntryExecute executes the request
	//  @return AclEntryResponse
	CreateAclEntryExecute(r APICreateAclEntryRequest) (*AclEntryResponse, *http.Response, error)

	/*
		DeleteAclEntry Delete an ACL entry

		Delete an ACL entry from a specified ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param aclId Alphanumeric string identifying a ACL.
		 @param aclEntryId Alphanumeric string identifying an ACL Entry.
		 @return APIDeleteAclEntryRequest
	*/
	DeleteAclEntry(ctx context.Context, serviceId string, aclId string, aclEntryId string) APIDeleteAclEntryRequest

	// DeleteAclEntryExecute executes the request
	//  @return InlineResponse200
	DeleteAclEntryExecute(r APIDeleteAclEntryRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetAclEntry Describe an ACL entry

		Retrieve a single ACL entry.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param aclId Alphanumeric string identifying a ACL.
		 @param aclEntryId Alphanumeric string identifying an ACL Entry.
		 @return APIGetAclEntryRequest
	*/
	GetAclEntry(ctx context.Context, serviceId string, aclId string, aclEntryId string) APIGetAclEntryRequest

	// GetAclEntryExecute executes the request
	//  @return AclEntryResponse
	GetAclEntryExecute(r APIGetAclEntryRequest) (*AclEntryResponse, *http.Response, error)

	/*
		ListAclEntries List ACL entries

		List ACL entries for a specified ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param aclId Alphanumeric string identifying a ACL.
		 @return APIListAclEntriesRequest
	*/
	ListAclEntries(ctx context.Context, serviceId string, aclId string) APIListAclEntriesRequest

	// ListAclEntriesExecute executes the request
	//  @return []AclEntryResponse
	ListAclEntriesExecute(r APIListAclEntriesRequest) ([]AclEntryResponse, *http.Response, error)

	/*
		UpdateAclEntry Update an ACL entry

		Update an ACL entry for a specified ACL.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param aclId Alphanumeric string identifying a ACL.
		 @param aclEntryId Alphanumeric string identifying an ACL Entry.
		 @return APIUpdateAclEntryRequest
	*/
	UpdateAclEntry(ctx context.Context, serviceId string, aclId string, aclEntryId string) APIUpdateAclEntryRequest

	// UpdateAclEntryExecute executes the request
	//  @return AclEntryResponse
	UpdateAclEntryExecute(r APIUpdateAclEntryRequest) (*AclEntryResponse, *http.Response, error)
}

// AclEntryAPIService AclEntryAPI service
type AclEntryAPIService service

// APIBulkUpdateAclEntriesRequest represents a request for the resource.
type APIBulkUpdateAclEntriesRequest struct {
	ctx                         context.Context
	APIService                  AclEntryAPI
	serviceId                   string
	aclId                       string
	bulkUpdateAclEntriesRequest *BulkUpdateAclEntriesRequest
}

// BulkUpdateAclEntriesRequest returns a pointer to a request.
func (r *APIBulkUpdateAclEntriesRequest) BulkUpdateAclEntriesRequest(bulkUpdateAclEntriesRequest BulkUpdateAclEntriesRequest) *APIBulkUpdateAclEntriesRequest {
	r.bulkUpdateAclEntriesRequest = &bulkUpdateAclEntriesRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkUpdateAclEntriesRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.BulkUpdateAclEntriesExecute(r)
}

/*
BulkUpdateAclEntries Update multiple ACL entries

Update multiple ACL entries on the same ACL. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 entries. [Contact support](https://support.fastly.com/) to discuss raising this limit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param aclId Alphanumeric string identifying a ACL.
 @return APIBulkUpdateAclEntriesRequest
*/
func (a *AclEntryAPIService) BulkUpdateAclEntries(ctx context.Context, serviceId string, aclId string) APIBulkUpdateAclEntriesRequest {
	return APIBulkUpdateAclEntriesRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		aclId:      aclId,
	}
}

// BulkUpdateAclEntriesExecute executes the request
//  @return InlineResponse200
func (a *AclEntryAPIService) BulkUpdateAclEntriesExecute(r APIBulkUpdateAclEntriesRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclEntryAPIService.BulkUpdateAclEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkUpdateAclEntriesRequest
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

// APICreateAclEntryRequest represents a request for the resource.
type APICreateAclEntryRequest struct {
	ctx        context.Context
	APIService AclEntryAPI
	serviceId  string
	aclId      string
	aclEntry   *AclEntry
}

// AclEntry returns a pointer to a request.
func (r *APICreateAclEntryRequest) AclEntry(aclEntry AclEntry) *APICreateAclEntryRequest {
	r.aclEntry = &aclEntry
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateAclEntryRequest) Execute() (*AclEntryResponse, *http.Response, error) {
	return r.APIService.CreateAclEntryExecute(r)
}

/*
CreateAclEntry Create an ACL entry

Add an ACL entry to an ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param aclId Alphanumeric string identifying a ACL.
 @return APICreateAclEntryRequest
*/
func (a *AclEntryAPIService) CreateAclEntry(ctx context.Context, serviceId string, aclId string) APICreateAclEntryRequest {
	return APICreateAclEntryRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		aclId:      aclId,
	}
}

// CreateAclEntryExecute executes the request
//  @return AclEntryResponse
func (a *AclEntryAPIService) CreateAclEntryExecute(r APICreateAclEntryRequest) (*AclEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AclEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclEntryAPIService.CreateAclEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
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

// APIDeleteAclEntryRequest represents a request for the resource.
type APIDeleteAclEntryRequest struct {
	ctx        context.Context
	APIService AclEntryAPI
	serviceId  string
	aclId      string
	aclEntryId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteAclEntryRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteAclEntryExecute(r)
}

/*
DeleteAclEntry Delete an ACL entry

Delete an ACL entry from a specified ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param aclId Alphanumeric string identifying a ACL.
 @param aclEntryId Alphanumeric string identifying an ACL Entry.
 @return APIDeleteAclEntryRequest
*/
func (a *AclEntryAPIService) DeleteAclEntry(ctx context.Context, serviceId string, aclId string, aclEntryId string) APIDeleteAclEntryRequest {
	return APIDeleteAclEntryRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		aclId:      aclId,
		aclEntryId: aclEntryId,
	}
}

// DeleteAclEntryExecute executes the request
//  @return InlineResponse200
func (a *AclEntryAPIService) DeleteAclEntryExecute(r APIDeleteAclEntryRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclEntryAPIService.DeleteAclEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_entry_id"+"}", gourl.PathEscape(parameterToString(r.aclEntryId, "")))

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

// APIGetAclEntryRequest represents a request for the resource.
type APIGetAclEntryRequest struct {
	ctx        context.Context
	APIService AclEntryAPI
	serviceId  string
	aclId      string
	aclEntryId string
}

// Execute calls the API using the request data configured.
func (r APIGetAclEntryRequest) Execute() (*AclEntryResponse, *http.Response, error) {
	return r.APIService.GetAclEntryExecute(r)
}

/*
GetAclEntry Describe an ACL entry

Retrieve a single ACL entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param aclId Alphanumeric string identifying a ACL.
 @param aclEntryId Alphanumeric string identifying an ACL Entry.
 @return APIGetAclEntryRequest
*/
func (a *AclEntryAPIService) GetAclEntry(ctx context.Context, serviceId string, aclId string, aclEntryId string) APIGetAclEntryRequest {
	return APIGetAclEntryRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		aclId:      aclId,
		aclEntryId: aclEntryId,
	}
}

// GetAclEntryExecute executes the request
//  @return AclEntryResponse
func (a *AclEntryAPIService) GetAclEntryExecute(r APIGetAclEntryRequest) (*AclEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AclEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclEntryAPIService.GetAclEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_entry_id"+"}", gourl.PathEscape(parameterToString(r.aclEntryId, "")))

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

// APIListAclEntriesRequest represents a request for the resource.
type APIListAclEntriesRequest struct {
	ctx        context.Context
	APIService AclEntryAPI
	serviceId  string
	aclId      string
	page       *int32
	perPage    *int32
	sort       *string
	direction  *string
}

// Page Current page.
func (r *APIListAclEntriesRequest) Page(page int32) *APIListAclEntriesRequest {
	r.page = &page
	return r
}

// PerPage Number of records per page.
func (r *APIListAclEntriesRequest) PerPage(perPage int32) *APIListAclEntriesRequest {
	r.perPage = &perPage
	return r
}

// Sort Field on which to sort.
func (r *APIListAclEntriesRequest) Sort(sort string) *APIListAclEntriesRequest {
	r.sort = &sort
	return r
}

// Direction Direction in which to sort results.
func (r *APIListAclEntriesRequest) Direction(direction string) *APIListAclEntriesRequest {
	r.direction = &direction
	return r
}

// Execute calls the API using the request data configured.
func (r APIListAclEntriesRequest) Execute() ([]AclEntryResponse, *http.Response, error) {
	return r.APIService.ListAclEntriesExecute(r)
}

/*
ListAclEntries List ACL entries

List ACL entries for a specified ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param aclId Alphanumeric string identifying a ACL.
 @return APIListAclEntriesRequest
*/
func (a *AclEntryAPIService) ListAclEntries(ctx context.Context, serviceId string, aclId string) APIListAclEntriesRequest {
	return APIListAclEntriesRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		aclId:      aclId,
	}
}

// ListAclEntriesExecute executes the request
//  @return []AclEntryResponse
func (a *AclEntryAPIService) ListAclEntriesExecute(r APIListAclEntriesRequest) ([]AclEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []AclEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclEntryAPIService.ListAclEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entries"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))

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

// APIUpdateAclEntryRequest represents a request for the resource.
type APIUpdateAclEntryRequest struct {
	ctx        context.Context
	APIService AclEntryAPI
	serviceId  string
	aclId      string
	aclEntryId string
	aclEntry   *AclEntry
}

// AclEntry returns a pointer to a request.
func (r *APIUpdateAclEntryRequest) AclEntry(aclEntry AclEntry) *APIUpdateAclEntryRequest {
	r.aclEntry = &aclEntry
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateAclEntryRequest) Execute() (*AclEntryResponse, *http.Response, error) {
	return r.APIService.UpdateAclEntryExecute(r)
}

/*
UpdateAclEntry Update an ACL entry

Update an ACL entry for a specified ACL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param aclId Alphanumeric string identifying a ACL.
 @param aclEntryId Alphanumeric string identifying an ACL Entry.
 @return APIUpdateAclEntryRequest
*/
func (a *AclEntryAPIService) UpdateAclEntry(ctx context.Context, serviceId string, aclId string, aclEntryId string) APIUpdateAclEntryRequest {
	return APIUpdateAclEntryRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		aclId:      aclId,
		aclEntryId: aclEntryId,
	}
}

// UpdateAclEntryExecute executes the request
//  @return AclEntryResponse
func (a *AclEntryAPIService) UpdateAclEntryExecute(r APIUpdateAclEntryRequest) (*AclEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AclEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclEntryAPIService.UpdateAclEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_id"+"}", gourl.PathEscape(parameterToString(r.aclId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"acl_entry_id"+"}", gourl.PathEscape(parameterToString(r.aclEntryId, "")))

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
