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

// TlsCertificatesAPI defines an interface for interacting with the resource.
type TlsCertificatesAPI interface {

	/*
		CreateTlsCert Create a TLS certificate

		Create a TLS certificate.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateTlsCertRequest
	*/
	CreateTlsCert(ctx context.Context) APICreateTlsCertRequest

	// CreateTlsCertExecute executes the request
	//  @return map[string]interface{}
	CreateTlsCertExecute(r APICreateTlsCertRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteTlsCert Delete a TLS certificate

		Destroy a TLS certificate. TLS certificates already enabled for a domain cannot be destroyed.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
		 @return APIDeleteTlsCertRequest
	*/
	DeleteTlsCert(ctx context.Context, tlsCertificateId string) APIDeleteTlsCertRequest

	// DeleteTlsCertExecute executes the request
	DeleteTlsCertExecute(r APIDeleteTlsCertRequest) (*http.Response, error)

	/*
		GetTlsCert Get a TLS certificate

		Show a TLS certificate.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
		 @return APIGetTlsCertRequest
	*/
	GetTlsCert(ctx context.Context, tlsCertificateId string) APIGetTlsCertRequest

	// GetTlsCertExecute executes the request
	//  @return TlsCertificateResponse
	GetTlsCertExecute(r APIGetTlsCertRequest) (*TlsCertificateResponse, *http.Response, error)

	/*
		GetTlsCertBlob Get a TLS certificate blob (Limited Availability)

		Retrieve a TLS certificate blob. This feature is part of a [limited availability](https://docs.fastly.com/products/fastly-product-lifecycle#limited-availability) release.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
		 @return APIGetTlsCertBlobRequest
	*/
	GetTlsCertBlob(ctx context.Context, tlsCertificateId string) APIGetTlsCertBlobRequest

	// GetTlsCertBlobExecute executes the request
	//  @return TlsCertificateBlobResponse
	GetTlsCertBlobExecute(r APIGetTlsCertBlobRequest) (*TlsCertificateBlobResponse, *http.Response, error)

	/*
		ListTlsCerts List TLS certificates

		List all TLS certificates.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTlsCertsRequest
	*/
	ListTlsCerts(ctx context.Context) APIListTlsCertsRequest

	// ListTlsCertsExecute executes the request
	//  @return TlsCertificatesResponse
	ListTlsCertsExecute(r APIListTlsCertsRequest) (*TlsCertificatesResponse, *http.Response, error)

	/*
		UpdateTlsCert Update a TLS certificate

		Replace a TLS certificate with a newly reissued TLS certificate, or update a TLS certificate's name. If replacing a TLS certificate, the new TLS certificate must contain all SAN entries as the current TLS certificate. It must either have an exact matching list or contain a superset.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
		 @return APIUpdateTlsCertRequest
	*/
	UpdateTlsCert(ctx context.Context, tlsCertificateId string) APIUpdateTlsCertRequest

	// UpdateTlsCertExecute executes the request
	//  @return TlsCertificateResponse
	UpdateTlsCertExecute(r APIUpdateTlsCertRequest) (*TlsCertificateResponse, *http.Response, error)
}

// TlsCertificatesAPIService TlsCertificatesAPI service
type TlsCertificatesAPIService service

// APICreateTlsCertRequest represents a request for the resource.
type APICreateTlsCertRequest struct {
	ctx            context.Context
	APIService     TlsCertificatesAPI
	tlsCertificate *TlsCertificate
}

// TlsCertificate returns a pointer to a request.
func (r *APICreateTlsCertRequest) TlsCertificate(tlsCertificate TlsCertificate) *APICreateTlsCertRequest {
	r.tlsCertificate = &tlsCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTlsCertRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.APIService.CreateTlsCertExecute(r)
}

/*
CreateTlsCert Create a TLS certificate

Create a TLS certificate.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTlsCertRequest
*/
func (a *TlsCertificatesAPIService) CreateTlsCert(ctx context.Context) APICreateTlsCertRequest {
	return APICreateTlsCertRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateTlsCertExecute executes the request
//  @return map[string]interface{}
func (a *TlsCertificatesAPIService) CreateTlsCertExecute(r APICreateTlsCertRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsCertificatesAPIService.CreateTlsCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates"

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
	localVarPostBody = r.tlsCertificate
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

// APIDeleteTlsCertRequest represents a request for the resource.
type APIDeleteTlsCertRequest struct {
	ctx              context.Context
	APIService       TlsCertificatesAPI
	tlsCertificateId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteTlsCertRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTlsCertExecute(r)
}

/*
DeleteTlsCert Delete a TLS certificate

Destroy a TLS certificate. TLS certificates already enabled for a domain cannot be destroyed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
 @return APIDeleteTlsCertRequest
*/
func (a *TlsCertificatesAPIService) DeleteTlsCert(ctx context.Context, tlsCertificateId string) APIDeleteTlsCertRequest {
	return APIDeleteTlsCertRequest{
		APIService:       a,
		ctx:              ctx,
		tlsCertificateId: tlsCertificateId,
	}
}

// DeleteTlsCertExecute executes the request
func (a *TlsCertificatesAPIService) DeleteTlsCertExecute(r APIDeleteTlsCertRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsCertificatesAPIService.DeleteTlsCert")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateId, "")))

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

// APIGetTlsCertRequest represents a request for the resource.
type APIGetTlsCertRequest struct {
	ctx              context.Context
	APIService       TlsCertificatesAPI
	tlsCertificateId string
}

// Execute calls the API using the request data configured.
func (r APIGetTlsCertRequest) Execute() (*TlsCertificateResponse, *http.Response, error) {
	return r.APIService.GetTlsCertExecute(r)
}

/*
GetTlsCert Get a TLS certificate

Show a TLS certificate.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
 @return APIGetTlsCertRequest
*/
func (a *TlsCertificatesAPIService) GetTlsCert(ctx context.Context, tlsCertificateId string) APIGetTlsCertRequest {
	return APIGetTlsCertRequest{
		APIService:       a,
		ctx:              ctx,
		tlsCertificateId: tlsCertificateId,
	}
}

// GetTlsCertExecute executes the request
//  @return TlsCertificateResponse
func (a *TlsCertificatesAPIService) GetTlsCertExecute(r APIGetTlsCertRequest) (*TlsCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsCertificatesAPIService.GetTlsCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateId, "")))

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

// APIGetTlsCertBlobRequest represents a request for the resource.
type APIGetTlsCertBlobRequest struct {
	ctx              context.Context
	APIService       TlsCertificatesAPI
	tlsCertificateId string
}

// Execute calls the API using the request data configured.
func (r APIGetTlsCertBlobRequest) Execute() (*TlsCertificateBlobResponse, *http.Response, error) {
	return r.APIService.GetTlsCertBlobExecute(r)
}

/*
GetTlsCertBlob Get a TLS certificate blob (Limited Availability)

Retrieve a TLS certificate blob. This feature is part of a [limited availability](https://docs.fastly.com/products/fastly-product-lifecycle#limited-availability) release.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
 @return APIGetTlsCertBlobRequest
*/
func (a *TlsCertificatesAPIService) GetTlsCertBlob(ctx context.Context, tlsCertificateId string) APIGetTlsCertBlobRequest {
	return APIGetTlsCertBlobRequest{
		APIService:       a,
		ctx:              ctx,
		tlsCertificateId: tlsCertificateId,
	}
}

// GetTlsCertBlobExecute executes the request
//  @return TlsCertificateBlobResponse
func (a *TlsCertificatesAPIService) GetTlsCertBlobExecute(r APIGetTlsCertBlobRequest) (*TlsCertificateBlobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsCertificateBlobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsCertificatesAPIService.GetTlsCertBlob")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}/blob"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateId, "")))

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

