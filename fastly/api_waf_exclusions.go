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

// WafExclusionsAPI defines an interface for interacting with the resource.
type WafExclusionsAPI interface {

	/*
	CreateWafRuleExclusion Create a WAF rule exclusion

	Create a WAF exclusion for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @return APICreateWafRuleExclusionRequest

	Deprecated
	*/
	CreateWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32) APICreateWafRuleExclusionRequest

	// CreateWafRuleExclusionExecute executes the request
	//  @return WafExclusionResponse
	// Deprecated
	CreateWafRuleExclusionExecute(r APICreateWafRuleExclusionRequest) (*WafExclusionResponse, *http.Response, error)

	/*
	DeleteWafRuleExclusion Delete a WAF rule exclusion

	Delete a WAF exclusion for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @param exclusionNumber A numeric ID identifying a WAF exclusion.
	 @return APIDeleteWafRuleExclusionRequest

	Deprecated
	*/
	DeleteWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32, exclusionNumber int32) APIDeleteWafRuleExclusionRequest

	// DeleteWafRuleExclusionExecute executes the request
	// Deprecated
	DeleteWafRuleExclusionExecute(r APIDeleteWafRuleExclusionRequest) (*http.Response, error)

	/*
	GetWafRuleExclusion Get a WAF rule exclusion

	Get a specific WAF exclusion object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @param exclusionNumber A numeric ID identifying a WAF exclusion.
	 @return APIGetWafRuleExclusionRequest

	Deprecated
	*/
	GetWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32, exclusionNumber int32) APIGetWafRuleExclusionRequest

	// GetWafRuleExclusionExecute executes the request
	//  @return WafExclusionResponse
	// Deprecated
	GetWafRuleExclusionExecute(r APIGetWafRuleExclusionRequest) (*WafExclusionResponse, *http.Response, error)

	/*
	ListWafRuleExclusions List WAF rule exclusions

	List all exclusions for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @return APIListWafRuleExclusionsRequest

	Deprecated
	*/
	ListWafRuleExclusions(ctx context.Context, firewallID string, firewallVersionNumber int32) APIListWafRuleExclusionsRequest

	// ListWafRuleExclusionsExecute executes the request
	//  @return WafExclusionsResponse
	// Deprecated
	ListWafRuleExclusionsExecute(r APIListWafRuleExclusionsRequest) (*WafExclusionsResponse, *http.Response, error)

	/*
	UpdateWafRuleExclusion Update a WAF rule exclusion

	Update a WAF exclusion for a particular firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @param exclusionNumber A numeric ID identifying a WAF exclusion.
	 @return APIUpdateWafRuleExclusionRequest

	Deprecated
	*/
	UpdateWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32, exclusionNumber int32) APIUpdateWafRuleExclusionRequest

	// UpdateWafRuleExclusionExecute executes the request
	//  @return WafExclusionResponse
	// Deprecated
	UpdateWafRuleExclusionExecute(r APIUpdateWafRuleExclusionRequest) (*WafExclusionResponse, *http.Response, error)
}

// WafExclusionsAPIService WafExclusionsAPI service
type WafExclusionsAPIService service

// APICreateWafRuleExclusionRequest represents a request for the resource.
type APICreateWafRuleExclusionRequest struct {
	ctx context.Context
	APIService WafExclusionsAPI
	firewallID string
	firewallVersionNumber int32
	wafExclusion *WafExclusion
}

// WafExclusion returns a pointer to a request.
func (r *APICreateWafRuleExclusionRequest) WafExclusion(wafExclusion WafExclusion) *APICreateWafRuleExclusionRequest {
	r.wafExclusion = &wafExclusion
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateWafRuleExclusionRequest) Execute() (*WafExclusionResponse, *http.Response, error) {
	return r.APIService.CreateWafRuleExclusionExecute(r)
}

/*
CreateWafRuleExclusion Create a WAF rule exclusion

Create a WAF exclusion for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @return APICreateWafRuleExclusionRequest

Deprecated
*/
func (a *WafExclusionsAPIService) CreateWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32) APICreateWafRuleExclusionRequest {
	return APICreateWafRuleExclusionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
	}
}

