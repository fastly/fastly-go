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
)

// Linger please
var (
	_ context.Context
)

// WholePlatformDdosHistoricalAPI defines an interface for interacting with the resource.
type WholePlatformDdosHistoricalAPI interface {

	/*
	GetPlatformDdosHistorical Get historical DDoS metrics for the entire Fastly platform

	Fetches historical DDoS metrics for the entire Fastly platform.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIGetPlatformDdosHistoricalRequest
	*/
	GetPlatformDdosHistorical(ctx context.Context) APIGetPlatformDdosHistoricalRequest

	// GetPlatformDdosHistoricalExecute executes the request
	//  @return PlatformDdosResponse
	GetPlatformDdosHistoricalExecute(r APIGetPlatformDdosHistoricalRequest) (*PlatformDdosResponse, *http.Response, error)
}

// WholePlatformDdosHistoricalAPIService WholePlatformDdosHistoricalAPI service
type WholePlatformDdosHistoricalAPIService service

// APIGetPlatformDdosHistoricalRequest represents a request for the resource.
type APIGetPlatformDdosHistoricalRequest struct {
	ctx context.Context
	APIService WholePlatformDdosHistoricalAPI
	start *string
	end *string
	downsample *string
}

// Start A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the inclusive start of the query time range. If not provided, a default is chosen based on the provided &#x60;downsample&#x60; value.
func (r *APIGetPlatformDdosHistoricalRequest) Start(start string) *APIGetPlatformDdosHistoricalRequest {
	r.start = &start
	return r
}
// End A valid ISO-8601-formatted date and time, or UNIX timestamp, indicating the exclusive end of the query time range. If not provided, a default is chosen based on the provided &#x60;downsample&#x60; value.
func (r *APIGetPlatformDdosHistoricalRequest) End(end string) *APIGetPlatformDdosHistoricalRequest {
	r.end = &end
	return r
}
// Downsample Duration of sample windows.
func (r *APIGetPlatformDdosHistoricalRequest) Downsample(downsample string) *APIGetPlatformDdosHistoricalRequest {
	r.downsample = &downsample
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetPlatformDdosHistoricalRequest) Execute() (*PlatformDdosResponse, *http.Response, error) {
	return r.APIService.GetPlatformDdosHistoricalExecute(r)
}

/*
GetPlatformDdosHistorical Get historical DDoS metrics for the entire Fastly platform

Fetches historical DDoS metrics for the entire Fastly platform.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetPlatformDdosHistoricalRequest
*/
func (a *WholePlatformDdosHistoricalAPIService) GetPlatformDdosHistorical(ctx context.Context) APIGetPlatformDdosHistoricalRequest {
	return APIGetPlatformDdosHistoricalRequest{
		APIService: a,
		ctx: ctx,
	}
}

// GetPlatformDdosHistoricalExecute executes the request
//  @return PlatformDdosResponse
func (a *WholePlatformDdosHistoricalAPIService) GetPlatformDdosHistoricalExecute(r APIGetPlatformDdosHistoricalRequest) (*PlatformDdosResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *PlatformDdosResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WholePlatformDdosHistoricalAPIService.GetPlatformDdosHistorical")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/platform/ddos"

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
