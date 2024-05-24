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

// LegacyWafRuleStatusAPI defines an interface for interacting with the resource.
type LegacyWafRuleStatusAPI interface {

	/*
	GetWafFirewallRuleStatus Get the status of a rule on a firewall

	Get a specific rule status object for a particular service, firewall, and rule.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIGetWafFirewallRuleStatusRequest

	Deprecated
	*/
	GetWafFirewallRuleStatus(ctx context.Context, serviceID string, firewallID string, wafRuleID string) APIGetWafFirewallRuleStatusRequest

	// GetWafFirewallRuleStatusExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetWafFirewallRuleStatusExecute(r APIGetWafFirewallRuleStatusRequest) (map[string]any, *http.Response, error)

	/*
	ListWafFirewallRuleStatuses List rule statuses

	List all rule statuses for a particular service and firewall.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIListWafFirewallRuleStatusesRequest

	Deprecated
	*/
	ListWafFirewallRuleStatuses(ctx context.Context, serviceID string, firewallID string) APIListWafFirewallRuleStatusesRequest

	// ListWafFirewallRuleStatusesExecute executes the request
	//  @return map[string]any
	// Deprecated
	ListWafFirewallRuleStatusesExecute(r APIListWafFirewallRuleStatusesRequest) (map[string]any, *http.Response, error)

	/*
	UpdateWafFirewallRuleStatus Update the status of a rule

	Update a rule status for a particular service, firewall, and rule.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @param wafRuleID Alphanumeric string identifying a WAF rule.
	 @return APIUpdateWafFirewallRuleStatusRequest

	Deprecated
	*/
	UpdateWafFirewallRuleStatus(ctx context.Context, serviceID string, firewallID string, wafRuleID string) APIUpdateWafFirewallRuleStatusRequest

	// UpdateWafFirewallRuleStatusExecute executes the request
	//  @return map[string]any
	// Deprecated
	UpdateWafFirewallRuleStatusExecute(r APIUpdateWafFirewallRuleStatusRequest) (map[string]any, *http.Response, error)

	/*
	UpdateWafFirewallRuleStatusesTag Create or update status of a tagged group of rules

	Create or update all rule statuses for a particular service and firewall, based on tag name. By default, only rule status for enabled rules (with status log or block) will be updated. To update rule statuses for disabled rules under the specified tag, use the force attribute.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIUpdateWafFirewallRuleStatusesTagRequest

	Deprecated
	*/
	UpdateWafFirewallRuleStatusesTag(ctx context.Context, serviceID string, firewallID string) APIUpdateWafFirewallRuleStatusesTagRequest

	// UpdateWafFirewallRuleStatusesTagExecute executes the request
	//  @return map[string]any
	// Deprecated
	UpdateWafFirewallRuleStatusesTagExecute(r APIUpdateWafFirewallRuleStatusesTagRequest) (map[string]any, *http.Response, error)
}

// LegacyWafRuleStatusAPIService LegacyWafRuleStatusAPI service
type LegacyWafRuleStatusAPIService service

// APIGetWafFirewallRuleStatusRequest represents a request for the resource.
type APIGetWafFirewallRuleStatusRequest struct {
	ctx context.Context
	APIService LegacyWafRuleStatusAPI
	serviceID string
	firewallID string
	wafRuleID string
}


// Execute calls the API using the request data configured.
func (r APIGetWafFirewallRuleStatusRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetWafFirewallRuleStatusExecute(r)
}

/*
GetWafFirewallRuleStatus Get the status of a rule on a firewall

Get a specific rule status object for a particular service, firewall, and rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param firewallID Alphanumeric string identifying a Firewall.
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIGetWafFirewallRuleStatusRequest

Deprecated
*/
func (a *LegacyWafRuleStatusAPIService) GetWafFirewallRuleStatus(ctx context.Context, serviceID string, firewallID string, wafRuleID string) APIGetWafFirewallRuleStatusRequest {
	return APIGetWafFirewallRuleStatusRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		firewallID: firewallID,
		wafRuleID: wafRuleID,
	}
}

// GetWafFirewallRuleStatusExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleStatusAPIService) GetWafFirewallRuleStatusExecute(r APIGetWafFirewallRuleStatusRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleStatusAPIService.GetWafFirewallRuleStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/wafs/{firewall_id}/rules/{waf_rule_id}/rule_status"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
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

// APIListWafFirewallRuleStatusesRequest represents a request for the resource.
type APIListWafFirewallRuleStatusesRequest struct {
	ctx context.Context
	APIService LegacyWafRuleStatusAPI
	serviceID string
	firewallID string
	filterStatus *string
	filterRuleMessage *string
	filterRuleRuleID *string
	filterRuleTags *string
	filterRuleTagsName *string
	include *string
	pageNumber *int32
	pageSize *int32
}

// FilterStatus Limit results to rule statuses with the specified status.
func (r *APIListWafFirewallRuleStatusesRequest) FilterStatus(filterStatus string) *APIListWafFirewallRuleStatusesRequest {
	r.filterStatus = &filterStatus
	return r
}
// FilterRuleMessage Limit results to rule statuses whose rules have the specified message.
func (r *APIListWafFirewallRuleStatusesRequest) FilterRuleMessage(filterRuleMessage string) *APIListWafFirewallRuleStatusesRequest {
	r.filterRuleMessage = &filterRuleMessage
	return r
}
// FilterRuleRuleID Limit results to rule statuses whose rules represent the specified ModSecurity rule_id.
func (r *APIListWafFirewallRuleStatusesRequest) FilterRuleRuleID(filterRuleRuleID string) *APIListWafFirewallRuleStatusesRequest {
	r.filterRuleRuleID = &filterRuleRuleID
	return r
}
// FilterRuleTags Limit results to rule statuses whose rules relate to the specified tag IDs.
func (r *APIListWafFirewallRuleStatusesRequest) FilterRuleTags(filterRuleTags string) *APIListWafFirewallRuleStatusesRequest {
	r.filterRuleTags = &filterRuleTags
	return r
}
// FilterRuleTagsName Limit results to rule statuses whose rules related to the named tags.
func (r *APIListWafFirewallRuleStatusesRequest) FilterRuleTagsName(filterRuleTagsName string) *APIListWafFirewallRuleStatusesRequest {
	r.filterRuleTagsName = &filterRuleTagsName
	return r
}
// Include Include relationships. Optional, comma separated values. Permitted values: &#x60;tags&#x60;. 
func (r *APIListWafFirewallRuleStatusesRequest) Include(include string) *APIListWafFirewallRuleStatusesRequest {
	r.include = &include
	return r
}
// PageNumber Current page.
func (r *APIListWafFirewallRuleStatusesRequest) PageNumber(pageNumber int32) *APIListWafFirewallRuleStatusesRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafFirewallRuleStatusesRequest) PageSize(pageSize int32) *APIListWafFirewallRuleStatusesRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafFirewallRuleStatusesRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListWafFirewallRuleStatusesExecute(r)
}

/*
ListWafFirewallRuleStatuses List rule statuses

List all rule statuses for a particular service and firewall.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIListWafFirewallRuleStatusesRequest

Deprecated
*/
func (a *LegacyWafRuleStatusAPIService) ListWafFirewallRuleStatuses(ctx context.Context, serviceID string, firewallID string) APIListWafFirewallRuleStatusesRequest {
	return APIListWafFirewallRuleStatusesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		firewallID: firewallID,
	}
}

// ListWafFirewallRuleStatusesExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleStatusAPIService) ListWafFirewallRuleStatusesExecute(r APIListWafFirewallRuleStatusesRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleStatusAPIService.ListWafFirewallRuleStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/wafs/{firewall_id}/rule_statuses"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterStatus != nil {
		localVarQueryParams.Add("filter[status]", parameterToString(*r.filterStatus, ""))
	}
	if r.filterRuleMessage != nil {
		localVarQueryParams.Add("filter[rule][message]", parameterToString(*r.filterRuleMessage, ""))
	}
	if r.filterRuleRuleID != nil {
		localVarQueryParams.Add("filter[rule][rule_id]", parameterToString(*r.filterRuleRuleID, ""))
	}
	if r.filterRuleTags != nil {
		localVarQueryParams.Add("filter[rule][tags]", parameterToString(*r.filterRuleTags, ""))
	}
	if r.filterRuleTagsName != nil {
		localVarQueryParams.Add("filter[rule][tags][name]", parameterToString(*r.filterRuleTagsName, ""))
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

// APIUpdateWafFirewallRuleStatusRequest represents a request for the resource.
type APIUpdateWafFirewallRuleStatusRequest struct {
	ctx context.Context
	APIService LegacyWafRuleStatusAPI
	serviceID string
	firewallID string
	wafRuleID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIUpdateWafFirewallRuleStatusRequest) RequestBody(requestBody map[string]map[string]any) *APIUpdateWafFirewallRuleStatusRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateWafFirewallRuleStatusRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.UpdateWafFirewallRuleStatusExecute(r)
}

/*
UpdateWafFirewallRuleStatus Update the status of a rule

Update a rule status for a particular service, firewall, and rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param firewallID Alphanumeric string identifying a Firewall.
 @param wafRuleID Alphanumeric string identifying a WAF rule.
 @return APIUpdateWafFirewallRuleStatusRequest

Deprecated
*/
func (a *LegacyWafRuleStatusAPIService) UpdateWafFirewallRuleStatus(ctx context.Context, serviceID string, firewallID string, wafRuleID string) APIUpdateWafFirewallRuleStatusRequest {
	return APIUpdateWafFirewallRuleStatusRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		firewallID: firewallID,
		wafRuleID: wafRuleID,
	}
}

// UpdateWafFirewallRuleStatusExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleStatusAPIService) UpdateWafFirewallRuleStatusExecute(r APIUpdateWafFirewallRuleStatusRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleStatusAPIService.UpdateWafFirewallRuleStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/wafs/{firewall_id}/rules/{waf_rule_id}/rule_status"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
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

// APIUpdateWafFirewallRuleStatusesTagRequest represents a request for the resource.
type APIUpdateWafFirewallRuleStatusesTagRequest struct {
	ctx context.Context
	APIService LegacyWafRuleStatusAPI
	serviceID string
	firewallID string
	name *string
	force *string
	requestBody *map[string]map[string]any
}

// Name The tag name to use to determine the set of rules to update. For example, OWASP or language-php.
func (r *APIUpdateWafFirewallRuleStatusesTagRequest) Name(name string) *APIUpdateWafFirewallRuleStatusesTagRequest {
	r.name = &name
	return r
}
// Force Whether or not to update rule statuses for disabled rules. Optional.
func (r *APIUpdateWafFirewallRuleStatusesTagRequest) Force(force string) *APIUpdateWafFirewallRuleStatusesTagRequest {
	r.force = &force
	return r
}
// RequestBody returns a pointer to a request.
func (r *APIUpdateWafFirewallRuleStatusesTagRequest) RequestBody(requestBody map[string]map[string]any) *APIUpdateWafFirewallRuleStatusesTagRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateWafFirewallRuleStatusesTagRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.UpdateWafFirewallRuleStatusesTagExecute(r)
}

/*
UpdateWafFirewallRuleStatusesTag Create or update status of a tagged group of rules

Create or update all rule statuses for a particular service and firewall, based on tag name. By default, only rule status for enabled rules (with status log or block) will be updated. To update rule statuses for disabled rules under the specified tag, use the force attribute.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIUpdateWafFirewallRuleStatusesTagRequest

Deprecated
*/
func (a *LegacyWafRuleStatusAPIService) UpdateWafFirewallRuleStatusesTag(ctx context.Context, serviceID string, firewallID string) APIUpdateWafFirewallRuleStatusesTagRequest {
	return APIUpdateWafFirewallRuleStatusesTagRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		firewallID: firewallID,
	}
}

// UpdateWafFirewallRuleStatusesTagExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafRuleStatusAPIService) UpdateWafFirewallRuleStatusesTagExecute(r APIUpdateWafFirewallRuleStatusesTagRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafRuleStatusAPIService.UpdateWafFirewallRuleStatusesTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/wafs/{firewall_id}/rule_statuses"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
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
