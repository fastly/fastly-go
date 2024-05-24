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

// TLSActivationsAPI defines an interface for interacting with the resource.
type TLSActivationsAPI interface {

	/*
	CreateTLSActivation Enable TLS for a domain using a custom certificate

	Enable TLS for a particular TLS domain and certificate combination. These relationships must be specified to create the TLS activation.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateTLSActivationRequest
	*/
	CreateTLSActivation(ctx context.Context) APICreateTLSActivationRequest

	// CreateTLSActivationExecute executes the request
	//  @return TLSActivationResponse
	CreateTLSActivationExecute(r APICreateTLSActivationRequest) (*TLSActivationResponse, *http.Response, error)

	/*
	DeleteTLSActivation Disable TLS on a domain

	Disable TLS on the domain associated with this TLS activation.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsActivationID Alphanumeric string identifying a TLS activation.
	 @return APIDeleteTLSActivationRequest
	*/
	DeleteTLSActivation(ctx context.Context, tlsActivationID string) APIDeleteTLSActivationRequest

	// DeleteTLSActivationExecute executes the request
	DeleteTLSActivationExecute(r APIDeleteTLSActivationRequest) (*http.Response, error)

	/*
	GetTLSActivation Get a TLS activation

	Show a TLS activation.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsActivationID Alphanumeric string identifying a TLS activation.
	 @return APIGetTLSActivationRequest
	*/
	GetTLSActivation(ctx context.Context, tlsActivationID string) APIGetTLSActivationRequest

	// GetTLSActivationExecute executes the request
	//  @return TLSActivationResponse
	GetTLSActivationExecute(r APIGetTLSActivationRequest) (*TLSActivationResponse, *http.Response, error)

	/*
	ListTLSActivations List TLS activations

	List all TLS activations.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListTLSActivationsRequest
	*/
	ListTLSActivations(ctx context.Context) APIListTLSActivationsRequest

	// ListTLSActivationsExecute executes the request
	//  @return TLSActivationsResponse
	ListTLSActivationsExecute(r APIListTLSActivationsRequest) (*TLSActivationsResponse, *http.Response, error)

	/*
	UpdateTLSActivation Update a certificate

	Update the certificate used to terminate TLS traffic for the domain associated with this TLS activation.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsActivationID Alphanumeric string identifying a TLS activation.
	 @return APIUpdateTLSActivationRequest
	*/
	UpdateTLSActivation(ctx context.Context, tlsActivationID string) APIUpdateTLSActivationRequest

	// UpdateTLSActivationExecute executes the request
	//  @return TLSActivationResponse
	UpdateTLSActivationExecute(r APIUpdateTLSActivationRequest) (*TLSActivationResponse, *http.Response, error)
}

// TLSActivationsAPIService TLSActivationsAPI service
type TLSActivationsAPIService service

// APICreateTLSActivationRequest represents a request for the resource.
type APICreateTLSActivationRequest struct {
	ctx context.Context
	APIService TLSActivationsAPI
	tlsActivation *TLSActivation
}

// TLSActivation returns a pointer to a request.
func (r *APICreateTLSActivationRequest) TLSActivation(tlsActivation TLSActivation) *APICreateTLSActivationRequest {
	r.tlsActivation = &tlsActivation
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTLSActivationRequest) Execute() (*TLSActivationResponse, *http.Response, error) {
	return r.APIService.CreateTLSActivationExecute(r)
}

