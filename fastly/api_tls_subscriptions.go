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

// TLSSubscriptionsAPI defines an interface for interacting with the resource.
type TLSSubscriptionsAPI interface {

	/*
	CreateGlobalsignEmailChallenge Creates a GlobalSign email challenge.

	Creates an email challenge for a domain on a GlobalSign subscription. An email challenge will generate an email that can be used to validate domain ownership. If this challenge is created, then the domain can only be validated using email for the given subscription.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
	 @param tlsAuthorizationID Alphanumeric string identifying a TLS subscription.
	 @return APICreateGlobalsignEmailChallengeRequest
	*/
	CreateGlobalsignEmailChallenge(ctx context.Context, tlsSubscriptionID string, tlsAuthorizationID string) APICreateGlobalsignEmailChallengeRequest

	// CreateGlobalsignEmailChallengeExecute executes the request
	//  @return map[string]any
	CreateGlobalsignEmailChallengeExecute(r APICreateGlobalsignEmailChallengeRequest) (map[string]any, *http.Response, error)

	/*
	CreateTLSSub Create a TLS subscription

	Create a new TLS subscription. This response includes a list of possible challenges to verify domain ownership.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APICreateTLSSubRequest
	*/
	CreateTLSSub(ctx context.Context) APICreateTLSSubRequest

	// CreateTLSSubExecute executes the request
	//  @return TLSSubscriptionResponse
	CreateTLSSubExecute(r APICreateTLSSubRequest) (*TLSSubscriptionResponse, *http.Response, error)

	/*
	DeleteGlobalsignEmailChallenge Delete a GlobalSign email challenge

	Deletes a GlobalSign email challenge. After a GlobalSign email challenge is deleted, the domain can use HTTP and DNS validation methods again.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
	 @param globalsignEmailChallengeID Alphanumeric string identifying a GlobalSign email challenge.
	 @param tlsAuthorizationID Alphanumeric string identifying a TLS subscription.
	 @return APIDeleteGlobalsignEmailChallengeRequest
	*/
	DeleteGlobalsignEmailChallenge(ctx context.Context, tlsSubscriptionID string, globalsignEmailChallengeID string, tlsAuthorizationID string) APIDeleteGlobalsignEmailChallengeRequest

	// DeleteGlobalsignEmailChallengeExecute executes the request
	DeleteGlobalsignEmailChallengeExecute(r APIDeleteGlobalsignEmailChallengeRequest) (*http.Response, error)

	/*
	DeleteTLSSub Delete a TLS subscription

	Destroy a TLS subscription. A subscription cannot be destroyed if there are domains in the TLS enabled state.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
	 @return APIDeleteTLSSubRequest
	*/
	DeleteTLSSub(ctx context.Context, tlsSubscriptionID string) APIDeleteTLSSubRequest

	// DeleteTLSSubExecute executes the request
	DeleteTLSSubExecute(r APIDeleteTLSSubRequest) (*http.Response, error)

	/*
	GetTLSSub Get a TLS subscription

	Show a TLS subscription.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
	 @return APIGetTLSSubRequest
	*/
	GetTLSSub(ctx context.Context, tlsSubscriptionID string) APIGetTLSSubRequest

	// GetTLSSubExecute executes the request
	//  @return TLSSubscriptionResponse
	GetTLSSubExecute(r APIGetTLSSubRequest) (*TLSSubscriptionResponse, *http.Response, error)

	/*
	ListTLSSubs List TLS subscriptions

	List all TLS subscriptions.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListTLSSubsRequest
	*/
	ListTLSSubs(ctx context.Context) APIListTLSSubsRequest

	// ListTLSSubsExecute executes the request
	//  @return TLSSubscriptionsResponse
	ListTLSSubsExecute(r APIListTLSSubsRequest) (*TLSSubscriptionsResponse, *http.Response, error)

	/*
	PatchTLSSub Update a TLS subscription

	Change the TLS domains or common name associated with this subscription, update the TLS configuration for this set of domains, or retry a subscription with state `failed` by setting the state to `retry`.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
	 @return APIPatchTLSSubRequest
	*/
	PatchTLSSub(ctx context.Context, tlsSubscriptionID string) APIPatchTLSSubRequest

	// PatchTLSSubExecute executes the request
	//  @return TLSSubscriptionResponse
	PatchTLSSubExecute(r APIPatchTLSSubRequest) (*TLSSubscriptionResponse, *http.Response, error)
}

