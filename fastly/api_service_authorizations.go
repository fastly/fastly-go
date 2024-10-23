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

// ServiceAuthorizationsAPI defines an interface for interacting with the resource.
type ServiceAuthorizationsAPI interface {

	/*
		CreateServiceAuthorization Create service authorization

		Create service authorization.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateServiceAuthorizationRequest
	*/
	CreateServiceAuthorization(ctx context.Context) APICreateServiceAuthorizationRequest

	// CreateServiceAuthorizationExecute executes the request
	//  @return ServiceAuthorizationResponse
	CreateServiceAuthorizationExecute(r APICreateServiceAuthorizationRequest) (*ServiceAuthorizationResponse, *http.Response, error)

	/*
		DeleteServiceAuthorization Delete service authorization

		Delete service authorization.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceAuthorizationID Alphanumeric string identifying a service authorization.
		 @return APIDeleteServiceAuthorizationRequest
	*/
	DeleteServiceAuthorization(ctx context.Context, serviceAuthorizationID string) APIDeleteServiceAuthorizationRequest

	// DeleteServiceAuthorizationExecute executes the request
	DeleteServiceAuthorizationExecute(r APIDeleteServiceAuthorizationRequest) (*http.Response, error)

	/*
		DeleteServiceAuthorization2 Delete service authorizations

		Delete service authorizations.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIDeleteServiceAuthorization2Request
	*/
	DeleteServiceAuthorization2(ctx context.Context) APIDeleteServiceAuthorization2Request

	// DeleteServiceAuthorization2Execute executes the request
	//  @return InlineResponse2007
	DeleteServiceAuthorization2Execute(r APIDeleteServiceAuthorization2Request) (*InlineResponse2007, *http.Response, error)

	/*
		ListServiceAuthorization List service authorizations

		List service authorizations.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListServiceAuthorizationRequest
	*/
	ListServiceAuthorization(ctx context.Context) APIListServiceAuthorizationRequest

	// ListServiceAuthorizationExecute executes the request
	//  @return ServiceAuthorizationsResponse
	ListServiceAuthorizationExecute(r APIListServiceAuthorizationRequest) (*ServiceAuthorizationsResponse, *http.Response, error)

	/*
		ShowServiceAuthorization Show service authorization

		Show service authorization.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceAuthorizationID Alphanumeric string identifying a service authorization.
		 @return APIShowServiceAuthorizationRequest
	*/
	ShowServiceAuthorization(ctx context.Context, serviceAuthorizationID string) APIShowServiceAuthorizationRequest

	// ShowServiceAuthorizationExecute executes the request
	//  @return ServiceAuthorizationResponse
	ShowServiceAuthorizationExecute(r APIShowServiceAuthorizationRequest) (*ServiceAuthorizationResponse, *http.Response, error)

	/*
		UpdateServiceAuthorization Update service authorization

		Update service authorization.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceAuthorizationID Alphanumeric string identifying a service authorization.
		 @return APIUpdateServiceAuthorizationRequest
	*/
	UpdateServiceAuthorization(ctx context.Context, serviceAuthorizationID string) APIUpdateServiceAuthorizationRequest

	// UpdateServiceAuthorizationExecute executes the request
	//  @return ServiceAuthorizationResponse
	UpdateServiceAuthorizationExecute(r APIUpdateServiceAuthorizationRequest) (*ServiceAuthorizationResponse, *http.Response, error)

	/*
		UpdateServiceAuthorization2 Update service authorizations

		Update service authorizations.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIUpdateServiceAuthorization2Request
	*/
	UpdateServiceAuthorization2(ctx context.Context) APIUpdateServiceAuthorization2Request

	// UpdateServiceAuthorization2Execute executes the request
	//  @return ServiceAuthorizationsResponse
	UpdateServiceAuthorization2Execute(r APIUpdateServiceAuthorization2Request) (*ServiceAuthorizationsResponse, *http.Response, error)
}

// ServiceAuthorizationsAPIService ServiceAuthorizationsAPI service
type ServiceAuthorizationsAPIService service

// APICreateServiceAuthorizationRequest represents a request for the resource.
type APICreateServiceAuthorizationRequest struct {
	ctx                  context.Context
	APIService           ServiceAuthorizationsAPI
	serviceAuthorization *ServiceAuthorization
}

// ServiceAuthorization returns a pointer to a request.
func (r *APICreateServiceAuthorizationRequest) ServiceAuthorization(serviceAuthorization ServiceAuthorization) *APICreateServiceAuthorizationRequest {
	r.serviceAuthorization = &serviceAuthorization
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateServiceAuthorizationRequest) Execute() (*ServiceAuthorizationResponse, *http.Response, error) {
	return r.APIService.CreateServiceAuthorizationExecute(r)
}