// CreateWafRuleExclusionExecute executes the request
//  @return WafExclusionResponse
// Deprecated
func (a *WafExclusionsAPIService) CreateWafRuleExclusionExecute(r APICreateWafRuleExclusionRequest) (*WafExclusionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafExclusionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafExclusionsAPIService.CreateWafRuleExclusion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))

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
	localVarPostBody = r.wafExclusion
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

// APIDeleteWafRuleExclusionRequest represents a request for the resource.
type APIDeleteWafRuleExclusionRequest struct {
	ctx context.Context
	APIService WafExclusionsAPI
	firewallID string
	firewallVersionNumber int32
	exclusionNumber int32
}


// Execute calls the API using the request data configured.
func (r APIDeleteWafRuleExclusionRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteWafRuleExclusionExecute(r)
}

/*
DeleteWafRuleExclusion Delete a WAF rule exclusion

Delete a WAF exclusion for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @param exclusionNumber A numeric ID identifying a WAF exclusion.
 @return APIDeleteWafRuleExclusionRequest

Deprecated
*/
func (a *WafExclusionsAPIService) DeleteWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32, exclusionNumber int32) APIDeleteWafRuleExclusionRequest {
	return APIDeleteWafRuleExclusionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
		exclusionNumber: exclusionNumber,
	}
}

// DeleteWafRuleExclusionExecute executes the request
// Deprecated
func (a *WafExclusionsAPIService) DeleteWafRuleExclusionExecute(r APIDeleteWafRuleExclusionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafExclusionsAPIService.DeleteWafRuleExclusion")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"exclusion_number"+"}", gourl.PathEscape(parameterToString(r.exclusionNumber, "")))

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

// APIGetWafRuleExclusionRequest represents a request for the resource.
type APIGetWafRuleExclusionRequest struct {
	ctx context.Context
	APIService WafExclusionsAPI
	firewallID string
	firewallVersionNumber int32
	exclusionNumber int32
}


// Execute calls the API using the request data configured.
func (r APIGetWafRuleExclusionRequest) Execute() (*WafExclusionResponse, *http.Response, error) {
	return r.APIService.GetWafRuleExclusionExecute(r)
}

/*
GetWafRuleExclusion Get a WAF rule exclusion

Get a specific WAF exclusion object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @param exclusionNumber A numeric ID identifying a WAF exclusion.
 @return APIGetWafRuleExclusionRequest

Deprecated
*/
func (a *WafExclusionsAPIService) GetWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32, exclusionNumber int32) APIGetWafRuleExclusionRequest {
	return APIGetWafRuleExclusionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
		exclusionNumber: exclusionNumber,
	}
}

// GetWafRuleExclusionExecute executes the request
//  @return WafExclusionResponse
// Deprecated
func (a *WafExclusionsAPIService) GetWafRuleExclusionExecute(r APIGetWafRuleExclusionRequest) (*WafExclusionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafExclusionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafExclusionsAPIService.GetWafRuleExclusion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"exclusion_number"+"}", gourl.PathEscape(parameterToString(r.exclusionNumber, "")))

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

// APIListWafRuleExclusionsRequest represents a request for the resource.
type APIListWafRuleExclusionsRequest struct {
	ctx context.Context
	APIService WafExclusionsAPI
	firewallID string
	firewallVersionNumber int32
	filterExclusionType *string
	filterName *string
	filterWafRulesModsecRuleID *int32
	pageNumber *int32
	pageSize *int32
	include *string
}

