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

// LoggingGcsAPI defines an interface for interacting with the resource.
type LoggingGcsAPI interface {

	/*
		CreateLogGcs Create a GCS log endpoint

		Create GCS logging for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APICreateLogGcsRequest
	*/
	CreateLogGcs(ctx context.Context, serviceId string, versionId int32) APICreateLogGcsRequest

	// CreateLogGcsExecute executes the request
	//  @return LoggingGcsResponse
	CreateLogGcsExecute(r APICreateLogGcsRequest) (*LoggingGcsResponse, *http.Response, error)

	/*
		DeleteLogGcs Delete a GCS log endpoint

		Delete the GCS Logging for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingGcsName The name for the real-time logging configuration.
		 @return APIDeleteLogGcsRequest
	*/
	DeleteLogGcs(ctx context.Context, serviceId string, versionId int32, loggingGcsName string) APIDeleteLogGcsRequest

	// DeleteLogGcsExecute executes the request
	//  @return InlineResponse200
	DeleteLogGcsExecute(r APIDeleteLogGcsRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogGcs Get a GCS log endpoint

		Get the GCS Logging for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingGcsName The name for the real-time logging configuration.
		 @return APIGetLogGcsRequest
	*/
	GetLogGcs(ctx context.Context, serviceId string, versionId int32, loggingGcsName string) APIGetLogGcsRequest

	// GetLogGcsExecute executes the request
	//  @return LoggingGcsResponse
	GetLogGcsExecute(r APIGetLogGcsRequest) (*LoggingGcsResponse, *http.Response, error)

	/*
		ListLogGcs List GCS log endpoints

		List all of the GCS log endpoints for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APIListLogGcsRequest
	*/
	ListLogGcs(ctx context.Context, serviceId string, versionId int32) APIListLogGcsRequest

	// ListLogGcsExecute executes the request
	//  @return []LoggingGcsResponse
	ListLogGcsExecute(r APIListLogGcsRequest) ([]LoggingGcsResponse, *http.Response, error)

	/*
		UpdateLogGcs Update a GCS log endpoint

		Update the GCS for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param loggingGcsName The name for the real-time logging configuration.
		 @return APIUpdateLogGcsRequest
	*/
	UpdateLogGcs(ctx context.Context, serviceId string, versionId int32, loggingGcsName string) APIUpdateLogGcsRequest

	// UpdateLogGcsExecute executes the request
	//  @return LoggingGcsResponse
	UpdateLogGcsExecute(r APIUpdateLogGcsRequest) (*LoggingGcsResponse, *http.Response, error)
}

// LoggingGcsAPIService LoggingGcsAPI service
type LoggingGcsAPIService service

// APICreateLogGcsRequest represents a request for the resource.
type APICreateLogGcsRequest struct {
	ctx                 context.Context
	APIService          LoggingGcsAPI
	serviceId           string
	versionId           int32
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	messageType         *string
	timestampFormat     *string
	compressionCodec    *string
	period              *int32
	gzipLevel           *int32
	user                *string
	secretKey           *string
	accountName         *string
	bucketName          *string
	path                *string
	publicKey           *string
	projectId           *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogGcsRequest) Name(name string) *APICreateLogGcsRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogGcsRequest) Placement(placement string) *APICreateLogGcsRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogGcsRequest) ResponseCondition(responseCondition string) *APICreateLogGcsRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogGcsRequest) Format(format string) *APICreateLogGcsRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogGcsRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogGcsRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogGcsRequest) FormatVersion(formatVersion int32) *APICreateLogGcsRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APICreateLogGcsRequest) MessageType(messageType string) *APICreateLogGcsRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APICreateLogGcsRequest) TimestampFormat(timestampFormat string) *APICreateLogGcsRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogGcsRequest) CompressionCodec(compressionCodec string) *APICreateLogGcsRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogGcsRequest) Period(period int32) *APICreateLogGcsRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogGcsRequest) GzipLevel(gzipLevel int32) *APICreateLogGcsRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// User Your Google Cloud Platform service account email address. The &#x60;client_email&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APICreateLogGcsRequest) User(user string) *APICreateLogGcsRequest {
	r.user = &user
	return r
}

