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

// TLSCertificatesAPI defines an interface for interacting with the resource.
type TLSCertificatesAPI interface {

	/*
	CreateTLSCert Create a TLS certificate

	Create a TLS certificate.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateTLSCertRequest
	*/
	CreateTLSCert(ctx context.Context) APICreateTLSCertRequest

	// CreateTLSCertExecute executes the request
	//  @return map[string]any
	CreateTLSCertExecute(r APICreateTLSCertRequest) (map[string]any, *http.Response, error)

	/*
	DeleteTLSCert Delete a TLS certificate

	Destroy a TLS certificate. TLS certificates already enabled for a domain cannot be destroyed.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
	 @return APIDeleteTLSCertRequest
	*/
	DeleteTLSCert(ctx context.Context, tlsCertificateID string) APIDeleteTLSCertRequest

	// DeleteTLSCertExecute executes the request
	DeleteTLSCertExecute(r APIDeleteTLSCertRequest) (*http.Response, error)

	/*
	GetTLSCert Get a TLS certificate

	Show a TLS certificate.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
	 @return APIGetTLSCertRequest
	*/
	GetTLSCert(ctx context.Context, tlsCertificateID string) APIGetTLSCertRequest

	// GetTLSCertExecute executes the request
	//  @return TLSCertificateResponse
	GetTLSCertExecute(r APIGetTLSCertRequest) (*TLSCertificateResponse, *http.Response, error)

	/*
	GetTLSCertBlob Get a TLS certificate blob (Limited Availability)

	Retrieve a TLS certificate blob. This feature is part of a [limited availability](https://docs.fastly.com/products/fastly-product-lifecycle#limited-availability) release.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
	 @return APIGetTLSCertBlobRequest
	*/
	GetTLSCertBlob(ctx context.Context, tlsCertificateID string) APIGetTLSCertBlobRequest

	// GetTLSCertBlobExecute executes the request
	//  @return TLSCertificateBlobResponse
	GetTLSCertBlobExecute(r APIGetTLSCertBlobRequest) (*TLSCertificateBlobResponse, *http.Response, error)

	/*
	ListTLSCerts List TLS certificates

	List all TLS certificates.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListTLSCertsRequest
	*/
	ListTLSCerts(ctx context.Context) APIListTLSCertsRequest

	// ListTLSCertsExecute executes the request
	//  @return TLSCertificatesResponse
	ListTLSCertsExecute(r APIListTLSCertsRequest) (*TLSCertificatesResponse, *http.Response, error)

	/*
	UpdateTLSCert Update a TLS certificate

	Replace a TLS certificate with a newly reissued TLS certificate, or update a TLS certificate's name. If replacing a TLS certificate, the new TLS certificate must contain all SAN entries as the current TLS certificate. It must either have an exact matching list or contain a superset.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
	 @return APIUpdateTLSCertRequest
	*/
	UpdateTLSCert(ctx context.Context, tlsCertificateID string) APIUpdateTLSCertRequest

	// UpdateTLSCertExecute executes the request
	//  @return TLSCertificateResponse
	UpdateTLSCertExecute(r APIUpdateTLSCertRequest) (*TLSCertificateResponse, *http.Response, error)
}

// TLSCertificatesAPIService TLSCertificatesAPI service
type TLSCertificatesAPIService service

// APICreateTLSCertRequest represents a request for the resource.
type APICreateTLSCertRequest struct {
	ctx context.Context
	APIService TLSCertificatesAPI
	tlsCertificate *TLSCertificate
}

// TLSCertificate returns a pointer to a request.
func (r *APICreateTLSCertRequest) TLSCertificate(tlsCertificate TLSCertificate) *APICreateTLSCertRequest {
	r.tlsCertificate = &tlsCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTLSCertRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.CreateTLSCertExecute(r)
}

