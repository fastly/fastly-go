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

// ProductDdosProtectionAPI defines an interface for interacting with the resource.
type ProductDdosProtectionAPI interface {

	/*
		DisableProductDdosProtection Disable product

		Disable the DDoS Protection product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIDisableProductDdosProtectionRequest
	*/
	DisableProductDdosProtection(ctx context.Context, serviceID string) APIDisableProductDdosProtectionRequest

	// DisableProductDdosProtectionExecute executes the request
	DisableProductDdosProtectionExecute(r APIDisableProductDdosProtectionRequest) (*http.Response, error)

	/*
		EnableProductDdosProtection Enable product

		Enable the DDoS Protection product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIEnableProductDdosProtectionRequest
	*/
	EnableProductDdosProtection(ctx context.Context, serviceID string) APIEnableProductDdosProtectionRequest

	// EnableProductDdosProtectionExecute executes the request
	//  @return DdosProtectionResponseEnable
	EnableProductDdosProtectionExecute(r APIEnableProductDdosProtectionRequest) (*DdosProtectionResponseEnable, *http.Response, error)

	/*
		GetProductDdosProtection Get product enablement status

		Get the enablement status of the DDoS Protection product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetProductDdosProtectionRequest
	*/
	GetProductDdosProtection(ctx context.Context, serviceID string) APIGetProductDdosProtectionRequest

	// GetProductDdosProtectionExecute executes the request
	//  @return DdosProtectionResponseEnable
	GetProductDdosProtectionExecute(r APIGetProductDdosProtectionRequest) (*DdosProtectionResponseEnable, *http.Response, error)

	/*
		GetProductDdosProtectionConfiguration Get configuration

		Get configuration of the DDoS Protection product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetProductDdosProtectionConfigurationRequest
	*/
	GetProductDdosProtectionConfiguration(ctx context.Context, serviceID string) APIGetProductDdosProtectionConfigurationRequest

	// GetProductDdosProtectionConfigurationExecute executes the request
	//  @return DdosProtectionResponseConfigure
	GetProductDdosProtectionConfigurationExecute(r APIGetProductDdosProtectionConfigurationRequest) (*DdosProtectionResponseConfigure, *http.Response, error)

	/*
		SetProductDdosProtectionConfiguration Update configuration

		Update configuration of the DDoS Protection product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APISetProductDdosProtectionConfigurationRequest
	*/
	SetProductDdosProtectionConfiguration(ctx context.Context, serviceID string) APISetProductDdosProtectionConfigurationRequest

	// SetProductDdosProtectionConfigurationExecute executes the request
	//  @return DdosProtectionResponseConfigure
	SetProductDdosProtectionConfigurationExecute(r APISetProductDdosProtectionConfigurationRequest) (*DdosProtectionResponseConfigure, *http.Response, error)
}

// ProductDdosProtectionAPIService ProductDdosProtectionAPI service
type ProductDdosProtectionAPIService service

// APIDisableProductDdosProtectionRequest represents a request for the resource.
type APIDisableProductDdosProtectionRequest struct {
	ctx        context.Context
	APIService ProductDdosProtectionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIDisableProductDdosProtectionRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductDdosProtectionExecute(r)
}

