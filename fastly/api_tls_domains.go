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
)

// Linger please
var (
	_ context.Context
)

// TLSDomainsAPI defines an interface for interacting with the resource.
type TLSDomainsAPI interface {

	/*
		ListTLSDomains List TLS domains

		List all TLS domains.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListTLSDomainsRequest
	*/
	ListTLSDomains(ctx context.Context) APIListTLSDomainsRequest

	// ListTLSDomainsExecute executes the request
	//  @return TLSDomainsResponse
	ListTLSDomainsExecute(r APIListTLSDomainsRequest) (*TLSDomainsResponse, *http.Response, error)
}

// TLSDomainsAPIService TLSDomainsAPI service
type TLSDomainsAPIService service

// APIListTLSDomainsRequest represents a request for the resource.
type APIListTLSDomainsRequest struct {
	ctx                      context.Context
	APIService               TLSDomainsAPI
	filterInUse              *string
	filterTLSCertificatesID  *string
	filterTLSSubscriptionsID *string
	include                  *string
	sort                     *string
	pageNumber               *int32
	pageSize                 *int32
}

// FilterInUse Optional. Limit the returned domains to those currently using Fastly to terminate TLS with SNI (that is, domains considered \&quot;in use\&quot;) Permitted values: true, false.
func (r *APIListTLSDomainsRequest) FilterInUse(filterInUse string) *APIListTLSDomainsRequest {
	r.filterInUse = &filterInUse
	return r
}

// FilterTLSCertificatesID Optional. Limit the returned domains to those listed in the given TLS certificate&#39;s SAN list.
func (r *APIListTLSDomainsRequest) FilterTLSCertificatesID(filterTLSCertificatesID string) *APIListTLSDomainsRequest {
	r.filterTLSCertificatesID = &filterTLSCertificatesID
	return r
}

// FilterTLSSubscriptionsID Optional. Limit the returned domains to those for a given TLS subscription.
func (r *APIListTLSDomainsRequest) FilterTLSSubscriptionsID(filterTLSSubscriptionsID string) *APIListTLSDomainsRequest {
	r.filterTLSSubscriptionsID = &filterTLSSubscriptionsID
	return r
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_activations&#x60;, &#x60;tls_certificates&#x60;, &#x60;tls_subscriptions&#x60;, &#x60;tls_subscriptions.tls_authorizations&#x60;, &#x60;tls_authorizations.globalsign_email_challenge&#x60;, and &#x60;tls_authorizations.self_managed_http_challenge&#x60;.
func (r *APIListTLSDomainsRequest) Include(include string) *APIListTLSDomainsRequest {
	r.include = &include
	return r
}

// Sort The order in which to list the results.
func (r *APIListTLSDomainsRequest) Sort(sort string) *APIListTLSDomainsRequest {
	r.sort = &sort
	return r
}

// PageNumber Current page.
func (r *APIListTLSDomainsRequest) PageNumber(pageNumber int32) *APIListTLSDomainsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListTLSDomainsRequest) PageSize(pageSize int32) *APIListTLSDomainsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSDomainsRequest) Execute() (*TLSDomainsResponse, *http.Response, error) {
	return r.APIService.ListTLSDomainsExecute(r)
}

/*
ListTLSDomains List TLS domains

List all TLS domains.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSDomainsRequest
*/
func (a *TLSDomainsAPIService) ListTLSDomains(ctx context.Context) APIListTLSDomainsRequest {
	return APIListTLSDomainsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListTLSDomainsExecute executes the request
//  @return TLSDomainsResponse
func (a *TLSDomainsAPIService) ListTLSDomainsExecute(r APIListTLSDomainsRequest) (*TLSDomainsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TLSDomainsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSDomainsAPIService.ListTLSDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/domains"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterInUse != nil {
		localVarQueryParams.Add("filter[in_use]", parameterToString(*r.filterInUse, ""))
	}
	if r.filterTLSCertificatesID != nil {
		localVarQueryParams.Add("filter[tls_certificates.id]", parameterToString(*r.filterTLSCertificatesID, ""))
	}
	if r.filterTLSSubscriptionsID != nil {
		localVarQueryParams.Add("filter[tls_subscriptions.id]", parameterToString(*r.filterTLSSubscriptionsID, ""))
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
