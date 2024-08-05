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

// ObservabilityCustomDashboardsAPI defines an interface for interacting with the resource.
type ObservabilityCustomDashboardsAPI interface {

	/*
	CreateDashboard Create a new dashboard

	Create a new dashboard

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateDashboardRequest
	*/
	CreateDashboard(ctx context.Context) APICreateDashboardRequest

	// CreateDashboardExecute executes the request
	//  @return Dashboard
	CreateDashboardExecute(r APICreateDashboardRequest) (*Dashboard, *http.Response, error)

	/*
	DeleteDashboard Delete an existing dashboard

	Delete an existing dashboard

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param dashboardID
	 @return APIDeleteDashboardRequest
	*/
	DeleteDashboard(ctx context.Context, dashboardID string) APIDeleteDashboardRequest

	// DeleteDashboardExecute executes the request
	DeleteDashboardExecute(r APIDeleteDashboardRequest) (*http.Response, error)

	/*
	GetDashboard Retrieve a dashboard by ID

	Retrieve a dashboard by ID

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param dashboardID
	 @return APIGetDashboardRequest
	*/
	GetDashboard(ctx context.Context, dashboardID string) APIGetDashboardRequest

	// GetDashboardExecute executes the request
	//  @return Dashboard
	GetDashboardExecute(r APIGetDashboardRequest) (*Dashboard, *http.Response, error)

	/*
	ListDashboards List all custom dashboards

	List all custom dashboards

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListDashboardsRequest
	*/
	ListDashboards(ctx context.Context) APIListDashboardsRequest

	// ListDashboardsExecute executes the request
	//  @return ListDashboardsResponse
	ListDashboardsExecute(r APIListDashboardsRequest) (*ListDashboardsResponse, *http.Response, error)

	/*
	UpdateDashboard Update an existing dashboard

	Update an existing dashboard

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param dashboardID
	 @return APIUpdateDashboardRequest
	*/
	UpdateDashboard(ctx context.Context, dashboardID string) APIUpdateDashboardRequest

	// UpdateDashboardExecute executes the request
	//  @return Dashboard
	UpdateDashboardExecute(r APIUpdateDashboardRequest) (*Dashboard, *http.Response, error)
}

// ObservabilityCustomDashboardsAPIService ObservabilityCustomDashboardsAPI service
type ObservabilityCustomDashboardsAPIService service

// APICreateDashboardRequest represents a request for the resource.
type APICreateDashboardRequest struct {
	ctx context.Context
	APIService ObservabilityCustomDashboardsAPI
	createDashboardRequest *CreateDashboardRequest
}

// CreateDashboardRequest returns a pointer to a request.
func (r *APICreateDashboardRequest) CreateDashboardRequest(createDashboardRequest CreateDashboardRequest) *APICreateDashboardRequest {
	r.createDashboardRequest = &createDashboardRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateDashboardRequest) Execute() (*Dashboard, *http.Response, error) {
	return r.APIService.CreateDashboardExecute(r)
}

