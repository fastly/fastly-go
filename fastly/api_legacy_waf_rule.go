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

// LegacyWafRuleAPI defines an interface for interacting with the resource.
type LegacyWafRuleAPI interface {

	/*
		GetLegacyWafFirewallRuleVcl Get VCL for a rule associated with a firewall

		Get associated VCL for a specific rule associated with a specific firewall.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param firewallID Alphanumeric string identifying a Firewall.
		 @param wafRuleID Alphanumeric string identifying a WAF rule.
		 @return APIGetLegacyWafFirewallRuleVclRequest

		Deprecated
	*/
	GetLegacyWafFirewallRuleVcl(ctx context.Context, firewallID string, wafRuleID string) APIGetLegacyWafFirewallRuleVclRequest

	// GetLegacyWafFirewallRuleVclExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetLegacyWafFirewallRuleVclExecute(r APIGetLegacyWafFirewallRuleVclRequest) (map[string]any, *http.Response, error)

	/*
		GetLegacyWafRule Get a rule

		Get a specific rule.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param wafRuleID Alphanumeric string identifying a WAF rule.
		 @return APIGetLegacyWafRuleRequest

		Deprecated
	*/
	GetLegacyWafRule(ctx context.Context, wafRuleID string) APIGetLegacyWafRuleRequest

	// GetLegacyWafRuleExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetLegacyWafRuleExecute(r APIGetLegacyWafRuleRequest) (map[string]any, *http.Response, error)

	/*
		GetLegacyWafRuleVcl Get VCL for a rule

		Get associated VCL for a specific rule.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param wafRuleID Alphanumeric string identifying a WAF rule.
		 @return APIGetLegacyWafRuleVclRequest

		Deprecated
	*/
	GetLegacyWafRuleVcl(ctx context.Context, wafRuleID string) APIGetLegacyWafRuleVclRequest

	// GetLegacyWafRuleVclExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetLegacyWafRuleVclExecute(r APIGetLegacyWafRuleVclRequest) (map[string]any, *http.Response, error)

	/*
		ListLegacyWafRules List rules in the latest configuration set

		List all rules in the latest configuration set.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListLegacyWafRulesRequest

		Deprecated
	*/
	ListLegacyWafRules(ctx context.Context) APIListLegacyWafRulesRequest

	// ListLegacyWafRulesExecute executes the request
	//  @return []map[string]any
	// Deprecated
	ListLegacyWafRulesExecute(r APIListLegacyWafRulesRequest) ([]map[string]any, *http.Response, error)
}

// LegacyWafRuleAPIService LegacyWafRuleAPI service
type LegacyWafRuleAPIService service

// APIGetLegacyWafFirewallRuleVclRequest represents a request for the resource.
type APIGetLegacyWafFirewallRuleVclRequest struct {
	ctx        context.Context
	APIService LegacyWafRuleAPI
	firewallID string
	wafRuleID  string
}

// Execute calls the API using the request data configured.
func (r APIGetLegacyWafFirewallRuleVclRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetLegacyWafFirewallRuleVclExecute(r)
}

/*
GetLegacyWafFirewallRuleVcl Get VCL for a rule associated with a firewall

Get associated VCL for a specific rule associated with a specific firewall.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a Firewall.
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIGetLegacyWafFirewallRuleVclRequest

Deprecated
*/
func (a *LegacyWafRuleAPIService) GetLegacyWafFirewallRuleVcl(ctx context.Context, firewallID string, wafRuleID string) APIGetLegacyWafFirewallRuleVclRequest {
	return APIGetLegacyWafFirewallRuleVclRequest{
		APIService: a,
		ctx:        ctx,
		firewallID: firewallID,
		wafRuleID:  wafRuleID,
	}
}

// GetLegacyWafFirewallRuleVclExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleAPIService) GetLegacyWafFirewallRuleVclExecute(r APIGetLegacyWafFirewallRuleVclRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleAPIService.GetLegacyWafFirewallRuleVcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/{firewall_id}/rules/{waf_rule_id}/vcl"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))

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

// APIGetLegacyWafRuleRequest represents a request for the resource.
type APIGetLegacyWafRuleRequest struct {
	ctx                      context.Context
	APIService               LegacyWafRuleAPI
	wafRuleID                string
	filterConfigurationSetID *string
	include                  *string
}

// FilterConfigurationSetID Optional. Limit rule to a specific configuration set or pass \&quot;all\&quot; to search all configuration sets, including stale ones.
func (r *APIGetLegacyWafRuleRequest) FilterConfigurationSetID(filterConfigurationSetID string) *APIGetLegacyWafRuleRequest {
	r.filterConfigurationSetID = &filterConfigurationSetID
	return r
}

// Include Include relationships. Optional. Comma separated values. Permitted values: &#x60;tags&#x60;, &#x60;rule_statuses&#x60;, &#x60;source&#x60;, and &#x60;vcl&#x60;.
func (r *APIGetLegacyWafRuleRequest) Include(include string) *APIGetLegacyWafRuleRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetLegacyWafRuleRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetLegacyWafRuleExecute(r)
}

