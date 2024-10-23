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

// KvStoreItemAPI defines an interface for interacting with the resource.
type KvStoreItemAPI interface {

	/*
		DeleteKeyFromStore Delete kv store item.

		Delete an item from an kv store

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @param keyName
		 @return APIDeleteKeyFromStoreRequest
	*/
	DeleteKeyFromStore(ctx context.Context, storeID string, keyName string) APIDeleteKeyFromStoreRequest

	// DeleteKeyFromStoreExecute executes the request
	DeleteKeyFromStoreExecute(r APIDeleteKeyFromStoreRequest) (*http.Response, error)

	/*
		GetKeys List kv store keys.

		List the keys of all items within an kv store.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @return APIGetKeysRequest
	*/
	GetKeys(ctx context.Context, storeID string) APIGetKeysRequest

	// GetKeysExecute executes the request
	//  @return InlineResponse2004
	GetKeysExecute(r APIGetKeysRequest) (*InlineResponse2004, *http.Response, error)

	/*
		GetValueForKey Get the value of an kv store item

		Get the value associated with a key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @param keyName
		 @return APIGetValueForKeyRequest
	*/
	GetValueForKey(ctx context.Context, storeID string, keyName string) APIGetValueForKeyRequest

	// GetValueForKeyExecute executes the request
	//  @return string
	GetValueForKeyExecute(r APIGetValueForKeyRequest) (string, *http.Response, error)

	/*
		SetValueForKey Insert an item into an kv store

		Set a new value for a new or existing key in an kv store.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @param keyName
		 @return APISetValueForKeyRequest
	*/
	SetValueForKey(ctx context.Context, storeID string, keyName string) APISetValueForKeyRequest

	// SetValueForKeyExecute executes the request
	//  @return string
	SetValueForKeyExecute(r APISetValueForKeyRequest) (string, *http.Response, error)
}

// KvStoreItemAPIService KvStoreItemAPI service
type KvStoreItemAPIService service

// APIDeleteKeyFromStoreRequest represents a request for the resource.
type APIDeleteKeyFromStoreRequest struct {
	ctx        context.Context
	APIService KvStoreItemAPI
	storeID    string
	keyName    string
}

// Execute calls the API using the request data configured.
func (r APIDeleteKeyFromStoreRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteKeyFromStoreExecute(r)
}

/*
DeleteKeyFromStore Delete kv store item.

Delete an item from an kv store

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param keyName
 @return APIDeleteKeyFromStoreRequest
*/
func (a *KvStoreItemAPIService) DeleteKeyFromStore(ctx context.Context, storeID string, keyName string) APIDeleteKeyFromStoreRequest {
	return APIDeleteKeyFromStoreRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
		keyName:    keyName,
	}
}

// DeleteKeyFromStoreExecute executes the request
func (a *KvStoreItemAPIService) DeleteKeyFromStoreExecute(r APIDeleteKeyFromStoreRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.DeleteKeyFromStore")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys/{key_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_name"+"}", gourl.PathEscape(parameterToString(r.keyName, "")))

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

// APIGetKeysRequest represents a request for the resource.
type APIGetKeysRequest struct {
	ctx        context.Context
	APIService KvStoreItemAPI
	storeID    string
	cursor     *string
	limit      *int32
	prefix     *string
}

// Cursor returns a pointer to a request.
func (r *APIGetKeysRequest) Cursor(cursor string) *APIGetKeysRequest {
	r.cursor = &cursor
	return r
}

// Limit returns a pointer to a request.
func (r *APIGetKeysRequest) Limit(limit int32) *APIGetKeysRequest {
	r.limit = &limit
	return r
}

// Prefix returns a pointer to a request.
func (r *APIGetKeysRequest) Prefix(prefix string) *APIGetKeysRequest {
	r.prefix = &prefix
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetKeysRequest) Execute() (*InlineResponse2004, *http.Response, error) {
	return r.APIService.GetKeysExecute(r)
}

/*
GetKeys List kv store keys.

List the keys of all items within an kv store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @return APIGetKeysRequest
*/
func (a *KvStoreItemAPIService) GetKeys(ctx context.Context, storeID string) APIGetKeysRequest {
	return APIGetKeysRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
	}
}

