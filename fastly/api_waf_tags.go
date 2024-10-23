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
)

// Linger please
var (
	_ context.Context
)

// WafTagsAPI defines an interface for interacting with the resource.
type WafTagsAPI interface {

	/*
		ListWafTags List tags

		List all tags.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListWafTagsRequest

		Deprecated
	*/
	ListWafTags(ctx context.Context) APIListWafTagsRequest

	// ListWafTagsExecute executes the request
	//  @return WafTagsResponse
	// Deprecated
	ListWafTagsExecute(r APIListWafTagsRequest) (*WafTagsResponse, *http.Response, error)
}

// WafTagsAPIService WafTagsAPI service
type WafTagsAPIService service

// APIListWafTagsRequest represents a request for the resource.
type APIListWafTagsRequest struct {
	ctx        context.Context
	APIService WafTagsAPI
	filterName *string
	pageNumber *int32
	pageSize   *int32
	include    *string
}

// FilterName Limit the returned tags to a specific name.
func (r *APIListWafTagsRequest) FilterName(filterName string) *APIListWafTagsRequest {
	r.filterName = &filterName
	return r
}

// PageNumber Current page.
func (r *APIListWafTagsRequest) PageNumber(pageNumber int32) *APIListWafTagsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListWafTagsRequest) PageSize(pageSize int32) *APIListWafTagsRequest {
	r.pageSize = &pageSize
	return r
}

// Include Include relationships. Optional.
func (r *APIListWafTagsRequest) Include(include string) *APIListWafTagsRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafTagsRequest) Execute() (*WafTagsResponse, *http.Response, error) {
	return r.APIService.ListWafTagsExecute(r)
}

/*
ListWafTags List tags

List all tags.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListWafTagsRequest

Deprecated
*/
func (a *WafTagsAPIService) ListWafTags(ctx context.Context) APIListWafTagsRequest {
	return APIListWafTagsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListWafTagsExecute executes the request
//  @return WafTagsResponse
// Deprecated
func (a *WafTagsAPIService) ListWafTagsExecute(r APIListWafTagsRequest) (*WafTagsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *WafTagsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafTagsAPIService.ListWafTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterName != nil {
		localVarQueryParams.Add("filter[name]", parameterToString(*r.filterName, ""))
	}
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
