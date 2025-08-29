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
)

// Linger please
var (
	_ context.Context
)

// ObservabilityTimeseriesAPI defines an interface for interacting with the resource.
type ObservabilityTimeseriesAPI interface {

	/*
		TimeseriesGet Retrieve observability data as a time series

		Retrieves observability data as a time series.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APITimeseriesGetRequest
	*/
	TimeseriesGet(ctx context.Context) APITimeseriesGetRequest

	// TimeseriesGetExecute executes the request
	//  @return TimeseriesGetResponse
	TimeseriesGetExecute(r APITimeseriesGetRequest) (*TimeseriesGetResponse, *http.Response, error)
}

// ObservabilityTimeseriesAPIService ObservabilityTimeseriesAPI service
type ObservabilityTimeseriesAPIService service

// APITimeseriesGetRequest represents a request for the resource.
type APITimeseriesGetRequest struct {
	ctx         context.Context
	APIService  ObservabilityTimeseriesAPI
	source      *string
	from        *string
	to          *string
	granularity *string
	series      *string
	dimensions  *string
	filter      *string
}

// Source returns a pointer to a request.
func (r *APITimeseriesGetRequest) Source(source string) *APITimeseriesGetRequest {
	r.source = &source
	return r
}

// From returns a pointer to a request.
func (r *APITimeseriesGetRequest) From(from string) *APITimeseriesGetRequest {
	r.from = &from
	return r
}

// To returns a pointer to a request.
func (r *APITimeseriesGetRequest) To(to string) *APITimeseriesGetRequest {
	r.to = &to
	return r
}

// Granularity returns a pointer to a request.
func (r *APITimeseriesGetRequest) Granularity(granularity string) *APITimeseriesGetRequest {
	r.granularity = &granularity
	return r
}

// Series returns a pointer to a request.
func (r *APITimeseriesGetRequest) Series(series string) *APITimeseriesGetRequest {
	r.series = &series
	return r
}

// Dimensions returns a pointer to a request.
func (r *APITimeseriesGetRequest) Dimensions(dimensions string) *APITimeseriesGetRequest {
	r.dimensions = &dimensions
	return r
}

// Filter returns a pointer to a request.
func (r *APITimeseriesGetRequest) Filter(filter string) *APITimeseriesGetRequest {
	r.filter = &filter
	return r
}

// Execute calls the API using the request data configured.
func (r APITimeseriesGetRequest) Execute() (*TimeseriesGetResponse, *http.Response, error) {
	return r.APIService.TimeseriesGetExecute(r)
}

/*
TimeseriesGet Retrieve observability data as a time series

Retrieves observability data as a time series.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APITimeseriesGetRequest
*/
func (a *ObservabilityTimeseriesAPIService) TimeseriesGet(ctx context.Context) APITimeseriesGetRequest {
	return APITimeseriesGetRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// TimeseriesGetExecute executes the request
//  @return TimeseriesGetResponse
func (a *ObservabilityTimeseriesAPIService) TimeseriesGetExecute(r APITimeseriesGetRequest) (*TimeseriesGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TimeseriesGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObservabilityTimeseriesAPIService.TimeseriesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/timeseries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.source == nil {
		return localVarReturnValue, nil, reportError("source is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}
	if r.series == nil {
		return localVarReturnValue, nil, reportError("series is required and must be specified")
	}

	localVarQueryParams.Add("source", parameterToString(*r.source, ""))
	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	localVarQueryParams.Add("granularity", parameterToString(*r.granularity, ""))
	if r.dimensions != nil {
		localVarQueryParams.Add("dimensions", parameterToString(*r.dimensions, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	localVarQueryParams.Add("series", parameterToString(*r.series, ""))
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
