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

// HistoricalAPI defines an interface for interacting with the resource.
type HistoricalAPI interface {

	/*
		GetHistStats Get historical stats

		Fetches historical stats for each of your Fastly services and groups the results by service ID.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetHistStatsRequest
	*/
	GetHistStats(ctx context.Context) APIGetHistStatsRequest

	// GetHistStatsExecute executes the request
	//  @return HistoricalStatsByServiceResponse
	GetHistStatsExecute(r APIGetHistStatsRequest) (*HistoricalStatsByServiceResponse, *http.Response, error)

	/*
		GetHistStatsAggregated Get aggregated historical stats

		Fetches historical stats information aggregated across all of your Fastly services.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetHistStatsAggregatedRequest
	*/
	GetHistStatsAggregated(ctx context.Context) APIGetHistStatsAggregatedRequest

	// GetHistStatsAggregatedExecute executes the request
	//  @return HistoricalStatsAggregatedResponse
	GetHistStatsAggregatedExecute(r APIGetHistStatsAggregatedRequest) (*HistoricalStatsAggregatedResponse, *http.Response, error)

	/*
		GetHistStatsField Get historical stats for a single field

		Fetches the specified field from the historical stats for each of your services and groups the results by service ID.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param field Name of the stats field.
		 @return APIGetHistStatsFieldRequest
	*/
	GetHistStatsField(ctx context.Context, field string) APIGetHistStatsFieldRequest

	// GetHistStatsFieldExecute executes the request
	//  @return HistoricalStatsByServiceResponse
	GetHistStatsFieldExecute(r APIGetHistStatsFieldRequest) (*HistoricalStatsByServiceResponse, *http.Response, error)

	/*
		GetHistStatsService Get historical stats for a single service

		Fetches historical stats for a given service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetHistStatsServiceRequest
	*/
	GetHistStatsService(ctx context.Context, serviceID string) APIGetHistStatsServiceRequest

	// GetHistStatsServiceExecute executes the request
	//  @return HistoricalStatsAggregatedResponse
	GetHistStatsServiceExecute(r APIGetHistStatsServiceRequest) (*HistoricalStatsAggregatedResponse, *http.Response, error)

	/*
		GetHistStatsServiceField Get historical stats for a single service/field combination

		Fetches the specified field from the historical stats for a given service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param field Name of the stats field.
		 @return APIGetHistStatsServiceFieldRequest
	*/
	GetHistStatsServiceField(ctx context.Context, serviceID string, field string) APIGetHistStatsServiceFieldRequest

	// GetHistStatsServiceFieldExecute executes the request
	//  @return HistoricalStatsAggregatedResponse
	GetHistStatsServiceFieldExecute(r APIGetHistStatsServiceFieldRequest) (*HistoricalStatsAggregatedResponse, *http.Response, error)

	/*
		GetRegions Get region codes

		Fetches the list of codes for regions that are covered by the Fastly CDN service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetRegionsRequest
	*/
	GetRegions(ctx context.Context) APIGetRegionsRequest

	// GetRegionsExecute executes the request
	//  @return HistoricalRegionsResponse
	GetRegionsExecute(r APIGetRegionsRequest) (*HistoricalRegionsResponse, *http.Response, error)

	/*
		GetUsage Get usage statistics

		Returns usage information aggregated across all Fastly services and grouped by region. To aggregate across all Fastly services by time period, see [`/stats/aggregate`](#get-hist-stats-aggregated).

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetUsageRequest
	*/
	GetUsage(ctx context.Context) APIGetUsageRequest

	// GetUsageExecute executes the request
	//  @return HistoricalUsageAggregatedResponse
	GetUsageExecute(r APIGetUsageRequest) (*HistoricalUsageAggregatedResponse, *http.Response, error)

	/*
		GetUsageMonth Get month-to-date usage statistics

		Returns month-to-date usage details for a given month and year. Usage details are aggregated by service and across all Fastly services, and then grouped by region. This endpoint does not use the `from` or `to` fields for selecting the date for which data is requested. Instead, it uses `month` and `year` integer fields. Both fields are optional and default to the current month and year respectively. When set, an optional `billable_units` field will convert bandwidth to GB and divide requests by 10,000.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetUsageMonthRequest
	*/
	GetUsageMonth(ctx context.Context) APIGetUsageMonthRequest

	// GetUsageMonthExecute executes the request
	//  @return HistoricalUsageMonthResponse
	GetUsageMonthExecute(r APIGetUsageMonthRequest) (*HistoricalUsageMonthResponse, *http.Response, error)

	/*
		GetUsageService Get usage statistics per service

		Returns usage information aggregated by service and grouped by service and region. For service stats by time period, see [`/stats`](#get-hist-stats) and [`/stats/field/:field`](#get-hist-stats-field).

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetUsageServiceRequest
	*/
	GetUsageService(ctx context.Context) APIGetUsageServiceRequest

	// GetUsageServiceExecute executes the request
	//  @return HistoricalUsageServiceResponse
	GetUsageServiceExecute(r APIGetUsageServiceRequest) (*HistoricalUsageServiceResponse, *http.Response, error)
}

