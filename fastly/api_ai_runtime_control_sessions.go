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
	"time"
)

// Linger please
var (
	_ context.Context
)

// AiRuntimeControlSessionsAPI defines an interface for interacting with the resource.
type AiRuntimeControlSessionsAPI interface {

	/*
		ArcGetSessionLogs Get session logs

		Returns session logs, including request and response data and metadata, for debugging and audit purposes.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcGetSessionLogsRequest
	*/
	ArcGetSessionLogs(ctx context.Context) APIArcGetSessionLogsRequest

	// ArcGetSessionLogsExecute executes the request
	//  @return ArcSessionsResponse
	ArcGetSessionLogsExecute(r APIArcGetSessionLogsRequest) (*ArcSessionsResponse, *http.Response, error)
}

// AiRuntimeControlSessionsAPIService AiRuntimeControlSessionsAPI service
type AiRuntimeControlSessionsAPIService service

// APIArcGetSessionLogsRequest represents a request for the resource.
type APIArcGetSessionLogsRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlSessionsAPI
	key        *string
	provider   *string
	model      *string
	from       *time.Time
	to         *time.Time
	cursor     *string
	limit      *int32
	sort       *string
}

// Key Filter results to a specific virtual key ID.
func (r *APIArcGetSessionLogsRequest) Key(key string) *APIArcGetSessionLogsRequest {
	r.key = &key
	return r
}

// Provider Filter results by AI provider connection ID.
func (r *APIArcGetSessionLogsRequest) Provider(provider string) *APIArcGetSessionLogsRequest {
	r.provider = &provider
	return r
}

// Model Filter results by AI model identifier.
func (r *APIArcGetSessionLogsRequest) Model(model string) *APIArcGetSessionLogsRequest {
	r.model = &model
	return r
}

// From The start of the time range, in RFC 3339 format.
func (r *APIArcGetSessionLogsRequest) From(from time.Time) *APIArcGetSessionLogsRequest {
	r.from = &from
	return r
}

// To The end of the time range, in RFC 3339 format. Defaults to &#x60;now&#x60;.
func (r *APIArcGetSessionLogsRequest) To(to time.Time) *APIArcGetSessionLogsRequest {
	r.to = &to
	return r
}

// Cursor Cursor for pagination. Use the &#x60;next_cursor&#x60; value from a previous response to request the next page of results, or the &#x60;prev_cursor&#x60; value to request the previous page.
func (r *APIArcGetSessionLogsRequest) Cursor(cursor string) *APIArcGetSessionLogsRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIArcGetSessionLogsRequest) Limit(limit int32) *APIArcGetSessionLogsRequest {
	r.limit = &limit
	return r
}

// Sort Sort field. Prefix with &#x60;-&#x60; for descending order (e.g., &#x60;sort&#x3D;-created_at&#x60;).
func (r *APIArcGetSessionLogsRequest) Sort(sort string) *APIArcGetSessionLogsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcGetSessionLogsRequest) Execute() (*ArcSessionsResponse, *http.Response, error) {
	return r.APIService.ArcGetSessionLogsExecute(r)
}

/*
ArcGetSessionLogs Get session logs

Returns session logs, including request and response data and metadata, for debugging and audit purposes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcGetSessionLogsRequest
*/
func (a *AiRuntimeControlSessionsAPIService) ArcGetSessionLogs(ctx context.Context) APIArcGetSessionLogsRequest {
	return APIArcGetSessionLogsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcGetSessionLogsExecute executes the request
//  @return ArcSessionsResponse
func (a *AiRuntimeControlSessionsAPIService) ArcGetSessionLogsExecute(r APIArcGetSessionLogsRequest) (*ArcSessionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcSessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlSessionsAPIService.ArcGetSessionLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.key != nil {
		localVarQueryParams.Add("key", parameterToString(*r.key, ""))
	}
	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	if r.model != nil {
		localVarQueryParams.Add("model", parameterToString(*r.model, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ArcError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ArcError
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
