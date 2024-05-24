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

// WafRulesAPI defines an interface for interacting with the resource.
type WafRulesAPI interface {

	/*
	GetWafRule Get a rule

	Get a specific rule. The `id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIGetWafRuleRequest

	Deprecated
	*/
	GetWafRule(ctx context.Context, wafRuleID string) APIGetWafRuleRequest

	// GetWafRuleExecute executes the request
	//  @return WafRuleResponse
	// Deprecated
	GetWafRuleExecute(r APIGetWafRuleRequest) (*WafRuleResponse, *http.Response, error)

	/*
	ListWafRules List available WAF rules

	List all available WAF rules.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListWafRulesRequest

	Deprecated
	*/
	ListWafRules(ctx context.Context) APIListWafRulesRequest

	// ListWafRulesExecute executes the request
	//  @return WafRulesResponse
	// Deprecated
	ListWafRulesExecute(r APIListWafRulesRequest) (*WafRulesResponse, *http.Response, error)
}

// WafRulesAPIService WafRulesAPI service
type WafRulesAPIService service

// APIGetWafRuleRequest represents a request for the resource.
type APIGetWafRuleRequest struct {
	ctx context.Context
	APIService WafRulesAPI
	wafRuleID string
	include *string
}

// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_tags&#x60; and &#x60;waf_rule_revisions&#x60;. 
func (r *APIGetWafRuleRequest) Include(include string) *APIGetWafRuleRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetWafRuleRequest) Execute() (*WafRuleResponse, *http.Response, error) {
	return r.APIService.GetWafRuleExecute(r)
}

/*
GetWafRule Get a rule

Get a specific rule. The `id` provided can be the ModSecurity Rule ID or the Fastly generated rule ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIGetWafRuleRequest

Deprecated
*/
func (a *WafRulesAPIService) GetWafRule(ctx context.Context, wafRuleID string) APIGetWafRuleRequest {
	return APIGetWafRuleRequest{
		APIService: a,
		ctx: ctx,
		wafRuleID: wafRuleID,
	}
}

// GetWafRuleExecute executes the request
//  @return WafRuleResponse
// Deprecated
func (a *WafRulesAPIService) GetWafRuleExecute(r APIGetWafRuleRequest) (*WafRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafRulesAPIService.GetWafRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/rules/{waf_rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))

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

// APIListWafRulesRequest represents a request for the resource.
type APIListWafRulesRequest struct {
	ctx context.Context
	APIService WafRulesAPI
	filterModsecRuleID *string
	filterWafTagsName *string
	filterWafRuleRevisionsSource *string
	filterWafFirewallIDNotMatch *string
	pageNumber *int32
	pageSize *int32
	include *string
}

// FilterModsecRuleID Limit the returned rules to a specific ModSecurity rule ID.
func (r *APIListWafRulesRequest) FilterModsecRuleID(filterModsecRuleID string) *APIListWafRulesRequest {
	r.filterModsecRuleID = &filterModsecRuleID
	return r
}
// FilterWafTagsName Limit the returned rules to a set linked to a tag by name.
func (r *APIListWafRulesRequest) FilterWafTagsName(filterWafTagsName string) *APIListWafRulesRequest {
	r.filterWafTagsName = &filterWafTagsName
	return r
}
// FilterWafRuleRevisionsSource Limit the returned rules to a set linked to a source.
func (r *APIListWafRulesRequest) FilterWafRuleRevisionsSource(filterWafRuleRevisionsSource string) *APIListWafRulesRequest {
	r.filterWafRuleRevisionsSource = &filterWafRuleRevisionsSource
	return r
}
// FilterWafFirewallIDNotMatch Limit the returned rules to a set not included in the active firewall version for a firewall.
func (r *APIListWafRulesRequest) FilterWafFirewallIDNotMatch(filterWafFirewallIDNotMatch string) *APIListWafRulesRequest {
	r.filterWafFirewallIDNotMatch = &filterWafFirewallIDNotMatch
	return r
}
// PageNumber Current page.
func (r *APIListWafRulesRequest) PageNumber(pageNumber int32) *APIListWafRulesRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafRulesRequest) PageSize(pageSize int32) *APIListWafRulesRequest {
	r.pageSize = &pageSize
	return r
}
// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_tags&#x60; and &#x60;waf_rule_revisions&#x60;. 
func (r *APIListWafRulesRequest) Include(include string) *APIListWafRulesRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafRulesRequest) Execute() (*WafRulesResponse, *http.Response, error) {
	return r.APIService.ListWafRulesExecute(r)
}

/*
ListWafRules List available WAF rules

List all available WAF rules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListWafRulesRequest

Deprecated
*/
func (a *WafRulesAPIService) ListWafRules(ctx context.Context) APIListWafRulesRequest {
	return APIListWafRulesRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListWafRulesExecute executes the request
//  @return WafRulesResponse
// Deprecated
func (a *WafRulesAPIService) ListWafRulesExecute(r APIListWafRulesRequest) (*WafRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafRulesAPIService.ListWafRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterModsecRuleID != nil {
		localVarQueryParams.Add("filter[modsec_rule_id]", parameterToString(*r.filterModsecRuleID, ""))
	}
	if r.filterWafTagsName != nil {
		localVarQueryParams.Add("filter[waf_tags][name]", parameterToString(*r.filterWafTagsName, ""))
	}
	if r.filterWafRuleRevisionsSource != nil {
		localVarQueryParams.Add("filter[waf_rule_revisions][source]", parameterToString(*r.filterWafRuleRevisionsSource, ""))
	}
	if r.filterWafFirewallIDNotMatch != nil {
		localVarQueryParams.Add("filter[waf_firewall.id][not][match]", parameterToString(*r.filterWafFirewallIDNotMatch, ""))
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