// HistoricalAPIService HistoricalAPI service
type HistoricalAPIService service

// APIGetHistStatsRequest represents a request for the resource.
type APIGetHistStatsRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	from       *string
	to         *string
	by         *string
	region     *string
	services   *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetHistStatsRequest) From(from string) *APIGetHistStatsRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetHistStatsRequest) To(to string) *APIGetHistStatsRequest {
	r.to = &to
	return r
}

// By Duration of sample windows. One of:   * &#x60;hour&#x60; - Group data by hour.   * &#x60;minute&#x60; - Group data by minute.   * &#x60;day&#x60; - Group data by day.
func (r *APIGetHistStatsRequest) By(by string) *APIGetHistStatsRequest {
	r.by = &by
	return r
}

// Region Limit query to a specific geographic region. One of:   * &#x60;usa&#x60; - North America.   * &#x60;europe&#x60; - Europe.   * &#x60;anzac&#x60; - Australia and New Zealand.   * &#x60;asia&#x60; - Asia.   * &#x60;asia_india&#x60; - India.   * &#x60;asia_southkorea&#x60; - South Korea.   * &#x60;africa_std&#x60; - Africa.   * &#x60;mexico&#x60; - Mexico.   * &#x60;southamerica_std&#x60; - South America.
func (r *APIGetHistStatsRequest) Region(region string) *APIGetHistStatsRequest {
	r.region = &region
	return r
}

// Services Limit the query to only the specified, comma-separated list of services.
func (r *APIGetHistStatsRequest) Services(services string) *APIGetHistStatsRequest {
	r.services = &services
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetHistStatsRequest) Execute() (*HistoricalStatsByServiceResponse, *http.Response, error) {
	return r.APIService.GetHistStatsExecute(r)
}

