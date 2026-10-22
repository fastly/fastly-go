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

// BotManagementAPI defines an interface for interacting with the resource.
type BotManagementAPI interface {

	/*
		BotManagementGetBot Get a bot

		Get the attributes of a bot in a category.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @param categoryId The ID of the bot category.
		 @param botId The ID of the bot.
		 @return APIBotManagementGetBotRequest
	*/
	BotManagementGetBot(ctx context.Context, workspaceId string, categoryId string, botId string) APIBotManagementGetBotRequest

	// BotManagementGetBotExecute executes the request
	//  @return BotManagementBotResponse
	BotManagementGetBotExecute(r APIBotManagementGetBotRequest) (*BotManagementBotResponse, *http.Response, error)

	/*
		BotManagementGetCategory Get a category

		Get the attributes of a bot category.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @param categoryId The ID of the bot category.
		 @return APIBotManagementGetCategoryRequest
	*/
	BotManagementGetCategory(ctx context.Context, workspaceId string, categoryId string) APIBotManagementGetCategoryRequest

	// BotManagementGetCategoryExecute executes the request
	//  @return BotManagementCategoryResponse
	BotManagementGetCategoryExecute(r APIBotManagementGetCategoryRequest) (*BotManagementCategoryResponse, *http.Response, error)

	/*
		BotManagementGetWorkspace Get a workspace

		Retrieve a bot management workspace.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @return APIBotManagementGetWorkspaceRequest
	*/
	BotManagementGetWorkspace(ctx context.Context, workspaceId string) APIBotManagementGetWorkspaceRequest

	// BotManagementGetWorkspaceExecute executes the request
	//  @return BotManagementWorkspaceResponse
	BotManagementGetWorkspaceExecute(r APIBotManagementGetWorkspaceRequest) (*BotManagementWorkspaceResponse, *http.Response, error)

	/*
		BotManagementListBots List all bots in a workspace

		List all bots in a workspace's policy across all categories.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @return APIBotManagementListBotsRequest
	*/
	BotManagementListBots(ctx context.Context, workspaceId string) APIBotManagementListBotsRequest

	// BotManagementListBotsExecute executes the request
	//  @return BotManagementWorkspaceBotsResponse
	BotManagementListBotsExecute(r APIBotManagementListBotsRequest) (*BotManagementWorkspaceBotsResponse, *http.Response, error)

	/*
		BotManagementListCategories List categories

		List all bot categories in a workspace's policy. Each category has an action that applies to every bot in the category. To list the bots themselves, use [list all bots in a workspace](#bot-management-list-bots) or [list bots in a category](#bot-management-list-category-bots).

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @return APIBotManagementListCategoriesRequest
	*/
	BotManagementListCategories(ctx context.Context, workspaceId string) APIBotManagementListCategoriesRequest

	// BotManagementListCategoriesExecute executes the request
	//  @return BotManagementCategoriesResponse
	BotManagementListCategoriesExecute(r APIBotManagementListCategoriesRequest) (*BotManagementCategoriesResponse, *http.Response, error)

	/*
		BotManagementListCategoryBots List bots in a category

		List all bots in a bot category.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @param categoryId The ID of the bot category.
		 @return APIBotManagementListCategoryBotsRequest
	*/
	BotManagementListCategoryBots(ctx context.Context, workspaceId string, categoryId string) APIBotManagementListCategoryBotsRequest

	// BotManagementListCategoryBotsExecute executes the request
	//  @return BotManagementBotsResponse
	BotManagementListCategoryBotsExecute(r APIBotManagementListCategoryBotsRequest) (*BotManagementBotsResponse, *http.Response, error)

	/*
		BotManagementListWorkspaces List workspaces

		List bot management workspaces.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return APIBotManagementListWorkspacesRequest
	*/
	BotManagementListWorkspaces(ctx context.Context) APIBotManagementListWorkspacesRequest

	// BotManagementListWorkspacesExecute executes the request
	//  @return BotManagementWorkspacesResponse
	BotManagementListWorkspacesExecute(r APIBotManagementListWorkspacesRequest) (*BotManagementWorkspacesResponse, *http.Response, error)

	/*
		BotManagementUpdateBot Set a bot action

		Set the action for a bot. Set the action to 'inherit' to apply the category's action or to a specific value to override the category's action.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @param categoryId The ID of the bot category.
		 @param botId The ID of the bot.
		 @return APIBotManagementUpdateBotRequest
	*/
	BotManagementUpdateBot(ctx context.Context, workspaceId string, categoryId string, botId string) APIBotManagementUpdateBotRequest

	// BotManagementUpdateBotExecute executes the request
	//  @return BotManagementBotResponse
	BotManagementUpdateBotExecute(r APIBotManagementUpdateBotRequest) (*BotManagementBotResponse, *http.Response, error)

	/*
		BotManagementUpdateCategory Set a category action

		Set the action for a bot category. Every bot in the category configured with the 'inherit' action uses the action set for the category.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @param categoryId The ID of the bot category.
		 @return APIBotManagementUpdateCategoryRequest
	*/
	BotManagementUpdateCategory(ctx context.Context, workspaceId string, categoryId string) APIBotManagementUpdateCategoryRequest

	// BotManagementUpdateCategoryExecute executes the request
	//  @return BotManagementCategoryResponse
	BotManagementUpdateCategoryExecute(r APIBotManagementUpdateCategoryRequest) (*BotManagementCategoryResponse, *http.Response, error)

	/*
		BotManagementUpdateWorkspace Update a workspace

		Update the attributes of a bot management workspace.

		 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param workspaceId The ID of the bot management workspace.
		 @return APIBotManagementUpdateWorkspaceRequest
	*/
	BotManagementUpdateWorkspace(ctx context.Context, workspaceId string) APIBotManagementUpdateWorkspaceRequest

	// BotManagementUpdateWorkspaceExecute executes the request
	//  @return BotManagementWorkspaceResponse
	BotManagementUpdateWorkspaceExecute(r APIBotManagementUpdateWorkspaceRequest) (*BotManagementWorkspaceResponse, *http.Response, error)
}

