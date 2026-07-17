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

// DmRoutingConfigsAPI defines an interface for interacting with the resource.
type DmRoutingConfigsAPI interface {

	/*
		ActivateDmRoutingConfigDraft Activate the draft

		Activate the current draft version. The previously active version, if any, becomes inactive but is retained in version history.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIActivateDmRoutingConfigDraftRequest
	*/
	ActivateDmRoutingConfigDraft(ctx context.Context, configId string) APIActivateDmRoutingConfigDraftRequest

	// ActivateDmRoutingConfigDraftExecute executes the request
	//  @return RoutingConfigVersionResponse
	ActivateDmRoutingConfigDraftExecute(r APIActivateDmRoutingConfigDraftRequest) (*RoutingConfigVersionResponse, *http.Response, error)

	/*
		CreateDmRoutingConfig Create a routing config

		Create a new routing config. An optional `initial_version` may be provided to seed the config with paths and rules in a single request, and may also be activated immediately.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APICreateDmRoutingConfigRequest
	*/
	CreateDmRoutingConfig(ctx context.Context) APICreateDmRoutingConfigRequest

	// CreateDmRoutingConfigExecute executes the request
	//  @return RoutingConfigResponse
	CreateDmRoutingConfigExecute(r APICreateDmRoutingConfigRequest) (*RoutingConfigResponse, *http.Response, error)

	/*
		CreateDmRoutingConfigPath Create a path

		Add a new path to the config's draft version. If no draft exists, one is created automatically by cloning the active version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APICreateDmRoutingConfigPathRequest
	*/
	CreateDmRoutingConfigPath(ctx context.Context, configId string) APICreateDmRoutingConfigPathRequest

	// CreateDmRoutingConfigPathExecute executes the request
	//  @return PathResponse
	CreateDmRoutingConfigPathExecute(r APICreateDmRoutingConfigPathRequest) (*PathResponse, *http.Response, error)

	/*
		CreateDmRoutingConfigRule Create a rule

		Add a new rule to a path on the config's draft version. If no draft exists, one is created automatically by cloning the active version. A rule with an empty `conditions` array is a default (catch-all) rule and there can be at most one default rule per path.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @return APICreateDmRoutingConfigRuleRequest
	*/
	CreateDmRoutingConfigRule(ctx context.Context, configId string, pathId string) APICreateDmRoutingConfigRuleRequest

	// CreateDmRoutingConfigRuleExecute executes the request
	//  @return RuleResponse
	CreateDmRoutingConfigRuleExecute(r APICreateDmRoutingConfigRuleRequest) (*RuleResponse, *http.Response, error)

	/*
		DeactivateDmRoutingConfig Deactivate a routing config

		Clear the active version designation. This is a bookkeeping operation only — it does not stop edge traffic. Minerva continues serving the last-activated version until the domain association is removed in Spotless. Only removing the routing config from the domain (via Spotless) triggers Neptune to drop the reference, which causes Minerva to stop fetching and eventually clean up the cached config. Idempotent: returns 200 even if already deactivated.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIDeactivateDmRoutingConfigRequest
	*/
	DeactivateDmRoutingConfig(ctx context.Context, configId string) APIDeactivateDmRoutingConfigRequest

	// DeactivateDmRoutingConfigExecute executes the request
	//  @return RoutingConfigResponse
	DeactivateDmRoutingConfigExecute(r APIDeactivateDmRoutingConfigRequest) (*RoutingConfigResponse, *http.Response, error)

	/*
		DeleteDmRoutingConfig Delete a routing config

		Delete a routing config. By default, configs that have an active version cannot be deleted. Pass `force=true` to bypass the active-version check — this is destructive and will immediately stop traffic routing for any paths the config serves. The `force` parameter does **not** bypass the domain-association check; if domains are still associated, deletion is rejected with 409 regardless of `force`.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIDeleteDmRoutingConfigRequest
	*/
	DeleteDmRoutingConfig(ctx context.Context, configId string) APIDeleteDmRoutingConfigRequest

	// DeleteDmRoutingConfigExecute executes the request
	DeleteDmRoutingConfigExecute(r APIDeleteDmRoutingConfigRequest) (*http.Response, error)

	/*
		DeleteDmRoutingConfigInactiveVersions Delete inactive versions

		Delete all inactive versions for a routing config. The currently active version, if any, is retained.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIDeleteDmRoutingConfigInactiveVersionsRequest
	*/
	DeleteDmRoutingConfigInactiveVersions(ctx context.Context, configId string) APIDeleteDmRoutingConfigInactiveVersionsRequest

	// DeleteDmRoutingConfigInactiveVersionsExecute executes the request
	DeleteDmRoutingConfigInactiveVersionsExecute(r APIDeleteDmRoutingConfigInactiveVersionsRequest) (*http.Response, error)

	/*
		DeleteDmRoutingConfigPath Delete a path

		Delete a path from the config's draft version. If no draft exists, one is created automatically by cloning the active version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @return APIDeleteDmRoutingConfigPathRequest
	*/
	DeleteDmRoutingConfigPath(ctx context.Context, configId string, pathId string) APIDeleteDmRoutingConfigPathRequest

	// DeleteDmRoutingConfigPathExecute executes the request
	DeleteDmRoutingConfigPathExecute(r APIDeleteDmRoutingConfigPathRequest) (*http.Response, error)

	/*
		DeleteDmRoutingConfigRule Delete a rule

		Delete a rule from the config's draft version. If no draft exists, one is created automatically by cloning the active version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @param ruleId
		 @return APIDeleteDmRoutingConfigRuleRequest
	*/
	DeleteDmRoutingConfigRule(ctx context.Context, configId string, pathId string, ruleId string) APIDeleteDmRoutingConfigRuleRequest

	// DeleteDmRoutingConfigRuleExecute executes the request
	DeleteDmRoutingConfigRuleExecute(r APIDeleteDmRoutingConfigRuleRequest) (*http.Response, error)

	/*
		DiscardDmRoutingConfigDraft Discard the draft

		Delete the current draft version, reverting any unactivated changes.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIDiscardDmRoutingConfigDraftRequest
	*/
	DiscardDmRoutingConfigDraft(ctx context.Context, configId string) APIDiscardDmRoutingConfigDraftRequest

	// DiscardDmRoutingConfigDraftExecute executes the request
	DiscardDmRoutingConfigDraftExecute(r APIDiscardDmRoutingConfigDraftRequest) (*http.Response, error)

	/*
		GetDmRoutingConfig Get a routing config

		Retrieve a single routing config by its identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIGetDmRoutingConfigRequest
	*/
	GetDmRoutingConfig(ctx context.Context, configId string) APIGetDmRoutingConfigRequest

	// GetDmRoutingConfigExecute executes the request
	//  @return RoutingConfigResponse
	GetDmRoutingConfigExecute(r APIGetDmRoutingConfigRequest) (*RoutingConfigResponse, *http.Response, error)

	/*
		GetDmRoutingConfigDraftDiff Get the draft diff

		Compare the current draft version against the active version and return the paths and rules that have been added, modified, or deleted.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIGetDmRoutingConfigDraftDiffRequest
	*/
	GetDmRoutingConfigDraftDiff(ctx context.Context, configId string) APIGetDmRoutingConfigDraftDiffRequest

	// GetDmRoutingConfigDraftDiffExecute executes the request
	//  @return DraftDiff
	GetDmRoutingConfigDraftDiffExecute(r APIGetDmRoutingConfigDraftDiffRequest) (*DraftDiff, *http.Response, error)

	/*
		GetDmRoutingConfigPath Get a path

		Retrieve a single path by its stable identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @return APIGetDmRoutingConfigPathRequest
	*/
	GetDmRoutingConfigPath(ctx context.Context, configId string, pathId string) APIGetDmRoutingConfigPathRequest

	// GetDmRoutingConfigPathExecute executes the request
	//  @return PathResponse
	GetDmRoutingConfigPathExecute(r APIGetDmRoutingConfigPathRequest) (*PathResponse, *http.Response, error)

	/*
		GetDmRoutingConfigRule Get a rule

		Retrieve a single rule by its stable identifier.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @param ruleId
		 @return APIGetDmRoutingConfigRuleRequest
	*/
	GetDmRoutingConfigRule(ctx context.Context, configId string, pathId string, ruleId string) APIGetDmRoutingConfigRuleRequest

	// GetDmRoutingConfigRuleExecute executes the request
	//  @return RuleResponse
	GetDmRoutingConfigRuleExecute(r APIGetDmRoutingConfigRuleRequest) (*RuleResponse, *http.Response, error)

	/*
		ListDmRoutingConfigPaths List paths

		List paths for the config. Returns paths from the active version if one exists, otherwise from the draft.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIListDmRoutingConfigPathsRequest
	*/
	ListDmRoutingConfigPaths(ctx context.Context, configId string) APIListDmRoutingConfigPathsRequest

	// ListDmRoutingConfigPathsExecute executes the request
	//  @return PathsResponse
	ListDmRoutingConfigPathsExecute(r APIListDmRoutingConfigPathsRequest) (*PathsResponse, *http.Response, error)

	/*
		ListDmRoutingConfigRules List rules

		List all rules for a path in evaluation order.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @return APIListDmRoutingConfigRulesRequest
	*/
	ListDmRoutingConfigRules(ctx context.Context, configId string, pathId string) APIListDmRoutingConfigRulesRequest

	// ListDmRoutingConfigRulesExecute executes the request
	//  @return RulesResponse
	ListDmRoutingConfigRulesExecute(r APIListDmRoutingConfigRulesRequest) (*RulesResponse, *http.Response, error)

	/*
		ListDmRoutingConfigVersions List versions

		List all versions for a routing config.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIListDmRoutingConfigVersionsRequest
	*/
	ListDmRoutingConfigVersions(ctx context.Context, configId string) APIListDmRoutingConfigVersionsRequest

	// ListDmRoutingConfigVersionsExecute executes the request
	//  @return VersionsResponse
	ListDmRoutingConfigVersionsExecute(r APIListDmRoutingConfigVersionsRequest) (*VersionsResponse, *http.Response, error)

	/*
		ListDmRoutingConfigs List routing configs

		List all routing configs for the authenticated customer.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIListDmRoutingConfigsRequest
	*/
	ListDmRoutingConfigs(ctx context.Context) APIListDmRoutingConfigsRequest

	// ListDmRoutingConfigsExecute executes the request
	//  @return RoutingConfigsResponse
	ListDmRoutingConfigsExecute(r APIListDmRoutingConfigsRequest) (*RoutingConfigsResponse, *http.Response, error)

	/*
		ReactivateDmRoutingConfigVersion Reactivate a version

		Reactivate a previously-active version. The currently active version, if any, becomes inactive but is retained in version history.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param versionId
		 @return APIReactivateDmRoutingConfigVersionRequest
	*/
	ReactivateDmRoutingConfigVersion(ctx context.Context, configId string, versionId string) APIReactivateDmRoutingConfigVersionRequest

	// ReactivateDmRoutingConfigVersionExecute executes the request
	//  @return RoutingConfigVersionResponse
	ReactivateDmRoutingConfigVersionExecute(r APIReactivateDmRoutingConfigVersionRequest) (*RoutingConfigVersionResponse, *http.Response, error)

	/*
		UpdateDmRoutingConfigDraft Update the draft

		Update metadata on the draft version, such as its comment. If no draft exists, one is created automatically by cloning the active version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @return APIUpdateDmRoutingConfigDraftRequest
	*/
	UpdateDmRoutingConfigDraft(ctx context.Context, configId string) APIUpdateDmRoutingConfigDraftRequest

	// UpdateDmRoutingConfigDraftExecute executes the request
	//  @return RoutingConfigVersionResponse
	UpdateDmRoutingConfigDraftExecute(r APIUpdateDmRoutingConfigDraftRequest) (*RoutingConfigVersionResponse, *http.Response, error)

	/*
		UpdateDmRoutingConfigPath Update a path

		Update a path on the config's draft version. If no draft exists, one is created automatically by cloning the active version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @return APIUpdateDmRoutingConfigPathRequest
	*/
	UpdateDmRoutingConfigPath(ctx context.Context, configId string, pathId string) APIUpdateDmRoutingConfigPathRequest

	// UpdateDmRoutingConfigPathExecute executes the request
	//  @return PathResponse
	UpdateDmRoutingConfigPathExecute(r APIUpdateDmRoutingConfigPathRequest) (*PathResponse, *http.Response, error)

	/*
		UpdateDmRoutingConfigRule Update a rule

		Update a rule on the config's draft version. If no draft exists, one is created automatically by cloning the active version.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param configId
		 @param pathId
		 @param ruleId
		 @return APIUpdateDmRoutingConfigRuleRequest
	*/
	UpdateDmRoutingConfigRule(ctx context.Context, configId string, pathId string, ruleId string) APIUpdateDmRoutingConfigRuleRequest

	// UpdateDmRoutingConfigRuleExecute executes the request
	//  @return RuleResponse
	UpdateDmRoutingConfigRuleExecute(r APIUpdateDmRoutingConfigRuleRequest) (*RuleResponse, *http.Response, error)
}

