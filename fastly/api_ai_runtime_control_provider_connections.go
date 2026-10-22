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

// AiRuntimeControlProviderConnectionsAPI defines an interface for interacting with the resource.
type AiRuntimeControlProviderConnectionsAPI interface {

	/*
		ArcCreateProviderConnection Create a provider connection

		Create a provider connection. Set `auth_type` to `api-key` (the default) to connect
	with a provider's own API key, or to `aws-iam` to connect to Amazon Bedrock using an
	AWS IAM role.

	Accounts are limited to 25 provider connections. Creating a 26th connection returns a
	409 response.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcCreateProviderConnectionRequest
	*/
	ArcCreateProviderConnection(ctx context.Context) APIArcCreateProviderConnectionRequest

	// ArcCreateProviderConnectionExecute executes the request
	//  @return ArcProviderConnectionResponse
	ArcCreateProviderConnectionExecute(r APIArcCreateProviderConnectionRequest) (*ArcProviderConnectionResponse, *http.Response, error)

	/*
		ArcDeleteProviderConnection Delete a provider connection

		Delete a provider connection. Returns a 409 response if the connection is still
	referenced by a virtual key's failover policy.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param id The ID of the provider connection.
		 @return APIArcDeleteProviderConnectionRequest
	*/
	ArcDeleteProviderConnection(ctx context.Context, id string) APIArcDeleteProviderConnectionRequest

	// ArcDeleteProviderConnectionExecute executes the request
	ArcDeleteProviderConnectionExecute(r APIArcDeleteProviderConnectionRequest) (*http.Response, error)

	/*
		ArcGetProviderConnection Get a provider connection

		Retrieve a single provider connection.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param id The ID of the provider connection.
		 @return APIArcGetProviderConnectionRequest
	*/
	ArcGetProviderConnection(ctx context.Context, id string) APIArcGetProviderConnectionRequest

	// ArcGetProviderConnectionExecute executes the request
	//  @return ArcGetProviderConnectionResponse
	ArcGetProviderConnectionExecute(r APIArcGetProviderConnectionRequest) (*ArcGetProviderConnectionResponse, *http.Response, error)

	/*
		ArcListProviderConnections List provider connections

		List the provider connections belonging to your account.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcListProviderConnectionsRequest
	*/
	ArcListProviderConnections(ctx context.Context) APIArcListProviderConnectionsRequest

	// ArcListProviderConnectionsExecute executes the request
	//  @return ArcProviderConnectionsResponse
	ArcListProviderConnectionsExecute(r APIArcListProviderConnectionsRequest) (*ArcProviderConnectionsResponse, *http.Response, error)

	/*
		ArcUpdateProviderConnection Update a provider connection

		Update the attributes of a provider connection. Only the fields you include are
	changed. Which fields you can set depends on the connection's existing `auth_type`:
	`base_url` and `api_key` apply only to `api-key` connections, and `customer_role_arn`
	and `region` apply only to `aws-iam` connections. `auth_type` itself cannot be changed.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param id The ID of the provider connection.
		 @return APIArcUpdateProviderConnectionRequest
	*/
	ArcUpdateProviderConnection(ctx context.Context, id string) APIArcUpdateProviderConnectionRequest

	// ArcUpdateProviderConnectionExecute executes the request
	//  @return ArcProviderConnectionResponse
	ArcUpdateProviderConnectionExecute(r APIArcUpdateProviderConnectionRequest) (*ArcProviderConnectionResponse, *http.Response, error)
}

// AiRuntimeControlProviderConnectionsAPIService AiRuntimeControlProviderConnectionsAPI service
type AiRuntimeControlProviderConnectionsAPIService service

// APIArcCreateProviderConnectionRequest represents a request for the resource.
type APIArcCreateProviderConnectionRequest struct {
	ctx                          context.Context
	APIService                   AiRuntimeControlProviderConnectionsAPI
	arcProviderConnectionRequest *ArcProviderConnectionRequest
}

