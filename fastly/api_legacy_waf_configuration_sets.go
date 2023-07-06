// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

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

// LegacyWafConfigurationSetsAPI defines an interface for interacting with the resource.
type LegacyWafConfigurationSetsAPI interface {

	/*
	ListWafConfigSets List configuration sets

	List all Configuration sets.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListWafConfigSetsRequest

	Deprecated
	*/
	ListWafConfigSets(ctx context.Context) APIListWafConfigSetsRequest

	// ListWafConfigSetsExecute executes the request
	//  @return map[string]any
	// Deprecated
	ListWafConfigSetsExecute(r APIListWafConfigSetsRequest) (map[string]any, *http.Response, error)

	/*
	ListWafsConfigSet List WAFs currently using a configuration set

	List the WAF objects currently using the specified configuration set.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configurationSetID Alphanumeric string identifying a WAF configuration set.
	 @return APIListWafsConfigSetRequest

	Deprecated
	*/
	ListWafsConfigSet(ctx context.Context, configurationSetID string) APIListWafsConfigSetRequest

	// ListWafsConfigSetExecute executes the request
	//  @return map[string]any
	// Deprecated
	ListWafsConfigSetExecute(r APIListWafsConfigSetRequest) (map[string]any, *http.Response, error)

	/*
	UseWafConfigSet Apply a configuration set to a WAF

	Update one or more WAF objects to use the specified configuration set.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configurationSetID Alphanumeric string identifying a WAF configuration set.
	 @return APIUseWafConfigSetRequest

	Deprecated
	*/
	UseWafConfigSet(ctx context.Context, configurationSetID string) APIUseWafConfigSetRequest

	// UseWafConfigSetExecute executes the request
	//  @return map[string]any
	// Deprecated
	UseWafConfigSetExecute(r APIUseWafConfigSetRequest) (map[string]any, *http.Response, error)
}

// LegacyWafConfigurationSetsAPIService LegacyWafConfigurationSetsAPI service
type LegacyWafConfigurationSetsAPIService service

// APIListWafConfigSetsRequest represents a request for the resource.
type APIListWafConfigSetsRequest struct {
	ctx context.Context
	APIService LegacyWafConfigurationSetsAPI
}


// Execute calls the API using the request data configured.
func (r APIListWafConfigSetsRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListWafConfigSetsExecute(r)
}

/*
ListWafConfigSets List configuration sets

List all Configuration sets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListWafConfigSetsRequest

Deprecated
*/
func (a *LegacyWafConfigurationSetsAPIService) ListWafConfigSets(ctx context.Context) APIListWafConfigSetsRequest {
	return APIListWafConfigSetsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListWafConfigSetsExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafConfigurationSetsAPIService) ListWafConfigSetsExecute(r APIListWafConfigSetsRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafConfigurationSetsAPIService.ListWafConfigSets")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/configuration_sets"

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

// APIListWafsConfigSetRequest represents a request for the resource.
type APIListWafsConfigSetRequest struct {
	ctx context.Context
	APIService LegacyWafConfigurationSetsAPI
	configurationSetID string
}


// Execute calls the API using the request data configured.
func (r APIListWafsConfigSetRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListWafsConfigSetExecute(r)
}

/*
ListWafsConfigSet List WAFs currently using a configuration set

List the WAF objects currently using the specified configuration set.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configurationSetID Alphanumeric string identifying a WAF configuration set.
 @return APIListWafsConfigSetRequest

Deprecated
*/
func (a *LegacyWafConfigurationSetsAPIService) ListWafsConfigSet(ctx context.Context, configurationSetID string) APIListWafsConfigSetRequest {
	return APIListWafsConfigSetRequest{
		APIService: a,
		ctx: ctx,
		configurationSetID: configurationSetID,
	}
}

// ListWafsConfigSetExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafConfigurationSetsAPIService) ListWafsConfigSetExecute(r APIListWafsConfigSetRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafConfigurationSetsAPIService.ListWafsConfigSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/configuration_sets/{configuration_set_id}/relationships/wafs"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"configuration_set_id"+"}", gourl.PathEscape(parameterToString(r.configurationSetID, "")))

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

// APIUseWafConfigSetRequest represents a request for the resource.
type APIUseWafConfigSetRequest struct {
	ctx context.Context
	APIService LegacyWafConfigurationSetsAPI
	configurationSetID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIUseWafConfigSetRequest) RequestBody(requestBody map[string]map[string]any) *APIUseWafConfigSetRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIUseWafConfigSetRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.UseWafConfigSetExecute(r)
}

/*
UseWafConfigSet Apply a configuration set to a WAF

Update one or more WAF objects to use the specified configuration set.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configurationSetID Alphanumeric string identifying a WAF configuration set.
 @return APIUseWafConfigSetRequest

Deprecated
*/
func (a *LegacyWafConfigurationSetsAPIService) UseWafConfigSet(ctx context.Context, configurationSetID string) APIUseWafConfigSetRequest {
	return APIUseWafConfigSetRequest{
		APIService: a,
		ctx: ctx,
		configurationSetID: configurationSetID,
	}
}

// UseWafConfigSetExecute executes the request
//  @return map[string]any
// Deprecated
func (a *LegacyWafConfigurationSetsAPIService) UseWafConfigSetExecute(r APIUseWafConfigSetRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWafConfigurationSetsAPIService.UseWafConfigSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wafs/configuration_sets/{configuration_set_id}/relationships/wafs"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"configuration_set_id"+"}", gourl.PathEscape(parameterToString(r.configurationSetID, "")))

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
