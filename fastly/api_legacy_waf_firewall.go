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

// LegacyWafFirewallAPI defines an interface for interacting with the resource.
type LegacyWafFirewallAPI interface {

	/*
	CreateLegacyWafFirewallService Create a firewall

	Create a firewall object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLegacyWafFirewallServiceRequest

	Deprecated
	*/
	CreateLegacyWafFirewallService(ctx context.Context, serviceID string, versionID int32) APICreateLegacyWafFirewallServiceRequest

	// CreateLegacyWafFirewallServiceExecute executes the request
	//  @return map[string]any
	// Deprecated
	CreateLegacyWafFirewallServiceExecute(r APICreateLegacyWafFirewallServiceRequest) (map[string]any, *http.Response, error)

	/*
	DisableLegacyWafFirewall Disable a firewall

	Disable a firewall for a particular service and version. This endpoint is intended to be used in an emergency. Disabling a firewall object for a specific service and version replaces your existing WAF ruleset with an empty ruleset. While disabled, your WAF ruleset will not be applied to your origin traffic. This endpoint is only available to users assigned the role of superuser or above. This is an asynchronous action. To check on the completion of this action, use the related link returned in the response to check on the Update Status of the action.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIDisableLegacyWafFirewallRequest

	Deprecated
	*/
	DisableLegacyWafFirewall(ctx context.Context, firewallID string) APIDisableLegacyWafFirewallRequest

	// DisableLegacyWafFirewallExecute executes the request
	//  @return map[string]any
	// Deprecated
	DisableLegacyWafFirewallExecute(r APIDisableLegacyWafFirewallRequest) (map[string]any, *http.Response, error)

	/*
	EnableLegacyWafFirewall Enable a firewall

	Re-enable a firewall object for a particular service and version after it has been disabled. This endpoint is intended to be used in an emergency. When a firewall object is re-enabled, a newly generated WAF ruleset VCL based on the current WAF configuration is used to replace the empty ruleset. This endpoint is only available to users assigned the role of superuser or above. This is an asynchronous action. To check on the completion of this action, use the related link returned in the response to check on the Update Status of the action.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIEnableLegacyWafFirewallRequest

	Deprecated
	*/
	EnableLegacyWafFirewall(ctx context.Context, firewallID string) APIEnableLegacyWafFirewallRequest

	// EnableLegacyWafFirewallExecute executes the request
	//  @return map[string]any
	// Deprecated
	EnableLegacyWafFirewallExecute(r APIEnableLegacyWafFirewallRequest) (map[string]any, *http.Response, error)

	/*
	GetLegacyWafFirewall Get a firewall object

	Get a specific firewall object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIGetLegacyWafFirewallRequest

	Deprecated
	*/
	GetLegacyWafFirewall(ctx context.Context, firewallID string) APIGetLegacyWafFirewallRequest

	// GetLegacyWafFirewallExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetLegacyWafFirewallExecute(r APIGetLegacyWafFirewallRequest) (map[string]any, *http.Response, error)

	/*
	GetLegacyWafFirewallService Get a firewall

	Get a specific firewall object.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIGetLegacyWafFirewallServiceRequest

	Deprecated
	*/
	GetLegacyWafFirewallService(ctx context.Context, serviceID string, versionID int32, firewallID string) APIGetLegacyWafFirewallServiceRequest

	// GetLegacyWafFirewallServiceExecute executes the request
	//  @return map[string]any
	// Deprecated
	GetLegacyWafFirewallServiceExecute(r APIGetLegacyWafFirewallServiceRequest) (map[string]any, *http.Response, error)

	/*
	ListLegacyWafFirewalls List active firewalls

	List all active firewall objects.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListLegacyWafFirewallsRequest

	Deprecated
	*/
	ListLegacyWafFirewalls(ctx context.Context) APIListLegacyWafFirewallsRequest

	// ListLegacyWafFirewallsExecute executes the request
	//  @return map[string]any
	// Deprecated
	ListLegacyWafFirewallsExecute(r APIListLegacyWafFirewallsRequest) (map[string]any, *http.Response, error)

	/*
	ListLegacyWafFirewallsService List firewalls

	List all firewall objects for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLegacyWafFirewallsServiceRequest

	Deprecated
	*/
	ListLegacyWafFirewallsService(ctx context.Context, serviceID string, versionID int32) APIListLegacyWafFirewallsServiceRequest

	// ListLegacyWafFirewallsServiceExecute executes the request
	//  @return map[string]any
	// Deprecated
	ListLegacyWafFirewallsServiceExecute(r APIListLegacyWafFirewallsServiceRequest) (map[string]any, *http.Response, error)

	/*
	UpdateLegacyWafFirewallService Update a firewall

	Update a firewall object for a particular service and version.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param firewallID Alphanumeric string identifying a Firewall.
	 @return APIUpdateLegacyWafFirewallServiceRequest

	Deprecated
	*/
	UpdateLegacyWafFirewallService(ctx context.Context, serviceID string, versionID int32, firewallID string) APIUpdateLegacyWafFirewallServiceRequest

	// UpdateLegacyWafFirewallServiceExecute executes the request
	//  @return map[string]any
	// Deprecated
	UpdateLegacyWafFirewallServiceExecute(r APIUpdateLegacyWafFirewallServiceRequest) (map[string]any, *http.Response, error)
}

