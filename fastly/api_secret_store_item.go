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

// SecretStoreItemAPI defines an interface for interacting with the resource.
type SecretStoreItemAPI interface {

	/*
	CreateSecret Create a new secret in a store.

	Create a new secret in a store.
Returns an error if a secret already exists with the same name.
See `PUT` and `PATCH` methods for ways to recreate an existing secret.

The `secret` field must be Base64-encoded because a secret can contain binary data.
In the example below, the unencoded secret is "Hello, world!"


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param storeID
	 @return APICreateSecretRequest
	*/
	CreateSecret(ctx context.Context, storeID string) APICreateSecretRequest

	// CreateSecretExecute executes the request
	//  @return SecretResponse
	CreateSecretExecute(r APICreateSecretRequest) (*SecretResponse, *http.Response, error)

	/*
	DeleteSecret Delete a secret from a store.

	Delete a secret from a store by name.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param storeID
	 @param secretName
	 @return APIDeleteSecretRequest
	*/
	DeleteSecret(ctx context.Context, storeID string, secretName string) APIDeleteSecretRequest

	// DeleteSecretExecute executes the request
	DeleteSecretExecute(r APIDeleteSecretRequest) (*http.Response, error)

	/*
	GetSecret Get secret metadata.

	Get metadata about a secret by name.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param storeID
	 @param secretName
	 @return APIGetSecretRequest
	*/
	GetSecret(ctx context.Context, storeID string, secretName string) APIGetSecretRequest

	// GetSecretExecute executes the request
	//  @return SecretResponse
	GetSecretExecute(r APIGetSecretRequest) (*SecretResponse, *http.Response, error)

	/*
	GetSecrets List secrets within a store.

	List all secrets within a store.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param storeID
	 @return APIGetSecretsRequest
	*/
	GetSecrets(ctx context.Context, storeID string) APIGetSecretsRequest

	// GetSecretsExecute executes the request
	//  @return InlineResponse2006
	GetSecretsExecute(r APIGetSecretsRequest) (*InlineResponse2006, *http.Response, error)

	/*
	MustRecreateSecret Recreate a secret in a store.

	Recreate a secret based on the secret's name.
Returns an error if there is no existing secret with the same name.

The `secret` field must be Base64-encoded because a secret can contain binary data.
In the example below, the unencoded secret is "Hello, world!"


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param storeID
	 @return APIMustRecreateSecretRequest
	*/
	MustRecreateSecret(ctx context.Context, storeID string) APIMustRecreateSecretRequest

	// MustRecreateSecretExecute executes the request
	//  @return SecretResponse
	MustRecreateSecretExecute(r APIMustRecreateSecretRequest) (*SecretResponse, *http.Response, error)

	/*
	RecreateSecret Create or recreate a secret in a store.

	Create or recreate a secret based on the secret's name.
The response object's `recreated` field will be true if the secret was recreated.

The `secret` field must be Base64-encoded because a secret can contain binary data.
In the example below, the unencoded secret is "Hello, world!"


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param storeID
	 @return APIRecreateSecretRequest
	*/
	RecreateSecret(ctx context.Context, storeID string) APIRecreateSecretRequest

	// RecreateSecretExecute executes the request
	//  @return SecretResponse
	RecreateSecretExecute(r APIRecreateSecretRequest) (*SecretResponse, *http.Response, error)
}

// SecretStoreItemAPIService SecretStoreItemAPI service
type SecretStoreItemAPIService service

// APICreateSecretRequest represents a request for the resource.
type APICreateSecretRequest struct {
	ctx context.Context
	APIService SecretStoreItemAPI
	storeID string
	secret *Secret
}

// Secret returns a pointer to a request.
func (r *APICreateSecretRequest) Secret(secret Secret) *APICreateSecretRequest {
	r.secret = &secret
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.APIService.CreateSecretExecute(r)
}

/*
CreateSecret Create a new secret in a store.

Create a new secret in a store.
Returns an error if a secret already exists with the same name.
See `PUT` and `PATCH` methods for ways to recreate an existing secret.

The `secret` field must be Base64-encoded because a secret can contain binary data.
In the example below, the unencoded secret is "Hello, world!"


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @return APICreateSecretRequest
*/
func (a *SecretStoreItemAPIService) CreateSecret(ctx context.Context, storeID string) APICreateSecretRequest {
	return APICreateSecretRequest{
		APIService: a,
		ctx: ctx,
		storeID: storeID,
	}
}