// ArcProviderConnectionRequest returns a pointer to a request.
func (r *APIArcCreateProviderConnectionRequest) ArcProviderConnectionRequest(arcProviderConnectionRequest ArcProviderConnectionRequest) *APIArcCreateProviderConnectionRequest {
	r.arcProviderConnectionRequest = &arcProviderConnectionRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcCreateProviderConnectionRequest) Execute() (*ArcProviderConnectionResponse, *http.Response, error) {
	return r.APIService.ArcCreateProviderConnectionExecute(r)
}

/*
ArcCreateProviderConnection Create a provider connection

Create a provider connection. Set `auth_type` to `api-key` (the default) to connect
with a provider's own API key, or to `aws-iam` to connect to Amazon Bedrock using an
AWS IAM role.

Accounts are limited to 25 provider connections. Creating a 26th connection returns a
409 response.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcCreateProviderConnectionRequest
*/
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcCreateProviderConnection(ctx context.Context) APIArcCreateProviderConnectionRequest {
	return APIArcCreateProviderConnectionRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcCreateProviderConnectionExecute executes the request
//  @return ArcProviderConnectionResponse
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcCreateProviderConnectionExecute(r APIArcCreateProviderConnectionRequest) (*ArcProviderConnectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcProviderConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProviderConnectionsAPIService.ArcCreateProviderConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/provider-connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.arcProviderConnectionRequest == nil {
		return localVarReturnValue, nil, reportError("arcProviderConnectionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.arcProviderConnectionRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcDeleteProviderConnectionRequest represents a request for the resource.
type APIArcDeleteProviderConnectionRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlProviderConnectionsAPI
	id         string
}

// Execute calls the API using the request data configured.
func (r APIArcDeleteProviderConnectionRequest) Execute() (*http.Response, error) {
	return r.APIService.ArcDeleteProviderConnectionExecute(r)
}

/*
ArcDeleteProviderConnection Delete a provider connection

Delete a provider connection. Returns a 409 response if the connection is still
referenced by a virtual key's failover policy.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the provider connection.
 @return APIArcDeleteProviderConnectionRequest
*/
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcDeleteProviderConnection(ctx context.Context, id string) APIArcDeleteProviderConnectionRequest {
	return APIArcDeleteProviderConnectionRequest{
		APIService: a,
		ctx:        ctx,
		id:         id,
	}
}

// ArcDeleteProviderConnectionExecute executes the request
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcDeleteProviderConnectionExecute(r APIArcDeleteProviderConnectionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProviderConnectionsAPIService.ArcDeleteProviderConnection")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/provider-connections/{id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"id"+"}", gourl.PathEscape(parameterToString(r.id, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
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

	return localVarHTTPResponse, nil
}

// APIArcGetProviderConnectionRequest represents a request for the resource.
type APIArcGetProviderConnectionRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlProviderConnectionsAPI
	id         string
}

// Execute calls the API using the request data configured.
func (r APIArcGetProviderConnectionRequest) Execute() (*ArcGetProviderConnectionResponse, *http.Response, error) {
	return r.APIService.ArcGetProviderConnectionExecute(r)
}

/*
ArcGetProviderConnection Get a provider connection

Retrieve a single provider connection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the provider connection.
 @return APIArcGetProviderConnectionRequest
*/
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcGetProviderConnection(ctx context.Context, id string) APIArcGetProviderConnectionRequest {
	return APIArcGetProviderConnectionRequest{
		APIService: a,
		ctx:        ctx,
		id:         id,
	}
}

// ArcGetProviderConnectionExecute executes the request
//  @return ArcGetProviderConnectionResponse
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcGetProviderConnectionExecute(r APIArcGetProviderConnectionRequest) (*ArcGetProviderConnectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcGetProviderConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProviderConnectionsAPIService.ArcGetProviderConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/provider-connections/{id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"id"+"}", gourl.PathEscape(parameterToString(r.id, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcListProviderConnectionsRequest represents a request for the resource.
type APIArcListProviderConnectionsRequest struct {
	ctx            context.Context
	APIService     AiRuntimeControlProviderConnectionsAPI
	includeDeleted *bool
	sort           *string
}

// IncludeDeleted Include deleted provider connections in the results. Deleted connections have a &#x60;deleted_at&#x60; value.
func (r *APIArcListProviderConnectionsRequest) IncludeDeleted(includeDeleted bool) *APIArcListProviderConnectionsRequest {
	r.includeDeleted = &includeDeleted
	return r
}

// Sort Sort field. Prefix with &#x60;-&#x60; for descending order (e.g., &#x60;sort&#x3D;-name&#x60;).
func (r *APIArcListProviderConnectionsRequest) Sort(sort string) *APIArcListProviderConnectionsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcListProviderConnectionsRequest) Execute() (*ArcProviderConnectionsResponse, *http.Response, error) {
	return r.APIService.ArcListProviderConnectionsExecute(r)
}

/*
ArcListProviderConnections List provider connections

List the provider connections belonging to your account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcListProviderConnectionsRequest
*/
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcListProviderConnections(ctx context.Context) APIArcListProviderConnectionsRequest {
	return APIArcListProviderConnectionsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcListProviderConnectionsExecute executes the request
//  @return ArcProviderConnectionsResponse
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcListProviderConnectionsExecute(r APIArcListProviderConnectionsRequest) (*ArcProviderConnectionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcProviderConnectionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProviderConnectionsAPIService.ArcListProviderConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/provider-connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.includeDeleted != nil {
		localVarQueryParams.Add("include_deleted", parameterToString(*r.includeDeleted, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// APIArcUpdateProviderConnectionRequest represents a request for the resource.
type APIArcUpdateProviderConnectionRequest struct {
	ctx                                context.Context
	APIService                         AiRuntimeControlProviderConnectionsAPI
	id                                 string
	arcUpdateProviderConnectionRequest *ArcUpdateProviderConnectionRequest
}

// ArcUpdateProviderConnectionRequest returns a pointer to a request.
func (r *APIArcUpdateProviderConnectionRequest) ArcUpdateProviderConnectionRequest(arcUpdateProviderConnectionRequest ArcUpdateProviderConnectionRequest) *APIArcUpdateProviderConnectionRequest {
	r.arcUpdateProviderConnectionRequest = &arcUpdateProviderConnectionRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcUpdateProviderConnectionRequest) Execute() (*ArcProviderConnectionResponse, *http.Response, error) {
	return r.APIService.ArcUpdateProviderConnectionExecute(r)
}

/*
ArcUpdateProviderConnection Update a provider connection

Update the attributes of a provider connection. Only the fields you include are
changed. Which fields you can set depends on the connection's existing `auth_type`:
`base_url` and `api_key` apply only to `api-key` connections, and `customer_role_arn`
and `region` apply only to `aws-iam` connections. `auth_type` itself cannot be changed.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the provider connection.
 @return APIArcUpdateProviderConnectionRequest
*/
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcUpdateProviderConnection(ctx context.Context, id string) APIArcUpdateProviderConnectionRequest {
	return APIArcUpdateProviderConnectionRequest{
		APIService: a,
		ctx:        ctx,
		id:         id,
	}
}

// ArcUpdateProviderConnectionExecute executes the request
//  @return ArcProviderConnectionResponse
func (a *AiRuntimeControlProviderConnectionsAPIService) ArcUpdateProviderConnectionExecute(r APIArcUpdateProviderConnectionRequest) (*ArcProviderConnectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcProviderConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProviderConnectionsAPIService.ArcUpdateProviderConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/provider-connections/{id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"id"+"}", gourl.PathEscape(parameterToString(r.id, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.arcUpdateProviderConnectionRequest == nil {
		return localVarReturnValue, nil, reportError("arcUpdateProviderConnectionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.arcUpdateProviderConnectionRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
