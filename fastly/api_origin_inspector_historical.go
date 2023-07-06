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

// OriginInspectorHistoricalAPI defines an interface for interacting with the resource.
type OriginInspectorHistoricalAPI interface {

	/*
	GetOriginInspectorHistorical Get historical origin data for a service

	Fetches historical origin metrics for a given Fastly service, optionally filtering and grouping the results by origin host, region, or POP.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @return APIGetOriginInspectorHistoricalRequest
	*/
	GetOriginInspectorHistorical(ctx context.Context, serviceID string) APIGetOriginInspectorHistoricalRequest

	// GetOriginInspectorHistoricalExecute executes the request
	//  @return HistoricalOriginsResponse
	GetOriginInspectorHistoricalExecute(r APIGetOriginInspectorHistoricalRequest) (*HistoricalOriginsResponse, *http.Response, error)
}

// OriginInspectorHistoricalAPIService OriginInspectorHistoricalAPI service
type OriginInspectorHistoricalAPIService service

// APIGetOriginInspectorHistoricalRequest represents a request for the resource.
type APIGetOriginInspectorHistoricalRequest struct {
	ctx context.Context
	APIService OriginInspectorHistoricalAPI
	serviceID string
	start *string
	end *string
	downsample *string
	metric *string
	groupBy *string
	limit *string
	cursor *string
	region *string
	datacenter *string
	host *string
}

// Start A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided &#x60;downsample&#x60; value.
func (r *APIGetOriginInspectorHistoricalRequest) Start(start string) *APIGetOriginInspectorHistoricalRequest {
	r.start = &start
	return r
}
// End A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided &#x60;downsample&#x60; value.
func (r *APIGetOriginInspectorHistoricalRequest) End(end string) *APIGetOriginInspectorHistoricalRequest {
	r.end = &end
	return r
}
// Downsample Duration of sample windows.
func (r *APIGetOriginInspectorHistoricalRequest) Downsample(downsample string) *APIGetOriginInspectorHistoricalRequest {
	r.downsample = &downsample
	return r
}
// Metric The metric to retrieve. Up to ten comma-separated metrics are accepted.
func (r *APIGetOriginInspectorHistoricalRequest) Metric(metric string) *APIGetOriginInspectorHistoricalRequest {
	r.metric = &metric
	return r
}
// GroupBy Dimensions to return in the query. Multiple dimensions may be separated by commas. For example, &#x60;group_by&#x3D;host&#x60; will return one timeseries for every origin host, as a total across all POPs. 
func (r *APIGetOriginInspectorHistoricalRequest) GroupBy(groupBy string) *APIGetOriginInspectorHistoricalRequest {
	r.groupBy = &groupBy
	return r
}
// Limit Number of results per page. The maximum is 200.
func (r *APIGetOriginInspectorHistoricalRequest) Limit(limit string) *APIGetOriginInspectorHistoricalRequest {
	r.limit = &limit
	return r
}
// Cursor Cursor value from a previous response to retrieve the next page. To request the first page, this should be empty.
func (r *APIGetOriginInspectorHistoricalRequest) Cursor(cursor string) *APIGetOriginInspectorHistoricalRequest {
	r.cursor = &cursor
	return r
}
// Region Limit query to one or more specific geographic regions. Values should be comma-separated. 
func (r *APIGetOriginInspectorHistoricalRequest) Region(region string) *APIGetOriginInspectorHistoricalRequest {
	r.region = &region
	return r
}
// Datacenter Limit query to one or more specific POPs. Values should be comma-separated.
func (r *APIGetOriginInspectorHistoricalRequest) Datacenter(datacenter string) *APIGetOriginInspectorHistoricalRequest {
	r.datacenter = &datacenter
	return r
}
// Host Limit query to one or more specific origin hosts. Values should be comma-separated.
func (r *APIGetOriginInspectorHistoricalRequest) Host(host string) *APIGetOriginInspectorHistoricalRequest {
	r.host = &host
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetOriginInspectorHistoricalRequest) Execute() (*HistoricalOriginsResponse, *http.Response, error) {
	return r.APIService.GetOriginInspectorHistoricalExecute(r)
}

/*
GetOriginInspectorHistorical Get historical origin data for a service

Fetches historical origin metrics for a given Fastly service, optionally filtering and grouping the results by origin host, region, or POP.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetOriginInspectorHistoricalRequest
*/
func (a *OriginInspectorHistoricalAPIService) GetOriginInspectorHistorical(ctx context.Context, serviceID string) APIGetOriginInspectorHistoricalRequest {
	return APIGetOriginInspectorHistoricalRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
	}
}

// GetOriginInspectorHistoricalExecute executes the request
//  @return HistoricalOriginsResponse
func (a *OriginInspectorHistoricalAPIService) GetOriginInspectorHistoricalExecute(r APIGetOriginInspectorHistoricalRequest) (*HistoricalOriginsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *HistoricalOriginsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginInspectorHistoricalAPIService.GetOriginInspectorHistorical")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/origins/services/{service_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	}
	if r.downsample != nil {
		localVarQueryParams.Add("downsample", parameterToString(*r.downsample, ""))
	}
	if r.metric != nil {
		localVarQueryParams.Add("metric", parameterToString(*r.metric, ""))
	}
	if r.groupBy != nil {
		localVarQueryParams.Add("group_by", parameterToString(*r.groupBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.datacenter != nil {
		localVarQueryParams.Add("datacenter", parameterToString(*r.datacenter, ""))
	}
	if r.host != nil {
		localVarQueryParams.Add("host", parameterToString(*r.host, ""))
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