// TLSSubscriptionsAPIService TLSSubscriptionsAPI service
type TLSSubscriptionsAPIService service

// APICreateGlobalsignEmailChallengeRequest represents a request for the resource.
type APICreateGlobalsignEmailChallengeRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	tlsSubscriptionID string
	tlsAuthorizationID string
	requestBody *map[string]map[string]any
}

// RequestBody returns a pointer to a request.
func (r *APICreateGlobalsignEmailChallengeRequest) RequestBody(requestBody map[string]map[string]any) *APICreateGlobalsignEmailChallengeRequest {
	r.requestBody = &requestBody
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateGlobalsignEmailChallengeRequest) Execute() (map[string]any, *http.Response, error) {
	return r.APIService.CreateGlobalsignEmailChallengeExecute(r)
}

/*
CreateGlobalsignEmailChallenge Creates a GlobalSign email challenge.

Creates an email challenge for a domain on a GlobalSign subscription. An email challenge will generate an email that can be used to validate domain ownership. If this challenge is created, then the domain can only be validated using email for the given subscription.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
 @param tlsAuthorizationID Alphanumeric string identifying a TLS subscription.
 @return APICreateGlobalsignEmailChallengeRequest
*/
func (a *TLSSubscriptionsAPIService) CreateGlobalsignEmailChallenge(ctx context.Context, tlsSubscriptionID string, tlsAuthorizationID string) APICreateGlobalsignEmailChallengeRequest {
	return APICreateGlobalsignEmailChallengeRequest{
		APIService: a,
		ctx: ctx,
		tlsSubscriptionID: tlsSubscriptionID,
		tlsAuthorizationID: tlsAuthorizationID,
	}
}

