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

// TLSPrivateKeysAPI defines an interface for interacting with the resource.
type TLSPrivateKeysAPI interface {

	/*
	CreateTLSKey Create a TLS private key

	Create a TLS private key.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateTLSKeyRequest
	*/
	CreateTLSKey(ctx context.Context) APICreateTLSKeyRequest

	// CreateTLSKeyExecute executes the request
	//  @return TLSPrivateKeyResponse
	CreateTLSKeyExecute(r APICreateTLSKeyRequest) (*TLSPrivateKeyResponse, *http.Response, error)

	/*
	DeleteTLSKey Delete a TLS private key

	Destroy a TLS private key. Only private keys not already matched to any certificates can be deleted.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsPrivateKeyID Alphanumeric string identifying a private Key.
	 @return APIDeleteTLSKeyRequest
	*/
	DeleteTLSKey(ctx context.Context, tlsPrivateKeyID string) APIDeleteTLSKeyRequest

	// DeleteTLSKeyExecute executes the request
	DeleteTLSKeyExecute(r APIDeleteTLSKeyRequest) (*http.Response, error)

	/*
	GetTLSKey Get a TLS private key

	Show a TLS private key.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsPrivateKeyID Alphanumeric string identifying a private Key.
	 @return APIGetTLSKeyRequest
	*/
	GetTLSKey(ctx context.Context, tlsPrivateKeyID string) APIGetTLSKeyRequest

	// GetTLSKeyExecute executes the request
	//  @return TLSPrivateKeyResponse
	GetTLSKeyExecute(r APIGetTLSKeyRequest) (*TLSPrivateKeyResponse, *http.Response, error)

	/*
	ListTLSKeys List TLS private keys

	List all TLS private keys.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListTLSKeysRequest
	*/
	ListTLSKeys(ctx context.Context) APIListTLSKeysRequest

	// ListTLSKeysExecute executes the request
	//  @return TLSPrivateKeysResponse
	ListTLSKeysExecute(r APIListTLSKeysRequest) (*TLSPrivateKeysResponse, *http.Response, error)
}

// TLSPrivateKeysAPIService TLSPrivateKeysAPI service
type TLSPrivateKeysAPIService service

// APICreateTLSKeyRequest represents a request for the resource.
type APICreateTLSKeyRequest struct {
	ctx context.Context
	APIService TLSPrivateKeysAPI
	tlsPrivateKey *TLSPrivateKey
}

// TLSPrivateKey returns a pointer to a request.
func (r *APICreateTLSKeyRequest) TLSPrivateKey(tlsPrivateKey TLSPrivateKey) *APICreateTLSKeyRequest {
	r.tlsPrivateKey = &tlsPrivateKey
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTLSKeyRequest) Execute() (*TLSPrivateKeyResponse, *http.Response, error) {
	return r.APIService.CreateTLSKeyExecute(r)
}

/*
CreateTLSKey Create a TLS private key

Create a TLS private key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTLSKeyRequest
*/
func (a *TLSPrivateKeysAPIService) CreateTLSKey(ctx context.Context) APICreateTLSKeyRequest {
	return APICreateTLSKeyRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateTLSKeyExecute executes the request
//  @return TLSPrivateKeyResponse
func (a *TLSPrivateKeysAPIService) CreateTLSKeyExecute(r APICreateTLSKeyRequest) (*TLSPrivateKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSPrivateKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSPrivateKeysAPIService.CreateTLSKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/private_keys"

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
	localVarPostBody = r.tlsPrivateKey
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

// APIDeleteTLSKeyRequest represents a request for the resource.
type APIDeleteTLSKeyRequest struct {
	ctx context.Context
	APIService TLSPrivateKeysAPI
	tlsPrivateKeyID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteTLSKeyRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTLSKeyExecute(r)
}

/*
DeleteTLSKey Delete a TLS private key

Destroy a TLS private key. Only private keys not already matched to any certificates can be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsPrivateKeyID Alphanumeric string identifying a private Key.
 @return APIDeleteTLSKeyRequest
*/
func (a *TLSPrivateKeysAPIService) DeleteTLSKey(ctx context.Context, tlsPrivateKeyID string) APIDeleteTLSKeyRequest {
	return APIDeleteTLSKeyRequest{
		APIService: a,
		ctx: ctx,
		tlsPrivateKeyID: tlsPrivateKeyID,
	}
}

// DeleteTLSKeyExecute executes the request
func (a *TLSPrivateKeysAPIService) DeleteTLSKeyExecute(r APIDeleteTLSKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSPrivateKeysAPIService.DeleteTLSKey")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/private_keys/{tls_private_key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_private_key_id"+"}", gourl.PathEscape(parameterToString(r.tlsPrivateKeyID, "")))

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

// APIGetTLSKeyRequest represents a request for the resource.
type APIGetTLSKeyRequest struct {
	ctx context.Context
	APIService TLSPrivateKeysAPI
	tlsPrivateKeyID string
}


// Execute calls the API using the request data configured.
func (r APIGetTLSKeyRequest) Execute() (*TLSPrivateKeyResponse, *http.Response, error) {
	return r.APIService.GetTLSKeyExecute(r)
}

/*
GetTLSKey Get a TLS private key

Show a TLS private key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsPrivateKeyID Alphanumeric string identifying a private Key.
 @return APIGetTLSKeyRequest
*/
func (a *TLSPrivateKeysAPIService) GetTLSKey(ctx context.Context, tlsPrivateKeyID string) APIGetTLSKeyRequest {
	return APIGetTLSKeyRequest{
		APIService: a,
		ctx: ctx,
		tlsPrivateKeyID: tlsPrivateKeyID,
	}
}

// GetTLSKeyExecute executes the request
//  @return TLSPrivateKeyResponse
func (a *TLSPrivateKeysAPIService) GetTLSKeyExecute(r APIGetTLSKeyRequest) (*TLSPrivateKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSPrivateKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSPrivateKeysAPIService.GetTLSKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/private_keys/{tls_private_key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_private_key_id"+"}", gourl.PathEscape(parameterToString(r.tlsPrivateKeyID, "")))

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

// APIListTLSKeysRequest represents a request for the resource.
type APIListTLSKeysRequest struct {
	ctx context.Context
	APIService TLSPrivateKeysAPI
	filterInUse *string
	pageNumber *int32
	pageSize *int32
}

// FilterInUse Limit the returned keys to those without any matching TLS certificates. The only valid value is false.
func (r *APIListTLSKeysRequest) FilterInUse(filterInUse string) *APIListTLSKeysRequest {
	r.filterInUse = &filterInUse
	return r
}
// PageNumber Current page.
func (r *APIListTLSKeysRequest) PageNumber(pageNumber int32) *APIListTLSKeysRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListTLSKeysRequest) PageSize(pageSize int32) *APIListTLSKeysRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSKeysRequest) Execute() (*TLSPrivateKeysResponse, *http.Response, error) {
	return r.APIService.ListTLSKeysExecute(r)
}

/*
ListTLSKeys List TLS private keys

List all TLS private keys.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSKeysRequest
*/
func (a *TLSPrivateKeysAPIService) ListTLSKeys(ctx context.Context) APIListTLSKeysRequest {
	return APIListTLSKeysRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListTLSKeysExecute executes the request
//  @return TLSPrivateKeysResponse
func (a *TLSPrivateKeysAPIService) ListTLSKeysExecute(r APIListTLSKeysRequest) (*TLSPrivateKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSPrivateKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSPrivateKeysAPIService.ListTLSKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/private_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterInUse != nil {
		localVarQueryParams.Add("filter[in_use]", parameterToString(*r.filterInUse, ""))
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
