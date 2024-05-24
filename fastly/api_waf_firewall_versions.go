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

// WafFirewallVersionsAPI defines an interface for interacting with the resource.
type WafFirewallVersionsAPI interface {

	/*
	CloneWafFirewallVersion Clone a firewall version

	Clone a specific, existing firewall version into a new, draft firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @return APICloneWafFirewallVersionRequest

	Deprecated
	*/
	CloneWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APICloneWafFirewallVersionRequest

	// CloneWafFirewallVersionExecute executes the request
	//  @return WafFirewallVersionResponse
	// Deprecated
	CloneWafFirewallVersionExecute(r APICloneWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error)

	/*
	CreateWafFirewallVersion Create a firewall version

	Create a new, draft firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @return APICreateWafFirewallVersionRequest

	Deprecated
	*/
	CreateWafFirewallVersion(ctx context.Context, firewallID string) APICreateWafFirewallVersionRequest

	// CreateWafFirewallVersionExecute executes the request
	//  @return WafFirewallVersionResponse
	// Deprecated
	CreateWafFirewallVersionExecute(r APICreateWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error)

	/*
	DeployActivateWafFirewallVersion Deploy or activate a firewall version

	Deploy or activate a specific firewall version. If a firewall has been disabled, deploying a firewall version will automatically enable the firewall again.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @return APIDeployActivateWafFirewallVersionRequest

	Deprecated
	*/
	DeployActivateWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APIDeployActivateWafFirewallVersionRequest

	// DeployActivateWafFirewallVersionExecute executes the request
	//  @return map[string]any
	// Deprecated
	DeployActivateWafFirewallVersionExecute(r APIDeployActivateWafFirewallVersionRequest) (map[string]any, *http.Response, error)

	/*
	GetWafFirewallVersion Get a firewall version

	Get details about a specific firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @return APIGetWafFirewallVersionRequest

	Deprecated
	*/
	GetWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APIGetWafFirewallVersionRequest

	// GetWafFirewallVersionExecute executes the request
	//  @return WafFirewallVersionResponse
	// Deprecated
	GetWafFirewallVersionExecute(r APIGetWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error)

	/*
	ListWafFirewallVersions List firewall versions

	Get a list of firewall versions associated with a specific firewall.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @return APIListWafFirewallVersionsRequest

	Deprecated
	*/
	ListWafFirewallVersions(ctx context.Context, firewallID string) APIListWafFirewallVersionsRequest

	// ListWafFirewallVersionsExecute executes the request
	//  @return WafFirewallVersionsResponse
	// Deprecated
	ListWafFirewallVersionsExecute(r APIListWafFirewallVersionsRequest) (*WafFirewallVersionsResponse, *http.Response, error)

	/*
	UpdateWafFirewallVersion Update a firewall version

	Update a specific firewall version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a WAF Firewall.
	 @param firewallVersionNumber Integer identifying a WAF firewall version.
	 @return APIUpdateWafFirewallVersionRequest

	Deprecated
	*/
	UpdateWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APIUpdateWafFirewallVersionRequest

	// UpdateWafFirewallVersionExecute executes the request
	//  @return WafFirewallVersionResponse
	// Deprecated
	UpdateWafFirewallVersionExecute(r APIUpdateWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error)
}

// WafFirewallVersionsAPIService WafFirewallVersionsAPI service
type WafFirewallVersionsAPIService service

// APICloneWafFirewallVersionRequest represents a request for the resource.
type APICloneWafFirewallVersionRequest struct {
	ctx context.Context
	APIService WafFirewallVersionsAPI
	firewallID string
	firewallVersionNumber int32
}


// Execute calls the API using the request data configured.
func (r APICloneWafFirewallVersionRequest) Execute() (*WafFirewallVersionResponse, *http.Response, error) {
	return r.APIService.CloneWafFirewallVersionExecute(r)
}

