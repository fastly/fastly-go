// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

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

// LoggingDigitaloceanAPI defines an interface for interacting with the resource.
type LoggingDigitaloceanAPI interface {

	/*
	CreateLogDigocean Create a DigitalOcean Spaces log endpoint

	Create a DigitalOcean Space for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogDigoceanRequest
	*/
	CreateLogDigocean(ctx context.Context, serviceID string, versionID int32) APICreateLogDigoceanRequest

	// CreateLogDigoceanExecute executes the request
	//  @return LoggingDigitaloceanResponse
	CreateLogDigoceanExecute(r APICreateLogDigoceanRequest) (*LoggingDigitaloceanResponse, *http.Response, error)

	/*
	DeleteLogDigocean Delete a DigitalOcean Spaces log endpoint

	Delete the DigitalOcean Space for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingDigitaloceanName The name for the real-time logging configuration.
	 @return APIDeleteLogDigoceanRequest
	*/
	DeleteLogDigocean(ctx context.Context, serviceID string, versionID int32, loggingDigitaloceanName string) APIDeleteLogDigoceanRequest

	// DeleteLogDigoceanExecute executes the request
	//  @return InlineResponse200
	DeleteLogDigoceanExecute(r APIDeleteLogDigoceanRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogDigocean Get a DigitalOcean Spaces log endpoint

	Get the DigitalOcean Space for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingDigitaloceanName The name for the real-time logging configuration.
	 @return APIGetLogDigoceanRequest
	*/
	GetLogDigocean(ctx context.Context, serviceID string, versionID int32, loggingDigitaloceanName string) APIGetLogDigoceanRequest

	// GetLogDigoceanExecute executes the request
	//  @return LoggingDigitaloceanResponse
	GetLogDigoceanExecute(r APIGetLogDigoceanRequest) (*LoggingDigitaloceanResponse, *http.Response, error)

	/*
	ListLogDigocean List DigitalOcean Spaces log endpoints

	List all of the DigitalOcean Spaces for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogDigoceanRequest
	*/
	ListLogDigocean(ctx context.Context, serviceID string, versionID int32) APIListLogDigoceanRequest

	// ListLogDigoceanExecute executes the request
	//  @return []LoggingDigitaloceanResponse
	ListLogDigoceanExecute(r APIListLogDigoceanRequest) ([]LoggingDigitaloceanResponse, *http.Response, error)

	/*
	UpdateLogDigocean Update a DigitalOcean Spaces log endpoint

	Update the DigitalOcean Space for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingDigitaloceanName The name for the real-time logging configuration.
	 @return APIUpdateLogDigoceanRequest
	*/
	UpdateLogDigocean(ctx context.Context, serviceID string, versionID int32, loggingDigitaloceanName string) APIUpdateLogDigoceanRequest

	// UpdateLogDigoceanExecute executes the request
	//  @return LoggingDigitaloceanResponse
	UpdateLogDigoceanExecute(r APIUpdateLogDigoceanRequest) (*LoggingDigitaloceanResponse, *http.Response, error)
}

// LoggingDigitaloceanAPIService LoggingDigitaloceanAPI service
type LoggingDigitaloceanAPIService service

// APICreateLogDigoceanRequest represents a request for the resource.
type APICreateLogDigoceanRequest struct {
	ctx context.Context
	APIService LoggingDigitaloceanAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	messageType *string
	timestampFormat *string
	compressionCodec *string
	period *int32
	gzipLevel *int32
	bucketName *string
	accessKey *string
	secretKey *string
	domain *string
	path *string
	publicKey *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogDigoceanRequest) Name(name string) *APICreateLogDigoceanRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogDigoceanRequest) Placement(placement string) *APICreateLogDigoceanRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogDigoceanRequest) ResponseCondition(responseCondition string) *APICreateLogDigoceanRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogDigoceanRequest) Format(format string) *APICreateLogDigoceanRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogDigoceanRequest) FormatVersion(formatVersion int32) *APICreateLogDigoceanRequest {
	r.formatVersion = &formatVersion
	return r
}
// MessageType How the message should be formatted.
func (r *APICreateLogDigoceanRequest) MessageType(messageType string) *APICreateLogDigoceanRequest {
	r.messageType = &messageType
	return r
}
// TimestampFormat A timestamp format
func (r *APICreateLogDigoceanRequest) TimestampFormat(timestampFormat string) *APICreateLogDigoceanRequest {
	r.timestampFormat = &timestampFormat
	return r
}
// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogDigoceanRequest) CompressionCodec(compressionCodec string) *APICreateLogDigoceanRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogDigoceanRequest) Period(period int32) *APICreateLogDigoceanRequest {
	r.period = &period
	return r
}
// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogDigoceanRequest) GzipLevel(gzipLevel int32) *APICreateLogDigoceanRequest {
	r.gzipLevel = &gzipLevel
	return r
}
// BucketName The name of the DigitalOcean Space.
func (r *APICreateLogDigoceanRequest) BucketName(bucketName string) *APICreateLogDigoceanRequest {
	r.bucketName = &bucketName
	return r
}
// AccessKey Your DigitalOcean Spaces account access key.
func (r *APICreateLogDigoceanRequest) AccessKey(accessKey string) *APICreateLogDigoceanRequest {
	r.accessKey = &accessKey
	return r
}
// SecretKey Your DigitalOcean Spaces account secret key.
func (r *APICreateLogDigoceanRequest) SecretKey(secretKey string) *APICreateLogDigoceanRequest {
	r.secretKey = &secretKey
	return r
}
// Domain The domain of the DigitalOcean Spaces endpoint.
func (r *APICreateLogDigoceanRequest) Domain(domain string) *APICreateLogDigoceanRequest {
	r.domain = &domain
	return r
}
// Path The path to upload logs to.
func (r *APICreateLogDigoceanRequest) Path(path string) *APICreateLogDigoceanRequest {
	r.path = &path
	return r
}
// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogDigoceanRequest) PublicKey(publicKey string) *APICreateLogDigoceanRequest {
	r.publicKey = &publicKey
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogDigoceanRequest) Execute() (*LoggingDigitaloceanResponse, *http.Response, error) {
	return r.APIService.CreateLogDigoceanExecute(r)
}

