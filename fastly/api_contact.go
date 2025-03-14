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

// ContactAPI defines an interface for interacting with the resource.
type ContactAPI interface {

	/*
		CreateContacts Add a new customer contact

		Create a contact.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param customerID Alphanumeric string identifying the customer.
		 @return APICreateContactsRequest
	*/
	CreateContacts(ctx context.Context, customerID string) APICreateContactsRequest

	// CreateContactsExecute executes the request
	//  @return ContactResponse
	CreateContactsExecute(r APICreateContactsRequest) (*ContactResponse, *http.Response, error)

	/*
		DeleteContact Delete a contact

		Delete a contact.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param customerID Alphanumeric string identifying the customer.
		 @param contactID An alphanumeric string identifying the customer contact.
		 @return APIDeleteContactRequest
	*/
	DeleteContact(ctx context.Context, customerID string, contactID string) APIDeleteContactRequest

	// DeleteContactExecute executes the request
	//  @return InlineResponse200
	DeleteContactExecute(r APIDeleteContactRequest) (*InlineResponse200, *http.Response, error)

	/*
		ListContacts List contacts

		List all contacts from a specified customer ID.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param customerID Alphanumeric string identifying the customer.
		 @return APIListContactsRequest
	*/
	ListContacts(ctx context.Context, customerID string) APIListContactsRequest

	// ListContactsExecute executes the request
	//  @return []SchemasContactResponse
	ListContactsExecute(r APIListContactsRequest) ([]SchemasContactResponse, *http.Response, error)
}

// ContactAPIService ContactAPI service
type ContactAPIService service

// APICreateContactsRequest represents a request for the resource.
type APICreateContactsRequest struct {
	ctx         context.Context
	APIService  ContactAPI
	customerID  string
	userID      *string
	contactType *string
	name        *string
	email       *string
	phone       *string
	customerID2 *string
}

// UserID The alphanumeric string representing the user for this customer contact.
func (r *APICreateContactsRequest) UserID(userID string) *APICreateContactsRequest {
	r.userID = &userID
	return r
}

// ContactType The type of contact.
func (r *APICreateContactsRequest) ContactType(contactType string) *APICreateContactsRequest {
	r.contactType = &contactType
	return r
}

// Name The name of this contact, when user_id is not provided.
func (r *APICreateContactsRequest) Name(name string) *APICreateContactsRequest {
	r.name = &name
	return r
}

// Email The email of this contact, when a user_id is not provided.
func (r *APICreateContactsRequest) Email(email string) *APICreateContactsRequest {
	r.email = &email
	return r
}

// Phone The phone number for this contact. Required for primary, technical, and security contact types.
func (r *APICreateContactsRequest) Phone(phone string) *APICreateContactsRequest {
	r.phone = &phone
	return r
}

// CustomerID2 The alphanumeric string representing the customer for this customer contact.
func (r *APICreateContactsRequest) CustomerID2(customerID2 string) *APICreateContactsRequest {
	r.customerID2 = &customerID2
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateContactsRequest) Execute() (*ContactResponse, *http.Response, error) {
	return r.APIService.CreateContactsExecute(r)
}

/*
CreateContacts Add a new customer contact

Create a contact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APICreateContactsRequest
*/
func (a *ContactAPIService) CreateContacts(ctx context.Context, customerID string) APICreateContactsRequest {
	return APICreateContactsRequest{
		APIService: a,
		ctx:        ctx,
		customerID: customerID,
	}
}

// CreateContactsExecute executes the request
//  @return ContactResponse
func (a *ContactAPIService) CreateContactsExecute(r APICreateContactsRequest) (*ContactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ContactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactAPIService.CreateContacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/contacts"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

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
	if r.userID != nil {
		localVarFormParams.Add("user_id", parameterToString(*r.userID, ""))
	}
	if r.contactType != nil {
		localVarFormParams.Add("contact_type", parameterToString(*r.contactType, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.email != nil {
		localVarFormParams.Add("email", parameterToString(*r.email, ""))
	}
	if r.phone != nil {
		localVarFormParams.Add("phone", parameterToString(*r.phone, ""))
	}
	if r.customerID2 != nil {
		localVarFormParams.Add("customer_id", parameterToString(*r.customerID2, ""))
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

// APIDeleteContactRequest represents a request for the resource.
type APIDeleteContactRequest struct {
	ctx        context.Context
	APIService ContactAPI
	customerID string
	contactID  string
}

// Execute calls the API using the request data configured.
func (r APIDeleteContactRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteContactExecute(r)
}

/*
DeleteContact Delete a contact

Delete a contact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @param contactID An alphanumeric string identifying the customer contact.
 @return APIDeleteContactRequest
*/
func (a *ContactAPIService) DeleteContact(ctx context.Context, customerID string, contactID string) APIDeleteContactRequest {
	return APIDeleteContactRequest{
		APIService: a,
		ctx:        ctx,
		customerID: customerID,
		contactID:  contactID,
	}
}

// DeleteContactExecute executes the request
//  @return InlineResponse200
func (a *ContactAPIService) DeleteContactExecute(r APIDeleteContactRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactAPIService.DeleteContact")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/contacts/{contact_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"contact_id"+"}", gourl.PathEscape(parameterToString(r.contactID, "")))

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

// APIListContactsRequest represents a request for the resource.
type APIListContactsRequest struct {
	ctx        context.Context
	APIService ContactAPI
	customerID string
}

// Execute calls the API using the request data configured.
func (r APIListContactsRequest) Execute() ([]SchemasContactResponse, *http.Response, error) {
	return r.APIService.ListContactsExecute(r)
}

/*
ListContacts List contacts

List all contacts from a specified customer ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIListContactsRequest
*/
func (a *ContactAPIService) ListContacts(ctx context.Context, customerID string) APIListContactsRequest {
	return APIListContactsRequest{
		APIService: a,
		ctx:        ctx,
		customerID: customerID,
	}
}

// ListContactsExecute executes the request
//  @return []SchemasContactResponse
func (a *ContactAPIService) ListContactsExecute(r APIListContactsRequest) ([]SchemasContactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []SchemasContactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactAPIService.ListContacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/contacts"
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
