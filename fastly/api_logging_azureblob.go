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

// LoggingAzureblobAPI defines an interface for interacting with the resource.
type LoggingAzureblobAPI interface {

	/*
		CreateLogAzure Create an Azure Blob Storage log endpoint

		Create an Azure Blob Storage logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogAzureRequest
	*/
	CreateLogAzure(ctx context.Context, serviceID string, versionID int32) APICreateLogAzureRequest

	// CreateLogAzureExecute executes the request
	//  @return LoggingAzureblobResponse
	CreateLogAzureExecute(r APICreateLogAzureRequest) (*LoggingAzureblobResponse, *http.Response, error)

	/*
		DeleteLogAzure Delete the Azure Blob Storage log endpoint

		Delete the Azure Blob Storage logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingAzureblobName The name for the real-time logging configuration.
		 @return APIDeleteLogAzureRequest
	*/
	DeleteLogAzure(ctx context.Context, serviceID string, versionID int32, loggingAzureblobName string) APIDeleteLogAzureRequest

	// DeleteLogAzureExecute executes the request
	//  @return InlineResponse200
	DeleteLogAzureExecute(r APIDeleteLogAzureRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetLogAzure Get an Azure Blob Storage log endpoint

		Get the Azure Blob Storage logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingAzureblobName The name for the real-time logging configuration.
		 @return APIGetLogAzureRequest
	*/
	GetLogAzure(ctx context.Context, serviceID string, versionID int32, loggingAzureblobName string) APIGetLogAzureRequest

	// GetLogAzureExecute executes the request
	//  @return LoggingAzureblobResponse
	GetLogAzureExecute(r APIGetLogAzureRequest) (*LoggingAzureblobResponse, *http.Response, error)

	/*
		ListLogAzure List Azure Blob Storage log endpoints

		List all of the Azure Blob Storage logging endpoints for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogAzureRequest
	*/
	ListLogAzure(ctx context.Context, serviceID string, versionID int32) APIListLogAzureRequest

	// ListLogAzureExecute executes the request
	//  @return []LoggingAzureblobResponse
	ListLogAzureExecute(r APIListLogAzureRequest) ([]LoggingAzureblobResponse, *http.Response, error)

	/*
		UpdateLogAzure Update an Azure Blob Storage log endpoint

		Update the Azure Blob Storage logging endpoint for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingAzureblobName The name for the real-time logging configuration.
		 @return APIUpdateLogAzureRequest
	*/
	UpdateLogAzure(ctx context.Context, serviceID string, versionID int32, loggingAzureblobName string) APIUpdateLogAzureRequest

	// UpdateLogAzureExecute executes the request
	//  @return LoggingAzureblobResponse
	UpdateLogAzureExecute(r APIUpdateLogAzureRequest) (*LoggingAzureblobResponse, *http.Response, error)
}

// LoggingAzureblobAPIService LoggingAzureblobAPI service
type LoggingAzureblobAPIService service

// APICreateLogAzureRequest represents a request for the resource.
type APICreateLogAzureRequest struct {
	ctx               context.Context
	APIService        LoggingAzureblobAPI
	serviceID         string
	versionID         int32
	name              *string
	placement         *string
	responseCondition *string
	format            *string
	formatVersion     *int32
	messageType       *string
	timestampFormat   *string
	compressionCodec  *string
	period            *int32
	gzipLevel         *int32
	path              *string
	accountName       *string
	container         *string
	sasToken          *string
	publicKey         *string
	fileMaxBytes      *int32
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogAzureRequest) Name(name string) *APICreateLogAzureRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogAzureRequest) Placement(placement string) *APICreateLogAzureRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogAzureRequest) ResponseCondition(responseCondition string) *APICreateLogAzureRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogAzureRequest) Format(format string) *APICreateLogAzureRequest {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogAzureRequest) FormatVersion(formatVersion int32) *APICreateLogAzureRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APICreateLogAzureRequest) MessageType(messageType string) *APICreateLogAzureRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APICreateLogAzureRequest) TimestampFormat(timestampFormat string) *APICreateLogAzureRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogAzureRequest) CompressionCodec(compressionCodec string) *APICreateLogAzureRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogAzureRequest) Period(period int32) *APICreateLogAzureRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogAzureRequest) GzipLevel(gzipLevel int32) *APICreateLogAzureRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// Path The path to upload logs to.
func (r *APICreateLogAzureRequest) Path(path string) *APICreateLogAzureRequest {
	r.path = &path
	return r
}

// AccountName The unique Azure Blob Storage namespace in which your data objects are stored. Required.
func (r *APICreateLogAzureRequest) AccountName(accountName string) *APICreateLogAzureRequest {
	r.accountName = &accountName
	return r
}

// Container The name of the Azure Blob Storage container in which to store logs. Required.
func (r *APICreateLogAzureRequest) Container(container string) *APICreateLogAzureRequest {
	r.container = &container
	return r
}

