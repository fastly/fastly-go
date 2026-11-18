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

// MutualAuthenticationAPI defines an interface for interacting with the resource.
type MutualAuthenticationAPI interface {

	/*
		CreateMutualTLSAuthentication Create a Mutual Authentication

		Create a mutual authentication using a bundle of certificates to enable client-to-server mutual TLS.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateMutualTLSAuthenticationRequest
	*/
	CreateMutualTLSAuthentication(ctx context.Context) APICreateMutualTLSAuthenticationRequest

	// CreateMutualTLSAuthenticationExecute executes the request
	//  @return MutualAuthenticationResponse
	CreateMutualTLSAuthenticationExecute(r APICreateMutualTLSAuthenticationRequest) (*MutualAuthenticationResponse, *http.Response, error)

	/*
		DeleteMutualTLS Delete a Mutual TLS

		Remove a Mutual TLS authentication

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param mutualAuthenticationID Alphanumeric string identifying a mutual authentication.
		 @return APIDeleteMutualTLSRequest
	*/
	DeleteMutualTLS(ctx context.Context, mutualAuthenticationID string) APIDeleteMutualTLSRequest

	// DeleteMutualTLSExecute executes the request
	DeleteMutualTLSExecute(r APIDeleteMutualTLSRequest) (*http.Response, error)

	/*
		GetMutualAuthentication Get a Mutual Authentication

		Show a Mutual Authentication.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param mutualAuthenticationID Alphanumeric string identifying a mutual authentication.
		 @return APIGetMutualAuthenticationRequest
	*/
	GetMutualAuthentication(ctx context.Context, mutualAuthenticationID string) APIGetMutualAuthenticationRequest

	// GetMutualAuthenticationExecute executes the request
	//  @return MutualAuthenticationResponse
	GetMutualAuthenticationExecute(r APIGetMutualAuthenticationRequest) (*MutualAuthenticationResponse, *http.Response, error)

	/*
		ListMutualAuthentications List Mutual Authentications

		List all mutual authentications.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListMutualAuthenticationsRequest
	*/
	ListMutualAuthentications(ctx context.Context) APIListMutualAuthenticationsRequest

	// ListMutualAuthenticationsExecute executes the request
	//  @return MutualAuthenticationsResponse
	ListMutualAuthenticationsExecute(r APIListMutualAuthenticationsRequest) (*MutualAuthenticationsResponse, *http.Response, error)

	/*
		PatchMutualAuthentication Update a Mutual Authentication

		Update a Mutual Authentication.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param mutualAuthenticationID Alphanumeric string identifying a mutual authentication.
		 @return APIPatchMutualAuthenticationRequest
	*/
	PatchMutualAuthentication(ctx context.Context, mutualAuthenticationID string) APIPatchMutualAuthenticationRequest

	// PatchMutualAuthenticationExecute executes the request
	//  @return MutualAuthenticationResponse
	PatchMutualAuthenticationExecute(r APIPatchMutualAuthenticationRequest) (*MutualAuthenticationResponse, *http.Response, error)
}

// MutualAuthenticationAPIService MutualAuthenticationAPI service
type MutualAuthenticationAPIService service

// APICreateMutualTLSAuthenticationRequest represents a request for the resource.
type APICreateMutualTLSAuthenticationRequest struct {
	ctx                  context.Context
	APIService           MutualAuthenticationAPI
	mutualAuthentication *MutualAuthentication
}

// MutualAuthentication returns a pointer to a request.
func (r *APICreateMutualTLSAuthenticationRequest) MutualAuthentication(mutualAuthentication MutualAuthentication) *APICreateMutualTLSAuthenticationRequest {
	r.mutualAuthentication = &mutualAuthentication
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateMutualTLSAuthenticationRequest) Execute() (*MutualAuthenticationResponse, *http.Response, error) {
	return r.APIService.CreateMutualTLSAuthenticationExecute(r)
}

