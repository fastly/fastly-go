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

// CacheSettingsAPI defines an interface for interacting with the resource.
type CacheSettingsAPI interface {

	/*
		CreateCacheSettings Create a cache settings object

		Create a cache settings object.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APICreateCacheSettingsRequest
	*/
	CreateCacheSettings(ctx context.Context, serviceId string, versionId int32) APICreateCacheSettingsRequest

	// CreateCacheSettingsExecute executes the request
	//  @return CacheSettingResponse
	CreateCacheSettingsExecute(r APICreateCacheSettingsRequest) (*CacheSettingResponse, *http.Response, error)

	/*
		DeleteCacheSettings Delete a cache settings object

		Delete a specific cache settings object.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param cacheSettingsName Name for the cache settings object.
		 @return APIDeleteCacheSettingsRequest
	*/
	DeleteCacheSettings(ctx context.Context, serviceId string, versionId int32, cacheSettingsName string) APIDeleteCacheSettingsRequest

	// DeleteCacheSettingsExecute executes the request
	//  @return InlineResponse200
	DeleteCacheSettingsExecute(r APIDeleteCacheSettingsRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetCacheSettings Get a cache settings object

		Get a specific cache settings object.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param cacheSettingsName Name for the cache settings object.
		 @return APIGetCacheSettingsRequest
	*/
	GetCacheSettings(ctx context.Context, serviceId string, versionId int32, cacheSettingsName string) APIGetCacheSettingsRequest

	// GetCacheSettingsExecute executes the request
	//  @return CacheSettingResponse
	GetCacheSettingsExecute(r APIGetCacheSettingsRequest) (*CacheSettingResponse, *http.Response, error)

	/*
		ListCacheSettings List cache settings objects

		Get a list of all cache settings for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APIListCacheSettingsRequest
	*/
	ListCacheSettings(ctx context.Context, serviceId string, versionId int32) APIListCacheSettingsRequest

	// ListCacheSettingsExecute executes the request
	//  @return []CacheSettingResponse
	ListCacheSettingsExecute(r APIListCacheSettingsRequest) ([]CacheSettingResponse, *http.Response, error)

	/*
		UpdateCacheSettings Update a cache settings object

		Update a specific cache settings object.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param cacheSettingsName Name for the cache settings object.
		 @return APIUpdateCacheSettingsRequest
	*/
	UpdateCacheSettings(ctx context.Context, serviceId string, versionId int32, cacheSettingsName string) APIUpdateCacheSettingsRequest

	// UpdateCacheSettingsExecute executes the request
	//  @return CacheSettingResponse
	UpdateCacheSettingsExecute(r APIUpdateCacheSettingsRequest) (*CacheSettingResponse, *http.Response, error)
}

// CacheSettingsAPIService CacheSettingsAPI service
type CacheSettingsAPIService service

// APICreateCacheSettingsRequest represents a request for the resource.
type APICreateCacheSettingsRequest struct {
	ctx            context.Context
	APIService     CacheSettingsAPI
	serviceId      string
	versionId      int32
	action         *string
	cacheCondition *string
	name           *string
	staleTtl       *string
	ttl            *string
}

// Action If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.
func (r *APICreateCacheSettingsRequest) Action(action string) *APICreateCacheSettingsRequest {
	r.action = &action
	return r
}

// CacheCondition Name of the cache condition controlling when this configuration applies.
func (r *APICreateCacheSettingsRequest) CacheCondition(cacheCondition string) *APICreateCacheSettingsRequest {
	r.cacheCondition = &cacheCondition
	return r
}

// Name Name for the cache settings object.
func (r *APICreateCacheSettingsRequest) Name(name string) *APICreateCacheSettingsRequest {
	r.name = &name
	return r
}

// StaleTtl Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;).
func (r *APICreateCacheSettingsRequest) StaleTtl(staleTtl string) *APICreateCacheSettingsRequest {
	r.staleTtl = &staleTtl
	return r
}

// Ttl Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;).
func (r *APICreateCacheSettingsRequest) Ttl(ttl string) *APICreateCacheSettingsRequest {
	r.ttl = &ttl
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateCacheSettingsRequest) Execute() (*CacheSettingResponse, *http.Response, error) {
	return r.APIService.CreateCacheSettingsExecute(r)
}

/*
CreateCacheSettings Create a cache settings object

Create a cache settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APICreateCacheSettingsRequest
*/
func (a *CacheSettingsAPIService) CreateCacheSettings(ctx context.Context, serviceId string, versionId int32) APICreateCacheSettingsRequest {
	return APICreateCacheSettingsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// CreateCacheSettingsExecute executes the request
//  @return CacheSettingResponse
func (a *CacheSettingsAPIService) CreateCacheSettingsExecute(r APICreateCacheSettingsRequest) (*CacheSettingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CacheSettingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsAPIService.CreateCacheSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/cache_settings"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.action != nil {
		localVarFormParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.cacheCondition != nil {
		localVarFormParams.Add("cache_condition", parameterToString(*r.cacheCondition, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.staleTtl != nil {
		localVarFormParams.Add("stale_ttl", parameterToString(*r.staleTtl, ""))
	}
	if r.ttl != nil {
		localVarFormParams.Add("ttl", parameterToString(*r.ttl, ""))
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

// APIDeleteCacheSettingsRequest represents a request for the resource.
type APIDeleteCacheSettingsRequest struct {
	ctx               context.Context
	APIService        CacheSettingsAPI
	serviceId         string
	versionId         int32
	cacheSettingsName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteCacheSettingsRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteCacheSettingsExecute(r)
}

/*
DeleteCacheSettings Delete a cache settings object

Delete a specific cache settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param cacheSettingsName Name for the cache settings object.
 @return APIDeleteCacheSettingsRequest
*/
func (a *CacheSettingsAPIService) DeleteCacheSettings(ctx context.Context, serviceId string, versionId int32, cacheSettingsName string) APIDeleteCacheSettingsRequest {
	return APIDeleteCacheSettingsRequest{
		APIService:        a,
		ctx:               ctx,
		serviceId:         serviceId,
		versionId:         versionId,
		cacheSettingsName: cacheSettingsName,
	}
}

// DeleteCacheSettingsExecute executes the request
//  @return InlineResponse200
func (a *CacheSettingsAPIService) DeleteCacheSettingsExecute(r APIDeleteCacheSettingsRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsAPIService.DeleteCacheSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cache_settings_name"+"}", gourl.PathEscape(parameterToString(r.cacheSettingsName, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

// APIGetCacheSettingsRequest represents a request for the resource.
type APIGetCacheSettingsRequest struct {
	ctx               context.Context
	APIService        CacheSettingsAPI
	serviceId         string
	versionId         int32
	cacheSettingsName string
}

// Execute calls the API using the request data configured.
func (r APIGetCacheSettingsRequest) Execute() (*CacheSettingResponse, *http.Response, error) {
	return r.APIService.GetCacheSettingsExecute(r)
}

/*
GetCacheSettings Get a cache settings object

Get a specific cache settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param cacheSettingsName Name for the cache settings object.
 @return APIGetCacheSettingsRequest
*/
func (a *CacheSettingsAPIService) GetCacheSettings(ctx context.Context, serviceId string, versionId int32, cacheSettingsName string) APIGetCacheSettingsRequest {
	return APIGetCacheSettingsRequest{
		APIService:        a,
		ctx:               ctx,
		serviceId:         serviceId,
		versionId:         versionId,
		cacheSettingsName: cacheSettingsName,
	}
}

// GetCacheSettingsExecute executes the request
//  @return CacheSettingResponse
func (a *CacheSettingsAPIService) GetCacheSettingsExecute(r APIGetCacheSettingsRequest) (*CacheSettingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CacheSettingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsAPIService.GetCacheSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cache_settings_name"+"}", gourl.PathEscape(parameterToString(r.cacheSettingsName, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

// APIListCacheSettingsRequest represents a request for the resource.
type APIListCacheSettingsRequest struct {
	ctx        context.Context
	APIService CacheSettingsAPI
	serviceId  string
	versionId  int32
}

// Execute calls the API using the request data configured.
func (r APIListCacheSettingsRequest) Execute() ([]CacheSettingResponse, *http.Response, error) {
	return r.APIService.ListCacheSettingsExecute(r)
}

/*
ListCacheSettings List cache settings objects

Get a list of all cache settings for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APIListCacheSettingsRequest
*/
func (a *CacheSettingsAPIService) ListCacheSettings(ctx context.Context, serviceId string, versionId int32) APIListCacheSettingsRequest {
	return APIListCacheSettingsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// ListCacheSettingsExecute executes the request
//  @return []CacheSettingResponse
func (a *CacheSettingsAPIService) ListCacheSettingsExecute(r APIListCacheSettingsRequest) ([]CacheSettingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []CacheSettingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsAPIService.ListCacheSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/cache_settings"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

// APIUpdateCacheSettingsRequest represents a request for the resource.
type APIUpdateCacheSettingsRequest struct {
	ctx               context.Context
	APIService        CacheSettingsAPI
	serviceId         string
	versionId         int32
	cacheSettingsName string
	action            *string
	cacheCondition    *string
	name              *string
	staleTtl          *string
	ttl               *string
}

// Action If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.
func (r *APIUpdateCacheSettingsRequest) Action(action string) *APIUpdateCacheSettingsRequest {
	r.action = &action
	return r
}

// CacheCondition Name of the cache condition controlling when this configuration applies.
func (r *APIUpdateCacheSettingsRequest) CacheCondition(cacheCondition string) *APIUpdateCacheSettingsRequest {
	r.cacheCondition = &cacheCondition
	return r
}

// Name Name for the cache settings object.
func (r *APIUpdateCacheSettingsRequest) Name(name string) *APIUpdateCacheSettingsRequest {
	r.name = &name
	return r
}

// StaleTtl Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;).
func (r *APIUpdateCacheSettingsRequest) StaleTtl(staleTtl string) *APIUpdateCacheSettingsRequest {
	r.staleTtl = &staleTtl
	return r
}

// Ttl Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;).
func (r *APIUpdateCacheSettingsRequest) Ttl(ttl string) *APIUpdateCacheSettingsRequest {
	r.ttl = &ttl
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateCacheSettingsRequest) Execute() (*CacheSettingResponse, *http.Response, error) {
	return r.APIService.UpdateCacheSettingsExecute(r)
}

/*
UpdateCacheSettings Update a cache settings object

Update a specific cache settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param cacheSettingsName Name for the cache settings object.
 @return APIUpdateCacheSettingsRequest
*/
func (a *CacheSettingsAPIService) UpdateCacheSettings(ctx context.Context, serviceId string, versionId int32, cacheSettingsName string) APIUpdateCacheSettingsRequest {
	return APIUpdateCacheSettingsRequest{
		APIService:        a,
		ctx:               ctx,
		serviceId:         serviceId,
		versionId:         versionId,
		cacheSettingsName: cacheSettingsName,
	}
}

// UpdateCacheSettingsExecute executes the request
//  @return CacheSettingResponse
func (a *CacheSettingsAPIService) UpdateCacheSettingsExecute(r APIUpdateCacheSettingsRequest) (*CacheSettingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CacheSettingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsAPIService.UpdateCacheSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cache_settings_name"+"}", gourl.PathEscape(parameterToString(r.cacheSettingsName, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.action != nil {
		localVarFormParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.cacheCondition != nil {
		localVarFormParams.Add("cache_condition", parameterToString(*r.cacheCondition, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.staleTtl != nil {
		localVarFormParams.Add("stale_ttl", parameterToString(*r.staleTtl, ""))
	}
	if r.ttl != nil {
		localVarFormParams.Add("ttl", parameterToString(*r.ttl, ""))
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
