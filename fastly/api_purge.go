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

// PurgeAPI defines an interface for interacting with the resource.
type PurgeAPI interface {

	/*
	BulkPurgeTag Purge multiple surrogate key tags

	Instant Purge a particular service of items tagged with surrogate keys. Up to 256 surrogate keys can be purged in one batch request. As an alternative to sending the keys in a JSON object in the body of the request, this endpoint also supports listing keys in a <code>Surrogate-Key</code> request header, e.g. <code>Surrogate-Key: key_1 key_2 key_3</code>.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @return APIBulkPurgeTagRequest
	*/
	BulkPurgeTag(ctx context.Context, serviceID string) APIBulkPurgeTagRequest

	// BulkPurgeTagExecute executes the request
	//  @return map[string]string
	BulkPurgeTagExecute(r APIBulkPurgeTagRequest) (map[string]string, *http.Response, error)

	/*
	PurgeAll Purge everything from a service

	Instant Purge everything from a service.

Purge-all requests cannot be done in soft mode and will always immediately invalidate all cached content associated with the service. To do a soft-purge-all, consider applying a constant [surrogate key](https://docs.fastly.com/en/guides/getting-started-with-surrogate-keys) tag (e.g., `"all"`) to all objects.


	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @return APIPurgeAllRequest
	*/
	PurgeAll(ctx context.Context, serviceID string) APIPurgeAllRequest

	// PurgeAllExecute executes the request
	//  @return InlineResponse200
	PurgeAllExecute(r APIPurgeAllRequest) (*InlineResponse200, *http.Response, error)

	/*
	PurgeSingleURL Purge a URL

	Instant Purge an individual URL.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param cachedURL URL of object in cache to be purged.
	 @return APIPurgeSingleURLRequest
	*/
	PurgeSingleURL(ctx context.Context, cachedURL string) APIPurgeSingleURLRequest

	// PurgeSingleURLExecute executes the request
	//  @return PurgeResponse
	PurgeSingleURLExecute(r APIPurgeSingleURLRequest) (*PurgeResponse, *http.Response, error)

	/*
	PurgeTag Purge by surrogate key tag

	Instant Purge a particular service of items tagged with a Surrogate Key. Only one surrogate key can be purged at a time. Multiple keys can be purged using a batch surrogate key purge request.

	 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceID Alphanumeric string identifying the service.
	 @param surrogateKey Surrogate keys are used to efficiently purge content from cache. Instead of purging your entire site or individual URLs, you can tag related assets (like all images and descriptions associated with a single product) with surrogate keys, and these grouped URLs can be purged in a single request.
	 @return APIPurgeTagRequest
	*/
	PurgeTag(ctx context.Context, serviceID string, surrogateKey string) APIPurgeTagRequest

	// PurgeTagExecute executes the request
	//  @return PurgeResponse
	PurgeTagExecute(r APIPurgeTagRequest) (*PurgeResponse, *http.Response, error)
}

// PurgeAPIService PurgeAPI service
type PurgeAPIService service

// APIBulkPurgeTagRequest represents a request for the resource.
type APIBulkPurgeTagRequest struct {
	ctx context.Context
	APIService PurgeAPI
	serviceID string
	fastlySoftPurge *int32
	surrogateKey *string
	purgeResponse *PurgeResponse
}

// FastlySoftPurge If present, this header triggers the purge to be &#39;soft&#39;, which marks the affected object as stale rather than making it inaccessible.  Typically set to \&quot;1\&quot; when used, but the value is not important.
func (r *APIBulkPurgeTagRequest) FastlySoftPurge(fastlySoftPurge int32) *APIBulkPurgeTagRequest {
	r.fastlySoftPurge = &fastlySoftPurge
	return r
}
// SurrogateKey Purge multiple surrogate key tags using a request header. Not required if a JSON POST body is specified.
func (r *APIBulkPurgeTagRequest) SurrogateKey(surrogateKey string) *APIBulkPurgeTagRequest {
	r.surrogateKey = &surrogateKey
	return r
}
// PurgeResponse returns a pointer to a request.
func (r *APIBulkPurgeTagRequest) PurgeResponse(purgeResponse PurgeResponse) *APIBulkPurgeTagRequest {
	r.purgeResponse = &purgeResponse
	return r
}

// Execute calls the API using the request data configured.
func (r APIBulkPurgeTagRequest) Execute() (map[string]string, *http.Response, error) {
	return r.APIService.BulkPurgeTagExecute(r)
}

/*
BulkPurgeTag Purge multiple surrogate key tags

Instant Purge a particular service of items tagged with surrogate keys. Up to 256 surrogate keys can be purged in one batch request. As an alternative to sending the keys in a JSON object in the body of the request, this endpoint also supports listing keys in a <code>Surrogate-Key</code> request header, e.g. <code>Surrogate-Key: key_1 key_2 key_3</code>.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIBulkPurgeTagRequest
*/
func (a *PurgeAPIService) BulkPurgeTag(ctx context.Context, serviceID string) APIBulkPurgeTagRequest {
	return APIBulkPurgeTagRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
	}
}

