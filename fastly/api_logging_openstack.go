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

// LoggingOpenstackAPI defines an interface for interacting with the resource.
type LoggingOpenstackAPI interface {

	/*
		CreateLogOpenstack Create an OpenStack log endpoint

		Create a openstack for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogOpenstackRequest
	*/
	CreateLogOpenstack(ctx context.Context, serviceID string, versionID int32) APICreateLogOpenstackRequest

	// CreateLogOpenstackExecute executes the request
	//  @return LoggingOpenstackResponse
	CreateLogOpenstackExecute(r APICreateLogOpenstackRequest) (*LoggingOpenstackResponse, *http.Response, error)

	/*
		DeleteLogOpenstack Delete an OpenStack log endpoint

		Delete the openstack for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingOpenstackName The name for the real-time logging configuration.
		 @return APIDeleteLogOpenstackRequest
	*/
	DeleteLogOpenstack(ctx context.Context, serviceID string, versionID int32, loggingOpenstackName string) APIDeleteLogOpenstackRequest

	// DeleteLogOpenstackExecute executes the request
	//  @return InlineResponse200
	DeleteLogOpenstackExecute(r APIDeleteLogOpenstackRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogOpenstack Get an OpenStack log endpoint

		Get the openstack for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingOpenstackName The name for the real-time logging configuration.
		 @return APIGetLogOpenstackRequest
	*/
	GetLogOpenstack(ctx context.Context, serviceID string, versionID int32, loggingOpenstackName string) APIGetLogOpenstackRequest

	// GetLogOpenstackExecute executes the request
	//  @return LoggingOpenstackResponse
	GetLogOpenstackExecute(r APIGetLogOpenstackRequest) (*LoggingOpenstackResponse, *http.Response, error)

	/*
		ListLogOpenstack List OpenStack log endpoints

		List all of the openstacks for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogOpenstackRequest
	*/
	ListLogOpenstack(ctx context.Context, serviceID string, versionID int32) APIListLogOpenstackRequest

	// ListLogOpenstackExecute executes the request
	//  @return []LoggingOpenstackResponse
	ListLogOpenstackExecute(r APIListLogOpenstackRequest) ([]LoggingOpenstackResponse, *http.Response, error)

	/*
		UpdateLogOpenstack Update an OpenStack log endpoint

		Update the openstack for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingOpenstackName The name for the real-time logging configuration.
		 @return APIUpdateLogOpenstackRequest
	*/
	UpdateLogOpenstack(ctx context.Context, serviceID string, versionID int32, loggingOpenstackName string) APIUpdateLogOpenstackRequest

	// UpdateLogOpenstackExecute executes the request
	//  @return LoggingOpenstackResponse
	UpdateLogOpenstackExecute(r APIUpdateLogOpenstackRequest) (*LoggingOpenstackResponse, *http.Response, error)
}

// LoggingOpenstackAPIService LoggingOpenstackAPI service
type LoggingOpenstackAPIService service

// APICreateLogOpenstackRequest represents a request for the resource.
type APICreateLogOpenstackRequest struct {
	ctx                 context.Context
	APIService          LoggingOpenstackAPI
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
	publicKey           *string
	url                 *string
	user                *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogOpenstackRequest) Name(name string) *APICreateLogOpenstackRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogOpenstackRequest) Placement(placement string) *APICreateLogOpenstackRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogOpenstackRequest) ResponseCondition(responseCondition string) *APICreateLogOpenstackRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APICreateLogOpenstackRequest) Format(format string) *APICreateLogOpenstackRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APICreateLogOpenstackRequest) LogProcessingRegion(logProcessingRegion string) *APICreateLogOpenstackRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogOpenstackRequest) FormatVersion(formatVersion int32) *APICreateLogOpenstackRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APICreateLogOpenstackRequest) MessageType(messageType string) *APICreateLogOpenstackRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APICreateLogOpenstackRequest) TimestampFormat(timestampFormat string) *APICreateLogOpenstackRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogOpenstackRequest) CompressionCodec(compressionCodec string) *APICreateLogOpenstackRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogOpenstackRequest) Period(period int32) *APICreateLogOpenstackRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogOpenstackRequest) GzipLevel(gzipLevel int32) *APICreateLogOpenstackRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// AccessKey Your OpenStack account access key.
func (r *APICreateLogOpenstackRequest) AccessKey(accessKey string) *APICreateLogOpenstackRequest {
	r.accessKey = &accessKey
	return r
}

// BucketName The name of your OpenStack container.
func (r *APICreateLogOpenstackRequest) BucketName(bucketName string) *APICreateLogOpenstackRequest {
	r.bucketName = &bucketName
	return r
}

// Path The path to upload logs to.
func (r *APICreateLogOpenstackRequest) Path(path string) *APICreateLogOpenstackRequest {
	r.path = &path
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogOpenstackRequest) PublicKey(publicKey string) *APICreateLogOpenstackRequest {
	r.publicKey = &publicKey
	return r
}

