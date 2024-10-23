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

// BillingUsageMetricsAPI defines an interface for interacting with the resource.
type BillingUsageMetricsAPI interface {

	/*
		GetServiceLevelUsage Retrieve service-level usage metrics for a product.

		Returns product usage, broken down by service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param customerID Alphanumeric string identifying the customer.
		 @return APIGetServiceLevelUsageRequest
	*/
	GetServiceLevelUsage(ctx context.Context, customerID string) APIGetServiceLevelUsageRequest

	// GetServiceLevelUsageExecute executes the request
	//  @return Serviceusagemetrics
	GetServiceLevelUsageExecute(r APIGetServiceLevelUsageRequest) (*Serviceusagemetrics, *http.Response, error)

	/*
		GetServiceLevelUsageTypes Retrieve product usage types for a customer.

		Returns product usage types reported by the customer's services.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param customerID Alphanumeric string identifying the customer.
		 @return APIGetServiceLevelUsageTypesRequest
	*/
	GetServiceLevelUsageTypes(ctx context.Context, customerID string) APIGetServiceLevelUsageTypesRequest

	// GetServiceLevelUsageTypesExecute executes the request
	//  @return Serviceusagetypes
	GetServiceLevelUsageTypesExecute(r APIGetServiceLevelUsageTypesRequest) (*Serviceusagetypes, *http.Response, error)
}

// BillingUsageMetricsAPIService BillingUsageMetricsAPI service
type BillingUsageMetricsAPIService service

// APIGetServiceLevelUsageRequest represents a request for the resource.
type APIGetServiceLevelUsageRequest struct {
	ctx             context.Context
	APIService      BillingUsageMetricsAPI
	customerID      string
	productID       *string
	usageTypeName   *string
	timeGranularity *string
	startDate       *string
	endDate         *string
	startMonth      *string
	endMonth        *string
	limit           *string
	cursor          *string
}

// ProductID The product identifier for the metrics returned (e.g., &#x60;cdn_usage&#x60;). This field is not required for CSV requests.
func (r *APIGetServiceLevelUsageRequest) ProductID(productID string) *APIGetServiceLevelUsageRequest {
	r.productID = &productID
	return r
}

// UsageTypeName The usage type name for the metrics returned (e.g., &#x60;North America Requests&#x60;). This field is not required for CSV requests.
func (r *APIGetServiceLevelUsageRequest) UsageTypeName(usageTypeName string) *APIGetServiceLevelUsageRequest {
	r.usageTypeName = &usageTypeName
	return r
}

// TimeGranularity returns a pointer to a request.
func (r *APIGetServiceLevelUsageRequest) TimeGranularity(timeGranularity string) *APIGetServiceLevelUsageRequest {
	r.timeGranularity = &timeGranularity
	return r
}

// StartDate returns a pointer to a request.
func (r *APIGetServiceLevelUsageRequest) StartDate(startDate string) *APIGetServiceLevelUsageRequest {
	r.startDate = &startDate
	return r
}

// EndDate returns a pointer to a request.
func (r *APIGetServiceLevelUsageRequest) EndDate(endDate string) *APIGetServiceLevelUsageRequest {
	r.endDate = &endDate
	return r
}

// StartMonth returns a pointer to a request.
func (r *APIGetServiceLevelUsageRequest) StartMonth(startMonth string) *APIGetServiceLevelUsageRequest {
	r.startMonth = &startMonth
	return r
}

// EndMonth returns a pointer to a request.
func (r *APIGetServiceLevelUsageRequest) EndMonth(endMonth string) *APIGetServiceLevelUsageRequest {
	r.endMonth = &endMonth
	return r
}

// Limit Number of results per page. The maximum is 100.
func (r *APIGetServiceLevelUsageRequest) Limit(limit string) *APIGetServiceLevelUsageRequest {
	r.limit = &limit
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIGetServiceLevelUsageRequest) Cursor(cursor string) *APIGetServiceLevelUsageRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetServiceLevelUsageRequest) Execute() (*Serviceusagemetrics, *http.Response, error) {
	return r.APIService.GetServiceLevelUsageExecute(r)
}

/*
GetServiceLevelUsage Retrieve service-level usage metrics for a product.

Returns product usage, broken down by service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIGetServiceLevelUsageRequest
*/
func (a *BillingUsageMetricsAPIService) GetServiceLevelUsage(ctx context.Context, customerID string) APIGetServiceLevelUsageRequest {
	return APIGetServiceLevelUsageRequest{
		APIService: a,
		ctx:        ctx,
		customerID: customerID,
	}
}

// GetServiceLevelUsageExecute executes the request
//  @return Serviceusagemetrics
func (a *BillingUsageMetricsAPIService) GetServiceLevelUsageExecute(r APIGetServiceLevelUsageRequest) (*Serviceusagemetrics, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Serviceusagemetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingUsageMetricsAPIService.GetServiceLevelUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing/v2/account_customers/{customer_id}/service-usage-metrics"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"customer_id"+"}", gourl.PathEscape(parameterToString(r.customerID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.productID == nil {
		return localVarReturnValue, nil, reportError("productID is required and must be specified")
	}
	if r.usageTypeName == nil {
		return localVarReturnValue, nil, reportError("usageTypeName is required and must be specified")
	}
	if r.timeGranularity == nil {
		return localVarReturnValue, nil, reportError("timeGranularity is required and must be specified")
	}

	localVarQueryParams.Add("product_id", parameterToString(*r.productID, ""))
	localVarQueryParams.Add("usage_type_name", parameterToString(*r.usageTypeName, ""))
	localVarQueryParams.Add("time_granularity", parameterToString(*r.timeGranularity, ""))
	if r.startDate != nil {
		localVarQueryParams.Add("start_date", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("end_date", parameterToString(*r.endDate, ""))
	}
	if r.startMonth != nil {
		localVarQueryParams.Add("start_month", parameterToString(*r.startMonth, ""))
	}
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", parameterToString(*r.endMonth, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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

// APIGetServiceLevelUsageTypesRequest represents a request for the resource.
type APIGetServiceLevelUsageTypesRequest struct {
	ctx        context.Context
	APIService BillingUsageMetricsAPI
	customerID string
}

// Execute calls the API using the request data configured.
func (r APIGetServiceLevelUsageTypesRequest) Execute() (*Serviceusagetypes, *http.Response, error) {
	return r.APIService.GetServiceLevelUsageTypesExecute(r)
}

/*
GetServiceLevelUsageTypes Retrieve product usage types for a customer.

Returns product usage types reported by the customer's services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerID Alphanumeric string identifying the customer.
 @return APIGetServiceLevelUsageTypesRequest
*/
func (a *BillingUsageMetricsAPIService) GetServiceLevelUsageTypes(ctx context.Context, customerID string) APIGetServiceLevelUsageTypesRequest {
	return APIGetServiceLevelUsageTypesRequest{
		APIService: a,
		ctx:        ctx,
		customerID: customerID,
	}
}

// GetServiceLevelUsageTypesExecute executes the request
//  @return Serviceusagetypes
func (a *BillingUsageMetricsAPIService) GetServiceLevelUsageTypesExecute(r APIGetServiceLevelUsageTypesRequest) (*Serviceusagetypes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Serviceusagetypes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingUsageMetricsAPIService.GetServiceLevelUsageTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing/v2/account_customers/{customer_id}/service-usage-types"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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
