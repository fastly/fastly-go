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

// TlsActivationsAPI defines an interface for interacting with the resource.
type TlsActivationsAPI interface {

	/*
		CreateTlsActivation Enable TLS for a domain using a custom certificate

		Enable TLS for a particular TLS domain and certificate combination. These relationships must be specified to create the TLS activation.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateTlsActivationRequest
	*/
	CreateTlsActivation(ctx context.Context) APICreateTlsActivationRequest

	// CreateTlsActivationExecute executes the request
	//  @return TlsActivationResponse
	CreateTlsActivationExecute(r APICreateTlsActivationRequest) (*TlsActivationResponse, *http.Response, error)

	/*
		DeleteTlsActivation Disable TLS on a domain

		Disable TLS on the domain associated with this TLS activation.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsActivationId Alphanumeric string identifying a TLS activation.
		 @return APIDeleteTlsActivationRequest
	*/
	DeleteTlsActivation(ctx context.Context, tlsActivationId string) APIDeleteTlsActivationRequest

	// DeleteTlsActivationExecute executes the request
	DeleteTlsActivationExecute(r APIDeleteTlsActivationRequest) (*http.Response, error)

	/*
		GetTlsActivation Get a TLS activation

		Show a TLS activation.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsActivationId Alphanumeric string identifying a TLS activation.
		 @return APIGetTlsActivationRequest
	*/
	GetTlsActivation(ctx context.Context, tlsActivationId string) APIGetTlsActivationRequest

	// GetTlsActivationExecute executes the request
	//  @return TlsActivationResponse
	GetTlsActivationExecute(r APIGetTlsActivationRequest) (*TlsActivationResponse, *http.Response, error)

	/*
		ListTlsActivations List TLS activations

		List all TLS activations.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTlsActivationsRequest
	*/
	ListTlsActivations(ctx context.Context) APIListTlsActivationsRequest

	// ListTlsActivationsExecute executes the request
	//  @return TlsActivationsResponse
	ListTlsActivationsExecute(r APIListTlsActivationsRequest) (*TlsActivationsResponse, *http.Response, error)

	/*
		UpdateTlsActivation Update a certificate

		Update the certificate used to terminate TLS traffic for the domain associated with this TLS activation.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsActivationId Alphanumeric string identifying a TLS activation.
		 @return APIUpdateTlsActivationRequest
	*/
	UpdateTlsActivation(ctx context.Context, tlsActivationId string) APIUpdateTlsActivationRequest

	// UpdateTlsActivationExecute executes the request
	//  @return TlsActivationResponse
	UpdateTlsActivationExecute(r APIUpdateTlsActivationRequest) (*TlsActivationResponse, *http.Response, error)
}

// TlsActivationsAPIService TlsActivationsAPI service
type TlsActivationsAPIService service

// APICreateTlsActivationRequest represents a request for the resource.
type APICreateTlsActivationRequest struct {
	ctx           context.Context
	APIService    TlsActivationsAPI
	tlsActivation *TlsActivation
}

// TlsActivation returns a pointer to a request.
func (r *APICreateTlsActivationRequest) TlsActivation(tlsActivation TlsActivation) *APICreateTlsActivationRequest {
	r.tlsActivation = &tlsActivation
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTlsActivationRequest) Execute() (*TlsActivationResponse, *http.Response, error) {
	return r.APIService.CreateTlsActivationExecute(r)
}