/*
GetLegacyWafRule Get a rule

Get a specific rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIGetLegacyWafRuleRequest

Deprecated
*/
func (a *LegacyWafRuleAPIService) GetLegacyWafRule(ctx context.Context, wafRuleID string) APIGetLegacyWafRuleRequest {
	return APIGetLegacyWafRuleRequest{
		APIService: a,
		ctx:        ctx,
		wafRuleID:  wafRuleID,
	}
}

// GetLegacyWafRuleExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleAPIService) GetLegacyWafRuleExecute(r APIGetLegacyWafRuleRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleAPIService.GetLegacyWafRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/rules/{waf_rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterConfigurationSetID != nil {
		localVarQueryParams.Add("filter[configuration_set_id]", parameterToString(*r.filterConfigurationSetID, ""))
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

// APIGetLegacyWafRuleVclRequest represents a request for the resource.
type APIGetLegacyWafRuleVclRequest struct {
	ctx        context.Context
	APIService LegacyWafRuleAPI
	wafRuleID  string
}

// Execute calls the API using the request data configured.
func (r APIGetLegacyWafRuleVclRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetLegacyWafRuleVclExecute(r)
}

/*
GetLegacyWafRuleVcl Get VCL for a rule

Get associated VCL for a specific rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIGetLegacyWafRuleVclRequest

Deprecated
*/
func (a *LegacyWafRuleAPIService) GetLegacyWafRuleVcl(ctx context.Context, wafRuleID string) APIGetLegacyWafRuleVclRequest {
	return APIGetLegacyWafRuleVclRequest{
		APIService: a,
		ctx:        ctx,
		wafRuleID:  wafRuleID,
	}
}

// GetLegacyWafRuleVclExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleAPIService) GetLegacyWafRuleVclExecute(r APIGetLegacyWafRuleVclRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleAPIService.GetLegacyWafRuleVcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/rules/{waf_rule_id}/vcl"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))

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

// APIListLegacyWafRulesRequest represents a request for the resource.
type APIListLegacyWafRulesRequest struct {
	ctx                      context.Context
	APIService               LegacyWafRuleAPI
	filterRuleID             *string
	filterSeverity           *string
	filterTagsName           *string
	filterConfigurationSetID *string
	pageNumber               *int32
	pageSize                 *int32
	include                  *string
}

// FilterRuleID Limit the returned rules to a specific rule ID.
func (r *APIListLegacyWafRulesRequest) FilterRuleID(filterRuleID string) *APIListLegacyWafRulesRequest {
	r.filterRuleID = &filterRuleID
	return r
}

// FilterSeverity Limit the returned rules to a specific severity.
func (r *APIListLegacyWafRulesRequest) FilterSeverity(filterSeverity string) *APIListLegacyWafRulesRequest {
	r.filterSeverity = &filterSeverity
	return r
}

// FilterTagsName Limit the returned rules to a set linked to a tag by name.
func (r *APIListLegacyWafRulesRequest) FilterTagsName(filterTagsName string) *APIListLegacyWafRulesRequest {
	r.filterTagsName = &filterTagsName
	return r
}

// FilterConfigurationSetID Optional. Limit rules to specific configuration set or pass \&quot;all\&quot; to search all configuration sets, including stale ones.
func (r *APIListLegacyWafRulesRequest) FilterConfigurationSetID(filterConfigurationSetID string) *APIListLegacyWafRulesRequest {
	r.filterConfigurationSetID = &filterConfigurationSetID
	return r
}

// PageNumber Current page.
func (r *APIListLegacyWafRulesRequest) PageNumber(pageNumber int32) *APIListLegacyWafRulesRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListLegacyWafRulesRequest) PageSize(pageSize int32) *APIListLegacyWafRulesRequest {
	r.pageSize = &pageSize
	return r
}

// Include Include relationships. Optional. Comma separated values. Permitted values: &#x60;tags&#x60;, &#x60;rule_statuses&#x60;, and &#x60;source&#x60;.
func (r *APIListLegacyWafRulesRequest) Include(include string) *APIListLegacyWafRulesRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListLegacyWafRulesRequest) Execute() ([]map[string]any, *http.Response, error) {
	return r.APIService.ListLegacyWafRulesExecute(r)
}

/*
ListLegacyWafRules List rules in the latest configuration set

List all rules in the latest configuration set.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListLegacyWafRulesRequest

Deprecated
*/
func (a *LegacyWafRuleAPIService) ListLegacyWafRules(ctx context.Context) APIListLegacyWafRulesRequest {
	return APIListLegacyWafRulesRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListLegacyWafRulesExecute executes the request
//  @return []map[string]any
// Deprecated
func (a *LegacyWafRuleAPIService) ListLegacyWafRulesExecute(r APIListLegacyWafRulesRequest) ([]map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleAPIService.ListLegacyWafRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterRuleID != nil {
		localVarQueryParams.Add("filter[rule_id]", parameterToString(*r.filterRuleID, ""))
	}
	if r.filterSeverity != nil {
		localVarQueryParams.Add("filter[severity]", parameterToString(*r.filterSeverity, ""))
	}
	if r.filterTagsName != nil {
		localVarQueryParams.Add("filter[tags][name]", parameterToString(*r.filterTagsName, ""))
	}
	if r.filterConfigurationSetID != nil {
		localVarQueryParams.Add("filter[configuration_set_id]", parameterToString(*r.filterConfigurationSetID, ""))
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
