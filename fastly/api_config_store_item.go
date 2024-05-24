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

// ConfigStoreItemAPI defines an interface for interacting with the resource.
type ConfigStoreItemAPI interface {

	/*
	BulkUpdateConfigStoreItem Update multiple entries in a config store

	Add multiple key-value pairs to an individual config store, specified by ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @return APIBulkUpdateConfigStoreItemRequest
	*/
	BulkUpdateConfigStoreItem(ctx context.Context, configStoreID string) APIBulkUpdateConfigStoreItemRequest

	// BulkUpdateConfigStoreItemExecute executes the request
	//  @return InlineResponse200
	BulkUpdateConfigStoreItemExecute(r APIBulkUpdateConfigStoreItemRequest) (*InlineResponse200, *http.Response, error)

	/*
	CreateConfigStoreItem Create an entry in a config store

	Add a single key-value pair to an individual config store, specified by ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @return APICreateConfigStoreItemRequest
	*/
	CreateConfigStoreItem(ctx context.Context, configStoreID string) APICreateConfigStoreItemRequest

	// CreateConfigStoreItemExecute executes the request
	//  @return ConfigStoreItemResponse
	CreateConfigStoreItemExecute(r APICreateConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error)

	/*
	DeleteConfigStoreItem Delete an item from a config store

	Delete an entry in a config store given a config store ID, and item key.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @param configStoreItemKey Item key, maximum 256 characters.
	 @return APIDeleteConfigStoreItemRequest
	*/
	DeleteConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIDeleteConfigStoreItemRequest

	// DeleteConfigStoreItemExecute executes the request
	//  @return InlineResponse200
	DeleteConfigStoreItemExecute(r APIDeleteConfigStoreItemRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetConfigStoreItem Get an item from a config store

	Retrieve a config store entry given a config store ID and item key.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @param configStoreItemKey Item key, maximum 256 characters.
	 @return APIGetConfigStoreItemRequest
	*/
	GetConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIGetConfigStoreItemRequest

	// GetConfigStoreItemExecute executes the request
	//  @return ConfigStoreItemResponse
	GetConfigStoreItemExecute(r APIGetConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error)

	/*
	ListConfigStoreItems List items in a config store

	List the key-value pairs associated with a given config store ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @return APIListConfigStoreItemsRequest
	*/
	ListConfigStoreItems(ctx context.Context, configStoreID string) APIListConfigStoreItemsRequest

	// ListConfigStoreItemsExecute executes the request
	//  @return []ConfigStoreItemResponse
	ListConfigStoreItemsExecute(r APIListConfigStoreItemsRequest) ([]ConfigStoreItemResponse, *http.Response, error)

	/*
	UpdateConfigStoreItem Update an entry in a config store

	Update an entry in a config store given a config store ID, item key, and item value.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @param configStoreItemKey Item key, maximum 256 characters.
	 @return APIUpdateConfigStoreItemRequest
	*/
	UpdateConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIUpdateConfigStoreItemRequest

	// UpdateConfigStoreItemExecute executes the request
	//  @return ConfigStoreItemResponse
	UpdateConfigStoreItemExecute(r APIUpdateConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error)

	/*
	UpsertConfigStoreItem Insert or update an entry in a config store

	Insert or update an entry in a config store given a config store ID, item key, and item value.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param configStoreID An alphanumeric string identifying the config store.
	 @param configStoreItemKey Item key, maximum 256 characters.
	 @return APIUpsertConfigStoreItemRequest
	*/
	UpsertConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIUpsertConfigStoreItemRequest

	// UpsertConfigStoreItemExecute executes the request
	//  @return ConfigStoreItemResponse
	UpsertConfigStoreItemExecute(r APIUpsertConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error)
}

// ConfigStoreItemAPIService ConfigStoreItemAPI service
type ConfigStoreItemAPIService service

// APIBulkUpdateConfigStoreItemRequest represents a request for the resource.
type APIBulkUpdateConfigStoreItemRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
	bulkUpdateConfigStoreListRequest *BulkUpdateConfigStoreListRequest
}

// BulkUpdateConfigStoreListRequest returns a pointer to a request.
func (r *APIBulkUpdateConfigStoreItemRequest) BulkUpdateConfigStoreListRequest(bulkUpdateConfigStoreListRequest BulkUpdateConfigStoreListRequest) *APIBulkUpdateConfigStoreItemRequest {
	r.bulkUpdateConfigStoreListRequest = &bulkUpdateConfigStoreListRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkUpdateConfigStoreItemRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.BulkUpdateConfigStoreItemExecute(r)
}

/*
BulkUpdateConfigStoreItem Update multiple entries in a config store

Add multiple key-value pairs to an individual config store, specified by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIBulkUpdateConfigStoreItemRequest
*/
func (a *ConfigStoreItemAPIService) BulkUpdateConfigStoreItem(ctx context.Context, configStoreID string) APIBulkUpdateConfigStoreItemRequest {
	return APIBulkUpdateConfigStoreItemRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
	}
}