// APIListTlsCertsRequest represents a request for the resource.
type APIListTlsCertsRequest struct {
	ctx                context.Context
	APIService         TlsCertificatesAPI
	filterInUse        *string
	filterNotAfter     *string
	filterTlsDomainsId *string
	include            *string
	sort               *string
	pageNumber         *int32
	pageSize           *int32
}

// FilterInUse Optional. Limit the returned certificates to those currently using Fastly to terminate TLS (that is, certificates associated with an activation). Permitted values: true, false.
func (r *APIListTlsCertsRequest) FilterInUse(filterInUse string) *APIListTlsCertsRequest {
	r.filterInUse = &filterInUse
	return r
}

// FilterNotAfter Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]&#x3D;2020-05-05).
func (r *APIListTlsCertsRequest) FilterNotAfter(filterNotAfter string) *APIListTlsCertsRequest {
	r.filterNotAfter = &filterNotAfter
	return r
}

// FilterTlsDomainsId Limit the returned certificates to those that include the specific domain.
func (r *APIListTlsCertsRequest) FilterTlsDomainsId(filterTlsDomainsId string) *APIListTlsCertsRequest {
	r.filterTlsDomainsId = &filterTlsDomainsId
	return r
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_activations&#x60;.
func (r *APIListTlsCertsRequest) Include(include string) *APIListTlsCertsRequest {
	r.include = &include
	return r
}

// Sort The order in which to list the results.
func (r *APIListTlsCertsRequest) Sort(sort string) *APIListTlsCertsRequest {
	r.sort = &sort
	return r
}

// PageNumber Current page.
func (r *APIListTlsCertsRequest) PageNumber(pageNumber int32) *APIListTlsCertsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTlsCertsRequest) PageSize(pageSize int32) *APIListTlsCertsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTlsCertsRequest) Execute() (*TlsCertificatesResponse, *http.Response, error) {
	return r.APIService.ListTlsCertsExecute(r)
}

