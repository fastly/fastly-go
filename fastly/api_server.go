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

// ServerAPI defines an interface for interacting with the resource.
type ServerAPI interface {

	/*
		CreatePoolServer Add a server to a pool

		Creates a single server for a particular service and pool.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param poolID Alphanumeric string identifying a Pool.
		 @return APICreatePoolServerRequest
	*/
	CreatePoolServer(ctx context.Context, serviceID string, poolID string) APICreatePoolServerRequest

	// CreatePoolServerExecute executes the request
	//  @return ServerResponse
	CreatePoolServerExecute(r APICreatePoolServerRequest) (*ServerResponse, *http.Response, error)

	/*
		DeletePoolServer Delete a server from a pool

		Deletes a single server for a particular service and pool.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param poolID Alphanumeric string identifying a Pool.
		 @param serverID Alphanumeric string identifying a Server.
		 @return APIDeletePoolServerRequest
	*/
	DeletePoolServer(ctx context.Context, serviceID string, poolID string, serverID string) APIDeletePoolServerRequest

	// DeletePoolServerExecute executes the request
	//  @return InlineResponse200
	DeletePoolServerExecute(r APIDeletePoolServerRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetPoolServer Get a pool server

		Gets a single server for a particular service and pool.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param poolID Alphanumeric string identifying a Pool.
		 @param serverID Alphanumeric string identifying a Server.
		 @return APIGetPoolServerRequest
	*/
	GetPoolServer(ctx context.Context, serviceID string, poolID string, serverID string) APIGetPoolServerRequest

	// GetPoolServerExecute executes the request
	//  @return ServerResponse
	GetPoolServerExecute(r APIGetPoolServerRequest) (*ServerResponse, *http.Response, error)

	/*
		ListPoolServers List servers in a pool

		Lists all servers for a particular service and pool.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param poolID Alphanumeric string identifying a Pool.
		 @return APIListPoolServersRequest
	*/
	ListPoolServers(ctx context.Context, serviceID string, poolID string) APIListPoolServersRequest

	// ListPoolServersExecute executes the request
	//  @return []ServerResponse
	ListPoolServersExecute(r APIListPoolServersRequest) ([]ServerResponse, *http.Response, error)

	/*
		UpdatePoolServer Update a server

		Updates a single server for a particular service and pool.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param poolID Alphanumeric string identifying a Pool.
		 @param serverID Alphanumeric string identifying a Server.
		 @return APIUpdatePoolServerRequest
	*/
	UpdatePoolServer(ctx context.Context, serviceID string, poolID string, serverID string) APIUpdatePoolServerRequest

	// UpdatePoolServerExecute executes the request
	//  @return ServerResponse
	UpdatePoolServerExecute(r APIUpdatePoolServerRequest) (*ServerResponse, *http.Response, error)
}

// ServerAPIService ServerAPI service
type ServerAPIService service

// APICreatePoolServerRequest represents a request for the resource.
type APICreatePoolServerRequest struct {
	ctx          context.Context
	APIService   ServerAPI
	serviceID    string
	poolID       string
	weight       *int32
	maxConn      *int32
	port         *int32
	address      *string
	comment      *string
	disabled     *bool
	overrideHost *string
}

// Weight Weight (&#x60;1-100&#x60;) used to load balance this server against others.
func (r *APICreatePoolServerRequest) Weight(weight int32) *APICreatePoolServerRequest {
	r.weight = &weight
	return r
}

// MaxConn Maximum number of connections. If the value is &#x60;0&#x60;, it inherits the value from pool&#39;s &#x60;max_conn_default&#x60;.
func (r *APICreatePoolServerRequest) MaxConn(maxConn int32) *APICreatePoolServerRequest {
	r.maxConn = &maxConn
	return r
}

// Port Port number. Setting port &#x60;443&#x60; does not force TLS. Set &#x60;use_tls&#x60; in pool to force TLS.
func (r *APICreatePoolServerRequest) Port(port int32) *APICreatePoolServerRequest {
	r.port = &port
	return r
}

// Address A hostname, IPv4, or IPv6 address for the server. Required.
func (r *APICreatePoolServerRequest) Address(address string) *APICreatePoolServerRequest {
	r.address = &address
	return r
}

// Comment A freeform descriptive note.
func (r *APICreatePoolServerRequest) Comment(comment string) *APICreatePoolServerRequest {
	r.comment = &comment
	return r
}

// Disabled Allows servers to be enabled and disabled in a pool.
func (r *APICreatePoolServerRequest) Disabled(disabled bool) *APICreatePoolServerRequest {
	r.disabled = &disabled
	return r
}

// OverrideHost The hostname to override the Host header. Defaults to &#x60;null&#x60; meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting.
func (r *APICreatePoolServerRequest) OverrideHost(overrideHost string) *APICreatePoolServerRequest {
	r.overrideHost = &overrideHost
	return r
}

// Execute calls the API using the request data configured.
func (r APICreatePoolServerRequest) Execute() (*ServerResponse, *http.Response, error) {
	return r.APIService.CreatePoolServerExecute(r)
}

/*
CreatePoolServer Add a server to a pool

Creates a single server for a particular service and pool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param poolID Alphanumeric string identifying a Pool.
 @return APICreatePoolServerRequest
*/
func (a *ServerAPIService) CreatePoolServer(ctx context.Context, serviceID string, poolID string) APICreatePoolServerRequest {
	return APICreatePoolServerRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		poolID:     poolID,
	}
}

