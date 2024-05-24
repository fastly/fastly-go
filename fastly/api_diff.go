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

// DiffAPI defines an interface for interacting with the resource.
type DiffAPI interface {

	/*
	DiffServiceVersions Diff two service versions

	Get diff between two versions.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param fromVersionID The version number of the service to which changes in the generated VCL are being compared. Can either be a positive number from 1 to your maximum version or a negative number from -1 down (-1 is latest version etc).
	 @param toVersionID The version number of the service from which changes in the generated VCL are being compared. Uses same numbering scheme as `from`.
	 @return APIDiffServiceVersionsRequest
	*/
	DiffServiceVersions(ctx context.Context, serviceID string, fromVersionID int32, toVersionID int32) APIDiffServiceVersionsRequest

	// DiffServiceVersionsExecute executes the request
	//  @return DiffResponse
	DiffServiceVersionsExecute(r APIDiffServiceVersionsRequest) (*DiffResponse, *http.Response, error)
}

// DiffAPIService DiffAPI service
type DiffAPIService service

// APIDiffServiceVersionsRequest represents a request for the resource.
type APIDiffServiceVersionsRequest struct {
	ctx context.Context
	APIService DiffAPI
	serviceID string
	fromVersionID int32
	toVersionID int32
	format *string
}

// Format Optional method to format the diff field.
func (r *APIDiffServiceVersionsRequest) Format(format string) *APIDiffServiceVersionsRequest {
	r.format = &format
	return r
}

// Execute calls the API using the request data configured.
func (r APIDiffServiceVersionsRequest) Execute() (*DiffResponse, *http.Response, error) {
	return r.APIService.DiffServiceVersionsExecute(r)
}

/*
DiffServiceVersions Diff two service versions

Get diff between two versions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param fromVersionID The version number of the service to which changes in the generated VCL are being compared. Can either be a positive number from 1 to your maximum version or a negative number from -1 down (-1 is latest version etc).
 @param toVersionID The version number of the service from which changes in the generated VCL are being compared. Uses same numbering scheme as `from`.
 @return APIDiffServiceVersionsRequest
*/
func (a *DiffAPIService) DiffServiceVersions(ctx context.Context, serviceID string, fromVersionID int32, toVersionID int32) APIDiffServiceVersionsRequest {
	return APIDiffServiceVersionsRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		fromVersionID: fromVersionID,
		toVersionID: toVersionID,
	}
}

// DiffServiceVersionsExecute executes the request
//  @return DiffResponse
func (a *DiffAPIService) DiffServiceVersionsExecute(r APIDiffServiceVersionsRequest) (*DiffResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *DiffResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiffAPIService.DiffServiceVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/diff/from/{from_version_id}/to/{to_version_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"from_version_id"+"}", gourl.PathEscape(parameterToString(r.fromVersionID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"to_version_id"+"}", gourl.PathEscape(parameterToString(r.toVersionID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
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
