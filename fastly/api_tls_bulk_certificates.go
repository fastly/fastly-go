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

// TLSBulkCertificatesAPI defines an interface for interacting with the resource.
type TLSBulkCertificatesAPI interface {

	/*
	DeleteBulkTLSCert Delete a certificate

	Destroy a certificate. This disables TLS for all domains listed as SAN entries.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param certificateID Alphanumeric string identifying a TLS bulk certificate.
	 @return APIDeleteBulkTLSCertRequest
	*/
	DeleteBulkTLSCert(ctx context.Context, certificateID string) APIDeleteBulkTLSCertRequest

	// DeleteBulkTLSCertExecute executes the request
	DeleteBulkTLSCertExecute(r APIDeleteBulkTLSCertRequest) (*http.Response, error)

	/*
	GetTLSBulkCert Get a certificate

	Retrieve a single certificate.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param certificateID Alphanumeric string identifying a TLS bulk certificate.
	 @return APIGetTLSBulkCertRequest
	*/
	GetTLSBulkCert(ctx context.Context, certificateID string) APIGetTLSBulkCertRequest

	// GetTLSBulkCertExecute executes the request
	//  @return TLSBulkCertificateResponse
	GetTLSBulkCertExecute(r APIGetTLSBulkCertRequest) (*TLSBulkCertificateResponse, *http.Response, error)

	/*
	ListTLSBulkCerts List certificates

	List all certificates.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListTLSBulkCertsRequest
	*/
	ListTLSBulkCerts(ctx context.Context) APIListTLSBulkCertsRequest

	// ListTLSBulkCertsExecute executes the request
	//  @return TLSBulkCertificatesResponse
	ListTLSBulkCertsExecute(r APIListTLSBulkCertsRequest) (*TLSBulkCertificatesResponse, *http.Response, error)

	/*
	UpdateBulkTLSCert Update a certificate

	Replace a certificate with a newly reissued certificate. By using this endpoint, the original certificate will cease to be used for future TLS handshakes. Thus, only SAN entries that appear in the replacement certificate will become TLS enabled. Any SAN entries that are missing in the replacement certificate will become disabled.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param certificateID Alphanumeric string identifying a TLS bulk certificate.
	 @return APIUpdateBulkTLSCertRequest
	*/
	UpdateBulkTLSCert(ctx context.Context, certificateID string) APIUpdateBulkTLSCertRequest

	// UpdateBulkTLSCertExecute executes the request
	//  @return TLSBulkCertificateResponse
	UpdateBulkTLSCertExecute(r APIUpdateBulkTLSCertRequest) (*TLSBulkCertificateResponse, *http.Response, error)

	/*
	UploadTLSBulkCert Upload a certificate

	Upload a new certificate. TLS domains are automatically enabled upon certificate creation. If a domain is already enabled on a previously uploaded certificate, that domain will be updated to use the new certificate for all future TLS handshake requests.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIUploadTLSBulkCertRequest
	*/
	UploadTLSBulkCert(ctx context.Context) APIUploadTLSBulkCertRequest

	// UploadTLSBulkCertExecute executes the request
	//  @return TLSBulkCertificateResponse
	UploadTLSBulkCertExecute(r APIUploadTLSBulkCertRequest) (*TLSBulkCertificateResponse, *http.Response, error)
}

// TLSBulkCertificatesAPIService TLSBulkCertificatesAPI service
type TLSBulkCertificatesAPIService service

// APIDeleteBulkTLSCertRequest represents a request for the resource.
type APIDeleteBulkTLSCertRequest struct {
	ctx context.Context
	APIService TLSBulkCertificatesAPI
	certificateID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteBulkTLSCertRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteBulkTLSCertExecute(r)
}

