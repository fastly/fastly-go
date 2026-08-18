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

// LoggingEndpointErrorsAPI defines an interface for interacting with the resource.
type LoggingEndpointErrorsAPI interface {

	/*
		GetLogEndpointErrors Stream Log Endpoint Errors

		Provides a near real-time stream of log errors through a hybrid short-polling model.
	A client should make an initial request using the `from` parameter to specify a start time.
	The `to` parameter should be used alongside the `from` parameter since the default bucket is 10 seconds.

	For pagination, use the URLs provided in the Link header of the response. These contain
	updated `from` timestamps for retrieving the next or previous page of logs.

	Defaults to `application/x-ndjson` format. Use `Accept: application/json` header
	to request standard JSON array format instead.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId
		 @return APIGetLogEndpointErrorsRequest
	*/
	GetLogEndpointErrors(ctx context.Context, serviceId string) APIGetLogEndpointErrorsRequest

	// GetLogEndpointErrorsExecute executes the request
	//  @return string
	GetLogEndpointErrorsExecute(r APIGetLogEndpointErrorsRequest) (string, *http.Response, error)
}

// LoggingEndpointErrorsAPIService LoggingEndpointErrorsAPI service
type LoggingEndpointErrorsAPIService service

// APIGetLogEndpointErrorsRequest represents a request for the resource.
type APIGetLogEndpointErrorsRequest struct {
	ctx            context.Context
	APIService     LoggingEndpointErrorsAPI
	serviceId      string
	from           *int64
	to             *int64
	filterEndpoint *string
}

// From returns a pointer to a request.
func (r *APIGetLogEndpointErrorsRequest) From(from int64) *APIGetLogEndpointErrorsRequest {
	r.from = &from
	return r
}

// To returns a pointer to a request.
func (r *APIGetLogEndpointErrorsRequest) To(to int64) *APIGetLogEndpointErrorsRequest {
	r.to = &to
	return r
}

// FilterEndpoint returns a pointer to a request.
func (r *APIGetLogEndpointErrorsRequest) FilterEndpoint(filterEndpoint string) *APIGetLogEndpointErrorsRequest {
	r.filterEndpoint = &filterEndpoint
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetLogEndpointErrorsRequest) Execute() (string, *http.Response, error) {
	return r.APIService.GetLogEndpointErrorsExecute(r)
}

/*
GetLogEndpointErrors Stream Log Endpoint Errors

Provides a near real-time stream of log errors through a hybrid short-polling model.
A client should make an initial request using the `from` parameter to specify a start time.
The `to` parameter should be used alongside the `from` parameter since the default bucket is 10 seconds.

For pagination, use the URLs provided in the Link header of the response. These contain
updated `from` timestamps for retrieving the next or previous page of logs.

Defaults to `application/x-ndjson` format. Use `Accept: application/json` header
to request standard JSON array format instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return APIGetLogEndpointErrorsRequest
*/
func (a *LoggingEndpointErrorsAPIService) GetLogEndpointErrors(ctx context.Context, serviceId string) APIGetLogEndpointErrorsRequest {
	return APIGetLogEndpointErrorsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// GetLogEndpointErrorsExecute executes the request
//  @return string
func (a *LoggingEndpointErrorsAPIService) GetLogEndpointErrorsExecute(r APIGetLogEndpointErrorsRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingEndpointErrorsAPIService.GetLogEndpointErrors")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/service/{service_id}/logging/errors"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.filterEndpoint != nil {
		localVarQueryParams.Add("filter[endpoint]", parameterToString(*r.filterEndpoint, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson", "application/json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
