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

// LoggingSftpAPI defines an interface for interacting with the resource.
type LoggingSftpAPI interface {

	/*
	CreateLogSftp Create an SFTP log endpoint

	Create a SFTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogSftpRequest
	*/
	CreateLogSftp(ctx context.Context, serviceID string, versionID int32) APICreateLogSftpRequest

	// CreateLogSftpExecute executes the request
	//  @return LoggingSftpResponse
	CreateLogSftpExecute(r APICreateLogSftpRequest) (*LoggingSftpResponse, *http.Response, error)

	/*
	DeleteLogSftp Delete an SFTP log endpoint

	Delete the SFTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSftpName The name for the real-time logging configuration.
	 @return APIDeleteLogSftpRequest
	*/
	DeleteLogSftp(ctx context.Context, serviceID string, versionID int32, loggingSftpName string) APIDeleteLogSftpRequest

	// DeleteLogSftpExecute executes the request
	//  @return InlineResponse200
	DeleteLogSftpExecute(r APIDeleteLogSftpRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogSftp Get an SFTP log endpoint

	Get the SFTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSftpName The name for the real-time logging configuration.
	 @return APIGetLogSftpRequest
	*/
	GetLogSftp(ctx context.Context, serviceID string, versionID int32, loggingSftpName string) APIGetLogSftpRequest

	// GetLogSftpExecute executes the request
	//  @return LoggingSftpResponse
	GetLogSftpExecute(r APIGetLogSftpRequest) (*LoggingSftpResponse, *http.Response, error)

	/*
	ListLogSftp List SFTP log endpoints

	List all of the SFTPs for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogSftpRequest
	*/
	ListLogSftp(ctx context.Context, serviceID string, versionID int32) APIListLogSftpRequest

	// ListLogSftpExecute executes the request
	//  @return []LoggingSftpResponse
	ListLogSftpExecute(r APIListLogSftpRequest) ([]LoggingSftpResponse, *http.Response, error)

	/*
	UpdateLogSftp Update an SFTP log endpoint

	Update the SFTP for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingSftpName The name for the real-time logging configuration.
	 @return APIUpdateLogSftpRequest
	*/
	UpdateLogSftp(ctx context.Context, serviceID string, versionID int32, loggingSftpName string) APIUpdateLogSftpRequest

	// UpdateLogSftpExecute executes the request
	//  @return LoggingSftpResponse
	UpdateLogSftpExecute(r APIUpdateLogSftpRequest) (*LoggingSftpResponse, *http.Response, error)
}

// LoggingSftpAPIService LoggingSftpAPI service
type LoggingSftpAPIService service

// APICreateLogSftpRequest represents a request for the resource.
type APICreateLogSftpRequest struct {
	ctx context.Context
	APIService LoggingSftpAPI
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
	port *int32
	password *string
	path *string
	publicKey *string
	secretKey *string
	sshKnownHosts *string
	user *string
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogSftpRequest) Name(name string) *APICreateLogSftpRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogSftpRequest) Placement(placement string) *APICreateLogSftpRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogSftpRequest) ResponseCondition(responseCondition string) *APICreateLogSftpRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogSftpRequest) Format(format string) *APICreateLogSftpRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogSftpRequest) FormatVersion(formatVersion int32) *APICreateLogSftpRequest {
	r.formatVersion = &formatVersion
	return r
}
// MessageType How the message should be formatted.
func (r *APICreateLogSftpRequest) MessageType(messageType string) *APICreateLogSftpRequest {
	r.messageType = &messageType
	return r
}
// TimestampFormat A timestamp format
func (r *APICreateLogSftpRequest) TimestampFormat(timestampFormat string) *APICreateLogSftpRequest {
	r.timestampFormat = &timestampFormat
	return r
}
// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogSftpRequest) CompressionCodec(compressionCodec string) *APICreateLogSftpRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogSftpRequest) Period(period int32) *APICreateLogSftpRequest {
	r.period = &period
	return r
}
// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogSftpRequest) GzipLevel(gzipLevel int32) *APICreateLogSftpRequest {
	r.gzipLevel = &gzipLevel
	return r
}
// Address A hostname or IPv4 address.
func (r *APICreateLogSftpRequest) Address(address string) *APICreateLogSftpRequest {
	r.address = &address
	return r
}
// Port The port number.
func (r *APICreateLogSftpRequest) Port(port int32) *APICreateLogSftpRequest {
	r.port = &port
	return r
}
// Password The password for the server. If both &#x60;password&#x60; and &#x60;secret_key&#x60; are passed, &#x60;secret_key&#x60; will be used in preference.
func (r *APICreateLogSftpRequest) Password(password string) *APICreateLogSftpRequest {
	r.password = &password
	return r
}
// Path The path to upload logs to.
func (r *APICreateLogSftpRequest) Path(path string) *APICreateLogSftpRequest {
	r.path = &path
	return r
}
// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogSftpRequest) PublicKey(publicKey string) *APICreateLogSftpRequest {
	r.publicKey = &publicKey
	return r
}
// SecretKey The SSH private key for the server. If both &#x60;password&#x60; and &#x60;secret_key&#x60; are passed, &#x60;secret_key&#x60; will be used in preference.
func (r *APICreateLogSftpRequest) SecretKey(secretKey string) *APICreateLogSftpRequest {
	r.secretKey = &secretKey
	return r
}
// SSHKnownHosts A list of host keys for all hosts we can connect to over SFTP.
func (r *APICreateLogSftpRequest) SSHKnownHosts(sshKnownHosts string) *APICreateLogSftpRequest {
	r.sshKnownHosts = &sshKnownHosts
	return r
}
// User The username for the server.
func (r *APICreateLogSftpRequest) User(user string) *APICreateLogSftpRequest {
	r.user = &user
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogSftpRequest) Execute() (*LoggingSftpResponse, *http.Response, error) {
	return r.APIService.CreateLogSftpExecute(r)
}

