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

// TlsBulkCertificatesAPI defines an interface for interacting with the resource.
type TlsBulkCertificatesAPI interface {

	/*
		DeleteBulkTlsCert Delete a certificate

		Destroy a certificate. This disables TLS for all domains listed as SAN entries.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param certificateId Alphanumeric string identifying a TLS bulk certificate.
		 @return APIDeleteBulkTlsCertRequest
	*/
	DeleteBulkTlsCert(ctx context.Context, certificateId string) APIDeleteBulkTlsCertRequest

	// DeleteBulkTlsCertExecute executes the request
	DeleteBulkTlsCertExecute(r APIDeleteBulkTlsCertRequest) (*http.Response, error)

	/*
		GetTlsBulkCert Get a certificate

		Retrieve a single certificate.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param certificateId Alphanumeric string identifying a TLS bulk certificate.
		 @return APIGetTlsBulkCertRequest
	*/
	GetTlsBulkCert(ctx context.Context, certificateId string) APIGetTlsBulkCertRequest

	// GetTlsBulkCertExecute executes the request
	//  @return TlsBulkCertificateResponse
	GetTlsBulkCertExecute(r APIGetTlsBulkCertRequest) (*TlsBulkCertificateResponse, *http.Response, error)

	/*
		ListTlsBulkCerts List certificates

		List all certificates.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTlsBulkCertsRequest
	*/
	ListTlsBulkCerts(ctx context.Context) APIListTlsBulkCertsRequest

	// ListTlsBulkCertsExecute executes the request
	//  @return TlsBulkCertificatesResponse
	ListTlsBulkCertsExecute(r APIListTlsBulkCertsRequest) (*TlsBulkCertificatesResponse, *http.Response, error)

	/*
		UpdateBulkTlsCert Update a certificate

		Replace a certificate with a newly reissued certificate. By using this endpoint, the original certificate will cease to be used for future TLS handshakes. Thus, only SAN entries that appear in the replacement certificate will become TLS enabled. Any SAN entries that are missing in the replacement certificate will become disabled.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param certificateId Alphanumeric string identifying a TLS bulk certificate.
		 @return APIUpdateBulkTlsCertRequest
	*/
	UpdateBulkTlsCert(ctx context.Context, certificateId string) APIUpdateBulkTlsCertRequest

	// UpdateBulkTlsCertExecute executes the request
	//  @return TlsBulkCertificateResponse
	UpdateBulkTlsCertExecute(r APIUpdateBulkTlsCertRequest) (*TlsBulkCertificateResponse, *http.Response, error)

	/*
		UploadTlsBulkCert Upload a certificate

		Upload a new certificate. TLS domains are automatically enabled upon certificate creation. If a domain is already enabled on a previously uploaded certificate, that domain will be updated to use the new certificate for all future TLS handshake requests.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIUploadTlsBulkCertRequest
	*/
	UploadTlsBulkCert(ctx context.Context) APIUploadTlsBulkCertRequest

	// UploadTlsBulkCertExecute executes the request
	//  @return TlsBulkCertificateResponse
	UploadTlsBulkCertExecute(r APIUploadTlsBulkCertRequest) (*TlsBulkCertificateResponse, *http.Response, error)
}

// TlsBulkCertificatesAPIService TlsBulkCertificatesAPI service
type TlsBulkCertificatesAPIService service

// APIDeleteBulkTlsCertRequest represents a request for the resource.
type APIDeleteBulkTlsCertRequest struct {
	ctx           context.Context
	APIService    TlsBulkCertificatesAPI
	certificateId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteBulkTlsCertRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteBulkTlsCertExecute(r)
}

/*
DeleteBulkTlsCert Delete a certificate

Destroy a certificate. This disables TLS for all domains listed as SAN entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certificateId Alphanumeric string identifying a TLS bulk certificate.
 @return APIDeleteBulkTlsCertRequest
*/
func (a *TlsBulkCertificatesAPIService) DeleteBulkTlsCert(ctx context.Context, certificateId string) APIDeleteBulkTlsCertRequest {
	return APIDeleteBulkTlsCertRequest{
		APIService:    a,
		ctx:           ctx,
		certificateId: certificateId,
	}
}