// CreatePoolServerExecute executes the request
//  @return ServerResponse
func (a *ServerAPIService) CreatePoolServerExecute(r APICreatePoolServerRequest) (*ServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerAPIService.CreatePoolServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/pool/{pool_id}/server"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_id"+"}", gourl.PathEscape(parameterToString(r.poolID, "")))

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
	if r.weight != nil {
		localVarFormParams.Add("weight", parameterToString(*r.weight, ""))
	}
	if r.maxConn != nil {
		localVarFormParams.Add("max_conn", parameterToString(*r.maxConn, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.disabled != nil {
		localVarFormParams.Add("disabled", parameterToString(*r.disabled, ""))
	}
	if r.overrideHost != nil {
		localVarFormParams.Add("override_host", parameterToString(*r.overrideHost, ""))
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

// APIDeletePoolServerRequest represents a request for the resource.
type APIDeletePoolServerRequest struct {
	ctx        context.Context
	APIService ServerAPI
	serviceID  string
	poolID     string
	serverID   string
}

// Execute calls the API using the request data configured.
func (r APIDeletePoolServerRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeletePoolServerExecute(r)
}

/*
DeletePoolServer Delete a server from a pool

Deletes a single server for a particular service and pool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param poolID Alphanumeric string identifying a Pool.
 @param serverID Alphanumeric string identifying a Server.
 @return APIDeletePoolServerRequest
*/
func (a *ServerAPIService) DeletePoolServer(ctx context.Context, serviceID string, poolID string, serverID string) APIDeletePoolServerRequest {
	return APIDeletePoolServerRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		poolID:     poolID,
		serverID:   serverID,
	}
}

// DeletePoolServerExecute executes the request
//  @return InlineResponse200
func (a *ServerAPIService) DeletePoolServerExecute(r APIDeletePoolServerRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerAPIService.DeletePoolServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/pool/{pool_id}/server/{server_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_id"+"}", gourl.PathEscape(parameterToString(r.poolID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"server_id"+"}", gourl.PathEscape(parameterToString(r.serverID, "")))

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

// APIGetPoolServerRequest represents a request for the resource.
type APIGetPoolServerRequest struct {
	ctx        context.Context
	APIService ServerAPI
	serviceID  string
	poolID     string
	serverID   string
}

// Execute calls the API using the request data configured.
func (r APIGetPoolServerRequest) Execute() (*ServerResponse, *http.Response, error) {
	return r.APIService.GetPoolServerExecute(r)
}

/*
GetPoolServer Get a pool server

Gets a single server for a particular service and pool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param poolID Alphanumeric string identifying a Pool.
 @param serverID Alphanumeric string identifying a Server.
 @return APIGetPoolServerRequest
*/
func (a *ServerAPIService) GetPoolServer(ctx context.Context, serviceID string, poolID string, serverID string) APIGetPoolServerRequest {
	return APIGetPoolServerRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		poolID:     poolID,
		serverID:   serverID,
	}
}

// GetPoolServerExecute executes the request
//  @return ServerResponse
func (a *ServerAPIService) GetPoolServerExecute(r APIGetPoolServerRequest) (*ServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerAPIService.GetPoolServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/pool/{pool_id}/server/{server_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_id"+"}", gourl.PathEscape(parameterToString(r.poolID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"server_id"+"}", gourl.PathEscape(parameterToString(r.serverID, "")))

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

// APIListPoolServersRequest represents a request for the resource.
type APIListPoolServersRequest struct {
	ctx        context.Context
	APIService ServerAPI
	serviceID  string
	poolID     string
}

// Execute calls the API using the request data configured.
func (r APIListPoolServersRequest) Execute() ([]ServerResponse, *http.Response, error) {
	return r.APIService.ListPoolServersExecute(r)
}

/*
ListPoolServers List servers in a pool

Lists all servers for a particular service and pool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param poolID Alphanumeric string identifying a Pool.
 @return APIListPoolServersRequest
*/
func (a *ServerAPIService) ListPoolServers(ctx context.Context, serviceID string, poolID string) APIListPoolServersRequest {
	return APIListPoolServersRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		poolID:     poolID,
	}
}

// ListPoolServersExecute executes the request
//  @return []ServerResponse
func (a *ServerAPIService) ListPoolServersExecute(r APIListPoolServersRequest) ([]ServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []ServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerAPIService.ListPoolServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/pool/{pool_id}/servers"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_id"+"}", gourl.PathEscape(parameterToString(r.poolID, "")))

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

// APIUpdatePoolServerRequest represents a request for the resource.
type APIUpdatePoolServerRequest struct {
	ctx          context.Context
	APIService   ServerAPI
	serviceID    string
	poolID       string
	serverID     string
	weight       *int32
	maxConn      *int32
	port         *int32
	address      *string
	comment      *string
	disabled     *bool
	overrideHost *string
}

// Weight Weight (&#x60;1-100&#x60;) used to load balance this server against others.
func (r *APIUpdatePoolServerRequest) Weight(weight int32) *APIUpdatePoolServerRequest {
	r.weight = &weight
	return r
}

// MaxConn Maximum number of connections. If the value is &#x60;0&#x60;, it inherits the value from pool&#39;s &#x60;max_conn_default&#x60;.
func (r *APIUpdatePoolServerRequest) MaxConn(maxConn int32) *APIUpdatePoolServerRequest {
	r.maxConn = &maxConn
	return r
}

// Port Port number. Setting port &#x60;443&#x60; does not force TLS. Set &#x60;use_tls&#x60; in pool to force TLS.
func (r *APIUpdatePoolServerRequest) Port(port int32) *APIUpdatePoolServerRequest {
	r.port = &port
	return r
}

// Address A hostname, IPv4, or IPv6 address for the server. Required.
func (r *APIUpdatePoolServerRequest) Address(address string) *APIUpdatePoolServerRequest {
	r.address = &address
	return r
}

// Comment A freeform descriptive note.
func (r *APIUpdatePoolServerRequest) Comment(comment string) *APIUpdatePoolServerRequest {
	r.comment = &comment
	return r
}

// Disabled Allows servers to be enabled and disabled in a pool.
func (r *APIUpdatePoolServerRequest) Disabled(disabled bool) *APIUpdatePoolServerRequest {
	r.disabled = &disabled
	return r
}

// OverrideHost The hostname to override the Host header. Defaults to &#x60;null&#x60; meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting.
func (r *APIUpdatePoolServerRequest) OverrideHost(overrideHost string) *APIUpdatePoolServerRequest {
	r.overrideHost = &overrideHost
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdatePoolServerRequest) Execute() (*ServerResponse, *http.Response, error) {
	return r.APIService.UpdatePoolServerExecute(r)
}

/*
UpdatePoolServer Update a server

Updates a single server for a particular service and pool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param poolID Alphanumeric string identifying a Pool.
 @param serverID Alphanumeric string identifying a Server.
 @return APIUpdatePoolServerRequest
*/
func (a *ServerAPIService) UpdatePoolServer(ctx context.Context, serviceID string, poolID string, serverID string) APIUpdatePoolServerRequest {
	return APIUpdatePoolServerRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		poolID:     poolID,
		serverID:   serverID,
	}
}

// UpdatePoolServerExecute executes the request
//  @return ServerResponse
func (a *ServerAPIService) UpdatePoolServerExecute(r APIUpdatePoolServerRequest) (*ServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerAPIService.UpdatePoolServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/pool/{pool_id}/server/{server_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"pool_id"+"}", gourl.PathEscape(parameterToString(r.poolID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"server_id"+"}", gourl.PathEscape(parameterToString(r.serverID, "")))

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
	if r.weight != nil {
		localVarFormParams.Add("weight", parameterToString(*r.weight, ""))
	}
	if r.maxConn != nil {
		localVarFormParams.Add("max_conn", parameterToString(*r.maxConn, ""))
	}
	if r.port != nil {
		localVarFormParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.address != nil {
		localVarFormParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.disabled != nil {
		localVarFormParams.Add("disabled", parameterToString(*r.disabled, ""))
	}
	if r.overrideHost != nil {
		localVarFormParams.Add("override_host", parameterToString(*r.overrideHost, ""))
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