/*
CreateTlsActivation Enable TLS for a domain using a custom certificate

Enable TLS for a particular TLS domain and certificate combination. These relationships must be specified to create the TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTlsActivationRequest
*/
func (a *TlsActivationsAPIService) CreateTlsActivation(ctx context.Context) APICreateTlsActivationRequest {
	return APICreateTlsActivationRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateTlsActivationExecute executes the request
//  @return TlsActivationResponse
func (a *TlsActivationsAPIService) CreateTlsActivationExecute(r APICreateTlsActivationRequest) (*TlsActivationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsActivationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsActivationsAPIService.CreateTlsActivation")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations"

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
	localVarPostBody = r.tlsActivation
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

// APIDeleteTlsActivationRequest represents a request for the resource.
type APIDeleteTlsActivationRequest struct {
	ctx             context.Context
	APIService      TlsActivationsAPI
	tlsActivationId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteTlsActivationRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTlsActivationExecute(r)
}

/*
DeleteTlsActivation Disable TLS on a domain

Disable TLS on the domain associated with this TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsActivationId Alphanumeric string identifying a TLS activation.
 @return APIDeleteTlsActivationRequest
*/
func (a *TlsActivationsAPIService) DeleteTlsActivation(ctx context.Context, tlsActivationId string) APIDeleteTlsActivationRequest {
	return APIDeleteTlsActivationRequest{
		APIService:      a,
		ctx:             ctx,
		tlsActivationId: tlsActivationId,
	}
}

// DeleteTlsActivationExecute executes the request
func (a *TlsActivationsAPIService) DeleteTlsActivationExecute(r APIDeleteTlsActivationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsActivationsAPIService.DeleteTlsActivation")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations/{tls_activation_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_activation_id"+"}", gourl.PathEscape(parameterToString(r.tlsActivationId, "")))

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

// APIGetTlsActivationRequest represents a request for the resource.
type APIGetTlsActivationRequest struct {
	ctx             context.Context
	APIService      TlsActivationsAPI
	tlsActivationId string
	include         *string
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_certificate&#x60;, &#x60;tls_configuration&#x60;, and &#x60;tls_domain&#x60;.
func (r *APIGetTlsActivationRequest) Include(include string) *APIGetTlsActivationRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetTlsActivationRequest) Execute() (*TlsActivationResponse, *http.Response, error) {
	return r.APIService.GetTlsActivationExecute(r)
}

/*
GetTlsActivation Get a TLS activation

Show a TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsActivationId Alphanumeric string identifying a TLS activation.
 @return APIGetTlsActivationRequest
*/
func (a *TlsActivationsAPIService) GetTlsActivation(ctx context.Context, tlsActivationId string) APIGetTlsActivationRequest {
	return APIGetTlsActivationRequest{
		APIService:      a,
		ctx:             ctx,
		tlsActivationId: tlsActivationId,
	}
}

// GetTlsActivationExecute executes the request
//  @return TlsActivationResponse
func (a *TlsActivationsAPIService) GetTlsActivationExecute(r APIGetTlsActivationRequest) (*TlsActivationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsActivationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsActivationsAPIService.GetTlsActivation")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations/{tls_activation_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_activation_id"+"}", gourl.PathEscape(parameterToString(r.tlsActivationId, "")))

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

// APIListTlsActivationsRequest represents a request for the resource.
type APIListTlsActivationsRequest struct {
	ctx                      context.Context
	APIService               TlsActivationsAPI
	filterTlsCertificateId   *string
	filterTlsConfigurationId *string
	filterTlsDomainId        *string
	include                  *string
	pageNumber               *int32
	pageSize                 *int32
}

// FilterTlsCertificateId Limit the returned activations to a specific certificate.
func (r *APIListTlsActivationsRequest) FilterTlsCertificateId(filterTlsCertificateId string) *APIListTlsActivationsRequest {
	r.filterTlsCertificateId = &filterTlsCertificateId
	return r
}

// FilterTlsConfigurationId Limit the returned activations to a specific TLS configuration.
func (r *APIListTlsActivationsRequest) FilterTlsConfigurationId(filterTlsConfigurationId string) *APIListTlsActivationsRequest {
	r.filterTlsConfigurationId = &filterTlsConfigurationId
	return r
}

// FilterTlsDomainId Limit the returned rules to a specific domain name.
func (r *APIListTlsActivationsRequest) FilterTlsDomainId(filterTlsDomainId string) *APIListTlsActivationsRequest {
	r.filterTlsDomainId = &filterTlsDomainId
	return r
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_certificate&#x60;, &#x60;tls_configuration&#x60;, and &#x60;tls_domain&#x60;.
func (r *APIListTlsActivationsRequest) Include(include string) *APIListTlsActivationsRequest {
	r.include = &include
	return r
}

// PageNumber Current page.
func (r *APIListTlsActivationsRequest) PageNumber(pageNumber int32) *APIListTlsActivationsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTlsActivationsRequest) PageSize(pageSize int32) *APIListTlsActivationsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTlsActivationsRequest) Execute() (*TlsActivationsResponse, *http.Response, error) {
	return r.APIService.ListTlsActivationsExecute(r)
}

/*
ListTlsActivations List TLS activations

List all TLS activations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTlsActivationsRequest
*/
func (a *TlsActivationsAPIService) ListTlsActivations(ctx context.Context) APIListTlsActivationsRequest {
	return APIListTlsActivationsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTlsActivationsExecute executes the request
//  @return TlsActivationsResponse
func (a *TlsActivationsAPIService) ListTlsActivationsExecute(r APIListTlsActivationsRequest) (*TlsActivationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsActivationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsActivationsAPIService.ListTlsActivations")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterTlsCertificateId != nil {
		localVarQueryParams.Add("filter[tls_certificate.id]", parameterToString(*r.filterTlsCertificateId, ""))
	}
	if r.filterTlsConfigurationId != nil {
		localVarQueryParams.Add("filter[tls_configuration.id]", parameterToString(*r.filterTlsConfigurationId, ""))
	}
	if r.filterTlsDomainId != nil {
		localVarQueryParams.Add("filter[tls_domain.id]", parameterToString(*r.filterTlsDomainId, ""))
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

// APIUpdateTlsActivationRequest represents a request for the resource.
type APIUpdateTlsActivationRequest struct {
	ctx             context.Context
	APIService      TlsActivationsAPI
	tlsActivationId string
	tlsActivation   *TlsActivation
}

// TlsActivation returns a pointer to a request.
func (r *APIUpdateTlsActivationRequest) TlsActivation(tlsActivation TlsActivation) *APIUpdateTlsActivationRequest {
	r.tlsActivation = &tlsActivation
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateTlsActivationRequest) Execute() (*TlsActivationResponse, *http.Response, error) {
	return r.APIService.UpdateTlsActivationExecute(r)
}

/*
UpdateTlsActivation Update a certificate

Update the certificate used to terminate TLS traffic for the domain associated with this TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsActivationId Alphanumeric string identifying a TLS activation.
 @return APIUpdateTlsActivationRequest
*/
func (a *TlsActivationsAPIService) UpdateTlsActivation(ctx context.Context, tlsActivationId string) APIUpdateTlsActivationRequest {
	return APIUpdateTlsActivationRequest{
		APIService:      a,
		ctx:             ctx,
		tlsActivationId: tlsActivationId,
	}
}

// UpdateTlsActivationExecute executes the request
//  @return TlsActivationResponse
func (a *TlsActivationsAPIService) UpdateTlsActivationExecute(r APIUpdateTlsActivationRequest) (*TlsActivationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsActivationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsActivationsAPIService.UpdateTlsActivation")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations/{tls_activation_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_activation_id"+"}", gourl.PathEscape(parameterToString(r.tlsActivationId, "")))

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
	localVarPostBody = r.tlsActivation
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
