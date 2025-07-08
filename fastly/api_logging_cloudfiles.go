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

// LoggingCloudfilesAPI defines an interface for interacting with the resource.
type LoggingCloudfilesAPI interface {

	/*
		CreateLogCloudfiles Create a Cloud Files log endpoint

		Create a Cloud Files log endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogCloudfilesRequest
	*/
	CreateLogCloudfiles(ctx context.Context, serviceID string, versionID int32) APICreateLogCloudfilesRequest

	// CreateLogCloudfilesExecute executes the request
	//  @return LoggingCloudfilesResponse
	CreateLogCloudfilesExecute(r APICreateLogCloudfilesRequest) (*LoggingCloudfilesResponse, *http.Response, error)

	/*
		DeleteLogCloudfiles Delete the Cloud Files log endpoint

		Delete the Cloud Files log endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingCloudfilesName The name for the real-time logging configuration.
		 @return APIDeleteLogCloudfilesRequest
	*/
	DeleteLogCloudfiles(ctx context.Context, serviceID string, versionID int32, loggingCloudfilesName string) APIDeleteLogCloudfilesRequest

	// DeleteLogCloudfilesExecute executes the request
	//  @return InlineResponse200
	DeleteLogCloudfilesExecute(r APIDeleteLogCloudfilesRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogCloudfiles Get a Cloud Files log endpoint

		Get the Cloud Files log endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingCloudfilesName The name for the real-time logging configuration.
		 @return APIGetLogCloudfilesRequest
	*/
	GetLogCloudfiles(ctx context.Context, serviceID string, versionID int32, loggingCloudfilesName string) APIGetLogCloudfilesRequest

	// GetLogCloudfilesExecute executes the request
	//  @return LoggingCloudfilesResponse
	GetLogCloudfilesExecute(r APIGetLogCloudfilesRequest) (*LoggingCloudfilesResponse, *http.Response, error)

	/*
		ListLogCloudfiles List Cloud Files log endpoints

		List all of the Cloud Files log endpoints for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogCloudfilesRequest
	*/
	ListLogCloudfiles(ctx context.Context, serviceID string, versionID int32) APIListLogCloudfilesRequest

	// ListLogCloudfilesExecute executes the request
	//  @return []LoggingCloudfilesResponse
	ListLogCloudfilesExecute(r APIListLogCloudfilesRequest) ([]LoggingCloudfilesResponse, *http.Response, error)

	/*
		UpdateLogCloudfiles Update the Cloud Files log endpoint

		Update the Cloud Files log endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingCloudfilesName The name for the real-time logging configuration.
		 @return APIUpdateLogCloudfilesRequest
	*/
	UpdateLogCloudfiles(ctx context.Context, serviceID string, versionID int32, loggingCloudfilesName string) APIUpdateLogCloudfilesRequest

	// UpdateLogCloudfilesExecute executes the request
	//  @return LoggingCloudfilesResponse
	UpdateLogCloudfilesExecute(r APIUpdateLogCloudfilesRequest) (*LoggingCloudfilesResponse, *http.Response, error)
}

// LoggingCloudfilesAPIService LoggingCloudfilesAPI service
type LoggingCloudfilesAPIService service

// APICreateLogCloudfilesRequest represents a request for the resource.
type APICreateLogCloudfilesRequest struct {
	ctx                 context.Context
	APIService          LoggingCloudfilesAPI
	serviceID           string
	versionID           int32
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
	accessKey           *string
	bucketName          *string
	path                *string
	region              *string
	publicKey           *string
	user                *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogCloudfilesRequest) Name(name string) *APICreateLogCloudfilesRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogCloudfilesRequest) Placement(placement string) *APICreateLogCloudfilesRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogCloudfilesRequest) ResponseCondition(responseCondition string) *APICreateLogCloudfilesRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogCloudfilesRequest) Format(format string) *APICreateLogCloudfilesRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogCloudfilesRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogCloudfilesRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogCloudfilesRequest) FormatVersion(formatVersion int32) *APICreateLogCloudfilesRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APICreateLogCloudfilesRequest) MessageType(messageType string) *APICreateLogCloudfilesRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APICreateLogCloudfilesRequest) TimestampFormat(timestampFormat string) *APICreateLogCloudfilesRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogCloudfilesRequest) CompressionCodec(compressionCodec string) *APICreateLogCloudfilesRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogCloudfilesRequest) Period(period int32) *APICreateLogCloudfilesRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogCloudfilesRequest) GzipLevel(gzipLevel int32) *APICreateLogCloudfilesRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// AccessKey Your Cloud Files account access key.
func (r *APICreateLogCloudfilesRequest) AccessKey(accessKey string) *APICreateLogCloudfilesRequest {
	r.accessKey = &accessKey
	return r
}

