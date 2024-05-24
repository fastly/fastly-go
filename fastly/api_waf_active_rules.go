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

// WafActiveRulesAPI defines an interface for interacting with the resource.
type WafActiveRulesAPI interface {

	/*
	BulkDeleteWafActiveRules Delete multiple active rules from a WAF

	Delete many active rules on a particular firewall version using the active rule ID. Limited to 500 rules per request.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @return APIBulkDeleteWafActiveRulesRequest

	Deprecated
	*/
	BulkDeleteWafActiveRules(ctx context.Context, firewallID string, versionID int32) APIBulkDeleteWafActiveRulesRequest

	// BulkDeleteWafActiveRulesExecute executes the request
	// Deprecated
	BulkDeleteWafActiveRulesExecute(r APIBulkDeleteWafActiveRulesRequest) (*http.Response, error)

	/*
	BulkUpdateWafActiveRules Update multiple active rules

	Bulk update all active rules on a [firewall version](https://www.fastly.com/documentation/reference/api/waf/firewall-version/). This endpoint will not add new active rules, only update existing active rules.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @return APIBulkUpdateWafActiveRulesRequest

	Deprecated
	*/
	BulkUpdateWafActiveRules(ctx context.Context, firewallID string, versionID int32) APIBulkUpdateWafActiveRulesRequest

	// BulkUpdateWafActiveRulesExecute executes the request
	// Deprecated
	BulkUpdateWafActiveRulesExecute(r APIBulkUpdateWafActiveRulesRequest) (*http.Response, error)

	/*
	CreateWafActiveRule Add a rule to a WAF as an active rule

	Create an active rule for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @return APICreateWafActiveRuleRequest

	Deprecated
	*/
	CreateWafActiveRule(ctx context.Context, firewallID string, versionID int32) APICreateWafActiveRuleRequest

	// CreateWafActiveRuleExecute executes the request
	//  @return WafActiveRuleCreationResponse
	// Deprecated
	CreateWafActiveRuleExecute(r APICreateWafActiveRuleRequest) (*WafActiveRuleCreationResponse, *http.Response, error)

	/*
	CreateWafActiveRulesTag Create active rules by tag

	Create active rules by tag. This endpoint will create active rules using the latest revision available for each rule.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @param wafTagName Name of the tag.
	 @return APICreateWafActiveRulesTagRequest

	Deprecated
	*/
	CreateWafActiveRulesTag(ctx context.Context, firewallID string, versionID int32, wafTagName string) APICreateWafActiveRulesTagRequest

	// CreateWafActiveRulesTagExecute executes the request
	// Deprecated
	CreateWafActiveRulesTagExecute(r APICreateWafActiveRulesTagRequest) (*http.Response, error)

	/*
	DeleteWafActiveRule Delete an active rule

	Delete an active rule for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIDeleteWafActiveRuleRequest

	Deprecated
	*/
	DeleteWafActiveRule(ctx context.Context, firewallID string, versionID int32, wafRuleID string) APIDeleteWafActiveRuleRequest

	// DeleteWafActiveRuleExecute executes the request
	// Deprecated
	DeleteWafActiveRuleExecute(r APIDeleteWafActiveRuleRequest) (*http.Response, error)

	/*
	GetWafActiveRule Get an active WAF rule object

	Get a specific active rule object. Includes details of the rule revision associated with the active rule object by default.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIGetWafActiveRuleRequest

	Deprecated
	*/
	GetWafActiveRule(ctx context.Context, firewallID string, versionID int32, wafRuleID string) APIGetWafActiveRuleRequest

	// GetWafActiveRuleExecute executes the request
	//  @return WafActiveRuleResponse
	// Deprecated
	GetWafActiveRuleExecute(r APIGetWafActiveRuleRequest) (*WafActiveRuleResponse, *http.Response, error)

	/*
	ListWafActiveRules List active rules on a WAF

	List all active rules for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @return APIListWafActiveRulesRequest

	Deprecated
	*/
	ListWafActiveRules(ctx context.Context, firewallID string, versionID int32) APIListWafActiveRulesRequest

	// ListWafActiveRulesExecute executes the request
	//  @return WafActiveRulesResponse
	// Deprecated
	ListWafActiveRulesExecute(r APIListWafActiveRulesRequest) (*WafActiveRulesResponse, *http.Response, error)

	/*
	UpdateWafActiveRule Update an active rule

	Update an active rule's status for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param versionID Integer identifying a service version.
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIUpdateWafActiveRuleRequest

	Deprecated
	*/
	UpdateWafActiveRule(ctx context.Context, firewallID string, versionID int32, wafRuleID string) APIUpdateWafActiveRuleRequest

	// UpdateWafActiveRuleExecute executes the request
	//  @return WafActiveRuleResponse
	// Deprecated
	UpdateWafActiveRuleExecute(r APIUpdateWafActiveRuleRequest) (*WafActiveRuleResponse, *http.Response, error)
}

