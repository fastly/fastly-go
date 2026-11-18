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

// LoggingBigqueryAPI defines an interface for interacting with the resource.
type LoggingBigqueryAPI interface {

	/*
		CreateLogBigquery Create a BigQuery log endpoint

		Create a BigQuery logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogBigqueryRequest
	*/
	CreateLogBigquery(ctx context.Context, serviceID string, versionID int32) APICreateLogBigqueryRequest

	// CreateLogBigqueryExecute executes the request
	//  @return LoggingBigqueryResponse
	CreateLogBigqueryExecute(r APICreateLogBigqueryRequest) (*LoggingBigqueryResponse, *http.Response, error)

	/*
		DeleteLogBigquery Delete a BigQuery log endpoint

		Delete a BigQuery logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingBigqueryName The name for the real-time logging configuration.
		 @return APIDeleteLogBigqueryRequest
	*/
	DeleteLogBigquery(ctx context.Context, serviceID string, versionID int32, loggingBigqueryName string) APIDeleteLogBigqueryRequest

	// DeleteLogBigqueryExecute executes the request
	//  @return InlineResponse200
	DeleteLogBigqueryExecute(r APIDeleteLogBigqueryRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogBigquery Get a BigQuery log endpoint

		Get the details for a BigQuery logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingBigqueryName The name for the real-time logging configuration.
		 @return APIGetLogBigqueryRequest
	*/
	GetLogBigquery(ctx context.Context, serviceID string, versionID int32, loggingBigqueryName string) APIGetLogBigqueryRequest

	// GetLogBigqueryExecute executes the request
	//  @return LoggingBigqueryResponse
	GetLogBigqueryExecute(r APIGetLogBigqueryRequest) (*LoggingBigqueryResponse, *http.Response, error)

	/*
		ListLogBigquery List BigQuery log endpoints

		List all of the BigQuery logging objects for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogBigqueryRequest
	*/
	ListLogBigquery(ctx context.Context, serviceID string, versionID int32) APIListLogBigqueryRequest

	// ListLogBigqueryExecute executes the request
	//  @return []LoggingBigqueryResponse
	ListLogBigqueryExecute(r APIListLogBigqueryRequest) ([]LoggingBigqueryResponse, *http.Response, error)

	/*
		UpdateLogBigquery Update a BigQuery log endpoint

		Update a BigQuery logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingBigqueryName The name for the real-time logging configuration.
		 @return APIUpdateLogBigqueryRequest
	*/
	UpdateLogBigquery(ctx context.Context, serviceID string, versionID int32, loggingBigqueryName string) APIUpdateLogBigqueryRequest

	// UpdateLogBigqueryExecute executes the request
	//  @return LoggingBigqueryResponse
	UpdateLogBigqueryExecute(r APIUpdateLogBigqueryRequest) (*LoggingBigqueryResponse, *http.Response, error)
}

// LoggingBigqueryAPIService LoggingBigqueryAPI service
type LoggingBigqueryAPIService service

// APICreateLogBigqueryRequest represents a request for the resource.
type APICreateLogBigqueryRequest struct {
	ctx                 context.Context
	APIService          LoggingBigqueryAPI
	serviceID           string
	versionID           int32
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	user                *string
	secretKey           *string
	accountName         *string
	dataset             *string
	table               *string
	templateSuffix      *string
	projectID           *string
}

// Name The name of the BigQuery logging object. Used as a primary key for API access.
func (r *APICreateLogBigqueryRequest) Name(name string) *APICreateLogBigqueryRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogBigqueryRequest) Placement(placement string) *APICreateLogBigqueryRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogBigqueryRequest) ResponseCondition(responseCondition string) *APICreateLogBigqueryRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce JSON that matches the schema of your BigQuery table.
func (r *APICreateLogBigqueryRequest) Format(format string) *APICreateLogBigqueryRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogBigqueryRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogBigqueryRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogBigqueryRequest) FormatVersion(formatVersion int32) *APICreateLogBigqueryRequest {
	r.formatVersion = &formatVersion
	return r
}

