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
	"time"
)

// Linger please
var (
	_ context.Context
)

// DdosProtectionAPI defines an interface for interacting with the resource.
type DdosProtectionAPI interface {

	/*
		DdosProtectionEventGet Get event by ID

		Get event by ID.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param eventId Unique ID of the event.
		 @return APIDdosProtectionEventGetRequest
	*/
	DdosProtectionEventGet(ctx context.Context, eventId string) APIDdosProtectionEventGetRequest

	// DdosProtectionEventGetExecute executes the request
	//  @return DdosProtectionEvent
	DdosProtectionEventGetExecute(r APIDdosProtectionEventGetRequest) (*DdosProtectionEvent, *http.Response, error)

	/*
		DdosProtectionEventList Get events

		Get events.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIDdosProtectionEventListRequest
	*/
	DdosProtectionEventList(ctx context.Context) APIDdosProtectionEventListRequest

	// DdosProtectionEventListExecute executes the request
	//  @return InlineResponse2002
	DdosProtectionEventListExecute(r APIDdosProtectionEventListRequest) (*InlineResponse2002, *http.Response, error)

	/*
		DdosProtectionEventRuleList Get all rules for an event

		Get all rules for an event.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param eventId Unique ID of the event.
		 @return APIDdosProtectionEventRuleListRequest
	*/
	DdosProtectionEventRuleList(ctx context.Context, eventId string) APIDdosProtectionEventRuleListRequest

	// DdosProtectionEventRuleListExecute executes the request
	//  @return InlineResponse2003
	DdosProtectionEventRuleListExecute(r APIDdosProtectionEventRuleListRequest) (*InlineResponse2003, *http.Response, error)

	/*
		DdosProtectionRuleGet Get a rule by ID

		Get a rule by ID.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param ruleId Unique ID of the rule.
		 @return APIDdosProtectionRuleGetRequest
	*/
	DdosProtectionRuleGet(ctx context.Context, ruleId string) APIDdosProtectionRuleGetRequest

	// DdosProtectionRuleGetExecute executes the request
	//  @return DdosProtectionRule
	DdosProtectionRuleGetExecute(r APIDdosProtectionRuleGetRequest) (*DdosProtectionRule, *http.Response, error)

	/*
		DdosProtectionRulePatch Update rule

		Update rule.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param ruleId Unique ID of the rule.
		 @return APIDdosProtectionRulePatchRequest
	*/
	DdosProtectionRulePatch(ctx context.Context, ruleId string) APIDdosProtectionRulePatchRequest

	// DdosProtectionRulePatchExecute executes the request
	//  @return DdosProtectionRule
	DdosProtectionRulePatchExecute(r APIDdosProtectionRulePatchRequest) (*DdosProtectionRule, *http.Response, error)

	/*
		DdosProtectionTrafficStatsRuleGet Get traffic stats for a rule

		Get traffic stats for a rule.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param eventId Unique ID of the event.
		 @param ruleId Unique ID of the rule.
		 @return APIDdosProtectionTrafficStatsRuleGetRequest
	*/
	DdosProtectionTrafficStatsRuleGet(ctx context.Context, eventId string, ruleId string) APIDdosProtectionTrafficStatsRuleGetRequest

	// DdosProtectionTrafficStatsRuleGetExecute executes the request
	//  @return DdosProtectionTrafficStats
	DdosProtectionTrafficStatsRuleGetExecute(r APIDdosProtectionTrafficStatsRuleGetRequest) (*DdosProtectionTrafficStats, *http.Response, error)
}

// DdosProtectionAPIService DdosProtectionAPI service
type DdosProtectionAPIService service

// APIDdosProtectionEventGetRequest represents a request for the resource.
type APIDdosProtectionEventGetRequest struct {
	ctx        context.Context
	APIService DdosProtectionAPI
	eventId    string
}

// Execute calls the API using the request data configured.
func (r APIDdosProtectionEventGetRequest) Execute() (*DdosProtectionEvent, *http.Response, error) {
	return r.APIService.DdosProtectionEventGetExecute(r)
}

/*
DdosProtectionEventGet Get event by ID

Get event by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Unique ID of the event.
 @return APIDdosProtectionEventGetRequest
*/
func (a *DdosProtectionAPIService) DdosProtectionEventGet(ctx context.Context, eventId string) APIDdosProtectionEventGetRequest {
	return APIDdosProtectionEventGetRequest{
		APIService: a,
		ctx:        ctx,
		eventId:    eventId,
	}
}

