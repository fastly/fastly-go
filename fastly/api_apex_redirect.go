// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

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
	"time"
)

// Linger please
var (
	_ context.Context
)

// ApexRedirectAPI defines an interface for interacting with the resource.
type ApexRedirectAPI interface {

	/*
	DeleteApexRedirect Delete an apex redirect

	Delete an apex redirect by its ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param apexRedirectID
	 @return APIDeleteApexRedirectRequest
	*/
	DeleteApexRedirect(ctx context.Context, apexRedirectID string) APIDeleteApexRedirectRequest

	// DeleteApexRedirectExecute executes the request
	//  @return InlineResponse200
	DeleteApexRedirectExecute(r APIDeleteApexRedirectRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetApexRedirect Get an apex redirect

	Get an apex redirect by its ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param apexRedirectID
	 @return APIGetApexRedirectRequest
	*/
	GetApexRedirect(ctx context.Context, apexRedirectID string) APIGetApexRedirectRequest

	// GetApexRedirectExecute executes the request
	//  @return ApexRedirect
	GetApexRedirectExecute(r APIGetApexRedirectRequest) (*ApexRedirect, *http.Response, error)

	/*
	ListApexRedirects List apex redirects

	List all apex redirects for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListApexRedirectsRequest
	*/
	ListApexRedirects(ctx context.Context, serviceID string, versionID int32) APIListApexRedirectsRequest

	// ListApexRedirectsExecute executes the request
	//  @return []ApexRedirect
	ListApexRedirectsExecute(r APIListApexRedirectsRequest) ([]ApexRedirect, *http.Response, error)

	/*
	UpdateApexRedirect Update an apex redirect

	Update an apex redirect by its ID.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param apexRedirectID
	 @return APIUpdateApexRedirectRequest
	*/
	UpdateApexRedirect(ctx context.Context, apexRedirectID string) APIUpdateApexRedirectRequest

	// UpdateApexRedirectExecute executes the request
	//  @return ApexRedirect
	UpdateApexRedirectExecute(r APIUpdateApexRedirectRequest) (*ApexRedirect, *http.Response, error)
}

// ApexRedirectAPIService ApexRedirectAPI service
type ApexRedirectAPIService service

// APIDeleteApexRedirectRequest represents a request for the resource.
type APIDeleteApexRedirectRequest struct {
	ctx context.Context
	APIService ApexRedirectAPI
	apexRedirectID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteApexRedirectRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteApexRedirectExecute(r)
}

/*
DeleteApexRedirect Delete an apex redirect

Delete an apex redirect by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apexRedirectID
 @return APIDeleteApexRedirectRequest
*/
func (a *ApexRedirectAPIService) DeleteApexRedirect(ctx context.Context, apexRedirectID string) APIDeleteApexRedirectRequest {
	return APIDeleteApexRedirectRequest{
		APIService: a,
		ctx: ctx,
		apexRedirectID: apexRedirectID,
	}
}

// DeleteApexRedirectExecute executes the request
//  @return InlineResponse200
func (a *ApexRedirectAPIService) DeleteApexRedirectExecute(r APIDeleteApexRedirectRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApexRedirectAPIService.DeleteApexRedirect")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apex-redirects/{apex_redirect_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"apex_redirect_id"+"}", gourl.PathEscape(parameterToString(r.apexRedirectID, "")))

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

// APIGetApexRedirectRequest represents a request for the resource.
type APIGetApexRedirectRequest struct {
	ctx context.Context
	APIService ApexRedirectAPI
	apexRedirectID string
}


// Execute calls the API using the request data configured.
func (r APIGetApexRedirectRequest) Execute() (*ApexRedirect, *http.Response, error) {
	return r.APIService.GetApexRedirectExecute(r)
}

/*
GetApexRedirect Get an apex redirect

Get an apex redirect by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apexRedirectID
 @return APIGetApexRedirectRequest
*/
func (a *ApexRedirectAPIService) GetApexRedirect(ctx context.Context, apexRedirectID string) APIGetApexRedirectRequest {
	return APIGetApexRedirectRequest{
		APIService: a,
		ctx: ctx,
		apexRedirectID: apexRedirectID,
	}
}

// GetApexRedirectExecute executes the request
//  @return ApexRedirect
func (a *ApexRedirectAPIService) GetApexRedirectExecute(r APIGetApexRedirectRequest) (*ApexRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ApexRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApexRedirectAPIService.GetApexRedirect")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apex-redirects/{apex_redirect_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"apex_redirect_id"+"}", gourl.PathEscape(parameterToString(r.apexRedirectID, "")))

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

// APIListApexRedirectsRequest represents a request for the resource.
type APIListApexRedirectsRequest struct {
	ctx context.Context
	APIService ApexRedirectAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListApexRedirectsRequest) Execute() ([]ApexRedirect, *http.Response, error) {
	return r.APIService.ListApexRedirectsExecute(r)
}

/*
ListApexRedirects List apex redirects

List all apex redirects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListApexRedirectsRequest
*/
func (a *ApexRedirectAPIService) ListApexRedirects(ctx context.Context, serviceID string, versionID int32) APIListApexRedirectsRequest {
	return APIListApexRedirectsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListApexRedirectsExecute executes the request
//  @return []ApexRedirect
func (a *ApexRedirectAPIService) ListApexRedirectsExecute(r APIListApexRedirectsRequest) ([]ApexRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []ApexRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApexRedirectAPIService.ListApexRedirects")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/apex-redirects"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

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

// APIUpdateApexRedirectRequest represents a request for the resource.
type APIUpdateApexRedirectRequest struct {
	ctx context.Context
	APIService ApexRedirectAPI
	apexRedirectID string
	serviceID *string
	version *int32
	createdAt *time.Time
	deletedAt *time.Time
	updatedAt *time.Time
	statusCode *int32
	domains *[]string
	featureRevision *int32
}

// ServiceID returns a pointer to a request.
func (r *APIUpdateApexRedirectRequest) ServiceID(serviceID string) *APIUpdateApexRedirectRequest {
	r.serviceID = &serviceID
	return r
}
// Version returns a pointer to a request.
func (r *APIUpdateApexRedirectRequest) Version(version int32) *APIUpdateApexRedirectRequest {
	r.version = &version
	return r
}
// CreatedAt Date and time in ISO 8601 format.
func (r *APIUpdateApexRedirectRequest) CreatedAt(createdAt time.Time) *APIUpdateApexRedirectRequest {
	r.createdAt = &createdAt
	return r
}
// DeletedAt Date and time in ISO 8601 format.
func (r *APIUpdateApexRedirectRequest) DeletedAt(deletedAt time.Time) *APIUpdateApexRedirectRequest {
	r.deletedAt = &deletedAt
	return r
}
// UpdatedAt Date and time in ISO 8601 format.
func (r *APIUpdateApexRedirectRequest) UpdatedAt(updatedAt time.Time) *APIUpdateApexRedirectRequest {
	r.updatedAt = &updatedAt
	return r
}
// StatusCode HTTP status code used to redirect the client.
func (r *APIUpdateApexRedirectRequest) StatusCode(statusCode int32) *APIUpdateApexRedirectRequest {
	r.statusCode = &statusCode
	return r
}
// Domains Array of apex domains that should redirect to their WWW subdomain.
func (r *APIUpdateApexRedirectRequest) Domains(domains []string) *APIUpdateApexRedirectRequest {
	r.domains = &domains
	return r
}
// FeatureRevision Revision number of the apex redirect feature implementation. Defaults to the most recent revision.
func (r *APIUpdateApexRedirectRequest) FeatureRevision(featureRevision int32) *APIUpdateApexRedirectRequest {
	r.featureRevision = &featureRevision
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateApexRedirectRequest) Execute() (*ApexRedirect, *http.Response, error) {
	return r.APIService.UpdateApexRedirectExecute(r)
}

/*
UpdateApexRedirect Update an apex redirect

Update an apex redirect by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apexRedirectID
 @return APIUpdateApexRedirectRequest
*/
func (a *ApexRedirectAPIService) UpdateApexRedirect(ctx context.Context, apexRedirectID string) APIUpdateApexRedirectRequest {
	return APIUpdateApexRedirectRequest{
		APIService: a,
		ctx: ctx,
		apexRedirectID: apexRedirectID,
	}
}

// UpdateApexRedirectExecute executes the request
//  @return ApexRedirect
func (a *ApexRedirectAPIService) UpdateApexRedirectExecute(r APIUpdateApexRedirectRequest) (*ApexRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ApexRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApexRedirectAPIService.UpdateApexRedirect")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apex-redirects/{apex_redirect_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"apex_redirect_id"+"}", gourl.PathEscape(parameterToString(r.apexRedirectID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.serviceID != nil {
		paramJSON, err := parameterToJSON(*r.serviceID)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("service_id", paramJSON)
	}
	if r.version != nil {
		paramJSON, err := parameterToJSON(*r.version)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("version", paramJSON)
	}
	if r.createdAt != nil {
		localVarFormParams.Add("created_at", parameterToString(*r.createdAt, ""))
	}
	if r.deletedAt != nil {
		localVarFormParams.Add("deleted_at", parameterToString(*r.deletedAt, ""))
	}
	if r.updatedAt != nil {
		localVarFormParams.Add("updated_at", parameterToString(*r.updatedAt, ""))
	}
	if r.statusCode != nil {
		localVarFormParams.Add("status_code", parameterToString(*r.statusCode, ""))
	}
	if r.domains != nil {
		localVarFormParams.Add("domains", parameterToString(*r.domains, "csv"))
	}
	if r.featureRevision != nil {
		localVarFormParams.Add("feature_revision", parameterToString(*r.featureRevision, ""))
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
