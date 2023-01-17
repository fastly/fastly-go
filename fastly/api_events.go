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
	"strings"
)

// Linger please
var (
	_ context.Context
)

// EventsAPI defines an interface for interacting with the resource.
type EventsAPI interface {

	/*
	GetEvent Get an event

	Get a specific event.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param eventID Alphanumeric string identifying an event.
	 @return APIGetEventRequest
	*/
	GetEvent(ctx context.Context, eventID string) APIGetEventRequest

	// GetEventExecute executes the request
	//  @return EventResponse
	GetEventExecute(r APIGetEventRequest) (*EventResponse, *http.Response, error)

	/*
	ListEvents List events

	List all events for a particular customer. Events can be filtered by user, customer and event type. Events can be sorted by date.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return APIListEventsRequest
	*/
	ListEvents(ctx context.Context) APIListEventsRequest

	// ListEventsExecute executes the request
	//  @return EventsResponse
	ListEventsExecute(r APIListEventsRequest) (*EventsResponse, *http.Response, error)
}

// EventsAPIService EventsAPI service
type EventsAPIService service

// APIGetEventRequest represents a request for the resource.
type APIGetEventRequest struct {
	ctx context.Context
	APIService EventsAPI
	eventID string
}


// Execute calls the API using the request data configured.
func (r APIGetEventRequest) Execute() (*EventResponse, *http.Response, error) {
	return r.APIService.GetEventExecute(r)
}

/*
GetEvent Get an event

Get a specific event.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventID Alphanumeric string identifying an event.
 @return APIGetEventRequest
*/
func (a *EventsAPIService) GetEvent(ctx context.Context, eventID string) APIGetEventRequest {
	return APIGetEventRequest{
		APIService: a,
		ctx: ctx,
		eventID: eventID,
	}
}

// GetEventExecute executes the request
//  @return EventResponse
func (a *EventsAPIService) GetEventExecute(r APIGetEventRequest) (*EventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *EventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.GetEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{event_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"event_id"+"}", gourl.PathEscape(parameterToString(r.eventID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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

// APIListEventsRequest represents a request for the resource.
type APIListEventsRequest struct {
	ctx context.Context
	APIService EventsAPI
	filterCustomerID *string
	filterEventType *string
	filterServiceID *string
	filterUserID *string
	filterTokenID *string
	pageNumber *int32
	pageSize *int32
	sort *string
}

// FilterCustomerID Limit the results returned to a specific customer.
func (r *APIListEventsRequest) FilterCustomerID(filterCustomerID string) *APIListEventsRequest {
	r.filterCustomerID = &filterCustomerID
	return r
}
// FilterEventType Limit the returned events to a specific &#x60;event_type&#x60;.
func (r *APIListEventsRequest) FilterEventType(filterEventType string) *APIListEventsRequest {
	r.filterEventType = &filterEventType
	return r
}
// FilterServiceID Limit the results returned to a specific service.
func (r *APIListEventsRequest) FilterServiceID(filterServiceID string) *APIListEventsRequest {
	r.filterServiceID = &filterServiceID
	return r
}
// FilterUserID Limit the results returned to a specific user.
func (r *APIListEventsRequest) FilterUserID(filterUserID string) *APIListEventsRequest {
	r.filterUserID = &filterUserID
	return r
}
// FilterTokenID Limit the returned events to a specific token.
func (r *APIListEventsRequest) FilterTokenID(filterTokenID string) *APIListEventsRequest {
	r.filterTokenID = &filterTokenID
	return r
}
// PageNumber Current page.
func (r *APIListEventsRequest) PageNumber(pageNumber int32) *APIListEventsRequest {
	r.pageNumber = &pageNumber
	return r
}
// PageSize Number of records per page.
func (r *APIListEventsRequest) PageSize(pageSize int32) *APIListEventsRequest {
	r.pageSize = &pageSize
	return r
}
// Sort The order in which to list the results by creation date.
func (r *APIListEventsRequest) Sort(sort string) *APIListEventsRequest {
	r.sort = &sort
	return r
}

// Execute calls the API using the request data configured.
func (r APIListEventsRequest) Execute() (*EventsResponse, *http.Response, error) {
	return r.APIService.ListEventsExecute(r)
}

/*
ListEvents List events

List all events for a particular customer. Events can be filtered by user, customer and event type. Events can be sorted by date.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListEventsRequest
*/
func (a *EventsAPIService) ListEvents(ctx context.Context) APIListEventsRequest {
	return APIListEventsRequest{
		APIService: a,
		ctx: ctx,
	}
}

// ListEventsExecute executes the request
//  @return EventsResponse
func (a *EventsAPIService) ListEventsExecute(r APIListEventsRequest) (*EventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *EventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.ListEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.filterCustomerID != nil {
		localVarQueryParams.Add("filter[customer_id]", parameterToString(*r.filterCustomerID, ""))
	}
	if r.filterEventType != nil {
		localVarQueryParams.Add("filter[event_type]", parameterToString(*r.filterEventType, ""))
	}
	if r.filterServiceID != nil {
		localVarQueryParams.Add("filter[service_id]", parameterToString(*r.filterServiceID, ""))
	}
	if r.filterUserID != nil {
		localVarQueryParams.Add("filter[user_id]", parameterToString(*r.filterUserID, ""))
	}
	if r.filterTokenID != nil {
		localVarQueryParams.Add("filter[token_id]", parameterToString(*r.filterTokenID, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

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
