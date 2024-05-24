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

// StatsAPI defines an interface for interacting with the resource.
type StatsAPI interface {

	/*
	GetServiceStats Get stats for a service

	Get the stats from a service for a block of time. This lists all stats by PoP location, starting with AMS. This call requires parameters to select block of time to query. Use either a timestamp range (using start_time and end_time) or a specified month/year combo (using month and year).

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @return APIGetServiceStatsRequest
	*/
	GetServiceStats(ctx context.Context, serviceID string) APIGetServiceStatsRequest

	// GetServiceStatsExecute executes the request
	//  @return Stats
	GetServiceStatsExecute(r APIGetServiceStatsRequest) (*Stats, *http.Response, error)
}

// StatsAPIService StatsAPI service
type StatsAPIService service

// APIGetServiceStatsRequest represents a request for the resource.
type APIGetServiceStatsRequest struct {
	ctx context.Context
	APIService StatsAPI
	serviceID string
	month *string
	year *string
	startTime *int32
	endTime *int32
}

// Month 2-digit month.
func (r *APIGetServiceStatsRequest) Month(month string) *APIGetServiceStatsRequest {
	r.month = &month
	return r
}
// Year 4-digit year.
func (r *APIGetServiceStatsRequest) Year(year string) *APIGetServiceStatsRequest {
	r.year = &year
	return r
}
// StartTime Epoch timestamp. Limits the results returned.
func (r *APIGetServiceStatsRequest) StartTime(startTime int32) *APIGetServiceStatsRequest {
	r.startTime = &startTime
	return r
}
// EndTime Epoch timestamp. Limits the results returned.
func (r *APIGetServiceStatsRequest) EndTime(endTime int32) *APIGetServiceStatsRequest {
	r.endTime = &endTime
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetServiceStatsRequest) Execute() (*Stats, *http.Response, error) {
	return r.APIService.GetServiceStatsExecute(r)
}

/*
GetServiceStats Get stats for a service

Get the stats from a service for a block of time. This lists all stats by PoP location, starting with AMS. This call requires parameters to select block of time to query. Use either a timestamp range (using start_time and end_time) or a specified month/year combo (using month and year).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetServiceStatsRequest
*/
func (a *StatsAPIService) GetServiceStats(ctx context.Context, serviceID string) APIGetServiceStatsRequest {
	return APIGetServiceStatsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
	}
}

// GetServiceStatsExecute executes the request
//  @return Stats
func (a *StatsAPIService) GetServiceStatsExecute(r APIGetServiceStatsRequest) (*Stats, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Stats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsAPIService.GetServiceStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/stats/summary"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.month != nil {
		localVarQueryParams.Add("month", parameterToString(*r.month, ""))
	}
	if r.year != nil {
		localVarQueryParams.Add("year", parameterToString(*r.year, ""))
	}
	if r.startTime != nil {
		localVarQueryParams.Add("start_time", parameterToString(*r.startTime, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("end_time", parameterToString(*r.endTime, ""))
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