// URL Your OpenStack auth url.
func (r *APICreateLogOpenstackRequest) URL(url string) *APICreateLogOpenstackRequest {
	r.url = &url
	return r
}

// User The username for your OpenStack account.
func (r *APICreateLogOpenstackRequest) User(user string) *APICreateLogOpenstackRequest {
	r.user = &user
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogOpenstackRequest) Execute() (*LoggingOpenstackResponse, *http.Response, error) {
	return r.APIService.CreateLogOpenstackExecute(r)
}

/*
CreateLogOpenstack Create an OpenStack log endpoint

Create a openstack for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogOpenstackRequest
*/
func (a *LoggingOpenstackAPIService) CreateLogOpenstack(ctx context.Context, serviceID string, versionID int32) APICreateLogOpenstackRequest {
	return APICreateLogOpenstackRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogOpenstackExecute executes the request
//  @return LoggingOpenstackResponse
func (a *LoggingOpenstackAPIService) CreateLogOpenstackExecute(r APICreateLogOpenstackRequest) (*LoggingOpenstackResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingOpenstackResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingOpenstackAPIService.CreateLogOpenstack")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/openstack"
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
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
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

// APIDeleteLogOpenstackRequest represents a request for the resource.
type APIDeleteLogOpenstackRequest struct {
	ctx                  context.Context
	APIService           LoggingOpenstackAPI
	serviceID            string
	versionID            int32
	loggingOpenstackName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogOpenstackRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogOpenstackExecute(r)
}

/*
DeleteLogOpenstack Delete an OpenStack log endpoint

Delete the openstack for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingOpenstackName The name for the real-time logging configuration.
 @return APIDeleteLogOpenstackRequest
*/
func (a *LoggingOpenstackAPIService) DeleteLogOpenstack(ctx context.Context, serviceID string, versionID int32, loggingOpenstackName string) APIDeleteLogOpenstackRequest {
	return APIDeleteLogOpenstackRequest{
		APIService:           a,
		ctx:                  ctx,
		serviceID:            serviceID,
		versionID:            versionID,
		loggingOpenstackName: loggingOpenstackName,
	}
}

// DeleteLogOpenstackExecute executes the request
//  @return InlineResponse200
func (a *LoggingOpenstackAPIService) DeleteLogOpenstackExecute(r APIDeleteLogOpenstackRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingOpenstackAPIService.DeleteLogOpenstack")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/openstack/{logging_openstack_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_openstack_name"+"}", gourl.PathEscape(parameterToString(r.loggingOpenstackName, "")))

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

// APIGetLogOpenstackRequest represents a request for the resource.
type APIGetLogOpenstackRequest struct {
	ctx                  context.Context
	APIService           LoggingOpenstackAPI
	serviceID            string
	versionID            int32
	loggingOpenstackName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogOpenstackRequest) Execute() (*LoggingOpenstackResponse, *http.Response, error) {
	return r.APIService.GetLogOpenstackExecute(r)
}

/*
GetLogOpenstack Get an OpenStack log endpoint

Get the openstack for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingOpenstackName The name for the real-time logging configuration.
 @return APIGetLogOpenstackRequest
*/
func (a *LoggingOpenstackAPIService) GetLogOpenstack(ctx context.Context, serviceID string, versionID int32, loggingOpenstackName string) APIGetLogOpenstackRequest {
	return APIGetLogOpenstackRequest{
		APIService:           a,
		ctx:                  ctx,
		serviceID:            serviceID,
		versionID:            versionID,
		loggingOpenstackName: loggingOpenstackName,
	}
}

// GetLogOpenstackExecute executes the request
//  @return LoggingOpenstackResponse
func (a *LoggingOpenstackAPIService) GetLogOpenstackExecute(r APIGetLogOpenstackRequest) (*LoggingOpenstackResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingOpenstackResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingOpenstackAPIService.GetLogOpenstack")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/openstack/{logging_openstack_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_openstack_name"+"}", gourl.PathEscape(parameterToString(r.loggingOpenstackName, "")))

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

// APIListLogOpenstackRequest represents a request for the resource.
type APIListLogOpenstackRequest struct {
	ctx        context.Context
	APIService LoggingOpenstackAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogOpenstackRequest) Execute() ([]LoggingOpenstackResponse, *http.Response, error) {
	return r.APIService.ListLogOpenstackExecute(r)
}

/*
ListLogOpenstack List OpenStack log endpoints

List all of the openstacks for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogOpenstackRequest
*/
func (a *LoggingOpenstackAPIService) ListLogOpenstack(ctx context.Context, serviceID string, versionID int32) APIListLogOpenstackRequest {
	return APIListLogOpenstackRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogOpenstackExecute executes the request
//  @return []LoggingOpenstackResponse
func (a *LoggingOpenstackAPIService) ListLogOpenstackExecute(r APIListLogOpenstackRequest) ([]LoggingOpenstackResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingOpenstackResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingOpenstackAPIService.ListLogOpenstack")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/openstack"
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

// APIUpdateLogOpenstackRequest represents a request for the resource.
type APIUpdateLogOpenstackRequest struct {
	ctx                  context.Context
	APIService           LoggingOpenstackAPI
	serviceID            string
	versionID            int32
	loggingOpenstackName string
	name                 *string
	placement            *string
	responseCondition    *string
	format               *string
	logProcessingRegion  *string
	formatVersion        *int32
	messageType          *string
	timestampFormat      *string
	compressionCodec     *string
	period               *int32
	gzipLevel            *int32
	accessKey            *string
	bucketName           *string
	path                 *string
	publicKey            *string
	url                  *string
	user                 *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogOpenstackRequest) Name(name string) *APIUpdateLogOpenstackRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogOpenstackRequest) Placement(placement string) *APIUpdateLogOpenstackRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogOpenstackRequest) ResponseCondition(responseCondition string) *APIUpdateLogOpenstackRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
func (r *APIUpdateLogOpenstackRequest) Format(format string) *APIUpdateLogOpenstackRequest {
	r.format = &format
	return r
}

// LogProcessingRegion The geographic region where the logs will be processed before streaming. Valid values are &#x60;us&#x60;, &#x60;eu&#x60;, and &#x60;none&#x60; for global.
func (r *APIUpdateLogOpenstackRequest) LogProcessingRegion(logProcessingRegion string) *APIUpdateLogOpenstackRequest {
	r.logProcessingRegion = &logProcessingRegion
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogOpenstackRequest) FormatVersion(formatVersion int32) *APIUpdateLogOpenstackRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APIUpdateLogOpenstackRequest) MessageType(messageType string) *APIUpdateLogOpenstackRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APIUpdateLogOpenstackRequest) TimestampFormat(timestampFormat string) *APIUpdateLogOpenstackRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogOpenstackRequest) CompressionCodec(compressionCodec string) *APIUpdateLogOpenstackRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogOpenstackRequest) Period(period int32) *APIUpdateLogOpenstackRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogOpenstackRequest) GzipLevel(gzipLevel int32) *APIUpdateLogOpenstackRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// AccessKey Your OpenStack account access key.
func (r *APIUpdateLogOpenstackRequest) AccessKey(accessKey string) *APIUpdateLogOpenstackRequest {
	r.accessKey = &accessKey
	return r
}

