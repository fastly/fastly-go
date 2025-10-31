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

// ProductNgwafAPI defines an interface for interacting with the resource.
type ProductNgwafAPI interface {

	/*
		DisableProductNgwaf Disable product

		Disable the Next-Gen WAF product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIDisableProductNgwafRequest
	*/
	DisableProductNgwaf(ctx context.Context, serviceId string) APIDisableProductNgwafRequest

	// DisableProductNgwafExecute executes the request
	DisableProductNgwafExecute(r APIDisableProductNgwafRequest) (*http.Response, error)

	/*
		EnableProductNgwaf Enable product

		Enable the Next-Gen WAF product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIEnableProductNgwafRequest
	*/
	EnableProductNgwaf(ctx context.Context, serviceId string) APIEnableProductNgwafRequest

	// EnableProductNgwafExecute executes the request
	//  @return NgwafResponseEnable
	EnableProductNgwafExecute(r APIEnableProductNgwafRequest) (*NgwafResponseEnable, *http.Response, error)

	/*
		GetProductNgwaf Get product enablement status

		Get the enablement status of the Next-Gen WAF product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIGetProductNgwafRequest
	*/
	GetProductNgwaf(ctx context.Context, serviceId string) APIGetProductNgwafRequest

	// GetProductNgwafExecute executes the request
	//  @return NgwafResponseEnable
	GetProductNgwafExecute(r APIGetProductNgwafRequest) (*NgwafResponseEnable, *http.Response, error)

	/*
		GetProductNgwafConfiguration Get configuration

		Get configuration of the Next-Gen WAF product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIGetProductNgwafConfigurationRequest
	*/
	GetProductNgwafConfiguration(ctx context.Context, serviceId string) APIGetProductNgwafConfigurationRequest

	// GetProductNgwafConfigurationExecute executes the request
	//  @return NgwafResponseConfigure
	GetProductNgwafConfigurationExecute(r APIGetProductNgwafConfigurationRequest) (*NgwafResponseConfigure, *http.Response, error)

	/*
		GetServicesProductNgwaf Get services with product enabled

		Get all the services which have the Next-Gen WAF product enabled.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetServicesProductNgwafRequest
	*/
	GetServicesProductNgwaf(ctx context.Context) APIGetServicesProductNgwafRequest

	// GetServicesProductNgwafExecute executes the request
	//  @return NgwafResponseBodyGetAllServices
	GetServicesProductNgwafExecute(r APIGetServicesProductNgwafRequest) (*NgwafResponseBodyGetAllServices, *http.Response, error)

	/*
		SetProductNgwafConfiguration Update configuration

		Update configuration of the Next-Gen WAF product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APISetProductNgwafConfigurationRequest
	*/
	SetProductNgwafConfiguration(ctx context.Context, serviceId string) APISetProductNgwafConfigurationRequest

	// SetProductNgwafConfigurationExecute executes the request
	//  @return NgwafResponseConfigure
	SetProductNgwafConfigurationExecute(r APISetProductNgwafConfigurationRequest) (*NgwafResponseConfigure, *http.Response, error)
}

// ProductNgwafAPIService ProductNgwafAPI service
type ProductNgwafAPIService service

// APIDisableProductNgwafRequest represents a request for the resource.
type APIDisableProductNgwafRequest struct {
	ctx        context.Context
	APIService ProductNgwafAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIDisableProductNgwafRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductNgwafExecute(r)
}