// SecretKey Your Google Cloud Platform account secret key. The &#x60;private_key&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APICreateLogGcsRequest) SecretKey(secretKey string) *APICreateLogGcsRequest {
	r.secretKey = &secretKey
	return r
}

// AccountName The name of the Google Cloud Platform service account associated with the target log collection service. Not required if &#x60;user&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APICreateLogGcsRequest) AccountName(accountName string) *APICreateLogGcsRequest {
	r.accountName = &accountName
	return r
}

// BucketName The name of the GCS bucket.
func (r *APICreateLogGcsRequest) BucketName(bucketName string) *APICreateLogGcsRequest {
	r.bucketName = &bucketName
	return r
}

// Path returns a pointer to a request.
func (r *APICreateLogGcsRequest) Path(path string) *APICreateLogGcsRequest {
	r.path = &path
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogGcsRequest) PublicKey(publicKey string) *APICreateLogGcsRequest {
	r.publicKey = &publicKey
	return r
}

// ProjectId Your Google Cloud Platform project ID. Required
func (r *APICreateLogGcsRequest) ProjectId(projectId string) *APICreateLogGcsRequest {
	r.projectId = &projectId
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogGcsRequest) Execute() (*LoggingGcsResponse, *http.Response, error) {
	return r.APIService.CreateLogGcsExecute(r)
}