// BotManagementAPIService BotManagementAPI service
type BotManagementAPIService service

// APIBotManagementGetBotRequest represents a request for the resource.
type APIBotManagementGetBotRequest struct {
	ctx         context.Context
	APIService  BotManagementAPI
	workspaceId string
	categoryId  string
	botId       string
}

// Execute calls the API using the request data configured.
func (r APIBotManagementGetBotRequest) Execute() (*BotManagementBotResponse, *http.Response, error) {
	return r.APIService.BotManagementGetBotExecute(r)
}

/*
BotManagementGetBot Get a bot

Get the attributes of a bot in a category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @param categoryId The ID of the bot category.
 @param botId The ID of the bot.
 @return APIBotManagementGetBotRequest
*/
func (a *BotManagementAPIService) BotManagementGetBot(ctx context.Context, workspaceId string, categoryId string, botId string) APIBotManagementGetBotRequest {
	return APIBotManagementGetBotRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
		categoryId:  categoryId,
		botId:       botId,
	}
}

// BotManagementGetBotExecute executes the request
//  @return BotManagementBotResponse
func (a *BotManagementAPIService) BotManagementGetBotExecute(r APIBotManagementGetBotRequest) (*BotManagementBotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementBotResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementGetBot")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots/{bot_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"category_id"+"}", gourl.PathEscape(parameterToString(r.categoryId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"bot_id"+"}", gourl.PathEscape(parameterToString(r.botId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementGetCategoryRequest represents a request for the resource.
type APIBotManagementGetCategoryRequest struct {
	ctx         context.Context
	APIService  BotManagementAPI
	workspaceId string
	categoryId  string
}

// Execute calls the API using the request data configured.
func (r APIBotManagementGetCategoryRequest) Execute() (*BotManagementCategoryResponse, *http.Response, error) {
	return r.APIService.BotManagementGetCategoryExecute(r)
}

/*
BotManagementGetCategory Get a category

Get the attributes of a bot category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @param categoryId The ID of the bot category.
 @return APIBotManagementGetCategoryRequest
*/
func (a *BotManagementAPIService) BotManagementGetCategory(ctx context.Context, workspaceId string, categoryId string) APIBotManagementGetCategoryRequest {
	return APIBotManagementGetCategoryRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
		categoryId:  categoryId,
	}
}

// BotManagementGetCategoryExecute executes the request
//  @return BotManagementCategoryResponse
func (a *BotManagementAPIService) BotManagementGetCategoryExecute(r APIBotManagementGetCategoryRequest) (*BotManagementCategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementGetCategory")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"category_id"+"}", gourl.PathEscape(parameterToString(r.categoryId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementGetWorkspaceRequest represents a request for the resource.
type APIBotManagementGetWorkspaceRequest struct {
	ctx         context.Context
	APIService  BotManagementAPI
	workspaceId string
}