// DeleteBulkTlsCertExecute executes the request
func (a *TlsBulkCertificatesAPIService) DeleteBulkTlsCertExecute(r APIDeleteBulkTlsCertRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsBulkCertificatesAPIService.DeleteBulkTlsCert")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates/{certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"certificate_id"+"}", gourl.PathEscape(parameterToString(r.certificateId, "")))

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

// APIGetTlsBulkCertRequest represents a request for the resource.
type APIGetTlsBulkCertRequest struct {
	ctx           context.Context
	APIService    TlsBulkCertificatesAPI
	certificateId string
}

// Execute calls the API using the request data configured.
func (r APIGetTlsBulkCertRequest) Execute() (*TlsBulkCertificateResponse, *http.Response, error) {
	return r.APIService.GetTlsBulkCertExecute(r)
}

/*
GetTlsBulkCert Get a certificate

Retrieve a single certificate.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certificateId Alphanumeric string identifying a TLS bulk certificate.
 @return APIGetTlsBulkCertRequest
*/
func (a *TlsBulkCertificatesAPIService) GetTlsBulkCert(ctx context.Context, certificateId string) APIGetTlsBulkCertRequest {
	return APIGetTlsBulkCertRequest{
		APIService:    a,
		ctx:           ctx,
		certificateId: certificateId,
	}
}

// GetTlsBulkCertExecute executes the request
//  @return TlsBulkCertificateResponse
func (a *TlsBulkCertificatesAPIService) GetTlsBulkCertExecute(r APIGetTlsBulkCertRequest) (*TlsBulkCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsBulkCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsBulkCertificatesAPIService.GetTlsBulkCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates/{certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"certificate_id"+"}", gourl.PathEscape(parameterToString(r.certificateId, "")))

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

// APIListTlsBulkCertsRequest represents a request for the resource.
type APIListTlsBulkCertsRequest struct {
	ctx               context.Context
	APIService        TlsBulkCertificatesAPI
	filterTlsDomainId *string
	filterNotBefore   *string
	filterNotAfter    *string
	pageNumber        *int32
	pageSize          *int32
	sort              *string
}

// FilterTlsDomainId Filter certificates by their matching, fully-qualified domain name.
func (r *APIListTlsBulkCertsRequest) FilterTlsDomainId(filterTlsDomainId string) *APIListTlsBulkCertsRequest {
	r.filterTlsDomainId = &filterTlsDomainId
	return r
}

// FilterNotBefore Filter the returned certificates by not_before date in UTC.  Accepts parameters: lt, lte, gt, gte (e.g., filter[not_before][gte]&#x3D;2020-05-05).
func (r *APIListTlsBulkCertsRequest) FilterNotBefore(filterNotBefore string) *APIListTlsBulkCertsRequest {
	r.filterNotBefore = &filterNotBefore
	return r
}

// FilterNotAfter Filter the returned certificates by expiry date in UTC.  Accepts parameters: lt, lte, gt, gte (e.g., filter[not_after][lte]&#x3D;2020-05-05).
func (r *APIListTlsBulkCertsRequest) FilterNotAfter(filterNotAfter string) *APIListTlsBulkCertsRequest {
	r.filterNotAfter = &filterNotAfter
	return r
}

// PageNumber Current page.
func (r *APIListTlsBulkCertsRequest) PageNumber(pageNumber int32) *APIListTlsBulkCertsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTlsBulkCertsRequest) PageSize(pageSize int32) *APIListTlsBulkCertsRequest {
	r.pageSize = &pageSize
	return r
}

// Sort The order in which to list the results by creation date.
func (r *APIListTlsBulkCertsRequest) Sort(sort string) *APIListTlsBulkCertsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTlsBulkCertsRequest) Execute() (*TlsBulkCertificatesResponse, *http.Response, error) {
	return r.APIService.ListTlsBulkCertsExecute(r)
}