/*
ListTlsCerts List TLS certificates

List all TLS certificates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTlsCertsRequest
*/
func (a *TlsCertificatesAPIService) ListTlsCerts(ctx context.Context) APIListTlsCertsRequest {
	return APIListTlsCertsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTlsCertsExecute executes the request
//  @return TlsCertificatesResponse
func (a *TlsCertificatesAPIService) ListTlsCertsExecute(r APIListTlsCertsRequest) (*TlsCertificatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsCertificatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsCertificatesAPIService.ListTlsCerts")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterInUse != nil {
		localVarQueryParams.Add("filter[in_use]", parameterToString(*r.filterInUse, ""))
	}
	if r.filterNotAfter != nil {
		localVarQueryParams.Add("filter[not_after]", parameterToString(*r.filterNotAfter, ""))
	}
	if r.filterTlsDomainsId != nil {
		localVarQueryParams.Add("filter[tls_domains.id]", parameterToString(*r.filterTlsDomainsId, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
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

// APIUpdateTlsCertRequest represents a request for the resource.
type APIUpdateTlsCertRequest struct {
	ctx              context.Context
	APIService       TlsCertificatesAPI
	tlsCertificateId string
	tlsCertificate   *TlsCertificate
}

// TlsCertificate returns a pointer to a request.
func (r *APIUpdateTlsCertRequest) TlsCertificate(tlsCertificate TlsCertificate) *APIUpdateTlsCertRequest {
	r.tlsCertificate = &tlsCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateTlsCertRequest) Execute() (*TlsCertificateResponse, *http.Response, error) {
	return r.APIService.UpdateTlsCertExecute(r)
}

/*
UpdateTlsCert Update a TLS certificate

Replace a TLS certificate with a newly reissued TLS certificate, or update a TLS certificate's name. If replacing a TLS certificate, the new TLS certificate must contain all SAN entries as the current TLS certificate. It must either have an exact matching list or contain a superset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateId Alphanumeric string identifying a TLS certificate.
 @return APIUpdateTlsCertRequest
*/
func (a *TlsCertificatesAPIService) UpdateTlsCert(ctx context.Context, tlsCertificateId string) APIUpdateTlsCertRequest {
	return APIUpdateTlsCertRequest{
		APIService:       a,
		ctx:              ctx,
		tlsCertificateId: tlsCertificateId,
	}
}

// UpdateTlsCertExecute executes the request
//  @return TlsCertificateResponse
func (a *TlsCertificatesAPIService) UpdateTlsCertExecute(r APIUpdateTlsCertRequest) (*TlsCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsCertificatesAPIService.UpdateTlsCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateId, "")))

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
	localVarPostBody = r.tlsCertificate
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