// CreateSecretExecute executes the request
//  @return SecretResponse
func (a *SecretStoreItemAPIService) CreateSecretExecute(r APICreateSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretStoreItemAPIService.CreateSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/secret/{store_id}/secrets"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.secret
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

// APIDeleteSecretRequest represents a request for the resource.
type APIDeleteSecretRequest struct {
	ctx context.Context
	APIService SecretStoreItemAPI
	storeID string
	secretName string
}


// Execute calls the API using the request data configured.
func (r APIDeleteSecretRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteSecretExecute(r)
}

/*
DeleteSecret Delete a secret from a store.

Delete a secret from a store by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param secretName
 @return APIDeleteSecretRequest
*/
func (a *SecretStoreItemAPIService) DeleteSecret(ctx context.Context, storeID string, secretName string) APIDeleteSecretRequest {
	return APIDeleteSecretRequest{
		APIService: a,
		ctx: ctx,
		storeID: storeID,
		secretName: secretName,
	}
}

// DeleteSecretExecute executes the request
func (a *SecretStoreItemAPIService) DeleteSecretExecute(r APIDeleteSecretRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretStoreItemAPIService.DeleteSecret")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/secret/{store_id}/secrets/{secret_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"secret_name"+"}", gourl.PathEscape(parameterToString(r.secretName, "")))

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
	localVarHTTPHeaderAccepts := []string{}

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

// APIGetSecretRequest represents a request for the resource.
type APIGetSecretRequest struct {
	ctx context.Context
	APIService SecretStoreItemAPI
	storeID string
	secretName string
}


// Execute calls the API using the request data configured.
func (r APIGetSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.APIService.GetSecretExecute(r)
}

/*
GetSecret Get secret metadata.

Get metadata about a secret by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @param secretName
 @return APIGetSecretRequest
*/
func (a *SecretStoreItemAPIService) GetSecret(ctx context.Context, storeID string, secretName string) APIGetSecretRequest {
	return APIGetSecretRequest{
		APIService: a,
		ctx: ctx,
		storeID: storeID,
		secretName: secretName,
	}
}

// GetSecretExecute executes the request
//  @return SecretResponse
func (a *SecretStoreItemAPIService) GetSecretExecute(r APIGetSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretStoreItemAPIService.GetSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/secret/{store_id}/secrets/{secret_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"secret_name"+"}", gourl.PathEscape(parameterToString(r.secretName, "")))

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

// APIGetSecretsRequest represents a request for the resource.
type APIGetSecretsRequest struct {
	ctx context.Context
	APIService SecretStoreItemAPI
	storeID string
	cursor *string
	limit *string
}

// Cursor Cursor value from a previous response to retrieve the next page. To request the first page, this should be empty.
func (r *APIGetSecretsRequest) Cursor(cursor string) *APIGetSecretsRequest {
	r.cursor = &cursor
	return r
}
// Limit Number of results per page. The maximum is 200.
func (r *APIGetSecretsRequest) Limit(limit string) *APIGetSecretsRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetSecretsRequest) Execute() (*InlineResponse2006, *http.Response, error) {
	return r.APIService.GetSecretsExecute(r)
}

/*
GetSecrets List secrets within a store.

List all secrets within a store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @return APIGetSecretsRequest
*/
func (a *SecretStoreItemAPIService) GetSecrets(ctx context.Context, storeID string) APIGetSecretsRequest {
	return APIGetSecretsRequest{
		APIService: a,
		ctx: ctx,
		storeID: storeID,
	}
}

// GetSecretsExecute executes the request
//  @return InlineResponse2006
func (a *SecretStoreItemAPIService) GetSecretsExecute(r APIGetSecretsRequest) (*InlineResponse2006, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2006
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretStoreItemAPIService.GetSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/secret/{store_id}/secrets"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
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

// APIMustRecreateSecretRequest represents a request for the resource.
type APIMustRecreateSecretRequest struct {
	ctx context.Context
	APIService SecretStoreItemAPI
	storeID string
	secret *Secret
}

// Secret returns a pointer to a request.
func (r *APIMustRecreateSecretRequest) Secret(secret Secret) *APIMustRecreateSecretRequest {
	r.secret = &secret
	return r
}

// Execute calls the API using the request data configured.
func (r APIMustRecreateSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.APIService.MustRecreateSecretExecute(r)
}

/*
MustRecreateSecret Recreate a secret in a store.

Recreate a secret based on the secret's name.
Returns an error if there is no existing secret with the same name.

The `secret` field must be Base64-encoded because a secret can contain binary data.
In the example below, the unencoded secret is "Hello, world!"


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @return APIMustRecreateSecretRequest
*/
func (a *SecretStoreItemAPIService) MustRecreateSecret(ctx context.Context, storeID string) APIMustRecreateSecretRequest {
	return APIMustRecreateSecretRequest{
		APIService: a,
		ctx: ctx,
		storeID: storeID,
	}
}

// MustRecreateSecretExecute executes the request
//  @return SecretResponse
func (a *SecretStoreItemAPIService) MustRecreateSecretExecute(r APIMustRecreateSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretStoreItemAPIService.MustRecreateSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/secret/{store_id}/secrets"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.secret
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

// APIRecreateSecretRequest represents a request for the resource.
type APIRecreateSecretRequest struct {
	ctx context.Context
	APIService SecretStoreItemAPI
	storeID string
	secret *Secret
}

// Secret returns a pointer to a request.
func (r *APIRecreateSecretRequest) Secret(secret Secret) *APIRecreateSecretRequest {
	r.secret = &secret
	return r
}

// Execute calls the API using the request data configured.
func (r APIRecreateSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.APIService.RecreateSecretExecute(r)
}

/*
RecreateSecret Create or recreate a secret in a store.

Create or recreate a secret based on the secret's name.
The response object's `recreated` field will be true if the secret was recreated.

The `secret` field must be Base64-encoded because a secret can contain binary data.
In the example below, the unencoded secret is "Hello, world!"


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param storeID
 @return APIRecreateSecretRequest
*/
func (a *SecretStoreItemAPIService) RecreateSecret(ctx context.Context, storeID string) APIRecreateSecretRequest {
	return APIRecreateSecretRequest{
		APIService: a,
		ctx: ctx,
		storeID: storeID,
	}
}

// RecreateSecretExecute executes the request
//  @return SecretResponse
func (a *SecretStoreItemAPIService) RecreateSecretExecute(r APIRecreateSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretStoreItemAPIService.RecreateSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/stores/secret/{store_id}/secrets"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"store_id"+"}", gourl.PathEscape(parameterToString(r.storeID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.secret
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