// DmRoutingConfigsAPIService DmRoutingConfigsAPI service
type DmRoutingConfigsAPIService service

// APIActivateDmRoutingConfigDraftRequest represents a request for the resource.
type APIActivateDmRoutingConfigDraftRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
}

// Execute calls the API using the request data configured.
func (r APIActivateDmRoutingConfigDraftRequest) Execute() (*RoutingConfigVersionResponse, *http.Response, error) {
	return r.APIService.ActivateDmRoutingConfigDraftExecute(r)
}

/*
ActivateDmRoutingConfigDraft Activate the draft

Activate the current draft version. The previously active version, if any, becomes inactive but is retained in version history.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIActivateDmRoutingConfigDraftRequest
*/
func (a *DmRoutingConfigsAPIService) ActivateDmRoutingConfigDraft(ctx context.Context, configId string) APIActivateDmRoutingConfigDraftRequest {
	return APIActivateDmRoutingConfigDraftRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// ActivateDmRoutingConfigDraftExecute executes the request
//  @return RoutingConfigVersionResponse
func (a *DmRoutingConfigsAPIService) ActivateDmRoutingConfigDraftExecute(r APIActivateDmRoutingConfigDraftRequest) (*RoutingConfigVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.ActivateDmRoutingConfigDraft")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/activate"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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

// APICreateDmRoutingConfigRequest represents a request for the resource.
type APICreateDmRoutingConfigRequest struct {
	ctx           context.Context
	APIService    DmRoutingConfigsAPI
	routingConfig *RoutingConfig
}

// RoutingConfig returns a pointer to a request.
func (r *APICreateDmRoutingConfigRequest) RoutingConfig(routingConfig RoutingConfig) *APICreateDmRoutingConfigRequest {
	r.routingConfig = &routingConfig
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateDmRoutingConfigRequest) Execute() (*RoutingConfigResponse, *http.Response, error) {
	return r.APIService.CreateDmRoutingConfigExecute(r)
}

/*
CreateDmRoutingConfig Create a routing config

Create a new routing config. An optional `initial_version` may be provided to seed the config with paths and rules in a single request, and may also be activated immediately.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APICreateDmRoutingConfigRequest
*/
func (a *DmRoutingConfigsAPIService) CreateDmRoutingConfig(ctx context.Context) APICreateDmRoutingConfigRequest {
	return APICreateDmRoutingConfigRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// CreateDmRoutingConfigExecute executes the request
//  @return RoutingConfigResponse
func (a *DmRoutingConfigsAPIService) CreateDmRoutingConfigExecute(r APICreateDmRoutingConfigRequest) (*RoutingConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.CreateDmRoutingConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs"

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
	localVarPostBody = r.routingConfig
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

// APICreateDmRoutingConfigPathRequest represents a request for the resource.
type APICreateDmRoutingConfigPathRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathCreate *PathCreate
}

// PathCreate returns a pointer to a request.
func (r *APICreateDmRoutingConfigPathRequest) PathCreate(pathCreate PathCreate) *APICreateDmRoutingConfigPathRequest {
	r.pathCreate = &pathCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateDmRoutingConfigPathRequest) Execute() (*PathResponse, *http.Response, error) {
	return r.APIService.CreateDmRoutingConfigPathExecute(r)
}

/*
CreateDmRoutingConfigPath Create a path

Add a new path to the config's draft version. If no draft exists, one is created automatically by cloning the active version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APICreateDmRoutingConfigPathRequest
*/
func (a *DmRoutingConfigsAPIService) CreateDmRoutingConfigPath(ctx context.Context, configId string) APICreateDmRoutingConfigPathRequest {
	return APICreateDmRoutingConfigPathRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// CreateDmRoutingConfigPathExecute executes the request
//  @return PathResponse
func (a *DmRoutingConfigsAPIService) CreateDmRoutingConfigPathExecute(r APICreateDmRoutingConfigPathRequest) (*PathResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PathResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.CreateDmRoutingConfigPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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
	localVarPostBody = r.pathCreate
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

// APICreateDmRoutingConfigRuleRequest represents a request for the resource.
type APICreateDmRoutingConfigRuleRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
	ruleCreate *RuleCreate
}

// RuleCreate returns a pointer to a request.
func (r *APICreateDmRoutingConfigRuleRequest) RuleCreate(ruleCreate RuleCreate) *APICreateDmRoutingConfigRuleRequest {
	r.ruleCreate = &ruleCreate
	return r
}

// Execute calls the API using the request data configured.
func (r APICreateDmRoutingConfigRuleRequest) Execute() (*RuleResponse, *http.Response, error) {
	return r.APIService.CreateDmRoutingConfigRuleExecute(r)
}

/*
CreateDmRoutingConfigRule Create a rule

Add a new rule to a path on the config's draft version. If no draft exists, one is created automatically by cloning the active version. A rule with an empty `conditions` array is a default (catch-all) rule and there can be at most one default rule per path.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @return APICreateDmRoutingConfigRuleRequest
*/
func (a *DmRoutingConfigsAPIService) CreateDmRoutingConfigRule(ctx context.Context, configId string, pathId string) APICreateDmRoutingConfigRuleRequest {
	return APICreateDmRoutingConfigRuleRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
	}
}

// CreateDmRoutingConfigRuleExecute executes the request
//  @return RuleResponse
func (a *DmRoutingConfigsAPIService) CreateDmRoutingConfigRuleExecute(r APICreateDmRoutingConfigRuleRequest) (*RuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.CreateDmRoutingConfigRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))

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
	localVarPostBody = r.ruleCreate
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

// APIDeactivateDmRoutingConfigRequest represents a request for the resource.
type APIDeactivateDmRoutingConfigRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
}

// Execute calls the API using the request data configured.
func (r APIDeactivateDmRoutingConfigRequest) Execute() (*RoutingConfigResponse, *http.Response, error) {
	return r.APIService.DeactivateDmRoutingConfigExecute(r)
}

/*
DeactivateDmRoutingConfig Deactivate a routing config

Clear the active version designation. This is a bookkeeping operation only — it does not stop edge traffic. Minerva continues serving the last-activated version until the domain association is removed in Spotless. Only removing the routing config from the domain (via Spotless) triggers Neptune to drop the reference, which causes Minerva to stop fetching and eventually clean up the cached config. Idempotent: returns 200 even if already deactivated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIDeactivateDmRoutingConfigRequest
*/
func (a *DmRoutingConfigsAPIService) DeactivateDmRoutingConfig(ctx context.Context, configId string) APIDeactivateDmRoutingConfigRequest {
	return APIDeactivateDmRoutingConfigRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// DeactivateDmRoutingConfigExecute executes the request
//  @return RoutingConfigResponse
func (a *DmRoutingConfigsAPIService) DeactivateDmRoutingConfigExecute(r APIDeactivateDmRoutingConfigRequest) (*RoutingConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.DeactivateDmRoutingConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/deactivate"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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

// APIDeleteDmRoutingConfigRequest represents a request for the resource.
type APIDeleteDmRoutingConfigRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	force      *bool
}

// Force When &#x60;true&#x60;, allows deleting a routing config that has an active version. This is destructive — traffic routing for any paths served by the config will stop immediately.
func (r *APIDeleteDmRoutingConfigRequest) Force(force bool) *APIDeleteDmRoutingConfigRequest {
	r.force = &force
	return r
}

// Execute calls the API using the request data configured.
func (r APIDeleteDmRoutingConfigRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteDmRoutingConfigExecute(r)
}

/*
DeleteDmRoutingConfig Delete a routing config

Delete a routing config. By default, configs that have an active version cannot be deleted. Pass `force=true` to bypass the active-version check — this is destructive and will immediately stop traffic routing for any paths the config serves. The `force` parameter does **not** bypass the domain-association check; if domains are still associated, deletion is rejected with 409 regardless of `force`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIDeleteDmRoutingConfigRequest
*/
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfig(ctx context.Context, configId string) APIDeleteDmRoutingConfigRequest {
	return APIDeleteDmRoutingConfigRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// DeleteDmRoutingConfigExecute executes the request
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigExecute(r APIDeleteDmRoutingConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.DeleteDmRoutingConfig")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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

// APIDeleteDmRoutingConfigInactiveVersionsRequest represents a request for the resource.
type APIDeleteDmRoutingConfigInactiveVersionsRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
}

// Execute calls the API using the request data configured.
func (r APIDeleteDmRoutingConfigInactiveVersionsRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteDmRoutingConfigInactiveVersionsExecute(r)
}

/*
DeleteDmRoutingConfigInactiveVersions Delete inactive versions

Delete all inactive versions for a routing config. The currently active version, if any, is retained.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIDeleteDmRoutingConfigInactiveVersionsRequest
*/
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigInactiveVersions(ctx context.Context, configId string) APIDeleteDmRoutingConfigInactiveVersionsRequest {
	return APIDeleteDmRoutingConfigInactiveVersionsRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// DeleteDmRoutingConfigInactiveVersionsExecute executes the request
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigInactiveVersionsExecute(r APIDeleteDmRoutingConfigInactiveVersionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.DeleteDmRoutingConfigInactiveVersions")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/versions/inactive"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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
	localVarHTTPHeaderAccepts := []string{}

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

// APIDeleteDmRoutingConfigPathRequest represents a request for the resource.
type APIDeleteDmRoutingConfigPathRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
}

// Execute calls the API using the request data configured.
func (r APIDeleteDmRoutingConfigPathRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteDmRoutingConfigPathExecute(r)
}

/*
DeleteDmRoutingConfigPath Delete a path

Delete a path from the config's draft version. If no draft exists, one is created automatically by cloning the active version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @return APIDeleteDmRoutingConfigPathRequest
*/
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigPath(ctx context.Context, configId string, pathId string) APIDeleteDmRoutingConfigPathRequest {
	return APIDeleteDmRoutingConfigPathRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
	}
}

// DeleteDmRoutingConfigPathExecute executes the request
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigPathExecute(r APIDeleteDmRoutingConfigPathRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.DeleteDmRoutingConfigPath")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))

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
	localVarHTTPHeaderAccepts := []string{}

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

// APIDeleteDmRoutingConfigRuleRequest represents a request for the resource.
type APIDeleteDmRoutingConfigRuleRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
	ruleId     string
}

// Execute calls the API using the request data configured.
func (r APIDeleteDmRoutingConfigRuleRequest) Execute() (*http.Response, error) {
	return r.APIService.DeleteDmRoutingConfigRuleExecute(r)
}

/*
DeleteDmRoutingConfigRule Delete a rule

Delete a rule from the config's draft version. If no draft exists, one is created automatically by cloning the active version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @param ruleId
 @return APIDeleteDmRoutingConfigRuleRequest
*/
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigRule(ctx context.Context, configId string, pathId string, ruleId string) APIDeleteDmRoutingConfigRuleRequest {
	return APIDeleteDmRoutingConfigRuleRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
		ruleId:     ruleId,
	}
}

// DeleteDmRoutingConfigRuleExecute executes the request
func (a *DmRoutingConfigsAPIService) DeleteDmRoutingConfigRuleExecute(r APIDeleteDmRoutingConfigRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.DeleteDmRoutingConfigRule")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules/{rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))
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
	localVarHTTPHeaderAccepts := []string{}

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

// APIDiscardDmRoutingConfigDraftRequest represents a request for the resource.
type APIDiscardDmRoutingConfigDraftRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
}

// Execute calls the API using the request data configured.
func (r APIDiscardDmRoutingConfigDraftRequest) Execute() (*http.Response, error) {
	return r.APIService.DiscardDmRoutingConfigDraftExecute(r)
}

/*
DiscardDmRoutingConfigDraft Discard the draft

Delete the current draft version, reverting any unactivated changes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIDiscardDmRoutingConfigDraftRequest
*/
func (a *DmRoutingConfigsAPIService) DiscardDmRoutingConfigDraft(ctx context.Context, configId string) APIDiscardDmRoutingConfigDraftRequest {
	return APIDiscardDmRoutingConfigDraftRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// DiscardDmRoutingConfigDraftExecute executes the request
func (a *DmRoutingConfigsAPIService) DiscardDmRoutingConfigDraftExecute(r APIDiscardDmRoutingConfigDraftRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.DiscardDmRoutingConfigDraft")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/draft"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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
	localVarHTTPHeaderAccepts := []string{}

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

// APIGetDmRoutingConfigRequest represents a request for the resource.
type APIGetDmRoutingConfigRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
}

// Execute calls the API using the request data configured.
func (r APIGetDmRoutingConfigRequest) Execute() (*RoutingConfigResponse, *http.Response, error) {
	return r.APIService.GetDmRoutingConfigExecute(r)
}

/*
GetDmRoutingConfig Get a routing config

Retrieve a single routing config by its identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIGetDmRoutingConfigRequest
*/
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfig(ctx context.Context, configId string) APIGetDmRoutingConfigRequest {
	return APIGetDmRoutingConfigRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// GetDmRoutingConfigExecute executes the request
//  @return RoutingConfigResponse
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigExecute(r APIGetDmRoutingConfigRequest) (*RoutingConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.GetDmRoutingConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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

// APIGetDmRoutingConfigDraftDiffRequest represents a request for the resource.
type APIGetDmRoutingConfigDraftDiffRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
}

// Execute calls the API using the request data configured.
func (r APIGetDmRoutingConfigDraftDiffRequest) Execute() (*DraftDiff, *http.Response, error) {
	return r.APIService.GetDmRoutingConfigDraftDiffExecute(r)
}

/*
GetDmRoutingConfigDraftDiff Get the draft diff

Compare the current draft version against the active version and return the paths and rules that have been added, modified, or deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIGetDmRoutingConfigDraftDiffRequest
*/
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigDraftDiff(ctx context.Context, configId string) APIGetDmRoutingConfigDraftDiffRequest {
	return APIGetDmRoutingConfigDraftDiffRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// GetDmRoutingConfigDraftDiffExecute executes the request
//  @return DraftDiff
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigDraftDiffExecute(r APIGetDmRoutingConfigDraftDiffRequest) (*DraftDiff, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DraftDiff
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.GetDmRoutingConfigDraftDiff")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/draft/diff"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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

// APIGetDmRoutingConfigPathRequest represents a request for the resource.
type APIGetDmRoutingConfigPathRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
}

// Execute calls the API using the request data configured.
func (r APIGetDmRoutingConfigPathRequest) Execute() (*PathResponse, *http.Response, error) {
	return r.APIService.GetDmRoutingConfigPathExecute(r)
}

/*
GetDmRoutingConfigPath Get a path

Retrieve a single path by its stable identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @return APIGetDmRoutingConfigPathRequest
*/
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigPath(ctx context.Context, configId string, pathId string) APIGetDmRoutingConfigPathRequest {
	return APIGetDmRoutingConfigPathRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
	}
}

// GetDmRoutingConfigPathExecute executes the request
//  @return PathResponse
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigPathExecute(r APIGetDmRoutingConfigPathRequest) (*PathResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PathResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.GetDmRoutingConfigPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))

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

// APIGetDmRoutingConfigRuleRequest represents a request for the resource.
type APIGetDmRoutingConfigRuleRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
	ruleId     string
}

// Execute calls the API using the request data configured.
func (r APIGetDmRoutingConfigRuleRequest) Execute() (*RuleResponse, *http.Response, error) {
	return r.APIService.GetDmRoutingConfigRuleExecute(r)
}

/*
GetDmRoutingConfigRule Get a rule

Retrieve a single rule by its stable identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @param ruleId
 @return APIGetDmRoutingConfigRuleRequest
*/
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigRule(ctx context.Context, configId string, pathId string, ruleId string) APIGetDmRoutingConfigRuleRequest {
	return APIGetDmRoutingConfigRuleRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
		ruleId:     ruleId,
	}
}

// GetDmRoutingConfigRuleExecute executes the request
//  @return RuleResponse
func (a *DmRoutingConfigsAPIService) GetDmRoutingConfigRuleExecute(r APIGetDmRoutingConfigRuleRequest) (*RuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.GetDmRoutingConfigRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules/{rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))
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

// APIListDmRoutingConfigPathsRequest represents a request for the resource.
type APIListDmRoutingConfigPathsRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	path       *string
	match      *string
	sort       *string
	cursor     *string
	limit      *int32
}

