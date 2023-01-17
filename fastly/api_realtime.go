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

// RealtimeAPI defines an interface for interacting with the resource.
type RealtimeAPI interface {

	/*
	GetStatsLast120Seconds Get real-time data for the last 120 seconds

	Get data for the 120 seconds preceding the latest timestamp available for a service.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @return APIGetStatsLast120SecondsRequest
	*/
	GetStatsLast120Seconds(ctx context.Context, serviceID string) APIGetStatsLast120SecondsRequest

	// GetStatsLast120SecondsExecute executes the request
	//  @return Realtime
	GetStatsLast120SecondsExecute(r APIGetStatsLast120SecondsRequest) (*Realtime, *http.Response, error)

	/*
	GetStatsLast120SecondsLimitEntries Get a limited number of real-time data entries

	Get data for the 120 seconds preceding the latest timestamp available for a service, up to a maximum of `max_entries` entries.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param maxEntries Maximum number of results to show.
	 @return APIGetStatsLast120SecondsLimitEntriesRequest
	*/
	GetStatsLast120SecondsLimitEntries(ctx context.Context, serviceID string, maxEntries int32) APIGetStatsLast120SecondsLimitEntriesRequest

	// GetStatsLast120SecondsLimitEntriesExecute executes the request
	//  @return Realtime
	GetStatsLast120SecondsLimitEntriesExecute(r APIGetStatsLast120SecondsLimitEntriesRequest) (*Realtime, *http.Response, error)

	/*
	GetStatsLastSecond Get real-time data from specified time

	Get real-time data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param timestampInSeconds Timestamp in seconds (Unix epoch time).
	 @return APIGetStatsLastSecondRequest
	*/
	GetStatsLastSecond(ctx context.Context, serviceID string, timestampInSeconds int32) APIGetStatsLastSecondRequest

	// GetStatsLastSecondExecute executes the request
	//  @return Realtime
	GetStatsLastSecondExecute(r APIGetStatsLastSecondRequest) (*Realtime, *http.Response, error)
}

// RealtimeAPIService RealtimeAPI service
type RealtimeAPIService service

// APIGetStatsLast120SecondsRequest represents a request for the resource.
type APIGetStatsLast120SecondsRequest struct {
	ctx context.Context
	APIService RealtimeAPI
	serviceID string
}


// Execute calls the API using the request data configured.
func (r APIGetStatsLast120SecondsRequest) Execute() (*Realtime, *http.Response, error) {
	return r.APIService.GetStatsLast120SecondsExecute(r)
}

/*
GetStatsLast120Seconds Get real-time data for the last 120 seconds

Get data for the 120 seconds preceding the latest timestamp available for a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetStatsLast120SecondsRequest
*/
func (a *RealtimeAPIService) GetStatsLast120Seconds(ctx context.Context, serviceID string) APIGetStatsLast120SecondsRequest {
	return APIGetStatsLast120SecondsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
	}
}

// GetStatsLast120SecondsExecute executes the request
//  @return Realtime
func (a *RealtimeAPIService) GetStatsLast120SecondsExecute(r APIGetStatsLast120SecondsRequest) (*Realtime, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Realtime
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealtimeAPIService.GetStatsLast120Seconds")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/channel/{service_id}/ts/h"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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

// APIGetStatsLast120SecondsLimitEntriesRequest represents a request for the resource.
type APIGetStatsLast120SecondsLimitEntriesRequest struct {
	ctx context.Context
	APIService RealtimeAPI
	serviceID string
	maxEntries int32
}


// Execute calls the API using the request data configured.
func (r APIGetStatsLast120SecondsLimitEntriesRequest) Execute() (*Realtime, *http.Response, error) {
	return r.APIService.GetStatsLast120SecondsLimitEntriesExecute(r)
}

/*
GetStatsLast120SecondsLimitEntries Get a limited number of real-time data entries

Get data for the 120 seconds preceding the latest timestamp available for a service, up to a maximum of `max_entries` entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param maxEntries Maximum number of results to show.
 @return APIGetStatsLast120SecondsLimitEntriesRequest
*/
func (a *RealtimeAPIService) GetStatsLast120SecondsLimitEntries(ctx context.Context, serviceID string, maxEntries int32) APIGetStatsLast120SecondsLimitEntriesRequest {
	return APIGetStatsLast120SecondsLimitEntriesRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		maxEntries: maxEntries,
	}
}

// GetStatsLast120SecondsLimitEntriesExecute executes the request
//  @return Realtime
func (a *RealtimeAPIService) GetStatsLast120SecondsLimitEntriesExecute(r APIGetStatsLast120SecondsLimitEntriesRequest) (*Realtime, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Realtime
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealtimeAPIService.GetStatsLast120SecondsLimitEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/channel/{service_id}/ts/h/limit/{max_entries}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"max_entries"+"}", gourl.PathEscape(parameterToString(r.maxEntries, "")))

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

// APIGetStatsLastSecondRequest represents a request for the resource.
type APIGetStatsLastSecondRequest struct {
	ctx context.Context
	APIService RealtimeAPI
	serviceID string
	timestampInSeconds int32
}


// Execute calls the API using the request data configured.
func (r APIGetStatsLastSecondRequest) Execute() (*Realtime, *http.Response, error) {
	return r.APIService.GetStatsLastSecondExecute(r)
}

/*
GetStatsLastSecond Get real-time data from specified time

Get real-time data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param timestampInSeconds Timestamp in seconds (Unix epoch time).
 @return APIGetStatsLastSecondRequest
*/
func (a *RealtimeAPIService) GetStatsLastSecond(ctx context.Context, serviceID string, timestampInSeconds int32) APIGetStatsLastSecondRequest {
	return APIGetStatsLastSecondRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		timestampInSeconds: timestampInSeconds,
	}
}

// GetStatsLastSecondExecute executes the request
//  @return Realtime
func (a *RealtimeAPIService) GetStatsLastSecondExecute(r APIGetStatsLastSecondRequest) (*Realtime, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *Realtime
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealtimeAPIService.GetStatsLastSecond")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/channel/{service_id}/ts/{timestamp_in_seconds}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"timestamp_in_seconds"+"}", gourl.PathEscape(parameterToString(r.timestampInSeconds, "")))

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
