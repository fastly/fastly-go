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

// NgwafSimulateAPI defines an interface for interacting with the resource.
type NgwafSimulateAPI interface {

	/*
		NgwafSimulateWafRequest Simulate a WAF request

		Simulates a request through the workspace's WAF configuration and returns
	the WAF response code and any signals that would be detected. The operation
	is stateless — no simulation data is persisted.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the workspace.
		 @return APINgwafSimulateWafRequestRequest
	*/
	NgwafSimulateWafRequest(ctx context.Context, workspaceId string) APINgwafSimulateWafRequestRequest

	// NgwafSimulateWafRequestExecute executes the request
	//  @return WafSimulateResponse
	NgwafSimulateWafRequestExecute(r APINgwafSimulateWafRequestRequest) (*WafSimulateResponse, *http.Response, error)
}

// NgwafSimulateAPIService NgwafSimulateAPI service
type NgwafSimulateAPIService service

// APINgwafSimulateWafRequestRequest represents a request for the resource.
type APINgwafSimulateWafRequestRequest struct {
	ctx                context.Context
	APIService         NgwafSimulateAPI
	workspaceId        string
	wafSimulateRequest *WafSimulateRequest
}

// WafSimulateRequest returns a pointer to a request.
func (r *APINgwafSimulateWafRequestRequest) WafSimulateRequest(wafSimulateRequest WafSimulateRequest) *APINgwafSimulateWafRequestRequest {
	r.wafSimulateRequest = &wafSimulateRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APINgwafSimulateWafRequestRequest) Execute() (*WafSimulateResponse, *http.Response, error) {
	return r.APIService.NgwafSimulateWafRequestExecute(r)
}

/*
NgwafSimulateWafRequest Simulate a WAF request

Simulates a request through the workspace's WAF configuration and returns
the WAF response code and any signals that would be detected. The operation
is stateless — no simulation data is persisted.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the workspace.
 @return APINgwafSimulateWafRequestRequest
*/
func (a *NgwafSimulateAPIService) NgwafSimulateWafRequest(ctx context.Context, workspaceId string) APINgwafSimulateWafRequestRequest {
	return APINgwafSimulateWafRequestRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// NgwafSimulateWafRequestExecute executes the request
//  @return WafSimulateResponse
func (a *NgwafSimulateAPIService) NgwafSimulateWafRequestExecute(r APINgwafSimulateWafRequestRequest) (*WafSimulateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *WafSimulateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NgwafSimulateAPIService.NgwafSimulateWafRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ngwaf/v1/workspaces/{workspace_id}/simulate"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.wafSimulateRequest == nil {
		return localVarReturnValue, nil, reportError("wafSimulateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.wafSimulateRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v interface{}
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