// BucketName The name of your Cloud Files container.
func (r *APICreateLogCloudfilesRequest) BucketName(bucketName string) *APICreateLogCloudfilesRequest {
	r.bucketName = &bucketName
	return r
}

// Path The path to upload logs to.
func (r *APICreateLogCloudfilesRequest) Path(path string) *APICreateLogCloudfilesRequest {
	r.path = &path
	return r
}

// Region The region to stream logs to.
func (r *APICreateLogCloudfilesRequest) Region(region string) *APICreateLogCloudfilesRequest {
	r.region = &region
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogCloudfilesRequest) PublicKey(publicKey string) *APICreateLogCloudfilesRequest {
	r.publicKey = &publicKey
	return r
}

// User The username for your Cloud Files account.
func (r *APICreateLogCloudfilesRequest) User(user string) *APICreateLogCloudfilesRequest {
	r.user = &user
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogCloudfilesRequest) Execute() (*LoggingCloudfilesResponse, *http.Response, error) {
	return r.APIService.CreateLogCloudfilesExecute(r)
}

/*
CreateLogCloudfiles Create a Cloud Files log endpoint

Create a Cloud Files log endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogCloudfilesRequest
*/
func (a *LoggingCloudfilesAPIService) CreateLogCloudfiles(ctx context.Context, serviceID string, versionID int32) APICreateLogCloudfilesRequest {
	return APICreateLogCloudfilesRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogCloudfilesExecute executes the request
//  @return LoggingCloudfilesResponse
func (a *LoggingCloudfilesAPIService) CreateLogCloudfilesExecute(r APICreateLogCloudfilesRequest) (*LoggingCloudfilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingCloudfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingCloudfilesAPIService.CreateLogCloudfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/cloudfiles"
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
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
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

// APIDeleteLogCloudfilesRequest represents a request for the resource.
type APIDeleteLogCloudfilesRequest struct {
	ctx                   context.Context
	APIService            LoggingCloudfilesAPI
	serviceID             string
	versionID             int32
	loggingCloudfilesName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogCloudfilesRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogCloudfilesExecute(r)
}

/*
DeleteLogCloudfiles Delete the Cloud Files log endpoint

Delete the Cloud Files log endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingCloudfilesName The name for the real-time logging configuration.
 @return APIDeleteLogCloudfilesRequest
*/
func (a *LoggingCloudfilesAPIService) DeleteLogCloudfiles(ctx context.Context, serviceID string, versionID int32, loggingCloudfilesName string) APIDeleteLogCloudfilesRequest {
	return APIDeleteLogCloudfilesRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceID:             serviceID,
		versionID:             versionID,
		loggingCloudfilesName: loggingCloudfilesName,
	}
}

// DeleteLogCloudfilesExecute executes the request
//  @return InlineResponse200
func (a *LoggingCloudfilesAPIService) DeleteLogCloudfilesExecute(r APIDeleteLogCloudfilesRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingCloudfilesAPIService.DeleteLogCloudfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_cloudfiles_name"+"}", gourl.PathEscape(parameterToString(r.loggingCloudfilesName, "")))

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

// APIGetLogCloudfilesRequest represents a request for the resource.
type APIGetLogCloudfilesRequest struct {
	ctx                   context.Context
	APIService            LoggingCloudfilesAPI
	serviceID             string
	versionID             int32
	loggingCloudfilesName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogCloudfilesRequest) Execute() (*LoggingCloudfilesResponse, *http.Response, error) {
	return r.APIService.GetLogCloudfilesExecute(r)
}

/*
GetLogCloudfiles Get a Cloud Files log endpoint

Get the Cloud Files log endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingCloudfilesName The name for the real-time logging configuration.
 @return APIGetLogCloudfilesRequest
*/
func (a *LoggingCloudfilesAPIService) GetLogCloudfiles(ctx context.Context, serviceID string, versionID int32, loggingCloudfilesName string) APIGetLogCloudfilesRequest {
	return APIGetLogCloudfilesRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceID:             serviceID,
		versionID:             versionID,
		loggingCloudfilesName: loggingCloudfilesName,
	}
}

// GetLogCloudfilesExecute executes the request
//  @return LoggingCloudfilesResponse
func (a *LoggingCloudfilesAPIService) GetLogCloudfilesExecute(r APIGetLogCloudfilesRequest) (*LoggingCloudfilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingCloudfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingCloudfilesAPIService.GetLogCloudfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_cloudfiles_name"+"}", gourl.PathEscape(parameterToString(r.loggingCloudfilesName, "")))

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

// APIListLogCloudfilesRequest represents a request for the resource.
type APIListLogCloudfilesRequest struct {
	ctx        context.Context
	APIService LoggingCloudfilesAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogCloudfilesRequest) Execute() ([]LoggingCloudfilesResponse, *http.Response, error) {
	return r.APIService.ListLogCloudfilesExecute(r)
}

