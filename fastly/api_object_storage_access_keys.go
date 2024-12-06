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

// ObjectStorageAccessKeysAPI defines an interface for interacting with the resource.
type ObjectStorageAccessKeysAPI interface {

	/*
		CreateAccessKey Create an access key

		Create an access key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateAccessKeyRequest
	*/
	CreateAccessKey(ctx context.Context) APICreateAccessKeyRequest

	// CreateAccessKeyExecute executes the request
	//  @return AccessKeyResponse
	CreateAccessKeyExecute(r APICreateAccessKeyRequest) (*AccessKeyResponse, *http.Response, error)

	/*
		DeleteAccessKey Delete an access key

		Delete an access key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param accessKey
		 @return APIDeleteAccessKeyRequest
	*/
	DeleteAccessKey(ctx context.Context, accessKey string) APIDeleteAccessKeyRequest

	// DeleteAccessKeyExecute executes the request
	DeleteAccessKeyExecute(r APIDeleteAccessKeyRequest) (*http.Response, error)

	/*
		GetAccessKey Get an access key

		Get an access key by its identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param accessKey
		 @return APIGetAccessKeyRequest
	*/
	GetAccessKey(ctx context.Context, accessKey string) APIGetAccessKeyRequest

	// GetAccessKeyExecute executes the request
	//  @return AccessKey
	GetAccessKeyExecute(r APIGetAccessKeyRequest) (*AccessKey, *http.Response, error)

	/*
		ListAccessKeys List access keys

		List access keys.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListAccessKeysRequest
	*/
	ListAccessKeys(ctx context.Context) APIListAccessKeysRequest

	// ListAccessKeysExecute executes the request
	//  @return AccessKeyResponse
	ListAccessKeysExecute(r APIListAccessKeysRequest) (*AccessKeyResponse, *http.Response, error)
}

// ObjectStorageAccessKeysAPIService ObjectStorageAccessKeysAPI service
type ObjectStorageAccessKeysAPIService service

// APICreateAccessKeyRequest represents a request for the resource.
type APICreateAccessKeyRequest struct {
	ctx        context.Context
	APIService ObjectStorageAccessKeysAPI
	accessKey  *AccessKey
}

// AccessKey returns a pointer to a request.
func (r *APICreateAccessKeyRequest) AccessKey(accessKey AccessKey) *APICreateAccessKeyRequest {
	r.accessKey = &accessKey
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateAccessKeyRequest) Execute() (*AccessKeyResponse, *http.Response, error) {
	return r.APIService.CreateAccessKeyExecute(r)
}

/*
CreateAccessKey Create an access key

Create an access key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateAccessKeyRequest
*/
func (a *ObjectStorageAccessKeysAPIService) CreateAccessKey(ctx context.Context) APICreateAccessKeyRequest {
	return APICreateAccessKeyRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateAccessKeyExecute executes the request
//  @return AccessKeyResponse
func (a *ObjectStorageAccessKeysAPIService) CreateAccessKeyExecute(r APICreateAccessKeyRequest) (*AccessKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AccessKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageAccessKeysAPIService.CreateAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/object-storage/access-keys"

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
	localVarPostBody = r.accessKey
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

// APIDeleteAccessKeyRequest represents a request for the resource.
type APIDeleteAccessKeyRequest struct {
	ctx        context.Context
	APIService ObjectStorageAccessKeysAPI
	accessKey  string
}

// Execute calls the API using the request data configured.
func (r APIDeleteAccessKeyRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteAccessKeyExecute(r)
}

/*
DeleteAccessKey Delete an access key

Delete an access key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessKey
 @return APIDeleteAccessKeyRequest
*/
func (a *ObjectStorageAccessKeysAPIService) DeleteAccessKey(ctx context.Context, accessKey string) APIDeleteAccessKeyRequest {
	return APIDeleteAccessKeyRequest{
		APIService: a,
		ctx:        ctx,
		accessKey:  accessKey,
	}
}

// DeleteAccessKeyExecute executes the request
func (a *ObjectStorageAccessKeysAPIService) DeleteAccessKeyExecute(r APIDeleteAccessKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageAccessKeysAPIService.DeleteAccessKey")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/object-storage/access-keys/{access_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"access_key"+"}", gourl.PathEscape(parameterToString(r.accessKey, "")))

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

// APIGetAccessKeyRequest represents a request for the resource.
type APIGetAccessKeyRequest struct {
	ctx        context.Context
	APIService ObjectStorageAccessKeysAPI
	accessKey  string
}

// Execute calls the API using the request data configured.
func (r APIGetAccessKeyRequest) Execute() (*AccessKey, *http.Response, error) {
	return r.APIService.GetAccessKeyExecute(r)
}

/*
GetAccessKey Get an access key

Get an access key by its identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessKey
 @return APIGetAccessKeyRequest
*/
func (a *ObjectStorageAccessKeysAPIService) GetAccessKey(ctx context.Context, accessKey string) APIGetAccessKeyRequest {
	return APIGetAccessKeyRequest{
		APIService: a,
		ctx:        ctx,
		accessKey:  accessKey,
	}
}

// GetAccessKeyExecute executes the request
//  @return AccessKey
func (a *ObjectStorageAccessKeysAPIService) GetAccessKeyExecute(r APIGetAccessKeyRequest) (*AccessKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AccessKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageAccessKeysAPIService.GetAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/object-storage/access-keys/{access_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"access_key"+"}", gourl.PathEscape(parameterToString(r.accessKey, "")))

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

// APIListAccessKeysRequest represents a request for the resource.
type APIListAccessKeysRequest struct {
	ctx        context.Context
	APIService ObjectStorageAccessKeysAPI
}

// Execute calls the API using the request data configured.
func (r APIListAccessKeysRequest) Execute() (*AccessKeyResponse, *http.Response, error) {
	return r.APIService.ListAccessKeysExecute(r)
}

/*
ListAccessKeys List access keys

List access keys.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListAccessKeysRequest
*/
func (a *ObjectStorageAccessKeysAPIService) ListAccessKeys(ctx context.Context) APIListAccessKeysRequest {
	return APIListAccessKeysRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListAccessKeysExecute executes the request
//  @return AccessKeyResponse
func (a *ObjectStorageAccessKeysAPIService) ListAccessKeysExecute(r APIListAccessKeysRequest) (*AccessKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AccessKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageAccessKeysAPIService.ListAccessKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/object-storage/access-keys"

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
