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

// WafFirewallsAPI defines an interface for interacting with the resource.
type WafFirewallsAPI interface {

	/*
	CreateWafFirewall Create a firewall

	Create a firewall object for a particular service and service version using a defined `prefetch_condition` and `response`. If the `prefetch_condition` or the `response` is missing from the request body, Fastly will generate a default object on your service.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateWafFirewallRequest

	Deprecated
	*/
	CreateWafFirewall(ctx context.Context) APICreateWafFirewallRequest

	// CreateWafFirewallExecute executes the request
	//  @return WafFirewallResponse
	// Deprecated
	CreateWafFirewallExecute(r APICreateWafFirewallRequest) (*WafFirewallResponse, *http.Response, error)

	/*
	DeleteWafFirewall Delete a firewall

	Delete the firewall object for a particular service and service version.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @return APIDeleteWafFirewallRequest

	Deprecated
	*/
	DeleteWafFirewall(ctx context.Context, firewallID string) APIDeleteWafFirewallRequest

	// DeleteWafFirewallExecute executes the request
	// Deprecated
	DeleteWafFirewallExecute(r APIDeleteWafFirewallRequest) (*http.Response, error)

	/*
	GetWafFirewall Get a firewall

	Get a specific firewall object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @return APIGetWafFirewallRequest

	Deprecated
	*/
	GetWafFirewall(ctx context.Context, firewallID string) APIGetWafFirewallRequest

	// GetWafFirewallExecute executes the request
	//  @return WafFirewallResponse
	// Deprecated
	GetWafFirewallExecute(r APIGetWafFirewallRequest) (*WafFirewallResponse, *http.Response, error)

	/*
	ListWafFirewalls List firewalls

	List all firewall objects.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListWafFirewallsRequest

	Deprecated
	*/
	ListWafFirewalls(ctx context.Context) APIListWafFirewallsRequest

	// ListWafFirewallsExecute executes the request
	//  @return WafFirewallsResponse
	// Deprecated
	ListWafFirewallsExecute(r APIListWafFirewallsRequest) (*WafFirewallsResponse, *http.Response, error)

	/*
	UpdateWafFirewall Update a firewall

	Update a firewall object for a particular service and service version. Specifying a `service_version_number` is required.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @return APIUpdateWafFirewallRequest

	Deprecated
	*/
	UpdateWafFirewall(ctx context.Context, firewallID string) APIUpdateWafFirewallRequest

	// UpdateWafFirewallExecute executes the request
	//  @return WafFirewallResponse
	// Deprecated
	UpdateWafFirewallExecute(r APIUpdateWafFirewallRequest) (*WafFirewallResponse, *http.Response, error)
}

// WafFirewallsAPIService WafFirewallsAPI service
type WafFirewallsAPIService service

// APICreateWafFirewallRequest represents a request for the resource.
type APICreateWafFirewallRequest struct {
	ctx context.Context
	APIService WafFirewallsAPI
	wafFirewall *WafFirewall
}

// WafFirewall returns a pointer to a request.
func (r *APICreateWafFirewallRequest) WafFirewall(wafFirewall WafFirewall) *APICreateWafFirewallRequest {
	r.wafFirewall = &wafFirewall
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateWafFirewallRequest) Execute() (*WafFirewallResponse, *http.Response, error) {
	return r.APIService.CreateWafFirewallExecute(r)
}

