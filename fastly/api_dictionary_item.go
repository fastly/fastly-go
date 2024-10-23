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

// DictionaryItemAPI defines an interface for interacting with the resource.
type DictionaryItemAPI interface {

	/*
		BulkUpdateDictionaryItem Update multiple entries in an edge dictionary

		Update multiple items in the same dictionary. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 items. [Contact support](https://support.fastly.com/) to discuss raising this limit.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @return APIBulkUpdateDictionaryItemRequest
	*/
	BulkUpdateDictionaryItem(ctx context.Context, serviceID string, dictionaryID string) APIBulkUpdateDictionaryItemRequest

	// BulkUpdateDictionaryItemExecute executes the request
	//  @return InlineResponse200
	BulkUpdateDictionaryItemExecute(r APIBulkUpdateDictionaryItemRequest) (*InlineResponse200, *http.Response, error)

	/*
		CreateDictionaryItem Create an entry in an edge dictionary

		Create DictionaryItem given service, dictionary ID, item key, and item value.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @return APICreateDictionaryItemRequest
	*/
	CreateDictionaryItem(ctx context.Context, serviceID string, dictionaryID string) APICreateDictionaryItemRequest

	// CreateDictionaryItemExecute executes the request
	//  @return DictionaryItemResponse
	CreateDictionaryItemExecute(r APICreateDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error)

	/*
		DeleteDictionaryItem Delete an item from an edge dictionary

		Delete DictionaryItem given service, dictionary ID, and item key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @param dictionaryItemKey Item key, maximum 256 characters.
		 @return APIDeleteDictionaryItemRequest
	*/
	DeleteDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIDeleteDictionaryItemRequest

	// DeleteDictionaryItemExecute executes the request
	//  @return InlineResponse200
	DeleteDictionaryItemExecute(r APIDeleteDictionaryItemRequest) (*InlineResponse200, *http.Response, error)

	/*
		GetDictionaryItem Get an item from an edge dictionary

		Retrieve a single DictionaryItem given service, dictionary ID and item key.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @param dictionaryItemKey Item key, maximum 256 characters.
		 @return APIGetDictionaryItemRequest
	*/
	GetDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIGetDictionaryItemRequest

	// GetDictionaryItemExecute executes the request
	//  @return DictionaryItemResponse
	GetDictionaryItemExecute(r APIGetDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error)

	/*
		ListDictionaryItems List items in an edge dictionary

		List of DictionaryItems given service and dictionary ID.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @return APIListDictionaryItemsRequest
	*/
	ListDictionaryItems(ctx context.Context, serviceID string, dictionaryID string) APIListDictionaryItemsRequest

	// ListDictionaryItemsExecute executes the request
	//  @return []DictionaryItemResponse
	ListDictionaryItemsExecute(r APIListDictionaryItemsRequest) ([]DictionaryItemResponse, *http.Response, error)

	/*
		UpdateDictionaryItem Update an entry in an edge dictionary

		Update DictionaryItem given service, dictionary ID, item key, and item value.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @param dictionaryItemKey Item key, maximum 256 characters.
		 @return APIUpdateDictionaryItemRequest
	*/
	UpdateDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIUpdateDictionaryItemRequest

	// UpdateDictionaryItemExecute executes the request
	//  @return DictionaryItemResponse
	UpdateDictionaryItemExecute(r APIUpdateDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error)

	/*
		UpsertDictionaryItem Insert or update an entry in an edge dictionary

		Upsert DictionaryItem given service, dictionary ID, item key, and item value.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param serviceID Alphanumeric string identifying the service.
		 @param dictionaryID Alphanumeric string identifying a Dictionary.
		 @param dictionaryItemKey Item key, maximum 256 characters.
		 @return APIUpsertDictionaryItemRequest
	*/
	UpsertDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIUpsertDictionaryItemRequest

	// UpsertDictionaryItemExecute executes the request
	//  @return DictionaryItemResponse
	UpsertDictionaryItemExecute(r APIUpsertDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error)
}

// DictionaryItemAPIService DictionaryItemAPI service
type DictionaryItemAPIService service

// APIBulkUpdateDictionaryItemRequest represents a request for the resource.
type APIBulkUpdateDictionaryItemRequest struct {
	ctx                             context.Context
	APIService                      DictionaryItemAPI
	serviceID                       string
	dictionaryID                    string
	bulkUpdateDictionaryListRequest *BulkUpdateDictionaryListRequest
}