// Path Filter results by path pattern. The match strategy is controlled by the &#x60;match&#x60; parameter.
func (r *APIListDmRoutingConfigPathsRequest) Path(path string) *APIListDmRoutingConfigPathsRequest {
	r.path = &path
	return r
}

// Match How to match the value of the &#x60;path&#x60; filter against existing path patterns. Has no effect unless &#x60;path&#x60; is also provided.
func (r *APIListDmRoutingConfigPathsRequest) Match(match string) *APIListDmRoutingConfigPathsRequest {
	r.match = &match
	return r
}

// Sort The order in which to list the results.
func (r *APIListDmRoutingConfigPathsRequest) Sort(sort string) *APIListDmRoutingConfigPathsRequest {
	r.sort = &sort
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIListDmRoutingConfigPathsRequest) Cursor(cursor string) *APIListDmRoutingConfigPathsRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIListDmRoutingConfigPathsRequest) Limit(limit int32) *APIListDmRoutingConfigPathsRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIListDmRoutingConfigPathsRequest) Execute() (*PathsResponse, *http.Response, error) {
	return r.APIService.ListDmRoutingConfigPathsExecute(r)
}

/*
ListDmRoutingConfigPaths List paths

List paths for the config. Returns paths from the active version if one exists, otherwise from the draft.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIListDmRoutingConfigPathsRequest
*/
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigPaths(ctx context.Context, configId string) APIListDmRoutingConfigPathsRequest {
	return APIListDmRoutingConfigPathsRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// ListDmRoutingConfigPathsExecute executes the request
//  @return PathsResponse
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigPathsExecute(r APIListDmRoutingConfigPathsRequest) (*PathsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PathsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.ListDmRoutingConfigPaths")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.path != nil {
		localVarQueryParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.match != nil {
		localVarQueryParams.Add("match", parameterToString(*r.match, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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

// APIListDmRoutingConfigRulesRequest represents a request for the resource.
type APIListDmRoutingConfigRulesRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
	sort       *string
	cursor     *string
	limit      *int32
}

// Sort The order in which to list the results.
func (r *APIListDmRoutingConfigRulesRequest) Sort(sort string) *APIListDmRoutingConfigRulesRequest {
	r.sort = &sort
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIListDmRoutingConfigRulesRequest) Cursor(cursor string) *APIListDmRoutingConfigRulesRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIListDmRoutingConfigRulesRequest) Limit(limit int32) *APIListDmRoutingConfigRulesRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIListDmRoutingConfigRulesRequest) Execute() (*RulesResponse, *http.Response, error) {
	return r.APIService.ListDmRoutingConfigRulesExecute(r)
}

/*
ListDmRoutingConfigRules List rules

List all rules for a path in evaluation order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @return APIListDmRoutingConfigRulesRequest
*/
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigRules(ctx context.Context, configId string, pathId string) APIListDmRoutingConfigRulesRequest {
	return APIListDmRoutingConfigRulesRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
	}
}

// ListDmRoutingConfigRulesExecute executes the request
//  @return RulesResponse
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigRulesExecute(r APIListDmRoutingConfigRulesRequest) (*RulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.ListDmRoutingConfigRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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

// APIListDmRoutingConfigVersionsRequest represents a request for the resource.
type APIListDmRoutingConfigVersionsRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	sort       *string
	cursor     *string
	limit      *int32
}

// Sort The order in which to list the results.
func (r *APIListDmRoutingConfigVersionsRequest) Sort(sort string) *APIListDmRoutingConfigVersionsRequest {
	r.sort = &sort
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIListDmRoutingConfigVersionsRequest) Cursor(cursor string) *APIListDmRoutingConfigVersionsRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIListDmRoutingConfigVersionsRequest) Limit(limit int32) *APIListDmRoutingConfigVersionsRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIListDmRoutingConfigVersionsRequest) Execute() (*VersionsResponse, *http.Response, error) {
	return r.APIService.ListDmRoutingConfigVersionsExecute(r)
}