/*
CreateTLSCert Create a TLS certificate

Create a TLS certificate.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTLSCertRequest
*/
func (a *TLSCertificatesAPIService) CreateTLSCert(ctx context.Context) APICreateTLSCertRequest {
	return APICreateTLSCertRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateTLSCertExecute executes the request
//  @return map[string]any
func (a *TLSCertificatesAPIService) CreateTLSCertExecute(r APICreateTLSCertRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSCertificatesAPIService.CreateTLSCert")
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

// APIDeleteTLSCertRequest represents a request for the resource.
type APIDeleteTLSCertRequest struct {
	ctx context.Context
	APIService TLSCertificatesAPI
	tlsCertificateID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteTLSCertRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTLSCertExecute(r)
}

/*
DeleteTLSCert Delete a TLS certificate

Destroy a TLS certificate. TLS certificates already enabled for a domain cannot be destroyed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
 @return APIDeleteTLSCertRequest
*/
func (a *TLSCertificatesAPIService) DeleteTLSCert(ctx context.Context, tlsCertificateID string) APIDeleteTLSCertRequest {
	return APIDeleteTLSCertRequest{
		APIService: a,
		ctx: ctx,
		tlsCertificateID: tlsCertificateID,
	}
}

// DeleteTLSCertExecute executes the request
func (a *TLSCertificatesAPIService) DeleteTLSCertExecute(r APIDeleteTLSCertRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSCertificatesAPIService.DeleteTLSCert")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateID, "")))

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

// APIGetTLSCertRequest represents a request for the resource.
type APIGetTLSCertRequest struct {
	ctx context.Context
	APIService TLSCertificatesAPI
	tlsCertificateID string
}


// Execute calls the API using the request data configured.
func (r APIGetTLSCertRequest) Execute() (*TLSCertificateResponse, *http.Response, error) {
	return r.APIService.GetTLSCertExecute(r)
}

/*
GetTLSCert Get a TLS certificate

Show a TLS certificate.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
 @return APIGetTLSCertRequest
*/
func (a *TLSCertificatesAPIService) GetTLSCert(ctx context.Context, tlsCertificateID string) APIGetTLSCertRequest {
	return APIGetTLSCertRequest{
		APIService: a,
		ctx: ctx,
		tlsCertificateID: tlsCertificateID,
	}
}

// GetTLSCertExecute executes the request
//  @return TLSCertificateResponse
func (a *TLSCertificatesAPIService) GetTLSCertExecute(r APIGetTLSCertRequest) (*TLSCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSCertificatesAPIService.GetTLSCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateID, "")))

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

// APIGetTLSCertBlobRequest represents a request for the resource.
type APIGetTLSCertBlobRequest struct {
	ctx context.Context
	APIService TLSCertificatesAPI
	tlsCertificateID string
}


// Execute calls the API using the request data configured.
func (r APIGetTLSCertBlobRequest) Execute() (*TLSCertificateBlobResponse, *http.Response, error) {
	return r.APIService.GetTLSCertBlobExecute(r)
}

/*
GetTLSCertBlob Get a TLS certificate blob (Limited Availability)

Retrieve a TLS certificate blob. This feature is part of a [limited availability](https://docs.fastly.com/products/fastly-product-lifecycle#limited-availability) release.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
 @return APIGetTLSCertBlobRequest
*/
func (a *TLSCertificatesAPIService) GetTLSCertBlob(ctx context.Context, tlsCertificateID string) APIGetTLSCertBlobRequest {
	return APIGetTLSCertBlobRequest{
		APIService: a,
		ctx: ctx,
		tlsCertificateID: tlsCertificateID,
	}
}

// GetTLSCertBlobExecute executes the request
//  @return TLSCertificateBlobResponse
func (a *TLSCertificatesAPIService) GetTLSCertBlobExecute(r APIGetTLSCertBlobRequest) (*TLSCertificateBlobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSCertificateBlobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSCertificatesAPIService.GetTLSCertBlob")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}/blob"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateID, "")))

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

// APIListTLSCertsRequest represents a request for the resource.
type APIListTLSCertsRequest struct {
	ctx context.Context
	APIService TLSCertificatesAPI
	filterInUse *string
	filterNotAfter *string
	filterTLSDomainsID *string
	include *string
	sort *string
	pageNumber *int32
	pageSize *int32
}

