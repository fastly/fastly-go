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

// LoggingFtpAPI defines an interface for interacting with the resource.
type LoggingFtpAPI interface {

	/*
	CreateLogFtp Create an FTP log endpoint

	Create a FTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogFtpRequest
	*/
	CreateLogFtp(ctx context.Context, serviceID string, versionID int32) APICreateLogFtpRequest

	// CreateLogFtpExecute executes the request
	//  @return LoggingFtpResponse
	CreateLogFtpExecute(r APICreateLogFtpRequest) (*LoggingFtpResponse, *http.Response, error)

	/*
	DeleteLogFtp Delete an FTP log endpoint

	Delete the FTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingFtpName The name for the real-time logging configuration.
	 @return APIDeleteLogFtpRequest
	*/
	DeleteLogFtp(ctx context.Context, serviceID string, versionID int32, loggingFtpName string) APIDeleteLogFtpRequest

	// DeleteLogFtpExecute executes the request
	//  @return InlineResponse200
	DeleteLogFtpExecute(r APIDeleteLogFtpRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogFtp Get an FTP log endpoint

	Get the FTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingFtpName The name for the real-time logging configuration.
	 @return APIGetLogFtpRequest
	*/
	GetLogFtp(ctx context.Context, serviceID string, versionID int32, loggingFtpName string) APIGetLogFtpRequest

	// GetLogFtpExecute executes the request
	//  @return LoggingFtpResponse
	GetLogFtpExecute(r APIGetLogFtpRequest) (*LoggingFtpResponse, *http.Response, error)

	/*
	ListLogFtp List FTP log endpoints

	List all of the FTPs for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogFtpRequest
	*/
	ListLogFtp(ctx context.Context, serviceID string, versionID int32) APIListLogFtpRequest

	// ListLogFtpExecute executes the request
	//  @return []LoggingFtpResponse
	ListLogFtpExecute(r APIListLogFtpRequest) ([]LoggingFtpResponse, *http.Response, error)

	/*
	UpdateLogFtp Update an FTP log endpoint

	Update the FTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingFtpName The name for the real-time logging configuration.
	 @return APIUpdateLogFtpRequest
	*/
	UpdateLogFtp(ctx context.Context, serviceID string, versionID int32, loggingFtpName string) APIUpdateLogFtpRequest

	// UpdateLogFtpExecute executes the request
	//  @return LoggingFtpResponse
	UpdateLogFtpExecute(r APIUpdateLogFtpRequest) (*LoggingFtpResponse, *http.Response, error)
}

// LoggingFtpAPIService LoggingFtpAPI service
type LoggingFtpAPIService service

// APICreateLogFtpRequest represents a request for the resource.
type APICreateLogFtpRequest struct {
	ctx context.Context
	APIService LoggingFtpAPI
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
	address *string
	hostname *string
	ipv4 *string
	password *string
	path *string
	publicKey *string
	user *string
	port *int32
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogFtpRequest) Name(name string) *APICreateLogFtpRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogFtpRequest) Placement(placement string) *APICreateLogFtpRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogFtpRequest) ResponseCondition(responseCondition string) *APICreateLogFtpRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogFtpRequest) Format(format string) *APICreateLogFtpRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogFtpRequest) FormatVersion(formatVersion int32) *APICreateLogFtpRequest {
	r.formatVersion = &formatVersion
	return r
}
// MessageType How the message should be formatted.
func (r *APICreateLogFtpRequest) MessageType(messageType string) *APICreateLogFtpRequest {
	r.messageType = &messageType
	return r
}
// TimestampFormat A timestamp format
func (r *APICreateLogFtpRequest) TimestampFormat(timestampFormat string) *APICreateLogFtpRequest {
	r.timestampFormat = &timestampFormat
	return r
}
// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogFtpRequest) CompressionCodec(compressionCodec string) *APICreateLogFtpRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogFtpRequest) Period(period int32) *APICreateLogFtpRequest {
	r.period = &period
	return r
}
// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogFtpRequest) GzipLevel(gzipLevel int32) *APICreateLogFtpRequest {
	r.gzipLevel = &gzipLevel
	return r
}
// Address An hostname or IPv4 address.
func (r *APICreateLogFtpRequest) Address(address string) *APICreateLogFtpRequest {
	r.address = &address
	return r
}
// Hostname Hostname used.
func (r *APICreateLogFtpRequest) Hostname(hostname string) *APICreateLogFtpRequest {
	r.hostname = &hostname
	return r
}
// Ipv4 IPv4 address of the host.
func (r *APICreateLogFtpRequest) Ipv4(ipv4 string) *APICreateLogFtpRequest {
	r.ipv4 = &ipv4
	return r
}
// Password The password for the server. For anonymous use an email address.
func (r *APICreateLogFtpRequest) Password(password string) *APICreateLogFtpRequest {
	r.password = &password
	return r
}
// Path The path to upload log files to. If the path ends in &#x60;/&#x60; then it is treated as a directory.
func (r *APICreateLogFtpRequest) Path(path string) *APICreateLogFtpRequest {
	r.path = &path
	return r
}
// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogFtpRequest) PublicKey(publicKey string) *APICreateLogFtpRequest {
	r.publicKey = &publicKey
	return r
}
// User The username for the server. Can be anonymous.
func (r *APICreateLogFtpRequest) User(user string) *APICreateLogFtpRequest {
	r.user = &user
	return r
}
// Port The port number.
func (r *APICreateLogFtpRequest) Port(port int32) *APICreateLogFtpRequest {
	r.port = &port
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogFtpRequest) Execute() (*LoggingFtpResponse, *http.Response, error) {
	return r.APIService.CreateLogFtpExecute(r)
}