/*
CreateTLSActivation Enable TLS for a domain using a custom certificate

Enable TLS for a particular TLS domain and certificate combination. These relationships must be specified to create the TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTLSActivationRequest
*/
func (a *TLSActivationsAPIService) CreateTLSActivation(ctx context.Context) APICreateTLSActivationRequest {
	return APICreateTLSActivationRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateTLSActivationExecute executes the request
//  @return TLSActivationResponse
func (a *TLSActivationsAPIService) CreateTLSActivationExecute(r APICreateTLSActivationRequest) (*TLSActivationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSActivationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSActivationsAPIService.CreateTLSActivation")
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

// APIDeleteTLSActivationRequest represents a request for the resource.
type APIDeleteTLSActivationRequest struct {
	ctx context.Context
	APIService TLSActivationsAPI
	tlsActivationID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteTLSActivationRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTLSActivationExecute(r)
}

/*
DeleteTLSActivation Disable TLS on a domain

Disable TLS on the domain associated with this TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsActivationID Alphanumeric string identifying a TLS activation.
 @return APIDeleteTLSActivationRequest
*/
func (a *TLSActivationsAPIService) DeleteTLSActivation(ctx context.Context, tlsActivationID string) APIDeleteTLSActivationRequest {
	return APIDeleteTLSActivationRequest{
		APIService: a,
		ctx: ctx,
		tlsActivationID: tlsActivationID,
	}
}

// DeleteTLSActivationExecute executes the request
func (a *TLSActivationsAPIService) DeleteTLSActivationExecute(r APIDeleteTLSActivationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSActivationsAPIService.DeleteTLSActivation")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations/{tls_activation_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_activation_id"+"}", gourl.PathEscape(parameterToString(r.tlsActivationID, "")))

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

// APIGetTLSActivationRequest represents a request for the resource.
type APIGetTLSActivationRequest struct {
	ctx context.Context
	APIService TLSActivationsAPI
	tlsActivationID string
	include *string
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_certificate&#x60;, &#x60;tls_configuration&#x60;, and &#x60;tls_domain&#x60;. 
func (r *APIGetTLSActivationRequest) Include(include string) *APIGetTLSActivationRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetTLSActivationRequest) Execute() (*TLSActivationResponse, *http.Response, error) {
	return r.APIService.GetTLSActivationExecute(r)
}

/*
GetTLSActivation Get a TLS activation

Show a TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsActivationID Alphanumeric string identifying a TLS activation.
 @return APIGetTLSActivationRequest
*/
func (a *TLSActivationsAPIService) GetTLSActivation(ctx context.Context, tlsActivationID string) APIGetTLSActivationRequest {
	return APIGetTLSActivationRequest{
		APIService: a,
		ctx: ctx,
		tlsActivationID: tlsActivationID,
	}
}

// GetTLSActivationExecute executes the request
//  @return TLSActivationResponse
func (a *TLSActivationsAPIService) GetTLSActivationExecute(r APIGetTLSActivationRequest) (*TLSActivationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSActivationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSActivationsAPIService.GetTLSActivation")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations/{tls_activation_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_activation_id"+"}", gourl.PathEscape(parameterToString(r.tlsActivationID, "")))

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

// APIListTLSActivationsRequest represents a request for the resource.
type APIListTLSActivationsRequest struct {
	ctx context.Context
	APIService TLSActivationsAPI
	filterTLSCertificateID *string
	filterTLSConfigurationID *string
	filterTLSDomainID *string
	include *string
	pageNumber *int32
	pageSize *int32
}

// FilterTLSCertificateID Limit the returned activations to a specific certificate.
func (r *APIListTLSActivationsRequest) FilterTLSCertificateID(filterTLSCertificateID string) *APIListTLSActivationsRequest {
	r.filterTLSCertificateID = &filterTLSCertificateID
	return r
}
// FilterTLSConfigurationID Limit the returned activations to a specific TLS configuration.
func (r *APIListTLSActivationsRequest) FilterTLSConfigurationID(filterTLSConfigurationID string) *APIListTLSActivationsRequest {
	r.filterTLSConfigurationID = &filterTLSConfigurationID
	return r
}
// FilterTLSDomainID Limit the returned rules to a specific domain name.
func (r *APIListTLSActivationsRequest) FilterTLSDomainID(filterTLSDomainID string) *APIListTLSActivationsRequest {
	r.filterTLSDomainID = &filterTLSDomainID
	return r
}
// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_certificate&#x60;, &#x60;tls_configuration&#x60;, and &#x60;tls_domain&#x60;. 
func (r *APIListTLSActivationsRequest) Include(include string) *APIListTLSActivationsRequest {
	r.include = &include
	return r
}
// PageNumber Current page.
func (r *APIListTLSActivationsRequest) PageNumber(pageNumber int32) *APIListTLSActivationsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListTLSActivationsRequest) PageSize(pageSize int32) *APIListTLSActivationsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSActivationsRequest) Execute() (*TLSActivationsResponse, *http.Response, error) {
	return r.APIService.ListTLSActivationsExecute(r)
}

/*
ListTLSActivations List TLS activations

List all TLS activations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSActivationsRequest
*/
func (a *TLSActivationsAPIService) ListTLSActivations(ctx context.Context) APIListTLSActivationsRequest {
	return APIListTLSActivationsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListTLSActivationsExecute executes the request
//  @return TLSActivationsResponse
func (a *TLSActivationsAPIService) ListTLSActivationsExecute(r APIListTLSActivationsRequest) (*TLSActivationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSActivationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSActivationsAPIService.ListTLSActivations")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterTLSCertificateID != nil {
		localVarQueryParams.Add("filter[tls_certificate.id]", parameterToString(*r.filterTLSCertificateID, ""))
	}
	if r.filterTLSConfigurationID != nil {
		localVarQueryParams.Add("filter[tls_configuration.id]", parameterToString(*r.filterTLSConfigurationID, ""))
	}
	if r.filterTLSDomainID != nil {
		localVarQueryParams.Add("filter[tls_domain.id]", parameterToString(*r.filterTLSDomainID, ""))
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

// APIUpdateTLSActivationRequest represents a request for the resource.
type APIUpdateTLSActivationRequest struct {
	ctx context.Context
	APIService TLSActivationsAPI
	tlsActivationID string
	tlsActivation *TLSActivation
}

// TLSActivation returns a pointer to a request.
func (r *APIUpdateTLSActivationRequest) TLSActivation(tlsActivation TLSActivation) *APIUpdateTLSActivationRequest {
	r.tlsActivation = &tlsActivation
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateTLSActivationRequest) Execute() (*TLSActivationResponse, *http.Response, error) {
	return r.APIService.UpdateTLSActivationExecute(r)
}

/*
UpdateTLSActivation Update a certificate

Update the certificate used to terminate TLS traffic for the domain associated with this TLS activation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsActivationID Alphanumeric string identifying a TLS activation.
 @return APIUpdateTLSActivationRequest
*/
func (a *TLSActivationsAPIService) UpdateTLSActivation(ctx context.Context, tlsActivationID string) APIUpdateTLSActivationRequest {
	return APIUpdateTLSActivationRequest{
		APIService: a,
		ctx: ctx,
		tlsActivationID: tlsActivationID,
	}
}

// UpdateTLSActivationExecute executes the request
//  @return TLSActivationResponse
func (a *TLSActivationsAPIService) UpdateTLSActivationExecute(r APIUpdateTLSActivationRequest) (*TLSActivationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSActivationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSActivationsAPIService.UpdateTLSActivation")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/activations/{tls_activation_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_activation_id"+"}", gourl.PathEscape(parameterToString(r.tlsActivationID, "")))

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
