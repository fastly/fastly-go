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

// GzipAPI defines an interface for interacting with the resource.
type GzipAPI interface {

	/*
	CreateGzipConfig Create a gzip configuration

	Create a named gzip configuration on a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateGzipConfigRequest
	*/
	CreateGzipConfig(ctx context.Context, serviceID string, versionID int32) APICreateGzipConfigRequest

	// CreateGzipConfigExecute executes the request
	//  @return GzipResponse
	CreateGzipConfigExecute(r APICreateGzipConfigRequest) (*GzipResponse, *http.Response, error)

	/*
	DeleteGzipConfig Delete a gzip configuration

	Delete a named gzip configuration on a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param gzipName Name of the gzip configuration.
	 @return APIDeleteGzipConfigRequest
	*/
	DeleteGzipConfig(ctx context.Context, serviceID string, versionID int32, gzipName string) APIDeleteGzipConfigRequest

	// DeleteGzipConfigExecute executes the request
	//  @return InlineResponse200
	DeleteGzipConfigExecute(r APIDeleteGzipConfigRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetGzipConfigs Get a gzip configuration

	Get the gzip configuration for a particular service, version, and name.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param gzipName Name of the gzip configuration.
	 @return APIGetGzipConfigsRequest
	*/
	GetGzipConfigs(ctx context.Context, serviceID string, versionID int32, gzipName string) APIGetGzipConfigsRequest

	// GetGzipConfigsExecute executes the request
	//  @return GzipResponse
	GetGzipConfigsExecute(r APIGetGzipConfigsRequest) (*GzipResponse, *http.Response, error)

	/*
	ListGzipConfigs List gzip configurations

	List all gzip configurations for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListGzipConfigsRequest
	*/
	ListGzipConfigs(ctx context.Context, serviceID string, versionID int32) APIListGzipConfigsRequest

	// ListGzipConfigsExecute executes the request
	//  @return []GzipResponse
	ListGzipConfigsExecute(r APIListGzipConfigsRequest) ([]GzipResponse, *http.Response, error)

	/*
	UpdateGzipConfig Update a gzip configuration

	Update a named gzip configuration on a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param gzipName Name of the gzip configuration.
	 @return APIUpdateGzipConfigRequest
	*/
	UpdateGzipConfig(ctx context.Context, serviceID string, versionID int32, gzipName string) APIUpdateGzipConfigRequest

	// UpdateGzipConfigExecute executes the request
	//  @return GzipResponse
	UpdateGzipConfigExecute(r APIUpdateGzipConfigRequest) (*GzipResponse, *http.Response, error)
}

// GzipAPIService GzipAPI service
type GzipAPIService service

// APICreateGzipConfigRequest represents a request for the resource.
type APICreateGzipConfigRequest struct {
	ctx context.Context
	APIService GzipAPI
	serviceID string
	versionID int32
	cacheCondition *string
	contentTypes *string
	extensions *string
	name *string
}

// CacheCondition Name of the cache condition controlling when this configuration applies.
func (r *APICreateGzipConfigRequest) CacheCondition(cacheCondition string) *APICreateGzipConfigRequest {
	r.cacheCondition = &cacheCondition
	return r
}
// ContentTypes Space-separated list of content types to compress. If you omit this field a default list will be used.
func (r *APICreateGzipConfigRequest) ContentTypes(contentTypes string) *APICreateGzipConfigRequest {
	r.contentTypes = &contentTypes
	return r
}
// Extensions Space-separated list of file extensions to compress. If you omit this field a default list will be used.
func (r *APICreateGzipConfigRequest) Extensions(extensions string) *APICreateGzipConfigRequest {
	r.extensions = &extensions
	return r
}
// Name Name of the gzip configuration.
func (r *APICreateGzipConfigRequest) Name(name string) *APICreateGzipConfigRequest {
	r.name = &name
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateGzipConfigRequest) Execute() (*GzipResponse, *http.Response, error) {
	return r.APIService.CreateGzipConfigExecute(r)
}

/*
CreateGzipConfig Create a gzip configuration

Create a named gzip configuration on a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateGzipConfigRequest
*/
func (a *GzipAPIService) CreateGzipConfig(ctx context.Context, serviceID string, versionID int32) APICreateGzipConfigRequest {
	return APICreateGzipConfigRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateGzipConfigExecute executes the request
//  @return GzipResponse
func (a *GzipAPIService) CreateGzipConfigExecute(r APICreateGzipConfigRequest) (*GzipResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *GzipResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GzipAPIService.CreateGzipConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/gzip"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))

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
	if r.cacheCondition != nil {
		localVarFormParams.Add("cache_condition", parameterToString(*r.cacheCondition, ""))
	}
	if r.contentTypes != nil {
		localVarFormParams.Add("content_types", parameterToString(*r.contentTypes, ""))
	}
	if r.extensions != nil {
		localVarFormParams.Add("extensions", parameterToString(*r.extensions, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
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

// APIDeleteGzipConfigRequest represents a request for the resource.
type APIDeleteGzipConfigRequest struct {
	ctx context.Context
	APIService GzipAPI
	serviceID string
	versionID int32
	gzipName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteGzipConfigRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteGzipConfigExecute(r)
}

/*
DeleteGzipConfig Delete a gzip configuration

Delete a named gzip configuration on a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param gzipName Name of the gzip configuration.
 @return APIDeleteGzipConfigRequest
*/
func (a *GzipAPIService) DeleteGzipConfig(ctx context.Context, serviceID string, versionID int32, gzipName string) APIDeleteGzipConfigRequest {
	return APIDeleteGzipConfigRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		gzipName: gzipName,
	}
}

// DeleteGzipConfigExecute executes the request
//  @return InlineResponse200
func (a *GzipAPIService) DeleteGzipConfigExecute(r APIDeleteGzipConfigRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GzipAPIService.DeleteGzipConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/gzip/{gzip_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"gzip_name"+"}", gourl.PathEscape(parameterToString(r.gzipName, "")))

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

// APIGetGzipConfigsRequest represents a request for the resource.
type APIGetGzipConfigsRequest struct {
	ctx context.Context
	APIService GzipAPI
	serviceID string
	versionID int32
	gzipName string
}


// Execute calls the API using the request data configured.
func (r APIGetGzipConfigsRequest) Execute() (*GzipResponse, *http.Response, error) {
	return r.APIService.GetGzipConfigsExecute(r)
}

/*
GetGzipConfigs Get a gzip configuration

Get the gzip configuration for a particular service, version, and name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param gzipName Name of the gzip configuration.
 @return APIGetGzipConfigsRequest
*/
func (a *GzipAPIService) GetGzipConfigs(ctx context.Context, serviceID string, versionID int32, gzipName string) APIGetGzipConfigsRequest {
	return APIGetGzipConfigsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		gzipName: gzipName,
	}
}

// GetGzipConfigsExecute executes the request
//  @return GzipResponse
func (a *GzipAPIService) GetGzipConfigsExecute(r APIGetGzipConfigsRequest) (*GzipResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *GzipResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GzipAPIService.GetGzipConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/gzip/{gzip_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"gzip_name"+"}", gourl.PathEscape(parameterToString(r.gzipName, "")))

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

// APIListGzipConfigsRequest represents a request for the resource.
type APIListGzipConfigsRequest struct {
	ctx context.Context
	APIService GzipAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListGzipConfigsRequest) Execute() ([]GzipResponse, *http.Response, error) {
	return r.APIService.ListGzipConfigsExecute(r)
}

/*
ListGzipConfigs List gzip configurations

List all gzip configurations for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListGzipConfigsRequest
*/
func (a *GzipAPIService) ListGzipConfigs(ctx context.Context, serviceID string, versionID int32) APIListGzipConfigsRequest {
	return APIListGzipConfigsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListGzipConfigsExecute executes the request
//  @return []GzipResponse
func (a *GzipAPIService) ListGzipConfigsExecute(r APIListGzipConfigsRequest) ([]GzipResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []GzipResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GzipAPIService.ListGzipConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/gzip"
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

// APIUpdateGzipConfigRequest represents a request for the resource.
type APIUpdateGzipConfigRequest struct {
	ctx context.Context
	APIService GzipAPI
	serviceID string
	versionID int32
	gzipName string
	cacheCondition *string
	contentTypes *string
	extensions *string
	name *string
}

// CacheCondition Name of the cache condition controlling when this configuration applies.
func (r *APIUpdateGzipConfigRequest) CacheCondition(cacheCondition string) *APIUpdateGzipConfigRequest {
	r.cacheCondition = &cacheCondition
	return r
}
// ContentTypes Space-separated list of content types to compress. If you omit this field a default list will be used.
func (r *APIUpdateGzipConfigRequest) ContentTypes(contentTypes string) *APIUpdateGzipConfigRequest {
	r.contentTypes = &contentTypes
	return r
}
// Extensions Space-separated list of file extensions to compress. If you omit this field a default list will be used.
func (r *APIUpdateGzipConfigRequest) Extensions(extensions string) *APIUpdateGzipConfigRequest {
	r.extensions = &extensions
	return r
}
// Name Name of the gzip configuration.
func (r *APIUpdateGzipConfigRequest) Name(name string) *APIUpdateGzipConfigRequest {
	r.name = &name
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateGzipConfigRequest) Execute() (*GzipResponse, *http.Response, error) {
	return r.APIService.UpdateGzipConfigExecute(r)
}

/*
UpdateGzipConfig Update a gzip configuration

Update a named gzip configuration on a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param gzipName Name of the gzip configuration.
 @return APIUpdateGzipConfigRequest
*/
func (a *GzipAPIService) UpdateGzipConfig(ctx context.Context, serviceID string, versionID int32, gzipName string) APIUpdateGzipConfigRequest {
	return APIUpdateGzipConfigRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		gzipName: gzipName,
	}
}

// UpdateGzipConfigExecute executes the request
//  @return GzipResponse
func (a *GzipAPIService) UpdateGzipConfigExecute(r APIUpdateGzipConfigRequest) (*GzipResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *GzipResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GzipAPIService.UpdateGzipConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/gzip/{gzip_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"gzip_name"+"}", gourl.PathEscape(parameterToString(r.gzipName, "")))

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
	if r.cacheCondition != nil {
		localVarFormParams.Add("cache_condition", parameterToString(*r.cacheCondition, ""))
	}
	if r.contentTypes != nil {
		localVarFormParams.Add("content_types", parameterToString(*r.contentTypes, ""))
	}
	if r.extensions != nil {
		localVarFormParams.Add("extensions", parameterToString(*r.extensions, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
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