/*
DeleteBulkTLSCert Delete a certificate

Destroy a certificate. This disables TLS for all domains listed as SAN entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certificateID Alphanumeric string identifying a TLS bulk certificate.
 @return APIDeleteBulkTLSCertRequest
*/
func (a *TLSBulkCertificatesAPIService) DeleteBulkTLSCert(ctx context.Context, certificateID string) APIDeleteBulkTLSCertRequest {
	return APIDeleteBulkTLSCertRequest{
		APIService: a,
		ctx: ctx,
		certificateID: certificateID,
	}
}

// DeleteBulkTLSCertExecute executes the request
func (a *TLSBulkCertificatesAPIService) DeleteBulkTLSCertExecute(r APIDeleteBulkTLSCertRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSBulkCertificatesAPIService.DeleteBulkTLSCert")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates/{certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"certificate_id"+"}", gourl.PathEscape(parameterToString(r.certificateID, "")))

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

// APIGetTLSBulkCertRequest represents a request for the resource.
type APIGetTLSBulkCertRequest struct {
	ctx context.Context
	APIService TLSBulkCertificatesAPI
	certificateID string
}


// Execute calls the API using the request data configured.
func (r APIGetTLSBulkCertRequest) Execute() (*TLSBulkCertificateResponse, *http.Response, error) {
	return r.APIService.GetTLSBulkCertExecute(r)
}

/*
GetTLSBulkCert Get a certificate

Retrieve a single certificate.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certificateID Alphanumeric string identifying a TLS bulk certificate.
 @return APIGetTLSBulkCertRequest
*/
func (a *TLSBulkCertificatesAPIService) GetTLSBulkCert(ctx context.Context, certificateID string) APIGetTLSBulkCertRequest {
	return APIGetTLSBulkCertRequest{
		APIService: a,
		ctx: ctx,
		certificateID: certificateID,
	}
}

// GetTLSBulkCertExecute executes the request
//  @return TLSBulkCertificateResponse
func (a *TLSBulkCertificatesAPIService) GetTLSBulkCertExecute(r APIGetTLSBulkCertRequest) (*TLSBulkCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSBulkCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSBulkCertificatesAPIService.GetTLSBulkCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates/{certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"certificate_id"+"}", gourl.PathEscape(parameterToString(r.certificateID, "")))

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

// APIListTLSBulkCertsRequest represents a request for the resource.
type APIListTLSBulkCertsRequest struct {
	ctx context.Context
	APIService TLSBulkCertificatesAPI
	filterTLSDomainID *string
	pageNumber *int32
	pageSize *int32
	sort *string
}

// FilterTLSDomainID Filter certificates by their matching, fully-qualified domain name.
func (r *APIListTLSBulkCertsRequest) FilterTLSDomainID(filterTLSDomainID string) *APIListTLSBulkCertsRequest {
	r.filterTLSDomainID = &filterTLSDomainID
	return r
}
// PageNumber Current page.
func (r *APIListTLSBulkCertsRequest) PageNumber(pageNumber int32) *APIListTLSBulkCertsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListTLSBulkCertsRequest) PageSize(pageSize int32) *APIListTLSBulkCertsRequest {
	r.pageSize = &pageSize
	return r
}
// Sort The order in which to list the results by creation date.
func (r *APIListTLSBulkCertsRequest) Sort(sort string) *APIListTLSBulkCertsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSBulkCertsRequest) Execute() (*TLSBulkCertificatesResponse, *http.Response, error) {
	return r.APIService.ListTLSBulkCertsExecute(r)
}

