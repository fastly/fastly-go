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

// MetricsPlatformAPI defines an interface for interacting with the resource.
type MetricsPlatformAPI interface {

	/*
		GetPlatformMetricsServiceHistorical Get historical time series metrics for a single service

		Fetches historical metrics for a single service for a given granularity.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param granularity Duration of sample windows.
		 @return APIGetPlatformMetricsServiceHistoricalRequest
	*/
	GetPlatformMetricsServiceHistorical(ctx context.Context, serviceId string, granularity string) APIGetPlatformMetricsServiceHistoricalRequest

	// GetPlatformMetricsServiceHistoricalExecute executes the request
	//  @return PlatformMetricsResponse
	GetPlatformMetricsServiceHistoricalExecute(r APIGetPlatformMetricsServiceHistoricalRequest) (*PlatformMetricsResponse, *http.Response, error)
}

// MetricsPlatformAPIService MetricsPlatformAPI service
type MetricsPlatformAPIService service

// APIGetPlatformMetricsServiceHistoricalRequest represents a request for the resource.
type APIGetPlatformMetricsServiceHistoricalRequest struct {
	ctx         context.Context
	APIService  MetricsPlatformAPI
	serviceId   string
	granularity string
	from        *string
	to          *string
	metric      *string
	metricSet   *string
	groupBy     *string
	region      *string
	datacenter  *string
	cursor      *string
	limit       *string
}

// From A valid RFC-8339-formatted date and time indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided &#x60;granularity&#x60; value.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) From(from string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.from = &from
	return r
}

// To A valid RFC-8339-formatted date and time indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided &#x60;granularity&#x60; value.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) To(to string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.to = &to
	return r
}

// Metric The metric(s) to retrieve. Multiple values should be comma-separated.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) Metric(metric string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.metric = &metric
	return r
}

// MetricSet The metric set(s) to retrieve. Multiple values should be comma-separated.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) MetricSet(metricSet string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.metricSet = &metricSet
	return r
}

// GroupBy Field to group_by in the query. For example, &#x60;group_by&#x3D;region&#x60; will return entries for grouped by timestamp and region.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) GroupBy(groupBy string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.groupBy = &groupBy
	return r
}

// Region Limit query to one or more specific geographic regions. Values should be comma-separated.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) Region(region string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.region = &region
	return r
}

// Datacenter Limit query to one or more specific POPs. Values should be comma-separated.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) Datacenter(datacenter string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.datacenter = &datacenter
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) Cursor(cursor string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.cursor = &cursor
	return r
}

// Limit Number of results per page. The maximum is 10000.
func (r *APIGetPlatformMetricsServiceHistoricalRequest) Limit(limit string) *APIGetPlatformMetricsServiceHistoricalRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetPlatformMetricsServiceHistoricalRequest) Execute() (*PlatformMetricsResponse, *http.Response, error) {
	return r.APIService.GetPlatformMetricsServiceHistoricalExecute(r)
}

/*
GetPlatformMetricsServiceHistorical Get historical time series metrics for a single service

Fetches historical metrics for a single service for a given granularity.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param granularity Duration of sample windows.
 @return APIGetPlatformMetricsServiceHistoricalRequest
*/
func (a *MetricsPlatformAPIService) GetPlatformMetricsServiceHistorical(ctx context.Context, serviceId string, granularity string) APIGetPlatformMetricsServiceHistoricalRequest {
	return APIGetPlatformMetricsServiceHistoricalRequest{
		APIService:  a,
		ctx:         ctx,
		serviceId:   serviceId,
		granularity: granularity,
	}
}

// GetPlatformMetricsServiceHistoricalExecute executes the request
//  @return PlatformMetricsResponse
func (a *MetricsPlatformAPIService) GetPlatformMetricsServiceHistoricalExecute(r APIGetPlatformMetricsServiceHistoricalRequest) (*PlatformMetricsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PlatformMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsPlatformAPIService.GetPlatformMetricsServiceHistorical")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/platform/services/{service_id}/{granularity}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"granularity"+"}", gourl.PathEscape(parameterToString(r.granularity, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.metric != nil {
		localVarQueryParams.Add("metric", parameterToString(*r.metric, ""))
	}
	if r.metricSet != nil {
		localVarQueryParams.Add("metric_set", parameterToString(*r.metricSet, ""))
	}
	if r.groupBy != nil {
		localVarQueryParams.Add("group_by", parameterToString(*r.groupBy, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.datacenter != nil {
		localVarQueryParams.Add("datacenter", parameterToString(*r.datacenter, ""))
	}
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
