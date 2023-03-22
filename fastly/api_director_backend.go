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
)

// Linger please
var (
	_ context.Context
)

// DirectorBackendAPI defines an interface for interacting with the resource.
type DirectorBackendAPI interface {

	/*
	CreateDirectorBackend Create a director-backend relationship

	Establishes a relationship between a Backend and a Director. The Backend is then considered a member of the Director and can be used to balance traffic onto.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param directorName Name for the Director.
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param backendName The name of the backend.
	 @return APICreateDirectorBackendRequest
	*/
	CreateDirectorBackend(ctx context.Context, directorName string, serviceID string, versionID int32, backendName string) APICreateDirectorBackendRequest

	// CreateDirectorBackendExecute executes the request
	//  @return DirectorBackend
	CreateDirectorBackendExecute(r APICreateDirectorBackendRequest) (*DirectorBackend, *http.Response, error)

	/*
	DeleteDirectorBackend Delete a director-backend relationship

	Deletes the relationship between a Backend and a Director. The Backend is no longer considered a member of the Director and thus will not have traffic balanced onto it from this Director.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param directorName Name for the Director.
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param backendName The name of the backend.
	 @return APIDeleteDirectorBackendRequest
	*/
	DeleteDirectorBackend(ctx context.Context, directorName string, serviceID string, versionID int32, backendName string) APIDeleteDirectorBackendRequest

	// DeleteDirectorBackendExecute executes the request
	//  @return InlineResponse200
	DeleteDirectorBackendExecute(r APIDeleteDirectorBackendRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetDirectorBackend Get a director-backend relationship

	Returns the relationship between a Backend and a Director. If the Backend has been associated with the Director, it returns a simple record indicating this. Otherwise, returns a 404.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param directorName Name for the Director.
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param backendName The name of the backend.
	 @return APIGetDirectorBackendRequest
	*/
	GetDirectorBackend(ctx context.Context, directorName string, serviceID string, versionID int32, backendName string) APIGetDirectorBackendRequest

	// GetDirectorBackendExecute executes the request
	//  @return DirectorBackend
	GetDirectorBackendExecute(r APIGetDirectorBackendRequest) (*DirectorBackend, *http.Response, error)
}

// DirectorBackendAPIService DirectorBackendAPI service
type DirectorBackendAPIService service

// APICreateDirectorBackendRequest represents a request for the resource.
type APICreateDirectorBackendRequest struct {
	ctx context.Context
	APIService DirectorBackendAPI
	directorName string
	serviceID string
	versionID int32
	backendName string
}


// Execute calls the API using the request data configured.
func (r APICreateDirectorBackendRequest) Execute() (*DirectorBackend, *http.Response, error) {
	return r.APIService.CreateDirectorBackendExecute(r)
}

/*
CreateDirectorBackend Create a director-backend relationship

Establishes a relationship between a Backend and a Director. The Backend is then considered a member of the Director and can be used to balance traffic onto.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directorName Name for the Director.
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param backendName The name of the backend.
 @return APICreateDirectorBackendRequest
*/
func (a *DirectorBackendAPIService) CreateDirectorBackend(ctx context.Context, directorName string, serviceID string, versionID int32, backendName string) APICreateDirectorBackendRequest {
	return APICreateDirectorBackendRequest{
		APIService: a,
		ctx: ctx,
		directorName: directorName,
		serviceID: serviceID,
		versionID: versionID,
		backendName: backendName,
	}
}

// CreateDirectorBackendExecute executes the request
//  @return DirectorBackend
func (a *DirectorBackendAPIService) CreateDirectorBackendExecute(r APICreateDirectorBackendRequest) (*DirectorBackend, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *DirectorBackend
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectorBackendAPIService.CreateDirectorBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"director_name"+"}", gourl.PathEscape(parameterToString(r.directorName, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"backend_name"+"}", gourl.PathEscape(parameterToString(r.backendName, "")))

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

// APIDeleteDirectorBackendRequest represents a request for the resource.
type APIDeleteDirectorBackendRequest struct {
	ctx context.Context
	APIService DirectorBackendAPI
	directorName string
	serviceID string
	versionID int32
	backendName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteDirectorBackendRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteDirectorBackendExecute(r)
}

/*
DeleteDirectorBackend Delete a director-backend relationship

Deletes the relationship between a Backend and a Director. The Backend is no longer considered a member of the Director and thus will not have traffic balanced onto it from this Director.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directorName Name for the Director.
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param backendName The name of the backend.
 @return APIDeleteDirectorBackendRequest
*/
func (a *DirectorBackendAPIService) DeleteDirectorBackend(ctx context.Context, directorName string, serviceID string, versionID int32, backendName string) APIDeleteDirectorBackendRequest {
	return APIDeleteDirectorBackendRequest{
		APIService: a,
		ctx: ctx,
		directorName: directorName,
		serviceID: serviceID,
		versionID: versionID,
		backendName: backendName,
	}
}

// DeleteDirectorBackendExecute executes the request
//  @return InlineResponse200
func (a *DirectorBackendAPIService) DeleteDirectorBackendExecute(r APIDeleteDirectorBackendRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectorBackendAPIService.DeleteDirectorBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"director_name"+"}", gourl.PathEscape(parameterToString(r.directorName, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"backend_name"+"}", gourl.PathEscape(parameterToString(r.backendName, "")))

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

// APIGetDirectorBackendRequest represents a request for the resource.
type APIGetDirectorBackendRequest struct {
	ctx context.Context
	APIService DirectorBackendAPI
	directorName string
	serviceID string
	versionID int32
	backendName string
}


// Execute calls the API using the request data configured.
func (r APIGetDirectorBackendRequest) Execute() (*DirectorBackend, *http.Response, error) {
	return r.APIService.GetDirectorBackendExecute(r)
}

/*
GetDirectorBackend Get a director-backend relationship

Returns the relationship between a Backend and a Director. If the Backend has been associated with the Director, it returns a simple record indicating this. Otherwise, returns a 404.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directorName Name for the Director.
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param backendName The name of the backend.
 @return APIGetDirectorBackendRequest
*/
func (a *DirectorBackendAPIService) GetDirectorBackend(ctx context.Context, directorName string, serviceID string, versionID int32, backendName string) APIGetDirectorBackendRequest {
	return APIGetDirectorBackendRequest{
		APIService: a,
		ctx: ctx,
		directorName: directorName,
		serviceID: serviceID,
		versionID: versionID,
		backendName: backendName,
	}
}

// GetDirectorBackendExecute executes the request
//  @return DirectorBackend
func (a *DirectorBackendAPIService) GetDirectorBackendExecute(r APIGetDirectorBackendRequest) (*DirectorBackend, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *DirectorBackend
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectorBackendAPIService.GetDirectorBackend")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"director_name"+"}", gourl.PathEscape(parameterToString(r.directorName, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"backend_name"+"}", gourl.PathEscape(parameterToString(r.backendName, "")))

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