/*
CreateLogSftp Create an SFTP log endpoint

Create a SFTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogSftpRequest
*/
func (a *LoggingSftpAPIService) CreateLogSftp(ctx context.Context, serviceID string, versionID int32) APICreateLogSftpRequest {
	return APICreateLogSftpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogSftpExecute executes the request
//  @return LoggingSftpResponse
func (a *LoggingSftpAPIService) CreateLogSftpExecute(r APICreateLogSftpRequest) (*LoggingSftpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSftpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSftpAPIService.CreateLogSftp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/sftp"
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
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
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
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.sshKnownHosts != nil {
		localVarFormParams.Add("ssh_known_hosts", parameterToString(*r.sshKnownHosts, ""))
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

// APIDeleteLogSftpRequest represents a request for the resource.
type APIDeleteLogSftpRequest struct {
	ctx context.Context
	APIService LoggingSftpAPI
	serviceID string
	versionID int32
	loggingSftpName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogSftpRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogSftpExecute(r)
}

/*
DeleteLogSftp Delete an SFTP log endpoint

Delete the SFTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSftpName The name for the real-time logging configuration.
 @return APIDeleteLogSftpRequest
*/
func (a *LoggingSftpAPIService) DeleteLogSftp(ctx context.Context, serviceID string, versionID int32, loggingSftpName string) APIDeleteLogSftpRequest {
	return APIDeleteLogSftpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSftpName: loggingSftpName,
	}
}

// DeleteLogSftpExecute executes the request
//  @return InlineResponse200
func (a *LoggingSftpAPIService) DeleteLogSftpExecute(r APIDeleteLogSftpRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSftpAPIService.DeleteLogSftp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_sftp_name"+"}", gourl.PathEscape(parameterToString(r.loggingSftpName, "")))

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

// APIGetLogSftpRequest represents a request for the resource.
type APIGetLogSftpRequest struct {
	ctx context.Context
	APIService LoggingSftpAPI
	serviceID string
	versionID int32
	loggingSftpName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogSftpRequest) Execute() (*LoggingSftpResponse, *http.Response, error) {
	return r.APIService.GetLogSftpExecute(r)
}

/*
GetLogSftp Get an SFTP log endpoint

Get the SFTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSftpName The name for the real-time logging configuration.
 @return APIGetLogSftpRequest
*/
func (a *LoggingSftpAPIService) GetLogSftp(ctx context.Context, serviceID string, versionID int32, loggingSftpName string) APIGetLogSftpRequest {
	return APIGetLogSftpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSftpName: loggingSftpName,
	}
}

// GetLogSftpExecute executes the request
//  @return LoggingSftpResponse
func (a *LoggingSftpAPIService) GetLogSftpExecute(r APIGetLogSftpRequest) (*LoggingSftpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSftpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSftpAPIService.GetLogSftp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_sftp_name"+"}", gourl.PathEscape(parameterToString(r.loggingSftpName, "")))

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

// APIListLogSftpRequest represents a request for the resource.
type APIListLogSftpRequest struct {
	ctx context.Context
	APIService LoggingSftpAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogSftpRequest) Execute() ([]LoggingSftpResponse, *http.Response, error) {
	return r.APIService.ListLogSftpExecute(r)
}

/*
ListLogSftp List SFTP log endpoints

List all of the SFTPs for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogSftpRequest
*/
func (a *LoggingSftpAPIService) ListLogSftp(ctx context.Context, serviceID string, versionID int32) APIListLogSftpRequest {
	return APIListLogSftpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogSftpExecute executes the request
//  @return []LoggingSftpResponse
func (a *LoggingSftpAPIService) ListLogSftpExecute(r APIListLogSftpRequest) ([]LoggingSftpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingSftpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSftpAPIService.ListLogSftp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/sftp"
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

// APIUpdateLogSftpRequest represents a request for the resource.
type APIUpdateLogSftpRequest struct {
	ctx context.Context
	APIService LoggingSftpAPI
	serviceID string
	versionID int32
	loggingSftpName string
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
	port *int32
	password *string
	path *string
	publicKey *string
	secretKey *string
	sshKnownHosts *string
	user *string
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogSftpRequest) Name(name string) *APIUpdateLogSftpRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APIUpdateLogSftpRequest) Placement(placement string) *APIUpdateLogSftpRequest {
	r.placement = &placement
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogSftpRequest) ResponseCondition(responseCondition string) *APIUpdateLogSftpRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogSftpRequest) Format(format string) *APIUpdateLogSftpRequest {
	r.format = &format
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APIUpdateLogSftpRequest) FormatVersion(formatVersion int32) *APIUpdateLogSftpRequest {
	r.formatVersion = &formatVersion
	return r
}
// MessageType How the message should be formatted.
func (r *APIUpdateLogSftpRequest) MessageType(messageType string) *APIUpdateLogSftpRequest {
	r.messageType = &messageType
	return r
}
// TimestampFormat A timestamp format
func (r *APIUpdateLogSftpRequest) TimestampFormat(timestampFormat string) *APIUpdateLogSftpRequest {
	r.timestampFormat = &timestampFormat
	return r
}
// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogSftpRequest) CompressionCodec(compressionCodec string) *APIUpdateLogSftpRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogSftpRequest) Period(period int32) *APIUpdateLogSftpRequest {
	r.period = &period
	return r
}
// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogSftpRequest) GzipLevel(gzipLevel int32) *APIUpdateLogSftpRequest {
	r.gzipLevel = &gzipLevel
	return r
}
// Address A hostname or IPv4 address.
func (r *APIUpdateLogSftpRequest) Address(address string) *APIUpdateLogSftpRequest {
	r.address = &address
	return r
}
// Port The port number.
func (r *APIUpdateLogSftpRequest) Port(port int32) *APIUpdateLogSftpRequest {
	r.port = &port
	return r
}
// Password The password for the server. If both &#x60;password&#x60; and &#x60;secret_key&#x60; are passed, &#x60;secret_key&#x60; will be used in preference.
func (r *APIUpdateLogSftpRequest) Password(password string) *APIUpdateLogSftpRequest {
	r.password = &password
	return r
}
// Path The path to upload logs to.
func (r *APIUpdateLogSftpRequest) Path(path string) *APIUpdateLogSftpRequest {
	r.path = &path
	return r
}
// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogSftpRequest) PublicKey(publicKey string) *APIUpdateLogSftpRequest {
	r.publicKey = &publicKey
	return r
}
// SecretKey The SSH private key for the server. If both &#x60;password&#x60; and &#x60;secret_key&#x60; are passed, &#x60;secret_key&#x60; will be used in preference.
func (r *APIUpdateLogSftpRequest) SecretKey(secretKey string) *APIUpdateLogSftpRequest {
	r.secretKey = &secretKey
	return r
}
// SSHKnownHosts A list of host keys for all hosts we can connect to over SFTP.
func (r *APIUpdateLogSftpRequest) SSHKnownHosts(sshKnownHosts string) *APIUpdateLogSftpRequest {
	r.sshKnownHosts = &sshKnownHosts
	return r
}
// User The username for the server.
func (r *APIUpdateLogSftpRequest) User(user string) *APIUpdateLogSftpRequest {
	r.user = &user
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogSftpRequest) Execute() (*LoggingSftpResponse, *http.Response, error) {
	return r.APIService.UpdateLogSftpExecute(r)
}

/*
UpdateLogSftp Update an SFTP log endpoint

Update the SFTP for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingSftpName The name for the real-time logging configuration.
 @return APIUpdateLogSftpRequest
*/
func (a *LoggingSftpAPIService) UpdateLogSftp(ctx context.Context, serviceID string, versionID int32, loggingSftpName string) APIUpdateLogSftpRequest {
	return APIUpdateLogSftpRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingSftpName: loggingSftpName,
	}
}

// UpdateLogSftpExecute executes the request
//  @return LoggingSftpResponse
func (a *LoggingSftpAPIService) UpdateLogSftpExecute(r APIUpdateLogSftpRequest) (*LoggingSftpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingSftpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingSftpAPIService.UpdateLogSftp")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_sftp_name"+"}", gourl.PathEscape(parameterToString(r.loggingSftpName, "")))

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
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
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
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.sshKnownHosts != nil {
		localVarFormParams.Add("ssh_known_hosts", parameterToString(*r.sshKnownHosts, ""))
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
