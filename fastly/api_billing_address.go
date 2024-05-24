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

// BillingAddressAPI defines an interface for interacting with the resource.
type BillingAddressAPI interface {

	/*
	AddBillingAddr Add a billing address to a customer

	Add a billing address to a customer.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIAddBillingAddrRequest
	*/
	AddBillingAddr(ctx context.Context, customerID string) APIAddBillingAddrRequest

	// AddBillingAddrExecute executes the request
	//  @return BillingAddressResponse
	AddBillingAddrExecute(r APIAddBillingAddrRequest) (*BillingAddressResponse, *http.Response, error)

	/*
	DeleteBillingAddr Delete a billing address

	Delete a customer's billing address.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIDeleteBillingAddrRequest
	*/
	DeleteBillingAddr(ctx context.Context, customerID string) APIDeleteBillingAddrRequest

	// DeleteBillingAddrExecute executes the request
	DeleteBillingAddrExecute(r APIDeleteBillingAddrRequest) (*http.Response, error)

	/*
	GetBillingAddr Get a billing address

	Get a customer's billing address.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIGetBillingAddrRequest
	*/
	GetBillingAddr(ctx context.Context, customerID string) APIGetBillingAddrRequest

	// GetBillingAddrExecute executes the request
	//  @return BillingAddressResponse
	GetBillingAddrExecute(r APIGetBillingAddrRequest) (*BillingAddressResponse, *http.Response, error)

	/*
	UpdateBillingAddr Update a billing address

	Update a customer's billing address. You may update only part of the customer's billing address.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIUpdateBillingAddrRequest
	*/
	UpdateBillingAddr(ctx context.Context, customerID string) APIUpdateBillingAddrRequest

	// UpdateBillingAddrExecute executes the request
	//  @return BillingAddressResponse
	UpdateBillingAddrExecute(r APIUpdateBillingAddrRequest) (*BillingAddressResponse, *http.Response, error)
}

// BillingAddressAPIService BillingAddressAPI service
type BillingAddressAPIService service

// APIAddBillingAddrRequest represents a request for the resource.
type APIAddBillingAddrRequest struct {
	ctx context.Context
	APIService BillingAddressAPI
	customerID string
	billingAddressRequest *BillingAddressRequest
}

// BillingAddressRequest Billing address
func (r *APIAddBillingAddrRequest) BillingAddressRequest(billingAddressRequest BillingAddressRequest) *APIAddBillingAddrRequest {
	r.billingAddressRequest = &billingAddressRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIAddBillingAddrRequest) Execute() (*BillingAddressResponse, *http.Response, error) {
	return r.APIService.AddBillingAddrExecute(r)
}

/*
AddBillingAddr Add a billing address to a customer

Add a billing address to a customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIAddBillingAddrRequest
*/
func (a *BillingAddressAPIService) AddBillingAddr(ctx context.Context, customerID string) APIAddBillingAddrRequest {
	return APIAddBillingAddrRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// AddBillingAddrExecute executes the request
//  @return BillingAddressResponse
func (a *BillingAddressAPIService) AddBillingAddrExecute(r APIAddBillingAddrRequest) (*BillingAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *BillingAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAddressAPIService.AddBillingAddr")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/billing_address"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.billingAddressRequest
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
			var v BillingAddressVerificationErrorResponse
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

// APIDeleteBillingAddrRequest represents a request for the resource.
type APIDeleteBillingAddrRequest struct {
	ctx context.Context
	APIService BillingAddressAPI
	customerID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteBillingAddrRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteBillingAddrExecute(r)
}

/*
DeleteBillingAddr Delete a billing address

Delete a customer's billing address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIDeleteBillingAddrRequest
*/
func (a *BillingAddressAPIService) DeleteBillingAddr(ctx context.Context, customerID string) APIDeleteBillingAddrRequest {
	return APIDeleteBillingAddrRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// DeleteBillingAddrExecute executes the request
func (a *BillingAddressAPIService) DeleteBillingAddrExecute(r APIDeleteBillingAddrRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAddressAPIService.DeleteBillingAddr")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/billing_address"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

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

// APIGetBillingAddrRequest represents a request for the resource.
type APIGetBillingAddrRequest struct {
	ctx context.Context
	APIService BillingAddressAPI
	customerID string
}


// Execute calls the API using the request data configured.
func (r APIGetBillingAddrRequest) Execute() (*BillingAddressResponse, *http.Response, error) {
	return r.APIService.GetBillingAddrExecute(r)
}

/*
GetBillingAddr Get a billing address

Get a customer's billing address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIGetBillingAddrRequest
*/
func (a *BillingAddressAPIService) GetBillingAddr(ctx context.Context, customerID string) APIGetBillingAddrRequest {
	return APIGetBillingAddrRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// GetBillingAddrExecute executes the request
//  @return BillingAddressResponse
func (a *BillingAddressAPIService) GetBillingAddrExecute(r APIGetBillingAddrRequest) (*BillingAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *BillingAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAddressAPIService.GetBillingAddr")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/billing_address"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

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

// APIUpdateBillingAddrRequest represents a request for the resource.
type APIUpdateBillingAddrRequest struct {
	ctx context.Context
	APIService BillingAddressAPI
	customerID string
	updateBillingAddressRequest *UpdateBillingAddressRequest
}

// UpdateBillingAddressRequest One or more billing address attributes
func (r *APIUpdateBillingAddrRequest) UpdateBillingAddressRequest(updateBillingAddressRequest UpdateBillingAddressRequest) *APIUpdateBillingAddrRequest {
	r.updateBillingAddressRequest = &updateBillingAddressRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateBillingAddrRequest) Execute() (*BillingAddressResponse, *http.Response, error) {
	return r.APIService.UpdateBillingAddrExecute(r)
}

/*
UpdateBillingAddr Update a billing address

Update a customer's billing address. You may update only part of the customer's billing address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIUpdateBillingAddrRequest
*/
func (a *BillingAddressAPIService) UpdateBillingAddr(ctx context.Context, customerID string) APIUpdateBillingAddrRequest {
	return APIUpdateBillingAddrRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// UpdateBillingAddrExecute executes the request
//  @return BillingAddressResponse
func (a *BillingAddressAPIService) UpdateBillingAddrExecute(r APIUpdateBillingAddrRequest) (*BillingAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *BillingAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAddressAPIService.UpdateBillingAddr")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/billing_address"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateBillingAddressRequest
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
			var v BillingAddressVerificationErrorResponse
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
