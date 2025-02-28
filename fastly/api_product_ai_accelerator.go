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
)

// Linger please
var (
	_ context.Context
)

// ProductAiAcceleratorAPI defines an interface for interacting with the resource.
type ProductAiAcceleratorAPI interface {

	/*
		DisableProductAiAccelerator Disable product

		Disable the AI Accelerator product

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIDisableProductAiAcceleratorRequest
	*/
	DisableProductAiAccelerator(ctx context.Context) APIDisableProductAiAcceleratorRequest

	// DisableProductAiAcceleratorExecute executes the request
	DisableProductAiAcceleratorExecute(r APIDisableProductAiAcceleratorRequest) (*http.Response, error)

	/*
		EnableAiAccelerator Enable product

		Enable the AI Accelerator product

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIEnableAiAcceleratorRequest
	*/
	EnableAiAccelerator(ctx context.Context) APIEnableAiAcceleratorRequest

	// EnableAiAcceleratorExecute executes the request
	//  @return AiAcceleratorResponseBodyEnable
	EnableAiAcceleratorExecute(r APIEnableAiAcceleratorRequest) (*AiAcceleratorResponseBodyEnable, *http.Response, error)

	/*
		GetAiAccelerator Get product enablement status

		Get the enablement status of the AI Accelerator product

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetAiAcceleratorRequest
	*/
	GetAiAccelerator(ctx context.Context) APIGetAiAcceleratorRequest

	// GetAiAcceleratorExecute executes the request
	//  @return AiAcceleratorResponseBodyEnable
	GetAiAcceleratorExecute(r APIGetAiAcceleratorRequest) (*AiAcceleratorResponseBodyEnable, *http.Response, error)
}

// ProductAiAcceleratorAPIService ProductAiAcceleratorAPI service
type ProductAiAcceleratorAPIService service

// APIDisableProductAiAcceleratorRequest represents a request for the resource.
type APIDisableProductAiAcceleratorRequest struct {
	ctx        context.Context
	APIService ProductAiAcceleratorAPI
}

// Execute calls the API using the request data configured.
func (r APIDisableProductAiAcceleratorRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductAiAcceleratorExecute(r)
}

/*
DisableProductAiAccelerator Disable product

Disable the AI Accelerator product

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIDisableProductAiAcceleratorRequest
*/
func (a *ProductAiAcceleratorAPIService) DisableProductAiAccelerator(ctx context.Context) APIDisableProductAiAcceleratorRequest {
	return APIDisableProductAiAcceleratorRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// DisableProductAiAcceleratorExecute executes the request
func (a *ProductAiAcceleratorAPIService) DisableProductAiAcceleratorExecute(r APIDisableProductAiAcceleratorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAiAcceleratorAPIService.DisableProductAiAccelerator")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ai_accelerator"

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

// APIEnableAiAcceleratorRequest represents a request for the resource.
type APIEnableAiAcceleratorRequest struct {
	ctx        context.Context
	APIService ProductAiAcceleratorAPI
}

// Execute calls the API using the request data configured.
func (r APIEnableAiAcceleratorRequest) Execute() (*AiAcceleratorResponseBodyEnable, *http.Response, error) {
	return r.APIService.EnableAiAcceleratorExecute(r)
}

/*
EnableAiAccelerator Enable product

Enable the AI Accelerator product

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIEnableAiAcceleratorRequest
*/
func (a *ProductAiAcceleratorAPIService) EnableAiAccelerator(ctx context.Context) APIEnableAiAcceleratorRequest {
	return APIEnableAiAcceleratorRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// EnableAiAcceleratorExecute executes the request
//  @return AiAcceleratorResponseBodyEnable
func (a *ProductAiAcceleratorAPIService) EnableAiAcceleratorExecute(r APIEnableAiAcceleratorRequest) (*AiAcceleratorResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiAcceleratorResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAiAcceleratorAPIService.EnableAiAccelerator")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ai_accelerator"

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

// APIGetAiAcceleratorRequest represents a request for the resource.
type APIGetAiAcceleratorRequest struct {
	ctx        context.Context
	APIService ProductAiAcceleratorAPI
}

// Execute calls the API using the request data configured.
func (r APIGetAiAcceleratorRequest) Execute() (*AiAcceleratorResponseBodyEnable, *http.Response, error) {
	return r.APIService.GetAiAcceleratorExecute(r)
}

/*
GetAiAccelerator Get product enablement status

Get the enablement status of the AI Accelerator product

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetAiAcceleratorRequest
*/
func (a *ProductAiAcceleratorAPIService) GetAiAccelerator(ctx context.Context) APIGetAiAcceleratorRequest {
	return APIGetAiAcceleratorRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetAiAcceleratorExecute executes the request
//  @return AiAcceleratorResponseBodyEnable
func (a *ProductAiAcceleratorAPIService) GetAiAcceleratorExecute(r APIGetAiAcceleratorRequest) (*AiAcceleratorResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiAcceleratorResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAiAcceleratorAPIService.GetAiAccelerator")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ai_accelerator"

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
