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
		KvStoreDeleteItem Delete an item.

		Delete an item.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @param key
		 @return APIKvStoreDeleteItemRequest
	*/
	KvStoreDeleteItem(ctx context.Context, storeID string, key string) APIKvStoreDeleteItemRequest

	// KvStoreDeleteItemExecute executes the request
	KvStoreDeleteItemExecute(r APIKvStoreDeleteItemRequest) (*http.Response, error)

	/*
		KvStoreGetItem Get an item.

		Get an item, including its value, metadata (if any), and generation marker.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @param key
		 @return APIKvStoreGetItemRequest
	*/
	KvStoreGetItem(ctx context.Context, storeID string, key string) APIKvStoreGetItemRequest

	// KvStoreGetItemExecute executes the request
	//  @return string
	KvStoreGetItemExecute(r APIKvStoreGetItemRequest) (string, *http.Response, error)

	/*
		KvStoreListItemKeys List item keys.

		Lists the matching item keys (or all item keys, if no prefix is supplied).

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @return APIKvStoreListItemKeysRequest
	*/
	KvStoreListItemKeys(ctx context.Context, storeID string) APIKvStoreListItemKeysRequest

	// KvStoreListItemKeysExecute executes the request
	//  @return InlineResponse2004
	KvStoreListItemKeysExecute(r APIKvStoreListItemKeysRequest) (*InlineResponse2004, *http.Response, error)

	/*
		KvStoreUpsertItem Insert or update an item.

		Inserts or updates an item's value and metadata.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param storeID
		 @param key
		 @return APIKvStoreUpsertItemRequest
	*/
	KvStoreUpsertItem(ctx context.Context, storeID string, key string) APIKvStoreUpsertItemRequest

	// KvStoreUpsertItemExecute executes the request
	KvStoreUpsertItemExecute(r APIKvStoreUpsertItemRequest) (*http.Response, error)
}

// KvStoreItemAPIService KvStoreItemAPI service
type KvStoreItemAPIService service

// APIKvStoreDeleteItemRequest represents a request for the resource.
type APIKvStoreDeleteItemRequest struct {
	ctx               context.Context
	APIService        KvStoreItemAPI
	storeID           string
	key               string
	ifGenerationMatch *int32
	force             *bool
}

// IfGenerationMatch returns a pointer to a request.
func (r *APIKvStoreDeleteItemRequest) IfGenerationMatch(ifGenerationMatch int32) *APIKvStoreDeleteItemRequest {
	r.ifGenerationMatch = &ifGenerationMatch
	return r
}

// Force returns a pointer to a request.
func (r *APIKvStoreDeleteItemRequest) Force(force bool) *APIKvStoreDeleteItemRequest {
	r.force = &force
	return r
}

// Execute calls the API using the request data configured.
func (r APIKvStoreDeleteItemRequest) Execute() (*http.Response, error) {
	return r.APIService.KvStoreDeleteItemExecute(r)
}

/*
KvStoreDeleteItem Delete an item.

Delete an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param key
 @return APIKvStoreDeleteItemRequest
*/
func (a *KvStoreItemAPIService) KvStoreDeleteItem(ctx context.Context, storeID string, key string) APIKvStoreDeleteItemRequest {
	return APIKvStoreDeleteItemRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
		key:        key,
	}
}

