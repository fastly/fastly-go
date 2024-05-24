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

// DomainInspectorHistoricalAPI defines an interface for interacting with the resource.
type DomainInspectorHistoricalAPI interface {

	/*
	GetDomainInspectorHistorical Get historical domain data for a service

	Fetches historical domain metrics for a given Fastly service, optionally filtering and grouping the results by domain, region, or POP.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @return APIGetDomainInspectorHistoricalRequest
	*/
	GetDomainInspectorHistorical(ctx context.Context, serviceID string) APIGetDomainInspectorHistoricalRequest

	// GetDomainInspectorHistoricalExecute executes the request
	//  @return HistoricalDomainsResponse
	GetDomainInspectorHistoricalExecute(r APIGetDomainInspectorHistoricalRequest) (*HistoricalDomainsResponse, *http.Response, error)
}

// DomainInspectorHistoricalAPIService DomainInspectorHistoricalAPI service
type DomainInspectorHistoricalAPIService service

// APIGetDomainInspectorHistoricalRequest represents a request for the resource.
type APIGetDomainInspectorHistoricalRequest struct {
	ctx context.Context
	APIService DomainInspectorHistoricalAPI
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
	domain *string
}

// Start A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided &#x60;downsample&#x60; value.
func (r *APIGetDomainInspectorHistoricalRequest) Start(start string) *APIGetDomainInspectorHistoricalRequest {
	r.start = &start
	return r
}
// End A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided &#x60;downsample&#x60; value.
func (r *APIGetDomainInspectorHistoricalRequest) End(end string) *APIGetDomainInspectorHistoricalRequest {
	r.end = &end
	return r
}
// Downsample Duration of sample windows.
func (r *APIGetDomainInspectorHistoricalRequest) Downsample(downsample string) *APIGetDomainInspectorHistoricalRequest {
	r.downsample = &downsample
	return r
}
// Metric The metrics to retrieve. Multiple values should be comma-separated.
func (r *APIGetDomainInspectorHistoricalRequest) Metric(metric string) *APIGetDomainInspectorHistoricalRequest {
	r.metric = &metric
	return r
}
// GroupBy Dimensions to return in the query. Multiple dimensions may be separated by commas. For example, &#x60;group_by&#x3D;domain&#x60; will return one timeseries for every domain, as a total across all datacenters (POPs). 
func (r *APIGetDomainInspectorHistoricalRequest) GroupBy(groupBy string) *APIGetDomainInspectorHistoricalRequest {
	r.groupBy = &groupBy
	return r
}
// Limit Number of results per page. The maximum is 200.
func (r *APIGetDomainInspectorHistoricalRequest) Limit(limit string) *APIGetDomainInspectorHistoricalRequest {
	r.limit = &limit
	return r
}
// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIGetDomainInspectorHistoricalRequest) Cursor(cursor string) *APIGetDomainInspectorHistoricalRequest {
	r.cursor = &cursor
	return r
}
// Region Limit query to one or more specific geographic regions. Values should be comma-separated. 
func (r *APIGetDomainInspectorHistoricalRequest) Region(region string) *APIGetDomainInspectorHistoricalRequest {
	r.region = &region
	return r
}
// Datacenter Limit query to one or more specific POPs. Values should be comma-separated.
func (r *APIGetDomainInspectorHistoricalRequest) Datacenter(datacenter string) *APIGetDomainInspectorHistoricalRequest {
	r.datacenter = &datacenter
	return r
}
// Domain Limit query to one or more specific domains. Values should be comma-separated.
func (r *APIGetDomainInspectorHistoricalRequest) Domain(domain string) *APIGetDomainInspectorHistoricalRequest {
	r.domain = &domain
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetDomainInspectorHistoricalRequest) Execute() (*HistoricalDomainsResponse, *http.Response, error) {
	return r.APIService.GetDomainInspectorHistoricalExecute(r)
}

/*
GetDomainInspectorHistorical Get historical domain data for a service

Fetches historical domain metrics for a given Fastly service, optionally filtering and grouping the results by domain, region, or POP.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetDomainInspectorHistoricalRequest
*/
func (a *DomainInspectorHistoricalAPIService) GetDomainInspectorHistorical(ctx context.Context, serviceID string) APIGetDomainInspectorHistoricalRequest {
	return APIGetDomainInspectorHistoricalRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
	}
}

// GetDomainInspectorHistoricalExecute executes the request
//  @return HistoricalDomainsResponse
func (a *DomainInspectorHistoricalAPIService) GetDomainInspectorHistoricalExecute(r APIGetDomainInspectorHistoricalRequest) (*HistoricalDomainsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *HistoricalDomainsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DomainInspectorHistoricalAPIService.GetDomainInspectorHistorical")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/domains/services/{service_id}"
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
	if r.domain != nil {
		localVarQueryParams.Add("domain", parameterToString(*r.domain, ""))
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
