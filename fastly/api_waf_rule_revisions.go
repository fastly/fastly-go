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

// WafRuleRevisionsAPI defines an interface for interacting with the resource.
type WafRuleRevisionsAPI interface {

	/*
	GetWafRuleRevision Get a revision of a rule

	Get a specific rule revision.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @param wafRuleRevisionNumber Revision number.
	 @return APIGetWafRuleRevisionRequest

	Deprecated
	*/
	GetWafRuleRevision(ctx context.Context, wafRuleID string, wafRuleRevisionNumber int32) APIGetWafRuleRevisionRequest

	// GetWafRuleRevisionExecute executes the request
	//  @return WafRuleRevisionResponse
	// Deprecated
	GetWafRuleRevisionExecute(r APIGetWafRuleRevisionRequest) (*WafRuleRevisionResponse, *http.Response, error)

	/*
	ListWafRuleRevisions List revisions for a rule

	List all revisions for a specific rule. The `rule_id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIListWafRuleRevisionsRequest

	Deprecated
	*/
	ListWafRuleRevisions(ctx context.Context, wafRuleID string) APIListWafRuleRevisionsRequest

	// ListWafRuleRevisionsExecute executes the request
	//  @return WafRuleRevisionsResponse
	// Deprecated
	ListWafRuleRevisionsExecute(r APIListWafRuleRevisionsRequest) (*WafRuleRevisionsResponse, *http.Response, error)
}

// WafRuleRevisionsAPIService WafRuleRevisionsAPI service
type WafRuleRevisionsAPIService service

// APIGetWafRuleRevisionRequest represents a request for the resource.
type APIGetWafRuleRevisionRequest struct {
	ctx context.Context
	APIService WafRuleRevisionsAPI
	wafRuleID string
	wafRuleRevisionNumber int32
	include *string
}

// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_rule&#x60;, &#x60;vcl&#x60;, and &#x60;source&#x60;. The &#x60;vcl&#x60; and &#x60;source&#x60; relationships show the WAF VCL and corresponding ModSecurity source. These fields are blank unless the relationship is included. 
func (r *APIGetWafRuleRevisionRequest) Include(include string) *APIGetWafRuleRevisionRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetWafRuleRevisionRequest) Execute() (*WafRuleRevisionResponse, *http.Response, error) {
	return r.APIService.GetWafRuleRevisionExecute(r)
}

/*
GetWafRuleRevision Get a revision of a rule

Get a specific rule revision.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @param wafRuleRevisionNumber Revision number.
 @return APIGetWafRuleRevisionRequest

Deprecated
*/
func (a *WafRuleRevisionsAPIService) GetWafRuleRevision(ctx context.Context, wafRuleID string, wafRuleRevisionNumber int32) APIGetWafRuleRevisionRequest {
	return APIGetWafRuleRevisionRequest{
		APIService: a,
		ctx: ctx,
		wafRuleID: wafRuleID,
		wafRuleRevisionNumber: wafRuleRevisionNumber,
	}
}

// GetWafRuleRevisionExecute executes the request
//  @return WafRuleRevisionResponse
// Deprecated
func (a *WafRuleRevisionsAPIService) GetWafRuleRevisionExecute(r APIGetWafRuleRevisionRequest) (*WafRuleRevisionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafRuleRevisionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafRuleRevisionsAPIService.GetWafRuleRevision")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/rules/{waf_rule_id}/revisions/{waf_rule_revision_number}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_revision_number"+"}", gourl.PathEscape(parameterToString(r.wafRuleRevisionNumber, "")))

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

// APIListWafRuleRevisionsRequest represents a request for the resource.
type APIListWafRuleRevisionsRequest struct {
	ctx context.Context
	APIService WafRuleRevisionsAPI
	wafRuleID string
	pageNumber *int32
	pageSize *int32
	include *string
}

// PageNumber Current page.
func (r *APIListWafRuleRevisionsRequest) PageNumber(pageNumber int32) *APIListWafRuleRevisionsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafRuleRevisionsRequest) PageSize(pageSize int32) *APIListWafRuleRevisionsRequest {
	r.pageSize = &pageSize
	return r
}
// Include Include relationships. Optional.
func (r *APIListWafRuleRevisionsRequest) Include(include string) *APIListWafRuleRevisionsRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafRuleRevisionsRequest) Execute() (*WafRuleRevisionsResponse, *http.Response, error) {
	return r.APIService.ListWafRuleRevisionsExecute(r)
}

/*
ListWafRuleRevisions List revisions for a rule

List all revisions for a specific rule. The `rule_id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIListWafRuleRevisionsRequest

Deprecated
*/
func (a *WafRuleRevisionsAPIService) ListWafRuleRevisions(ctx context.Context, wafRuleID string) APIListWafRuleRevisionsRequest {
	return APIListWafRuleRevisionsRequest{
		APIService: a,
		ctx: ctx,
		wafRuleID: wafRuleID,
	}
}

// ListWafRuleRevisionsExecute executes the request
//  @return WafRuleRevisionsResponse
// Deprecated
func (a *WafRuleRevisionsAPIService) ListWafRuleRevisionsExecute(r APIListWafRuleRevisionsRequest) (*WafRuleRevisionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafRuleRevisionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafRuleRevisionsAPIService.ListWafRuleRevisions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/rules/{waf_rule_id}/revisions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))

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