// BucketName The name of your OpenStack container.
func (r *APIUpdateLogOpenstackRequest) BucketName(bucketName string) *APIUpdateLogOpenstackRequest {
	r.bucketName = &bucketName
	return r
}

// Path The path to upload logs to.
func (r *APIUpdateLogOpenstackRequest) Path(path string) *APIUpdateLogOpenstackRequest {
	r.path = &path
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogOpenstackRequest) PublicKey(publicKey string) *APIUpdateLogOpenstackRequest {
	r.publicKey = &publicKey
	return r
}

// URL Your OpenStack auth url.
func (r *APIUpdateLogOpenstackRequest) URL(url string) *APIUpdateLogOpenstackRequest {
	r.url = &url
	return r
}

// User The username for your OpenStack account.
func (r *APIUpdateLogOpenstackRequest) User(user string) *APIUpdateLogOpenstackRequest {
	r.user = &user
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogOpenstackRequest) Execute() (*LoggingOpenstackResponse, *http.Response, error) {
	return r.APIService.UpdateLogOpenstackExecute(r)
}

/*
UpdateLogOpenstack Update an OpenStack log endpoint

Update the openstack for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingOpenstackName The name for the real-time logging configuration.
 @return APIUpdateLogOpenstackRequest
*/
func (a *LoggingOpenstackAPIService) UpdateLogOpenstack(ctx context.Context, serviceID string, versionID int32, loggingOpenstackName string) APIUpdateLogOpenstackRequest {
	return APIUpdateLogOpenstackRequest{
		APIService:           a,
		ctx:                  ctx,
		serviceID:            serviceID,
		versionID:            versionID,
		loggingOpenstackName: loggingOpenstackName,
	}
}

// UpdateLogOpenstackExecute executes the request
//  @return LoggingOpenstackResponse
func (a *LoggingOpenstackAPIService) UpdateLogOpenstackExecute(r APIUpdateLogOpenstackRequest) (*LoggingOpenstackResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingOpenstackResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingOpenstackAPIService.UpdateLogOpenstack")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/openstack/{logging_openstack_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_openstack_name"+"}", gourl.PathEscape(parameterToString(r.loggingOpenstackName, "")))

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
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.url != nil {
		localVarFormParams.Add("url", parameterToString(*r.url, ""))
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
