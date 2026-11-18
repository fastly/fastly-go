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

// ConfigStoreAPI defines an interface for interacting with the resource.
type ConfigStoreAPI interface {

	/*
		CreateConfigStore Create a config store

		Create a config store.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateConfigStoreRequest
	*/
	CreateConfigStore(ctx context.Context) APICreateConfigStoreRequest

	// CreateConfigStoreExecute executes the request
	//  @return ConfigStoreResponse
	CreateConfigStoreExecute(r APICreateConfigStoreRequest) (*ConfigStoreResponse, *http.Response, error)

	/*
		DeleteConfigStore Delete a config store

		Delete a config store.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configStoreID An alphanumeric string identifying the config store.
		 @return APIDeleteConfigStoreRequest
	*/
	DeleteConfigStore(ctx context.Context, configStoreID string) APIDeleteConfigStoreRequest

	// DeleteConfigStoreExecute executes the request
	//  @return InlineResponse200
	DeleteConfigStoreExecute(r APIDeleteConfigStoreRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetConfigStore Describe a config store

		Describe a config store by its identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configStoreID An alphanumeric string identifying the config store.
		 @return APIGetConfigStoreRequest
	*/
	GetConfigStore(ctx context.Context, configStoreID string) APIGetConfigStoreRequest

	// GetConfigStoreExecute executes the request
	//  @return ConfigStoreResponse
	GetConfigStoreExecute(r APIGetConfigStoreRequest) (*ConfigStoreResponse, *http.Response, error)

	/*
		GetConfigStoreInfo Get config store metadata

		Retrieve metadata for a single config store.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configStoreID An alphanumeric string identifying the config store.
		 @return APIGetConfigStoreInfoRequest
	*/
	GetConfigStoreInfo(ctx context.Context, configStoreID string) APIGetConfigStoreInfoRequest

	// GetConfigStoreInfoExecute executes the request
	//  @return ConfigStoreInfoResponse
	GetConfigStoreInfoExecute(r APIGetConfigStoreInfoRequest) (*ConfigStoreInfoResponse, *http.Response, error)

	/*
		ListConfigStoreServices List linked services

		List services linked to a config store

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configStoreID An alphanumeric string identifying the config store.
		 @return APIListConfigStoreServicesRequest
	*/
	ListConfigStoreServices(ctx context.Context, configStoreID string) APIListConfigStoreServicesRequest

	// ListConfigStoreServicesExecute executes the request
	//  @return map[string]any
	ListConfigStoreServicesExecute(r APIListConfigStoreServicesRequest) (map[string]any, *http.Response, error)

	/*
		ListConfigStores List config stores

		List config stores.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListConfigStoresRequest
	*/
	ListConfigStores(ctx context.Context) APIListConfigStoresRequest

	// ListConfigStoresExecute executes the request
	//  @return []ConfigStoreResponse
	ListConfigStoresExecute(r APIListConfigStoresRequest) ([]ConfigStoreResponse, *http.Response, error)

	/*
		UpdateConfigStore Update a config store

		Update a config store.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configStoreID An alphanumeric string identifying the config store.
		 @return APIUpdateConfigStoreRequest
	*/
	UpdateConfigStore(ctx context.Context, configStoreID string) APIUpdateConfigStoreRequest

	// UpdateConfigStoreExecute executes the request
	//  @return ConfigStoreResponse
	UpdateConfigStoreExecute(r APIUpdateConfigStoreRequest) (*ConfigStoreResponse, *http.Response, error)
}

// ConfigStoreAPIService ConfigStoreAPI service
type ConfigStoreAPIService service

// APICreateConfigStoreRequest represents a request for the resource.
type APICreateConfigStoreRequest struct {
	ctx        context.Context
	APIService ConfigStoreAPI
	name       *string
}

// Name The name of the config store.
func (r *APICreateConfigStoreRequest) Name(name string) *APICreateConfigStoreRequest {
	r.name = &name
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateConfigStoreRequest) Execute() (*ConfigStoreResponse, *http.Response, error) {
	return r.APIService.CreateConfigStoreExecute(r)
}

/*
CreateConfigStore Create a config store

Create a config store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateConfigStoreRequest
*/
func (a *ConfigStoreAPIService) CreateConfigStore(ctx context.Context) APICreateConfigStoreRequest {
	return APICreateConfigStoreRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateConfigStoreExecute executes the request
//  @return ConfigStoreResponse
func (a *ConfigStoreAPIService) CreateConfigStoreExecute(r APICreateConfigStoreRequest) (*ConfigStoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConfigStoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.CreateConfigStore")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config"

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
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
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

// APIDeleteConfigStoreRequest represents a request for the resource.
type APIDeleteConfigStoreRequest struct {
	ctx           context.Context
	APIService    ConfigStoreAPI
	configStoreID string
}

// Execute calls the API using the request data configured.
func (r APIDeleteConfigStoreRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteConfigStoreExecute(r)
}

/*
DeleteConfigStore Delete a config store

Delete a config store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIDeleteConfigStoreRequest
*/
func (a *ConfigStoreAPIService) DeleteConfigStore(ctx context.Context, configStoreID string) APIDeleteConfigStoreRequest {
	return APIDeleteConfigStoreRequest{
		APIService:    a,
		ctx:           ctx,
		configStoreID: configStoreID,
	}
}

// DeleteConfigStoreExecute executes the request
//  @return InlineResponse200
func (a *ConfigStoreAPIService) DeleteConfigStoreExecute(r APIDeleteConfigStoreRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.DeleteConfigStore")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))

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

// APIGetConfigStoreRequest represents a request for the resource.
type APIGetConfigStoreRequest struct {
	ctx           context.Context
	APIService    ConfigStoreAPI
	configStoreID string
}

// Execute calls the API using the request data configured.
func (r APIGetConfigStoreRequest) Execute() (*ConfigStoreResponse, *http.Response, error) {
	return r.APIService.GetConfigStoreExecute(r)
}

/*
GetConfigStore Describe a config store

Describe a config store by its identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIGetConfigStoreRequest
*/
func (a *ConfigStoreAPIService) GetConfigStore(ctx context.Context, configStoreID string) APIGetConfigStoreRequest {
	return APIGetConfigStoreRequest{
		APIService:    a,
		ctx:           ctx,
		configStoreID: configStoreID,
	}
}

// GetConfigStoreExecute executes the request
//  @return ConfigStoreResponse
func (a *ConfigStoreAPIService) GetConfigStoreExecute(r APIGetConfigStoreRequest) (*ConfigStoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConfigStoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.GetConfigStore")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))

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

// APIGetConfigStoreInfoRequest represents a request for the resource.
type APIGetConfigStoreInfoRequest struct {
	ctx           context.Context
	APIService    ConfigStoreAPI
	configStoreID string
}

// Execute calls the API using the request data configured.
func (r APIGetConfigStoreInfoRequest) Execute() (*ConfigStoreInfoResponse, *http.Response, error) {
	return r.APIService.GetConfigStoreInfoExecute(r)
}

/*
GetConfigStoreInfo Get config store metadata

Retrieve metadata for a single config store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIGetConfigStoreInfoRequest
*/
func (a *ConfigStoreAPIService) GetConfigStoreInfo(ctx context.Context, configStoreID string) APIGetConfigStoreInfoRequest {
	return APIGetConfigStoreInfoRequest{
		APIService:    a,
		ctx:           ctx,
		configStoreID: configStoreID,
	}
}

// GetConfigStoreInfoExecute executes the request
//  @return ConfigStoreInfoResponse
func (a *ConfigStoreAPIService) GetConfigStoreInfoExecute(r APIGetConfigStoreInfoRequest) (*ConfigStoreInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConfigStoreInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.GetConfigStoreInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/info"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))

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

// APIListConfigStoreServicesRequest represents a request for the resource.
type APIListConfigStoreServicesRequest struct {
	ctx           context.Context
	APIService    ConfigStoreAPI
	configStoreID string
}

// Execute calls the API using the request data configured.
func (r APIListConfigStoreServicesRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.ListConfigStoreServicesExecute(r)
}

/*
ListConfigStoreServices List linked services

List services linked to a config store

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIListConfigStoreServicesRequest
*/
func (a *ConfigStoreAPIService) ListConfigStoreServices(ctx context.Context, configStoreID string) APIListConfigStoreServicesRequest {
	return APIListConfigStoreServicesRequest{
		APIService:    a,
		ctx:           ctx,
		configStoreID: configStoreID,
	}
}

// ListConfigStoreServicesExecute executes the request
//  @return map[string]any
func (a *ConfigStoreAPIService) ListConfigStoreServicesExecute(r APIListConfigStoreServicesRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.ListConfigStoreServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/services"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))

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

// APIListConfigStoresRequest represents a request for the resource.
type APIListConfigStoresRequest struct {
	ctx        context.Context
	APIService ConfigStoreAPI
	name       *string
}

// Name Returns a one-element array containing the details for the named config store.
func (r *APIListConfigStoresRequest) Name(name string) *APIListConfigStoresRequest {
	r.name = &name
	return r
}

// Execute calls the API using the request data configured.
func (r APIListConfigStoresRequest) Execute() ([]ConfigStoreResponse, *http.Response, error) {
	return r.APIService.ListConfigStoresExecute(r)
}

/*
ListConfigStores List config stores

List config stores.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListConfigStoresRequest
*/
func (a *ConfigStoreAPIService) ListConfigStores(ctx context.Context) APIListConfigStoresRequest {
	return APIListConfigStoresRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListConfigStoresExecute executes the request
//  @return []ConfigStoreResponse
func (a *ConfigStoreAPIService) ListConfigStoresExecute(r APIListConfigStoresRequest) ([]ConfigStoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []ConfigStoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.ListConfigStores")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
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

// APIUpdateConfigStoreRequest represents a request for the resource.
type APIUpdateConfigStoreRequest struct {
	ctx           context.Context
	APIService    ConfigStoreAPI
	configStoreID string
	name          *string
}

// Name The name of the config store.
func (r *APIUpdateConfigStoreRequest) Name(name string) *APIUpdateConfigStoreRequest {
	r.name = &name
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateConfigStoreRequest) Execute() (*ConfigStoreResponse, *http.Response, error) {
	return r.APIService.UpdateConfigStoreExecute(r)
}

/*
UpdateConfigStore Update a config store

Update a config store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIUpdateConfigStoreRequest
*/
func (a *ConfigStoreAPIService) UpdateConfigStore(ctx context.Context, configStoreID string) APIUpdateConfigStoreRequest {
	return APIUpdateConfigStoreRequest{
		APIService:    a,
		ctx:           ctx,
		configStoreID: configStoreID,
	}
}

// UpdateConfigStoreExecute executes the request
//  @return ConfigStoreResponse
func (a *ConfigStoreAPIService) UpdateConfigStoreExecute(r APIUpdateConfigStoreRequest) (*ConfigStoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConfigStoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.UpdateConfigStore")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))

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
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
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
