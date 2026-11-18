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

// AiRuntimeControlVirtualKeysAPI defines an interface for interacting with the resource.
type AiRuntimeControlVirtualKeysAPI interface {

	/*
		ArcCreateVirtualKey Create a virtual key

		Create a new virtual key for authenticating requests through the AI Runtime Control gateway.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcCreateVirtualKeyRequest
	*/
	ArcCreateVirtualKey(ctx context.Context) APIArcCreateVirtualKeyRequest

	// ArcCreateVirtualKeyExecute executes the request
	//  @return ArcCreateVirtualKeyResponse
	ArcCreateVirtualKeyExecute(r APIArcCreateVirtualKeyRequest) (*ArcCreateVirtualKeyResponse, *http.Response, error)

	/*
		ArcDeleteVirtualKey Delete a virtual key

		Delete a virtual key. This also deletes any security configuration associated with the key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param keyID The ID of the virtual key.
		 @return APIArcDeleteVirtualKeyRequest
	*/
	ArcDeleteVirtualKey(ctx context.Context, keyID string) APIArcDeleteVirtualKeyRequest

	// ArcDeleteVirtualKeyExecute executes the request
	ArcDeleteVirtualKeyExecute(r APIArcDeleteVirtualKeyRequest) (*http.Response, error)

	/*
		ArcGetVirtualKey Get a virtual key

		Retrieve a single virtual key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param keyID The ID of the virtual key.
		 @return APIArcGetVirtualKeyRequest
	*/
	ArcGetVirtualKey(ctx context.Context, keyID string) APIArcGetVirtualKeyRequest

	// ArcGetVirtualKeyExecute executes the request
	//  @return ArcVirtualKeyResponse
	ArcGetVirtualKeyExecute(r APIArcGetVirtualKeyRequest) (*ArcVirtualKeyResponse, *http.Response, error)

	/*
		ArcGetVirtualKeyFailover Get a virtual key's failover policy

		Retrieve the failover policy for a virtual key. A failover policy lists alternate
	provider connections to retry a request against, and which error conditions trigger
	a retry.

	Failover is a per-account feature. If it isn't enabled on your account, this endpoint
	returns a 404 response.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param keyID The ID of the virtual key.
		 @return APIArcGetVirtualKeyFailoverRequest
	*/
	ArcGetVirtualKeyFailover(ctx context.Context, keyID string) APIArcGetVirtualKeyFailoverRequest

	// ArcGetVirtualKeyFailoverExecute executes the request
	//  @return ArcVirtualKeyFailoverResponse
	ArcGetVirtualKeyFailoverExecute(r APIArcGetVirtualKeyFailoverRequest) (*ArcVirtualKeyFailoverResponse, *http.Response, error)

	/*
		ArcListVirtualKeys List virtual keys

		List the virtual keys belonging to your account.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcListVirtualKeysRequest
	*/
	ArcListVirtualKeys(ctx context.Context) APIArcListVirtualKeysRequest

	// ArcListVirtualKeysExecute executes the request
	//  @return ArcVirtualKeysResponse
	ArcListVirtualKeysExecute(r APIArcListVirtualKeysRequest) (*ArcVirtualKeysResponse, *http.Response, error)

	/*
		ArcRotateVirtualKey Rotate a virtual key

		Issue a new access token for a virtual key, invalidating the previous one.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param keyID The ID of the virtual key.
		 @return APIArcRotateVirtualKeyRequest
	*/
	ArcRotateVirtualKey(ctx context.Context, keyID string) APIArcRotateVirtualKeyRequest

	// ArcRotateVirtualKeyExecute executes the request
	//  @return ArcRotateVirtualKeyResponse
	ArcRotateVirtualKeyExecute(r APIArcRotateVirtualKeyRequest) (*ArcRotateVirtualKeyResponse, *http.Response, error)

	/*
		ArcUpdateVirtualKey Update a virtual key

		Update the attributes of a virtual key. Only the fields you include are changed.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param keyID The ID of the virtual key.
		 @return APIArcUpdateVirtualKeyRequest
	*/
	ArcUpdateVirtualKey(ctx context.Context, keyID string) APIArcUpdateVirtualKeyRequest

	// ArcUpdateVirtualKeyExecute executes the request
	//  @return ArcUpdateVirtualKeyResponse
	ArcUpdateVirtualKeyExecute(r APIArcUpdateVirtualKeyRequest) (*ArcUpdateVirtualKeyResponse, *http.Response, error)

	/*
		ArcUpdateVirtualKeyFailover Update a virtual key's failover policy

		Replace the failover policy for a virtual key.

	Failover is a per-account feature. If it isn't enabled on your account, this endpoint
	returns a 404 response.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param keyID The ID of the virtual key.
		 @return APIArcUpdateVirtualKeyFailoverRequest
	*/
	ArcUpdateVirtualKeyFailover(ctx context.Context, keyID string) APIArcUpdateVirtualKeyFailoverRequest

	// ArcUpdateVirtualKeyFailoverExecute executes the request
	//  @return ArcVirtualKeyFailoverResponse
	ArcUpdateVirtualKeyFailoverExecute(r APIArcUpdateVirtualKeyFailoverRequest) (*ArcVirtualKeyFailoverResponse, *http.Response, error)
}