/*
DisableProductNgwaf Disable product

Disable the Next-Gen WAF product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIDisableProductNgwafRequest
*/
func (a *ProductNgwafAPIService) DisableProductNgwaf(ctx context.Context, serviceId string) APIDisableProductNgwafRequest {
	return APIDisableProductNgwafRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// DisableProductNgwafExecute executes the request
func (a *ProductNgwafAPIService) DisableProductNgwafExecute(r APIDisableProductNgwafRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductNgwafAPIService.DisableProductNgwaf")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ngwaf/services/{service_id}"
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

// APIEnableProductNgwafRequest represents a request for the resource.
type APIEnableProductNgwafRequest struct {
	ctx                context.Context
	APIService         ProductNgwafAPI
	serviceId          string
	ngwafRequestEnable *NgwafRequestEnable
}

// NgwafRequestEnable returns a pointer to a request.
func (r *APIEnableProductNgwafRequest) NgwafRequestEnable(ngwafRequestEnable NgwafRequestEnable) *APIEnableProductNgwafRequest {
	r.ngwafRequestEnable = &ngwafRequestEnable
	return r
}

// Execute calls the API using the request data configured.
func (r APIEnableProductNgwafRequest) Execute() (*NgwafResponseEnable, *http.Response, error) {
	return r.APIService.EnableProductNgwafExecute(r)
}

/*
EnableProductNgwaf Enable product

Enable the Next-Gen WAF product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIEnableProductNgwafRequest
*/
func (a *ProductNgwafAPIService) EnableProductNgwaf(ctx context.Context, serviceId string) APIEnableProductNgwafRequest {
	return APIEnableProductNgwafRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// EnableProductNgwafExecute executes the request
//  @return NgwafResponseEnable
func (a *ProductNgwafAPIService) EnableProductNgwafExecute(r APIEnableProductNgwafRequest) (*NgwafResponseEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NgwafResponseEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductNgwafAPIService.EnableProductNgwaf")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ngwaf/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

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
	localVarPostBody = r.ngwafRequestEnable
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

// APIGetProductNgwafRequest represents a request for the resource.
type APIGetProductNgwafRequest struct {
	ctx        context.Context
	APIService ProductNgwafAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductNgwafRequest) Execute() (*NgwafResponseEnable, *http.Response, error) {
	return r.APIService.GetProductNgwafExecute(r)
}

/*
GetProductNgwaf Get product enablement status

Get the enablement status of the Next-Gen WAF product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIGetProductNgwafRequest
*/
func (a *ProductNgwafAPIService) GetProductNgwaf(ctx context.Context, serviceId string) APIGetProductNgwafRequest {
	return APIGetProductNgwafRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// GetProductNgwafExecute executes the request
//  @return NgwafResponseEnable
func (a *ProductNgwafAPIService) GetProductNgwafExecute(r APIGetProductNgwafRequest) (*NgwafResponseEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NgwafResponseEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductNgwafAPIService.GetProductNgwaf")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ngwaf/services/{service_id}"
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

// APIGetProductNgwafConfigurationRequest represents a request for the resource.
type APIGetProductNgwafConfigurationRequest struct {
	ctx        context.Context
	APIService ProductNgwafAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductNgwafConfigurationRequest) Execute() (*NgwafResponseConfigure, *http.Response, error) {
	return r.APIService.GetProductNgwafConfigurationExecute(r)
}

/*
GetProductNgwafConfiguration Get configuration

Get configuration of the Next-Gen WAF product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIGetProductNgwafConfigurationRequest
*/
func (a *ProductNgwafAPIService) GetProductNgwafConfiguration(ctx context.Context, serviceId string) APIGetProductNgwafConfigurationRequest {
	return APIGetProductNgwafConfigurationRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// GetProductNgwafConfigurationExecute executes the request
//  @return NgwafResponseConfigure
func (a *ProductNgwafAPIService) GetProductNgwafConfigurationExecute(r APIGetProductNgwafConfigurationRequest) (*NgwafResponseConfigure, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NgwafResponseConfigure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductNgwafAPIService.GetProductNgwafConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ngwaf/services/{service_id}/configuration"
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

// APIGetServicesProductNgwafRequest represents a request for the resource.
type APIGetServicesProductNgwafRequest struct {
	ctx        context.Context
	APIService ProductNgwafAPI
}

// Execute calls the API using the request data configured.
func (r APIGetServicesProductNgwafRequest) Execute() (*NgwafResponseBodyGetAllServices, *http.Response, error) {
	return r.APIService.GetServicesProductNgwafExecute(r)
}

/*
GetServicesProductNgwaf Get services with product enabled

Get all the services which have the Next-Gen WAF product enabled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetServicesProductNgwafRequest
*/
func (a *ProductNgwafAPIService) GetServicesProductNgwaf(ctx context.Context) APIGetServicesProductNgwafRequest {
	return APIGetServicesProductNgwafRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetServicesProductNgwafExecute executes the request
//  @return NgwafResponseBodyGetAllServices
func (a *ProductNgwafAPIService) GetServicesProductNgwafExecute(r APIGetServicesProductNgwafRequest) (*NgwafResponseBodyGetAllServices, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NgwafResponseBodyGetAllServices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductNgwafAPIService.GetServicesProductNgwaf")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ngwaf/services"

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

// APISetProductNgwafConfigurationRequest represents a request for the resource.
type APISetProductNgwafConfigurationRequest struct {
	ctx                             context.Context
	APIService                      ProductNgwafAPI
	serviceId                       string
	ngwafRequestUpdateConfiguration *NgwafRequestUpdateConfiguration
}

// NgwafRequestUpdateConfiguration returns a pointer to a request.
func (r *APISetProductNgwafConfigurationRequest) NgwafRequestUpdateConfiguration(ngwafRequestUpdateConfiguration NgwafRequestUpdateConfiguration) *APISetProductNgwafConfigurationRequest {
	r.ngwafRequestUpdateConfiguration = &ngwafRequestUpdateConfiguration
	return r
}

// Execute calls the API using the request data configured.
func (r APISetProductNgwafConfigurationRequest) Execute() (*NgwafResponseConfigure, *http.Response, error) {
	return r.APIService.SetProductNgwafConfigurationExecute(r)
}

/*
SetProductNgwafConfiguration Update configuration

Update configuration of the Next-Gen WAF product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APISetProductNgwafConfigurationRequest
*/
func (a *ProductNgwafAPIService) SetProductNgwafConfiguration(ctx context.Context, serviceId string) APISetProductNgwafConfigurationRequest {
	return APISetProductNgwafConfigurationRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// SetProductNgwafConfigurationExecute executes the request
//  @return NgwafResponseConfigure
func (a *ProductNgwafAPIService) SetProductNgwafConfigurationExecute(r APISetProductNgwafConfigurationRequest) (*NgwafResponseConfigure, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NgwafResponseConfigure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductNgwafAPIService.SetProductNgwafConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ngwaf/services/{service_id}/configuration"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

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
	localVarPostBody = r.ngwafRequestUpdateConfiguration
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