// Execute calls the API using the request data configured.
func (r APIBotManagementGetWorkspaceRequest) Execute() (*BotManagementWorkspaceResponse, *http.Response, error) {
	return r.APIService.BotManagementGetWorkspaceExecute(r)
}

/*
BotManagementGetWorkspace Get a workspace

Retrieve a bot management workspace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @return APIBotManagementGetWorkspaceRequest
*/
func (a *BotManagementAPIService) BotManagementGetWorkspace(ctx context.Context, workspaceId string) APIBotManagementGetWorkspaceRequest {
	return APIBotManagementGetWorkspaceRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// BotManagementGetWorkspaceExecute executes the request
//  @return BotManagementWorkspaceResponse
func (a *BotManagementAPIService) BotManagementGetWorkspaceExecute(r APIBotManagementGetWorkspaceRequest) (*BotManagementWorkspaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementWorkspaceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementGetWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementListBotsRequest represents a request for the resource.
type APIBotManagementListBotsRequest struct {
	ctx         context.Context
	APIService  BotManagementAPI
	workspaceId string
	limit       *int32
	cursor      *string
}

// Limit Limit how many results are returned.
func (r *APIBotManagementListBotsRequest) Limit(limit int32) *APIBotManagementListBotsRequest {
	r.limit = &limit
	return r
}

// Cursor Cursor for pagination. Use the &#x60;next_cursor&#x60; value from a previous response to request the next page of results.
func (r *APIBotManagementListBotsRequest) Cursor(cursor string) *APIBotManagementListBotsRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementListBotsRequest) Execute() (*BotManagementWorkspaceBotsResponse, *http.Response, error) {
	return r.APIService.BotManagementListBotsExecute(r)
}

/*
BotManagementListBots List all bots in a workspace

List all bots in a workspace's policy across all categories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @return APIBotManagementListBotsRequest
*/
func (a *BotManagementAPIService) BotManagementListBots(ctx context.Context, workspaceId string) APIBotManagementListBotsRequest {
	return APIBotManagementListBotsRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// BotManagementListBotsExecute executes the request
//  @return BotManagementWorkspaceBotsResponse
func (a *BotManagementAPIService) BotManagementListBotsExecute(r APIBotManagementListBotsRequest) (*BotManagementWorkspaceBotsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementWorkspaceBotsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementListBots")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/bots"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))

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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementListCategoriesRequest represents a request for the resource.
type APIBotManagementListCategoriesRequest struct {
	ctx         context.Context
	APIService  BotManagementAPI
	workspaceId string
	limit       *int32
	cursor      *string
}

// Limit Limit how many results are returned.
func (r *APIBotManagementListCategoriesRequest) Limit(limit int32) *APIBotManagementListCategoriesRequest {
	r.limit = &limit
	return r
}

// Cursor Cursor for pagination. Use the &#x60;next_cursor&#x60; value from a previous response to request the next page of results.
func (r *APIBotManagementListCategoriesRequest) Cursor(cursor string) *APIBotManagementListCategoriesRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementListCategoriesRequest) Execute() (*BotManagementCategoriesResponse, *http.Response, error) {
	return r.APIService.BotManagementListCategoriesExecute(r)
}

/*
BotManagementListCategories List categories

List all bot categories in a workspace's policy. Each category has an action that applies to every bot in the category. To list the bots themselves, use [list all bots in a workspace](#bot-management-list-bots) or [list bots in a category](#bot-management-list-category-bots).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @return APIBotManagementListCategoriesRequest
*/
func (a *BotManagementAPIService) BotManagementListCategories(ctx context.Context, workspaceId string) APIBotManagementListCategoriesRequest {
	return APIBotManagementListCategoriesRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// BotManagementListCategoriesExecute executes the request
//  @return BotManagementCategoriesResponse
func (a *BotManagementAPIService) BotManagementListCategoriesExecute(r APIBotManagementListCategoriesRequest) (*BotManagementCategoriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementCategoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementListCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/categories"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))

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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementListCategoryBotsRequest represents a request for the resource.
type APIBotManagementListCategoryBotsRequest struct {
	ctx         context.Context
	APIService  BotManagementAPI
	workspaceId string
	categoryId  string
	limit       *int32
	cursor      *string
}

// Limit Limit how many results are returned.
func (r *APIBotManagementListCategoryBotsRequest) Limit(limit int32) *APIBotManagementListCategoryBotsRequest {
	r.limit = &limit
	return r
}

// Cursor Cursor for pagination. Use the &#x60;next_cursor&#x60; value from a previous response to request the next page of results.
func (r *APIBotManagementListCategoryBotsRequest) Cursor(cursor string) *APIBotManagementListCategoryBotsRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementListCategoryBotsRequest) Execute() (*BotManagementBotsResponse, *http.Response, error) {
	return r.APIService.BotManagementListCategoryBotsExecute(r)
}

/*
BotManagementListCategoryBots List bots in a category

List all bots in a bot category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @param categoryId The ID of the bot category.
 @return APIBotManagementListCategoryBotsRequest
*/
func (a *BotManagementAPIService) BotManagementListCategoryBots(ctx context.Context, workspaceId string, categoryId string) APIBotManagementListCategoryBotsRequest {
	return APIBotManagementListCategoryBotsRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
		categoryId:  categoryId,
	}
}