// User Your Google Cloud Platform service account email address. The &#x60;client_email&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APICreateLogBigqueryRequest) User(user string) *APICreateLogBigqueryRequest {
	r.user = &user
	return r
}

// SecretKey Your Google Cloud Platform account secret key. The &#x60;private_key&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APICreateLogBigqueryRequest) SecretKey(secretKey string) *APICreateLogBigqueryRequest {
	r.secretKey = &secretKey
	return r
}

// AccountName The name of the Google Cloud Platform service account associated with the target log collection service. Not required if &#x60;user&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APICreateLogBigqueryRequest) AccountName(accountName string) *APICreateLogBigqueryRequest {
	r.accountName = &accountName
	return r
}

// Dataset Your BigQuery dataset.
func (r *APICreateLogBigqueryRequest) Dataset(dataset string) *APICreateLogBigqueryRequest {
	r.dataset = &dataset
	return r
}

// Table Your BigQuery table.
func (r *APICreateLogBigqueryRequest) Table(table string) *APICreateLogBigqueryRequest {
	r.table = &table
	return r
}

// TemplateSuffix BigQuery table name suffix template. Optional.
func (r *APICreateLogBigqueryRequest) TemplateSuffix(templateSuffix string) *APICreateLogBigqueryRequest {
	r.templateSuffix = &templateSuffix
	return r
}

// ProjectID Your Google Cloud Platform project ID. Required
func (r *APICreateLogBigqueryRequest) ProjectID(projectID string) *APICreateLogBigqueryRequest {
	r.projectID = &projectID
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogBigqueryRequest) Execute() (*LoggingBigqueryResponse, *http.Response, error) {
	return r.APIService.CreateLogBigqueryExecute(r)
}

