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
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextHTTPSignatureAuth takes HTTPSignatureAuth as authentication for the request.
	ContextHTTPSignatureAuth = contextKey("httpsignature")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "fastly-go/v1.0.0-beta.24",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "https://api.fastly.com",
				Description: "No description provided",
			},
			{
				URL: "https://rt.fastly.com",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
			"ACLAPIService.CreateACL": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLAPIService.DeleteACL": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLAPIService.GetACL": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLAPIService.ListACLs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLAPIService.UpdateACL": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLEntryAPIService.BulkUpdateACLEntries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLEntryAPIService.CreateACLEntry": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLEntryAPIService.DeleteACLEntry": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLEntryAPIService.GetACLEntry": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLEntryAPIService.ListACLEntries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ACLEntryAPIService.UpdateACLEntry": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ApexRedirectAPIService.CreateApexRedirect": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ApexRedirectAPIService.DeleteApexRedirect": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ApexRedirectAPIService.GetApexRedirect": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ApexRedirectAPIService.ListApexRedirects": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ApexRedirectAPIService.UpdateApexRedirect": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"AutomationTokensAPIService.CreateAutomationToken": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"AutomationTokensAPIService.GetAutomationTokenID": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"AutomationTokensAPIService.GetAutomationTokensIDServices": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"AutomationTokensAPIService.ListAutomationTokens": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"AutomationTokensAPIService.RevokeAutomationTokenID": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BackendAPIService.CreateBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BackendAPIService.DeleteBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BackendAPIService.GetBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BackendAPIService.ListBackends": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BackendAPIService.UpdateBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAPIService.GetInvoice": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAPIService.GetInvoiceByID": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAPIService.GetInvoiceMtd": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAddressAPIService.AddBillingAddr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAddressAPIService.DeleteBillingAddr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAddressAPIService.GetBillingAddr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"BillingAddressAPIService.UpdateBillingAddr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CacheSettingsAPIService.CreateCacheSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CacheSettingsAPIService.DeleteCacheSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CacheSettingsAPIService.GetCacheSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CacheSettingsAPIService.ListCacheSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CacheSettingsAPIService.UpdateCacheSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConditionAPIService.CreateCondition": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConditionAPIService.DeleteCondition": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConditionAPIService.GetCondition": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConditionAPIService.ListConditions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConditionAPIService.UpdateCondition": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.CreateConfigStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.DeleteConfigStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.GetConfigStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.GetConfigStoreInfo": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.ListConfigStoreServices": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.ListConfigStores": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreAPIService.UpdateConfigStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.BulkUpdateConfigStoreItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.CreateConfigStoreItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.DeleteConfigStoreItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.GetConfigStoreItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.ListConfigStoreItems": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.UpdateConfigStoreItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ConfigStoreItemAPIService.UpsertConfigStoreItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ContactAPIService.CreateContacts": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ContactAPIService.DeleteContact": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ContactAPIService.ListContacts": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ContentAPIService.ContentCheck": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CustomerAPIService.DeleteCustomer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CustomerAPIService.GetCustomer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CustomerAPIService.GetLoggedInCustomer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CustomerAPIService.ListUsers": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"CustomerAPIService.UpdateCustomer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryAPIService.CreateDictionary": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryAPIService.DeleteDictionary": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryAPIService.GetDictionary": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryAPIService.ListDictionaries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryAPIService.UpdateDictionary": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryInfoAPIService.GetDictionaryInfo": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.BulkUpdateDictionaryItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.CreateDictionaryItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.DeleteDictionaryItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.GetDictionaryItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.ListDictionaryItems": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.UpdateDictionaryItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DictionaryItemAPIService.UpsertDictionaryItem": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DiffAPIService.DiffServiceVersions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorAPIService.CreateDirector": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorAPIService.DeleteDirector": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorAPIService.GetDirector": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorAPIService.ListDirectors": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorAPIService.UpdateDirector": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorBackendAPIService.CreateDirectorBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorBackendAPIService.DeleteDirectorBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DirectorBackendAPIService.GetDirectorBackend": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.CheckDomain": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.CheckDomains": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.CreateDomain": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.DeleteDomain": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.GetDomain": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.ListDomains": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainAPIService.UpdateDomain": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainInspectorHistoricalAPIService.GetDomainInspectorHistorical": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainInspectorRealtimeAPIService.GetDomainInspectorLast120Seconds": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainInspectorRealtimeAPIService.GetDomainInspectorLastMaxEntries": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainInspectorRealtimeAPIService.GetDomainInspectorLastSecond": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"DomainOwnershipsAPIService.ListDomainOwnerships": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"EnabledProductsAPIService.DisableProduct": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"EnabledProductsAPIService.EnableProduct": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"EnabledProductsAPIService.GetEnabledProduct": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"EventsAPIService.GetEvent": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"EventsAPIService.ListEvents": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"GzipAPIService.CreateGzipConfig": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"GzipAPIService.DeleteGzipConfig": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"GzipAPIService.GetGzipConfigs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"GzipAPIService.ListGzipConfigs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"GzipAPIService.UpdateGzipConfig": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HeaderAPIService.CreateHeaderObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HeaderAPIService.DeleteHeaderObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HeaderAPIService.GetHeaderObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HeaderAPIService.ListHeaderObjects": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HeaderAPIService.UpdateHeaderObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HealthcheckAPIService.CreateHealthcheck": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HealthcheckAPIService.DeleteHealthcheck": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HealthcheckAPIService.GetHealthcheck": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HealthcheckAPIService.ListHealthchecks": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HealthcheckAPIService.UpdateHealthcheck": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetHistStats": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetHistStatsAggregated": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetHistStatsField": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetHistStatsService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetHistStatsServiceField": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetRegions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetUsage": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetUsageMonth": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HistoricalAPIService.GetUsageService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HTTP3APIService.CreateHTTP3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HTTP3APIService.DeleteHTTP3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"HTTP3APIService.GetHTTP3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamPermissionsAPIService.ListPermissions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.AddRolePermissions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.CreateARole": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.DeleteARole": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.GetARole": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.ListRolePermissions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.ListRoles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.RemoveRolePermissions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamRolesAPIService.UpdateARole": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.AddServiceGroupServices": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.CreateAServiceGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.DeleteAServiceGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.GetAServiceGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.ListServiceGroupServices": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.ListServiceGroups": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.RemoveServiceGroupServices": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamServiceGroupsAPIService.UpdateAServiceGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.AddUserGroupMembers": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.AddUserGroupRoles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.AddUserGroupServiceGroups": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.CreateAUserGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.DeleteAUserGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.GetAUserGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.ListUserGroupMembers": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.ListUserGroupRoles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.ListUserGroupServiceGroups": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.ListUserGroups": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.RemoveUserGroupMembers": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.RemoveUserGroupRoles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.RemoveUserGroupServiceGroups": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"IamUserGroupsAPIService.UpdateAUserGroup": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"InvitationsAPIService.CreateInvitation": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"InvitationsAPIService.DeleteInvitation": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"InvitationsAPIService.ListInvitations": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreAPIService.CreateStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreAPIService.DeleteStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreAPIService.GetStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreAPIService.GetStores": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreItemAPIService.DeleteKeyFromStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreItemAPIService.GetKeys": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreItemAPIService.GetValueForKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"KvStoreItemAPIService.SetValueForKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafConfigurationSetsAPIService.ListWafConfigSets": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafConfigurationSetsAPIService.ListWafsConfigSet": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafConfigurationSetsAPIService.UseWafConfigSet": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.CreateLegacyWafFirewallService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.DisableLegacyWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.EnableLegacyWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.GetLegacyWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.GetLegacyWafFirewallService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.ListLegacyWafFirewalls": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.ListLegacyWafFirewallsService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafFirewallAPIService.UpdateLegacyWafFirewallService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafOwaspAPIService.CreateOwaspSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafOwaspAPIService.GetOwaspSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafOwaspAPIService.UpdateOwaspSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleAPIService.GetLegacyWafFirewallRuleVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleAPIService.GetLegacyWafRule": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleAPIService.GetLegacyWafRuleVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleAPIService.ListLegacyWafRules": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleStatusAPIService.GetWafFirewallRuleStatus": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleStatusAPIService.ListWafFirewallRuleStatuses": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleStatusAPIService.UpdateWafFirewallRuleStatus": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRuleStatusAPIService.UpdateWafFirewallRuleStatusesTag": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRulesetAPIService.GetWafRuleset": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRulesetAPIService.GetWafRulesetVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafRulesetAPIService.UpdateWafRuleset": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafTagAPIService.ListLegacyWafTags": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafUpdateStatusAPIService.GetWafUpdateStatus": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LegacyWafUpdateStatusAPIService.ListWafUpdateStatuses": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingAzureblobAPIService.CreateLogAzure": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingAzureblobAPIService.DeleteLogAzure": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingAzureblobAPIService.GetLogAzure": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingAzureblobAPIService.ListLogAzure": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingAzureblobAPIService.UpdateLogAzure": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingBigqueryAPIService.CreateLogBigquery": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingBigqueryAPIService.DeleteLogBigquery": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingBigqueryAPIService.GetLogBigquery": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingBigqueryAPIService.ListLogBigquery": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingBigqueryAPIService.UpdateLogBigquery": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingCloudfilesAPIService.CreateLogCloudfiles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingCloudfilesAPIService.DeleteLogCloudfiles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingCloudfilesAPIService.GetLogCloudfiles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingCloudfilesAPIService.ListLogCloudfiles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingCloudfilesAPIService.UpdateLogCloudfiles": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDatadogAPIService.CreateLogDatadog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDatadogAPIService.DeleteLogDatadog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDatadogAPIService.GetLogDatadog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDatadogAPIService.ListLogDatadog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDatadogAPIService.UpdateLogDatadog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDigitaloceanAPIService.CreateLogDigocean": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDigitaloceanAPIService.DeleteLogDigocean": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDigitaloceanAPIService.GetLogDigocean": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDigitaloceanAPIService.ListLogDigocean": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingDigitaloceanAPIService.UpdateLogDigocean": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingElasticsearchAPIService.CreateLogElasticsearch": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingElasticsearchAPIService.DeleteLogElasticsearch": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingElasticsearchAPIService.GetLogElasticsearch": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingElasticsearchAPIService.ListLogElasticsearch": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingElasticsearchAPIService.UpdateLogElasticsearch": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingFtpAPIService.CreateLogFtp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingFtpAPIService.DeleteLogFtp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingFtpAPIService.GetLogFtp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingFtpAPIService.ListLogFtp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingFtpAPIService.UpdateLogFtp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingGcsAPIService.CreateLogGcs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingGcsAPIService.DeleteLogGcs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingGcsAPIService.GetLogGcs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingGcsAPIService.ListLogGcs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingGcsAPIService.UpdateLogGcs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHerokuAPIService.CreateLogHeroku": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHerokuAPIService.DeleteLogHeroku": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHerokuAPIService.GetLogHeroku": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHerokuAPIService.ListLogHeroku": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHerokuAPIService.UpdateLogHeroku": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHoneycombAPIService.CreateLogHoneycomb": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHoneycombAPIService.DeleteLogHoneycomb": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHoneycombAPIService.GetLogHoneycomb": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHoneycombAPIService.ListLogHoneycomb": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHoneycombAPIService.UpdateLogHoneycomb": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHTTPSAPIService.CreateLogHTTPS": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHTTPSAPIService.DeleteLogHTTPS": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHTTPSAPIService.GetLogHTTPS": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHTTPSAPIService.ListLogHTTPS": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingHTTPSAPIService.UpdateLogHTTPS": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKafkaAPIService.CreateLogKafka": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKafkaAPIService.DeleteLogKafka": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKafkaAPIService.GetLogKafka": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKafkaAPIService.ListLogKafka": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKafkaAPIService.UpdateLogKafka": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKinesisAPIService.CreateLogKinesis": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKinesisAPIService.DeleteLogKinesis": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKinesisAPIService.GetLogKinesis": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKinesisAPIService.ListLogKinesis": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingKinesisAPIService.UpdateLogKinesis": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogentriesAPIService.CreateLogLogentries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogentriesAPIService.DeleteLogLogentries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogentriesAPIService.GetLogLogentries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogentriesAPIService.ListLogLogentries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogentriesAPIService.UpdateLogLogentries": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogglyAPIService.CreateLogLoggly": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogglyAPIService.DeleteLogLoggly": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogglyAPIService.GetLogLoggly": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogglyAPIService.ListLogLoggly": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogglyAPIService.UpdateLogLoggly": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogshuttleAPIService.CreateLogLogshuttle": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogshuttleAPIService.DeleteLogLogshuttle": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogshuttleAPIService.GetLogLogshuttle": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogshuttleAPIService.ListLogLogshuttle": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingLogshuttleAPIService.UpdateLogLogshuttle": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicAPIService.CreateLogNewrelic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicAPIService.DeleteLogNewrelic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicAPIService.GetLogNewrelic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicAPIService.ListLogNewrelic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicAPIService.UpdateLogNewrelic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicotlpAPIService.CreateLogNewrelicotlp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicotlpAPIService.DeleteLogNewrelicotlp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicotlpAPIService.GetLogNewrelicotlp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicotlpAPIService.ListLogNewrelicotlp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingNewrelicotlpAPIService.UpdateLogNewrelicotlp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingOpenstackAPIService.CreateLogOpenstack": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingOpenstackAPIService.DeleteLogOpenstack": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingOpenstackAPIService.GetLogOpenstack": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingOpenstackAPIService.ListLogOpenstack": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingOpenstackAPIService.UpdateLogOpenstack": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPapertrailAPIService.CreateLogPapertrail": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPapertrailAPIService.DeleteLogPapertrail": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPapertrailAPIService.GetLogPapertrail": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPapertrailAPIService.ListLogPapertrail": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPapertrailAPIService.UpdateLogPapertrail": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPubsubAPIService.CreateLogGcpPubsub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPubsubAPIService.DeleteLogGcpPubsub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPubsubAPIService.GetLogGcpPubsub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPubsubAPIService.ListLogGcpPubsub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingPubsubAPIService.UpdateLogGcpPubsub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingS3APIService.CreateLogAwsS3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingS3APIService.DeleteLogAwsS3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingS3APIService.GetLogAwsS3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingS3APIService.ListLogAwsS3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingS3APIService.UpdateLogAwsS3": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingScalyrAPIService.CreateLogScalyr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingScalyrAPIService.DeleteLogScalyr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingScalyrAPIService.GetLogScalyr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingScalyrAPIService.ListLogScalyr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingScalyrAPIService.UpdateLogScalyr": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSftpAPIService.CreateLogSftp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSftpAPIService.DeleteLogSftp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSftpAPIService.GetLogSftp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSftpAPIService.ListLogSftp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSftpAPIService.UpdateLogSftp": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSplunkAPIService.CreateLogSplunk": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSplunkAPIService.DeleteLogSplunk": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSplunkAPIService.GetLogSplunk": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSplunkAPIService.ListLogSplunk": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSplunkAPIService.UpdateLogSplunk": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSumologicAPIService.CreateLogSumologic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSumologicAPIService.DeleteLogSumologic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSumologicAPIService.GetLogSumologic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSumologicAPIService.ListLogSumologic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSumologicAPIService.UpdateLogSumologic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSyslogAPIService.CreateLogSyslog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSyslogAPIService.DeleteLogSyslog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSyslogAPIService.GetLogSyslog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSyslogAPIService.ListLogSyslog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"LoggingSyslogAPIService.UpdateLogSyslog": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"MutualAuthenticationAPIService.CreateMutualTLSAuthentication": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"MutualAuthenticationAPIService.DeleteMutualTLS": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"MutualAuthenticationAPIService.GetMutualAuthentication": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"MutualAuthenticationAPIService.ListMutualAuthentications": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"MutualAuthenticationAPIService.PatchMutualAuthentication": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"OriginInspectorHistoricalAPIService.GetOriginInspectorHistorical": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"OriginInspectorRealtimeAPIService.GetOriginInspectorLast120Seconds": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"OriginInspectorRealtimeAPIService.GetOriginInspectorLastMaxEntries": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"OriginInspectorRealtimeAPIService.GetOriginInspectorLastSecond": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"PackageAPIService.GetPackage": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PackageAPIService.PutPackage": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PoolAPIService.CreateServerPool": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PoolAPIService.DeleteServerPool": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PoolAPIService.GetServerPool": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PoolAPIService.ListServerPools": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PoolAPIService.UpdateServerPool": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PopAPIService.ListPops": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PublicIPListAPIService.ListFastlyIps": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PublishAPIService.Publish": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PurgeAPIService.BulkPurgeTag": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PurgeAPIService.PurgeAll": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PurgeAPIService.PurgeSingleURL": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"PurgeAPIService.PurgeTag": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RateLimiterAPIService.CreateRateLimiter": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RateLimiterAPIService.DeleteRateLimiter": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RateLimiterAPIService.GetRateLimiter": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RateLimiterAPIService.ListRateLimiters": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RateLimiterAPIService.UpdateRateLimiter": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RealtimeAPIService.GetStatsLast120Seconds": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"RealtimeAPIService.GetStatsLast120SecondsLimitEntries": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"RealtimeAPIService.GetStatsLastSecond": {
				{
					URL: "https://rt.fastly.com",
					Description: "No description provided",
				},
			},
			"RequestSettingsAPIService.CreateRequestSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RequestSettingsAPIService.DeleteRequestSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RequestSettingsAPIService.GetRequestSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RequestSettingsAPIService.ListRequestSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"RequestSettingsAPIService.UpdateRequestSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResourceAPIService.CreateResource": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResourceAPIService.DeleteResource": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResourceAPIService.GetResource": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResourceAPIService.ListResources": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResourceAPIService.UpdateResource": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResponseObjectAPIService.CreateResponseObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResponseObjectAPIService.DeleteResponseObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResponseObjectAPIService.GetResponseObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResponseObjectAPIService.ListResponseObjects": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ResponseObjectAPIService.UpdateResponseObject": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreAPIService.ClientKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreAPIService.CreateSecretStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreAPIService.DeleteSecretStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreAPIService.GetSecretStore": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreAPIService.GetSecretStores": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreAPIService.SigningKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreItemAPIService.CreateSecret": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreItemAPIService.DeleteSecret": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreItemAPIService.GetSecret": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreItemAPIService.GetSecrets": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreItemAPIService.MustRecreateSecret": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SecretStoreItemAPIService.RecreateSecret": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServerAPIService.CreatePoolServer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServerAPIService.DeletePoolServer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServerAPIService.GetPoolServer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServerAPIService.ListPoolServers": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServerAPIService.UpdatePoolServer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.CreateService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.DeleteService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.GetService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.GetServiceDetail": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.ListServiceDomains": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.ListServices": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.SearchService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAPIService.UpdateService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.CreateServiceAuthorization": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.DeleteServiceAuthorization": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.DeleteServiceAuthorization2": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.ListServiceAuthorization": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.ShowServiceAuthorization": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.UpdateServiceAuthorization": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"ServiceAuthorizationsAPIService.UpdateServiceAuthorization2": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SettingsAPIService.GetServiceSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SettingsAPIService.UpdateServiceSettings": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.CreateSnippet": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.DeleteSnippet": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.GetSnippet": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.GetSnippetDynamic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.ListSnippets": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.UpdateSnippet": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SnippetAPIService.UpdateSnippetDynamic": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"StarAPIService.CreateServiceStar": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"StarAPIService.DeleteServiceStar": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"StarAPIService.GetServiceStar": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"StarAPIService.ListServiceStars": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"StatsAPIService.GetServiceStats": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"SudoAPIService.RequestSudoAccess": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSActivationsAPIService.CreateTLSActivation": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSActivationsAPIService.DeleteTLSActivation": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSActivationsAPIService.GetTLSActivation": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSActivationsAPIService.ListTLSActivations": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSActivationsAPIService.UpdateTLSActivation": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSBulkCertificatesAPIService.DeleteBulkTLSCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSBulkCertificatesAPIService.GetTLSBulkCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSBulkCertificatesAPIService.ListTLSBulkCerts": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSBulkCertificatesAPIService.UpdateBulkTLSCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSBulkCertificatesAPIService.UploadTLSBulkCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSCertificatesAPIService.CreateTLSCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSCertificatesAPIService.DeleteTLSCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSCertificatesAPIService.GetTLSCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSCertificatesAPIService.ListTLSCerts": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSCertificatesAPIService.UpdateTLSCert": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSConfigurationsAPIService.GetTLSConfig": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSConfigurationsAPIService.ListTLSConfigs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSConfigurationsAPIService.UpdateTLSConfig": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSDomainsAPIService.ListTLSDomains": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSPrivateKeysAPIService.CreateTLSKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSPrivateKeysAPIService.DeleteTLSKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSPrivateKeysAPIService.GetTLSKey": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSPrivateKeysAPIService.ListTLSKeys": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.CreateGlobalsignEmailChallenge": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.CreateTLSSub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.DeleteGlobalsignEmailChallenge": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.DeleteTLSSub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.GetTLSSub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.ListTLSSubs": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TLSSubscriptionsAPIService.PatchTLSSub": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.BulkRevokeTokens": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.CreateToken": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.GetToken": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.GetTokenCurrent": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.ListTokensCustomer": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.ListTokensUser": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.RevokeToken": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"TokensAPIService.RevokeTokenCurrent": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.CreateUser": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.DeleteUser": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.GetCurrentUser": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.GetUser": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.RequestPasswordReset": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.UpdateUser": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"UserAPIService.UpdateUserPassword": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.CreateCustomVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.DeleteCustomVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.GetCustomVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.GetCustomVclBoilerplate": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.GetCustomVclGenerated": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.GetCustomVclGeneratedHighlighted": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.GetCustomVclHighlighted": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.GetCustomVclRaw": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.LintVclDefault": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.LintVclForService": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.ListCustomVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.SetCustomVclMain": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclAPIService.UpdateCustomVcl": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VclDiffAPIService.VclDiffServiceVersions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.ActivateServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.CloneServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.CreateServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.DeactivateServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.GetServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.ListServiceVersions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.LockServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.UpdateServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"VersionAPIService.ValidateServiceVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.BulkDeleteWafActiveRules": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.BulkUpdateWafActiveRules": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.CreateWafActiveRule": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.CreateWafActiveRulesTag": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.DeleteWafActiveRule": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.GetWafActiveRule": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.ListWafActiveRules": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafActiveRulesAPIService.UpdateWafActiveRule": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafExclusionsAPIService.CreateWafRuleExclusion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafExclusionsAPIService.DeleteWafRuleExclusion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafExclusionsAPIService.GetWafRuleExclusion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafExclusionsAPIService.ListWafRuleExclusions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafExclusionsAPIService.UpdateWafRuleExclusion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallVersionsAPIService.CloneWafFirewallVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallVersionsAPIService.CreateWafFirewallVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallVersionsAPIService.DeployActivateWafFirewallVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallVersionsAPIService.GetWafFirewallVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallVersionsAPIService.ListWafFirewallVersions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallVersionsAPIService.UpdateWafFirewallVersion": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallsAPIService.CreateWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallsAPIService.DeleteWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallsAPIService.GetWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallsAPIService.ListWafFirewalls": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafFirewallsAPIService.UpdateWafFirewall": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafRuleRevisionsAPIService.GetWafRuleRevision": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafRuleRevisionsAPIService.ListWafRuleRevisions": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafRulesAPIService.GetWafRule": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafRulesAPIService.ListWafRules": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WafTagsAPIService.ListWafTags": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
			"WholePlatformDdosHistoricalAPIService.GetPlatformDdosHistorical": {
				{
					URL: "https://api.fastly.com",
					Description: "No description provided",
				},
			},
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.ReplaceAll(url, "{"+name+"}", value)
		} else {
			url = strings.ReplaceAll(url, "{"+name+"}", variable.DefaultValue)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		operationIndices, ok := osi.(map[string]int)
		if !ok {
			return 0, reportError("invalid type %T should be map[string]int", osi)
		}
		if index, ok := operationIndices[endpoint]; ok {
			return index, nil
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		operationVariables, ok := osv.(map[string]map[string]string)
		if !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		}
		if variables, ok := operationVariables[endpoint]; ok {
			return variables, nil
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// NewAPIKeyContext returns a context pre-configured with the provided Fastly API key.
func NewAPIKeyContext(apiKey string) context.Context {
	return context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"token": {
			Key: apiKey,
		},
	})
}

// NewAPIKeyContextFromEnv returns a context pre-configured with a Fastly API key from the provided environment variable.
func NewAPIKeyContextFromEnv(envKey string) context.Context {
	return context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"token": {
			Key: os.Getenv(envKey),
		},
	})
}