/*
CreateLogGcs Create a GCS log endpoint

Create GCS logging for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APICreateLogGcsRequest
*/
func (a *LoggingGcsAPIService) CreateLogGcs(ctx context.Context, serviceId string, versionId int32) APICreateLogGcsRequest {
	return APICreateLogGcsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// CreateLogGcsExecute executes the request
//  @return LoggingGcsResponse
func (a *LoggingGcsAPIService) CreateLogGcsExecute(r APICreateLogGcsRequest) (*LoggingGcsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingGcsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGcsAPIService.CreateLogGcs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/gcs"
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
	if r.messageType != nil {
		localVarFormParams.Add("message_type", parameterToString(*r.messageType, ""))
	}
	if r.timestampFormat != nil {
		localVarFormParams.Add("timestamp_format", parameterToString(*r.timestampFormat, ""))
	}
	if r.compressionCodec != nil {
		localVarFormParams.Add("compression_codec", parameterToString(*r.compressionCodec, ""))
	}
	if r.period != nil {
		localVarFormParams.Add("period", parameterToString(*r.period, ""))
	}
	if r.gzipLevel != nil {
		localVarFormParams.Add("gzip_level", parameterToString(*r.gzipLevel, ""))
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
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.path != nil {
		paramJson, err := parameterToJSON(*r.path)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("path", paramJson)
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.projectId != nil {
		localVarFormParams.Add("project_id", parameterToString(*r.projectId, ""))
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

// APIDeleteLogGcsRequest represents a request for the resource.
type APIDeleteLogGcsRequest struct {
	ctx            context.Context
	APIService     LoggingGcsAPI
	serviceId      string
	versionId      int32
	loggingGcsName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogGcsRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogGcsExecute(r)
}

/*
DeleteLogGcs Delete a GCS log endpoint

Delete the GCS Logging for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingGcsName The name for the real-time logging configuration.
 @return APIDeleteLogGcsRequest
*/
func (a *LoggingGcsAPIService) DeleteLogGcs(ctx context.Context, serviceId string, versionId int32, loggingGcsName string) APIDeleteLogGcsRequest {
	return APIDeleteLogGcsRequest{
		APIService:     a,
		ctx:            ctx,
		serviceId:      serviceId,
		versionId:      versionId,
		loggingGcsName: loggingGcsName,
	}
}

// DeleteLogGcsExecute executes the request
//  @return InlineResponse200
func (a *LoggingGcsAPIService) DeleteLogGcsExecute(r APIDeleteLogGcsRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGcsAPIService.DeleteLogGcs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/gcs/{logging_gcs_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_gcs_name"+"}", gourl.PathEscape(parameterToString(r.loggingGcsName, "")))

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

// APIGetLogGcsRequest represents a request for the resource.
type APIGetLogGcsRequest struct {
	ctx            context.Context
	APIService     LoggingGcsAPI
	serviceId      string
	versionId      int32
	loggingGcsName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogGcsRequest) Execute() (*LoggingGcsResponse, *http.Response, error) {
	return r.APIService.GetLogGcsExecute(r)
}

/*
GetLogGcs Get a GCS log endpoint

Get the GCS Logging for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingGcsName The name for the real-time logging configuration.
 @return APIGetLogGcsRequest
*/
func (a *LoggingGcsAPIService) GetLogGcs(ctx context.Context, serviceId string, versionId int32, loggingGcsName string) APIGetLogGcsRequest {
	return APIGetLogGcsRequest{
		APIService:     a,
		ctx:            ctx,
		serviceId:      serviceId,
		versionId:      versionId,
		loggingGcsName: loggingGcsName,
	}
}

// GetLogGcsExecute executes the request
//  @return LoggingGcsResponse
func (a *LoggingGcsAPIService) GetLogGcsExecute(r APIGetLogGcsRequest) (*LoggingGcsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingGcsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGcsAPIService.GetLogGcs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/gcs/{logging_gcs_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_gcs_name"+"}", gourl.PathEscape(parameterToString(r.loggingGcsName, "")))

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

// APIListLogGcsRequest represents a request for the resource.
type APIListLogGcsRequest struct {
	ctx        context.Context
	APIService LoggingGcsAPI
	serviceId  string
	versionId  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogGcsRequest) Execute() ([]LoggingGcsResponse, *http.Response, error) {
	return r.APIService.ListLogGcsExecute(r)
}

/*
ListLogGcs List GCS log endpoints

List all of the GCS log endpoints for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APIListLogGcsRequest
*/
func (a *LoggingGcsAPIService) ListLogGcs(ctx context.Context, serviceId string, versionId int32) APIListLogGcsRequest {
	return APIListLogGcsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// ListLogGcsExecute executes the request
//  @return []LoggingGcsResponse
func (a *LoggingGcsAPIService) ListLogGcsExecute(r APIListLogGcsRequest) ([]LoggingGcsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingGcsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGcsAPIService.ListLogGcs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/gcs"
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

// APIUpdateLogGcsRequest represents a request for the resource.
type APIUpdateLogGcsRequest struct {
	ctx                 context.Context
	APIService          LoggingGcsAPI
	serviceId           string
	versionId           int32
	loggingGcsName      string
	name                *string
	placement           *string
	responseCondition   *string
	format              *string
	logProcessingRegion *string
	formatVersion       *int32
	messageType         *string
	timestampFormat     *string
	compressionCodec    *string
	period              *int32
	gzipLevel           *int32
	user                *string
	secretKey           *string
	accountName         *string
	bucketName          *string
	path                *string
	publicKey           *string
	projectId           *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogGcsRequest) Name(name string) *APIUpdateLogGcsRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogGcsRequest) Placement(placement string) *APIUpdateLogGcsRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogGcsRequest) ResponseCondition(responseCondition string) *APIUpdateLogGcsRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APIUpdateLogGcsRequest) Format(format string) *APIUpdateLogGcsRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogGcsRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogGcsRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogGcsRequest) FormatVersion(formatVersion int32) *APIUpdateLogGcsRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APIUpdateLogGcsRequest) MessageType(messageType string) *APIUpdateLogGcsRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APIUpdateLogGcsRequest) TimestampFormat(timestampFormat string) *APIUpdateLogGcsRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogGcsRequest) CompressionCodec(compressionCodec string) *APIUpdateLogGcsRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogGcsRequest) Period(period int32) *APIUpdateLogGcsRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogGcsRequest) GzipLevel(gzipLevel int32) *APIUpdateLogGcsRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// User Your Google Cloud Platform service account email address. The &#x60;client_email&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APIUpdateLogGcsRequest) User(user string) *APIUpdateLogGcsRequest {
	r.user = &user
	return r
}