/*
CreateDashboard Create a new dashboard

Create a new dashboard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateDashboardRequest
*/
func (a *ObservabilityCustomDashboardsAPIService) CreateDashboard(ctx context.Context) APICreateDashboardRequest {
	return APICreateDashboardRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateDashboardExecute executes the request
//  @return Dashboard
func (a *ObservabilityCustomDashboardsAPIService) CreateDashboardExecute(r APICreateDashboardRequest) (*Dashboard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Dashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObservabilityCustomDashboardsAPIService.CreateDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/dashboards"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createDashboardRequest
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

// APIDeleteDashboardRequest represents a request for the resource.
type APIDeleteDashboardRequest struct {
	ctx context.Context
	APIService ObservabilityCustomDashboardsAPI
	dashboardID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteDashboardRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteDashboardExecute(r)
}

/*
DeleteDashboard Delete an existing dashboard

Delete an existing dashboard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardID
 @return APIDeleteDashboardRequest
*/
func (a *ObservabilityCustomDashboardsAPIService) DeleteDashboard(ctx context.Context, dashboardID string) APIDeleteDashboardRequest {
	return APIDeleteDashboardRequest{
		APIService: a,
		ctx: ctx,
		dashboardID: dashboardID,
	}
}

// DeleteDashboardExecute executes the request
func (a *ObservabilityCustomDashboardsAPIService) DeleteDashboardExecute(r APIDeleteDashboardRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObservabilityCustomDashboardsAPIService.DeleteDashboard")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/dashboards/{dashboard_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dashboard_id"+"}", gourl.PathEscape(parameterToString(r.dashboardID, "")))

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

// APIGetDashboardRequest represents a request for the resource.
type APIGetDashboardRequest struct {
	ctx context.Context
	APIService ObservabilityCustomDashboardsAPI
	dashboardID string
}


// Execute calls the API using the request data configured.
func (r APIGetDashboardRequest) Execute() (*Dashboard, *http.Response, error) {
	return r.APIService.GetDashboardExecute(r)
}

/*
GetDashboard Retrieve a dashboard by ID

Retrieve a dashboard by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardID
 @return APIGetDashboardRequest
*/
func (a *ObservabilityCustomDashboardsAPIService) GetDashboard(ctx context.Context, dashboardID string) APIGetDashboardRequest {
	return APIGetDashboardRequest{
		APIService: a,
		ctx: ctx,
		dashboardID: dashboardID,
	}
}

// GetDashboardExecute executes the request
//  @return Dashboard
func (a *ObservabilityCustomDashboardsAPIService) GetDashboardExecute(r APIGetDashboardRequest) (*Dashboard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Dashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObservabilityCustomDashboardsAPIService.GetDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/dashboards/{dashboard_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dashboard_id"+"}", gourl.PathEscape(parameterToString(r.dashboardID, "")))

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

// APIListDashboardsRequest represents a request for the resource.
type APIListDashboardsRequest struct {
	ctx context.Context
	APIService ObservabilityCustomDashboardsAPI
}


// Execute calls the API using the request data configured.
func (r APIListDashboardsRequest) Execute() (*ListDashboardsResponse, *http.Response, error) {
	return r.APIService.ListDashboardsExecute(r)
}

/*
ListDashboards List all custom dashboards

List all custom dashboards

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListDashboardsRequest
*/
func (a *ObservabilityCustomDashboardsAPIService) ListDashboards(ctx context.Context) APIListDashboardsRequest {
	return APIListDashboardsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListDashboardsExecute executes the request
//  @return ListDashboardsResponse
func (a *ObservabilityCustomDashboardsAPIService) ListDashboardsExecute(r APIListDashboardsRequest) (*ListDashboardsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *ListDashboardsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObservabilityCustomDashboardsAPIService.ListDashboards")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/dashboards"

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

// APIUpdateDashboardRequest represents a request for the resource.
type APIUpdateDashboardRequest struct {
	ctx context.Context
	APIService ObservabilityCustomDashboardsAPI
	dashboardID string
	updateDashboardRequest *UpdateDashboardRequest
}

// UpdateDashboardRequest returns a pointer to a request.
func (r *APIUpdateDashboardRequest) UpdateDashboardRequest(updateDashboardRequest UpdateDashboardRequest) *APIUpdateDashboardRequest {
	r.updateDashboardRequest = &updateDashboardRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateDashboardRequest) Execute() (*Dashboard, *http.Response, error) {
	return r.APIService.UpdateDashboardExecute(r)
}

/*
UpdateDashboard Update an existing dashboard

Update an existing dashboard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardID
 @return APIUpdateDashboardRequest
*/
func (a *ObservabilityCustomDashboardsAPIService) UpdateDashboard(ctx context.Context, dashboardID string) APIUpdateDashboardRequest {
	return APIUpdateDashboardRequest{
		APIService: a,
		ctx: ctx,
		dashboardID: dashboardID,
	}
}

// UpdateDashboardExecute executes the request
//  @return Dashboard
func (a *ObservabilityCustomDashboardsAPIService) UpdateDashboardExecute(r APIUpdateDashboardRequest) (*Dashboard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Dashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObservabilityCustomDashboardsAPIService.UpdateDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/dashboards/{dashboard_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dashboard_id"+"}", gourl.PathEscape(parameterToString(r.dashboardID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateDashboardRequest
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
