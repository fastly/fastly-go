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

// LoggingPubsubAPI defines an interface for interacting with the resource.
type LoggingPubsubAPI interface {

	/*
	CreateLogGcpPubsub Create a GCP Cloud Pub/Sub log endpoint

	Create a Pub/Sub logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogGcpPubsubRequest
	*/
	CreateLogGcpPubsub(ctx context.Context, serviceID string, versionID int32) APICreateLogGcpPubsubRequest

	// CreateLogGcpPubsubExecute executes the request
	//  @return LoggingGooglePubsubResponse
	CreateLogGcpPubsubExecute(r APICreateLogGcpPubsubRequest) (*LoggingGooglePubsubResponse, *http.Response, error)

	/*
	DeleteLogGcpPubsub Delete a GCP Cloud Pub/Sub log endpoint

	Delete a Pub/Sub logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingGooglePubsubName The name for the real-time logging configuration.
	 @return APIDeleteLogGcpPubsubRequest
	*/
	DeleteLogGcpPubsub(ctx context.Context, serviceID string, versionID int32, loggingGooglePubsubName string) APIDeleteLogGcpPubsubRequest

	// DeleteLogGcpPubsubExecute executes the request
	//  @return InlineResponse200
	DeleteLogGcpPubsubExecute(r APIDeleteLogGcpPubsubRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogGcpPubsub Get a GCP Cloud Pub/Sub log endpoint

	Get the details for a Pub/Sub logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingGooglePubsubName The name for the real-time logging configuration.
	 @return APIGetLogGcpPubsubRequest
	*/
	GetLogGcpPubsub(ctx context.Context, serviceID string, versionID int32, loggingGooglePubsubName string) APIGetLogGcpPubsubRequest

	// GetLogGcpPubsubExecute executes the request
	//  @return LoggingGooglePubsubResponse
	GetLogGcpPubsubExecute(r APIGetLogGcpPubsubRequest) (*LoggingGooglePubsubResponse, *http.Response, error)

	/*
	ListLogGcpPubsub List GCP Cloud Pub/Sub log endpoints

	List all of the Pub/Sub logging objects for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogGcpPubsubRequest
	*/
	ListLogGcpPubsub(ctx context.Context, serviceID string, versionID int32) APIListLogGcpPubsubRequest

	// ListLogGcpPubsubExecute executes the request
	//  @return []LoggingGooglePubsubResponse
	ListLogGcpPubsubExecute(r APIListLogGcpPubsubRequest) ([]LoggingGooglePubsubResponse, *http.Response, error)

	/*
	UpdateLogGcpPubsub Update a GCP Cloud Pub/Sub log endpoint

	Update a Pub/Sub logging object for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingGooglePubsubName The name for the real-time logging configuration.
	 @return APIUpdateLogGcpPubsubRequest
	*/
	UpdateLogGcpPubsub(ctx context.Context, serviceID string, versionID int32, loggingGooglePubsubName string) APIUpdateLogGcpPubsubRequest

	// UpdateLogGcpPubsubExecute executes the request
	//  @return LoggingGooglePubsubResponse
	UpdateLogGcpPubsubExecute(r APIUpdateLogGcpPubsubRequest) (*LoggingGooglePubsubResponse, *http.Response, error)
}

// LoggingPubsubAPIService LoggingPubsubAPI service
type LoggingPubsubAPIService service

// APICreateLogGcpPubsubRequest represents a request for the resource.
type APICreateLogGcpPubsubRequest struct {
	ctx context.Context
	APIService LoggingPubsubAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	user *string
	secretKey *string
	accountName *string
	topic *string
	projectID *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogGcpPubsubRequest) Name(name string) *APICreateLogGcpPubsubRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogGcpPubsubRequest) Placement(placement string) *APICreateLogGcpPubsubRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogGcpPubsubRequest) ResponseCondition(responseCondition string) *APICreateLogGcpPubsubRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogGcpPubsubRequest) Format(format string) *APICreateLogGcpPubsubRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogGcpPubsubRequest) FormatVersion(formatVersion int32) *APICreateLogGcpPubsubRequest {
	r.formatVersion = &formatVersion
	return r
}
// User Your Google Cloud Platform service account email address. The &#x60;client_email&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APICreateLogGcpPubsubRequest) User(user string) *APICreateLogGcpPubsubRequest {
	r.user = &user
	return r
}
// SecretKey Your Google Cloud Platform account secret key. The &#x60;private_key&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APICreateLogGcpPubsubRequest) SecretKey(secretKey string) *APICreateLogGcpPubsubRequest {
	r.secretKey = &secretKey
	return r
}
// AccountName The name of the Google Cloud Platform service account associated with the target log collection service. Not required if &#x60;user&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APICreateLogGcpPubsubRequest) AccountName(accountName string) *APICreateLogGcpPubsubRequest {
	r.accountName = &accountName
	return r
}
// Topic The Google Cloud Pub/Sub topic to which logs will be published. Required.
func (r *APICreateLogGcpPubsubRequest) Topic(topic string) *APICreateLogGcpPubsubRequest {
	r.topic = &topic
	return r
}
// ProjectID Your Google Cloud Platform project ID. Required
func (r *APICreateLogGcpPubsubRequest) ProjectID(projectID string) *APICreateLogGcpPubsubRequest {
	r.projectID = &projectID
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogGcpPubsubRequest) Execute() (*LoggingGooglePubsubResponse, *http.Response, error) {
	return r.APIService.CreateLogGcpPubsubExecute(r)
}