// FilterExclusionType Filters the results based on this exclusion type.
func (r *APIListWafRuleExclusionsRequest) FilterExclusionType(filterExclusionType string) *APIListWafRuleExclusionsRequest {
	r.filterExclusionType = &filterExclusionType
	return r
}
// FilterName Filters the results based on name.
func (r *APIListWafRuleExclusionsRequest) FilterName(filterName string) *APIListWafRuleExclusionsRequest {
	r.filterName = &filterName
	return r
}
// FilterWafRulesModsecRuleID Filters the results based on this ModSecurity rule ID.
func (r *APIListWafRuleExclusionsRequest) FilterWafRulesModsecRuleID(filterWafRulesModsecRuleID int32) *APIListWafRuleExclusionsRequest {
	r.filterWafRulesModsecRuleID = &filterWafRulesModsecRuleID
	return r
}
// PageNumber Current page.
func (r *APIListWafRuleExclusionsRequest) PageNumber(pageNumber int32) *APIListWafRuleExclusionsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafRuleExclusionsRequest) PageSize(pageSize int32) *APIListWafRuleExclusionsRequest {
	r.pageSize = &pageSize
	return r
}
// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_rules&#x60; and &#x60;waf_rule_revisions&#x60;. 
func (r *APIListWafRuleExclusionsRequest) Include(include string) *APIListWafRuleExclusionsRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafRuleExclusionsRequest) Execute() (*WafExclusionsResponse, *http.Response, error) {
	return r.APIService.ListWafRuleExclusionsExecute(r)
}

/*
ListWafRuleExclusions List WAF rule exclusions

List all exclusions for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @return APIListWafRuleExclusionsRequest

Deprecated
*/
func (a *WafExclusionsAPIService) ListWafRuleExclusions(ctx context.Context, firewallID string, firewallVersionNumber int32) APIListWafRuleExclusionsRequest {
	return APIListWafRuleExclusionsRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
	}
}

// ListWafRuleExclusionsExecute executes the request
//  @return WafExclusionsResponse
// Deprecated
func (a *WafExclusionsAPIService) ListWafRuleExclusionsExecute(r APIListWafRuleExclusionsRequest) (*WafExclusionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafExclusionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafExclusionsAPIService.ListWafRuleExclusions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterExclusionType != nil {
		localVarQueryParams.Add("filter[exclusion_type]", parameterToString(*r.filterExclusionType, ""))
	}
	if r.filterName != nil {
		localVarQueryParams.Add("filter[name]", parameterToString(*r.filterName, ""))
	}
	if r.filterWafRulesModsecRuleID != nil {
		localVarQueryParams.Add("filter[waf_rules.modsec_rule_id]", parameterToString(*r.filterWafRulesModsecRuleID, ""))
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

// APIUpdateWafRuleExclusionRequest represents a request for the resource.
type APIUpdateWafRuleExclusionRequest struct {
	ctx context.Context
	APIService WafExclusionsAPI
	firewallID string
	firewallVersionNumber int32
	exclusionNumber int32
	wafExclusion *WafExclusion
}

// WafExclusion returns a pointer to a request.
func (r *APIUpdateWafRuleExclusionRequest) WafExclusion(wafExclusion WafExclusion) *APIUpdateWafRuleExclusionRequest {
	r.wafExclusion = &wafExclusion
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateWafRuleExclusionRequest) Execute() (*WafExclusionResponse, *http.Response, error) {
	return r.APIService.UpdateWafRuleExclusionExecute(r)
}

/*
UpdateWafRuleExclusion Update a WAF rule exclusion

Update a WAF exclusion for a particular firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @param exclusionNumber A numeric ID identifying a WAF exclusion.
 @return APIUpdateWafRuleExclusionRequest

Deprecated
*/
func (a *WafExclusionsAPIService) UpdateWafRuleExclusion(ctx context.Context, firewallID string, firewallVersionNumber int32, exclusionNumber int32) APIUpdateWafRuleExclusionRequest {
	return APIUpdateWafRuleExclusionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
		exclusionNumber: exclusionNumber,
	}
}

// UpdateWafRuleExclusionExecute executes the request
//  @return WafExclusionResponse
// Deprecated
func (a *WafExclusionsAPIService) UpdateWafRuleExclusionExecute(r APIUpdateWafRuleExclusionRequest) (*WafExclusionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafExclusionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafExclusionsAPIService.UpdateWafRuleExclusion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"exclusion_number"+"}", gourl.PathEscape(parameterToString(r.exclusionNumber, "")))

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
	localVarPostBody = r.wafExclusion
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