// CreateGlobalsignEmailChallengeExecute executes the request
//  @return map[string]any
func (a *TLSSubscriptionsAPIService) CreateGlobalsignEmailChallengeExecute(r APICreateGlobalsignEmailChallengeRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.CreateGlobalsignEmailChallenge")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_subscription_id"+"}", gourl.PathEscape(parameterToString(r.tlsSubscriptionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_authorization_id"+"}", gourl.PathEscape(parameterToString(r.tlsAuthorizationID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	localVarPostBody = r.requestBody
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

// APICreateTLSSubRequest represents a request for the resource.
type APICreateTLSSubRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	force *bool
	tlsSubscription *TLSSubscription
}

// Force A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain. 
func (r *APICreateTLSSubRequest) Force(force bool) *APICreateTLSSubRequest {
	r.force = &force
	return r
}
// TLSSubscription returns a pointer to a request.
func (r *APICreateTLSSubRequest) TLSSubscription(tlsSubscription TLSSubscription) *APICreateTLSSubRequest {
	r.tlsSubscription = &tlsSubscription
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateTLSSubRequest) Execute() (*TLSSubscriptionResponse, *http.Response, error) {
	return r.APIService.CreateTLSSubExecute(r)
}

/*
CreateTLSSub Create a TLS subscription

Create a new TLS subscription. This response includes a list of possible challenges to verify domain ownership.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateTLSSubRequest
*/
func (a *TLSSubscriptionsAPIService) CreateTLSSub(ctx context.Context) APICreateTLSSubRequest {
	return APICreateTLSSubRequest{
		APIService: a,
		ctx: ctx,
	}
}

// CreateTLSSubExecute executes the request
//  @return TLSSubscriptionResponse
func (a *TLSSubscriptionsAPIService) CreateTLSSubExecute(r APICreateTLSSubRequest) (*TLSSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.CreateTLSSub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tlsSubscription
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

// APIDeleteGlobalsignEmailChallengeRequest represents a request for the resource.
type APIDeleteGlobalsignEmailChallengeRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	tlsSubscriptionID string
	globalsignEmailChallengeID string
	tlsAuthorizationID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteGlobalsignEmailChallengeRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteGlobalsignEmailChallengeExecute(r)
}

/*
DeleteGlobalsignEmailChallenge Delete a GlobalSign email challenge

Deletes a GlobalSign email challenge. After a GlobalSign email challenge is deleted, the domain can use HTTP and DNS validation methods again.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
 @param globalsignEmailChallengeID Alphanumeric string identifying a GlobalSign email challenge.
 @param tlsAuthorizationID Alphanumeric string identifying a TLS subscription.
 @return APIDeleteGlobalsignEmailChallengeRequest
*/
func (a *TLSSubscriptionsAPIService) DeleteGlobalsignEmailChallenge(ctx context.Context, tlsSubscriptionID string, globalsignEmailChallengeID string, tlsAuthorizationID string) APIDeleteGlobalsignEmailChallengeRequest {
	return APIDeleteGlobalsignEmailChallengeRequest{
		APIService: a,
		ctx: ctx,
		tlsSubscriptionID: tlsSubscriptionID,
		globalsignEmailChallengeID: globalsignEmailChallengeID,
		tlsAuthorizationID: tlsAuthorizationID,
	}
}

// DeleteGlobalsignEmailChallengeExecute executes the request
func (a *TLSSubscriptionsAPIService) DeleteGlobalsignEmailChallengeExecute(r APIDeleteGlobalsignEmailChallengeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.DeleteGlobalsignEmailChallenge")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges/{globalsign_email_challenge_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_subscription_id"+"}", gourl.PathEscape(parameterToString(r.tlsSubscriptionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"globalsign_email_challenge_id"+"}", gourl.PathEscape(parameterToString(r.globalsignEmailChallengeID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_authorization_id"+"}", gourl.PathEscape(parameterToString(r.tlsAuthorizationID, "")))

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

// APIDeleteTLSSubRequest represents a request for the resource.
type APIDeleteTLSSubRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	tlsSubscriptionID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteTLSSubRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteTLSSubExecute(r)
}

/*
DeleteTLSSub Delete a TLS subscription

Destroy a TLS subscription. A subscription cannot be destroyed if there are domains in the TLS enabled state.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
 @return APIDeleteTLSSubRequest
*/
func (a *TLSSubscriptionsAPIService) DeleteTLSSub(ctx context.Context, tlsSubscriptionID string) APIDeleteTLSSubRequest {
	return APIDeleteTLSSubRequest{
		APIService: a,
		ctx: ctx,
		tlsSubscriptionID: tlsSubscriptionID,
	}
}

// DeleteTLSSubExecute executes the request
func (a *TLSSubscriptionsAPIService) DeleteTLSSubExecute(r APIDeleteTLSSubRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.DeleteTLSSub")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions/{tls_subscription_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_subscription_id"+"}", gourl.PathEscape(parameterToString(r.tlsSubscriptionID, "")))

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

// APIGetTLSSubRequest represents a request for the resource.
type APIGetTLSSubRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	tlsSubscriptionID string
	include *string
}

// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_authorizations&#x60; and &#x60;tls_authorizations.globalsign_email_challenge&#x60;. 
func (r *APIGetTLSSubRequest) Include(include string) *APIGetTLSSubRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetTLSSubRequest) Execute() (*TLSSubscriptionResponse, *http.Response, error) {
	return r.APIService.GetTLSSubExecute(r)
}

/*
GetTLSSub Get a TLS subscription

Show a TLS subscription.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
 @return APIGetTLSSubRequest
*/
func (a *TLSSubscriptionsAPIService) GetTLSSub(ctx context.Context, tlsSubscriptionID string) APIGetTLSSubRequest {
	return APIGetTLSSubRequest{
		APIService: a,
		ctx: ctx,
		tlsSubscriptionID: tlsSubscriptionID,
	}
}

// GetTLSSubExecute executes the request
//  @return TLSSubscriptionResponse
func (a *TLSSubscriptionsAPIService) GetTLSSubExecute(r APIGetTLSSubRequest) (*TLSSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.GetTLSSub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions/{tls_subscription_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_subscription_id"+"}", gourl.PathEscape(parameterToString(r.tlsSubscriptionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIListTLSSubsRequest represents a request for the resource.
type APIListTLSSubsRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	filterState *string
	filterTLSDomainsID *string
	filterHasActiveOrder *bool
	include *string
	pageNumber *int32
	pageSize *int32
	sort *string
}

// FilterState Limit the returned subscriptions by state. Valid values are &#x60;pending&#x60;, &#x60;processing&#x60;, &#x60;issued&#x60;, &#x60;renewing&#x60;, and &#x60;failed&#x60;. Accepts parameters: &#x60;not&#x60; (e.g., &#x60;filter[state][not]&#x3D;renewing&#x60;). 
func (r *APIListTLSSubsRequest) FilterState(filterState string) *APIListTLSSubsRequest {
	r.filterState = &filterState
	return r
}
// FilterTLSDomainsID Limit the returned subscriptions to those that include the specific domain.
func (r *APIListTLSSubsRequest) FilterTLSDomainsID(filterTLSDomainsID string) *APIListTLSSubsRequest {
	r.filterTLSDomainsID = &filterTLSDomainsID
	return r
}
// FilterHasActiveOrder Limit the returned subscriptions to those that have currently active orders. Permitted values: &#x60;true&#x60;. 
func (r *APIListTLSSubsRequest) FilterHasActiveOrder(filterHasActiveOrder bool) *APIListTLSSubsRequest {
	r.filterHasActiveOrder = &filterHasActiveOrder
	return r
}
// Include Include related objects. Optional, comma-separated values. Permitted values: &#x60;tls_authorizations&#x60; and &#x60;tls_authorizations.globalsign_email_challenge&#x60;. 
func (r *APIListTLSSubsRequest) Include(include string) *APIListTLSSubsRequest {
	r.include = &include
	return r
}
// PageNumber Current page.
func (r *APIListTLSSubsRequest) PageNumber(pageNumber int32) *APIListTLSSubsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListTLSSubsRequest) PageSize(pageSize int32) *APIListTLSSubsRequest {
	r.pageSize = &pageSize
	return r
}
// Sort The order in which to list the results by creation date.
func (r *APIListTLSSubsRequest) Sort(sort string) *APIListTLSSubsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIListTLSSubsRequest) Execute() (*TLSSubscriptionsResponse, *http.Response, error) {
	return r.APIService.ListTLSSubsExecute(r)
}

