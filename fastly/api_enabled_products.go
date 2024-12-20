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

// EnabledProductsAPI defines an interface for interacting with the resource.
type EnabledProductsAPI interface {

	/*
		DisableProduct Disable a product

		Disable a product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, `websockets`, `bot_management`, `ngwaf`, `ddos_protection`, and `log_explorer_insights`.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param productID
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIDisableProductRequest
	*/
	DisableProduct(ctx context.Context, productID string, serviceID string) APIDisableProductRequest

	// DisableProductExecute executes the request
	DisableProductExecute(r APIDisableProductRequest) (*http.Response, error)

	/*
		EnableProduct Enable a product

		Enable a product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, `websockets`, `bot_management`, `ngwaf`, `ddos_protection`, and `log_explorer_insights`.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param productID
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIEnableProductRequest
	*/
	EnableProduct(ctx context.Context, productID string, serviceID string) APIEnableProductRequest

	// EnableProductExecute executes the request
	//  @return EnabledProductResponse
	EnableProductExecute(r APIEnableProductRequest) (*EnabledProductResponse, *http.Response, error)

	/*
		GetEnabledProduct Get enabled product

		Get enabled product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, `websockets`, `bot_management`, `ngwaf`, `ddos_protection`, and `log_explorer_insights`.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param productID
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetEnabledProductRequest
	*/
	GetEnabledProduct(ctx context.Context, productID string, serviceID string) APIGetEnabledProductRequest

	// GetEnabledProductExecute executes the request
	//  @return EnabledProductResponse
	GetEnabledProductExecute(r APIGetEnabledProductRequest) (*EnabledProductResponse, *http.Response, error)

	/*
		GetProductConfiguration Get configuration for a product

		Get configuration for an enabled product on a service. Supported product IDs: `ngwaf` and `ddos_protection`.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param productID
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetProductConfigurationRequest
	*/
	GetProductConfiguration(ctx context.Context, productID string, serviceID string) APIGetProductConfigurationRequest

	// GetProductConfigurationExecute executes the request
	//  @return ConfiguredProductResponse
	GetProductConfigurationExecute(r APIGetProductConfigurationRequest) (*ConfiguredProductResponse, *http.Response, error)

	/*
		SetProductConfiguration Update configuration for a product

		Update configuration for an enabled product on a service. Supported product IDs: `ngwaf` and `ddos_protection`.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param productID
		 @param serviceID Alphanumeric string identifying the service.
		 @return APISetProductConfigurationRequest
	*/
	SetProductConfiguration(ctx context.Context, productID string, serviceID string) APISetProductConfigurationRequest

	// SetProductConfigurationExecute executes the request
	//  @return ConfiguredProductResponse
	SetProductConfigurationExecute(r APISetProductConfigurationRequest) (*ConfiguredProductResponse, *http.Response, error)
}

// EnabledProductsAPIService EnabledProductsAPI service
type EnabledProductsAPIService service

// APIDisableProductRequest represents a request for the resource.
type APIDisableProductRequest struct {
	ctx        context.Context
	APIService EnabledProductsAPI
	productID  string
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIDisableProductRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductExecute(r)
}

/*
DisableProduct Disable a product

Disable a product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, `websockets`, `bot_management`, `ngwaf`, `ddos_protection`, and `log_explorer_insights`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productID
 @param serviceID Alphanumeric string identifying the service.
 @return APIDisableProductRequest
*/
func (a *EnabledProductsAPIService) DisableProduct(ctx context.Context, productID string, serviceID string) APIDisableProductRequest {
	return APIDisableProductRequest{
		APIService: a,
		ctx:        ctx,
		productID:  productID,
		serviceID:  serviceID,
	}
}

// DisableProductExecute executes the request
func (a *EnabledProductsAPIService) DisableProductExecute(r APIDisableProductRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnabledProductsAPIService.DisableProduct")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/{product_id}/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"product_id"+"}", gourl.PathEscape(parameterToString(r.productID, "")))
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

// APIEnableProductRequest represents a request for the resource.
type APIEnableProductRequest struct {
	ctx            context.Context
	APIService     EnabledProductsAPI
	productID      string
	serviceID      string
	setWorkspaceID *SetWorkspaceID
}

// SetWorkspaceID returns a pointer to a request.
func (r *APIEnableProductRequest) SetWorkspaceID(setWorkspaceID SetWorkspaceID) *APIEnableProductRequest {
	r.setWorkspaceID = &setWorkspaceID
	return r
}

// Execute calls the API using the request data configured.
func (r APIEnableProductRequest) Execute() (*EnabledProductResponse, *http.Response, error) {
	return r.APIService.EnableProductExecute(r)
}

/*
EnableProduct Enable a product

Enable a product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, `websockets`, `bot_management`, `ngwaf`, `ddos_protection`, and `log_explorer_insights`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productID
 @param serviceID Alphanumeric string identifying the service.
 @return APIEnableProductRequest
*/
func (a *EnabledProductsAPIService) EnableProduct(ctx context.Context, productID string, serviceID string) APIEnableProductRequest {
	return APIEnableProductRequest{
		APIService: a,
		ctx:        ctx,
		productID:  productID,
		serviceID:  serviceID,
	}
}

