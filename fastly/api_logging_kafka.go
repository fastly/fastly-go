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

// LoggingKafkaAPI defines an interface for interacting with the resource.
type LoggingKafkaAPI interface {

	/*
	CreateLogKafka Create a Kafka log endpoint

	Create a Kafka logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APICreateLogKafkaRequest
	*/
	CreateLogKafka(ctx context.Context, serviceID string, versionID int32) APICreateLogKafkaRequest

	// CreateLogKafkaExecute executes the request
	//  @return LoggingKafkaResponse
	CreateLogKafkaExecute(r APICreateLogKafkaRequest) (*LoggingKafkaResponse, *http.Response, error)

	/*
	DeleteLogKafka Delete the Kafka log endpoint

	Delete the Kafka logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingKafkaName The name for the real-time logging configuration.
	 @return APIDeleteLogKafkaRequest
	*/
	DeleteLogKafka(ctx context.Context, serviceID string, versionID int32, loggingKafkaName string) APIDeleteLogKafkaRequest

	// DeleteLogKafkaExecute executes the request
	//  @return InlineResponse200
	DeleteLogKafkaExecute(r APIDeleteLogKafkaRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetLogKafka Get a Kafka log endpoint

	Get the Kafka logging endpoint for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @param loggingKafkaName The name for the real-time logging configuration.
	 @return APIGetLogKafkaRequest
	*/
	GetLogKafka(ctx context.Context, serviceID string, versionID int32, loggingKafkaName string) APIGetLogKafkaRequest

	// GetLogKafkaExecute executes the request
	//  @return LoggingKafkaResponse
	GetLogKafkaExecute(r APIGetLogKafkaRequest) (*LoggingKafkaResponse, *http.Response, error)

	/*
	ListLogKafka List Kafka log endpoints

	List all of the Kafka logging endpoints for a particular service and version.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param versionID Integer identifying a service version.
	 @return APIListLogKafkaRequest
	*/
	ListLogKafka(ctx context.Context, serviceID string, versionID int32) APIListLogKafkaRequest

	// ListLogKafkaExecute executes the request
	//  @return []LoggingKafkaResponse
	ListLogKafkaExecute(r APIListLogKafkaRequest) ([]LoggingKafkaResponse, *http.Response, error)
}

// LoggingKafkaAPIService LoggingKafkaAPI service
type LoggingKafkaAPIService service

// APICreateLogKafkaRequest represents a request for the resource.
type APICreateLogKafkaRequest struct {
	ctx context.Context
	APIService LoggingKafkaAPI
	serviceID string
	versionID int32
	name *string
	placement *string
	formatVersion *int32
	responseCondition *string
	format *string
	tlsCaCert *string
	tlsClientCert *string
	tlsClientKey *string
	tlsHostname *string
	topic *string
	brokers *string
	compressionCodec *string
	requiredAcks *int32
	requestMaxBytes *int32
	parseLogKeyvals *bool
	authMethod *string
	user *string
	password *string
	useTLS *LoggingUseTLS
}

// Name The name for the real-time logging configuration.
func (r *APICreateLogKafkaRequest) Name(name string) *APICreateLogKafkaRequest {
	r.name = &name
	return r
}
// Placement Where in the generated VCL the logging call should be placed. If not set, endpoints with &#x60;format_version&#x60; of 2 are placed in &#x60;vcl_log&#x60; and those with &#x60;format_version&#x60; of 1 are placed in &#x60;vcl_deliver&#x60;. 
func (r *APICreateLogKafkaRequest) Placement(placement string) *APICreateLogKafkaRequest {
	r.placement = &placement
	return r
}
// FormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in &#x60;vcl_log&#x60; if &#x60;format_version&#x60; is set to &#x60;2&#x60; and in &#x60;vcl_deliver&#x60; if &#x60;format_version&#x60; is set to &#x60;1&#x60;. 
func (r *APICreateLogKafkaRequest) FormatVersion(formatVersion int32) *APICreateLogKafkaRequest {
	r.formatVersion = &formatVersion
	return r
}
// ResponseCondition The name of an existing condition in the configured endpoint, or leave blank to always execute.
func (r *APICreateLogKafkaRequest) ResponseCondition(responseCondition string) *APICreateLogKafkaRequest {
	r.responseCondition = &responseCondition
	return r
}
// Format A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
func (r *APICreateLogKafkaRequest) Format(format string) *APICreateLogKafkaRequest {
	r.format = &format
	return r
}
// TLSCaCert A secure certificate to authenticate a server with. Must be in PEM format.
func (r *APICreateLogKafkaRequest) TLSCaCert(tlsCaCert string) *APICreateLogKafkaRequest {
	r.tlsCaCert = &tlsCaCert
	return r
}
// TLSClientCert The client certificate used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogKafkaRequest) TLSClientCert(tlsClientCert string) *APICreateLogKafkaRequest {
	r.tlsClientCert = &tlsClientCert
	return r
}
// TLSClientKey The client private key used to make authenticated requests. Must be in PEM format.
func (r *APICreateLogKafkaRequest) TLSClientKey(tlsClientKey string) *APICreateLogKafkaRequest {
	r.tlsClientKey = &tlsClientKey
	return r
}
// TLSHostname The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
func (r *APICreateLogKafkaRequest) TLSHostname(tlsHostname string) *APICreateLogKafkaRequest {
	r.tlsHostname = &tlsHostname
	return r
}
// Topic The Kafka topic to send logs to. Required.
func (r *APICreateLogKafkaRequest) Topic(topic string) *APICreateLogKafkaRequest {
	r.topic = &topic
	return r
}
// Brokers A comma-separated list of IP addresses or hostnames of Kafka brokers. Required.
func (r *APICreateLogKafkaRequest) Brokers(brokers string) *APICreateLogKafkaRequest {
	r.brokers = &brokers
	return r
}
// CompressionCodec The codec used for compression of your logs.
func (r *APICreateLogKafkaRequest) CompressionCodec(compressionCodec string) *APICreateLogKafkaRequest {
	r.compressionCodec = &compressionCodec
	return r
}
// RequiredAcks The number of acknowledgements a leader must receive before a write is considered successful.
func (r *APICreateLogKafkaRequest) RequiredAcks(requiredAcks int32) *APICreateLogKafkaRequest {
	r.requiredAcks = &requiredAcks
	return r
}
// RequestMaxBytes The maximum number of bytes sent in one request. Defaults &#x60;0&#x60; (no limit).
func (r *APICreateLogKafkaRequest) RequestMaxBytes(requestMaxBytes int32) *APICreateLogKafkaRequest {
	r.requestMaxBytes = &requestMaxBytes
	return r
}
// ParseLogKeyvals Enables parsing of key&#x3D;value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers).
func (r *APICreateLogKafkaRequest) ParseLogKeyvals(parseLogKeyvals bool) *APICreateLogKafkaRequest {
	r.parseLogKeyvals = &parseLogKeyvals
	return r
}
// AuthMethod SASL authentication method.
func (r *APICreateLogKafkaRequest) AuthMethod(authMethod string) *APICreateLogKafkaRequest {
	r.authMethod = &authMethod
	return r
}
// User SASL user.
func (r *APICreateLogKafkaRequest) User(user string) *APICreateLogKafkaRequest {
	r.user = &user
	return r
}
// Password SASL password.
func (r *APICreateLogKafkaRequest) Password(password string) *APICreateLogKafkaRequest {
	r.password = &password
	return r
}
// UseTLS returns a pointer to a request.
func (r *APICreateLogKafkaRequest) UseTLS(useTLS LoggingUseTLS) *APICreateLogKafkaRequest {
	r.useTLS = &useTLS
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateLogKafkaRequest) Execute() (*LoggingKafkaResponse, *http.Response, error) {
	return r.APIService.CreateLogKafkaExecute(r)
}

