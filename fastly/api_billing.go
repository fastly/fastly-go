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

// BillingAPI defines an interface for interacting with the resource.
type BillingAPI interface {

	/*
	GetInvoice Get an invoice

	Get the invoice for a given year and month. Can be any month from when the Customer was created to the current month.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param month 2-digit month.
	 @param year 4-digit year.
	 @return APIGetInvoiceRequest
	*/
	GetInvoice(ctx context.Context, month string, year string) APIGetInvoiceRequest

	// GetInvoiceExecute executes the request
	//  @return BillingResponse
	GetInvoiceExecute(r APIGetInvoiceRequest) (*BillingResponse, *http.Response, error)

	/*
	GetInvoiceByID Get an invoice

	Get the invoice for the given invoice_id.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @param invoiceID Alphanumeric string identifying the invoice.
	 @return APIGetInvoiceByIDRequest
	*/
	GetInvoiceByID(ctx context.Context, customerID string, invoiceID string) APIGetInvoiceByIDRequest

	// GetInvoiceByIDExecute executes the request
	//  @return BillingResponse
	GetInvoiceByIDExecute(r APIGetInvoiceByIDRequest) (*BillingResponse, *http.Response, error)

	/*
	GetInvoiceMtd Get month-to-date billing estimate

	Get the current month-to-date estimate. This endpoint has two different responses. Under normal circumstances, it generally takes less than 5 seconds to generate but in certain cases can take up to 60 seconds. Once generated the month-to-date estimate is cached for 4 hours, and is available the next request will return the JSON representation of the month-to-date estimate. While a report is being generated in the background, this endpoint will return a `202 Accepted` response. The full format of which can be found in detail in our [billing calculation guide](https://docs.fastly.com/en/guides/how-we-calculate-your-bill). There are certain accounts for which we are unable to generate a month-to-date estimate. For example, accounts who have parent-pay are unable to generate an MTD estimate. The parent accounts are able to generate a month-to-date estimate but that estimate will not include the child accounts amounts at this time.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param customerID Alphanumeric string identifying the customer.
	 @return APIGetInvoiceMtdRequest
	*/
	GetInvoiceMtd(ctx context.Context, customerID string) APIGetInvoiceMtdRequest

	// GetInvoiceMtdExecute executes the request
	//  @return BillingEstimateResponse
	GetInvoiceMtdExecute(r APIGetInvoiceMtdRequest) (*BillingEstimateResponse, *http.Response, error)
}

// BillingAPIService BillingAPI service
type BillingAPIService service

// APIGetInvoiceRequest represents a request for the resource.
type APIGetInvoiceRequest struct {
	ctx context.Context
	APIService BillingAPI
	month string
	year string
}


// Execute calls the API using the request data configured.
func (r APIGetInvoiceRequest) Execute() (*BillingResponse, *http.Response, error) {
	return r.APIService.GetInvoiceExecute(r)
}

/*
GetInvoice Get an invoice

Get the invoice for a given year and month. Can be any month from when the Customer was created to the current month.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param month 2-digit month.
 @param year 4-digit year.
 @return APIGetInvoiceRequest
*/
func (a *BillingAPIService) GetInvoice(ctx context.Context, month string, year string) APIGetInvoiceRequest {
	return APIGetInvoiceRequest{
		APIService: a,
		ctx: ctx,
		month: month,
		year: year,
	}
}

// GetInvoiceExecute executes the request
//  @return BillingResponse
func (a *BillingAPIService) GetInvoiceExecute(r APIGetInvoiceRequest) (*BillingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *BillingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAPIService.GetInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing/v2/year/{year}/month/{month}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"month"+"}", gourl.PathEscape(parameterToString(r.month, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"year"+"}", gourl.PathEscape(parameterToString(r.year, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv", "application/pdf"}

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

// APIGetInvoiceByIDRequest represents a request for the resource.
type APIGetInvoiceByIDRequest struct {
	ctx context.Context
	APIService BillingAPI
	customerID string
	invoiceID string
}


// Execute calls the API using the request data configured.
func (r APIGetInvoiceByIDRequest) Execute() (*BillingResponse, *http.Response, error) {
	return r.APIService.GetInvoiceByIDExecute(r)
}

/*
GetInvoiceByID Get an invoice

Get the invoice for the given invoice_id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @param invoiceID Alphanumeric string identifying the invoice.
 @return APIGetInvoiceByIDRequest
*/
func (a *BillingAPIService) GetInvoiceByID(ctx context.Context, customerID string, invoiceID string) APIGetInvoiceByIDRequest {
	return APIGetInvoiceByIDRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
		invoiceID: invoiceID,
	}
}

// GetInvoiceByIDExecute executes the request
//  @return BillingResponse
func (a *BillingAPIService) GetInvoiceByIDExecute(r APIGetInvoiceByIDRequest) (*BillingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *BillingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAPIService.GetInvoiceByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing/v2/account_customers/{customer_id}/invoices/{invoice_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"invoice_id"+"}", gourl.PathEscape(parameterToString(r.invoiceID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv", "application/pdf"}

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

// APIGetInvoiceMtdRequest represents a request for the resource.
type APIGetInvoiceMtdRequest struct {
	ctx context.Context
	APIService BillingAPI
	customerID string
	month *string
	year *string
}

// Month 2-digit month.
func (r *APIGetInvoiceMtdRequest) Month(month string) *APIGetInvoiceMtdRequest {
	r.month = &month
	return r
}
// Year 4-digit year.
func (r *APIGetInvoiceMtdRequest) Year(year string) *APIGetInvoiceMtdRequest {
	r.year = &year
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetInvoiceMtdRequest) Execute() (*BillingEstimateResponse, *http.Response, error) {
	return r.APIService.GetInvoiceMtdExecute(r)
}

/*
GetInvoiceMtd Get month-to-date billing estimate

Get the current month-to-date estimate. This endpoint has two different responses. Under normal circumstances, it generally takes less than 5 seconds to generate but in certain cases can take up to 60 seconds. Once generated the month-to-date estimate is cached for 4 hours, and is available the next request will return the JSON representation of the month-to-date estimate. While a report is being generated in the background, this endpoint will return a `202 Accepted` response. The full format of which can be found in detail in our [billing calculation guide](https://docs.fastly.com/en/guides/how-we-calculate-your-bill). There are certain accounts for which we are unable to generate a month-to-date estimate. For example, accounts who have parent-pay are unable to generate an MTD estimate. The parent accounts are able to generate a month-to-date estimate but that estimate will not include the child accounts amounts at this time.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIGetInvoiceMtdRequest
*/
func (a *BillingAPIService) GetInvoiceMtd(ctx context.Context, customerID string) APIGetInvoiceMtdRequest {
	return APIGetInvoiceMtdRequest{
		APIService: a,
		ctx: ctx,
		customerID: customerID,
	}
}

// GetInvoiceMtdExecute executes the request
//  @return BillingEstimateResponse
func (a *BillingAPIService) GetInvoiceMtdExecute(r APIGetInvoiceMtdRequest) (*BillingEstimateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *BillingEstimateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingAPIService.GetInvoiceMtd")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing/v2/account_customers/{customer_id}/mtd_invoice"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.month != nil {
		localVarQueryParams.Add("month", parameterToString(*r.month, ""))
	}
	if r.year != nil {
		localVarQueryParams.Add("year", parameterToString(*r.year, ""))
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