/*
CreateLogGcpPubsub Create a GCP Cloud Pub/Sub log endpoint

Create a Pub/Sub logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogGcpPubsubRequest
*/
func (a *LoggingPubsubAPIService) CreateLogGcpPubsub(ctx context.Context, serviceID string, versionID int32) APICreateLogGcpPubsubRequest {
	return APICreateLogGcpPubsubRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogGcpPubsubExecute executes the request
//  @return LoggingGooglePubsubResponse
func (a *LoggingPubsubAPIService) CreateLogGcpPubsubExecute(r APICreateLogGcpPubsubRequest) (*LoggingGooglePubsubResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingGooglePubsubResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPubsubAPIService.CreateLogGcpPubsub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/pubsub"
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
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.accountName != nil {
		localVarFormParams.Add("account_name", parameterToString(*r.accountName, ""))
	}
	if r.topic != nil {
		localVarFormParams.Add("topic", parameterToString(*r.topic, ""))
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

// APIDeleteLogGcpPubsubRequest represents a request for the resource.
type APIDeleteLogGcpPubsubRequest struct {
	ctx context.Context
	APIService LoggingPubsubAPI
	serviceID string
	versionID int32
	loggingGooglePubsubName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogGcpPubsubRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogGcpPubsubExecute(r)
}

/*
DeleteLogGcpPubsub Delete a GCP Cloud Pub/Sub log endpoint

Delete a Pub/Sub logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingGooglePubsubName The name for the real-time logging configuration.
 @return APIDeleteLogGcpPubsubRequest
*/
func (a *LoggingPubsubAPIService) DeleteLogGcpPubsub(ctx context.Context, serviceID string, versionID int32, loggingGooglePubsubName string) APIDeleteLogGcpPubsubRequest {
	return APIDeleteLogGcpPubsubRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingGooglePubsubName: loggingGooglePubsubName,
	}
}

// DeleteLogGcpPubsubExecute executes the request
//  @return InlineResponse200
func (a *LoggingPubsubAPIService) DeleteLogGcpPubsubExecute(r APIDeleteLogGcpPubsubRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPubsubAPIService.DeleteLogGcpPubsub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/pubsub/{logging_google_pubsub_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_google_pubsub_name"+"}", gourl.PathEscape(parameterToString(r.loggingGooglePubsubName, "")))

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

// APIGetLogGcpPubsubRequest represents a request for the resource.
type APIGetLogGcpPubsubRequest struct {
	ctx context.Context
	APIService LoggingPubsubAPI
	serviceID string
	versionID int32
	loggingGooglePubsubName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogGcpPubsubRequest) Execute() (*LoggingGooglePubsubResponse, *http.Response, error) {
	return r.APIService.GetLogGcpPubsubExecute(r)
}

/*
GetLogGcpPubsub Get a GCP Cloud Pub/Sub log endpoint

Get the details for a Pub/Sub logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingGooglePubsubName The name for the real-time logging configuration.
 @return APIGetLogGcpPubsubRequest
*/
func (a *LoggingPubsubAPIService) GetLogGcpPubsub(ctx context.Context, serviceID string, versionID int32, loggingGooglePubsubName string) APIGetLogGcpPubsubRequest {
	return APIGetLogGcpPubsubRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingGooglePubsubName: loggingGooglePubsubName,
	}
}

// GetLogGcpPubsubExecute executes the request
//  @return LoggingGooglePubsubResponse
func (a *LoggingPubsubAPIService) GetLogGcpPubsubExecute(r APIGetLogGcpPubsubRequest) (*LoggingGooglePubsubResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingGooglePubsubResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPubsubAPIService.GetLogGcpPubsub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/pubsub/{logging_google_pubsub_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_google_pubsub_name"+"}", gourl.PathEscape(parameterToString(r.loggingGooglePubsubName, "")))

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

// APIListLogGcpPubsubRequest represents a request for the resource.
type APIListLogGcpPubsubRequest struct {
	ctx context.Context
	APIService LoggingPubsubAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogGcpPubsubRequest) Execute() ([]LoggingGooglePubsubResponse, *http.Response, error) {
	return r.APIService.ListLogGcpPubsubExecute(r)
}

/*
ListLogGcpPubsub List GCP Cloud Pub/Sub log endpoints

List all of the Pub/Sub logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogGcpPubsubRequest
*/
func (a *LoggingPubsubAPIService) ListLogGcpPubsub(ctx context.Context, serviceID string, versionID int32) APIListLogGcpPubsubRequest {
	return APIListLogGcpPubsubRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogGcpPubsubExecute executes the request
//  @return []LoggingGooglePubsubResponse
func (a *LoggingPubsubAPIService) ListLogGcpPubsubExecute(r APIListLogGcpPubsubRequest) ([]LoggingGooglePubsubResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingGooglePubsubResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPubsubAPIService.ListLogGcpPubsub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/pubsub"
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

// APIUpdateLogGcpPubsubRequest represents a request for the resource.
type APIUpdateLogGcpPubsubRequest struct {
	ctx context.Context
	APIService LoggingPubsubAPI
	serviceID string
	versionID int32
	loggingGooglePubsubName string
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	user *string
	secretKey *string
	accountName *string
	topic *string
	projectID *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogGcpPubsubRequest) Name(name string) *APIUpdateLogGcpPubsubRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogGcpPubsubRequest) Placement(placement string) *APIUpdateLogGcpPubsubRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogGcpPubsubRequest) ResponseCondition(responseCondition string) *APIUpdateLogGcpPubsubRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogGcpPubsubRequest) Format(format string) *APIUpdateLogGcpPubsubRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogGcpPubsubRequest) FormatVersion(formatVersion int32) *APIUpdateLogGcpPubsubRequest {
	r.formatVersion = &formatVersion
	return r
}
// User Your Google Cloud Platform service account email address. The &#x60;client_email&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APIUpdateLogGcpPubsubRequest) User(user string) *APIUpdateLogGcpPubsubRequest {
	r.user = &user
	return r
}
// SecretKey Your Google Cloud Platform account secret key. The &#x60;private_key&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APIUpdateLogGcpPubsubRequest) SecretKey(secretKey string) *APIUpdateLogGcpPubsubRequest {
	r.secretKey = &secretKey
	return r
}
// AccountName The name of the Google Cloud Platform service account associated with the target log collection service. Not required if &#x60;user&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APIUpdateLogGcpPubsubRequest) AccountName(accountName string) *APIUpdateLogGcpPubsubRequest {
	r.accountName = &accountName
	return r
}
// Topic The Google Cloud Pub/Sub topic to which logs will be published. Required.
func (r *APIUpdateLogGcpPubsubRequest) Topic(topic string) *APIUpdateLogGcpPubsubRequest {
	r.topic = &topic
	return r
}
// ProjectID Your Google Cloud Platform project ID. Required
func (r *APIUpdateLogGcpPubsubRequest) ProjectID(projectID string) *APIUpdateLogGcpPubsubRequest {
	r.projectID = &projectID
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogGcpPubsubRequest) Execute() (*LoggingGooglePubsubResponse, *http.Response, error) {
	return r.APIService.UpdateLogGcpPubsubExecute(r)
}

