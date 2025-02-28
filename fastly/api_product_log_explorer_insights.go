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

// ProductLogExplorerInsightsAPI defines an interface for interacting with the resource.
type ProductLogExplorerInsightsAPI interface {

	/*
		DisableProductLogExplorerInsights Disable product

		Disable the Log Explorer & Insights product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIDisableProductLogExplorerInsightsRequest
	*/
	DisableProductLogExplorerInsights(ctx context.Context, serviceID string) APIDisableProductLogExplorerInsightsRequest

	// DisableProductLogExplorerInsightsExecute executes the request
	DisableProductLogExplorerInsightsExecute(r APIDisableProductLogExplorerInsightsRequest) (*http.Response, error)

	/*
		EnableProductLogExplorerInsights Enable product

		Enable the Log Explorer & Insights product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIEnableProductLogExplorerInsightsRequest
	*/
	EnableProductLogExplorerInsights(ctx context.Context, serviceID string) APIEnableProductLogExplorerInsightsRequest

	// EnableProductLogExplorerInsightsExecute executes the request
	//  @return LogExplorerInsightsResponseBodyEnable
	EnableProductLogExplorerInsightsExecute(r APIEnableProductLogExplorerInsightsRequest) (*LogExplorerInsightsResponseBodyEnable, *http.Response, error)

	/*
		GetProductLogExplorerInsights Get product enablement status

		Get the enablement status of the Log Explorer & Insights product on a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetProductLogExplorerInsightsRequest
	*/
	GetProductLogExplorerInsights(ctx context.Context, serviceID string) APIGetProductLogExplorerInsightsRequest

	// GetProductLogExplorerInsightsExecute executes the request
	//  @return LogExplorerInsightsResponseBodyEnable
	GetProductLogExplorerInsightsExecute(r APIGetProductLogExplorerInsightsRequest) (*LogExplorerInsightsResponseBodyEnable, *http.Response, error)
}

// ProductLogExplorerInsightsAPIService ProductLogExplorerInsightsAPI service
type ProductLogExplorerInsightsAPIService service

// APIDisableProductLogExplorerInsightsRequest represents a request for the resource.
type APIDisableProductLogExplorerInsightsRequest struct {
	ctx        context.Context
	APIService ProductLogExplorerInsightsAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIDisableProductLogExplorerInsightsRequest) Execute() (*http.Response, error) {
	return r.APIService.DisableProductLogExplorerInsightsExecute(r)
}

/*
DisableProductLogExplorerInsights Disable product

Disable the Log Explorer & Insights product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIDisableProductLogExplorerInsightsRequest
*/
func (a *ProductLogExplorerInsightsAPIService) DisableProductLogExplorerInsights(ctx context.Context, serviceID string) APIDisableProductLogExplorerInsightsRequest {
	return APIDisableProductLogExplorerInsightsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// DisableProductLogExplorerInsightsExecute executes the request
func (a *ProductLogExplorerInsightsAPIService) DisableProductLogExplorerInsightsExecute(r APIDisableProductLogExplorerInsightsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductLogExplorerInsightsAPIService.DisableProductLogExplorerInsights")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/log_explorer_insights/services/{service_id}"
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

// APIEnableProductLogExplorerInsightsRequest represents a request for the resource.
type APIEnableProductLogExplorerInsightsRequest struct {
	ctx        context.Context
	APIService ProductLogExplorerInsightsAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIEnableProductLogExplorerInsightsRequest) Execute() (*LogExplorerInsightsResponseBodyEnable, *http.Response, error) {
	return r.APIService.EnableProductLogExplorerInsightsExecute(r)
}

/*
EnableProductLogExplorerInsights Enable product

Enable the Log Explorer & Insights product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIEnableProductLogExplorerInsightsRequest
*/
func (a *ProductLogExplorerInsightsAPIService) EnableProductLogExplorerInsights(ctx context.Context, serviceID string) APIEnableProductLogExplorerInsightsRequest {
	return APIEnableProductLogExplorerInsightsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// EnableProductLogExplorerInsightsExecute executes the request
//  @return LogExplorerInsightsResponseBodyEnable
func (a *ProductLogExplorerInsightsAPIService) EnableProductLogExplorerInsightsExecute(r APIEnableProductLogExplorerInsightsRequest) (*LogExplorerInsightsResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LogExplorerInsightsResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductLogExplorerInsightsAPIService.EnableProductLogExplorerInsights")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/log_explorer_insights/services/{service_id}"
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

// APIGetProductLogExplorerInsightsRequest represents a request for the resource.
type APIGetProductLogExplorerInsightsRequest struct {
	ctx        context.Context
	APIService ProductLogExplorerInsightsAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetProductLogExplorerInsightsRequest) Execute() (*LogExplorerInsightsResponseBodyEnable, *http.Response, error) {
	return r.APIService.GetProductLogExplorerInsightsExecute(r)
}

/*
GetProductLogExplorerInsights Get product enablement status

Get the enablement status of the Log Explorer & Insights product on a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetProductLogExplorerInsightsRequest
*/
func (a *ProductLogExplorerInsightsAPIService) GetProductLogExplorerInsights(ctx context.Context, serviceID string) APIGetProductLogExplorerInsightsRequest {
	return APIGetProductLogExplorerInsightsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// GetProductLogExplorerInsightsExecute executes the request
//  @return LogExplorerInsightsResponseBodyEnable
func (a *ProductLogExplorerInsightsAPIService) GetProductLogExplorerInsightsExecute(r APIGetProductLogExplorerInsightsRequest) (*LogExplorerInsightsResponseBodyEnable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LogExplorerInsightsResponseBodyEnable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductLogExplorerInsightsAPIService.GetProductLogExplorerInsights")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enabled-products/v1/log_explorer_insights/services/{service_id}"
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