// AiRuntimeControlVirtualKeysAPIService AiRuntimeControlVirtualKeysAPI service
type AiRuntimeControlVirtualKeysAPIService service

// APIArcCreateVirtualKeyRequest represents a request for the resource.
type APIArcCreateVirtualKeyRequest struct {
	ctx           context.Context
	APIService    AiRuntimeControlVirtualKeysAPI
	arcVirtualKey *ArcVirtualKey
}

// ArcVirtualKey returns a pointer to a request.
func (r *APIArcCreateVirtualKeyRequest) ArcVirtualKey(arcVirtualKey ArcVirtualKey) *APIArcCreateVirtualKeyRequest {
	r.arcVirtualKey = &arcVirtualKey
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcCreateVirtualKeyRequest) Execute() (*ArcCreateVirtualKeyResponse, *http.Response, error) {
	return r.APIService.ArcCreateVirtualKeyExecute(r)
}

/*
ArcCreateVirtualKey Create a virtual key

Create a new virtual key for authenticating requests through the AI Runtime Control gateway.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcCreateVirtualKeyRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcCreateVirtualKey(ctx context.Context) APIArcCreateVirtualKeyRequest {
	return APIArcCreateVirtualKeyRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcCreateVirtualKeyExecute executes the request
//  @return ArcCreateVirtualKeyResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcCreateVirtualKeyExecute(r APIArcCreateVirtualKeyRequest) (*ArcCreateVirtualKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcCreateVirtualKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcCreateVirtualKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.arcVirtualKey == nil {
		return localVarReturnValue, nil, reportError("arcVirtualKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.arcVirtualKey
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcDeleteVirtualKeyRequest represents a request for the resource.
type APIArcDeleteVirtualKeyRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlVirtualKeysAPI
	keyID      string
}

// Execute calls the API using the request data configured.
func (r APIArcDeleteVirtualKeyRequest) Execute() (*http.Response, error) {
	return r.APIService.ArcDeleteVirtualKeyExecute(r)
}

/*
ArcDeleteVirtualKey Delete a virtual key

Delete a virtual key. This also deletes any security configuration associated with the key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The ID of the virtual key.
 @return APIArcDeleteVirtualKeyRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcDeleteVirtualKey(ctx context.Context, keyID string) APIArcDeleteVirtualKeyRequest {
	return APIArcDeleteVirtualKeyRequest{
		APIService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// ArcDeleteVirtualKeyExecute executes the request
func (a *AiRuntimeControlVirtualKeysAPIService) ArcDeleteVirtualKeyExecute(r APIArcDeleteVirtualKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcDeleteVirtualKey")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys/{key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_id"+"}", gourl.PathEscape(parameterToString(r.keyID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcGetVirtualKeyRequest represents a request for the resource.
type APIArcGetVirtualKeyRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlVirtualKeysAPI
	keyID      string
}

// Execute calls the API using the request data configured.
func (r APIArcGetVirtualKeyRequest) Execute() (*ArcVirtualKeyResponse, *http.Response, error) {
	return r.APIService.ArcGetVirtualKeyExecute(r)
}

/*
ArcGetVirtualKey Get a virtual key

Retrieve a single virtual key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The ID of the virtual key.
 @return APIArcGetVirtualKeyRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcGetVirtualKey(ctx context.Context, keyID string) APIArcGetVirtualKeyRequest {
	return APIArcGetVirtualKeyRequest{
		APIService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// ArcGetVirtualKeyExecute executes the request
//  @return ArcVirtualKeyResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcGetVirtualKeyExecute(r APIArcGetVirtualKeyRequest) (*ArcVirtualKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcVirtualKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcGetVirtualKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys/{key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_id"+"}", gourl.PathEscape(parameterToString(r.keyID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcGetVirtualKeyFailoverRequest represents a request for the resource.
type APIArcGetVirtualKeyFailoverRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlVirtualKeysAPI
	keyID      string
}

// Execute calls the API using the request data configured.
func (r APIArcGetVirtualKeyFailoverRequest) Execute() (*ArcVirtualKeyFailoverResponse, *http.Response, error) {
	return r.APIService.ArcGetVirtualKeyFailoverExecute(r)
}

/*
ArcGetVirtualKeyFailover Get a virtual key's failover policy

Retrieve the failover policy for a virtual key. A failover policy lists alternate
provider connections to retry a request against, and which error conditions trigger
a retry.

Failover is a per-account feature. If it isn't enabled on your account, this endpoint
returns a 404 response.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The ID of the virtual key.
 @return APIArcGetVirtualKeyFailoverRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcGetVirtualKeyFailover(ctx context.Context, keyID string) APIArcGetVirtualKeyFailoverRequest {
	return APIArcGetVirtualKeyFailoverRequest{
		APIService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// ArcGetVirtualKeyFailoverExecute executes the request
//  @return ArcVirtualKeyFailoverResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcGetVirtualKeyFailoverExecute(r APIArcGetVirtualKeyFailoverRequest) (*ArcVirtualKeyFailoverResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcVirtualKeyFailoverResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcGetVirtualKeyFailover")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys/{key_id}/failover"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_id"+"}", gourl.PathEscape(parameterToString(r.keyID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcListVirtualKeysRequest represents a request for the resource.
type APIArcListVirtualKeysRequest struct {
	ctx            context.Context
	APIService     AiRuntimeControlVirtualKeysAPI
	includeDeleted *bool
	search         *string
	cursor         *string
	limit          *int32
	sort           *string
}

// IncludeDeleted Include deleted virtual keys in the results.
func (r *APIArcListVirtualKeysRequest) IncludeDeleted(includeDeleted bool) *APIArcListVirtualKeysRequest {
	r.includeDeleted = &includeDeleted
	return r
}

// Search Filter results to virtual keys whose name contains this substring.
func (r *APIArcListVirtualKeysRequest) Search(search string) *APIArcListVirtualKeysRequest {
	r.search = &search
	return r
}

// Cursor Cursor for pagination. Use the &#x60;next_cursor&#x60; value from a previous response to request the next page of results, or the &#x60;prev_cursor&#x60; value to request the previous page.
func (r *APIArcListVirtualKeysRequest) Cursor(cursor string) *APIArcListVirtualKeysRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIArcListVirtualKeysRequest) Limit(limit int32) *APIArcListVirtualKeysRequest {
	r.limit = &limit
	return r
}

// Sort Sort field. Prefix with &#x60;-&#x60; for descending order (e.g., &#x60;sort&#x3D;-created_at&#x60;).
func (r *APIArcListVirtualKeysRequest) Sort(sort string) *APIArcListVirtualKeysRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcListVirtualKeysRequest) Execute() (*ArcVirtualKeysResponse, *http.Response, error) {
	return r.APIService.ArcListVirtualKeysExecute(r)
}

/*
ArcListVirtualKeys List virtual keys

List the virtual keys belonging to your account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcListVirtualKeysRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcListVirtualKeys(ctx context.Context) APIArcListVirtualKeysRequest {
	return APIArcListVirtualKeysRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcListVirtualKeysExecute executes the request
//  @return ArcVirtualKeysResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcListVirtualKeysExecute(r APIArcListVirtualKeysRequest) (*ArcVirtualKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcVirtualKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcListVirtualKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.includeDeleted != nil {
		localVarQueryParams.Add("include_deleted", parameterToString(*r.includeDeleted, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcRotateVirtualKeyRequest represents a request for the resource.
type APIArcRotateVirtualKeyRequest struct {
	ctx                        context.Context
	APIService                 AiRuntimeControlVirtualKeysAPI
	keyID                      string
	arcRotateVirtualKeyRequest *ArcRotateVirtualKeyRequest
}

// ArcRotateVirtualKeyRequest returns a pointer to a request.
func (r *APIArcRotateVirtualKeyRequest) ArcRotateVirtualKeyRequest(arcRotateVirtualKeyRequest ArcRotateVirtualKeyRequest) *APIArcRotateVirtualKeyRequest {
	r.arcRotateVirtualKeyRequest = &arcRotateVirtualKeyRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcRotateVirtualKeyRequest) Execute() (*ArcRotateVirtualKeyResponse, *http.Response, error) {
	return r.APIService.ArcRotateVirtualKeyExecute(r)
}

/*
ArcRotateVirtualKey Rotate a virtual key

Issue a new access token for a virtual key, invalidating the previous one.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The ID of the virtual key.
 @return APIArcRotateVirtualKeyRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcRotateVirtualKey(ctx context.Context, keyID string) APIArcRotateVirtualKeyRequest {
	return APIArcRotateVirtualKeyRequest{
		APIService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// ArcRotateVirtualKeyExecute executes the request
//  @return ArcRotateVirtualKeyResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcRotateVirtualKeyExecute(r APIArcRotateVirtualKeyRequest) (*ArcRotateVirtualKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcRotateVirtualKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcRotateVirtualKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys/{key_id}/rotate"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_id"+"}", gourl.PathEscape(parameterToString(r.keyID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.arcRotateVirtualKeyRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcUpdateVirtualKeyRequest represents a request for the resource.
type APIArcUpdateVirtualKeyRequest struct {
	ctx                        context.Context
	APIService                 AiRuntimeControlVirtualKeysAPI
	keyID                      string
	arcUpdateVirtualKeyRequest *ArcUpdateVirtualKeyRequest
}

// ArcUpdateVirtualKeyRequest returns a pointer to a request.
func (r *APIArcUpdateVirtualKeyRequest) ArcUpdateVirtualKeyRequest(arcUpdateVirtualKeyRequest ArcUpdateVirtualKeyRequest) *APIArcUpdateVirtualKeyRequest {
	r.arcUpdateVirtualKeyRequest = &arcUpdateVirtualKeyRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcUpdateVirtualKeyRequest) Execute() (*ArcUpdateVirtualKeyResponse, *http.Response, error) {
	return r.APIService.ArcUpdateVirtualKeyExecute(r)
}

/*
ArcUpdateVirtualKey Update a virtual key

Update the attributes of a virtual key. Only the fields you include are changed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The ID of the virtual key.
 @return APIArcUpdateVirtualKeyRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcUpdateVirtualKey(ctx context.Context, keyID string) APIArcUpdateVirtualKeyRequest {
	return APIArcUpdateVirtualKeyRequest{
		APIService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// ArcUpdateVirtualKeyExecute executes the request
//  @return ArcUpdateVirtualKeyResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcUpdateVirtualKeyExecute(r APIArcUpdateVirtualKeyRequest) (*ArcUpdateVirtualKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcUpdateVirtualKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcUpdateVirtualKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys/{key_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_id"+"}", gourl.PathEscape(parameterToString(r.keyID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.arcUpdateVirtualKeyRequest == nil {
		return localVarReturnValue, nil, reportError("arcUpdateVirtualKeyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.arcUpdateVirtualKeyRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcUpdateVirtualKeyFailoverRequest represents a request for the resource.
type APIArcUpdateVirtualKeyFailoverRequest struct {
	ctx                          context.Context
	APIService                   AiRuntimeControlVirtualKeysAPI
	keyID                        string
	arcVirtualKeyFailoverRequest *ArcVirtualKeyFailoverRequest
}

// ArcVirtualKeyFailoverRequest returns a pointer to a request.
func (r *APIArcUpdateVirtualKeyFailoverRequest) ArcVirtualKeyFailoverRequest(arcVirtualKeyFailoverRequest ArcVirtualKeyFailoverRequest) *APIArcUpdateVirtualKeyFailoverRequest {
	r.arcVirtualKeyFailoverRequest = &arcVirtualKeyFailoverRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcUpdateVirtualKeyFailoverRequest) Execute() (*ArcVirtualKeyFailoverResponse, *http.Response, error) {
	return r.APIService.ArcUpdateVirtualKeyFailoverExecute(r)
}

/*
ArcUpdateVirtualKeyFailover Update a virtual key's failover policy

Replace the failover policy for a virtual key.

Failover is a per-account feature. If it isn't enabled on your account, this endpoint
returns a 404 response.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The ID of the virtual key.
 @return APIArcUpdateVirtualKeyFailoverRequest
*/
func (a *AiRuntimeControlVirtualKeysAPIService) ArcUpdateVirtualKeyFailover(ctx context.Context, keyID string) APIArcUpdateVirtualKeyFailoverRequest {
	return APIArcUpdateVirtualKeyFailoverRequest{
		APIService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// ArcUpdateVirtualKeyFailoverExecute executes the request
//  @return ArcVirtualKeyFailoverResponse
func (a *AiRuntimeControlVirtualKeysAPIService) ArcUpdateVirtualKeyFailoverExecute(r APIArcUpdateVirtualKeyFailoverRequest) (*ArcVirtualKeyFailoverResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcVirtualKeyFailoverResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlVirtualKeysAPIService.ArcUpdateVirtualKeyFailover")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/keys/{key_id}/failover"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"key_id"+"}", gourl.PathEscape(parameterToString(r.keyID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.arcVirtualKeyFailoverRequest == nil {
		return localVarReturnValue, nil, reportError("arcVirtualKeyFailoverRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.arcVirtualKeyFailoverRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