/*
CreateLogFtp Create an FTP log endpoint

Create a FTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogFtpRequest
*/
func (a *LoggingFtpAPIService) CreateLogFtp(ctx context.Context, serviceID string, versionID int32) APICreateLogFtpRequest {
	return APICreateLogFtpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogFtpExecute executes the request
//  @return LoggingFtpResponse
func (a *LoggingFtpAPIService) CreateLogFtpExecute(r APICreateLogFtpRequest) (*LoggingFtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingFtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingFtpAPIService.CreateLogFtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/ftp"
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
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.hostname != nil {
		localVarFormParams.Add("hostname", parameterToString(*r.hostname, ""))
	}
	if r.ipv4 != nil {
		localVarFormParams.Add("ipv4", parameterToString(*r.ipv4, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
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

// APIDeleteLogFtpRequest represents a request for the resource.
type APIDeleteLogFtpRequest struct {
	ctx context.Context
	APIService LoggingFtpAPI
	serviceID string
	versionID int32
	loggingFtpName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogFtpRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogFtpExecute(r)
}

/*
DeleteLogFtp Delete an FTP log endpoint

Delete the FTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingFtpName The name for the real-time logging configuration.
 @return APIDeleteLogFtpRequest
*/
func (a *LoggingFtpAPIService) DeleteLogFtp(ctx context.Context, serviceID string, versionID int32, loggingFtpName string) APIDeleteLogFtpRequest {
	return APIDeleteLogFtpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingFtpName: loggingFtpName,
	}
}

// DeleteLogFtpExecute executes the request
//  @return InlineResponse200
func (a *LoggingFtpAPIService) DeleteLogFtpExecute(r APIDeleteLogFtpRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingFtpAPIService.DeleteLogFtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/ftp/{logging_ftp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_ftp_name"+"}", gourl.PathEscape(parameterToString(r.loggingFtpName, "")))

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

// APIGetLogFtpRequest represents a request for the resource.
type APIGetLogFtpRequest struct {
	ctx context.Context
	APIService LoggingFtpAPI
	serviceID string
	versionID int32
	loggingFtpName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogFtpRequest) Execute() (*LoggingFtpResponse, *http.Response, error) {
	return r.APIService.GetLogFtpExecute(r)
}

/*
GetLogFtp Get an FTP log endpoint

Get the FTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingFtpName The name for the real-time logging configuration.
 @return APIGetLogFtpRequest
*/
func (a *LoggingFtpAPIService) GetLogFtp(ctx context.Context, serviceID string, versionID int32, loggingFtpName string) APIGetLogFtpRequest {
	return APIGetLogFtpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingFtpName: loggingFtpName,
	}
}

// GetLogFtpExecute executes the request
//  @return LoggingFtpResponse
func (a *LoggingFtpAPIService) GetLogFtpExecute(r APIGetLogFtpRequest) (*LoggingFtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingFtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingFtpAPIService.GetLogFtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/ftp/{logging_ftp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_ftp_name"+"}", gourl.PathEscape(parameterToString(r.loggingFtpName, "")))

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

// APIListLogFtpRequest represents a request for the resource.
type APIListLogFtpRequest struct {
	ctx context.Context
	APIService LoggingFtpAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogFtpRequest) Execute() ([]LoggingFtpResponse, *http.Response, error) {
	return r.APIService.ListLogFtpExecute(r)
}

/*
ListLogFtp List FTP log endpoints

List all of the FTPs for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogFtpRequest
*/
func (a *LoggingFtpAPIService) ListLogFtp(ctx context.Context, serviceID string, versionID int32) APIListLogFtpRequest {
	return APIListLogFtpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogFtpExecute executes the request
//  @return []LoggingFtpResponse
func (a *LoggingFtpAPIService) ListLogFtpExecute(r APIListLogFtpRequest) ([]LoggingFtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingFtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingFtpAPIService.ListLogFtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/ftp"
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

// APIUpdateLogFtpRequest represents a request for the resource.
type APIUpdateLogFtpRequest struct {
	ctx context.Context
	APIService LoggingFtpAPI
	serviceID string
	versionID int32
	loggingFtpName string
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
	address *string
	hostname *string
	ipv4 *string
	password *string
	path *string
	publicKey *string
	user *string
	port *int32
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogFtpRequest) Name(name string) *APIUpdateLogFtpRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogFtpRequest) Placement(placement string) *APIUpdateLogFtpRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogFtpRequest) ResponseCondition(responseCondition string) *APIUpdateLogFtpRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogFtpRequest) Format(format string) *APIUpdateLogFtpRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogFtpRequest) FormatVersion(formatVersion int32) *APIUpdateLogFtpRequest {
	r.formatVersion = &formatVersion
	return r
}
// MessageType How the message should be formatted.
func (r *APIUpdateLogFtpRequest) MessageType(messageType string) *APIUpdateLogFtpRequest {
	r.messageType = &messageType
	return r
}
// TimestampFormat A timestamp format
func (r *APIUpdateLogFtpRequest) TimestampFormat(timestampFormat string) *APIUpdateLogFtpRequest {
	r.timestampFormat = &timestampFormat
	return r
}
// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogFtpRequest) CompressionCodec(compressionCodec string) *APIUpdateLogFtpRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogFtpRequest) Period(period int32) *APIUpdateLogFtpRequest {
	r.period = &period
	return r
}
// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogFtpRequest) GzipLevel(gzipLevel int32) *APIUpdateLogFtpRequest {
	r.gzipLevel = &gzipLevel
	return r
}
// Address An hostname or IPv4 address.
func (r *APIUpdateLogFtpRequest) Address(address string) *APIUpdateLogFtpRequest {
	r.address = &address
	return r
}
// Hostname Hostname used.
func (r *APIUpdateLogFtpRequest) Hostname(hostname string) *APIUpdateLogFtpRequest {
	r.hostname = &hostname
	return r
}
// Ipv4 IPv4 address of the host.
func (r *APIUpdateLogFtpRequest) Ipv4(ipv4 string) *APIUpdateLogFtpRequest {
	r.ipv4 = &ipv4
	return r
}
// Password The password for the server. For anonymous use an email address.
func (r *APIUpdateLogFtpRequest) Password(password string) *APIUpdateLogFtpRequest {
	r.password = &password
	return r
}
// Path The path to upload log files to. If the path ends in &#x60;/&#x60; then it is treated as a directory.
func (r *APIUpdateLogFtpRequest) Path(path string) *APIUpdateLogFtpRequest {
	r.path = &path
	return r
}
// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogFtpRequest) PublicKey(publicKey string) *APIUpdateLogFtpRequest {
	r.publicKey = &publicKey
	return r
}
// User The username for the server. Can be anonymous.
func (r *APIUpdateLogFtpRequest) User(user string) *APIUpdateLogFtpRequest {
	r.user = &user
	return r
}
// Port The port number.
func (r *APIUpdateLogFtpRequest) Port(port int32) *APIUpdateLogFtpRequest {
	r.port = &port
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogFtpRequest) Execute() (*LoggingFtpResponse, *http.Response, error) {
	return r.APIService.UpdateLogFtpExecute(r)
}

/*
UpdateLogFtp Update an FTP log endpoint

Update the FTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingFtpName The name for the real-time logging configuration.
 @return APIUpdateLogFtpRequest
*/
func (a *LoggingFtpAPIService) UpdateLogFtp(ctx context.Context, serviceID string, versionID int32, loggingFtpName string) APIUpdateLogFtpRequest {
	return APIUpdateLogFtpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingFtpName: loggingFtpName,
	}
}

// UpdateLogFtpExecute executes the request
//  @return LoggingFtpResponse
func (a *LoggingFtpAPIService) UpdateLogFtpExecute(r APIUpdateLogFtpRequest) (*LoggingFtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingFtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingFtpAPIService.UpdateLogFtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/ftp/{logging_ftp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_ftp_name"+"}", gourl.PathEscape(parameterToString(r.loggingFtpName, "")))

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
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.hostname != nil {
		localVarFormParams.Add("hostname", parameterToString(*r.hostname, ""))
	}
	if r.ipv4 != nil {
		localVarFormParams.Add("ipv4", parameterToString(*r.ipv4, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
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