// KvStoreDeleteItemExecute executes the request
func (a *KvStoreItemAPIService) KvStoreDeleteItemExecute(r APIKvStoreDeleteItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.KvStoreDeleteItem")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys/{key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key"+"}", gourl.PathEscape(parameterToString(r.key, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
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
	if r.ifGenerationMatch != nil {
		localVarHeaderParams["if-generation-match"] = parameterToString(*r.ifGenerationMatch, "")
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

// APIKvStoreGetItemRequest represents a request for the resource.
type APIKvStoreGetItemRequest struct {
	ctx        context.Context
	APIService KvStoreItemAPI
	storeID    string
	key        string
}

// Execute calls the API using the request data configured.
func (r APIKvStoreGetItemRequest) Execute() (string, *http.Response, error) {
	return r.APIService.KvStoreGetItemExecute(r)
}

/*
KvStoreGetItem Get an item.

Get an item, including its value, metadata (if any), and generation marker.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param key
 @return APIKvStoreGetItemRequest
*/
func (a *KvStoreItemAPIService) KvStoreGetItem(ctx context.Context, storeID string, key string) APIKvStoreGetItemRequest {
	return APIKvStoreGetItemRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
		key:        key,
	}
}

// KvStoreGetItemExecute executes the request
//  @return string
func (a *KvStoreItemAPIService) KvStoreGetItemExecute(r APIKvStoreGetItemRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.KvStoreGetItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys/{key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key"+"}", gourl.PathEscape(parameterToString(r.key, "")))

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

// APIKvStoreListItemKeysRequest represents a request for the resource.
type APIKvStoreListItemKeysRequest struct {
	ctx         context.Context
	APIService  KvStoreItemAPI
	storeID     string
	cursor      *string
	limit       *int32
	prefix      *string
	consistency *string
}

// Cursor returns a pointer to a request.
func (r *APIKvStoreListItemKeysRequest) Cursor(cursor string) *APIKvStoreListItemKeysRequest {
	r.cursor = &cursor
	return r
}

// Limit returns a pointer to a request.
func (r *APIKvStoreListItemKeysRequest) Limit(limit int32) *APIKvStoreListItemKeysRequest {
	r.limit = &limit
	return r
}

// Prefix returns a pointer to a request.
func (r *APIKvStoreListItemKeysRequest) Prefix(prefix string) *APIKvStoreListItemKeysRequest {
	r.prefix = &prefix
	return r
}

// Consistency returns a pointer to a request.
func (r *APIKvStoreListItemKeysRequest) Consistency(consistency string) *APIKvStoreListItemKeysRequest {
	r.consistency = &consistency
	return r
}

// Execute calls the API using the request data configured.
func (r APIKvStoreListItemKeysRequest) Execute() (*InlineResponse2004, *http.Response, error) {
	return r.APIService.KvStoreListItemKeysExecute(r)
}

/*
KvStoreListItemKeys List item keys.

Lists the matching item keys (or all item keys, if no prefix is supplied).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @return APIKvStoreListItemKeysRequest
*/
func (a *KvStoreItemAPIService) KvStoreListItemKeys(ctx context.Context, storeID string) APIKvStoreListItemKeysRequest {
	return APIKvStoreListItemKeysRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
	}
}

// KvStoreListItemKeysExecute executes the request
//  @return InlineResponse2004
func (a *KvStoreItemAPIService) KvStoreListItemKeysExecute(r APIKvStoreListItemKeysRequest) (*InlineResponse2004, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse2004
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.KvStoreListItemKeys")
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
	if r.consistency != nil {
		localVarQueryParams.Add("consistency", parameterToString(*r.consistency, ""))
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

// APIKvStoreUpsertItemRequest represents a request for the resource.
type APIKvStoreUpsertItemRequest struct {
	ctx               context.Context
	APIService        KvStoreItemAPI
	storeID           string
	key               string
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
func (r *APIKvStoreUpsertItemRequest) IfGenerationMatch(ifGenerationMatch int32) *APIKvStoreUpsertItemRequest {
	r.ifGenerationMatch = &ifGenerationMatch
	return r
}

// TimeToLiveSec returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) TimeToLiveSec(timeToLiveSec int32) *APIKvStoreUpsertItemRequest {
	r.timeToLiveSec = &timeToLiveSec
	return r
}

// Metadata returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) Metadata(metadata string) *APIKvStoreUpsertItemRequest {
	r.metadata = &metadata
	return r
}

// Add returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) Add(add bool) *APIKvStoreUpsertItemRequest {
	r.add = &add
	return r
}

// Append returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) Append(append bool) *APIKvStoreUpsertItemRequest {
	r.append = &append
	return r
}

// Prepend returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) Prepend(prepend bool) *APIKvStoreUpsertItemRequest {
	r.prepend = &prepend
	return r
}

// BackgroundFetch returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) BackgroundFetch(backgroundFetch bool) *APIKvStoreUpsertItemRequest {
	r.backgroundFetch = &backgroundFetch
	return r
}

// Body returns a pointer to a request.
func (r *APIKvStoreUpsertItemRequest) Body(body string) *APIKvStoreUpsertItemRequest {
	r.body = &body
	return r
}

// Execute calls the API using the request data configured.
func (r APIKvStoreUpsertItemRequest) Execute() (*http.Response, error) {
	return r.APIService.KvStoreUpsertItemExecute(r)
}

/*
KvStoreUpsertItem Insert or update an item.

Inserts or updates an item's value and metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param key
 @return APIKvStoreUpsertItemRequest
*/
func (a *KvStoreItemAPIService) KvStoreUpsertItem(ctx context.Context, storeID string, key string) APIKvStoreUpsertItemRequest {
	return APIKvStoreUpsertItemRequest{
		APIService: a,
		ctx:        ctx,
		storeID:    storeID,
		key:        key,
	}
}

// KvStoreUpsertItemExecute executes the request
func (a *KvStoreItemAPIService) KvStoreUpsertItemExecute(r APIKvStoreUpsertItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvStoreItemAPIService.KvStoreUpsertItem")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/kv/{store_id}/keys/{key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key"+"}", gourl.PathEscape(parameterToString(r.key, "")))

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
	localVarHTTPHeaderAccepts := []string{}

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
