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

// LoggingGrafanacloudlogsAPI defines an interface for interacting with the resource.
type LoggingGrafanacloudlogsAPI interface {

	/*
		CreateLogGrafanacloudlogs Create a Grafana Cloud Logs log endpoint

		Create a Grafana Cloud Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogGrafanacloudlogsRequest
	*/
	CreateLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32) APICreateLogGrafanacloudlogsRequest

	// CreateLogGrafanacloudlogsExecute executes the request
	//  @return LoggingGrafanacloudlogsResponse
	CreateLogGrafanacloudlogsExecute(r APICreateLogGrafanacloudlogsRequest) (*LoggingGrafanacloudlogsResponse, *http.Response, error)

	/*
		DeleteLogGrafanacloudlogs Delete the Grafana Cloud Logs log endpoint

		Delete the Grafana Cloud Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingGrafanacloudlogsName The name for the real-time logging configuration.
		 @return APIDeleteLogGrafanacloudlogsRequest
	*/
	DeleteLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32, loggingGrafanacloudlogsName string) APIDeleteLogGrafanacloudlogsRequest

	// DeleteLogGrafanacloudlogsExecute executes the request
	//  @return InlineResponse200
	DeleteLogGrafanacloudlogsExecute(r APIDeleteLogGrafanacloudlogsRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogGrafanacloudlogs Get a Grafana Cloud Logs log endpoint

		Get the details of a Grafana Cloud Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingGrafanacloudlogsName The name for the real-time logging configuration.
		 @return APIGetLogGrafanacloudlogsRequest
	*/
	GetLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32, loggingGrafanacloudlogsName string) APIGetLogGrafanacloudlogsRequest

	// GetLogGrafanacloudlogsExecute executes the request
	//  @return LoggingGrafanacloudlogsResponse
	GetLogGrafanacloudlogsExecute(r APIGetLogGrafanacloudlogsRequest) (*LoggingGrafanacloudlogsResponse, *http.Response, error)

	/*
		ListLogGrafanacloudlogs List Grafana Cloud Logs log endpoints

		List all of the Grafana Cloud Logs logging objects for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogGrafanacloudlogsRequest
	*/
	ListLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32) APIListLogGrafanacloudlogsRequest

	// ListLogGrafanacloudlogsExecute executes the request
	//  @return []LoggingGrafanacloudlogsResponse
	ListLogGrafanacloudlogsExecute(r APIListLogGrafanacloudlogsRequest) ([]LoggingGrafanacloudlogsResponse, *http.Response, error)

	/*
		UpdateLogGrafanacloudlogs Update a Grafana Cloud Logs log endpoint

		Update a Grafana Cloud Logs logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingGrafanacloudlogsName The name for the real-time logging configuration.
		 @return APIUpdateLogGrafanacloudlogsRequest
	*/
	UpdateLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32, loggingGrafanacloudlogsName string) APIUpdateLogGrafanacloudlogsRequest

	// UpdateLogGrafanacloudlogsExecute executes the request
	//  @return LoggingGrafanacloudlogsResponse
	UpdateLogGrafanacloudlogsExecute(r APIUpdateLogGrafanacloudlogsRequest) (*LoggingGrafanacloudlogsResponse, *http.Response, error)
}

// LoggingGrafanacloudlogsAPIService LoggingGrafanacloudlogsAPI service
type LoggingGrafanacloudlogsAPIService service

// APICreateLogGrafanacloudlogsRequest represents a request for the resource.
type APICreateLogGrafanacloudlogsRequest struct {
	ctx               context.Context
	APIService        LoggingGrafanacloudlogsAPI
	serviceID         string
	versionID         int32
	name              *string
	placement         *string
	responseCondition *string
	format            *string
	formatVersion     *int32
	user              *string
	url               *string
	token             *string
	index             *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogGrafanacloudlogsRequest) Name(name string) *APICreateLogGrafanacloudlogsRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogGrafanacloudlogsRequest) Placement(placement string) *APICreateLogGrafanacloudlogsRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogGrafanacloudlogsRequest) ResponseCondition(responseCondition string) *APICreateLogGrafanacloudlogsRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogGrafanacloudlogsRequest) Format(format string) *APICreateLogGrafanacloudlogsRequest {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogGrafanacloudlogsRequest) FormatVersion(formatVersion int32) *APICreateLogGrafanacloudlogsRequest {
	r.formatVersion = &formatVersion
	return r
}

// User The Grafana Cloud Logs Dataset you want to log to.
func (r *APICreateLogGrafanacloudlogsRequest) User(user string) *APICreateLogGrafanacloudlogsRequest {
	r.user = &user
	return r
}

// URL The URL of the Loki instance in your Grafana stack.
func (r *APICreateLogGrafanacloudlogsRequest) URL(url string) *APICreateLogGrafanacloudlogsRequest {
	r.url = &url
	return r
}

// Token The Grafana Access Policy token with &#x60;logs:write&#x60; access scoped to your Loki instance.
func (r *APICreateLogGrafanacloudlogsRequest) Token(token string) *APICreateLogGrafanacloudlogsRequest {
	r.token = &token
	return r
}