// BulkUpdateConfigStoreItemExecute executes the request
//  @return InlineResponse200
func (a *ConfigStoreItemAPIService) BulkUpdateConfigStoreItemExecute(r APIBulkUpdateConfigStoreItemRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.BulkUpdateConfigStoreItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/items"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.bulkUpdateConfigStoreListRequest
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

// APICreateConfigStoreItemRequest represents a request for the resource.
type APICreateConfigStoreItemRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
	itemKey *string
	itemValue *string
}

// ItemKey Item key, maximum 256 characters.
func (r *APICreateConfigStoreItemRequest) ItemKey(itemKey string) *APICreateConfigStoreItemRequest {
	r.itemKey = &itemKey
	return r
}
// ItemValue Item value, maximum 8000 characters.
func (r *APICreateConfigStoreItemRequest) ItemValue(itemValue string) *APICreateConfigStoreItemRequest {
	r.itemValue = &itemValue
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateConfigStoreItemRequest) Execute() (*ConfigStoreItemResponse, *http.Response, error) {
	return r.APIService.CreateConfigStoreItemExecute(r)
}

/*
CreateConfigStoreItem Create an entry in a config store

Add a single key-value pair to an individual config store, specified by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APICreateConfigStoreItemRequest
*/
func (a *ConfigStoreItemAPIService) CreateConfigStoreItem(ctx context.Context, configStoreID string) APICreateConfigStoreItemRequest {
	return APICreateConfigStoreItemRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
	}
}

// CreateConfigStoreItemExecute executes the request
//  @return ConfigStoreItemResponse
func (a *ConfigStoreItemAPIService) CreateConfigStoreItemExecute(r APICreateConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ConfigStoreItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.CreateConfigStoreItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/item"
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
	if r.itemKey != nil {
		localVarFormParams.Add("item_key", parameterToString(*r.itemKey, ""))
	}
	if r.itemValue != nil {
		localVarFormParams.Add("item_value", parameterToString(*r.itemValue, ""))
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

// APIDeleteConfigStoreItemRequest represents a request for the resource.
type APIDeleteConfigStoreItemRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
	configStoreItemKey string
}


// Execute calls the API using the request data configured.
func (r APIDeleteConfigStoreItemRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteConfigStoreItemExecute(r)
}

/*
DeleteConfigStoreItem Delete an item from a config store

Delete an entry in a config store given a config store ID, and item key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @param configStoreItemKey Item key, maximum 256 characters.
 @return APIDeleteConfigStoreItemRequest
*/
func (a *ConfigStoreItemAPIService) DeleteConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIDeleteConfigStoreItemRequest {
	return APIDeleteConfigStoreItemRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
		configStoreItemKey: configStoreItemKey,
	}
}

// DeleteConfigStoreItemExecute executes the request
//  @return InlineResponse200
func (a *ConfigStoreItemAPIService) DeleteConfigStoreItemExecute(r APIDeleteConfigStoreItemRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.DeleteConfigStoreItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/item/{config_store_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_item_key"+"}", gourl.PathEscape(parameterToString(r.configStoreItemKey, "")))

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

// APIGetConfigStoreItemRequest represents a request for the resource.
type APIGetConfigStoreItemRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
	configStoreItemKey string
}


// Execute calls the API using the request data configured.
func (r APIGetConfigStoreItemRequest) Execute() (*ConfigStoreItemResponse, *http.Response, error) {
	return r.APIService.GetConfigStoreItemExecute(r)
}

/*
GetConfigStoreItem Get an item from a config store

Retrieve a config store entry given a config store ID and item key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @param configStoreItemKey Item key, maximum 256 characters.
 @return APIGetConfigStoreItemRequest
*/
func (a *ConfigStoreItemAPIService) GetConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIGetConfigStoreItemRequest {
	return APIGetConfigStoreItemRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
		configStoreItemKey: configStoreItemKey,
	}
}

// GetConfigStoreItemExecute executes the request
//  @return ConfigStoreItemResponse
func (a *ConfigStoreItemAPIService) GetConfigStoreItemExecute(r APIGetConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ConfigStoreItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.GetConfigStoreItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/item/{config_store_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_item_key"+"}", gourl.PathEscape(parameterToString(r.configStoreItemKey, "")))

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

// APIListConfigStoreItemsRequest represents a request for the resource.
type APIListConfigStoreItemsRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
}


// Execute calls the API using the request data configured.
func (r APIListConfigStoreItemsRequest) Execute() ([]ConfigStoreItemResponse, *http.Response, error) {
	return r.APIService.ListConfigStoreItemsExecute(r)
}