/*
CreateLogDigocean Create a DigitalOcean Spaces log endpoint

Create a DigitalOcean Space for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogDigoceanRequest
*/
func (a *LoggingDigitaloceanAPIService) CreateLogDigocean(ctx context.Context, serviceID string, versionID int32) APICreateLogDigoceanRequest {
	return APICreateLogDigoceanRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogDigoceanExecute executes the request
//  @return LoggingDigitaloceanResponse
func (a *LoggingDigitaloceanAPIService) CreateLogDigoceanExecute(r APICreateLogDigoceanRequest) (*LoggingDigitaloceanResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingDigitaloceanResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingDigitaloceanAPIService.CreateLogDigocean")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/digitalocean"
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
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.domain != nil {
		localVarFormParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
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

// APIDeleteLogDigoceanRequest represents a request for the resource.
type APIDeleteLogDigoceanRequest struct {
	ctx context.Context
	APIService LoggingDigitaloceanAPI
	serviceID string
	versionID int32
	loggingDigitaloceanName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogDigoceanRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogDigoceanExecute(r)
}

/*
DeleteLogDigocean Delete a DigitalOcean Spaces log endpoint

Delete the DigitalOcean Space for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingDigitaloceanName The name for the real-time logging configuration.
 @return APIDeleteLogDigoceanRequest
*/
func (a *LoggingDigitaloceanAPIService) DeleteLogDigocean(ctx context.Context, serviceID string, versionID int32, loggingDigitaloceanName string) APIDeleteLogDigoceanRequest {
	return APIDeleteLogDigoceanRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingDigitaloceanName: loggingDigitaloceanName,
	}
}

// DeleteLogDigoceanExecute executes the request
//  @return InlineResponse200
func (a *LoggingDigitaloceanAPIService) DeleteLogDigoceanExecute(r APIDeleteLogDigoceanRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingDigitaloceanAPIService.DeleteLogDigocean")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_digitalocean_name"+"}", gourl.PathEscape(parameterToString(r.loggingDigitaloceanName, "")))

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

// APIGetLogDigoceanRequest represents a request for the resource.
type APIGetLogDigoceanRequest struct {
	ctx context.Context
	APIService LoggingDigitaloceanAPI
	serviceID string
	versionID int32
	loggingDigitaloceanName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogDigoceanRequest) Execute() (*LoggingDigitaloceanResponse, *http.Response, error) {
	return r.APIService.GetLogDigoceanExecute(r)
}

/*
GetLogDigocean Get a DigitalOcean Spaces log endpoint

Get the DigitalOcean Space for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingDigitaloceanName The name for the real-time logging configuration.
 @return APIGetLogDigoceanRequest
*/
func (a *LoggingDigitaloceanAPIService) GetLogDigocean(ctx context.Context, serviceID string, versionID int32, loggingDigitaloceanName string) APIGetLogDigoceanRequest {
	return APIGetLogDigoceanRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingDigitaloceanName: loggingDigitaloceanName,
	}
}

// GetLogDigoceanExecute executes the request
//  @return LoggingDigitaloceanResponse
func (a *LoggingDigitaloceanAPIService) GetLogDigoceanExecute(r APIGetLogDigoceanRequest) (*LoggingDigitaloceanResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingDigitaloceanResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingDigitaloceanAPIService.GetLogDigocean")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_digitalocean_name"+"}", gourl.PathEscape(parameterToString(r.loggingDigitaloceanName, "")))

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

// APIListLogDigoceanRequest represents a request for the resource.
type APIListLogDigoceanRequest struct {
	ctx context.Context
	APIService LoggingDigitaloceanAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogDigoceanRequest) Execute() ([]LoggingDigitaloceanResponse, *http.Response, error) {
	return r.APIService.ListLogDigoceanExecute(r)
}

/*
ListLogDigocean List DigitalOcean Spaces log endpoints

List all of the DigitalOcean Spaces for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogDigoceanRequest
*/
func (a *LoggingDigitaloceanAPIService) ListLogDigocean(ctx context.Context, serviceID string, versionID int32) APIListLogDigoceanRequest {
	return APIListLogDigoceanRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogDigoceanExecute executes the request
//  @return []LoggingDigitaloceanResponse
func (a *LoggingDigitaloceanAPIService) ListLogDigoceanExecute(r APIListLogDigoceanRequest) ([]LoggingDigitaloceanResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingDigitaloceanResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingDigitaloceanAPIService.ListLogDigocean")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/digitalocean"
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

// APIUpdateLogDigoceanRequest represents a request for the resource.
type APIUpdateLogDigoceanRequest struct {
	ctx context.Context
	APIService LoggingDigitaloceanAPI
	serviceID string
	versionID int32
	loggingDigitaloceanName string
	name *string
	placement *string
	responseCondition *string
	format *string
	formatVersion *int32
	messageType *string
	timestampFormat *string
	compressionCodec *string
	period *int32
	gzipLevel *int32
	bucketName *string
	accessKey *string
	secretKey *string
	domain *string
	path *string
	publicKey *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogDigoceanRequest) Name(name string) *APIUpdateLogDigoceanRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogDigoceanRequest) Placement(placement string) *APIUpdateLogDigoceanRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogDigoceanRequest) ResponseCondition(responseCondition string) *APIUpdateLogDigoceanRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogDigoceanRequest) Format(format string) *APIUpdateLogDigoceanRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogDigoceanRequest) FormatVersion(formatVersion int32) *APIUpdateLogDigoceanRequest {
	r.formatVersion = &formatVersion
	return r
}
// MessageType How the message should be formatted.
func (r *APIUpdateLogDigoceanRequest) MessageType(messageType string) *APIUpdateLogDigoceanRequest {
	r.messageType = &messageType
	return r
}
// TimestampFormat A timestamp format
func (r *APIUpdateLogDigoceanRequest) TimestampFormat(timestampFormat string) *APIUpdateLogDigoceanRequest {
	r.timestampFormat = &timestampFormat
	return r
}
// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogDigoceanRequest) CompressionCodec(compressionCodec string) *APIUpdateLogDigoceanRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogDigoceanRequest) Period(period int32) *APIUpdateLogDigoceanRequest {
	r.period = &period
	return r
}
// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogDigoceanRequest) GzipLevel(gzipLevel int32) *APIUpdateLogDigoceanRequest {
	r.gzipLevel = &gzipLevel
	return r
}
// BucketName The name of the DigitalOcean Space.
func (r *APIUpdateLogDigoceanRequest) BucketName(bucketName string) *APIUpdateLogDigoceanRequest {
	r.bucketName = &bucketName
	return r
}
// AccessKey Your DigitalOcean Spaces account access key.
func (r *APIUpdateLogDigoceanRequest) AccessKey(accessKey string) *APIUpdateLogDigoceanRequest {
	r.accessKey = &accessKey
	return r
}
// SecretKey Your DigitalOcean Spaces account secret key.
func (r *APIUpdateLogDigoceanRequest) SecretKey(secretKey string) *APIUpdateLogDigoceanRequest {
	r.secretKey = &secretKey
	return r
}
// Domain The domain of the DigitalOcean Spaces endpoint.
func (r *APIUpdateLogDigoceanRequest) Domain(domain string) *APIUpdateLogDigoceanRequest {
	r.domain = &domain
	return r
}
// Path The path to upload logs to.
func (r *APIUpdateLogDigoceanRequest) Path(path string) *APIUpdateLogDigoceanRequest {
	r.path = &path
	return r
}
// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogDigoceanRequest) PublicKey(publicKey string) *APIUpdateLogDigoceanRequest {
	r.publicKey = &publicKey
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogDigoceanRequest) Execute() (*LoggingDigitaloceanResponse, *http.Response, error) {
	return r.APIService.UpdateLogDigoceanExecute(r)
}

/*
UpdateLogDigocean Update a DigitalOcean Spaces log endpoint

Update the DigitalOcean Space for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingDigitaloceanName The name for the real-time logging configuration.
 @return APIUpdateLogDigoceanRequest
*/
func (a *LoggingDigitaloceanAPIService) UpdateLogDigocean(ctx context.Context, serviceID string, versionID int32, loggingDigitaloceanName string) APIUpdateLogDigoceanRequest {
	return APIUpdateLogDigoceanRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingDigitaloceanName: loggingDigitaloceanName,
	}
}

// UpdateLogDigoceanExecute executes the request
//  @return LoggingDigitaloceanResponse
func (a *LoggingDigitaloceanAPIService) UpdateLogDigoceanExecute(r APIUpdateLogDigoceanRequest) (*LoggingDigitaloceanResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingDigitaloceanResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingDigitaloceanAPIService.UpdateLogDigocean")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_digitalocean_name"+"}", gourl.PathEscape(parameterToString(r.loggingDigitaloceanName, "")))

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
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.domain != nil {
		localVarFormParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
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