// BotManagementListCategoryBotsExecute executes the request
//  @return BotManagementBotsResponse
func (a *BotManagementAPIService) BotManagementListCategoryBotsExecute(r APIBotManagementListCategoryBotsRequest) (*BotManagementBotsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementBotsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementListCategoryBots")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"category_id"+"}", gourl.PathEscape(parameterToString(r.categoryId, "")))

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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementListWorkspacesRequest represents a request for the resource.
type APIBotManagementListWorkspacesRequest struct {
	ctx        context.Context
	APIService BotManagementAPI
	serviceId  *string
	limit      *int32
	cursor     *string
}

// ServiceId Filter results to the workspace attached to this service. Returns an empty list if the service is not attached to a workspace.
func (r *APIBotManagementListWorkspacesRequest) ServiceId(serviceId string) *APIBotManagementListWorkspacesRequest {
	r.serviceId = &serviceId
	return r
}

// Limit Limit how many results are returned.
func (r *APIBotManagementListWorkspacesRequest) Limit(limit int32) *APIBotManagementListWorkspacesRequest {
	r.limit = &limit
	return r
}

// Cursor Cursor for pagination. Use the &#x60;next_cursor&#x60; value from a previous response to request the next page of results.
func (r *APIBotManagementListWorkspacesRequest) Cursor(cursor string) *APIBotManagementListWorkspacesRequest {
	r.cursor = &cursor
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementListWorkspacesRequest) Execute() (*BotManagementWorkspacesResponse, *http.Response, error) {
	return r.APIService.BotManagementListWorkspacesExecute(r)
}

/*
BotManagementListWorkspaces List workspaces

List bot management workspaces.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return APIBotManagementListWorkspacesRequest
*/
func (a *BotManagementAPIService) BotManagementListWorkspaces(ctx context.Context) APIBotManagementListWorkspacesRequest {
	return APIBotManagementListWorkspacesRequest{
		APIService: a,
		ctx:        ctx,
	}
}

// BotManagementListWorkspacesExecute executes the request
//  @return BotManagementWorkspacesResponse
func (a *BotManagementAPIService) BotManagementListWorkspacesExecute(r APIBotManagementListWorkspacesRequest) (*BotManagementWorkspacesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementWorkspacesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementListWorkspaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}

	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementUpdateBotRequest represents a request for the resource.
type APIBotManagementUpdateBotRequest struct {
	ctx              context.Context
	APIService       BotManagementAPI
	workspaceId      string
	categoryId       string
	botId            string
	botManagementBot *BotManagementBot
}

// BotManagementBot returns a pointer to a request.
func (r *APIBotManagementUpdateBotRequest) BotManagementBot(botManagementBot BotManagementBot) *APIBotManagementUpdateBotRequest {
	r.botManagementBot = &botManagementBot
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementUpdateBotRequest) Execute() (*BotManagementBotResponse, *http.Response, error) {
	return r.APIService.BotManagementUpdateBotExecute(r)
}

/*
BotManagementUpdateBot Set a bot action

Set the action for a bot. Set the action to 'inherit' to apply the category's action or to a specific value to override the category's action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @param categoryId The ID of the bot category.
 @param botId The ID of the bot.
 @return APIBotManagementUpdateBotRequest
*/
func (a *BotManagementAPIService) BotManagementUpdateBot(ctx context.Context, workspaceId string, categoryId string, botId string) APIBotManagementUpdateBotRequest {
	return APIBotManagementUpdateBotRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
		categoryId:  categoryId,
		botId:       botId,
	}
}