/*
ListDmRoutingConfigVersions List versions

List all versions for a routing config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIListDmRoutingConfigVersionsRequest
*/
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigVersions(ctx context.Context, configId string) APIListDmRoutingConfigVersionsRequest {
	return APIListDmRoutingConfigVersionsRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// ListDmRoutingConfigVersionsExecute executes the request
//  @return VersionsResponse
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigVersionsExecute(r APIListDmRoutingConfigVersionsRequest) (*VersionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *VersionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.ListDmRoutingConfigVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/versions"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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

// APIListDmRoutingConfigsRequest represents a request for the resource.
type APIListDmRoutingConfigsRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	state      *[]string
	sort       *string
	cursor     *string
	limit      *int32
}

// State Filter configs by lifecycle state. Accepts a comma-separated list of state values (e.g. &#x60;?state&#x3D;active,active-with-draft&#x60;). Returns only configs whose current state matches one of the provided values. Returns 400 if any value is not a recognised state.
func (r *APIListDmRoutingConfigsRequest) State(state []string) *APIListDmRoutingConfigsRequest {
	r.state = &state
	return r
}

// Sort The order in which to list the results.
func (r *APIListDmRoutingConfigsRequest) Sort(sort string) *APIListDmRoutingConfigsRequest {
	r.sort = &sort
	return r
}

// Cursor Cursor value from the &#x60;next_cursor&#x60; field of a previous response, used to retrieve the next page. To request the first page, this should be empty.
func (r *APIListDmRoutingConfigsRequest) Cursor(cursor string) *APIListDmRoutingConfigsRequest {
	r.cursor = &cursor
	return r
}

// Limit Limit how many results are returned.
func (r *APIListDmRoutingConfigsRequest) Limit(limit int32) *APIListDmRoutingConfigsRequest {
	r.limit = &limit
	return r
}

// Execute calls the API using the request data configured.
func (r APIListDmRoutingConfigsRequest) Execute() (*RoutingConfigsResponse, *http.Response, error) {
	return r.APIService.ListDmRoutingConfigsExecute(r)
}

/*
ListDmRoutingConfigs List routing configs

List all routing configs for the authenticated customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIListDmRoutingConfigsRequest
*/
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigs(ctx context.Context) APIListDmRoutingConfigsRequest {
	return APIListDmRoutingConfigsRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// ListDmRoutingConfigsExecute executes the request
//  @return RoutingConfigsResponse
func (a *DmRoutingConfigsAPIService) ListDmRoutingConfigsExecute(r APIListDmRoutingConfigsRequest) (*RoutingConfigsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.ListDmRoutingConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.state != nil {
		localVarQueryParams.Add("state", parameterToString(*r.state, "csv"))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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

// APIReactivateDmRoutingConfigVersionRequest represents a request for the resource.
type APIReactivateDmRoutingConfigVersionRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	versionId  string
}

// Execute calls the API using the request data configured.
func (r APIReactivateDmRoutingConfigVersionRequest) Execute() (*RoutingConfigVersionResponse, *http.Response, error) {
	return r.APIService.ReactivateDmRoutingConfigVersionExecute(r)
}

/*
ReactivateDmRoutingConfigVersion Reactivate a version

Reactivate a previously-active version. The currently active version, if any, becomes inactive but is retained in version history.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param versionId
 @return APIReactivateDmRoutingConfigVersionRequest
*/
func (a *DmRoutingConfigsAPIService) ReactivateDmRoutingConfigVersion(ctx context.Context, configId string, versionId string) APIReactivateDmRoutingConfigVersionRequest {
	return APIReactivateDmRoutingConfigVersionRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		versionId:  versionId,
	}
}

// ReactivateDmRoutingConfigVersionExecute executes the request
//  @return RoutingConfigVersionResponse
func (a *DmRoutingConfigsAPIService) ReactivateDmRoutingConfigVersionExecute(r APIReactivateDmRoutingConfigVersionRequest) (*RoutingConfigVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.ReactivateDmRoutingConfigVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/versions/{version_id}/activate"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
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

// APIUpdateDmRoutingConfigDraftRequest represents a request for the resource.
type APIUpdateDmRoutingConfigDraftRequest struct {
	ctx         context.Context
	APIService  DmRoutingConfigsAPI
	configId    string
	draftUpdate *DraftUpdate
}

// DraftUpdate returns a pointer to a request.
func (r *APIUpdateDmRoutingConfigDraftRequest) DraftUpdate(draftUpdate DraftUpdate) *APIUpdateDmRoutingConfigDraftRequest {
	r.draftUpdate = &draftUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateDmRoutingConfigDraftRequest) Execute() (*RoutingConfigVersionResponse, *http.Response, error) {
	return r.APIService.UpdateDmRoutingConfigDraftExecute(r)
}

/*
UpdateDmRoutingConfigDraft Update the draft

Update metadata on the draft version, such as its comment. If no draft exists, one is created automatically by cloning the active version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @return APIUpdateDmRoutingConfigDraftRequest
*/
func (a *DmRoutingConfigsAPIService) UpdateDmRoutingConfigDraft(ctx context.Context, configId string) APIUpdateDmRoutingConfigDraftRequest {
	return APIUpdateDmRoutingConfigDraftRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// UpdateDmRoutingConfigDraftExecute executes the request
//  @return RoutingConfigVersionResponse
func (a *DmRoutingConfigsAPIService) UpdateDmRoutingConfigDraftExecute(r APIUpdateDmRoutingConfigDraftRequest) (*RoutingConfigVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RoutingConfigVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.UpdateDmRoutingConfigDraft")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/draft"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))

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
	localVarPostBody = r.draftUpdate
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

// APIUpdateDmRoutingConfigPathRequest represents a request for the resource.
type APIUpdateDmRoutingConfigPathRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
	pathUpdate *PathUpdate
}

// PathUpdate returns a pointer to a request.
func (r *APIUpdateDmRoutingConfigPathRequest) PathUpdate(pathUpdate PathUpdate) *APIUpdateDmRoutingConfigPathRequest {
	r.pathUpdate = &pathUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateDmRoutingConfigPathRequest) Execute() (*PathResponse, *http.Response, error) {
	return r.APIService.UpdateDmRoutingConfigPathExecute(r)
}

/*
UpdateDmRoutingConfigPath Update a path

Update a path on the config's draft version. If no draft exists, one is created automatically by cloning the active version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @return APIUpdateDmRoutingConfigPathRequest
*/
func (a *DmRoutingConfigsAPIService) UpdateDmRoutingConfigPath(ctx context.Context, configId string, pathId string) APIUpdateDmRoutingConfigPathRequest {
	return APIUpdateDmRoutingConfigPathRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
	}
}

// UpdateDmRoutingConfigPathExecute executes the request
//  @return PathResponse
func (a *DmRoutingConfigsAPIService) UpdateDmRoutingConfigPathExecute(r APIUpdateDmRoutingConfigPathRequest) (*PathResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PathResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.UpdateDmRoutingConfigPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))

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
	localVarPostBody = r.pathUpdate
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

