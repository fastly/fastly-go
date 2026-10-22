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

// ProductAiRuntimeControlAPI defines an interface for interacting with the resource.
type ProductAiRuntimeControlAPI interface {

	/*
		DisableProductAiRuntimeControl Disable product

		Disable the AI Runtime Control product

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIDisableProductAiRuntimeControlRequest
	*/
	DisableProductAiRuntimeControl(ctx context.Context) APIDisableProductAiRuntimeControlRequest

	// DisableProductAiRuntimeControlExecute executes the request
	DisableProductAiRuntimeControlExecute(r APIDisableProductAiRuntimeControlRequest) (*http.Response, error)

	/*
		EnableProductAiRuntimeControl Enable product

		Enable the AI Runtime Control product

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIEnableProductAiRuntimeControlRequest
	*/
	EnableProductAiRuntimeControl(ctx context.Context) APIEnableProductAiRuntimeControlRequest

	// EnableProductAiRuntimeControlExecute executes the request
	//  @return AiRuntimeControlResponseBodyEnable
	EnableProductAiRuntimeControlExecute(r APIEnableProductAiRuntimeControlRequest) (*AiRuntimeControlResponseBodyEnable, *http.Response, error)

	/*
		GetProductAiRuntimeControl Get product enablement status

		Get the enablement status of the AI Runtime Control product

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetProductAiRuntimeControlRequest
	*/
	GetProductAiRuntimeControl(ctx context.Context) APIGetProductAiRuntimeControlRequest

	// GetProductAiRuntimeControlExecute executes the request
	//  @return AiRuntimeControlResponseBodyEnable
	GetProductAiRuntimeControlExecute(r APIGetProductAiRuntimeControlRequest) (*AiRuntimeControlResponseBodyEnable, *http.Response, error)
}

// ProductAiRuntimeControlAPIService ProductAiRuntimeControlAPI service
type ProductAiRuntimeControlAPIService service

// APIDisableProductAiRuntimeControlRequest represents a request for the resource.
type APIDisableProductAiRuntimeControlRequest struct {
	ctx        context.Context
	APIService ProductAiRuntimeControlAPI
}

// Execute calls the API using the request data configured.
func (r APIDisableProductAiRuntimeControlRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductAiRuntimeControlExecute(r)
}

/*
DisableProductAiRuntimeControl Disable product

Disable the AI Runtime Control product

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIDisableProductAiRuntimeControlRequest
*/
func (a *ProductAiRuntimeControlAPIService) DisableProductAiRuntimeControl(ctx context.Context) APIDisableProductAiRuntimeControlRequest {
	return APIDisableProductAiRuntimeControlRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// DisableProductAiRuntimeControlExecute executes the request
func (a *ProductAiRuntimeControlAPIService) DisableProductAiRuntimeControlExecute(r APIDisableProductAiRuntimeControlRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAiRuntimeControlAPIService.DisableProductAiRuntimeControl")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ai_runtime_control"

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

// APIEnableProductAiRuntimeControlRequest represents a request for the resource.
type APIEnableProductAiRuntimeControlRequest struct {
	ctx        context.Context
	APIService ProductAiRuntimeControlAPI
}

// Execute calls the API using the request data configured.
func (r APIEnableProductAiRuntimeControlRequest) Execute() (*AiRuntimeControlResponseBodyEnable, *http.Response, error) {
	return r.APIService.EnableProductAiRuntimeControlExecute(r)
}

/*
EnableProductAiRuntimeControl Enable product

Enable the AI Runtime Control product

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIEnableProductAiRuntimeControlRequest
*/
func (a *ProductAiRuntimeControlAPIService) EnableProductAiRuntimeControl(ctx context.Context) APIEnableProductAiRuntimeControlRequest {
	return APIEnableProductAiRuntimeControlRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// EnableProductAiRuntimeControlExecute executes the request
//  @return AiRuntimeControlResponseBodyEnable
func (a *ProductAiRuntimeControlAPIService) EnableProductAiRuntimeControlExecute(r APIEnableProductAiRuntimeControlRequest) (*AiRuntimeControlResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiRuntimeControlResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAiRuntimeControlAPIService.EnableProductAiRuntimeControl")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ai_runtime_control"

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

// APIGetProductAiRuntimeControlRequest represents a request for the resource.
type APIGetProductAiRuntimeControlRequest struct {
	ctx        context.Context
	APIService ProductAiRuntimeControlAPI
}

// Execute calls the API using the request data configured.
func (r APIGetProductAiRuntimeControlRequest) Execute() (*AiRuntimeControlResponseBodyEnable, *http.Response, error) {
	return r.APIService.GetProductAiRuntimeControlExecute(r)
}

/*
GetProductAiRuntimeControl Get product enablement status

Get the enablement status of the AI Runtime Control product

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetProductAiRuntimeControlRequest
*/
func (a *ProductAiRuntimeControlAPIService) GetProductAiRuntimeControl(ctx context.Context) APIGetProductAiRuntimeControlRequest {
	return APIGetProductAiRuntimeControlRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetProductAiRuntimeControlExecute executes the request
//  @return AiRuntimeControlResponseBodyEnable
func (a *ProductAiRuntimeControlAPIService) GetProductAiRuntimeControlExecute(r APIGetProductAiRuntimeControlRequest) (*AiRuntimeControlResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiRuntimeControlResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAiRuntimeControlAPIService.GetProductAiRuntimeControl")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ai_runtime_control"

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
