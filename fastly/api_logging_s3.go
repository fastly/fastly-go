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

// LoggingS3API defines an interface for interacting with the resource.
type LoggingS3API interface {

	/*
		CreateLogAwsS3 Create an AWS S3 log endpoint

		Create a S3 for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APICreateLogAwsS3Request
	*/
	CreateLogAwsS3(ctx context.Context, serviceID string, versionID int32) APICreateLogAwsS3Request

	// CreateLogAwsS3Execute executes the request
	//  @return LoggingS3Response
	CreateLogAwsS3Execute(r APICreateLogAwsS3Request) (*LoggingS3Response, *http.Response, error)

	/*
		DeleteLogAwsS3 Delete an AWS S3 log endpoint

		Delete the S3 for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingS3Name The name for the real-time logging configuration.
		 @return APIDeleteLogAwsS3Request
	*/
	DeleteLogAwsS3(ctx context.Context, serviceID string, versionID int32, loggingS3Name string) APIDeleteLogAwsS3Request

	// DeleteLogAwsS3Execute executes the request
	//  @return InlineResponse200
	DeleteLogAwsS3Execute(r APIDeleteLogAwsS3Request) (*InlineResponse200, *http.Response, error)

	/*
		GetLogAwsS3 Get an AWS S3 log endpoint

		Get the S3 for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingS3Name The name for the real-time logging configuration.
		 @return APIGetLogAwsS3Request
	*/
	GetLogAwsS3(ctx context.Context, serviceID string, versionID int32, loggingS3Name string) APIGetLogAwsS3Request

	// GetLogAwsS3Execute executes the request
	//  @return LoggingS3Response
	GetLogAwsS3Execute(r APIGetLogAwsS3Request) (*LoggingS3Response, *http.Response, error)

	/*
		ListLogAwsS3 List AWS S3 log endpoints

		List all of the S3s for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @return APIListLogAwsS3Request
	*/
	ListLogAwsS3(ctx context.Context, serviceID string, versionID int32) APIListLogAwsS3Request

	// ListLogAwsS3Execute executes the request
	//  @return []LoggingS3Response
	ListLogAwsS3Execute(r APIListLogAwsS3Request) ([]LoggingS3Response, *http.Response, error)

	/*
		UpdateLogAwsS3 Update an AWS S3 log endpoint

		Update the S3 for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param versionID Integer identifying a service version.
		 @param loggingS3Name The name for the real-time logging configuration.
		 @return APIUpdateLogAwsS3Request
	*/
	UpdateLogAwsS3(ctx context.Context, serviceID string, versionID int32, loggingS3Name string) APIUpdateLogAwsS3Request

	// UpdateLogAwsS3Execute executes the request
	//  @return LoggingS3Response
	UpdateLogAwsS3Execute(r APIUpdateLogAwsS3Request) (*LoggingS3Response, *http.Response, error)
}

// LoggingS3APIService LoggingS3API service
type LoggingS3APIService service

// APICreateLogAwsS3Request represents a request for the resource.
type APICreateLogAwsS3Request struct {
	ctx                          context.Context
	APIService                   LoggingS3API
	serviceID                    string
	versionID                    int32
	name                         *string
	placement                    *string
	responseCondition            *string
	format                       *string
	formatVersion                *int32
	messageType                  *string
	timestampFormat              *string
	compressionCodec             *string
	period                       *int32
	gzipLevel                    *int32
	accessKey                    *string
	acl                          *string
	bucketName                   *string
	domain                       *string
	iamRole                      *string
	path                         *string
	publicKey                    *string
	redundancy                   *string
	secretKey                    *string
	serverSideEncryptionKmsKeyID *string
	serverSideEncryption         *string
	fileMaxBytes                 *int32
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogAwsS3Request) Name(name string) *APICreateLogAwsS3Request {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APICreateLogAwsS3Request) Placement(placement string) *APICreateLogAwsS3Request {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogAwsS3Request) ResponseCondition(responseCondition string) *APICreateLogAwsS3Request {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogAwsS3Request) Format(format string) *APICreateLogAwsS3Request {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APICreateLogAwsS3Request) FormatVersion(formatVersion int32) *APICreateLogAwsS3Request {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APICreateLogAwsS3Request) MessageType(messageType string) *APICreateLogAwsS3Request {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APICreateLogAwsS3Request) TimestampFormat(timestampFormat string) *APICreateLogAwsS3Request {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogAwsS3Request) CompressionCodec(compressionCodec string) *APICreateLogAwsS3Request {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APICreateLogAwsS3Request) Period(period int32) *APICreateLogAwsS3Request {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APICreateLogAwsS3Request) GzipLevel(gzipLevel int32) *APICreateLogAwsS3Request {
	r.gzipLevel = &gzipLevel
	return r
}

