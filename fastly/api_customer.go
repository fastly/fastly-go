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

// CustomerAPI defines an interface for interacting with the resource.
type CustomerAPI interface {

	/*
	DeleteCustomer Delete a customer

	Delete a customer.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIDeleteCustomerRequest
	*/
	DeleteCustomer(ctx context.Context, customerID string) APIDeleteCustomerRequest

	// DeleteCustomerExecute executes the request
	//  @return InlineResponse200
	DeleteCustomerExecute(r APIDeleteCustomerRequest) (*InlineResponse200, *http.Response, error)

	/*
	GetCustomer Get a customer

	Get a specific customer.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIGetCustomerRequest
	*/
	GetCustomer(ctx context.Context, customerID string) APIGetCustomerRequest

	// GetCustomerExecute executes the request
	//  @return CustomerResponse
	GetCustomerExecute(r APIGetCustomerRequest) (*CustomerResponse, *http.Response, error)

	/*
	GetLoggedInCustomer Get the logged in customer

	Get the logged in customer.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIGetLoggedInCustomerRequest
	*/
	GetLoggedInCustomer(ctx context.Context) APIGetLoggedInCustomerRequest

	// GetLoggedInCustomerExecute executes the request
	//  @return CustomerResponse
	GetLoggedInCustomerExecute(r APIGetLoggedInCustomerRequest) (*CustomerResponse, *http.Response, error)

	/*
	ListUsers List users

	List all users from a specified customer id.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIListUsersRequest
	*/
	ListUsers(ctx context.Context, customerID string) APIListUsersRequest

	// ListUsersExecute executes the request
	//  @return []SchemasUserResponse
	ListUsersExecute(r APIListUsersRequest) ([]SchemasUserResponse, *http.Response, error)

	/*
	UpdateCustomer Update a customer

	Update a customer.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIUpdateCustomerRequest
	*/
	UpdateCustomer(ctx context.Context, customerID string) APIUpdateCustomerRequest

	// UpdateCustomerExecute executes the request
	//  @return CustomerResponse
	UpdateCustomerExecute(r APIUpdateCustomerRequest) (*CustomerResponse, *http.Response, error)
}

// CustomerAPIService CustomerAPI service
type CustomerAPIService service

// APIDeleteCustomerRequest represents a request for the resource.
type APIDeleteCustomerRequest struct {
	ctx context.Context
	APIService CustomerAPI
	customerID string
}


// Execute calls the API using the request data configured.
func (r APIDeleteCustomerRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteCustomerExecute(r)
}

