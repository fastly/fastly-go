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

// BotRequestsMetricsAPI defines an interface for interacting with the resource.
type BotRequestsMetricsAPI interface {

	/*
		GetBotRequestMetricsServiceHistorical Get per-service historical time series metrics for bot requests broken down by bot-type.

		Return historical time series metrics for bot requests broken down by bot-type.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param granularity Duration of sample windows.
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetBotRequestMetricsServiceHistoricalRequest
	*/
	GetBotRequestMetricsServiceHistorical(ctx context.Context, granularity string, serviceID string) APIGetBotRequestMetricsServiceHistoricalRequest

	// GetBotRequestMetricsServiceHistoricalExecute executes the request
	//  @return BotRequestMetricsResponse
	GetBotRequestMetricsServiceHistoricalExecute(r APIGetBotRequestMetricsServiceHistoricalRequest) (*BotRequestMetricsResponse, *http.Response, error)
}

// BotRequestsMetricsAPIService BotRequestsMetricsAPI service
type BotRequestsMetricsAPIService service

// APIGetBotRequestMetricsServiceHistoricalRequest represents a request for the resource.
type APIGetBotRequestMetricsServiceHistoricalRequest struct {
	ctx         context.Context
	APIService  BotRequestsMetricsAPI
	granularity string
	serviceID   string
	from        *string
	to          *string
	groupBy     *string
	datacenter  *string
	bot         *string
	region      *string
	cursor      *string
	metric      *string
	metricSet   *string
	limit       *string
}

// From A valid RFC-8339-formatted date and time indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided &#x60;granularity&#x60; value.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) From(from string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.from = &from
	return r
}

// To A valid RFC-8339-formatted date and time indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided &#x60;granularity&#x60; value.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) To(to string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.to = &to
	return r
}

// GroupBy Field to group by in the query. For example, &#x60;group_by&#x3D;bot_id&#x60; will return entries grouped by &#x60;timestamp&#x60; and &#x60;bot_id&#x60;.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) GroupBy(groupBy string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.groupBy = &groupBy
	return r
}

// Datacenter Limit query to one or more specific POPs. Values should be comma-separated.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) Datacenter(datacenter string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.datacenter = &datacenter
	return r
}

// Bot Filter results based by Bot ID.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) Bot(bot string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.bot = &bot
	return r
}

// Region Limit query to one or more specific geographic regions. Values should be comma-separated.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) Region(region string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.region = &region
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) Cursor(cursor string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.cursor = &cursor
	return r
}

// Metric One or more metrics to retrieve. Multiple values should be comma-separated.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) Metric(metric string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.metric = &metric
	return r
}

// MetricSet The metric set to retrieve.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) MetricSet(metricSet string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.metricSet = &metricSet
	return r
}

// Limit Number of results per page. The maximum is 10000.
func (r *APIGetBotRequestMetricsServiceHistoricalRequest) Limit(limit string) *APIGetBotRequestMetricsServiceHistoricalRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetBotRequestMetricsServiceHistoricalRequest) Execute() (*BotRequestMetricsResponse, *http.Response, error) {
	return r.APIService.GetBotRequestMetricsServiceHistoricalExecute(r)
}

/*
GetBotRequestMetricsServiceHistorical Get per-service historical time series metrics for bot requests broken down by bot-type.

Return historical time series metrics for bot requests broken down by bot-type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param granularity Duration of sample windows.
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetBotRequestMetricsServiceHistoricalRequest
*/
func (a *BotRequestsMetricsAPIService) GetBotRequestMetricsServiceHistorical(ctx context.Context, granularity string, serviceID string) APIGetBotRequestMetricsServiceHistoricalRequest {
	return APIGetBotRequestMetricsServiceHistoricalRequest{
		APIService:  a,
		ctx:         ctx,
		granularity: granularity,
		serviceID:   serviceID,
	}
}

// GetBotRequestMetricsServiceHistoricalExecute executes the request
//  @return BotRequestMetricsResponse
func (a *BotRequestsMetricsAPIService) GetBotRequestMetricsServiceHistoricalExecute(r APIGetBotRequestMetricsServiceHistoricalRequest) (*BotRequestMetricsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotRequestMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotRequestsMetricsAPIService.GetBotRequestMetricsServiceHistorical")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/bots/entities/{service_id}/{granularity}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"granularity"+"}", gourl.PathEscape(parameterToString(r.granularity, "")))
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
	if r.groupBy != nil {
		localVarQueryParams.Add("group_by", parameterToString(*r.groupBy, ""))
	}
	if r.datacenter != nil {
		localVarQueryParams.Add("datacenter", parameterToString(*r.datacenter, ""))
	}
	if r.bot != nil {
		localVarQueryParams.Add("bot", parameterToString(*r.bot, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.metric != nil {
		localVarQueryParams.Add("metric", parameterToString(*r.metric, ""))
	}
	if r.metricSet != nil {
		localVarQueryParams.Add("metric_set", parameterToString(*r.metricSet, ""))
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