// AccessKey The access key for your S3 account. Not required if &#x60;iam_role&#x60; is provided.
func (r *APICreateLogAwsS3Request) AccessKey(accessKey string) *APICreateLogAwsS3Request {
	r.accessKey = &accessKey
	return r
}

// ACL The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information.
func (r *APICreateLogAwsS3Request) ACL(acl string) *APICreateLogAwsS3Request {
	r.acl = &acl
	return r
}

// BucketName The bucket name for S3 account.
func (r *APICreateLogAwsS3Request) BucketName(bucketName string) *APICreateLogAwsS3Request {
	r.bucketName = &bucketName
	return r
}

// Domain The domain of the Amazon S3 endpoint.
func (r *APICreateLogAwsS3Request) Domain(domain string) *APICreateLogAwsS3Request {
	r.domain = &domain
	return r
}

// IamRole The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if &#x60;access_key&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APICreateLogAwsS3Request) IamRole(iamRole string) *APICreateLogAwsS3Request {
	r.iamRole = &iamRole
	return r
}

// Path The path to upload logs to.
func (r *APICreateLogAwsS3Request) Path(path string) *APICreateLogAwsS3Request {
	r.path = &path
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APICreateLogAwsS3Request) PublicKey(publicKey string) *APICreateLogAwsS3Request {
	r.publicKey = &publicKey
	return r
}

// Redundancy The S3 redundancy level.
func (r *APICreateLogAwsS3Request) Redundancy(redundancy string) *APICreateLogAwsS3Request {
	r.redundancy = &redundancy
	return r
}

// SecretKey The secret key for your S3 account. Not required if &#x60;iam_role&#x60; is provided.
func (r *APICreateLogAwsS3Request) SecretKey(secretKey string) *APICreateLogAwsS3Request {
	r.secretKey = &secretKey
	return r
}

// ServerSideEncryptionKmsKeyID Optional server-side KMS Key ID. Must be set if &#x60;server_side_encryption&#x60; is set to &#x60;aws:kms&#x60; or &#x60;AES256&#x60;.
func (r *APICreateLogAwsS3Request) ServerSideEncryptionKmsKeyID(serverSideEncryptionKmsKeyID string) *APICreateLogAwsS3Request {
	r.serverSideEncryptionKmsKeyID = &serverSideEncryptionKmsKeyID
	return r
}

// ServerSideEncryption Set this to &#x60;AES256&#x60; or &#x60;aws:kms&#x60; to enable S3 Server Side Encryption.
func (r *APICreateLogAwsS3Request) ServerSideEncryption(serverSideEncryption string) *APICreateLogAwsS3Request {
	r.serverSideEncryption = &serverSideEncryption
	return r
}

// FileMaxBytes The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
func (r *APICreateLogAwsS3Request) FileMaxBytes(fileMaxBytes int32) *APICreateLogAwsS3Request {
	r.fileMaxBytes = &fileMaxBytes
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogAwsS3Request) Execute() (*LoggingS3Response, *http.Response, error) {
	return r.APIService.CreateLogAwsS3Execute(r)
}