/*
UpdateLogGcpPubsub Update a GCP Cloud Pub/Sub log endpoint

Update a Pub/Sub logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingGooglePubsubName The name for the real-time logging configuration.
 @return APIUpdateLogGcpPubsubRequest
*/
func (a *LoggingPubsubAPIService) UpdateLogGcpPubsub(ctx context.Context, serviceID string, versionID int32, loggingGooglePubsubName string) APIUpdateLogGcpPubsubRequest {
	return APIUpdateLogGcpPubsubRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingGooglePubsubName: loggingGooglePubsubName,
	}
}

// UpdateLogGcpPubsubExecute executes the request
//  @return LoggingGooglePubsubResponse
func (a *LoggingPubsubAPIService) UpdateLogGcpPubsubExecute(r APIUpdateLogGcpPubsubRequest) (*LoggingGooglePubsubResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingGooglePubsubResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingPubsubAPIService.UpdateLogGcpPubsub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/pubsub/{logging_google_pubsub_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_google_pubsub_name"+"}", gourl.PathEscape(parameterToString(r.loggingGooglePubsubName, "")))

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
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.accountName != nil {
		localVarFormParams.Add("account_name", parameterToString(*r.accountName, ""))
	}
	if r.topic != nil {
		localVarFormParams.Add("topic", parameterToString(*r.topic, ""))
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