/*
CreateServiceAuthorization Create service authorization

Create service authorization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateServiceAuthorizationRequest
*/
func (a *ServiceAuthorizationsAPIService) CreateServiceAuthorization(ctx context.Context) APICreateServiceAuthorizationRequest {
	return APICreateServiceAuthorizationRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateServiceAuthorizationExecute executes the request
//  @return ServiceAuthorizationResponse
func (a *ServiceAuthorizationsAPIService) CreateServiceAuthorizationExecute(r APICreateServiceAuthorizationRequest) (*ServiceAuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.CreateServiceAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations"

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
	localVarPostBody = r.serviceAuthorization
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

// APIDeleteServiceAuthorizationRequest represents a request for the resource.
type APIDeleteServiceAuthorizationRequest struct {
	ctx                    context.Context
	APIService             ServiceAuthorizationsAPI
	serviceAuthorizationID string
}

// Execute calls the API using the request data configured.
func (r APIDeleteServiceAuthorizationRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteServiceAuthorizationExecute(r)
}

/*
DeleteServiceAuthorization Delete service authorization

Delete service authorization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceAuthorizationID Alphanumeric string identifying a service authorization.
 @return APIDeleteServiceAuthorizationRequest
*/
func (a *ServiceAuthorizationsAPIService) DeleteServiceAuthorization(ctx context.Context, serviceAuthorizationID string) APIDeleteServiceAuthorizationRequest {
	return APIDeleteServiceAuthorizationRequest{
		APIService:             a,
		ctx:                    ctx,
		serviceAuthorizationID: serviceAuthorizationID,
	}
}

// DeleteServiceAuthorizationExecute executes the request
func (a *ServiceAuthorizationsAPIService) DeleteServiceAuthorizationExecute(r APIDeleteServiceAuthorizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.DeleteServiceAuthorization")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations/{service_authorization_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_authorization_id"+"}", gourl.PathEscape(parameterToString(r.serviceAuthorizationID, "")))

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

// APIDeleteServiceAuthorization2Request represents a request for the resource.
type APIDeleteServiceAuthorization2Request struct {
	ctx         context.Context
	APIService  ServiceAuthorizationsAPI
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIDeleteServiceAuthorization2Request) RequestBody(requestBody map[string]map[string]any) *APIDeleteServiceAuthorization2Request {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIDeleteServiceAuthorization2Request) Execute() (*InlineResponse2007, *http.Response, error) {
	return r.APIService.DeleteServiceAuthorization2Execute(r)
}

/*
DeleteServiceAuthorization2 Delete service authorizations

Delete service authorizations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIDeleteServiceAuthorization2Request
*/
func (a *ServiceAuthorizationsAPIService) DeleteServiceAuthorization2(ctx context.Context) APIDeleteServiceAuthorization2Request {
	return APIDeleteServiceAuthorization2Request{
		APIService: a,
		ctx:        ctx,
	}
}

// DeleteServiceAuthorization2Execute executes the request
//  @return InlineResponse2007
func (a *ServiceAuthorizationsAPIService) DeleteServiceAuthorization2Execute(r APIDeleteServiceAuthorization2Request) (*InlineResponse2007, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse2007
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.DeleteServiceAuthorization2")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json; ext=bulk"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json; ext=bulk"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
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

// APIListServiceAuthorizationRequest represents a request for the resource.
type APIListServiceAuthorizationRequest struct {
	ctx        context.Context
	APIService ServiceAuthorizationsAPI
	pageNumber *int32
	pageSize   *int32
}

// PageNumber Current page.
func (r *APIListServiceAuthorizationRequest) PageNumber(pageNumber int32) *APIListServiceAuthorizationRequest {
	r.pageNumber = &pageNumber
	return r
}

// PageSize Number of records per page.
func (r *APIListServiceAuthorizationRequest) PageSize(pageSize int32) *APIListServiceAuthorizationRequest {
	r.pageSize = &pageSize
	return r
}

// Execute calls the API using the request data configured.
func (r APIListServiceAuthorizationRequest) Execute() (*ServiceAuthorizationsResponse, *http.Response, error) {
	return r.APIService.ListServiceAuthorizationExecute(r)
}

/*
ListServiceAuthorization List service authorizations

List service authorizations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListServiceAuthorizationRequest
*/
func (a *ServiceAuthorizationsAPIService) ListServiceAuthorization(ctx context.Context) APIListServiceAuthorizationRequest {
	return APIListServiceAuthorizationRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListServiceAuthorizationExecute executes the request
//  @return ServiceAuthorizationsResponse
func (a *ServiceAuthorizationsAPIService) ListServiceAuthorizationExecute(r APIListServiceAuthorizationRequest) (*ServiceAuthorizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAuthorizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.ListServiceAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

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

// APIShowServiceAuthorizationRequest represents a request for the resource.
type APIShowServiceAuthorizationRequest struct {
	ctx                    context.Context
	APIService             ServiceAuthorizationsAPI
	serviceAuthorizationID string
}

// Execute calls the API using the request data configured.
func (r APIShowServiceAuthorizationRequest) Execute() (*ServiceAuthorizationResponse, *http.Response, error) {
	return r.APIService.ShowServiceAuthorizationExecute(r)
}

/*
ShowServiceAuthorization Show service authorization

Show service authorization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceAuthorizationID Alphanumeric string identifying a service authorization.
 @return APIShowServiceAuthorizationRequest
*/
func (a *ServiceAuthorizationsAPIService) ShowServiceAuthorization(ctx context.Context, serviceAuthorizationID string) APIShowServiceAuthorizationRequest {
	return APIShowServiceAuthorizationRequest{
		APIService:             a,
		ctx:                    ctx,
		serviceAuthorizationID: serviceAuthorizationID,
	}
}

// ShowServiceAuthorizationExecute executes the request
//  @return ServiceAuthorizationResponse
func (a *ServiceAuthorizationsAPIService) ShowServiceAuthorizationExecute(r APIShowServiceAuthorizationRequest) (*ServiceAuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.ShowServiceAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations/{service_authorization_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_authorization_id"+"}", gourl.PathEscape(parameterToString(r.serviceAuthorizationID, "")))

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

// APIUpdateServiceAuthorizationRequest represents a request for the resource.
type APIUpdateServiceAuthorizationRequest struct {
	ctx                    context.Context
	APIService             ServiceAuthorizationsAPI
	serviceAuthorizationID string
	serviceAuthorization   *ServiceAuthorization
}

// ServiceAuthorization returns a pointer to a request.
func (r *APIUpdateServiceAuthorizationRequest) ServiceAuthorization(serviceAuthorization ServiceAuthorization) *APIUpdateServiceAuthorizationRequest {
	r.serviceAuthorization = &serviceAuthorization
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateServiceAuthorizationRequest) Execute() (*ServiceAuthorizationResponse, *http.Response, error) {
	return r.APIService.UpdateServiceAuthorizationExecute(r)
}

/*
UpdateServiceAuthorization Update service authorization

Update service authorization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceAuthorizationID Alphanumeric string identifying a service authorization.
 @return APIUpdateServiceAuthorizationRequest
*/
func (a *ServiceAuthorizationsAPIService) UpdateServiceAuthorization(ctx context.Context, serviceAuthorizationID string) APIUpdateServiceAuthorizationRequest {
	return APIUpdateServiceAuthorizationRequest{
		APIService:             a,
		ctx:                    ctx,
		serviceAuthorizationID: serviceAuthorizationID,
	}
}

// UpdateServiceAuthorizationExecute executes the request
//  @return ServiceAuthorizationResponse
func (a *ServiceAuthorizationsAPIService) UpdateServiceAuthorizationExecute(r APIUpdateServiceAuthorizationRequest) (*ServiceAuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.UpdateServiceAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations/{service_authorization_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_authorization_id"+"}", gourl.PathEscape(parameterToString(r.serviceAuthorizationID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.serviceAuthorization
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

// APIUpdateServiceAuthorization2Request represents a request for the resource.
type APIUpdateServiceAuthorization2Request struct {
	ctx         context.Context
	APIService  ServiceAuthorizationsAPI
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APIUpdateServiceAuthorization2Request) RequestBody(requestBody map[string]map[string]any) *APIUpdateServiceAuthorization2Request {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateServiceAuthorization2Request) Execute() (*ServiceAuthorizationsResponse, *http.Response, error) {
	return r.APIService.UpdateServiceAuthorization2Execute(r)
}

/*
UpdateServiceAuthorization2 Update service authorizations

Update service authorizations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIUpdateServiceAuthorization2Request
*/
func (a *ServiceAuthorizationsAPIService) UpdateServiceAuthorization2(ctx context.Context) APIUpdateServiceAuthorization2Request {
	return APIUpdateServiceAuthorization2Request{
		APIService: a,
		ctx:        ctx,
	}
}

// UpdateServiceAuthorization2Execute executes the request
//  @return ServiceAuthorizationsResponse
func (a *ServiceAuthorizationsAPIService) UpdateServiceAuthorization2Execute(r APIUpdateServiceAuthorization2Request) (*ServiceAuthorizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAuthorizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAuthorizationsAPIService.UpdateServiceAuthorization2")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-authorizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json; ext=bulk"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json; ext=bulk"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
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