/*
ListTlsBulkCerts List certificates

List all certificates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTlsBulkCertsRequest
*/
func (a *TlsBulkCertificatesAPIService) ListTlsBulkCerts(ctx context.Context) APIListTlsBulkCertsRequest {
	return APIListTlsBulkCertsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTlsBulkCertsExecute executes the request
//  @return TlsBulkCertificatesResponse
func (a *TlsBulkCertificatesAPIService) ListTlsBulkCertsExecute(r APIListTlsBulkCertsRequest) (*TlsBulkCertificatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsBulkCertificatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsBulkCertificatesAPIService.ListTlsBulkCerts")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterTlsDomainId != nil {
		localVarQueryParams.Add("filter[tls_domain.id]", parameterToString(*r.filterTlsDomainId, ""))
	}
	if r.filterNotBefore != nil {
		localVarQueryParams.Add("filter[not_before]", parameterToString(*r.filterNotBefore, ""))
	}
	if r.filterNotAfter != nil {
		localVarQueryParams.Add("filter[not_after]", parameterToString(*r.filterNotAfter, ""))
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

// APIUpdateBulkTlsCertRequest represents a request for the resource.
type APIUpdateBulkTlsCertRequest struct {
	ctx                context.Context
	APIService         TlsBulkCertificatesAPI
	certificateId      string
	tlsBulkCertificate *TlsBulkCertificate
}

// TlsBulkCertificate returns a pointer to a request.
func (r *APIUpdateBulkTlsCertRequest) TlsBulkCertificate(tlsBulkCertificate TlsBulkCertificate) *APIUpdateBulkTlsCertRequest {
	r.tlsBulkCertificate = &tlsBulkCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateBulkTlsCertRequest) Execute() (*TlsBulkCertificateResponse, *http.Response, error) {
	return r.APIService.UpdateBulkTlsCertExecute(r)
}

/*
UpdateBulkTlsCert Update a certificate

Replace a certificate with a newly reissued certificate. By using this endpoint, the original certificate will cease to be used for future TLS handshakes. Thus, only SAN entries that appear in the replacement certificate will become TLS enabled. Any SAN entries that are missing in the replacement certificate will become disabled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certificateId Alphanumeric string identifying a TLS bulk certificate.
 @return APIUpdateBulkTlsCertRequest
*/
func (a *TlsBulkCertificatesAPIService) UpdateBulkTlsCert(ctx context.Context, certificateId string) APIUpdateBulkTlsCertRequest {
	return APIUpdateBulkTlsCertRequest{
		APIService:    a,
		ctx:           ctx,
		certificateId: certificateId,
	}
}

// UpdateBulkTlsCertExecute executes the request
//  @return TlsBulkCertificateResponse
func (a *TlsBulkCertificatesAPIService) UpdateBulkTlsCertExecute(r APIUpdateBulkTlsCertRequest) (*TlsBulkCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsBulkCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsBulkCertificatesAPIService.UpdateBulkTlsCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/bulk/certificates/{certificate_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"certificate_id"+"}", gourl.PathEscape(parameterToString(r.certificateId, "")))

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

// APIUploadTlsBulkCertRequest represents a request for the resource.
type APIUploadTlsBulkCertRequest struct {
	ctx                context.Context
	APIService         TlsBulkCertificatesAPI
	tlsBulkCertificate *TlsBulkCertificate
}

// TlsBulkCertificate returns a pointer to a request.
func (r *APIUploadTlsBulkCertRequest) TlsBulkCertificate(tlsBulkCertificate TlsBulkCertificate) *APIUploadTlsBulkCertRequest {
	r.tlsBulkCertificate = &tlsBulkCertificate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUploadTlsBulkCertRequest) Execute() (*TlsBulkCertificateResponse, *http.Response, error) {
	return r.APIService.UploadTlsBulkCertExecute(r)
}

/*
UploadTlsBulkCert Upload a certificate

Upload a new certificate. TLS domains are automatically enabled upon certificate creation. If a domain is already enabled on a previously uploaded certificate, that domain will be updated to use the new certificate for all future TLS handshake requests.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIUploadTlsBulkCertRequest
*/
func (a *TlsBulkCertificatesAPIService) UploadTlsBulkCert(ctx context.Context) APIUploadTlsBulkCertRequest {
	return APIUploadTlsBulkCertRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// UploadTlsBulkCertExecute executes the request
//  @return TlsBulkCertificateResponse
func (a *TlsBulkCertificatesAPIService) UploadTlsBulkCertExecute(r APIUploadTlsBulkCertRequest) (*TlsBulkCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TlsBulkCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TlsBulkCertificatesAPIService.UploadTlsBulkCert")
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