/*
CreateWafFirewall Create a firewall

Create a firewall object for a particular service and service version using a defined `prefetch_condition` and `response`. If the `prefetch_condition` or the `response` is missing from the request body, Fastly will generate a default object on your service.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateWafFirewallRequest

Deprecated
*/
func (a *WafFirewallsAPIService) CreateWafFirewall(ctx context.Context) APICreateWafFirewallRequest {
	return APICreateWafFirewallRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateWafFirewallExecute executes the request
//  @return WafFirewallResponse
// Deprecated
func (a *WafFirewallsAPIService) CreateWafFirewallExecute(r APICreateWafFirewallRequest) (*WafFirewallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallsAPIService.CreateWafFirewall")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.wafFirewall
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

// APIDeleteWafFirewallRequest represents a request for the resource.
type APIDeleteWafFirewallRequest struct {
	ctx context.Context
	APIService WafFirewallsAPI
	firewallID string
	wafFirewall *WafFirewall
}

// WafFirewall returns a pointer to a request.
func (r *APIDeleteWafFirewallRequest) WafFirewall(wafFirewall WafFirewall) *APIDeleteWafFirewallRequest {
	r.wafFirewall = &wafFirewall
	return r
}

// Execute calls the API using the request data configured.
func (r APIDeleteWafFirewallRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteWafFirewallExecute(r)
}

/*
DeleteWafFirewall Delete a firewall

Delete the firewall object for a particular service and service version.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @return APIDeleteWafFirewallRequest

Deprecated
*/
func (a *WafFirewallsAPIService) DeleteWafFirewall(ctx context.Context, firewallID string) APIDeleteWafFirewallRequest {
	return APIDeleteWafFirewallRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// DeleteWafFirewallExecute executes the request
// Deprecated
func (a *WafFirewallsAPIService) DeleteWafFirewallExecute(r APIDeleteWafFirewallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallsAPIService.DeleteWafFirewall")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

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
	localVarPostBody = r.wafFirewall
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

// APIGetWafFirewallRequest represents a request for the resource.
type APIGetWafFirewallRequest struct {
	ctx context.Context
	APIService WafFirewallsAPI
	firewallID string
	filterServiceVersionNumber *string
	include *string
}

// FilterServiceVersionNumber Limit the results returned to a specific service version.
func (r *APIGetWafFirewallRequest) FilterServiceVersionNumber(filterServiceVersionNumber string) *APIGetWafFirewallRequest {
	r.filterServiceVersionNumber = &filterServiceVersionNumber
	return r
}
// Include Include related objects. Optional.
func (r *APIGetWafFirewallRequest) Include(include string) *APIGetWafFirewallRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetWafFirewallRequest) Execute() (*WafFirewallResponse, *http.Response, error) {
	return r.APIService.GetWafFirewallExecute(r)
}

/*
GetWafFirewall Get a firewall

Get a specific firewall object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @return APIGetWafFirewallRequest

Deprecated
*/
func (a *WafFirewallsAPIService) GetWafFirewall(ctx context.Context, firewallID string) APIGetWafFirewallRequest {
	return APIGetWafFirewallRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// GetWafFirewallExecute executes the request
//  @return WafFirewallResponse
// Deprecated
func (a *WafFirewallsAPIService) GetWafFirewallExecute(r APIGetWafFirewallRequest) (*WafFirewallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallsAPIService.GetWafFirewall")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterServiceVersionNumber != nil {
		localVarQueryParams.Add("filter[service_version_number]", parameterToString(*r.filterServiceVersionNumber, ""))
	}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIListWafFirewallsRequest represents a request for the resource.
type APIListWafFirewallsRequest struct {
	ctx context.Context
	APIService WafFirewallsAPI
	pageNumber *int32
	pageSize *int32
	filterServiceID *string
	filterServiceVersionNumber *string
	include *string
}

// PageNumber Current page.
func (r *APIListWafFirewallsRequest) PageNumber(pageNumber int32) *APIListWafFirewallsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafFirewallsRequest) PageSize(pageSize int32) *APIListWafFirewallsRequest {
	r.pageSize = &pageSize
	return r
}
// FilterServiceID Limit the results returned to a specific service.
func (r *APIListWafFirewallsRequest) FilterServiceID(filterServiceID string) *APIListWafFirewallsRequest {
	r.filterServiceID = &filterServiceID
	return r
}
// FilterServiceVersionNumber Limit the results returned to a specific service version.
func (r *APIListWafFirewallsRequest) FilterServiceVersionNumber(filterServiceVersionNumber string) *APIListWafFirewallsRequest {
	r.filterServiceVersionNumber = &filterServiceVersionNumber
	return r
}
// Include Include related objects. Optional.
func (r *APIListWafFirewallsRequest) Include(include string) *APIListWafFirewallsRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafFirewallsRequest) Execute() (*WafFirewallsResponse, *http.Response, error) {
	return r.APIService.ListWafFirewallsExecute(r)
}

/*
ListWafFirewalls List firewalls

List all firewall objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListWafFirewallsRequest

Deprecated
*/
func (a *WafFirewallsAPIService) ListWafFirewalls(ctx context.Context) APIListWafFirewallsRequest {
	return APIListWafFirewallsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListWafFirewallsExecute executes the request
//  @return WafFirewallsResponse
// Deprecated
func (a *WafFirewallsAPIService) ListWafFirewallsExecute(r APIListWafFirewallsRequest) (*WafFirewallsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallsAPIService.ListWafFirewalls")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.filterServiceID != nil {
		localVarQueryParams.Add("filter[service_id]", parameterToString(*r.filterServiceID, ""))
	}
	if r.filterServiceVersionNumber != nil {
		localVarQueryParams.Add("filter[service_version_number]", parameterToString(*r.filterServiceVersionNumber, ""))
	}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIUpdateWafFirewallRequest represents a request for the resource.
type APIUpdateWafFirewallRequest struct {
	ctx context.Context
	APIService WafFirewallsAPI
	firewallID string
	wafFirewall *WafFirewall
}

// WafFirewall returns a pointer to a request.
func (r *APIUpdateWafFirewallRequest) WafFirewall(wafFirewall WafFirewall) *APIUpdateWafFirewallRequest {
	r.wafFirewall = &wafFirewall
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateWafFirewallRequest) Execute() (*WafFirewallResponse, *http.Response, error) {
	return r.APIService.UpdateWafFirewallExecute(r)
}

/*
UpdateWafFirewall Update a firewall

Update a firewall object for a particular service and service version. Specifying a `service_version_number` is required.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @return APIUpdateWafFirewallRequest

Deprecated
*/
func (a *WafFirewallsAPIService) UpdateWafFirewall(ctx context.Context, firewallID string) APIUpdateWafFirewallRequest {
	return APIUpdateWafFirewallRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// UpdateWafFirewallExecute executes the request
//  @return WafFirewallResponse
// Deprecated
func (a *WafFirewallsAPIService) UpdateWafFirewallExecute(r APIUpdateWafFirewallRequest) (*WafFirewallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallsAPIService.UpdateWafFirewall")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.wafFirewall
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
