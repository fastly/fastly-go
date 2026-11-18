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

// LoggingKinesisAPI defines an interface for interacting with the resource.
type LoggingKinesisAPI interface {

	/*
		CreateLogKinesis Create  an Amazon Kinesis log endpoint

		Create an Amazon Kinesis Data Streams logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogKinesisRequest
	*/
	CreateLogKinesis(ctx context.Context, serviceID string, versionID int32) APICreateLogKinesisRequest

	// CreateLogKinesisExecute executes the request
	//  @return LoggingKinesisResponse
	CreateLogKinesisExecute(r APICreateLogKinesisRequest) (*LoggingKinesisResponse, *http.Response, error)

	/*
		DeleteLogKinesis Delete the Amazon Kinesis log endpoint

		Delete an Amazon Kinesis Data Streams logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingKinesisName The name for the real-time logging configuration.
		 @return APIDeleteLogKinesisRequest
	*/
	DeleteLogKinesis(ctx context.Context, serviceID string, versionID int32, loggingKinesisName string) APIDeleteLogKinesisRequest

	// DeleteLogKinesisExecute executes the request
	//  @return InlineResponse200
	DeleteLogKinesisExecute(r APIDeleteLogKinesisRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogKinesis Get an Amazon Kinesis log endpoint

		Get the details for an Amazon Kinesis Data Streams logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingKinesisName The name for the real-time logging configuration.
		 @return APIGetLogKinesisRequest
	*/
	GetLogKinesis(ctx context.Context, serviceID string, versionID int32, loggingKinesisName string) APIGetLogKinesisRequest

	// GetLogKinesisExecute executes the request
	//  @return LoggingKinesisResponse
	GetLogKinesisExecute(r APIGetLogKinesisRequest) (*LoggingKinesisResponse, *http.Response, error)

	/*
		ListLogKinesis List Amazon Kinesis log endpoints

		List all of the Amazon Kinesis Data Streams logging objects for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogKinesisRequest
	*/
	ListLogKinesis(ctx context.Context, serviceID string, versionID int32) APIListLogKinesisRequest

	// ListLogKinesisExecute executes the request
	//  @return []LoggingKinesisResponse
	ListLogKinesisExecute(r APIListLogKinesisRequest) ([]LoggingKinesisResponse, *http.Response, error)

	/*
		UpdateLogKinesis Update the Amazon Kinesis log endpoint

		Update an Amazon Kinesis Data Streams logging object for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingKinesisName The name for the real-time logging configuration.
		 @return APIUpdateLogKinesisRequest
	*/
	UpdateLogKinesis(ctx context.Context, serviceID string, versionID int32, loggingKinesisName string) APIUpdateLogKinesisRequest

	// UpdateLogKinesisExecute executes the request
	//  @return LoggingKinesisResponse
	UpdateLogKinesisExecute(r APIUpdateLogKinesisRequest) (*LoggingKinesisResponse, *http.Response, error)
}

// LoggingKinesisAPIService LoggingKinesisAPI service
type LoggingKinesisAPIService service

// APICreateLogKinesisRequest represents a request for the resource.
type APICreateLogKinesisRequest struct {
	ctx           context.Context
	APIService    LoggingKinesisAPI
	serviceID     string
	versionID     int32
	name          *string
	placement     *LoggingPlacement
	format        *string
	topic         *string
	region        *AwsRegion
	secretKey     *string
	accessKey     *string
	iamRole       *string
	formatVersion *int32
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogKinesisRequest) Name(name string) *APICreateLogKinesisRequest {
	r.name = &name
	return r
}

// Placement returns a pointer to a request.
func (r *APICreateLogKinesisRequest) Placement(placement LoggingPlacement) *APICreateLogKinesisRequest {
	r.placement = &placement
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogKinesisRequest) Format(format string) *APICreateLogKinesisRequest {
	r.format = &format
	return r
}

// Topic The Amazon Kinesis stream to send logs to. Required.
func (r *APICreateLogKinesisRequest) Topic(topic string) *APICreateLogKinesisRequest {
	r.topic = &topic
	return r
}

// Region returns a pointer to a request.
func (r *APICreateLogKinesisRequest) Region(region AwsRegion) *APICreateLogKinesisRequest {
	r.region = &region
	return r
}

// SecretKey The secret key associated with the target Amazon Kinesis stream. Not required if &#x60;iam_role&#x60; is specified.
func (r *APICreateLogKinesisRequest) SecretKey(secretKey string) *APICreateLogKinesisRequest {
	r.secretKey = &secretKey
	return r
}

// AccessKey The access key associated with the target Amazon Kinesis stream. Not required if &#x60;iam_role&#x60; is specified.
func (r *APICreateLogKinesisRequest) AccessKey(accessKey string) *APICreateLogKinesisRequest {
	r.accessKey = &accessKey
	return r
}

// IamRole The ARN for an IAM role granting Fastly access to the target Amazon Kinesis stream. Not required if &#x60;access_key&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APICreateLogKinesisRequest) IamRole(iamRole string) *APICreateLogKinesisRequest {
	r.iamRole = &iamRole
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogKinesisRequest) FormatVersion(formatVersion int32) *APICreateLogKinesisRequest {
	r.formatVersion = &formatVersion
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogKinesisRequest) Execute() (*LoggingKinesisResponse, *http.Response, error) {
	return r.APIService.CreateLogKinesisExecute(r)
}

