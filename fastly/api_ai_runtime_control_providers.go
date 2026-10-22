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

// AiRuntimeControlProvidersAPI defines an interface for interacting with the resource.
type AiRuntimeControlProvidersAPI interface {

	/*
		ArcGetBedrockSetup Get Bedrock setup information

		Returns the AI Runtime Control IAM role Amazon Resource Name (ARN) and external ID to reference in your AWS
	account's IAM trust policy before creating an `aws-iam` [provider connection](https://www.fastly.com/documentation/reference/api/ai-runtime-control/provider-connections#arc-create-provider-connection)
	to Amazon Bedrock.

	Requires AWS IAM authentication to be enabled on your account. If it isn't, this
	endpoint returns a 404 response.


		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcGetBedrockSetupRequest
	*/
	ArcGetBedrockSetup(ctx context.Context) APIArcGetBedrockSetupRequest

	// ArcGetBedrockSetupExecute executes the request
	//  @return ArcBedrockSetupResponse
	ArcGetBedrockSetupExecute(r APIArcGetBedrockSetupRequest) (*ArcBedrockSetupResponse, *http.Response, error)

	/*
		ArcGetModelsForProvider Get models for a provider

		Returns the models available for a specific provider.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param providerId The ID of the provider.
		 @return APIArcGetModelsForProviderRequest
	*/
	ArcGetModelsForProvider(ctx context.Context, providerId string) APIArcGetModelsForProviderRequest

	// ArcGetModelsForProviderExecute executes the request
	//  @return ArcModelsResponse
	ArcGetModelsForProviderExecute(r APIArcGetModelsForProviderRequest) (*ArcModelsResponse, *http.Response, error)

	/*
		ArcGetSupportedProviders Get supported providers

		Returns the AI providers supported by AI Runtime Control, with each provider's available models nested within.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIArcGetSupportedProvidersRequest
	*/
	ArcGetSupportedProviders(ctx context.Context) APIArcGetSupportedProvidersRequest

	// ArcGetSupportedProvidersExecute executes the request
	//  @return ArcProvidersResponse
	ArcGetSupportedProvidersExecute(r APIArcGetSupportedProvidersRequest) (*ArcProvidersResponse, *http.Response, error)
}

// AiRuntimeControlProvidersAPIService AiRuntimeControlProvidersAPI service
type AiRuntimeControlProvidersAPIService service

// APIArcGetBedrockSetupRequest represents a request for the resource.
type APIArcGetBedrockSetupRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlProvidersAPI
}

// Execute calls the API using the request data configured.
func (r APIArcGetBedrockSetupRequest) Execute() (*ArcBedrockSetupResponse, *http.Response, error) {
	return r.APIService.ArcGetBedrockSetupExecute(r)
}

/*
ArcGetBedrockSetup Get Bedrock setup information

Returns the AI Runtime Control IAM role Amazon Resource Name (ARN) and external ID to reference in your AWS
account's IAM trust policy before creating an `aws-iam` [provider connection](https://www.fastly.com/documentation/reference/api/ai-runtime-control/provider-connections#arc-create-provider-connection)
to Amazon Bedrock.

Requires AWS IAM authentication to be enabled on your account. If it isn't, this
endpoint returns a 404 response.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcGetBedrockSetupRequest
*/
func (a *AiRuntimeControlProvidersAPIService) ArcGetBedrockSetup(ctx context.Context) APIArcGetBedrockSetupRequest {
	return APIArcGetBedrockSetupRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcGetBedrockSetupExecute executes the request
//  @return ArcBedrockSetupResponse
func (a *AiRuntimeControlProvidersAPIService) ArcGetBedrockSetupExecute(r APIArcGetBedrockSetupRequest) (*ArcBedrockSetupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcBedrockSetupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProvidersAPIService.ArcGetBedrockSetup")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/providers/bedrock/setup"

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

// APIArcGetModelsForProviderRequest represents a request for the resource.
type APIArcGetModelsForProviderRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlProvidersAPI
	providerId string
}

// Execute calls the API using the request data configured.
func (r APIArcGetModelsForProviderRequest) Execute() (*ArcModelsResponse, *http.Response, error) {
	return r.APIService.ArcGetModelsForProviderExecute(r)
}

/*
ArcGetModelsForProvider Get models for a provider

Returns the models available for a specific provider.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param providerId The ID of the provider.
 @return APIArcGetModelsForProviderRequest
*/
func (a *AiRuntimeControlProvidersAPIService) ArcGetModelsForProvider(ctx context.Context, providerId string) APIArcGetModelsForProviderRequest {
	return APIArcGetModelsForProviderRequest{
		APIService: a,
		ctx:        ctx,
		providerId: providerId,
	}
}

// ArcGetModelsForProviderExecute executes the request
//  @return ArcModelsResponse
func (a *AiRuntimeControlProvidersAPIService) ArcGetModelsForProviderExecute(r APIArcGetModelsForProviderRequest) (*ArcModelsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcModelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProvidersAPIService.ArcGetModelsForProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/providers/{provider_id}/models"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"provider_id"+"}", gourl.PathEscape(parameterToString(r.providerId, "")))

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

// APIArcGetSupportedProvidersRequest represents a request for the resource.
type APIArcGetSupportedProvidersRequest struct {
	ctx        context.Context
	APIService AiRuntimeControlProvidersAPI
}

// Execute calls the API using the request data configured.
func (r APIArcGetSupportedProvidersRequest) Execute() (*ArcProvidersResponse, *http.Response, error) {
	return r.APIService.ArcGetSupportedProvidersExecute(r)
}

/*
ArcGetSupportedProviders Get supported providers

Returns the AI providers supported by AI Runtime Control, with each provider's available models nested within.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIArcGetSupportedProvidersRequest
*/
func (a *AiRuntimeControlProvidersAPIService) ArcGetSupportedProviders(ctx context.Context) APIArcGetSupportedProvidersRequest {
	return APIArcGetSupportedProvidersRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ArcGetSupportedProvidersExecute executes the request
//  @return ArcProvidersResponse
func (a *AiRuntimeControlProvidersAPIService) ArcGetSupportedProvidersExecute(r APIArcGetSupportedProvidersRequest) (*ArcProvidersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ArcProvidersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AiRuntimeControlProvidersAPIService.ArcGetSupportedProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ai-runtime-control/v1/providers"

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