// FilterInUse Optional. Limit the returned certificates to those currently using Fastly to terminate TLS (that is, certificates associated with an activation). Permitted values: true, false.
func (r *APIListTLSCertsRequest) FilterInUse(filterInUse string) *APIListTLSCertsRequest {
	r.filterInUse = &filterInUse
	return r
}
// FilterNotAfter Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]&#x3D;2020-05-05). 
func (r *APIListTLSCertsRequest) FilterNotAfter(filterNotAfter string) *APIListTLSCertsRequest {
	r.filterNotAfter = &filterNotAfter
	return r
}
// FilterTLSDomainsID Limit the returned certificates to those that include the specific domain.
func (r *APIListTLSCertsRequest) FilterTLSDomainsID(filterTLSDomainsID string) *APIListTLSCertsRequest {
	r.filterTLSDomainsID = &filterTLSDomainsID
	return r
}
// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_activations&#x60;. 
func (r *APIListTLSCertsRequest) Include(include string) *APIListTLSCertsRequest {
	r.include = &include
	return r
}
// Sort The order in which to list the results.
func (r *APIListTLSCertsRequest) Sort(sort string) *APIListTLSCertsRequest {
	r.sort = &sort
	return r
}
// PageNumber Current page.
func (r *APIListTLSCertsRequest) PageNumber(pageNumber int32) *APIListTLSCertsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListTLSCertsRequest) PageSize(pageSize int32) *APIListTLSCertsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSCertsRequest) Execute() (*TLSCertificatesResponse, *http.Response, error) {
	return r.APIService.ListTLSCertsExecute(r)
}

/*
ListTLSCerts List TLS certificates

List all TLS certificates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSCertsRequest
*/
func (a *TLSCertificatesAPIService) ListTLSCerts(ctx context.Context) APIListTLSCertsRequest {
	return APIListTLSCertsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListTLSCertsExecute executes the request
//  @return TLSCertificatesResponse
func (a *TLSCertificatesAPIService) ListTLSCertsExecute(r APIListTLSCertsRequest) (*TLSCertificatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSCertificatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSCertificatesAPIService.ListTLSCerts")
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
	if r.filterTLSDomainsID != nil {
		localVarQueryParams.Add("filter[tls_domains.id]", parameterToString(*r.filterTLSDomainsID, ""))
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

// APIUpdateTLSCertRequest represents a request for the resource.
type APIUpdateTLSCertRequest struct {
	ctx context.Context
	APIService TLSCertificatesAPI
	tlsCertificateID string
	tlsCertificate *TLSCertificate
}

// TLSCertificate returns a pointer to a request.
func (r *APIUpdateTLSCertRequest) TLSCertificate(tlsCertificate TLSCertificate) *APIUpdateTLSCertRequest {
	r.tlsCertificate = &tlsCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateTLSCertRequest) Execute() (*TLSCertificateResponse, *http.Response, error) {
	return r.APIService.UpdateTLSCertExecute(r)
}

/*
UpdateTLSCert Update a TLS certificate

Replace a TLS certificate with a newly reissued TLS certificate, or update a TLS certificate's name. If replacing a TLS certificate, the new TLS certificate must contain all SAN entries as the current TLS certificate. It must either have an exact matching list or contain a superset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsCertificateID Alphanumeric string identifying a TLS certificate.
 @return APIUpdateTLSCertRequest
*/
func (a *TLSCertificatesAPIService) UpdateTLSCert(ctx context.Context, tlsCertificateID string) APIUpdateTLSCertRequest {
	return APIUpdateTLSCertRequest{
		APIService: a,
		ctx: ctx,
		tlsCertificateID: tlsCertificateID,
	}
}

// UpdateTLSCertExecute executes the request
//  @return TLSCertificateResponse
func (a *TLSCertificatesAPIService) UpdateTLSCertExecute(r APIUpdateTLSCertRequest) (*TLSCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSCertificatesAPIService.UpdateTLSCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/certificates/{tls_certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_certificate_id"+"}", gourl.PathEscape(parameterToString(r.tlsCertificateID, "")))

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