// BulkUpdateDictionaryListRequest returns a pointer to a request.
func (r *APIBulkUpdateDictionaryItemRequest) BulkUpdateDictionaryListRequest(bulkUpdateDictionaryListRequest BulkUpdateDictionaryListRequest) *APIBulkUpdateDictionaryItemRequest {
	r.bulkUpdateDictionaryListRequest = &bulkUpdateDictionaryListRequest
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkUpdateDictionaryItemRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.BulkUpdateDictionaryItemExecute(r)
}

/*
BulkUpdateDictionaryItem Update multiple entries in an edge dictionary

Update multiple items in the same dictionary. For faster updates to your service, group your changes into large batches. The maximum batch size is 1000 items. [Contact support](https://support.fastly.com/) to discuss raising this limit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @return APIBulkUpdateDictionaryItemRequest
*/
func (a *DictionaryItemAPIService) BulkUpdateDictionaryItem(ctx context.Context, serviceID string, dictionaryID string) APIBulkUpdateDictionaryItemRequest {
	return APIBulkUpdateDictionaryItemRequest{
		APIService:   a,
		ctx:          ctx,
		serviceID:    serviceID,
		dictionaryID: dictionaryID,
	}
}

// BulkUpdateDictionaryItemExecute executes the request
//  @return InlineResponse200
func (a *DictionaryItemAPIService) BulkUpdateDictionaryItemExecute(r APIBulkUpdateDictionaryItemRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.BulkUpdateDictionaryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/items"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkUpdateDictionaryListRequest
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

// APICreateDictionaryItemRequest represents a request for the resource.
type APICreateDictionaryItemRequest struct {
	ctx          context.Context
	APIService   DictionaryItemAPI
	serviceID    string
	dictionaryID string
	itemKey      *string
	itemValue    *string
}

// ItemKey Item key, maximum 256 characters.
func (r *APICreateDictionaryItemRequest) ItemKey(itemKey string) *APICreateDictionaryItemRequest {
	r.itemKey = &itemKey
	return r
}

// ItemValue Item value, maximum 8000 characters.
func (r *APICreateDictionaryItemRequest) ItemValue(itemValue string) *APICreateDictionaryItemRequest {
	r.itemValue = &itemValue
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateDictionaryItemRequest) Execute() (*DictionaryItemResponse, *http.Response, error) {
	return r.APIService.CreateDictionaryItemExecute(r)
}

/*
CreateDictionaryItem Create an entry in an edge dictionary

Create DictionaryItem given service, dictionary ID, item key, and item value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @return APICreateDictionaryItemRequest
*/
func (a *DictionaryItemAPIService) CreateDictionaryItem(ctx context.Context, serviceID string, dictionaryID string) APICreateDictionaryItemRequest {
	return APICreateDictionaryItemRequest{
		APIService:   a,
		ctx:          ctx,
		serviceID:    serviceID,
		dictionaryID: dictionaryID,
	}
}

// CreateDictionaryItemExecute executes the request
//  @return DictionaryItemResponse
func (a *DictionaryItemAPIService) CreateDictionaryItemExecute(r APICreateDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DictionaryItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.CreateDictionaryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/item"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))

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
	if r.itemKey != nil {
		localVarFormParams.Add("item_key", parameterToString(*r.itemKey, ""))
	}
	if r.itemValue != nil {
		localVarFormParams.Add("item_value", parameterToString(*r.itemValue, ""))
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

// APIDeleteDictionaryItemRequest represents a request for the resource.
type APIDeleteDictionaryItemRequest struct {
	ctx               context.Context
	APIService        DictionaryItemAPI
	serviceID         string
	dictionaryID      string
	dictionaryItemKey string
}

// Execute calls the API using the request data configured.
func (r APIDeleteDictionaryItemRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.DeleteDictionaryItemExecute(r)
}

/*
DeleteDictionaryItem Delete an item from an edge dictionary

Delete DictionaryItem given service, dictionary ID, and item key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @param dictionaryItemKey Item key, maximum 256 characters.
 @return APIDeleteDictionaryItemRequest
*/
func (a *DictionaryItemAPIService) DeleteDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIDeleteDictionaryItemRequest {
	return APIDeleteDictionaryItemRequest{
		APIService:        a,
		ctx:               ctx,
		serviceID:         serviceID,
		dictionaryID:      dictionaryID,
		dictionaryItemKey: dictionaryItemKey,
	}
}

// DeleteDictionaryItemExecute executes the request
//  @return InlineResponse200
func (a *DictionaryItemAPIService) DeleteDictionaryItemExecute(r APIDeleteDictionaryItemRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.DeleteDictionaryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_item_key"+"}", gourl.PathEscape(parameterToString(r.dictionaryItemKey, "")))

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

// APIGetDictionaryItemRequest represents a request for the resource.
type APIGetDictionaryItemRequest struct {
	ctx               context.Context
	APIService        DictionaryItemAPI
	serviceID         string
	dictionaryID      string
	dictionaryItemKey string
}

// Execute calls the API using the request data configured.
func (r APIGetDictionaryItemRequest) Execute() (*DictionaryItemResponse, *http.Response, error) {
	return r.APIService.GetDictionaryItemExecute(r)
}

/*
GetDictionaryItem Get an item from an edge dictionary

Retrieve a single DictionaryItem given service, dictionary ID and item key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @param dictionaryItemKey Item key, maximum 256 characters.
 @return APIGetDictionaryItemRequest
*/
func (a *DictionaryItemAPIService) GetDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIGetDictionaryItemRequest {
	return APIGetDictionaryItemRequest{
		APIService:        a,
		ctx:               ctx,
		serviceID:         serviceID,
		dictionaryID:      dictionaryID,
		dictionaryItemKey: dictionaryItemKey,
	}
}

// GetDictionaryItemExecute executes the request
//  @return DictionaryItemResponse
func (a *DictionaryItemAPIService) GetDictionaryItemExecute(r APIGetDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DictionaryItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.GetDictionaryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_item_key"+"}", gourl.PathEscape(parameterToString(r.dictionaryItemKey, "")))

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

// APIListDictionaryItemsRequest represents a request for the resource.
type APIListDictionaryItemsRequest struct {
	ctx          context.Context
	APIService   DictionaryItemAPI
	serviceID    string
	dictionaryID string
	page         *int32
	perPage      *int32
	sort         *string
	direction    *string
}

// Page Current page.
func (r *APIListDictionaryItemsRequest) Page(page int32) *APIListDictionaryItemsRequest {
	r.page = &page
	return r
}

// PerPage Number of records per page.
func (r *APIListDictionaryItemsRequest) PerPage(perPage int32) *APIListDictionaryItemsRequest {
	r.perPage = &perPage
	return r
}

// Sort Field on which to sort.
func (r *APIListDictionaryItemsRequest) Sort(sort string) *APIListDictionaryItemsRequest {
	r.sort = &sort
	return r
}

// Direction Direction in which to sort results.
func (r *APIListDictionaryItemsRequest) Direction(direction string) *APIListDictionaryItemsRequest {
	r.direction = &direction
	return r
}

// Execute calls the API using the request data configured.
func (r APIListDictionaryItemsRequest) Execute() ([]DictionaryItemResponse, *http.Response, error) {
	return r.APIService.ListDictionaryItemsExecute(r)
}

/*
ListDictionaryItems List items in an edge dictionary

List of DictionaryItems given service and dictionary ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @return APIListDictionaryItemsRequest
*/
func (a *DictionaryItemAPIService) ListDictionaryItems(ctx context.Context, serviceID string, dictionaryID string) APIListDictionaryItemsRequest {
	return APIListDictionaryItemsRequest{
		APIService:   a,
		ctx:          ctx,
		serviceID:    serviceID,
		dictionaryID: dictionaryID,
	}
}

// ListDictionaryItemsExecute executes the request
//  @return []DictionaryItemResponse
func (a *DictionaryItemAPIService) ListDictionaryItemsExecute(r APIListDictionaryItemsRequest) ([]DictionaryItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []DictionaryItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.ListDictionaryItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/items"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
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

// APIUpdateDictionaryItemRequest represents a request for the resource.
type APIUpdateDictionaryItemRequest struct {
	ctx               context.Context
	APIService        DictionaryItemAPI
	serviceID         string
	dictionaryID      string
	dictionaryItemKey string
	itemKey           *string
	itemValue         *string
}

// ItemKey Item key, maximum 256 characters.
func (r *APIUpdateDictionaryItemRequest) ItemKey(itemKey string) *APIUpdateDictionaryItemRequest {
	r.itemKey = &itemKey
	return r
}

// ItemValue Item value, maximum 8000 characters.
func (r *APIUpdateDictionaryItemRequest) ItemValue(itemValue string) *APIUpdateDictionaryItemRequest {
	r.itemValue = &itemValue
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateDictionaryItemRequest) Execute() (*DictionaryItemResponse, *http.Response, error) {
	return r.APIService.UpdateDictionaryItemExecute(r)
}

/*
UpdateDictionaryItem Update an entry in an edge dictionary

Update DictionaryItem given service, dictionary ID, item key, and item value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @param dictionaryItemKey Item key, maximum 256 characters.
 @return APIUpdateDictionaryItemRequest
*/
func (a *DictionaryItemAPIService) UpdateDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIUpdateDictionaryItemRequest {
	return APIUpdateDictionaryItemRequest{
		APIService:        a,
		ctx:               ctx,
		serviceID:         serviceID,
		dictionaryID:      dictionaryID,
		dictionaryItemKey: dictionaryItemKey,
	}
}

// UpdateDictionaryItemExecute executes the request
//  @return DictionaryItemResponse
func (a *DictionaryItemAPIService) UpdateDictionaryItemExecute(r APIUpdateDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DictionaryItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.UpdateDictionaryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_item_key"+"}", gourl.PathEscape(parameterToString(r.dictionaryItemKey, "")))

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
	if r.itemKey != nil {
		localVarFormParams.Add("item_key", parameterToString(*r.itemKey, ""))
	}
	if r.itemValue != nil {
		localVarFormParams.Add("item_value", parameterToString(*r.itemValue, ""))
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

// APIUpsertDictionaryItemRequest represents a request for the resource.
type APIUpsertDictionaryItemRequest struct {
	ctx               context.Context
	APIService        DictionaryItemAPI
	serviceID         string
	dictionaryID      string
	dictionaryItemKey string
	itemKey           *string
	itemValue         *string
}

// ItemKey Item key, maximum 256 characters.
func (r *APIUpsertDictionaryItemRequest) ItemKey(itemKey string) *APIUpsertDictionaryItemRequest {
	r.itemKey = &itemKey
	return r
}

// ItemValue Item value, maximum 8000 characters.
func (r *APIUpsertDictionaryItemRequest) ItemValue(itemValue string) *APIUpsertDictionaryItemRequest {
	r.itemValue = &itemValue
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpsertDictionaryItemRequest) Execute() (*DictionaryItemResponse, *http.Response, error) {
	return r.APIService.UpsertDictionaryItemExecute(r)
}

/*
UpsertDictionaryItem Insert or update an entry in an edge dictionary

Upsert DictionaryItem given service, dictionary ID, item key, and item value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param dictionaryID Alphanumeric string identifying a Dictionary.
 @param dictionaryItemKey Item key, maximum 256 characters.
 @return APIUpsertDictionaryItemRequest
*/
func (a *DictionaryItemAPIService) UpsertDictionaryItem(ctx context.Context, serviceID string, dictionaryID string, dictionaryItemKey string) APIUpsertDictionaryItemRequest {
	return APIUpsertDictionaryItemRequest{
		APIService:        a,
		ctx:               ctx,
		serviceID:         serviceID,
		dictionaryID:      dictionaryID,
		dictionaryItemKey: dictionaryItemKey,
	}
}

// UpsertDictionaryItemExecute executes the request
//  @return DictionaryItemResponse
func (a *DictionaryItemAPIService) UpsertDictionaryItemExecute(r APIUpsertDictionaryItemRequest) (*DictionaryItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DictionaryItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DictionaryItemAPIService.UpsertDictionaryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_id"+"}", gourl.PathEscape(parameterToString(r.dictionaryID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"dictionary_item_key"+"}", gourl.PathEscape(parameterToString(r.dictionaryItemKey, "")))

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
	if r.itemKey != nil {
		localVarFormParams.Add("item_key", parameterToString(*r.itemKey, ""))
	}
	if r.itemValue != nil {
		localVarFormParams.Add("item_value", parameterToString(*r.itemValue, ""))
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
