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

// NetworkProtectAPI defines an interface for interacting with the resource.
type NetworkProtectAPI interface {

	/*
		CreateNetworkProtectEndpoint Create a Network Protect endpoint

		Creates a GRE delivery destination. A new endpoint is not attached to a network until it is added to an endpoint policy.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateNetworkProtectEndpointRequest
	*/
	CreateNetworkProtectEndpoint(ctx context.Context) APICreateNetworkProtectEndpointRequest

	// CreateNetworkProtectEndpointExecute executes the request
	//  @return NetworkProtectEndpointResponse
	CreateNetworkProtectEndpointExecute(r APICreateNetworkProtectEndpointRequest) (*NetworkProtectEndpointResponse, *http.Response, error)

	/*
		CreateNetworkProtectNetwork Create a Network Protect network

		Registers customer-owned IP space for protection. The network starts with advertisement disabled and Fastly does not announce its prefix. Turn advertisement on after LOA and ROA authorization are ready. The create request cannot include advertise; sending it returns a `400 Bad Request` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateNetworkProtectNetworkRequest
	*/
	CreateNetworkProtectNetwork(ctx context.Context) APICreateNetworkProtectNetworkRequest

	// CreateNetworkProtectNetworkExecute executes the request
	//  @return NetworkProtectNetworkResponse
	CreateNetworkProtectNetworkExecute(r APICreateNetworkProtectNetworkRequest) (*NetworkProtectNetworkResponse, *http.Response, error)

	/*
		CreateNetworkProtectPrefixList Create a Network Protect prefix list

		Creates a typed list of addresses that rules can reference. The family is required and cannot be changed. The addresses list may remain empty.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateNetworkProtectPrefixListRequest
	*/
	CreateNetworkProtectPrefixList(ctx context.Context) APICreateNetworkProtectPrefixListRequest

	// CreateNetworkProtectPrefixListExecute executes the request
	//  @return NetworkProtectPrefixListResponse
	CreateNetworkProtectPrefixListExecute(r APICreateNetworkProtectPrefixListRequest) (*NetworkProtectPrefixListResponse, *http.Response, error)

	/*
		CreateNetworkProtectRuleSet Create a Network Protect rule set

		Creates an ordered rule set and stores its filtering intent. A rule set has at most 64 rules and uses the default action when no rule matches. Custom attachment and detachment are restricted in this release.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateNetworkProtectRuleSetRequest
	*/
	CreateNetworkProtectRuleSet(ctx context.Context) APICreateNetworkProtectRuleSetRequest

	// CreateNetworkProtectRuleSetExecute executes the request
	//  @return NetworkProtectRuleSetResponse
	CreateNetworkProtectRuleSetExecute(r APICreateNetworkProtectRuleSetRequest) (*NetworkProtectRuleSetResponse, *http.Response, error)

	/*
		DeleteNetworkProtectEndpoint Delete a Network Protect endpoint

		Deletes an endpoint. The endpoint must first be removed from every network policy; active references return a `409 Conflict` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param endpointId Unique identifier of the Network Protect endpoint.
		 @return APIDeleteNetworkProtectEndpointRequest
	*/
	DeleteNetworkProtectEndpoint(ctx context.Context, endpointId string) APIDeleteNetworkProtectEndpointRequest

	// DeleteNetworkProtectEndpointExecute executes the request
	DeleteNetworkProtectEndpointExecute(r APIDeleteNetworkProtectEndpointRequest) (*http.Response, error)

	/*
		DeleteNetworkProtectNetwork Delete a Network Protect network

		Deletes a network and stops its advertisement and filtering. Endpoints and rule sets used by the network are not deleted.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param networkId Unique identifier of the Network Protect network.
		 @return APIDeleteNetworkProtectNetworkRequest
	*/
	DeleteNetworkProtectNetwork(ctx context.Context, networkId string) APIDeleteNetworkProtectNetworkRequest

	// DeleteNetworkProtectNetworkExecute executes the request
	DeleteNetworkProtectNetworkExecute(r APIDeleteNetworkProtectNetworkRequest) (*http.Response, error)

	/*
		DeleteNetworkProtectPrefixList Delete a Network Protect prefix list

		Deletes a prefix list. Active rule references must be removed first; referenced lists return a `409 Conflict` error. Unknown and foreign IDs are indistinguishable and return a `404 Not Found` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param prefixListId Unique identifier of the Network Protect prefix list.
		 @return APIDeleteNetworkProtectPrefixListRequest
	*/
	DeleteNetworkProtectPrefixList(ctx context.Context, prefixListId string) APIDeleteNetworkProtectPrefixListRequest

	// DeleteNetworkProtectPrefixListExecute executes the request
	DeleteNetworkProtectPrefixListExecute(r APIDeleteNetworkProtectPrefixListRequest) (*http.Response, error)

	/*
		DeleteNetworkProtectRuleSet Delete a Network Protect rule set

		Deletes a rule set. Attached rule sets cannot be deleted and return a `409 Conflict` error until their references are removed.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param ruleSetId Unique identifier of the Network Protect rule set.
		 @return APIDeleteNetworkProtectRuleSetRequest
	*/
	DeleteNetworkProtectRuleSet(ctx context.Context, ruleSetId string) APIDeleteNetworkProtectRuleSetRequest

	// DeleteNetworkProtectRuleSetExecute executes the request
	DeleteNetworkProtectRuleSetExecute(r APIDeleteNetworkProtectRuleSetRequest) (*http.Response, error)

	/*
		GetNetworkProtectEndpoint Get a Network Protect endpoint

		Retrieves a single endpoint by its identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param endpointId Unique identifier of the Network Protect endpoint.
		 @return APIGetNetworkProtectEndpointRequest
	*/
	GetNetworkProtectEndpoint(ctx context.Context, endpointId string) APIGetNetworkProtectEndpointRequest

	// GetNetworkProtectEndpointExecute executes the request
	//  @return NetworkProtectEndpointResponse
	GetNetworkProtectEndpointExecute(r APIGetNetworkProtectEndpointRequest) (*NetworkProtectEndpointResponse, *http.Response, error)

	/*
		GetNetworkProtectNetwork Get a Network Protect network

		Retrieves a single network by its identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param networkId Unique identifier of the Network Protect network.
		 @return APIGetNetworkProtectNetworkRequest
	*/
	GetNetworkProtectNetwork(ctx context.Context, networkId string) APIGetNetworkProtectNetworkRequest

	// GetNetworkProtectNetworkExecute executes the request
	//  @return NetworkProtectNetworkResponse
	GetNetworkProtectNetworkExecute(r APIGetNetworkProtectNetworkRequest) (*NetworkProtectNetworkResponse, *http.Response, error)

	/*
		GetNetworkProtectPrefixList Get a Network Protect prefix list

		Retrieves a single prefix list by its identifier. Unknown and foreign IDs are indistinguishable and return a `404 Not Found` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param prefixListId Unique identifier of the Network Protect prefix list.
		 @return APIGetNetworkProtectPrefixListRequest
	*/
	GetNetworkProtectPrefixList(ctx context.Context, prefixListId string) APIGetNetworkProtectPrefixListRequest

	// GetNetworkProtectPrefixListExecute executes the request
	//  @return NetworkProtectPrefixListResponse
	GetNetworkProtectPrefixListExecute(r APIGetNetworkProtectPrefixListRequest) (*NetworkProtectPrefixListResponse, *http.Response, error)

	/*
		GetNetworkProtectRuleSet Get a Network Protect rule set

		Retrieves a single rule set with rules in evaluation order.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param ruleSetId Unique identifier of the Network Protect rule set.
		 @return APIGetNetworkProtectRuleSetRequest
	*/
	GetNetworkProtectRuleSet(ctx context.Context, ruleSetId string) APIGetNetworkProtectRuleSetRequest

	// GetNetworkProtectRuleSetExecute executes the request
	//  @return NetworkProtectRuleSetResponse
	GetNetworkProtectRuleSetExecute(r APIGetNetworkProtectRuleSetRequest) (*NetworkProtectRuleSetResponse, *http.Response, error)

	/*
		ListNetworkProtectEndpoints List Network Protect endpoints

		Lists the customer's GRE delivery endpoints in ascending ID order. The response is cursor-paginated and includes the effective page limit.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListNetworkProtectEndpointsRequest
	*/
	ListNetworkProtectEndpoints(ctx context.Context) APIListNetworkProtectEndpointsRequest

	// ListNetworkProtectEndpointsExecute executes the request
	//  @return NetworkProtectEndpointsResponse
	ListNetworkProtectEndpointsExecute(r APIListNetworkProtectEndpointsRequest) (*NetworkProtectEndpointsResponse, *http.Response, error)

	/*
		ListNetworkProtectNetworks List Network Protect networks

		Lists the customer's networks in ascending ID order. The response is cursor-paginated and includes the effective page limit.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListNetworkProtectNetworksRequest
	*/
	ListNetworkProtectNetworks(ctx context.Context) APIListNetworkProtectNetworksRequest

	// ListNetworkProtectNetworksExecute executes the request
	//  @return NetworkProtectNetworksResponse
	ListNetworkProtectNetworksExecute(r APIListNetworkProtectNetworksRequest) (*NetworkProtectNetworksResponse, *http.Response, error)

	/*
		ListNetworkProtectPrefixLists List Network Protect prefix lists

		Lists prefix lists in ascending ID order. Positive limits above 1000 are accepted and clamped to 1000. Invalid limits return a `400 Bad Request` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListNetworkProtectPrefixListsRequest
	*/
	ListNetworkProtectPrefixLists(ctx context.Context) APIListNetworkProtectPrefixListsRequest

	// ListNetworkProtectPrefixListsExecute executes the request
	//  @return NetworkProtectPrefixListsResponse
	ListNetworkProtectPrefixListsExecute(r APIListNetworkProtectPrefixListsRequest) (*NetworkProtectPrefixListsResponse, *http.Response, error)

	/*
		ListNetworkProtectRuleSets List Network Protect rule sets

		Lists rule sets in ascending ID order. Positive limits above 1000 are accepted and clamped to 1000; invalid limits return a `400 Bad Request` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListNetworkProtectRuleSetsRequest
	*/
	ListNetworkProtectRuleSets(ctx context.Context) APIListNetworkProtectRuleSetsRequest

	// ListNetworkProtectRuleSetsExecute executes the request
	//  @return NetworkProtectRuleSetsResponse
	ListNetworkProtectRuleSetsExecute(r APIListNetworkProtectRuleSetsRequest) (*NetworkProtectRuleSetsResponse, *http.Response, error)

	/*
		UpdateNetworkProtectEndpoint Update a Network Protect endpoint

		Updates an endpoint. Omitted fields remain unchanged. The IP address cannot be changed; create a new endpoint if the destination changes.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param endpointId Unique identifier of the Network Protect endpoint.
		 @return APIUpdateNetworkProtectEndpointRequest
	*/
	UpdateNetworkProtectEndpoint(ctx context.Context, endpointId string) APIUpdateNetworkProtectEndpointRequest

	// UpdateNetworkProtectEndpointExecute executes the request
	//  @return NetworkProtectEndpointResponse
	UpdateNetworkProtectEndpointExecute(r APIUpdateNetworkProtectEndpointRequest) (*NetworkProtectEndpointResponse, *http.Response, error)

	/*
		UpdateNetworkProtectNetwork Update a Network Protect network

		Updates a network. Omitted fields remain unchanged. The prefix cannot be changed. Sending `endpoint_policies` replaces the complete policy list; an empty array removes every policy.
	Advertisement is desired state. Enabling it requires an approved LOA, a valid ROA, and a ready default-region delivery policy. A readiness failure returns a `409 Conflict` error.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param networkId Unique identifier of the Network Protect network.
		 @return APIUpdateNetworkProtectNetworkRequest
	*/
	UpdateNetworkProtectNetwork(ctx context.Context, networkId string) APIUpdateNetworkProtectNetworkRequest

	// UpdateNetworkProtectNetworkExecute executes the request
	//  @return NetworkProtectNetworkResponse
	UpdateNetworkProtectNetworkExecute(r APIUpdateNetworkProtectNetworkRequest) (*NetworkProtectNetworkResponse, *http.Response, error)

	/*
		UpdateNetworkProtectPrefixList Update a Network Protect prefix list

		Updates a prefix list. The family cannot be changed. Sending addresses replaces the complete list; an empty array clears it when no active source-list restriction applies. Addresses are canonicalized, deduplicated, and sorted by the backend.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param prefixListId Unique identifier of the Network Protect prefix list.
		 @return APIUpdateNetworkProtectPrefixListRequest
	*/
	UpdateNetworkProtectPrefixList(ctx context.Context, prefixListId string) APIUpdateNetworkProtectPrefixListRequest

	// UpdateNetworkProtectPrefixListExecute executes the request
	//  @return NetworkProtectPrefixListResponse
	UpdateNetworkProtectPrefixListExecute(r APIUpdateNetworkProtectPrefixListRequest) (*NetworkProtectPrefixListResponse, *http.Response, error)

	/*
		UpdateNetworkProtectRuleSet Update a Network Protect rule set

		Updates a rule set. Omitted fields remain unchanged. Sending rules replaces the complete ordered list; an empty array removes every rule. Custom attachment and detachment remain restricted in this release.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param ruleSetId Unique identifier of the Network Protect rule set.
		 @return APIUpdateNetworkProtectRuleSetRequest
	*/
	UpdateNetworkProtectRuleSet(ctx context.Context, ruleSetId string) APIUpdateNetworkProtectRuleSetRequest

	// UpdateNetworkProtectRuleSetExecute executes the request
	//  @return NetworkProtectRuleSetResponse
	UpdateNetworkProtectRuleSetExecute(r APIUpdateNetworkProtectRuleSetRequest) (*NetworkProtectRuleSetResponse, *http.Response, error)
}