/*
DisableProductDdosProtection Disable product

Disable the DDoS Protection product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIDisableProductDdosProtectionRequest
*/
func (a *ProductDdosProtectionAPIService) DisableProductDdosProtection(ctx context.Context, serviceID string) APIDisableProductDdosProtectionRequest {
	return APIDisableProductDdosProtectionRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// DisableProductDdosProtectionExecute executes the request
func (a *ProductDdosProtectionAPIService) DisableProductDdosProtectionExecute(r APIDisableProductDdosProtectionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductDdosProtectionAPIService.DisableProductDdosProtection")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ddos_protection/services/{service_id}"
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

// APIEnableProductDdosProtectionRequest represents a request for the resource.
type APIEnableProductDdosProtectionRequest struct {
	ctx        context.Context
	APIService ProductDdosProtectionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIEnableProductDdosProtectionRequest) Execute() (*DdosProtectionResponseEnable, *http.Response, error) {
	return r.APIService.EnableProductDdosProtectionExecute(r)
}

/*
EnableProductDdosProtection Enable product

Enable the DDoS Protection product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIEnableProductDdosProtectionRequest
*/
func (a *ProductDdosProtectionAPIService) EnableProductDdosProtection(ctx context.Context, serviceID string) APIEnableProductDdosProtectionRequest {
	return APIEnableProductDdosProtectionRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// EnableProductDdosProtectionExecute executes the request
//  @return DdosProtectionResponseEnable
func (a *ProductDdosProtectionAPIService) EnableProductDdosProtectionExecute(r APIEnableProductDdosProtectionRequest) (*DdosProtectionResponseEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionResponseEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductDdosProtectionAPIService.EnableProductDdosProtection")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ddos_protection/services/{service_id}"
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

// APIGetProductDdosProtectionRequest represents a request for the resource.
type APIGetProductDdosProtectionRequest struct {
	ctx        context.Context
	APIService ProductDdosProtectionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductDdosProtectionRequest) Execute() (*DdosProtectionResponseEnable, *http.Response, error) {
	return r.APIService.GetProductDdosProtectionExecute(r)
}

/*
GetProductDdosProtection Get product enablement status

Get the enablement status of the DDoS Protection product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetProductDdosProtectionRequest
*/
func (a *ProductDdosProtectionAPIService) GetProductDdosProtection(ctx context.Context, serviceID string) APIGetProductDdosProtectionRequest {
	return APIGetProductDdosProtectionRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// GetProductDdosProtectionExecute executes the request
//  @return DdosProtectionResponseEnable
func (a *ProductDdosProtectionAPIService) GetProductDdosProtectionExecute(r APIGetProductDdosProtectionRequest) (*DdosProtectionResponseEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionResponseEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductDdosProtectionAPIService.GetProductDdosProtection")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ddos_protection/services/{service_id}"
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

// APIGetProductDdosProtectionConfigurationRequest represents a request for the resource.
type APIGetProductDdosProtectionConfigurationRequest struct {
	ctx        context.Context
	APIService ProductDdosProtectionAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductDdosProtectionConfigurationRequest) Execute() (*DdosProtectionResponseConfigure, *http.Response, error) {
	return r.APIService.GetProductDdosProtectionConfigurationExecute(r)
}

/*
GetProductDdosProtectionConfiguration Get configuration

Get configuration of the DDoS Protection product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetProductDdosProtectionConfigurationRequest
*/
func (a *ProductDdosProtectionAPIService) GetProductDdosProtectionConfiguration(ctx context.Context, serviceID string) APIGetProductDdosProtectionConfigurationRequest {
	return APIGetProductDdosProtectionConfigurationRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// GetProductDdosProtectionConfigurationExecute executes the request
//  @return DdosProtectionResponseConfigure
func (a *ProductDdosProtectionAPIService) GetProductDdosProtectionConfigurationExecute(r APIGetProductDdosProtectionConfigurationRequest) (*DdosProtectionResponseConfigure, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionResponseConfigure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductDdosProtectionAPIService.GetProductDdosProtectionConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ddos_protection/services/{service_id}/configuration"
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

// APISetProductDdosProtectionConfigurationRequest represents a request for the resource.
type APISetProductDdosProtectionConfigurationRequest struct {
	ctx                                      context.Context
	APIService                               ProductDdosProtectionAPI
	serviceID                                string
	ddosProtectionRequestUpdateConfiguration *DdosProtectionRequestUpdateConfiguration
}

// DdosProtectionRequestUpdateConfiguration returns a pointer to a request.
func (r *APISetProductDdosProtectionConfigurationRequest) DdosProtectionRequestUpdateConfiguration(ddosProtectionRequestUpdateConfiguration DdosProtectionRequestUpdateConfiguration) *APISetProductDdosProtectionConfigurationRequest {
	r.ddosProtectionRequestUpdateConfiguration = &ddosProtectionRequestUpdateConfiguration
	return r
}

// Execute calls the API using the request data configured.
func (r APISetProductDdosProtectionConfigurationRequest) Execute() (*DdosProtectionResponseConfigure, *http.Response, error) {
	return r.APIService.SetProductDdosProtectionConfigurationExecute(r)
}

/*
SetProductDdosProtectionConfiguration Update configuration

Update configuration of the DDoS Protection product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APISetProductDdosProtectionConfigurationRequest
*/
func (a *ProductDdosProtectionAPIService) SetProductDdosProtectionConfiguration(ctx context.Context, serviceID string) APISetProductDdosProtectionConfigurationRequest {
	return APISetProductDdosProtectionConfigurationRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// SetProductDdosProtectionConfigurationExecute executes the request
//  @return DdosProtectionResponseConfigure
func (a *ProductDdosProtectionAPIService) SetProductDdosProtectionConfigurationExecute(r APISetProductDdosProtectionConfigurationRequest) (*DdosProtectionResponseConfigure, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionResponseConfigure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductDdosProtectionAPIService.SetProductDdosProtectionConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/ddos_protection/services/{service_id}/configuration"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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
	localVarPostBody = r.ddosProtectionRequestUpdateConfiguration
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
