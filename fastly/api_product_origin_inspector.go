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

// ProductOriginInspectorAPI defines an interface for interacting with the resource.
type ProductOriginInspectorAPI interface {

	/*
		DisableProductOriginInspector Disable product

		Disable the Origin Inspector product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIDisableProductOriginInspectorRequest
	*/
	DisableProductOriginInspector(ctx context.Context, serviceId string) APIDisableProductOriginInspectorRequest

	// DisableProductOriginInspectorExecute executes the request
	DisableProductOriginInspectorExecute(r APIDisableProductOriginInspectorRequest) (*http.Response, error)

	/*
		EnableProductOriginInspector Enable product

		Enable the Origin Inspector product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIEnableProductOriginInspectorRequest
	*/
	EnableProductOriginInspector(ctx context.Context, serviceId string) APIEnableProductOriginInspectorRequest

	// EnableProductOriginInspectorExecute executes the request
	//  @return OriginInspectorResponseBodyEnable
	EnableProductOriginInspectorExecute(r APIEnableProductOriginInspectorRequest) (*OriginInspectorResponseBodyEnable, *http.Response, error)

	/*
		GetProductOriginInspector Get product enablement status

		Get the enablement status of the Origin Inspector product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIGetProductOriginInspectorRequest
	*/
	GetProductOriginInspector(ctx context.Context, serviceId string) APIGetProductOriginInspectorRequest

	// GetProductOriginInspectorExecute executes the request
	//  @return OriginInspectorResponseBodyEnable
	GetProductOriginInspectorExecute(r APIGetProductOriginInspectorRequest) (*OriginInspectorResponseBodyEnable, *http.Response, error)

	/*
		GetServicesProductOriginInspector Get services with product enabled

		Get all the services which have the Origin Inspector product enabled.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetServicesProductOriginInspectorRequest
	*/
	GetServicesProductOriginInspector(ctx context.Context) APIGetServicesProductOriginInspectorRequest

	// GetServicesProductOriginInspectorExecute executes the request
	//  @return OriginInspectorResponseBodyGetAllServices
	GetServicesProductOriginInspectorExecute(r APIGetServicesProductOriginInspectorRequest) (*OriginInspectorResponseBodyGetAllServices, *http.Response, error)
}

// ProductOriginInspectorAPIService ProductOriginInspectorAPI service
type ProductOriginInspectorAPIService service

// APIDisableProductOriginInspectorRequest represents a request for the resource.
type APIDisableProductOriginInspectorRequest struct {
	ctx        context.Context
	APIService ProductOriginInspectorAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIDisableProductOriginInspectorRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductOriginInspectorExecute(r)
}

/*
DisableProductOriginInspector Disable product

Disable the Origin Inspector product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIDisableProductOriginInspectorRequest
*/
func (a *ProductOriginInspectorAPIService) DisableProductOriginInspector(ctx context.Context, serviceId string) APIDisableProductOriginInspectorRequest {
	return APIDisableProductOriginInspectorRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// DisableProductOriginInspectorExecute executes the request
func (a *ProductOriginInspectorAPIService) DisableProductOriginInspectorExecute(r APIDisableProductOriginInspectorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductOriginInspectorAPIService.DisableProductOriginInspector")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/origin_inspector/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

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

// APIEnableProductOriginInspectorRequest represents a request for the resource.
type APIEnableProductOriginInspectorRequest struct {
	ctx        context.Context
	APIService ProductOriginInspectorAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIEnableProductOriginInspectorRequest) Execute() (*OriginInspectorResponseBodyEnable, *http.Response, error) {
	return r.APIService.EnableProductOriginInspectorExecute(r)
}

/*
EnableProductOriginInspector Enable product

Enable the Origin Inspector product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIEnableProductOriginInspectorRequest
*/
func (a *ProductOriginInspectorAPIService) EnableProductOriginInspector(ctx context.Context, serviceId string) APIEnableProductOriginInspectorRequest {
	return APIEnableProductOriginInspectorRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// EnableProductOriginInspectorExecute executes the request
//  @return OriginInspectorResponseBodyEnable
func (a *ProductOriginInspectorAPIService) EnableProductOriginInspectorExecute(r APIEnableProductOriginInspectorRequest) (*OriginInspectorResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OriginInspectorResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductOriginInspectorAPIService.EnableProductOriginInspector")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/origin_inspector/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

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

// APIGetProductOriginInspectorRequest represents a request for the resource.
type APIGetProductOriginInspectorRequest struct {
	ctx        context.Context
	APIService ProductOriginInspectorAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductOriginInspectorRequest) Execute() (*OriginInspectorResponseBodyEnable, *http.Response, error) {
	return r.APIService.GetProductOriginInspectorExecute(r)
}

/*
GetProductOriginInspector Get product enablement status

Get the enablement status of the Origin Inspector product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIGetProductOriginInspectorRequest
*/
func (a *ProductOriginInspectorAPIService) GetProductOriginInspector(ctx context.Context, serviceId string) APIGetProductOriginInspectorRequest {
	return APIGetProductOriginInspectorRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// GetProductOriginInspectorExecute executes the request
//  @return OriginInspectorResponseBodyEnable
func (a *ProductOriginInspectorAPIService) GetProductOriginInspectorExecute(r APIGetProductOriginInspectorRequest) (*OriginInspectorResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OriginInspectorResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductOriginInspectorAPIService.GetProductOriginInspector")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/origin_inspector/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

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

// APIGetServicesProductOriginInspectorRequest represents a request for the resource.
type APIGetServicesProductOriginInspectorRequest struct {
	ctx        context.Context
	APIService ProductOriginInspectorAPI
}

// Execute calls the API using the request data configured.
func (r APIGetServicesProductOriginInspectorRequest) Execute() (*OriginInspectorResponseBodyGetAllServices, *http.Response, error) {
	return r.APIService.GetServicesProductOriginInspectorExecute(r)
}

/*
GetServicesProductOriginInspector Get services with product enabled

Get all the services which have the Origin Inspector product enabled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetServicesProductOriginInspectorRequest
*/
func (a *ProductOriginInspectorAPIService) GetServicesProductOriginInspector(ctx context.Context) APIGetServicesProductOriginInspectorRequest {
	return APIGetServicesProductOriginInspectorRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetServicesProductOriginInspectorExecute executes the request
//  @return OriginInspectorResponseBodyGetAllServices
func (a *ProductOriginInspectorAPIService) GetServicesProductOriginInspectorExecute(r APIGetServicesProductOriginInspectorRequest) (*OriginInspectorResponseBodyGetAllServices, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OriginInspectorResponseBodyGetAllServices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductOriginInspectorAPIService.GetServicesProductOriginInspector")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/origin_inspector/services"

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