// NetworkProtectAPIService NetworkProtectAPI service
type NetworkProtectAPIService service

// APICreateNetworkProtectEndpointRequest represents a request for the resource.
type APICreateNetworkProtectEndpointRequest struct {
	ctx                          context.Context
	APIService                   NetworkProtectAPI
	networkProtectEndpointCreate *NetworkProtectEndpointCreate
}

// NetworkProtectEndpointCreate returns a pointer to a request.
func (r *APICreateNetworkProtectEndpointRequest) NetworkProtectEndpointCreate(networkProtectEndpointCreate NetworkProtectEndpointCreate) *APICreateNetworkProtectEndpointRequest {
	r.networkProtectEndpointCreate = &networkProtectEndpointCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateNetworkProtectEndpointRequest) Execute() (*NetworkProtectEndpointResponse, *http.Response, error) {
	return r.APIService.CreateNetworkProtectEndpointExecute(r)
}

/*
CreateNetworkProtectEndpoint Create a Network Protect endpoint

Creates a GRE delivery destination. A new endpoint is not attached to a network until it is added to an endpoint policy.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateNetworkProtectEndpointRequest
*/
func (a *NetworkProtectAPIService) CreateNetworkProtectEndpoint(ctx context.Context) APICreateNetworkProtectEndpointRequest {
	return APICreateNetworkProtectEndpointRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateNetworkProtectEndpointExecute executes the request
//  @return NetworkProtectEndpointResponse
func (a *NetworkProtectAPIService) CreateNetworkProtectEndpointExecute(r APICreateNetworkProtectEndpointRequest) (*NetworkProtectEndpointResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectEndpointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.CreateNetworkProtectEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/endpoints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectEndpointCreate == nil {
		return localVarReturnValue, nil, reportError("networkProtectEndpointCreate is required and must be specified")
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
	localVarPostBody = r.networkProtectEndpointCreate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APICreateNetworkProtectNetworkRequest represents a request for the resource.
type APICreateNetworkProtectNetworkRequest struct {
	ctx                         context.Context
	APIService                  NetworkProtectAPI
	networkProtectNetworkCreate *NetworkProtectNetworkCreate
}

// NetworkProtectNetworkCreate returns a pointer to a request.
func (r *APICreateNetworkProtectNetworkRequest) NetworkProtectNetworkCreate(networkProtectNetworkCreate NetworkProtectNetworkCreate) *APICreateNetworkProtectNetworkRequest {
	r.networkProtectNetworkCreate = &networkProtectNetworkCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateNetworkProtectNetworkRequest) Execute() (*NetworkProtectNetworkResponse, *http.Response, error) {
	return r.APIService.CreateNetworkProtectNetworkExecute(r)
}

/*
CreateNetworkProtectNetwork Create a Network Protect network

Registers customer-owned IP space for protection. The network starts with advertisement disabled and Fastly does not announce its prefix. Turn advertisement on after LOA and ROA authorization are ready. The create request cannot include advertise; sending it returns a `400 Bad Request` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateNetworkProtectNetworkRequest
*/
func (a *NetworkProtectAPIService) CreateNetworkProtectNetwork(ctx context.Context) APICreateNetworkProtectNetworkRequest {
	return APICreateNetworkProtectNetworkRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateNetworkProtectNetworkExecute executes the request
//  @return NetworkProtectNetworkResponse
func (a *NetworkProtectAPIService) CreateNetworkProtectNetworkExecute(r APICreateNetworkProtectNetworkRequest) (*NetworkProtectNetworkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.CreateNetworkProtectNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectNetworkCreate == nil {
		return localVarReturnValue, nil, reportError("networkProtectNetworkCreate is required and must be specified")
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
	localVarPostBody = r.networkProtectNetworkCreate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APICreateNetworkProtectPrefixListRequest represents a request for the resource.
type APICreateNetworkProtectPrefixListRequest struct {
	ctx                            context.Context
	APIService                     NetworkProtectAPI
	networkProtectPrefixListCreate *NetworkProtectPrefixListCreate
}

// NetworkProtectPrefixListCreate returns a pointer to a request.
func (r *APICreateNetworkProtectPrefixListRequest) NetworkProtectPrefixListCreate(networkProtectPrefixListCreate NetworkProtectPrefixListCreate) *APICreateNetworkProtectPrefixListRequest {
	r.networkProtectPrefixListCreate = &networkProtectPrefixListCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateNetworkProtectPrefixListRequest) Execute() (*NetworkProtectPrefixListResponse, *http.Response, error) {
	return r.APIService.CreateNetworkProtectPrefixListExecute(r)
}

/*
CreateNetworkProtectPrefixList Create a Network Protect prefix list

Creates a typed list of addresses that rules can reference. The family is required and cannot be changed. The addresses list may remain empty.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateNetworkProtectPrefixListRequest
*/
func (a *NetworkProtectAPIService) CreateNetworkProtectPrefixList(ctx context.Context) APICreateNetworkProtectPrefixListRequest {
	return APICreateNetworkProtectPrefixListRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateNetworkProtectPrefixListExecute executes the request
//  @return NetworkProtectPrefixListResponse
func (a *NetworkProtectAPIService) CreateNetworkProtectPrefixListExecute(r APICreateNetworkProtectPrefixListRequest) (*NetworkProtectPrefixListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectPrefixListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.CreateNetworkProtectPrefixList")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/prefix-lists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectPrefixListCreate == nil {
		return localVarReturnValue, nil, reportError("networkProtectPrefixListCreate is required and must be specified")
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
	localVarPostBody = r.networkProtectPrefixListCreate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APICreateNetworkProtectRuleSetRequest represents a request for the resource.
type APICreateNetworkProtectRuleSetRequest struct {
	ctx                         context.Context
	APIService                  NetworkProtectAPI
	networkProtectRuleSetCreate *NetworkProtectRuleSetCreate
}

// NetworkProtectRuleSetCreate returns a pointer to a request.
func (r *APICreateNetworkProtectRuleSetRequest) NetworkProtectRuleSetCreate(networkProtectRuleSetCreate NetworkProtectRuleSetCreate) *APICreateNetworkProtectRuleSetRequest {
	r.networkProtectRuleSetCreate = &networkProtectRuleSetCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateNetworkProtectRuleSetRequest) Execute() (*NetworkProtectRuleSetResponse, *http.Response, error) {
	return r.APIService.CreateNetworkProtectRuleSetExecute(r)
}

/*
CreateNetworkProtectRuleSet Create a Network Protect rule set

Creates an ordered rule set and stores its filtering intent. A rule set has at most 64 rules and uses the default action when no rule matches. Custom attachment and detachment are restricted in this release.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateNetworkProtectRuleSetRequest
*/
func (a *NetworkProtectAPIService) CreateNetworkProtectRuleSet(ctx context.Context) APICreateNetworkProtectRuleSetRequest {
	return APICreateNetworkProtectRuleSetRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateNetworkProtectRuleSetExecute executes the request
//  @return NetworkProtectRuleSetResponse
func (a *NetworkProtectAPIService) CreateNetworkProtectRuleSetExecute(r APICreateNetworkProtectRuleSetRequest) (*NetworkProtectRuleSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectRuleSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.CreateNetworkProtectRuleSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/rule-sets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectRuleSetCreate == nil {
		return localVarReturnValue, nil, reportError("networkProtectRuleSetCreate is required and must be specified")
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
	localVarPostBody = r.networkProtectRuleSetCreate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APIDeleteNetworkProtectEndpointRequest represents a request for the resource.
type APIDeleteNetworkProtectEndpointRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	endpointId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteNetworkProtectEndpointRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteNetworkProtectEndpointExecute(r)
}

/*
DeleteNetworkProtectEndpoint Delete a Network Protect endpoint

Deletes an endpoint. The endpoint must first be removed from every network policy; active references return a `409 Conflict` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param endpointId Unique identifier of the Network Protect endpoint.
 @return APIDeleteNetworkProtectEndpointRequest
*/
func (a *NetworkProtectAPIService) DeleteNetworkProtectEndpoint(ctx context.Context, endpointId string) APIDeleteNetworkProtectEndpointRequest {
	return APIDeleteNetworkProtectEndpointRequest{
		APIService: a,
		ctx:        ctx,
		endpointId: endpointId,
	}
}

// DeleteNetworkProtectEndpointExecute executes the request
func (a *NetworkProtectAPIService) DeleteNetworkProtectEndpointExecute(r APIDeleteNetworkProtectEndpointRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.DeleteNetworkProtectEndpoint")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/endpoints/{endpoint_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"endpoint_id"+"}", gourl.PathEscape(parameterToString(r.endpointId, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
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

	return localVarHTTPResponse, nil
}

// APIDeleteNetworkProtectNetworkRequest represents a request for the resource.
type APIDeleteNetworkProtectNetworkRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	networkId  string
}

// Execute calls the API using the request data configured.
func (r APIDeleteNetworkProtectNetworkRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteNetworkProtectNetworkExecute(r)
}

/*
DeleteNetworkProtectNetwork Delete a Network Protect network

Deletes a network and stops its advertisement and filtering. Endpoints and rule sets used by the network are not deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Unique identifier of the Network Protect network.
 @return APIDeleteNetworkProtectNetworkRequest
*/
func (a *NetworkProtectAPIService) DeleteNetworkProtectNetwork(ctx context.Context, networkId string) APIDeleteNetworkProtectNetworkRequest {
	return APIDeleteNetworkProtectNetworkRequest{
		APIService: a,
		ctx:        ctx,
		networkId:  networkId,
	}
}

// DeleteNetworkProtectNetworkExecute executes the request
func (a *NetworkProtectAPIService) DeleteNetworkProtectNetworkExecute(r APIDeleteNetworkProtectNetworkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.DeleteNetworkProtectNetwork")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/networks/{network_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"network_id"+"}", gourl.PathEscape(parameterToString(r.networkId, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
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

	return localVarHTTPResponse, nil
}

// APIDeleteNetworkProtectPrefixListRequest represents a request for the resource.
type APIDeleteNetworkProtectPrefixListRequest struct {
	ctx          context.Context
	APIService   NetworkProtectAPI
	prefixListId string
}

// Execute calls the API using the request data configured.
func (r APIDeleteNetworkProtectPrefixListRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteNetworkProtectPrefixListExecute(r)
}

/*
DeleteNetworkProtectPrefixList Delete a Network Protect prefix list

Deletes a prefix list. Active rule references must be removed first; referenced lists return a `409 Conflict` error. Unknown and foreign IDs are indistinguishable and return a `404 Not Found` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param prefixListId Unique identifier of the Network Protect prefix list.
 @return APIDeleteNetworkProtectPrefixListRequest
*/
func (a *NetworkProtectAPIService) DeleteNetworkProtectPrefixList(ctx context.Context, prefixListId string) APIDeleteNetworkProtectPrefixListRequest {
	return APIDeleteNetworkProtectPrefixListRequest{
		APIService:   a,
		ctx:          ctx,
		prefixListId: prefixListId,
	}
}

// DeleteNetworkProtectPrefixListExecute executes the request
func (a *NetworkProtectAPIService) DeleteNetworkProtectPrefixListExecute(r APIDeleteNetworkProtectPrefixListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.DeleteNetworkProtectPrefixList")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/prefix-lists/{prefix_list_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"prefix_list_id"+"}", gourl.PathEscape(parameterToString(r.prefixListId, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
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

	return localVarHTTPResponse, nil
}

// APIDeleteNetworkProtectRuleSetRequest represents a request for the resource.
type APIDeleteNetworkProtectRuleSetRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	ruleSetId  string
}

// Execute calls the API using the request data configured.
func (r APIDeleteNetworkProtectRuleSetRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteNetworkProtectRuleSetExecute(r)
}

/*
DeleteNetworkProtectRuleSet Delete a Network Protect rule set

Deletes a rule set. Attached rule sets cannot be deleted and return a `409 Conflict` error until their references are removed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleSetId Unique identifier of the Network Protect rule set.
 @return APIDeleteNetworkProtectRuleSetRequest
*/
func (a *NetworkProtectAPIService) DeleteNetworkProtectRuleSet(ctx context.Context, ruleSetId string) APIDeleteNetworkProtectRuleSetRequest {
	return APIDeleteNetworkProtectRuleSetRequest{
		APIService: a,
		ctx:        ctx,
		ruleSetId:  ruleSetId,
	}
}

// DeleteNetworkProtectRuleSetExecute executes the request
func (a *NetworkProtectAPIService) DeleteNetworkProtectRuleSetExecute(r APIDeleteNetworkProtectRuleSetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.DeleteNetworkProtectRuleSet")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/rule-sets/{rule_set_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rule_set_id"+"}", gourl.PathEscape(parameterToString(r.ruleSetId, "")))

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
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

	return localVarHTTPResponse, nil
}

// APIGetNetworkProtectEndpointRequest represents a request for the resource.
type APIGetNetworkProtectEndpointRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	endpointId string
}

// Execute calls the API using the request data configured.
func (r APIGetNetworkProtectEndpointRequest) Execute() (*NetworkProtectEndpointResponse, *http.Response, error) {
	return r.APIService.GetNetworkProtectEndpointExecute(r)
}

/*
GetNetworkProtectEndpoint Get a Network Protect endpoint

Retrieves a single endpoint by its identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param endpointId Unique identifier of the Network Protect endpoint.
 @return APIGetNetworkProtectEndpointRequest
*/
func (a *NetworkProtectAPIService) GetNetworkProtectEndpoint(ctx context.Context, endpointId string) APIGetNetworkProtectEndpointRequest {
	return APIGetNetworkProtectEndpointRequest{
		APIService: a,
		ctx:        ctx,
		endpointId: endpointId,
	}
}

// GetNetworkProtectEndpointExecute executes the request
//  @return NetworkProtectEndpointResponse
func (a *NetworkProtectAPIService) GetNetworkProtectEndpointExecute(r APIGetNetworkProtectEndpointRequest) (*NetworkProtectEndpointResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectEndpointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.GetNetworkProtectEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/endpoints/{endpoint_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"endpoint_id"+"}", gourl.PathEscape(parameterToString(r.endpointId, "")))

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
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIGetNetworkProtectNetworkRequest represents a request for the resource.
type APIGetNetworkProtectNetworkRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	networkId  string
}

// Execute calls the API using the request data configured.
func (r APIGetNetworkProtectNetworkRequest) Execute() (*NetworkProtectNetworkResponse, *http.Response, error) {
	return r.APIService.GetNetworkProtectNetworkExecute(r)
}

/*
GetNetworkProtectNetwork Get a Network Protect network

Retrieves a single network by its identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Unique identifier of the Network Protect network.
 @return APIGetNetworkProtectNetworkRequest
*/
func (a *NetworkProtectAPIService) GetNetworkProtectNetwork(ctx context.Context, networkId string) APIGetNetworkProtectNetworkRequest {
	return APIGetNetworkProtectNetworkRequest{
		APIService: a,
		ctx:        ctx,
		networkId:  networkId,
	}
}

// GetNetworkProtectNetworkExecute executes the request
//  @return NetworkProtectNetworkResponse
func (a *NetworkProtectAPIService) GetNetworkProtectNetworkExecute(r APIGetNetworkProtectNetworkRequest) (*NetworkProtectNetworkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.GetNetworkProtectNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/networks/{network_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"network_id"+"}", gourl.PathEscape(parameterToString(r.networkId, "")))

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
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIGetNetworkProtectPrefixListRequest represents a request for the resource.
type APIGetNetworkProtectPrefixListRequest struct {
	ctx          context.Context
	APIService   NetworkProtectAPI
	prefixListId string
}

// Execute calls the API using the request data configured.
func (r APIGetNetworkProtectPrefixListRequest) Execute() (*NetworkProtectPrefixListResponse, *http.Response, error) {
	return r.APIService.GetNetworkProtectPrefixListExecute(r)
}

/*
GetNetworkProtectPrefixList Get a Network Protect prefix list

Retrieves a single prefix list by its identifier. Unknown and foreign IDs are indistinguishable and return a `404 Not Found` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param prefixListId Unique identifier of the Network Protect prefix list.
 @return APIGetNetworkProtectPrefixListRequest
*/
func (a *NetworkProtectAPIService) GetNetworkProtectPrefixList(ctx context.Context, prefixListId string) APIGetNetworkProtectPrefixListRequest {
	return APIGetNetworkProtectPrefixListRequest{
		APIService:   a,
		ctx:          ctx,
		prefixListId: prefixListId,
	}
}

// GetNetworkProtectPrefixListExecute executes the request
//  @return NetworkProtectPrefixListResponse
func (a *NetworkProtectAPIService) GetNetworkProtectPrefixListExecute(r APIGetNetworkProtectPrefixListRequest) (*NetworkProtectPrefixListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectPrefixListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.GetNetworkProtectPrefixList")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/prefix-lists/{prefix_list_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"prefix_list_id"+"}", gourl.PathEscape(parameterToString(r.prefixListId, "")))

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
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIGetNetworkProtectRuleSetRequest represents a request for the resource.
type APIGetNetworkProtectRuleSetRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	ruleSetId  string
}

// Execute calls the API using the request data configured.
func (r APIGetNetworkProtectRuleSetRequest) Execute() (*NetworkProtectRuleSetResponse, *http.Response, error) {
	return r.APIService.GetNetworkProtectRuleSetExecute(r)
}

/*
GetNetworkProtectRuleSet Get a Network Protect rule set

Retrieves a single rule set with rules in evaluation order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleSetId Unique identifier of the Network Protect rule set.
 @return APIGetNetworkProtectRuleSetRequest
*/
func (a *NetworkProtectAPIService) GetNetworkProtectRuleSet(ctx context.Context, ruleSetId string) APIGetNetworkProtectRuleSetRequest {
	return APIGetNetworkProtectRuleSetRequest{
		APIService: a,
		ctx:        ctx,
		ruleSetId:  ruleSetId,
	}
}

// GetNetworkProtectRuleSetExecute executes the request
//  @return NetworkProtectRuleSetResponse
func (a *NetworkProtectAPIService) GetNetworkProtectRuleSetExecute(r APIGetNetworkProtectRuleSetRequest) (*NetworkProtectRuleSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectRuleSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.GetNetworkProtectRuleSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/rule-sets/{rule_set_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rule_set_id"+"}", gourl.PathEscape(parameterToString(r.ruleSetId, "")))

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
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIListNetworkProtectEndpointsRequest represents a request for the resource.
type APIListNetworkProtectEndpointsRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	limit      *int64
	cursor     *string
}

// Limit Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a &#x60;400 Bad Request&#x60; error.
func (r *APIListNetworkProtectEndpointsRequest) Limit(limit int64) *APIListNetworkProtectEndpointsRequest {
	r.limit = &limit
	return r
}

// Cursor Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page.
func (r *APIListNetworkProtectEndpointsRequest) Cursor(cursor string) *APIListNetworkProtectEndpointsRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIListNetworkProtectEndpointsRequest) Execute() (*NetworkProtectEndpointsResponse, *http.Response, error) {
	return r.APIService.ListNetworkProtectEndpointsExecute(r)
}

/*
ListNetworkProtectEndpoints List Network Protect endpoints

Lists the customer's GRE delivery endpoints in ascending ID order. The response is cursor-paginated and includes the effective page limit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListNetworkProtectEndpointsRequest
*/
func (a *NetworkProtectAPIService) ListNetworkProtectEndpoints(ctx context.Context) APIListNetworkProtectEndpointsRequest {
	return APIListNetworkProtectEndpointsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListNetworkProtectEndpointsExecute executes the request
//  @return NetworkProtectEndpointsResponse
func (a *NetworkProtectAPIService) ListNetworkProtectEndpointsExecute(r APIListNetworkProtectEndpointsRequest) (*NetworkProtectEndpointsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectEndpointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.ListNetworkProtectEndpoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/endpoints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIListNetworkProtectNetworksRequest represents a request for the resource.
type APIListNetworkProtectNetworksRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	limit      *int64
	cursor     *string
}

// Limit Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a &#x60;400 Bad Request&#x60; error.
func (r *APIListNetworkProtectNetworksRequest) Limit(limit int64) *APIListNetworkProtectNetworksRequest {
	r.limit = &limit
	return r
}

// Cursor Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page.
func (r *APIListNetworkProtectNetworksRequest) Cursor(cursor string) *APIListNetworkProtectNetworksRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIListNetworkProtectNetworksRequest) Execute() (*NetworkProtectNetworksResponse, *http.Response, error) {
	return r.APIService.ListNetworkProtectNetworksExecute(r)
}

/*
ListNetworkProtectNetworks List Network Protect networks

Lists the customer's networks in ascending ID order. The response is cursor-paginated and includes the effective page limit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListNetworkProtectNetworksRequest
*/
func (a *NetworkProtectAPIService) ListNetworkProtectNetworks(ctx context.Context) APIListNetworkProtectNetworksRequest {
	return APIListNetworkProtectNetworksRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListNetworkProtectNetworksExecute executes the request
//  @return NetworkProtectNetworksResponse
func (a *NetworkProtectAPIService) ListNetworkProtectNetworksExecute(r APIListNetworkProtectNetworksRequest) (*NetworkProtectNetworksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectNetworksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.ListNetworkProtectNetworks")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIListNetworkProtectPrefixListsRequest represents a request for the resource.
type APIListNetworkProtectPrefixListsRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	limit      *int64
	cursor     *string
}

// Limit Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a &#x60;400 Bad Request&#x60; error.
func (r *APIListNetworkProtectPrefixListsRequest) Limit(limit int64) *APIListNetworkProtectPrefixListsRequest {
	r.limit = &limit
	return r
}

// Cursor Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page.
func (r *APIListNetworkProtectPrefixListsRequest) Cursor(cursor string) *APIListNetworkProtectPrefixListsRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIListNetworkProtectPrefixListsRequest) Execute() (*NetworkProtectPrefixListsResponse, *http.Response, error) {
	return r.APIService.ListNetworkProtectPrefixListsExecute(r)
}

/*
ListNetworkProtectPrefixLists List Network Protect prefix lists

Lists prefix lists in ascending ID order. Positive limits above 1000 are accepted and clamped to 1000. Invalid limits return a `400 Bad Request` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListNetworkProtectPrefixListsRequest
*/
func (a *NetworkProtectAPIService) ListNetworkProtectPrefixLists(ctx context.Context) APIListNetworkProtectPrefixListsRequest {
	return APIListNetworkProtectPrefixListsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListNetworkProtectPrefixListsExecute executes the request
//  @return NetworkProtectPrefixListsResponse
func (a *NetworkProtectAPIService) ListNetworkProtectPrefixListsExecute(r APIListNetworkProtectPrefixListsRequest) (*NetworkProtectPrefixListsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectPrefixListsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.ListNetworkProtectPrefixLists")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/prefix-lists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIListNetworkProtectRuleSetsRequest represents a request for the resource.
type APIListNetworkProtectRuleSetsRequest struct {
	ctx        context.Context
	APIService NetworkProtectAPI
	limit      *int64
	cursor     *string
}

// Limit Maximum resources to return. Defaults to 100. Positive values above 1000 are accepted and clamped to 1000; the effective value is reported in meta.limit. Zero, negative, and non-integer values return a &#x60;400 Bad Request&#x60; error.
func (r *APIListNetworkProtectRuleSetsRequest) Limit(limit int64) *APIListNetworkProtectRuleSetsRequest {
	r.limit = &limit
	return r
}

// Cursor Continuation value from a previous response&#39;s meta.next_cursor. Omit this parameter to request the first page.
func (r *APIListNetworkProtectRuleSetsRequest) Cursor(cursor string) *APIListNetworkProtectRuleSetsRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIListNetworkProtectRuleSetsRequest) Execute() (*NetworkProtectRuleSetsResponse, *http.Response, error) {
	return r.APIService.ListNetworkProtectRuleSetsExecute(r)
}

/*
ListNetworkProtectRuleSets List Network Protect rule sets

Lists rule sets in ascending ID order. Positive limits above 1000 are accepted and clamped to 1000; invalid limits return a `400 Bad Request` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListNetworkProtectRuleSetsRequest
*/
func (a *NetworkProtectAPIService) ListNetworkProtectRuleSets(ctx context.Context) APIListNetworkProtectRuleSetsRequest {
	return APIListNetworkProtectRuleSetsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListNetworkProtectRuleSetsExecute executes the request
//  @return NetworkProtectRuleSetsResponse
func (a *NetworkProtectAPIService) ListNetworkProtectRuleSetsExecute(r APIListNetworkProtectRuleSetsRequest) (*NetworkProtectRuleSetsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectRuleSetsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.ListNetworkProtectRuleSets")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/rule-sets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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

// APIUpdateNetworkProtectEndpointRequest represents a request for the resource.
type APIUpdateNetworkProtectEndpointRequest struct {
	ctx                          context.Context
	APIService                   NetworkProtectAPI
	endpointId                   string
	networkProtectEndpointUpdate *NetworkProtectEndpointUpdate
}

// NetworkProtectEndpointUpdate returns a pointer to a request.
func (r *APIUpdateNetworkProtectEndpointRequest) NetworkProtectEndpointUpdate(networkProtectEndpointUpdate NetworkProtectEndpointUpdate) *APIUpdateNetworkProtectEndpointRequest {
	r.networkProtectEndpointUpdate = &networkProtectEndpointUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateNetworkProtectEndpointRequest) Execute() (*NetworkProtectEndpointResponse, *http.Response, error) {
	return r.APIService.UpdateNetworkProtectEndpointExecute(r)
}

/*
UpdateNetworkProtectEndpoint Update a Network Protect endpoint

Updates an endpoint. Omitted fields remain unchanged. The IP address cannot be changed; create a new endpoint if the destination changes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param endpointId Unique identifier of the Network Protect endpoint.
 @return APIUpdateNetworkProtectEndpointRequest
*/
func (a *NetworkProtectAPIService) UpdateNetworkProtectEndpoint(ctx context.Context, endpointId string) APIUpdateNetworkProtectEndpointRequest {
	return APIUpdateNetworkProtectEndpointRequest{
		APIService: a,
		ctx:        ctx,
		endpointId: endpointId,
	}
}

// UpdateNetworkProtectEndpointExecute executes the request
//  @return NetworkProtectEndpointResponse
func (a *NetworkProtectAPIService) UpdateNetworkProtectEndpointExecute(r APIUpdateNetworkProtectEndpointRequest) (*NetworkProtectEndpointResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectEndpointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.UpdateNetworkProtectEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/endpoints/{endpoint_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"endpoint_id"+"}", gourl.PathEscape(parameterToString(r.endpointId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectEndpointUpdate == nil {
		return localVarReturnValue, nil, reportError("networkProtectEndpointUpdate is required and must be specified")
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
	localVarPostBody = r.networkProtectEndpointUpdate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APIUpdateNetworkProtectNetworkRequest represents a request for the resource.
type APIUpdateNetworkProtectNetworkRequest struct {
	ctx                         context.Context
	APIService                  NetworkProtectAPI
	networkId                   string
	networkProtectNetworkUpdate *NetworkProtectNetworkUpdate
}

// NetworkProtectNetworkUpdate returns a pointer to a request.
func (r *APIUpdateNetworkProtectNetworkRequest) NetworkProtectNetworkUpdate(networkProtectNetworkUpdate NetworkProtectNetworkUpdate) *APIUpdateNetworkProtectNetworkRequest {
	r.networkProtectNetworkUpdate = &networkProtectNetworkUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateNetworkProtectNetworkRequest) Execute() (*NetworkProtectNetworkResponse, *http.Response, error) {
	return r.APIService.UpdateNetworkProtectNetworkExecute(r)
}

/*
UpdateNetworkProtectNetwork Update a Network Protect network

Updates a network. Omitted fields remain unchanged. The prefix cannot be changed. Sending `endpoint_policies` replaces the complete policy list; an empty array removes every policy.
Advertisement is desired state. Enabling it requires an approved LOA, a valid ROA, and a ready default-region delivery policy. A readiness failure returns a `409 Conflict` error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Unique identifier of the Network Protect network.
 @return APIUpdateNetworkProtectNetworkRequest
*/
func (a *NetworkProtectAPIService) UpdateNetworkProtectNetwork(ctx context.Context, networkId string) APIUpdateNetworkProtectNetworkRequest {
	return APIUpdateNetworkProtectNetworkRequest{
		APIService: a,
		ctx:        ctx,
		networkId:  networkId,
	}
}

// UpdateNetworkProtectNetworkExecute executes the request
//  @return NetworkProtectNetworkResponse
func (a *NetworkProtectAPIService) UpdateNetworkProtectNetworkExecute(r APIUpdateNetworkProtectNetworkRequest) (*NetworkProtectNetworkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.UpdateNetworkProtectNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/networks/{network_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"network_id"+"}", gourl.PathEscape(parameterToString(r.networkId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectNetworkUpdate == nil {
		return localVarReturnValue, nil, reportError("networkProtectNetworkUpdate is required and must be specified")
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
	localVarPostBody = r.networkProtectNetworkUpdate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APIUpdateNetworkProtectPrefixListRequest represents a request for the resource.
type APIUpdateNetworkProtectPrefixListRequest struct {
	ctx                            context.Context
	APIService                     NetworkProtectAPI
	prefixListId                   string
	networkProtectPrefixListUpdate *NetworkProtectPrefixListUpdate
}

// NetworkProtectPrefixListUpdate returns a pointer to a request.
func (r *APIUpdateNetworkProtectPrefixListRequest) NetworkProtectPrefixListUpdate(networkProtectPrefixListUpdate NetworkProtectPrefixListUpdate) *APIUpdateNetworkProtectPrefixListRequest {
	r.networkProtectPrefixListUpdate = &networkProtectPrefixListUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateNetworkProtectPrefixListRequest) Execute() (*NetworkProtectPrefixListResponse, *http.Response, error) {
	return r.APIService.UpdateNetworkProtectPrefixListExecute(r)
}

/*
UpdateNetworkProtectPrefixList Update a Network Protect prefix list

Updates a prefix list. The family cannot be changed. Sending addresses replaces the complete list; an empty array clears it when no active source-list restriction applies. Addresses are canonicalized, deduplicated, and sorted by the backend.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param prefixListId Unique identifier of the Network Protect prefix list.
 @return APIUpdateNetworkProtectPrefixListRequest
*/
func (a *NetworkProtectAPIService) UpdateNetworkProtectPrefixList(ctx context.Context, prefixListId string) APIUpdateNetworkProtectPrefixListRequest {
	return APIUpdateNetworkProtectPrefixListRequest{
		APIService:   a,
		ctx:          ctx,
		prefixListId: prefixListId,
	}
}

// UpdateNetworkProtectPrefixListExecute executes the request
//  @return NetworkProtectPrefixListResponse
func (a *NetworkProtectAPIService) UpdateNetworkProtectPrefixListExecute(r APIUpdateNetworkProtectPrefixListRequest) (*NetworkProtectPrefixListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectPrefixListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.UpdateNetworkProtectPrefixList")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/prefix-lists/{prefix_list_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"prefix_list_id"+"}", gourl.PathEscape(parameterToString(r.prefixListId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectPrefixListUpdate == nil {
		return localVarReturnValue, nil, reportError("networkProtectPrefixListUpdate is required and must be specified")
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
	localVarPostBody = r.networkProtectPrefixListUpdate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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

// APIUpdateNetworkProtectRuleSetRequest represents a request for the resource.
type APIUpdateNetworkProtectRuleSetRequest struct {
	ctx                         context.Context
	APIService                  NetworkProtectAPI
	ruleSetId                   string
	networkProtectRuleSetUpdate *NetworkProtectRuleSetUpdate
}

// NetworkProtectRuleSetUpdate returns a pointer to a request.
func (r *APIUpdateNetworkProtectRuleSetRequest) NetworkProtectRuleSetUpdate(networkProtectRuleSetUpdate NetworkProtectRuleSetUpdate) *APIUpdateNetworkProtectRuleSetRequest {
	r.networkProtectRuleSetUpdate = &networkProtectRuleSetUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateNetworkProtectRuleSetRequest) Execute() (*NetworkProtectRuleSetResponse, *http.Response, error) {
	return r.APIService.UpdateNetworkProtectRuleSetExecute(r)
}

/*
UpdateNetworkProtectRuleSet Update a Network Protect rule set

Updates a rule set. Omitted fields remain unchanged. Sending rules replaces the complete ordered list; an empty array removes every rule. Custom attachment and detachment remain restricted in this release.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleSetId Unique identifier of the Network Protect rule set.
 @return APIUpdateNetworkProtectRuleSetRequest
*/
func (a *NetworkProtectAPIService) UpdateNetworkProtectRuleSet(ctx context.Context, ruleSetId string) APIUpdateNetworkProtectRuleSetRequest {
	return APIUpdateNetworkProtectRuleSetRequest{
		APIService: a,
		ctx:        ctx,
		ruleSetId:  ruleSetId,
	}
}

// UpdateNetworkProtectRuleSetExecute executes the request
//  @return NetworkProtectRuleSetResponse
func (a *NetworkProtectAPIService) UpdateNetworkProtectRuleSetExecute(r APIUpdateNetworkProtectRuleSetRequest) (*NetworkProtectRuleSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkProtectRuleSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkProtectAPIService.UpdateNetworkProtectRuleSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-protect/v1/rule-sets/{rule_set_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"rule_set_id"+"}", gourl.PathEscape(parameterToString(r.ruleSetId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.networkProtectRuleSetUpdate == nil {
		return localVarReturnValue, nil, reportError("networkProtectRuleSetUpdate is required and must be specified")
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
	localVarPostBody = r.networkProtectRuleSetUpdate
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
			var v NetworkProtectInvalidRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v NetworkProtectConflict
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
