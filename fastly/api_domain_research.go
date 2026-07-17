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

// DomainResearchAPI defines an interface for interacting with the resource.
type DomainResearchAPI interface {

	/*
		DomainStatus Domain status

		The `Status` method checks the availability status of a single domain name.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIDomainStatusRequest
	*/
	DomainStatus(ctx context.Context) APIDomainStatusRequest

	// DomainStatusExecute executes the request
	//  @return Status
	DomainStatusExecute(r APIDomainStatusRequest) (*Status, *http.Response, error)

	/*
		SuggestDomains Suggest domains

		The `Suggest` method performs a real-time query of the search term(s) against the [known zone database](http://zonedb.org),
	making recommendations, stemming, and applying Unicode folding, IDN normalization, registrar supported-zone
	restrictions, and other refinements. **Note:** `Suggest` method responses do not include domain availability status.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APISuggestDomainsRequest
	*/
	SuggestDomains(ctx context.Context) APISuggestDomainsRequest

	// SuggestDomainsExecute executes the request
	//  @return InlineResponse20016
	SuggestDomainsExecute(r APISuggestDomainsRequest) (*InlineResponse20016, *http.Response, error)
}

// DomainResearchAPIService DomainResearchAPI service
type DomainResearchAPIService service

// APIDomainStatusRequest represents a request for the resource.
type APIDomainStatusRequest struct {
	ctx        context.Context
	APIService DomainResearchAPI
	domain     *string
	scope      *string
}

// Domain returns a pointer to a request.
func (r *APIDomainStatusRequest) Domain(domain string) *APIDomainStatusRequest {
	r.domain = &domain
	return r
}

// Scope returns a pointer to a request.
func (r *APIDomainStatusRequest) Scope(scope string) *APIDomainStatusRequest {
	r.scope = &scope
	return r
}

// Execute calls the API using the request data configured.
func (r APIDomainStatusRequest) Execute() (*Status, *http.Response, error) {
	return r.APIService.DomainStatusExecute(r)
}

/*
DomainStatus Domain status

The `Status` method checks the availability status of a single domain name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIDomainStatusRequest
*/
func (a *DomainResearchAPIService) DomainStatus(ctx context.Context) APIDomainStatusRequest {
	return APIDomainStatusRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// DomainStatusExecute executes the request
//  @return Status
func (a *DomainResearchAPIService) DomainStatusExecute(r APIDomainStatusRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DomainResearchAPIService.DomainStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/tools/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.domain == nil {
		return localVarReturnValue, nil, reportError("domain is required and must be specified")
	}

	localVarQueryParams.Add("domain", parameterToString(*r.domain, ""))
	if r.scope != nil {
		localVarQueryParams.Add("scope", parameterToString(*r.scope, ""))
	}
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

// APISuggestDomainsRequest represents a request for the resource.
type APISuggestDomainsRequest struct {
	ctx        context.Context
	APIService DomainResearchAPI
	query      *string
	defaults   *string
	keywords   *string
	location   *string
	vendor     *string
}

// Query returns a pointer to a request.
func (r *APISuggestDomainsRequest) Query(query string) *APISuggestDomainsRequest {
	r.query = &query
	return r
}

// Defaults returns a pointer to a request.
func (r *APISuggestDomainsRequest) Defaults(defaults string) *APISuggestDomainsRequest {
	r.defaults = &defaults
	return r
}

// Keywords returns a pointer to a request.
func (r *APISuggestDomainsRequest) Keywords(keywords string) *APISuggestDomainsRequest {
	r.keywords = &keywords
	return r
}

// Location returns a pointer to a request.
func (r *APISuggestDomainsRequest) Location(location string) *APISuggestDomainsRequest {
	r.location = &location
	return r
}

// Vendor returns a pointer to a request.
func (r *APISuggestDomainsRequest) Vendor(vendor string) *APISuggestDomainsRequest {
	r.vendor = &vendor
	return r
}

// Execute calls the API using the request data configured.
func (r APISuggestDomainsRequest) Execute() (*InlineResponse20016, *http.Response, error) {
	return r.APIService.SuggestDomainsExecute(r)
}

/*
SuggestDomains Suggest domains

The `Suggest` method performs a real-time query of the search term(s) against the [known zone database](http://zonedb.org),
making recommendations, stemming, and applying Unicode folding, IDN normalization, registrar supported-zone
restrictions, and other refinements. **Note:** `Suggest` method responses do not include domain availability status.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APISuggestDomainsRequest
*/
func (a *DomainResearchAPIService) SuggestDomains(ctx context.Context) APISuggestDomainsRequest {
	return APISuggestDomainsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// SuggestDomainsExecute executes the request
//  @return InlineResponse20016
func (a *DomainResearchAPIService) SuggestDomainsExecute(r APISuggestDomainsRequest) (*InlineResponse20016, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse20016
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DomainResearchAPIService.SuggestDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/tools/suggest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.query == nil {
		return localVarReturnValue, nil, reportError("query is required and must be specified")
	}

	localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	if r.defaults != nil {
		localVarQueryParams.Add("defaults", parameterToString(*r.defaults, ""))
	}
	if r.keywords != nil {
		localVarQueryParams.Add("keywords", parameterToString(*r.keywords, ""))
	}
	if r.location != nil {
		localVarQueryParams.Add("location", parameterToString(*r.location, ""))
	}
	if r.vendor != nil {
		localVarQueryParams.Add("vendor", parameterToString(*r.vendor, ""))
	}
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