// DdosProtectionEventGetExecute executes the request
//  @return DdosProtectionEvent
func (a *DdosProtectionAPIService) DdosProtectionEventGetExecute(r APIDdosProtectionEventGetRequest) (*DdosProtectionEvent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DdosProtectionAPIService.DdosProtectionEventGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ddos-protection/v1/events/{event_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"event_id"+"}", gourl.PathEscape(parameterToString(r.eventId, "")))

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DdosProtectionError
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

// APIDdosProtectionEventListRequest represents a request for the resource.
type APIDdosProtectionEventListRequest struct {
	ctx        context.Context
	APIService DdosProtectionAPI
	cursor     *string
	limit      *int32
	serviceId  *string
	from       *time.Time
	to         *time.Time
	name       *string
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIDdosProtectionEventListRequest) Cursor(cursor string) *APIDdosProtectionEventListRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIDdosProtectionEventListRequest) Limit(limit int32) *APIDdosProtectionEventListRequest {
	r.limit = &limit
	return r
}

// ServiceId Filter results based on a service_id.
func (r *APIDdosProtectionEventListRequest) ServiceId(serviceId string) *APIDdosProtectionEventListRequest {
	r.serviceId = &serviceId
	return r
}

// From Represents the start of a date-time range expressed in RFC 3339 format.
func (r *APIDdosProtectionEventListRequest) From(from time.Time) *APIDdosProtectionEventListRequest {
	r.from = &from
	return r
}

// To Represents the end of a date-time range expressed in RFC 3339 format.
func (r *APIDdosProtectionEventListRequest) To(to time.Time) *APIDdosProtectionEventListRequest {
	r.to = &to
	return r
}

// Name returns a pointer to a request.
func (r *APIDdosProtectionEventListRequest) Name(name string) *APIDdosProtectionEventListRequest {
	r.name = &name
	return r
}

// Execute calls the API using the request data configured.
func (r APIDdosProtectionEventListRequest) Execute() (*InlineResponse2002, *http.Response, error) {
	return r.APIService.DdosProtectionEventListExecute(r)
}