/*
ListTLSBulkCerts List certificates

List all certificates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSBulkCertsRequest
*/
func (a *TLSBulkCertificatesAPIService) ListTLSBulkCerts(ctx context.Context) APIListTLSBulkCertsRequest {
	return APIListTLSBulkCertsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListTLSBulkCertsExecute executes the request
//  @return TLSBulkCertificatesResponse
func (a *TLSBulkCertificatesAPIService) ListTLSBulkCertsExecute(r APIListTLSBulkCertsRequest) (*TLSBulkCertificatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSBulkCertificatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSBulkCertificatesAPIService.ListTLSBulkCerts")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterTLSDomainID != nil {
		localVarQueryParams.Add("filter[tls_domain.id]", parameterToString(*r.filterTLSDomainID, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
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

// APIUpdateBulkTLSCertRequest represents a request for the resource.
type APIUpdateBulkTLSCertRequest struct {
	ctx context.Context
	APIService TLSBulkCertificatesAPI
	certificateID string
	tlsBulkCertificate *TLSBulkCertificate
}

// TLSBulkCertificate returns a pointer to a request.
func (r *APIUpdateBulkTLSCertRequest) TLSBulkCertificate(tlsBulkCertificate TLSBulkCertificate) *APIUpdateBulkTLSCertRequest {
	r.tlsBulkCertificate = &tlsBulkCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateBulkTLSCertRequest) Execute() (*TLSBulkCertificateResponse, *http.Response, error) {
	return r.APIService.UpdateBulkTLSCertExecute(r)
}

/*
UpdateBulkTLSCert Update a certificate

Replace a certificate with a newly reissued certificate. By using this endpoint, the original certificate will cease to be used for future TLS handshakes. Thus, only SAN entries that appear in the replacement certificate will become TLS enabled. Any SAN entries that are missing in the replacement certificate will become disabled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certificateID Alphanumeric string identifying a TLS bulk certificate.
 @return APIUpdateBulkTLSCertRequest
*/
func (a *TLSBulkCertificatesAPIService) UpdateBulkTLSCert(ctx context.Context, certificateID string) APIUpdateBulkTLSCertRequest {
	return APIUpdateBulkTLSCertRequest{
		APIService: a,
		ctx: ctx,
		certificateID: certificateID,
	}
}

// UpdateBulkTLSCertExecute executes the request
//  @return TLSBulkCertificateResponse
func (a *TLSBulkCertificatesAPIService) UpdateBulkTLSCertExecute(r APIUpdateBulkTLSCertRequest) (*TLSBulkCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSBulkCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSBulkCertificatesAPIService.UpdateBulkTLSCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates/{certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"certificate_id"+"}", gourl.PathEscape(parameterToString(r.certificateID, "")))

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
	localVarPostBody = r.tlsBulkCertificate
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

// APIUploadTLSBulkCertRequest represents a request for the resource.
type APIUploadTLSBulkCertRequest struct {
	ctx context.Context
	APIService TLSBulkCertificatesAPI
	tlsBulkCertificate *TLSBulkCertificate
}

// TLSBulkCertificate returns a pointer to a request.
func (r *APIUploadTLSBulkCertRequest) TLSBulkCertificate(tlsBulkCertificate TLSBulkCertificate) *APIUploadTLSBulkCertRequest {
	r.tlsBulkCertificate = &tlsBulkCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUploadTLSBulkCertRequest) Execute() (*TLSBulkCertificateResponse, *http.Response, error) {
	return r.APIService.UploadTLSBulkCertExecute(r)
}

/*
UploadTLSBulkCert Upload a certificate

Upload a new certificate. TLS domains are automatically enabled upon certificate creation. If a domain is already enabled on a previously uploaded certificate, that domain will be updated to use the new certificate for all future TLS handshake requests.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIUploadTLSBulkCertRequest
*/
func (a *TLSBulkCertificatesAPIService) UploadTLSBulkCert(ctx context.Context) APIUploadTLSBulkCertRequest {
	return APIUploadTLSBulkCertRequest{
		APIService: a,
		ctx: ctx,
	}
}

// UploadTLSBulkCertExecute executes the request
//  @return TLSBulkCertificateResponse
func (a *TLSBulkCertificatesAPIService) UploadTLSBulkCertExecute(r APIUploadTLSBulkCertRequest) (*TLSBulkCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSBulkCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSBulkCertificatesAPIService.UploadTLSBulkCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates"

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
	localVarPostBody = r.tlsBulkCertificate
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