// BotManagementUpdateBotExecute executes the request
//  @return BotManagementBotResponse
func (a *BotManagementAPIService) BotManagementUpdateBotExecute(r APIBotManagementUpdateBotRequest) (*BotManagementBotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementBotResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementUpdateBot")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}/bots/{bot_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"category_id"+"}", gourl.PathEscape(parameterToString(r.categoryId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"bot_id"+"}", gourl.PathEscape(parameterToString(r.botId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.botManagementBot == nil {
		return localVarReturnValue, nil, reportError("botManagementBot is required and must be specified")
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
	localVarPostBody = r.botManagementBot
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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementUpdateCategoryRequest represents a request for the resource.
type APIBotManagementUpdateCategoryRequest struct {
	ctx                   context.Context
	APIService            BotManagementAPI
	workspaceId           string
	categoryId            string
	botManagementCategory *BotManagementCategory
}

// BotManagementCategory returns a pointer to a request.
func (r *APIBotManagementUpdateCategoryRequest) BotManagementCategory(botManagementCategory BotManagementCategory) *APIBotManagementUpdateCategoryRequest {
	r.botManagementCategory = &botManagementCategory
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementUpdateCategoryRequest) Execute() (*BotManagementCategoryResponse, *http.Response, error) {
	return r.APIService.BotManagementUpdateCategoryExecute(r)
}

/*
BotManagementUpdateCategory Set a category action

Set the action for a bot category. Every bot in the category configured with the 'inherit' action uses the action set for the category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @param categoryId The ID of the bot category.
 @return APIBotManagementUpdateCategoryRequest
*/
func (a *BotManagementAPIService) BotManagementUpdateCategory(ctx context.Context, workspaceId string, categoryId string) APIBotManagementUpdateCategoryRequest {
	return APIBotManagementUpdateCategoryRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
		categoryId:  categoryId,
	}
}

// BotManagementUpdateCategoryExecute executes the request
//  @return BotManagementCategoryResponse
func (a *BotManagementAPIService) BotManagementUpdateCategoryExecute(r APIBotManagementUpdateCategoryRequest) (*BotManagementCategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementUpdateCategory")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}/policy/categories/{category_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"category_id"+"}", gourl.PathEscape(parameterToString(r.categoryId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.botManagementCategory == nil {
		return localVarReturnValue, nil, reportError("botManagementCategory is required and must be specified")
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
	localVarPostBody = r.botManagementCategory
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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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

// APIBotManagementUpdateWorkspaceRequest represents a request for the resource.
type APIBotManagementUpdateWorkspaceRequest struct {
	ctx                    context.Context
	APIService             BotManagementAPI
	workspaceId            string
	botManagementWorkspace *BotManagementWorkspace
}

// BotManagementWorkspace returns a pointer to a request.
func (r *APIBotManagementUpdateWorkspaceRequest) BotManagementWorkspace(botManagementWorkspace BotManagementWorkspace) *APIBotManagementUpdateWorkspaceRequest {
	r.botManagementWorkspace = &botManagementWorkspace
	return r
}

// Execute calls the API using the request data configured.
func (r APIBotManagementUpdateWorkspaceRequest) Execute() (*BotManagementWorkspaceResponse, *http.Response, error) {
	return r.APIService.BotManagementUpdateWorkspaceExecute(r)
}

/*
BotManagementUpdateWorkspace Update a workspace

Update the attributes of a bot management workspace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The ID of the bot management workspace.
 @return APIBotManagementUpdateWorkspaceRequest
*/
func (a *BotManagementAPIService) BotManagementUpdateWorkspace(ctx context.Context, workspaceId string) APIBotManagementUpdateWorkspaceRequest {
	return APIBotManagementUpdateWorkspaceRequest{
		APIService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// BotManagementUpdateWorkspaceExecute executes the request
//  @return BotManagementWorkspaceResponse
func (a *BotManagementAPIService) BotManagementUpdateWorkspaceExecute(r APIBotManagementUpdateWorkspaceRequest) (*BotManagementWorkspaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BotManagementWorkspaceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotManagementAPIService.BotManagementUpdateWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bot-management/v1/workspaces/{workspace_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"workspace_id"+"}", gourl.PathEscape(parameterToString(r.workspaceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := gourl.Values{}
	localVarFormParams := gourl.Values{}
	if r.botManagementWorkspace == nil {
		return localVarReturnValue, nil, reportError("botManagementWorkspace is required and must be specified")
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
	localVarPostBody = r.botManagementWorkspace
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
			var v BotManagementInvalidRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BotManagementNotAuthenticatedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BotManagementNotAllowedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BotManagementNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v BotManagementRateLimitError
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