// SecretKey Your Google Cloud Platform account secret key. The &#x60;private_key&#x60; field in your service account authentication JSON. Not required if &#x60;account_name&#x60; is specified.
func (r *APIUpdateLogGcsRequest) SecretKey(secretKey string) *APIUpdateLogGcsRequest {
	r.secretKey = &secretKey
	return r
}

// AccountName The name of the Google Cloud Platform service account associated with the target log collection service. Not required if &#x60;user&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APIUpdateLogGcsRequest) AccountName(accountName string) *APIUpdateLogGcsRequest {
	r.accountName = &accountName
	return r
}

// BucketName The name of the GCS bucket.
func (r *APIUpdateLogGcsRequest) BucketName(bucketName string) *APIUpdateLogGcsRequest {
	r.bucketName = &bucketName
	return r
}

// Path returns a pointer to a request.
func (r *APIUpdateLogGcsRequest) Path(path string) *APIUpdateLogGcsRequest {
	r.path = &path
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogGcsRequest) PublicKey(publicKey string) *APIUpdateLogGcsRequest {
	r.publicKey = &publicKey
	return r
}

// ProjectId Your Google Cloud Platform project ID. Required
func (r *APIUpdateLogGcsRequest) ProjectId(projectId string) *APIUpdateLogGcsRequest {
	r.projectId = &projectId
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogGcsRequest) Execute() (*LoggingGcsResponse, *http.Response, error) {
	return r.APIService.UpdateLogGcsExecute(r)
}

/*
UpdateLogGcs Update a GCS log endpoint

Update the GCS for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param loggingGcsName The name for the real-time logging configuration.
 @return APIUpdateLogGcsRequest
*/
func (a *LoggingGcsAPIService) UpdateLogGcs(ctx context.Context, serviceId string, versionId int32, loggingGcsName string) APIUpdateLogGcsRequest {
	return APIUpdateLogGcsRequest{
		APIService:     a,
		ctx:            ctx,
		serviceId:      serviceId,
		versionId:      versionId,
		loggingGcsName: loggingGcsName,
	}
}

// UpdateLogGcsExecute executes the request
//  @return LoggingGcsResponse
func (a *LoggingGcsAPIService) UpdateLogGcsExecute(r APIUpdateLogGcsRequest) (*LoggingGcsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingGcsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingGcsAPIService.UpdateLogGcs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/gcs/{logging_gcs_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_gcs_name"+"}", gourl.PathEscape(parameterToString(r.loggingGcsName, "")))

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
	if r.messageType != nil {
		localVarFormParams.Add("message_type", parameterToString(*r.messageType, ""))
	}
	if r.timestampFormat != nil {
		localVarFormParams.Add("timestamp_format", parameterToString(*r.timestampFormat, ""))
	}
	if r.compressionCodec != nil {
		localVarFormParams.Add("compression_codec", parameterToString(*r.compressionCodec, ""))
	}
	if r.period != nil {
		localVarFormParams.Add("period", parameterToString(*r.period, ""))
	}
	if r.gzipLevel != nil {
		localVarFormParams.Add("gzip_level", parameterToString(*r.gzipLevel, ""))
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
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.path != nil {
		paramJson, err := parameterToJSON(*r.path)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("path", paramJson)
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.projectId != nil {
		localVarFormParams.Add("project_id", parameterToString(*r.projectId, ""))
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
