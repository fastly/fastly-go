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

// ConditionAPI defines an interface for interacting with the resource.
type ConditionAPI interface {

	/*
		CreateCondition Create a condition

		Creates a new condition.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APICreateConditionRequest
	*/
	CreateCondition(ctx context.Context, serviceId string, versionId int32) APICreateConditionRequest

	// CreateConditionExecute executes the request
	//  @return ConditionResponse
	CreateConditionExecute(r APICreateConditionRequest) (*ConditionResponse, *http.Response, error)

	/*
		DeleteCondition Delete a condition

		Deletes the specified condition.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param conditionName Name of the condition. Required.
		 @return APIDeleteConditionRequest
	*/
	DeleteCondition(ctx context.Context, serviceId string, versionId int32, conditionName string) APIDeleteConditionRequest

	// DeleteConditionExecute executes the request
	//  @return InlineResponse200
	DeleteConditionExecute(r APIDeleteConditionRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetCondition Describe a condition

		Gets the specified condition.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param conditionName Name of the condition. Required.
		 @return APIGetConditionRequest
	*/
	GetCondition(ctx context.Context, serviceId string, versionId int32, conditionName string) APIGetConditionRequest

	// GetConditionExecute executes the request
	//  @return ConditionResponse
	GetConditionExecute(r APIGetConditionRequest) (*ConditionResponse, *http.Response, error)

	/*
		ListConditions List conditions

		Gets all conditions for a particular service and version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @return APIListConditionsRequest
	*/
	ListConditions(ctx context.Context, serviceId string, versionId int32) APIListConditionsRequest

	// ListConditionsExecute executes the request
	//  @return []ConditionResponse
	ListConditionsExecute(r APIListConditionsRequest) ([]ConditionResponse, *http.Response, error)

	/*
		UpdateCondition Update a condition

		Updates the specified condition.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceId Alphanumeric string identifying the service.
		 @param versionId Integer identifying a service version.
		 @param conditionName Name of the condition. Required.
		 @return APIUpdateConditionRequest
	*/
	UpdateCondition(ctx context.Context, serviceId string, versionId int32, conditionName string) APIUpdateConditionRequest

	// UpdateConditionExecute executes the request
	//  @return ConditionResponse
	UpdateConditionExecute(r APIUpdateConditionRequest) (*ConditionResponse, *http.Response, error)
}

// ConditionAPIService ConditionAPI service
type ConditionAPIService service

// APICreateConditionRequest represents a request for the resource.
type APICreateConditionRequest struct {
	ctx        context.Context
	APIService ConditionAPI
	serviceId  string
	versionId  int32
	comment    *string
	name       *string
	priority   *string
	statement  *string
	serviceId2 *string
	version    *string
	type_      *string
}

// Comment A freeform descriptive note.
func (r *APICreateConditionRequest) Comment(comment string) *APICreateConditionRequest {
	r.comment = &comment
	return r
}

// Name Name of the condition. Required.
func (r *APICreateConditionRequest) Name(name string) *APICreateConditionRequest {
	r.name = &name
	return r
}

// Priority A numeric string. Priority determines execution order. Lower numbers execute first.
func (r *APICreateConditionRequest) Priority(priority string) *APICreateConditionRequest {
	r.priority = &priority
	return r
}

// Statement A conditional expression in VCL used to determine if the condition is met.
func (r *APICreateConditionRequest) Statement(statement string) *APICreateConditionRequest {
	r.statement = &statement
	return r
}

// ServiceId2 returns a pointer to a request.
func (r *APICreateConditionRequest) ServiceId2(serviceId2 string) *APICreateConditionRequest {
	r.serviceId2 = &serviceId2
	return r
}

// Version A numeric string that represents the service version.
func (r *APICreateConditionRequest) Version(version string) *APICreateConditionRequest {
	r.version = &version
	return r
}

// Type_ Type of the condition. Required.
func (r *APICreateConditionRequest) Type_(type_ string) *APICreateConditionRequest {
	r.type_ = &type_
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateConditionRequest) Execute() (*ConditionResponse, *http.Response, error) {
	return r.APIService.CreateConditionExecute(r)
}

/*
CreateCondition Create a condition

Creates a new condition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APICreateConditionRequest
*/
func (a *ConditionAPIService) CreateCondition(ctx context.Context, serviceId string, versionId int32) APICreateConditionRequest {
	return APICreateConditionRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// CreateConditionExecute executes the request
//  @return ConditionResponse
func (a *ConditionAPIService) CreateConditionExecute(r APICreateConditionRequest) (*ConditionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConditionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConditionAPIService.CreateCondition")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/condition"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
	}
	if r.statement != nil {
		localVarFormParams.Add("statement", parameterToString(*r.statement, ""))
	}
	if r.serviceId2 != nil {
		paramJson, err := parameterToJSON(*r.serviceId2)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("service_id", paramJson)
	}
	if r.version != nil {
		localVarFormParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.type_ != nil {
		localVarFormParams.Add("type", parameterToString(*r.type_, ""))
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

// APIDeleteConditionRequest represents a request for the resource.
type APIDeleteConditionRequest struct {
	ctx           context.Context
	APIService    ConditionAPI
	serviceId     string
	versionId     int32
	conditionName string
}

// Execute calls the API using the request data configured.
func (r APIDeleteConditionRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteConditionExecute(r)
}

/*
DeleteCondition Delete a condition

Deletes the specified condition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param conditionName Name of the condition. Required.
 @return APIDeleteConditionRequest
*/
func (a *ConditionAPIService) DeleteCondition(ctx context.Context, serviceId string, versionId int32, conditionName string) APIDeleteConditionRequest {
	return APIDeleteConditionRequest{
		APIService:    a,
		ctx:           ctx,
		serviceId:     serviceId,
		versionId:     versionId,
		conditionName: conditionName,
	}
}

// DeleteConditionExecute executes the request
//  @return InlineResponse200
func (a *ConditionAPIService) DeleteConditionExecute(r APIDeleteConditionRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConditionAPIService.DeleteCondition")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/condition/{condition_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"condition_name"+"}", gourl.PathEscape(parameterToString(r.conditionName, "")))

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

// APIGetConditionRequest represents a request for the resource.
type APIGetConditionRequest struct {
	ctx           context.Context
	APIService    ConditionAPI
	serviceId     string
	versionId     int32
	conditionName string
}

// Execute calls the API using the request data configured.
func (r APIGetConditionRequest) Execute() (*ConditionResponse, *http.Response, error) {
	return r.APIService.GetConditionExecute(r)
}

/*
GetCondition Describe a condition

Gets the specified condition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param conditionName Name of the condition. Required.
 @return APIGetConditionRequest
*/
func (a *ConditionAPIService) GetCondition(ctx context.Context, serviceId string, versionId int32, conditionName string) APIGetConditionRequest {
	return APIGetConditionRequest{
		APIService:    a,
		ctx:           ctx,
		serviceId:     serviceId,
		versionId:     versionId,
		conditionName: conditionName,
	}
}

// GetConditionExecute executes the request
//  @return ConditionResponse
func (a *ConditionAPIService) GetConditionExecute(r APIGetConditionRequest) (*ConditionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConditionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConditionAPIService.GetCondition")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/condition/{condition_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"condition_name"+"}", gourl.PathEscape(parameterToString(r.conditionName, "")))

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

// APIListConditionsRequest represents a request for the resource.
type APIListConditionsRequest struct {
	ctx        context.Context
	APIService ConditionAPI
	serviceId  string
	versionId  int32
}

// Execute calls the API using the request data configured.
func (r APIListConditionsRequest) Execute() ([]ConditionResponse, *http.Response, error) {
	return r.APIService.ListConditionsExecute(r)
}

/*
ListConditions List conditions

Gets all conditions for a particular service and version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @return APIListConditionsRequest
*/
func (a *ConditionAPIService) ListConditions(ctx context.Context, serviceId string, versionId int32) APIListConditionsRequest {
	return APIListConditionsRequest{
		APIService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		versionId:  versionId,
	}
}

// ListConditionsExecute executes the request
//  @return []ConditionResponse
func (a *ConditionAPIService) ListConditionsExecute(r APIListConditionsRequest) ([]ConditionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []ConditionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConditionAPIService.ListConditions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/condition"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))

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

// APIUpdateConditionRequest represents a request for the resource.
type APIUpdateConditionRequest struct {
	ctx           context.Context
	APIService    ConditionAPI
	serviceId     string
	versionId     int32
	conditionName string
	comment       *string
	name          *string
	priority      *string
	statement     *string
	serviceId2    *string
	version       *string
	type_         *string
}

// Comment A freeform descriptive note.
func (r *APIUpdateConditionRequest) Comment(comment string) *APIUpdateConditionRequest {
	r.comment = &comment
	return r
}

// Name Name of the condition. Required.
func (r *APIUpdateConditionRequest) Name(name string) *APIUpdateConditionRequest {
	r.name = &name
	return r
}

// Priority A numeric string. Priority determines execution order. Lower numbers execute first.
func (r *APIUpdateConditionRequest) Priority(priority string) *APIUpdateConditionRequest {
	r.priority = &priority
	return r
}

// Statement A conditional expression in VCL used to determine if the condition is met.
func (r *APIUpdateConditionRequest) Statement(statement string) *APIUpdateConditionRequest {
	r.statement = &statement
	return r
}

// ServiceId2 returns a pointer to a request.
func (r *APIUpdateConditionRequest) ServiceId2(serviceId2 string) *APIUpdateConditionRequest {
	r.serviceId2 = &serviceId2
	return r
}

// Version A numeric string that represents the service version.
func (r *APIUpdateConditionRequest) Version(version string) *APIUpdateConditionRequest {
	r.version = &version
	return r
}

// Type_ Type of the condition. Required.
func (r *APIUpdateConditionRequest) Type_(type_ string) *APIUpdateConditionRequest {
	r.type_ = &type_
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateConditionRequest) Execute() (*ConditionResponse, *http.Response, error) {
	return r.APIService.UpdateConditionExecute(r)
}

/*
UpdateCondition Update a condition

Updates the specified condition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Alphanumeric string identifying the service.
 @param versionId Integer identifying a service version.
 @param conditionName Name of the condition. Required.
 @return APIUpdateConditionRequest
*/
func (a *ConditionAPIService) UpdateCondition(ctx context.Context, serviceId string, versionId int32, conditionName string) APIUpdateConditionRequest {
	return APIUpdateConditionRequest{
		APIService:    a,
		ctx:           ctx,
		serviceId:     serviceId,
		versionId:     versionId,
		conditionName: conditionName,
	}
}

// UpdateConditionExecute executes the request
//  @return ConditionResponse
func (a *ConditionAPIService) UpdateConditionExecute(r APIUpdateConditionRequest) (*ConditionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ConditionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConditionAPIService.UpdateCondition")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/version/{version_id}/condition/{condition_name}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"version_id"+"}", gourl.PathEscape(parameterToString(r.versionId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"condition_name"+"}", gourl.PathEscape(parameterToString(r.conditionName, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.comment != nil {
		localVarFormParams.Add("comment", parameterToString(*r.comment, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
	}
	if r.statement != nil {
		localVarFormParams.Add("statement", parameterToString(*r.statement, ""))
	}
	if r.serviceId2 != nil {
		paramJson, err := parameterToJSON(*r.serviceId2)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("service_id", paramJson)
	}
	if r.version != nil {
		localVarFormParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.type_ != nil {
		localVarFormParams.Add("type", parameterToString(*r.type_, ""))
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
