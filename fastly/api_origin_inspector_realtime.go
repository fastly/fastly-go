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

// OriginInspectorRealtimeAPI defines an interface for interacting with the resource.
type OriginInspectorRealtimeAPI interface {

	/*
		GetOriginInspectorLast120Seconds Get real-time origin data for the last 120 seconds

		Get data for the 120 seconds preceding the latest timestamp available for a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @return APIGetOriginInspectorLast120SecondsRequest
	*/
	GetOriginInspectorLast120Seconds(ctx context.Context, serviceID string) APIGetOriginInspectorLast120SecondsRequest

	// GetOriginInspectorLast120SecondsExecute executes the request
	//  @return OriginInspector
	GetOriginInspectorLast120SecondsExecute(r APIGetOriginInspectorLast120SecondsRequest) (*OriginInspector, *http.Response, error)

	/*
		GetOriginInspectorLastMaxEntries Get a limited number of real-time origin data entries

		Get data for the `max_entries` seconds preceding the latest timestamp available for a service, up to a maximum of 120 entries.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param maxEntries Maximum number of results to display.
		 @return APIGetOriginInspectorLastMaxEntriesRequest
	*/
	GetOriginInspectorLastMaxEntries(ctx context.Context, serviceID string, maxEntries int32) APIGetOriginInspectorLastMaxEntriesRequest

	// GetOriginInspectorLastMaxEntriesExecute executes the request
	//  @return OriginInspector
	GetOriginInspectorLastMaxEntriesExecute(r APIGetOriginInspectorLastMaxEntriesRequest) (*OriginInspector, *http.Response, error)

	/*
		GetOriginInspectorLastSecond Get real-time origin data from specific time.

		Get real-time origin data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.
	Due to processing latency, the earliest entry in the response dataset may be earlier than `start_timestamp` by the value of `AggregateDelay`.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param startTimestamp Timestamp in seconds (Unix epoch time).
		 @return APIGetOriginInspectorLastSecondRequest
	*/
	GetOriginInspectorLastSecond(ctx context.Context, serviceID string, startTimestamp int32) APIGetOriginInspectorLastSecondRequest

	// GetOriginInspectorLastSecondExecute executes the request
	//  @return OriginInspector
	GetOriginInspectorLastSecondExecute(r APIGetOriginInspectorLastSecondRequest) (*OriginInspector, *http.Response, error)
}

// OriginInspectorRealtimeAPIService OriginInspectorRealtimeAPI service
type OriginInspectorRealtimeAPIService service

// APIGetOriginInspectorLast120SecondsRequest represents a request for the resource.
type APIGetOriginInspectorLast120SecondsRequest struct {
	ctx        context.Context
	APIService OriginInspectorRealtimeAPI
	serviceID  string
}

// Execute calls the API using the request data configured.
func (r APIGetOriginInspectorLast120SecondsRequest) Execute() (*OriginInspector, *http.Response, error) {
	return r.APIService.GetOriginInspectorLast120SecondsExecute(r)
}

/*
GetOriginInspectorLast120Seconds Get real-time origin data for the last 120 seconds

Get data for the 120 seconds preceding the latest timestamp available for a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIGetOriginInspectorLast120SecondsRequest
*/
func (a *OriginInspectorRealtimeAPIService) GetOriginInspectorLast120Seconds(ctx context.Context, serviceID string) APIGetOriginInspectorLast120SecondsRequest {
	return APIGetOriginInspectorLast120SecondsRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
	}
}

// GetOriginInspectorLast120SecondsExecute executes the request
//  @return OriginInspector
func (a *OriginInspectorRealtimeAPIService) GetOriginInspectorLast120SecondsExecute(r APIGetOriginInspectorLast120SecondsRequest) (*OriginInspector, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OriginInspector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginInspectorRealtimeAPIService.GetOriginInspectorLast120Seconds")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/origins/{service_id}/ts/h"
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

// APIGetOriginInspectorLastMaxEntriesRequest represents a request for the resource.
type APIGetOriginInspectorLastMaxEntriesRequest struct {
	ctx        context.Context
	APIService OriginInspectorRealtimeAPI
	serviceID  string
	maxEntries int32
}

// Execute calls the API using the request data configured.
func (r APIGetOriginInspectorLastMaxEntriesRequest) Execute() (*OriginInspector, *http.Response, error) {
	return r.APIService.GetOriginInspectorLastMaxEntriesExecute(r)
}

/*
GetOriginInspectorLastMaxEntries Get a limited number of real-time origin data entries

Get data for the `max_entries` seconds preceding the latest timestamp available for a service, up to a maximum of 120 entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param maxEntries Maximum number of results to display.
 @return APIGetOriginInspectorLastMaxEntriesRequest
*/
func (a *OriginInspectorRealtimeAPIService) GetOriginInspectorLastMaxEntries(ctx context.Context, serviceID string, maxEntries int32) APIGetOriginInspectorLastMaxEntriesRequest {
	return APIGetOriginInspectorLastMaxEntriesRequest{
		APIService: a,
		ctx:        ctx,
		serviceID:  serviceID,
		maxEntries: maxEntries,
	}
}

// GetOriginInspectorLastMaxEntriesExecute executes the request
//  @return OriginInspector
func (a *OriginInspectorRealtimeAPIService) GetOriginInspectorLastMaxEntriesExecute(r APIGetOriginInspectorLastMaxEntriesRequest) (*OriginInspector, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OriginInspector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginInspectorRealtimeAPIService.GetOriginInspectorLastMaxEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/origins/{service_id}/ts/h/limit/{max_entries}"
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

// APIGetOriginInspectorLastSecondRequest represents a request for the resource.
type APIGetOriginInspectorLastSecondRequest struct {
	ctx            context.Context
	APIService     OriginInspectorRealtimeAPI
	serviceID      string
	startTimestamp int32
}

// Execute calls the API using the request data configured.
func (r APIGetOriginInspectorLastSecondRequest) Execute() (*OriginInspector, *http.Response, error) {
	return r.APIService.GetOriginInspectorLastSecondExecute(r)
}

/*
GetOriginInspectorLastSecond Get real-time origin data from specific time.

Get real-time origin data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.
Due to processing latency, the earliest entry in the response dataset may be earlier than `start_timestamp` by the value of `AggregateDelay`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param startTimestamp Timestamp in seconds (Unix epoch time).
 @return APIGetOriginInspectorLastSecondRequest
*/
func (a *OriginInspectorRealtimeAPIService) GetOriginInspectorLastSecond(ctx context.Context, serviceID string, startTimestamp int32) APIGetOriginInspectorLastSecondRequest {
	return APIGetOriginInspectorLastSecondRequest{
		APIService:     a,
		ctx:            ctx,
		serviceID:      serviceID,
		startTimestamp: startTimestamp,
	}
}

// GetOriginInspectorLastSecondExecute executes the request
//  @return OriginInspector
func (a *OriginInspectorRealtimeAPIService) GetOriginInspectorLastSecondExecute(r APIGetOriginInspectorLastSecondRequest) (*OriginInspector, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OriginInspector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginInspectorRealtimeAPIService.GetOriginInspectorLastSecond")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/origins/{service_id}/ts/{start_timestamp}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"start_timestamp"+"}", gourl.PathEscape(parameterToString(r.startTimestamp, "")))

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