// GetKeysExecute executes the request
//  @return InlineResponse2004
func (a *KvStoreItemAPIService) GetKeysExecute(r APIGetKeysRequest) (*InlineResponse2004, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse2004
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.GetKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
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

// APIGetValueForKeyRequest represents a request for the resource.
type APIGetValueForKeyRequest struct {
	ctx        context.Context
	APIService KvStoreItemAPI
	storeID    string
	keyName    string
}

// Execute calls the API using the request data configured.
func (r APIGetValueForKeyRequest) Execute() (string, *http.Response, error) {
	return r.APIService.GetValueForKeyExecute(r)
}

/*
GetValueForKey Get the value of an kv store item

Get the value associated with a key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param keyName
 @return APIGetValueForKeyRequest
*/
func (a *KvStoreItemAPIService) GetValueForKey(ctx context.Context, storeID string, keyName string) APIGetValueForKeyRequest {
	return APIGetValueForKeyRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
		keyName:    keyName,
	}
}

// GetValueForKeyExecute executes the request
//  @return string
func (a *KvStoreItemAPIService) GetValueForKeyExecute(r APIGetValueForKeyRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.GetValueForKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys/{key_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_name"+"}", gourl.PathEscape(parameterToString(r.keyName, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

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

// APISetValueForKeyRequest represents a request for the resource.
type APISetValueForKeyRequest struct {
	ctx               context.Context
	APIService        KvStoreItemAPI
	storeID           string
	keyName           string
	ifGenerationMatch *int32
	timeToLiveSec     *int32
	metadata          *string
	add               *bool
	append            *bool
	prepend           *bool
	backgroundFetch   *bool
	body              *string
}

// IfGenerationMatch returns a pointer to a request.
func (r *APISetValueForKeyRequest) IfGenerationMatch(ifGenerationMatch int32) *APISetValueForKeyRequest {
	r.ifGenerationMatch = &ifGenerationMatch
	return r
}

// TimeToLiveSec returns a pointer to a request.
func (r *APISetValueForKeyRequest) TimeToLiveSec(timeToLiveSec int32) *APISetValueForKeyRequest {
	r.timeToLiveSec = &timeToLiveSec
	return r
}

// Metadata returns a pointer to a request.
func (r *APISetValueForKeyRequest) Metadata(metadata string) *APISetValueForKeyRequest {
	r.metadata = &metadata
	return r
}

// Add returns a pointer to a request.
func (r *APISetValueForKeyRequest) Add(add bool) *APISetValueForKeyRequest {
	r.add = &add
	return r
}

// Append returns a pointer to a request.
func (r *APISetValueForKeyRequest) Append(append bool) *APISetValueForKeyRequest {
	r.append = &append
	return r
}

// Prepend returns a pointer to a request.
func (r *APISetValueForKeyRequest) Prepend(prepend bool) *APISetValueForKeyRequest {
	r.prepend = &prepend
	return r
}

// BackgroundFetch returns a pointer to a request.
func (r *APISetValueForKeyRequest) BackgroundFetch(backgroundFetch bool) *APISetValueForKeyRequest {
	r.backgroundFetch = &backgroundFetch
	return r
}

// Body returns a pointer to a request.
func (r *APISetValueForKeyRequest) Body(body string) *APISetValueForKeyRequest {
	r.body = &body
	return r
}

// Execute calls the API using the request data configured.
func (r APISetValueForKeyRequest) Execute() (string, *http.Response, error) {
	return r.APIService.SetValueForKeyExecute(r)
}

/*
SetValueForKey Insert an item into an kv store

Set a new value for a new or existing key in an kv store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param keyName
 @return APISetValueForKeyRequest
*/
func (a *KvStoreItemAPIService) SetValueForKey(ctx context.Context, storeID string, keyName string) APISetValueForKeyRequest {
	return APISetValueForKeyRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
		keyName:    keyName,
	}
}

// SetValueForKeyExecute executes the request
//  @return string
func (a *KvStoreItemAPIService) SetValueForKeyExecute(r APISetValueForKeyRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.SetValueForKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys/{key_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_name"+"}", gourl.PathEscape(parameterToString(r.keyName, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.add != nil {
		localVarQueryParams.Add("add", parameterToString(*r.add, ""))
	}
	if r.append != nil {
		localVarQueryParams.Add("append", parameterToString(*r.append, ""))
	}
	if r.prepend != nil {
		localVarQueryParams.Add("prepend", parameterToString(*r.prepend, ""))
	}
	if r.backgroundFetch != nil {
		localVarQueryParams.Add("background_fetch", parameterToString(*r.backgroundFetch, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifGenerationMatch != nil {
		localVarHeaderParams["if-generation-match"] = parameterToString(*r.ifGenerationMatch, "")
	}
	if r.timeToLiveSec != nil {
		localVarHeaderParams["time_to_live_sec"] = parameterToString(*r.timeToLiveSec, "")
	}
	if r.metadata != nil {
		localVarHeaderParams["metadata"] = parameterToString(*r.metadata, "")
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