/*
CloneWafFirewallVersion Clone a firewall version

Clone a specific, existing firewall version into a new, draft firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @return APICloneWafFirewallVersionRequest

Deprecated
*/
func (a *WafFirewallVersionsAPIService) CloneWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APICloneWafFirewallVersionRequest {
	return APICloneWafFirewallVersionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
	}
}

// CloneWafFirewallVersionExecute executes the request
//  @return WafFirewallVersionResponse
// Deprecated
func (a *WafFirewallVersionsAPIService) CloneWafFirewallVersionExecute(r APICloneWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallVersionsAPIService.CloneWafFirewallVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/clone"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))

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

// APICreateWafFirewallVersionRequest represents a request for the resource.
type APICreateWafFirewallVersionRequest struct {
	ctx context.Context
	APIService WafFirewallVersionsAPI
	firewallID string
	wafFirewallVersion *WafFirewallVersion
}

// WafFirewallVersion returns a pointer to a request.
func (r *APICreateWafFirewallVersionRequest) WafFirewallVersion(wafFirewallVersion WafFirewallVersion) *APICreateWafFirewallVersionRequest {
	r.wafFirewallVersion = &wafFirewallVersion
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateWafFirewallVersionRequest) Execute() (*WafFirewallVersionResponse, *http.Response, error) {
	return r.APIService.CreateWafFirewallVersionExecute(r)
}

/*
CreateWafFirewallVersion Create a firewall version

Create a new, draft firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @return APICreateWafFirewallVersionRequest

Deprecated
*/
func (a *WafFirewallVersionsAPIService) CreateWafFirewallVersion(ctx context.Context, firewallID string) APICreateWafFirewallVersionRequest {
	return APICreateWafFirewallVersionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// CreateWafFirewallVersionExecute executes the request
//  @return WafFirewallVersionResponse
// Deprecated
func (a *WafFirewallVersionsAPIService) CreateWafFirewallVersionExecute(r APICreateWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallVersionsAPIService.CreateWafFirewallVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions"
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
	localVarPostBody = r.wafFirewallVersion
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

// APIDeployActivateWafFirewallVersionRequest represents a request for the resource.
type APIDeployActivateWafFirewallVersionRequest struct {
	ctx context.Context
	APIService WafFirewallVersionsAPI
	firewallID string
	firewallVersionNumber int32
}


// Execute calls the API using the request data configured.
func (r APIDeployActivateWafFirewallVersionRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.DeployActivateWafFirewallVersionExecute(r)
}

/*
DeployActivateWafFirewallVersion Deploy or activate a firewall version

Deploy or activate a specific firewall version. If a firewall has been disabled, deploying a firewall version will automatically enable the firewall again.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @return APIDeployActivateWafFirewallVersionRequest

Deprecated
*/
func (a *WafFirewallVersionsAPIService) DeployActivateWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APIDeployActivateWafFirewallVersionRequest {
	return APIDeployActivateWafFirewallVersionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
	}
}

// DeployActivateWafFirewallVersionExecute executes the request
//  @return map[string]any
// Deprecated
func (a *WafFirewallVersionsAPIService) DeployActivateWafFirewallVersionExecute(r APIDeployActivateWafFirewallVersionRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallVersionsAPIService.DeployActivateWafFirewallVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/activate"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))

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

// APIGetWafFirewallVersionRequest represents a request for the resource.
type APIGetWafFirewallVersionRequest struct {
	ctx context.Context
	APIService WafFirewallVersionsAPI
	firewallID string
	firewallVersionNumber int32
	include *string
}

// Include Include relationships. Optional, comma-separated values. Permitted values: &#x60;waf_firewall&#x60; and &#x60;waf_active_rules&#x60;. 
func (r *APIGetWafFirewallVersionRequest) Include(include string) *APIGetWafFirewallVersionRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetWafFirewallVersionRequest) Execute() (*WafFirewallVersionResponse, *http.Response, error) {
	return r.APIService.GetWafFirewallVersionExecute(r)
}

/*
GetWafFirewallVersion Get a firewall version

Get details about a specific firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @return APIGetWafFirewallVersionRequest

Deprecated
*/
func (a *WafFirewallVersionsAPIService) GetWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APIGetWafFirewallVersionRequest {
	return APIGetWafFirewallVersionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
	}
}

// GetWafFirewallVersionExecute executes the request
//  @return WafFirewallVersionResponse
// Deprecated
func (a *WafFirewallVersionsAPIService) GetWafFirewallVersionExecute(r APIGetWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallVersionsAPIService.GetWafFirewallVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_version_number"+"}", gourl.PathEscape(parameterToString(r.firewallVersionNumber, "")))

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

// APIListWafFirewallVersionsRequest represents a request for the resource.
type APIListWafFirewallVersionsRequest struct {
	ctx context.Context
	APIService WafFirewallVersionsAPI
	firewallID string
	include *string
	pageNumber *int32
	pageSize *int32
}

// Include Include relationships. Optional.
func (r *APIListWafFirewallVersionsRequest) Include(include string) *APIListWafFirewallVersionsRequest {
	r.include = &include
	return r
}
// PageNumber Current page.
func (r *APIListWafFirewallVersionsRequest) PageNumber(pageNumber int32) *APIListWafFirewallVersionsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListWafFirewallVersionsRequest) PageSize(pageSize int32) *APIListWafFirewallVersionsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListWafFirewallVersionsRequest) Execute() (*WafFirewallVersionsResponse, *http.Response, error) {
	return r.APIService.ListWafFirewallVersionsExecute(r)
}

