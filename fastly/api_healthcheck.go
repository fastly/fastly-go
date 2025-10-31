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

// HealthcheckAPI defines an interface for interacting with the resource.
type HealthcheckAPI interface {

	/*
		CreateHealthcheck Create a health check

		Create a health check for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APICreateHealthcheckRequest
	*/
	CreateHealthcheck(ctx context.Context, serviceId string, versionId int32) APICreateHealthcheckRequest

	// CreateHealthcheckExecute executes the request
	//  @return HealthcheckResponse
	CreateHealthcheckExecute(r APICreateHealthcheckRequest) (*HealthcheckResponse, *http.Response, error)

	/*
		DeleteHealthcheck Delete a health check

		Delete the health check for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param healthcheckName The name of the health check.
		 @return APIDeleteHealthcheckRequest
	*/
	DeleteHealthcheck(ctx context.Context, serviceId string, versionId int32, healthcheckName string) APIDeleteHealthcheckRequest

	// DeleteHealthcheckExecute executes the request
	//  @return InlineResponse200
	DeleteHealthcheckExecute(r APIDeleteHealthcheckRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetHealthcheck Get a health check

		Get the health check for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param healthcheckName The name of the health check.
		 @return APIGetHealthcheckRequest
	*/
	GetHealthcheck(ctx context.Context, serviceId string, versionId int32, healthcheckName string) APIGetHealthcheckRequest

	// GetHealthcheckExecute executes the request
	//  @return HealthcheckResponse
	GetHealthcheckExecute(r APIGetHealthcheckRequest) (*HealthcheckResponse, *http.Response, error)

	/*
		ListHealthchecks List health checks

		List all of the health checks for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APIListHealthchecksRequest
	*/
	ListHealthchecks(ctx context.Context, serviceId string, versionId int32) APIListHealthchecksRequest

	// ListHealthchecksExecute executes the request
	//  @return []HealthcheckResponse
	ListHealthchecksExecute(r APIListHealthchecksRequest) ([]HealthcheckResponse, *http.Response, error)

	/*
		UpdateHealthcheck Update a health check

		Update the health check for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param healthcheckName The name of the health check.
		 @return APIUpdateHealthcheckRequest
	*/
	UpdateHealthcheck(ctx context.Context, serviceId string, versionId int32, healthcheckName string) APIUpdateHealthcheckRequest

	// UpdateHealthcheckExecute executes the request
	//  @return HealthcheckResponse
	UpdateHealthcheckExecute(r APIUpdateHealthcheckRequest) (*HealthcheckResponse, *http.Response, error)
}

// HealthcheckAPIService HealthcheckAPI service
type HealthcheckAPIService service

// APICreateHealthcheckRequest represents a request for the resource.
type APICreateHealthcheckRequest struct {
	ctx              context.Context
	APIService       HealthcheckAPI
	serviceId        string
	versionId        int32
	checkInterval    *int32
	comment          *string
	expectedResponse *int32
	headers          *[]string
	host             *string
	httpVersion      *string
	initial          *int32
	method           *string
	name             *string
	path             *string
	threshold        *int32
	timeout          *int32
	window           *int32
}

// CheckInterval How often to run the health check in milliseconds. Minimum 1 second, maximum 1 hour.
func (r *APICreateHealthcheckRequest) CheckInterval(checkInterval int32) *APICreateHealthcheckRequest {
	r.checkInterval = &checkInterval
	return r
}

// Comment A freeform descriptive note.
func (r *APICreateHealthcheckRequest) Comment(comment string) *APICreateHealthcheckRequest {
	r.comment = &comment
	return r
}

// ExpectedResponse The status code expected from the host.
func (r *APICreateHealthcheckRequest) ExpectedResponse(expectedResponse int32) *APICreateHealthcheckRequest {
	r.expectedResponse = &expectedResponse
	return r
}

// Headers Array of custom headers that will be added to the health check probes.
func (r *APICreateHealthcheckRequest) Headers(headers []string) *APICreateHealthcheckRequest {
	r.headers = &headers
	return r
}

// Host Which host to check.
func (r *APICreateHealthcheckRequest) Host(host string) *APICreateHealthcheckRequest {
	r.host = &host
	return r
}

// HttpVersion Whether to use version 1.0 or 1.1 HTTP.
func (r *APICreateHealthcheckRequest) HttpVersion(httpVersion string) *APICreateHealthcheckRequest {
	r.httpVersion = &httpVersion
	return r
}

// Initial When loading a config, the initial number of probes to be seen as OK.
func (r *APICreateHealthcheckRequest) Initial(initial int32) *APICreateHealthcheckRequest {
	r.initial = &initial
	return r
}

// Method Which HTTP method to use.
func (r *APICreateHealthcheckRequest) Method(method string) *APICreateHealthcheckRequest {
	r.method = &method
	return r
}

// Name The name of the health check.
func (r *APICreateHealthcheckRequest) Name(name string) *APICreateHealthcheckRequest {
	r.name = &name
	return r
}

// Path The path to check.
func (r *APICreateHealthcheckRequest) Path(path string) *APICreateHealthcheckRequest {
	r.path = &path
	return r
}

// Threshold How many health checks must succeed to be considered healthy.
func (r *APICreateHealthcheckRequest) Threshold(threshold int32) *APICreateHealthcheckRequest {
	r.threshold = &threshold
	return r
}

// Timeout Timeout in milliseconds.
func (r *APICreateHealthcheckRequest) Timeout(timeout int32) *APICreateHealthcheckRequest {
	r.timeout = &timeout
	return r
}

// Window The number of most recent health check queries to keep for this health check.
func (r *APICreateHealthcheckRequest) Window(window int32) *APICreateHealthcheckRequest {
	r.window = &window
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateHealthcheckRequest) Execute() (*HealthcheckResponse, *http.Response, error) {
	return r.APIService.CreateHealthcheckExecute(r)
}

/*
CreateHealthcheck Create a health check

Create a health check for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APICreateHealthcheckRequest
*/
func (a *HealthcheckAPIService) CreateHealthcheck(ctx context.Context, serviceId string, versionId int32) APICreateHealthcheckRequest {
	return APICreateHealthcheckRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// CreateHealthcheckExecute executes the request
//  @return HealthcheckResponse
func (a *HealthcheckAPIService) CreateHealthcheckExecute(r APICreateHealthcheckRequest) (*HealthcheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HealthcheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HealthcheckAPIService.CreateHealthcheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/healthcheck"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

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
	if r.checkInterval != nil {
		localVarFormParams.Add("check_interval", parameterToString(*r.checkInterval, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.expectedResponse != nil {
		localVarFormParams.Add("expected_response", parameterToString(*r.expectedResponse, ""))
	}
	if r.headers != nil {
		localVarFormParams.Add("headers", parameterToString(*r.headers, "csv"))
	}
	if r.host != nil {
		localVarFormParams.Add("host", parameterToString(*r.host, ""))
	}
	if r.httpVersion != nil {
		localVarFormParams.Add("http_version", parameterToString(*r.httpVersion, ""))
	}
	if r.initial != nil {
		localVarFormParams.Add("initial", parameterToString(*r.initial, ""))
	}
	if r.method != nil {
		localVarFormParams.Add("method", parameterToString(*r.method, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.threshold != nil {
		localVarFormParams.Add("threshold", parameterToString(*r.threshold, ""))
	}
	if r.timeout != nil {
		localVarFormParams.Add("timeout", parameterToString(*r.timeout, ""))
	}
	if r.window != nil {
		localVarFormParams.Add("window", parameterToString(*r.window, ""))
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

// APIDeleteHealthcheckRequest represents a request for the resource.
type APIDeleteHealthcheckRequest struct {
	ctx             context.Context
	APIService      HealthcheckAPI
	serviceId       string
	versionId       int32
	healthcheckName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteHealthcheckRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteHealthcheckExecute(r)
}

/*
DeleteHealthcheck Delete a health check

Delete the health check for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param healthcheckName The name of the health check.
 @return APIDeleteHealthcheckRequest
*/
func (a *HealthcheckAPIService) DeleteHealthcheck(ctx context.Context, serviceId string, versionId int32, healthcheckName string) APIDeleteHealthcheckRequest {
	return APIDeleteHealthcheckRequest{
		APIService:      a,
		ctx:             ctx,
		serviceId:       serviceId,
		versionId:       versionId,
		healthcheckName: healthcheckName,
	}
}

// DeleteHealthcheckExecute executes the request
//  @return InlineResponse200
func (a *HealthcheckAPIService) DeleteHealthcheckExecute(r APIDeleteHealthcheckRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HealthcheckAPIService.DeleteHealthcheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"healthcheck_name"+"}", gourl.PathEscape(parameterToString(r.healthcheckName, "")))

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

// APIGetHealthcheckRequest represents a request for the resource.
type APIGetHealthcheckRequest struct {
	ctx             context.Context
	APIService      HealthcheckAPI
	serviceId       string
	versionId       int32
	healthcheckName string
}

// Execute calls the API using the request data configured.
func (r APIGetHealthcheckRequest) Execute() (*HealthcheckResponse, *http.Response, error) {
	return r.APIService.GetHealthcheckExecute(r)
}

/*
GetHealthcheck Get a health check

Get the health check for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param healthcheckName The name of the health check.
 @return APIGetHealthcheckRequest
*/
func (a *HealthcheckAPIService) GetHealthcheck(ctx context.Context, serviceId string, versionId int32, healthcheckName string) APIGetHealthcheckRequest {
	return APIGetHealthcheckRequest{
		APIService:      a,
		ctx:             ctx,
		serviceId:       serviceId,
		versionId:       versionId,
		healthcheckName: healthcheckName,
	}
}

// GetHealthcheckExecute executes the request
//  @return HealthcheckResponse
func (a *HealthcheckAPIService) GetHealthcheckExecute(r APIGetHealthcheckRequest) (*HealthcheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HealthcheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HealthcheckAPIService.GetHealthcheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"healthcheck_name"+"}", gourl.PathEscape(parameterToString(r.healthcheckName, "")))

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

// APIListHealthchecksRequest represents a request for the resource.
type APIListHealthchecksRequest struct {
	ctx        context.Context
	APIService HealthcheckAPI
	serviceId  string
	versionId  int32
}

// Execute calls the API using the request data configured.
func (r APIListHealthchecksRequest) Execute() ([]HealthcheckResponse, *http.Response, error) {
	return r.APIService.ListHealthchecksExecute(r)
}

/*
ListHealthchecks List health checks

List all of the health checks for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APIListHealthchecksRequest
*/
func (a *HealthcheckAPIService) ListHealthchecks(ctx context.Context, serviceId string, versionId int32) APIListHealthchecksRequest {
	return APIListHealthchecksRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// ListHealthchecksExecute executes the request
//  @return []HealthcheckResponse
func (a *HealthcheckAPIService) ListHealthchecksExecute(r APIListHealthchecksRequest) ([]HealthcheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []HealthcheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HealthcheckAPIService.ListHealthchecks")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/healthcheck"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

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

// APIUpdateHealthcheckRequest represents a request for the resource.
type APIUpdateHealthcheckRequest struct {
	ctx              context.Context
	APIService       HealthcheckAPI
	serviceId        string
	versionId        int32
	healthcheckName  string
	checkInterval    *int32
	comment          *string
	expectedResponse *int32
	headers          *[]string
	host             *string
	httpVersion      *string
	initial          *int32
	method           *string
	name             *string
	path             *string
	threshold        *int32
	timeout          *int32
	window           *int32
}

// CheckInterval How often to run the health check in milliseconds. Minimum 1 second, maximum 1 hour.
func (r *APIUpdateHealthcheckRequest) CheckInterval(checkInterval int32) *APIUpdateHealthcheckRequest {
	r.checkInterval = &checkInterval
	return r
}

// Comment A freeform descriptive note.
func (r *APIUpdateHealthcheckRequest) Comment(comment string) *APIUpdateHealthcheckRequest {
	r.comment = &comment
	return r
}

// ExpectedResponse The status code expected from the host.
func (r *APIUpdateHealthcheckRequest) ExpectedResponse(expectedResponse int32) *APIUpdateHealthcheckRequest {
	r.expectedResponse = &expectedResponse
	return r
}

// Headers Array of custom headers that will be added to the health check probes.
func (r *APIUpdateHealthcheckRequest) Headers(headers []string) *APIUpdateHealthcheckRequest {
	r.headers = &headers
	return r
}

// Host Which host to check.
func (r *APIUpdateHealthcheckRequest) Host(host string) *APIUpdateHealthcheckRequest {
	r.host = &host
	return r
}

// HttpVersion Whether to use version 1.0 or 1.1 HTTP.
func (r *APIUpdateHealthcheckRequest) HttpVersion(httpVersion string) *APIUpdateHealthcheckRequest {
	r.httpVersion = &httpVersion
	return r
}

// Initial When loading a config, the initial number of probes to be seen as OK.
func (r *APIUpdateHealthcheckRequest) Initial(initial int32) *APIUpdateHealthcheckRequest {
	r.initial = &initial
	return r
}

// Method Which HTTP method to use.
func (r *APIUpdateHealthcheckRequest) Method(method string) *APIUpdateHealthcheckRequest {
	r.method = &method
	return r
}

// Name The name of the health check.
func (r *APIUpdateHealthcheckRequest) Name(name string) *APIUpdateHealthcheckRequest {
	r.name = &name
	return r
}

// Path The path to check.
func (r *APIUpdateHealthcheckRequest) Path(path string) *APIUpdateHealthcheckRequest {
	r.path = &path
	return r
}

// Threshold How many health checks must succeed to be considered healthy.
func (r *APIUpdateHealthcheckRequest) Threshold(threshold int32) *APIUpdateHealthcheckRequest {
	r.threshold = &threshold
	return r
}

// Timeout Timeout in milliseconds.
func (r *APIUpdateHealthcheckRequest) Timeout(timeout int32) *APIUpdateHealthcheckRequest {
	r.timeout = &timeout
	return r
}

// Window The number of most recent health check queries to keep for this health check.
func (r *APIUpdateHealthcheckRequest) Window(window int32) *APIUpdateHealthcheckRequest {
	r.window = &window
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateHealthcheckRequest) Execute() (*HealthcheckResponse, *http.Response, error) {
	return r.APIService.UpdateHealthcheckExecute(r)
}

/*
UpdateHealthcheck Update a health check

Update the health check for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param healthcheckName The name of the health check.
 @return APIUpdateHealthcheckRequest
*/
func (a *HealthcheckAPIService) UpdateHealthcheck(ctx context.Context, serviceId string, versionId int32, healthcheckName string) APIUpdateHealthcheckRequest {
	return APIUpdateHealthcheckRequest{
		APIService:      a,
		ctx:             ctx,
		serviceId:       serviceId,
		versionId:       versionId,
		healthcheckName: healthcheckName,
	}
}

// UpdateHealthcheckExecute executes the request
//  @return HealthcheckResponse
func (a *HealthcheckAPIService) UpdateHealthcheckExecute(r APIUpdateHealthcheckRequest) (*HealthcheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HealthcheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HealthcheckAPIService.UpdateHealthcheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"healthcheck_name"+"}", gourl.PathEscape(parameterToString(r.healthcheckName, "")))

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
	if r.checkInterval != nil {
		localVarFormParams.Add("check_interval", parameterToString(*r.checkInterval, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.expectedResponse != nil {
		localVarFormParams.Add("expected_response", parameterToString(*r.expectedResponse, ""))
	}
	if r.headers != nil {
		localVarFormParams.Add("headers", parameterToString(*r.headers, "csv"))
	}
	if r.host != nil {
		localVarFormParams.Add("host", parameterToString(*r.host, ""))
	}
	if r.httpVersion != nil {
		localVarFormParams.Add("http_version", parameterToString(*r.httpVersion, ""))
	}
	if r.initial != nil {
		localVarFormParams.Add("initial", parameterToString(*r.initial, ""))
	}
	if r.method != nil {
		localVarFormParams.Add("method", parameterToString(*r.method, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.threshold != nil {
		localVarFormParams.Add("threshold", parameterToString(*r.threshold, ""))
	}
	if r.timeout != nil {
		localVarFormParams.Add("timeout", parameterToString(*r.timeout, ""))
	}
	if r.window != nil {
		localVarFormParams.Add("window", parameterToString(*r.window, ""))
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
