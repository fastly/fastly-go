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
	"time"
)

// Linger please
var (
	_ context.Context
)

// HTTP3API defines an interface for interacting with the resource.
type HTTP3API interface {

	/*
	CreateHTTP3 Enable support for HTTP/3

	Enable HTTP/3 (QUIC) support for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateHTTP3Request
	*/
	CreateHTTP3(ctx context.Context, serviceID string, versionID int32) APICreateHTTP3Request

	// CreateHTTP3Execute executes the request
	//  @return HTTP3
	CreateHTTP3Execute(r APICreateHTTP3Request) (*HTTP3, *http.Response, error)

	/*
	DeleteHTTP3 Disable support for HTTP/3

	Disable HTTP/3 (QUIC) support for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIDeleteHTTP3Request
	*/
	DeleteHTTP3(ctx context.Context, serviceID string, versionID int32) APIDeleteHTTP3Request

	// DeleteHTTP3Execute executes the request
	//  @return InlineResponse200
	DeleteHTTP3Execute(r APIDeleteHTTP3Request) (*InlineResponse200, *http.Response, error)

	/*
	GetHTTP3 Get HTTP/3 status

	Get the status of HTTP/3 (QUIC) support for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIGetHTTP3Request
	*/
	GetHTTP3(ctx context.Context, serviceID string, versionID int32) APIGetHTTP3Request

	// GetHTTP3Execute executes the request
	//  @return HTTP3
	GetHTTP3Execute(r APIGetHTTP3Request) (*HTTP3, *http.Response, error)
}

// HTTP3APIService HTTP3API service
type HTTP3APIService service

// APICreateHTTP3Request represents a request for the resource.
type APICreateHTTP3Request struct {
	ctx context.Context
	APIService HTTP3API
	serviceID string
	versionID int32
	serviceID2 *string
	version *int32
	createdAt *time.Time
	deletedAt *time.Time
	updatedAt *time.Time
	featureRevision *int32
}

// ServiceID2 returns a pointer to a request.
func (r *APICreateHTTP3Request) ServiceID2(serviceID2 string) *APICreateHTTP3Request {
	r.serviceID2 = &serviceID2
	return r
}
// Version returns a pointer to a request.
func (r *APICreateHTTP3Request) Version(version int32) *APICreateHTTP3Request {
	r.version = &version
	return r
}
// CreatedAt Date and time in ISO 8601 format.
func (r *APICreateHTTP3Request) CreatedAt(createdAt time.Time) *APICreateHTTP3Request {
	r.createdAt = &createdAt
	return r
}
// DeletedAt Date and time in ISO 8601 format.
func (r *APICreateHTTP3Request) DeletedAt(deletedAt time.Time) *APICreateHTTP3Request {
	r.deletedAt = &deletedAt
	return r
}
// UpdatedAt Date and time in ISO 8601 format.
func (r *APICreateHTTP3Request) UpdatedAt(updatedAt time.Time) *APICreateHTTP3Request {
	r.updatedAt = &updatedAt
	return r
}
// FeatureRevision Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision.
func (r *APICreateHTTP3Request) FeatureRevision(featureRevision int32) *APICreateHTTP3Request {
	r.featureRevision = &featureRevision
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateHTTP3Request) Execute() (*HTTP3, *http.Response, error) {
	return r.APIService.CreateHTTP3Execute(r)
}

/*
CreateHTTP3 Enable support for HTTP/3

Enable HTTP/3 (QUIC) support for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateHTTP3Request
*/
func (a *HTTP3APIService) CreateHTTP3(ctx context.Context, serviceID string, versionID int32) APICreateHTTP3Request {
	return APICreateHTTP3Request{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateHTTP3Execute executes the request
//  @return HTTP3
func (a *HTTP3APIService) CreateHTTP3Execute(r APICreateHTTP3Request) (*HTTP3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *HTTP3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HTTP3APIService.CreateHTTP3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/http3"
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
	if r.serviceID2 != nil {
		paramJSON, err := parameterToJSON(*r.serviceID2)
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

// APIDeleteHTTP3Request represents a request for the resource.
type APIDeleteHTTP3Request struct {
	ctx context.Context
	APIService HTTP3API
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIDeleteHTTP3Request) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteHTTP3Execute(r)
}

/*
DeleteHTTP3 Disable support for HTTP/3

Disable HTTP/3 (QUIC) support for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIDeleteHTTP3Request
*/
func (a *HTTP3APIService) DeleteHTTP3(ctx context.Context, serviceID string, versionID int32) APIDeleteHTTP3Request {
	return APIDeleteHTTP3Request{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// DeleteHTTP3Execute executes the request
//  @return InlineResponse200
func (a *HTTP3APIService) DeleteHTTP3Execute(r APIDeleteHTTP3Request) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HTTP3APIService.DeleteHTTP3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/http3"
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

// APIGetHTTP3Request represents a request for the resource.
type APIGetHTTP3Request struct {
	ctx context.Context
	APIService HTTP3API
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIGetHTTP3Request) Execute() (*HTTP3, *http.Response, error) {
	return r.APIService.GetHTTP3Execute(r)
}

/*
GetHTTP3 Get HTTP/3 status

Get the status of HTTP/3 (QUIC) support for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIGetHTTP3Request
*/
func (a *HTTP3APIService) GetHTTP3(ctx context.Context, serviceID string, versionID int32) APIGetHTTP3Request {
	return APIGetHTTP3Request{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// GetHTTP3Execute executes the request
//  @return HTTP3
func (a *HTTP3APIService) GetHTTP3Execute(r APIGetHTTP3Request) (*HTTP3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *HTTP3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HTTP3APIService.GetHTTP3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/http3"
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