/*
ListConfigStoreItems List items in a config store

List the key-value pairs associated with a given config store ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @return APIListConfigStoreItemsRequest
*/
func (a *ConfigStoreItemAPIService) ListConfigStoreItems(ctx context.Context, configStoreID string) APIListConfigStoreItemsRequest {
	return APIListConfigStoreItemsRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
	}
}

// ListConfigStoreItemsExecute executes the request
//  @return []ConfigStoreItemResponse
func (a *ConfigStoreItemAPIService) ListConfigStoreItemsExecute(r APIListConfigStoreItemsRequest) ([]ConfigStoreItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []ConfigStoreItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.ListConfigStoreItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/items"
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

// APIUpdateConfigStoreItemRequest represents a request for the resource.
type APIUpdateConfigStoreItemRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
	configStoreItemKey string
	itemKey *string
	itemValue *string
}

// ItemKey Item key, maximum 256 characters.
func (r *APIUpdateConfigStoreItemRequest) ItemKey(itemKey string) *APIUpdateConfigStoreItemRequest {
	r.itemKey = &itemKey
	return r
}
// ItemValue Item value, maximum 8000 characters.
func (r *APIUpdateConfigStoreItemRequest) ItemValue(itemValue string) *APIUpdateConfigStoreItemRequest {
	r.itemValue = &itemValue
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateConfigStoreItemRequest) Execute() (*ConfigStoreItemResponse, *http.Response, error) {
	return r.APIService.UpdateConfigStoreItemExecute(r)
}

/*
UpdateConfigStoreItem Update an entry in a config store

Update an entry in a config store given a config store ID, item key, and item value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @param configStoreItemKey Item key, maximum 256 characters.
 @return APIUpdateConfigStoreItemRequest
*/
func (a *ConfigStoreItemAPIService) UpdateConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIUpdateConfigStoreItemRequest {
	return APIUpdateConfigStoreItemRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
		configStoreItemKey: configStoreItemKey,
	}
}

// UpdateConfigStoreItemExecute executes the request
//  @return ConfigStoreItemResponse
func (a *ConfigStoreItemAPIService) UpdateConfigStoreItemExecute(r APIUpdateConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ConfigStoreItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.UpdateConfigStoreItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/item/{config_store_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_item_key"+"}", gourl.PathEscape(parameterToString(r.configStoreItemKey, "")))

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
	if r.itemKey != nil {
		localVarFormParams.Add("item_key", parameterToString(*r.itemKey, ""))
	}
	if r.itemValue != nil {
		localVarFormParams.Add("item_value", parameterToString(*r.itemValue, ""))
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

// APIUpsertConfigStoreItemRequest represents a request for the resource.
type APIUpsertConfigStoreItemRequest struct {
	ctx context.Context
	APIService ConfigStoreItemAPI
	configStoreID string
	configStoreItemKey string
	itemKey *string
	itemValue *string
}

// ItemKey Item key, maximum 256 characters.
func (r *APIUpsertConfigStoreItemRequest) ItemKey(itemKey string) *APIUpsertConfigStoreItemRequest {
	r.itemKey = &itemKey
	return r
}
// ItemValue Item value, maximum 8000 characters.
func (r *APIUpsertConfigStoreItemRequest) ItemValue(itemValue string) *APIUpsertConfigStoreItemRequest {
	r.itemValue = &itemValue
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpsertConfigStoreItemRequest) Execute() (*ConfigStoreItemResponse, *http.Response, error) {
	return r.APIService.UpsertConfigStoreItemExecute(r)
}

/*
UpsertConfigStoreItem Insert or update an entry in a config store

Insert or update an entry in a config store given a config store ID, item key, and item value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configStoreID An alphanumeric string identifying the config store.
 @param configStoreItemKey Item key, maximum 256 characters.
 @return APIUpsertConfigStoreItemRequest
*/
func (a *ConfigStoreItemAPIService) UpsertConfigStoreItem(ctx context.Context, configStoreID string, configStoreItemKey string) APIUpsertConfigStoreItemRequest {
	return APIUpsertConfigStoreItemRequest{
		APIService: a,
		ctx: ctx,
		configStoreID: configStoreID,
		configStoreItemKey: configStoreItemKey,
	}
}

// UpsertConfigStoreItemExecute executes the request
//  @return ConfigStoreItemResponse
func (a *ConfigStoreItemAPIService) UpsertConfigStoreItemExecute(r APIUpsertConfigStoreItemRequest) (*ConfigStoreItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ConfigStoreItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreItemAPIService.UpsertConfigStoreItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/config/{config_store_id}/item/{config_store_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_id"+"}", gourl.PathEscape(parameterToString(r.configStoreID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_store_item_key"+"}", gourl.PathEscape(parameterToString(r.configStoreItemKey, "")))

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
	if r.itemKey != nil {
		localVarFormParams.Add("item_key", parameterToString(*r.itemKey, ""))
	}
	if r.itemValue != nil {
		localVarFormParams.Add("item_value", parameterToString(*r.itemValue, ""))
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
