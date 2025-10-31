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

// DmDomainsAPI defines an interface for interacting with the resource.
type DmDomainsAPI interface {

	/*
		CreateDmDomain Create a domain

		Create a domain

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateDmDomainRequest
	*/
	CreateDmDomain(ctx context.Context) APICreateDmDomainRequest

	// CreateDmDomainExecute executes the request
	//  @return SuccessfulResponseAsObject
	CreateDmDomainExecute(r APICreateDmDomainRequest) (*SuccessfulResponseAsObject, *http.Response, error)

	/*
		DeleteDmDomain Delete a domain

		Delete a domain

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param domainId
		 @return APIDeleteDmDomainRequest
	*/
	DeleteDmDomain(ctx context.Context, domainId string) APIDeleteDmDomainRequest

	// DeleteDmDomainExecute executes the request
	DeleteDmDomainExecute(r APIDeleteDmDomainRequest) (*http.Response, error)

	/*
		GetDmDomain Get a domain

		Show a domain

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param domainId
		 @return APIGetDmDomainRequest
	*/
	GetDmDomain(ctx context.Context, domainId string) APIGetDmDomainRequest

	// GetDmDomainExecute executes the request
	//  @return SuccessfulResponseAsObject
	GetDmDomainExecute(r APIGetDmDomainRequest) (*SuccessfulResponseAsObject, *http.Response, error)

	/*
		ListDmDomains List domains

		List all domains

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListDmDomainsRequest
	*/
	ListDmDomains(ctx context.Context) APIListDmDomainsRequest

	// ListDmDomainsExecute executes the request
	//  @return InlineResponse2004
	ListDmDomainsExecute(r APIListDmDomainsRequest) (*InlineResponse2004, *http.Response, error)

	/*
		UpdateDmDomain Update a domain

		Update a domain

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param domainId
		 @return APIUpdateDmDomainRequest
	*/
	UpdateDmDomain(ctx context.Context, domainId string) APIUpdateDmDomainRequest

	// UpdateDmDomainExecute executes the request
	//  @return SuccessfulResponseAsObject
	UpdateDmDomainExecute(r APIUpdateDmDomainRequest) (*SuccessfulResponseAsObject, *http.Response, error)
}

// DmDomainsAPIService DmDomainsAPI service
type DmDomainsAPIService service

// APICreateDmDomainRequest represents a request for the resource.
type APICreateDmDomainRequest struct {
	ctx                  context.Context
	APIService           DmDomainsAPI
	requestBodyForCreate *RequestBodyForCreate
}

// RequestBodyForCreate returns a pointer to a request.
func (r *APICreateDmDomainRequest) RequestBodyForCreate(requestBodyForCreate RequestBodyForCreate) *APICreateDmDomainRequest {
	r.requestBodyForCreate = &requestBodyForCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateDmDomainRequest) Execute() (*SuccessfulResponseAsObject, *http.Response, error) {
	return r.APIService.CreateDmDomainExecute(r)
}