/*
DdosProtectionEventList Get events

Get events.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIDdosProtectionEventListRequest
*/
func (a *DdosProtectionAPIService) DdosProtectionEventList(ctx context.Context) APIDdosProtectionEventListRequest {
	return APIDdosProtectionEventListRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// DdosProtectionEventListExecute executes the request
//  @return InlineResponse2002
func (a *DdosProtectionAPIService) DdosProtectionEventListExecute(r APIDdosProtectionEventListRequest) (*InlineResponse2002, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse2002
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DdosProtectionAPIService.DdosProtectionEventList")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ddos-protection/v1/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DdosProtectionError
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

// APIDdosProtectionEventRuleListRequest represents a request for the resource.
type APIDdosProtectionEventRuleListRequest struct {
	ctx        context.Context
	APIService DdosProtectionAPI
	eventId    string
	cursor     *string
	limit      *int32
	include    *string
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIDdosProtectionEventRuleListRequest) Cursor(cursor string) *APIDdosProtectionEventRuleListRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIDdosProtectionEventRuleListRequest) Limit(limit int32) *APIDdosProtectionEventRuleListRequest {
	r.limit = &limit
	return r
}

// Include Include relationships. Optional. Comma-separated values.
func (r *APIDdosProtectionEventRuleListRequest) Include(include string) *APIDdosProtectionEventRuleListRequest {
	r.include = &include
	return r
}

// Execute calls the API using the request data configured.
func (r APIDdosProtectionEventRuleListRequest) Execute() (*InlineResponse2003, *http.Response, error) {
	return r.APIService.DdosProtectionEventRuleListExecute(r)
}

/*
DdosProtectionEventRuleList Get all rules for an event

Get all rules for an event.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Unique ID of the event.
 @return APIDdosProtectionEventRuleListRequest
*/
func (a *DdosProtectionAPIService) DdosProtectionEventRuleList(ctx context.Context, eventId string) APIDdosProtectionEventRuleListRequest {
	return APIDdosProtectionEventRuleListRequest{
		APIService: a,
		ctx:        ctx,
		eventId:    eventId,
	}
}

// DdosProtectionEventRuleListExecute executes the request
//  @return InlineResponse2003
func (a *DdosProtectionAPIService) DdosProtectionEventRuleListExecute(r APIDdosProtectionEventRuleListRequest) (*InlineResponse2003, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse2003
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DdosProtectionAPIService.DdosProtectionEventRuleList")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ddos-protection/v1/events/{event_id}/rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"event_id"+"}", gourl.PathEscape(parameterToString(r.eventId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DdosProtectionError
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

// APIDdosProtectionRuleGetRequest represents a request for the resource.
type APIDdosProtectionRuleGetRequest struct {
	ctx        context.Context
	APIService DdosProtectionAPI
	ruleId     string
}

// Execute calls the API using the request data configured.
func (r APIDdosProtectionRuleGetRequest) Execute() (*DdosProtectionRule, *http.Response, error) {
	return r.APIService.DdosProtectionRuleGetExecute(r)
}

/*
DdosProtectionRuleGet Get a rule by ID

Get a rule by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleId Unique ID of the rule.
 @return APIDdosProtectionRuleGetRequest
*/
func (a *DdosProtectionAPIService) DdosProtectionRuleGet(ctx context.Context, ruleId string) APIDdosProtectionRuleGetRequest {
	return APIDdosProtectionRuleGetRequest{
		APIService: a,
		ctx:        ctx,
		ruleId:     ruleId,
	}
}

// DdosProtectionRuleGetExecute executes the request
//  @return DdosProtectionRule
func (a *DdosProtectionAPIService) DdosProtectionRuleGetExecute(r APIDdosProtectionRuleGetRequest) (*DdosProtectionRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DdosProtectionAPIService.DdosProtectionRuleGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ddos-protection/v1/rules/{rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rule_id"+"}", gourl.PathEscape(parameterToString(r.ruleId, "")))

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DdosProtectionError
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

// APIDdosProtectionRulePatchRequest represents a request for the resource.
type APIDdosProtectionRulePatchRequest struct {
	ctx                     context.Context
	APIService              DdosProtectionAPI
	ruleId                  string
	ddosProtectionRulePatch *DdosProtectionRulePatch
}

// DdosProtectionRulePatch returns a pointer to a request.
func (r *APIDdosProtectionRulePatchRequest) DdosProtectionRulePatch(ddosProtectionRulePatch DdosProtectionRulePatch) *APIDdosProtectionRulePatchRequest {
	r.ddosProtectionRulePatch = &ddosProtectionRulePatch
	return r
}

// Execute calls the API using the request data configured.
func (r APIDdosProtectionRulePatchRequest) Execute() (*DdosProtectionRule, *http.Response, error) {
	return r.APIService.DdosProtectionRulePatchExecute(r)
}

/*
DdosProtectionRulePatch Update rule

Update rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleId Unique ID of the rule.
 @return APIDdosProtectionRulePatchRequest
*/
func (a *DdosProtectionAPIService) DdosProtectionRulePatch(ctx context.Context, ruleId string) APIDdosProtectionRulePatchRequest {
	return APIDdosProtectionRulePatchRequest{
		APIService: a,
		ctx:        ctx,
		ruleId:     ruleId,
	}
}

// DdosProtectionRulePatchExecute executes the request
//  @return DdosProtectionRule
func (a *DdosProtectionAPIService) DdosProtectionRulePatchExecute(r APIDdosProtectionRulePatchRequest) (*DdosProtectionRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DdosProtectionAPIService.DdosProtectionRulePatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ddos-protection/v1/rules/{rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rule_id"+"}", gourl.PathEscape(parameterToString(r.ruleId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

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
	localVarPostBody = r.ddosProtectionRulePatch
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
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DdosProtectionError
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

// APIDdosProtectionTrafficStatsRuleGetRequest represents a request for the resource.
type APIDdosProtectionTrafficStatsRuleGetRequest struct {
	ctx        context.Context
	APIService DdosProtectionAPI
	eventId    string
	ruleId     string
}

// Execute calls the API using the request data configured.
func (r APIDdosProtectionTrafficStatsRuleGetRequest) Execute() (*DdosProtectionTrafficStats, *http.Response, error) {
	return r.APIService.DdosProtectionTrafficStatsRuleGetExecute(r)
}

/*
DdosProtectionTrafficStatsRuleGet Get traffic stats for a rule

Get traffic stats for a rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId Unique ID of the event.
 @param ruleId Unique ID of the rule.
 @return APIDdosProtectionTrafficStatsRuleGetRequest
*/
func (a *DdosProtectionAPIService) DdosProtectionTrafficStatsRuleGet(ctx context.Context, eventId string, ruleId string) APIDdosProtectionTrafficStatsRuleGetRequest {
	return APIDdosProtectionTrafficStatsRuleGetRequest{
		APIService: a,
		ctx:        ctx,
		eventId:    eventId,
		ruleId:     ruleId,
	}
}

// DdosProtectionTrafficStatsRuleGetExecute executes the request
//  @return DdosProtectionTrafficStats
func (a *DdosProtectionAPIService) DdosProtectionTrafficStatsRuleGetExecute(r APIDdosProtectionTrafficStatsRuleGetRequest) (*DdosProtectionTrafficStats, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DdosProtectionTrafficStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DdosProtectionAPIService.DdosProtectionTrafficStatsRuleGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ddos-protection/v1/events/{event_id}/rules/{rule_id}/traffic-stats"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"event_id"+"}", gourl.PathEscape(parameterToString(r.eventId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rule_id"+"}", gourl.PathEscape(parameterToString(r.ruleId, "")))

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v DdosProtectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DdosProtectionError
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