/*
CreateLogBigquery Create a BigQuery log endpoint

Create a BigQuery logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogBigqueryRequest
*/
func (a *LoggingBigqueryAPIService) CreateLogBigquery(ctx context.Context, serviceID string, versionID int32) APICreateLogBigqueryRequest {
	return APICreateLogBigqueryRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogBigqueryExecute executes the request
//  @return LoggingBigqueryResponse
func (a *LoggingBigqueryAPIService) CreateLogBigqueryExecute(r APICreateLogBigqueryRequest) (*LoggingBigqueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingBigqueryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingBigqueryAPIService.CreateLogBigquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/bigquery"
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
	if r.logProcessingRegion != nil {
		localVarFormParams.Add("log_processing_region", parameterToString(*r.logProcessingRegion, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.accountName != nil {
		localVarFormParams.Add("account_name", parameterToString(*r.accountName, ""))
	}
	if r.dataset != nil {
		localVarFormParams.Add("dataset", parameterToString(*r.dataset, ""))
	}
	if r.table != nil {
		localVarFormParams.Add("table", parameterToString(*r.table, ""))
	}
	if r.templateSuffix != nil {
		localVarFormParams.Add("template_suffix", parameterToString(*r.templateSuffix, ""))
	}
	if r.projectID != nil {
		localVarFormParams.Add("project_id", parameterToString(*r.projectID, ""))
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

// APIDeleteLogBigqueryRequest represents a request for the resource.
type APIDeleteLogBigqueryRequest struct {
	ctx                 context.Context
	APIService          LoggingBigqueryAPI
	serviceID           string
	versionID           int32
	loggingBigqueryName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogBigqueryRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogBigqueryExecute(r)
}

/*
DeleteLogBigquery Delete a BigQuery log endpoint

Delete a BigQuery logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingBigqueryName The name for the real-time logging configuration.
 @return APIDeleteLogBigqueryRequest
*/
func (a *LoggingBigqueryAPIService) DeleteLogBigquery(ctx context.Context, serviceID string, versionID int32, loggingBigqueryName string) APIDeleteLogBigqueryRequest {
	return APIDeleteLogBigqueryRequest{
		APIService:          a,
		ctx:                 ctx,
		serviceID:           serviceID,
		versionID:           versionID,
		loggingBigqueryName: loggingBigqueryName,
	}
}

// DeleteLogBigqueryExecute executes the request
//  @return InlineResponse200
func (a *LoggingBigqueryAPIService) DeleteLogBigqueryExecute(r APIDeleteLogBigqueryRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingBigqueryAPIService.DeleteLogBigquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_bigquery_name"+"}", gourl.PathEscape(parameterToString(r.loggingBigqueryName, "")))

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

// APIGetLogBigqueryRequest represents a request for the resource.
type APIGetLogBigqueryRequest struct {
	ctx                 context.Context
	APIService          LoggingBigqueryAPI
	serviceID           string
	versionID           int32
	loggingBigqueryName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogBigqueryRequest) Execute() (*LoggingBigqueryResponse, *http.Response, error) {
	return r.APIService.GetLogBigqueryExecute(r)
}

/*
GetLogBigquery Get a BigQuery log endpoint

Get the details for a BigQuery logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingBigqueryName The name for the real-time logging configuration.
 @return APIGetLogBigqueryRequest
*/
func (a *LoggingBigqueryAPIService) GetLogBigquery(ctx context.Context, serviceID string, versionID int32, loggingBigqueryName string) APIGetLogBigqueryRequest {
	return APIGetLogBigqueryRequest{
		APIService:          a,
		ctx:                 ctx,
		serviceID:           serviceID,
		versionID:           versionID,
		loggingBigqueryName: loggingBigqueryName,
	}
}

// GetLogBigqueryExecute executes the request
//  @return LoggingBigqueryResponse
func (a *LoggingBigqueryAPIService) GetLogBigqueryExecute(r APIGetLogBigqueryRequest) (*LoggingBigqueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingBigqueryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingBigqueryAPIService.GetLogBigquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_bigquery_name"+"}", gourl.PathEscape(parameterToString(r.loggingBigqueryName, "")))

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

// APIListLogBigqueryRequest represents a request for the resource.
type APIListLogBigqueryRequest struct {
	ctx        context.Context
	APIService LoggingBigqueryAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogBigqueryRequest) Execute() ([]LoggingBigqueryResponse, *http.Response, error) {
	return r.APIService.ListLogBigqueryExecute(r)
}

/*
ListLogBigquery List BigQuery log endpoints

List all of the BigQuery logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogBigqueryRequest
*/
func (a *LoggingBigqueryAPIService) ListLogBigquery(ctx context.Context, serviceID string, versionID int32) APIListLogBigqueryRequest {
	return APIListLogBigqueryRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogBigqueryExecute executes the request
//  @return []LoggingBigqueryResponse
func (a *LoggingBigqueryAPIService) ListLogBigqueryExecute(r APIListLogBigqueryRequest) ([]LoggingBigqueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingBigqueryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingBigqueryAPIService.ListLogBigquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/bigquery"
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

// APIUpdateLogBigqueryRequest represents a request for the resource.
type APIUpdateLogBigqueryRequest struct {
	ctx                 context.Context
	APIService          LoggingBigqueryAPI
	serviceID           string
	versionID           int32
	loggingBigqueryName string
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	user                *string
	secretKey           *string
	accountName         *string
	dataset             *string
	table               *string
	templateSuffix      *string
	projectID           *string
}

// Name The name of the BigQuery logging object. Used as a primary key for API access.
func (r *APIUpdateLogBigqueryRequest) Name(name string) *APIUpdateLogBigqueryRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogBigqueryRequest) Placement(placement string) *APIUpdateLogBigqueryRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogBigqueryRequest) ResponseCondition(responseCondition string) *APIUpdateLogBigqueryRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce JSON that matches the schema of your BigQuery table.
func (r *APIUpdateLogBigqueryRequest) Format(format string) *APIUpdateLogBigqueryRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogBigqueryRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogBigqueryRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogBigqueryRequest) FormatVersion(formatVersion int32) *APIUpdateLogBigqueryRequest {
	r.formatVersion = &formatVersion
	return r
}

// User Your Google Cloud Platform service account email address. The &#x60;client_email&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APIUpdateLogBigqueryRequest) User(user string) *APIUpdateLogBigqueryRequest {
	r.user = &user
	return r
}

// SecretKey Your Google Cloud Platform account secret key. The &#x60;private_key&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APIUpdateLogBigqueryRequest) SecretKey(secretKey string) *APIUpdateLogBigqueryRequest {
	r.secretKey = &secretKey
	return r
}

// AccountName The name of the Google Cloud Platform service account associated with the target log collection service. Not required if &#x60;user&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APIUpdateLogBigqueryRequest) AccountName(accountName string) *APIUpdateLogBigqueryRequest {
	r.accountName = &accountName
	return r
}

// Dataset Your BigQuery dataset.
func (r *APIUpdateLogBigqueryRequest) Dataset(dataset string) *APIUpdateLogBigqueryRequest {
	r.dataset = &dataset
	return r
}

// Table Your BigQuery table.
func (r *APIUpdateLogBigqueryRequest) Table(table string) *APIUpdateLogBigqueryRequest {
	r.table = &table
	return r
}

// TemplateSuffix BigQuery table name suffix template. Optional.
func (r *APIUpdateLogBigqueryRequest) TemplateSuffix(templateSuffix string) *APIUpdateLogBigqueryRequest {
	r.templateSuffix = &templateSuffix
	return r
}

// ProjectID Your Google Cloud Platform project ID. Required
func (r *APIUpdateLogBigqueryRequest) ProjectID(projectID string) *APIUpdateLogBigqueryRequest {
	r.projectID = &projectID
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogBigqueryRequest) Execute() (*LoggingBigqueryResponse, *http.Response, error) {
	return r.APIService.UpdateLogBigqueryExecute(r)
}

/*
UpdateLogBigquery Update a BigQuery log endpoint

Update a BigQuery logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingBigqueryName The name for the real-time logging configuration.
 @return APIUpdateLogBigqueryRequest
*/
func (a *LoggingBigqueryAPIService) UpdateLogBigquery(ctx context.Context, serviceID string, versionID int32, loggingBigqueryName string) APIUpdateLogBigqueryRequest {
	return APIUpdateLogBigqueryRequest{
		APIService:          a,
		ctx:                 ctx,
		serviceID:           serviceID,
		versionID:           versionID,
		loggingBigqueryName: loggingBigqueryName,
	}
}

// UpdateLogBigqueryExecute executes the request
//  @return LoggingBigqueryResponse
func (a *LoggingBigqueryAPIService) UpdateLogBigqueryExecute(r APIUpdateLogBigqueryRequest) (*LoggingBigqueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingBigqueryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingBigqueryAPIService.UpdateLogBigquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_bigquery_name"+"}", gourl.PathEscape(parameterToString(r.loggingBigqueryName, "")))

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
	if r.logProcessingRegion != nil {
		localVarFormParams.Add("log_processing_region", parameterToString(*r.logProcessingRegion, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.accountName != nil {
		localVarFormParams.Add("account_name", parameterToString(*r.accountName, ""))
	}
	if r.dataset != nil {
		localVarFormParams.Add("dataset", parameterToString(*r.dataset, ""))
	}
	if r.table != nil {
		localVarFormParams.Add("table", parameterToString(*r.table, ""))
	}
	if r.templateSuffix != nil {
		localVarFormParams.Add("template_suffix", parameterToString(*r.templateSuffix, ""))
	}
	if r.projectID != nil {
		localVarFormParams.Add("project_id", parameterToString(*r.projectID, ""))
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