// BulkPurgeTagExecute executes the request
//  @return map[string]string
func (a *PurgeAPIService) BulkPurgeTagExecute(r APIBulkPurgeTagRequest) (map[string]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  map[string]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurgeAPIService.BulkPurgeTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/purge"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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
	if r.fastlySoftPurge != nil {
		localVarHeaderParams["fastly-soft-purge"] = parameterToString(*r.fastlySoftPurge, "")
	}
	if r.surrogateKey != nil {
		localVarHeaderParams["surrogate-key"] = parameterToString(*r.surrogateKey, "")
	}
	// body params
	localVarPostBody = r.purgeResponse
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

// APIPurgeAllRequest represents a request for the resource.
type APIPurgeAllRequest struct {
	ctx context.Context
	APIService PurgeAPI
	serviceID string
}


// Execute calls the API using the request data configured.
func (r APIPurgeAllRequest) Execute() (*InlineResponse200, *http.Response, error) {
	return r.APIService.PurgeAllExecute(r)
}

/*
PurgeAll Purge everything from a service

Instant Purge everything from a service.

Purge-all requests cannot be done in soft mode and will always immediately invalidate all cached content associated with the service. To do a soft-purge-all, consider applying a constant [surrogate key](https://docs.fastly.com/en/guides/getting-started-with-surrogate-keys) tag (e.g., `"all"`) to all objects.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @return APIPurgeAllRequest
*/
func (a *PurgeAPIService) PurgeAll(ctx context.Context, serviceID string) APIPurgeAllRequest {
	return APIPurgeAllRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
	}
}

// PurgeAllExecute executes the request
//  @return InlineResponse200
func (a *PurgeAPIService) PurgeAllExecute(r APIPurgeAllRequest) (*InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurgeAPIService.PurgeAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/purge_all"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))

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

// APIPurgeSingleURLRequest represents a request for the resource.
type APIPurgeSingleURLRequest struct {
	ctx context.Context
	APIService PurgeAPI
	cachedURL string
	fastlySoftPurge *int32
}

// FastlySoftPurge If present, this header triggers the purge to be &#39;soft&#39;, which marks the affected object as stale rather than making it inaccessible.  Typically set to \&quot;1\&quot; when used, but the value is not important.
func (r *APIPurgeSingleURLRequest) FastlySoftPurge(fastlySoftPurge int32) *APIPurgeSingleURLRequest {
	r.fastlySoftPurge = &fastlySoftPurge
	return r
}

// Execute calls the API using the request data configured.
func (r APIPurgeSingleURLRequest) Execute() (*PurgeResponse, *http.Response, error) {
	return r.APIService.PurgeSingleURLExecute(r)
}

/*
PurgeSingleURL Purge a URL

Instant Purge an individual URL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cachedURL URL of object in cache to be purged.
 @return APIPurgeSingleURLRequest
*/
func (a *PurgeAPIService) PurgeSingleURL(ctx context.Context, cachedURL string) APIPurgeSingleURLRequest {
	return APIPurgeSingleURLRequest{
		APIService: a,
		ctx: ctx,
		cachedURL: cachedURL,
	}
}

// PurgeSingleURLExecute executes the request
//  @return PurgeResponse
func (a *PurgeAPIService) PurgeSingleURLExecute(r APIPurgeSingleURLRequest) (*PurgeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *PurgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurgeAPIService.PurgeSingleURL")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purge/{cached_url}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cached_url"+"}", parameterToString(r.cachedURL, ""))

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
	if r.fastlySoftPurge != nil {
		localVarHeaderParams["fastly-soft-purge"] = parameterToString(*r.fastlySoftPurge, "")
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

// APIPurgeTagRequest represents a request for the resource.
type APIPurgeTagRequest struct {
	ctx context.Context
	APIService PurgeAPI
	serviceID string
	surrogateKey string
	fastlySoftPurge *int32
}

// FastlySoftPurge If present, this header triggers the purge to be &#39;soft&#39;, which marks the affected object as stale rather than making it inaccessible.  Typically set to \&quot;1\&quot; when used, but the value is not important.
func (r *APIPurgeTagRequest) FastlySoftPurge(fastlySoftPurge int32) *APIPurgeTagRequest {
	r.fastlySoftPurge = &fastlySoftPurge
	return r
}

// Execute calls the API using the request data configured.
func (r APIPurgeTagRequest) Execute() (*PurgeResponse, *http.Response, error) {
	return r.APIService.PurgeTagExecute(r)
}

/*
PurgeTag Purge by surrogate key tag

Instant Purge a particular service of items tagged with a Surrogate Key. Only one surrogate key can be purged at a time. Multiple keys can be purged using a batch surrogate key purge request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceID Alphanumeric string identifying the service.
 @param surrogateKey Surrogate keys are used to efficiently purge content from cache. Instead of purging your entire site or individual URLs, you can tag related assets (like all images and descriptions associated with a single product) with surrogate keys, and these grouped URLs can be purged in a single request.
 @return APIPurgeTagRequest
*/
func (a *PurgeAPIService) PurgeTag(ctx context.Context, serviceID string, surrogateKey string) APIPurgeTagRequest {
	return APIPurgeTagRequest{
		APIService: a,
		ctx: ctx,
		serviceID: serviceID,
		surrogateKey: surrogateKey,
	}
}

// PurgeTagExecute executes the request
//  @return PurgeResponse
func (a *PurgeAPIService) PurgeTagExecute(r APIPurgeTagRequest) (*PurgeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     any
		formFiles            []formFile
		localVarReturnValue  *PurgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurgeAPIService.PurgeTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{service_id}/purge/{surrogate_key}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"service_id"+"}", gourl.PathEscape(parameterToString(r.serviceID, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"surrogate_key"+"}", parameterToString(r.surrogateKey, ""))

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
	if r.fastlySoftPurge != nil {
		localVarHeaderParams["fastly-soft-purge"] = parameterToString(*r.fastlySoftPurge, "")
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
