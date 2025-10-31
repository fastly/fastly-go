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

// DomainInspectorRealtimeAPI defines an interface for interacting with the resource.
type DomainInspectorRealtimeAPI interface {

	/*
		GetDomainInspectorLast120Seconds Get real-time domain data for the last 120 seconds

		Get data for the 120 seconds preceding the latest timestamp available for a service.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @return APIGetDomainInspectorLast120SecondsRequest
	*/
	GetDomainInspectorLast120Seconds(ctx context.Context, serviceId string) APIGetDomainInspectorLast120SecondsRequest

	// GetDomainInspectorLast120SecondsExecute executes the request
	//  @return DomainInspector
	GetDomainInspectorLast120SecondsExecute(r APIGetDomainInspectorLast120SecondsRequest) (*DomainInspector, *http.Response, error)

	/*
		GetDomainInspectorLastMaxEntries Get a limited number of real-time domain data entries

		Get data for the `max_entries` seconds preceding the latest timestamp available for a service, up to a maximum of 120 entries.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param maxEntries Maximum number of results to show.
		 @return APIGetDomainInspectorLastMaxEntriesRequest
	*/
	GetDomainInspectorLastMaxEntries(ctx context.Context, serviceId string, maxEntries int32) APIGetDomainInspectorLastMaxEntriesRequest

	// GetDomainInspectorLastMaxEntriesExecute executes the request
	//  @return DomainInspector
	GetDomainInspectorLastMaxEntriesExecute(r APIGetDomainInspectorLastMaxEntriesRequest) (*DomainInspector, *http.Response, error)

	/*
		GetDomainInspectorLastSecond Get real-time domain data from a specified time

		Get real-time domain data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.
	Due to processing latency, the earliest entry in the response dataset may be earlier than `start_timestamp` by the value of `AggregateDelay`.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param startTimestamp Timestamp in seconds (Unix epoch time).
		 @return APIGetDomainInspectorLastSecondRequest
	*/
	GetDomainInspectorLastSecond(ctx context.Context, serviceId string, startTimestamp int32) APIGetDomainInspectorLastSecondRequest

	// GetDomainInspectorLastSecondExecute executes the request
	//  @return DomainInspector
	GetDomainInspectorLastSecondExecute(r APIGetDomainInspectorLastSecondRequest) (*DomainInspector, *http.Response, error)
}

// DomainInspectorRealtimeAPIService DomainInspectorRealtimeAPI service
type DomainInspectorRealtimeAPIService service

// APIGetDomainInspectorLast120SecondsRequest represents a request for the resource.
type APIGetDomainInspectorLast120SecondsRequest struct {
	ctx        context.Context
	APIService DomainInspectorRealtimeAPI
	serviceId  string
}

// Execute calls the API using the request data configured.
func (r APIGetDomainInspectorLast120SecondsRequest) Execute() (*DomainInspector, *http.Response, error) {
	return r.APIService.GetDomainInspectorLast120SecondsExecute(r)
}

/*
GetDomainInspectorLast120Seconds Get real-time domain data for the last 120 seconds

Get data for the 120 seconds preceding the latest timestamp available for a service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @return APIGetDomainInspectorLast120SecondsRequest
*/
func (a *DomainInspectorRealtimeAPIService) GetDomainInspectorLast120Seconds(ctx context.Context, serviceId string) APIGetDomainInspectorLast120SecondsRequest {
	return APIGetDomainInspectorLast120SecondsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// GetDomainInspectorLast120SecondsExecute executes the request
//  @return DomainInspector
func (a *DomainInspectorRealtimeAPIService) GetDomainInspectorLast120SecondsExecute(r APIGetDomainInspectorLast120SecondsRequest) (*DomainInspector, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DomainInspector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DomainInspectorRealtimeAPIService.GetDomainInspectorLast120Seconds")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/domains/{service_id}/ts/h"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

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

// APIGetDomainInspectorLastMaxEntriesRequest represents a request for the resource.
type APIGetDomainInspectorLastMaxEntriesRequest struct {
	ctx        context.Context
	APIService DomainInspectorRealtimeAPI
	serviceId  string
	maxEntries int32
}

// Execute calls the API using the request data configured.
func (r APIGetDomainInspectorLastMaxEntriesRequest) Execute() (*DomainInspector, *http.Response, error) {
	return r.APIService.GetDomainInspectorLastMaxEntriesExecute(r)
}

/*
GetDomainInspectorLastMaxEntries Get a limited number of real-time domain data entries

Get data for the `max_entries` seconds preceding the latest timestamp available for a service, up to a maximum of 120 entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param maxEntries Maximum number of results to show.
 @return APIGetDomainInspectorLastMaxEntriesRequest
*/
func (a *DomainInspectorRealtimeAPIService) GetDomainInspectorLastMaxEntries(ctx context.Context, serviceId string, maxEntries int32) APIGetDomainInspectorLastMaxEntriesRequest {
	return APIGetDomainInspectorLastMaxEntriesRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		maxEntries: maxEntries,
	}
}

// GetDomainInspectorLastMaxEntriesExecute executes the request
//  @return DomainInspector
func (a *DomainInspectorRealtimeAPIService) GetDomainInspectorLastMaxEntriesExecute(r APIGetDomainInspectorLastMaxEntriesRequest) (*DomainInspector, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DomainInspector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DomainInspectorRealtimeAPIService.GetDomainInspectorLastMaxEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/domains/{service_id}/ts/h/limit/{max_entries}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
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

// APIGetDomainInspectorLastSecondRequest represents a request for the resource.
type APIGetDomainInspectorLastSecondRequest struct {
	ctx            context.Context
	APIService     DomainInspectorRealtimeAPI
	serviceId      string
	startTimestamp int32
}

// Execute calls the API using the request data configured.
func (r APIGetDomainInspectorLastSecondRequest) Execute() (*DomainInspector, *http.Response, error) {
	return r.APIService.GetDomainInspectorLastSecondExecute(r)
}

/*
GetDomainInspectorLastSecond Get real-time domain data from a specified time

Get real-time domain data for the specified reporting period. Specify `0` to get a single entry for the last complete second. The `Timestamp` field included in the response provides the time index of the latest entry in the dataset and can be provided as the `start_timestamp` of the next request for a seamless continuation of the dataset from one request to the next.
Due to processing latency, the earliest entry in the response dataset may be earlier than `start_timestamp` by the value of `AggregateDelay`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param startTimestamp Timestamp in seconds (Unix epoch time).
 @return APIGetDomainInspectorLastSecondRequest
*/
func (a *DomainInspectorRealtimeAPIService) GetDomainInspectorLastSecond(ctx context.Context, serviceId string, startTimestamp int32) APIGetDomainInspectorLastSecondRequest {
	return APIGetDomainInspectorLastSecondRequest{
		APIService:     a,
		ctx:            ctx,
		serviceId:      serviceId,
		startTimestamp: startTimestamp,
	}
}

// GetDomainInspectorLastSecondExecute executes the request
//  @return DomainInspector
func (a *DomainInspectorRealtimeAPIService) GetDomainInspectorLastSecondExecute(r APIGetDomainInspectorLastSecondRequest) (*DomainInspector, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DomainInspector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DomainInspectorRealtimeAPIService.GetDomainInspectorLastSecond")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/domains/{service_id}/ts/{start_timestamp}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
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