/*
CreateLogKinesis Create  an Amazon Kinesis log endpoint

Create an Amazon Kinesis Data Streams logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogKinesisRequest
*/
func (a *LoggingKinesisAPIService) CreateLogKinesis(ctx context.Context, serviceID string, versionID int32) APICreateLogKinesisRequest {
	return APICreateLogKinesisRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogKinesisExecute executes the request
//  @return LoggingKinesisResponse
func (a *LoggingKinesisAPIService) CreateLogKinesisExecute(r APICreateLogKinesisRequest) (*LoggingKinesisResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingKinesisResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKinesisAPIService.CreateLogKinesis")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kinesis"
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
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.topic != nil {
		localVarFormParams.Add("topic", parameterToString(*r.topic, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.iamRole != nil {
		localVarFormParams.Add("iam_role", parameterToString(*r.iamRole, ""))
	}
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
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

// APIDeleteLogKinesisRequest represents a request for the resource.
type APIDeleteLogKinesisRequest struct {
	ctx                context.Context
	APIService         LoggingKinesisAPI
	serviceID          string
	versionID          int32
	loggingKinesisName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogKinesisRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogKinesisExecute(r)
}

/*
DeleteLogKinesis Delete the Amazon Kinesis log endpoint

Delete an Amazon Kinesis Data Streams logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingKinesisName The name for the real-time logging configuration.
 @return APIDeleteLogKinesisRequest
*/
func (a *LoggingKinesisAPIService) DeleteLogKinesis(ctx context.Context, serviceID string, versionID int32, loggingKinesisName string) APIDeleteLogKinesisRequest {
	return APIDeleteLogKinesisRequest{
		APIService:         a,
		ctx:                ctx,
		serviceID:          serviceID,
		versionID:          versionID,
		loggingKinesisName: loggingKinesisName,
	}
}

// DeleteLogKinesisExecute executes the request
//  @return InlineResponse200
func (a *LoggingKinesisAPIService) DeleteLogKinesisExecute(r APIDeleteLogKinesisRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKinesisAPIService.DeleteLogKinesis")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kinesis/{logging_kinesis_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_kinesis_name"+"}", gourl.PathEscape(parameterToString(r.loggingKinesisName, "")))

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

// APIGetLogKinesisRequest represents a request for the resource.
type APIGetLogKinesisRequest struct {
	ctx                context.Context
	APIService         LoggingKinesisAPI
	serviceID          string
	versionID          int32
	loggingKinesisName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogKinesisRequest) Execute() (*LoggingKinesisResponse, *http.Response, error) {
	return r.APIService.GetLogKinesisExecute(r)
}

/*
GetLogKinesis Get an Amazon Kinesis log endpoint

Get the details for an Amazon Kinesis Data Streams logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingKinesisName The name for the real-time logging configuration.
 @return APIGetLogKinesisRequest
*/
func (a *LoggingKinesisAPIService) GetLogKinesis(ctx context.Context, serviceID string, versionID int32, loggingKinesisName string) APIGetLogKinesisRequest {
	return APIGetLogKinesisRequest{
		APIService:         a,
		ctx:                ctx,
		serviceID:          serviceID,
		versionID:          versionID,
		loggingKinesisName: loggingKinesisName,
	}
}

// GetLogKinesisExecute executes the request
//  @return LoggingKinesisResponse
func (a *LoggingKinesisAPIService) GetLogKinesisExecute(r APIGetLogKinesisRequest) (*LoggingKinesisResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingKinesisResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKinesisAPIService.GetLogKinesis")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kinesis/{logging_kinesis_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_kinesis_name"+"}", gourl.PathEscape(parameterToString(r.loggingKinesisName, "")))

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

// APIListLogKinesisRequest represents a request for the resource.
type APIListLogKinesisRequest struct {
	ctx        context.Context
	APIService LoggingKinesisAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogKinesisRequest) Execute() ([]LoggingKinesisResponse, *http.Response, error) {
	return r.APIService.ListLogKinesisExecute(r)
}

/*
ListLogKinesis List Amazon Kinesis log endpoints

List all of the Amazon Kinesis Data Streams logging objects for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogKinesisRequest
*/
func (a *LoggingKinesisAPIService) ListLogKinesis(ctx context.Context, serviceID string, versionID int32) APIListLogKinesisRequest {
	return APIListLogKinesisRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogKinesisExecute executes the request
//  @return []LoggingKinesisResponse
func (a *LoggingKinesisAPIService) ListLogKinesisExecute(r APIListLogKinesisRequest) ([]LoggingKinesisResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingKinesisResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKinesisAPIService.ListLogKinesis")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kinesis"
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

// APIUpdateLogKinesisRequest represents a request for the resource.
type APIUpdateLogKinesisRequest struct {
	ctx                context.Context
	APIService         LoggingKinesisAPI
	serviceID          string
	versionID          int32
	loggingKinesisName string
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogKinesisRequest) Execute() (*LoggingKinesisResponse, *http.Response, error) {
	return r.APIService.UpdateLogKinesisExecute(r)
}

/*
UpdateLogKinesis Update the Amazon Kinesis log endpoint

Update an Amazon Kinesis Data Streams logging object for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingKinesisName The name for the real-time logging configuration.
 @return APIUpdateLogKinesisRequest
*/
func (a *LoggingKinesisAPIService) UpdateLogKinesis(ctx context.Context, serviceID string, versionID int32, loggingKinesisName string) APIUpdateLogKinesisRequest {
	return APIUpdateLogKinesisRequest{
		APIService:         a,
		ctx:                ctx,
		serviceID:          serviceID,
		versionID:          versionID,
		loggingKinesisName: loggingKinesisName,
	}
}

// UpdateLogKinesisExecute executes the request
//  @return LoggingKinesisResponse
func (a *LoggingKinesisAPIService) UpdateLogKinesisExecute(r APIUpdateLogKinesisRequest) (*LoggingKinesisResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingKinesisResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKinesisAPIService.UpdateLogKinesis")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kinesis/{logging_kinesis_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_kinesis_name"+"}", gourl.PathEscape(parameterToString(r.loggingKinesisName, "")))

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
