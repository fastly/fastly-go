# fastly-go

A Go client library for interacting with most facets of the [Fastly API](https://developer.fastly.com/reference/api).

## Requirements

Go version 1.18

## Installation

Add the following to your project's `go.mod`:

```go.mod
require (
	github.com/fastly/fastly-go v1.0.0-beta.11
)
```

## Usage

> **NOTE**: The Fastly API requires an [API token](https://developer.fastly.com/reference/api/#authentication) for most operations.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fastly/fastly-go/fastly"
)

func main() {
  cfg := fastly.NewConfiguration()
  apiClient := fastly.NewAPIClient(cfg)
  ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")

  req := apiClient.ServiceAPI.CreateService(ctx)

  res, httpResp, err := req.Comment("comment_example").Name("name_example").CustomerID(os.Getenv("FASTLY_CUSTOMER_ID")).ResourceType("vcl").Execute()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.CreateService`: %+v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %+v\n", httpResp)
    return
  }

  fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.CreateService`:\n\n%+v\n\n", res)

  serviceID = *res.ID
  fmt.Printf("Service Name: %s\nService ID: %s\n\n", *res.Name, serviceID)
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.fastly.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ACLAPI* | [**CreateACL**](docs/AclAPI.md#createacl) | **POST** `/service/{service_id}/version/{version_id}/acl` | Create a new ACL
*ACLAPI* | [**DeleteACL**](docs/AclAPI.md#deleteacl) | **DELETE** `/service/{service_id}/version/{version_id}/acl/{acl_name}` | Delete an ACL
*ACLAPI* | [**GetACL**](docs/AclAPI.md#getacl) | **GET** `/service/{service_id}/version/{version_id}/acl/{acl_name}` | Describe an ACL
*ACLAPI* | [**ListACLs**](docs/AclAPI.md#listacls) | **GET** `/service/{service_id}/version/{version_id}/acl` | List ACLs
*ACLAPI* | [**UpdateACL**](docs/AclAPI.md#updateacl) | **PUT** `/service/{service_id}/version/{version_id}/acl/{acl_name}` | Update an ACL
*ACLEntryAPI* | [**BulkUpdateACLEntries**](docs/AclEntryAPI.md#bulkupdateaclentries) | **PATCH** `/service/{service_id}/acl/{acl_id}/entries` | Update multiple ACL entries
*ACLEntryAPI* | [**CreateACLEntry**](docs/AclEntryAPI.md#createaclentry) | **POST** `/service/{service_id}/acl/{acl_id}/entry` | Create an ACL entry
*ACLEntryAPI* | [**DeleteACLEntry**](docs/AclEntryAPI.md#deleteaclentry) | **DELETE** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Delete an ACL entry
*ACLEntryAPI* | [**GetACLEntry**](docs/AclEntryAPI.md#getaclentry) | **GET** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Describe an ACL entry
*ACLEntryAPI* | [**ListACLEntries**](docs/AclEntryAPI.md#listaclentries) | **GET** `/service/{service_id}/acl/{acl_id}/entries` | List ACL entries
*ACLEntryAPI* | [**UpdateACLEntry**](docs/AclEntryAPI.md#updateaclentry) | **PATCH** `/service/{service_id}/acl/{acl_id}/entry/{acl_entry_id}` | Update an ACL entry
*ApexRedirectAPI* | [**DeleteApexRedirect**](docs/ApexRedirectAPI.md#deleteapexredirect) | **DELETE** `/apex-redirects/{apex_redirect_id}` | Delete an apex redirect
*ApexRedirectAPI* | [**GetApexRedirect**](docs/ApexRedirectAPI.md#getapexredirect) | **GET** `/apex-redirects/{apex_redirect_id}` | Get an apex redirect
*ApexRedirectAPI* | [**ListApexRedirects**](docs/ApexRedirectAPI.md#listapexredirects) | **GET** `/service/{service_id}/version/{version_id}/apex-redirects` | List apex redirects
*ApexRedirectAPI* | [**UpdateApexRedirect**](docs/ApexRedirectAPI.md#updateapexredirect) | **PUT** `/apex-redirects/{apex_redirect_id}` | Update an apex redirect
*BackendAPI* | [**CreateBackend**](docs/BackendAPI.md#createbackend) | **POST** `/service/{service_id}/version/{version_id}/backend` | Create a backend
*BackendAPI* | [**DeleteBackend**](docs/BackendAPI.md#deletebackend) | **DELETE** `/service/{service_id}/version/{version_id}/backend/{backend_name}` | Delete a backend
*BackendAPI* | [**GetBackend**](docs/BackendAPI.md#getbackend) | **GET** `/service/{service_id}/version/{version_id}/backend/{backend_name}` | Describe a backend
*BackendAPI* | [**ListBackends**](docs/BackendAPI.md#listbackends) | **GET** `/service/{service_id}/version/{version_id}/backend` | List backends
*BackendAPI* | [**UpdateBackend**](docs/BackendAPI.md#updatebackend) | **PUT** `/service/{service_id}/version/{version_id}/backend/{backend_name}` | Update a backend
*BillingAPI* | [**GetInvoice**](docs/BillingAPI.md#getinvoice) | **GET** `/billing/v2/year/{year}/month/{month}` | Get an invoice
*BillingAPI* | [**GetInvoiceByID**](docs/BillingAPI.md#getinvoicebyid) | **GET** `/billing/v2/account_customers/{customer_id}/invoices/{invoice_id}` | Get an invoice
*BillingAPI* | [**GetInvoiceMtd**](docs/BillingAPI.md#getinvoicemtd) | **GET** `/billing/v2/account_customers/{customer_id}/mtd_invoice` | Get month-to-date billing estimate
*BillingAddressAPI* | [**AddBillingAddr**](docs/BillingAddressAPI.md#addbillingaddr) | **POST** `/customer/{customer_id}/billing_address` | Add a billing address to a customer
*BillingAddressAPI* | [**DeleteBillingAddr**](docs/BillingAddressAPI.md#deletebillingaddr) | **DELETE** `/customer/{customer_id}/billing_address` | Delete a billing address
*BillingAddressAPI* | [**GetBillingAddr**](docs/BillingAddressAPI.md#getbillingaddr) | **GET** `/customer/{customer_id}/billing_address` | Get a billing address
*BillingAddressAPI* | [**UpdateBillingAddr**](docs/BillingAddressAPI.md#updatebillingaddr) | **PATCH** `/customer/{customer_id}/billing_address` | Update a billing address
*CacheSettingsAPI* | [**CreateCacheSettings**](docs/CacheSettingsAPI.md#createcachesettings) | **POST** `/service/{service_id}/version/{version_id}/cache_settings` | Create a cache settings object
*CacheSettingsAPI* | [**DeleteCacheSettings**](docs/CacheSettingsAPI.md#deletecachesettings) | **DELETE** `/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}` | Delete a cache settings object
*CacheSettingsAPI* | [**GetCacheSettings**](docs/CacheSettingsAPI.md#getcachesettings) | **GET** `/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}` | Get a cache settings object
*CacheSettingsAPI* | [**ListCacheSettings**](docs/CacheSettingsAPI.md#listcachesettings) | **GET** `/service/{service_id}/version/{version_id}/cache_settings` | List cache settings objects
*CacheSettingsAPI* | [**UpdateCacheSettings**](docs/CacheSettingsAPI.md#updatecachesettings) | **PUT** `/service/{service_id}/version/{version_id}/cache_settings/{cache_settings_name}` | Update a cache settings object
*ConditionAPI* | [**CreateCondition**](docs/ConditionAPI.md#createcondition) | **POST** `/service/{service_id}/version/{version_id}/condition` | Create a condition
*ConditionAPI* | [**DeleteCondition**](docs/ConditionAPI.md#deletecondition) | **DELETE** `/service/{service_id}/version/{version_id}/condition/{condition_name}` | Delete a condition
*ConditionAPI* | [**GetCondition**](docs/ConditionAPI.md#getcondition) | **GET** `/service/{service_id}/version/{version_id}/condition/{condition_name}` | Describe a condition
*ConditionAPI* | [**ListConditions**](docs/ConditionAPI.md#listconditions) | **GET** `/service/{service_id}/version/{version_id}/condition` | List conditions
*ConditionAPI* | [**UpdateCondition**](docs/ConditionAPI.md#updatecondition) | **PUT** `/service/{service_id}/version/{version_id}/condition/{condition_name}` | Update a condition
*ConfigStoreAPI* | [**CreateConfigStore**](docs/ConfigStoreAPI.md#createconfigstore) | **POST** `/resources/stores/config` | Create a config store
*ConfigStoreAPI* | [**DeleteConfigStore**](docs/ConfigStoreAPI.md#deleteconfigstore) | **DELETE** `/resources/stores/config/{config_store_id}` | Delete a config store
*ConfigStoreAPI* | [**GetConfigStore**](docs/ConfigStoreAPI.md#getconfigstore) | **GET** `/resources/stores/config/{config_store_id}` | Describe a config store
*ConfigStoreAPI* | [**GetConfigStoreInfo**](docs/ConfigStoreAPI.md#getconfigstoreinfo) | **GET** `/resources/stores/config/{config_store_id}/info` | Get config store metadata
*ConfigStoreAPI* | [**ListConfigStoreServices**](docs/ConfigStoreAPI.md#listconfigstoreservices) | **GET** `/resources/stores/config/{config_store_id}/services` | List linked services
*ConfigStoreAPI* | [**ListConfigStores**](docs/ConfigStoreAPI.md#listconfigstores) | **GET** `/resources/stores/config` | List config stores
*ConfigStoreAPI* | [**UpdateConfigStore**](docs/ConfigStoreAPI.md#updateconfigstore) | **PUT** `/resources/stores/config/{config_store_id}` | Update a config store
*ConfigStoreItemAPI* | [**BulkUpdateConfigStoreItem**](docs/ConfigStoreItemAPI.md#bulkupdateconfigstoreitem) | **PATCH** `/resources/stores/config/{config_store_id}/items` | Update multiple entries in a config store
*ConfigStoreItemAPI* | [**CreateConfigStoreItem**](docs/ConfigStoreItemAPI.md#createconfigstoreitem) | **POST** `/resources/stores/config/{config_store_id}/item` | Create an entry in a config store
*ConfigStoreItemAPI* | [**DeleteConfigStoreItem**](docs/ConfigStoreItemAPI.md#deleteconfigstoreitem) | **DELETE** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Delete an item from a config store
*ConfigStoreItemAPI* | [**GetConfigStoreItem**](docs/ConfigStoreItemAPI.md#getconfigstoreitem) | **GET** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Get an item from a config store
*ConfigStoreItemAPI* | [**ListConfigStoreItems**](docs/ConfigStoreItemAPI.md#listconfigstoreitems) | **GET** `/resources/stores/config/{config_store_id}/items` | List items in a config store
*ConfigStoreItemAPI* | [**UpdateConfigStoreItem**](docs/ConfigStoreItemAPI.md#updateconfigstoreitem) | **PATCH** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Update an entry in a config store
*ConfigStoreItemAPI* | [**UpsertConfigStoreItem**](docs/ConfigStoreItemAPI.md#upsertconfigstoreitem) | **PUT** `/resources/stores/config/{config_store_id}/item/{config_store_item_key}` | Insert or update an entry in a config store
*ContactAPI* | [**DeleteContact**](docs/ContactAPI.md#deletecontact) | **DELETE** `/customer/{customer_id}/contact/{contact_id}` | Delete a contact
*ContactAPI* | [**ListContacts**](docs/ContactAPI.md#listcontacts) | **GET** `/customer/{customer_id}/contacts` | List contacts
*ContentAPI* | [**ContentCheck**](docs/ContentAPI.md#contentcheck) | **GET** `/content/edge_check` | Check status of content in each POP's cache
*CustomerAPI* | [**DeleteCustomer**](docs/CustomerAPI.md#deletecustomer) | **DELETE** `/customer/{customer_id}` | Delete a customer
*CustomerAPI* | [**GetCustomer**](docs/CustomerAPI.md#getcustomer) | **GET** `/customer/{customer_id}` | Get a customer
*CustomerAPI* | [**GetLoggedInCustomer**](docs/CustomerAPI.md#getloggedincustomer) | **GET** `/current_customer` | Get the logged in customer
*CustomerAPI* | [**ListUsers**](docs/CustomerAPI.md#listusers) | **GET** `/customer/{customer_id}/users` | List users
*CustomerAPI* | [**UpdateCustomer**](docs/CustomerAPI.md#updatecustomer) | **PUT** `/customer/{customer_id}` | Update a customer
*DictionaryAPI* | [**CreateDictionary**](docs/DictionaryAPI.md#createdictionary) | **POST** `/service/{service_id}/version/{version_id}/dictionary` | Create an edge dictionary
*DictionaryAPI* | [**DeleteDictionary**](docs/DictionaryAPI.md#deletedictionary) | **DELETE** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_name}` | Delete an edge dictionary
*DictionaryAPI* | [**GetDictionary**](docs/DictionaryAPI.md#getdictionary) | **GET** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_name}` | Get an edge dictionary
*DictionaryAPI* | [**ListDictionaries**](docs/DictionaryAPI.md#listdictionaries) | **GET** `/service/{service_id}/version/{version_id}/dictionary` | List edge dictionaries
*DictionaryAPI* | [**UpdateDictionary**](docs/DictionaryAPI.md#updatedictionary) | **PUT** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_name}` | Update an edge dictionary
*DictionaryInfoAPI* | [**GetDictionaryInfo**](docs/DictionaryInfoAPI.md#getdictionaryinfo) | **GET** `/service/{service_id}/version/{version_id}/dictionary/{dictionary_id}/info` | Get edge dictionary metadata
*DictionaryItemAPI* | [**BulkUpdateDictionaryItem**](docs/DictionaryItemAPI.md#bulkupdatedictionaryitem) | **PATCH** `/service/{service_id}/dictionary/{dictionary_id}/items` | Update multiple entries in an edge dictionary
*DictionaryItemAPI* | [**CreateDictionaryItem**](docs/DictionaryItemAPI.md#createdictionaryitem) | **POST** `/service/{service_id}/dictionary/{dictionary_id}/item` | Create an entry in an edge dictionary
*DictionaryItemAPI* | [**DeleteDictionaryItem**](docs/DictionaryItemAPI.md#deletedictionaryitem) | **DELETE** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Delete an item from an edge dictionary
*DictionaryItemAPI* | [**GetDictionaryItem**](docs/DictionaryItemAPI.md#getdictionaryitem) | **GET** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Get an item from an edge dictionary
*DictionaryItemAPI* | [**ListDictionaryItems**](docs/DictionaryItemAPI.md#listdictionaryitems) | **GET** `/service/{service_id}/dictionary/{dictionary_id}/items` | List items in an edge dictionary
*DictionaryItemAPI* | [**UpdateDictionaryItem**](docs/DictionaryItemAPI.md#updatedictionaryitem) | **PATCH** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Update an entry in an edge dictionary
*DictionaryItemAPI* | [**UpsertDictionaryItem**](docs/DictionaryItemAPI.md#upsertdictionaryitem) | **PUT** `/service/{service_id}/dictionary/{dictionary_id}/item/{dictionary_item_key}` | Insert or update an entry in an edge dictionary
*DiffAPI* | [**DiffServiceVersions**](docs/DiffAPI.md#diffserviceversions) | **GET** `/service/{service_id}/diff/from/{from_version_id}/to/{to_version_id}` | Diff two service versions
*DirectorAPI* | [**CreateDirector**](docs/DirectorAPI.md#createdirector) | **POST** `/service/{service_id}/version/{version_id}/director` | Create a director
*DirectorAPI* | [**DeleteDirector**](docs/DirectorAPI.md#deletedirector) | **DELETE** `/service/{service_id}/version/{version_id}/director/{director_name}` | Delete a director
*DirectorAPI* | [**GetDirector**](docs/DirectorAPI.md#getdirector) | **GET** `/service/{service_id}/version/{version_id}/director/{director_name}` | Get a director
*DirectorAPI* | [**ListDirectors**](docs/DirectorAPI.md#listdirectors) | **GET** `/service/{service_id}/version/{version_id}/director` | List directors
*DirectorBackendAPI* | [**CreateDirectorBackend**](docs/DirectorBackendAPI.md#createdirectorbackend) | **POST** `/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}` | Create a director-backend relationship
*DirectorBackendAPI* | [**DeleteDirectorBackend**](docs/DirectorBackendAPI.md#deletedirectorbackend) | **DELETE** `/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}` | Delete a director-backend relationship
*DirectorBackendAPI* | [**GetDirectorBackend**](docs/DirectorBackendAPI.md#getdirectorbackend) | **GET** `/service/{service_id}/version/{version_id}/director/{director_name}/backend/{backend_name}` | Get a director-backend relationship
*DomainAPI* | [**CheckDomain**](docs/DomainAPI.md#checkdomain) | **GET** `/service/{service_id}/version/{version_id}/domain/{domain_name}/check` | Validate DNS configuration for a single domain on a service
*DomainAPI* | [**CheckDomains**](docs/DomainAPI.md#checkdomains) | **GET** `/service/{service_id}/version/{version_id}/domain/check_all` | Validate DNS configuration for all domains on a service
*DomainAPI* | [**CreateDomain**](docs/DomainAPI.md#createdomain) | **POST** `/service/{service_id}/version/{version_id}/domain` | Add a domain name to a service
*DomainAPI* | [**DeleteDomain**](docs/DomainAPI.md#deletedomain) | **DELETE** `/service/{service_id}/version/{version_id}/domain/{domain_name}` | Remove a domain from a service
*DomainAPI* | [**GetDomain**](docs/DomainAPI.md#getdomain) | **GET** `/service/{service_id}/version/{version_id}/domain/{domain_name}` | Describe a domain
*DomainAPI* | [**ListDomains**](docs/DomainAPI.md#listdomains) | **GET** `/service/{service_id}/version/{version_id}/domain` | List domains
*DomainAPI* | [**UpdateDomain**](docs/DomainAPI.md#updatedomain) | **PUT** `/service/{service_id}/version/{version_id}/domain/{domain_name}` | Update a domain
*DomainOwnershipsAPI* | [**ListDomainOwnerships**](docs/DomainOwnershipsAPI.md#listdomainownerships) | **GET** `/domain-ownerships` | List domain-ownerships
*EnabledProductsAPI* | [**DisableProduct**](docs/EnabledProductsAPI.md#disableproduct) | **DELETE** `/enabled-products/{product_id}/services/{service_id}` | Disable a product
*EnabledProductsAPI* | [**EnableProduct**](docs/EnabledProductsAPI.md#enableproduct) | **PUT** `/enabled-products/{product_id}/services/{service_id}` | Enable a product
*EnabledProductsAPI* | [**GetEnabledProduct**](docs/EnabledProductsAPI.md#getenabledproduct) | **GET** `/enabled-products/{product_id}/services/{service_id}` | Get enabled product
*EventsAPI* | [**GetEvent**](docs/EventsAPI.md#getevent) | **GET** `/events/{event_id}` | Get an event
*EventsAPI* | [**ListEvents**](docs/EventsAPI.md#listevents) | **GET** `/events` | List events
*GzipAPI* | [**CreateGzipConfig**](docs/GzipAPI.md#creategzipconfig) | **POST** `/service/{service_id}/version/{version_id}/gzip` | Create a gzip configuration
*GzipAPI* | [**DeleteGzipConfig**](docs/GzipAPI.md#deletegzipconfig) | **DELETE** `/service/{service_id}/version/{version_id}/gzip/{gzip_name}` | Delete a gzip configuration
*GzipAPI* | [**GetGzipConfigs**](docs/GzipAPI.md#getgzipconfigs) | **GET** `/service/{service_id}/version/{version_id}/gzip/{gzip_name}` | Get a gzip configuration
*GzipAPI* | [**ListGzipConfigs**](docs/GzipAPI.md#listgzipconfigs) | **GET** `/service/{service_id}/version/{version_id}/gzip` | List gzip configurations
*GzipAPI* | [**UpdateGzipConfig**](docs/GzipAPI.md#updategzipconfig) | **PUT** `/service/{service_id}/version/{version_id}/gzip/{gzip_name}` | Update a gzip configuration
*HeaderAPI* | [**CreateHeaderObject**](docs/HeaderAPI.md#createheaderobject) | **POST** `/service/{service_id}/version/{version_id}/header` | Create a Header object
*HeaderAPI* | [**DeleteHeaderObject**](docs/HeaderAPI.md#deleteheaderobject) | **DELETE** `/service/{service_id}/version/{version_id}/header/{header_name}` | Delete a Header object
*HeaderAPI* | [**GetHeaderObject**](docs/HeaderAPI.md#getheaderobject) | **GET** `/service/{service_id}/version/{version_id}/header/{header_name}` | Get a Header object
*HeaderAPI* | [**ListHeaderObjects**](docs/HeaderAPI.md#listheaderobjects) | **GET** `/service/{service_id}/version/{version_id}/header` | List Header objects
*HeaderAPI* | [**UpdateHeaderObject**](docs/HeaderAPI.md#updateheaderobject) | **PUT** `/service/{service_id}/version/{version_id}/header/{header_name}` | Update a Header object
*HealthcheckAPI* | [**CreateHealthcheck**](docs/HealthcheckAPI.md#createhealthcheck) | **POST** `/service/{service_id}/version/{version_id}/healthcheck` | Create a health check
*HealthcheckAPI* | [**DeleteHealthcheck**](docs/HealthcheckAPI.md#deletehealthcheck) | **DELETE** `/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}` | Delete a health check
*HealthcheckAPI* | [**GetHealthcheck**](docs/HealthcheckAPI.md#gethealthcheck) | **GET** `/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}` | Get a health check
*HealthcheckAPI* | [**ListHealthchecks**](docs/HealthcheckAPI.md#listhealthchecks) | **GET** `/service/{service_id}/version/{version_id}/healthcheck` | List health checks
*HealthcheckAPI* | [**UpdateHealthcheck**](docs/HealthcheckAPI.md#updatehealthcheck) | **PUT** `/service/{service_id}/version/{version_id}/healthcheck/{healthcheck_name}` | Update a health check
*HistoricalAPI* | [**GetHistStats**](docs/HistoricalAPI.md#gethiststats) | **GET** `/stats` | Get historical stats
*HistoricalAPI* | [**GetHistStatsAggregated**](docs/HistoricalAPI.md#gethiststatsaggregated) | **GET** `/stats/aggregate` | Get aggregated historical stats
*HistoricalAPI* | [**GetHistStatsField**](docs/HistoricalAPI.md#gethiststatsfield) | **GET** `/stats/field/{field}` | Get historical stats for a single field
*HistoricalAPI* | [**GetHistStatsService**](docs/HistoricalAPI.md#gethiststatsservice) | **GET** `/stats/service/{service_id}` | Get historical stats for a single service
*HistoricalAPI* | [**GetHistStatsServiceField**](docs/HistoricalAPI.md#gethiststatsservicefield) | **GET** `/stats/service/{service_id}/field/{field}` | Get historical stats for a single service/field combination
*HistoricalAPI* | [**GetRegions**](docs/HistoricalAPI.md#getregions) | **GET** `/stats/regions` | Get region codes
*HistoricalAPI* | [**GetUsage**](docs/HistoricalAPI.md#getusage) | **GET** `/stats/usage` | Get usage statistics
*HistoricalAPI* | [**GetUsageMonth**](docs/HistoricalAPI.md#getusagemonth) | **GET** `/stats/usage_by_month` | Get month-to-date usage statistics
*HistoricalAPI* | [**GetUsageService**](docs/HistoricalAPI.md#getusageservice) | **GET** `/stats/usage_by_service` | Get usage statistics per service
*HTTP3API* | [**CreateHTTP3**](docs/Http3API.md#createhttp3) | **POST** `/service/{service_id}/version/{version_id}/http3` | Enable support for HTTP/3
*HTTP3API* | [**DeleteHTTP3**](docs/Http3API.md#deletehttp3) | **DELETE** `/service/{service_id}/version/{version_id}/http3` | Disable support for HTTP/3
*HTTP3API* | [**GetHTTP3**](docs/Http3API.md#gethttp3) | **GET** `/service/{service_id}/version/{version_id}/http3` | Get HTTP/3 status
*IamPermissionsAPI* | [**ListPermissions**](docs/IamPermissionsAPI.md#listpermissions) | **GET** `/permissions` | List permissions
*IamRolesAPI* | [**DeleteARole**](docs/IamRolesAPI.md#deletearole) | **DELETE** `/roles/{role_id}` | Delete a role
*IamRolesAPI* | [**GetARole**](docs/IamRolesAPI.md#getarole) | **GET** `/roles/{role_id}` | Get a role
*IamRolesAPI* | [**ListRolePermissions**](docs/IamRolesAPI.md#listrolepermissions) | **GET** `/roles/{role_id}/permissions` | List permissions in a role
*IamRolesAPI* | [**ListRoles**](docs/IamRolesAPI.md#listroles) | **GET** `/roles` | List roles
*IamServiceGroupsAPI* | [**DeleteAServiceGroup**](docs/IamServiceGroupsAPI.md#deleteaservicegroup) | **DELETE** `/service-groups/{service_group_id}` | Delete a service group
*IamServiceGroupsAPI* | [**GetAServiceGroup**](docs/IamServiceGroupsAPI.md#getaservicegroup) | **GET** `/service-groups/{service_group_id}` | Get a service group
*IamServiceGroupsAPI* | [**ListServiceGroupServices**](docs/IamServiceGroupsAPI.md#listservicegroupservices) | **GET** `/service-groups/{service_group_id}/services` | List services to a service group
*IamServiceGroupsAPI* | [**ListServiceGroups**](docs/IamServiceGroupsAPI.md#listservicegroups) | **GET** `/service-groups` | List service groups
*IamUserGroupsAPI* | [**DeleteAUserGroup**](docs/IamUserGroupsAPI.md#deleteausergroup) | **DELETE** `/user-groups/{user_group_id}` | Delete a user group
*IamUserGroupsAPI* | [**GetAUserGroup**](docs/IamUserGroupsAPI.md#getausergroup) | **GET** `/user-groups/{user_group_id}` | Get a user group
*IamUserGroupsAPI* | [**ListUserGroupMembers**](docs/IamUserGroupsAPI.md#listusergroupmembers) | **GET** `/user-groups/{user_group_id}/members` | List members of a user group
*IamUserGroupsAPI* | [**ListUserGroupRoles**](docs/IamUserGroupsAPI.md#listusergrouproles) | **GET** `/user-groups/{user_group_id}/roles` | List roles in a user group
*IamUserGroupsAPI* | [**ListUserGroupServiceGroups**](docs/IamUserGroupsAPI.md#listusergroupservicegroups) | **GET** `/user-groups/{user_group_id}/service-groups` | List service groups in a user group
*IamUserGroupsAPI* | [**ListUserGroups**](docs/IamUserGroupsAPI.md#listusergroups) | **GET** `/user-groups` | List user groups
*InvitationsAPI* | [**CreateInvitation**](docs/InvitationsAPI.md#createinvitation) | **POST** `/invitations` | Create an invitation
*InvitationsAPI* | [**DeleteInvitation**](docs/InvitationsAPI.md#deleteinvitation) | **DELETE** `/invitations/{invitation_id}` | Delete an invitation
*InvitationsAPI* | [**ListInvitations**](docs/InvitationsAPI.md#listinvitations) | **GET** `/invitations` | List invitations
*KvStoreAPI* | [**CreateStore**](docs/KvStoreAPI.md#createstore) | **POST** `/resources/stores/kv` | Create an kv store.
*KvStoreAPI* | [**DeleteStore**](docs/KvStoreAPI.md#deletestore) | **DELETE** `/resources/stores/kv/{store_id}` | Delete an kv store.
*KvStoreAPI* | [**GetStore**](docs/KvStoreAPI.md#getstore) | **GET** `/resources/stores/kv/{store_id}` | Describe an kv store.
*KvStoreAPI* | [**GetStores**](docs/KvStoreAPI.md#getstores) | **GET** `/resources/stores/kv` | List kv stores.
*KvStoreItemAPI* | [**DeleteKeyFromStore**](docs/KvStoreItemAPI.md#deletekeyfromstore) | **DELETE** `/resources/stores/kv/{store_id}/keys/{key_name}` | Delete kv store item.
*KvStoreItemAPI* | [**GetKeys**](docs/KvStoreItemAPI.md#getkeys) | **GET** `/resources/stores/kv/{store_id}/keys` | List kv store keys.
*KvStoreItemAPI* | [**GetValueForKey**](docs/KvStoreItemAPI.md#getvalueforkey) | **GET** `/resources/stores/kv/{store_id}/keys/{key_name}` | Get the value of an kv store item
*KvStoreItemAPI* | [**SetValueForKey**](docs/KvStoreItemAPI.md#setvalueforkey) | **PUT** `/resources/stores/kv/{store_id}/keys/{key_name}` | Insert an item into an kv store
*LoggingAzureblobAPI* | [**CreateLogAzure**](docs/LoggingAzureblobAPI.md#createlogazure) | **POST** `/service/{service_id}/version/{version_id}/logging/azureblob` | Create an Azure Blob Storage log endpoint
*LoggingAzureblobAPI* | [**DeleteLogAzure**](docs/LoggingAzureblobAPI.md#deletelogazure) | **DELETE** `/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}` | Delete the Azure Blob Storage log endpoint
*LoggingAzureblobAPI* | [**GetLogAzure**](docs/LoggingAzureblobAPI.md#getlogazure) | **GET** `/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}` | Get an Azure Blob Storage log endpoint
*LoggingAzureblobAPI* | [**ListLogAzure**](docs/LoggingAzureblobAPI.md#listlogazure) | **GET** `/service/{service_id}/version/{version_id}/logging/azureblob` | List Azure Blob Storage log endpoints
*LoggingAzureblobAPI* | [**UpdateLogAzure**](docs/LoggingAzureblobAPI.md#updatelogazure) | **PUT** `/service/{service_id}/version/{version_id}/logging/azureblob/{logging_azureblob_name}` | Update an Azure Blob Storage log endpoint
*LoggingBigqueryAPI* | [**CreateLogBigquery**](docs/LoggingBigqueryAPI.md#createlogbigquery) | **POST** `/service/{service_id}/version/{version_id}/logging/bigquery` | Create a BigQuery log endpoint
*LoggingBigqueryAPI* | [**DeleteLogBigquery**](docs/LoggingBigqueryAPI.md#deletelogbigquery) | **DELETE** `/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}` | Delete a BigQuery log endpoint
*LoggingBigqueryAPI* | [**GetLogBigquery**](docs/LoggingBigqueryAPI.md#getlogbigquery) | **GET** `/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}` | Get a BigQuery log endpoint
*LoggingBigqueryAPI* | [**ListLogBigquery**](docs/LoggingBigqueryAPI.md#listlogbigquery) | **GET** `/service/{service_id}/version/{version_id}/logging/bigquery` | List BigQuery log endpoints
*LoggingBigqueryAPI* | [**UpdateLogBigquery**](docs/LoggingBigqueryAPI.md#updatelogbigquery) | **PUT** `/service/{service_id}/version/{version_id}/logging/bigquery/{logging_bigquery_name}` | Update a BigQuery log endpoint
*LoggingCloudfilesAPI* | [**CreateLogCloudfiles**](docs/LoggingCloudfilesAPI.md#createlogcloudfiles) | **POST** `/service/{service_id}/version/{version_id}/logging/cloudfiles` | Create a Cloud Files log endpoint
*LoggingCloudfilesAPI* | [**DeleteLogCloudfiles**](docs/LoggingCloudfilesAPI.md#deletelogcloudfiles) | **DELETE** `/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}` | Delete the Cloud Files log endpoint
*LoggingCloudfilesAPI* | [**GetLogCloudfiles**](docs/LoggingCloudfilesAPI.md#getlogcloudfiles) | **GET** `/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}` | Get a Cloud Files log endpoint
*LoggingCloudfilesAPI* | [**ListLogCloudfiles**](docs/LoggingCloudfilesAPI.md#listlogcloudfiles) | **GET** `/service/{service_id}/version/{version_id}/logging/cloudfiles` | List Cloud Files log endpoints
*LoggingCloudfilesAPI* | [**UpdateLogCloudfiles**](docs/LoggingCloudfilesAPI.md#updatelogcloudfiles) | **PUT** `/service/{service_id}/version/{version_id}/logging/cloudfiles/{logging_cloudfiles_name}` | Update the Cloud Files log endpoint
*LoggingDatadogAPI* | [**CreateLogDatadog**](docs/LoggingDatadogAPI.md#createlogdatadog) | **POST** `/service/{service_id}/version/{version_id}/logging/datadog` | Create a Datadog log endpoint
*LoggingDatadogAPI* | [**DeleteLogDatadog**](docs/LoggingDatadogAPI.md#deletelogdatadog) | **DELETE** `/service/{service_id}/version/{version_id}/logging/datadog/{logging_datadog_name}` | Delete a Datadog log endpoint
*LoggingDatadogAPI* | [**GetLogDatadog**](docs/LoggingDatadogAPI.md#getlogdatadog) | **GET** `/service/{service_id}/version/{version_id}/logging/datadog/{logging_datadog_name}` | Get a Datadog log endpoint
*LoggingDatadogAPI* | [**ListLogDatadog**](docs/LoggingDatadogAPI.md#listlogdatadog) | **GET** `/service/{service_id}/version/{version_id}/logging/datadog` | List Datadog log endpoints
*LoggingDatadogAPI* | [**UpdateLogDatadog**](docs/LoggingDatadogAPI.md#updatelogdatadog) | **PUT** `/service/{service_id}/version/{version_id}/logging/datadog/{logging_datadog_name}` | Update a Datadog log endpoint
*LoggingDigitaloceanAPI* | [**CreateLogDigocean**](docs/LoggingDigitaloceanAPI.md#createlogdigocean) | **POST** `/service/{service_id}/version/{version_id}/logging/digitalocean` | Create a DigitalOcean Spaces log endpoint
*LoggingDigitaloceanAPI* | [**DeleteLogDigocean**](docs/LoggingDigitaloceanAPI.md#deletelogdigocean) | **DELETE** `/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}` | Delete a DigitalOcean Spaces log endpoint
*LoggingDigitaloceanAPI* | [**GetLogDigocean**](docs/LoggingDigitaloceanAPI.md#getlogdigocean) | **GET** `/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}` | Get a DigitalOcean Spaces log endpoint
*LoggingDigitaloceanAPI* | [**ListLogDigocean**](docs/LoggingDigitaloceanAPI.md#listlogdigocean) | **GET** `/service/{service_id}/version/{version_id}/logging/digitalocean` | List DigitalOcean Spaces log endpoints
*LoggingDigitaloceanAPI* | [**UpdateLogDigocean**](docs/LoggingDigitaloceanAPI.md#updatelogdigocean) | **PUT** `/service/{service_id}/version/{version_id}/logging/digitalocean/{logging_digitalocean_name}` | Update a DigitalOcean Spaces log endpoint
*LoggingElasticsearchAPI* | [**CreateLogElasticsearch**](docs/LoggingElasticsearchAPI.md#createlogelasticsearch) | **POST** `/service/{service_id}/version/{version_id}/logging/elasticsearch` | Create an Elasticsearch log endpoint
*LoggingElasticsearchAPI* | [**DeleteLogElasticsearch**](docs/LoggingElasticsearchAPI.md#deletelogelasticsearch) | **DELETE** `/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}` | Delete an Elasticsearch log endpoint
*LoggingElasticsearchAPI* | [**GetLogElasticsearch**](docs/LoggingElasticsearchAPI.md#getlogelasticsearch) | **GET** `/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}` | Get an Elasticsearch log endpoint
*LoggingElasticsearchAPI* | [**ListLogElasticsearch**](docs/LoggingElasticsearchAPI.md#listlogelasticsearch) | **GET** `/service/{service_id}/version/{version_id}/logging/elasticsearch` | List Elasticsearch log endpoints
*LoggingElasticsearchAPI* | [**UpdateLogElasticsearch**](docs/LoggingElasticsearchAPI.md#updatelogelasticsearch) | **PUT** `/service/{service_id}/version/{version_id}/logging/elasticsearch/{logging_elasticsearch_name}` | Update an Elasticsearch log endpoint
*LoggingFtpAPI* | [**CreateLogFtp**](docs/LoggingFtpAPI.md#createlogftp) | **POST** `/service/{service_id}/version/{version_id}/logging/ftp` | Create an FTP log endpoint
*LoggingFtpAPI* | [**DeleteLogFtp**](docs/LoggingFtpAPI.md#deletelogftp) | **DELETE** `/service/{service_id}/version/{version_id}/logging/ftp/{logging_ftp_name}` | Delete an FTP log endpoint
*LoggingFtpAPI* | [**GetLogFtp**](docs/LoggingFtpAPI.md#getlogftp) | **GET** `/service/{service_id}/version/{version_id}/logging/ftp/{logging_ftp_name}` | Get an FTP log endpoint
*LoggingFtpAPI* | [**ListLogFtp**](docs/LoggingFtpAPI.md#listlogftp) | **GET** `/service/{service_id}/version/{version_id}/logging/ftp` | List FTP log endpoints
*LoggingFtpAPI* | [**UpdateLogFtp**](docs/LoggingFtpAPI.md#updatelogftp) | **PUT** `/service/{service_id}/version/{version_id}/logging/ftp/{logging_ftp_name}` | Update an FTP log endpoint
*LoggingGcsAPI* | [**CreateLogGcs**](docs/LoggingGcsAPI.md#createloggcs) | **POST** `/service/{service_id}/version/{version_id}/logging/gcs` | Create a GCS log endpoint
*LoggingGcsAPI* | [**DeleteLogGcs**](docs/LoggingGcsAPI.md#deleteloggcs) | **DELETE** `/service/{service_id}/version/{version_id}/logging/gcs/{logging_gcs_name}` | Delete a GCS log endpoint
*LoggingGcsAPI* | [**GetLogGcs**](docs/LoggingGcsAPI.md#getloggcs) | **GET** `/service/{service_id}/version/{version_id}/logging/gcs/{logging_gcs_name}` | Get a GCS log endpoint
*LoggingGcsAPI* | [**ListLogGcs**](docs/LoggingGcsAPI.md#listloggcs) | **GET** `/service/{service_id}/version/{version_id}/logging/gcs` | List GCS log endpoints
*LoggingGcsAPI* | [**UpdateLogGcs**](docs/LoggingGcsAPI.md#updateloggcs) | **PUT** `/service/{service_id}/version/{version_id}/logging/gcs/{logging_gcs_name}` | Update a GCS log endpoint
*LoggingHerokuAPI* | [**CreateLogHeroku**](docs/LoggingHerokuAPI.md#createlogheroku) | **POST** `/service/{service_id}/version/{version_id}/logging/heroku` | Create a Heroku log endpoint
*LoggingHerokuAPI* | [**DeleteLogHeroku**](docs/LoggingHerokuAPI.md#deletelogheroku) | **DELETE** `/service/{service_id}/version/{version_id}/logging/heroku/{logging_heroku_name}` | Delete the Heroku log endpoint
*LoggingHerokuAPI* | [**GetLogHeroku**](docs/LoggingHerokuAPI.md#getlogheroku) | **GET** `/service/{service_id}/version/{version_id}/logging/heroku/{logging_heroku_name}` | Get a Heroku log endpoint
*LoggingHerokuAPI* | [**ListLogHeroku**](docs/LoggingHerokuAPI.md#listlogheroku) | **GET** `/service/{service_id}/version/{version_id}/logging/heroku` | List Heroku log endpoints
*LoggingHerokuAPI* | [**UpdateLogHeroku**](docs/LoggingHerokuAPI.md#updatelogheroku) | **PUT** `/service/{service_id}/version/{version_id}/logging/heroku/{logging_heroku_name}` | Update the Heroku log endpoint
*LoggingHoneycombAPI* | [**CreateLogHoneycomb**](docs/LoggingHoneycombAPI.md#createloghoneycomb) | **POST** `/service/{service_id}/version/{version_id}/logging/honeycomb` | Create a Honeycomb log endpoint
*LoggingHoneycombAPI* | [**DeleteLogHoneycomb**](docs/LoggingHoneycombAPI.md#deleteloghoneycomb) | **DELETE** `/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}` | Delete the Honeycomb log endpoint
*LoggingHoneycombAPI* | [**GetLogHoneycomb**](docs/LoggingHoneycombAPI.md#getloghoneycomb) | **GET** `/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}` | Get a Honeycomb log endpoint
*LoggingHoneycombAPI* | [**ListLogHoneycomb**](docs/LoggingHoneycombAPI.md#listloghoneycomb) | **GET** `/service/{service_id}/version/{version_id}/logging/honeycomb` | List Honeycomb log endpoints
*LoggingHoneycombAPI* | [**UpdateLogHoneycomb**](docs/LoggingHoneycombAPI.md#updateloghoneycomb) | **PUT** `/service/{service_id}/version/{version_id}/logging/honeycomb/{logging_honeycomb_name}` | Update a Honeycomb log endpoint
*LoggingHTTPSAPI* | [**CreateLogHTTPS**](docs/LoggingHTTPSAPI.md#createloghttps) | **POST** `/service/{service_id}/version/{version_id}/logging/https` | Create an HTTPS log endpoint
*LoggingHTTPSAPI* | [**DeleteLogHTTPS**](docs/LoggingHTTPSAPI.md#deleteloghttps) | **DELETE** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Delete an HTTPS log endpoint
*LoggingHTTPSAPI* | [**GetLogHTTPS**](docs/LoggingHTTPSAPI.md#getloghttps) | **GET** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Get an HTTPS log endpoint
*LoggingHTTPSAPI* | [**ListLogHTTPS**](docs/LoggingHTTPSAPI.md#listloghttps) | **GET** `/service/{service_id}/version/{version_id}/logging/https` | List HTTPS log endpoints
*LoggingHTTPSAPI* | [**UpdateLogHTTPS**](docs/LoggingHTTPSAPI.md#updateloghttps) | **PUT** `/service/{service_id}/version/{version_id}/logging/https/{logging_https_name}` | Update an HTTPS log endpoint
*LoggingKafkaAPI* | [**CreateLogKafka**](docs/LoggingKafkaAPI.md#createlogkafka) | **POST** `/service/{service_id}/version/{version_id}/logging/kafka` | Create a Kafka log endpoint
*LoggingKafkaAPI* | [**DeleteLogKafka**](docs/LoggingKafkaAPI.md#deletelogkafka) | **DELETE** `/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}` | Delete the Kafka log endpoint
*LoggingKafkaAPI* | [**GetLogKafka**](docs/LoggingKafkaAPI.md#getlogkafka) | **GET** `/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}` | Get a Kafka log endpoint
*LoggingKafkaAPI* | [**ListLogKafka**](docs/LoggingKafkaAPI.md#listlogkafka) | **GET** `/service/{service_id}/version/{version_id}/logging/kafka` | List Kafka log endpoints
*LoggingKinesisAPI* | [**CreateLogKinesis**](docs/LoggingKinesisAPI.md#createlogkinesis) | **POST** `/service/{service_id}/version/{version_id}/logging/kinesis` | Create  an Amazon Kinesis log endpoint
*LoggingKinesisAPI* | [**DeleteLogKinesis**](docs/LoggingKinesisAPI.md#deletelogkinesis) | **DELETE** `/service/{service_id}/version/{version_id}/logging/kinesis/{logging_kinesis_name}` | Delete the Amazon Kinesis log endpoint
*LoggingKinesisAPI* | [**GetLogKinesis**](docs/LoggingKinesisAPI.md#getlogkinesis) | **GET** `/service/{service_id}/version/{version_id}/logging/kinesis/{logging_kinesis_name}` | Get an Amazon Kinesis log endpoint
*LoggingKinesisAPI* | [**ListLogKinesis**](docs/LoggingKinesisAPI.md#listlogkinesis) | **GET** `/service/{service_id}/version/{version_id}/logging/kinesis` | List Amazon Kinesis log endpoints
*LoggingLogentriesAPI* | [**CreateLogLogentries**](docs/LoggingLogentriesAPI.md#createloglogentries) | **POST** `/service/{service_id}/version/{version_id}/logging/logentries` | Create a Logentries log endpoint
*LoggingLogentriesAPI* | [**DeleteLogLogentries**](docs/LoggingLogentriesAPI.md#deleteloglogentries) | **DELETE** `/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}` | Delete a Logentries log endpoint
*LoggingLogentriesAPI* | [**GetLogLogentries**](docs/LoggingLogentriesAPI.md#getloglogentries) | **GET** `/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}` | Get a Logentries log endpoint
*LoggingLogentriesAPI* | [**ListLogLogentries**](docs/LoggingLogentriesAPI.md#listloglogentries) | **GET** `/service/{service_id}/version/{version_id}/logging/logentries` | List Logentries log endpoints
*LoggingLogentriesAPI* | [**UpdateLogLogentries**](docs/LoggingLogentriesAPI.md#updateloglogentries) | **PUT** `/service/{service_id}/version/{version_id}/logging/logentries/{logging_logentries_name}` | Update a Logentries log endpoint
*LoggingLogglyAPI* | [**CreateLogLoggly**](docs/LoggingLogglyAPI.md#createlogloggly) | **POST** `/service/{service_id}/version/{version_id}/logging/loggly` | Create a Loggly log endpoint
*LoggingLogglyAPI* | [**DeleteLogLoggly**](docs/LoggingLogglyAPI.md#deletelogloggly) | **DELETE** `/service/{service_id}/version/{version_id}/logging/loggly/{logging_loggly_name}` | Delete a Loggly log endpoint
*LoggingLogglyAPI* | [**GetLogLoggly**](docs/LoggingLogglyAPI.md#getlogloggly) | **GET** `/service/{service_id}/version/{version_id}/logging/loggly/{logging_loggly_name}` | Get a Loggly log endpoint
*LoggingLogglyAPI* | [**ListLogLoggly**](docs/LoggingLogglyAPI.md#listlogloggly) | **GET** `/service/{service_id}/version/{version_id}/logging/loggly` | List Loggly log endpoints
*LoggingLogglyAPI* | [**UpdateLogLoggly**](docs/LoggingLogglyAPI.md#updatelogloggly) | **PUT** `/service/{service_id}/version/{version_id}/logging/loggly/{logging_loggly_name}` | Update a Loggly log endpoint
*LoggingLogshuttleAPI* | [**CreateLogLogshuttle**](docs/LoggingLogshuttleAPI.md#createloglogshuttle) | **POST** `/service/{service_id}/version/{version_id}/logging/logshuttle` | Create a Log Shuttle log endpoint
*LoggingLogshuttleAPI* | [**DeleteLogLogshuttle**](docs/LoggingLogshuttleAPI.md#deleteloglogshuttle) | **DELETE** `/service/{service_id}/version/{version_id}/logging/logshuttle/{logging_logshuttle_name}` | Delete a Log Shuttle log endpoint
*LoggingLogshuttleAPI* | [**GetLogLogshuttle**](docs/LoggingLogshuttleAPI.md#getloglogshuttle) | **GET** `/service/{service_id}/version/{version_id}/logging/logshuttle/{logging_logshuttle_name}` | Get a Log Shuttle log endpoint
*LoggingLogshuttleAPI* | [**ListLogLogshuttle**](docs/LoggingLogshuttleAPI.md#listloglogshuttle) | **GET** `/service/{service_id}/version/{version_id}/logging/logshuttle` | List Log Shuttle log endpoints
*LoggingLogshuttleAPI* | [**UpdateLogLogshuttle**](docs/LoggingLogshuttleAPI.md#updateloglogshuttle) | **PUT** `/service/{service_id}/version/{version_id}/logging/logshuttle/{logging_logshuttle_name}` | Update a Log Shuttle log endpoint
*LoggingNewrelicAPI* | [**CreateLogNewrelic**](docs/LoggingNewrelicAPI.md#createlognewrelic) | **POST** `/service/{service_id}/version/{version_id}/logging/newrelic` | Create a New Relic log endpoint
*LoggingNewrelicAPI* | [**DeleteLogNewrelic**](docs/LoggingNewrelicAPI.md#deletelognewrelic) | **DELETE** `/service/{service_id}/version/{version_id}/logging/newrelic/{logging_newrelic_name}` | Delete a New Relic log endpoint
*LoggingNewrelicAPI* | [**GetLogNewrelic**](docs/LoggingNewrelicAPI.md#getlognewrelic) | **GET** `/service/{service_id}/version/{version_id}/logging/newrelic/{logging_newrelic_name}` | Get a New Relic log endpoint
*LoggingNewrelicAPI* | [**ListLogNewrelic**](docs/LoggingNewrelicAPI.md#listlognewrelic) | **GET** `/service/{service_id}/version/{version_id}/logging/newrelic` | List New Relic log endpoints
*LoggingNewrelicAPI* | [**UpdateLogNewrelic**](docs/LoggingNewrelicAPI.md#updatelognewrelic) | **PUT** `/service/{service_id}/version/{version_id}/logging/newrelic/{logging_newrelic_name}` | Update a New Relic log endpoint
*LoggingOpenstackAPI* | [**CreateLogOpenstack**](docs/LoggingOpenstackAPI.md#createlogopenstack) | **POST** `/service/{service_id}/version/{version_id}/logging/openstack` | Create an OpenStack log endpoint
*LoggingOpenstackAPI* | [**DeleteLogOpenstack**](docs/LoggingOpenstackAPI.md#deletelogopenstack) | **DELETE** `/service/{service_id}/version/{version_id}/logging/openstack/{logging_openstack_name}` | Delete an OpenStack log endpoint
*LoggingOpenstackAPI* | [**GetLogOpenstack**](docs/LoggingOpenstackAPI.md#getlogopenstack) | **GET** `/service/{service_id}/version/{version_id}/logging/openstack/{logging_openstack_name}` | Get an OpenStack log endpoint
*LoggingOpenstackAPI* | [**ListLogOpenstack**](docs/LoggingOpenstackAPI.md#listlogopenstack) | **GET** `/service/{service_id}/version/{version_id}/logging/openstack` | List OpenStack log endpoints
*LoggingOpenstackAPI* | [**UpdateLogOpenstack**](docs/LoggingOpenstackAPI.md#updatelogopenstack) | **PUT** `/service/{service_id}/version/{version_id}/logging/openstack/{logging_openstack_name}` | Update an OpenStack log endpoint
*LoggingPapertrailAPI* | [**CreateLogPapertrail**](docs/LoggingPapertrailAPI.md#createlogpapertrail) | **POST** `/service/{service_id}/version/{version_id}/logging/papertrail` | Create a Papertrail log endpoint
*LoggingPapertrailAPI* | [**DeleteLogPapertrail**](docs/LoggingPapertrailAPI.md#deletelogpapertrail) | **DELETE** `/service/{service_id}/version/{version_id}/logging/papertrail/{logging_papertrail_name}` | Delete a Papertrail log endpoint
*LoggingPapertrailAPI* | [**GetLogPapertrail**](docs/LoggingPapertrailAPI.md#getlogpapertrail) | **GET** `/service/{service_id}/version/{version_id}/logging/papertrail/{logging_papertrail_name}` | Get a Papertrail log endpoint
*LoggingPapertrailAPI* | [**ListLogPapertrail**](docs/LoggingPapertrailAPI.md#listlogpapertrail) | **GET** `/service/{service_id}/version/{version_id}/logging/papertrail` | List Papertrail log endpoints
*LoggingPapertrailAPI* | [**UpdateLogPapertrail**](docs/LoggingPapertrailAPI.md#updatelogpapertrail) | **PUT** `/service/{service_id}/version/{version_id}/logging/papertrail/{logging_papertrail_name}` | Update a Papertrail log endpoint
*LoggingPubsubAPI* | [**CreateLogGcpPubsub**](docs/LoggingPubsubAPI.md#createloggcppubsub) | **POST** `/service/{service_id}/version/{version_id}/logging/pubsub` | Create a GCP Cloud Pub/Sub log endpoint
*LoggingPubsubAPI* | [**DeleteLogGcpPubsub**](docs/LoggingPubsubAPI.md#deleteloggcppubsub) | **DELETE** `/service/{service_id}/version/{version_id}/logging/pubsub/{logging_google_pubsub_name}` | Delete a GCP Cloud Pub/Sub log endpoint
*LoggingPubsubAPI* | [**GetLogGcpPubsub**](docs/LoggingPubsubAPI.md#getloggcppubsub) | **GET** `/service/{service_id}/version/{version_id}/logging/pubsub/{logging_google_pubsub_name}` | Get a GCP Cloud Pub/Sub log endpoint
*LoggingPubsubAPI* | [**ListLogGcpPubsub**](docs/LoggingPubsubAPI.md#listloggcppubsub) | **GET** `/service/{service_id}/version/{version_id}/logging/pubsub` | List GCP Cloud Pub/Sub log endpoints
*LoggingPubsubAPI* | [**UpdateLogGcpPubsub**](docs/LoggingPubsubAPI.md#updateloggcppubsub) | **PUT** `/service/{service_id}/version/{version_id}/logging/pubsub/{logging_google_pubsub_name}` | Update a GCP Cloud Pub/Sub log endpoint
*LoggingS3API* | [**CreateLogAwsS3**](docs/LoggingS3API.md#createlogawss3) | **POST** `/service/{service_id}/version/{version_id}/logging/s3` | Create an AWS S3 log endpoint
*LoggingS3API* | [**DeleteLogAwsS3**](docs/LoggingS3API.md#deletelogawss3) | **DELETE** `/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}` | Delete an AWS S3 log endpoint
*LoggingS3API* | [**GetLogAwsS3**](docs/LoggingS3API.md#getlogawss3) | **GET** `/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}` | Get an AWS S3 log endpoint
*LoggingS3API* | [**ListLogAwsS3**](docs/LoggingS3API.md#listlogawss3) | **GET** `/service/{service_id}/version/{version_id}/logging/s3` | List AWS S3 log endpoints
*LoggingS3API* | [**UpdateLogAwsS3**](docs/LoggingS3API.md#updatelogawss3) | **PUT** `/service/{service_id}/version/{version_id}/logging/s3/{logging_s3_name}` | Update an AWS S3 log endpoint
*LoggingScalyrAPI* | [**CreateLogScalyr**](docs/LoggingScalyrAPI.md#createlogscalyr) | **POST** `/service/{service_id}/version/{version_id}/logging/scalyr` | Create a Scalyr log endpoint
*LoggingScalyrAPI* | [**DeleteLogScalyr**](docs/LoggingScalyrAPI.md#deletelogscalyr) | **DELETE** `/service/{service_id}/version/{version_id}/logging/scalyr/{logging_scalyr_name}` | Delete the Scalyr log endpoint
*LoggingScalyrAPI* | [**GetLogScalyr**](docs/LoggingScalyrAPI.md#getlogscalyr) | **GET** `/service/{service_id}/version/{version_id}/logging/scalyr/{logging_scalyr_name}` | Get a Scalyr log endpoint
*LoggingScalyrAPI* | [**ListLogScalyr**](docs/LoggingScalyrAPI.md#listlogscalyr) | **GET** `/service/{service_id}/version/{version_id}/logging/scalyr` | List Scalyr log endpoints
*LoggingScalyrAPI* | [**UpdateLogScalyr**](docs/LoggingScalyrAPI.md#updatelogscalyr) | **PUT** `/service/{service_id}/version/{version_id}/logging/scalyr/{logging_scalyr_name}` | Update the Scalyr log endpoint
*LoggingSftpAPI* | [**CreateLogSftp**](docs/LoggingSftpAPI.md#createlogsftp) | **POST** `/service/{service_id}/version/{version_id}/logging/sftp` | Create an SFTP log endpoint
*LoggingSftpAPI* | [**DeleteLogSftp**](docs/LoggingSftpAPI.md#deletelogsftp) | **DELETE** `/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}` | Delete an SFTP log endpoint
*LoggingSftpAPI* | [**GetLogSftp**](docs/LoggingSftpAPI.md#getlogsftp) | **GET** `/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}` | Get an SFTP log endpoint
*LoggingSftpAPI* | [**ListLogSftp**](docs/LoggingSftpAPI.md#listlogsftp) | **GET** `/service/{service_id}/version/{version_id}/logging/sftp` | List SFTP log endpoints
*LoggingSftpAPI* | [**UpdateLogSftp**](docs/LoggingSftpAPI.md#updatelogsftp) | **PUT** `/service/{service_id}/version/{version_id}/logging/sftp/{logging_sftp_name}` | Update an SFTP log endpoint
*LoggingSplunkAPI* | [**CreateLogSplunk**](docs/LoggingSplunkAPI.md#createlogsplunk) | **POST** `/service/{service_id}/version/{version_id}/logging/splunk` | Create a Splunk log endpoint
*LoggingSplunkAPI* | [**DeleteLogSplunk**](docs/LoggingSplunkAPI.md#deletelogsplunk) | **DELETE** `/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}` | Delete a Splunk log endpoint
*LoggingSplunkAPI* | [**GetLogSplunk**](docs/LoggingSplunkAPI.md#getlogsplunk) | **GET** `/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}` | Get a Splunk log endpoint
*LoggingSplunkAPI* | [**ListLogSplunk**](docs/LoggingSplunkAPI.md#listlogsplunk) | **GET** `/service/{service_id}/version/{version_id}/logging/splunk` | List Splunk log endpoints
*LoggingSplunkAPI* | [**UpdateLogSplunk**](docs/LoggingSplunkAPI.md#updatelogsplunk) | **PUT** `/service/{service_id}/version/{version_id}/logging/splunk/{logging_splunk_name}` | Update a Splunk log endpoint
*LoggingSumologicAPI* | [**CreateLogSumologic**](docs/LoggingSumologicAPI.md#createlogsumologic) | **POST** `/service/{service_id}/version/{version_id}/logging/sumologic` | Create a Sumologic log endpoint
*LoggingSumologicAPI* | [**DeleteLogSumologic**](docs/LoggingSumologicAPI.md#deletelogsumologic) | **DELETE** `/service/{service_id}/version/{version_id}/logging/sumologic/{logging_sumologic_name}` | Delete a Sumologic log endpoint
*LoggingSumologicAPI* | [**GetLogSumologic**](docs/LoggingSumologicAPI.md#getlogsumologic) | **GET** `/service/{service_id}/version/{version_id}/logging/sumologic/{logging_sumologic_name}` | Get a Sumologic log endpoint
*LoggingSumologicAPI* | [**ListLogSumologic**](docs/LoggingSumologicAPI.md#listlogsumologic) | **GET** `/service/{service_id}/version/{version_id}/logging/sumologic` | List Sumologic log endpoints
*LoggingSumologicAPI* | [**UpdateLogSumologic**](docs/LoggingSumologicAPI.md#updatelogsumologic) | **PUT** `/service/{service_id}/version/{version_id}/logging/sumologic/{logging_sumologic_name}` | Update a Sumologic log endpoint
*LoggingSyslogAPI* | [**CreateLogSyslog**](docs/LoggingSyslogAPI.md#createlogsyslog) | **POST** `/service/{service_id}/version/{version_id}/logging/syslog` | Create a syslog log endpoint
*LoggingSyslogAPI* | [**DeleteLogSyslog**](docs/LoggingSyslogAPI.md#deletelogsyslog) | **DELETE** `/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}` | Delete a syslog log endpoint
*LoggingSyslogAPI* | [**GetLogSyslog**](docs/LoggingSyslogAPI.md#getlogsyslog) | **GET** `/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}` | Get a syslog log endpoint
*LoggingSyslogAPI* | [**ListLogSyslog**](docs/LoggingSyslogAPI.md#listlogsyslog) | **GET** `/service/{service_id}/version/{version_id}/logging/syslog` | List Syslog log endpoints
*LoggingSyslogAPI* | [**UpdateLogSyslog**](docs/LoggingSyslogAPI.md#updatelogsyslog) | **PUT** `/service/{service_id}/version/{version_id}/logging/syslog/{logging_syslog_name}` | Update a syslog log endpoint
*MutualAuthenticationAPI* | [**CreateMutualTLSAuthentication**](docs/MutualAuthenticationAPI.md#createmutualtlsauthentication) | **POST** `/tls/mutual_authentications` | Create a Mutual Authentication
*MutualAuthenticationAPI* | [**DeleteMutualTLS**](docs/MutualAuthenticationAPI.md#deletemutualtls) | **DELETE** `/tls/mutual_authentications/{mutual_authentication_id}` | Delete a Mutual TLS
*MutualAuthenticationAPI* | [**GetMutualAuthentication**](docs/MutualAuthenticationAPI.md#getmutualauthentication) | **GET** `/tls/mutual_authentications/{mutual_authentication_id}` | Get a Mutual Authentication
*MutualAuthenticationAPI* | [**ListMutualAuthentications**](docs/MutualAuthenticationAPI.md#listmutualauthentications) | **GET** `/tls/mutual_authentications` | List Mutual Authentications
*MutualAuthenticationAPI* | [**PatchMutualAuthentication**](docs/MutualAuthenticationAPI.md#patchmutualauthentication) | **PATCH** `/tls/mutual_authentications/{mutual_authentication_id}` | Update a Mutual Authentication
*PackageAPI* | [**GetPackage**](docs/PackageAPI.md#getpackage) | **GET** `/service/{service_id}/version/{version_id}/package` | Get details of the service's Compute@Edge package.
*PackageAPI* | [**PutPackage**](docs/PackageAPI.md#putpackage) | **PUT** `/service/{service_id}/version/{version_id}/package` | Upload a Compute@Edge package.
*PoolAPI* | [**CreateServerPool**](docs/PoolAPI.md#createserverpool) | **POST** `/service/{service_id}/version/{version_id}/pool` | Create a server pool
*PoolAPI* | [**DeleteServerPool**](docs/PoolAPI.md#deleteserverpool) | **DELETE** `/service/{service_id}/version/{version_id}/pool/{pool_name}` | Delete a server pool
*PoolAPI* | [**GetServerPool**](docs/PoolAPI.md#getserverpool) | **GET** `/service/{service_id}/version/{version_id}/pool/{pool_name}` | Get a server pool
*PoolAPI* | [**ListServerPools**](docs/PoolAPI.md#listserverpools) | **GET** `/service/{service_id}/version/{version_id}/pool` | List server pools
*PoolAPI* | [**UpdateServerPool**](docs/PoolAPI.md#updateserverpool) | **PUT** `/service/{service_id}/version/{version_id}/pool/{pool_name}` | Update a server pool
*PopAPI* | [**ListPops**](docs/PopAPI.md#listpops) | **GET** `/datacenters` | List Fastly POPs
*PublicIPListAPI* | [**ListFastlyIps**](docs/PublicIPListAPI.md#listfastlyips) | **GET** `/public-ip-list` | List Fastly's public IPs
*PublishAPI* | [**Publish**](docs/PublishAPI.md#publish) | **POST** `/service/{service_id}/publish/` | Send messages to Fanout subscribers
*PurgeAPI* | [**PurgeAll**](docs/PurgeAPI.md#purgeall) | **POST** `/service/{service_id}/purge_all` | Purge everything from a service
*PurgeAPI* | [**PurgeSingleURL**](docs/PurgeAPI.md#purgesingleurl) | **POST** `/purge/{cached_url}` | Purge a URL
*PurgeAPI* | [**PurgeTag**](docs/PurgeAPI.md#purgetag) | **POST** `/service/{service_id}/purge/{surrogate_key}` | Purge by surrogate key tag
*RateLimiterAPI* | [**DeleteRateLimiter**](docs/RateLimiterAPI.md#deleteratelimiter) | **DELETE** `/rate-limiters/{rate_limiter_id}` | Delete a rate limiter
*RateLimiterAPI* | [**GetRateLimiter**](docs/RateLimiterAPI.md#getratelimiter) | **GET** `/rate-limiters/{rate_limiter_id}` | Get a rate limiter
*RateLimiterAPI* | [**ListRateLimiters**](docs/RateLimiterAPI.md#listratelimiters) | **GET** `/service/{service_id}/version/{version_id}/rate-limiters` | List rate limiters
*RealtimeAPI* | [**GetStatsLast120Seconds**](docs/RealtimeAPI.md#getstatslast120seconds) | **GET** `/v1/channel/{service_id}/ts/h` | Get real-time data for the last 120 seconds
*RealtimeAPI* | [**GetStatsLast120SecondsLimitEntries**](docs/RealtimeAPI.md#getstatslast120secondslimitentries) | **GET** `/v1/channel/{service_id}/ts/h/limit/{max_entries}` | Get a limited number of real-time data entries
*RealtimeAPI* | [**GetStatsLastSecond**](docs/RealtimeAPI.md#getstatslastsecond) | **GET** `/v1/channel/{service_id}/ts/{timestamp_in_seconds}` | Get real-time data from specified time
*RequestSettingsAPI* | [**DeleteRequestSettings**](docs/RequestSettingsAPI.md#deleterequestsettings) | **DELETE** `/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}` | Delete a Request Settings object
*RequestSettingsAPI* | [**GetRequestSettings**](docs/RequestSettingsAPI.md#getrequestsettings) | **GET** `/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}` | Get a Request Settings object
*RequestSettingsAPI* | [**ListRequestSettings**](docs/RequestSettingsAPI.md#listrequestsettings) | **GET** `/service/{service_id}/version/{version_id}/request_settings` | List Request Settings objects
*RequestSettingsAPI* | [**UpdateRequestSettings**](docs/RequestSettingsAPI.md#updaterequestsettings) | **PUT** `/service/{service_id}/version/{version_id}/request_settings/{request_settings_name}` | Update a Request Settings object
*ResourceAPI* | [**CreateResource**](docs/ResourceAPI.md#createresource) | **POST** `/service/{service_id}/version/{version_id}/resource` | Create a resource link
*ResourceAPI* | [**DeleteResource**](docs/ResourceAPI.md#deleteresource) | **DELETE** `/service/{service_id}/version/{version_id}/resource/{id}` | Delete a resource link
*ResourceAPI* | [**GetResource**](docs/ResourceAPI.md#getresource) | **GET** `/service/{service_id}/version/{version_id}/resource/{id}` | Display a resource link
*ResourceAPI* | [**ListResources**](docs/ResourceAPI.md#listresources) | **GET** `/service/{service_id}/version/{version_id}/resource` | List resource links
*ResourceAPI* | [**UpdateResource**](docs/ResourceAPI.md#updateresource) | **PUT** `/service/{service_id}/version/{version_id}/resource/{id}` | Update a resource link
*ResponseObjectAPI* | [**DeleteResponseObject**](docs/ResponseObjectAPI.md#deleteresponseobject) | **DELETE** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Delete a Response Object
*ResponseObjectAPI* | [**GetResponseObject**](docs/ResponseObjectAPI.md#getresponseobject) | **GET** `/service/{service_id}/version/{version_id}/response_object/{response_object_name}` | Get a Response object
*ResponseObjectAPI* | [**ListResponseObjects**](docs/ResponseObjectAPI.md#listresponseobjects) | **GET** `/service/{service_id}/version/{version_id}/response_object` | List Response objects
*ServerAPI* | [**CreatePoolServer**](docs/ServerAPI.md#createpoolserver) | **POST** `/service/{service_id}/pool/{pool_id}/server` | Add a server to a pool
*ServerAPI* | [**DeletePoolServer**](docs/ServerAPI.md#deletepoolserver) | **DELETE** `/service/{service_id}/pool/{pool_id}/server/{server_id}` | Delete a server from a pool
*ServerAPI* | [**GetPoolServer**](docs/ServerAPI.md#getpoolserver) | **GET** `/service/{service_id}/pool/{pool_id}/server/{server_id}` | Get a pool server
*ServerAPI* | [**ListPoolServers**](docs/ServerAPI.md#listpoolservers) | **GET** `/service/{service_id}/pool/{pool_id}/servers` | List servers in a pool
*ServerAPI* | [**UpdatePoolServer**](docs/ServerAPI.md#updatepoolserver) | **PUT** `/service/{service_id}/pool/{pool_id}/server/{server_id}` | Update a server
*ServiceAPI* | [**CreateService**](docs/ServiceAPI.md#createservice) | **POST** `/service` | Create a service
*ServiceAPI* | [**DeleteService**](docs/ServiceAPI.md#deleteservice) | **DELETE** `/service/{service_id}` | Delete a service
*ServiceAPI* | [**GetService**](docs/ServiceAPI.md#getservice) | **GET** `/service/{service_id}` | Get a service
*ServiceAPI* | [**GetServiceDetail**](docs/ServiceAPI.md#getservicedetail) | **GET** `/service/{service_id}/details` | Get service details
*ServiceAPI* | [**ListServiceDomains**](docs/ServiceAPI.md#listservicedomains) | **GET** `/service/{service_id}/domain` | List the domains within a service
*ServiceAPI* | [**ListServices**](docs/ServiceAPI.md#listservices) | **GET** `/service` | List services
*ServiceAPI* | [**SearchService**](docs/ServiceAPI.md#searchservice) | **GET** `/service/search` | Search for a service by name
*ServiceAPI* | [**UpdateService**](docs/ServiceAPI.md#updateservice) | **PUT** `/service/{service_id}` | Update a service
*ServiceAuthorizationsAPI* | [**CreateServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#createserviceauthorization) | **POST** `/service-authorizations` | Create service authorization
*ServiceAuthorizationsAPI* | [**DeleteServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#deleteserviceauthorization) | **DELETE** `/service-authorizations/{service_authorization_id}` | Delete service authorization
*ServiceAuthorizationsAPI* | [**ListServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#listserviceauthorization) | **GET** `/service-authorizations` | List service authorizations
*ServiceAuthorizationsAPI* | [**ShowServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#showserviceauthorization) | **GET** `/service-authorizations/{service_authorization_id}` | Show service authorization
*ServiceAuthorizationsAPI* | [**UpdateServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#updateserviceauthorization) | **PATCH** `/service-authorizations/{service_authorization_id}` | Update service authorization
*SettingsAPI* | [**GetServiceSettings**](docs/SettingsAPI.md#getservicesettings) | **GET** `/service/{service_id}/version/{version_id}/settings` | Get service settings
*SettingsAPI* | [**UpdateServiceSettings**](docs/SettingsAPI.md#updateservicesettings) | **PUT** `/service/{service_id}/version/{version_id}/settings` | Update service settings
*SnippetAPI* | [**CreateSnippet**](docs/SnippetAPI.md#createsnippet) | **POST** `/service/{service_id}/version/{version_id}/snippet` | Create a snippet
*SnippetAPI* | [**DeleteSnippet**](docs/SnippetAPI.md#deletesnippet) | **DELETE** `/service/{service_id}/version/{version_id}/snippet/{snippet_name}` | Delete a snippet
*SnippetAPI* | [**GetSnippet**](docs/SnippetAPI.md#getsnippet) | **GET** `/service/{service_id}/version/{version_id}/snippet/{snippet_name}` | Get a versioned snippet
*SnippetAPI* | [**GetSnippetDynamic**](docs/SnippetAPI.md#getsnippetdynamic) | **GET** `/service/{service_id}/snippet/{snippet_id}` | Get a dynamic snippet
*SnippetAPI* | [**ListSnippets**](docs/SnippetAPI.md#listsnippets) | **GET** `/service/{service_id}/version/{version_id}/snippet` | List snippets
*SnippetAPI* | [**UpdateSnippetDynamic**](docs/SnippetAPI.md#updatesnippetdynamic) | **PUT** `/service/{service_id}/snippet/{snippet_id}` | Update a dynamic snippet
*StarAPI* | [**CreateServiceStar**](docs/StarAPI.md#createservicestar) | **POST** `/stars` | Create a star
*StarAPI* | [**DeleteServiceStar**](docs/StarAPI.md#deleteservicestar) | **DELETE** `/stars/{star_id}` | Delete a star
*StarAPI* | [**GetServiceStar**](docs/StarAPI.md#getservicestar) | **GET** `/stars/{star_id}` | Get a star
*StarAPI* | [**ListServiceStars**](docs/StarAPI.md#listservicestars) | **GET** `/stars` | List stars
*StatsAPI* | [**GetServiceStats**](docs/StatsAPI.md#getservicestats) | **GET** `/service/{service_id}/stats/summary` | Get stats for a service
*TLSActivationsAPI* | [**CreateTLSActivation**](docs/TlsActivationsAPI.md#createtlsactivation) | **POST** `/tls/activations` | Enable TLS for a domain using a custom certificate
*TLSActivationsAPI* | [**DeleteTLSActivation**](docs/TlsActivationsAPI.md#deletetlsactivation) | **DELETE** `/tls/activations/{tls_activation_id}` | Disable TLS on a domain
*TLSActivationsAPI* | [**GetTLSActivation**](docs/TlsActivationsAPI.md#gettlsactivation) | **GET** `/tls/activations/{tls_activation_id}` | Get a TLS activation
*TLSActivationsAPI* | [**ListTLSActivations**](docs/TlsActivationsAPI.md#listtlsactivations) | **GET** `/tls/activations` | List TLS activations
*TLSActivationsAPI* | [**UpdateTLSActivation**](docs/TlsActivationsAPI.md#updatetlsactivation) | **PATCH** `/tls/activations/{tls_activation_id}` | Update a certificate
*TLSBulkCertificatesAPI* | [**DeleteBulkTLSCert**](docs/TlsBulkCertificatesAPI.md#deletebulktlscert) | **DELETE** `/tls/bulk/certificates/{certificate_id}` | Delete a certificate
*TLSBulkCertificatesAPI* | [**GetTLSBulkCert**](docs/TlsBulkCertificatesAPI.md#gettlsbulkcert) | **GET** `/tls/bulk/certificates/{certificate_id}` | Get a certificate
*TLSBulkCertificatesAPI* | [**ListTLSBulkCerts**](docs/TlsBulkCertificatesAPI.md#listtlsbulkcerts) | **GET** `/tls/bulk/certificates` | List certificates
*TLSBulkCertificatesAPI* | [**UpdateBulkTLSCert**](docs/TlsBulkCertificatesAPI.md#updatebulktlscert) | **PATCH** `/tls/bulk/certificates/{certificate_id}` | Update a certificate
*TLSBulkCertificatesAPI* | [**UploadTLSBulkCert**](docs/TlsBulkCertificatesAPI.md#uploadtlsbulkcert) | **POST** `/tls/bulk/certificates` | Upload a certificate
*TLSCertificatesAPI* | [**CreateTLSCert**](docs/TlsCertificatesAPI.md#createtlscert) | **POST** `/tls/certificates` | Create a TLS certificate
*TLSCertificatesAPI* | [**DeleteTLSCert**](docs/TlsCertificatesAPI.md#deletetlscert) | **DELETE** `/tls/certificates/{tls_certificate_id}` | Delete a TLS certificate
*TLSCertificatesAPI* | [**GetTLSCert**](docs/TlsCertificatesAPI.md#gettlscert) | **GET** `/tls/certificates/{tls_certificate_id}` | Get a TLS certificate
*TLSCertificatesAPI* | [**ListTLSCerts**](docs/TlsCertificatesAPI.md#listtlscerts) | **GET** `/tls/certificates` | List TLS certificates
*TLSCertificatesAPI* | [**UpdateTLSCert**](docs/TlsCertificatesAPI.md#updatetlscert) | **PATCH** `/tls/certificates/{tls_certificate_id}` | Update a TLS certificate
*TLSConfigurationsAPI* | [**GetTLSConfig**](docs/TlsConfigurationsAPI.md#gettlsconfig) | **GET** `/tls/configurations/{tls_configuration_id}` | Get a TLS configuration
*TLSConfigurationsAPI* | [**ListTLSConfigs**](docs/TlsConfigurationsAPI.md#listtlsconfigs) | **GET** `/tls/configurations` | List TLS configurations
*TLSConfigurationsAPI* | [**UpdateTLSConfig**](docs/TlsConfigurationsAPI.md#updatetlsconfig) | **PATCH** `/tls/configurations/{tls_configuration_id}` | Update a TLS configuration
*TLSDomainsAPI* | [**ListTLSDomains**](docs/TlsDomainsAPI.md#listtlsdomains) | **GET** `/tls/domains` | List TLS domains
*TLSPrivateKeysAPI* | [**CreateTLSKey**](docs/TlsPrivateKeysAPI.md#createtlskey) | **POST** `/tls/private_keys` | Create a TLS private key
*TLSPrivateKeysAPI* | [**DeleteTLSKey**](docs/TlsPrivateKeysAPI.md#deletetlskey) | **DELETE** `/tls/private_keys/{tls_private_key_id}` | Delete a TLS private key
*TLSPrivateKeysAPI* | [**GetTLSKey**](docs/TlsPrivateKeysAPI.md#gettlskey) | **GET** `/tls/private_keys/{tls_private_key_id}` | Get a TLS private key
*TLSPrivateKeysAPI* | [**ListTLSKeys**](docs/TlsPrivateKeysAPI.md#listtlskeys) | **GET** `/tls/private_keys` | List TLS private keys
*TLSSubscriptionsAPI* | [**CreateGlobalsignEmailChallenge**](docs/TlsSubscriptionsAPI.md#createglobalsignemailchallenge) | **POST** `/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges` | Creates a GlobalSign email challenge.
*TLSSubscriptionsAPI* | [**CreateTLSSub**](docs/TlsSubscriptionsAPI.md#createtlssub) | **POST** `/tls/subscriptions` | Create a TLS subscription
*TLSSubscriptionsAPI* | [**DeleteGlobalsignEmailChallenge**](docs/TlsSubscriptionsAPI.md#deleteglobalsignemailchallenge) | **DELETE** `/tls/subscriptions/{tls_subscription_id}/authorizations/{tls_authorization_id}/globalsign_email_challenges/{globalsign_email_challenge_id}` | Delete a GlobalSign email challenge
*TLSSubscriptionsAPI* | [**DeleteTLSSub**](docs/TlsSubscriptionsAPI.md#deletetlssub) | **DELETE** `/tls/subscriptions/{tls_subscription_id}` | Delete a TLS subscription
*TLSSubscriptionsAPI* | [**GetTLSSub**](docs/TlsSubscriptionsAPI.md#gettlssub) | **GET** `/tls/subscriptions/{tls_subscription_id}` | Get a TLS subscription
*TLSSubscriptionsAPI* | [**ListTLSSubs**](docs/TlsSubscriptionsAPI.md#listtlssubs) | **GET** `/tls/subscriptions` | List TLS subscriptions
*TLSSubscriptionsAPI* | [**PatchTLSSub**](docs/TlsSubscriptionsAPI.md#patchtlssub) | **PATCH** `/tls/subscriptions/{tls_subscription_id}` | Update a TLS subscription
*TokensAPI* | [**GetToken**](docs/TokensAPI.md#gettoken) | **GET** `/tokens/{token_id}` | Get a token
*TokensAPI* | [**GetTokenCurrent**](docs/TokensAPI.md#gettokencurrent) | **GET** `/tokens/self` | Get the current token
*TokensAPI* | [**ListTokensCustomer**](docs/TokensAPI.md#listtokenscustomer) | **GET** `/customer/{customer_id}/tokens` | List tokens for a customer
*TokensAPI* | [**ListTokensUser**](docs/TokensAPI.md#listtokensuser) | **GET** `/tokens` | List tokens for the authenticated user
*TokensAPI* | [**RevokeToken**](docs/TokensAPI.md#revoketoken) | **DELETE** `/tokens/{token_id}` | Revoke a token
*TokensAPI* | [**RevokeTokenCurrent**](docs/TokensAPI.md#revoketokencurrent) | **DELETE** `/tokens/self` | Revoke the current token
*UserAPI* | [**CreateUser**](docs/UserAPI.md#createuser) | **POST** `/user` | Create a user
*UserAPI* | [**DeleteUser**](docs/UserAPI.md#deleteuser) | **DELETE** `/user/{user_id}` | Delete a user
*UserAPI* | [**GetCurrentUser**](docs/UserAPI.md#getcurrentuser) | **GET** `/current_user` | Get the current user
*UserAPI* | [**GetUser**](docs/UserAPI.md#getuser) | **GET** `/user/{user_id}` | Get a user
*UserAPI* | [**RequestPasswordReset**](docs/UserAPI.md#requestpasswordreset) | **POST** `/user/{user_login}/password/request_reset` | Request a password reset
*UserAPI* | [**UpdateUser**](docs/UserAPI.md#updateuser) | **PUT** `/user/{user_id}` | Update a user
*UserAPI* | [**UpdateUserPassword**](docs/UserAPI.md#updateuserpassword) | **POST** `/current_user/password` | Update the user's password
*VclAPI* | [**CreateCustomVcl**](docs/VclAPI.md#createcustomvcl) | **POST** `/service/{service_id}/version/{version_id}/vcl` | Create a custom VCL file
*VclAPI* | [**DeleteCustomVcl**](docs/VclAPI.md#deletecustomvcl) | **DELETE** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}` | Delete a custom VCL file
*VclAPI* | [**GetCustomVcl**](docs/VclAPI.md#getcustomvcl) | **GET** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}` | Get a custom VCL file
*VclAPI* | [**GetCustomVclBoilerplate**](docs/VclAPI.md#getcustomvclboilerplate) | **GET** `/service/{service_id}/version/{version_id}/boilerplate` | Get boilerplate VCL
*VclAPI* | [**GetCustomVclGenerated**](docs/VclAPI.md#getcustomvclgenerated) | **GET** `/service/{service_id}/version/{version_id}/generated_vcl` | Get the generated VCL for a service
*VclAPI* | [**GetCustomVclRaw**](docs/VclAPI.md#getcustomvclraw) | **GET** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}/download` | Download a custom VCL file
*VclAPI* | [**ListCustomVcl**](docs/VclAPI.md#listcustomvcl) | **GET** `/service/{service_id}/version/{version_id}/vcl` | List custom VCL files
*VclAPI* | [**SetCustomVclMain**](docs/VclAPI.md#setcustomvclmain) | **PUT** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}/main` | Set a custom VCL file as main
*VclAPI* | [**UpdateCustomVcl**](docs/VclAPI.md#updatecustomvcl) | **PUT** `/service/{service_id}/version/{version_id}/vcl/{vcl_name}` | Update a custom VCL file
*VclDiffAPI* | [**VclDiffServiceVersions**](docs/VclDiffAPI.md#vcldiffserviceversions) | **GET** `/service/{service_id}/vcl/diff/from/{from_version_id}/to/{to_version_id}` | Get a comparison of the VCL changes between two service versions
*VersionAPI* | [**ActivateServiceVersion**](docs/VersionAPI.md#activateserviceversion) | **PUT** `/service/{service_id}/version/{version_id}/activate` | Activate a service version
*VersionAPI* | [**CloneServiceVersion**](docs/VersionAPI.md#cloneserviceversion) | **PUT** `/service/{service_id}/version/{version_id}/clone` | Clone a service version
*VersionAPI* | [**CreateServiceVersion**](docs/VersionAPI.md#createserviceversion) | **POST** `/service/{service_id}/version` | Create a service version
*VersionAPI* | [**DeactivateServiceVersion**](docs/VersionAPI.md#deactivateserviceversion) | **PUT** `/service/{service_id}/version/{version_id}/deactivate` | Deactivate a service version
*VersionAPI* | [**GetServiceVersion**](docs/VersionAPI.md#getserviceversion) | **GET** `/service/{service_id}/version/{version_id}` | Get a version of a service
*VersionAPI* | [**ListServiceVersions**](docs/VersionAPI.md#listserviceversions) | **GET** `/service/{service_id}/version` | List versions of a service
*VersionAPI* | [**LockServiceVersion**](docs/VersionAPI.md#lockserviceversion) | **PUT** `/service/{service_id}/version/{version_id}/lock` | Lock a service version
*VersionAPI* | [**UpdateServiceVersion**](docs/VersionAPI.md#updateserviceversion) | **PUT** `/service/{service_id}/version/{version_id}` | Update a service version
*VersionAPI* | [**ValidateServiceVersion**](docs/VersionAPI.md#validateserviceversion) | **GET** `/service/{service_id}/version/{version_id}/validate` | Validate a service version
*WafActiveRulesAPI* | [**BulkUpdateWafActiveRules**](docs/WafActiveRulesAPI.md#bulkupdatewafactiverules) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/bulk` | Update multiple active rules
*WafActiveRulesAPI* | [**CreateWafActiveRule**](docs/WafActiveRulesAPI.md#createwafactiverule) | **POST** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules` | Add a rule to a WAF as an active rule
*WafActiveRulesAPI* | [**CreateWafActiveRulesTag**](docs/WafActiveRulesAPI.md#createwafactiverulestag) | **POST** `/waf/firewalls/{firewall_id}/versions/{version_id}/tags/{waf_tag_name}/active-rules` | Create active rules by tag
*WafActiveRulesAPI* | [**DeleteWafActiveRule**](docs/WafActiveRulesAPI.md#deletewafactiverule) | **DELETE** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}` | Delete an active rule
*WafActiveRulesAPI* | [**GetWafActiveRule**](docs/WafActiveRulesAPI.md#getwafactiverule) | **GET** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}` | Get an active WAF rule object
*WafActiveRulesAPI* | [**ListWafActiveRules**](docs/WafActiveRulesAPI.md#listwafactiverules) | **GET** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules` | List active rules on a WAF
*WafActiveRulesAPI* | [**UpdateWafActiveRule**](docs/WafActiveRulesAPI.md#updatewafactiverule) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}` | Update an active rule
*WafExclusionsAPI* | [**CreateWafRuleExclusion**](docs/WafExclusionsAPI.md#createwafruleexclusion) | **POST** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions` | Create a WAF rule exclusion
*WafExclusionsAPI* | [**DeleteWafRuleExclusion**](docs/WafExclusionsAPI.md#deletewafruleexclusion) | **DELETE** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}` | Delete a WAF rule exclusion
*WafExclusionsAPI* | [**GetWafRuleExclusion**](docs/WafExclusionsAPI.md#getwafruleexclusion) | **GET** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}` | Get a WAF rule exclusion
*WafExclusionsAPI* | [**ListWafRuleExclusions**](docs/WafExclusionsAPI.md#listwafruleexclusions) | **GET** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions` | List WAF rule exclusions
*WafExclusionsAPI* | [**UpdateWafRuleExclusion**](docs/WafExclusionsAPI.md#updatewafruleexclusion) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/exclusions/{exclusion_number}` | Update a WAF rule exclusion
*WafFirewallVersionsAPI* | [**CloneWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#clonewaffirewallversion) | **PUT** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/clone` | Clone a firewall version
*WafFirewallVersionsAPI* | [**CreateWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#createwaffirewallversion) | **POST** `/waf/firewalls/{firewall_id}/versions` | Create a firewall version
*WafFirewallVersionsAPI* | [**DeployActivateWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#deployactivatewaffirewallversion) | **PUT** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}/activate` | Deploy or activate a firewall version
*WafFirewallVersionsAPI* | [**GetWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#getwaffirewallversion) | **GET** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}` | Get a firewall version
*WafFirewallVersionsAPI* | [**ListWafFirewallVersions**](docs/WafFirewallVersionsAPI.md#listwaffirewallversions) | **GET** `/waf/firewalls/{firewall_id}/versions` | List firewall versions
*WafFirewallVersionsAPI* | [**UpdateWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#updatewaffirewallversion) | **PATCH** `/waf/firewalls/{firewall_id}/versions/{firewall_version_number}` | Update a firewall version
*WafFirewallsAPI* | [**CreateWafFirewall**](docs/WafFirewallsAPI.md#createwaffirewall) | **POST** `/waf/firewalls` | Create a firewall
*WafFirewallsAPI* | [**DeleteWafFirewall**](docs/WafFirewallsAPI.md#deletewaffirewall) | **DELETE** `/waf/firewalls/{firewall_id}` | Delete a firewall
*WafFirewallsAPI* | [**GetWafFirewall**](docs/WafFirewallsAPI.md#getwaffirewall) | **GET** `/waf/firewalls/{firewall_id}` | Get a firewall
*WafFirewallsAPI* | [**ListWafFirewalls**](docs/WafFirewallsAPI.md#listwaffirewalls) | **GET** `/waf/firewalls` | List firewalls
*WafFirewallsAPI* | [**UpdateWafFirewall**](docs/WafFirewallsAPI.md#updatewaffirewall) | **PATCH** `/waf/firewalls/{firewall_id}` | Update a firewall
*WafRuleRevisionsAPI* | [**GetWafRuleRevision**](docs/WafRuleRevisionsAPI.md#getwafrulerevision) | **GET** `/waf/rules/{waf_rule_id}/revisions/{waf_rule_revision_number}` | Get a revision of a rule
*WafRuleRevisionsAPI* | [**ListWafRuleRevisions**](docs/WafRuleRevisionsAPI.md#listwafrulerevisions) | **GET** `/waf/rules/{waf_rule_id}/revisions` | List revisions for a rule
*WafRulesAPI* | [**GetWafRule**](docs/WafRulesAPI.md#getwafrule) | **GET** `/waf/rules/{waf_rule_id}` | Get a rule
*WafRulesAPI* | [**ListWafRules**](docs/WafRulesAPI.md#listwafrules) | **GET** `/waf/rules` | List available WAF rules
*WafTagsAPI* | [**ListWafTags**](docs/WafTagsAPI.md#listwaftags) | **GET** `/waf/tags` | List tags


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Issues

The fastly-go API client currently does not support the following endpoints:

- [`/automation-tokens/{id}/services`](https://developer.fastly.com/reference/api/auth-tokens/automation) (GET)
- [`/automation-tokens/{id}`](https://developer.fastly.com/reference/api/auth-tokens/automation) (DELETE, GET)
- [`/automation-tokens`](https://developer.fastly.com/reference/api/auth-tokens/automation) (GET, POST)
- [`/customer/{customer_id}/contacts`](https://developer.fastly.com/reference/api/account/contact) (POST)
- [`/docs/section/{section}`](https://developer.fastly.com/reference/api/utils/docs) (GET)
- [`/docs/subject/{subject}`](https://developer.fastly.com/reference/api/utils/docs) (GET)
- [`/docs`](https://developer.fastly.com/reference/api/utils/docs) (GET)
- [`/metrics/domains/services/{service_id}`](https://developer.fastly.com/reference/api/metrics-stats/domain-inspector/historical) (GET)
- [`/metrics/origins/services/{service_id}`](https://developer.fastly.com/reference/api/metrics-stats/origin-inspector/historical) (GET)
- [`/rate-limiters/{rate_limiter_id}`](https://developer.fastly.com/reference/api/vcl-services/rate-limiter) (PUT)
- [`/resources/stores/secret/client-key`](https://developer.fastly.com/reference/api/services/resources/secret-store) (POST)
- [`/resources/stores/secret/signing-key`](https://developer.fastly.com/reference/api/services/resources/secret-store) (GET)
- [`/resources/stores/secret/{store_id}/secrets/{secret_name}`](https://developer.fastly.com/reference/api/services/resources/secret) (DELETE, GET)
- [`/resources/stores/secret/{store_id}/secrets`](https://developer.fastly.com/reference/api/services/resources/secret) (GET, PATCH, POST, PUT)
- [`/resources/stores/secret/{store_id}`](https://developer.fastly.com/reference/api/services/resources/secret-store) (DELETE, GET)
- [`/resources/stores/secret`](https://developer.fastly.com/reference/api/services/resources/secret-store) (GET, POST)
- [`/roles/{role_id}/permissions`](https://developer.fastly.com/reference/api/account/roles) (DELETE, POST)
- [`/roles/{role_id}`](https://developer.fastly.com/reference/api/account/roles) (PATCH)
- [`/roles`](https://developer.fastly.com/reference/api/account/roles) (POST)
- [`/service-authorizations`](https://developer.fastly.com/reference/api/account/service-authorization) (DELETE, PATCH)
- [`/service-groups/{service_group_id}/services`](https://developer.fastly.com/reference/api/account/service-groups) (DELETE, POST)
- [`/service-groups/{service_group_id}`](https://developer.fastly.com/reference/api/account/service-groups) (PATCH)
- [`/service-groups`](https://developer.fastly.com/reference/api/account/service-groups) (POST)
- [`/service/{service_id}/lint`](https://developer.fastly.com/reference/api/vcl-services/vcl) (POST)
- [`/service/{service_id}/purge`](https://developer.fastly.com/reference/api/purging) (POST)
- [`/service/{service_id}/version/{version_id}/apex-redirects`](https://developer.fastly.com/reference/api/vcl-services/apex-redirect) (POST)
- [`/service/{service_id}/version/{version_id}/director/{director_name}`](https://developer.fastly.com/reference/api/load-balancing/directors/director) (PUT)
- [`/service/{service_id}/version/{version_id}/generated_vcl/content`](https://developer.fastly.com/reference/api/vcl-services/vcl) (GET)
- [`/service/{service_id}/version/{version_id}/logging/kafka/{logging_kafka_name}`](https://developer.fastly.com/reference/api/logging/kafka) (PUT)
- [`/service/{service_id}/version/{version_id}/logging/kinesis/{logging_kinesis_name}`](https://developer.fastly.com/reference/api/logging/kinesis) (PUT)
- [`/service/{service_id}/version/{version_id}/rate-limiters`](https://developer.fastly.com/reference/api/vcl-services/rate-limiter) (POST)
- [`/service/{service_id}/version/{version_id}/request_settings`](https://developer.fastly.com/reference/api/vcl-services/request-settings) (POST)
- [`/service/{service_id}/version/{version_id}/response_object/{response_object_name}`](https://developer.fastly.com/reference/api/vcl-services/response-object) (PUT)
- [`/service/{service_id}/version/{version_id}/response_object`](https://developer.fastly.com/reference/api/vcl-services/response-object) (POST)
- [`/service/{service_id}/version/{version_id}/snippet/{snippet_name}`](https://developer.fastly.com/reference/api/vcl-services/snippet) (PUT)
- [`/service/{service_id}/version/{version_id}/vcl/{vcl_name}/content`](https://developer.fastly.com/reference/api/vcl-services/vcl) (GET)
- [`/service/{service_id}/version/{version_id}/wafs/{firewall_id}`](https://developer.fastly.com/reference/api/legacy-waf/firewall) (GET, PATCH)
- [`/service/{service_id}/version/{version_id}/wafs`](https://developer.fastly.com/reference/api/legacy-waf/firewall) (GET, POST)
- [`/service/{service_id}/wafs/{firewall_id}/owasp`](https://developer.fastly.com/reference/api/legacy-waf/owasp) (GET, PATCH, POST)
- [`/service/{service_id}/wafs/{firewall_id}/rule_statuses`](https://developer.fastly.com/reference/api/legacy-waf/rule-status) (GET, POST)
- [`/service/{service_id}/wafs/{firewall_id}/rules/{waf_rule_id}/rule_status`](https://developer.fastly.com/reference/api/legacy-waf/rule-status) (GET, PATCH)
- [`/service/{service_id}/wafs/{firewall_id}/ruleset/preview`](https://developer.fastly.com/reference/api/legacy-waf/ruleset) (GET)
- [`/service/{service_id}/wafs/{firewall_id}/ruleset`](https://developer.fastly.com/reference/api/legacy-waf/ruleset) (GET, PATCH)
- [`/service/{service_id}/wafs/{firewall_id}/update_statuses/{update_status_id}`](https://developer.fastly.com/reference/api/legacy-waf/update-status) (GET)
- [`/service/{service_id}/wafs/{firewall_id}/update_statuses`](https://developer.fastly.com/reference/api/legacy-waf/update-status) (GET)
- [`/sudo`](https://developer.fastly.com/reference/api/utils/sudo) (POST)
- [`/tls/activations/{tls_activation_id}`](https://developer.fastly.com/reference/api/tls/mutual-tls/activations) (GET, PATCH)
- [`/tls/activations`](https://developer.fastly.com/reference/api/tls/mutual-tls/activations) (GET)
- [`/tokens`](https://developer.fastly.com/reference/api/auth-tokens/user) (DELETE, POST)
- [`/user-groups/{user_group_id}/members`](https://developer.fastly.com/reference/api/account/user-groups) (DELETE, POST)
- [`/user-groups/{user_group_id}/roles`](https://developer.fastly.com/reference/api/account/user-groups) (DELETE, POST)
- [`/user-groups/{user_group_id}/service-groups`](https://developer.fastly.com/reference/api/account/user-groups) (DELETE, POST)
- [`/user-groups/{user_group_id}`](https://developer.fastly.com/reference/api/account/user-groups) (PATCH)
- [`/user-groups`](https://developer.fastly.com/reference/api/account/user-groups) (POST)
- [`/v1/channel/{service_id}/ts/h/limit/{max_entries}`](https://developer.fastly.com/reference/api/metrics-stats/origin-insights) (GET)
- [`/v1/channel/{service_id}/ts/h`](https://developer.fastly.com/reference/api/metrics-stats/origin-insights) (GET)
- [`/v1/channel/{service_id}/ts/{start_timestamp}`](https://developer.fastly.com/reference/api/metrics-stats/origin-insights) (GET)
- [`/v1/domains/{service_id}/ts/h/limit/{max_entries}`](https://developer.fastly.com/reference/api/metrics-stats/domain-inspector/real-time) (GET)
- [`/v1/domains/{service_id}/ts/h`](https://developer.fastly.com/reference/api/metrics-stats/domain-inspector/real-time) (GET)
- [`/v1/domains/{service_id}/ts/{start_timestamp}`](https://developer.fastly.com/reference/api/metrics-stats/domain-inspector/real-time) (GET)
- [`/v1/origins/{service_id}/ts/h/limit/{max_entries}`](https://developer.fastly.com/reference/api/metrics-stats/origin-inspector/real-time) (GET)
- [`/v1/origins/{service_id}/ts/h`](https://developer.fastly.com/reference/api/metrics-stats/origin-inspector/real-time) (GET)
- [`/v1/origins/{service_id}/ts/{start_timestamp}`](https://developer.fastly.com/reference/api/metrics-stats/origin-inspector/real-time) (GET)
- [`/vcl_lint`](https://developer.fastly.com/reference/api/vcl-services/vcl) (POST)
- [`/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules`](https://developer.fastly.com/reference/api/waf/rules/active) (DELETE)
- [`/wafs/configuration_sets/{configuration_set_id}/relationships/wafs`](https://developer.fastly.com/reference/api/legacy-waf/configuration-set) (GET, PATCH)
- [`/wafs/configuration_sets`](https://developer.fastly.com/reference/api/legacy-waf/configuration-set) (GET)
- [`/wafs/rules/{waf_rule_id}/vcl`](https://developer.fastly.com/reference/api/legacy-waf/rule) (GET)
- [`/wafs/rules/{waf_rule_id}`](https://developer.fastly.com/reference/api/legacy-waf/rule) (GET)
- [`/wafs/rules`](https://developer.fastly.com/reference/api/legacy-waf/rule) (GET)
- [`/wafs/tags`](https://developer.fastly.com/reference/api/legacy-waf/tag) (GET)
- [`/wafs/{firewall_id}/disable`](https://developer.fastly.com/reference/api/legacy-waf/firewall) (PATCH)
- [`/wafs/{firewall_id}/enable`](https://developer.fastly.com/reference/api/legacy-waf/firewall) (PATCH)
- [`/wafs/{firewall_id}/rules/{waf_rule_id}/vcl`](https://developer.fastly.com/reference/api/legacy-waf/rule) (GET)
- [`/wafs/{firewall_id}`](https://developer.fastly.com/reference/api/legacy-waf/firewall) (GET)
- [`/wafs`](https://developer.fastly.com/reference/api/legacy-waf/firewall) (GET)


If you encounter any non-security-related bug or unexpected behavior, please [file an issue][bug]
using the bug report template.

[bug]: https://github.com/fastly/fastly-rust/issues/new?labels=bug

### Security issues

Please see our [SECURITY.md](./SECURITY.md) for guidance on reporting security-related issues.

## License

[MIT](./LICENSE).