/*
ListWafFirewallVersions List firewall versions

Get a list of firewall versions associated with a specific firewall.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @return APIListWafFirewallVersionsRequest

Deprecated
*/
func (a *WafFirewallVersionsAPIService) ListWafFirewallVersions(ctx context.Context, firewallID string) APIListWafFirewallVersionsRequest {
	return APIListWafFirewallVersionsRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// ListWafFirewallVersionsExecute executes the request
//  @return WafFirewallVersionsResponse
// Deprecated
func (a *WafFirewallVersionsAPIService) ListWafFirewallVersionsExecute(r APIListWafFirewallVersionsRequest) (*WafFirewallVersionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallVersionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallVersionsAPIService.ListWafFirewallVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

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

// APIUpdateWafFirewallVersionRequest represents a request for the resource.
type APIUpdateWafFirewallVersionRequest struct {
	ctx context.Context
	APIService WafFirewallVersionsAPI
	firewallID string
	firewallVersionNumber int32
	wafFirewallVersion *WafFirewallVersion
}

// WafFirewallVersion returns a pointer to a request.
func (r *APIUpdateWafFirewallVersionRequest) WafFirewallVersion(wafFirewallVersion WafFirewallVersion) *APIUpdateWafFirewallVersionRequest {
	r.wafFirewallVersion = &wafFirewallVersion
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateWafFirewallVersionRequest) Execute() (*WafFirewallVersionResponse, *http.Response, error) {
	return r.APIService.UpdateWafFirewallVersionExecute(r)
}

/*
UpdateWafFirewallVersion Update a firewall version

Update a specific firewall version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a WAF Firewall.
 @param firewallVersionNumber Integer identifying a WAF firewall version.
 @return APIUpdateWafFirewallVersionRequest

Deprecated
*/
func (a *WafFirewallVersionsAPIService) UpdateWafFirewallVersion(ctx context.Context, firewallID string, firewallVersionNumber int32) APIUpdateWafFirewallVersionRequest {
	return APIUpdateWafFirewallVersionRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
		firewallVersionNumber: firewallVersionNumber,
	}
}

// UpdateWafFirewallVersionExecute executes the request
//  @return WafFirewallVersionResponse
// Deprecated
func (a *WafFirewallVersionsAPIService) UpdateWafFirewallVersionExecute(r APIUpdateWafFirewallVersionRequest) (*WafFirewallVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *WafFirewallVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WafFirewallVersionsAPIService.UpdateWafFirewallVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/firewalls/{firewall_id}/versions/{firewall_version_number}"
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
	localVarPostBody = r.wafFirewallVersion
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