/*
CreateMutualTLSAuthentication Create a Mutual Authentication

Create a mutual authentication using a bundle of certificates to enable client-to-server mutual TLS.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateMutualTLSAuthenticationRequest
*/
func (a *MutualAuthenticationAPIService) CreateMutualTLSAuthentication(ctx context.Context) APICreateMutualTLSAuthenticationRequest {
	return APICreateMutualTLSAuthenticationRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateMutualTLSAuthenticationExecute executes the request
//  @return MutualAuthenticationResponse
func (a *MutualAuthenticationAPIService) CreateMutualTLSAuthenticationExecute(r APICreateMutualTLSAuthenticationRequest) (*MutualAuthenticationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MutualAuthenticationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualAuthenticationAPIService.CreateMutualTLSAuthentication")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/mutual_authentications"

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
	localVarPostBody = r.mutualAuthentication
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

// APIDeleteMutualTLSRequest represents a request for the resource.
type APIDeleteMutualTLSRequest struct {
	ctx                    context.Context
	APIService             MutualAuthenticationAPI
	mutualAuthenticationID string
}

// Execute calls the API using the request data configured.
func (r APIDeleteMutualTLSRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteMutualTLSExecute(r)
}

/*
DeleteMutualTLS Delete a Mutual TLS

Remove a Mutual TLS authentication

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mutualAuthenticationID Alphanumeric string identifying a mutual authentication.
 @return APIDeleteMutualTLSRequest
*/
func (a *MutualAuthenticationAPIService) DeleteMutualTLS(ctx context.Context, mutualAuthenticationID string) APIDeleteMutualTLSRequest {
	return APIDeleteMutualTLSRequest{
		APIService:             a,
		ctx:                    ctx,
		mutualAuthenticationID: mutualAuthenticationID,
	}
}

// DeleteMutualTLSExecute executes the request
func (a *MutualAuthenticationAPIService) DeleteMutualTLSExecute(r APIDeleteMutualTLSRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualAuthenticationAPIService.DeleteMutualTLS")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/mutual_authentications/{mutual_authentication_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"mutual_authentication_id"+"}", gourl.PathEscape(parameterToString(r.mutualAuthenticationID, "")))

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

// APIGetMutualAuthenticationRequest represents a request for the resource.
type APIGetMutualAuthenticationRequest struct {
	ctx                    context.Context
	APIService             MutualAuthenticationAPI
	mutualAuthenticationID string
	include                *string
}

// Include Comma-separated list of related objects to include (optional). Permitted values: &#x60;tls_activations&#x60;. Including TLS activations will provide you with the TLS domain names that are related to your Mutual TLS authentication.
func (r *APIGetMutualAuthenticationRequest) Include(include string) *APIGetMutualAuthenticationRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetMutualAuthenticationRequest) Execute() (*MutualAuthenticationResponse, *http.Response, error) {
	return r.APIService.GetMutualAuthenticationExecute(r)
}

/*
GetMutualAuthentication Get a Mutual Authentication

Show a Mutual Authentication.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mutualAuthenticationID Alphanumeric string identifying a mutual authentication.
 @return APIGetMutualAuthenticationRequest
*/
func (a *MutualAuthenticationAPIService) GetMutualAuthentication(ctx context.Context, mutualAuthenticationID string) APIGetMutualAuthenticationRequest {
	return APIGetMutualAuthenticationRequest{
		APIService:             a,
		ctx:                    ctx,
		mutualAuthenticationID: mutualAuthenticationID,
	}
}

// GetMutualAuthenticationExecute executes the request
//  @return MutualAuthenticationResponse
func (a *MutualAuthenticationAPIService) GetMutualAuthenticationExecute(r APIGetMutualAuthenticationRequest) (*MutualAuthenticationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MutualAuthenticationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualAuthenticationAPIService.GetMutualAuthentication")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/mutual_authentications/{mutual_authentication_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"mutual_authentication_id"+"}", gourl.PathEscape(parameterToString(r.mutualAuthenticationID, "")))

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

// APIListMutualAuthenticationsRequest represents a request for the resource.
type APIListMutualAuthenticationsRequest struct {
	ctx        context.Context
	APIService MutualAuthenticationAPI
	include    *string
	pageNumber *int32
	pageSize   *int32
}

// Include Comma-separated list of related objects to include (optional). Permitted values: &#x60;tls_activations&#x60;. Including TLS activations will provide you with the TLS domain names that are related to your Mutual TLS authentication.
func (r *APIListMutualAuthenticationsRequest) Include(include string) *APIListMutualAuthenticationsRequest {
	r.include = &include
	return r
}

// PageNumber Current page.
func (r *APIListMutualAuthenticationsRequest) PageNumber(pageNumber int32) *APIListMutualAuthenticationsRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListMutualAuthenticationsRequest) PageSize(pageSize int32) *APIListMutualAuthenticationsRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListMutualAuthenticationsRequest) Execute() (*MutualAuthenticationsResponse, *http.Response, error) {
	return r.APIService.ListMutualAuthenticationsExecute(r)
}