/*
CreateLogAwsS3 Create an AWS S3 log endpoint

Create a S3 for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogAwsS3Request
*/
func (a *LoggingS3APIService) CreateLogAwsS3(ctx context.Context, serviceID string, versionID int32) APICreateLogAwsS3Request {
	return APICreateLogAwsS3Request{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// CreateLogAwsS3Execute executes the request
//  @return LoggingS3Response
func (a *LoggingS3APIService) CreateLogAwsS3Execute(r APICreateLogAwsS3Request) (*LoggingS3Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingS3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingS3APIService.CreateLogAwsS3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/s3"
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
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.acl != nil {
		localVarFormParams.Add("acl", parameterToString(*r.acl, ""))
	}
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.domain != nil {
		localVarFormParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.iamRole != nil {
		localVarFormParams.Add("iam_role", parameterToString(*r.iamRole, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.redundancy != nil {
		localVarFormParams.Add("redundancy", parameterToString(*r.redundancy, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.serverSideEncryptionKmsKeyID != nil {
		localVarFormParams.Add("server_side_encryption_kms_key_id", parameterToString(*r.serverSideEncryptionKmsKeyID, ""))
	}
	if r.serverSideEncryption != nil {
		localVarFormParams.Add("server_side_encryption", parameterToString(*r.serverSideEncryption, ""))
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

// APIDeleteLogAwsS3Request represents a request for the resource.
type APIDeleteLogAwsS3Request struct {
	ctx           context.Context
	APIService    LoggingS3API
	serviceID     string
	versionID     int32
	loggingS3Name string
}

// Execute calls the API using the request data configured.
func (r APIDeleteLogAwsS3Request) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogAwsS3Execute(r)
}

/*
DeleteLogAwsS3 Delete an AWS S3 log endpoint

Delete the S3 for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingS3Name The name for the real-time logging configuration.
 @return APIDeleteLogAwsS3Request
*/
func (a *LoggingS3APIService) DeleteLogAwsS3(ctx context.Context, serviceID string, versionID int32, loggingS3Name string) APIDeleteLogAwsS3Request {
	return APIDeleteLogAwsS3Request{
		APIService:    a,
		ctx:           ctx,
		serviceID:     serviceID,
		versionID:     versionID,
		loggingS3Name: loggingS3Name,
	}
}

// DeleteLogAwsS3Execute executes the request
//  @return InlineResponse200
func (a *LoggingS3APIService) DeleteLogAwsS3Execute(r APIDeleteLogAwsS3Request) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingS3APIService.DeleteLogAwsS3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_s3_name"+"}", gourl.PathEscape(parameterToString(r.loggingS3Name, "")))

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

// APIGetLogAwsS3Request represents a request for the resource.
type APIGetLogAwsS3Request struct {
	ctx           context.Context
	APIService    LoggingS3API
	serviceID     string
	versionID     int32
	loggingS3Name string
}

// Execute calls the API using the request data configured.
func (r APIGetLogAwsS3Request) Execute() (*LoggingS3Response, *http.Response, error) {
	return r.APIService.GetLogAwsS3Execute(r)
}

/*
GetLogAwsS3 Get an AWS S3 log endpoint

Get the S3 for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingS3Name The name for the real-time logging configuration.
 @return APIGetLogAwsS3Request
*/
func (a *LoggingS3APIService) GetLogAwsS3(ctx context.Context, serviceID string, versionID int32, loggingS3Name string) APIGetLogAwsS3Request {
	return APIGetLogAwsS3Request{
		APIService:    a,
		ctx:           ctx,
		serviceID:     serviceID,
		versionID:     versionID,
		loggingS3Name: loggingS3Name,
	}
}

// GetLogAwsS3Execute executes the request
//  @return LoggingS3Response
func (a *LoggingS3APIService) GetLogAwsS3Execute(r APIGetLogAwsS3Request) (*LoggingS3Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingS3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingS3APIService.GetLogAwsS3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_s3_name"+"}", gourl.PathEscape(parameterToString(r.loggingS3Name, "")))

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

// APIListLogAwsS3Request represents a request for the resource.
type APIListLogAwsS3Request struct {
	ctx        context.Context
	APIService LoggingS3API
	serviceID  string
	versionID  int32
}

// Execute calls the API using the request data configured.
func (r APIListLogAwsS3Request) Execute() ([]LoggingS3Response, *http.Response, error) {
	return r.APIService.ListLogAwsS3Execute(r)
}

/*
ListLogAwsS3 List AWS S3 log endpoints

List all of the S3s for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogAwsS3Request
*/
func (a *LoggingS3APIService) ListLogAwsS3(ctx context.Context, serviceID string, versionID int32) APIListLogAwsS3Request {
	return APIListLogAwsS3Request{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		versionID:  versionID,
	}
}

// ListLogAwsS3Execute executes the request
//  @return []LoggingS3Response
func (a *LoggingS3APIService) ListLogAwsS3Execute(r APIListLogAwsS3Request) ([]LoggingS3Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []LoggingS3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingS3APIService.ListLogAwsS3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/s3"
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

// APIUpdateLogAwsS3Request represents a request for the resource.
type APIUpdateLogAwsS3Request struct {
	ctx                          context.Context
	APIService                   LoggingS3API
	serviceID                    string
	versionID                    int32
	loggingS3Name                string
	name                         *string
	placement                    *string
	responseCondition            *string
	format                       *string
	formatVersion                *int32
	messageType                  *string
	timestampFormat              *string
	compressionCodec             *string
	period                       *int32
	gzipLevel                    *int32
	accessKey                    *string
	acl                          *string
	bucketName                   *string
	domain                       *string
	iamRole                      *string
	path                         *string
	publicKey                    *string
	redundancy                   *string
	secretKey                    *string
	serverSideEncryptionKmsKeyID *string
	serverSideEncryption         *string
	fileMaxBytes                 *int32
}

// Name The name for the real-time logging configuration.
func (r *APIUpdateLogAwsS3Request) Name(name string) *APIUpdateLogAwsS3Request {
	r.name = &name
	return r
}

// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;.
func (r *APIUpdateLogAwsS3Request) Placement(placement string) *APIUpdateLogAwsS3Request {
	r.placement = &placement
	return r
}

// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APIUpdateLogAwsS3Request) ResponseCondition(responseCondition string) *APIUpdateLogAwsS3Request {
	r.responseCondition = &responseCondition
	return r
}

// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APIUpdateLogAwsS3Request) Format(format string) *APIUpdateLogAwsS3Request {
	r.format = &format
	return r
}

// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;.
func (r *APIUpdateLogAwsS3Request) FormatVersion(formatVersion int32) *APIUpdateLogAwsS3Request {
	r.formatVersion = &formatVersion
	return r
}

// MessageType How the message should be formatted.
func (r *APIUpdateLogAwsS3Request) MessageType(messageType string) *APIUpdateLogAwsS3Request {
	r.messageType = &messageType
	return r
}

// TimestampFormat A timestamp format
func (r *APIUpdateLogAwsS3Request) TimestampFormat(timestampFormat string) *APIUpdateLogAwsS3Request {
	r.timestampFormat = &timestampFormat
	return r
}

// CompressionCodec The codec used for compressing your logs. Valid values are &#x60;zstd&#x60;, &#x60;snappy&#x60;, and &#x60;gzip&#x60;. Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogAwsS3Request) CompressionCodec(compressionCodec string) *APIUpdateLogAwsS3Request {
	r.compressionCodec = &compressionCodec
	return r
}

// Period How frequently log files are finalized so they can be available for reading (in seconds).
func (r *APIUpdateLogAwsS3Request) Period(period int32) *APIUpdateLogAwsS3Request {
	r.period = &period
	return r
}

// GzipLevel The level of gzip encoding when sending logs (default &#x60;0&#x60;, no compression). Specifying both &#x60;compression_codec&#x60; and &#x60;gzip_level&#x60; in the same API request will result in an error.
func (r *APIUpdateLogAwsS3Request) GzipLevel(gzipLevel int32) *APIUpdateLogAwsS3Request {
	r.gzipLevel = &gzipLevel
	return r
}

// AccessKey The access key for your S3 account. Not required if &#x60;iam_role&#x60; is provided.
func (r *APIUpdateLogAwsS3Request) AccessKey(accessKey string) *APIUpdateLogAwsS3Request {
	r.accessKey = &accessKey
	return r
}

// ACL The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information.
func (r *APIUpdateLogAwsS3Request) ACL(acl string) *APIUpdateLogAwsS3Request {
	r.acl = &acl
	return r
}

// BucketName The bucket name for S3 account.
func (r *APIUpdateLogAwsS3Request) BucketName(bucketName string) *APIUpdateLogAwsS3Request {
	r.bucketName = &bucketName
	return r
}

// Domain The domain of the Amazon S3 endpoint.
func (r *APIUpdateLogAwsS3Request) Domain(domain string) *APIUpdateLogAwsS3Request {
	r.domain = &domain
	return r
}

// IamRole The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if &#x60;access_key&#x60; and &#x60;secret_key&#x60; are provided.
func (r *APIUpdateLogAwsS3Request) IamRole(iamRole string) *APIUpdateLogAwsS3Request {
	r.iamRole = &iamRole
	return r
}

// Path The path to upload logs to.
func (r *APIUpdateLogAwsS3Request) Path(path string) *APIUpdateLogAwsS3Request {
	r.path = &path
	return r
}

// PublicKey A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
func (r *APIUpdateLogAwsS3Request) PublicKey(publicKey string) *APIUpdateLogAwsS3Request {
	r.publicKey = &publicKey
	return r
}

// Redundancy The S3 redundancy level.
func (r *APIUpdateLogAwsS3Request) Redundancy(redundancy string) *APIUpdateLogAwsS3Request {
	r.redundancy = &redundancy
	return r
}

// SecretKey The secret key for your S3 account. Not required if &#x60;iam_role&#x60; is provided.
func (r *APIUpdateLogAwsS3Request) SecretKey(secretKey string) *APIUpdateLogAwsS3Request {
	r.secretKey = &secretKey
	return r
}

// ServerSideEncryptionKmsKeyID Optional server-side KMS Key ID. Must be set if &#x60;server_side_encryption&#x60; is set to &#x60;aws:kms&#x60; or &#x60;AES256&#x60;.
func (r *APIUpdateLogAwsS3Request) ServerSideEncryptionKmsKeyID(serverSideEncryptionKmsKeyID string) *APIUpdateLogAwsS3Request {
	r.serverSideEncryptionKmsKeyID = &serverSideEncryptionKmsKeyID
	return r
}

// ServerSideEncryption Set this to &#x60;AES256&#x60; or &#x60;aws:kms&#x60; to enable S3 Server Side Encryption.
func (r *APIUpdateLogAwsS3Request) ServerSideEncryption(serverSideEncryption string) *APIUpdateLogAwsS3Request {
	r.serverSideEncryption = &serverSideEncryption
	return r
}

// FileMaxBytes The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
func (r *APIUpdateLogAwsS3Request) FileMaxBytes(fileMaxBytes int32) *APIUpdateLogAwsS3Request {
	r.fileMaxBytes = &fileMaxBytes
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateLogAwsS3Request) Execute() (*LoggingS3Response, *http.Response, error) {
	return r.APIService.UpdateLogAwsS3Execute(r)
}

/*
UpdateLogAwsS3 Update an AWS S3 log endpoint

Update the S3 for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingS3Name The name for the real-time logging configuration.
 @return APIUpdateLogAwsS3Request
*/
func (a *LoggingS3APIService) UpdateLogAwsS3(ctx context.Context, serviceID string, versionID int32, loggingS3Name string) APIUpdateLogAwsS3Request {
	return APIUpdateLogAwsS3Request{
		APIService:    a,
		ctx:           ctx,
		serviceID:     serviceID,
		versionID:     versionID,
		loggingS3Name: loggingS3Name,
	}
}

// UpdateLogAwsS3Execute executes the request
//  @return LoggingS3Response
func (a *LoggingS3APIService) UpdateLogAwsS3Execute(r APIUpdateLogAwsS3Request) (*LoggingS3Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LoggingS3Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingS3APIService.UpdateLogAwsS3")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_s3_name"+"}", gourl.PathEscape(parameterToString(r.loggingS3Name, "")))

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
	if r.accessKey != nil {
		localVarFormParams.Add("access_key", parameterToString(*r.accessKey, ""))
	}
	if r.acl != nil {
		localVarFormParams.Add("acl", parameterToString(*r.acl, ""))
	}
	if r.bucketName != nil {
		localVarFormParams.Add("bucket_name", parameterToString(*r.bucketName, ""))
	}
	if r.domain != nil {
		localVarFormParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.iamRole != nil {
		localVarFormParams.Add("iam_role", parameterToString(*r.iamRole, ""))
	}
	if r.path != nil {
		localVarFormParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.publicKey != nil {
		localVarFormParams.Add("public_key", parameterToString(*r.publicKey, ""))
	}
	if r.redundancy != nil {
		localVarFormParams.Add("redundancy", parameterToString(*r.redundancy, ""))
	}
	if r.secretKey != nil {
		localVarFormParams.Add("secret_key", parameterToString(*r.secretKey, ""))
	}
	if r.serverSideEncryptionKmsKeyID != nil {
		localVarFormParams.Add("server_side_encryption_kms_key_id", parameterToString(*r.serverSideEncryptionKmsKeyID, ""))
	}
	if r.serverSideEncryption != nil {
		localVarFormParams.Add("server_side_encryption", parameterToString(*r.serverSideEncryption, ""))
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
