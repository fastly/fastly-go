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

// AiRuntimeControlUsageMetricsAPI defines an interface for interacting with the resource.
type AiRuntimeControlUsageMetricsAPI interface {

	/*
		ArcExportUsageMetrics Export usage metrics

		Returns usage metrics for AI Runtime Control requests. Accepts the same filter
	parameters as [get usage metrics](#arc-get-usage-metrics), except `sort` and
	`granularity`, which aren't supported here.

	Set the `Accept` header to `text/csv` to receive a CSV file download instead of a
	JSON response. In the CSV output, the column named `time` corresponds to the `date`
	field in the JSON response.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcExportUsageMetricsRequest
	*/
	ArcExportUsageMetrics(ctx context.Context) APIArcExportUsageMetricsRequest

	// ArcExportUsageMetricsExecute executes the request
	//  @return ArcUsageMetricsResponse
	ArcExportUsageMetricsExecute(r APIArcExportUsageMetricsRequest) (*ArcUsageMetricsResponse, *http.Response, error)

	/*
		ArcGetUsageMetrics Get usage metrics

		Returns usage metrics for AI Runtime Control requests, including request counts, session counts, input token counts, and output token counts.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcGetUsageMetricsRequest
	*/
	ArcGetUsageMetrics(ctx context.Context) APIArcGetUsageMetricsRequest

	// ArcGetUsageMetricsExecute executes the request
	//  @return ArcUsageMetricsResponse
	ArcGetUsageMetricsExecute(r APIArcGetUsageMetricsRequest) (*ArcUsageMetricsResponse, *http.Response, error)
}

// AiRuntimeControlUsageMetricsAPIService AiRuntimeControlUsageMetricsAPI service
type AiRuntimeControlUsageMetricsAPIService service

// APIArcExportUsageMetricsRequest represents a request for the resource.
type APIArcExportUsageMetricsRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlUsageMetricsAPI
	key        *string
	provider   *string
	model      *string
	from       *time.Time
	to         *time.Time
}

// Key Filter results to a specific virtual key ID.
func (r *APIArcExportUsageMetricsRequest) Key(key string) *APIArcExportUsageMetricsRequest {
	r.key = &key
	return r
}

// Provider Filter results by AI provider connection ID.
func (r *APIArcExportUsageMetricsRequest) Provider(provider string) *APIArcExportUsageMetricsRequest {
	r.provider = &provider
	return r
}

// Model Filter results by AI model identifier.
func (r *APIArcExportUsageMetricsRequest) Model(model string) *APIArcExportUsageMetricsRequest {
	r.model = &model
	return r
}

// From The start of the time range, in RFC 3339 format.
func (r *APIArcExportUsageMetricsRequest) From(from time.Time) *APIArcExportUsageMetricsRequest {
	r.from = &from
	return r
}

// To The end of the time range, in RFC 3339 format. Defaults to now.
func (r *APIArcExportUsageMetricsRequest) To(to time.Time) *APIArcExportUsageMetricsRequest {
	r.to = &to
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcExportUsageMetricsRequest) Execute() (*ArcUsageMetricsResponse, *http.Response, error) {
	return r.APIService.ArcExportUsageMetricsExecute(r)
}

/*
ArcExportUsageMetrics Export usage metrics

Returns usage metrics for AI Runtime Control requests. Accepts the same filter
parameters as [get usage metrics](#arc-get-usage-metrics), except `sort` and
`granularity`, which aren't supported here.

Set the `Accept` header to `text/csv` to receive a CSV file download instead of a
JSON response. In the CSV output, the column named `time` corresponds to the `date`
field in the JSON response.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcExportUsageMetricsRequest
*/
func (a *AiRuntimeControlUsageMetricsAPIService) ArcExportUsageMetrics(ctx context.Context) APIArcExportUsageMetricsRequest {
	return APIArcExportUsageMetricsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcExportUsageMetricsExecute executes the request
//  @return ArcUsageMetricsResponse
func (a *AiRuntimeControlUsageMetricsAPIService) ArcExportUsageMetricsExecute(r APIArcExportUsageMetricsRequest) (*ArcUsageMetricsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcUsageMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlUsageMetricsAPIService.ArcExportUsageMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/usage-metrics/export"

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
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv", "application/problem+json"}

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

// APIArcGetUsageMetricsRequest represents a request for the resource.
type APIArcGetUsageMetricsRequest struct {
	ctx         context.Context
	APIService  AiRuntimeControlUsageMetricsAPI
	key         *string
	provider    *string
	model       *string
	from        *time.Time
	to          *time.Time
	granularity *string
	sort        *string
}

// Key Filter results to a specific virtual key ID.
func (r *APIArcGetUsageMetricsRequest) Key(key string) *APIArcGetUsageMetricsRequest {
	r.key = &key
	return r
}

// Provider Filter results by AI provider connection ID.
func (r *APIArcGetUsageMetricsRequest) Provider(provider string) *APIArcGetUsageMetricsRequest {
	r.provider = &provider
	return r
}

// Model Filter results by AI model identifier.
func (r *APIArcGetUsageMetricsRequest) Model(model string) *APIArcGetUsageMetricsRequest {
	r.model = &model
	return r
}

// From The start of the time range, in RFC 3339 format.
func (r *APIArcGetUsageMetricsRequest) From(from time.Time) *APIArcGetUsageMetricsRequest {
	r.from = &from
	return r
}

// To The end of the time range, in RFC 3339 format. Defaults to now.
func (r *APIArcGetUsageMetricsRequest) To(to time.Time) *APIArcGetUsageMetricsRequest {
	r.to = &to
	return r
}

// Granularity The time granularity to aggregate usage by.
func (r *APIArcGetUsageMetricsRequest) Granularity(granularity string) *APIArcGetUsageMetricsRequest {
	r.granularity = &granularity
	return r
}

// Sort Sort field. Prefix with &#x60;-&#x60; for descending order (e.g., &#x60;sort&#x3D;-date&#x60;).
func (r *APIArcGetUsageMetricsRequest) Sort(sort string) *APIArcGetUsageMetricsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIArcGetUsageMetricsRequest) Execute() (*ArcUsageMetricsResponse, *http.Response, error) {
	return r.APIService.ArcGetUsageMetricsExecute(r)
}

/*
ArcGetUsageMetrics Get usage metrics

Returns usage metrics for AI Runtime Control requests, including request counts, session counts, input token counts, and output token counts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcGetUsageMetricsRequest
*/
func (a *AiRuntimeControlUsageMetricsAPIService) ArcGetUsageMetrics(ctx context.Context) APIArcGetUsageMetricsRequest {
	return APIArcGetUsageMetricsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcGetUsageMetricsExecute executes the request
//  @return ArcUsageMetricsResponse
func (a *AiRuntimeControlUsageMetricsAPIService) ArcGetUsageMetricsExecute(r APIArcGetUsageMetricsRequest) (*ArcUsageMetricsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcUsageMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlUsageMetricsAPIService.ArcGetUsageMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/usage-metrics"

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
	if r.granularity != nil {
		localVarQueryParams.Add("granularity", parameterToString(*r.granularity, ""))
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