/*
ListLogCloudfiles List Cloud Files log endpoints

List all of the Cloud Files log endpoints for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogCloudfilesRequest
*/
func (a *LoggingCloudfilesAPIService) ListLogCloudfiles(ctx context.Context, serviceID string, versionID int32) APIListLogCloudfilesRequest {
	return APIListLogCloudfilesRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogCloudfilesExecute executes the request
//  @return []LoggingCloudfilesResponse
func (a *LoggingCloudfilesAPIService) ListLogCloudfilesExecute(r APIListLogCloudfilesRequest) ([]LoggingCloudfilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingCloudfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingCloudfilesAPIService.ListLogCloudfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/cloudfiles"
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

// APIUpdateLogCloudfilesRequest represents a request for the resource.
type APIUpdateLogCloudfilesRequest struct {
	ctx                   context.Context
	APIService            LoggingCloudfilesAPI
	serviceID             string
	versionID             int32
	loggingCloudfilesName string
	name                  *string
	placement             *string
	responseCondition     *string
	format                *string
	logProcessingRegion   *string
	formatVersion         *int32
	messageType           *string
	timestampFormat       *string
	compressionCodec      *string
	period                *int32
	gzipLevel             *int32
	accessKey             *string
	bucketName            *string
	path                  *string
	region                *string
	publicKey             *string
	user                  *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogCloudfilesRequest) Name(name string) *APIUpdateLogCloudfilesRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogCloudfilesRequest) Placement(placement string) *APIUpdateLogCloudfilesRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogCloudfilesRequest) ResponseCondition(responseCondition string) *APIUpdateLogCloudfilesRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APIUpdateLogCloudfilesRequest) Format(format string) *APIUpdateLogCloudfilesRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogCloudfilesRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogCloudfilesRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogCloudfilesRequest) FormatVersion(formatVersion int32) *APIUpdateLogCloudfilesRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APIUpdateLogCloudfilesRequest) MessageType(messageType string) *APIUpdateLogCloudfilesRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APIUpdateLogCloudfilesRequest) TimestampFormat(timestampFormat string) *APIUpdateLogCloudfilesRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogCloudfilesRequest) CompressionCodec(compressionCodec string) *APIUpdateLogCloudfilesRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogCloudfilesRequest) Period(period int32) *APIUpdateLogCloudfilesRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogCloudfilesRequest) GzipLevel(gzipLevel int32) *APIUpdateLogCloudfilesRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// AccessKey Your Cloud Files account access key.
func (r *APIUpdateLogCloudfilesRequest) AccessKey(accessKey string) *APIUpdateLogCloudfilesRequest {
	r.accessKey = &accessKey
	return r
}

// BucketName The name of your Cloud Files container.
func (r *APIUpdateLogCloudfilesRequest) BucketName(bucketName string) *APIUpdateLogCloudfilesRequest {
	r.bucketName = &bucketName
	return r
}

// Path The path to upload logs to.
func (r *APIUpdateLogCloudfilesRequest) Path(path string) *APIUpdateLogCloudfilesRequest {
	r.path = &path
	return r
}

// Region The region to stream logs to.
func (r *APIUpdateLogCloudfilesRequest) Region(region string) *APIUpdateLogCloudfilesRequest {
	r.region = &region
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogCloudfilesRequest) PublicKey(publicKey string) *APIUpdateLogCloudfilesRequest {
	r.publicKey = &publicKey
	return r
}

// User The username for your Cloud Files account.
func (r *APIUpdateLogCloudfilesRequest) User(user string) *APIUpdateLogCloudfilesRequest {
	r.user = &user
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogCloudfilesRequest) Execute() (*LoggingCloudfilesResponse, *http.Response, error) {
	return r.APIService.UpdateLogCloudfilesExecute(r)
}

/*
UpdateLogCloudfiles Update the Cloud Files log endpoint

Update the Cloud Files log endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingCloudfilesName The name for the real-time logging configuration.
 @return APIUpdateLogCloudfilesRequest
*/
func (a *LoggingCloudfilesAPIService) UpdateLogCloudfiles(ctx context.Context, serviceID string, versionID int32, loggingCloudfilesName string) APIUpdateLogCloudfilesRequest {
	return APIUpdateLogCloudfilesRequest{
		APIService:            a,
		ctx:                   ctx,
		serviceID:             serviceID,
		versionID:             versionID,
		loggingCloudfilesName: loggingCloudfilesName,
	}
}

// UpdateLogCloudfilesExecute executes the request
//  @return LoggingCloudfilesResponse
func (a *LoggingCloudfilesAPIService) UpdateLogCloudfilesExecute(r APIUpdateLogCloudfilesRequest) (*LoggingCloudfilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingCloudfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingCloudfilesAPIService.UpdateLogCloudfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_cloudfiles_name"+"}", gourl.PathEscape(parameterToString(r.loggingCloudfilesName, "")))

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
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.region != nil {
		localVarFormParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
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
