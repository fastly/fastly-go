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

// TLSConfigurationsAPI defines an interface for interacting with the resource.
type TLSConfigurationsAPI interface {

	/*
		GetTLSConfig Get a TLS configuration

		Show a TLS configuration.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsConfigurationID Alphanumeric string identifying a TLS configuration.
		 @return APIGetTLSConfigRequest
	*/
	GetTLSConfig(ctx context.Context, tlsConfigurationID string) APIGetTLSConfigRequest

	// GetTLSConfigExecute executes the request
	//  @return TLSConfigurationResponse
	GetTLSConfigExecute(r APIGetTLSConfigRequest) (*TLSConfigurationResponse, *http.Response, error)

	/*
		ListTLSConfigs List TLS configurations

		List all TLS configurations.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTLSConfigsRequest
	*/
	ListTLSConfigs(ctx context.Context) APIListTLSConfigsRequest

	// ListTLSConfigsExecute executes the request
	//  @return TLSConfigurationsResponse
	ListTLSConfigsExecute(r APIListTLSConfigsRequest) (*TLSConfigurationsResponse, *http.Response, error)

	/*
		UpdateTLSConfig Update a TLS configuration

		Update a TLS configuration.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsConfigurationID Alphanumeric string identifying a TLS configuration.
		 @return APIUpdateTLSConfigRequest
	*/
	UpdateTLSConfig(ctx context.Context, tlsConfigurationID string) APIUpdateTLSConfigRequest

	// UpdateTLSConfigExecute executes the request
	//  @return TLSConfigurationResponse
	UpdateTLSConfigExecute(r APIUpdateTLSConfigRequest) (*TLSConfigurationResponse, *http.Response, error)
}

// TLSConfigurationsAPIService TLSConfigurationsAPI service
type TLSConfigurationsAPIService service

// APIGetTLSConfigRequest represents a request for the resource.
type APIGetTLSConfigRequest struct {
	ctx                context.Context
	APIService         TLSConfigurationsAPI
	tlsConfigurationID string
	include            *string
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;dns_records&#x60;.
func (r *APIGetTLSConfigRequest) Include(include string) *APIGetTLSConfigRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetTLSConfigRequest) Execute() (*TLSConfigurationResponse, *http.Response, error) {
	return r.APIService.GetTLSConfigExecute(r)
}

/*
GetTLSConfig Get a TLS configuration

Show a TLS configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsConfigurationID Alphanumeric string identifying a TLS configuration.
 @return APIGetTLSConfigRequest
*/
func (a *TLSConfigurationsAPIService) GetTLSConfig(ctx context.Context, tlsConfigurationID string) APIGetTLSConfigRequest {
	return APIGetTLSConfigRequest{
		APIService:         a,
		ctx:                ctx,
		tlsConfigurationID: tlsConfigurationID,
	}
}

// GetTLSConfigExecute executes the request
//  @return TLSConfigurationResponse
func (a *TLSConfigurationsAPIService) GetTLSConfigExecute(r APIGetTLSConfigRequest) (*TLSConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TLSConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSConfigurationsAPIService.GetTLSConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/configurations/{tls_configuration_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_configuration_id"+"}", gourl.PathEscape(parameterToString(r.tlsConfigurationID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIListTLSConfigsRequest represents a request for the resource.
type APIListTLSConfigsRequest struct {
	ctx        context.Context
	APIService TLSConfigurationsAPI
	filterBulk *string
	include    *string
	pageNumber *int32
	pageSize   *int32
}

// FilterBulk Optionally filters by the bulk attribute.
func (r *APIListTLSConfigsRequest) FilterBulk(filterBulk string) *APIListTLSConfigsRequest {
	r.filterBulk = &filterBulk
	return r
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;dns_records&#x60;.
func (r *APIListTLSConfigsRequest) Include(include string) *APIListTLSConfigsRequest {
	r.include = &include
	return r
}

// PageNumber Current page.
func (r *APIListTLSConfigsRequest) PageNumber(pageNumber int32) *APIListTLSConfigsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTLSConfigsRequest) PageSize(pageSize int32) *APIListTLSConfigsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSConfigsRequest) Execute() (*TLSConfigurationsResponse, *http.Response, error) {
	return r.APIService.ListTLSConfigsExecute(r)
}

/*
ListTLSConfigs List TLS configurations

List all TLS configurations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSConfigsRequest
*/
func (a *TLSConfigurationsAPIService) ListTLSConfigs(ctx context.Context) APIListTLSConfigsRequest {
	return APIListTLSConfigsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTLSConfigsExecute executes the request
//  @return TLSConfigurationsResponse
func (a *TLSConfigurationsAPIService) ListTLSConfigsExecute(r APIListTLSConfigsRequest) (*TLSConfigurationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TLSConfigurationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSConfigurationsAPIService.ListTLSConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/configurations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterBulk != nil {
		localVarQueryParams.Add("filter[bulk]", parameterToString(*r.filterBulk, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIUpdateTLSConfigRequest represents a request for the resource.
type APIUpdateTLSConfigRequest struct {
	ctx                context.Context
	APIService         TLSConfigurationsAPI
	tlsConfigurationID string
	tlsConfiguration   *TLSConfiguration
}

// TLSConfiguration returns a pointer to a request.
func (r *APIUpdateTLSConfigRequest) TLSConfiguration(tlsConfiguration TLSConfiguration) *APIUpdateTLSConfigRequest {
	r.tlsConfiguration = &tlsConfiguration
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateTLSConfigRequest) Execute() (*TLSConfigurationResponse, *http.Response, error) {
	return r.APIService.UpdateTLSConfigExecute(r)
}

/*
UpdateTLSConfig Update a TLS configuration

Update a TLS configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsConfigurationID Alphanumeric string identifying a TLS configuration.
 @return APIUpdateTLSConfigRequest
*/
func (a *TLSConfigurationsAPIService) UpdateTLSConfig(ctx context.Context, tlsConfigurationID string) APIUpdateTLSConfigRequest {
	return APIUpdateTLSConfigRequest{
		APIService:         a,
		ctx:                ctx,
		tlsConfigurationID: tlsConfigurationID,
	}
}

// UpdateTLSConfigExecute executes the request
//  @return TLSConfigurationResponse
func (a *TLSConfigurationsAPIService) UpdateTLSConfigExecute(r APIUpdateTLSConfigRequest) (*TLSConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TLSConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSConfigurationsAPIService.UpdateTLSConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/configurations/{tls_configuration_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_configuration_id"+"}", gourl.PathEscape(parameterToString(r.tlsConfigurationID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tlsConfiguration
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
