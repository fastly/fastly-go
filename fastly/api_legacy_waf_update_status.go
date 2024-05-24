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

// LegacyWafUpdateStatusAPI defines an interface for interacting with the resource.
type LegacyWafUpdateStatusAPI interface {

	/*
	GetWafUpdateStatus Get the status of a WAF update

	Get a specific update status object for a particular service and firewall object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @param updateStatusID Alphanumeric string identifying a WAF update status.
	 @return APIGetWafUpdateStatusRequest

	Deprecated
	*/
	GetWafUpdateStatus(ctx context.Context, serviceID string, firewallID string, updateStatusID string) APIGetWafUpdateStatusRequest

	// GetWafUpdateStatusExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetWafUpdateStatusExecute(r APIGetWafUpdateStatusRequest) (map[string]any, *http.Response, error)

	/*
	ListWafUpdateStatuses List update statuses

	List all update statuses for a particular service and firewall object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIListWafUpdateStatusesRequest

	Deprecated
	*/
	ListWafUpdateStatuses(ctx context.Context, serviceID string, firewallID string) APIListWafUpdateStatusesRequest

	// ListWafUpdateStatusesExecute executes the request
	//  @return map[string]any
	// Deprecated
	ListWafUpdateStatusesExecute(r APIListWafUpdateStatusesRequest) (map[string]any, *http.Response, error)
}

// LegacyWafUpdateStatusAPIService LegacyWafUpdateStatusAPI service
type LegacyWafUpdateStatusAPIService service

// APIGetWafUpdateStatusRequest represents a request for the resource.
type APIGetWafUpdateStatusRequest struct {
	ctx context.Context
	APIService LegacyWafUpdateStatusAPI
	serviceID string
	firewallID string
	updateStatusID string
}


// Execute calls the API using the request data configured.
func (r APIGetWafUpdateStatusRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetWafUpdateStatusExecute(r)
}

/*
GetWafUpdateStatus Get the status of a WAF update

Get a specific update status object for a particular service and firewall object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param firewallID Alphanumeric string identifying a Firewall.
 @param updateStatusID Alphanumeric string identifying a WAF update status.
 @return APIGetWafUpdateStatusRequest

Deprecated
*/
func (a *LegacyWafUpdateStatusAPIService) GetWafUpdateStatus(ctx context.Context, serviceID string, firewallID string, updateStatusID string) APIGetWafUpdateStatusRequest {
	return APIGetWafUpdateStatusRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		firewallID: firewallID,
		updateStatusID: updateStatusID,
	}
}

// GetWafUpdateStatusExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafUpdateStatusAPIService) GetWafUpdateStatusExecute(r APIGetWafUpdateStatusRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafUpdateStatusAPIService.GetWafUpdateStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/wafs/{firewall_id}/update_statuses/{update_status_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"update_status_id"+"}", gourl.PathEscape(parameterToString(r.updateStatusID, "")))

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

// APIListWafUpdateStatusesRequest represents a request for the resource.
type APIListWafUpdateStatusesRequest struct {
	ctx context.Context
	APIService LegacyWafUpdateStatusAPI
	serviceID string
	firewallID string
	pageNumber *int32
	pageSize *int32
	include *string
}

// PageNumber Current page.
func (r *APIListWafUpdateStatusesRequest) PageNumber(pageNumber int32) *APIListWafUpdateStatusesRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafUpdateStatusesRequest) PageSize(pageSize int32) *APIListWafUpdateStatusesRequest {
	r.pageSize = &pageSize
	return r
}
// Include Include relationships. Optional, comma separated values. Permitted values: &#x60;waf&#x60;. 
func (r *APIListWafUpdateStatusesRequest) Include(include string) *APIListWafUpdateStatusesRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafUpdateStatusesRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListWafUpdateStatusesExecute(r)
}

/*
ListWafUpdateStatuses List update statuses

List all update statuses for a particular service and firewall object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIListWafUpdateStatusesRequest

Deprecated
*/
func (a *LegacyWafUpdateStatusAPIService) ListWafUpdateStatuses(ctx context.Context, serviceID string, firewallID string) APIListWafUpdateStatusesRequest {
	return APIListWafUpdateStatusesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		firewallID: firewallID,
	}
}

// ListWafUpdateStatusesExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafUpdateStatusAPIService) ListWafUpdateStatusesExecute(r APIListWafUpdateStatusesRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafUpdateStatusAPIService.ListWafUpdateStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/wafs/{firewall_id}/update_statuses"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
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