// WafActiveRulesAPIService WafActiveRulesAPI service
type WafActiveRulesAPIService service

// APIBulkDeleteWafActiveRulesRequest represents a request for the resource.
type APIBulkDeleteWafActiveRulesRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIBulkDeleteWafActiveRulesRequest) RequestBody(requestBody map[string]map[string]any) *APIBulkDeleteWafActiveRulesRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkDeleteWafActiveRulesRequest) Execute() (*http.Response, error) {
	return r.APIService.BulkDeleteWafActiveRulesExecute(r)
}

/*
BulkDeleteWafActiveRules Delete multiple active rules from a WAF

Delete many active rules on a particular firewall version using the active rule ID. Limited to 500 rules per request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @return APIBulkDeleteWafActiveRulesRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) BulkDeleteWafActiveRules(ctx context.Context, firewallID string, versionID int32) APIBulkDeleteWafActiveRulesRequest {
	return APIBulkDeleteWafActiveRulesRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
	}
}

// BulkDeleteWafActiveRulesExecute executes the request
// Deprecated
func (a *WafActiveRulesAPIService) BulkDeleteWafActiveRulesExecute(r APIBulkDeleteWafActiveRulesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.BulkDeleteWafActiveRules")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json; ext=bulk"}

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
	localVarPostBody = r.requestBody
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

// APIBulkUpdateWafActiveRulesRequest represents a request for the resource.
type APIBulkUpdateWafActiveRulesRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	body *WafActiveRuleData
}

// Body returns a pointer to a request.
func (r *APIBulkUpdateWafActiveRulesRequest) Body(body WafActiveRuleData) *APIBulkUpdateWafActiveRulesRequest {
	r.body = &body
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkUpdateWafActiveRulesRequest) Execute() (*http.Response, error) {
	return r.APIService.BulkUpdateWafActiveRulesExecute(r)
}

/*
BulkUpdateWafActiveRules Update multiple active rules

Bulk update all active rules on a [firewall version](https://www.fastly.com/documentation/reference/api/waf/firewall-version/). This endpoint will not add new active rules, only update existing active rules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @return APIBulkUpdateWafActiveRulesRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) BulkUpdateWafActiveRules(ctx context.Context, firewallID string, versionID int32) APIBulkUpdateWafActiveRulesRequest {
	return APIBulkUpdateWafActiveRulesRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
	}
}

// BulkUpdateWafActiveRulesExecute executes the request
// Deprecated
func (a *WafActiveRulesAPIService) BulkUpdateWafActiveRulesExecute(r APIBulkUpdateWafActiveRulesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.BulkUpdateWafActiveRules")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/bulk"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
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

// APICreateWafActiveRuleRequest represents a request for the resource.
type APICreateWafActiveRuleRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	wafActiveRule *WafActiveRule
}

// WafActiveRule returns a pointer to a request.
func (r *APICreateWafActiveRuleRequest) WafActiveRule(wafActiveRule WafActiveRule) *APICreateWafActiveRuleRequest {
	r.wafActiveRule = &wafActiveRule
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateWafActiveRuleRequest) Execute() (*WafActiveRuleCreationResponse, *http.Response, error) {
	return r.APIService.CreateWafActiveRuleExecute(r)
}

/*
CreateWafActiveRule Add a rule to a WAF as an active rule

Create an active rule for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @return APICreateWafActiveRuleRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) CreateWafActiveRule(ctx context.Context, firewallID string, versionID int32) APICreateWafActiveRuleRequest {
	return APICreateWafActiveRuleRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
	}
}

// CreateWafActiveRuleExecute executes the request
//  @return WafActiveRuleCreationResponse
// Deprecated
func (a *WafActiveRulesAPIService) CreateWafActiveRuleExecute(r APICreateWafActiveRuleRequest) (*WafActiveRuleCreationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafActiveRuleCreationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.CreateWafActiveRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json", "application/vnd.api+json; ext=bulk"}

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
	localVarPostBody = r.wafActiveRule
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

// APICreateWafActiveRulesTagRequest represents a request for the resource.
type APICreateWafActiveRulesTagRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	wafTagName string
	wafActiveRule *WafActiveRule
}

// WafActiveRule returns a pointer to a request.
func (r *APICreateWafActiveRulesTagRequest) WafActiveRule(wafActiveRule WafActiveRule) *APICreateWafActiveRulesTagRequest {
	r.wafActiveRule = &wafActiveRule
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateWafActiveRulesTagRequest) Execute() (*http.Response, error) {
	return r.APIService.CreateWafActiveRulesTagExecute(r)
}

/*
CreateWafActiveRulesTag Create active rules by tag

Create active rules by tag. This endpoint will create active rules using the latest revision available for each rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @param wafTagName Name of the tag.
 @return APICreateWafActiveRulesTagRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) CreateWafActiveRulesTag(ctx context.Context, firewallID string, versionID int32, wafTagName string) APICreateWafActiveRulesTagRequest {
	return APICreateWafActiveRulesTagRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
		wafTagName: wafTagName,
	}
}

// CreateWafActiveRulesTagExecute executes the request
// Deprecated
func (a *WafActiveRulesAPIService) CreateWafActiveRulesTagExecute(r APICreateWafActiveRulesTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.CreateWafActiveRulesTag")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/tags/{waf_tag_name}/active-rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_tag_name"+"}", gourl.PathEscape(parameterToString(r.wafTagName, "")))

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.wafActiveRule
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

// APIDeleteWafActiveRuleRequest represents a request for the resource.
type APIDeleteWafActiveRuleRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	wafRuleID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteWafActiveRuleRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteWafActiveRuleExecute(r)
}

/*
DeleteWafActiveRule Delete an active rule

Delete an active rule for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIDeleteWafActiveRuleRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) DeleteWafActiveRule(ctx context.Context, firewallID string, versionID int32, wafRuleID string) APIDeleteWafActiveRuleRequest {
	return APIDeleteWafActiveRuleRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
		wafRuleID: wafRuleID,
	}
}

// DeleteWafActiveRuleExecute executes the request
// Deprecated
func (a *WafActiveRulesAPIService) DeleteWafActiveRuleExecute(r APIDeleteWafActiveRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.DeleteWafActiveRule")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
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

// APIGetWafActiveRuleRequest represents a request for the resource.
type APIGetWafActiveRuleRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	wafRuleID string
	include *string
}

// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_rule_revision&#x60; and &#x60;waf_firewall_version&#x60;. 
func (r *APIGetWafActiveRuleRequest) Include(include string) *APIGetWafActiveRuleRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetWafActiveRuleRequest) Execute() (*WafActiveRuleResponse, *http.Response, error) {
	return r.APIService.GetWafActiveRuleExecute(r)
}

/*
GetWafActiveRule Get an active WAF rule object

Get a specific active rule object. Includes details of the rule revision associated with the active rule object by default.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIGetWafActiveRuleRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) GetWafActiveRule(ctx context.Context, firewallID string, versionID int32, wafRuleID string) APIGetWafActiveRuleRequest {
	return APIGetWafActiveRuleRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
		wafRuleID: wafRuleID,
	}
}

// GetWafActiveRuleExecute executes the request
//  @return WafActiveRuleResponse
// Deprecated
func (a *WafActiveRulesAPIService) GetWafActiveRuleExecute(r APIGetWafActiveRuleRequest) (*WafActiveRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafActiveRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.GetWafActiveRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
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

// APIListWafActiveRulesRequest represents a request for the resource.
type APIListWafActiveRulesRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	filterStatus *string
	filterWafRuleRevisionMessage *string
	filterWafRuleRevisionModsecRuleID *string
	filterOutdated *string
	include *string
	pageNumber *int32
	pageSize *int32
}

// FilterStatus Limit results to active rules with the specified status.
func (r *APIListWafActiveRulesRequest) FilterStatus(filterStatus string) *APIListWafActiveRulesRequest {
	r.filterStatus = &filterStatus
	return r
}
// FilterWafRuleRevisionMessage Limit results to active rules with the specified message.
func (r *APIListWafActiveRulesRequest) FilterWafRuleRevisionMessage(filterWafRuleRevisionMessage string) *APIListWafActiveRulesRequest {
	r.filterWafRuleRevisionMessage = &filterWafRuleRevisionMessage
	return r
}
// FilterWafRuleRevisionModsecRuleID Limit results to active rules that represent the specified ModSecurity modsec_rule_id.
func (r *APIListWafActiveRulesRequest) FilterWafRuleRevisionModsecRuleID(filterWafRuleRevisionModsecRuleID string) *APIListWafActiveRulesRequest {
	r.filterWafRuleRevisionModsecRuleID = &filterWafRuleRevisionModsecRuleID
	return r
}
// FilterOutdated Limit results to active rules referencing an outdated rule revision.
func (r *APIListWafActiveRulesRequest) FilterOutdated(filterOutdated string) *APIListWafActiveRulesRequest {
	r.filterOutdated = &filterOutdated
	return r
}
// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_rule_revision&#x60; and &#x60;waf_firewall_version&#x60;. 
func (r *APIListWafActiveRulesRequest) Include(include string) *APIListWafActiveRulesRequest {
	r.include = &include
	return r
}
// PageNumber Current page.
func (r *APIListWafActiveRulesRequest) PageNumber(pageNumber int32) *APIListWafActiveRulesRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafActiveRulesRequest) PageSize(pageSize int32) *APIListWafActiveRulesRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafActiveRulesRequest) Execute() (*WafActiveRulesResponse, *http.Response, error) {
	return r.APIService.ListWafActiveRulesExecute(r)
}

/*
ListWafActiveRules List active rules on a WAF

List all active rules for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @return APIListWafActiveRulesRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) ListWafActiveRules(ctx context.Context, firewallID string, versionID int32) APIListWafActiveRulesRequest {
	return APIListWafActiveRulesRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
	}
}

// ListWafActiveRulesExecute executes the request
//  @return WafActiveRulesResponse
// Deprecated
func (a *WafActiveRulesAPIService) ListWafActiveRulesExecute(r APIListWafActiveRulesRequest) (*WafActiveRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafActiveRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.ListWafActiveRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterStatus != nil {
		localVarQueryParams.Add("filter[status]", parameterToString(*r.filterStatus, ""))
	}
	if r.filterWafRuleRevisionMessage != nil {
		localVarQueryParams.Add("filter[waf_rule_revision][message]", parameterToString(*r.filterWafRuleRevisionMessage, ""))
	}
	if r.filterWafRuleRevisionModsecRuleID != nil {
		localVarQueryParams.Add("filter[waf_rule_revision][modsec_rule_id]", parameterToString(*r.filterWafRuleRevisionModsecRuleID, ""))
	}
	if r.filterOutdated != nil {
		localVarQueryParams.Add("filter[outdated]", parameterToString(*r.filterOutdated, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
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

// APIUpdateWafActiveRuleRequest represents a request for the resource.
type APIUpdateWafActiveRuleRequest struct {
	ctx context.Context
	APIService WafActiveRulesAPI
	firewallID string
	versionID int32
	wafRuleID string
	wafActiveRule *WafActiveRule
}

// WafActiveRule returns a pointer to a request.
func (r *APIUpdateWafActiveRuleRequest) WafActiveRule(wafActiveRule WafActiveRule) *APIUpdateWafActiveRuleRequest {
	r.wafActiveRule = &wafActiveRule
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateWafActiveRuleRequest) Execute() (*WafActiveRuleResponse, *http.Response, error) {
	return r.APIService.UpdateWafActiveRuleExecute(r)
}

/*
UpdateWafActiveRule Update an active rule

Update an active rule's status for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param versionID Integer identifying a service version.
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIUpdateWafActiveRuleRequest

Deprecated
*/
func (a *WafActiveRulesAPIService) UpdateWafActiveRule(ctx context.Context, firewallID string, versionID int32, wafRuleID string) APIUpdateWafActiveRuleRequest {
	return APIUpdateWafActiveRuleRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		versionID: versionID,
		wafRuleID: wafRuleID,
	}
}

// UpdateWafActiveRuleExecute executes the request
//  @return WafActiveRuleResponse
// Deprecated
func (a *WafActiveRulesAPIService) UpdateWafActiveRuleExecute(r APIUpdateWafActiveRuleRequest) (*WafActiveRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafActiveRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafActiveRulesAPIService.UpdateWafActiveRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"waf_rule_id"+"}", gourl.PathEscape(parameterToString(r.wafRuleID, "")))

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
	localVarPostBody = r.wafActiveRule
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