// SasToken The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required.
func (r *APICreateLogAzureRequest) SasToken(sasToken string) *APICreateLogAzureRequest {
	r.sasToken = &sasToken
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogAzureRequest) PublicKey(publicKey string) *APICreateLogAzureRequest {
	r.publicKey = &publicKey
	return r
}

// FileMaxBytes The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
func (r *APICreateLogAzureRequest) FileMaxBytes(fileMaxBytes int32) *APICreateLogAzureRequest {
	r.fileMaxBytes = &fileMaxBytes
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogAzureRequest) Execute() (*LoggingAzureblobResponse, *http.Response, error) {
	return r.APIService.CreateLogAzureExecute(r)
}

/*
CreateLogAzure Create an Azure Blob Storage log endpoint

Create an Azure Blob Storage logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogAzureRequest
*/
func (a *LoggingAzureblobAPIService) CreateLogAzure(ctx context.Context, serviceID string, versionID int32) APICreateLogAzureRequest {
	return APICreateLogAzureRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogAzureExecute executes the request
//  @return LoggingAzureblobResponse
func (a *LoggingAzureblobAPIService) CreateLogAzureExecute(r APICreateLogAzureRequest) (*LoggingAzureblobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingAzureblobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingAzureblobAPIService.CreateLogAzure")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/azureblob"
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
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.accountName != nil {
		localVarFormParams.Add("account_name", parameterToString(*r.accountName, ""))
	}
	if r.container != nil {
		localVarFormParams.Add("container", parameterToString(*r.container, ""))
	}
	if r.sasToken != nil {
		localVarFormParams.Add("sas_token", parameterToString(*r.sasToken, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.fileMaxBytes != nil {
		localVarFormParams.Add("file_max_bytes", parameterToString(*r.fileMaxBytes, ""))
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

// APIDeleteLogAzureRequest represents a request for the resource.
type APIDeleteLogAzureRequest struct {
	ctx                  context.Context
	APIService           LoggingAzureblobAPI
	serviceID            string
	versionID            int32
	loggingAzureblobName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogAzureRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogAzureExecute(r)
}

/*
DeleteLogAzure Delete the Azure Blob Storage log endpoint

Delete the Azure Blob Storage logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingAzureblobName The name for the real-time logging configuration.
 @return APIDeleteLogAzureRequest
*/
func (a *LoggingAzureblobAPIService) DeleteLogAzure(ctx context.Context, serviceID string, versionID int32, loggingAzureblobName string) APIDeleteLogAzureRequest {
	return APIDeleteLogAzureRequest{
		APIService:           a,
		ctx:                  ctx,
		serviceID:            serviceID,
		versionID:            versionID,
		loggingAzureblobName: loggingAzureblobName,
	}
}

// DeleteLogAzureExecute executes the request
//  @return InlineResponse200
func (a *LoggingAzureblobAPIService) DeleteLogAzureExecute(r APIDeleteLogAzureRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingAzureblobAPIService.DeleteLogAzure")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_azureblob_name"+"}", gourl.PathEscape(parameterToString(r.loggingAzureblobName, "")))

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

// APIGetLogAzureRequest represents a request for the resource.
type APIGetLogAzureRequest struct {
	ctx                  context.Context
	APIService           LoggingAzureblobAPI
	serviceID            string
	versionID            int32
	loggingAzureblobName string
}

// Execute calls the API using the request data configured.
func (r APIGetLogAzureRequest) Execute() (*LoggingAzureblobResponse, *http.Response, error) {
	return r.APIService.GetLogAzureExecute(r)
}

/*
GetLogAzure Get an Azure Blob Storage log endpoint

Get the Azure Blob Storage logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingAzureblobName The name for the real-time logging configuration.
 @return APIGetLogAzureRequest
*/
func (a *LoggingAzureblobAPIService) GetLogAzure(ctx context.Context, serviceID string, versionID int32, loggingAzureblobName string) APIGetLogAzureRequest {
	return APIGetLogAzureRequest{
		APIService:           a,
		ctx:                  ctx,
		serviceID:            serviceID,
		versionID:            versionID,
		loggingAzureblobName: loggingAzureblobName,
	}
}

// GetLogAzureExecute executes the request
//  @return LoggingAzureblobResponse
func (a *LoggingAzureblobAPIService) GetLogAzureExecute(r APIGetLogAzureRequest) (*LoggingAzureblobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingAzureblobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingAzureblobAPIService.GetLogAzure")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_azureblob_name"+"}", gourl.PathEscape(parameterToString(r.loggingAzureblobName, "")))

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

// APIListLogAzureRequest represents a request for the resource.
type APIListLogAzureRequest struct {
	ctx        context.Context
	APIService LoggingAzureblobAPI
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogAzureRequest) Execute() ([]LoggingAzureblobResponse, *http.Response, error) {
	return r.APIService.ListLogAzureExecute(r)
}

/*
ListLogAzure List Azure Blob Storage log endpoints

List all of the Azure Blob Storage logging endpoints for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogAzureRequest
*/
func (a *LoggingAzureblobAPIService) ListLogAzure(ctx context.Context, serviceID string, versionID int32) APIListLogAzureRequest {
	return APIListLogAzureRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogAzureExecute executes the request
//  @return []LoggingAzureblobResponse
func (a *LoggingAzureblobAPIService) ListLogAzureExecute(r APIListLogAzureRequest) ([]LoggingAzureblobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingAzureblobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingAzureblobAPIService.ListLogAzure")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/azureblob"
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

// APIUpdateLogAzureRequest represents a request for the resource.
type APIUpdateLogAzureRequest struct {
	ctx                  context.Context
	APIService           LoggingAzureblobAPI
	serviceID            string
	versionID            int32
	loggingAzureblobName string
	name                 *string
	placement            *string
	responseCondition    *string
	format               *string
	formatVersion        *int32
	messageType          *string
	timestampFormat      *string
	compressionCodec     *string
	period               *int32
	gzipLevel            *int32
	path                 *string
	accountName          *string
	container            *string
	sasToken             *string
	publicKey            *string
	fileMaxBytes         *int32
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogAzureRequest) Name(name string) *APIUpdateLogAzureRequest {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogAzureRequest) Placement(placement string) *APIUpdateLogAzureRequest {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogAzureRequest) ResponseCondition(responseCondition string) *APIUpdateLogAzureRequest {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogAzureRequest) Format(format string) *APIUpdateLogAzureRequest {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogAzureRequest) FormatVersion(formatVersion int32) *APIUpdateLogAzureRequest {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APIUpdateLogAzureRequest) MessageType(messageType string) *APIUpdateLogAzureRequest {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APIUpdateLogAzureRequest) TimestampFormat(timestampFormat string) *APIUpdateLogAzureRequest {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogAzureRequest) CompressionCodec(compressionCodec string) *APIUpdateLogAzureRequest {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogAzureRequest) Period(period int32) *APIUpdateLogAzureRequest {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogAzureRequest) GzipLevel(gzipLevel int32) *APIUpdateLogAzureRequest {
	r.gzipLevel = &gzipLevel
	return r
}

// Path The path to upload logs to.
func (r *APIUpdateLogAzureRequest) Path(path string) *APIUpdateLogAzureRequest {
	r.path = &path
	return r
}

// AccountName The unique Azure Blob Storage namespace in which your data objects are stored. Required.
func (r *APIUpdateLogAzureRequest) AccountName(accountName string) *APIUpdateLogAzureRequest {
	r.accountName = &accountName
	return r
}

// Container The name of the Azure Blob Storage container in which to store logs. Required.
func (r *APIUpdateLogAzureRequest) Container(container string) *APIUpdateLogAzureRequest {
	r.container = &container
	return r
}

// SasToken The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required.
func (r *APIUpdateLogAzureRequest) SasToken(sasToken string) *APIUpdateLogAzureRequest {
	r.sasToken = &sasToken
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogAzureRequest) PublicKey(publicKey string) *APIUpdateLogAzureRequest {
	r.publicKey = &publicKey
	return r
}

// FileMaxBytes The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
func (r *APIUpdateLogAzureRequest) FileMaxBytes(fileMaxBytes int32) *APIUpdateLogAzureRequest {
	r.fileMaxBytes = &fileMaxBytes
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogAzureRequest) Execute() (*LoggingAzureblobResponse, *http.Response, error) {
	return r.APIService.UpdateLogAzureExecute(r)
}

/*
UpdateLogAzure Update an Azure Blob Storage log endpoint

Update the Azure Blob Storage logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingAzureblobName The name for the real-time logging configuration.
 @return APIUpdateLogAzureRequest
*/
func (a *LoggingAzureblobAPIService) UpdateLogAzure(ctx context.Context, serviceID string, versionID int32, loggingAzureblobName string) APIUpdateLogAzureRequest {
	return APIUpdateLogAzureRequest{
		APIService:           a,
		ctx:                  ctx,
		serviceID:            serviceID,
		versionID:            versionID,
		loggingAzureblobName: loggingAzureblobName,
	}
}

// UpdateLogAzureExecute executes the request
//  @return LoggingAzureblobResponse
func (a *LoggingAzureblobAPIService) UpdateLogAzureExecute(r APIUpdateLogAzureRequest) (*LoggingAzureblobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingAzureblobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingAzureblobAPIService.UpdateLogAzure")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_azureblob_name"+"}", gourl.PathEscape(parameterToString(r.loggingAzureblobName, "")))

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
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.accountName != nil {
		localVarFormParams.Add("account_name", parameterToString(*r.accountName, ""))
	}
	if r.container != nil {
		localVarFormParams.Add("container", parameterToString(*r.container, ""))
	}
	if r.sasToken != nil {
		localVarFormParams.Add("sas_token", parameterToString(*r.sasToken, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.fileMaxBytes != nil {
		localVarFormParams.Add("file_max_bytes", parameterToString(*r.fileMaxBytes, ""))
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