/*
CreateLogKafka Create a Kafka log endpoint

Create a Kafka logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APICreateLogKafkaRequest
*/
func (a *LoggingKafkaAPIService) CreateLogKafka(ctx context.Context, serviceID string, versionID int32) APICreateLogKafkaRequest {
	return APICreateLogKafkaRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// CreateLogKafkaExecute executes the request
//  @return LoggingKafkaResponse
func (a *LoggingKafkaAPIService) CreateLogKafkaExecute(r APICreateLogKafkaRequest) (*LoggingKafkaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingKafkaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKafkaAPIService.CreateLogKafka")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kafka"
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
	if r.formatVersion != nil {
		localVarFormParams.Add("format_version", parameterToString(*r.formatVersion, ""))
	}
	if r.responseCondition != nil {
		localVarFormParams.Add("response_condition", parameterToString(*r.responseCondition, ""))
	}
	if r.format != nil {
		localVarFormParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.tlsCaCert != nil {
		localVarFormParams.Add("tls_ca_cert", parameterToString(*r.tlsCaCert, ""))
	}
	if r.tlsClientCert != nil {
		localVarFormParams.Add("tls_client_cert", parameterToString(*r.tlsClientCert, ""))
	}
	if r.tlsClientKey != nil {
		localVarFormParams.Add("tls_client_key", parameterToString(*r.tlsClientKey, ""))
	}
	if r.tlsHostname != nil {
		localVarFormParams.Add("tls_hostname", parameterToString(*r.tlsHostname, ""))
	}
	if r.topic != nil {
		localVarFormParams.Add("topic", parameterToString(*r.topic, ""))
	}
	if r.brokers != nil {
		localVarFormParams.Add("brokers", parameterToString(*r.brokers, ""))
	}
	if r.compressionCodec != nil {
		localVarFormParams.Add("compression_codec", parameterToString(*r.compressionCodec, ""))
	}
	if r.requiredAcks != nil {
		localVarFormParams.Add("required_acks", parameterToString(*r.requiredAcks, ""))
	}
	if r.requestMaxBytes != nil {
		localVarFormParams.Add("request_max_bytes", parameterToString(*r.requestMaxBytes, ""))
	}
	if r.parseLogKeyvals != nil {
		localVarFormParams.Add("parse_log_keyvals", parameterToString(*r.parseLogKeyvals, ""))
	}
	if r.authMethod != nil {
		localVarFormParams.Add("auth_method", parameterToString(*r.authMethod, ""))
	}
	if r.user != nil {
		localVarFormParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
	}
	if r.useTLS != nil {
		localVarFormParams.Add("use_tls", parameterToString(*r.useTLS, ""))
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

// APIDeleteLogKafkaRequest represents a request for the resource.
type APIDeleteLogKafkaRequest struct {
	ctx context.Context
	APIService LoggingKafkaAPI
	serviceID string
	versionID int32
	loggingKafkaName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteLogKafkaRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteLogKafkaExecute(r)
}

/*
DeleteLogKafka Delete the Kafka log endpoint

Delete the Kafka logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingKafkaName The name for the real-time logging configuration.
 @return APIDeleteLogKafkaRequest
*/
func (a *LoggingKafkaAPIService) DeleteLogKafka(ctx context.Context, serviceID string, versionID int32, loggingKafkaName string) APIDeleteLogKafkaRequest {
	return APIDeleteLogKafkaRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingKafkaName: loggingKafkaName,
	}
}

// DeleteLogKafkaExecute executes the request
//  @return InlineResponse200
func (a *LoggingKafkaAPIService) DeleteLogKafkaExecute(r APIDeleteLogKafkaRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKafkaAPIService.DeleteLogKafka")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_kafka_name"+"}", gourl.PathEscape(parameterToString(r.loggingKafkaName, "")))

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

// APIGetLogKafkaRequest represents a request for the resource.
type APIGetLogKafkaRequest struct {
	ctx context.Context
	APIService LoggingKafkaAPI
	serviceID string
	versionID int32
	loggingKafkaName string
}


// Execute calls the API using the request data configured.
func (r APIGetLogKafkaRequest) Execute() (*LoggingKafkaResponse, *http.Response, error) {
	return r.APIService.GetLogKafkaExecute(r)
}

/*
GetLogKafka Get a Kafka log endpoint

Get the Kafka logging endpoint for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @param loggingKafkaName The name for the real-time logging configuration.
 @return APIGetLogKafkaRequest
*/
func (a *LoggingKafkaAPIService) GetLogKafka(ctx context.Context, serviceID string, versionID int32, loggingKafkaName string) APIGetLogKafkaRequest {
	return APIGetLogKafkaRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
		loggingKafkaName: loggingKafkaName,
	}
}

// GetLogKafkaExecute executes the request
//  @return LoggingKafkaResponse
func (a *LoggingKafkaAPIService) GetLogKafkaExecute(r APIGetLogKafkaRequest) (*LoggingKafkaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *LoggingKafkaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKafkaAPIService.GetLogKafka")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"logging_kafka_name"+"}", gourl.PathEscape(parameterToString(r.loggingKafkaName, "")))

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

// APIListLogKafkaRequest represents a request for the resource.
type APIListLogKafkaRequest struct {
	ctx context.Context
	APIService LoggingKafkaAPI
	serviceID string
	versionID int32
}


// Execute calls the API using the request data configured.
func (r APIListLogKafkaRequest) Execute() ([]LoggingKafkaResponse, *http.Response, error) {
	return r.APIService.ListLogKafkaExecute(r)
}

/*
ListLogKafka List Kafka log endpoints

List all of the Kafka logging endpoints for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param versionID Integer identifying a service version.
 @return APIListLogKafkaRequest
*/
func (a *LoggingKafkaAPIService) ListLogKafka(ctx context.Context, serviceID string, versionID int32) APIListLogKafkaRequest {
	return APIListLogKafkaRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		versionID: versionID,
	}
}

// ListLogKafkaExecute executes the request
//  @return []LoggingKafkaResponse
func (a *LoggingKafkaAPIService) ListLogKafkaExecute(r APIListLogKafkaRequest) ([]LoggingKafkaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []LoggingKafkaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingKafkaAPIService.ListLogKafka")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/logging/kafka"
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