/*
ListMutualAuthentications List Mutual Authentications

List all mutual authentications.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListMutualAuthenticationsRequest
*/
func (a *MutualAuthenticationAPIService) ListMutualAuthentications(ctx context.Context) APIListMutualAuthenticationsRequest {
	return APIListMutualAuthenticationsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListMutualAuthenticationsExecute executes the request
//  @return MutualAuthenticationsResponse
func (a *MutualAuthenticationAPIService) ListMutualAuthenticationsExecute(r APIListMutualAuthenticationsRequest) (*MutualAuthenticationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MutualAuthenticationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualAuthenticationAPIService.ListMutualAuthentications")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/mutual_authentications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

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

// APIPatchMutualAuthenticationRequest represents a request for the resource.
type APIPatchMutualAuthenticationRequest struct {
	ctx                    context.Context
	APIService             MutualAuthenticationAPI
	mutualAuthenticationID string
	mutualAuthentication   *MutualAuthentication
}

// MutualAuthentication returns a pointer to a request.
func (r *APIPatchMutualAuthenticationRequest) MutualAuthentication(mutualAuthentication MutualAuthentication) *APIPatchMutualAuthenticationRequest {
	r.mutualAuthentication = &mutualAuthentication
	return r
}

// Execute calls the API using the request data configured.
func (r APIPatchMutualAuthenticationRequest) Execute() (*MutualAuthenticationResponse, *http.Response, error) {
	return r.APIService.PatchMutualAuthenticationExecute(r)
}

/*
PatchMutualAuthentication Update a Mutual Authentication

Update a Mutual Authentication.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mutualAuthenticationID Alphanumeric string identifying a mutual authentication.
 @return APIPatchMutualAuthenticationRequest
*/
func (a *MutualAuthenticationAPIService) PatchMutualAuthentication(ctx context.Context, mutualAuthenticationID string) APIPatchMutualAuthenticationRequest {
	return APIPatchMutualAuthenticationRequest{
		APIService:             a,
		ctx:                    ctx,
		mutualAuthenticationID: mutualAuthenticationID,
	}
}

// PatchMutualAuthenticationExecute executes the request
//  @return MutualAuthenticationResponse
func (a *MutualAuthenticationAPIService) PatchMutualAuthenticationExecute(r APIPatchMutualAuthenticationRequest) (*MutualAuthenticationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MutualAuthenticationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualAuthenticationAPIService.PatchMutualAuthentication")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/mutual_authentications/{mutual_authentication_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"mutual_authentication_id"+"}", gourl.PathEscape(parameterToString(r.mutualAuthenticationID, "")))

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
	localVarPostBody = r.mutualAuthentication
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
