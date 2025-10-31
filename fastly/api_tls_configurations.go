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

// TlsConfigurationsAPI defines an interface for interacting with the resource.
type TlsConfigurationsAPI interface {

	/*
		GetTlsConfig Get a TLS configuration

		Show a TLS configuration.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsConfigurationId Alphanumeric string identifying a TLS configuration.
		 @return APIGetTlsConfigRequest
	*/
	GetTlsConfig(ctx context.Context, tlsConfigurationId string) APIGetTlsConfigRequest

	// GetTlsConfigExecute executes the request
	//  @return TlsConfigurationResponse
	GetTlsConfigExecute(r APIGetTlsConfigRequest) (*TlsConfigurationResponse, *http.Response, error)

	/*
		ListTlsConfigs List TLS configurations

		List all TLS configurations.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTlsConfigsRequest
	*/
	ListTlsConfigs(ctx context.Context) APIListTlsConfigsRequest

	// ListTlsConfigsExecute executes the request
	//  @return TlsConfigurationsResponse
	ListTlsConfigsExecute(r APIListTlsConfigsRequest) (*TlsConfigurationsResponse, *http.Response, error)

	/*
		UpdateTlsConfig Update a TLS configuration

		Update a TLS configuration.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsConfigurationId Alphanumeric string identifying a TLS configuration.
		 @return APIUpdateTlsConfigRequest
	*/
	UpdateTlsConfig(ctx context.Context, tlsConfigurationId string) APIUpdateTlsConfigRequest

	// UpdateTlsConfigExecute executes the request
	//  @return TlsConfigurationResponse
	UpdateTlsConfigExecute(r APIUpdateTlsConfigRequest) (*TlsConfigurationResponse, *http.Response, error)
}

// TlsConfigurationsAPIService TlsConfigurationsAPI service
type TlsConfigurationsAPIService service

// APIGetTlsConfigRequest represents a request for the resource.
type APIGetTlsConfigRequest struct {
	ctx                context.Context
	APIService         TlsConfigurationsAPI
	tlsConfigurationId string
	include            *string
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;dns_records&#x60;.
func (r *APIGetTlsConfigRequest) Include(include string) *APIGetTlsConfigRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetTlsConfigRequest) Execute() (*TlsConfigurationResponse, *http.Response, error) {
	return r.APIService.GetTlsConfigExecute(r)
}

/*
GetTlsConfig Get a TLS configuration

Show a TLS configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsConfigurationId Alphanumeric string identifying a TLS configuration.
 @return APIGetTlsConfigRequest
*/
func (a *TlsConfigurationsAPIService) GetTlsConfig(ctx context.Context, tlsConfigurationId string) APIGetTlsConfigRequest {
	return APIGetTlsConfigRequest{
		APIService:         a,
		ctx:                ctx,
		tlsConfigurationId: tlsConfigurationId,
	}
}

// GetTlsConfigExecute executes the request
//  @return TlsConfigurationResponse
func (a *TlsConfigurationsAPIService) GetTlsConfigExecute(r APIGetTlsConfigRequest) (*TlsConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsConfigurationsAPIService.GetTlsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/configurations/{tls_configuration_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_configuration_id"+"}", gourl.PathEscape(parameterToString(r.tlsConfigurationId, "")))

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

// APIListTlsConfigsRequest represents a request for the resource.
type APIListTlsConfigsRequest struct {
	ctx        context.Context
	APIService TlsConfigurationsAPI
	filterBulk *string
	include    *string
	pageNumber *int32
	pageSize   *int32
}

// FilterBulk Optionally filters by the bulk attribute.
func (r *APIListTlsConfigsRequest) FilterBulk(filterBulk string) *APIListTlsConfigsRequest {
	r.filterBulk = &filterBulk
	return r
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;dns_records&#x60;.
func (r *APIListTlsConfigsRequest) Include(include string) *APIListTlsConfigsRequest {
	r.include = &include
	return r
}

// PageNumber Current page.
func (r *APIListTlsConfigsRequest) PageNumber(pageNumber int32) *APIListTlsConfigsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTlsConfigsRequest) PageSize(pageSize int32) *APIListTlsConfigsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTlsConfigsRequest) Execute() (*TlsConfigurationsResponse, *http.Response, error) {
	return r.APIService.ListTlsConfigsExecute(r)
}

/*
ListTlsConfigs List TLS configurations

List all TLS configurations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTlsConfigsRequest
*/
func (a *TlsConfigurationsAPIService) ListTlsConfigs(ctx context.Context) APIListTlsConfigsRequest {
	return APIListTlsConfigsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTlsConfigsExecute executes the request
//  @return TlsConfigurationsResponse
func (a *TlsConfigurationsAPIService) ListTlsConfigsExecute(r APIListTlsConfigsRequest) (*TlsConfigurationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsConfigurationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsConfigurationsAPIService.ListTlsConfigs")
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

// APIUpdateTlsConfigRequest represents a request for the resource.
type APIUpdateTlsConfigRequest struct {
	ctx                context.Context
	APIService         TlsConfigurationsAPI
	tlsConfigurationId string
	tlsConfiguration   *TlsConfiguration
}

// TlsConfiguration returns a pointer to a request.
func (r *APIUpdateTlsConfigRequest) TlsConfiguration(tlsConfiguration TlsConfiguration) *APIUpdateTlsConfigRequest {
	r.tlsConfiguration = &tlsConfiguration
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateTlsConfigRequest) Execute() (*TlsConfigurationResponse, *http.Response, error) {
	return r.APIService.UpdateTlsConfigExecute(r)
}

/*
UpdateTlsConfig Update a TLS configuration

Update a TLS configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsConfigurationId Alphanumeric string identifying a TLS configuration.
 @return APIUpdateTlsConfigRequest
*/
func (a *TlsConfigurationsAPIService) UpdateTlsConfig(ctx context.Context, tlsConfigurationId string) APIUpdateTlsConfigRequest {
	return APIUpdateTlsConfigRequest{
		APIService:         a,
		ctx:                ctx,
		tlsConfigurationId: tlsConfigurationId,
	}
}

// UpdateTlsConfigExecute executes the request
//  @return TlsConfigurationResponse
func (a *TlsConfigurationsAPIService) UpdateTlsConfigExecute(r APIUpdateTlsConfigRequest) (*TlsConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsConfigurationsAPIService.UpdateTlsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/configurations/{tls_configuration_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_configuration_id"+"}", gourl.PathEscape(parameterToString(r.tlsConfigurationId, "")))

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
