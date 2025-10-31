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

// TlsPrivateKeysAPI defines an interface for interacting with the resource.
type TlsPrivateKeysAPI interface {

	/*
		CreateTlsKey Create a TLS private key

		Create a TLS private key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateTlsKeyRequest
	*/
	CreateTlsKey(ctx context.Context) APICreateTlsKeyRequest

	// CreateTlsKeyExecute executes the request
	//  @return TlsPrivateKeyResponse
	CreateTlsKeyExecute(r APICreateTlsKeyRequest) (*TlsPrivateKeyResponse, *http.Response, error)

	/*
		DeleteTlsKey Delete a TLS private key

		Destroy a TLS private key. Only private keys not already matched to any certificates can be deleted.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsPrivateKeyId Alphanumeric string identifying a private Key.
		 @return APIDeleteTlsKeyRequest
	*/
	DeleteTlsKey(ctx context.Context, tlsPrivateKeyId string) APIDeleteTlsKeyRequest

	// DeleteTlsKeyExecute executes the request
	DeleteTlsKeyExecute(r APIDeleteTlsKeyRequest) (*http.Response, error)

	/*
		GetTlsKey Get a TLS private key

		Show a TLS private key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsPrivateKeyId Alphanumeric string identifying a private Key.
		 @return APIGetTlsKeyRequest
	*/
	GetTlsKey(ctx context.Context, tlsPrivateKeyId string) APIGetTlsKeyRequest

	// GetTlsKeyExecute executes the request
	//  @return TlsPrivateKeyResponse
	GetTlsKeyExecute(r APIGetTlsKeyRequest) (*TlsPrivateKeyResponse, *http.Response, error)

	/*
		ListTlsKeys List TLS private keys

		List all TLS private keys.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTlsKeysRequest
	*/
	ListTlsKeys(ctx context.Context) APIListTlsKeysRequest

	// ListTlsKeysExecute executes the request
	//  @return TlsPrivateKeysResponse
	ListTlsKeysExecute(r APIListTlsKeysRequest) (*TlsPrivateKeysResponse, *http.Response, error)
}

// TlsPrivateKeysAPIService TlsPrivateKeysAPI service
type TlsPrivateKeysAPIService service

// APICreateTlsKeyRequest represents a request for the resource.
type APICreateTlsKeyRequest struct {
	ctx           context.Context
	APIService    TlsPrivateKeysAPI
	tlsPrivateKey *TlsPrivateKey
}

// TlsPrivateKey returns a pointer to a request.
func (r *APICreateTlsKeyRequest) TlsPrivateKey(tlsPrivateKey TlsPrivateKey) *APICreateTlsKeyRequest {
	r.tlsPrivateKey = &tlsPrivateKey
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTlsKeyRequest) Execute() (*TlsPrivateKeyResponse, *http.Response, error) {
	return r.APIService.CreateTlsKeyExecute(r)
}

/*
CreateTlsKey Create a TLS private key

Create a TLS private key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTlsKeyRequest
*/
func (a *TlsPrivateKeysAPIService) CreateTlsKey(ctx context.Context) APICreateTlsKeyRequest {
	return APICreateTlsKeyRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateTlsKeyExecute executes the request
//  @return TlsPrivateKeyResponse
func (a *TlsPrivateKeysAPIService) CreateTlsKeyExecute(r APICreateTlsKeyRequest) (*TlsPrivateKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsPrivateKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsPrivateKeysAPIService.CreateTlsKey")
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

// APIDeleteTlsKeyRequest represents a request for the resource.
type APIDeleteTlsKeyRequest struct {
	ctx             context.Context
	APIService      TlsPrivateKeysAPI
	tlsPrivateKeyId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteTlsKeyRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTlsKeyExecute(r)
}

/*
DeleteTlsKey Delete a TLS private key

Destroy a TLS private key. Only private keys not already matched to any certificates can be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsPrivateKeyId Alphanumeric string identifying a private Key.
 @return APIDeleteTlsKeyRequest
*/
func (a *TlsPrivateKeysAPIService) DeleteTlsKey(ctx context.Context, tlsPrivateKeyId string) APIDeleteTlsKeyRequest {
	return APIDeleteTlsKeyRequest{
		APIService:      a,
		ctx:             ctx,
		tlsPrivateKeyId: tlsPrivateKeyId,
	}
}

// DeleteTlsKeyExecute executes the request
func (a *TlsPrivateKeysAPIService) DeleteTlsKeyExecute(r APIDeleteTlsKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsPrivateKeysAPIService.DeleteTlsKey")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/private_keys/{tls_private_key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_private_key_id"+"}", gourl.PathEscape(parameterToString(r.tlsPrivateKeyId, "")))

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

// APIGetTlsKeyRequest represents a request for the resource.
type APIGetTlsKeyRequest struct {
	ctx             context.Context
	APIService      TlsPrivateKeysAPI
	tlsPrivateKeyId string
}

// Execute calls the API using the request data configured.
func (r APIGetTlsKeyRequest) Execute() (*TlsPrivateKeyResponse, *http.Response, error) {
	return r.APIService.GetTlsKeyExecute(r)
}

/*
GetTlsKey Get a TLS private key

Show a TLS private key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsPrivateKeyId Alphanumeric string identifying a private Key.
 @return APIGetTlsKeyRequest
*/
func (a *TlsPrivateKeysAPIService) GetTlsKey(ctx context.Context, tlsPrivateKeyId string) APIGetTlsKeyRequest {
	return APIGetTlsKeyRequest{
		APIService:      a,
		ctx:             ctx,
		tlsPrivateKeyId: tlsPrivateKeyId,
	}
}

// GetTlsKeyExecute executes the request
//  @return TlsPrivateKeyResponse
func (a *TlsPrivateKeysAPIService) GetTlsKeyExecute(r APIGetTlsKeyRequest) (*TlsPrivateKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsPrivateKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsPrivateKeysAPIService.GetTlsKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/private_keys/{tls_private_key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_private_key_id"+"}", gourl.PathEscape(parameterToString(r.tlsPrivateKeyId, "")))

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

// APIListTlsKeysRequest represents a request for the resource.
type APIListTlsKeysRequest struct {
	ctx         context.Context
	APIService  TlsPrivateKeysAPI
	filterInUse *string
	pageNumber  *int32
	pageSize    *int32
}

// FilterInUse Limit the returned keys to those without any matching TLS certificates. The only valid value is false.
func (r *APIListTlsKeysRequest) FilterInUse(filterInUse string) *APIListTlsKeysRequest {
	r.filterInUse = &filterInUse
	return r
}

// PageNumber Current page.
func (r *APIListTlsKeysRequest) PageNumber(pageNumber int32) *APIListTlsKeysRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTlsKeysRequest) PageSize(pageSize int32) *APIListTlsKeysRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTlsKeysRequest) Execute() (*TlsPrivateKeysResponse, *http.Response, error) {
	return r.APIService.ListTlsKeysExecute(r)
}

/*
ListTlsKeys List TLS private keys

List all TLS private keys.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTlsKeysRequest
*/
func (a *TlsPrivateKeysAPIService) ListTlsKeys(ctx context.Context) APIListTlsKeysRequest {
	return APIListTlsKeysRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTlsKeysExecute executes the request
//  @return TlsPrivateKeysResponse
func (a *TlsPrivateKeysAPIService) ListTlsKeysExecute(r APIListTlsKeysRequest) (*TlsPrivateKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsPrivateKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsPrivateKeysAPIService.ListTlsKeys")
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