/*
DeleteCustomer Delete a customer

Delete a customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIDeleteCustomerRequest
*/
func (a *CustomerAPIService) DeleteCustomer(ctx context.Context, customerID string) APIDeleteCustomerRequest {
	return APIDeleteCustomerRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// DeleteCustomerExecute executes the request
//  @return InlineResponse200
func (a *CustomerAPIService) DeleteCustomerExecute(r APIDeleteCustomerRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.DeleteCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}"
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

// APIGetCustomerRequest represents a request for the resource.
type APIGetCustomerRequest struct {
	ctx context.Context
	APIService CustomerAPI
	customerID string
}


// Execute calls the API using the request data configured.
func (r APIGetCustomerRequest) Execute() (*CustomerResponse, *http.Response, error) {
	return r.APIService.GetCustomerExecute(r)
}

/*
GetCustomer Get a customer

Get a specific customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIGetCustomerRequest
*/
func (a *CustomerAPIService) GetCustomer(ctx context.Context, customerID string) APIGetCustomerRequest {
	return APIGetCustomerRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// GetCustomerExecute executes the request
//  @return CustomerResponse
func (a *CustomerAPIService) GetCustomerExecute(r APIGetCustomerRequest) (*CustomerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *CustomerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.GetCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}"
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

// APIGetLoggedInCustomerRequest represents a request for the resource.
type APIGetLoggedInCustomerRequest struct {
	ctx context.Context
	APIService CustomerAPI
}


// Execute calls the API using the request data configured.
func (r APIGetLoggedInCustomerRequest) Execute() (*CustomerResponse, *http.Response, error) {
	return r.APIService.GetLoggedInCustomerExecute(r)
}

/*
GetLoggedInCustomer Get the logged in customer

Get the logged in customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetLoggedInCustomerRequest
*/
func (a *CustomerAPIService) GetLoggedInCustomer(ctx context.Context) APIGetLoggedInCustomerRequest {
	return APIGetLoggedInCustomerRequest{
		APIService: a,
		ctx: ctx,
	}
}

// GetLoggedInCustomerExecute executes the request
//  @return CustomerResponse
func (a *CustomerAPIService) GetLoggedInCustomerExecute(r APIGetLoggedInCustomerRequest) (*CustomerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *CustomerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.GetLoggedInCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/current_customer"

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

// APIListUsersRequest represents a request for the resource.
type APIListUsersRequest struct {
	ctx context.Context
	APIService CustomerAPI
	customerID string
}


// Execute calls the API using the request data configured.
func (r APIListUsersRequest) Execute() ([]SchemasUserResponse, *http.Response, error) {
	return r.APIService.ListUsersExecute(r)
}

/*
ListUsers List users

List all users from a specified customer id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIListUsersRequest
*/
func (a *CustomerAPIService) ListUsers(ctx context.Context, customerID string) APIListUsersRequest {
	return APIListUsersRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// ListUsersExecute executes the request
//  @return []SchemasUserResponse
func (a *CustomerAPIService) ListUsersExecute(r APIListUsersRequest) ([]SchemasUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  []SchemasUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.ListUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}/users"
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

// APIUpdateCustomerRequest represents a request for the resource.
type APIUpdateCustomerRequest struct {
	ctx context.Context
	APIService CustomerAPI
	customerID string
	billingContactID *string
	billingNetworkType *string
	billingRef *string
	canConfigureWordpress *bool
	canResetPasswords *bool
	canUploadVcl *bool
	force2fa *bool
	forceSso *bool
	hasAccountPanel *bool
	hasImprovedEvents *bool
	hasImprovedSslConfig *bool
	hasOpenstackLogging *bool
	hasPci *bool
	hasPciPasswords *bool
	ipWhitelist *string
	legalContactID *string
	name *string
	ownerID *string
	phoneNumber *string
	postalAddress *string
	pricingPlan *string
	pricingPlanID *string
	securityContactID *string
	technicalContactID *string
}

// BillingContactID The alphanumeric string representing the primary billing contact.
func (r *APIUpdateCustomerRequest) BillingContactID(billingContactID string) *APIUpdateCustomerRequest {
	r.billingContactID = &billingContactID
	return r
}
// BillingNetworkType Customer&#39;s current network revenue type.
func (r *APIUpdateCustomerRequest) BillingNetworkType(billingNetworkType string) *APIUpdateCustomerRequest {
	r.billingNetworkType = &billingNetworkType
	return r
}
// BillingRef Used for adding purchased orders to customer&#39;s account.
func (r *APIUpdateCustomerRequest) BillingRef(billingRef string) *APIUpdateCustomerRequest {
	r.billingRef = &billingRef
	return r
}
// CanConfigureWordpress Whether this customer can view or edit wordpress.
func (r *APIUpdateCustomerRequest) CanConfigureWordpress(canConfigureWordpress bool) *APIUpdateCustomerRequest {
	r.canConfigureWordpress = &canConfigureWordpress
	return r
}
// CanResetPasswords Whether this customer can reset passwords.
func (r *APIUpdateCustomerRequest) CanResetPasswords(canResetPasswords bool) *APIUpdateCustomerRequest {
	r.canResetPasswords = &canResetPasswords
	return r
}
// CanUploadVcl Whether this customer can upload VCL.
func (r *APIUpdateCustomerRequest) CanUploadVcl(canUploadVcl bool) *APIUpdateCustomerRequest {
	r.canUploadVcl = &canUploadVcl
	return r
}
// Force2fa Specifies whether 2FA is forced or not forced on the customer account. Logs out non-2FA users once 2FA is force enabled.
func (r *APIUpdateCustomerRequest) Force2fa(force2fa bool) *APIUpdateCustomerRequest {
	r.force2fa = &force2fa
	return r
}
// ForceSso Specifies whether SSO is forced or not forced on the customer account.
func (r *APIUpdateCustomerRequest) ForceSso(forceSso bool) *APIUpdateCustomerRequest {
	r.forceSso = &forceSso
	return r
}
// HasAccountPanel Specifies whether the account has access or does not have access to the account panel.
func (r *APIUpdateCustomerRequest) HasAccountPanel(hasAccountPanel bool) *APIUpdateCustomerRequest {
	r.hasAccountPanel = &hasAccountPanel
	return r
}
// HasImprovedEvents Specifies whether the account has access or does not have access to the improved events.
func (r *APIUpdateCustomerRequest) HasImprovedEvents(hasImprovedEvents bool) *APIUpdateCustomerRequest {
	r.hasImprovedEvents = &hasImprovedEvents
	return r
}
// HasImprovedSslConfig Whether this customer can view or edit the SSL config.
func (r *APIUpdateCustomerRequest) HasImprovedSslConfig(hasImprovedSslConfig bool) *APIUpdateCustomerRequest {
	r.hasImprovedSslConfig = &hasImprovedSslConfig
	return r
}
// HasOpenstackLogging Specifies whether the account has enabled or not enabled openstack logging.
func (r *APIUpdateCustomerRequest) HasOpenstackLogging(hasOpenstackLogging bool) *APIUpdateCustomerRequest {
	r.hasOpenstackLogging = &hasOpenstackLogging
	return r
}
// HasPci Specifies whether the account can edit PCI for a service.
func (r *APIUpdateCustomerRequest) HasPci(hasPci bool) *APIUpdateCustomerRequest {
	r.hasPci = &hasPci
	return r
}
// HasPciPasswords Specifies whether PCI passwords are required for the account.
func (r *APIUpdateCustomerRequest) HasPciPasswords(hasPciPasswords bool) *APIUpdateCustomerRequest {
	r.hasPciPasswords = &hasPciPasswords
	return r
}
// IPWhitelist The range of IP addresses authorized to access the customer account.
func (r *APIUpdateCustomerRequest) IPWhitelist(ipWhitelist string) *APIUpdateCustomerRequest {
	r.ipWhitelist = &ipWhitelist
	return r
}
// LegalContactID The alphanumeric string identifying the account&#39;s legal contact.
func (r *APIUpdateCustomerRequest) LegalContactID(legalContactID string) *APIUpdateCustomerRequest {
	r.legalContactID = &legalContactID
	return r
}
// Name The name of the customer, generally the company name.
func (r *APIUpdateCustomerRequest) Name(name string) *APIUpdateCustomerRequest {
	r.name = &name
	return r
}
// OwnerID The alphanumeric string identifying the account owner.
func (r *APIUpdateCustomerRequest) OwnerID(ownerID string) *APIUpdateCustomerRequest {
	r.ownerID = &ownerID
	return r
}
// PhoneNumber The phone number associated with the account.
func (r *APIUpdateCustomerRequest) PhoneNumber(phoneNumber string) *APIUpdateCustomerRequest {
	r.phoneNumber = &phoneNumber
	return r
}
// PostalAddress The postal address associated with the account.
func (r *APIUpdateCustomerRequest) PostalAddress(postalAddress string) *APIUpdateCustomerRequest {
	r.postalAddress = &postalAddress
	return r
}
// PricingPlan The pricing plan this customer is under.
func (r *APIUpdateCustomerRequest) PricingPlan(pricingPlan string) *APIUpdateCustomerRequest {
	r.pricingPlan = &pricingPlan
	return r
}
// PricingPlanID The alphanumeric string identifying the pricing plan.
func (r *APIUpdateCustomerRequest) PricingPlanID(pricingPlanID string) *APIUpdateCustomerRequest {
	r.pricingPlanID = &pricingPlanID
	return r
}
// SecurityContactID The alphanumeric string identifying the account&#39;s security contact.
func (r *APIUpdateCustomerRequest) SecurityContactID(securityContactID string) *APIUpdateCustomerRequest {
	r.securityContactID = &securityContactID
	return r
}
// TechnicalContactID The alphanumeric string identifying the account&#39;s technical contact.
func (r *APIUpdateCustomerRequest) TechnicalContactID(technicalContactID string) *APIUpdateCustomerRequest {
	r.technicalContactID = &technicalContactID
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateCustomerRequest) Execute() (*CustomerResponse, *http.Response, error) {
	return r.APIService.UpdateCustomerExecute(r)
}

/*
UpdateCustomer Update a customer

Update a customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIUpdateCustomerRequest
*/
func (a *CustomerAPIService) UpdateCustomer(ctx context.Context, customerID string) APIUpdateCustomerRequest {
	return APIUpdateCustomerRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// UpdateCustomerExecute executes the request
//  @return CustomerResponse
func (a *CustomerAPIService) UpdateCustomerExecute(r APIUpdateCustomerRequest) (*CustomerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *CustomerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.UpdateCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customer_id}"
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
	if r.billingContactID != nil {
		localVarFormParams.Add("billing_contact_id", parameterToString(*r.billingContactID, ""))
	}
	if r.billingNetworkType != nil {
		localVarFormParams.Add("billing_network_type", parameterToString(*r.billingNetworkType, ""))
	}
	if r.billingRef != nil {
		localVarFormParams.Add("billing_ref", parameterToString(*r.billingRef, ""))
	}
	if r.canConfigureWordpress != nil {
		localVarFormParams.Add("can_configure_wordpress", parameterToString(*r.canConfigureWordpress, ""))
	}
	if r.canResetPasswords != nil {
		localVarFormParams.Add("can_reset_passwords", parameterToString(*r.canResetPasswords, ""))
	}
	if r.canUploadVcl != nil {
		localVarFormParams.Add("can_upload_vcl", parameterToString(*r.canUploadVcl, ""))
	}
	if r.force2fa != nil {
		localVarFormParams.Add("force_2fa", parameterToString(*r.force2fa, ""))
	}
	if r.forceSso != nil {
		localVarFormParams.Add("force_sso", parameterToString(*r.forceSso, ""))
	}
	if r.hasAccountPanel != nil {
		localVarFormParams.Add("has_account_panel", parameterToString(*r.hasAccountPanel, ""))
	}
	if r.hasImprovedEvents != nil {
		localVarFormParams.Add("has_improved_events", parameterToString(*r.hasImprovedEvents, ""))
	}
	if r.hasImprovedSslConfig != nil {
		localVarFormParams.Add("has_improved_ssl_config", parameterToString(*r.hasImprovedSslConfig, ""))
	}
	if r.hasOpenstackLogging != nil {
		localVarFormParams.Add("has_openstack_logging", parameterToString(*r.hasOpenstackLogging, ""))
	}
	if r.hasPci != nil {
		localVarFormParams.Add("has_pci", parameterToString(*r.hasPci, ""))
	}
	if r.hasPciPasswords != nil {
		localVarFormParams.Add("has_pci_passwords", parameterToString(*r.hasPciPasswords, ""))
	}
	if r.ipWhitelist != nil {
		localVarFormParams.Add("ip_whitelist", parameterToString(*r.ipWhitelist, ""))
	}
	if r.legalContactID != nil {
		localVarFormParams.Add("legal_contact_id", parameterToString(*r.legalContactID, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.ownerID != nil {
		localVarFormParams.Add("owner_id", parameterToString(*r.ownerID, ""))
	}
	if r.phoneNumber != nil {
		localVarFormParams.Add("phone_number", parameterToString(*r.phoneNumber, ""))
	}
	if r.postalAddress != nil {
		localVarFormParams.Add("postal_address", parameterToString(*r.postalAddress, ""))
	}
	if r.pricingPlan != nil {
		localVarFormParams.Add("pricing_plan", parameterToString(*r.pricingPlan, ""))
	}
	if r.pricingPlanID != nil {
		localVarFormParams.Add("pricing_plan_id", parameterToString(*r.pricingPlanID, ""))
	}
	if r.securityContactID != nil {
		localVarFormParams.Add("security_contact_id", parameterToString(*r.securityContactID, ""))
	}
	if r.technicalContactID != nil {
		localVarFormParams.Add("technical_contact_id", parameterToString(*r.technicalContactID, ""))
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
