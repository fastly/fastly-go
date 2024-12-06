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

// InsightsAPI defines an interface for interacting with the resource.
type InsightsAPI interface {

	/*
		GetLogInsights Retrieve log insights

		Retrieves statistics from sampled log records.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIGetLogInsightsRequest
	*/
	GetLogInsights(ctx context.Context) APIGetLogInsightsRequest

	// GetLogInsightsExecute executes the request
	//  @return GetLogInsightsResponse
	GetLogInsightsExecute(r APIGetLogInsightsRequest) (*GetLogInsightsResponse, *http.Response, error)
}

// InsightsAPIService InsightsAPI service
type InsightsAPIService service

// APIGetLogInsightsRequest represents a request for the resource.
type APIGetLogInsightsRequest struct {
	ctx              context.Context
	APIService       InsightsAPI
	visualization    *string
	serviceID        *string
	start            *string
	end              *string
	pops             *string
	domain           *string
	domainExactMatch *bool
	limit            *float32
}

// Visualization returns a pointer to a request.
func (r *APIGetLogInsightsRequest) Visualization(visualization string) *APIGetLogInsightsRequest {
	r.visualization = &visualization
	return r
}

// ServiceID returns a pointer to a request.
func (r *APIGetLogInsightsRequest) ServiceID(serviceID string) *APIGetLogInsightsRequest {
	r.serviceID = &serviceID
	return r
}

// Start returns a pointer to a request.
func (r *APIGetLogInsightsRequest) Start(start string) *APIGetLogInsightsRequest {
	r.start = &start
	return r
}

// End returns a pointer to a request.
func (r *APIGetLogInsightsRequest) End(end string) *APIGetLogInsightsRequest {
	r.end = &end
	return r
}

// Pops returns a pointer to a request.
func (r *APIGetLogInsightsRequest) Pops(pops string) *APIGetLogInsightsRequest {
	r.pops = &pops
	return r
}

// Domain returns a pointer to a request.
func (r *APIGetLogInsightsRequest) Domain(domain string) *APIGetLogInsightsRequest {
	r.domain = &domain
	return r
}

// DomainExactMatch returns a pointer to a request.
func (r *APIGetLogInsightsRequest) DomainExactMatch(domainExactMatch bool) *APIGetLogInsightsRequest {
	r.domainExactMatch = &domainExactMatch
	return r
}

// Limit returns a pointer to a request.
func (r *APIGetLogInsightsRequest) Limit(limit float32) *APIGetLogInsightsRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIGetLogInsightsRequest) Execute() (*GetLogInsightsResponse, *http.Response, error) {
	return r.APIService.GetLogInsightsExecute(r)
}

/*
GetLogInsights Retrieve log insights

Retrieves statistics from sampled log records.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIGetLogInsightsRequest
*/
func (a *InsightsAPIService) GetLogInsights(ctx context.Context) APIGetLogInsightsRequest {
	return APIGetLogInsightsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// GetLogInsightsExecute executes the request
//  @return GetLogInsightsResponse
func (a *InsightsAPIService) GetLogInsightsExecute(r APIGetLogInsightsRequest) (*GetLogInsightsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GetLogInsightsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightsAPIService.GetLogInsights")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/observability/log-insights"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.visualization == nil {
		return localVarReturnValue, nil, reportError("visualization is required and must be specified")
	}
	if r.serviceID == nil {
		return localVarReturnValue, nil, reportError("serviceID is required and must be specified")
	}
	if r.start == nil {
		return localVarReturnValue, nil, reportError("start is required and must be specified")
	}
	if r.end == nil {
		return localVarReturnValue, nil, reportError("end is required and must be specified")
	}

	localVarQueryParams.Add("visualization", parameterToString(*r.visualization, ""))
	localVarQueryParams.Add("service_id", parameterToString(*r.serviceID, ""))
	localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	if r.pops != nil {
		localVarQueryParams.Add("pops", parameterToString(*r.pops, ""))
	}
	if r.domain != nil {
		localVarQueryParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.domainExactMatch != nil {
		localVarQueryParams.Add("domain_exact_match", parameterToString(*r.domainExactMatch, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