/*
GetHistStats Get historical stats

Fetches historical stats for each of your Fastly services and groups the results by service ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetHistStatsRequest
*/
func (a *HistoricalAPIService) GetHistStats(ctx context.Context) APIGetHistStatsRequest {
	return APIGetHistStatsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetHistStatsExecute executes the request
//  @return HistoricalStatsByServiceResponse
func (a *HistoricalAPIService) GetHistStatsExecute(r APIGetHistStatsRequest) (*HistoricalStatsByServiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalStatsByServiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetHistStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.by != nil {
		localVarQueryParams.Add("by", parameterToString(*r.by, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.services != nil {
		localVarQueryParams.Add("services", parameterToString(*r.services, ""))
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

// APIGetHistStatsAggregatedRequest represents a request for the resource.
type APIGetHistStatsAggregatedRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	from       *string
	to         *string
	by         *string
	region     *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetHistStatsAggregatedRequest) From(from string) *APIGetHistStatsAggregatedRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetHistStatsAggregatedRequest) To(to string) *APIGetHistStatsAggregatedRequest {
	r.to = &to
	return r
}

// By Duration of sample windows. One of:   * &#x60;hour&#x60; - Group data by hour.   * &#x60;minute&#x60; - Group data by minute.   * &#x60;day&#x60; - Group data by day.
func (r *APIGetHistStatsAggregatedRequest) By(by string) *APIGetHistStatsAggregatedRequest {
	r.by = &by
	return r
}

// Region Limit query to a specific geographic region. One of:   * &#x60;usa&#x60; - North America.   * &#x60;europe&#x60; - Europe.   * &#x60;anzac&#x60; - Australia and New Zealand.   * &#x60;asia&#x60; - Asia.   * &#x60;asia_india&#x60; - India.   * &#x60;asia_southkorea&#x60; - South Korea.   * &#x60;africa_std&#x60; - Africa.   * &#x60;mexico&#x60; - Mexico.   * &#x60;southamerica_std&#x60; - South America.
func (r *APIGetHistStatsAggregatedRequest) Region(region string) *APIGetHistStatsAggregatedRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetHistStatsAggregatedRequest) Execute() (*HistoricalStatsAggregatedResponse, *http.Response, error) {
	return r.APIService.GetHistStatsAggregatedExecute(r)
}

/*
GetHistStatsAggregated Get aggregated historical stats

Fetches historical stats information aggregated across all of your Fastly services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetHistStatsAggregatedRequest
*/
func (a *HistoricalAPIService) GetHistStatsAggregated(ctx context.Context) APIGetHistStatsAggregatedRequest {
	return APIGetHistStatsAggregatedRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetHistStatsAggregatedExecute executes the request
//  @return HistoricalStatsAggregatedResponse
func (a *HistoricalAPIService) GetHistStatsAggregatedExecute(r APIGetHistStatsAggregatedRequest) (*HistoricalStatsAggregatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalStatsAggregatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetHistStatsAggregated")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.by != nil {
		localVarQueryParams.Add("by", parameterToString(*r.by, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
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

// APIGetHistStatsFieldRequest represents a request for the resource.
type APIGetHistStatsFieldRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	field      string
	from       *string
	to         *string
	by         *string
	region     *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetHistStatsFieldRequest) From(from string) *APIGetHistStatsFieldRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetHistStatsFieldRequest) To(to string) *APIGetHistStatsFieldRequest {
	r.to = &to
	return r
}

// By Duration of sample windows. One of:   * &#x60;hour&#x60; - Group data by hour.   * &#x60;minute&#x60; - Group data by minute.   * &#x60;day&#x60; - Group data by day.
func (r *APIGetHistStatsFieldRequest) By(by string) *APIGetHistStatsFieldRequest {
	r.by = &by
	return r
}

// Region Limit query to a specific geographic region. One of:   * &#x60;usa&#x60; - North America.   * &#x60;europe&#x60; - Europe.   * &#x60;anzac&#x60; - Australia and New Zealand.   * &#x60;asia&#x60; - Asia.   * &#x60;asia_india&#x60; - India.   * &#x60;asia_southkorea&#x60; - South Korea.   * &#x60;africa_std&#x60; - Africa.   * &#x60;mexico&#x60; - Mexico.   * &#x60;southamerica_std&#x60; - South America.
func (r *APIGetHistStatsFieldRequest) Region(region string) *APIGetHistStatsFieldRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetHistStatsFieldRequest) Execute() (*HistoricalStatsByServiceResponse, *http.Response, error) {
	return r.APIService.GetHistStatsFieldExecute(r)
}

/*
GetHistStatsField Get historical stats for a single field

Fetches the specified field from the historical stats for each of your services and groups the results by service ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param field Name of the stats field.
 @return APIGetHistStatsFieldRequest
*/
func (a *HistoricalAPIService) GetHistStatsField(ctx context.Context, field string) APIGetHistStatsFieldRequest {
	return APIGetHistStatsFieldRequest{
		APIService: a,
		ctx:        ctx,
		field:      field,
	}
}

// GetHistStatsFieldExecute executes the request
//  @return HistoricalStatsByServiceResponse
func (a *HistoricalAPIService) GetHistStatsFieldExecute(r APIGetHistStatsFieldRequest) (*HistoricalStatsByServiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalStatsByServiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetHistStatsField")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/field/{field}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"field"+"}", gourl.PathEscape(parameterToString(r.field, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.by != nil {
		localVarQueryParams.Add("by", parameterToString(*r.by, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
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

// APIGetHistStatsServiceRequest represents a request for the resource.
type APIGetHistStatsServiceRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	serviceID  string
	from       *string
	to         *string
	by         *string
	region     *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetHistStatsServiceRequest) From(from string) *APIGetHistStatsServiceRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetHistStatsServiceRequest) To(to string) *APIGetHistStatsServiceRequest {
	r.to = &to
	return r
}

// By Duration of sample windows. One of:   * &#x60;hour&#x60; - Group data by hour.   * &#x60;minute&#x60; - Group data by minute.   * &#x60;day&#x60; - Group data by day.
func (r *APIGetHistStatsServiceRequest) By(by string) *APIGetHistStatsServiceRequest {
	r.by = &by
	return r
}

// Region Limit query to a specific geographic region. One of:   * &#x60;usa&#x60; - North America.   * &#x60;europe&#x60; - Europe.   * &#x60;anzac&#x60; - Australia and New Zealand.   * &#x60;asia&#x60; - Asia.   * &#x60;asia_india&#x60; - India.   * &#x60;asia_southkorea&#x60; - South Korea.   * &#x60;africa_std&#x60; - Africa.   * &#x60;mexico&#x60; - Mexico.   * &#x60;southamerica_std&#x60; - South America.
func (r *APIGetHistStatsServiceRequest) Region(region string) *APIGetHistStatsServiceRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetHistStatsServiceRequest) Execute() (*HistoricalStatsAggregatedResponse, *http.Response, error) {
	return r.APIService.GetHistStatsServiceExecute(r)
}

/*
GetHistStatsService Get historical stats for a single service

Fetches historical stats for a given service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetHistStatsServiceRequest
*/
func (a *HistoricalAPIService) GetHistStatsService(ctx context.Context, serviceID string) APIGetHistStatsServiceRequest {
	return APIGetHistStatsServiceRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// GetHistStatsServiceExecute executes the request
//  @return HistoricalStatsAggregatedResponse
func (a *HistoricalAPIService) GetHistStatsServiceExecute(r APIGetHistStatsServiceRequest) (*HistoricalStatsAggregatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalStatsAggregatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetHistStatsService")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/service/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.by != nil {
		localVarQueryParams.Add("by", parameterToString(*r.by, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
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

// APIGetHistStatsServiceFieldRequest represents a request for the resource.
type APIGetHistStatsServiceFieldRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	serviceID  string
	field      string
	from       *string
	to         *string
	by         *string
	region     *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetHistStatsServiceFieldRequest) From(from string) *APIGetHistStatsServiceFieldRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetHistStatsServiceFieldRequest) To(to string) *APIGetHistStatsServiceFieldRequest {
	r.to = &to
	return r
}

// By Duration of sample windows. One of:   * &#x60;hour&#x60; - Group data by hour.   * &#x60;minute&#x60; - Group data by minute.   * &#x60;day&#x60; - Group data by day.
func (r *APIGetHistStatsServiceFieldRequest) By(by string) *APIGetHistStatsServiceFieldRequest {
	r.by = &by
	return r
}

// Region Limit query to a specific geographic region. One of:   * &#x60;usa&#x60; - North America.   * &#x60;europe&#x60; - Europe.   * &#x60;anzac&#x60; - Australia and New Zealand.   * &#x60;asia&#x60; - Asia.   * &#x60;asia_india&#x60; - India.   * &#x60;asia_southkorea&#x60; - South Korea.   * &#x60;africa_std&#x60; - Africa.   * &#x60;mexico&#x60; - Mexico.   * &#x60;southamerica_std&#x60; - South America.
func (r *APIGetHistStatsServiceFieldRequest) Region(region string) *APIGetHistStatsServiceFieldRequest {
	r.region = &region
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetHistStatsServiceFieldRequest) Execute() (*HistoricalStatsAggregatedResponse, *http.Response, error) {
	return r.APIService.GetHistStatsServiceFieldExecute(r)
}

/*
GetHistStatsServiceField Get historical stats for a single service/field combination

Fetches the specified field from the historical stats for a given service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param field Name of the stats field.
 @return APIGetHistStatsServiceFieldRequest
*/
func (a *HistoricalAPIService) GetHistStatsServiceField(ctx context.Context, serviceID string, field string) APIGetHistStatsServiceFieldRequest {
	return APIGetHistStatsServiceFieldRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		field:      field,
	}
}

// GetHistStatsServiceFieldExecute executes the request
//  @return HistoricalStatsAggregatedResponse
func (a *HistoricalAPIService) GetHistStatsServiceFieldExecute(r APIGetHistStatsServiceFieldRequest) (*HistoricalStatsAggregatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalStatsAggregatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetHistStatsServiceField")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/service/{service_id}/field/{field}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"field"+"}", gourl.PathEscape(parameterToString(r.field, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.by != nil {
		localVarQueryParams.Add("by", parameterToString(*r.by, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
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

// APIGetRegionsRequest represents a request for the resource.
type APIGetRegionsRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
}

// Execute calls the API using the request data configured.
func (r APIGetRegionsRequest) Execute() (*HistoricalRegionsResponse, *http.Response, error) {
	return r.APIService.GetRegionsExecute(r)
}

/*
GetRegions Get region codes

Fetches the list of codes for regions that are covered by the Fastly CDN service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetRegionsRequest
*/
func (a *HistoricalAPIService) GetRegions(ctx context.Context) APIGetRegionsRequest {
	return APIGetRegionsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetRegionsExecute executes the request
//  @return HistoricalRegionsResponse
func (a *HistoricalAPIService) GetRegionsExecute(r APIGetRegionsRequest) (*HistoricalRegionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalRegionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetRegions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/regions"

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

// APIGetUsageRequest represents a request for the resource.
type APIGetUsageRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	from       *string
	to         *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetUsageRequest) From(from string) *APIGetUsageRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetUsageRequest) To(to string) *APIGetUsageRequest {
	r.to = &to
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetUsageRequest) Execute() (*HistoricalUsageAggregatedResponse, *http.Response, error) {
	return r.APIService.GetUsageExecute(r)
}

/*
GetUsage Get usage statistics

Returns usage information aggregated across all Fastly services and grouped by region. To aggregate across all Fastly services by time period, see [`/stats/aggregate`](#get-hist-stats-aggregated).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetUsageRequest
*/
func (a *HistoricalAPIService) GetUsage(ctx context.Context) APIGetUsageRequest {
	return APIGetUsageRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetUsageExecute executes the request
//  @return HistoricalUsageAggregatedResponse
func (a *HistoricalAPIService) GetUsageExecute(r APIGetUsageRequest) (*HistoricalUsageAggregatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalUsageAggregatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
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

// APIGetUsageMonthRequest represents a request for the resource.
type APIGetUsageMonthRequest struct {
	ctx           context.Context
	APIService    HistoricalAPI
	year          *string
	month         *string
	billableUnits *bool
}

// Year 4-digit year.
func (r *APIGetUsageMonthRequest) Year(year string) *APIGetUsageMonthRequest {
	r.year = &year
	return r
}

// Month 2-digit month.
func (r *APIGetUsageMonthRequest) Month(month string) *APIGetUsageMonthRequest {
	r.month = &month
	return r
}

// BillableUnits If &#x60;true&#x60;, return results as billable units.
func (r *APIGetUsageMonthRequest) BillableUnits(billableUnits bool) *APIGetUsageMonthRequest {
	r.billableUnits = &billableUnits
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetUsageMonthRequest) Execute() (*HistoricalUsageMonthResponse, *http.Response, error) {
	return r.APIService.GetUsageMonthExecute(r)
}

/*
GetUsageMonth Get month-to-date usage statistics

Returns month-to-date usage details for a given month and year. Usage details are aggregated by service and across all Fastly services, and then grouped by region. This endpoint does not use the `from` or `to` fields for selecting the date for which data is requested. Instead, it uses `month` and `year` integer fields. Both fields are optional and default to the current month and year respectively. When set, an optional `billable_units` field will convert bandwidth to GB and divide requests by 10,000.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetUsageMonthRequest
*/
func (a *HistoricalAPIService) GetUsageMonth(ctx context.Context) APIGetUsageMonthRequest {
	return APIGetUsageMonthRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetUsageMonthExecute executes the request
//  @return HistoricalUsageMonthResponse
func (a *HistoricalAPIService) GetUsageMonthExecute(r APIGetUsageMonthRequest) (*HistoricalUsageMonthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalUsageMonthResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetUsageMonth")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/usage_by_month"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.year != nil {
		localVarQueryParams.Add("year", parameterToString(*r.year, ""))
	}
	if r.month != nil {
		localVarQueryParams.Add("month", parameterToString(*r.month, ""))
	}
	if r.billableUnits != nil {
		localVarQueryParams.Add("billable_units", parameterToString(*r.billableUnits, ""))
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

// APIGetUsageServiceRequest represents a request for the resource.
type APIGetUsageServiceRequest struct {
	ctx        context.Context
	APIService HistoricalAPI
	from       *string
	to         *string
}

// From Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as &#39;yesterday&#39;, or &#39;two weeks ago&#39;. Default varies based on the value of &#x60;by&#x60;.
func (r *APIGetUsageServiceRequest) From(from string) *APIGetUsageServiceRequest {
	r.from = &from
	return r
}

// To Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as &#x60;from&#x60;.
func (r *APIGetUsageServiceRequest) To(to string) *APIGetUsageServiceRequest {
	r.to = &to
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetUsageServiceRequest) Execute() (*HistoricalUsageServiceResponse, *http.Response, error) {
	return r.APIService.GetUsageServiceExecute(r)
}

/*
GetUsageService Get usage statistics per service

Returns usage information aggregated by service and grouped by service and region. For service stats by time period, see [`/stats`](#get-hist-stats) and [`/stats/field/:field`](#get-hist-stats-field).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetUsageServiceRequest
*/
func (a *HistoricalAPIService) GetUsageService(ctx context.Context) APIGetUsageServiceRequest {
	return APIGetUsageServiceRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetUsageServiceExecute executes the request
//  @return HistoricalUsageServiceResponse
func (a *HistoricalAPIService) GetUsageServiceExecute(r APIGetUsageServiceRequest) (*HistoricalUsageServiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *HistoricalUsageServiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricalAPIService.GetUsageService")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/usage_by_service"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
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