// EnableProductExecute executes the request
//  @return EnabledProductResponse
func (a *EnabledProductsAPIService) EnableProductExecute(r APIEnableProductRequest) (*EnabledProductResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EnabledProductResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnabledProductsAPIService.EnableProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/{product_id}/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"product_id"+"}", gourl.PathEscape(parameterToString(r.productID, "")))
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
	localVarPostBody = r.setWorkspaceID
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

// APIGetEnabledProductRequest represents a request for the resource.
type APIGetEnabledProductRequest struct {
	ctx        context.Context
	APIService EnabledProductsAPI
	productID  string
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetEnabledProductRequest) Execute() (*EnabledProductResponse, *http.Response, error) {
	return r.APIService.GetEnabledProductExecute(r)
}

/*
GetEnabledProduct Get enabled product

Get enabled product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, `websockets`, `bot_management`, `ngwaf`, `ddos_protection`, and `log_explorer_insights`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productID
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetEnabledProductRequest
*/
func (a *EnabledProductsAPIService) GetEnabledProduct(ctx context.Context, productID string, serviceID string) APIGetEnabledProductRequest {
	return APIGetEnabledProductRequest{
		APIService: a,
		ctx:        ctx,
		productID:  productID,
		serviceID:  serviceID,
	}
}

// GetEnabledProductExecute executes the request
//  @return EnabledProductResponse
func (a *EnabledProductsAPIService) GetEnabledProductExecute(r APIGetEnabledProductRequest) (*EnabledProductResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EnabledProductResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnabledProductsAPIService.GetEnabledProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/{product_id}/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"product_id"+"}", gourl.PathEscape(parameterToString(r.productID, "")))
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

// APIGetProductConfigurationRequest represents a request for the resource.
type APIGetProductConfigurationRequest struct {
	ctx        context.Context
	APIService EnabledProductsAPI
	productID  string
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductConfigurationRequest) Execute() (*ConfiguredProductResponse, *http.Response, error) {
	return r.APIService.GetProductConfigurationExecute(r)
}

/*
GetProductConfiguration Get configuration for a product

Get configuration for an enabled product on a service. Supported product IDs: `ngwaf` and `ddos_protection`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productID
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetProductConfigurationRequest
*/
func (a *EnabledProductsAPIService) GetProductConfiguration(ctx context.Context, productID string, serviceID string) APIGetProductConfigurationRequest {
	return APIGetProductConfigurationRequest{
		APIService: a,
		ctx:        ctx,
		productID:  productID,
		serviceID:  serviceID,
	}
}

// GetProductConfigurationExecute executes the request
//  @return ConfiguredProductResponse
func (a *EnabledProductsAPIService) GetProductConfigurationExecute(r APIGetProductConfigurationRequest) (*ConfiguredProductResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConfiguredProductResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnabledProductsAPIService.GetProductConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/{product_id}/services/{service_id}/configuration"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"product_id"+"}", gourl.PathEscape(parameterToString(r.productID, "")))
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

// APISetProductConfigurationRequest represents a request for the resource.
type APISetProductConfigurationRequest struct {
	ctx              context.Context
	APIService       EnabledProductsAPI
	productID        string
	serviceID        string
	setConfiguration *SetConfiguration
}

// SetConfiguration returns a pointer to a request.
func (r *APISetProductConfigurationRequest) SetConfiguration(setConfiguration SetConfiguration) *APISetProductConfigurationRequest {
	r.setConfiguration = &setConfiguration
	return r
}

// Execute calls the API using the request data configured.
func (r APISetProductConfigurationRequest) Execute() (*ConfiguredProductResponse, *http.Response, error) {
	return r.APIService.SetProductConfigurationExecute(r)
}

/*
SetProductConfiguration Update configuration for a product

Update configuration for an enabled product on a service. Supported product IDs: `ngwaf` and `ddos_protection`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productID
 @param serviceID Alphanumeric string identifying the service.
 @return APISetProductConfigurationRequest
*/
func (a *EnabledProductsAPIService) SetProductConfiguration(ctx context.Context, productID string, serviceID string) APISetProductConfigurationRequest {
	return APISetProductConfigurationRequest{
		APIService: a,
		ctx:        ctx,
		productID:  productID,
		serviceID:  serviceID,
	}
}

// SetProductConfigurationExecute executes the request
//  @return ConfiguredProductResponse
func (a *EnabledProductsAPIService) SetProductConfigurationExecute(r APISetProductConfigurationRequest) (*ConfiguredProductResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConfiguredProductResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnabledProductsAPIService.SetProductConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/{product_id}/services/{service_id}/configuration"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"product_id"+"}", gourl.PathEscape(parameterToString(r.productID, "")))
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
	localVarPostBody = r.setConfiguration
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