/*
ListTLSSubs List TLS subscriptions

List all TLS subscriptions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListTLSSubsRequest
*/
func (a *TLSSubscriptionsAPIService) ListTLSSubs(ctx context.Context) APIListTLSSubsRequest {
	return APIListTLSSubsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListTLSSubsExecute executes the request
//  @return TLSSubscriptionsResponse
func (a *TLSSubscriptionsAPIService) ListTLSSubsExecute(r APIListTLSSubsRequest) (*TLSSubscriptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.ListTLSSubs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterState != nil {
		localVarQueryParams.Add("filter[state]", parameterToString(*r.filterState, ""))
	}
	if r.filterTLSDomainsID != nil {
		localVarQueryParams.Add("filter[tls_domains.id]", parameterToString(*r.filterTLSDomainsID, ""))
	}
	if r.filterHasActiveOrder != nil {
		localVarQueryParams.Add("filter[has_active_order]", parameterToString(*r.filterHasActiveOrder, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIPatchTLSSubRequest represents a request for the resource.
type APIPatchTLSSubRequest struct {
	ctx context.Context
	APIService TLSSubscriptionsAPI
	tlsSubscriptionID string
	force *bool
	tlsSubscription *TLSSubscription
}

// Force A flag that allows you to edit and delete a subscription with active domains. Valid to use on PATCH and DELETE actions. As a warning, removing an active domain from a subscription or forcing the deletion of a subscription may result in breaking TLS termination to that domain. 
func (r *APIPatchTLSSubRequest) Force(force bool) *APIPatchTLSSubRequest {
	r.force = &force
	return r
}
// TLSSubscription returns a pointer to a request.
func (r *APIPatchTLSSubRequest) TLSSubscription(tlsSubscription TLSSubscription) *APIPatchTLSSubRequest {
	r.tlsSubscription = &tlsSubscription
	return r
}

// Execute calls the API using the request data configured.
func (r APIPatchTLSSubRequest) Execute() (*TLSSubscriptionResponse, *http.Response, error) {
	return r.APIService.PatchTLSSubExecute(r)
}

/*
PatchTLSSub Update a TLS subscription

Change the TLS domains or common name associated with this subscription, update the TLS configuration for this set of domains, or retry a subscription with state `failed` by setting the state to `retry`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tlsSubscriptionID Alphanumeric string identifying a TLS subscription.
 @return APIPatchTLSSubRequest
*/
func (a *TLSSubscriptionsAPIService) PatchTLSSub(ctx context.Context, tlsSubscriptionID string) APIPatchTLSSubRequest {
	return APIPatchTLSSubRequest{
		APIService: a,
		ctx: ctx,
		tlsSubscriptionID: tlsSubscriptionID,
	}
}

// PatchTLSSubExecute executes the request
//  @return TLSSubscriptionResponse
func (a *TLSSubscriptionsAPIService) PatchTLSSubExecute(r APIPatchTLSSubRequest) (*TLSSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *TLSSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TLSSubscriptionsAPIService.PatchTLSSub")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tls/subscriptions/{tls_subscription_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"tls_subscription_id"+"}", gourl.PathEscape(parameterToString(r.tlsSubscriptionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tlsSubscription
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