// APIUpdateDmRoutingConfigRuleRequest represents a request for the resource.
type APIUpdateDmRoutingConfigRuleRequest struct {
	ctx        context.Context
	APIService DmRoutingConfigsAPI
	configId   string
	pathId     string
	ruleId     string
	ruleUpdate *RuleUpdate
}

// RuleUpdate returns a pointer to a request.
func (r *APIUpdateDmRoutingConfigRuleRequest) RuleUpdate(ruleUpdate RuleUpdate) *APIUpdateDmRoutingConfigRuleRequest {
	r.ruleUpdate = &ruleUpdate
	return r
}

// Execute calls the API using the request data configured.
func (r APIUpdateDmRoutingConfigRuleRequest) Execute() (*RuleResponse, *http.Response, error) {
	return r.APIService.UpdateDmRoutingConfigRuleExecute(r)
}

/*
UpdateDmRoutingConfigRule Update a rule

Update a rule on the config's draft version. If no draft exists, one is created automatically by cloning the active version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId
 @param pathId
 @param ruleId
 @return APIUpdateDmRoutingConfigRuleRequest
*/
func (a *DmRoutingConfigsAPIService) UpdateDmRoutingConfigRule(ctx context.Context, configId string, pathId string, ruleId string) APIUpdateDmRoutingConfigRuleRequest {
	return APIUpdateDmRoutingConfigRuleRequest{
		APIService: a,
		ctx:        ctx,
		configId:   configId,
		pathId:     pathId,
		ruleId:     ruleId,
	}
}

// UpdateDmRoutingConfigRuleExecute executes the request
//  @return RuleResponse
func (a *DmRoutingConfigsAPIService) UpdateDmRoutingConfigRuleExecute(r APIUpdateDmRoutingConfigRuleRequest) (*RuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *RuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmRoutingConfigsAPIService.UpdateDmRoutingConfigRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domain-management/v1/routing-configs/{config_id}/paths/{path_id}/rules/{rule_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"config_id"+"}", gourl.PathEscape(parameterToString(r.configId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"path_id"+"}", gourl.PathEscape(parameterToString(r.pathId, "")))
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.ruleUpdate
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