// LegacyWafFirewallAPIService LegacyWafFirewallAPI service
type LegacyWafFirewallAPIService service

// APICreateLegacyWafFirewallServiceRequest represents a request for the resource.
type APICreateLegacyWafFirewallServiceRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	serviceID string
	versionID int32
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APICreateLegacyWafFirewallServiceRequest) RequestBody(requestBody map[string]map[string]any) *APICreateLegacyWafFirewallServiceRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLegacyWafFirewallServiceRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.CreateLegacyWafFirewallServiceExecute(r)
}

/*
CreateLegacyWafFirewallService Create a firewall

Create a firewall object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLegacyWafFirewallServiceRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) CreateLegacyWafFirewallService(ctx context.Context, serviceID string, versionID int32) APICreateLegacyWafFirewallServiceRequest {
	return APICreateLegacyWafFirewallServiceRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLegacyWafFirewallServiceExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) CreateLegacyWafFirewallServiceExecute(r APICreateLegacyWafFirewallServiceRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.CreateLegacyWafFirewallService")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/wafs"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
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

// APIDisableLegacyWafFirewallRequest represents a request for the resource.
type APIDisableLegacyWafFirewallRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	firewallID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIDisableLegacyWafFirewallRequest) RequestBody(requestBody map[string]map[string]any) *APIDisableLegacyWafFirewallRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIDisableLegacyWafFirewallRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.DisableLegacyWafFirewallExecute(r)
}

/*
DisableLegacyWafFirewall Disable a firewall

Disable a firewall for a particular service and version. This endpoint is intended to be used in an emergency. Disabling a firewall object for a specific service and version replaces your existing WAF ruleset with an empty ruleset. While disabled, your WAF ruleset will not be applied to your origin traffic. This endpoint is only available to users assigned the role of superuser or above. This is an asynchronous action. To check on the completion of this action, use the related link returned in the response to check on the Update Status of the action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIDisableLegacyWafFirewallRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) DisableLegacyWafFirewall(ctx context.Context, firewallID string) APIDisableLegacyWafFirewallRequest {
	return APIDisableLegacyWafFirewallRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// DisableLegacyWafFirewallExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) DisableLegacyWafFirewallExecute(r APIDisableLegacyWafFirewallRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.DisableLegacyWafFirewall")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/{firewall_id}/disable"
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

// APIEnableLegacyWafFirewallRequest represents a request for the resource.
type APIEnableLegacyWafFirewallRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	firewallID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIEnableLegacyWafFirewallRequest) RequestBody(requestBody map[string]map[string]any) *APIEnableLegacyWafFirewallRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIEnableLegacyWafFirewallRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.EnableLegacyWafFirewallExecute(r)
}

/*
EnableLegacyWafFirewall Enable a firewall

Re-enable a firewall object for a particular service and version after it has been disabled. This endpoint is intended to be used in an emergency. When a firewall object is re-enabled, a newly generated WAF ruleset VCL based on the current WAF configuration is used to replace the empty ruleset. This endpoint is only available to users assigned the role of superuser or above. This is an asynchronous action. To check on the completion of this action, use the related link returned in the response to check on the Update Status of the action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIEnableLegacyWafFirewallRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) EnableLegacyWafFirewall(ctx context.Context, firewallID string) APIEnableLegacyWafFirewallRequest {
	return APIEnableLegacyWafFirewallRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// EnableLegacyWafFirewallExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) EnableLegacyWafFirewallExecute(r APIEnableLegacyWafFirewallRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.EnableLegacyWafFirewall")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/{firewall_id}/enable"
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

// APIGetLegacyWafFirewallRequest represents a request for the resource.
type APIGetLegacyWafFirewallRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	firewallID string
	include *string
}

// Include Include relationships. Optional, comma separated values. Permitted values: &#x60;configuration_set&#x60;, &#x60;owasp&#x60;, &#x60;rules&#x60;, &#x60;rule_statuses&#x60;. 
func (r *APIGetLegacyWafFirewallRequest) Include(include string) *APIGetLegacyWafFirewallRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetLegacyWafFirewallRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetLegacyWafFirewallExecute(r)
}

/*
GetLegacyWafFirewall Get a firewall object

Get a specific firewall object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIGetLegacyWafFirewallRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) GetLegacyWafFirewall(ctx context.Context, firewallID string) APIGetLegacyWafFirewallRequest {
	return APIGetLegacyWafFirewallRequest{
		APIService: a,
		ctx: ctx,
		firewallID: firewallID,
	}
}

// GetLegacyWafFirewallExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) GetLegacyWafFirewallExecute(r APIGetLegacyWafFirewallRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.GetLegacyWafFirewall")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/{firewall_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

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

// APIGetLegacyWafFirewallServiceRequest represents a request for the resource.
type APIGetLegacyWafFirewallServiceRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	serviceID string
	versionID int32
	firewallID string
}


// Execute calls the API using the request data configured.
func (r APIGetLegacyWafFirewallServiceRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.GetLegacyWafFirewallServiceExecute(r)
}

/*
GetLegacyWafFirewallService Get a firewall

Get a specific firewall object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIGetLegacyWafFirewallServiceRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) GetLegacyWafFirewallService(ctx context.Context, serviceID string, versionID int32, firewallID string) APIGetLegacyWafFirewallServiceRequest {
	return APIGetLegacyWafFirewallServiceRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		firewallID: firewallID,
	}
}

// GetLegacyWafFirewallServiceExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) GetLegacyWafFirewallServiceExecute(r APIGetLegacyWafFirewallServiceRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.GetLegacyWafFirewallService")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/wafs/{firewall_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"firewall_id"+"}", gourl.PathEscape(parameterToString(r.firewallID, "")))

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

// APIListLegacyWafFirewallsRequest represents a request for the resource.
type APIListLegacyWafFirewallsRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	filterRulesRuleID *string
	pageNumber *int32
	pageSize *int32
	include *string
}

// FilterRulesRuleID Limit the returned firewalls to a specific rule ID.
func (r *APIListLegacyWafFirewallsRequest) FilterRulesRuleID(filterRulesRuleID string) *APIListLegacyWafFirewallsRequest {
	r.filterRulesRuleID = &filterRulesRuleID
	return r
}
// PageNumber Current page.
func (r *APIListLegacyWafFirewallsRequest) PageNumber(pageNumber int32) *APIListLegacyWafFirewallsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListLegacyWafFirewallsRequest) PageSize(pageSize int32) *APIListLegacyWafFirewallsRequest {
	r.pageSize = &pageSize
	return r
}
// Include Include relationships. Optional, comma separated values. Permitted values: &#x60;configuration_set&#x60;, &#x60;owasp&#x60;. 
func (r *APIListLegacyWafFirewallsRequest) Include(include string) *APIListLegacyWafFirewallsRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListLegacyWafFirewallsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListLegacyWafFirewallsExecute(r)
}

/*
ListLegacyWafFirewalls List active firewalls

List all active firewall objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListLegacyWafFirewallsRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) ListLegacyWafFirewalls(ctx context.Context) APIListLegacyWafFirewallsRequest {
	return APIListLegacyWafFirewallsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListLegacyWafFirewallsExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) ListLegacyWafFirewallsExecute(r APIListLegacyWafFirewallsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.ListLegacyWafFirewalls")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterRulesRuleID != nil {
		localVarQueryParams.Add("filter[rules][rule_id]", parameterToString(*r.filterRulesRuleID, ""))
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

// APIListLegacyWafFirewallsServiceRequest represents a request for the resource.
type APIListLegacyWafFirewallsServiceRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	serviceID string
	versionID int32
	pageNumber *int32
	pageSize *int32
	include *string
}

// PageNumber Current page.
func (r *APIListLegacyWafFirewallsServiceRequest) PageNumber(pageNumber int32) *APIListLegacyWafFirewallsServiceRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListLegacyWafFirewallsServiceRequest) PageSize(pageSize int32) *APIListLegacyWafFirewallsServiceRequest {
	r.pageSize = &pageSize
	return r
}
// Include Include relationships. Optional, comma separated values. Permitted values: &#x60;configuration_set&#x60;, &#x60;owasp&#x60;. 
func (r *APIListLegacyWafFirewallsServiceRequest) Include(include string) *APIListLegacyWafFirewallsServiceRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIListLegacyWafFirewallsServiceRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListLegacyWafFirewallsServiceExecute(r)
}

/*
ListLegacyWafFirewallsService List firewalls

List all firewall objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLegacyWafFirewallsServiceRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) ListLegacyWafFirewallsService(ctx context.Context, serviceID string, versionID int32) APIListLegacyWafFirewallsServiceRequest {
	return APIListLegacyWafFirewallsServiceRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLegacyWafFirewallsServiceExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) ListLegacyWafFirewallsServiceExecute(r APIListLegacyWafFirewallsServiceRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.ListLegacyWafFirewallsService")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/wafs"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

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

// APIUpdateLegacyWafFirewallServiceRequest represents a request for the resource.
type APIUpdateLegacyWafFirewallServiceRequest struct {
	ctx context.Context
	APIService LegacyWafFirewallAPI
	serviceID string
	versionID int32
	firewallID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIUpdateLegacyWafFirewallServiceRequest) RequestBody(requestBody map[string]map[string]any) *APIUpdateLegacyWafFirewallServiceRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLegacyWafFirewallServiceRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.UpdateLegacyWafFirewallServiceExecute(r)
}

/*
UpdateLegacyWafFirewallService Update a firewall

Update a firewall object for a particular service and version.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param firewallID Alphanumeric string identifying a Firewall.
 @return APIUpdateLegacyWafFirewallServiceRequest

Deprecated
*/
func (a *LegacyWafFirewallAPIService) UpdateLegacyWafFirewallService(ctx context.Context, serviceID string, versionID int32, firewallID string) APIUpdateLegacyWafFirewallServiceRequest {
	return APIUpdateLegacyWafFirewallServiceRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		firewallID: firewallID,
	}
}

// UpdateLegacyWafFirewallServiceExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafFirewallAPIService) UpdateLegacyWafFirewallServiceExecute(r APIUpdateLegacyWafFirewallServiceRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafFirewallAPIService.UpdateLegacyWafFirewallService")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/wafs/{firewall_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
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