// Index The Stream Labels, a JSON string used to identify the stream.
func (r *APICreateLogGrafanacloudlogsRequest) Index(index string) *APICreateLogGrafanacloudlogsRequest {
	r.index = &index
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogGrafanacloudlogsRequest) Execute() (*LoggingGrafanacloudlogsResponse, *http.Response, error) {
	return r.APIService.CreateLogGrafanacloudlogsExecute(r)
}

/*
CreateLogGrafanacloudlogs Create a Grafana Cloud Logs log endpoint

Create a Grafana Cloud Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogGrafanacloudlogsRequest
*/
func (a *LoggingGrafanacloudlogsAPIService) CreateLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32) APICreateLogGrafanacloudlogsRequest {
	return APICreateLogGrafanacloudlogsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogGrafanacloudlogsExecute executes the request
//  @return LoggingGrafanacloudlogsResponse
func (a *LoggingGrafanacloudlogsAPIService) CreateLogGrafanacloudlogsExecute(r APICreateLogGrafanacloudlogsRequest) (*LoggingGrafanacloudlogsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingGrafanacloudlogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGrafanacloudlogsAPIService.CreateLogGrafanacloudlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/grafanacloudlogs"
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
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.placement != nil {
		localVarFormParams.Add("placement", parameterToString(*r.placement, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.index != nil {
		localVarFormParams.Add("index", parameterToString(*r.index, ""))
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

// APIDeleteLogGrafanacloudlogsRequest represents a request for the resource.
type APIDeleteLogGrafanacloudlogsRequest struct {
	ctx                         context.Context
	APIService                  LoggingGrafanacloudlogsAPI
	serviceID                   string
	versionID                   int32
	loggingGrafanacloudlogsName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogGrafanacloudlogsRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogGrafanacloudlogsExecute(r)
}

/*
DeleteLogGrafanacloudlogs Delete the Grafana Cloud Logs log endpoint

Delete the Grafana Cloud Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingGrafanacloudlogsName The name for the real-time logging configuration.
 @return APIDeleteLogGrafanacloudlogsRequest
*/
func (a *LoggingGrafanacloudlogsAPIService) DeleteLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32, loggingGrafanacloudlogsName string) APIDeleteLogGrafanacloudlogsRequest {
	return APIDeleteLogGrafanacloudlogsRequest{
		APIService:                  a,
		ctx:                         ctx,
		serviceID:                   serviceID,
		versionID:                   versionID,
		loggingGrafanacloudlogsName: loggingGrafanacloudlogsName,
	}
}

// DeleteLogGrafanacloudlogsExecute executes the request
//  @return InlineResponse200
func (a *LoggingGrafanacloudlogsAPIService) DeleteLogGrafanacloudlogsExecute(r APIDeleteLogGrafanacloudlogsRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGrafanacloudlogsAPIService.DeleteLogGrafanacloudlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/grafanacloudlogs/{logging_grafanacloudlogs_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_grafanacloudlogs_name"+"}", gourl.PathEscape(parameterToString(r.loggingGrafanacloudlogsName, "")))

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

// APIGetLogGrafanacloudlogsRequest represents a request for the resource.
type APIGetLogGrafanacloudlogsRequest struct {
	ctx                         context.Context
	APIService                  LoggingGrafanacloudlogsAPI
	serviceID                   string
	versionID                   int32
	loggingGrafanacloudlogsName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogGrafanacloudlogsRequest) Execute() (*LoggingGrafanacloudlogsResponse, *http.Response, error) {
	return r.APIService.GetLogGrafanacloudlogsExecute(r)
}

/*
GetLogGrafanacloudlogs Get a Grafana Cloud Logs log endpoint

Get the details of a Grafana Cloud Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingGrafanacloudlogsName The name for the real-time logging configuration.
 @return APIGetLogGrafanacloudlogsRequest
*/
func (a *LoggingGrafanacloudlogsAPIService) GetLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32, loggingGrafanacloudlogsName string) APIGetLogGrafanacloudlogsRequest {
	return APIGetLogGrafanacloudlogsRequest{
		APIService:                  a,
		ctx:                         ctx,
		serviceID:                   serviceID,
		versionID:                   versionID,
		loggingGrafanacloudlogsName: loggingGrafanacloudlogsName,
	}
}

// GetLogGrafanacloudlogsExecute executes the request
//  @return LoggingGrafanacloudlogsResponse
func (a *LoggingGrafanacloudlogsAPIService) GetLogGrafanacloudlogsExecute(r APIGetLogGrafanacloudlogsRequest) (*LoggingGrafanacloudlogsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingGrafanacloudlogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGrafanacloudlogsAPIService.GetLogGrafanacloudlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/grafanacloudlogs/{logging_grafanacloudlogs_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_grafanacloudlogs_name"+"}", gourl.PathEscape(parameterToString(r.loggingGrafanacloudlogsName, "")))

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

// APIListLogGrafanacloudlogsRequest represents a request for the resource.
type APIListLogGrafanacloudlogsRequest struct {
	ctx        context.Context
	APIService LoggingGrafanacloudlogsAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogGrafanacloudlogsRequest) Execute() ([]LoggingGrafanacloudlogsResponse, *http.Response, error) {
	return r.APIService.ListLogGrafanacloudlogsExecute(r)
}

/*
ListLogGrafanacloudlogs List Grafana Cloud Logs log endpoints

List all of the Grafana Cloud Logs logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogGrafanacloudlogsRequest
*/
func (a *LoggingGrafanacloudlogsAPIService) ListLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32) APIListLogGrafanacloudlogsRequest {
	return APIListLogGrafanacloudlogsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogGrafanacloudlogsExecute executes the request
//  @return []LoggingGrafanacloudlogsResponse
func (a *LoggingGrafanacloudlogsAPIService) ListLogGrafanacloudlogsExecute(r APIListLogGrafanacloudlogsRequest) ([]LoggingGrafanacloudlogsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingGrafanacloudlogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGrafanacloudlogsAPIService.ListLogGrafanacloudlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/grafanacloudlogs"
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

// APIUpdateLogGrafanacloudlogsRequest represents a request for the resource.
type APIUpdateLogGrafanacloudlogsRequest struct {
	ctx                         context.Context
	APIService                  LoggingGrafanacloudlogsAPI
	serviceID                   string
	versionID                   int32
	loggingGrafanacloudlogsName string
	name                        *string
	placement                   *string
	responseCondition           *string
	format                      *string
	formatVersion               *int32
	user                        *string
	url                         *string
	token                       *string
	index                       *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogGrafanacloudlogsRequest) Name(name string) *APIUpdateLogGrafanacloudlogsRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogGrafanacloudlogsRequest) Placement(placement string) *APIUpdateLogGrafanacloudlogsRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogGrafanacloudlogsRequest) ResponseCondition(responseCondition string) *APIUpdateLogGrafanacloudlogsRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogGrafanacloudlogsRequest) Format(format string) *APIUpdateLogGrafanacloudlogsRequest {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogGrafanacloudlogsRequest) FormatVersion(formatVersion int32) *APIUpdateLogGrafanacloudlogsRequest {
	r.formatVersion = &formatVersion
	return r
}

// User The Grafana Cloud Logs Dataset you want to log to.
func (r *APIUpdateLogGrafanacloudlogsRequest) User(user string) *APIUpdateLogGrafanacloudlogsRequest {
	r.user = &user
	return r
}

// URL The URL of the Loki instance in your Grafana stack.
func (r *APIUpdateLogGrafanacloudlogsRequest) URL(url string) *APIUpdateLogGrafanacloudlogsRequest {
	r.url = &url
	return r
}

// Token The Grafana Access Policy token with &#x60;logs:write&#x60; access scoped to your Loki instance.
func (r *APIUpdateLogGrafanacloudlogsRequest) Token(token string) *APIUpdateLogGrafanacloudlogsRequest {
	r.token = &token
	return r
}

// Index The Stream Labels, a JSON string used to identify the stream.
func (r *APIUpdateLogGrafanacloudlogsRequest) Index(index string) *APIUpdateLogGrafanacloudlogsRequest {
	r.index = &index
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogGrafanacloudlogsRequest) Execute() (*LoggingGrafanacloudlogsResponse, *http.Response, error) {
	return r.APIService.UpdateLogGrafanacloudlogsExecute(r)
}

/*
UpdateLogGrafanacloudlogs Update a Grafana Cloud Logs log endpoint

Update a Grafana Cloud Logs logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingGrafanacloudlogsName The name for the real-time logging configuration.
 @return APIUpdateLogGrafanacloudlogsRequest
*/
func (a *LoggingGrafanacloudlogsAPIService) UpdateLogGrafanacloudlogs(ctx context.Context, serviceID string, versionID int32, loggingGrafanacloudlogsName string) APIUpdateLogGrafanacloudlogsRequest {
	return APIUpdateLogGrafanacloudlogsRequest{
		APIService:                  a,
		ctx:                         ctx,
		serviceID:                   serviceID,
		versionID:                   versionID,
		loggingGrafanacloudlogsName: loggingGrafanacloudlogsName,
	}
}

// UpdateLogGrafanacloudlogsExecute executes the request
//  @return LoggingGrafanacloudlogsResponse
func (a *LoggingGrafanacloudlogsAPIService) UpdateLogGrafanacloudlogsExecute(r APIUpdateLogGrafanacloudlogsRequest) (*LoggingGrafanacloudlogsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingGrafanacloudlogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGrafanacloudlogsAPIService.UpdateLogGrafanacloudlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/grafanacloudlogs/{logging_grafanacloudlogs_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_grafanacloudlogs_name"+"}", gourl.PathEscape(parameterToString(r.loggingGrafanacloudlogsName, "")))

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
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.placement != nil {
		localVarFormParams.Add("placement", parameterToString(*r.placement, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
	}
	if r.token != nil {
		localVarFormParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.index != nil {
		localVarFormParams.Add("index", parameterToString(*r.index, ""))
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
