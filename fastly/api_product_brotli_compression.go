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

// ProductBrotliCompressionAPI defines an interface for interacting with the resource.
type ProductBrotliCompressionAPI interface {

	/*
		DisableProductBrotliCompression Disable product

		Disable the Brotli Compression product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIDisableProductBrotliCompressionRequest
	*/
	DisableProductBrotliCompression(ctx context.Context, serviceID string) APIDisableProductBrotliCompressionRequest

	// DisableProductBrotliCompressionExecute executes the request
	DisableProductBrotliCompressionExecute(r APIDisableProductBrotliCompressionRequest) (*http.Response, error)

	/*
		EnableProductBrotliCompression Enable product

		Enable the Brotli Compression product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIEnableProductBrotliCompressionRequest
	*/
	EnableProductBrotliCompression(ctx context.Context, serviceID string) APIEnableProductBrotliCompressionRequest

	// EnableProductBrotliCompressionExecute executes the request
	//  @return BrotliCompressionResponseBodyEnable
	EnableProductBrotliCompressionExecute(r APIEnableProductBrotliCompressionRequest) (*BrotliCompressionResponseBodyEnable, *http.Response, error)

	/*
		GetProductBrotliCompression Get product enablement status

		Get the enablement status of the Brotli Compression product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetProductBrotliCompressionRequest
	*/
	GetProductBrotliCompression(ctx context.Context, serviceID string) APIGetProductBrotliCompressionRequest

	// GetProductBrotliCompressionExecute executes the request
	//  @return BrotliCompressionResponseBodyEnable
	GetProductBrotliCompressionExecute(r APIGetProductBrotliCompressionRequest) (*BrotliCompressionResponseBodyEnable, *http.Response, error)

	/*
		GetServicesProductBrotliCompression Get services with product enabled

		Get all the services which have the Brotli Compression product enabled.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetServicesProductBrotliCompressionRequest
	*/
	GetServicesProductBrotliCompression(ctx context.Context) APIGetServicesProductBrotliCompressionRequest

	// GetServicesProductBrotliCompressionExecute executes the request
	//  @return BrotliCompressionResponseBodyGetAllServices
	GetServicesProductBrotliCompressionExecute(r APIGetServicesProductBrotliCompressionRequest) (*BrotliCompressionResponseBodyGetAllServices, *http.Response, error)
}

// ProductBrotliCompressionAPIService ProductBrotliCompressionAPI service
type ProductBrotliCompressionAPIService service

// APIDisableProductBrotliCompressionRequest represents a request for the resource.
type APIDisableProductBrotliCompressionRequest struct {
	ctx        context.Context
	APIService ProductBrotliCompressionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIDisableProductBrotliCompressionRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductBrotliCompressionExecute(r)
}

/*
DisableProductBrotliCompression Disable product

Disable the Brotli Compression product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIDisableProductBrotliCompressionRequest
*/
func (a *ProductBrotliCompressionAPIService) DisableProductBrotliCompression(ctx context.Context, serviceID string) APIDisableProductBrotliCompressionRequest {
	return APIDisableProductBrotliCompressionRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// DisableProductBrotliCompressionExecute executes the request
func (a *ProductBrotliCompressionAPIService) DisableProductBrotliCompressionExecute(r APIDisableProductBrotliCompressionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductBrotliCompressionAPIService.DisableProductBrotliCompression")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/brotli_compression/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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

// APIEnableProductBrotliCompressionRequest represents a request for the resource.
type APIEnableProductBrotliCompressionRequest struct {
	ctx        context.Context
	APIService ProductBrotliCompressionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIEnableProductBrotliCompressionRequest) Execute() (*BrotliCompressionResponseBodyEnable, *http.Response, error) {
	return r.APIService.EnableProductBrotliCompressionExecute(r)
}

/*
EnableProductBrotliCompression Enable product

Enable the Brotli Compression product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIEnableProductBrotliCompressionRequest
*/
func (a *ProductBrotliCompressionAPIService) EnableProductBrotliCompression(ctx context.Context, serviceID string) APIEnableProductBrotliCompressionRequest {
	return APIEnableProductBrotliCompressionRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// EnableProductBrotliCompressionExecute executes the request
//  @return BrotliCompressionResponseBodyEnable
func (a *ProductBrotliCompressionAPIService) EnableProductBrotliCompressionExecute(r APIEnableProductBrotliCompressionRequest) (*BrotliCompressionResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BrotliCompressionResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductBrotliCompressionAPIService.EnableProductBrotliCompression")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/brotli_compression/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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

// APIGetProductBrotliCompressionRequest represents a request for the resource.
type APIGetProductBrotliCompressionRequest struct {
	ctx        context.Context
	APIService ProductBrotliCompressionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductBrotliCompressionRequest) Execute() (*BrotliCompressionResponseBodyEnable, *http.Response, error) {
	return r.APIService.GetProductBrotliCompressionExecute(r)
}

/*
GetProductBrotliCompression Get product enablement status

Get the enablement status of the Brotli Compression product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetProductBrotliCompressionRequest
*/
func (a *ProductBrotliCompressionAPIService) GetProductBrotliCompression(ctx context.Context, serviceID string) APIGetProductBrotliCompressionRequest {
	return APIGetProductBrotliCompressionRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// GetProductBrotliCompressionExecute executes the request
//  @return BrotliCompressionResponseBodyEnable
func (a *ProductBrotliCompressionAPIService) GetProductBrotliCompressionExecute(r APIGetProductBrotliCompressionRequest) (*BrotliCompressionResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BrotliCompressionResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductBrotliCompressionAPIService.GetProductBrotliCompression")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/brotli_compression/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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

// APIGetServicesProductBrotliCompressionRequest represents a request for the resource.
type APIGetServicesProductBrotliCompressionRequest struct {
	ctx        context.Context
	APIService ProductBrotliCompressionAPI
}

// Execute calls the API using the request data configured.
func (r APIGetServicesProductBrotliCompressionRequest) Execute() (*BrotliCompressionResponseBodyGetAllServices, *http.Response, error) {
	return r.APIService.GetServicesProductBrotliCompressionExecute(r)
}

/*
GetServicesProductBrotliCompression Get services with product enabled

Get all the services which have the Brotli Compression product enabled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetServicesProductBrotliCompressionRequest
*/
func (a *ProductBrotliCompressionAPIService) GetServicesProductBrotliCompression(ctx context.Context) APIGetServicesProductBrotliCompressionRequest {
	return APIGetServicesProductBrotliCompressionRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetServicesProductBrotliCompressionExecute executes the request
//  @return BrotliCompressionResponseBodyGetAllServices
func (a *ProductBrotliCompressionAPIService) GetServicesProductBrotliCompressionExecute(r APIGetServicesProductBrotliCompressionRequest) (*BrotliCompressionResponseBodyGetAllServices, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BrotliCompressionResponseBodyGetAllServices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductBrotliCompressionAPIService.GetServicesProductBrotliCompression")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/brotli_compression/services"

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