/*
CreateDmDomain Create a domain

Create a domain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateDmDomainRequest
*/
func (a *DmDomainsAPIService) CreateDmDomain(ctx context.Context) APICreateDmDomainRequest {
	return APICreateDmDomainRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateDmDomainExecute executes the request
//  @return SuccessfulResponseAsObject
func (a *DmDomainsAPIService) CreateDmDomainExecute(r APICreateDmDomainRequest) (*SuccessfulResponseAsObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SuccessfulResponseAsObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmDomainsAPIService.CreateDmDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/domains"

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
	localVarPostBody = r.requestBodyForCreate
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

// APIDeleteDmDomainRequest represents a request for the resource.
type APIDeleteDmDomainRequest struct {
	ctx        context.Context
	APIService DmDomainsAPI
	domainId   string
}

// Execute calls the API using the request data configured.
func (r APIDeleteDmDomainRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteDmDomainExecute(r)
}

/*
DeleteDmDomain Delete a domain

Delete a domain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainId
 @return APIDeleteDmDomainRequest
*/
func (a *DmDomainsAPIService) DeleteDmDomain(ctx context.Context, domainId string) APIDeleteDmDomainRequest {
	return APIDeleteDmDomainRequest{
		APIService: a,
		ctx:        ctx,
		domainId:   domainId,
	}
}

// DeleteDmDomainExecute executes the request
func (a *DmDomainsAPIService) DeleteDmDomainExecute(r APIDeleteDmDomainRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmDomainsAPIService.DeleteDmDomain")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/domains/{domain_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"domain_id"+"}", gourl.PathEscape(parameterToString(r.domainId, "")))

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

// APIGetDmDomainRequest represents a request for the resource.
type APIGetDmDomainRequest struct {
	ctx        context.Context
	APIService DmDomainsAPI
	domainId   string
}

// Execute calls the API using the request data configured.
func (r APIGetDmDomainRequest) Execute() (*SuccessfulResponseAsObject, *http.Response, error) {
	return r.APIService.GetDmDomainExecute(r)
}

/*
GetDmDomain Get a domain

Show a domain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainId
 @return APIGetDmDomainRequest
*/
func (a *DmDomainsAPIService) GetDmDomain(ctx context.Context, domainId string) APIGetDmDomainRequest {
	return APIGetDmDomainRequest{
		APIService: a,
		ctx:        ctx,
		domainId:   domainId,
	}
}

// GetDmDomainExecute executes the request
//  @return SuccessfulResponseAsObject
func (a *DmDomainsAPIService) GetDmDomainExecute(r APIGetDmDomainRequest) (*SuccessfulResponseAsObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SuccessfulResponseAsObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmDomainsAPIService.GetDmDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/domains/{domain_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"domain_id"+"}", gourl.PathEscape(parameterToString(r.domainId, "")))

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

// APIListDmDomainsRequest represents a request for the resource.
type APIListDmDomainsRequest struct {
	ctx        context.Context
	APIService DmDomainsAPI
	fqdn       *string
	serviceId  *string
	sort       *string
	activated  *bool
	verified   *bool
	cursor     *string
	limit      *int32
}

// Fqdn returns a pointer to a request.
func (r *APIListDmDomainsRequest) Fqdn(fqdn string) *APIListDmDomainsRequest {
	r.fqdn = &fqdn
	return r
}

// ServiceId Filter results based on a service_id.
func (r *APIListDmDomainsRequest) ServiceId(serviceId string) *APIListDmDomainsRequest {
	r.serviceId = &serviceId
	return r
}

// Sort The order in which to list the results.
func (r *APIListDmDomainsRequest) Sort(sort string) *APIListDmDomainsRequest {
	r.sort = &sort
	return r
}

// Activated returns a pointer to a request.
func (r *APIListDmDomainsRequest) Activated(activated bool) *APIListDmDomainsRequest {
	r.activated = &activated
	return r
}

// Verified returns a pointer to a request.
func (r *APIListDmDomainsRequest) Verified(verified bool) *APIListDmDomainsRequest {
	r.verified = &verified
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIListDmDomainsRequest) Cursor(cursor string) *APIListDmDomainsRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIListDmDomainsRequest) Limit(limit int32) *APIListDmDomainsRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIListDmDomainsRequest) Execute() (*InlineResponse2004, *http.Response, error) {
	return r.APIService.ListDmDomainsExecute(r)
}

/*
ListDmDomains List domains

List all domains

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListDmDomainsRequest
*/
func (a *DmDomainsAPIService) ListDmDomains(ctx context.Context) APIListDmDomainsRequest {
	return APIListDmDomainsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListDmDomainsExecute executes the request
//  @return InlineResponse2004
func (a *DmDomainsAPIService) ListDmDomainsExecute(r APIListDmDomainsRequest) (*InlineResponse2004, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse2004
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmDomainsAPIService.ListDmDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/domains"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.fqdn != nil {
		localVarQueryParams.Add("fqdn", parameterToString(*r.fqdn, ""))
	}
	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.activated != nil {
		localVarQueryParams.Add("activated", parameterToString(*r.activated, ""))
	}
	if r.verified != nil {
		localVarQueryParams.Add("verified", parameterToString(*r.verified, ""))
	}
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

// APIUpdateDmDomainRequest represents a request for the resource.
type APIUpdateDmDomainRequest struct {
	ctx                  context.Context
	APIService           DmDomainsAPI
	domainId             string
	requestBodyForUpdate *RequestBodyForUpdate
}

// RequestBodyForUpdate returns a pointer to a request.
func (r *APIUpdateDmDomainRequest) RequestBodyForUpdate(requestBodyForUpdate RequestBodyForUpdate) *APIUpdateDmDomainRequest {
	r.requestBodyForUpdate = &requestBodyForUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateDmDomainRequest) Execute() (*SuccessfulResponseAsObject, *http.Response, error) {
	return r.APIService.UpdateDmDomainExecute(r)
}

/*
UpdateDmDomain Update a domain

Update a domain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainId
 @return APIUpdateDmDomainRequest
*/
func (a *DmDomainsAPIService) UpdateDmDomain(ctx context.Context, domainId string) APIUpdateDmDomainRequest {
	return APIUpdateDmDomainRequest{
		APIService: a,
		ctx:        ctx,
		domainId:   domainId,
	}
}

// UpdateDmDomainExecute executes the request
//  @return SuccessfulResponseAsObject
func (a *DmDomainsAPIService) UpdateDmDomainExecute(r APIUpdateDmDomainRequest) (*SuccessfulResponseAsObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SuccessfulResponseAsObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmDomainsAPIService.UpdateDmDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/domains/{domain_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"domain_id"+"}", gourl.PathEscape(parameterToString(r.domainId, "")))

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
	localVarPostBody = r.requestBodyForUpdate
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
